package foundation

import (
	"github.com/hsiafan/cocoa/objc"
	"sync"
	"sync/atomic"
	"unsafe"
)

type Resource interface {
	Ptr() unsafe.Pointer
	Dealloc()
}

// ResourceRegistry maintains resources, and their ids.
type ResourceRegistry struct {
	currentId int64
	resources map[int64]interface{}
	lock      sync.RWMutex
}

// NewResourceRegistry create new ResourceRegistry
func NewResourceRegistry() *ResourceRegistry {
	return &ResourceRegistry{
		resources: make(map[int64]interface{}),
	}
}

// NextId return next id for resource
func (r *ResourceRegistry) NextId() int64 {
	return atomic.AddInt64(&r.currentId, 1)
}

// Register register resource, and return it's id
func (r *ResourceRegistry) Register(obj Resource) int64 {
	id := r.NextId()
	r.lock.Lock()
	if or, ok := r.resources[id]; ok {
		if or.(Resource).Ptr() != obj.Ptr() {
			panic("resource with this id already exists")
		}
		return id
	}
	r.resources[id] = obj
	r.lock.Unlock()
	objc.AddDeallocHook(obj, func() {
		r.Delete(id)
	})
	return id
}

// Store store resource with it's id
func (r *ResourceRegistry) Store(id int64, obj interface{}) {
	r.lock.Lock()
	defer r.lock.Unlock()
	if _, ok := r.resources[id]; ok {
		panic("resource with this id already exists")
	}
	r.resources[id] = obj

}

// Get get resource by id
func (r *ResourceRegistry) Get(id int64) interface{} {
	r.lock.RLock()
	defer r.lock.RUnlock()
	return r.resources[id]
}

// Delete delete resource by id
func (r *ResourceRegistry) Delete(id int64) interface{} {
	r.lock.Lock()
	defer r.lock.Unlock()
	if _, ok := r.resources[id]; !ok {
		panic("resource with this id not exists")
	}
	obj := r.resources[id]
	delete(r.resources, id)
	return obj
}
