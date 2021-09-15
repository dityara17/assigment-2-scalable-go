package model

import "gorm.io/gorm"

type Order struct {
	gorm.Model
	ID           int
	CustomerName string
	OrderAt      string
}

type Items struct {
	gorm.Model
	ID          int
	ItemCode    int
	Description string
	Quantity    int
	OrderId     int
	Order       Order `gorm:"foreignKey:OrderId"`
}
