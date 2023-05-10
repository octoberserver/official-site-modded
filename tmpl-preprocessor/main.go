package main

import (
	"fmt"
	"html/template"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

var template_data = map[string]map[string]string{
	"src/home.html": {
		"key": "value",
	},
	"src/downloads.html": {
		"Main_Modpack_Name":             "All The Mods - ATM7",
		"Main_Modpack_Launcher_Link":    "",
		"Main_ModPack_Traditional_Link": "",
		"Main_ModPack_Curseforge_Link":  "",
	},
}

func main() {
	cmd := exec.Command("cp", "-r", "../src/", "../dist/src/")
	cmd.Output()

	err := filepath.Walk("../src", func(raw_path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() || !strings.HasSuffix(info.Name(), ".html") {
			return nil
		}

		path := strings.Trim(raw_path, "../")
		fmt.Printf("Processed: %s\n", path)

		f, err := os.OpenFile("../dist/"+path, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0600)
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
