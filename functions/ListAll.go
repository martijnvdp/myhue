package functions

import (
	"fmt"

	"github.com/amimof/huego"
)

//ListAll lights
func ListAll(token string, b *huego.Bridge) {

	l, err := b.GetLights()
	if err != nil {
		panic(err)
	}
	for _, lights := range l {
		fmt.Println("[")
		fmt.Println("Name:        : ", lights.Name)
		fmt.Println("ID:          : ", lights.ID)
		fmt.Println("ModelID      : ", lights.ModelID)
		fmt.Println("On           : ", lights.State.On)
		fmt.Println("Bright level : ", lights.State.Bri)
		fmt.Println("Reachable    : ", lights.State.Reachable)
		fmt.Println("]")
	}
	fmt.Printf("Found %d lights", len(l))
}
