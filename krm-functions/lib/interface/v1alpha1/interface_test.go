/*
 Copyright 2023 Nephio.

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

package v1alpha1

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	nephioreqv1alpha1 "github.com/nephio-project/api/nf_requirements/v1alpha1"
)

func TestInterface(t *testing.T) {

	f := `apiVersion: req.nephio.org/v1alpha1
kind: Interface
metadata:
  name: n3
  annotations:
    config.kubernetes.io/local-config: "true"
spec:
  networkInstance:
    name: vpc-ran
  cniType: sriov
  attachmentType: vlan
`

	x := NewMutator(f)
	itfce, err := x.UnMarshal()
	if err != nil {
		t.Errorf("cannot unmarshal file: %s", err.Error())
	}

	cases := map[string]struct {
		fn   func(*nephioreqv1alpha1.Interface) string
		want string
	}{
		"NetworkInstanceName": {
			fn: func(itfce *nephioreqv1alpha1.Interface) string {
				return itfce.Spec.NetworkInstance.Name
			},
			want: "vpc-ran",
		},
		"CNITYpe": {
			fn: func(itfce *nephioreqv1alpha1.Interface) string {
				return string(itfce.Spec.CNIType)
			},
			want: "sriov",
		},
		"AttachementType": {
			fn: func(itfce *nephioreqv1alpha1.Interface) string {
				return string(itfce.Spec.AttachmentType)
			},
			want: "vlan",
		},
	}

	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			got := tc.fn(itfce)
			if diff := cmp.Diff(tc.want, got); diff != "" {
				t.Errorf("TestInterface: -want, +got:\n%s", diff)
			}
		})
	}
}
