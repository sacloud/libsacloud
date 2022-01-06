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

package naked

import (
	"encoding/json"
	"reflect"
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

func TestProxyLBACMESetting_UnmarshalJSON(t *testing.T) {
	type fields struct {
		Enabled         bool
		CommonName      string
		SubjectAltNames []string
	}
	type args struct {
		data []byte
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "unmarshal",
			fields: fields{
				Enabled:         true,
				CommonName:      "example.com",
				SubjectAltNames: []string{"test1.example.com", "test2.example.com", "test3.example.com", "test4.example.com"},
			},
			args: args{
				data: []byte(`{"Enabled":true,"CommonName":"example.com","SubjectAltNames":"test1.example.com\ntest2.example.com,test3.example.com test4.example.com"}`),
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &ProxyLBACMESetting{
				Enabled:         tt.fields.Enabled,
				CommonName:      tt.fields.CommonName,
				SubjectAltNames: tt.fields.SubjectAltNames,
			}
			if err := p.UnmarshalJSON(tt.args.data); (err != nil) != tt.wantErr {
				t.Errorf("UnmarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
			}
			expect := &ProxyLBACMESetting{
				Enabled:         tt.fields.Enabled,
				CommonName:      tt.fields.CommonName,
				SubjectAltNames: tt.fields.SubjectAltNames,
			}
			if !reflect.DeepEqual(p, expect) {
				t.Errorf("MarshalJSON() got = %v, want %v", p, expect)
			}
		})
	}
}

func TestProxyLBACMESetting_MarshalJSON(t *testing.T) {
	type fields struct {
		Enabled         bool
		CommonName      string
		SubjectAltNames []string
	}
	tests := []struct {
		name    string
		fields  fields
		want    []byte
		wantErr bool
	}{
		{
			name: "marshal",
			fields: fields{
				Enabled:         true,
				CommonName:      "example.com",
				SubjectAltNames: []string{"test1.example.com", "test2.example.com", "test3.example.com", "test4.example.com"},
			},
			want:    []byte(`{"Enabled":true,"CommonName":"example.com","SubjectAltNames":"test1.example.com,test2.example.com,test3.example.com,test4.example.com"}`),
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := ProxyLBACMESetting{
				Enabled:         tt.fields.Enabled,
				CommonName:      tt.fields.CommonName,
				SubjectAltNames: tt.fields.SubjectAltNames,
			}
			got, err := p.MarshalJSON()
			if (err != nil) != tt.wantErr {
				t.Errorf("MarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MarshalJSON() got = %v, want %v", string(got), string(tt.want))
			}
		})
	}
}
