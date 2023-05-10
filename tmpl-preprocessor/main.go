package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	err := filepath.Walk("../src", func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		fmt.Printf("dir: %v: name: %s\n", !info.IsDir() && strings.HasSuffix(info.Name(), ".html"), path)
		return nil
	})
	if err != nil {
		fmt.Println(err)
	}
}
