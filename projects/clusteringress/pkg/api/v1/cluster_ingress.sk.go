// Code generated by solo-kit. DO NOT EDIT.

package v1

import (
	"sort"

	"github.com/solo-io/go-utils/hashutils"
	"github.com/solo-io/solo-kit/pkg/api/v1/clients/kube/crd"
	"github.com/solo-io/solo-kit/pkg/api/v1/resources"
	"github.com/solo-io/solo-kit/pkg/api/v1/resources/core"
	"github.com/solo-io/solo-kit/pkg/errors"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

func NewClusterIngress(namespace, name string) *ClusterIngress {
	clusteringress := &ClusterIngress{}
	clusteringress.SetMetadata(core.Metadata{
		Name:      name,
		Namespace: namespace,
	})
	return clusteringress
}

func (r *ClusterIngress) SetMetadata(meta core.Metadata) {
	r.Metadata = meta
}

func (r *ClusterIngress) SetStatus(status core.Status) {
	r.Status = status
}

func (r *ClusterIngress) Hash() uint64 {
	metaCopy := r.GetMetadata()
	metaCopy.ResourceVersion = ""
	return hashutils.HashAll(
		metaCopy,
		r.ClusterIngressSpec,
	)
}

type ClusterIngressList []*ClusterIngress

// namespace is optional, if left empty, names can collide if the list contains more than one with the same name
func (list ClusterIngressList) Find(namespace, name string) (*ClusterIngress, error) {
	for _, clusterIngress := range list {
		if clusterIngress.GetMetadata().Name == name {
			if namespace == "" || clusterIngress.GetMetadata().Namespace == namespace {
				return clusterIngress, nil
			}
		}
	}
	return nil, errors.Errorf("list did not find clusterIngress %v.%v", namespace, name)
}

func (list ClusterIngressList) AsResources() resources.ResourceList {
	var ress resources.ResourceList
	for _, clusterIngress := range list {
		ress = append(ress, clusterIngress)
	}
	return ress
}

func (list ClusterIngressList) AsInputResources() resources.InputResourceList {
	var ress resources.InputResourceList
	for _, clusterIngress := range list {
		ress = append(ress, clusterIngress)
	}
	return ress
}

func (list ClusterIngressList) Names() []string {
	var names []string
	for _, clusterIngress := range list {
		names = append(names, clusterIngress.GetMetadata().Name)
	}
	return names
}

func (list ClusterIngressList) NamespacesDotNames() []string {
	var names []string
	for _, clusterIngress := range list {
		names = append(names, clusterIngress.GetMetadata().Namespace+"."+clusterIngress.GetMetadata().Name)
	}
	return names
}

func (list ClusterIngressList) Sort() ClusterIngressList {
	sort.SliceStable(list, func(i, j int) bool {
		return list[i].GetMetadata().Less(list[j].GetMetadata())
	})
	return list
}

func (list ClusterIngressList) Clone() ClusterIngressList {
	var clusterIngressList ClusterIngressList
	for _, clusterIngress := range list {
		clusterIngressList = append(clusterIngressList, resources.Clone(clusterIngress).(*ClusterIngress))
	}
	return clusterIngressList
}

func (list ClusterIngressList) Each(f func(element *ClusterIngress)) {
	for _, clusterIngress := range list {
		f(clusterIngress)
	}
}

func (list ClusterIngressList) EachResource(f func(element resources.Resource)) {
	for _, clusterIngress := range list {
		f(clusterIngress)
	}
}

func (list ClusterIngressList) AsInterfaces() []interface{} {
	var asInterfaces []interface{}
	list.Each(func(element *ClusterIngress) {
		asInterfaces = append(asInterfaces, element)
	})
	return asInterfaces
}

var _ resources.Resource = &ClusterIngress{}

// Kubernetes Adapter for ClusterIngress

func (o *ClusterIngress) GetObjectKind() schema.ObjectKind {
	t := ClusterIngressCrd.TypeMeta()
	return &t
}

func (o *ClusterIngress) DeepCopyObject() runtime.Object {
	return resources.Clone(o).(*ClusterIngress)
}

var ClusterIngressCrd = crd.NewCrd("clusteringress.gloo.solo.io",
	"clusteringresses",
	"clusteringress.gloo.solo.io",
	"v1",
	"ClusterIngress",
	"cig",
	true,
	&ClusterIngress{})
