package main

import (
	"fmt"

	"github.com/SergeyParamoshkin/Quickstart/version"
	"github.com/SergeyParamoshkin/rest/client"
)

func main() {
	// nolint
	d := client.Client{}
	d.Ping()
	fmt.Println("Hello Golang World", version.VERSION)
}
