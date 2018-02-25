package main

import (
	"context"
	"crypto/tls"
	"fmt"
	"log"
	"net"
	"net/http"
	"strings"

	rafia "github.com/Stymphalian/ikuaki/rafia"
	"github.com/Stymphalian/ikuaki/rafia/cert"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

func grpcHandlerFunc(rpcServer *grpc.Server, other http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ct := r.Header.Get("Content-Type")
		if r.ProtoMajor == 2 && strings.Contains(ct, "application/grpc") {
			log.Printf("Serving the RPC")
			rpcServer.ServeHTTP(w, r)
		} else {
			log.Printf("Serving the HTTP")
			other.ServeHTTP(w, r)
		}
	})
}

type Farewell struct{}
type Greet struct{}
type Statusz struct{}

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
func (this *Statusz) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Here is my status"))
}

func createGrpcOptions(pub, priv string) []grpc.ServerOption {
	creds, err := credentials.NewServerTLSFromFile(pub, priv)
	if err != nil {
		log.Fatal(err)
	}
	return []grpc.ServerOption{grpc.Creds(creds)}
}

func createTLSConfig(publicFilepath string) *tls.Config {
	certPool, err := cert.CreateCertPoolFromPubKeys([]string{publicFilepath})
	if err != nil {
		panic(err)
	}
	config := &tls.Config{
		// ClientAuth: tls.RequireAndVerifyClientCert,
		ClientCAs: certPool,
	}
	return config
}

// openssl genrsa -out server.key 2048
// openssl ecparam -genkey -name secp384r1 -out server.key
// openssl req -new -x509 -sha256 -key server.key -out server.crt -days 3650
// openssl req -x509 -nodes -days 365 -newkey rsa:2048 -keyout key.pem -out cert.pem

func main() {
	addr := "localhost:8080"
	clientPublicFilepath := "../data/client.public.pem"
	privateFilepath := "../data/server.private.pem"
	publicFilepath := "../data/server.public.pem"

	// Generate a self-signed certificate
	keyPair := cert.GenerateRSAKeyOrDie()
	_, certBytes := cert.GenerateRootCertOrDie(keyPair, []string{"localhost", "127.0.0.1"})
	cert.WritePrivateKeyAsPEM(keyPair, privateFilepath)
	cert.WriteCertAsPEM(certBytes, publicFilepath)

	// RPC server
	rpcServer := grpc.NewServer(createGrpcOptions(publicFilepath, privateFilepath)...)
	rafia.RegisterFarewellerServer(rpcServer, &Farewell{})
	rafia.RegisterGreeterServer(rpcServer, &Greet{})

	// Http Server
	httpServer := http.NewServeMux()
	httpServer.Handle("/statusz", &Statusz{})

	s := http.Server{
		Addr:      addr,
		Handler:   grpcHandlerFunc(rpcServer, httpServer),
		TLSConfig: createTLSConfig(clientPublicFilepath),
	}

	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Serving on %s\n", lis.Addr().String())
	if err := s.ServeTLS(lis, publicFilepath, privateFilepath); err != nil {
		log.Fatal(err)
	}
}
