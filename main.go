package main

import (
	// "encoding/json"
	"flag"
	"fmt"
	"go.evanpurkhiser.com/netgear"
	"os"
	"time"
)

var (
	host     = flag.String("host", "192.168.1.1", "Your netgear router address")
	username = flag.String("username", os.Getenv("RT_USERNAME"), "Your netgear router username")
	password = flag.String("password", os.Getenv("RT_PASSWORD"), "Your netgear router password")
)

var output = map[netgear.DeviceChange]string{
	netgear.DeviceAdded:   "Device Added",
	netgear.DeviceRemoved: "Device Removed",
}

func listener(change *netgear.ChangedDevice, err error) {
	if err != nil {
		fmt.Printf("Failed to query for devices: %s\n", err)
		return
	}
	InsertDeviceConnectionLog(change)
	fmt.Printf(output[change.Change]+": %s\n", change.Device.MAC)
}

func main() {
	flag.Parse()

	client := netgear.NewClient(*host, *username, *password)

	pollTime := time.Second * 10
	client.OnDeviceChanged(pollTime, listener)

	<-make(chan bool)
}
