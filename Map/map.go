package main

import (
	"fmt"
)

func main() {
	colors := map[string]string{
		"red":   "#ff0000",
		"white": "#ffffff",
		"black": "#000000",
	}
	fmt.Println(colors)
	printMap(colors)

	newColors := make(map[int]string)
	newColors[10] = "#4bf745"
	fmt.Println(newColors)

	delete(newColors, 10)
	fmt.Println(newColors)

}

func printMap(c map[string]string) {
	for key, value := range c {
		fmt.Println(key, value)
	}
}
