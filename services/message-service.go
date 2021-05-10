package services

import (
	"aapanavyapar-service-smssender/data-services/cash"
	"aapanavyapar-service-smssender/pb"
	"aapanavyapar-service-smssender/registers"
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"io"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"
)

type SmsSenderServiceServer struct {
	mutex sync.RWMutex
	Cash  *cash.RedisDataBase
	//ConnectionsRequest []*pb.ReadyToSendMessageRequest
	ConnectionsStream []pb.SmsSenderService_ReadyToSendMessageServer
	CurrentWorker     int
}

func NewSmsService() *SmsSenderServiceServer {

	server := &SmsSenderServiceServer{
		Cash: cash.NewRedisClient(),
		//ConnectionsRequest: nil,
		ConnectionsStream: nil,
		CurrentWorker:     -1,
	}

	err := server.Cash.InitSmsStream(context.TODO())
	if err != nil {
		panic(err)
	}

	checkOldMessages, err := strconv.ParseBool(os.Getenv("REDIS_STREAM_CHECK_BACKLOG"))
	if err != nil {
		panic(err)
	}

	go server.SyncPerformer(checkOldMessages)

	return server
}

func (smsService *SmsSenderServiceServer) ReadyToSendMessage(stream pb.SmsSenderService_ReadyToSendMessageServer) error {
	counter := 0
	for {
		request, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}

		_, ok := registers.APPS[request.GetApiKey()]
		if !ok {
			return status.Errorf(codes.Unauthenticated, "Invalid API Key")
		}

		if counter == 0 {
			if smsService.ConnectionsStream != nil {
				//smsService.ConnectionsRequest = append(smsService.ConnectionsRequest, request)
				smsService.ConnectionsStream = append(smsService.ConnectionsStream, stream)

			} else {
				//smsService.ConnectionsRequest = []*pb.ReadyToSendMessageRequest{request}
				smsService.ConnectionsStream = []pb.SmsSenderService_ReadyToSendMessageServer{stream}

			}
		}
		counter += 1
	}
	return nil
}

func (smsService *SmsSenderServiceServer) SendSms(ctx context.Context, request *pb.SendSmsRequest) (*pb.SendSmsResponse, error) {
	if os.Getenv("API_KEY") != request.GetApiKey() {
		return nil, status.Errorf(codes.Unauthenticated, "Invalid API Key")
	}

	smsService.mutex.Lock()
	mid := time.Now().UnixNano()
	smsService.mutex.Unlock()

	err := smsService.Cash.AddToSmsStream(ctx, request.GetMobileNo(), strconv.FormatInt(mid, 10), request.GetMessage())
	if err != nil {
		return nil, err
	}

	return &pb.SendSmsResponse{Status: true}, nil
}

func (smsService *SmsSenderServiceServer) AckToSms(ctx context.Context, request *pb.AckToSmsRequest) (*pb.AckToSmsResponse, error) {
	_, ok := registers.APPS[request.GetApiKey()]
	if !ok {
		return nil, status.Errorf(codes.Unauthenticated, "Invalid API Key")
	}

	err := smsService.Cash.AckToSmsStream(ctx, request.GetMessageId())
	if err != nil {
		return nil, err
	}

	err = smsService.Cash.DelFromSmsStream(ctx, request.GetMessageId())
	if err != nil {
		return nil, err
	}

	return &pb.AckToSmsResponse{Status: true}, nil
}

func (smsService *SmsSenderServiceServer) SyncPerformer(checkOldMessages bool) {
	lastId := "0"

	for {
		myKeyId := ">" //For Undelivered Ids So that Each Consumer Get Unique Id.
		if checkOldMessages {
			myKeyId = lastId

		}

		readGroup := smsService.Cash.ReadFromSmsStream(context.TODO(), 10, 200, myKeyId, os.Getenv("REDIS_STREAM_MSG_WORKER"))
		if readGroup.Err() != nil {
			if strings.Contains(readGroup.Err().Error(), "timeout") {
				fmt.Println("MSG : TimeOUT")
				continue

			}
			if readGroup.Err() == redis.Nil {
				fmt.Println("MSG : No Data Available")
				continue

			}
			panic(readGroup.Err())

		}

		data, err := readGroup.Result()
		if err != nil {
			panic(err)

		}

		if len(data[0].Messages) == 0 {
			checkOldMessages = false
			fmt.Println("MSG : Started Checking For New Messages ..!!")
			continue

		}

		var val redis.XMessage
		fmt.Println("\n\n\n NEW : ", lastId)
		for _, val = range data[0].Messages {
			err = smsService.SendSmsToSend(context.TODO(), &val)
			if err != nil {
				fmt.Println("MSG : Context Error Please Check For Sms Sending Device Connectivity, Network Error Or Any Other Dependency Problem")

				for len(smsService.ConnectionsStream) == 0 {
					fmt.Println("Waiting For Apps To Come Online")
					time.Sleep(time.Second)
				}

				checkOldMessages = true
				val.ID = "0"
				break

			}

		}
		lastId = val.ID
	}

}

func (smsService *SmsSenderServiceServer) SendSmsToSend(ctx context.Context, val *redis.XMessage) error {

	if len(smsService.ConnectionsStream) == 0 {
		return fmt.Errorf("no sms sending sevice online")

	} else {
		smsService.CurrentWorker = (smsService.CurrentWorker + 1) % len(smsService.ConnectionsStream)

		messageId := val.ID
		mobileNo := val.Values["mobileNo"].(string)
		message := val.Values["message"].(string)

		stream := smsService.ConnectionsStream[smsService.CurrentWorker]

		err := stream.Send(&pb.ReadyToSendMessageResponse{
			MobileNo:  mobileNo,
			MessageId: messageId,
			Message:   message,
		})
		if err != nil {
			smsService.ConnectionsStream = append(smsService.ConnectionsStream[:smsService.CurrentWorker], smsService.ConnectionsStream[smsService.CurrentWorker+1:]...)
			fmt.Println(err)
			return err
		}

		return nil
	}

}
