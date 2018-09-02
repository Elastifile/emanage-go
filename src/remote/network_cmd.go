package remote

import (
	"fmt"
	"net"
	"regexp"
	"strings"
	"time"
	"unicode"

	"github.com/go-errors/errors"

	"github.com/elastifile/emanage-go/src/helputils"
	terrors "github.com/elastifile/emanage-go/src/tools/errors"
)

const (
	defaultTimeout    = 30 * time.Second
	vHeadVlanPattern  = "v%v."
	loaderVlanPattern = ".%v"
)

// GetNetworkInterfaces returns all the available network interface names
func (rem *Remote) GetNetworkInterfaces() (netDevices []string, err error) {
	var output string

	// Get list of interfaces
	cmd := "ls -1 /sys/class/net"
	if output, err = rem.Run(cmd); err != nil {
		return []string{}, errors.WrapPrefix(err, fmt.Sprintf("Failed getting list of interfaces on host %v. "+
			"Output: %v", rem.Address(), output), 0)
	}

	netDevices = strings.Split(strings.TrimSpace(output), "\n")
	return
}

// GetNetworkInterfaceByVlan returns the network interface name by VLAN
func (rem *Remote) GetNetworkInterfaceByVlan(vlan int) (netDevice string, err error) {
	netDevices, err := rem.GetNetworkInterfaces()
	if err != nil {
		return
	}

	// Get matching VLAN interface
	for _, dev := range netDevices {
		if strings.Contains(dev, fmt.Sprintf(vHeadVlanPattern, vlan)) {
			netDevice = dev
			break
		} else if strings.Contains(dev, fmt.Sprintf(loaderVlanPattern, vlan)) {
			netDevice = dev
			break
		}
	}

	if netDevice == "" {
		err = errors.Errorf("Failed to find net device for VLAN %v on host %v. Available devices: %v",
			vlan, rem.Address(), netDevices)
	} else {
		logger.Debug(fmt.Sprintf("Translating VLAN %v to net device %v on host %v", vlan, netDevice, rem.Address()))
	}

	return
}

// GetNetworkInterfaceByMac returns the network interface name by MAC address
func (rem *Remote) GetNetworkInterfaceByMac(mac string) (string, error) {
	if mac == "" {
		return "", errors.New("Unable to comply. Received an empty MAC address from caller.")
	}
	nic, err := rem.Run(fmt.Sprintf("ip a | awk -v MAC=%v 'BEGIN {IGNORECASE = 1} /BROADCAST/ {nic=$2} "+
		"/link/ {mac=$2} {if (mac==MAC) {print nic;exit}}' | awk -F: '{print $1}'", mac))
	if err != nil {
		logger.Error("Failed to get NIC name using MAC", "err", err, "mac", mac, "cmd output", nic)
		return nic, err
	}
	nic = strings.TrimSpace(nic)
	if nic == "" {
		return "", errors.Errorf("Got empty NIC name by MAC address %v", mac)
	}
	logger.Info("Detected NIC name", "nic", nic)
	return nic, nil
}

// GetInterfaceLinkStatusByDev returns link status of network interface, specified by device name
func (rem *Remote) GetInterfaceLinkStatusByDev(deviceName string) (status string, err error) {
	cmd := fmt.Sprintf("ip a | awk -e '/ %v[@:]/ {print $9}'", deviceName)
	status, err = rem.Run(cmd)
	status = strings.TrimSpace(status)
	if err != nil {
		err = errors.WrapPrefix(err, fmt.Sprintf("Failed to get network interface status for %v. Output: %v",
			deviceName, status), 0)
	}
	return
}

// GetInterfaceLinkStatusByVlan returns link status of network interface, specified by VLAN
func (rem *Remote) GetInterfaceLinkStatusByVlan(vlan int) (status string, err error) {
	cmd := fmt.Sprintf("ip a | awk -e '/: v%v\\./ {print $9}'", vlan)
	status, err = rem.Run(cmd)
	status = strings.TrimSpace(status)
	if err != nil {
		err = errors.WrapPrefix(err, fmt.Sprintf("Failed to get network interface status for VLAN %v. Output: %v",
			vlan, status), 0)
	}
	return
}

