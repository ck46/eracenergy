package web

import (
	"eracenergy/utils"
	"net/http"
	"path"
)

func ServeFileHandler(res http.ResponseWriter, req *http.Request) {
	fname := path.Base(req.URL.Path)
	filePath := "assets/files/" + fname
	http.ServeFile(res, req, utils.AppFilePath(filePath))
}

func GenericHandler(w http.ResponseWriter, r *http.Request, tpl string, title string) {
	args := utils.Context{
		Title: title,
	}
	utils.TemplateHandler(w, r, tpl, args)
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	title := "ERAC Energy - Home"
	tpl := "templates/home.tmpl"
	GenericHandler(w, r, tpl, title)
}

func AboutHandler(w http.ResponseWriter, r *http.Request) {
	title := "ERAC Energy - About Us"
	tpl := "templates/about.tmpl"
	GenericHandler(w, r, tpl, title)
}

func SolutionsHandler(w http.ResponseWriter, r *http.Request) {
	title := "ERAC Energy - Our Solutions"
	tpl := "templates/solutions.tmpl"
	GenericHandler(w, r, tpl, title)
}

func OffGridHandler(w http.ResponseWriter, r *http.Request) {
	title := "ERAC Energy - OFF-GRID SOLAR PV SOLUTIONS"
	tpl := "templates/off-grid-solar.tmpl"
	GenericHandler(w, r, tpl, title)
}

func SolarHybridHandler(w http.ResponseWriter, r *http.Request) {
	title := "ERAC Energy - SOLAR HYBRID AND BATTERY BACKUP SOLUTIONS"
	tpl := "templates/solar-hybrid.tmpl"
	GenericHandler(w, r, tpl, title)
}

func SolarLightsHandler(w http.ResponseWriter, r *http.Request) {
	title := "ERAC Energy - SOLAR STREET LIGHTS"
	tpl := "templates/solar-street-lights.tmpl"
	GenericHandler(w, r, tpl, title)
}

func SolarPumpsHandler(w http.ResponseWriter, r *http.Request) {
	title := "ERAC Energy - SOLAR PUMPS AND GEYSERS"
	tpl := "templates/solar-pumps.tmpl"
	GenericHandler(w, r, tpl, title)
}

func SolarEnergyHandler(w http.ResponseWriter, r *http.Request) {
	title := "ERAC Energy - SOLAR ENERGY KITS AND SOLAR HOME SYSTEMS"
	tpl := "templates/solar-energy-kits.tmpl"
	GenericHandler(w, r, tpl, title)
}

func TechnicalHandler(w http.ResponseWriter, r *http.Request) {
	title := "ERAC Energy - TECHNICAL & CONSULTANCY SERVICES"
	tpl := "templates/technical-consultancy.tmpl"
	GenericHandler(w, r, tpl, title)
}
