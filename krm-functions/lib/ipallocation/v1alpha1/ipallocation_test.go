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

	"github.com/nokia/k8s-ipam/apis/ipam/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func TestParseKubeObjectNonEmpty(t *testing.T) {

	ipalloc := NewGenerator(
		v1.ObjectMeta{
			Name:      "obj",
			Namespace: "kube",
			Annotations: map[string]string{
				"test": "annotations",
			},
			Finalizers: []string{
				"finalizer1",
				"finalizer2",
			},
		}, v1alpha1.IPAllocationSpec{
			PrefixKind: v1alpha1.PrefixKindNetwork,
		})

	kObj, err := ipalloc.ParseKubeObject()
	if err != nil {
		t.Errorf(err.Error())
	}
	if kObj.GetAPIVersion() != "ipam.nephio.org/v1alpha1" {
		t.Errorf("api version not correct expected :%v got: %v", "ipam.nephio.org/v1alpha1", kObj.GetAPIVersion())
	}

	if len(kObj.GetAnnotations()) != 2 {
		t.Errorf("annotations size not correct expected :%v got: %v", 1, len(kObj.GetAnnotations()))
	}

}

func TestParseKubeObjectEmpty(t *testing.T) {
	ipallocObj := NewGenerator(v1.ObjectMeta{}, v1alpha1.IPAllocationSpec{})

	kObj, err := ipallocObj.ParseKubeObject()
	if err != nil {
		t.Errorf(err.Error())
	}
	if kObj.GetAPIVersion() != "ipam.nephio.org/v1alpha1" {
		t.Errorf("api version not correct expected :%v got: %v", "ipam.nephio.org/v1alpha1", kObj.GetAPIVersion())
	}
	if len(kObj.GetAnnotations()) != 1 {
		t.Errorf("annotations size not correct expected :%v got: %v", 1, len(kObj.GetAnnotations()))
	}
}
