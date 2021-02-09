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
		switch {
		case strings.Contains(err2.Error(), "ERROR 3"):
			log.Fatal("Error, light not available")
		case strings.Contains(err2.Error(), "unauthorized user"):
			log.Fatal("User not authorized, push button on the hue bridge before run")
		default:
			panic(err2)
		}
	}
	if *a || (*br != 0) {
		light.On()
		brl := uint8(*br)
		light.Bri(brl)
	} else {
		light.Off()
	}
	return err
}
