//vm_scheme.go
package govpsnet

type IpAddresses []struct {
	IpAddress `json:"ip_addresses"`
}

type IpAddress struct {
	CloudID   int    `json:"cloud_id"`
	CreatedAt string `json:"created_at"`
	ID        int    `json:"id"`
	IpAddress string `json:"ip_address"`
	UpdatedAt string `json:"updated_at"`
}

type VirtualMachines []struct {
	VirtualMachine `json:"virtual_machine"`
}

type VirtualMachine struct {
	AdditionalIpAddresses []interface{} `json:"additional_ip_addresses"`
	AvailableDiskSize     interface{}   `json:"available_disk_size"`
	BackupLicenses        struct {
		R1softLicense interface{} `json:"r1soft_license"`
		RsyncLicense  interface{} `json:"rsync_license"`
	} `json:"backup_licenses"`
	BackupsEnabled      bool        `json:"backups_enabled"`
	BandwidthNodesCount int         `json:"bandwidth_nodes_count"`
	BandwidthUsed       int         `json:"bandwidth_used"`
	BetaNodesCount      int         `json:"beta_nodes_count"`
	Built               bool        `json:"built"`
	CloudID             int         `json:"cloud_id"`
	ConsumerID          int         `json:"consumer_id"`
	CreatedAt           string      `json:"created_at"`
	DailyNodesCount     int         `json:"daily_nodes_count"`
	DelayedStorage      bool        `json:"delayed_storage"`
	DelayedStorageUntil interface{} `json:"delayed_storage_until"`
	DeployedDiskSize    string      `json:"deployed_disk_size"`
	DomainName          string      `json:"domain_name"`
	FusionNodesCount    int         `json:"fusion_nodes_count"`
	Hostname            string      `json:"hostname"`
	ID                  int         `json:"id"`
	IpAddresses         IpAddresses `json:"ip_addresses"`
	Label               string      `json:"label"`
	Password            string      `json:"password"`
	PowerActionPending  bool        `json:"power_action_pending"`
	PrimaryIpAddress    struct {
		IpAddress struct {
			CloudID          int         `json:"cloud_id"`
			ConsumerID       int         `json:"consumer_id"`
			CreatedAt        string      `json:"created_at"`
			Description      interface{} `json:"description"`
			ID               int         `json:"id"`
			IpAddress        string      `json:"ip_address"`
			Netmask          string      `json:"netmask"`
			Network          string      `json:"network"`
			UpdatedAt        string      `json:"updated_at"`
			VirtualMachineID int         `json:"virtual_machine_id"`
		} `json:"ip_address"`
	} `json:"primary_ip_address"`
	RamNodesCount          int           `json:"ram_nodes_count"`
	RegularNodesCount      int           `json:"regular_nodes_count"`
	Running                bool          `json:"running"`
	SlicesCount            int           `json:"slices_count"`
	StorageNodesCount      int           `json:"storage_nodes_count"`
	SystemTemplateID       int           `json:"system_template_id"`
	UpdatedAt              string        `json:"updated_at"`
	UpgradeSchedules       []interface{} `json:"upgrade_schedules"`
	VirtualMachineLicenses []interface{} `json:"virtual_machine_licenses"`
	VirtualMachineStatus   string        `json:"virtual_machine_status"`
}

type NewVMInstant struct {
	Label             string `json:"label"`
	Hostname          string `json:"hostname"`
	SystemTemplateID  int    `json:"system_template_id"`
	CloudID           int    `json:"cloud_id"`
	BackupsEnabled    bool   `json:"backups_enabled"`
	RsyncBackupEnable bool   `json:"rsync_backups_enabled"`
	SlicesRequired    int    `json:"slices_required"`
}
