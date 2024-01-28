package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/bokiledobri/pig/internal/generator"
)

var genTypes = "project"
var projectTypes = "web|api|cli"

func (a *app) parseArgs() {
	g := generator.New(a.infoLog, a.successLog, a.errorLog)
	fp := flag.NewFlagSet("project", flag.ExitOnError)
	fp.StringVar(&g.AppType, "type", "cli", fmt.Sprintf("Type of project (%s)", projectTypes))

	if len(os.Args) < 2 {
		a.errorLog.Printf("Please provide a type of generator (%s)\n", genTypes)
		return
	}
	switch os.Args[1] {
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
		a.generator = g
		return
	default:
		a.errorLog.Println("Invalid generator type")
		return
	}

}
