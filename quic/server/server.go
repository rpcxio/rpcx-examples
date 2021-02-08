//go run -tags quic .
package main

import (
	"crypto/tls"
	"crypto/x509"
	"flag"
	"io/ioutil"

	"github.com/smallnest/rpcx/server"
)

var (
	addr = flag.String("addr", "localhost:8972", "server address")
)

func main() {
	flag.Parse()

	//CA
	caCertPEM, err := ioutil.ReadFile("../ca.pem")
	if err != nil {
		panic(err)
	}

	roots := x509.NewCertPool()
	ok := roots.AppendCertsFromPEM(caCertPEM)
	if !ok {
		panic("failed to parse root certificate")
	}

	cert, err := tls.LoadX509KeyPair("server.pem", "server.key")
	if err != nil {
		panic(err)
	}

	config := &tls.Config{
		Certificates: []tls.Certificate{cert},
		RootCAs:      roots,
	}

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
