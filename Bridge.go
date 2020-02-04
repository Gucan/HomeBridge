package main
import (
	"github.com/brutella/hc"
	"github.com/brutella/hc/accessory"
	"time"
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
*/
import "C"
func main() {
	PIN :=[]C.int{148,149,150,151,152,153,154,155}
	NAME:=[]string{"插孔1","插孔2","插孔3","插孔4","插孔5","插孔6","插孔7","插孔8"}
	for i:=0;i<len(PIN);i++ {
		C.init(PIN[i])
	}

	bridge := accessory.New(accessory.Info{Name:"Gucan-Bridge",Manufacturer:"Gucan",SerialNumber:"GCBRIDGE",Model:"BR"},accessory.TypeBridge)
	container := accessory.NewContainer()

	for i:=0;i<len(PIN);i++ {
		P:=PIN[i]
		N:=NAME[i]
		acc := accessory.NewSwitch(accessory.Info{Name:N,Manufacturer:"Gucan",SerialNumber:"GCSWITCH",Model:"A"})
		acc.Switch.On.OnValueRemoteUpdate(func(on bool) { if on == true { C.set(P,1) } else { print(P,0) } })
		container.AddAccessory(acc.Accessory)
	}

	config := hc.Config{Pin:"52037521",Port:"12345",StoragePath:"./db"}
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

	hc.OnTermination(func() {
		t.Stop()
	})

	t.Start()
}
