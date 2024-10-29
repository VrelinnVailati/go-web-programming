package main

import (
	"net/http"

	"github.com/VrelinnVailati/oreilly-go-web-programming/chit-chat/routes"
)

func main() {
	mux := http.NewServeMux()
	files := http.FileServer(http.Dir("/public"))

	mux.Handle("/static/", http.StripPrefix("/static/", files))

	// Root routes
	mux.HandleFunc("/", routes.Index)
	mux.HandleFunc("/err", routes.Error)

	// Auth routes
	mux.HandleFunc("/login", routes.Login)
	mux.HandleFunc("/logout", routes.Logout)
	mux.HandleFunc("/signup", routes.Signup)
	mux.HandleFunc("/signup_account", routes.SignupAccount)
	mux.HandleFunc("/authenticate", routes.Authenticate)

	// Threads routes
	mux.HandleFunc("/thread/new", routes.NewThread)
	mux.HandleFunc("/thread/create", routes.CreateThread)
	mux.HandleFunc("/thread/post", routes.PostThread)
	mux.HandleFunc("/thread/read", routes.ReadThread)

	server := &http.Server{
		Addr:    "0.0.0.0:8080",
		Handler: mux,
	}
	_ = server.ListenAndServe()
}
