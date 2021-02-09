package functions

import (
	"log"
	"strings"

	"github.com/amimof/huego"
)

//SetLight state on or off and brightness level
func SetLight(n *int, a *bool, br *int, token string, b *huego.Bridge) (err string) {

	light, err2 := b.GetLight(*n)
	if err2 != nil {
		if strings.Contains(err2.Error(), "unauthorized user") {
			log.Fatal("User not authorized, push button on the hue bridge before run")
		} else {
			panic(err2)
		}
	}
	if *a {
		light.On()
		brl := uint8(*br)
		light.Bri(brl)
	} else {
		light.Off()
	}
	return err
}
