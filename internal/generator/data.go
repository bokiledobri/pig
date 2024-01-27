package generator

type Data struct {
	AppName string
    ModName string
	AppType string
	GenType string
}

func NewData() *Data {
	return &Data{
		AppType: "cli",
	}
}
