### map_usb_devices

Map each plugged devices with the relative com port opened, providing `idProduct:idVendor` and the location of device files `/dev`. The output is a `map[string]string` with the com port and the devices information achieved through `libusb`.  

## Required

Install the development version of libusb

`sudo apt-get install libusb-1.0-0-dev`

## How it works

Take a look here

`https://dev.to/enbis/how-udev-rules-can-help-us-to-recognize-a-usb-to-serial-device-over-dev-tty-interface-pbk`

Get the attached device 

``` bash
$dmesg
[10561.087277] usb 1-8: USB disconnect, device number 11
[10561.087711] ftdi_sio ttyUSB0: FTDI USB Serial Device converter now disconnected from ttyUSB0
[10561.087812] ftdi_sio 1-8:1.0: device disconnected
[10573.097500] usb 1-8: new full-speed USB device number 12 using xhci_hcd
[10573.251701] usb 1-8: New USB device found, idVendor=0403, idProduct=6001
[10573.251703] usb 1-8: New USB device strings: Mfr=1, Product=2, SerialNumber=3
[10573.251705] usb 1-8: Product: LABS Board1
[10573.251706] usb 1-8: Manufacturer: LABS
[10573.251706] usb 1-8: SerialNumber: A14EHLMI
[10573.254404] ftdi_sio 1-8:1.0: FTDI USB Serial Device converter detected
[10573.254436] usb 1-8: Detected FT232RL
[10573.254615] usb 1-8: FTDI USB Serial Device converter now attached to ttyUSB0
[11096.231736] perf: interrupt took too long (3266 > 3185), lowering kernel.perf_event_max_sample_rate to 6100
``` 

and a bunch of its information 

``` bash
$udevadm info /dev/ttyUSB0
P: /devices/pci0000:00/0000:00:14.0/usb1/1-8/1-8:1.0/ttyUSB0/tty/ttyUSB0
N: ttyUSB0
S: serial/by-id/usb-LABS_LABS_Board1_A14EHLMI-if00-port0
S: serial/by-path/pci-0000:00:14.0-usb-0:8:1.0-port0
E: DEVLINKS=/dev/serial/by-id/usb-LABS_LABS_Board1_A14EHLMI-if00-port0 /dev/serial/by-path/pci-0000:00:14.0-usb-0:8:1.0-port0
E: DEVNAME=/dev/ttyUSB0
E: DEVPATH=/devices/pci0000:00/0000:00:14.0/usb1/1-8/1-8:1.0/ttyUSB0/tty/ttyUSB0
E: ID_BUS=usb
E: ID_MM_CANDIDATE=1
E: ID_MODEL=LABS_Board1
E: ID_MODEL_ENC=LABS\x20Board1
E: ID_MODEL_FROM_DATABASE=FT232 Serial (UART) IC
E: ID_MODEL_ID=6001
E: ID_PATH=pci-0000:00:14.0-usb-0:8:1.0
E: ID_PATH_TAG=pci-0000_00_14_0-usb-0_8_1_0
E: ID_PCI_CLASS_FROM_DATABASE=Serial bus controller
E: ID_PCI_INTERFACE_FROM_DATABASE=XHCI
E: ID_PCI_SUBCLASS_FROM_DATABASE=USB controller
E: ID_REVISION=0600
E: ID_SERIAL=LABS_LABS_Board1_A14EHLMI
E: ID_SERIAL_SHORT=A14EHLMI
E: ID_TYPE=generic
E: ID_USB_DRIVER=ftdi_sio
E: ID_USB_INTERFACES=:ffffff:
E: ID_USB_INTERFACE_NUM=00
E: ID_VENDOR=LABS
E: ID_VENDOR_ENC=LABS
E: ID_VENDOR_FROM_DATABASE=Future Technology Devices International, Ltd
E: ID_VENDOR_ID=0403
E: MAJOR=188
E: MINOR=0
E: SUBSYSTEM=tty
E: TAGS=:snap_network-manager_networkmanager:systemd:
E: USEC_INITIALIZED=10572736847
```

Take a look at `ID_VENDOR_ID=0403` and `ID_MODEL_ID=6001`, these data will be using on the package to map the devices on the com port. `product=0x6001` and `vendor=0x0403`

for Linux environment:
* root: `/dev`
  
* dev: `ttyUSB`

```
GetUsbDeviceMap(root, dev string, product, vendor uint16)
```