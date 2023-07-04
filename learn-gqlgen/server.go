package main

import (
	"os"

	"learn-gqlgen/di"
	"learn-gqlgen/infra/request"

	"github.com/gin-gonic/gin"
)

const defaultPort = "8080"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	di.InitializeApp()
	r := gin.Default()
	r.POST("/query", request.GraphqlHandler())
	r.GET("/", request.GlaygroundHandler())
	r.Run()
}
