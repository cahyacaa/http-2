package client

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"golang.org/x/net/http2"
)

func transport2() *http2.Transport {
	return &http2.Transport{
		TLSClientConfig:    tlsConfigClient(),
		DisableCompression: true,
		AllowHTTP:          false,
	}
}

func tlsConfigClient() *tls.Config {
	crt, err := os.ReadFile("./cert/public.crt")
	if err != nil {
		log.Fatal(err)
	}

	rootCAs := x509.NewCertPool()
	rootCAs.AppendCertsFromPEM(crt)

	return &tls.Config{
		RootCAs:            rootCAs,
		InsecureSkipVerify: true,
		ServerName:         "localhost",
	}
}

func Http2() {
	client := &http.Client{Transport: transport2()}

	res, err := client.Get("https://localhost:8000/welcome")
	if err != nil {
		log.Fatal(err)
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	res.Body.Close()

	fmt.Printf("Code: %d\n", res.StatusCode)
	fmt.Printf("Body: %s\n", body)
}

func Http1() {
	client := &http.Client{}

	res, err := client.Get("https://localhost:8000/welcome")
	if err != nil {
		log.Fatal(err)
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	res.Body.Close()

	fmt.Printf("Code: %d\n", res.StatusCode)
	fmt.Printf("Body: %s\n", body)
}
