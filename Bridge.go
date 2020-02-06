package main
import (
	"github.com/brutella/hc"
	"github.com/brutella/hc/accessory"
	//"time"
)
/*
#include<stdio.h>
#include<stdlib.h>
char CMD[40];
int init(int p) {
	char C[88];
	snprintf(C,88,"echo %d > /sys/class/gpio/export && echo out > /sys/class/gpio/gpio%d/direction",p,p);
	return system(C);
}
int set(int p,int v) {
	snprintf(CMD,40,"echo %d > /sys/class/gpio/gpio%d/value",v,p);
	return system(CMD);
}
int get(int p) {
	snprintf(CMD,40,"cat /sys/class/gpio/gpio%d/value|grep 1",p);
	return system(CMD);
}
*/
import "C"
func main() {
	PIN :=[]C.int{12,11,1,0,3,15,16,14,198,199,7,19,18,2,13,10}
	NAME:=[]string{"开关1","开关2","开关3","开关4","开关5","开关6","开关7","开关8","开关9","开关10","开关11","开关12","开关13","开关14","开关15","开关16"}
	for i:=0;i<len(PIN);i++ {
		C.init(PIN[i])
	}

	bridge := accessory.New(accessory.Info{Name:"Gucan-Bridge",Manufacturer:"Gucan",SerialNumber:"GCBRIDGE",Model:"BR"},accessory.TypeBridge)
	container := accessory.NewContainer()

	for i:=0;i<len(PIN);i++ {
		P:=PIN[i]
		N:=NAME[i]
		acc := accessory.NewSwitch(accessory.Info{Name:N,Manufacturer:"Gucan",SerialNumber:"GCSWITCH",Model:"A"})
		acc.Switch.On.OnValueRemoteUpdate(func(on bool) { if on == true { C.set(P,1) } else { C.set(P,0) } })
		container.AddAccessory(acc.Accessory)
	}

	config := hc.Config{Pin:"52037521",Port:"12345",StoragePath:"/usr/data"}
	t, err := hc.NewIPTransport(config, bridge, container.Accessories...)
	if err != nil {
		print(err)
	}

	//go func() {
	//	for y:=0;y<len(PIN);y++ {
	//		//on1 := !acc1.Switch.On.GetValue()
	//		//if on1 == true { print("acc1 is on\n") } else { print("acc1 is off\n") }
	//		acc1.Switch.On.SetValue(on1)
	//
	//		time.Sleep(1*time.Second)
	//	}
	//}()

	//hc.OnTermination(func() {
	//	t.Stop()
	//})

	t.Start()
}
