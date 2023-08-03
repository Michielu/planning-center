package main

import "net/http"

type REQUEST string

const (
	APP_ID  = "APP_ID"
	APP_PWD = "PASSWORD"
)

const (
	LIST_ARRANGEMENT REQUEST = "list-arrangement"
	OTHER            REQUEST = "other"
)

type PlanningCenter struct {
	AppID       string
	AppPassword string
	Client      *http.Client
}
