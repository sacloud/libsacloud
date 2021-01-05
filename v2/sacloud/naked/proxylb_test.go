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

package naked

import (
	"encoding/json"
	"testing"
)

func TestProxyLBCertificates_UnmarshalJSON(t *testing.T) {
	cases := []struct {
		in     *ProxyLBCertificates
		expect string
	}{
		{
			in: &ProxyLBCertificates{
				PrimaryCert: &ProxyLBCertificate{
					ServerCertificate:       "aaa",
					IntermediateCertificate: "bbb",
					PrivateKey:              "ccc",
				},
			},
			expect: `{"PrimaryCert":{"ServerCertificate":"aaa","IntermediateCertificate":"bbb","PrivateKey":"ccc"},"AdditionalCerts":[]}`,
		},
		{
			in: &ProxyLBCertificates{
				PrimaryCert: &ProxyLBCertificate{
					ServerCertificate:       "aaa",
					IntermediateCertificate: "bbb",
					PrivateKey:              "ccc",
				},
				AdditionalCerts: ProxyLBAdditionalCerts{
					{
						ServerCertificate:       "aaa",
						IntermediateCertificate: "bbb",
						PrivateKey:              "ccc",
					},
				},
			},
			expect: `{"PrimaryCert":{"ServerCertificate":"aaa","IntermediateCertificate":"bbb","PrivateKey":"ccc"},"AdditionalCerts":[{"ServerCertificate":"aaa","IntermediateCertificate":"bbb","PrivateKey":"ccc"}]}`,
		},
	}

	for _, tc := range cases {
		data, err := json.Marshal(tc.in)
		if err != nil {
			t.Fatal(err)
		}
		if string(data) != tc.expect {
			t.Fatalf("got unexpected JSON: expected: %s got: %s", tc.expect, data)
		}
	}
}
