package generator

type Generator struct {
	name string
}

func New(name string) *Generator {
	g := &Generator{name: name}
	return g
}

func (g *Generator) GenerateProject(projectType string) error {
	// Generate files and folders for the new project
	var data struct {
		AppName string
	}
	data.AppName = g.name
	suffix := "min"
	//Generate a Makefile
	err := makeFile(g.name, "Makefile", suffix, data)
	if err != nil {
		return err
	}

	//Generate a .gitignore file
	err = makeFile(g.name, ".gitignore", suffix, data)
	if err != nil {
		return err
	}

	//Generate main.go file
	err = makeFile(g.name, "cmd/"+projectType+"/main.go", suffix, data)
	if err != nil {
		return err
	}

	if projectType == "web" || projectType == "api" {
		//Generate handlers.go file
		err = makeFile(g.name, "cmd/"+projectType+"/handlers.go", suffix, data)
		if err != nil {
			return err
		}
	}
    if projectType == "web" {
        //Generate home template
        err = makeFile(g.name, "ui/html/pages/home.tmpl", suffix, data)
        if err != nil {
            return err
        }

    }
	return nil
}
