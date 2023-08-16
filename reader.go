package debinterface

import (
	"net"
	"strings"
)

func Unmarshal() *Interfaces {
	return UnmarshalWith("/etc/network/interfaces")
}

func UnmarshalWith(path string) *Interfaces {
	if !isExist(path) {
		return nil
	}
	interfaces := &Interfaces{interfacesPath: path}
	lines, err := readFileByLine(path)
	if err != nil {
		return nil
	}
	var autoSet = make(Set)
	var hotplugSet = make(Set)
	for _, line := range lines {
		if len(line) == 0 {
			continue
		} else if strings.HasPrefix(line, "#") {
			continue
		} else if strings.HasPrefix(line, "auto") {
			autoSet.add(readIndex(line, 1))
			continue
		} else if strings.HasPrefix(line, "allow-hotplug") {
			hotplugSet.add(readIndex(line, 1))
			continue
		} else if strings.HasPrefix(line, "iface") {
			readIface(line, interfaces)
			continue
		} else if strings.HasPrefix(line, "address") {
			readAddress(line, interfaces)
			continue
		} else if strings.HasPrefix(line, "netmask") {
			readNetmask(line, interfaces)
			continue
		} else if strings.HasPrefix(line, "network") {
			readNetwork(line, interfaces)
			continue
		} else if strings.HasPrefix(line, "broadcast") {
			readBroadcast(line, interfaces)
			continue
		} else if strings.HasPrefix(line, "gateway") {
			readGateway(line, interfaces)
			continue
		} else if strings.HasPrefix(line, "source") {
			readSource(line, interfaces)
			continue
		}
	}
	for _, adapter := range interfaces.adapters {
		adapter.Auto = autoSet.has(adapter.Name)
		adapter.Hotplug = hotplugSet.has(adapter.Name)
	}
	return interfaces
}

func readIndex(l string, index int) string {
	a := split(l)
	if len(a) > index {
		return a[index]
	}
	return ""
}

func readIface(line string, i *Interfaces) {
	i.adapters = append(i.adapters, &Adapter{})
	i.count++
	iLine := split(line)
	size := len(iLine)
	if len(i.adapters) == i.count {
		if size > 1 {
			i.adapters[i.count-1].Name = iLine[1]
		}
		if size > 2 {
			i.adapters[i.count-1].Family = toFamily(iLine[2])
		}
		if size > 3 {
			i.adapters[i.count-1].Protocol = toProtocol(iLine[3])
		}
	}
}

func readAddress(line string, i *Interfaces) {
	aLine := split(line)
	if len(i.adapters) == i.count && len(aLine) > 1 {
		address := aLine[1]
		if strings.ContainsAny(address, "/") {
			cidr, ipNet, _ := net.ParseCIDR(address)
			i.adapters[i.count-1].Address = cidr
			i.adapters[i.count-1].Netmask = ipv4Mask(ipNet.Mask)
		} else {
			i.adapters[i.count-1].Address = net.ParseIP(address)
		}

	}
}

func readNetmask(line string, i *Interfaces) {
	mLine := split(line)
	if len(i.adapters) == i.count && len(mLine) > 1 {
		i.adapters[i.count-1].Netmask = net.ParseIP(mLine[1])
	}
}

func readNetwork(line string, i *Interfaces) {
	nLine := split(line)
	if len(i.adapters) == i.count && len(nLine) > 1 {
		i.adapters[i.count-1].Network = net.ParseIP(nLine[1])
	}
}

func readBroadcast(line string, i *Interfaces) {
	bLine := split(line)
	if len(i.adapters) == i.count && len(bLine) > 1 {
		i.adapters[i.count-1].Broadcast = net.ParseIP(bLine[1])
	}
}

func readGateway(line string, i *Interfaces) {
	gLine := split(line)
	if len(i.adapters) == i.count && len(gLine) > 1 {
		i.adapters[i.count-1].Gateway = net.ParseIP(gLine[1])
	}
}

func readSource(line string, i *Interfaces) {
	sLine := split(line)
	if len(sLine) > 1 {
		i.sources = append(i.sources, sLine[1])
	}
}
