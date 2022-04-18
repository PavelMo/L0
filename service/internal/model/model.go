package main

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"time"
)

type CacheStruct struct {
	uid       string
	OrderInfo OrderInfo
}

type OrderInfo struct {
	OrderUID    string `json:"order_uid" validate:"required"`
	TrackNumber string `json:"track_number" validate:"required"`
	Entry       string `json:"entry" validate:"required"`
	Delivery    struct {
		Name    string `json:"name" validate:"required"`
		Phone   string `json:"phone" validate:"required"`
		Zip     string `json:"zip" validate:"required"`
		City    string `json:"city" validate:"required"`
		Address string `json:"address" validate:"required"`
		Region  string `json:"region" validate:"required"`
		Email   string `json:"email" validate:"required,email"`
	} `json:"delivery"`
	Payment struct {
		Transaction  string `json:"transaction" validate:"required"`
		RequestID    string `json:"request_id" validate:"required"`
		Currency     string `json:"currency" validate:"required"`
		Provider     string `json:"provider" validate:"required"`
		Amount       int    `json:"amount" validate:"required,gt=-1"`
		PaymentDt    int    `json:"payment_dt" validate:"required,gt=-1"`
		Bank         string `json:"bank" validate:"required"`
		DeliveryCost int    `json:"delivery_cost" validate:"required,gt=-1"`
		GoodsTotal   int    `json:"goods_total" validate:"required,gt=-1"`
		CustomFee    int    `json:"custom_fee" validate:"required,gt=-1"`
	} `json:"payment"`
	Items             []Items
	Locale            string    `json:"locale" validate:"required"`
	InternalSignature string    `json:"internal_signature" validate:"required"`
	CustomerID        string    `json:"customer_id" validate:"required"`
	DeliveryService   string    `json:"delivery_service" validate:"required"`
	Shardkey          string    `json:"shardkey" validate:"required"`
	SmID              int       `json:"sm_id" validate:"required,numeric"`
	DateCreated       time.Time `json:"date_created" validate:"required"`
	OofShard          string    `json:"oof_shard" validate:"required"`
}
type Items struct {
	ChrtID      int    `json:"chrt_id" validate:"required,numeric"`
	TrackNumber string `json:"track_number" validate:"required"`
	Price       int    `json:"price" validate:"required,numeric"`
	Rid         string `json:"rid" validate:"required"`
	Name        string `json:"name" validate:"required"`
	Sale        int    `json:"sale" validate:"required,numeric"`
	Size        string `json:"size" validate:"required"`
	TotalPrice  int    `json:"total_price" validate:"required,numeric"`
	NmID        int    `json:"nm_id" validate:"required,numeric"`
	Brand       string `json:"brand" validate:"required"`
	Status      int    `json:"status" validate:"required,numeric"`
}

func (o OrderInfo) Value() (driver.Value, error) {
	return json.Marshal(o)
}
func (o *OrderInfo) Scan(value interface{}) error {
	b, ok := value.([]byte)
	if !ok {
		return errors.New("type assertion to []byte failed")
	}

	return json.Unmarshal(b, &o)
}
