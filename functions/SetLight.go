package functions

import (
	"github.com/amimof/huego"
)

//GetAll get lights from hue bridge
func SetLight(n *int, a *bool, br *int, token string, b *huego.Bridge) (err string) {

	light, _ := b.GetLight(*n)
	if *a {
		light.On()
		brl := uint8(*br)
		light.Bri(brl)
	} else {
		light.Off()
	}
	return err
}
