package inventory

type Product struct {
	ID     int64  `json:"id"`
	Name   string `json:"name"`
	Price  int    `json:"price"`
	Amount int    `json:"amount"`
}

func GetTotalAmount(products []Product) int {
	var sum int
	for _, v := range products {
		sum += v.Amount * v.Price
	}
	return sum
}
