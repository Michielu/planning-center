package main

import (
	"errors"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/joho/godotenv"
)

func main() {
	pc, err := setupPC()
	if err != nil {
		fmt.Println("error loading env", err.Error())
		os.Exit(1)
	}

	command := REQUEST(os.Args[1])

	switch command {
	case LIST_ARRANGEMENT:
		pc.listArragements()
		break
	case OTHER:
		fmt.Println("Something else")
		break
	default:
		fmt.Println("Unknown command:", command)
		os.Exit(1)
	}
}

func setupPC() (*PlanningCenter, error) {
	err := godotenv.Load(".env")
	if err != nil {
		return nil, err
	}

	account := os.Getenv(PC_ACCOUNT)
	auth := os.Getenv(PC_AUTH)

	if account == "" || auth == "" {
		return nil, errors.New(".env file not setup correctly")
	}
	return &PlanningCenter{
		Account:        account,
		Authentication: auth,
		Client:         &http.Client{Timeout: 30 * time.Second},
	}, nil
}
