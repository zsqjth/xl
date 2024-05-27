package pojo

type Product struct {
	ID          int     `json:"id"`
	ProductName string  `json:"name"`
	Img         string  `json:"img"`
	Price       float64 `json:"price"`
}
