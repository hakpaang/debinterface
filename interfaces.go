package debinterface

import "net"

type Interfaces struct {
	interfacesPath string
	adapters       []*Adapter
	sources        []string
	count          int
}

type Family int

const (
	INET Family = 1 + iota
	INET6
)

func toFamily(f string) Family {
	switch f {
	case "inet":
		return INET
	case "inet6":
		return INET6
	default:
		return 0
	}
}

type Protocol int

const (
	DHCP Protocol = 1 + iota
	STATIC
	LOOPBACK
	MANUAL
)

func toProtocol(f string) Protocol {
	switch f {
	case "dhcp":
		return DHCP
	case "static":
		return STATIC
	case "loopback":
		return LOOPBACK
	case "manual":
		return MANUAL
	default:
		return 0
	}
}

type Adapter struct {
	Name      string
	Auto      bool
	Hotplug   bool
	Address   net.IP
	Netmask   net.IP
	Network   net.IP
	Broadcast net.IP
	Gateway   net.IP
	Family    Family
	Protocol  Protocol
}
