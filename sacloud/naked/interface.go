package naked

import (
	"encoding/json"

	"github.com/sacloud/libsacloud/v2/sacloud/types"
)

// Interface サーバなどに接続されているNICの情報
type Interface struct {
	ID            types.ID          `json:",omitempty" yaml:"id,omitempty" structs:",omitempty"`
	MACAddress    string            `json:",omitempty" yaml:"mac_address,omitempty" structs:",omitempty"`
	IPAddress     string            `json:",omitempty" yaml:"ip_address,omitempty" structs:",omitempty"`
	UserIPAddress string            `json:",omitempty" yaml:"user_ip_address,omitempty" structs:",omitempty"`
	HostName      string            `json:",omitempty" yaml:"host_name,omitempty" structs:",omitempty"`
	Switch        *Switch           `json:",omitempty" yaml:"switch,omitempty" structs:",omitempty"`
	PacketFilter  *PacketFilterInfo `json:",omitempty" yaml:"packet_filter,omitempty" structs:",omitempty"`
	Server        *Server           `json:",omitempty" yaml:"server,omitempty" structs:",omitempty"`

	// Index 仮想フィールド、VPCルータなどでInterfaces(実体は[]*Interface)を扱う場合にUnmarshalJSONの中で設定される
	//
	// Findした際のAPIからの応答にも同名のフィールドが含まれるが無関係。
	Index int `json:"-" yaml:"-" structs:"-"`
}

// Interfaces Interface配列
//
// 配列中にnullが返ってくる(VPCルータなど)への対応のためのtype
type Interfaces []*Interface

// UnmarshalJSON 配列中にnullが返ってくる(VPCルータなど)への対応
func (i *Interfaces) UnmarshalJSON(b []byte) error {
	type alias Interfaces
	var a alias
	if err := json.Unmarshal(b, &a); err != nil {
		return err
	}

	var dest []*Interface
	for i, v := range a {
		if v != nil {
			v.Index = i
			dest = append(dest, v)
		}
	}

	*i = Interfaces(dest)
	return nil
}

// MarshalJSON 配列中にnullが入る場合(VPCルータなど)への対応
func (i *Interfaces) MarshalJSON() ([]byte, error) {
	max := 0
	for _, iface := range *i {
		if max < iface.Index {
			max = iface.Index
		}
	}

	var dest = make([]*Interface, max)
	for _, iface := range *i {
		dest[iface.Index] = iface
	}

	return json.Marshal(dest)
}
