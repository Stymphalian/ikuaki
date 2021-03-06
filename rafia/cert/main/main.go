// Binary to create self-signed x509 Certificates and write them to files
package main

import (
	"log"

	"github.com/Stymphalian/ikuaki/rafia/cert"
)

func main() {
	keyPair := cert.GenerateRSAKeyOrDie()
	_, certBytes := cert.GenerateRootCertOrDie(keyPair, []string{"localhost:8080"})
	cert.WritePrivateKeyAsPEM(keyPair, "private.key.pem")
	cert.WriteCertAsPEM(certBytes, "cert.public.pem")

	pool, err := cert.CreateCertPoolFromPubKeys([]string{"cert.public.pem"})
	if err != nil {
		log.Fatal("Failed to create cert pool from private and public PEM files ", err)
	}
	log.Println(pool)
}
