package lasagna

// TODO: define the 'PreparationTime()' function
func PreparationTime(layers []string, time int) int {
	if time != 0 {
		return len(layers) * time
	} else {
		return len(layers) * 2
	}
}

// TODO: define the 'Quantities()' function
func Quantities(layers []string) (int, float64) {
	var noodles int
	var sauce float64
	for i := 0; i < len(layers); i++ {
		if layers[i] == "sauce" {
			sauce += 0.2
		} else if layers[i] == "noodles" {
			noodles += 50
		}
	}
	return noodles, sauce
}

// TODO: define the 'AddSecretIngredient()' function
func AddSecretIngredient(friendList, myList []string) {
	myList[len(myList)-1] = friendList[len(friendList)-1]
}

// TODO: define the 'ScaleRecipe()' function
func ScaleRecipe(amounts []float64, portions int) []float64 {
	var scaledAmount []float64
	for i := 0; i < len(amounts); i++ {
		scaledAmount = append(scaledAmount, amounts[i]*(float64(portions)/2))
	}
	return scaledAmount
}
