/*
Copyright 2016 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package github

import (
	"testing"
)

var globalSecret = `
'*':
  - value: abc
    created_at: 2020-10-02T15:00:00Z
  - value: key2
    created_at: 2018-10-02T15:00:00Z
`

var defaultTokenGenerator = func() []byte {
	return []byte(globalSecret)
}

// echo -n 'BODY' | openssl dgst -sha1 -hmac KEY
func TestValidatePayload(t *testing.T) {
	var testcases = []struct {
		payload        string
		sig            string
		tokenGenerator func() []byte
		valid          bool
	}{
		{
			"{}",
			"sha1=db5c76f4264d0ad96cf21baec394964b4b8ce580",
			defaultTokenGenerator,
			true,
		},
		{
			"{}",
			"db5c76f4264d0ad96cf21baec394964b4b8ce580",
			defaultTokenGenerator,
			false,
		},
		{
			"{}",
			"",
			defaultTokenGenerator,
			false,
		},
		{
			"{}",
			"",
			defaultTokenGenerator,
			false,
		},
	}
	for _, tc := range testcases {
		if ValidatePayload([]byte(tc.payload), tc.sig, tc.tokenGenerator) != tc.valid {
			t.Errorf("Wrong validation for %+v", tc)
		}
	}
}
