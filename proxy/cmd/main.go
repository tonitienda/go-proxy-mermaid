package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/tonitienda/go-proxy-mermaid/controllers"
	"github.com/tonitienda/go-proxy-mermaid/templates"
)

func main() {

	router := gin.Default()
	router.GET("/healthz", func(c *gin.Context) {
		c.Data(200, "text/html; charset=utf-8", []byte("OK"))
	})

	router.GET("/", func(c *gin.Context) {

		fmt.Println("Requesting", c.Request.URL.Path)

		requestorID := c.Query("requestorID")
		fmt.Println(c.Request.URL)
		data := controllers.GetServiceData(requestorID, c.Request.Host)

		fmt.Println("Data", data)
		requestedContentType := c.Request.Header.Get("Content-Type")
		fmt.Println("Request headers", c.Request.Header)
		switch requestedContentType {
		case gin.MIMEJSON:
			fmt.Println("JSON", data)
			c.JSON(200, data)
		default:
			page := templates.GetPage(data)
			fmt.Println(page)
			c.Data(200, "text/html; charset=utf-8", []byte(page))

		}
	})

	router.Run() // listen and serve on 0.0.0.0:8080
}
