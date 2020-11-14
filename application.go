package main

import (
	"fmt"
	"net/http"
	"net/url"
	"time"
	"strings"
	"os/user"
)

type Application struct {
  Config      *Config
  Postman     *http.ServeMux
	Version			float32
	Finger			string
	Artifact	
	Gathering
}


func ( this *Application ) logger ( logMessage string ) {

	fmt.Printf ( "%t, %s\n", time.Now(), logMessage )

}

func (this *Application ) checkOrigin ( r *http.Request ) bool {

	if r.Host == this.Config.Origin {
		fmt.Printf ( "Origin %v granted\n", r.Host )
		return true
	}

	return false
}

func ( this *Application ) setUser () {
	usr, err := user.Current()
	if err !=nil {
		fmt.Printf ( "Cant detect user in system: explain: %v\n", err )
		return
	}

	this.Config.Userhome = usr.HomeDir
	this.logger( fmt.Sprintf ( "Userhome: %v\n", usr.HomeDir ) )
}

func (this *Application) Init () *Application {
	this.Config = new(Config).Load()
	this.Postman = nil
	this.Artifact = *new(Artifact)
	this.Artifact.buildBasePath( this.Config.Bamboo)
	this.Gathering.URLStatus = true
	this.Version = 0.2

	return this
}


func ( this *Application ) URLParse ( urlFileName string) {

	this.Artifact.newArtifact()
	this.Artifact.buildBasePath( this.Config.Bamboo )

	fmt.Printf ( "Parsing: %v\n", urlFileName )
	p, err := url.ParseRequestURI ( urlFileName ) 
	
	if err != nil {
		fmt.Printf ( "[error] Can't validate data: %v\n", err )
		this.Gathering.URLStatus = false
		return
	} else {
		this.Gathering.URLStatus = true
	}
 
	pURL := strings.Split( p.Path, "/" )	
	fmt.Printf ( "Extracted:\n%v\n",  pURL )
	for _,v := range pURL {
		fmt.Printf ( "Val: %s\n", v )
	}

	this.Artifact.setFile ( pURL[ len(pURL) - 1 ])
	this.Artifact.getProject ( pURL[ 2 ] )
	this.Artifact.addPath ( pURL[ len(pURL) - 2 ] )

	this.Artifact.Remotefile = this.Artifact.Filename 
	pathToFile, isFind := this.Artifact.findFile ()
	if isFind {
		this.Artifact.setFile ( pathToFile )
	}

}

