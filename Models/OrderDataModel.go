package Models


type Order struct{
	id string `json:"id"`
	product_id string `json:"product_id"`
	quantity int `json:"quantity"`
	status string `json:"status"`
}

func (o *Order)TableName ()string{
	return "order"
}
