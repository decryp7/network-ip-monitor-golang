package main

import (
	"fmt"
	"time"
)

type networkMonitor struct {
	networkMonitorStop    *chan bool
	networkMonitorStopped *chan bool
}

func (n *networkMonitor) stop() {
	*n.networkMonitorStop <- true
	<-*n.networkMonitorStopped
}

var logr *logger

func (n *networkMonitor) start() {
	l := logger{}
	logr = &l

	m, err := getMainNetworkInterface()

	if err != nil {
		logr.writeError(fmt.Sprintf("Unable to get main network interface: %v\n", err))
	}

	logr.writeLog(fmt.Sprintf("Main network interface: %v", m))

	ticker := time.NewTicker(1 * time.Second)
	//make a channel for notification to stop
	s := make(chan bool)
	n.networkMonitorStop = &s
	ns := make(chan bool)
	n.networkMonitorStopped = &ns

	//go routine. hell yeah!
	//loop forever until stop
	go func() {
		for {
			select {
			case <-*n.networkMonitorStop:
				//cleanup
				logr.writeLog("Stopping network ip monitoring...")
				*n.networkMonitorStopped <- true
				logr.writeLog("Network ip monitoring stopped.")
				return
			case <-ticker.C:
				checkAndUpdateMainNetworkInterfaceIPAddress(&m)
			}
		}
	}()
	logr.writeLog("Network ip monitoring started!")
}

func checkAndUpdateMainNetworkInterfaceIPAddress(m *networkInterface) {
	currentIpAddress, err := getIPAddress(m.name)

	if err != nil {
		fmt.Printf("Unable to check and update ip address for interface: %v %v\n", m.name, err.Error())
		return
	}

	fmt.Printf("Existing IP: %v  Current IP: %v\n", currentIpAddress, m.ipAddress)

	if currentIpAddress != m.ipAddress {
		message := fmt.Sprintf("IP address has changed from %v to %v", m.ipAddress, currentIpAddress)
		logr.writeLog(message)
		sendEmailUsingOutlook("IP Address Changed!", message)
		m.ipAddress = currentIpAddress
	}
}
