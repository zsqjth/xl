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

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusNoContent)
		return
	}
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
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusNoContent)
		return
	}
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
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusNoContent)
		return
	}
	var product pojo.Product
	name := r.URL.Query().Get("name")
	fmt.Println(name)
	product, err := service.Find(name)
	fmt.Println(product)
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
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusNoContent)
		return
	}
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
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusNoContent)
		return
	}
	var product pojo.Product
	json.NewDecoder(r.Body).Decode(&product)
	fmt.Println(product)
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
