# debinterface

 A Golang library for dealing with the interfaces configuration file at /etc/network/interfaces.

# example

```go
func TestUnmarshalWith(t *testing.T) {
	type args struct {
		path string
	}
	tests := []struct {
		name string
		args args
		want *Interfaces
	}{
		{
			name: "input",
			args: args{
				path: "input",
			},
			want: &Interfaces{
				interfacesPath: "input",
				sources:        []string{"/etc/network/interfaces.d/*"},
				count:          3,
				adapters: []*Adapter{
					{Name: "lo", Auto: true, Family: INET, Protocol: LOOPBACK},
					{Name: "enp1s0", Auto: true, Family: INET, Protocol: STATIC, Address: net.ParseIP("192.168.50.235"), Netmask: net.ParseIP("255.255.255.0")},
					{Name: "enp2s0", Auto: true, Family: INET, Protocol: STATIC, Address: net.ParseIP("192.168.50.234"), Netmask: net.ParseIP("255.255.255.0"), Gateway: net.ParseIP("192.168.50.1")},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := UnmarshalWith(tt.args.path); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UnmarshalWith() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMarshal(t *testing.T) {
	type args struct {
		interfaces *Interfaces
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name:    "output",
			wantErr: false,
			args: args{interfaces: &Interfaces{
				interfacesPath: "output",
				sources:        []string{"/etc/network/interfaces.d/*"},
				count:          3,
				adapters: []*Adapter{
					{Name: "lo", Auto: true, Family: INET, Protocol: LOOPBACK},
					{Name: "enp1s0", Auto: true, Family: INET, Protocol: STATIC, Address: net.ParseIP("192.168.50.235"), Netmask: net.ParseIP("255.255.255.0")},
					{Name: "enp2s0", Auto: true, Family: INET, Protocol: STATIC, Address: net.ParseIP("192.168.50.234"), Netmask: net.ParseIP("255.255.255.0"), Gateway: net.ParseIP("192.168.50.1")},
				},
			},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := Marshal(tt.args.interfaces); (err != nil) != tt.wantErr {
				t.Errorf("Marshal() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
```
