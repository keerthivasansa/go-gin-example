package main

import (
	"log"
	"net/http"
	"strconv"

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
	r.GET("/vote/:age", func(ctx *gin.Context) {
		ageStr := ctx.Param("age")
		age, err := strconv.Atoi(ageStr)
		if err != nil || age < 18 {
			ctx.String(http.StatusForbidden, "You can't vote")
			return
		}
		ctx.String(http.StatusOK, "You can vote! 🥳")
	})
	log.Fatal(r.Run(":5000"))
}
