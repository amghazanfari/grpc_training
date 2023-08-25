package main

import (
	"fmt"
	"log"
	"net"
	pb "server_streaming_server/proto_out"

	"google.golang.org/grpc"
)

const (
	port = ":50051"
)

type server struct {
	orderMap map[string]*pb.Order
}

var orderMap = []pb.Order{
	{
		Id:          "sample",
		Items:       []string{"one", "two"},
		Description: "something",
		Price:       float32(12.5),
		Destination: "germany",
	},
	{
		Id:          "sample",
		Items:       []string{"three"},
		Description: "something2",
		Price:       float32(13.5),
		Destination: "england",
	},
}

func (s *server) SearchOrders(searchQuery *pb.OrderName, stream pb.ProductInfo_SearchOrdersServer) error {
	for _, order := range orderMap {
		if order.Id == searchQuery.Name {
			err := stream.Send(&order)
			if err != nil {
				return fmt.Errorf("error sending message to stream : %v", err)
			}
		}
	}
	return nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatal("failed to listen %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterProductInfoServer(s, &server{})
	log.Printf("Starting gRPC listener on port " + port)
	if err := s.Serve(lis); err != nil {
		log.Fatal("failed to listen %v", err)
	}

}
