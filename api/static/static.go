package static

import (
	"path/filepath"
	"runtime"
)

func GetCurrentAbPathByCaller() string {
	var abPath string
	_, file, _, ok := runtime.Caller(0)
	if ok {
		abPath = filepath.Dir(file)
	}
	return abPath
}
