package main

import (
	"flag"
	"go-mongo/controller"

	"github.com/gin-gonic/gin"
)

var port string

func main() {
	flag.StringVar(&port, "p", ":6666", "port")
	app := gin.Default()
	app.POST("/insert", controller.MongoInsert)

	app.Run(port)
}
