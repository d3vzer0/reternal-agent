package modules

import (
	"encoding/base64"
	"encoding/json"
	"net"
)

type NetIface struct {
	Name string `json:"name"`
	Mac  string `json:"mac"`
	IP   string `json:"ip"`
}

func GetIfaces(input string) string {
	ifaces, _ := net.Interfaces()
	var netifaces []NetIface
	for _, i := range ifaces {
		addrs, _ := i.Addrs()
		for _, a := range addrs {
			iface := NetIface{Name: i.Name, Mac: i.HardwareAddr.String(), IP: a.String()}
			netifaces = append(netifaces, iface)
		}
	}
	json_content, _ := json.Marshal(netifaces)
	encode_base64 := base64.StdEncoding.EncodeToString(json_content[:])
	return encode_base64
}
