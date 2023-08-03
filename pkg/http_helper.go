package main

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
)

const BASE_URL = "https://api.planningcenteronline.com"

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

func basicAuth(username, password string) string {
	auth := username + ":" + password
	return base64.StdEncoding.EncodeToString([]byte(auth))
}
