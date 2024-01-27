package main

import (
	"errors"
	"flag"
	"os"
	"strings"

	"github.com/bokiledobri/pig/internal/generator"
)

func parseArgs() (error, *generator.Data) {
	data := generator.NewData()
	fp := flag.NewFlagSet("project", flag.ExitOnError)
	fp.StringVar(&data.AppType, "type", "cli", "Type of project (web|api|cli)")
    
        if len(os.Args)<2{
            return errors.New("Please specify generator type (project)"), nil
        }
	switch os.Args[1] {
	case "project":
        if len(os.Args)<3{
            return errors.New("Please provide the name of the project"), nil
        }
        data.ModName = os.Args[2]
        s := strings.Split(data.ModName, "/")
        data.AppName = s[len(s)-1]
		err := fp.Parse(os.Args[3:])
        data.GenType = "project"

		if data.AppType != "web" && data.AppType != "api" && data.AppType != "cli" {
			data.AppType = "cli"
		}
		return err, data
	}
    return errors.New("Invalid generator type"), nil

}
