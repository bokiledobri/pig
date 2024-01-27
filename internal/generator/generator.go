package generator

import (
	"fmt"
	"os"
	"os/exec"
)

type Generator struct {
}

// Creates new Generator instance, the name param represents the project name
func New() *Generator {
	return &Generator{}
}

// Generates a new project. projectType can be "web", "api" or "cli",
// and defines the type of project to generate
func (g *Generator) GenerateProject(data *Data) error {
	// Generate files and folders for the new project

	projectType := data.AppType
	suffix := "min"
	//Generate a Makefile
	err := makeFile(data.AppName, "Makefile", suffix, data)
	if err != nil {
		return err
	}

	//Generate a .gitignore file
	err = makeFile(data.AppName, ".gitignore", suffix, data)
	if err != nil {
		return err
	}

	//Generate main.go file
	err = makeFile(data.AppName, "cmd/"+projectType+"/main.go", suffix, data)
	if err != nil {
		return err
	}

	if projectType == "web" || projectType == "api" {
		//Generate handlers.go file
		err = makeFile(data.AppName, "cmd/"+projectType+"/handlers.go", suffix, data)
		if err != nil {
			return err
		}
	}
	if projectType == "web" {
		//Generate home template
		err = makeFile(data.AppName, "ui/html/pages/home.tmpl", suffix, data)
		if err != nil {
			return err
		}

	}
    err = os.Chdir(data.AppName)
    if err !=nil{
        return err
    }
    err = exec.Command("go", "mod", "init", data.ModName).Run()
    err = exec.Command("git",  "init").Run()


	return err
}
