package main

import (
	"fmt"
)

func main() {
	err, g := parseArgs()
    if err!=nil{
        panic(err)
    }
	switch g.GenType {
	case "project":
		fmt.Println(g.GenerateProject())
	}
}
