package main

import (
	"fmt"
	"net/http"
)

type TestType interface {
}

func (pc PlanningCenter) listArragements() {
	fmt.Println("account", pc.Account)
	fmt.Println("auth", pc.Authentication)

	fmt.Println("I AM RETURNING ARRAGEMENTS")

	// Example
	// Source: https://www.makeuseof.com/go-make-http-requests/
	// https://reqres.in/
	apiUrl := "https://reqres.in/api/users"
	req, err := http.NewRequest(http.MethodGet, apiUrl, nil)

	// Should set account and authentication somewhere here
	req.Header.Set("Content-Type", "application/json; charset=utf-8")

	res, err := pc.Client.Do(req)
	if err != nil {
		fmt.Println("err: ", err.Error())
		return
	}

	if res.StatusCode != http.StatusOK {
		fmt.Println("unexpected status code: ", res.Status)
	}

	var test TestType

	getJson(*res, test)

	fmt.Println("test", test)

	// a := req.Response.Body
	// fmt.Println("Success request: ", a)

}
