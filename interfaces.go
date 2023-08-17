package debinterface

import (
  "net"
)

type Interfaces struct {
  InterfacesPath string
  Adapters       []*Adapter
  Sources        []string
  Count          int
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

func (f Family) fromFamily() string {
  switch f {
  case INET:
    return "inet"
  case INET6:
    return "inet6"
  default:
    return ""
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

func (p Protocol) fromProtocol() string {
  switch p {
  case DHCP:
    return "dhcp"
  case STATIC:
    return "static"
  case LOOPBACK:
    return "loopback"
  case MANUAL:
    return "manual"
  default:
    return ""
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

func (adapter *Adapter) validateAll() bool {
  if !adapter.validateName() {
    return false
  }
  if !adapter.validateAddress() {
    return false
  }
  if !adapter.validateNetmask() {
    return false
  }
  if !adapter.validateNetwork() {
    return false
  }
  if !adapter.validateBroadcast() {
    return false
  }
  if !adapter.validateGateway() {
    return false
  }
  if !adapter.validateFamily() {
    return false
  }
  if !adapter.validateProtocol() {
    return false
  }
  return true
}

func (adapter *Adapter) validateName() bool {
  if len(adapter.Name) == 0 {
    return false
  }
  return true
}

func (adapter *Adapter) validateAddress() bool {
  return true
}

func (adapter *Adapter) validateNetmask() bool {
  return true
}

func (adapter *Adapter) validateNetwork() bool {
  return true
}

func (adapter *Adapter) validateBroadcast() bool {
  return true
}

func (adapter *Adapter) validateGateway() bool {
  return true
}

func (adapter *Adapter) validateFamily() bool {
  switch adapter.Family {
  case INET:
    return true
  case INET6:
    return true
  }
  return false
}

func (adapter *Adapter) validateProtocol() bool {
  switch adapter.Protocol {
  case DHCP:
    return true
  case STATIC:
    return true
  case MANUAL:
    return true
  case LOOPBACK:
    return true
  }
  return false
}
