package main

import (
	"fmt"
	"github.com/sevlyar/go-daemon"
	"html"
	"log"
	"net/http"
)

func main() {
	cntext := daemon.Context{
		PidFileName: "daemon.pid",
		PidFilePerm: 0644,
		LogFileName: "",
		LogFilePerm: 0644,
		WorkDir:     "",
		Args:        []string{"[daemon example]"},
	}
	d, err := cntext.Reborn()
	if err != nil {
		log.Fatal("Unable to run: ", err)
	}
	if d != nil {
		return
	}
	http.HandleFunc("/", httpHandler)
	err = http.ListenAndServe("127.0.0.1:8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}

func httpHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("request from %s: %s %q", r.RemoteAddr, r.Method, r.URL)
	_, err := fmt.Fprintf(w, "go-daemon: %q", html.EscapeString(r.URL.Path))
	if err != nil {
		log.Fatal(err)
	}
}
