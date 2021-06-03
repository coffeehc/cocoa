package objc

// #import "object.h"
import "C"
import (
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
	if o == nil {
		return nil
	}
	return o.Ptr()
}

// NSObject is wrapper for objc-NSObject
type NSObject struct {
	ptr unsafe.Pointer
}

func (o NSObject) Dealloc() {
	if o.Ptr() == nil {
		panic("objc object is null")
	}
	C.Object_Dealloc(o.Ptr())
}

func (o NSObject) Ptr() unsafe.Pointer {
	return o.ptr
}

func MakeObject(ptr unsafe.Pointer) NSObject {
	return NSObject{ptr}
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
