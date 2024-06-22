package service_usb

import (
	"fmt"
	"log"

	"github.com/google/gousb"
)

func SetupUSB() {

	fmt.Println("Start Scanning Devices.")

	err := GoogleUSB()
	if err != nil {
		fmt.Println("USB Devices Error:", err)
		panic(err)
	}

	fmt.Println("Finish Scanning Devices.")

}

func GoogleUSB() error {
	usb_context := gousb.NewContext()
	defer usb_context.Close()

	// List all USB devices connected
	devices, err := usb_context.OpenDevices(func(desc *gousb.DeviceDesc) bool {
		// This function filters the devices. You can filter based on VID/PID, etc.
		// For now, returning true means we want to include all devices.
		return true
	})

	if err != nil {
		fmt.Println("Scanning Error:", err.Error())
		return err
	}

	if len(devices) == 0 {
		fmt.Println("No Devices founded.")
		return nil
	}

	for i := range devices {
		current_device := devices[i]
		defer current_device.Close()

		seial, err := current_device.SerialNumber()
		if err != nil {
			log.Printf("Failed to open serial :" + err.Error())
			return err
		}

		product, err := current_device.Product()
		if err != nil {
			log.Printf("Failed to open product :" + err.Error())
			return err
		}
		manu_fac, err := current_device.Manufacturer()
		if err != nil {
			log.Printf("Failed to open manufacturer :" + err.Error())
			return err
		}

		current_device_info := "Device: " + seial + " " + product + " " + manu_fac
		fmt.Println(current_device_info, "had connected.")

		// Claim an interface
		// Assuming interface 0 and endpoint 1 for the demonstration
		intf, done, err := current_device.DefaultInterface()
		if err != nil {
			log.Printf(current_device_info, "Error claiming default interface:", err.Error())
			continue
		}
		defer done()

		// Find the appropriate endpoints
		var endpointIn, endpointOut *gousb.EndpointDesc
		for _, endpoint := range intf.Setting.Endpoints {
			if endpoint.Direction == gousb.EndpointDirectionIn {
				endpointIn = &endpoint
			} else {
				endpointOut = &endpoint
			}
		}

		if endpointIn == nil || endpointOut == nil {
			log.Printf(current_device_info, "Could not find suitable endpoints on interface.")
			continue
		}

		fmt.Println("  endpoint_in", endpointIn)
		fmt.Println("  endpoint_out", endpointOut)

	}
	return nil
}
