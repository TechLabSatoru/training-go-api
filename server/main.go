package main

import(
	"fmt"
	"net/http"
	
	"server/post"
)

func main() {
	fmt.Println(post.Sample)

	http.HandleFunc("/api/upload", post.HandlePostRequest)
	http.HandleFunc("/v1/upload", post.SamplePostRequest)
	http.HandleFunc("/", post.HandleGetRequest)

	fmt.Println("Starting server on port 18080...")
	err := http.ListenAndServe(":18080", nil)
	if err != nil {
		fmt.Println("Failed to start server:", err)
	}

	fmt.Println("Hello World")
}
