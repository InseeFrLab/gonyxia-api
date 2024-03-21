package main

type configuration struct {
	Authentication authentication
	RootPath       string
}

type authentication struct {
	IssuerURI string
	Audience  string
}
