package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/martijnxd/myhue/functions"
)

func main() {
	l := flag.Int("light", 1, "an int")
	a := flag.Bool("on", false, "a bool")
	flag.Parse()
	token, bridge := functions.ConnectHUE()
	if os.Getenv("HUETOKEN") == "" {
		fmt.Println("set ENV HUETOKEN=", token)
	}
	functions.ListAll(token, bridge)

	functions.SetLight(l, a, token, bridge)
}
