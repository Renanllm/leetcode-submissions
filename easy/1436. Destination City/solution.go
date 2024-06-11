package main

func DestCity(paths [][]string) string {
	leftCities := map[string]string{}
	rightCities := map[string]string{}
	for _, paths := range paths {
		a, b := paths[0], paths[1]
		leftCities[a] = ""
		rightCities[b] = ""
	}
	var destiny string
	for city := range rightCities {
		if _, ok := leftCities[city]; !ok {
			destiny = city
		}
	}
	return destiny
}
