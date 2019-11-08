package map_usb_devices

import (
	"fmt"
	"os"
	"testing"
)

var testCases = []struct {
	description string
	input1      string
	input2      string
	want        interface{}
}{
	{
		description: "linux ttyUSB",
		input1:      "/dev",
		input2:      "ttyUSB",
		want:        2,
	},
}

var implementations = []struct {
	descr string
	f     func(r, d string) ([]os.FileInfo, error)
}{
	{
		descr: "PluggedPort",
		f:     GetPluggedDevices,
	},
}

func TestGetPluggedDevices(t *testing.T) {
	for _, impl := range implementations {
		t.Run(impl.descr, func(t *testing.T) {
			for _, tc := range testCases {
				t.Run(tc.description, func(t *testing.T) {
					got, _ := impl.f(tc.input1, tc.input2)
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
