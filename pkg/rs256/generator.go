package rs256

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/asn1"
	"encoding/gob"
	"encoding/pem"
	"fmt"
	"log"
	"os"
	"path"
)

// Generate generate and save RSA 256 key files
func Generate(savePath string, bitSize int) {
	reader := rand.Reader

	key, err := rsa.GenerateKey(reader, bitSize)
	checkError(err)

	publicKey := key.PublicKey

	saveGobKey(path.Join(savePath, "private.key"), key)
	savePEMKey(path.Join(savePath, "private.pem"), key)

	saveGobKey(path.Join(savePath, "public.key"), publicKey)
	savePublicPEMKey(path.Join(savePath, "public.pem"), publicKey)
}

// saveGobKey save gob key(*.key file)
func saveGobKey(fileName string, key interface{}) {
	outFile, err := os.Create(fileName)
	checkError(err)
	defer func(outFile *os.File) {
		err := outFile.Close()
		if err != nil {
			checkError(err)
		}
	}(outFile)

	encoder := gob.NewEncoder(outFile)
	err = encoder.Encode(key)
	checkError(err)
}

// savePEMKey save private pem key(*.pem)
func savePEMKey(fileName string, key *rsa.PrivateKey) {
	outFile, err := os.Create(fileName)
	checkError(err)
	defer func(outFile *os.File) {
		err := outFile.Close()
		if err != nil {
			checkError(err)
		}
	}(outFile)

	var privateKey = &pem.Block{
		Type:  "PRIVATE KEY",
		Bytes: x509.MarshalPKCS1PrivateKey(key),
	}

	err = pem.Encode(outFile, privateKey)
	checkError(err)
}

// savePublicPEMKey save public pem key(*.pem)
func savePublicPEMKey(fileName string, pubKey rsa.PublicKey) {
	asn1Bytes, err := asn1.Marshal(pubKey)
	checkError(err)

	var pemKey = &pem.Block{
		Type:  "PUBLIC KEY",
		Bytes: asn1Bytes,
	}

	pemFile, err := os.Create(fileName)
	checkError(err)
	defer func(pemFile *os.File) {
		err := pemFile.Close()
		if err != nil {
			checkError(err)
		}
	}(pemFile)

	err = pem.Encode(pemFile, pemKey)
	checkError(err)
}

// checkError 예외 처리
func checkError(err error) {
	if err != nil {
		fmt.Println("Fatal error ", err.Error())
		os.Exit(1)
	}
}

// GobDecode 생성된 Gob 키 파일 Decoding
func GobDecode(filename string) map[string]string {
	var pubKey map[string]string

	open, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}

	decoder := gob.NewDecoder(open)

	err = decoder.Decode(&pubKey)
	if err != nil {
		log.Fatal(err)
	}

	return pubKey
}

// PrivatePemDecode 생성된 Private PEM 키 Decoding
func PrivatePemDecode(filename string) *rsa.PrivateKey {
	privatePem, err := os.ReadFile(filename)
	if err != nil {
		log.Println("not exists private.pem")
	}

	b, _ := pem.Decode(privatePem)
	privateKey, _ := x509.ParsePKCS1PrivateKey(b.Bytes)

	return privateKey
}

// PublicPemDecode 생성된 public PEM 키 Decoding
func PublicPemDecode(filename string) *rsa.PublicKey {
	publicPem, err := os.ReadFile(filename)
	publicKey := new(rsa.PublicKey)
	if err != nil {
		log.Fatal(err)
	}

	b, _ := pem.Decode(publicPem)
	_, err = asn1.Unmarshal(b.Bytes, &publicKey)

	if err != nil {
		log.Fatal(err)
	}

	return publicKey
}
