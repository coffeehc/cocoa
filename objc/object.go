package objc

// #cgo CFLAGS: -x objective-c
// #cgo LDFLAGS: -framework Foundation
// #import "object.h"
import "C"
import (
	"github.com/hsiafan/cocoa/internal/utils"
	"sync"
	"unsafe"
)

// PointerHolder is a interface for holding a objc pointer
type PointerHolder interface {
	// Ptr return the delegate objc objc pointer
	Ptr() unsafe.Pointer
}

// Object is interface for all objc NSObject type
type Object interface {
	PointerHolder
	Dealloc()
}

// ExtractPtr return the objc ptr hold by Object. If is nil, or contains a nil pointer, return nil
func ExtractPtr(o PointerHolder) unsafe.Pointer {
	if utils.InterfaceIsNil(o) {
		return nil
	}
	return o.Ptr()
}

var _ Object = (*NSObject)(nil)

// NSObject is wrapper for objc-NSObject
type NSObject struct {
	ptr unsafe.Pointer
}

func (o *NSObject) Dealloc() {
	panic("to be implemented")
}

func (o *NSObject) Ptr() unsafe.Pointer {
	return o.ptr
}

func MakeObject(ptr unsafe.Pointer) *NSObject {
	if ptr == nil {
		return nil
	}
	return &NSObject{ptr}
}

var tasks = make(map[int64]func())
var taskLock sync.RWMutex
var currentTaskId int64

// AddDeallocHook add cocoa object dealloc hook
func AddDeallocHook(obj Object, hook func()) {
	if obj.Ptr() == nil {
		panic("cocoa pointer is nil")
	}
	taskLock.Lock()
	currentTaskId++
	id := currentTaskId
	tasks[id] = hook
	taskLock.Unlock()
	C.Dealloc_AddHook(obj.Ptr(), C.long(id))
}

//export runDeallocTask
func runDeallocTask(id int64) {

	taskLock.RLock()
	task := tasks[id]
	taskLock.RUnlock()

	task()

	taskLock.Lock()
	delete(tasks, id)
	taskLock.Unlock()
}
