package Order


type Order struct{
	Id string `json:"id"`
	Cust_id string `json:"cust_id"`
	Product_id string `json:"product_id"`
	Quantity int `json:"quantity"`
	Status string `json:"status"`
}

func (o *Order)TableName ()string{
	return "order"
}
