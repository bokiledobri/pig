package generator

import (
	"os"
	"strings"

	"github.com/bokiledobri/gomakegen/internal/template"
)


func makeFile(projectName string, fileName string, suffix string, args any) error {

	nameSlice := strings.Split(fileName, "/")
	dirs := nameSlice[:len(nameSlice)-1]
    f := nameSlice[len(nameSlice)-1]
	dirName := projectName + "/" + strings.Join(dirs, "/")
	t := template.New(f+ "." + suffix)
	t, err := t.Parse(fileName + "." + suffix)
	if err != nil {
		return err
	}
	err = os.MkdirAll(dirName, os.ModePerm)
	if err != nil {
		return err
	}
	mk, err := os.Create(projectName + "/" + fileName)
	if err != nil {
		return err
	}
	return t.Execute(mk, args)
}
