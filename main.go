package main

import (
	"fmt"

	"github.com/plloi/go-http-api-wrapper-template/api"
	"github.com/ufoscout/go-up"
)

func main() {
	config, err := go_up.NewGoUp().
		AddFile("./.env", true).
		Build()

	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s", err))
	}

	api := cylance.NewClient(nil)
	_ = config.GetString("string-setting")
	api.GetItems()
}
