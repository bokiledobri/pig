package main

import (
	"log"
	"os"

	"github.com/bokiledobri/pig/internal/generator"
)

type app struct {
	infoLog    *log.Logger
	successLog *log.Logger
	errorLog   *log.Logger
	generator  *generator.Generator
}

func main() {
	i := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	s := log.New(os.Stdout, "SUCCESS\t", log.Ldate|log.Ltime)
	e := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)
	a := &app{
		infoLog:    i,
		errorLog:   e,
		successLog: s,
	}
	a.parseArgs()
	a.generator.Generate()
}
