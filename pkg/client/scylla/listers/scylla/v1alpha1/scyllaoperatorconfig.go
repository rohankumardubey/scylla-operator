// Code generated by lister-gen. DO NOT EDIT.

package v1alpha1

import (
	v1alpha1 "github.com/scylladb/scylla-operator/pkg/api/scylla/v1alpha1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// ScyllaOperatorConfigLister helps list ScyllaOperatorConfigs.
// All objects returned here must be treated as read-only.
type ScyllaOperatorConfigLister interface {
	// List lists all ScyllaOperatorConfigs in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.ScyllaOperatorConfig, err error)
	// Get retrieves the ScyllaOperatorConfig from the index for a given name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.ScyllaOperatorConfig, error)
	ScyllaOperatorConfigListerExpansion
}

// scyllaOperatorConfigLister implements the ScyllaOperatorConfigLister interface.
type scyllaOperatorConfigLister struct {
	indexer cache.Indexer
}

// NewScyllaOperatorConfigLister returns a new ScyllaOperatorConfigLister.
func NewScyllaOperatorConfigLister(indexer cache.Indexer) ScyllaOperatorConfigLister {
	return &scyllaOperatorConfigLister{indexer: indexer}
}

// List lists all ScyllaOperatorConfigs in the indexer.
func (s *scyllaOperatorConfigLister) List(selector labels.Selector) (ret []*v1alpha1.ScyllaOperatorConfig, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.ScyllaOperatorConfig))
	})
	return ret, err
}

// Get retrieves the ScyllaOperatorConfig from the index for a given name.
func (s *scyllaOperatorConfigLister) Get(name string) (*v1alpha1.ScyllaOperatorConfig, error) {
	obj, exists, err := s.indexer.GetByKey(name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("scyllaoperatorconfig"), name)
	}
	return obj.(*v1alpha1.ScyllaOperatorConfig), nil
}
