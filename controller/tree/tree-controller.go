package tree

import (
	DIRECTORY "file-viewer/domain/tree"
	"file-viewer/services/tree"

	"github.com/gin-gonic/gin"
)

var dir = DIRECTORY.Directory{}

func StreamFiles(c *gin.Context) {
	response := tree.StreamFiles(&dir)
	c.IndentedJSON(200, response)
}

func MoveTo(c *gin.Context) {
	response := tree.MoveTo(&dir, c.Param("dir"))
	c.IndentedJSON(200, response)
}
