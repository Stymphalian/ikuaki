package world

import (
	"context"
	"fmt"
	"io"

	pb "github.com/Stymphalian/ikuaki/api/protos"
)

type World struct {
}

func (this *World) Enter(ctx context.Context, r *pb.EnterReq) (*pb.EnterRes, error) {
	return &pb.EnterRes{}, nil
}
func (this *World) Exit(ctx context.Context, r *pb.ExitReq) (*pb.ExitRes, error) {
	return &pb.ExitRes{}, nil
}
func (this *World) Inform(stream pb.World_InformServer) error {
	for {
		in, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}
		fmt.Printf("%s told world %s\n", in.AgentName, in.Text)
	}
	return nil
}
