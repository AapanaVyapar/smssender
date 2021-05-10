package main

import (
	"aapanavyapar-service-smssender/pb"
	"context"
	"fmt"
	_ "github.com/joho/godotenv/autoload"
	"google.golang.org/grpc"
	"io"
	"log"
	"os"
	"time"
)

func main() {
	log.Printf("Stating server on port  :  %d", os.Getenv("Port"))

	fmt.Println("Environmental Variables Loaded .. !!")

	serverAddress := fmt.Sprintf("0.0.0.0:%s", os.Getenv("Port"))
	log.Printf("dialing to server  : %s", serverAddress)

	conn, err := grpc.Dial(serverAddress, grpc.WithInsecure())
	if err != nil {
		log.Fatal("Cannot dial server ", err)
	}

	client := pb.NewSmsSenderServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), 50*time.Hour)
	defer cancel()

	stream, err := client.ReadyToSendMessage(ctx)
	if err != nil {
		panic(err)
	}

	err = stream.Send(&pb.ReadyToSendMessageRequest{
		ApiKey: "fdfdsb&*h3uhfdskjwrhufds998Aihwihvbjfjhiur2732wefiuhsd7e98fdsa",
	})
	if err != nil {
		panic(err)
	}

	for {
		data, err := stream.Recv()
		fmt.Println("data : ", data)
		fmt.Println("error : ", err)
		if err == io.EOF {
			break
		}
		if err != nil {
			panic(err)
		}

		fmt.Println("Sending Message To : ", data.GetMobileNo())
		resp, err := client.AckToSms(ctx, &pb.AckToSmsRequest{
			ApiKey:    "fdfdsb&*h3uhfdskjwrhufds998Aihwihvbjfjhiur2732wefiuhsd7e98fdsa",
			MessageId: data.GetMessageId(),
		})
		if err != nil {
			panic(err)
		}
		fmt.Println(resp.GetStatus())

	}

}
