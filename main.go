package main

import (
	"fmt"
	"os"

	"github.com/martijnxd/myhue/functions"
)

func main() {
	token, bridge := functions.ConnectHUE()
	if os.Getenv("HUETOKEN") == "" {
		fmt.Println("set ENV HUETOKEN=", token)
	}
	functions.SetLight(1, true, token, bridge)
	functions.ListAll(token, bridge)
}
