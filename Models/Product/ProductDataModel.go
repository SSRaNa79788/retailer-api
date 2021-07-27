package Product


type Product struct{
	Id           string `json:"id"`
	Product_name string `json:"product_name"`
	Price        int    `json:"price"`
	Quantity     int    `json:"quantity"`
}

func (p *Product)TableName() string{
	return "product"
}