// GetInterfaceLinkStatusByMac returns link status of network interface, specified by MAC address
func (rem *Remote) GetInterfaceLinkStatusByMac(mac string) (string, error) {
	status, err := rem.Run(fmt.Sprintf("ip a | awk -v MAC=%v 'BEGIN {IGNORECASE = 1} /BROADCAST/ {status=$9} "+
		"/link/ {mac=$2}  {if (mac==MAC) {print status;exit}}'", mac))
	status = strings.TrimSpace(status)
	if err != nil {
		logger.Error("Failed to get the nic status", "ERROR", err, "Mac", mac, "output", status)
		return "", err
	}
	return status, nil
}

// GetAllNetworkInterfaceLinkStatuses returns network interface link statuses for all devices
func (rem *Remote) GetAllNetworkInterfaceLinkStatuses() (statuses map[string]string, err error) {
	statuses = make(map[string]string)
	deviceNames, err := rem.GetNetworkInterfaces()
	if err != nil {
		return
	}
	for _, deviceName := range deviceNames {
		statuses[deviceName], err = rem.GetInterfaceLinkStatusByDev(deviceName)
		if err != nil {
			return nil, errors.WrapPrefix(err, fmt.Sprintf("Failed getting link status for NIC %v", deviceName), 0)
		}
	}
	return
}

// logAllNetworkInterfaceLinkStatuses logs all network interface link statuses for debugging purposes
func (rem *Remote) logAllNetworkInterfaceLinkStatuses() (err error) {
	statuses, err := rem.GetAllNetworkInterfaceLinkStatuses()
	if err != nil {
		return
	}
	logger.Debug(fmt.Sprintf("Interface statuses for %v: %v", rem.Address(), statuses))
	return
}

// ipLinkSet sets link state to the desired state on network interface, specified by device name
func (rem *Remote) ipLinkSet(netDevice string, status string) error {
	cmd := fmt.Sprintf("ip link set %v %v", netDevice, status)
	output, err := rem.Run(cmd)
	if err != nil {
		logger.Debug("Failed to set interface link status", "nodeAddress", rem.Address(),
			"netDevice", netDevice, "status", status, "cmd", cmd, "output", output, "err", err)
	}
	rem.logAllNetworkInterfaceLinkStatuses()
	return err
}

// GetIpAddrByDev returns the host's IP address assigned to the specified device
func (rem *Remote) GetIpAddrByDev(deviceName string) (ipAddr string, err error) {
	// Get device's CIDR
	cmd := fmt.Sprintf("ip -o -f inet address show dev %v | awk '{printf $4}'", deviceName)
	output, err := rem.Run(cmd)
	if err != nil {
		logger.Debug("Failed to get device's CIDR", "host address", rem.Address(), "deviceName", deviceName,
			"cmd", cmd, "output", output, "err", err)
		return
	}

	// Extract IP address from CIDR
	ipBytes, _, err := net.ParseCIDR(output)
	if err != nil {
		logger.Debug("Failed to parse CIDR", "host address", rem.Address(), "deviceName", deviceName,
			"cmd", cmd, "output", output, "err", err)
		return
	}
	ipAddr = ipBytes.String()

	return
}

// GetIpAddrByDev returns the host's IP address on the specified VLAN
func (rem *Remote) GetIpAddrByVlan(vlan int) (ipAddr string, err error) {
	device, err := rem.GetNetworkInterfaceByVlan(vlan)
	if err != nil {
		return "", errors.WrapPrefix(err, fmt.Sprintf("Failed to get net device for VLAN %v on host %v",
			vlan, rem.Address()), 0)
	}
	return rem.GetIpAddrByDev(device)
}

