package main

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
	"time"
	"strings"
	"os/user"
	"io/ioutil"
)

type Application struct {
  Config      *Config
  Postman     *http.ServeMux
	Version			float32
	Finger			string
	Artifact	
	Gathering
	Status			bool
	Error				string
	PrivateKey	[]byte
}


func (this *Application ) PrintClients () bool{

  for _, v := range this.Config.Clients {
      fmt.Printf ( "Client %s\n:", v.IP )
  }
  return true
}


func ( this *Application ) logger ( logMessage string ) {
  log.Println ( fmt.Sprintf ( "%t, %s", time.Now(), logMessage ) )
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
	this.Artifact.buildBasePath( this.Config.Bamboo )
	this.Gathering.URLStatus = true

	this.Version = 0.5
	p, err := ioutil.ReadFile( strings.Join ( []string{this.Config.SSHKeysDir, this.Config.HostKey}, "/" ) )
	if err != nil {

		this.Error = fmt.Sprintf ( "%s", err )
		this.Status = false

	} else {

		this.Status = true
		this.PrivateKey = p

	}

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
	fmt.Printf ( "Extracted:%v\n",  pURL )
	for idx,v := range pURL {
		fmt.Printf ( "[%v] %s", idx, v )
	}
	fmt.Printf ( "Total: %v\n", len( pURL ) )
	this.Artifact.SetConfig ( *this.Config )

	this.Artifact.setFile ( pURL[ len(pURL) - 1 ])
	this.Artifact.setLabel ( pURL[ 5 ] )
	this.Artifact.getProject ( pURL[ 2 ] )

	this.Artifact.Remotefile = strings.ReplaceAll ( this.Artifact.Filename, " ", "\\ ")

	pathToFile, isFind := this.Artifact.findFile ()
	fmt.Printf ( "Will be copied: %v\n", this.Artifact.Filename )
	fmt.Printf ( "Path: %v\n", pathToFile )
	if isFind {
		this.Artifact.setFile ( pathToFile )
	}

}

