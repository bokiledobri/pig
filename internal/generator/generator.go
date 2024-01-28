package generator

import (
	"os"
	"os/exec"
)

type Generator struct {
	AppName string
    ModName string
	AppType string
	GenType string
}

func New() *Generator {
	return &Generator{
		AppType: "cli",
	}
}

// Generates a new project. projectType can be "web", "api" or "cli",
// and defines the type of project to generate
func (g *Generator) GenerateProject() error {
	// Generate files and folders for the new project

	projectType := g.AppType
	suffix := "min"
	//Generate a Makefile
	err := makeFile(g.AppName, "Makefile", suffix, g)
	if err != nil {
		return err
	}

	//Generate a .gitignore file
	err = makeFile(g.AppName, ".gitignore", suffix, g)
	if err != nil {
		return err
	}

	//Generate main.go file
	err = makeFile(g.AppName, "cmd/"+projectType+"/main.go", suffix, g)
	if err != nil {
		return err
	}

	if projectType == "web" || projectType == "api" {
		//Generate handlers.go file
		err = makeFile(g.AppName, "cmd/"+projectType+"/handlers.go", suffix, g)
		if err != nil {
			return err
		}
	}
	if projectType == "web" {
		//Generate home template
		err = makeFile(g.AppName, "ui/html/pages/home.tmpl", suffix, g)
		if err != nil {
			return err
		}

	}
    err = os.Chdir(g.AppName)
    if err !=nil{
        return err
    }
    err = exec.Command("go", "mod", "init", g.ModName).Run()
    if err !=nil{
        return err
    }
    err = exec.Command("git",  "init").Run()


	return err
}
