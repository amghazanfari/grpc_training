package main

import (
	"context"
	"io"
	"log"
	"time"

	pb "server_streaming_client/proto_out"

	"google.golang.org/grpc"
)

const (
	address = "localhost:50051"
)

func main() {
	conn, err := grpc.Dial(address, grpc.WithInsecure())

	if err != nil {
		log.Fatalf("did not connect to %v", err)
	}
	defer conn.Close()

	c := pb.NewProductInfoClient(conn)
	name := "sample"

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	r, _ := c.SearchOrders(ctx, &pb.OrderName{Name: name})

	for {
		searchOrder, err := r.Recv()
		if err == io.EOF {
			break
		}
		log.Println("we get", searchOrder)
	}
}
