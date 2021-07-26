package Models


type Product struct{
	id string `json:"id"`
	product_name string `json:"product_name"`
	price int `json:"price"`
	quantity int `json:"quantity"`
}

func (p *Product)TableName() string{
	return "product"
}
