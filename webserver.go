package main


import (
	"fmt"
	"net/http"
	"time"
	"html/template"
)

type WebServer struct {
  Server      *http.ServeMux
  Parent      *Application
  Port        int
  Index       *template.Template
}

func ( this *WebServer ) Init ( parent *Application ) {

	this.Parent = parent
	fmt.Printf ( "Web Server started %v\n", time.Now().Format("01-02-2006 03:04") )

	this.Server = http.NewServeMux()

	this.Server.Handle("/resources/", http.StripPrefix("/resources/", http.FileServer(http.Dir("web/resources"))))

  this.Server.HandleFunc("/", this.DefaultHandler )
	this.Index = template.New( "index")
	
} 


func (this *WebServer) LoadTemplates () bool {
	ptr, err := template.ParseFiles ( this.Parent.Config.Web.RootPath + "/index.tpl")

	if err != nil {
		return true
	}

	this.Index = ptr
	return false
}


func (this *WebServer) Start () {

	fmt.Sprintf ( "Server started on interface: %s\n", fmt.Sprintf ( "%s:%d", "", this.Parent.Config.Web.Port )  )
	http.ListenAndServe( fmt.Sprintf ( "%s:%d", "", this.Parent.Config.Web.Port ), this.Server )
	fmt.Printf ( "Server stopped\n" )

}

func ( this *WebServer ) DefaultHandler ( w http.ResponseWriter, r *http.Request ) {

	err := this.LoadTemplates()
	if err {
		fmt.Fprintf ( w, "[error] unable to load resources: %v\n", err )
		return
	}

	err2 := this.Index.Execute(w, this.Parent )
	if err2 != nil {
		fmt.Printf ( "error execute template: %v\n", err2 ) 
		return
	}
	fmt.Printf ( "error execute template: %v\n", err2 )

}
