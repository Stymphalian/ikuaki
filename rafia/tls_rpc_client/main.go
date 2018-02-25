package main

import (
	"context"
	"fmt"
	"log"

	rafia "github.com/Stymphalian/ikuaki/rafia"
	"github.com/Stymphalian/ikuaki/rafia/cert"
	"github.com/kr/pretty"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

func createGrpcOptions(serverCerts []string) []grpc.DialOption {
	pool, err := cert.CreateCertPoolFromPubKeys(serverCerts)
	if err != nil {
		log.Fatal(err)
	}

	creds := credentials.NewClientTLSFromCert(pool, "")
	return []grpc.DialOption{grpc.WithTransportCredentials(creds)}
}

func main() {
	addr := "localhost:8080"
	publicFile := "../data/server.public.pem"

	conn, err := grpc.Dial(addr, createGrpcOptions([]string{publicFile})...)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	greet := rafia.NewGreeterClient(conn)
	farewell := rafia.NewFarewellerClient(conn)
	ctx := context.Background()

	res, err := greet.Greet(ctx, &rafia.Empty{Text: "One"})
	if err != nil {
		log.Fatal(err)
	}
	res2, err := farewell.Farewell(ctx, &rafia.Empty{Text: "Two"})
	if err != nil {
		log.Fatal(err)
	}

	PrettyPrint(res)
	PrettyPrint(res2)
}

// A helper method for Pretty printing any object
func PrettyPrint(v interface{}) {
	fmt.Printf("%# v\n", pretty.Formatter(v))
}
