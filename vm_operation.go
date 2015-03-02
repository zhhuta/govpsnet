//vm_operation.go
package govpsnet

import (
	"fmt"
)

// Create new Instant
func (c *Client) CreatInstant(op *NewVMInstant) (*VirtualMachine, error) {
	new_vm := struct {
		VirtualMachine `json:"virtual_machine"`
	}{}
	err := c.post(vm_list, op, &new_vm)
	if err != nil {
		return nil, err
	}
	return &new_vm.VirtualMachine, nil
}

// Start vm
func (c *Client) Start(id int) error {
	route := fmt.Sprintf(vm_action_start, id)
	return c.post(route, nil, nil)
}

//Stop vm
func (c *Client) Stop(id int) error {
	route := fmt.Sprintf(vm_shutdown, id)
	return c.post(route, nil, nil)
}

// Reboot vm
func (c *Client) Reboot(id int) error {
	route := fmt.Sprintf(vm_reboot, id)
	return c.post(route, nil, nil)
}

//Reboot recovery
func (c *Client) RebootRecovery(id int) error {
	route := fmt.Sprintf(vm_reboot_recovery, id)
	return c.post(route, Params{"mode": "recovery"}, nil)
}

// Rebuid VM
func (c *Client) Rebuild(id int) error {
	route := fmt.Sprintf(vm_rebuild, id)
	return c.post(route, nil, nil)
}

// Rebuild Network of VM
func (c *Client) RebuildNetwork(id int) error {
	route := fmt.Sprintf(vm_network_rebuild, id)
	return c.post(route, nil, nil)
}

//Delete VM
func (c *Client) Delete(id int) error {
	route := fmt.Sprintf(vm_info, id)
	return c.delete(route)
}

// Get VM properties
func (c *Client) VirtualMachineInfo(id int) (*VirtualMachine, error) {
	url := fmt.Sprintf(vm_info, id)
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

// Get VMs list
func (c *Client) Virtual_Machines_List() (VirtualMachines, error) {
	// Inetersting return option - check how to init r struct
	r := []struct {
		VirtualMachine `json:"virtual_machine"`
	}{}
	err := c.get(vm_list, &r)
	if err != nil {
		return nil, err
	}
	//fmt.Println(r[0])
	return r, nil
}
