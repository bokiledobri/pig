package generator

import (
	"os"
	"strings"

	"github.com/bokiledobri/gomakegen/internal/template"
)

type Generator struct {
	name string
}

func New(name string) *Generator {
	g := &Generator{name: name}
	return g
}

func (g *Generator) GenerateProject(projectType string) error {
	// Generate files and folders for the new project

	// Create Makefile
	return makeFile(g.name, "Makefile", "min")
}

func makeFile(projectName string, fileName string, suffix string) error {

	t := template.New(fileName+"."+suffix)
	t, err := t.Parse(fileName+"."+suffix)
	if err != nil {
		return err
	}
	nameSlice := strings.Split(fileName, "/")
	dirs := nameSlice[:len(nameSlice)-1]
	dirName :=projectName+"/"+ strings.Join(dirs, "/")
	err = os.MkdirAll(dirName, os.ModePerm)
	if err != nil {
		return err
	}
	mk, err := os.Create(projectName+"/"+fileName)
	if err != nil {
		return err
	}
	return t.Execute(mk, nil)
}
