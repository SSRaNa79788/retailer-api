package Customer

import "time"

type Customer struct{
	Id string `json:"id"`
	LastOrderTime time.Time `json:"lastordertime"`
}

func (c *Customer)TableName ()string{
	return "customer"
}
