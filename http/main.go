package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

type logWriter struct{}

func main() {
	res, err := http.Get("https://google.com")

	if err != nil {
		log.Fatalf("ERROR: %v", err)
	}

	lw := logWriter{}

	io.Copy(lw, res.Body)
}

// logWriter now implements the Writer interface
func (logWriter) Write(bs []byte) (int, error) {
	fmt.Println(string(bs))
	fmt.Println("Bytes written", len(bs))
	return len(bs), nil
}
