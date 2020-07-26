package main

import (
	"context"
	grpc "google.golang.org/grpc"
	"log"
	"simple-grpc/pkg/proto/credit"
	"time"
)

func main() {
	log.Println("Client running ...")

	conn, err := grpc.Dial(":5001", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalln(err)
	}
	defer conn.Close()

	client := credit.NewCreditServiceClient(conn)

	request := &credit.CreditRequest{Amount: 190.01}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	response, err := client.Credit(ctx, request)
	if err != nil {
		log.Fatalln(err)
	}

	log.Println("Response:", response.GetConfirmation())
}
