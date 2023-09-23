package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/tonitienda/go-proxy-mermaid/controllers"
	"github.com/tonitienda/go-proxy-mermaid/templates"
)

func main() {

	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		data := controllers.GetServiceData()

		fmt.Println("Requesting", c.Request.URL.Path)
		switch c.NegotiateFormat(gin.MIMEHTML, gin.MIMEJSON) {
		case gin.MIMEHTML:
			page := templates.GetPage(data)
			c.Data(200, "text/html; charset=utf-8", []byte(page))
		case gin.MIMEJSON:
			c.JSON(200, data)
		}
	})

	router.Run() // listen and serve on 0.0.0.0:8080
}
