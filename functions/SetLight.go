package functions

import (
	"github.com/amimof/huego"
)

//GetAll get lights from hue bridge
func SetLight(n *int, a *bool, token string, b *huego.Bridge) (err string) {

	light, _ := b.GetLight(*n)
	if *a {
		light.On()
	} else {
		light.Off()
	}
	return err
}
