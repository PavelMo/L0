package main

import (
	"encoding/json"
	"fmt"
	"github.com/nats-io/stan.go"
	"log"
	"math/rand"
	"time"
)

func main() {

	sc, err := stan.Connect("test-cluster", "simple-pub")
	if err != nil {
		log.Println(err)
	}

	defer sc.Close()
	for i := 0; i < 10; i++ {
		//Генерируем структуру заказа
		RandStruct := GetRandJson()

		bytesStruct, err := json.Marshal(RandStruct)
		if err != nil {
			log.Println(err)
		}

		err = sc.Publish("testi", bytesStruct)

		if err != nil {
			log.Fatalf("Error during publish: %v\n", err)
		}

		fmt.Println(RandStruct.OrderUID)
		time.Sleep(500 * time.Millisecond)
	}
	time.Sleep(1 * time.Second)
}

func GetRandJson() *OrderInfo {
	rand.Seed(time.Now().UnixNano())

	s := new(OrderInfo)

	UID := getRandId(19)
	s.OrderUID = UID

	s.TrackNumber = getRandId(13)
	s.Entry = "WBIL"

	s.Delivery.Name = fmt.Sprintf("%s %s", RandName[rand.Intn(len(RandName)-1)], RandSurname[rand.Intn(len(RandSurname)-1)])
	s.Delivery.Phone = "+79" + getRandDigit(9)
	s.Delivery.Zip = getRandId(7)
	s.Delivery.City, s.Delivery.Region = getRandCity()
	s.Delivery.Address = RandAddress[rand.Intn(len(RandAddress)-1)]
	s.Delivery.Email = getRandId(5) + RandMail[rand.Intn(len(RandMail)-1)]

	s.Payment.Transaction = UID
	s.Payment.RequestID = getRandDigit(3)
	s.Payment.Currency = RandCurrency[rand.Intn(len(RandCurrency)-1)]
	s.Payment.Provider = "wbpay"
	s.Payment.PaymentDt = rand.Intn(9999999999-1000000000+1) + 1000000000
	s.Payment.Bank = RandBank[rand.Intn(len(RandBank)-1)]
	s.Payment.DeliveryCost = rand.Intn(2000)
	s.Payment.GoodsTotal = 1
	s.Payment.CustomFee = rand.Intn(20)

	item := Items{
		ChrtID:      rand.Intn(9999999-1000000+1) + 1000000,
		TrackNumber: s.TrackNumber,
		Price:       rand.Intn(999999-100+1) + 100,
		Rid:         getRandId(17),
		Name:        ProductName[rand.Intn(len(ProductName)-1)],
		Sale:        rand.Intn(90),
		Size:        getRandId(2),
		NmID:        rand.Intn(999999-100000+1) + 100000,
		Brand:       ProductBrand[rand.Intn(len(ProductBrand)-1)],
		Status:      202,
	}
	item.TotalPrice = int(float64(item.Price) * (1 - float64(item.Sale)/100))
	s.Payment.Amount = item.TotalPrice
	s.Items = append(s.Items, item)

	s.Locale = "en"
	s.InternalSignature = getRandId(10)
	s.CustomerID = getRandId(9)
	s.DeliveryService = "dhl"
	s.Shardkey = getRandId(1)
	s.SmID = rand.Intn(100)
	s.DateCreated = time.Now()
	s.OofShard = getRandId(1)
	return s
}
func getRandDigit(n int) string {
	id := make([]rune, n)
	for i := 0; i < n; i++ {
		id[i] = RandDigits[rand.Intn(len(RandDigits)-1)]
	}
	return string(id)
}
func getRandId(n int) string {
	orderUID := make([]rune, n)
	for i := 0; i < n; i++ {
		orderUID[i] = RunesRand[rand.Intn(len(RunesRand)-1)]
	}
	return string(orderUID)
}
func getRandCity() (string, string) {
	randD := rand.Intn(len(RandCity) - 1)
	j := 0
	for i, val := range RandCity {
		if j == randD {
			return i, val
		}
		j++
	}
	return "", ""
}
