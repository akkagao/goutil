package goutil

import (
	"path/filepath"
	"strings"
)

func CurrentDirName() string {
	path, err := filepath.Abs("./")
	ChkErr(err)
	index := strings.LastIndex(path, "/")
	return path[index+1:]
}
