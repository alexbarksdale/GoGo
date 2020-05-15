package main

import "fmt"

func main() {
	// 1. 'var' usage to create a map
	// var colors map[string]string

	// 2. 'make' usage to create map
	// colors := make(map[string]string)

	// 3. literal way to create a map
	//            key    value
	colors := map[string]string{
		"red":   "#ff0000",
		"green": "#4bf745",
	}
	// delete(colors, "red")

	printMap(colors)
}

func printMap(c map[string]string) {
	for key, val := range c {
		fmt.Println("Hex code for", key, "is", val)
	}
}