// GetIpAddrByMac returns the IP address associated with the specified MAC address
func (rem *Remote) GetIpAddrByMac(mac string) (ipAddr string, err error) {
	if mac == "" {
		return "", errors.New("Unable to comply. Received an empty MAC address from caller.")
	}
	device, err := rem.GetNetworkInterfaceByMac(mac)
	if err != nil {
		return "", errors.WrapPrefix(err, fmt.Sprintf("Failed to get net device by MAC %v on host %v",
			mac, rem.Address()), 0)
	}
	if device == "" {
		return "", errors.Errorf("Got empty device name by MAC address %v", mac)
	}
	ipAddr, err = rem.GetIpAddrByDev(device)
	if err != nil {
		return "", errors.WrapPrefix(err, fmt.Sprintf("Failed to get device %v IP address on host %v",
			device, rem.Address()), 0)
	}
	return
}

// GetArpCachedMacForIpAddr returns the MAC address cached by ARP for the specified IP address
func (rem *Remote) GetArpCachedMacForIpAddr(destAddr string) (mac string, err error) {
	cmd := fmt.Sprintf("ip neighbour show to %v | awk '{printf $5}'", destAddr)
	mac, err = rem.Run(cmd)
	if err != nil {
		logger.Debug("Failed to get MAC address from ARP cache", "host address", rem.Address(),
			"destAddr", destAddr, "cmd", cmd, "output", mac, "err", err)
	}
	return
}

// ClearArpCacheByDev removes ARP cache entry for a specific IP/device pair
func (rem *Remote) ClearArpCacheByDev(destAddr string, deviceName string) error {
	cmd := fmt.Sprintf("ip neighbour delete to %v dev %v", destAddr, deviceName)
	output, err := rem.Run(cmd)
	// After IP is changed, "ip neighbour delete" fails with ENOENT. Since ARP cache is empty, this can be ignored.
	if err != nil && !strings.Contains(output, "No such file or directory") {
		logger.Debug("Failed to clear ARP cache", "on host", rem.Address(), "destAddr", destAddr,
			"deviceName", deviceName, "cmd", cmd, "output", output, "err", err)
		return errors.WrapPrefix(err, "Failed to clear ARP cache", 0)
	}
	return nil
}

// ClearArpCacheByVlan removes ARP cache entry for a specific IP/device pair by VLAN
func (rem *Remote) ClearArpCacheByVlan(destAddr string, vlan int) error {
	device, err := rem.GetNetworkInterfaceByVlan(vlan)
	if err != nil {
		return errors.WrapPrefix(err, fmt.Sprintf("Failed to get net device for VLAN %v on host %v",
			vlan, rem.Address()), 0)
	}
	return rem.ClearArpCacheByDev(destAddr, device)
}

// PopulateArpCache make sure ARP cache contains a value for the specified IP address
func (rem *Remote) PopulateArpCache(ipAddr string) error {
	cmd := fmt.Sprintf("ping -fc 10 %v", ipAddr)
	output, err := rem.Run(cmd)
	if err != nil {
		logger.Debug("Failed to populate ARP cache", "on host", rem.Address(), "ipAddr", ipAddr,
			"cmd", cmd, "output", output)
	}
	return err
}

// SetIpAddrByDev sets device's IP address/netmask
func (rem *Remote) SetIpAddrByDev(deviceName string, ipAddr string) (err error) {
	cmd := fmt.Sprintf("ifconfig %v %v up", deviceName, ipAddr)
	output, err := rem.Run(cmd)
	if err != nil {
		logger.Debug("Failed to set ip address", "host address", rem.Address(),
			"deviceName", deviceName, "ipAddr to set", ipAddr, "cmd", cmd, "output", output)
	}
	return
}

