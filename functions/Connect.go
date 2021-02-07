package functions

import (
	"bufio"
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
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("No HUEUSER env set, enter username: ")
		user, _ = reader.ReadString('\n')
	}
	if hostip == "" {
		bridge, _ = huego.Discover()
	} else {
		bridge = huego.New(hostip, user)
	}
	if token == "" {
		token, _ = bridge.CreateUser(user)
		viper.SetDefault("token", token)
		viper.SetDefault("user", user)
		viper.SafeWriteConfig()
	}
	bridge.Login(token)
	return token, bridge
}
