package controller

import (
	"go-mongo/server"

	"github.com/gin-gonic/gin"
)

func MongoInsert(ctx *gin.Context) {
	var data map[string]string
	ctx.BindJSON(&data)

	col := server.NewCollection("test", "users")
	col.Insert(data)
	ctx.JSON(200, data)
}

func MongoDelete(ctx *gin.Context) {
	var data map[string]string
	ctx.BindJSON(&data)
	col := server.NewCollection("test", "users")
	col.Delete(data)

	ctx.JSON(200, data)
}
