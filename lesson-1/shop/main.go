package main

import (
	"GB/lesson-1/shop/pkg/notification"
	"GB/lesson-1/shop/repository"
	"GB/lesson-1/shop/service"
	"flag"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	var tokenStr string
	flag.StringVar(&tokenStr, "t", "", "token for TelegramApi")
	flag.Parse()
	notif, err := notification.NewTelegramBot(tokenStr, 38266)
	if err != nil {
		log.Fatal(err)
	}
	rep := repository.NewMapDB()
	service := service.NewService(rep, notif)
	s := &server{

		service: service,
		rep:     rep,
	}

	router := mux.NewRouter()

	router.HandleFunc("/items", s.listItemHandler).Methods("GET")
	router.HandleFunc("/items", s.createItemHandler).Methods("POST")
	router.HandleFunc("/items/{id}", s.getItemHandler).Methods("GET")
	router.HandleFunc("/items/{id}", s.deleteItemHandler).Methods("DELETE")
	router.HandleFunc("/items/{id}", s.updateItemHandler).Methods("PUT")

	router.HandleFunc("/orders", s.listOrdersHandler).Methods("GET")
	router.HandleFunc("/orders", s.createOrderHandler).Methods("POST")

	srv := &http.Server{
		Addr:    ":8081",
		Handler: router,
	}
	log.Fatal(srv.ListenAndServe())
}
