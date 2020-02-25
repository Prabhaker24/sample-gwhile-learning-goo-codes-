package main

import "fmt"

func main() {
	colors := map[string]string{
		"red":  "#13212",
		"blue": "#32131",
		"grey": "#54345",
	}
	colors["x"] = "y"
	for i, tempmap := range colors {
		fmt.Println(i, tempmap)
	}

	fmt.Println(colors)
}
