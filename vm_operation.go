//vm_operation.go
package govpsnet

import (
	"fmt"
)

func (c *Client) CreatInstant() error {

}

// Start vm
func (c *Client) Start(id int) error {
	route := fmt.Sprintf("virtual_machines/%d/power_on.api10json", id)
	return c.post(route, nil, nil)
}

//Stop vm
func (c *Client) Stop(id int) error {
	route := fmt.Sprintf("virtual_machines/%d/shutdown.api10json", id)
	return c.post(route, nil, nil)
}

// Reboot vm
func (c *Client) Reboot(id int) error {
	route := fmt.Sprintf("virtual_machines/%d/reboot.api10json", id)
	return c.post(route, nil, nil)
}

//Reboot recovery
func (c *Client) RebootRecovery(id int) error {
	route := fmt.Sprintf("virtual_machines/%d/reboot.api10json", id)
	return c.post(route, Params{"mode": "recovery"}, nil)
}
func (c *Client) Rebuild(id int) error {
	route := fmt.Sprintf("virtual_machines/%d/rebuild.api10json", id)
	return c.post(route, nil, nil)
}
func (c *Client) RebuildNetwork(id int) error {
	route := fmt.Sprintf("virtual_machines/{ID}/rebuild_network.api10json", id)
	return c.post(route, nil, nil)
}
func (c *Client) VirtualMachineInfo(id int) (*VirtualMachine, error) {
	url := fmt.Sprintf("virtual_machines/%d.api10json", id)
	r := struct {
		VirtualMachine `json:"virtual_machine"`
	}{}
	//var r VirtualMachine
	err := c.get(url, &r)
	if err != nil {
		return nil, err
	}
	return &r.VirtualMachine, nil
}

func (c *Client) Virtual_Machines_List() (VirtualMachines, error) {
	// Inetersting return option - check how to init r struct
	r := []struct {
		VirtualMachine `json:"virtual_machine"`
	}{}
	err := c.get("virtual_machines.api10json", &r)
	if err != nil {
		return nil, err
	}
	//fmt.Println(r[0])
	return r, nil
}
