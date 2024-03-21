package configuration

type Region struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type Configuration struct {
	Authentication Authentication
	RootPath       string
	Regions        []Region
}

type Authentication struct {
	IssuerURI string
	Audience  string
}
