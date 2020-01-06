package main
import (
	"github.com/brutella/hc"
	"github.com/brutella/hc/accessory"
)
func main() {
	bridge := accessory.New(accessory.Info{Name:"Bridge",Manufacturer:"Gucan",SerialNumber:"GCBRIDGE",Model:"BR"},accessory.TypeBridge)
	c := accessory.NewContainer()

	acc1 := accessory.NewSwitch(accessory.Info{Name:"开关1",Manufacturer:"Gucan",SerialNumber:"GCSWITCH1",Model:"A"})
	acc1.Switch.On.OnValueRemoteUpdate(func(on bool) { if on == true { print("a1 on\n") } else { print("a1 off\n") } })
	c.AddAccessory(acc1.Accessory)
	acc2 := accessory.NewSwitch(accessory.Info{Name:"开关2",Manufacturer:"Gucan",SerialNumber:"GCSWITCH2",Model:"B"})
	acc2.Switch.On.OnValueRemoteUpdate(func(on bool) { if on == true { print("a2 on\n") } else { print("a2 off\n") } })
	c.AddAccessory(acc2.Accessory)
	acc3 := accessory.NewSwitch(accessory.Info{Name:"开关3",Manufacturer:"Gucan",SerialNumber:"GCSWITCH3",Model:"C"})
	acc3.Switch.On.OnValueRemoteUpdate(func(on bool) { if on == true { print("a3 on\n") } else { print("a3 off\n") } })
	c.AddAccessory(acc3.Accessory)

	config := hc.Config{Pin:"52037521",Port:"12345",StoragePath:"./db"}
	t, err := hc.NewIPTransport(config, bridge, c.Accessories...)
	if err != nil {
		print(err)
	}
	hc.OnTermination(func() {
		t.Stop()
	})
	t.Start()
}
