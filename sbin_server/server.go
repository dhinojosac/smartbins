package main

import (
	"context"
	"fmt"
	"log"
	"net"

	sbin_pb "github.com/dhinojosac/smartbins/sbin_pb"
	"google.golang.org/grpc"
)

type server struct{}

func (*server) CreateBin(ctx context.Context, req *sbin_pb.BinRequest) (*sbin_pb.BinResponse, error) {
	fmt.Printf("CreateBin function was called with %v\n", req)

	log.Printf("Bin: %v", req.Bin)

	//TODO: Store in database

	result := true
	res := &sbin_pb.BinResponse{
		Status: result,
	}
	return res, nil
}

func (*server) GetFullBins(ctx context.Context, req *sbin_pb.GetFullBinsRequest) (*sbin_pb.GetFullBinsResponse, error) {
	fmt.Printf("GetNote streaming  server function was called with %v\n", req)

	log.Printf("GetFullBins in the city: %v", req.City)

	// Search on DB to get all full (>80%) bins in the city

	loc := &sbin_pb.Location{
		Region: "RM",
		City:   "Santiago",
		Place:  "Open World",
		Lat:    "-33.3875089",
		Long:   "-70.6168167",
	}
	bin := &sbin_pb.Bin{
		Id:       "12nd12a",
		Name:     "Contenedor 1",
		State:    "Full",
		Level:    82,
		Location: loc,
	}

	bins := []*sbin_pb.Bin{bin}
	//TODO: Get from db to return
	responses := &sbin_pb.GetFullBinsResponse{
		Bins: bins,
	}

	return responses, nil

}

func main() {
	fmt.Printf("Go Server is running\n")
	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("Failed to listen %v", err)

	}

	s := grpc.NewServer()

	sbin_pb.RegisterBinServiceServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve %v", err)

	}
}
