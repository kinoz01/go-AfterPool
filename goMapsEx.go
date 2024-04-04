package main

import "fmt"

// Units stores the Gross Store unit measurements.
func Units() map[string]int {
	return map[string]int{
		"quarter_of_a_dozen": 3,
		"half_of_a_dozen":    6,
		"dozen":              12,
		"small_gross":        120,
		"gross":              144,
		"great_gross":        1728,
	}
}

// NewBill creates a new bill.
func NewBill() map[string]int {
	return map[string]int{}
}

// AddItem adds an item to customer bill.
func AddItem(bill, units map[string]int, item, unit string) bool {
	c := bill[item]
	for key, value := range units {
		if key == unit {
			bill[item] = c + value
			return true
		}
	}
	return false
}

// RemoveItem removes an item from customer bill.
func RemoveItem(bill, units map[string]int, item, unit string) bool {
	qInBill, itemExist  := bill[item]
	qInUnits, unitExist := units[unit]
	if !itemExist || !unitExist || qInBill < qInUnits {
		return false
	} 
	if qInBill > qInUnits {
		bill[item] -= units[unit]
	} else  {
		delete(bill, item)
	}
	return true
}

// GetItem returns the quantity of an item that the customer has in his/her bill.
func GetItem(bill map[string]int, item string) (int, bool) {
	Item, ItemExist := bill[item]
	return Item, ItemExist
}

func main() {
    fmt.Println(Units())
	
	fmt.Println(NewBill())

	bill := NewBill()
	units := Units()
	ok := AddItem(bill, units, "carrot", "dozen")
	fmt.Println(ok)
	fmt.Println(bill)

	ok = RemoveItem(bill, units, "carrot", "dozen")
	fmt.Println(ok)
	fmt.Println(bill)

	bill = map[string]int{"carrot": 12, "grapes": 3}
	fmt.Println(GetItem(bill, "carrot"))
}

// Gross Store
