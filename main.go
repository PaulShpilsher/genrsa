package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"flag"
	"log"
	"os"

	"golang.org/x/crypto/ssh"
)

func main() {
	bits, pemFile, pubFile := initConfig()
	log.Printf("generating RSA key (bit size %d) -%s, %s", bits, pemFile, pubFile)

	privateKey := createPrivateKey(bits)

	pemBinary := getPemBinary(privateKey)

	pubBinary := getPublicKeyBinary(&privateKey.PublicKey)

	writeFile(pemFile, pemBinary)
	log.Printf("private key saved to \"%s\" file", pemFile)

	writeFile(pubFile, pubBinary)
	log.Printf("public key saved to \"%s\" file", pubFile)

}

// initConfig reads options from the command line arguments
func initConfig() (bits int, pemFile string, pubFile string) {
	var output string

	flag.IntVar(&bits, "bits", 4096, "Key size (eg. 4096)")
	flag.StringVar(&output, "output", "rsa", "Output file name base. (eg. specifying \"foo\" writes files \"foo.prv\", \"foo.pem\", and \"foo.pub\")")
	flag.Parse()

	pemFile = output + ".pem"
	pubFile = output + ".pub"
	return bits, pemFile, pubFile
}

// createPrivateKey generates a random RSA private key of the given bit size.
func createPrivateKey(bits int) *rsa.PrivateKey {
	privateKey, err := rsa.GenerateKey(rand.Reader, bits)
	if err != nil {
		log.Fatalf("private key generation failed. err: %v\n", err)
	}

	if err = privateKey.Validate(); err != nil {
		log.Fatalf("private key validation failed. err: %v\n", err)
	}

	log.Println("private key created")
	return privateKey
}

// getPublicKeyBinary take a rsa.PublicKey and return bytes suitable for writing to .pub file
// returns in the format "ssh-rsa ..."
func getPublicKeyBinary(publicKey *rsa.PublicKey) []byte {
	publicRsaKey, err := ssh.NewPublicKey(publicKey)
	if err != nil {
		log.Fatalf("public key creation failed. err: %v\n", err)
	}

	pubKeyBytes := ssh.MarshalAuthorizedKey(publicRsaKey)

	log.Println("public key created")
	return pubKeyBytes
}

// getPemBinary encodes Private Key from RSA to PEM format
func getPemBinary(privateKey *rsa.PrivateKey) []byte {
	pemBinary := pem.EncodeToMemory(&pem.Block{
		Type:    "RSA PRIVATE KEY",
		Headers: nil,
		Bytes:   x509.MarshalPKCS1PrivateKey(privateKey), // Get ASN.1 DER format
	})

	log.Println("PEM created")
	return pemBinary
}

func writeFile(fileName string, data []byte) {
	if err := os.WriteFile(fileName, data, 0600); err != nil {
		log.Fatalf("writing \"%s\" file failed. err: %v\n", fileName, err)
	}
}
