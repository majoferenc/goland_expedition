package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/", func(context *gin.Context) {
		context.JSON(200, gin.H{
			"message": "Hello from server",
		})
	})
	fmt.Print("Hello from Go!")
	err := r.Run()
	if err != nil {
		fmt.Print(err)
	}
}
