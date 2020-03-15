package main

import (
	"github.com/brutella/hc"
	"github.com/brutella/hc/accessory"
)
/*
*/
import "C"

func main() {

	acc := accessory.NewSwitch(accessory.Info{Name:"Switch",Manufacturer:"Gucan",SerialNumber:"GCSWITCH",Model:"SI"})
	acc.Switch.On.OnValueRemoteUpdate(func(on bool) { if on == true { print("a1 on\n") } else { print("a1 off\n") } })

	config := hc.Config{Pin:"52037521",Port:"12345",StoragePath:"./db"}
	t, err := hc.NewIPTransport(config, acc.Accessory)
	if err != nil {
		print(err)
	}
	t.Start()
}
