package map_usb_devices

import (
	"errors"
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func GetUsbDeviceMap(root, dev string, product, vendor uint16) (map[string]string, error) {

	pPorts, err := GetPluggedDevices(root, dev)
	if err != nil {
		return nil, err
	}

	fmt.Printf("pre %v", pPorts)
	sort.Slice(pPorts, func(i, j int) bool {
		return pPorts[i].ModTime().Before(pPorts[j].ModTime())
	})
	fmt.Printf("post %v", pPorts)

	resMap := make(map[string]string, len(pPorts))

	pDevices, err := GetDevicesInfo(product, vendor)
	if err != nil {
		return nil, err
	}

	tempMap := make(map[int]string, len(pDevices))
	var orderKeys []int
	for _, dev := range pDevices {
		iAddr := strings.LastIndex(dev, "addr=")
		if iAddr != -1 {
			split := strings.Split(dev[iAddr:], "=")
			kAddr, _ := strconv.Atoi(split[1])
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
