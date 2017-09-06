// This file was automatically generated by informer-gen

package internalversion

import (
	network "github.com/openshift/origin/pkg/network/apis/network"
	internalinterfaces "github.com/openshift/origin/pkg/network/generated/informers/internalversion/internalinterfaces"
	internalclientset "github.com/openshift/origin/pkg/network/generated/internalclientset"
	internalversion "github.com/openshift/origin/pkg/network/generated/listers/network/internalversion"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
	time "time"
)

// ClusterNetworkInformer provides access to a shared informer and lister for
// ClusterNetworks.
type ClusterNetworkInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() internalversion.ClusterNetworkLister
}

type clusterNetworkInformer struct {
	factory internalinterfaces.SharedInformerFactory
}

func newClusterNetworkInformer(client internalclientset.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	sharedIndexInformer := cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				return client.Network().ClusterNetworks().List(options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				return client.Network().ClusterNetworks().Watch(options)
			},
		},
		&network.ClusterNetwork{},
		resyncPeriod,
		cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc},
	)

	return sharedIndexInformer
}

func (f *clusterNetworkInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&network.ClusterNetwork{}, newClusterNetworkInformer)
}

func (f *clusterNetworkInformer) Lister() internalversion.ClusterNetworkLister {
	return internalversion.NewClusterNetworkLister(f.Informer().GetIndexer())
}
