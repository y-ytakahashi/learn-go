package main

import (
	"fmt"
	"net/http"
)

func main() {

	server := http.Server{
		Addr:    "127.0.0.1:8080",
		Handler: nil,
	}

	server.ListenAndServe()

	fmt.Println("hello go")
}
