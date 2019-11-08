package map_usb_devices

import (
	"fmt"
	"testing"
)

var testCasesDevices = []struct {
	description string
	product     uint16
	vendor      uint16
	want        interface{}
}{
	{
		description: "devices info",
		product:     0x6001,
		vendor:      0x0403,
		want:        2,
	},
}

var implementationsDevices = []struct {
	descr string
	f     func(p, v uint16) ([]string, error)
}{
	{
		descr: "DevicesInfo",
		f:     GetDevicesInfo,
	},
}

func TestGetDevicesInfo(t *testing.T) {
	for _, impl := range implementationsDevices {
		t.Run(impl.descr, func(t *testing.T) {
			for _, tc := range testCasesDevices {
				t.Run(tc.description, func(t *testing.T) {
					got, _ := impl.f(tc.product, tc.vendor)
					if len(got) != tc.want {
						t.Fatalf("Test result %v \ngot: %d\nwant: %d", got, len(got), tc.want)
					} else {
						fmt.Printf("Test result %v", got)
					}
				})
			}
		})
	}
}
