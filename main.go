package main

import (
	"file-viewer/router"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.New()
	router.MapUrls(r)
	r.Run(":8080")
}
