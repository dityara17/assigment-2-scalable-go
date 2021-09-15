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

	//o.DB.Create(&request)
	//o.DB.Save(&request)

	response.Status = "created"
	response.Code = 201
	response.Data = request
	return response
}
