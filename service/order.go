package service

import (
	"assigment-2-scalable-go/model/web"
	"encoding/json"
	"gorm.io/gorm"
	"io/ioutil"
	"net/http"
)

type Order struct {
	DB *gorm.DB
}

func (o *Order) InsertOrder(w http.ResponseWriter, r *http.Request) web.WebResponse {

	// set security CORS
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	// read body
	c, errRead := ioutil.ReadAll(r.Body)

	//init response
	var response web.WebResponse

	if errRead != nil {
		response.Status = "internal server error"
		response.Code = 500
		response.Data = nil
		return response
	}
	var request web.Order

	errMarshal := json.Unmarshal(c, &request)

	if errMarshal != nil {
		response.Status = "internal server error"
		response.Code = 500
		response.Data = nil
		return response
	}
	// save to db
	o.DB.Create(&request)
	//return success
	response.Status = "created"
	response.Code = 201
	response.Data = "success add new order"
	return response
}

func (o *Order) UpdateOrder(w http.ResponseWriter, r *http.Request) web.WebResponse {
	return web.WebResponse{}
}

func (o *Order) GetOrders(w http.ResponseWriter, r *http.Request) web.WebResponse {
	var response web.WebResponse
	var orders []web.Order

	rows, err := o.DB.Table("orders").
		Joins("left join items i on orders.id = i.order_id").
		Select("orders.id , orders.ordered_at, orders.customer_name, " +
			"i.item_code, i.description,i.quantity").
		Rows()

	if err != nil {
		response.Code = 500
		response.Status = "something wrong"
		response.Data = nil
		return response
	}

	//if len(rows) <= 0 {
	//	response.Code = 204
	//	response.Status = "no content"
	//	response.Data = nil
	//	return response
	//}

	var order web.Order
	order.Items = make([]web.Items, 0)
	for rows.Next() {
		orderItem := web.Items{}
		err := rows.Scan(&order.ID, &order.OrderedAt, &order.CustomerName,
			&orderItem.ItemCode, &orderItem.Description, &orderItem.Quantity)
		if err != nil {
			return web.WebResponse{
				Code:   500,
				Status: "something wrong",
				Data:   nil,
			}
		}
		order.Items = append(order.Items, orderItem)
		orders = append(orders, order)
	}

	response.Code = 200
	response.Status = "success"
	response.Data = orders
	return response
}

func (o *Order) GetOrder(w http.ResponseWriter, r *http.Request) web.WebResponse {
	return web.WebResponse{}
}

func (o *Order) DeleteOrder(w http.ResponseWriter, r *http.Request) web.WebResponse {
	return web.WebResponse{}
}
