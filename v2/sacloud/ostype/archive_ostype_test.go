// Copyright 2016-2022 The Libsacloud Authors
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

package ostype

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestArchiveOSTypeDefinitions(t *testing.T) {
	// OSTypeShortNamesへの追加忘れを防ぐ
	require.Equal(t, len(OSTypeShortNames), len(ArchiveOSTypes)+1) // miracleのエイリアス分で+1
}
