package api

import (
	"fmt"
	"log"
	"net"

	"github.com/Stymphalian/ikuaki/api/freeport"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const (
	WORLD_PORT = 40713
)

type RegisterServerFn func(s *grpc.Server)

func RunServerPortOrDie(port int, registerServer RegisterServerFn) {

	lis, err := net.Listen("tcp", fmt.Sprintf(":%v", port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	fmt.Printf("Listening on port %v\n", lis.Addr().String())
	s := grpc.NewServer()

	registerServer(s)

	// Register reflection service on gRPC server.
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func RunServerOrDie(registerServer RegisterServerFn) {
	port, err := freeport.GetPort()
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	RunServerPortOrDie(port, registerServer)
}