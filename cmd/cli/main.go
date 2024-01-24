package main

import (
	"fmt"

	"github.com/bokiledobri/gomakegen/internal/generator"
)


func main() {
    g := generator.New("myProject")
	fmt.Println(g.GenerateProject("cli"))
}
