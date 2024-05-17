package tree

import (
	DIRECTORY "file-viewer/domain/tree"
	"file-viewer/services/tree"
	"html/template"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

var dir = DIRECTORY.Directory{}

// TODO: send the Directoy struct to the html instead of rendering the HTML every time

func StreamFiles(c *gin.Context) {
	response := tree.StreamFiles(&dir)
	RenderIndex(c, &response)
	c.HTML(http.StatusOK, "index.html", nil)
}

func MoveTo(c *gin.Context) {
	log.Print(c.Params)
	response := tree.MoveTo(&dir, c.Param("dir"))
	RenderIndex(c, &response)
	c.HTML(http.StatusOK, "index.html", nil)
}

func MoveUp(c *gin.Context) {
	response := tree.MoveTo(&dir, "..")
	RenderIndex(c, &response)
	c.HTML(http.StatusOK, "index.html", nil)
}

func RenderIndex(c *gin.Context, r *DIRECTORY.Directory) {
	tmpl := template.Must(template.ParseFiles("views/tree/index.html"))
	tmpl.Execute(c.Writer, r)
}
