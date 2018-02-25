package main

import (
	"crypto/tls"
	"log"
	"net/http"

	"github.com/Stymphalian/ikuaki/rafia/cert"
)

type Server struct{}

func (this *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// w.Header().Add("Content", "encoding/text")
	w.Write([]byte("Hello Vin!"))
}

func createClientTLSConfig(publicFilepath string) *tls.Config {
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

func createServerTLSConfig(pub, priv string) *tls.Config {
	cer, err := tls.LoadX509KeyPair(pub, priv)
	if err != nil {
		panic(err)
	}
	config := &tls.Config{Certificates: []tls.Certificate{cer}}
	return config
}

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

	// Create a routing mux
	m := http.NewServeMux()
	m.Handle("/statusz", &Server{})

	// Create the server
	s := http.Server{
		Addr:      addr,
		Handler:   m,
		TLSConfig: createClientTLSConfig(clientPublicFilepath),
	}

	// Create the server cert for the listen
	serverTLSConfig := createServerTLSConfig(publicFilepath, privateFilepath)

	// Listen and serve
	lis, err := tls.Listen("tcp", addr, serverTLSConfig)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Serving on ", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatal(err)
	}
}
