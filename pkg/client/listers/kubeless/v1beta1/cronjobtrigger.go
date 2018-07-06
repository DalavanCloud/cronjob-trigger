/*
Copyright (c) 2016-2017 Bitnami

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

// This file was automatically generated by lister-gen

package v1beta1

import (
	v1beta1 "github.com/kubeless/cronjob-trigger/pkg/apis/kubeless/v1beta1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// CronJobTriggerLister helps list CronJobTriggers.
type CronJobTriggerLister interface {
	// List lists all CronJobTriggers in the indexer.
	List(selector labels.Selector) (ret []*v1beta1.CronJobTrigger, err error)
	// CronJobTriggers returns an object that can list and get CronJobTriggers.
	CronJobTriggers(namespace string) CronJobTriggerNamespaceLister
	CronJobTriggerListerExpansion
}

// cronJobTriggerLister implements the CronJobTriggerLister interface.
type cronJobTriggerLister struct {
	indexer cache.Indexer
}

// NewCronJobTriggerLister returns a new CronJobTriggerLister.
func NewCronJobTriggerLister(indexer cache.Indexer) CronJobTriggerLister {
	return &cronJobTriggerLister{indexer: indexer}
}

// List lists all CronJobTriggers in the indexer.
func (s *cronJobTriggerLister) List(selector labels.Selector) (ret []*v1beta1.CronJobTrigger, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1beta1.CronJobTrigger))
	})
	return ret, err
}

// CronJobTriggers returns an object that can list and get CronJobTriggers.
func (s *cronJobTriggerLister) CronJobTriggers(namespace string) CronJobTriggerNamespaceLister {
	return cronJobTriggerNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// CronJobTriggerNamespaceLister helps list and get CronJobTriggers.
type CronJobTriggerNamespaceLister interface {
	// List lists all CronJobTriggers in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1beta1.CronJobTrigger, err error)
	// Get retrieves the CronJobTrigger from the indexer for a given namespace and name.
	Get(name string) (*v1beta1.CronJobTrigger, error)
	CronJobTriggerNamespaceListerExpansion
}

// cronJobTriggerNamespaceLister implements the CronJobTriggerNamespaceLister
// interface.
type cronJobTriggerNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all CronJobTriggers in the indexer for a given namespace.
func (s cronJobTriggerNamespaceLister) List(selector labels.Selector) (ret []*v1beta1.CronJobTrigger, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1beta1.CronJobTrigger))
	})
	return ret, err
}

// Get retrieves the CronJobTrigger from the indexer for a given namespace and name.
func (s cronJobTriggerNamespaceLister) Get(name string) (*v1beta1.CronJobTrigger, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1beta1.Resource("cronjobtrigger"), name)
	}
	return obj.(*v1beta1.CronJobTrigger), nil
}