// SetIpAddrByDev sets device's IP address/netmask by VLAN
func (rem *Remote) SetIpAddrByVlan(vlan int, ipAddr string) (err error) {
	device, err := rem.GetNetworkInterfaceByVlan(vlan)
	if err != nil {
		return errors.WrapPrefix(err, fmt.Sprintf("Failed to get net device for VLAN %v on host %v",
			vlan, rem.Address()), 0)
	}
	return rem.SetIpAddrByDev(device, ipAddr)
}

// CheckNicLinkStatusByDev compares the actual device's (specified by device name) link status to the expected value
func (rem *Remote) CheckNicLinkStatusByDev(deviceName string, expectedNicLinkStatus string) (err error) {
	linkStatus, err := rem.GetInterfaceLinkStatusByDev(deviceName)
	if err != nil {
		return errors.WrapPrefix(err, fmt.Sprintf("Failed to get Data NIC link status on %v for device %v",
			rem.Address(), deviceName), 0)
	}

	logger.Debug(fmt.Sprintf("Data NIC link status for device %v on %v is %v", deviceName, rem.Address(), linkStatus))
	if linkStatus != expectedNicLinkStatus {
		return errors.WrapPrefix(err, fmt.Sprintf("Data NIC link status for device %v on node %v is %v, "+
			"but expected to be %v", deviceName, rem.Address(), linkStatus, expectedNicLinkStatus), 0)
	}

	return
}

// CheckNicLinkStatusByVlan compares the actual device's (specified by VLAN) link status to the expected value
func (rem *Remote) CheckNicLinkStatusByVlan(vlan int, expectedNicLinkStatus string) (err error) {
	linkStatus, err := rem.GetInterfaceLinkStatusByVlan(vlan)
	if err != nil {
		return errors.WrapPrefix(err, fmt.Sprintf("Failed to get Data NIC link status on %v for VLAN %v",
			rem.Address(), vlan), 0)
	}

	logger.Debug(fmt.Sprintf("Data NIC link status for vlan %v on %v is %v", vlan, rem.Address(), linkStatus))
	if linkStatus != expectedNicLinkStatus {
		return errors.WrapPrefix(err, fmt.Sprintf("Data NIC link status for vlan %v on node %v is %v, "+
			"but expected to be %v", vlan, rem.Address(), linkStatus, expectedNicLinkStatus), 0)
	}

	return
}

// BringInterfaceLinkDownByDev brings the interface's link down by device name
func (rem *Remote) BringInterfaceLinkDownByDev(deviceName string) error {
	logger.Action("Bringing network interface link DOWN by device name",
		"nodeAddress", rem.Address(), "deviceName", deviceName)
	return rem.ipLinkSet(deviceName, "down")
}

// BringInterfaceLinkUpByDev brings the interface's link up by device name
func (rem *Remote) BringInterfaceLinkUpByDev(deviceName string) error {
	logger.Action("Bringing network interface link UP by device name",
		"nodeAddress", rem.Address(), "deviceName", deviceName)
	return rem.ipLinkSet(deviceName, "up")
}

// BringInterfaceLinkDownByVlan brings the interface's link down by VLAN
func (rem *Remote) BringInterfaceLinkDownByVlan(vlan int) (err error) {
	logger.Info("Bringing network interface link DOWN by VLAN", "nodeAddress", rem.Address(), "vlan", vlan)
	device, err := rem.GetNetworkInterfaceByVlan(vlan)
	if err != nil {
		return errors.WrapPrefix(err, fmt.Sprintf("Failed to get net device for VLAN %v on host %v",
			vlan, rem.Address()), 0)
	}
	return rem.BringInterfaceLinkDownByDev(device)
}

// BringInterfaceLinkUpByVlan brings the interface's link up by VLAN
func (rem *Remote) BringInterfaceLinkUpByVlan(vlan int) (err error) {
	logger.Info("Bringing network interface link UP by VLAN", "nodeAddress", rem.Address(), "vlan", vlan)
	device, err := rem.GetNetworkInterfaceByVlan(vlan)
	if err != nil {
		return errors.WrapPrefix(err, fmt.Sprintf("Failed to get net device for VLAN %v on host %v",
			vlan, rem.Address()), 0)
	}
	return rem.BringInterfaceLinkUpByDev(device)
}

