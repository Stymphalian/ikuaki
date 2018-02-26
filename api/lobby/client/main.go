package main

import (
	"context"
	"log"

	pb "github.com/Stymphalian/ikuaki/api/protos"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:42657", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}

	client := pb.NewLobbyClient(conn)
	resp, err := client.CreateWorld(
		context.Background(),
		&pb.CreateWorldReq{Name: "World1"})
	if err != nil {
		log.Fatal(err)
	}

	log.Println("CreateWorldResp = ", resp.Addr)
}
