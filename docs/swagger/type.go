package swagger

type Spec struct {
	Version          string
	Host             string
	BasePath         string
	Schemes          []string
	Title            string
	Description      string
	InfoInstanceName string
	SwaggerTemplate  string
	LeftDelim        string
	RightDelim       string
}
