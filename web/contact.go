package web

import "net/http"

func ContactHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method == "GET" {
		title := "ERAC Energy - Contact Us"
		tpl := "templates/contact.tmpl"
		GenericHandler(w, r, tpl, title)
	}
}
