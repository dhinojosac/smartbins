package main

import (
	"context"
	"fmt"
	"log"

	sbinpb "github.com/dhinojosac/smartbins/sbin_pb"
	"google.golang.org/grpc"
)

func main() {
	fmt.Printf("Go client is running\n")

	cc, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}

	defer cc.Close()

	c := sbinpb.NewBinServiceClient(cc)

	//todoUnary(c)

	todoServerStreaming(c)

}

func todoUnary(c sbinpb.BinServiceClient) {
	fmt.Println("Starting unary RPC Todo")

	req := &sbinpb.BinRequest{
		Bin: &sbinpb.Bin{
			Name: "Primer Bin",
			Id:   "221asd1i3",
		},
	}

	res, err := c.CreateBin(context.Background(), req)
	if err != nil {
		log.Fatalf("Error, calling Todo unary RPC: %v\n", err)
	}

	log.Printf("Response Todo: %v", res.Status)
}

func todoServerStreaming(c sbinpb.BinServiceClient) {
	fmt.Println("Starting streaming RPC Todo client")

	req := &sbinpb.GetFullBinsRequest{
		City: "Santiago",
	}

	res, err := c.GetFullBins(context.Background(), req)
	if err != nil {
		log.Fatalf("Error, calling Todo Streming RPC")
	}

	log.Printf("Response Todo: %v", res.Bins)

}
