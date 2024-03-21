package configuration

type Configuration struct {
	Authentication Authentication
	RootPath       string
	Regions        interface{}
}

type Authentication struct {
	IssuerURI string
	Audience  string
}
