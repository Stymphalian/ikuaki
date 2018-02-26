package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"time"

	"github.com/Stymphalian/ikuaki/api"
	"github.com/Stymphalian/ikuaki/api/agent"
	pb "github.com/Stymphalian/ikuaki/api/protos"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

var (
	fWorldAddr = flag.String("world_addr", "", "The address to the world server")
	fAgentName = flag.String("agent_name", "", "The name of this agent")
	fPort      = flag.Int("port", 0, "The port to run this server on")
)

func runWorldInform(c pb.WorldClient, a *agent.Agent) {
	stream, err := c.Inform(context.Background())
	if err != nil {
		log.Fatalf("Failed to create Inform stream")
	}
	waitc := make(chan struct{})
	go func() {
		for {
			in, err := stream.Recv()
			if err == io.EOF {
				// read done.
				close(waitc)
				return
			}
			if err != nil {
				log.Fatalf("Failed to receive info from world : %v", err)
			}
			log.Printf("Receive info from the world %v\n", in)
		}
	}()
	go func() {
		for {
			log.Println("Sleeeping 5 seconds, sending update to world")
			time.Sleep(5 * time.Second)

			err := stream.Send(&pb.InformReq{
				AgentName: a.AgentName,
				Text:      fmt.Sprintf("the time is %v", time.Now()),
			})
			if err != nil {
				log.Fatalf("Failed to send a message.")
			}
		}
	}()

	for {
		// just loop forever
	}
	// stream.CloseSend()
	// <-waitc
}

func main() {
	flag.Parse()
	if *fWorldAddr == "" {
		*fWorldAddr = fmt.Sprintf("localhost:%v", api.WORLD_PORT)
	}
	if *fAgentName == "" {
		log.Fatalf("--agent_name must be specified")
	}
	if *fPort == 0 {
		log.Fatalf("--port must be specified")
	}

	a := &agent.Agent{AgentName: *fAgentName}

	// Connect and run a thread to the world server
	conn, err := grpc.Dial(*fWorldAddr, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	worldClient := pb.NewWorldClient(conn)
	log.Printf("Connected to world %v", *fWorldAddr)
	go runWorldInform(worldClient, a)

	api.RunServerPortOrDie(*fPort, func(s *grpc.Server) {
		pb.RegisterAgentServer(s, a)
	})
}
