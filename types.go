package main

type Gathering struct {
	URLStatus			bool
	
}

type Client struct {
	IP				string
	Username	string
	Version		string
	Finger		string
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

type DBConfig struct {
	Username			string
	Password			string
	Host					string
	DatabaseName	string
}

type Artifact struct {
	Filename				string
	Remotefile			string
	Project					string
	Build						string
	LabelPath				string
	Label						string
	Plan						string
	Buildname				string
	SQLPattern			string
	Path						[]string
	Config
	Tools	
}
