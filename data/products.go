package data

import "time"

type Product struct {
	ID          int
	Name        string
	Description string
	Price       float32
	SKU         string
	CreatedOn   string
	UpdatedOn   string
	DeletedOn   string
}

// we just created a data structure - Product
// Let's now create a slice of Product

var productList = []*Product{
	&Product{
		ID:          1,
		Name:        "Chole Bhature",
		Description: "Tasty",
		Price:       80,
		SKU:         "1",
		CreatedOn:   time.Now().UTC.String(),
		UpdatedOn:   time.Now().UTC.String(),
	},
	&Product{
		ID:          2,
		Name:        "Rajma Chawal",
		Description: "Spicy",
		Price:       100,
		SKU:         "10",
		CreatedOn:   time.Now().UTC.String(),
		UpdatedOn:   time.Now().UTC.String(),
	},
	&Product{
		ID:          3,
		Name:        "Pastry",
		Description: "Sweet",
		Price:       50,
		SKU:         "100",
		CreatedOn:   time.Now().UTC.String(),
		UpdatedOn:   time.Now().UTC.String(),
	},
}

// productList is a slice of pointers to Product structs. 