package controllers

import (
	"html/template"
	"net/http"
	"path"
)

func (server *Server) LeafletBasicMap(w http.ResponseWriter, r *http.Request) {
	leafletTemplate := append(mainTemplateString, path.Join("views", "maps", "map-leaflet.html"))
	var tmpl = template.Must(template.ParseFiles(leafletTemplate...))

	var data = M{
		"title": baseTitle + "Map Basic",
		"sidebar": "map-basic",
	}

	err := tmpl.ExecuteTemplate(w, "map-leaflet", data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
