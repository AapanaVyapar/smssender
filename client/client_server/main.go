package main

import (
	"aapanavyapar-service-smssender/pb"
	"context"
	"fmt"
	_ "github.com/joho/godotenv/autoload"
	"google.golang.org/grpc"
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

	for i := 0; i < 5; i++ {
		resp1, err := client.SendSms(ctx, &pb.SendSmsRequest{
			ApiKey:   os.Getenv("API_KEY"),
			MobileNo: "9172879779",
			Message:  "Jai Shri Ram ..!!",
		})
		if err != nil {
			panic(err)
		}
		fmt.Println(resp1.GetStatus())
	}
}
