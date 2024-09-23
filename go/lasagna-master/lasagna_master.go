package lasagna

func PreparationTime(layers []string, time int) int {
	if time == 0 {
		time = 2
	}
	return time * len(layers)
}

func Quantities(layers []string) (int, float64) {
	n, s := 0, 0
	for _, layer := range layers {
		if layer == "noodles" {
			n++
		} else if layer == "sauce" {
			s++
		}
	}
	return n * 50, float64(s) * 0.2
}

func AddSecretIngredient(a, b []string) {
	m := make(map[string]int)
	for _, v := range a {
		m[v]++
	}
	for _, v := range b {
		m[v]--
	}
	for k, v := range m {
		if v > 0 && k != "béchamel" {
			b[len(b)-1] = k
			break
		}
	}
}

func ScaleRecipe(s []float64, n int) []float64 {
	scaleFactor := float64(n) / 2
	newS := make([]float64, len(s))
	copy(newS, s)
	for i := range newS {
		newS[i] *= scaleFactor
	}
	return newS
}
