package service

import (
	"assigment-2-scalable-go/model/web"
	"encoding/json"
	"github.com/gorilla/mux"
	"gorm.io/gorm"
	"io/ioutil"
	"net/http"
	"strconv"
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
	// set security CORS
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	// read body
	c, errRead := ioutil.ReadAll(r.Body)

	var response web.WebResponse

	if errRead != nil {
		response.Code = 500
		response.Status = "invalid body request"
		return response
	}

	var request web.Order
	var order web.Order
	var items web.Items

	errMarshal := json.Unmarshal(c, &request)
	if errMarshal != nil {
		response.Code = 500
		response.Status = "error marshal"
		return response
	}

	err := o.DB.Model(order).Where("orders.id = ?", request.ID).Updates(request).Error
	if err != nil {
		response.Code = 500
		response.Status = "error update to db"
		return response
	}

	for _, x := range request.Items {

		mp := make(map[string]interface{})
		mp["item_code"] = x.ItemCode
		mp["description"] = x.Description
		mp["quantity"] = x.Quantity

		err = o.DB.Model(items).Where("items.id = ?", request.ID).Updates(mp).Error
		if err != nil {
			response.Code = 500
			response.Status = "error update to db"
			return response
		}
	}

	response.Code = 202
	response.Status = "success"
	response.Data = "success update"
	return response
}

func (o *Order) GetOrders(w http.ResponseWriter, r *http.Request) web.WebResponse {
	var response web.WebResponse
	var orders []web.Order
	// set security CORS
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	o.DB.Preload("Items").
		Find(&orders)

	response.Code = 200
	response.Status = "success"
	response.Data = orders
	return response
}

func (o *Order) GetOrder(w http.ResponseWriter, r *http.Request) web.WebResponse {
	// set security CORS
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	// response
	response := web.WebResponse{}

	// handling path ID
	vars := mux.Vars(r)
	orderId, err := strconv.Atoi(vars["id"])
	if err != nil {
		response.Code = 500
		response.Status = "invalid uri"
		response.Data = nil
		return response
	}

	orderData := web.Order{}
	// preload one to many
	o.DB.Preload("Items").Where("orders.id = ?", orderId).
		Find(&orderData)

	// if no result
	if orderData.ID == 0 {
		response.Status = "no result"
		response.Code = 200
		response.Data = nil
		return response
	}

	response.Status = "success"
	response.Code = 200
	response.Data = orderData
	return response
}

func (o *Order) DeleteOrder(w http.ResponseWriter, r *http.Request) web.WebResponse {
	// set security CORS
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	// response
	response := web.WebResponse{}

	// handling path ID
	vars := mux.Vars(r)
	orderId, err := strconv.Atoi(vars["id"])
	if err != nil {
		response.Code = 500
		response.Status = "invalid parameter"
		response.Data = nil
		return response
	}

	order := web.Order{}
	items := web.Items{}

	o.DB.Where("items.order_id = ?", orderId).Delete(&items)
	o.DB.Where("orders.id = ?", orderId).Delete(&order)
	response.Code = 200
	response.Status = "success"
	response.Data = "order success deleted"
	return response
}
