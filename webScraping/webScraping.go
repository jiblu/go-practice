package main

import (
	// "io"
	"log"
	"net/http"
	// "os"
)

func main() {
	response, err := http.Get("https://www.devdungeon.com")
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()

}