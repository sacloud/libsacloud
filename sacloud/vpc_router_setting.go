package sacloud

type VPCRouterSetting struct {
	Interfaces         []*VPCRouterInterface        `json:",omitempty"`
	StaticNAT          *VPCRouterStaticNAT          `json:",omitempty"`
	PortForwarding     *VPCRouterPortForwarding     `json:",omitempty"`
	Firewall           *VPCRouterFirewall           `json:",omitempty"`
	DHCPServer         *VPCRouterDHCPServer         `json:",omitempty"`
	DHCPStaticMapping  *VPCRouterDHCPStaticMapping  `json:",omitempty"`
	L2TPIPsecServer    *VPCRouterL2TPIPsecServer    `json:",omitempty"`
	PPTPServer         *VPCRouterPPTPServer         `json:",omitempty"`
	RemoteAccessUsers  *VPCRouterRemoteAccessUsers  `json:",omitempty"`
	SiteToSiteIPsecVPN *VPCRouterSiteToSiteIPsecVPN `json:",omitempty"`
	VRID               int                          `json:",omitempty"`
}

type VPCRouterInterface struct {
	IPAddress        []string `json:",omitempty"`
	NetworkMaskLen   int      `json:",omitempty"`
	VirtualIPAddress string   `json:",omitempty"`
}

type VPCRouterStaticNAT struct {
	Config  []*VPCRouterStaticNATConfig `json:",omitempty"`
	Enabled string                      `json:",omitempty"`
}
type VPCRouterStaticNATConfig struct {
	GlobalAddress  string `json:",omitempty"`
	PrivateAddress string `json:",omitempty"`
}

type VPCRouterPortForwarding struct {
	Config  []*VPCRouterPortForwarding `json:",omitempty"`
	Enabled string                     `json:",omitempty"`
}
type VPCRouterPortForwardingConfig struct {
	Protocol       string `json:",omitempty"` // tcp/udp only
	GlobalPort     string `json:",omitempty"`
	PrivateAddress string `json:",omitempty"`
	PrivatePort    string `json:",omitempty"`
}

type VPCRouterFirewall struct {
	Config  []*VPCRouterFirewallSetting `json:",omitempty"`
	Enabled string                      `json:",omitempty"`
}
type VPCRouterFirewallSetting struct {
	Receive []*VPCRouterFirewallRule `json:",omitempty"`
	Send    []*VPCRouterFirewallRule `json:",omitempty"`
}
type VPCRouterFirewallRule struct {
	Action             string `json:",omitempty"`
	Protocol           string `json:",omitempty"`
	SourceNetwork      string `json:",omitempty"`
	SourcePort         string `json:",omitempty"`
	DestinationNetwork string `json:",omitempty"`
	DestinationPort    string `json:",omitempty"`
}

type VPCRouterDHCPServer struct {
	Config  []*VPCRouterDHCPServerConfig `json:",omitempty"`
	Enabled string                       `json:",omitempty"`
}
type VPCRouterDHCPServerConfig struct {
	Interface  string `json:",omitempty"`
	RangeStart string `json:",omitempty"`
	RangeStop  string `json:",omitempty"`
}

type VPCRouterDHCPStaticMapping struct {
	Config  []*VPCRouterDHCPStaticMappingConfig `json:",omitempty"`
	Enabled string                              `json:",omitempty"`
}
type VPCRouterDHCPStaticMappingConfig struct {
	IPAddress  string `json:",omitempty"`
	MACAddress string `json:",omitempty"`
}

type VPCRouterL2TPIPsecServer struct {
	Config  *VPCRouterL2TPIPsecServerConfig `json:",omitempty"`
	Enabled string                          `json:",omitempty"`
}

type VPCRouterL2TPIPsecServerConfig struct {
	PreSharedSecret string `json:",omitempty"`
	RangeStart      string `json:",omitempty"`
	RangeStop       string `json:",omitempty"`
}

type VPCRouterPPTPServer struct {
	Config  *VPCRouterPPTPServerConfig `json:",omitempty"`
	Enabled string                     `json:",omitempty"`
}
type VPCRouterPPTPServerConfig struct {
	RangeStart string `json:",omitempty"`
	RangeStop  string `json:",omitempty"`
}

type VPCRouterRemoteAccessUsers struct {
	Config  []*VPCRouterRemoteAccessUsers `json:",omitempty"`
	Enabled string                        `json:",omitempty"`
}
type VPCRouterRemoteAccessUsersConfig struct {
	UserName string `json:",omitempty"`
	Password string `json:",omitempty"`
}

type VPCRouterSiteToSiteIPsecVPN struct {
	Config  []*VPCRouterSiteToSiteIPsecVPNConfig `json:",omitempty"`
	Enabled string                               `json:",omitempty"`
}

type VPCRouterSiteToSiteIPsecVPNConfig struct {
	LocalPrefix     []string `json:",omitempty"`
	Peer            string   `json:",omitempty"`
	PreSharedSecret string   `json:",omitempty"`
	RemoteID        string   `json:",omitempty"`
	Routes          []string `json:",omitempty"`
}
