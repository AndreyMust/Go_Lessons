
package main

import (
	"context"
	"fmt"
	"log"

	"google.golang.org/grpc"
	pb "path/to/your/protobuf" // Замените на путь к вашему файлу protobuf
)

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewColorServiceClient(conn)

	response, err := client.GetColor(context.Background(), &pb.Empty{})
	if err != nil {
		log.Fatalf("Error calling GetColor: %v", err)
	}

	fmt.Printf("Response from server: %s\n", response.Color)
}