// BringInterfaceDown brings the NIC down by MAC address using ifdown - deprecated
func (rem *Remote) BringInterfaceDown(mac string) error {
	logger.Action("bring interface down", "nodeAddress", rem.Address(), "mac", mac)
	logger.Warn("Using 'ifdown' is deprecated - test should be updated to use 'ip link set' instead")
	nic, err := rem.GetNetworkInterfaceByMac(mac)
	if err != nil {
		return err
	}
	output, err := rem.Run(fmt.Sprintf("ifdown %v", nic))
	if err != nil {
		logger.Debug("Failed to bring NIC down", "nic", nic, "output", output)
	}
	return err
}

// BringInterfaceUp brings the NIC up by MAC address using ifup - deprecated
func (rem *Remote) BringInterfaceUp(mac string) error {
	logger.Action("bring interface up", "nodeAddress", rem.Address(), "mac", mac)
	logger.Warn("Using 'ifup' is deprecated - test should be updated to use 'ip link set' instead")
	nic, err := rem.GetNetworkInterfaceByMac(mac)
	if err != nil {
		return err
	}
	output, err := rem.Run(fmt.Sprintf("ifup %v", nic))
	if err != nil {
		logger.Debug("Failed to bring NIC up", "nic", nic, "output", output)
	}
	return err
}

func (rem *Remote) getRouteInfoByDestinationAddr(destinationAddr string) (routeInfo string, err error) {
	cmd := fmt.Sprintf("ip -o route get %v", destinationAddr)
	routeInfo, err = rem.RunWithTimeout(cmd, defaultTimeout)
	if err != nil {
		err = errors.WrapPrefix(err, fmt.Sprintf(
			"Filed to get routing info from %v to %v", rem.address, destinationAddr), 0)
		return
	}
	return
}

func (rem *Remote) GetDevByDestinationAddr(destinationAddr string) (dev string, err error) {
	routeInfo, err := rem.getRouteInfoByDestinationAddr(destinationAddr)
	dev, err = helputils.GetValueByFieldName("dev", routeInfo)
	if err != nil {
		err = errors.WrapPrefix(err, "Failed to get device name by destination address", 0)
	}
	return
}

func (rem *Remote) GetSourceAddrByDestination(destinationAddr string) (ipAddr string, err error) {
	routeInfo, err := rem.getRouteInfoByDestinationAddr(destinationAddr)
	ipAddr, err = helputils.GetValueByFieldName("src", routeInfo)
	if err != nil {
		err = errors.WrapPrefix(err, "Failed to get source address by destination address", 0)
	}
	return
}

func (rem *Remote) WriteIpNeighbours(targetFile string) error {
	cmd := "ip -s neigh > " + targetFile
	rem.RunWithTimeout(cmd, defaultTimeout)
	logger.Info("Dumped IP neighbours", "file", targetFile, "host", rem.Address())
	return nil
}

func (rem *Remote) IpRoutDel(match bool, ipRe *regexp.Regexp) error {
	out, err := rem.Run("ip rout")
	if err != nil {
		return errors.New(err)
	}

	ipLines := helputils.FilterStr(
		strings.Split(out, "\n"),
		func(line string) bool {
			return len(line) > 0 && unicode.IsDigit(rune(line[0])) && match == ipRe.MatchString(line)
		},
	)

	errChan := make(chan error)
	for _, ipLine := range ipLines {
		ip := strings.Split(ipLine, " ")[0]
		go func() {
			_, err := rem.Run("ip rout del " + ip)
			errChan <- err
		}()
	}
	if err := terrors.CollectAll(errChan, len(ipLines)); err != nil {
		return errors.Wrap(err, 0)
	}

	return nil
}
