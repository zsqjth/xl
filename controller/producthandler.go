package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"xl/pojo"
	"xl/service"
	"xl/utils"
)

func InsertHandler(w http.ResponseWriter, r *http.Request) {
	var product pojo.Product
	json.NewDecoder(r.Body).Decode(&product)
	fmt.Println(product)
	st, err := service.Insert(product.ProductName, product.Img, product.Price)
	if err != nil {
		return
	}
	if st {
		utils.RespondWithJSON(w, 0, "添加成功", nil)
	} else {
		utils.RespondWithJSON(w, 1, "添加失败", nil)
	}

}
func ShowAllHandler(w http.ResponseWriter, r *http.Request) {
	var products []pojo.Product
	products, err := service.ShowAll()
	if err != nil {
		return
	}
	if products != nil {
		utils.RespondWithJSON(w, 0, "success", products)
	} else {
		utils.RespondWithJSON(w, 1, "false", nil)
	}

}
func ShowHandler(w http.ResponseWriter, r *http.Request) {
	var product pojo.Product
	name := r.URL.Query().Get("name")
	product, err := service.Find(name)
	if err != nil {
		return
	}
	if product.ProductName != "" {
		utils.RespondWithJSON(w, 0, "success", product)
	} else {
		utils.RespondWithJSON(w, 1, "false", nil)
	}

}
func UpdateHandler(w http.ResponseWriter, r *http.Request) {
	var product pojo.Product
	json.NewDecoder(r.Body).Decode(&product)
	st, err := service.Update(product.ProductName, product.Img, product.Price)
	if err != nil {
		return
	}
	if st {
		utils.RespondWithJSON(w, 0, "success", nil)
	} else {
		utils.RespondWithJSON(w, 1, "false", nil)
	}
}

func DestoryHandler(w http.ResponseWriter, r *http.Request) {
	var product pojo.Product
	json.NewDecoder(r.Body).Decode(&product)
	st, err := service.Destory(product.ProductName)
	if err != nil {
		fmt.Println("有错误")
		return
	}
	if st {
		utils.RespondWithJSON(w, 0, "success", nil)
	} else {
		utils.RespondWithJSON(w, 1, "false", nil)
	}

}
