package api

type IPv6NetAPI struct {
	*baseAPI
}

func NewIPv6NetAPI(client *Client) *IPv6NetAPI {
	return &IPv6NetAPI{
		&baseAPI{
			client: client,
			FuncGetResourceURL: func() string {
				return "ipv6net"
			},
		},
	}
}
