package debinterface

import (
  "errors"
  "fmt"
)

func Marshal(interfaces *Interfaces) error {
  path := interfaces.InterfacesPath
  if len(path) == 0 {
    return errors.New("empty path")
  }
  var lines []string

  for _, source := range interfaces.Sources {
    if len(source) > 0 {
      lines = append(lines, fmt.Sprintf("source %s", source))
    }
  }
  lines = append(lines, "")

  for _, adapter := range interfaces.Adapters {
    if !adapter.validateAll() {
      continue
    }
    if adapter.Auto {
      lines = append(lines, fmt.Sprintf("auto %s", adapter.Name))
    }
    if adapter.Hotplug {
      lines = append(lines, fmt.Sprintf("allow-hotplug %s", adapter.Name))
    }
    lines = append(lines, fmt.Sprintf("iface %s %s %s", adapter.Name, adapter.Family.fromFamily(), adapter.Protocol.fromProtocol()))
    if adapter.Address != nil {
      lines = append(lines, fmt.Sprintf("address %s", adapter.Address))
    }
    if adapter.Netmask != nil {
      lines = append(lines, fmt.Sprintf("netmask %s", adapter.Netmask))
    }
    if adapter.Network != nil {
      lines = append(lines, fmt.Sprintf("network %s", adapter.Network))
    }
    if adapter.Broadcast != nil {
      lines = append(lines, fmt.Sprintf("broadcast %s", adapter.Broadcast))
    }
    if adapter.Gateway != nil {
      lines = append(lines, fmt.Sprintf("gateway %s", adapter.Gateway))
    }
    lines = append(lines, "")
  }

  return writeFileByLine(path, lines)
}
