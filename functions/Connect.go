package functions

import (
	"fmt"
	"os"

	"github.com/amimof/huego"
	"github.com/spf13/viper"
)

//ConnectHUE to HUE and return TOKEN
func ConnectHUE() (token string, bridge *huego.Bridge) {
	token = os.Getenv("HUETOKEN")
	user := os.Getenv("HUEUSER")
	hostip := os.Getenv("HUEIP")
	if user == "" {
		fmt.Print("No HUEUSER env set, enter username: ")
		fmt.Scanln(&user)
	}
	if hostip == "" {
		bridge, _ = huego.Discover()
		bconfig, err := bridge.GetConfig()
		if err == nil {
			hostip = bconfig.IPAddress
			fmt.Println(hostip)
		}
	} else {
		bridge = huego.New(hostip, user)
	}
	if token == "" {
		token, _ = bridge.CreateUser(user)
		fmt.Println("token for user: ", user)
		fmt.Println("token: ", token)
	}
	bridge.Login(token)
	viper.SetDefault("token", token)
	viper.SetDefault("user", user)
	viper.SafeWriteConfig()
	return token, bridge
}
