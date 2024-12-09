package server

import (
	"fmt"
	"net/http"
)

const DefaultLiveReloadPath = "/_livereload"

type Config struct {
	WebRoot        string // path to directory to be served
	Port           int64
	LiveReloadPath string
}

type Server struct {
	cnf *Config
	srv http.Server
	cg  *connectionGroup
}

func New(cnf *Config) *Server {

	if cnf.LiveReloadPath == "" {
		cnf.LiveReloadPath = DefaultLiveReloadPath
	}

	s := &Server{
		cnf: cnf,
		cg:  newConnectionGroup(),
	}

	mux := http.NewServeMux()
	mux.Handle("/", http.FileServer(http.Dir(s.cnf.WebRoot))) // fs server
	mux.HandleFunc(cnf.LiveReloadPath, s.liveReloadHandler)   // livereload handler

	s.srv = http.Server{
		Addr:    fmt.Sprintf("localhost:%d", cnf.Port),
		Handler: mux,
	}

	return s
}
