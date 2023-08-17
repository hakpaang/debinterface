package debinterface

import (
  "net"
  "strings"
)

func Unmarshal() *Interfaces {
  return UnmarshalWith("/etc/network/interfaces")
}

func UnmarshalWith(path string) *Interfaces {
  if err := isExist(path); err != nil {
    return nil
  }
  interfaces := &Interfaces{InterfacesPath: path}
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
  for _, adapter := range interfaces.Adapters {
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
  i.Adapters = append(i.Adapters, &Adapter{})
  i.Count++
  iLine := split(line)
  size := len(iLine)
  if len(i.Adapters) == i.Count {
    if size > 1 {
      i.Adapters[i.Count-1].Name = iLine[1]
    }
    if size > 2 {
      i.Adapters[i.Count-1].Family = toFamily(iLine[2])
    }
    if size > 3 {
      i.Adapters[i.Count-1].Protocol = toProtocol(iLine[3])
    }
  }
}

func readAddress(line string, i *Interfaces) {
  aLine := split(line)
  if len(i.Adapters) == i.Count && len(aLine) > 1 {
    address := aLine[1]
    if strings.ContainsAny(address, "/") {
      cidr, ipNet, _ := net.ParseCIDR(address)
      i.Adapters[i.Count-1].Address = cidr
      i.Adapters[i.Count-1].Netmask = ipv4Mask(ipNet.Mask)
    } else {
      i.Adapters[i.Count-1].Address = net.ParseIP(address)
    }

  }
}

func readNetmask(line string, i *Interfaces) {
  mLine := split(line)
  if len(i.Adapters) == i.Count && len(mLine) > 1 {
    i.Adapters[i.Count-1].Netmask = net.ParseIP(mLine[1])
  }
}

func readNetwork(line string, i *Interfaces) {
  nLine := split(line)
  if len(i.Adapters) == i.Count && len(nLine) > 1 {
    i.Adapters[i.Count-1].Network = net.ParseIP(nLine[1])
  }
}

func readBroadcast(line string, i *Interfaces) {
  bLine := split(line)
  if len(i.Adapters) == i.Count && len(bLine) > 1 {
    i.Adapters[i.Count-1].Broadcast = net.ParseIP(bLine[1])
  }
}

func readGateway(line string, i *Interfaces) {
  gLine := split(line)
  if len(i.Adapters) == i.Count && len(gLine) > 1 {
    i.Adapters[i.Count-1].Gateway = net.ParseIP(gLine[1])
  }
}

func readSource(line string, i *Interfaces) {
  sLine := split(line)
  if len(sLine) > 1 {
    i.Sources = append(i.Sources, sLine[1])
  }
}
