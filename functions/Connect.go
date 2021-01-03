package functions

import (
	"os"

	"github.com/amimof/huego"
)

//ConnectHUE to HUE and return TOKEN
func ConnectHUE() (token string, bridge *huego.Bridge) {
	token = os.Getenv("HUETOKEN")
	user := os.Getenv("HUEUSER")
	hostip := os.Getenv("HUEIP")

	if hostip == "" {
		bridge, _ = huego.Discover()
	} else {
		bridge = huego.New(hostip, user)
	}
	if token == "" {
		token, _ = bridge.CreateUser(user)
	}
	bridge.Login(token)
	return token, bridge
}
