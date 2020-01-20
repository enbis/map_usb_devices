package map_usb_devices

import (
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
		prod, _ := dev.Product()
		res = append(res, prod)

		defer dev.Close()
	}

	return res, nil
}

func findDevices(product, vendor uint16) func(desc *gousb.DeviceDesc) bool {
	return func(desc *gousb.DeviceDesc) bool {
		return desc.Product == gousb.ID(product) && desc.Vendor == gousb.ID(vendor)
	}
}
