package engine

import (
	"html/template"
	"log"
	"os"
	"path/filepath"
	gtpl "text/template"

	"github.com/Masterminds/sprig"
	"github.com/redhatrises/dockerfile-template/pkg/dft/api"
)

func funcMap() gtpl.FuncMap {
	f := sprig.TxtFuncMap()

	// Add some extra functionality
	extra := template.FuncMap{}

	for k, v := range extra {
		f[k] = v
	}

	return f
}

func FileTemplate(container api.Container, template string, files string, stdout bool) {
	funcMap := funcMap()
	tname := filepath.Base(template)

	tpl, err := gtpl.New(tname).Funcs(funcMap).ParseFiles(template)
	if err != nil {
		log.Println("template create: ", err)
		return
	}

	if !stdout {
		f, err := os.Create(files)
		if err != nil {
			log.Println("create file: ", err)
			return
		}

		err = tpl.Execute(f, container)
		if err != nil {
			log.Println("template execute: ", err)
			return
		}
		f.Close()
	} else {
		err = tpl.Execute(os.Stdout, container)
		if err != nil {
			log.Println("template execute: ", err)
			return
		}
	}
}
