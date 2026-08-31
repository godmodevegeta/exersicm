package gross

// Units stores the Gross Store unit measurements.
func Units() map[string]int {
	gross_map := map[string]int{}
	gross_map["quarter_of_a_dozen"] = 3
	gross_map["half_of_a_dozen"] = 6
	gross_map["dozen"] = 12
	gross_map["small_gross"] = 120
	gross_map["gross"] = 144
	gross_map["great_gross"] = 1728

	return gross_map
}

// NewBill creates a new bill.
func NewBill() map[string]int {
	return make(map[string]int)
}

// AddItem adds an item to customer bill.
func AddItem(bill, units map[string]int, item, unit string) bool {
	_, exists := units[unit]
	if !exists {
		return false
	}
	// adding item to bill map
	
	// if item already exist in bill
	value, exists := bill[item]
	if !exists {
		bill[item] = units[unit]
	}
	bill[item] = units[unit] + value
	return true
}

// RemoveItem removes an item from customer bill.
func RemoveItem(bill, units map[string]int, item, unit string) bool {

	unit_value, exists := units[unit]
	if !exists {
		return false
	}

	item_value, exists := bill[item]
	if !exists {
		return false
	}

	newQuantity := item_value - unit_value
	if newQuantity <  0 {
		return false	
	} else if newQuantity == 0 {
		delete(bill, item)
		return true
	} 

	bill[item] = newQuantity
	return true
}

// GetItem returns the quantity of an item that the customer has in his/her bill.
func GetItem(bill map[string]int, item string) (int, bool) {
	item_value, exists := bill[item]
	if !exists {
		return 0, false
	}
	return item_value, true
}
