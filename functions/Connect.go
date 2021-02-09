package functions

import (
	"fmt"

	"github.com/amimof/huego"
	"github.com/spf13/viper"
)

//ConnectHUE to HUE and return TOKEN
func ConnectHUE() (token string, bridge *huego.Bridge) {
	user := viper.GetString("hueconfig.user")
	token = viper.GetString("hueconfig.token")
	hostip := viper.GetString("hueconfig.ip")
	if user == "" {
		fmt.Print("No HUEUSER env set, enter username: ")
		fmt.Scanln(&user)
		if user != "" {
			viper.SetDefault("hueconfig.user", user)
			viper.WriteConfig()
		}
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
		if token != "" {
			viper.SetDefault("hueconfig.token", token)
			viper.WriteConfig()
		}
	}
	bridge.Login(token)
	return token, bridge
}
