package main

import (
	// "github.com/ev001/go-webservice/controllers"
	// "github.com/ev001/go-webservice/models"
	"/home/student/go/src/github.com/ev001/go-webservice/controllers"
	"github.com/gin-gonic/gin"
	"models"
	"net/http"
)

func main() {
	r := gin.Default()
	// connection db 이후에 코드를 추가해야한다.
	models.ConnectDataBase()
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"data": "hello world"})
	})

	r.GET("/books", controllers.FindBooks) // new

	r.Run()
}
