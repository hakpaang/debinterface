package example

import (
  "github.com/hakpaang/debinterface"
  "net"
  "testing"
)

func TestMarshal(t *testing.T) {
  type args struct {
    interfaces *debinterface.Interfaces
  }
  tests := []struct {
    name    string
    args    args
    wantErr bool
  }{
    {
      name:    "output",
      wantErr: false,
      args: args{interfaces: &debinterface.Interfaces{
        InterfacesPath: "output",
        Sources:        []string{"/etc/network/interfaces.d/*"},
        Count:          3,
        Adapters: []*debinterface.Adapter{
          {Name: "lo", Auto: true, Family: debinterface.INET, Protocol: debinterface.LOOPBACK},
          {Name: "enp1s0", Auto: true, Family: debinterface.INET, Protocol: debinterface.STATIC, Address: net.ParseIP("192.168.50.235"), Netmask: net.ParseIP("255.255.255.0")},
          {Name: "enp2s0", Auto: true, Family: debinterface.INET, Protocol: debinterface.STATIC, Address: net.ParseIP("192.168.50.234"), Netmask: net.ParseIP("255.255.255.0"), Gateway: net.ParseIP("192.168.50.1")},
        },
      },
      },
    },
  }
  for _, tt := range tests {
    t.Run(tt.name, func(t *testing.T) {
      if err := debinterface.Marshal(tt.args.interfaces); (err != nil) != tt.wantErr {
        t.Errorf("Marshal() error = %v, wantErr %v", err, tt.wantErr)
      }
    })
  }
}
