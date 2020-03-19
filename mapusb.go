package map_usb_devices

import (
	"errors"
	"sort"
	"strings"
)

type DeviceInfo struct {
	Manufacturer string `json:"manufacturer"`
	Product      string `json:"product"`
	SerialNumber string `json:"serialNumber`
	Description  string `json:"description"`
}

func GetUsbDeviceMap(root, dev string, product, vendor uint16) (map[string]string, error) {

	pPorts, err := GetPluggedDevices(root, dev)
	if err != nil {
		return nil, err
	}

	sort.Strings(pPorts)

	resMap := make(map[string]string, len(pPorts))

	pDevices, err := GetDevicesInfo(product, vendor)
	if err != nil {
		return nil, err
	}

	sort.Strings(pDevices)

	if len(pPorts) != len(pDevices) {
		return nil, errors.New("Error comparison ports and devices")
	}

	for i, d := range pDevices {

		if !strings.Contains(pPorts[i], d) {
			return nil, errors.New("Error between ports and devices: values not aligned, check rules.d file")
		}
		resMap[pPorts[i]] = d
	}

	return resMap, nil
}
