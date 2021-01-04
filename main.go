package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/martijnxd/myhue/functions"
)

func main() {
	l := flag.Int("light", 0, "an int")
	list := flag.Bool("list", false, "list all lights")
	a := flag.Bool("on", false, "a bool")
	flag.Parse()
	token, bridge := functions.ConnectHUE()
	if os.Getenv("HUETOKEN") == "" {
		fmt.Println("set ENV HUETOKEN=", token)
	}
	if *list {
		functions.ListAll(token, bridge)
	}
	if *l != 0 {
		functions.SetLight(l, a, token, bridge)
	}
}
