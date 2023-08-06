package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/", func(ctx *gin.Context) {
		msg := "Routes:\nPing:`\\ping`\nParam:`\\{name}\\`"
		ctx.String(http.StatusOK, msg)
	})
	r.GET("/ping", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "Pong")
	})
	r.GET("/:name", func(ctx *gin.Context) {
		name := ctx.Param("name")
		ctx.String(http.StatusOK, "Hello %s! 👋", name)
	})
	log.Fatal(r.Run(":5000"))
}
