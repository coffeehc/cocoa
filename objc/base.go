package objc

// #cgo CFLAGS: -x objective-c
// #cgo LDFLAGS: -framework Foundation
// #import "base.h"
import "C"
import "runtime/cgo"

// DispatchAsyncToMainQueue schedule task to run in main thread async, it does the same thing as
// 	dispatch_async(dispatch_get_main_queue(), ^{
// 	})
// This method can be used to do ui operations in UI thread, for appkit etc..
func DispatchAsyncToMainQueue(task func()) {
	id := cgo.NewHandle(task)
	C.Dispatch_MainQueueAsync(C.uintptr_t(id))
}

// WithAutoreleasePool run code in a new auto release pool.
func WithAutoreleasePool(task func()) {
	id := cgo.NewHandle(task)
	C.Run_WithAutoreleasePool(C.uintptr_t(id))
}

//export runTask
func runTask(p C.uintptr_t) {
	h := cgo.Handle(p)
	task := h.Value().(func())
	task()
	h.Delete()
}
