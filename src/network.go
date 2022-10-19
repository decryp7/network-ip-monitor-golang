package main

import (
	"fmt"
	"net"
	"regexp"
)

// match 10.*.*.*
var networkIPAddressRegex = regexp.MustCompile(`^10\..*$`)

// match IP4 ip address
var ip4AddressRegex = regexp.MustCompile(`^(?:[0-9]{1,3}\.){3}[0-9]{1,3}$`)

type networkInterface struct {
	name      string
	ipAddress string
}

func getMainNetworkInterface() (n networkInterface, err error) {
	ifaces, err := net.Interfaces()

	if err != nil {
		logr.writeError(fmt.Sprintf("error occurred during get interfaces: %v", err.Error()))
		return
	}

	for _, i := range ifaces {
		addrs, err := i.Addrs()

		if err != nil {
			logr.writeError(fmt.Sprintf("error occured during getting addresses of interface(%v): %v", i.Name, err.Error()))
		}

		for _, a := range addrs {
			switch v := a.(type) {
			// case *net.IPAddr:
			// 	fmt.Printf("%v : %s (%s)\n", i.Name, v, v.IP.DefaultMask())
			case *net.IPNet:
				//fmt.Printf("%v : %v %v\n", i.Name, v.IP, i.Flags.String())
				if v.IP != nil && networkIPAddressRegex.MatchString(v.IP.String()) {
					return networkInterface{
						name:      i.Name,
						ipAddress: v.IP.String()}, nil
				}
			}
		}
	}

	return networkInterface{}, fmt.Errorf("unable to get main network interface based on regex: %v", networkIPAddressRegex)
}

func getIPAddress(interfaceName string) (ipAddress string, err error) {
	i, err := net.InterfaceByName(interfaceName)

	if err != nil {
		return "", fmt.Errorf("unable to get ip address of network interface: %v %v", interfaceName, err.Error())
	}

	addrs, err := i.Addrs()

	if err != nil {
		fmt.Println(fmt.Errorf("error occured during getting addresses of interface(%v): %v", i.Name, err.Error()))
	}

	for _, a := range addrs {
		switch v := a.(type) {
		case *net.IPNet:
			if v.IP != nil && ip4AddressRegex.MatchString(v.IP.String()) {
				return v.IP.String(), nil
			}
		}
	}

	return "", fmt.Errorf("unable to get ip address of network interface: %v", interfaceName)
}
