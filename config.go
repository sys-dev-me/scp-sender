package main

//import "path/filepath"
import "os"
import "encoding/json"
import "fmt"
import "errors"

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


