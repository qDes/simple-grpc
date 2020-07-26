package main

import (
	"context"
	"fmt"
	grpc "google.golang.org/grpc"
	"log"
	"net"
	"simple-grpc/pkg/proto/credit"
)

type server struct {
	credit.UnimplementedCreditServiceServer
}
func main() {
	log.Println("Server running ...")

	lis, err := net.Listen("tcp", ":5001")
	if err != nil {
		log.Fatalln(err)
	}

	srv := grpc.NewServer()
	credit.RegisterCreditServiceServer(srv, &server{})

	log.Fatalln(srv.Serve(lis))
}

func (s *server) Credit(ctx context.Context, request *credit.CreditRequest) (*credit.CreditResponse, error) {
	log.Println(fmt.Sprintf("Request: %g", request.GetAmount()))

	return &credit.CreditResponse{Confirmation: fmt.Sprintf("Credited %g", request.GetAmount())}, nil
}
