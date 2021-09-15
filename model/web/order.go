package web

import "gorm.io/gorm"

type Order struct {
	gorm.Model
	ID           int    `json:"orderId"`
	CustomerName string `json:"customerName"`
	OrderAt      string `json:"orderedAt"`
}

type Items struct {
	gorm.Model
	ID          int    `json:"itemId"`
	ItemCode    int    `json:"itemCode"`
	Description string `json:"description"`
	Quantity    int    `json:"quantity"`
	OrderId     Order  `json:"orderId" gorm:"foreignKey:OrderId"`
}

// 		  "ID": 0,
//        "CreatedAt": "2021-09-16T00:07:10.406+07:00",
//        "UpdatedAt": "2021-09-16T00:07:10.406+07:00",
//        "DeletedAt": null,
//        "itemId": 2,
//        "itemCode": 0,
//        "description": "",
//        "quantity": 0,
//        "orderId": {
//            "ID": 0,
//            "CreatedAt": "0001-01-01T00:00:00Z",
//            "UpdatedAt": "0001-01-01T00:00:00Z",
//            "DeletedAt": null,
//            "orderId": 0,
//            "customerName": "",
//            "orderAt": ""
//        }
