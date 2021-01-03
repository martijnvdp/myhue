package functions

import "github.com/amimof/huego"

//AuthUser authorize user
func AuthUser(b *huego.Bridge) {
	//bridge, _ := huego.Discover()
	user, _ := b.CreateUser(b.User) // Link button needs to be pressed
	b = b.Login(user)
	//light, _ := b.GetLight(3)
	//light.Off()
}
