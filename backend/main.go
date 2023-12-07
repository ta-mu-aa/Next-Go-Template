package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/ta-mu-aa/next-go-template/go-app/db"
)

func main() {
	log.Println("start server...")

	// DB接続
	dsn, err := db.Init()
	if err != nil {
		log.Fatal(err)
	}
	_, err = db.Connect(dsn)
	if err != nil {
		log.Fatal(err)
	}

	r := gin.Default()
	r.GET("/hello", func(context *gin.Context) {
		context.JSON(200, gin.H{
			"message": "Hello World!",
		})
	})
	log.Fatal(r.Run(":8080"))
}