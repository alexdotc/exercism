package lasagna

func PreparationTime(layers []string, minutes int) int {
	if minutes == 0 { minutes = 2 }
	return len(layers) * minutes
}

func Quantities(layers []string) (noodles int, sauce float64) {
	for _,v := range layers {
		if v == "noodles" {
			noodles += 50
		} else if v == "sauce" {
			sauce += 0.2
		}
	}
	return noodles, sauce
}

func AddSecretIngredient(friend []string, mine []string) {
	mine[len(mine)-1] = friend[len(friend)-1]
}

func ScaleRecipe(quantities []float64, portions int) []float64 {
	scaledQuantities := make([]float64, len(quantities))
	for i,v := range quantities {
		scaledQuantities[i] = v*(float64(portions)/2.0)
	}
	return scaledQuantities
}
