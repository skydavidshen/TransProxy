package testing

import (
	"TransProxy/manager"
	"TransProxy/manager/config"
)

const rootDir = "/Users/davidshen/Documents/webroot/TransProxy"

type Testing struct {
	RootDir string
}

func NEW() *Testing {
	t := new(Testing)
	t.RootDir = rootDir
	return t
}

func (t *Testing) Init() {
	manager.TP_ROOT_DIR = t.RootDir
	manager.TP_CONFIG = config.Viper()
}

func (t *Testing) InitConfig() {
	manager.TP_ROOT_DIR = t.RootDir
	manager.TP_CONFIG = config.Viper()
}
