package freeport

import (
	"log"
	"net"
)

// GetFreePort asks the kernel for a free open port that is ready to use.
func GetPort() (int, error) {
	addr, err := net.ResolveTCPAddr("tcp", "localhost:0")
	if err != nil {
		return 0, err
	}

	l, err := net.ListenTCP("tcp", addr)
	if err != nil {
		return 0, err
	}
	defer l.Close()
	return l.Addr().(*net.TCPAddr).Port, nil
}

func GetPortOrDie() int {
	port, err := GetPort()
	if err != nil {
		log.Fatalf("GetPortOrDie: Failed to get a free port %v", err)
	}
	return port
}
