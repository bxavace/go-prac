// JSON Activity - Mimicking Customer Order
package main

import (
	"fmt"
	"encoding/json"
)

type address struct {
	Street		string `json:"street"`
	City		string `json:"city"`
	ZipCode		int `json:"zipcode"`
}

type item struct {
	Name		string `json:"itemname"`
	Description	string `json:"desc,omitempty"`
	Quantity	int `json:"qty"`
	Price		int `json:"price"`
}

type order struct {
	TotalPrice	int `json:"totalprice"`
	IsPaid		bool `json:"paid"`
	Fragile		bool `json:"fragile,omitempty"`
	OrderDetail	[]item `json:"orderdetail"`
}

type customer struct {
	UserName	string `json:"username"`
	Password	string `json:"-"`
	Token		string `json:"-"`
	ShipTo		address `json:"shipto"`
	PurchaseOrder	order `json:"order"`
}

func main() {
	jsonData := []byte(`
	{
		"username": "blackhat",
		"shipto": {
			"street": "Sulphur Springs Rd",
			"city": "Park City",
			"state": "VA",
			"zipcode": 12345
		},
		"order": {
			"paid": false,
			"orderdetail": [{
				"itemname": "Lordaeron and Archimonde",
				"desc": "book",
				"qty": 3,
				"price": 50
			},
			{
				"itemname": "Archimonde's Ascent",
				"desc": "cd",
				"qty": 1,
				"price": 65
			},
			{
				"itemname": "Network+ study guide",
				"desc": "book",
				"qty": 2,
				"price": 300
			}]
		}
	}
	`)

	var c customer

	if !json.Valid(jsonData) {
		fmt.Printf("JSON not valid: %s", jsonData)
	}

	err := json.Unmarshal(jsonData, &c)
	if err != nil {
		fmt.Println(err)
	}
	game := item{Name: "Final Fantasy", Description: "game", Quantity: 1, Price: 79}
	glass := item{Name: "Wine Glass", Quantity: 10, Price: 50}
	c.PurchaseOrder.OrderDetail = append(c.PurchaseOrder.OrderDetail, game)
	c.PurchaseOrder.OrderDetail = append(c.PurchaseOrder.OrderDetail, glass)

	c.Total()
	c.PurchaseOrder.IsPaid = true
	c.PurchaseOrder.Fragile = true

	customerOrder, err := json.MarshalIndent(c, "", "    ")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(customerOrder))
}

func (c *customer) Total() {
	price := 0
	for _, item := range c.PurchaseOrder.OrderDetail {
		price += item.Quantity * item.Price
	}

	c.PurchaseOrder.TotalPrice = price
}

