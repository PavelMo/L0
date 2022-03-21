package main

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/patrickmn/go-cache"
	"net/http"
)

const (
	host     = "localhost"
	user     = "postgres"
	password = "2202"
	dbname   = "L0"
)

var connectionString = fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable", host, user, password, dbname)
var Cache *cache.Cache
var router = chi.NewRouter()
var s = &http.Server{
	Addr:    ":8080",
	Handler: router,
}
