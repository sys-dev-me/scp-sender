package main

import (
	"net/http"
	"html/template"
	"golang.org/x/crypto/ssh"
)

type WebServer struct {
	Server			*http.ServeMux
	Parent			*Application
	Port				int
	Index				*template.Template
}

type Gathering struct {
	URLStatus			bool
	
}

type Client struct {
	IP				string
	Username	string
	Version		string
	Finger		string
}

type Config struct {	

	Origin			string
	LogFile			string
	Userhome		string
	Clients			[]Client
	Port				int
	HostKey			string
	Web					WebSettings
	Bamboo			BambooSettings
	sshConfig		ssh.Config
	
}

type BambooSettings struct {
	MountedDir			string
	ArtifactFolder	string
}

type WebSettings struct {
	Port						int
	RootPath				string
	URLVariableName	string
	RootURL					string
}

type Error struct {
	Code				int
	Description	string
}

type Answer struct {
	Terminated	bool
	Errors			[]Error		
	Client			Client
	Filename		string
	RemoteFile	string
	Connected		bool
	Status			bool
}

type Respond struct {
	Result		[]Answer
}

type Artifact struct {
	Filename				string
	Remotefile			string
	Project					string
	Build						string
	Buildname				string
	Path						[]string
	Tools	
}
