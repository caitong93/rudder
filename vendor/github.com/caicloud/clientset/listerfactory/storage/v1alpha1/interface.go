/*
Copyright 2018 caicloud authors. All rights reserved.
*/

// Code generated by listerfactory-gen. DO NOT EDIT.

package v1alpha1

import (
	internalinterfaces "github.com/caicloud/clientset/listerfactory/internalinterfaces"
	informers "k8s.io/client-go/informers"
	kubernetes "k8s.io/client-go/kubernetes"
	v1alpha1 "k8s.io/client-go/listers/storage/v1alpha1"
)

// Interface provides access to all the listers in this group version.
type Interface interface { // VolumeAttachments returns a VolumeAttachmentLister
	VolumeAttachments() v1alpha1.VolumeAttachmentLister
}

type version struct {
	client           kubernetes.Interface
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

type infromerVersion struct {
	factory informers.SharedInformerFactory
}

// New returns a new Interface.
func New(client kubernetes.Interface, tweakListOptions internalinterfaces.TweakListOptionsFunc) Interface {
	return &version{client: client, tweakListOptions: tweakListOptions}
}

// NewFrom returns a new Interface.
func NewFrom(factory informers.SharedInformerFactory) Interface {
	return &infromerVersion{factory: factory}
}

// VolumeAttachments returns a VolumeAttachmentLister.
func (v *version) VolumeAttachments() v1alpha1.VolumeAttachmentLister {
	return &volumeAttachmentLister{client: v.client, tweakListOptions: v.tweakListOptions}
}

// VolumeAttachments returns a VolumeAttachmentLister.
func (v *infromerVersion) VolumeAttachments() v1alpha1.VolumeAttachmentLister {
	return v.factory.Storage().V1alpha1().VolumeAttachments().Lister()
}
