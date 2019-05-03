package main

import "C"

import (
	"log"

	"github.com/brutella/hc"
	"github.com/brutella/hc/accessory"
)

func main() {}

//export StartServer
func StartServer(pin string) {
	bridgeInfo := accessory.Info{
		Name:         "TheosBridge",
		Manufacturer: "tskippervold",
	}

	bridge := accessory.NewBridge(bridgeInfo)

	config := hc.Config{
		Pin: pin,
	}

	light := addLight().Accessory

	t, err := hc.NewIPTransport(config, bridge.Accessory, light)
	if err != nil {
		log.Fatal(err)
	}

	hc.OnTermination(func() {
		log.Println("On termination")
		<-t.Stop()
	})

	t.Start()
}

func addLight() *accessory.Lightbulb {
	info := accessory.Info{
		Name: "Light",
	}

	acc := accessory.NewLightbulb(info)

	acc.Lightbulb.On.OnValueRemoteUpdate(func(on bool) {
		log.Printf("Light on %s", on)
	})

	acc.Lightbulb.On.OnValueRemoteGet(func() bool {
		return false
	})

	acc.Lightbulb.Brightness.OnValueRemoteUpdate(func(val int) {
		log.Printf("Light brightness %d", val)
	})

	acc.OnIdentify(func() {
		log.Printf("Identify %s", info.Name)
	})

	return acc
}
