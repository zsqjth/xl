package service

import (
	"fmt"
	"xl/dao"
)

func Login(username, password string) (bool, error) {
	user, err := dao.Select(username, password)
	if err != nil {
		return false, err
	}
	fmt.Println(user)
	if user.UserName == "" {
		return false, nil
	} else {
		return true, nil
	}
}

func Register(username, password string) error {
	err := dao.Add(username, password)
	if err != nil {
		return err
	} else {
		return nil
	}
}
