package agent

import (
	pb "github.com/Stymphalian/ikuaki/api/protos"
)

type Agent struct {
	AgentName string
}

func (this *Agent) Update(stream pb.Agent_UpdateServer) error {
	return nil
}
