package web

import (
	"gorm.io/gorm"
	"time"
)

type Order struct {
	gorm.Model
	ID           int       `json:"orderId"`
	OrderedAt    time.Time `json:"orderedAt"`
	CustomerName string    `json:"customerName"`
	Items        []Items   `json:"items"`
}

type Items struct {
	gorm.Model
	ID          int    `json:"lineItemId"`
	ItemCode    string `json:"itemCode"`
	Description string `json:"description"`
	Quantity    int    `json:"quantity"`
	OrderId     int    `json:"orderId"`
}
