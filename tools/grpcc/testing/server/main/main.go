package main

import (
	"context"
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"strings"
	"time"

	"github.com/Stymphalian/ikuaki/freeport"
	pb "github.com/Stymphalian/ikuaki/tools/grpcc/testing/server"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type Lobby struct{}

func (this *Lobby) CreateWorld(ctx context.Context, r *pb.CreateWorldReq) (
	*pb.CreateWorldRes, error) {
	return &pb.CreateWorldRes{
		Addr: fmt.Sprintf("localhost:%d", 8080),
	}, nil
}

func (this *Lobby) ClientStream(stream pb.Lobby_ClientStreamServer) error {
	names := make([]string, 0)
	for {
		in, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(&pb.CreateWorldRes{
				Addr: fmt.Sprintf("%s:%d", strings.Join(names, " "), 8080),
			})
		}
		if err != nil {
			return err
		}
		names = append(names, in.Name)
		fmt.Printf("client streaming: %v\n", in)
	}
}

func (this *Lobby) ServerStream(req *pb.CreateWorldReq,
	stream pb.Lobby_ServerStreamServer) error {

	for i := 0; i < 10; i++ {
		stream.Send(&pb.CreateWorldRes{
			Addr: fmt.Sprintf("%s:%d", req.GetName(), i),
		})
	}
	return nil
}

func (this *Lobby) BidiStream(stream pb.Lobby_BidiStreamServer) error {
	done := make(chan bool)
	lastName := ""
	go func() {
		for {
			in, err := stream.Recv()
			if err == io.EOF {
				done <- true
			}
			if err != nil {
				continue
			}
			fmt.Printf("client streaming: %v\n", in)
			lastName = in.Name
		}
	}()
	go func() {
		sleeptime := 0
		for {
			time.Sleep(time.Second * 5)
			sleeptime += 5
			stream.Send(&pb.CreateWorldRes{
				Addr: fmt.Sprintf("%s is sleeping %d", lastName, sleeptime),
			})

			select {
			case <-done:
				return
			default:
				continue
			}
		}
	}()

	<-done
	return nil
}

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

	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatal(err)
	}

	s := grpc.NewServer()
	pb.RegisterLobbyServer(s, &Lobby{})
	reflection.Register(s)

	log.Println("Serving lobby: ", lis.Addr().String())
	if err := s.Serve(lis); err != nil {
		log.Fatal(err)
	}
}
