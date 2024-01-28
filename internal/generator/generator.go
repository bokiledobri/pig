package generator

import (
	"log"
	"os"
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
	data := ProjectData{
		AppName: g.AppName,
		AppType: g.AppType,
	}

	projectType := g.AppType
	suffix := "min"
	//Generate a Makefile
	err := g.makeFile("Makefile", suffix, data)
	if err != nil {
		return
	}

	//Generate a .gitignore file

	err = g.makeFile(".gitignore", suffix, data)
	if err != nil {
		return
	}
	//Generate main.go file
	err = g.makeFile("cmd/"+projectType+"/main.go", suffix, data)
	if err != nil {
		return
	}

	if projectType == "web" || projectType == "api" {
		//Generate handlers.go file
		err = g.makeFile("cmd/"+projectType+"/handlers.go", suffix, data)
		if err != nil {
			return
		}
	}
	if projectType == "web" {
		//Generate home template
		err = g.makeFile("ui/html/pages/home.tmpl", suffix, data)
		if err != nil {
			return
		}

	}
	g.InfoLog.Printf("cd %s\n", g.AppName)
	err = os.Chdir(g.AppName)
	if err != nil {
		g.ErrorLog.Printf("Could not cd into project directory: %v\n", err.Error())
		return
	}
	g.executeCommand("go mod init")
	g.executeCommand("git init")
	g.SuccessLog.Printf("Successfully created %s\n", g.AppName)
}
