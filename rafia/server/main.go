package main

import (
	"context"
	"crypto/tls"
	"fmt"
	"log"
	"net"
	"net/http"
	"strings"

	pb "github.com/Stymphalian/ikuaki/rafia"
	"github.com/kr/pretty"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

func statuszHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello world"))
}

func grpcHandlerFunc(rpcServer *grpc.Server, other http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		PrettyPrint(r)
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

func (this *Farewell) Farewell(ctx context.Context, r *pb.Empty) (*pb.Empty, error) {
	return &pb.Empty{}, nil
}
func (this *Greet) Greet(ctx context.Context, r *pb.Empty) (*pb.Empty, error) {
	return &pb.Empty{}, nil
}
func (this *Statusz) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Here is my status"))
}

// A helper method for Pretty printing any object
func PrettyPrint(v interface{}) {
	log.Printf("%# v\n", pretty.Formatter(v))
}

func getServerOpts() []grpc.ServerOption {
	// certFile := "cert.pem"
	// keyFile := "key.pem"
	// creds, err := credentials.NewServerTLSFromFile(certFile, keyFile)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// return []grpc.ServerOption{grpc.Creds(creds)}
	// return []grpc.ServerOption{grpc.Creds(credentials.NewClientTLSFromCert(pb.KDemoCertPool, "localhost:8080"))}

	certFile := "cert.pem"
	keyFile := "key.pem"
	creds, err := credentials.NewServerTLSFromFile(certFile, keyFile)
	if err != nil {
		log.Fatal(err)
	}
	return []grpc.ServerOption{grpc.Creds(creds)}
}

func getTLSConfig() *tls.Config {
	// caCert, err := ioutil.ReadFile("client.crt")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// caCertPool := x509.NewCertPool()
	// caCertPool.AppendCertsFromPEM(pb.KCert)
	cfg := &tls.Config{
		ClientAuth: tls.RequireAndVerifyClientCert,
		ClientCAs:  pb.KDemoCertPool,
	}
	return cfg
}

// openssl genrsa -out server.key 2048
// openssl ecparam -genkey -name secp384r1 -out server.key
// openssl req -new -x509 -sha256 -key server.key -out server.crt -days 3650
// openssl req -x509 -nodes -days 365 -newkey rsa:2048 -keyout key.pem -out cert.pem

func main() {
	fmt.Println("Starting http server")

	rpcServer := grpc.NewServer(getServerOpts()...)
	pb.RegisterFarewellerServer(rpcServer, &Farewell{})
	pb.RegisterGreeterServer(rpcServer, &Greet{})

	// lis, err := net.Listen("tcp", "localhost:8080")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// if err := rpcServer.Serve(lis); err != nil {
	// 	log.Fatal(err)
	// }

	normalServer := http.NewServeMux()
	normalServer.Handle("/statusz", &Statusz{})

	s := http.Server{
		Addr:      pb.KDemoAddr,
		Handler:   grpcHandlerFunc(rpcServer, normalServer),
		TLSConfig: getTLSConfig(),
	}

	lis, err := net.Listen("tcp", pb.KDemoAddr)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Serving on %s\n", lis.Addr().String())
	if err := s.Serve(lis); err != nil {
		log.Fatal(err)
	}

	// fmt.Println("Serving on ", lis.Addr)
	// if err := s.ListenAndServe(); err != nil {
	// 	log.Fatal(err)
	// }
}
