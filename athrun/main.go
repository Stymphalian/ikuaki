// Usage:
// go run main.go --port=11111 --output_dir=.
package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"os"

	pb "github.com/Stymphalian/ikuaki/athrun/protos"
	"github.com/Stymphalian/ikuaki/athrun/server"
	"github.com/Stymphalian/ikuaki/freeport"
	"github.com/mgutz/ansi"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var (
	fPort = flag.Int("port", 0, "The port to run this server on. Leave as 0 "+
		"to auto assign a port")
	fGoSrc = flag.String("go_path", "", "The GOPATH/src which is used as the "+
		"root directory from which build requests should be made. If left as "+
		"empty string then look in the environment variables")
	fOutputDir = flag.String("output_dir", "", "The destination directory in "+
		"which to save the built binaries to.")
)

// A simple function which makes sure all required flags are provided and set
// to sensible values.
func checkAndSetFlags() {
	if *fPort == 0 {
		port, err := freeport.GetPort()
		if err != nil {
			log.Fatalf("failed to listen: %v", err)
		}
		*fPort = port
	}
	if *fGoSrc == "" {
		path, ok := os.LookupEnv("GOPATH")
		if !ok {
			log.Fatalf("Failed to find GOPATH in environment vairables")
		}
		*fGoSrc = path + "/src"
	}
	if *fOutputDir == "" {
		log.Fatal("Must specify --output_dir")
	}
}

func main() {
	// check for all the flags
	flag.Parse()
	checkAndSetFlags()

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *fPort))
	if err != nil {
		log.Fatalf("Failed to listen on port %v, %v", *fPort, err)
	}
	srv, err := server.NewServer(lis.Addr().String(), *fGoSrc, *fOutputDir)
	if err != nil {
		log.Fatal(err)
	}

	s := grpc.NewServer()
	pb.RegisterAthrunServer(s, srv)
	reflection.Register(s)

	log.Println(ansi.Yellow, "Serving:\n",
		ansi.Blue, "Address:", ansi.LightWhite, srv.Addr, "\n",
		ansi.Blue, "Output Directory:", ansi.LightWhite, srv.OutputDir, "\n",
		ansi.Blue, "GoSrc:", ansi.LightWhite, srv.GoSrc, ansi.Reset)

	if err := s.Serve(lis); err != nil {
		log.Fatal("Failed to serve ", err)
	}
}
