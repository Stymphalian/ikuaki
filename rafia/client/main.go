package main

import (
	"context"
	"fmt"
	"log"

	pb "github.com/Stymphalian/ikuaki/rafia"
	"github.com/kr/pretty"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

func getRpcOptions() []grpc.DialOption {
	var opts []grpc.DialOption
	creds := credentials.NewClientTLSFromCert(pb.KDemoCertPool, pb.KDemoAddr)
	opts = append(opts, grpc.WithTransportCredentials(creds))
	return opts
}

func main() {
	conn, err := grpc.Dial(pb.KDemoAddr, getRpcOptions()...)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	greet := pb.NewGreeterClient(conn)
	farewell := pb.NewFarewellerClient(conn)
	ctx := context.Background()

	_, err = greet.Greet(ctx, &pb.Empty{})
	if err != nil {
		log.Fatal(err)
	}
	_, err = farewell.Farewell(ctx, &pb.Empty{})
	if err != nil {
		log.Fatal(err)
	}
}

// A helper method for Pretty printing any object
func PrettyPrint(v interface{}) {
	fmt.Printf("%# v\n", pretty.Formatter(v))
}
