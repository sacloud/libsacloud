// Copyright 2016-2020 The Libsacloud Authors
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

package archive

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestDirector_Builder(t *testing.T) {
	dummySource := bytes.NewBufferString("dummy")

	cases := []struct {
		msg string
		in  *Director
		out Builder
	}{
		{
			msg: "BlankBuilder",
			in: &Director{
				Name:         "blank",
				SizeGB:       20,
				SourceReader: dummySource,
			},
			out: &BlankArchiveBuilder{
				Name:         "blank",
				SizeGB:       20,
				SourceReader: dummySource,
			},
		},
		{
			msg: "FromSharedArchiveBuilder",
			in: &Director{
				Name:            "shared",
				SourceSharedKey: "is1a:0000:xxxx",
			},
			out: &FromSharedArchiveBuilder{
				Name:            "shared",
				SourceSharedKey: "is1a:0000:xxxx",
			},
		},
		{
			msg: "TransferArchiveBuilder",
			in: &Director{
				Name:              "transfer",
				SourceArchiveID:   1,
				SourceArchiveZone: "is1a",
			},
			out: &TransferArchiveBuilder{
				Name:              "transfer",
				SourceArchiveID:   1,
				SourceArchiveZone: "is1a",
			},
		},
		{
			msg: "StandardArchiveBuilder_with_ArchiveID",
			in: &Director{
				Name:            "standard-with-archive-id",
				SourceArchiveID: 1,
			},
			out: &StandardArchiveBuilder{
				Name:            "standard-with-archive-id",
				SourceArchiveID: 1,
			},
		},
		{
			msg: "StandardArchiveBuilder_with_DiskID",
			in: &Director{
				Name:         "standard-with-disk-id",
				SourceDiskID: 1,
			},
			out: &StandardArchiveBuilder{
				Name:         "standard-with-disk-id",
				SourceDiskID: 1,
			},
		},
	}

	for _, tc := range cases {
		require.EqualValues(t, tc.out, tc.in.Builder(), tc.msg)
	}
}
