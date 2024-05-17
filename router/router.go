package router

import (
	"file-viewer/controller/tree"

	"github.com/gin-gonic/gin"
)

func MapUrls(engine *gin.Engine) {
	engine.LoadHTMLGlob("views/**/*.html")
	engine.GET("/", tree.StreamFiles)
	engine.GET("/:dir", tree.MoveTo)
	engine.GET("/up", tree.MoveUp)
	// engine.GET("/favicon.co", func(c *gin.Context) {
	// 	c.Writer.WriteHeader(http.StatusNoContent)
	// })
}
