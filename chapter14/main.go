package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func main() {
	// Error handling
	// Ex 1
	var ans1 string

	fmt.Print("Name: ")
	_, err := fmt.Scan(&ans1)
	if err != nil {
		panic(err)
	}
	fmt.Println(ans1)

	// Ex 2
	f, err := os.Create("names.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()

	r := strings.NewReader("Wassup")
	fmt.Println("Created file")
	io.Copy(f, r)

	// Ex: 3
	_, err = os.Open("no-file.txt")
	if err != nil {
		fmt.Println("Error:", err)
		log.Println("Error:", err)
		log.Panicln("Panic error:", err)
		log.Fatalln("Fatal Error:", err)
		panic(err)
	}
}
