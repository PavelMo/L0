package main

import (
	"database/sql"
	"fmt"
	"github.com/patrickmn/go-cache"
	"log"
)

func getCache(db *sql.DB, c *cache.Cache) {
	rows, err := db.Query("SELECT * FROM json")
	if err != nil {
		fmt.Println(err)
	}
	defer rows.Close()
	

	for rows.Next() {
		cacheStruct := new(CacheStruct)
		if err = rows.Scan(&cacheStruct.uid, &cacheStruct.OrderInfo); err != nil {
			log.Println("Error occurred while scanning rows:", err)
		}
		c.Set(cacheStruct.uid, cacheStruct.OrderInfo, cache.NoExpiration)
	}

}
func Insert(db *sql.DB, o *OrderInfo) {
	_, err := db.Exec("INSERT INTO json (uid,data_json) VALUES($1,$2)", o.OrderUID, o)
	if err != nil {
		log.Println(err)
	}
}
