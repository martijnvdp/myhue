package functions

import (
	"encoding/base64"
	"fmt"
	"os"

	"github.com/amimof/huego"
	"github.com/spf13/viper"
)

//ConnectHUE to HUE and return TOKEN
func ConnectHUE() (token string, bridge *huego.Bridge) {
	user := viper.GetString("hueconfig.user")
	token64, err := base64.StdEncoding.DecodeString(viper.GetString("hueconfig.token"))
	token = string(token64)
	if err != nil {
		panic(err)
	}
	hostip := viper.GetString("hueconfig.ip")
	if token == "" {
		token = os.Getenv("HUE_TOKEN")
		user = os.Getenv("HUE_USER")
	}
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
			viper.SetDefault("hueconfig.token", base64.StdEncoding.EncodeToString([]byte(token)))
			viper.WriteConfig()
		}
	}
	bridge.Login(token)
	return token, bridge
}
