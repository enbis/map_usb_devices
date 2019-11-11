package map_usb_devices

import (
	"encoding/json"
	"errors"
	"sort"
	"strconv"
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

	sort.Slice(pPorts, func(i, j int) bool {
		return pPorts[i].ModTime().Before(pPorts[j].ModTime())
	})

	resMap := make(map[string]string, len(pPorts))

	pDevices, err := GetDevicesInfo(product, vendor)
	if err != nil {
		return nil, err
	}

	tempMap := make(map[int]string, len(pDevices))
	var orderKeys []int
	for _, dev := range pDevices {
		deviceInfo := new(DeviceInfo)
		err := json.Unmarshal([]byte(dev), deviceInfo)
		if err != nil {
			return nil, err
		}
		split := strings.Split(deviceInfo.Description, "=")
		if len(split) > 0 {
			kAddr, err := strconv.Atoi(split[len(split)-1])
			if err != nil {
				return nil, err
			}
			tempMap[kAddr] = dev
			orderKeys = append(orderKeys, kAddr)
		}
	}

	sort.Ints(orderKeys)

	if len(pPorts) != len(orderKeys) {
		return nil, errors.New("Error comparison ports and devices")
	}
	for i, k := range orderKeys {
		resMap[pPorts[i].Name()] = tempMap[k]
	}

	return resMap, nil
}
