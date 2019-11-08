package map_usb_devices

import (
	"fmt"
	"testing"
)

func TestGetUsbDeviceMap(t *testing.T) {
	res, err := GetUsbDeviceMap("/dev", "ttyUSB", 0x6001, 0x0403)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("map res %v \n", res)
}
