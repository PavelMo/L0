package main

import "fmt"

const (
	host     = "localhost"
	user     = "postgres"
	password = "2202"
	dbname   = "L0"
)

var connectionString = fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable", host, user, password, dbname)
