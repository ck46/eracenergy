package web

import "net/http"

func BlogHandler(w http.ResponseWriter, r *http.Request) {
	title := "ERAC Energy - Blog"
	tpl := "templates/blog.tmpl"
	GenericHandler(w, r, tpl, title)
}
