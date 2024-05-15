package tree

type File struct {
	Name string `json:"name"`
	Type string `json:"type"`
	Size int64  `json:"size"`
}

type Directory struct {
	Path string `json:"path"`
	Name string `json:"name"`
}

type Tree struct {
	Path string
}

func (t *Tree) MoveTo(d Directory) {

}
