package main

import (
	"eracenergy/utils"
	"eracenergy/web"
	"net/http"

	"github.com/gorilla/mux"
)

func router() *mux.Router {
	r := mux.NewRouter()
	// serving static files
	staticFileDirectory := http.Dir(utils.AppFilePath("assets/"))

	staticFileHandler := http.StripPrefix("/assets/",
		http.FileServer(staticFileDirectory))
	r.PathPrefix("/assets/").Handler(staticFileHandler).Methods("GET")

	// serve a file
	r.HandleFunc("/favicon.ico", web.ServeFileHandler)
	r.HandleFunc("/robots.txt", web.ServeFileHandler)
	r.HandleFunc("/ads.txt", web.ServeFileHandler)

	// main routes
	r.HandleFunc("/", web.HomeHandler).Methods("GET")
	r.HandleFunc("/about", web.AboutHandler).Methods("GET")
	r.HandleFunc("/solutions", web.SolutionsHandler).Methods("GET")
	r.HandleFunc("/blog", web.BlogHandler).Methods("GET")
	r.HandleFunc("/contact", web.ContactHandler).Methods("GET", "POST")

	return r
}
