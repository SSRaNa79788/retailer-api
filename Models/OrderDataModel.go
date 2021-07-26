package Models


type Order struct{
	id string `json:"id"`
	cust_id string `json:"cust_id"`
	product_id string `json:"product_id"`
	quantity int `json:"quantity"`
	status string `json:"status"`
}

func (o *Order)TableName ()string{
	return "order"
}
