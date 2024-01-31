package template

import (
	"embed"
	"io"
	"text/template"
)

//go:embed templates/*
var fs embed.FS

type Template struct {
	fs       embed.FS
	template *template.Template
}

func New(name string) *Template {
	t := template.New(name)
    t.Delims("<%", "%>")
	rt := &Template{
		fs:       fs,
		template: t,
	}
	return rt
}

func (t *Template) Parse(patterns ...string) (*Template, error) {
	templatesDir := "templates/"
	for i, p := range patterns {
		patterns[i] = templatesDir + p
	}
	temp, err := t.template.ParseFS(t.fs, patterns...)
	tm := *t
	tm.template = temp
	return &tm, err
}

func (t *Template) Execute(w io.Writer, data interface{}) error {
	return t.template.Execute(w, data)
}
