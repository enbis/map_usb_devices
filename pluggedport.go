package map_usb_devices

import (
	"io/ioutil"
	"os"
	"strings"
)

func GetPluggedDevices(root, dev string) ([]os.FileInfo, error) {
	res := []os.FileInfo{}

	files, err := ioutil.ReadDir(root)
	if err != nil {
		return nil, err
	}

	for _, f := range files {
		if strings.Contains(f.Name(), dev) {
			res = append(res, f)
		}
	}

	return res, nil
}
