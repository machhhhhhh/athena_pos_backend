package utils

import (
	"net"

	"github.com/gin-gonic/gin"
	"github.com/gofiber/fiber/v2"
)

func GetIPAdressGIN(context *gin.Context) string {

	master_ip := context.ClientIP()

	ifaces, err := net.Interfaces()
	if err != nil || len(ifaces) == 0 {
		return master_ip
	}
	for _, iface := range ifaces {
		if iface.Flags&net.FlagUp == 0 {
			continue // interface down
		}
		if iface.Flags&net.FlagLoopback != 0 {
			continue // loopback interface
		}
		addrs, err := iface.Addrs()
		if err != nil {
			return master_ip
		}
		for _, addr := range addrs {
			var ip net.IP
			switch v := addr.(type) {
			case *net.IPNet:
				ip = v.IP
			case *net.IPAddr:
				ip = v.IP
			}
			if ip == nil || ip.IsLoopback() {
				continue
			}
			ip = ip.To4()
			if ip == nil {
				continue // not an ipv4 address
			}
			return master_ip + " | " + ip.String()
		}
	}
	return master_ip
}

func GetIPAdressFiber(context *fiber.Ctx) string {

	master_ip := context.IP()

	ifaces, err := net.Interfaces()
	if err != nil || len(ifaces) == 0 {
		return master_ip
	}
	for _, iface := range ifaces {
		if iface.Flags&net.FlagUp == 0 {
			continue // interface down
		}
		if iface.Flags&net.FlagLoopback != 0 {
			continue // loopback interface
		}
		addrs, err := iface.Addrs()
		if err != nil {
			return master_ip
		}
		for _, addr := range addrs {
			var ip net.IP
			switch v := addr.(type) {
			case *net.IPNet:
				ip = v.IP
			case *net.IPAddr:
				ip = v.IP
			}
			if ip == nil || ip.IsLoopback() {
				continue
			}
			ip = ip.To4()
			if ip == nil {
				continue // not an ipv4 address
			}
			return master_ip + " | " + ip.String()
		}
	}
	return master_ip
}
