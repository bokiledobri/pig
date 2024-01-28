package main

import (
	"errors"
	"flag"
	"os"
	"strings"

	"github.com/bokiledobri/pig/internal/generator"
)

func parseArgs() (error, *generator.Generator) {
	g := generator.New()
	fp := flag.NewFlagSet("project", flag.ExitOnError)
	fp.StringVar(&g.AppType, "type", "cli", "Type of project (web|api|cli)")

	if len(os.Args) < 2 {
		return errors.New("Please specify generator type (project)"), nil
	}
	switch os.Args[1] {
	case "project":
		if len(os.Args) < 3 {
			return errors.New("Please provide the name of the project"), nil
		}
		g.ModName = os.Args[2]
		s := strings.Split(g.ModName, "/")
		g.AppName = s[len(s)-1]
		err := fp.Parse(os.Args[3:])
		g.GenType = "project"

		if g.AppType != "web" && g.AppType != "api" && g.AppType != "cli" {
			g.AppType = "cli"
		}
		return err, g
	}
	return errors.New("Invalid generator type"), nil

}
