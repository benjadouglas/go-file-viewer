package tree

import (
	"file-viewer/domain/tree"
)

func StreamFiles(d *tree.Directory) tree.Directory {
	d.Start()
	return *d
}

func MoveTo(d *tree.Directory, dir string) tree.Directory {
	d.MoveDir(dir)
	return *d
}
