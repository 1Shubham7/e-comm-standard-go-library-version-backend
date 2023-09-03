package data

import ( 
	"time"
	"encoding/json"
	"io"
)

type Product struct {
	// stuff encoded inside `` is called struct tags, they have a specific usecase
	ID          int 		`json:"id"` // now the ID is 'id' during the output, you can call it "hello" or anything
	Name        string 		`json:"name"`
	Description string 		`json:"description"`
	Price       float32		`json:"price"`
	SKU         string 		`json:"sku"`
	CreatedOn   string 		`json"-"` // t	his will not output, we will use it internally
	UpdatedOn   string 		`json"-"`
	DeletedOn   string 		`json"-"`
}

// instead of json.Marshal we want to use this method instead
type Products []*Product

func (p *Products) ToJSON(w io.Writer) error { //name of function is ToJSON and it returns an error
	encoder := json.NewEncoder(w)
	return encoder.Encode(p)
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
		CreatedOn:   time.Now().UTC().String(),
		UpdatedOn:   time.Now().UTC().String(),
	},
	&Product{
		ID:          2,
		Name:        "Rajma Chawal",
		Description: "Spicy",
		Price:       100,
		SKU:         "10",
		CreatedOn:   time.Now().UTC().String(),
		UpdatedOn:   time.Now().UTC().String(),
	},
	&Product{
		ID:          3,
		Name:        "Pastry",
		Description: "Sweet",
		Price:       50,
		SKU:         "100",
		CreatedOn:   time.Now().UTC().String(),
		UpdatedOn:   time.Now().UTC().String(),
	},
}

// productList is a slice of pointers to Product structs. 

// abstracting the products by a function
func GetProducts() Products {
	return productList
}
