package lobby

import (
	"context"
	"fmt"

	pb "github.com/Stymphalian/ikuaki/protos"
)

const (
	kDefaultWorldName = "Earth"
)

type Lobby struct {
	worldAddrs map[string]string
	agentAddrs map[string]string
}

func NewLobby() *Lobby {
	return &Lobby{
		make(map[string]string),
		make(map[string]string),
	}
}

func (this *Lobby) Create(ctx context.Context, r *pb.CreateReq) (*pb.CreateRes, error) {
	return &pb.CreateRes{
		Addr: &pb.Addr{
			Hostport: fmt.Sprintf("localhost:%d", 8080),
		},
	}, nil
}

func (this *Lobby) Destroy(ctx context.Context, r *pb.DestroyReq) (*pb.DestroyRes, error) {
	return &pb.DestroyRes{}, nil
}

func (this *Lobby) List(ctx context.Context, r *pb.ListReq) (*pb.ListRes, error) {
	return &pb.ListRes{}, nil
}

// func (this *Lobby) CreateWorld(ctx context.Context, r *pb.CreateWorldReq) (*pb.CreateWorldRes, error) {
// 	if _, ok := this.worldAddrs[kDefaultWorldName]; ok {
// 		return nil, fmt.Errorf("World already exists")
// 	}

// 	path, err := exec.LookPath("go")
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	port := freeport.GetPortOrDie()
// 	args := []string{
// 		"run",
// 		"../../world/main/main.go",
// 		fmt.Sprintf("--port=%d", port),
// 	}
// 	cmd := exec.Command(path, args...)
// 	cmd.Stdout = os.Stdout
// 	cmd.Stderr = os.Stderr
// 	err = cmd.Start()
// 	if err != nil {
// 		return nil, err
// 	}

// 	resp := &pb.CreateWorldRes{
// 		Addr: fmt.Sprintf("localhost:%d", port),
// 	}
// 	log.Println("Created world on addr: ", resp.Addr)
// 	this.worldAddrs[kDefaultWorldName] = resp.Addr
// 	return resp, nil
// }

// func (this *Lobby) CreateAgent(ctx context.Context, r *pb.CreateAgentReq) (*pb.CreateAgentRes, error) {
// 	if _, ok := this.agentAddrs[r.Name]; ok {
// 		return nil, fmt.Errorf("Agent with name %s already exists", r.Name)
// 	}

// 	path, err := exec.LookPath("go")
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	port := freeport.GetPortOrDie()
// 	args := []string{
// 		"run",
// 		"../../agent/main/main.go",
// 		fmt.Sprintf("--port=%d", port),
// 		fmt.Sprintf("--world_addr=%s", this.worldAddrs[kDefaultWorldName]),
// 		fmt.Sprintf("--agent_name=%s", r.Name),
// 	}
// 	cmd := exec.Command(path, args...)
// 	cmd.Stdout = os.Stdout
// 	cmd.Stderr = os.Stderr
// 	err = cmd.Start()
// 	if err != nil {
// 		return nil, err
// 	}

// 	agentIdBytes := make([]byte, 64)
// 	n, err := rand.Read(agentIdBytes)
// 	if err != nil {
// 		return nil, fmt.Errorf("Failed to generate an agent id")
// 	}
// 	if n != 64 {
// 		return nil, fmt.Errorf("Failed to generate a agent id, not enough bits")
// 	}

// 	resp := &pb.CreateAgentRes{
// 		AgentId: &pb.AgentId{
// 			Id: base64.StdEncoding.EncodeToString(agentIdBytes),
// 		},
// 		Addr: fmt.Sprintf("localhost:%d", port),
// 	}
// 	this.agentAddrs[r.Name] = resp.Addr
// 	log.Println("Created agent on addr: ", resp.Addr)
// 	return resp, nil
// }
