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
