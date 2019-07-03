package schema

import (
	"testing"

	"github.com/sacloud/libsacloud/internal/schema/meta"
	"github.com/sacloud/libsacloud/sacloud"
	"github.com/stretchr/testify/require"
)

func TestResource_LowerName(t *testing.T) {

	expects := []struct {
		resourceName string
		lowerName    string
	}{
		{
			resourceName: "Foo",
			lowerName:    "foo",
		},
		{
			resourceName: "IPAddress",
			lowerName:    "ip_address",
		},
		{
			resourceName: "ProxyLB",
			lowerName:    "proxy_lb",
		},
		{
			resourceName: "MonitorCPU",
			lowerName:    "monitor_cpu",
		},
		{
			resourceName: "SSDDisk",
			lowerName:    "ssd_disk",
		},
	}

	for _, expect := range expects {
		r := &Resource{
			Name: expect.resourceName,
		}
		require.Equal(t, expect.lowerName, r.FileSafeName())
	}

}

func TestResource_ImportStatements(t *testing.T) {

	var emptyList []string

	expects := []struct {
		resource          *Resource
		additionalImports []string
		imports           []string
	}{
		{
			resource:          &Resource{},
			additionalImports: emptyList,
			imports:           emptyList,
		},
		{
			resource: &Resource{
				operations: []*Operation{
					{
						arguments: []*Argument{
							{
								Type: meta.Static(sacloud.Client{}),
							},
						},
					},
				},
			},
			additionalImports: []string{"context"},
			imports:           wrapByDoubleQuote("context", "github.com/sacloud/libsacloud/sacloud"),
		},
	}

	for _, expect := range expects {
		require.Equal(t, expect.imports, expect.resource.ImportStatements(expect.additionalImports...))
	}
}

func TestResources_ImportStatements(t *testing.T) {
	var emptyList []string

	expects := []struct {
		resources         Resources
		additionalImports []string
		imports           []string
	}{
		{
			resources:         Resources{},
			additionalImports: emptyList,
			imports:           emptyList,
		},
		{
			resources: Resources([]*Resource{
				{
					operations: []*Operation{
						{
							arguments: []*Argument{
								{
									Type: meta.Static(sacloud.Client{}),
								},
							},
						},
					},
				},
				{
					operations: []*Operation{
						{
							arguments: []*Argument{
								{
									Type: meta.Static(sacloud.Client{}),
								},
							},
							results: Results([]*Result{
								{
									Model: &Model{
										Name: "Note",
									},
								},
							}),
						},
					},
				},
			}),
			additionalImports: []string{"context"},
			imports:           wrapByDoubleQuote("context", "github.com/sacloud/libsacloud/sacloud"),
		},
	}

	for _, expect := range expects {
		require.Equal(t, expect.imports, expect.resources.ImportStatements(expect.additionalImports...))
	}
}
