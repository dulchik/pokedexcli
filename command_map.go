package main

import (
	"fmt"
	"net/http"
	"log"
	"io"
	"strconv"
)

func commandMap() error {
	offset := 0
	url := "https://pokeapi.co/api/v2/location-area?limit=20&offset=" + strconv.Itoa(offset)
	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	body, err := io.ReadAll(res.Body)
	if res.StatusCode > 299 {
		log.Fatalf("Rresponse failed with status code: %d and\nbody: %s\n", res.StatusCode, body)
	}
	if err != nil {
		log.Fatal(err)
	}

	defer res.Body.Close()

	

	fmt.Printf("%s\n", body)
	
	return nil
}
