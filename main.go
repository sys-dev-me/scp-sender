package main


import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	scp "github.com/bramvdbogaerde/go-scp"
	"golang.org/x/crypto/ssh"
	"github.com/gorilla/websocket"
)

var srv = websocket.Upgrader{}

func (this *Application ) Respond ( v Client ) (*Answer, error) {

		filename := this.Artifact.Filename
		remoteFileName := this.Artifact.Remotefile
		remoteDestination := "/var/db/filestorage"
		answ := new(Answer)
		answ.Client = v
		answ.Terminated = false

		sshConfig := new(ssh.ClientConfig)
		sshConfig.HostKeyCallback = ssh.InsecureIgnoreHostKey()
		sshConfig.HostKeyAlgorithms = []string{ssh.KeyAlgoRSA}

		signer, err := ssh.ParsePrivateKey( this.PrivateKey )

		sshConfig.User = "DEFAULT_USER_CHANGE_IT_PLEASE"
		sshConfig.Auth = []ssh.AuthMethod{ ssh.PublicKeys ( signer ) }
		sshConfig.Ciphers = []string{"aes128-ctr", "aes192-ctr", "aes256-ctr"}

			sshConfig.User = v.Username
			client := scp.NewClient( fmt.Sprintf ( "%s:22", v.IP ), sshConfig)
			f, _ := os.Open( filename )

			answ.Client = v
			answ.Filename = filename
			answ.RemoteFile = strings.Join([]string{remoteDestination, remoteFileName}, "/")

			err = client.Connect()

			if err != nil {
				fmt.Printf ( "Conection error: %v\n", err )

				answ.Connected = false
				return answ, err

				} else {

				answ.Connected = true

				}

			err = client.CopyFile(f, strings.Join([]string{remoteDestination, remoteFileName}, "/"), "0650")

      if err != nil {
        fmt.Println("Error while copying file\n", err)
				answ.Status = false
      } else {
        fmt.Println ("Finished")
				answ.Status = true
      }

			defer f.Close()
			defer client.Close()

	if this.Gathering.URLStatus == false {
		answ.Status = false
	}

	return answ, nil

}

func (this *Application) DefaultHandler ( w http.ResponseWriter, r *http.Request ) {

	fmt.Printf ( "Received data\n" )

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Sec-WebSocket-Protocol", "Upgrade")
	srv.CheckOrigin = this.checkOrigin
	srv.Subprotocols = []string{ "hey", "Upgrade"}
	c, err := srv.Upgrade(w, r, nil)

	if err != nil {
		this.logger ( fmt.Sprintf ( "Error while create updagrader: %v\n", err ) )
	}

	_, fileName, err := c.ReadMessage()
	this.URLParse( string ( fileName ) )


	for _, v := range this.Config.Clients {

		answer, err := this.Respond( v )
		if err != nil {
			fmt.Printf ( "Can't send due to error: %v\n", err )
			answer.Errors = append (answer.Errors, Error{ Code: 11, Description:fmt.Sprintf ( "%s" , err ) } )
		}
		c.WriteJSON( answer )

	}
	defer c.Close()

}

func main () {
	webserver	:= new( WebServer )
	app				:= new( Application )

	app.Init()
	if !app.Status {
		log.Println ( "Erorr while initializing application configuration: %v\n", app.Error )
		os.Exit( 1 )
	}
	webserver.Init ( app )
	go webserver.Start()

	app.Postman = http.NewServeMux()
	//app.Postman.HandleFunc ( "/parser", app.URLParse )
	app.Postman.HandleFunc ( "/gathering", app.DefaultHandler )
	log.Fatal ( http.ListenAndServe(":3001", app.Postman ) )

}


// custom error type

func (this *Error ) getError ( code int, dsc string ) *Error {
	this.Code = code
	this.Description = dsc

	return this
}
