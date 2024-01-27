package main

import (
	"fmt"

	"github.com/bokiledobri/pig/internal/generator"
)

func main() {
	g := generator.New()
	err, data := parseArgs()
    if err!=nil{
        panic(err)
    }
	switch data.GenType {
	case "project":
		fmt.Println(g.GenerateProject(data))
	}
}
