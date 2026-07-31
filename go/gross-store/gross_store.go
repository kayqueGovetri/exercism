package gross

var mapUnits = map[string]int{
	"quarter_of_a_dozen": 3,
	"half_of_a_dozen":    6,
	"dozen":              12,
	"small_gross":        120,
	"gross":              144,
	"great_gross":        1728,
}

// Units stores the Gross Store unit measurements.
func Units() map[string]int {
	return mapUnits
}

// NewBill creates a new bill.
func NewBill() map[string]int {
	bill := map[string]int{}
	return bill
}

// AddItem adds an item to customer bill.
func AddItem(bill, units map[string]int, item, unit string) bool {
	_, exists := units[unit]
	if !exists {
		return false
	}
	bill[item] = bill[item] + units[unit]
	return true
}

// RemoveItem removes an item from customer bill.
func RemoveItem(bill, units map[string]int, item, unit string) bool {
	_, exists := bill[item]
	if !exists {
		return false
	}
	_, exists_unit := units[unit]
	if !exists_unit {
		return false
	}
	new_quantity := bill[item] - units[unit]
	if new_quantity < 0 {
		return false
	}
	if new_quantity == 0 {
		delete(bill, item)
		return true
	}
	bill[item] = new_quantity
	return true
}

// GetItem returns the quantity of an item that the customer has in his/her bill.
func GetItem(bill map[string]int, item string) (int, bool) {
	value, exists := bill[item]
	return value, exists
}
