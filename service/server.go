package main

import (
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"html/template"
	"log"
	"net/http"
)

func RunServer() {
	router := chi.NewRouter()

	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)

	router.Get("/{id}", GetOrderInfo)
	log.Fatal(http.ListenAndServe(":8080", router))
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
		tmpl, err := template.ParseFiles("html_template/order_info.html")
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
