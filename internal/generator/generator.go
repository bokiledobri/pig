package generator

import (
	"log"
	"os"
	"os/exec"
)

type Generator struct {
	AppName    string
	ModName    string
	AppType    string
	GenType    string
	InfoLog    *log.Logger
	SuccessLog *log.Logger
	ErrorLog   *log.Logger
}

func New(info, success, err *log.Logger) *Generator {
	return &Generator{
		AppType:    "cli",
		InfoLog:    info,
		SuccessLog: success,
		ErrorLog:   err,
	}
}
func (g *Generator) Generate() {
	switch g.GenType {
	case "project":
		g.generateProject()
	default:
		g.ErrorLog.Println("Invalid generator type")
	}
}

// Generates a new project. projectType can be "web", "api" or "cli",
// and defines the type of project to generate
func (g *Generator) generateProject() {
	// Generate files and folders for the new project

	projectType := g.AppType
	suffix := "min"
	//Generate a Makefile
	g.InfoLog.Println("Creating Makefile")
	err := makeFile(g.AppName, "Makefile", suffix, g)
	if err != nil {
		g.ErrorLog.Printf("Could not create Makefile: %v\n", err.Error())
		return
	}

	//Generate a .gitignore file
	g.InfoLog.Println("Creating .gitignore")
	err = makeFile(g.AppName, ".gitignore", suffix, g)
	if err != nil {
		g.ErrorLog.Printf("Could not create .gitignore: %v\n", err.Error())
		return
	}

	//Generate main.go file
	g.InfoLog.Println("Creating main.go")
	err = makeFile(g.AppName, "cmd/"+projectType+"/main.go", suffix, g)
	if err != nil {
		g.ErrorLog.Printf("Could not create main.go: %v\n", err.Error())
		return
	}

	g.InfoLog.Println("Creating handlers.go")
	if projectType == "web" || projectType == "api" {
		//Generate handlers.go file
		err = makeFile(g.AppName, "cmd/"+projectType+"/handlers.go", suffix, g)
		if err != nil {
			g.ErrorLog.Printf("Could not create handlers.go: %v\n", err.Error())
			return
		}
	}
	if projectType == "web" {
		//Generate home template
		g.InfoLog.Println("Creating home.tmpl")
		err = makeFile(g.AppName, "ui/html/pages/home.tmpl", suffix, g)
		if err != nil {
			g.ErrorLog.Printf("Could not create home.tmpl: %v\n", err.Error())
			return
		}

	}
	g.InfoLog.Printf("cd %s\n", g.AppName)
	err = os.Chdir(g.AppName)
	if err != nil {
		g.ErrorLog.Printf("Could not cd into project directory: %v\n", err.Error())
		return
	}
	g.InfoLog.Println("running \"go mod init\"")
	err = exec.Command("go", "mod", "init", g.ModName).Run()
	if err != nil {
		g.ErrorLog.Printf("Could not initialize go module: %v\n", err.Error())
	}

	g.InfoLog.Println("running \"git init\"")
	err = exec.Command("git", "init").Run()

	if err != nil {
	g.ErrorLog.Printf("Could not initialize git repository: %v\n", err.Error())
}
    g.SuccessLog.Printf("Successfully created %s\n", g.AppName)
}
