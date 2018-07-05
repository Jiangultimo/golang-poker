package main

import (
	"log"
	"net/http"
	"apiserver/router"
	"github.com/gin-gonic/gin"
)

func main() {
	g := gin.New()
	middlewares := []gin.HandlerFunc{}

	router.Load(
		g,
		middlewares...
	)

	log.Printf("start to listening the incoming requests on http address: %s", ":8080")
	log.Printf()
}