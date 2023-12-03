package util

import (
	"fmt"
	"os"
	"strings"
)

var quiet bool

func SetQuiet(val bool) {
	quiet = val
}

func ReadFile(path string) []string {
	content, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}
	contentStr := string(content)
	return strings.Split(contentStr, "\n")
}

func Println(out any) {
	if !quiet {
		fmt.Println(out)
	}
}

func PrintLines(lines []string) {
	if !quiet {
		Println(fmt.Sprintf("File contents: \n%s", strings.Join(lines, "\n")))
	}
}
