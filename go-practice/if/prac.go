package main

import "fmt"

func main() {
	statePopulation := map[string]int{
		"california":   39250017,
		"texas":        27864916,
		"florida":      27593457,
		"new york":     19723456,
		"pennsylvania": 12622367,
		"Illinois":     12068965,
		"ohio":         11613456,
	}
	if pop, ok := statePopulation["florida"]; ok {
		fmt.Println(pop)
	}
}
