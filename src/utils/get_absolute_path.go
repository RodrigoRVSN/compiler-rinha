package utils

import (
	"path/filepath"
	"runtime"
)

func GetAbsolutePath(relativePath string) string {
	_, filename, _, ok := runtime.Caller(0)

	if !ok {
		panic("Error getting the runtime file path")
	}

	dir := filepath.Dir(filename)
	absolutePath := filepath.Join(dir, relativePath)

	return absolutePath
}
