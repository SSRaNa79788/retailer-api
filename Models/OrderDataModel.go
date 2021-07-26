package Models


type Order struct{
	id string `json:"id"`
	produc_name string `json:"product_name"`
	price int `json:"price"`
	quantity int `json:"quantity"`
}

func (o *Order)TableName ()string{
	return "order"
}
