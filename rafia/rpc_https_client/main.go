package main

import (
	"context"
	"crypto/tls"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	rafia "github.com/Stymphalian/ikuaki/rafia"
	"github.com/Stymphalian/ikuaki/rafia/cert"
	"github.com/kr/pretty"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

var (
	kAddr                       = "127.0.0.1:8080"
	kClientHttpsPrivateFilepath = "../data/client.https.private.pem"
	kClientHttpsPublicFilepath  = "../data/client.https.public.pem"
	kServerPublicFilepath       = "../data/server.public.pem"
)

func createHttpsClient() *http.Client {
	// Generate a self-signed certificate
	keyPair := cert.GenerateRSAKeyOrDie()
	_, certBytes := cert.GenerateRootCertOrDie(keyPair, []string{kAddr})
	cert.WritePrivateKeyAsPEM(keyPair, kClientHttpsPrivateFilepath)
	cert.WriteCertAsPEM(certBytes, kClientHttpsPublicFilepath)

	// Create the server certificate pool
	serverCAPool, err := cert.CreateCertPoolFromPubKeys(
		[]string{kServerPublicFilepath})
	if err != nil {
		log.Fatal(err)
	}

	// Create the certificate for this client
	clientCert, err := tls.LoadX509KeyPair(kClientHttpsPublicFilepath,
		kClientHttpsPrivateFilepath)
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
	return client
}

func createRpcClient() *grpc.ClientConn {
	serverCerts := []string{kServerPublicFilepath}
	pool, err := cert.CreateCertPoolFromPubKeys(serverCerts)
	if err != nil {
		log.Fatal(err)
	}

	creds := credentials.NewClientTLSFromCert(pool, "")
	dialOptions := []grpc.DialOption{grpc.WithTransportCredentials(creds)}

	conn, err := grpc.Dial(kAddr, dialOptions...)
	if err != nil {
		log.Fatal(err)
	}
	return conn
}

func main() {
	// Create the clients
	httpClient := createHttpsClient()
	conn := createRpcClient()
	defer conn.Close()
	greet := rafia.NewGreeterClient(conn)
	farewell := rafia.NewFarewellerClient(conn)

	// Do http calls
	httpResp, err := httpClient.Get(fmt.Sprintf("https://%s/statusz", kAddr))
	if err != nil {
		log.Fatal(err)
	}
	defer httpResp.Body.Close()
	httpRespText, err := ioutil.ReadAll(httpResp.Body)
	if err != nil {
		log.Fatal(err)
	}

	// Do rpc calls
	ctx := context.Background()
	greetResp, err := greet.Greet(ctx, &rafia.Empty{Text: "RPC Greetings"})
	if err != nil {
		log.Fatal(err)
	}
	farewellResp, err := farewell.Farewell(ctx, &rafia.Empty{Text: "RPC Farewell"})
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("%s\n", string(httpRespText))
	PrettyPrint(greetResp)
	PrettyPrint(farewellResp)
}

// A helper method for Pretty printing any object
func PrettyPrint(v interface{}) {
	fmt.Printf("%# v\n", pretty.Formatter(v))
}
