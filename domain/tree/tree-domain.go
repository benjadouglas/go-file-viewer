package tree

import (
	"log"
	"os"
)

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
	t.update()
}

func (t *Directory) MoveDir(dir string) {
	t.Path = t.Path + "/" + dir
	log.Println(t.Path)
	t.update()
}

func (t *Directory) update() { // private class method if first letter lowercase
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
