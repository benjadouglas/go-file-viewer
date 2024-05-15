package tree

import (
	"fmt"
	"log"
	"os"
)

func Tree(path string) {
	rootDir := os.Getenv("HOME")
	dirEntries, err := os.ReadDir(rootDir)
	if err != nil {
		fmt.Println("Error reading directory:", err)
		return
	}
	for _, entry := range dirEntries {
		log.Println(entry.Name())
	}
}
