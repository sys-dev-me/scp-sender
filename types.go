package main

import (
	"net/http"
	"html/template"
)

type WebServer struct {
	Server			*http.ServeMux
	Parent			*Application
	Port				int
	Index				*template.Template
}

type Application struct {
	Config			*Config
	Postman			*http.ServeMux
}

type Client struct {
	IP				string
	Username	string
}

type Config struct {	

	Origin			string
	LogFile			string
	Userhome		string
	Clients			[]Client
	Port				int
	HostKey			string
	Web					WebSettings
	
}

type WebSettings struct {
	Port				int
	RootPath		string
}

type Error struct {
	Code				int
	Description	string
}

type Answer struct {
	Terminated	bool
	Errors			[]Error		
	Client			string
	Filename		string
	RemoteFile	string
	Connected		bool
	Status			bool
}

type Respond struct {
	Result		[]Answer
}
