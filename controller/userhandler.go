package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"xl/dao"
	"xl/pojo"
	"xl/service"
	"xl/utils"
)

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusNoContent)
		return
	}
	var user pojo.User
	err := json.NewDecoder(r.Body).Decode(&user)
	fmt.Println(user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	st, err := service.Login(user.UserName, user.PassWord)
	fmt.Println(st)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if st {
		cookie := &http.Cookie{
			Name:  "user_auth",
			Value: fmt.Sprintf("%s:%s", user.UserName, user.PassWord),
			Path:  "/",
		}
		http.SetCookie(w, cookie)
		utils.RespondWithJSON(w, 0, "success", "登录成功")
		return
	} else {
		utils.RespondWithJSON(w, 1, "false", "登录失败")
	}
}

func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusNoContent)
		return
	}
	var user pojo.User
	err := json.NewDecoder(r.Body).Decode(&user)
	fmt.Println(user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err1 := dao.Add(user.UserName, user.PassWord)
	if err1 != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	utils.RespondWithJSON(w, 0, "success", "注册成功")

}
