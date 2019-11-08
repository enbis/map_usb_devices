package map_usb_devices

import (
	"fmt"

	"github.com/google/gousb"
)

func GetDevicesInfo(product, vendor uint16) ([]string, error) {

	res := []string{}
	ctx := gousb.NewContext()
	defer ctx.Close()

	devices, err := ctx.OpenDevices(findDevices(product, vendor))
	if err != nil {
		return nil, err
	}

	for _, dev := range devices {
		manuf, _ := dev.Manufacturer()
		prod, _ := dev.Product()
		sn, _ := dev.SerialNumber()
		res = append(res, fmt.Sprintf("Manufacturer@%s; Product@%s; SerialNumber@%s; Description@%s", manuf, prod, sn, dev.String()))

		defer dev.Close()
	}

	return res, nil
}

func findDevices(product, vendor uint16) func(desc *gousb.DeviceDesc) bool {
	return func(desc *gousb.DeviceDesc) bool {
		return desc.Product == gousb.ID(product) && desc.Vendor == gousb.ID(vendor)
	}
}
