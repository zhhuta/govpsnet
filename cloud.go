//cloud.go

package govpsnet

var cloud_lists = "available_clouds.api10json"

type Cloud struct {
	Available             bool    `json:"available"`
	BandwidthNodesEnabled bool    `json:"bandwidth_nodes_enabled"`
	CloudVersion          int     `json:"cloud_version"`
	CurrencyCode          string  `json:"currency_code"`
	DailyNodeBandwidth    int     `json:"daily_node_bandwidth"`
	FusionIo              bool    `json:"fusion_io"`
	ID                    int     `json:"id"`
	Label                 string  `json:"label"`
	MonthlyNodeBandwidth  int     `json:"monthly_node_bandwidth"`
	PremiumPrice          float64 `json:",string"`
	SystemTemplates       []struct {
		ID    int    `json:"id"`
		Label string `json:"label"`
	} `json:"system_templates"`
}
type Clouds []struct {
	Cloud `json:"cloud"`
}

func (c *Client) Get_Clouds() (Clouds, error) {
	r := []struct {
		Cloud `json:"cloud"`
	}{}
	err := c.get(cloud_lists, &r)
	if err != nil {
		return nil, err
	}
	return r, nil

}
