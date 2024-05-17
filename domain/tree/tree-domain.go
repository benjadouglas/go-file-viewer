package tree

import (
	"os"
	"strings"
)

// TODO: Move the struct to tree-services, if you wanna move to another directory
// you shouldnt be using string manipulation to do so, its prompt to error and adds
// an extra layer of complexity
type Directory struct {
	Path   string  `json:"Path"`
	Entrys []Entry `json:"entrys"`
}

type Entry struct {
	Name  string `json:"name"`
	IsDir bool   `json:"isDir"`
}

func (t *Directory) Start() {
	t.Path = os.Getenv("HOME")
	t.refresh()
}

func (t *Directory) MoveDir(dir string) {
	if dir == ".." {
		lastSlashIndex := strings.LastIndex(t.Path, "/")
		if lastSlashIndex != -1 {
			t.Path = t.Path[:lastSlashIndex]
		}
	} else {
		t.Path = t.Path + "/" + dir
	}
	t.refresh()
}

func (t *Directory) refresh() { // private class method if first letter lowercase
	t.Entrys = nil
	e, err := os.ReadDir(t.Path)
	if err != nil {
		return
	}
	for _, i := range e {
		if i.Name()[0] != '.' {
			t.Entrys = append(t.Entrys, Entry{Name: i.Name(), IsDir: i.IsDir()})
		}
	}
}
