package tree

// TODO: Por ahora vamos a usar una clase para ambos tipos de archivos, luego vamos a implementar una
// clase especifica para c/u
// type FileClass struct {
// 	Name string `json:"name"`
// 	Type string `json:"type"`
// 	Size int64  `json:"size"`
// }

// type DirectoryClass struct {
// 	Path string `json:"path"`
// 	Name string `json:"name"`
// }

// type Tree struct {
// 	Path string
// }

// func (t *Tree) MoveTo(d DirectoryClass) {

// }

type EntryClass struct {
	Name string `json:"name"`
	Type string `json:"type"`
}
