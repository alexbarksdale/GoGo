package main

import "fmt"

func main() {
	// Callbacks
	ii := []int{1, 2, 3, 4, 5, 6}
	fmt.Println(sum(ii...))
	fmt.Println(even(sum, ii...))

}

// Callback Section
// xi = Slice of type Int
func sum(xi ...int) int {
	fmt.Printf("%T\n", xi)
	total := 0
	for _, v := range xi {
		total += v
	}
	return total
}

func even(f func(xi ...int) int, vi ...int) int {
	var yi []int

	for _, v := range vi {
		if v%2 == 0 {
			yi = append(yi, v)

		}
	}
	return f(yi...)
}

// End Callback
