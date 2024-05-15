package tree

import (
	"file-viewer/domain/tree"

	"github.com/gin-gonic/gin"
)

func Tree(c *gin.Context) {
	var dir tree.Directory
	c.BindJSON(&dir)
}
