package test

import (
	"context"
	"fmt"

	"github.com/sacloud/libsacloud/v2/sacloud"
	"github.com/sacloud/libsacloud/v2/sacloud/search"
)

func lookupDNSByName(caller sacloud.APICaller, zoneName string) (*sacloud.DNS, error) {
	dnsOp := sacloud.NewDNSOp(caller)
	searched, err := dnsOp.Find(context.Background(), &sacloud.FindCondition{
		Count: 1,
		Filter: search.Filter{
			search.Key("Name"): zoneName,
		},
	})
	if err != nil {
		return nil, err
	}
	if searched.Count == 0 {
		return nil, fmt.Errorf("dns zone %q is not found", zoneName)
	}

	// 部分一致などにより予期せぬゾーンとマッチしていないかチェック
	if searched.DNS[0].Name != zoneName {
		return nil, fmt.Errorf("fetched dns zone does not match to desired: param: %s, actual: %s", zoneName, searched.DNS[0].Name)
	}

	return searched.DNS[0], nil
}
