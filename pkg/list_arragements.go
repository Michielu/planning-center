package main

import (
	"fmt"
	"net/http"
)

type TestType interface {
}

func (pc PlanningCenter) listArragements() {
	fmt.Println("AppID", pc.AppID)
	fmt.Println("AppPassword", pc.AppPassword)

	fmt.Println("I AM RETURNING ARRAGEMENTS")

	apiUrl := fmt.Sprintf("%s/people/v2/people", BASE_URL)
	req, err := http.NewRequest(http.MethodGet, apiUrl, nil)

	// Should set account and authentication somewhere here
	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	fmt.Println("basicauth: ", basicAuth(pc.AppID, pc.AppPassword)) // This is the creds
	req.Header.Add("Authorization", "Basic "+basicAuth(pc.AppID, pc.AppPassword))
	// req.Header.Add("Authorization", "Basic "+pc.Credentials)

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
