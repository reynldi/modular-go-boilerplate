package util

import (
	"flag"
	"fmt"
	"github.com/reynldi/modular-go-boilerplate/pkg/application/config"
	"path/filepath"
	"runtime"
)

var (
	_, b, _, _ = runtime.Caller(0)
	Root = filepath.Join(filepath.Dir(b), "../..")
)

func LoadEnv() {
	var configFile string
	flag.StringVar(&configFile, "config", Root + "/secret.yaml", "path to config file")
	flag.Parse()

	if err := config.ValidateConfigPath(configFile); err != nil {
		fmt.Println(err.Error())
	}

	_, err := config.New(configFile)
	if err != nil {
		fmt.Println(err.Error())
	}
}