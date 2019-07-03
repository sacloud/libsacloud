package sacloud

import (
	"testing"

	"github.com/sacloud/libsacloud/v2/sacloud/accessor"
)

func TestAccessor(t *testing.T) {

	var v interface{}
	v = &Server{}
	if _, ok := v.(accessor.ID); !ok {
		t.Fatal("target is not implements accessor.ID")
	}
	if _, ok := v.(accessor.Availability); !ok {
		t.Fatal("target is not implements accessor.Availability")
	}
	if _, ok := v.(accessor.CreatedAt); !ok {
		t.Fatal("target is not implements accessor.CreatedAt")
	}
	if _, ok := v.(accessor.ModifiedAt); !ok {
		t.Fatal("target is not implements accessor.ModifiedAt")
	}
	if _, ok := v.(accessor.Instance); !ok {
		t.Fatal("target is not implements accessor.Instance")
	}
	if _, ok := v.(accessor.InstanceStatus); !ok {
		t.Fatal("target is not implements accessor.InstanceStatus")
	}
	if _, ok := v.(accessor.MemoryMB); !ok {
		t.Fatal("target is not implements accessor.MemoryMB")
	}

	v = &Disk{}
	if _, ok := v.(accessor.DiskMigratable); !ok {
		t.Fatal("target is not implements accessor.DiskMigratable")
	}
	if _, ok := v.(accessor.DiskPlan); !ok {
		t.Fatal("target is not implements accessor.DiskPlan")
	}
	if _, ok := v.(accessor.SizeMB); !ok {
		t.Fatal("target is not implements accessor.SizeMB")
	}
	if _, ok := v.(accessor.MigratedMB); !ok {
		t.Fatal("target is not implements accessor.MigratedMB")
	}

	v = &Note{}
	if _, ok := v.(accessor.Scope); !ok {
		t.Fatal("target is not implements accessor.Scope")
	}
}
