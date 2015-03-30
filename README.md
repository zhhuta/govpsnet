# govpsnet
It's still alpha 

## Example 
``` 
package main

import (
	"fmt"
	"govpsnet"
)

func main() {
	c, _ := govpsnet.NewConnection("zhhuta@gmail.com", "0CQ0lad5LNV0QQ==")
	
	vm_list, err := c.Virtual_Machines_List()
	if err != nil {
		fmt.Println(err)
	}
	for _, vms := range vm_list {
		fmt.Println(vms.Label, vms.SystemTemplateID, vms.PrimaryIpAddress.IpAddress.IpAddress)	
	}
}
```
