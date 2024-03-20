package main

type configuration struct {
	Authentication authentication
}

type authentication struct {
	IssuerURI string
	Audience  string
}
