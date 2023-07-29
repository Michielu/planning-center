package main

import "net/http"

type REQUEST string

const (
	PC_ACCOUNT = "PLANNING_CENTER_ACCOUNT"
	PC_AUTH    = "PLANNING_CENTER_AUTH"
)

const (
	LIST_ARRANGEMENT REQUEST = "list-arrangement"
	OTHER            REQUEST = "other"
)

type PlanningCenter struct {
	Account        string
	Authentication string
	Client         *http.Client
}
