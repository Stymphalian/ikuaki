package main

import (
	"github.com/Stymphalian/ikuaki/api"
	pb "github.com/Stymphalian/ikuaki/api/protos"
	"github.com/Stymphalian/ikuaki/api/world/world"
	"google.golang.org/grpc"
)

func main() {
	api.RunServerPortOrDie(api.WORLD_PORT, func(s *grpc.Server) {
		pb.RegisterWorldServer(s, &world.World{})
	})
}
