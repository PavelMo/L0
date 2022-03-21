package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/go-playground/validator/v10"
	_ "github.com/lib/pq"
	"github.com/nats-io/stan.go"
	"github.com/patrickmn/go-cache"
	"log"
	"os"
	"os/signal"
	"time"
)

var Cache *cache.Cache

func main() {
	var connectionString = fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable", host, user, password, dbname)
	Cache = cache.New(5*time.Minute, 10*time.Minute)

	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		log.Fatal("Error occurred while connecting to database:", err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal("Error occurred while checking connection to database:", err)
	}
	//Восстанавливаем кеш из бд
	getCache(db, Cache)
	//Запускаем сервер
	go RunServer()

	sc, err := stan.Connect("test-cluster", "simple-sub")
	if err != nil {
		log.Fatal(err)
	}

	_, err = sc.QueueSubscribe("testi", "group", func(m *stan.Msg) {
		orderStruct := new(OrderInfo)
		err = json.Unmarshal(m.Data, &orderStruct)
		if err != nil {
			log.Println("Error occurred while unmarshalling into json struct:", err)
			return
		}

		err = Validate(orderStruct)

		if err != nil {
			fmt.Println(err)
		} else {

			fmt.Println(orderStruct.OrderUID)
			Insert(db, orderStruct)

			Cache.Set(orderStruct.OrderUID, orderStruct, cache.NoExpiration)
		}

	}, stan.DurableName("my-durable"))
	if err != nil {
		log.Fatal(err)
	}

	signalChan := make(chan os.Signal, 1)
	Done := make(chan bool)
	signal.Notify(signalChan, os.Interrupt)
	go func() {
		for range signalChan {
			fmt.Println("Closing connection...")
			sc.Close()
			Done <- true
		}
	}()
	<-Done

}
func Validate(data *OrderInfo) error {
	validate := validator.New()
	return validate.Struct(data)
}
