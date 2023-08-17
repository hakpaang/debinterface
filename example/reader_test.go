package example

import (
  "github.com/hakpaang/debinterface"
  "net"
  "reflect"
  "testing"
)

func TestUnmarshalWith(t *testing.T) {
  type args struct {
    path string
  }
  tests := []struct {
    name string
    args args
    want *debinterface.Interfaces
  }{
    {
      name: "input",
      args: args{
        path: "input",
      },
      want: &debinterface.Interfaces{
        InterfacesPath: "input",
        Sources:        []string{"/etc/network/interfaces.d/*"},
        Count:          3,
        Adapters: []*debinterface.Adapter{
          {Name: "lo", Auto: true, Family: debinterface.INET, Protocol: debinterface.LOOPBACK},
          {Name: "enp1s0", Auto: true, Family: debinterface.INET, Protocol: debinterface.STATIC, Address: net.ParseIP("192.168.50.235"), Netmask: net.ParseIP("255.255.255.0")},
          {Name: "enp2s0", Auto: true, Family: debinterface.INET, Protocol: debinterface.STATIC, Address: net.ParseIP("192.168.50.234"), Netmask: net.ParseIP("255.255.255.0"), Gateway: net.ParseIP("192.168.50.1")},
        },
      },
    },
  }
  for _, tt := range tests {
    t.Run(tt.name, func(t *testing.T) {
      if got := debinterface.UnmarshalWith(tt.args.path); !reflect.DeepEqual(got, tt.want) {
        t.Errorf("UnmarshalWith() = %v, want %v", got, tt.want)
      }
    })
  }
}
