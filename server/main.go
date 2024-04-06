// package main

// import(
// 	"fmt"
// 	"net/http"

// 	"server/post"
// )

// func main() {
// 	fmt.Println(post.Sample)

// 	http.HandleFunc("/api/upload", post.HandlePostRequest)
// 	http.HandleFunc("/v1/upload", post.SamplePostRequest)
// 	http.HandleFunc("/", post.HandleGetRequest)

// 	fmt.Println("Starting server on port 18080...")
// 	// err := http.ListenAndServe(":18080", nil)
// 	err := http.ListenAndServeTLS(
// 		":18080",
// 		"./input/server.crt",
// 		"./input/yourdomain.key",
// 		nil)
// 	if err != nil {
// 		fmt.Println("Failed to start server:", err)
// 	}

// 	fmt.Println("Hello World")
// }

package main

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"server/post"
)

func main() {
	fmt.Println(post.Sample)

	http.HandleFunc("/api/upload", post.HandlePostRequest)
	http.HandleFunc("/v1/upload", post.SamplePostRequest)
	http.HandleFunc("/", post.HandleGetRequest)

	// クライアントからの証明書を検証するためのルート証明書プールをロードします。
	caCert, err := ioutil.ReadFile("./input/ca.pem")
	if err != nil {
		log.Fatalf("読み込みに失敗しました: %s", err)
	}
	caCertPool := x509.NewCertPool()
	caCertPool.AppendCertsFromPEM(caCert)

	// TLS設定
	tlsConfig := &tls.Config{
		ClientCAs:  caCertPool,
		ClientAuth: tls.RequireAndVerifyClientCert, // クライアント証明書が必須であり、検証も行う
	}

	tlsConfig.BuildNameToCertificate()

	server := &http.Server{
		Addr:      ":18080",
		TLSConfig: tlsConfig,
	}

	fmt.Println("Starting mTLS server on port 18080...")
	err = server.ListenAndServeTLS("./input/server.crt", "./input/server.key")
	if err != nil {
		log.Fatalf("サーバーの起動に失敗しました: %v", err)
	}
}
