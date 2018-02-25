// Helper functions for creating self-signed x509.Certificates
//
// To generate a cert/private key pair and write them to file
//   keyPair := cert.GenerateRSAKeyOrDie()
//   _, certBytes := cert.GenerateRootCertOrDie(keyPair, []string{"localhost:8080"})
//   cert.WritePrivateKeyAsPEM(keyPair, "private.key.pem")
//   cert.WriteCertAsPEM(certBytes, "cert.public.pem")
//
// To create a CertPool for a given set of cert/private key pairs
//   pool, err := cert.CreateCertPoolFromFiles("cert.pem", "private.pem")
//   if err != nil {
//   	log.Fatal(err)
//   }
//
package cert

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/tls"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/pem"
	"io/ioutil"
	"log"
	"math/big"
	"net"
	"os"
	"time"
)

func GenerateRSAKeyOrDie() *rsa.PrivateKey {
	// generate a new key-pair
	rootKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		log.Fatalf("failed to generate private key: %s", err)
	}
	return rootKey
}

func GenerateRootCertOrDie(rootKey *rsa.PrivateKey, hosts []string) (*x509.Certificate, []byte) {
	numList := new(big.Int).Lsh(big.NewInt(1), 128)
	serialNumber, err := rand.Int(rand.Reader, numList)
	if err != nil {
		log.Fatalf("failed to create serial number: %v\n", err)
	}

	// Very basic template of a x509 certificate
	template := x509.Certificate{
		SerialNumber: serialNumber,
		Subject: pkix.Name{
			Organization: []string{"Jordan Yu"},
		},
		SignatureAlgorithm:    x509.SHA256WithRSA,
		NotBefore:             time.Now(),
		NotAfter:              time.Now().Add(24 * time.Hour),
		BasicConstraintsValid: true,

		IsCA:        true,
		KeyUsage:    x509.KeyUsageKeyEncipherment | x509.KeyUsageDigitalSignature | x509.KeyUsageCertSign,
		ExtKeyUsage: []x509.ExtKeyUsage{x509.ExtKeyUsageServerAuth, x509.ExtKeyUsageClientAuth},
	}
	// Add whatever host the user requests
	for _, h := range hosts {
		if ip := net.ParseIP(h); ip != nil {
			template.IPAddresses = append(template.IPAddresses, ip)
		} else {
			template.DNSNames = append(template.DNSNames, h)
		}
	}

	// Actually create the certificate as a bytes array
	certBytes, err := x509.CreateCertificate(rand.Reader, &template, &template, &rootKey.PublicKey, rootKey)
	if err != nil {
		log.Fatalf("Faield to create self-signed certificate %v\n", err)
	}

	cert, err := x509.ParseCertificate(certBytes)
	if err != nil {
		log.Fatalf("Failed to parse certificate bytes back into cert: %v", err)
	}
	return cert, certBytes
}

func WritePrivateKeyAsPEM(key *rsa.PrivateKey, outputFilepath string) error {
	keyOut, err := os.OpenFile(outputFilepath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0600)
	if err != nil {
		log.Print("failed to open "+outputFilepath+" for writing:", err)
		return err
	}
	defer keyOut.Close()

	RSAPemBlock := &pem.Block{Type: "RSA PRIVATE KEY", Bytes: x509.MarshalPKCS1PrivateKey(key)}
	pem.Encode(keyOut, RSAPemBlock)
	return nil
}

func WriteCertAsPEM(certBytes []byte, outputFilepath string) error {
	certOut, err := os.Create(outputFilepath)
	if err != nil {
		log.Printf("failed to open "+outputFilepath+" for writing: %s", err)
		return err
	}
	defer certOut.Close()

	pem.Encode(certOut, &pem.Block{Type: "CERTIFICATE", Bytes: certBytes})
	return nil
}

// Reads the public and private PEM filespaths and create a certPool from them.
func CreateCertPoolFromFiles(publicKeyPath string, privateKeyPath string) (
	*x509.CertPool, error) {
	_, err := tls.LoadX509KeyPair(publicKeyPath, privateKeyPath)
	if err != nil {
		return nil, err
	}

	publicKeyBytes, err := ioutil.ReadFile(publicKeyPath)
	if err != nil {
		return nil, err
	}

	certPool := x509.NewCertPool()
	ok := certPool.AppendCertsFromPEM([]byte(publicKeyBytes))
	if !ok {
		panic("Bad certificates")
	}

	return certPool, nil
}
