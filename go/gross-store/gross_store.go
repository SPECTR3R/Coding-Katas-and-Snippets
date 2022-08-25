package gross

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
	return make(map[string]int)
}

// AddItem adds an item to customer bill.
func AddItem(bill, units map[string]int, item, unit string) bool {
	uValue, uExists := units[unit]

	if !uExists {
		return false
	}

	bill[item] += uValue

	return true
}

// RemoveItem removes an item from customer bill.
func RemoveItem(bill, units map[string]int, item, unit string) bool {
	uValue, uExists := units[unit]
	bValue, bExists := bill[item]

	newValue := bValue - uValue

	if !uExists || !bExists || newValue < 0 {
		return false
	}

	if newValue == 0 {
		delete(bill, item)
	} else {
		bill[item] = newValue
	}

	return true
}

// GetItem returns the quantity of an item that the customer has in his/her bill.
func GetItem(bill map[string]int, item string) (int, bool) {
	bValue, bExists := bill[item]
	return bValue, bExists
}
