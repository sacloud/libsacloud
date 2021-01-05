// Copyright 2016-2021 The Libsacloud Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package dsl

import (
	"testing"

	"github.com/sacloud/libsacloud/v2/internal/dsl/meta"
	"github.com/sacloud/libsacloud/v2/sacloud"
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
				Operations: []*Operation{
					{
						Arguments: []*Argument{
							{
								Type: meta.Static(sacloud.Client{}),
							},
						},
					},
				},
			},
			additionalImports: []string{"context"},
			imports:           wrapByDoubleQuote("context", "github.com/sacloud/libsacloud/v2/sacloud"),
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
					Operations: []*Operation{
						{
							Arguments: []*Argument{
								{
									Type: meta.Static(sacloud.Client{}),
								},
							},
						},
					},
				},
				{
					Operations: []*Operation{
						{
							Arguments: []*Argument{
								{
									Type: meta.Static(sacloud.Client{}),
								},
							},
							Results: Results([]*Result{
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
			imports:           wrapByDoubleQuote("context", "github.com/sacloud/libsacloud/v2/sacloud"),
		},
	}

	for _, expect := range expects {
		require.Equal(t, expect.imports, expect.resources.ImportStatements(expect.additionalImports...))
	}
}
