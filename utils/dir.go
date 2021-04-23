package utils

import (
	"os"
	"strings"
)

//Get application root dir
func GetRootDir() string {
	dir, _ := os.Getwd()
	return strings.Replace(dir, "\\", "/", -1)
}
