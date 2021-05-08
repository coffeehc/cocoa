package objc

// #cgo CFLAGS: -x objective-c
// #cgo LDFLAGS: -framework Foundation
// #import "base.h"
import "C"

// DispatchAsyncToMainQueue schedule task to run in main thread async, it does the same thing as
// 	dispatch_async(dispatch_get_main_queue(), ^{
// 	})
// This method can be used to do ui operations in UI thread, for appkit etc..
func DispatchAsyncToMainQueue(task func()) {
	taskLock.Lock()
	currentTaskId++
	id := currentTaskId
	tasks[id] = task
	taskLock.Unlock()
	C.Dispatch_MainQueueAsync(C.long(id))
}

// WithAutoreleasePool run code in a new auto release pool.
func WithAutoreleasePool(task func()) {
	taskLock.Lock()
	currentTaskId++
	id := currentTaskId
	tasks[id] = task
	taskLock.Unlock()
	C.Run_WithAutoreleasePool(C.long(id))
}

//export runTask
func runTask(id int64) {

	taskLock.RLock()
	task := tasks[id]
	taskLock.RUnlock()

	task()

	taskLock.Lock()
	delete(tasks, id)
	taskLock.Unlock()
}
