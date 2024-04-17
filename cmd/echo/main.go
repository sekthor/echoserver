package main

import (
	"io"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type data struct {
	Method  string
	Host    string
	Path    string
	Headers map[string][]string
	Body    string
}

func handleRequest(c *gin.Context) {
	var d data

	d.Method = c.Request.Method
	d.Headers = c.Request.Header
	d.Host = c.Request.Host
	d.Path = c.Request.URL.Path

	body, err := io.ReadAll(c.Request.Body)
	if err != nil {
		log.Println("could not read body")
	}
	d.Body = string(body)
	c.HTML(http.StatusOK, "index.html", d)
}

func main() {
	router := gin.Default()
	router.LoadHTMLGlob("../../templates/*")
	router.NoRoute(handleRequest)
	router.Run(":8080")
}
