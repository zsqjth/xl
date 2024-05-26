package dao

import (
	"xl/pojo"
	"xl/utils"
)

func Select(username string, password string) (pojo.User, error) {
	var user pojo.User
	sql := "select id,username,password from user where username = ? and password = ?"
	row, err := utils.Db.Query(sql, username, password)
	if err != nil {
		return user, err
	}
	defer row.Close()
	if row.Next() {
		row.Scan(&user.ID, &user.UserName, &user.PassWord)
	}
	return user, nil
}

func Add(username string, password string) error {
	sql := "insert into user(username,password) values(?,?)"
	_, err := utils.Db.Exec(sql, username, password)
	if err != nil {
		return err
	}
	return nil
}
