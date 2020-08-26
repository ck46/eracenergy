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

	// solutions
	r.HandleFunc("/off-grid-solar-pv-solutions", web.OffGridHandler).Methods("GET")
	r.HandleFunc("/solar-hybrid-and-battery-backup", web.SolarHybridHandler).Methods("GET")
	r.HandleFunc("/solar-street-lights", web.SolarLightsHandler).Methods("GET")
	r.HandleFunc("/solar-pumps-and-geysers", web.SolarPumpsHandler).Methods("GET")
	r.HandleFunc("/solar-energy-kits-and-home-systems", web.SolarEnergyHandler).Methods("GET")
	r.HandleFunc("/technical-consultancy", web.TechnicalHandler).Methods("GET")

	return r
}
