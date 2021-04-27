package utils

import "fmt"

func PrintObj(obj interface{}, keys ...string) {
	key := ""
	for _, k := range keys {
		key = k
	}

	if key == "" {
		key = "obj"
	}
	fmt.Printf("%s: %v", key, obj)
	println()
}
