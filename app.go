package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	// create a router with default middleware
	r := gin.Default()

	// add handlers
	r.StaticFile("/", "./static/index.html")

	// listen and serve on 0.0.0.0:8888
	r.Run(":8888")
}
