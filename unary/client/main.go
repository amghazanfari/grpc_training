package main

import (
	"context"
	"log"
	"time"

	pb "unary_client/proto_out"
	"unary_client/random"

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

	for {
		name := random.RandomProduct()
		description := random.RandomDescription()

		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()

		r, err := c.AddProduct(ctx, &pb.Product{Name: name, Description: description})

		if err != nil {
			log.Fatalf("Could not add product: %v", err)
		}
		log.Printf("Product ID: %s added successfully", r.Value)

		time.Sleep(time.Second)

	}
}
