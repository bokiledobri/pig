package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/bokiledobri/pig/internal/generator"
	"github.com/bokiledobri/pig/internal/template"
)

//Represents types of generator to be presented to user in help/error messager
var genTypes = "project"
//Represents types of project to be presented to user in help/error messager
var projectTypes = "web|api|cli"

//Creates new *generator.Generator instance based on given args and assigns it
//to a.generator
func (a *app) parseArgs() {
	g := generator.New(a.infoLog, a.successLog, a.errorLog)
	a.generator = g
	fp := flag.NewFlagSet("project", flag.ExitOnError)
	fp.StringVar(&g.AppType, "type", "cli", fmt.Sprintf("Type of project (%s)", projectTypes))
    //if  run without arguments, print a friendly error message
	if len(os.Args) < 2 {
		a.errorLog.Printf("Please provide a type of generator (%s)\n", genTypes)
		return
	}
	switch os.Args[1] {
    //if first argument is "project", provide params to Generator to generate
    //a new project
	case "project":
		if len(os.Args) < 3 {
			a.errorLog.Println("Please provide the name of the project")
			return
		}
		g.ModName = os.Args[2]
		s := strings.Split(g.ModName, "/")
		g.AppName = s[len(s)-1]
		err := fp.Parse(os.Args[3:])
		if err != nil {
			a.errorLog.Print(err.Error())
		}
		g.GenType = "project"

		if g.AppType != "web" && g.AppType != "api" && g.AppType != "cli" {
			g.AppType = "cli"
		}
		return
	case "-v":
		fallthrough
	case "version":
		fallthrough
	case "--version":
		a.standardLog.Printf("Pig Is Generator\nversion %s\n", a.version)
		return
	case "-h":
		fallthrough
	case "--help":
		fallthrough
	case "help":
		t := template.New("usage.md")
        t, err :=t.Parse("usage.md")
        if err != nil{
            a.errorLog.Println(err.Error())
        }
        err =t.Execute(os.Stdout, nil)
        if err != nil{
            a.errorLog.Println(err.Error())
        }
	default:
		a.errorLog.Println("Invalid generator type")
		return
	}

}
