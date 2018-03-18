package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"path/filepath"

	"github.com/Stymphalian/ikuaki/api"
	"github.com/Stymphalian/ikuaki/api/world"
	pb "github.com/Stymphalian/ikuaki/protos"
	"google.golang.org/grpc"
)

var (
	fPort = flag.Int("port", 0, "The port to run this server on")
)

func ReadFilePrint(filename string) {
	str, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatal("Failed to read api_key", err)
	}
	fmt.Printf("%s contents: = %s\n", filename, str)
}

func main() {
	flag.Parse()
	var port int
	if *fPort == 0 {
		port = api.WORLD_PORT
	} else {
		port = *fPort
	}

	ReadFilePrint("/data/secrets/api_key")
	ReadFilePrint("/data/secrets/api_key2")
	ReadFilePrint("/data/configs/undine.properties")
	files, err := filepath.Glob("/data/configs/*")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(files)

	api.RunServerPortOrDie(port, func(s *grpc.Server) {
		pb.RegisterWorldServer(s, &world.World{})
	})
}
