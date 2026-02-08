package models

type Product struct {
	ID 	int     `json:"id"`
	Name 	string  `json:"name"`
	Price 	float64 `json:"price"`
	Stock 	int     `json:"stock"`
	CategoryID int `json:"category_id"`
	Active  string	`json:"active"`
	CategoryName string `json:"category_name"`
}

type ProdukFilter struct{
	Name string
	Active string
}
