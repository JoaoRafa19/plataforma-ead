package main

import (
	"fmt"
	"log"
	"net"

	grpcroutes "github.com/JoaoRafa19/plataforma-ead-service-core/internal/grpc-routes"
	"github.com/JoaoRafa19/plataforma-ead-service-core/pb"
	"google.golang.org/grpc"
)



func main() {

	fmt.Println("Start gRPC server!")

	lis, err := net.Listen("tcp", "0.0.0.0:9000")
	if err != nil {
		log.Fatal("error to listen server in port")
	}

	server := grpcroutes.NewServer()

	grpcServer := grpc.NewServer()

	pb.RegisterUserServiceServer(grpcServer, server )


	if err := grpcServer.Serve(lis); err != nil {
		log.Fatal("Erro ao iniciar server gRPC")
	}
}