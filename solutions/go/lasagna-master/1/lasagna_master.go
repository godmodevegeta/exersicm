package lasagnamaster

// TODO: define the 'PreparationTime()' function
func PreparationTime(layers []string, time_per_layer int) int {
	length_of_layer := len(layers)
	if time_per_layer == 0 {
		return length_of_layer * 2;
	}
	return length_of_layer * time_per_layer
}

// TODO: define the 'Quantities()' function
func Quantities(layers []string) (int, float64) {
	noodle := 0
	sauce := 0.0
	for i := 0; i < len(layers); i++ {
		layer := layers[i] 
		if layer == "noodles" {
			noodle = noodle + 50
		} else if layer == "sauce" {
			sauce = sauce + 0.2
		}
	}
	return noodle, sauce
}

// TODO: define the 'AddSecretIngredient()' function
func AddSecretIngredient(friendsList, myList []string) {
	friends_last_item := friendsList[len(friendsList) - 1]
	myList[len(myList) - 1] = friends_last_item
}

// TODO: define the 'ScaleRecipe()' function
func ScaleRecipe(quantities []float64, no_of_portions int) []float64 {
	scaler := float64(no_of_portions)/2.0
	scaled_quantities := make([]float64, len(quantities))
	for i := 0; i < len(quantities); i++ {
		scaled_quantities[i] = quantities[i] * scaler
	}
	return scaled_quantities
}
// Your first steps could be to read through the tasks, and create
// these functions with their correct parameter lists and return types.
// The function body only needs to contain `panic("")`.
//
// This will make the tests compile, but they will fail.
// You can then implement the function logic one by one and see
// an increasing number of tests passing as you implement more
// functionality.
