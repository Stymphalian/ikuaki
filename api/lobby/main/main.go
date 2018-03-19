package main

import (
	"flag"
	"fmt"
	"log"
	"net"

	"github.com/Stymphalian/ikuaki/api/lobby"
	"github.com/Stymphalian/ikuaki/freeport"

	pb "github.com/Stymphalian/ikuaki/protos"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var (
	fPort = flag.Int("port", 0, "The port to run the server on.")
)

func main() {
	// Parse all the flags
	flag.Parse()
	var port int
	if *fPort != 0 {
		port = *fPort
	} else {
		port = freeport.GetPortOrDie()
	}

	// Create a listener
	addr := fmt.Sprintf(":%d", port)
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
