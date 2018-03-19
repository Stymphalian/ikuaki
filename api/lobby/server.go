package lobby

import (
	"context"
	"fmt"
	"strings"

	"github.com/Stymphalian/ikuaki/cloud"
	"github.com/Stymphalian/ikuaki/cloud/kube"
	pb "github.com/Stymphalian/ikuaki/protos"
)

const (
	kDefaultWorldName = "Earth"
)

type Lobby struct {
	worldAddrs map[string]string
	agentAddrs map[string]string
	client     cloud.Clouder
}

func NewLobby() *Lobby {
	// Create a client to our cloud provider. This is where we launch all of
	// our machines.
	client, err := kube.NewKubeCloud()
	if err != nil {
		panic(err)
	}

	return &Lobby{
		make(map[string]string),
		make(map[string]string),
		client,
	}
}

func (this *Lobby) Create(ctx context.Context, r *pb.CreateReq) (*pb.CreateRes, error) {
	var addr cloud.Addr
	var err error
	switch r.ServerType {
	case pb.ServerTypeEnum_WORLD_SERVER:
		addr, err = this.client.Start("../../../folau/world_p.yaml", r.GetArgs())
		if err != nil {
			return nil, err
		}
	case pb.ServerTypeEnum_AGENT_SERVER:
		addr, err = this.client.Start("../../../folau/agent_p.yaml", r.GetArgs())
		if err != nil {
			return nil, err
		}
	default:
		return nil, fmt.Errorf("Unsupported server type %s",
			pb.ServerTypeEnum_name[int32(r.ServerType)])
	}
	fmt.Printf("Created server: %s, %s\n", addr.Id(), addr.Hostport())

	return &pb.CreateRes{
		Id: addr.Id(),
		Addr: &pb.Addr{
			Hostport: addr.Hostport(),
		},
	}, nil
}

func (this *Lobby) Destroy(ctx context.Context, r *pb.DestroyReq) (*pb.DestroyRes, error) {
	err := this.client.Stop(r.Id)
	if err != nil {
		return nil, err
	}
	return &pb.DestroyRes{}, nil
}

func (this *Lobby) List(ctx context.Context, r *pb.ListReq) (*pb.ListRes, error) {
	// Get the addrs from the cloud client
	addrs, err := this.client.List()
	if err != nil {
		return nil, err
	}

	// Depending on the server type we will filter based on the prefix of the
	// ID
	prefix := ""
	switch r.ServerType {
	case pb.ServerTypeEnum_WORLD_SERVER:
		prefix = "ikuaki-world"
	case pb.ServerTypeEnum_AGENT_SERVER:
		prefix = "ikuaki-agent"
	default:
		return nil, fmt.Errorf("Unsupported server type %s",
			pb.ServerTypeEnum_name[int32(r.ServerType)])
	}

	// Extract out all the Addrs into a map
	resp := &pb.ListRes{}
	resp.Servers = make(map[string]*pb.Addr)
	for _, v := range addrs {
		if strings.HasPrefix(v.Id(), prefix) {
			resp.Servers[v.Id()] = &pb.Addr{
				Hostport: v.Hostport(),
			}
		}
	}
	return resp, nil
}
