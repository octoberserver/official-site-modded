package main

import (
	"fmt"
	"html/template"
	"os"
	"path/filepath"
	"strings"
)

var template_data = map[string]any{
	"home.html": struct {
		TestText string
	}{
		"Ducky",
	},
}

func main() {
	err := filepath.Walk("../src", func(raw_path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() || !strings.HasSuffix(info.Name(), ".html") {
			return nil
		}

		path := strings.Trim(raw_path, "../")
		fmt.Printf("name: %s\n", path)

		f, err := os.OpenFile("./dist/"+path, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0600)
		if err != nil {
			return err
		}

		var tmpl = template.Must(template.ParseFiles(raw_path))
		tmpl.Execute(f, template_data[path])

		f.Close()
		return nil
	})
	if err != nil {
		fmt.Println(err)
	}
}
