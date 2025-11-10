package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
	pb "path/to/your/protobuf" // Замените на путь к вашему файлу protobuf

)

type server struct{}

func (s *server) GetColor(ctx context.Context, request *pb.Empty) (*pb.ColorResponse, error) {
	return &pb.ColorResponse{Color: "color"}, nil
}

func main() {
	listen, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	srv := grpc.NewServer()
	pb.RegisterColorServiceServer(srv, &server{})

	fmt.Println("gRPC server is listening on port 50051...")
	if err := srv.Serve(listen); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}