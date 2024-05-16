package router

import (
	"file-viewer/controller/tree"

	"github.com/gin-gonic/gin"
)

func MapUrls(engine *gin.Engine) {
	engine.GET("/", tree.StreamFiles)
	engine.GET("/:dir", tree.MoveTo)
}
