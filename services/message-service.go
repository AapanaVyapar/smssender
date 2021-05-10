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
	"os"
	"strconv"
	"strings"
	"sync"
	"time"
)

type SmsSenderServiceServer struct {
	mutex         sync.RWMutex
	Cash          *cash.RedisDataBase
	Connections   map[*pb.ReadyToSendMessageRequest]pb.SmsSenderService_ReadyToSendMessageServer
	CurrentWorker int
}

func NewSmsService() *SmsSenderServiceServer {

	server := &SmsSenderServiceServer{
		Cash:          cash.NewRedisClient(),
		Connections:   make(map[*pb.ReadyToSendMessageRequest]pb.SmsSenderService_ReadyToSendMessageServer, 10),
		CurrentWorker: 0,
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

func (smsService *SmsSenderServiceServer) ReadyToSendMessage(request *pb.ReadyToSendMessageRequest, stream pb.SmsSenderService_ReadyToSendMessageServer) error {
	_, ok := registers.APPS[request.GetApiKey()]
	if !ok {
		return status.Errorf(codes.Unauthenticated, "Invalid API Key")
	}
	smsService.Connections[request] = stream
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

		fmt.Println(smsService.CurrentWorker)

		readGroup := smsService.Cash.ReadFromSmsStream(context.TODO(), 10, 200, myKeyId, registers.APPS[registers.WORKERS_ARRAY[smsService.CurrentWorker]])
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
			fmt.Println("FAV : Started Checking For New Messages ..!!")
			continue

		}

		var val redis.XMessage
		fmt.Println("\n\n\n NEW : ", lastId)
		for _, val = range data[0].Messages {
			err = smsService.SendSmsToSend(context.TODO(), &val)
			if err != nil {
				fmt.Println("MSG : Context Error Please Check For Data Base Connectivity, Network Error Or Any Other Dependency Problem")
				checkOldMessages = true
				val.ID = "0"
				break

			}

		}
		lastId = val.ID
		if smsService.CurrentWorker >= len(registers.WORKERS_ARRAY)-1 {
			smsService.CurrentWorker = 0

		} else {
			smsService.CurrentWorker += 1

		}
	}

}

func (smsService *SmsSenderServiceServer) SendSmsToSend(ctx context.Context, val *redis.XMessage) error {
	messageId := val.ID
	mobileNo := val.Values["mobileNo"].(string)
	message := val.Values["message"].(string)

	fmt.Println(messageId)
	fmt.Println(mobileNo)
	fmt.Println(message)

	return nil
}
