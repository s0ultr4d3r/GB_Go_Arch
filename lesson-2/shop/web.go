package main

import (
	"GB/lesson-2/shop/models"
	"html/template"
	"net/http"
)

type serverWeb struct {
	*server
}

type itemListResponce struct {
	items []*models.Item
}

func (sw *serverWeb) webItemListHandler(w http.ResponseWriter, r *http.Request) {
	filter := sw.parseItemFilterQuery(r)

	items, err := s.rep.ListItems(filter)
	if err != nil && err != repository.ErrNotFound {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if err == repository.ErrNotFound {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	sw.controller.listItemHandler(w, r)
	tmpl, err := template.New("template.html").ParseFiles("template.html")
	tmpl.Execute(w, map[string]interface{}{
		"title": "Magazin",
		"items": items,
	})
}
