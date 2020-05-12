package main

import (
	// "errors"
	"fmt"
	"log"
)

func main() {
	_, err := sqrt(-10)
	if err != nil {
		log.Fatalln(err)
	}
}

func sqrt(f float64) (float64, error) {
	if f < 0 {
		//     float64, error
		// return 0, errors.New("Norgate math: square root of negative number")
		return 0, fmt.Errorf("Norgate math: square root of negative number: %v", f)
	}
	return 42, nil
}
