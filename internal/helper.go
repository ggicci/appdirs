package internal

import (
	"path/filepath"
)

func PathItemAppend(path, dir string) string {
	return filepath.Join(path, dir)
}

func PathItemAppendList(path []string, dir string) []string {
	var newPath []string
	for _, p := range path {
		newPath = append(newPath, PathItemAppend(p, dir))
	}
	return newPath
}
