package map_usb_devices

import (
	"io/ioutil"
	"strings"
)

func GetPluggedDevices(root, dev string) ([]string, error) {
	res := []string{}

	files, err := ioutil.ReadDir(root)
	if err != nil {
		return nil, err
	}

	for _, f := range files {

		if strings.Contains(f.Name(), dev) {
			res = append(res, f.Name())
		}
	}

	return res, nil
}
