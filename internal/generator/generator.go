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

// Returns a pointer to the basic generator.Generator instance
func New(info, success, err *log.Logger) *Generator {
	return &Generator{
		GenType:    "",
		AppType:    "cli",
		InfoLog:    info,
		SuccessLog: success,
		ErrorLog:   err,
	}
}

// invokes correct generator based on the value of GenType
func (g *Generator) Generate() {
	switch g.GenType {
	case "":
		return
	case "project":
		g.generateProject()
	default:
		g.ErrorLog.Println("Invalid generator type")
	}
}

// Generates a new project
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

	if projectType == "cli" {
		//Generate main_test.go file
		err = g.makeFile("cmd/"+projectType+"/main_test.go", suffix, data)
		if err != nil {
			return
		}
	}

	if projectType == "web" || projectType == "api" {
		//Generate handlers.go file
		err = g.makeFile("cmd/"+projectType+"/handlers.go", suffix, data)
		if err != nil {
			return
		}
	}
	if projectType == "web" {
		//Generate templates
		err = g.makeFile("ui/html/base.tmpl", suffix, data)
		if err != nil {
			return
		}

		err = g.makeFile("ui/html/pages/home.tmpl", suffix, data)
		if err != nil {
			return
		}

		err = g.makeFile("ui/html/partials/header.tmpl", suffix, data)
		if err != nil {
			return
		}

		err = g.makeFile("ui/html/partials/footer.tmpl", suffix, data)
		if err != nil {
			return
		}

        g.generateAssets()


	}
	//Run "go mod init" and "git init" in project directory
	g.InfoLog.Printf("cd %s\n", g.AppName)
	err = os.Chdir(g.AppName)
	if err != nil {
		g.ErrorLog.Printf("Could not cd into project directory: %v\n", err.Error())
		return
	}
	g.executeCommand("go mod init " + g.ModName)
	g.executeCommand("git init")
	g.SuccessLog.Printf("Successfully created %s\n", g.AppName)
}
