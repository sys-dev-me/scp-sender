package main

import (
"os"
"encoding/json"
"fmt"
"errors"
"golang.org/x/crypto/ssh"
)

type Config struct {

  RemoteLocation  string
  Origin          string
  LogFile         string
  Userhome        string
  Clients         []Client
  Port            int
	SSHKeysDir			string
  HostKey         string
  Web             WebSettings
  Bamboo          BambooSettings
  sshConfig       ssh.Config

}

type BambooSettings struct {
  MountedDir      string
  ArtifactFolder  string
	DB							DBConfig
}

type WebSettings struct {
  Port            int
	Build						int
  RootPath        string
  URLVariableName string
  RootURL         string
}


func (this *Config) Load() *Config {

	configFile := "./.settings"

	if _, err := os.Stat ( configFile ); os.IsNotExist ( err ) {

		fmt.Printf ( "Unable to read config: %v\n", err  )
		os.Exit (1)

	}	

	jsonConfig, err := os.Open ( configFile )
	if err != nil {
		fmt.Printf ( "Unable to read config file: %v, probably because of: %v", configFile, err  )
		os.Exit ( 1 )
	}

	defer jsonConfig.Close()

	json.NewDecoder( jsonConfig ).Decode( &this )

	fmt.Printf ( "Loaded: %+v\n", this.Bamboo )

	return this
}

func (this *Config) getLogFileName () string {
	return this.LogFile
}

func (this *Config) Print () {
	fmt.Printf ( "%v\n", this )
}

func (this *Config) readJSON ( fileName string ) (*os.File, error) {

	if fileName == "" {
		return nil, errors.New( "File name must be filled" )
	}

	jsonFile, err := os.Open( fileName )

	if err != nil {
  	  fmt.Println(err)
	}
	//fmt.Println("Successfully Opened %s", fileName)

	defer jsonFile.Close()
	return jsonFile, nil
}


