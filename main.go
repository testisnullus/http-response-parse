package main

import (
	"fmt"
	"json"
	"net/http"
)

type ParsedResponse struct {
	Respone string `json:"respone"`
}

func main() {
	parsedResponse := new(ParsedResponse)
	fmt.Println(getJson("https://sokyra.party/uk/material/907", parsedResponse))
}

func getJson(url string, target interface{}) error {
	r, err := http.Get(url)
	if err != nil {
		return err
	}
	defer r.Body.Close()

	return json.NewDecoder(r.Body).Decode(target)
}
