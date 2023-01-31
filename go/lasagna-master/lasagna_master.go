package lasagna

// TODO: define the 'PreparationTime()' function
func PreparationTime(layers []string, layerPrepTime int) int {
	if layerPrepTime <= 0 {
		layerPrepTime = 2
	}

	return len(layers) * layerPrepTime
}

// TODO: define the 'Quantities()' function
func Quantities(layers []string) (noodles int, sauce float64) {
	for _, v := range layers {
		if v == "noodles" {
			noodles += 50
		} else if v == "sauce" {
			sauce += 0.2
		}
	}

	return noodles, sauce
}

// TODO: define the 'AddSecretIngredient()' function
func AddSecretIngredient(friendList []string, ingredientList []string) {
	ingredientList[len(ingredientList)-1] = friendList[len(friendList)-1]
}

// TODO: define the 'ScaleRecipe()'
func ScaleRecipe(amounts []float64, portions int) []float64 {
	scaled := make([]float64, len(amounts))

	for i, v := range amounts {
		scaled[i] = (v * float64(portions)) / 2
	}

	return scaled
}
