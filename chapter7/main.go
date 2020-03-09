package main

import (
	"fmt"
)

func main() {
	// Arrays - more recommended to use Slices
	fmt.Println("\nArrays:")
	var x [5]int
	fmt.Println(x)
	fmt.Println(len(x))
	x[3] = 42
	fmt.Println(x, "\n")

	// Slices - Allows you to grp together the values of the same type
	// SLICES ARE DYNAMIC
	fmt.Println("\nSlicing:")
	// y := []type{values} // composite literal
	y := []int{4, 5, 6, 7, 8}
	fmt.Println(y)

	for i, v := range y {
		fmt.Println(i, v)
	}

	// Slicing a Slice
	fmt.Println("\nSlice a Slice:")
	a := []int{4, 5, 6, 7, 8}
	fmt.Println(a)
	fmt.Println(a[0:3])

	for i := 0; i < len(a); i += 1 {
		fmt.Println(i, a[i])
	}

	// Range over a slice
	xi := []int{4, 5, 6, 7, 9}
	for key, value := range xi {
		fmt.Println(key, value)
	}

	// Append to a slice
	fmt.Println("\nAppend:")
	b := []int{1, 2, 3, 4, 5}
	fmt.Println(b)
	b = append(b, 10, 20, 30, 40)
	fmt.Println(b)
	// Spread operator
	fmt.Println("\nSpread:")
	c := []int{50, 60, 70, 80}
	b = append(b, c...)
	fmt.Println(b)
	// Remove from an array
	fmt.Println("\nRemove item:")
	b = append(b[:2], b[4:]...)
	fmt.Println(b)
	// Make - Instead of recreating a slice for appending, Make() creates
	// a slice of a certain size and cap and doesn't create it over and over.
	fmt.Println("\nMake:")
	d := make([]int, 10, 50)
	fmt.Println(d, len(d), cap(d))
	d = append(d, 50)
	fmt.Println(d, len(d), cap(d))
	// Multi dimen slice
	fmt.Println("\nMulti Dimen Slice:")
	jb := []string{"James", "Bond", "Chocolate", "Martini"}
	fmt.Println(jb)
	mp := []string{"Miss", "Moneypenny", "Strawberry", "Hazelnut"}
	fmt.Println(mp)

	xp := [][]string{jb, mp}
	fmt.Println(xp)

	// Maps
	fmt.Println("\nMaps:")
	m := map[string]int{
		// Must have trailing comma
		"james": 32,
		"bob":   25,
	}
	fmt.Println(m)
	fmt.Println(m["james"])
	fmt.Println(m["Noname"])
	v, ok := m["Noname"]
	fmt.Println(v)
	fmt.Println(ok)

	if v, ok := m["Noname"]; ok {
		fmt.Println("exists", v)
	}

	if v, ok := m["james"]; ok {
		fmt.Println("exists", v)
	}
	// Add to a Map
	m["alex"] = 33
	for key, value := range m {
		fmt.Println(key, value)
	}

	// Delete a key from an item
	delete(m, "bob")
	fmt.Println("Deleted 'bob'", m)
	// check if it exists and then delete
	if v, ok := m["alex"]; ok { // Comma Okay Idiom
		fmt.Println(m)
		fmt.Println("Exists, now deleting", v)
		delete(m, "alex")
		fmt.Println(m)
	}

}
