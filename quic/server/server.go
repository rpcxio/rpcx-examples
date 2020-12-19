//go run -tags quic server.go
package main

import (
	"crypto/tls"
	"flag"
	"log"

	"github.com/smallnest/rpcx/server"
)

var (
	addr = flag.String("addr", "localhost:8972", "server address")
)

func main() {
	flag.Parse()

	cert, err := tls.LoadX509KeyPair("server.pem", "server.key")
	if err != nil {
		log.Print(err)
		return
	}

	config := &tls.Config{Certificates: []tls.Certificate{cert}}

	s := server.NewServer(server.WithTLSConfig(config))
	s.RegisterName("Arith", new(Arith), "")

	err = s.Serve("quic", *addr)
	if err != nil {
		panic(err)
	}
}

// func generateTLSConfig() (*tls.Config, error) {
// 	key, err := rsa.GenerateKey(rand.Reader, 2048)
// 	if err != nil {
// 		return nil, err
// 	}
// 	template := x509.Certificate{
// 		SerialNumber: big.NewInt(1),
// 		NotBefore:    time.Now(),
// 		NotAfter:     time.Now().Add(time.Hour),
// 		KeyUsage:     x509.KeyUsageCertSign | x509.KeyUsageDigitalSignature,
// 	}
// 	certDER, err := x509.CreateCertificate(rand.Reader, &template, &template, &key.PublicKey, key)
// 	if err != nil {
// 		fmt.Println(err)
// 		return nil, err
// 	}
// 	keyPEM := pem.EncodeToMemory(&pem.Block{
// 		Type: "RSA PRIVATE KEY", Bytes: x509.MarshalPKCS1PrivateKey(key),
// 	})
// 	b := pem.Block{Type: "CERTIFICATE", Bytes: certDER}
// 	certPEM := pem.EncodeToMemory(&b)

// 	tlsCert, err := tls.X509KeyPair(certPEM, keyPEM)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return &tls.Config{
// 		Certificates: []tls.Certificate{tlsCert},
// 	}, nil
// }
