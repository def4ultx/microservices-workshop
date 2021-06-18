package inventory

import "testing"

func TestGetTotalAmount_ReturnZeroAmount(t *testing.T) {

	p := make([]Product, 0)
	actual := GetTotalAmount(p)
	expected := 0

	if actual != expected {
		t.Errorf("want: (%v), got: (%v)", expected, actual)
	}
}

func TestGetTotalAmount_Return100(t *testing.T) {

	products := make([]Product, 0)
	p := Product{
		Price:  25,
		Amount: 2,
	}
	products = append(products, p, p)

	actual := GetTotalAmount(products)
	expected := 100

	if actual != expected {
		t.Errorf("want: (%v), got: (%v)", expected, actual)
	}
}
