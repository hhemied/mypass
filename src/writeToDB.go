package main

import (
	"log"
	"os"
)

func main() {
	db := "file"
	f, err := os.OpenFile(db, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	if _, err := f.WriteString("Hello\n"); err != nil {
		log.Fatal(err)
	}
}
