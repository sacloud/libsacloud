package disk

import (
	"context"
	"testing"

	"github.com/sacloud/libsacloud/v2/pkg/size"
	"github.com/sacloud/libsacloud/v2/sacloud"
)

func TestCreateRequest_ToRequestParameter(t *testing.T) {
	in := &CreateRequest{SizeGB: 2}
	expect := &sacloud.DiskCreateRequest{
		SizeMB: size.GiBToMiB(2),
	}

	req, err := in.ToRequestParameter(context.Background(), nil)
	if err != nil {
		t.Fatal(err)
	}
	if req.SizeMB != expect.SizeMB {
		t.Error("ToRequestParameter returns unexpected value")
	}
}
