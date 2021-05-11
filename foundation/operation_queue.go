package foundation

// #cgo CFLAGS: -x objective-c
// #cgo LDFLAGS: -framework Foundation
// #include "operation_queue.h"
import "C"
import (
	"github.com/hsiafan/cocoa/objc"
	"unsafe"
)

type OperationQueue interface {
	objc.Object
	AddOperation(op Operation)
	CancelAllOperations()
	WaitUntilAllOperationsAreFinished()
	MaxConcurrentOperationCount() int
	SetMaxConcurrentOperationCount(value int)
	IsSuspended() bool
	SetSuspended(value bool)
	Name() string
	SetName(value string)
	Progress() Progress
}

type NSOperationQueue struct {
	objc.NSObject
}

func MakeOperationQueue(ptr unsafe.Pointer) *NSOperationQueue {
	if ptr == nil {
		return nil
	}
	return &NSOperationQueue{
		NSObject: *objc.MakeObject(ptr),
	}
}

func AllocOperationQueue() *NSOperationQueue {
	return MakeOperationQueue(C.C_OperationQueue_Alloc())
}

func (n *NSOperationQueue) Init() OperationQueue {
	result := C.C_NSOperationQueue_Init(n.Ptr())
	return MakeOperationQueue(result)
}

func (n *NSOperationQueue) AddOperation(op Operation) {
	C.C_NSOperationQueue_AddOperation(n.Ptr(), objc.ExtractPtr(op))
}

func (n *NSOperationQueue) CancelAllOperations() {
	C.C_NSOperationQueue_CancelAllOperations(n.Ptr())
}

func (n *NSOperationQueue) WaitUntilAllOperationsAreFinished() {
	C.C_NSOperationQueue_WaitUntilAllOperationsAreFinished(n.Ptr())
}

func (n *NSOperationQueue) MaxConcurrentOperationCount() int {
	result := C.C_NSOperationQueue_MaxConcurrentOperationCount(n.Ptr())
	return int(result)
}

func (n *NSOperationQueue) SetMaxConcurrentOperationCount(value int) {
	C.C_NSOperationQueue_SetMaxConcurrentOperationCount(n.Ptr(), C.int(value))
}

func (n *NSOperationQueue) IsSuspended() bool {
	result := C.C_NSOperationQueue_IsSuspended(n.Ptr())
	return bool(result)
}

func (n *NSOperationQueue) SetSuspended(value bool) {
	C.C_NSOperationQueue_SetSuspended(n.Ptr(), C.bool(value))
}

func (n *NSOperationQueue) Name() string {
	result := C.C_NSOperationQueue_Name(n.Ptr())
	return MakeString(result).String()
}

func (n *NSOperationQueue) SetName(value string) {
	C.C_NSOperationQueue_SetName(n.Ptr(), NewString(value).Ptr())
}

func (n *NSOperationQueue) Progress() Progress {
	result := C.C_NSOperationQueue_Progress(n.Ptr())
	return MakeProgress(result)
}
