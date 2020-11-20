package main

import (
	"os"
	"path/filepath"
)

type Tools struct {

	LookupDirectory			string
	LookupFile					string
	FilePath						string
	Error								error
}


func ( this *Tools ) PadStringLeft (s, p string, count int) string {

    ret := make( []byte, len( p ) * count + len( s ) )

    b := ret[:len( p ) * count ]
    bp := copy( b, p )
    for bp < len( b ) {
        copy( b [bp:], b[:bp] )
        bp *= 2
    }
    copy( ret [ len( b ):], s )
    return string( ret )
}

func (this *Tools ) isProjectDirExist( whatWeAreLooking string ) bool {

	if _, err := os.Stat( whatWeAreLooking ); os.IsNotExist(err) {
		return false
	}

	this.LookupDirectory = whatWeAreLooking
	return true
} 


func (this *Tools) checkFile( path string, info os.FileInfo, err error ) error {
	if err != nil {

			this.Error = err
			return err
	}

	if !info.IsDir() {
		if info.Name() == this.LookupFile {
			this.FilePath = path
			return nil
		}
	}

	return nil
}

func (this *Tools ) Lookup ( fileName string ) (string, bool) {

	this.LookupFile = fileName

	err := filepath.Walk( this.LookupDirectory, this.checkFile )

	if err != nil {
		return this.FilePath, true
	}

	return this.FilePath, false

}

func ( this *Tools ) fillStringLeft ( s string, count int, fillingChar string) string {

	strLen := len ( s )
	fillCount := 0
	if count > strLen {
		fillCount = count - strLen
	}

	result := ""
	for i:=0; i < fillCount; i++ {
		result += fillingChar
	}

	result += s
	return result
}
