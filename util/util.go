package util

import "os"

func ReadFile(path string) string {
	content, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}
	return string(content)
}
