package functions

import (
	"fmt"
	"log"
	"strings"

	"github.com/amimof/huego"
)

type lights struct {
	NAME      string `mapstructure:"NAME"`
	ID        int    `mapstructure:"ID"`
	ModelID   string `mapstructure:"ModelID"`
	On        bool   `mapstructure:"On"`
	Bright    uint8  `mapstructure:"Bright"`
	Reachable bool   `mapstructure:"Reachable"`
}

//ListAll lights
func ListAll(r, w bool, f, token string, b *huego.Bridge) {
	var list []lights
	l, err := b.GetLights()
	if err != nil {
		if strings.Contains(err.Error(), "unauthorized user") {
			log.Fatal("User not authorized, push button on the hue bridge before run")
		} else {
			panic(err)
		}
	}
	for _, ll := range l {
		if r && ll.State.Reachable {
			if f != "" && strings.Contains(ll.Name, f) {
				list = append(list, lights{
					ID: ll.ID, NAME: ll.Name, ModelID: ll.ModelID, On: ll.State.On, Bright: ll.State.Bri, Reachable: ll.State.Reachable,
				})
			}
			if f == "" {
				list = append(list, lights{
					ID: ll.ID, NAME: ll.Name, ModelID: ll.ModelID, On: ll.State.On, Bright: ll.State.Bri, Reachable: ll.State.Reachable,
				})
			}
		}
		if !r && f == "" {
			list = append(list, lights{
				ID: ll.ID, NAME: ll.Name, ModelID: ll.ModelID, On: ll.State.On, Bright: ll.State.Bri, Reachable: ll.State.Reachable,
			})
		}
		if !r && f != "" && strings.Contains(ll.Name, f) {
			list = append(list, lights{
				ID: ll.ID, NAME: ll.Name, ModelID: ll.ModelID, On: ll.State.On, Bright: ll.State.Bri, Reachable: ll.State.Reachable,
			})
		}

	}
	if w {
		fmt.Println("Name - ID - Model - State - Reachable")
	}
	for _, ll := range list {
		if w {
			fmt.Println(ll.NAME, " | ", ll.ID, " | ", ll.ModelID, " | ", ll.On, " | ", ll.Reachable)
		} else {
			fmt.Println("Name      : ", ll.NAME)
			fmt.Println("ID        : ", ll.ID)
			fmt.Println("Model     : ", ll.ModelID)
			fmt.Println("State     : ", ll.On)
			fmt.Println("Reachable : ", ll.Reachable)
			fmt.Println("-----")
		}
	}
	fmt.Printf("Found %d lights", len(list))
	fmt.Println("")
}
