package main

import (
	"flag"
	"fmt"
	"log"
	"net"

	"github.com/Stymphalian/ikuaki/api/freeport"
	"github.com/Stymphalian/ikuaki/api/lobby"

	// googog "github.com/Stymphalian/ikuaki/api/lobby/client/protos"
	pb "github.com/Stymphalian/ikuaki/api/protos"
	// pb2 "github.com/Stymphalian/ikuaki/api/protosfat"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var (
	fPort = flag.Int("port", 0, "The port to run the server on.")
)

func main() {
	flag.Parse()
	var port int
	if *fPort != 0 {
		port = *fPort
	} else {
		port = freeport.GetPortOrDie()
	}
	addr := fmt.Sprintf(":%d", port)

	// fmt.Println(googog.Empty{})
	// fmt.Println(pb2.Fat{})

	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatal(err)
	}

	s := grpc.NewServer()
	pb.RegisterLobbyServer(s, lobby.NewLobby())
	reflection.Register(s)

	log.Println("Serving lobby: ", lis.Addr().String())
	if err := s.Serve(lis); err != nil {
		log.Fatal(err)
	}
}
