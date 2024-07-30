package config

import (
	"fmt"
	"github.com/BurntSushi/toml"
	"path/filepath"
	"runtime"
)

func init() {
	currentAbPath := getCurrentAbPathByCaller()
	abs, err := filepath.Abs(currentAbPath + "/configV21.toml")
	if err != nil {
		panic("read toml file err: " + err.Error())
		return
	}

	_, err = toml.DecodeFile(abs, &Config)
	if err != nil {
		panic("read toml file err: " + err.Error())
		return
	}

	fmt.Println("init Config: ", Config)
}

func getCurrentAbPathByCaller() string {
	var abPath string
	_, file, _, ok := runtime.Caller(0)
	if ok {
		abPath = filepath.Dir(file)
	}
	return abPath
}
