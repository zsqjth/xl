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

func LogoutHandler(w http.ResponseWriter, r *http.Request) {
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
		utils.RespondWithJSON(w, "登录成功")
		return
	} else {
		utils.RespondWithJSON(w, "登录失败")
	}
}

func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	var user pojo.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err1 := dao.Add(user.UserName, user.PassWord)
	if err1 != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	utils.RespondWithJSON(w, "注册成功")

}
