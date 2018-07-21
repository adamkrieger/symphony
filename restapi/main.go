package main

import (
	"github.com/gin-gonic/gin"
	"github.com/adamkrieger/symphony/common"
	"os"
)

var (
	instanceID = common.RandASCIIBytes(6)
)

func main(){
	ngn := gin.Default()
	ngn.Use(corsMiddleware())

	ngn.GET("/", root)
	ngn.GET("/ping", ping)
	ngn.Run(":" + os.Getenv("RUNPORT"))
}

func corsMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context){
		ctx.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		ctx.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		ctx.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		ctx.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

		if ctx.Request.Method == "OPTIONS" {
			ctx.AbortWithStatus(204)
			return
		} else {
			ctx.Next()
		}
	}
}

func root(ctx *gin.Context) {
	perReqID:= common.RandASCIIBytes(6)
	ctx.JSON(200, gin.H{"instance":instanceID, "message":perReqID})
}

func ping(ctx *gin.Context) {
	ctx.JSON(200, gin.H{"message":"pong"})
}
