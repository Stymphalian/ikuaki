package main

import (
	"context"
	"fmt"
	"log"
	"net"

	rafia "github.com/Stymphalian/ikuaki/rafia"
	"github.com/Stymphalian/ikuaki/rafia/cert"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

type Farewell struct{}
type Greet struct{}

func (this *Farewell) Farewell(ctx context.Context, r *rafia.Empty) (*rafia.Empty, error) {
	return &rafia.Empty{
		Text: fmt.Sprintf("Farewell %s", r.Text),
	}, nil
}
func (this *Greet) Greet(ctx context.Context, r *rafia.Empty) (*rafia.Empty, error) {
	return &rafia.Empty{
		Text: fmt.Sprintf("Greetings %s", r.Text),
	}, nil
}

func createGrpcOptions(pub, priv string) []grpc.ServerOption {
	creds, err := credentials.NewServerTLSFromFile(pub, priv)
	if err != nil {
		log.Fatal(err)
	}
	return []grpc.ServerOption{grpc.Creds(creds)}
}

func main() {
	addr := "localhost:8080"
	privateFile := "../data/server.private.pem"
	publicFile := "../data/server.public.pem"

	// Generate a self-signed certificate
	keyPair := cert.GenerateRSAKeyOrDie()
	_, certBytes := cert.GenerateRootCertOrDie(keyPair, []string{"localhost", "127.0.0.1"})
	cert.WritePrivateKeyAsPEM(keyPair, privateFile)
	cert.WriteCertAsPEM(certBytes, publicFile)

	s := grpc.NewServer(createGrpcOptions(publicFile, privateFile)...)
	rafia.RegisterFarewellerServer(s, &Farewell{})
	rafia.RegisterGreeterServer(s, &Greet{})

	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Serving on %s", lis.Addr().String())
	if err := s.Serve(lis); err != nil {
		log.Fatal(err)
	}
}
