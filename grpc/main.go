package main

import (
	"log"
	"net"

	"google.golang.org/grpc"
	"github.com/tutorialedge/go-grpc-beginners-tutorial/chat"
)

func main()  {

	log.Println("gRPC Server Tutorial")

	lis, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Println(err)
	}

	s := chat.Server{}
	grpcServer := grpc.NewServer()

	chat.RegisterChatServiceServer(grpcServer, &s)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("faild to serve : %s", err)
	}

}
