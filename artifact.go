package main

import (
	"fmt"
	"strings"
	"os"
)


func ( this *Artifact ) newArtifact() *Artifact {
	this.Filename = ""
	this.Path = make([]string, 0 )
	this.Project = ""
	this.Build = ""
	this.Tools = *new(Tools)

	return this
}

func ( this *Artifact ) buildBasePath ( BambooConfig BambooSettings ) {
	this.addPath ( BambooConfig.MountedDir )
	this.addPath ( BambooConfig.ArtifactFolder )
}

func ( this *Artifact ) findFile () (string, bool) {
	
	fmt.Printf ( "Try find %v/%v\n", this.getPath(), this.Filename )
	_, err := os.Lstat( strings.Join([]string{this.getPath(), this.Filename}, "/"))
	if err != nil {
			fmt.Printf ( "Error while stat file: %v\n", err )
			return "", false
		} else {
			return strings.Join([]string{this.getPath(), this.Filename}, "/"), true
	}
}


func ( this *Artifact ) setFile ( filename string ) {
	this.Filename = filename
}

func ( this *Artifact ) addPath ( partOfPath string ) {
	this.Path = append ( this.Path, partOfPath )
}

func ( this *Artifact ) getProject ( rawString string ) {
	tmp := strings.Split( rawString, "-")
	this.Project = strings.Join ( tmp [ 0:len( tmp ) - 1 ], "-" )
	this.addPath( this.Project + "/shared" )
	buildNumber := strings.Join( tmp[ len(tmp) - 1:len( tmp ) ], "") 
	fmt.Printf ( "Build number is %v\n", buildNumber )

	this.Build = "build-" + this.Tools.fillStringLeft ( fmt.Sprintf ( "%s", buildNumber ) , 5, "0" )
	fmt.Printf ( "Template is: %v\n", this.Build )
	
	this.addPath ( this.Build )
}

func ( this *Artifact ) getPath () string {
	return strings.Join ( this.Path, "/" )
}
