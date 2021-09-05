/*
MIT License

Copyright (c) 2018 Amadeus s.a.s.

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
*/

// Code generated by lister-gen. DO NOT EDIT.

package v1

import (
	v1 "github.com/bjosv/redis-operator/pkg/api/redis/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// RedisClusterLister helps list RedisClusters.
// All objects returned here must be treated as read-only.
type RedisClusterLister interface {
	// List lists all RedisClusters in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1.RedisCluster, err error)
	// RedisClusters returns an object that can list and get RedisClusters.
	RedisClusters(namespace string) RedisClusterNamespaceLister
	RedisClusterListerExpansion
}

// redisClusterLister implements the RedisClusterLister interface.
type redisClusterLister struct {
	indexer cache.Indexer
}

// NewRedisClusterLister returns a new RedisClusterLister.
func NewRedisClusterLister(indexer cache.Indexer) RedisClusterLister {
	return &redisClusterLister{indexer: indexer}
}

// List lists all RedisClusters in the indexer.
func (s *redisClusterLister) List(selector labels.Selector) (ret []*v1.RedisCluster, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.RedisCluster))
	})
	return ret, err
}

// RedisClusters returns an object that can list and get RedisClusters.
func (s *redisClusterLister) RedisClusters(namespace string) RedisClusterNamespaceLister {
	return redisClusterNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// RedisClusterNamespaceLister helps list and get RedisClusters.
// All objects returned here must be treated as read-only.
type RedisClusterNamespaceLister interface {
	// List lists all RedisClusters in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1.RedisCluster, err error)
	// Get retrieves the RedisCluster from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1.RedisCluster, error)
	RedisClusterNamespaceListerExpansion
}

// redisClusterNamespaceLister implements the RedisClusterNamespaceLister
// interface.
type redisClusterNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all RedisClusters in the indexer for a given namespace.
func (s redisClusterNamespaceLister) List(selector labels.Selector) (ret []*v1.RedisCluster, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.RedisCluster))
	})
	return ret, err
}

// Get retrieves the RedisCluster from the indexer for a given namespace and name.
func (s redisClusterNamespaceLister) Get(name string) (*v1.RedisCluster, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1.Resource("rediscluster"), name)
	}
	return obj.(*v1.RedisCluster), nil
}
