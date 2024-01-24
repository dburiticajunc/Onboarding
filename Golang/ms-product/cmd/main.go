package main

import (
	"fmt"
	"ms-product/internal/application"
	"os"
)

func main() {

	cfg := &application.ConfigAppDefault{
		ServerAddr: os.Getenv("SERVER_ADDR"),
		DbFile:     os.Getenv("DB_FILE"),
	}

	app := application.NewApplicationDefault(cfg)

	// - setup
	err := app.SetUp()
	if err != nil {
		fmt.Println(err)
		return
	}

	// -run
	err = app.Run()
	if err != nil {
		fmt.Println(err)
		return
	}
}
