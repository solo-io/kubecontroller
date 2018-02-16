package kubecontroller

import "k8s.io/client-go/tools/cache"

// returns a handler that runs f() every time an update occurs,
// regardless of which type of update
func NewSyncHandler(f func()) cache.ResourceEventHandler {
	return &cache.ResourceEventHandlerFuncs{
		AddFunc: func(obj interface{}) {
			f()
		},
		UpdateFunc: func(oldObj, newObj interface{}) {
			f()
		},
		DeleteFunc: func(obj interface{}) {
			f()
		},
	}
}
