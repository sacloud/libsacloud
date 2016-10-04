package api

type SubnetAPI struct {
	*baseAPI
}

func NewSubnetAPI(client *Client) *SubnetAPI {
	return &SubnetAPI{
		&baseAPI{
			client: client,
			FuncGetResourceURL: func() string {
				return "subnet"
			},
		},
	}
}
