package functions

import (
	"fmt"

	"github.com/amimof/huego"
)

//ListAll lights
func ListAll(token string, b *huego.Bridge) {

	bridge := huego.New(b.Host, b.User)
	l, err := bridge.GetLights()

	if err != nil {
		panic(err)
	}
	fmt.Printf("Found %d lights", len(l))
	for _, lights := range l {
		fmt.Println("%l", lights.Name)
		fmt.Println("%l", lights.State)
	}
}
