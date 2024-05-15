package tree

import (
	"file-viewer/services/tree"

	"github.com/gin-gonic/gin"
)

func Root(c *gin.Context) {
	response := tree.Tree("")
	c.IndentedJSON(200, response)
}
