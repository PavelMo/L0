package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"html/template"
	"log"
	"net/http"
	"os"
	"os/signal"
)

func RunServer() {
	var router = chi.NewRouter()
	var s = &http.Server{
		Addr:    ":8080",
		Handler: router,
	}
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)

	router.Get("/{id}", GetOrderInfo)

	done := make(chan struct{})
	go func() {
		sigint := make(chan os.Signal, 1)
		signal.Notify(sigint, os.Interrupt)
		<-sigint

		if err := s.Shutdown(context.Background()); err != nil {
			log.Printf("HTTP server Shutdown: %v", err)
		}
		close(done)
	}()

	if err := s.ListenAndServe(); err != http.ErrServerClosed {
		log.Printf("HTTP server ListenAndServe: %v", err)
	}
	<-done
}
func GetOrderInfo(w http.ResponseWriter, r *http.Request) {
	OrderStruct := new(OrderInfo)

	params := chi.URLParam(r, "id")

	data, ok := Cache.Get(params)
	if ok {

		jsonByte, err := json.Marshal(data)
		if err != nil {
			fmt.Println(err)
		}
		err = json.Unmarshal(jsonByte, &OrderStruct)
		if err != nil {
			fmt.Println(err)
		}
		tmpl, err := template.ParseFiles("web\static/order_info.html")
		if err != nil {
			fmt.Println(err)
		}
		err = tmpl.Execute(w, OrderStruct)
		if err != nil {
			log.Println(err)
		}
	} else {

		_, err := fmt.Fprintf(w, `<h1>Пользоваетль не найден</h1>`)
		if err != nil {
			log.Println(err)
		}

	}

}
