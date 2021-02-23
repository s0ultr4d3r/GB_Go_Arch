package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"GB/lesson-1/shop/models"
	"GB/lesson-1/shop/repository"
	"GB/lesson-1/shop/service"

	"github.com/gorilla/mux"
)

type server struct {
	rep     repository.Repository
	service service.Service
}

func (s *server) createOrderHandler(w http.ResponseWriter, r *http.Request) {
	order := new(models.Order)
	err := json.NewDecoder(r.Body).Decode(order)
	if err != nil {
		log.Println(err)
		json.NewEncoder(w).Encode(map[string]bool{"ok": false})
		return
	}

	order, err = s.service.CreateOrder(order)
	if err != nil {
		log.Println(err)
		json.NewEncoder(w).Encode(map[string]bool{"ok": false})
		return
	}
	json.NewEncoder(w).Encode(map[string]bool{"ok": true})
}

func (s *server) parseOrderFilterQuery(r *http.Request) *repository.OrderFilter {
	filter := &repository.OrderFilter{}

	if limitRaw := r.FormValue("limit"); limitRaw != "" {
		if limitInput, err := strconv.Atoi(limitRaw); err == nil {
			filter.Limit = limitInput
		}
	}
	if filter.Limit == 0 {
		filter.Limit = 5
	}

	if offsetRaw := r.FormValue("offset"); offsetRaw != "" {
		if offsetInput, err := strconv.Atoi(offsetRaw); err == nil {
			filter.Offset = offsetInput
		}
	}

	return filter
}

func (s *server) listOrdersHandler(w http.ResponseWriter, r *http.Request) {
	filter := s.parseOrderFilterQuery(r)

	orders, err := s.rep.ListOrders(filter)
	if err != nil && err != repository.ErrNotFound {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if err == repository.ErrNotFound {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	resp := &ListResponse{
		Payload: orders,
		Limit:   filter.Limit,
		Offset:  filter.Offset,
	}
	if err := json.NewEncoder(w).Encode(resp); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func (s *server) createItemHandler(w http.ResponseWriter, r *http.Request) {
	item := new(models.Item)
	if err := json.NewDecoder(r.Body).Decode(item); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	item, err := s.rep.CreateItem(item)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if err := json.NewEncoder(w).Encode(item); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func (s *server) updateItemHandler(w http.ResponseWriter, r *http.Request) {
	item := new(models.Item)
	if err := json.NewDecoder(r.Body).Decode(item); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	item, err := s.rep.UpdateItem(item)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if err := json.NewEncoder(w).Encode(item); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func (s *server) deleteItemHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr := vars["id"]

	itemID, err := strconv.ParseInt(idStr, 10, 32)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = s.rep.DeleteItem(int32(itemID))
	if err != nil && err != repository.ErrNotFound {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if err == repository.ErrNotFound {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

func (s *server) getItemHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr := vars["id"]

	itemID, err := strconv.ParseInt(idStr, 10, 32)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	item, err := s.rep.GetItem(int32(itemID))
	if err != nil && err != repository.ErrNotFound {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if err == repository.ErrNotFound {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	if err := json.NewEncoder(w).Encode(item); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func (s *server) parseItemFilterQuery(r *http.Request) *repository.ItemFilter {
	filter := &repository.ItemFilter{}

	if limitRaw := r.FormValue("limit"); limitRaw != "" {
		if limitInput, err := strconv.Atoi(limitRaw); err == nil {
			filter.Limit = limitInput
		}
	}
	if filter.Limit == 0 {
		filter.Limit = 5
	}

	if offsetRaw := r.FormValue("offset"); offsetRaw != "" {
		if offsetInput, err := strconv.Atoi(offsetRaw); err == nil {
			filter.Offset = offsetInput
		}
	}

	if priceRightRaw := r.FormValue("price_right"); priceRightRaw != "" {
		if priceRightInput, err := strconv.ParseInt(priceRightRaw, 10, 64); err == nil {
			filter.PriceRight = &priceRightInput
		}
	}

	if priceLeftRaw := r.FormValue("price_left"); priceLeftRaw != "" {
		if priceLeftInput, err := strconv.ParseInt(priceLeftRaw, 10, 64); err == nil {
			filter.PriceLeft = &priceLeftInput
		}
	}
	return filter
}

type ListResponse struct {
	Payload interface{} `json:"payload"`
	Limit   int         `json:"limit"`
	Offset  int         `json:"offset"`
}

func (s *server) listItemHandler(w http.ResponseWriter, r *http.Request) {
	filter := s.parseItemFilterQuery(r)

	items, err := s.rep.ListItems(filter)
	if err != nil && err != repository.ErrNotFound {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if err == repository.ErrNotFound {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	resp := &ListResponse{
		Payload: items,
		Limit:   filter.Limit,
		Offset:  filter.Offset,
	}
	if err := json.NewEncoder(w).Encode(resp); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
