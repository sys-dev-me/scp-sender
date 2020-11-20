package main

import (
	"fmt"
	"strings"
	"os"
	"database/sql"
	//"context"
	"time"
	_ "github.com/go-sql-driver/mysql"
)


func (this *Artifact ) SetConfig ( cfg Config ) {
	this.Config = cfg
}

func ( this *Artifact ) newArtifact() *Artifact {
	this.Filename = ""
	this.Path = make([]string, 0 )
	this.Project = ""
	this.Build = ""
	this.Tools = *new(Tools)
	this.SQLPattern = "select BUILD.BUILD_TYPE, FULL_KEY, ARTIFACT.BUILD_NUMBER, ARTIFACT.CHAIN_ARTIFACT as SHARED, STORAGE_TAG from BUILD join ARTIFACT on BUILD.FULL_KEY = ARTIFACT.PLAN_KEY join BRS_ARTIFACT_LINK on ARTIFACT.ARTIFACT_ID = BRS_ARTIFACT_LINK.ARTIFACT_ID join BUILDRESULTSUMMARY on BUILDRESULTSUMMARY.BUILDRESULTSUMMARY_ID = BRS_ARTIFACT_LINK.PRODUCERJOBRESULT_ID where GLOBALLY_STORED = FALSE AND ARTIFACT.BUILD_NUMBER=%s AND ARTIFACT.PLAN_KEY=\"%s\" AND ARTIFACT.LABEL=\"%s\"  order by PLAN_KEY, BUILD_NUMBER, SHARED;"

	return this
}

func ( this *Artifact ) buildBasePath ( BambooConfig BambooSettings ) {
	this.addPath ( BambooConfig.MountedDir )
	this.addPath ( BambooConfig.ArtifactFolder )
}

func ( this *Artifact ) findFile () (string, bool) {
	
	fmt.Printf ( "Looking: %v\n", this.getPath() + "/" + this.Filename )

	_, err := os.Lstat( strings.Join([]string{this.getPath(), this.Filename}, "/"))
	if err != nil {
			fmt.Printf ( "Error while stat file: %v\n", err )
			return "", false
		} else {
			//this.Filename = strings.ReplaceAll ( this.Filename, " ", "\\ ")
			return strings.Join([]string{this.getPath(), this.Filename}, "/"), true
	}
}


func (this *Artifact ) ExecSQL() string {

	fmt.Printf ( "DB Config: %+v\n", this.Config.Bamboo.DB)

	host := fmt.Sprintf ( "tcp(%s:3306)", this.Config.Bamboo.DB.Host )
	db, err := sql.Open("mysql", fmt.Sprintf ( "%s:%s@%s/%s", this.Config.Bamboo.DB.Username, this.Config.Bamboo.DB.Password, host, this.Config.Bamboo.DB.DatabaseName ) )

	fmt.Printf ( "Status of connection: %v\n", db )
	defer db.Close()

	if err != nil {
		fmt.Printf ( "Error DB open: %s\n", err )
		panic(err)
	}
	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	//var ctx context.Context

	var planName string
	var buildType string
  var fullKey string
  var buildNumber string
  var isShared string

	fmt.Printf ( "Execute query: %s\n", fmt.Sprintf( this.SQLPattern, this.Build, this.Project, this.Label)  )
	err = db.QueryRow(fmt.Sprintf( this.SQLPattern, this.Build, this.Project, this.Label) ).Scan( &buildType, &fullKey, &buildNumber, &isShared, &planName )
	fmt.Printf ( "Possible err: %v\n", err )

	if err != nil {
		fmt.Printf ( "Error while query: %s\n", err)
	}

	fmt.Printf ( "PLAN path: %s\n", planName )
 	return planName	

}

func ( this *Artifact ) setFile ( filename string ) {
	this.Filename = filename
}

func ( this *Artifact ) addPath ( partOfPath string ) {
	this.Path = append ( this.Path, partOfPath )
}

func ( this *Artifact ) setLabel ( rawString string ) {

	this.LabelPath = rawString
	this.Label = strings.ReplaceAll ( rawString, "-", " ")

}

func ( this *Artifact ) getProject ( rawString string ) {
	tmp := strings.Split( rawString, "-")

	this.Project = strings.Join ( tmp [ 0:len( tmp ) - 1 ], "-" )

		buildNumber := strings.Join( tmp[ len(tmp) - 1:len( tmp ) ], "") 
		fmt.Printf ( "Build number is %v\n", buildNumber )
		fmt.Printf ( "Project Key is %s\n", this.Project )
		fmt.Printf ( "Label is: %s\n", this.Label )
		this.Build = buildNumber
		this.Plan = this.ExecSQL()
		fmt.Printf ( "Paln is: %s\n", this.Plan )
		fmt.Printf ( "Template is: %v\n", this.Build )

		this.addPath ( this.Plan )
		this.addPath ( "shared" )
		this.addPath ( "build-"+ this.Tools.fillStringLeft( this.Build,5,"0" ) )
		this.addPath ( this.LabelPath )

}

func ( this *Artifact ) getPath () string {
	return strings.Join ( this.Path, "/" )
}
