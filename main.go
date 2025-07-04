package main

import (
	"iot-web/handlers"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	r.LoadHTMLGlob("templates/*")

	r.Static("/assets", "./assets")

	r.GET("/", handlers.PageHandler)

	log.Println("Starting server on 0.0.0.0:8085")
	log.Fatal(http.ListenAndServe("0.0.0.0:8085", r))
}
