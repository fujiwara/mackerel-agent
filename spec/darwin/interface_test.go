// +build darwin

package darwin

import (
	"os"
	"testing"
)

func TestInterfaceKey(t *testing.T) {
	g := &InterfaceGenerator{}

	if g.Key() != "interface" {
		t.Error("key should be interface")
	}
}

func TestInterfaceGenerate(t *testing.T) {
	g := &InterfaceGenerator{}
	value, err := g.Generate()
	if err != nil {
		t.Errorf("should not raise error: %v", err)
	}

	interfaces, typeOk := value.([]map[string]interface{})
	if !typeOk {
		t.Errorf("value should be slice of map. %+v", value)
	}

	if os.Getenv("TRAVIS") != "" {
		t.Skip("Skip in Travis for now")
	}

	if len(interfaces) == 0 {
		t.Error("should have at least 1 interface")
	}

	iface := interfaces[0]
	if _, ok := iface["ipAddress"]; !ok {
		t.Error("interface should have ipAddress")
	}
	if _, ok := iface["macAddress"]; !ok {
		t.Error("interface should have macAddress")
	}
	if _, ok := iface["netmask"]; !ok {
		t.Error("interface should have netmask")
	}
	if _, ok := iface["defaultGateway"]; !ok {
		t.Error("interface should have defaultGateway")
	}
}

func TestGenerateByIfconfigCommand(t *testing.T) {
	g := &InterfaceGenerator{}
	interfaces, err := g.generateByIfconfigCommand()
	if err != nil {
		t.Log("Skip: should not raise error")
	}

	name := "eth0"
	if _, ok := interfaces[name]; !ok {
		t.Log("Skip: should have interfaces")
	}

	iface := interfaces[name]
	if len(iface) == 0 {
		t.Log("Skip: should have item")
	}
	if _, ok := iface["ipAddress"]; !ok {
		t.Log("Skip: interface should have ipAddress")
	}
	if _, ok := iface["macAddress"]; !ok {
		t.Log("Skip: interface should have macAddress")
	}
	if _, ok := iface["netmask"]; !ok {
		t.Log("Skip: interface should have netmask")
	}
	if _, ok := iface["defaultGateway"]; !ok {
		t.Log("Skip: interface should have defaultGateway")
	}
}
