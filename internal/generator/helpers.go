package generator

import (
	"os"
	"strings"

	"github.com/bokiledobri/pig/internal/template"
)

func (g *Generator) makeFile(fileName string, suffix string, args any) error {

	nameSlice := strings.Split(fileName, "/")
	dirs := nameSlice[:len(nameSlice)-1]
	f := nameSlice[len(nameSlice)-1]
	dirName := g.AppName + "/" + strings.Join(dirs, "/")
    fullFileName := g.AppName+"/"+fileName
    g.InfoLog.Printf("Creating %q\n", fullFileName)
	t := template.New(f + "." + suffix)
	t, err := t.Parse(fileName + "." + suffix)
	if err != nil {
		g.ErrorLog.Printf("Could not create %q: %v\n",fullFileName, err.Error())
		return err
	}
	err = os.MkdirAll(dirName, os.ModePerm)
	if err != nil {
		g.ErrorLog.Printf("Could not create %q: %v\n", fullFileName, err.Error())
		return err
	}
	mk, err := os.Create(fullFileName)
	if err != nil {
		g.ErrorLog.Printf("Could not create %q: %v\n", fullFileName, err.Error())
		return err
	}

	err = t.Execute(mk, args)
	if err != nil {
		g.ErrorLog.Printf("Could not create %q: %v\n", fullFileName, err.Error())
		return err
	}
    g.SuccessLog.Printf("%q successfully created", fullFileName)
    return nil
}
