package main

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"

	"gopkg.in/yaml.v3"
)

var template_data = map[string]map[string]string{}

// "src/home.html": {
// 	"key": "value",
// },
// "src/downloads.html": {
// 	"Main_Modpack_Name":             "All The Mods - ATM7",
// 	"Main_Modpack_Launcher_Link":    "",
// 	"Main_ModPack_Traditional_Link": "",
// 	"Main_ModPack_Curseforge_Link":  "",
// },

func main() {
	exe_path, err := os.Executable()
	if err != nil {
		log.Println(err)
	}
	// Strip off the executable
	idx := strings.LastIndex(exe_path, "/")
	var preprocessor_root string
	if idx != -1 {
		preprocessor_root = exe_path[:idx]
	} else {
		preprocessor_root = exe_path
	}
	// Strip of the parent directory of the executable to get the project root
	var proj_root string
	idx = strings.LastIndex(preprocessor_root, "/")
	if idx != -1 {
		proj_root = preprocessor_root[:idx]
	} else {
		preprocessor_root = proj_root
	}

	data_buf, err := ioutil.ReadFile(preprocessor_root + "/values.yml")
	if err != nil {
		fmt.Println(err)
	}

	err = yaml.Unmarshal(data_buf, &template_data)
	if err != nil {
		fmt.Println(err)
	}

	/* ------------- */

	fmt.Printf("Executable Path: %s\n", exe_path)
	fmt.Printf("Project Root: %s\n", proj_root)

	// exec.Command("cp", "-r", proj_root+"/src/", proj_root+"/dist/src/").Output()

	err = filepath.Walk(proj_root+"/src", func(raw_path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() || !strings.HasSuffix(info.Name(), ".html") {
			return nil
		}

		path := strings.TrimPrefix(raw_path, proj_root+"/")

		fmt.Printf("Processed: %s\n", path)

		f, err := os.OpenFile(proj_root+"/dist/"+path, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0600)
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
