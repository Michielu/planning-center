package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// Returns JSON struct of response body
func getJson(r http.Response, target interface{}) error {
	defer r.Body.Close()

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&target)
	if err != nil {
		fmt.Printf("%T\n%s\n%#v\n", err, err, err)

		return err
	}

	fmt.Println("target", target)
	return nil
}
