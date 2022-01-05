package hander

import (
	"crypto/tls"
	"crypto/x509"
	"io/ioutil"

	"google.golang.org/grpc/credentials"
)

//获取客户端证书
func GetClientCreds() credentials.TransportCredentials {
	// creds, err := credentials.NewClientTLSFromFile("keys/server.crt", "localhost")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	cert, _ := tls.LoadX509KeyPair("cert/client.crt", "cert/client.key")
	certPool := x509.NewCertPool()
	ca, _ := ioutil.ReadFile("cert/ca.crt")
	certPool.AppendCertsFromPEM(ca)

	creds := credentials.NewTLS(&tls.Config{
		Certificates: []tls.Certificate{cert}, //客服端证书
		ServerName:   "localhost",
		RootCAs:      certPool,
	})
	return creds
}
