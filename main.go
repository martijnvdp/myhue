package main

import "github.com/martijnxd/myhue/functions"

func main() {
	token, bridge := functions.ConnectHUE()
	functions.SetLight(1, true, token, bridge)
	functions.ListAll(token, bridge)
}
