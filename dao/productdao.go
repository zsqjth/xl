package dao

import (
	"xl/pojo"
	"xl/utils"
)

func SelectAll() ([]pojo.Product, error) {

	var products []pojo.Product
	sql := "select id,productname,img,price from product "
	row, err := utils.Db.Query(sql)
	if err != nil {
		return nil, err
	}
	defer row.Close()
	for row.Next() {
		var product pojo.Product
		err := row.Scan(&product.ID, &product.ProductName, &product.Img, &product.Price)
		if err != nil {
			return nil, err
		}
		products = append(products, product)
	}
	return products, nil

}

func SelectProduct(productname string) (pojo.Product, error) {
	var product pojo.Product
	sql := "select id,productname,img,price from product where productname = ?"
	row, err := utils.Db.Query(sql, productname)
	if err != nil {
		return product, err
	}
	defer row.Close()
	if row.Next() {
		row.Scan(&product.ID, &product.ProductName, &product.Img, &product.Price)
	}
	return product, nil

}
func AddProduct(productname string, img string, price float64) error {
	sql := "insert into product (productname,img,price) values (?,?,?)"
	_, err := utils.Db.Exec(sql, productname, img, price)
	if err != nil {
		return err
	}
	return nil

}

func Delete(productname string) error {
	sql := "delete from product where productname = ?"
	_, err := utils.Db.Exec(sql, productname)
	if err != nil {
		return err
	}
	return nil
}
func Update(productname string, img string, price float64) error {
	sql := "update product set productname=?, img=?, price=? where productname = ?"
	_, err := utils.Db.Exec(sql, productname, img, price, productname)
	if err != nil {
		return err
	}
	return nil
}
