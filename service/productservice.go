package service

import (
	"fmt"
	"xl/dao"
	"xl/pojo"
)

func ShowAll() ([]pojo.Product, error) {
	products, err := dao.SelectAll()
	if err != nil {
		return nil, err
	}
	return products, nil
}
func Find(name string) (pojo.Product, error) {
	product, err := dao.SelectProduct(name)
	if err != nil {
		return pojo.Product{}, err
	}
	return product, nil
}
func Destory(name string) (bool, error) {

	err := dao.Delete(name)
	if err != nil {
		fmt.Println("错误了")
		return false, err
	}
	return true, nil
}
func Insert(name string, img string, price float64) (bool, error) {
	err := dao.AddProduct(name, img, price)
	if err != nil {
		return false, err
	}
	return true, nil
}

func Update(name string, img string, price float64) (bool, error) {
	err := dao.Update(name, img, price)
	if err != nil {
		return false, err
	}
	return true, nil
}
