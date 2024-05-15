package tree

import (
	"file-viewer/domain/tree"
	"fmt"
	"os"
)

func Tree(path string) []tree.EntryClass {
	var files []tree.EntryClass
	rootDir := os.Getenv("HOME")
	dirEntries, err := os.ReadDir(rootDir)
	if err != nil {
		fmt.Println("Error reading directory:", err)
		return files
	}
	for _, entry := range dirEntries {
		if entry.Name()[0] != '.' {
			entryType := "file"
			if entry.IsDir() {
				entryType = "dir"
			}
			files = append(files, tree.EntryClass{
				Name: entry.Name(),
				Type: entryType,
			})
		}
	}
	return files
}
