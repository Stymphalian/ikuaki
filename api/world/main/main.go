package main

import (
	"flag"

	"github.com/Stymphalian/ikuaki/api"
	pb "github.com/Stymphalian/ikuaki/protos"
	"github.com/Stymphalian/ikuaki/api/world"
	"google.golang.org/grpc"
)

var (
	fPort = flag.Int("port", 0, "The port to run this server on")
)

func main() {
	flag.Parse()
	var port int
	if *fPort == 0 {
		port = api.WORLD_PORT
	} else {
		port = *fPort
	}

	api.RunServerPortOrDie(port, func(s *grpc.Server) {
		pb.RegisterWorldServer(s, &world.World{})
	})
}
