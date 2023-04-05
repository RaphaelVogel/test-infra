/*
Copyright SAP SE or an SAP affiliate company. All rights reserved. This file is licensed under the Apache Software License, v. 2 except as noted otherwise in the LICENSE file

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

// Code generated by lister-gen. DO NOT EDIT.

package v1alpha1

import (
	v1alpha1 "github.com/gardener/gardener/pkg/apis/core/v1alpha1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// ControllerRegistrationLister helps list ControllerRegistrations.
// All objects returned here must be treated as read-only.
type ControllerRegistrationLister interface {
	// List lists all ControllerRegistrations in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.ControllerRegistration, err error)
	// Get retrieves the ControllerRegistration from the index for a given name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.ControllerRegistration, error)
	ControllerRegistrationListerExpansion
}

// controllerRegistrationLister implements the ControllerRegistrationLister interface.
type controllerRegistrationLister struct {
	indexer cache.Indexer
}

// NewControllerRegistrationLister returns a new ControllerRegistrationLister.
func NewControllerRegistrationLister(indexer cache.Indexer) ControllerRegistrationLister {
	return &controllerRegistrationLister{indexer: indexer}
}

// List lists all ControllerRegistrations in the indexer.
func (s *controllerRegistrationLister) List(selector labels.Selector) (ret []*v1alpha1.ControllerRegistration, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.ControllerRegistration))
	})
	return ret, err
}

// Get retrieves the ControllerRegistration from the index for a given name.
func (s *controllerRegistrationLister) Get(name string) (*v1alpha1.ControllerRegistration, error) {
	obj, exists, err := s.indexer.GetByKey(name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("controllerregistration"), name)
	}
	return obj.(*v1alpha1.ControllerRegistration), nil
}
