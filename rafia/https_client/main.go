package main

import (
	"crypto/tls"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/Stymphalian/ikuaki/rafia/cert"
)

func main() {
	addr := "127.0.0.1:8080"
	clientPrivateFilepath := "../data/client.private.pem"
	clientPublicFilepath := "../data/client.public.pem"
	serverPublicFilepath := "../data/server.public.pem"

	// Generate a self-signed certificate
	keyPair := cert.GenerateRSAKeyOrDie()
	_, certBytes := cert.GenerateRootCertOrDie(keyPair, []string{addr})
	cert.WritePrivateKeyAsPEM(keyPair, clientPrivateFilepath)
	cert.WriteCertAsPEM(certBytes, clientPublicFilepath)

	// Create the server certificate pool
	serverCAPool, err := cert.CreateCertPoolFromPubKeys([]string{serverPublicFilepath})
	if err != nil {
		log.Fatal(err)
	}

	// Create the certificate for this client
	clientCert, err := tls.LoadX509KeyPair(clientPublicFilepath, clientPrivateFilepath)
	if err != nil {
		log.Fatal(err)
	}

	// Create simple http client
	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				RootCAs:      serverCAPool,
				Certificates: []tls.Certificate{clientCert},
			},
		},
	}

	resp, err := client.Get(fmt.Sprintf("https://%s/statusz", addr))
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	val, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("%s\n", string(val))
}
