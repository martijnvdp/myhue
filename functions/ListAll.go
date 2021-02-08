package functions

import (
	"fmt"
	"log"
	"strings"

	"github.com/amimof/huego"
)

//ListAll lights
func ListAll(token string, b *huego.Bridge) {

	l, err := b.GetLights()
	if err != nil {
		if strings.Contains(err.Error(), "unauthorized user") {
			log.Fatal("User not authorized, push button on the hue bridge before run")
		} else {
			panic(err)
		}
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
