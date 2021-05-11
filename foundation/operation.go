package foundation

// #cgo CFLAGS: -x objective-c
// #cgo LDFLAGS: -framework Foundation
// #include "operation.h"
import "C"
import (
	"github.com/hsiafan/cocoa/objc"
	"unsafe"
)

type Operation interface {
	objc.Object
	Start()
	Main()
	Cancel()
	AddDependency(op Operation)
	RemoveDependency(op Operation)
	WaitUntilFinished()
	IsCancelled() bool
	IsExecuting() bool
	IsFinished() bool
	IsConcurrent() bool
	IsAsynchronous() bool
	IsReady() bool
	Name() string
	SetName(value string)
}

type NSOperation struct {
	objc.NSObject
}

func MakeOperation(ptr unsafe.Pointer) *NSOperation {
	if ptr == nil {
		return nil
	}
	return &NSOperation{
		NSObject: *objc.MakeObject(ptr),
	}
}

func AllocOperation() *NSOperation {
	return MakeOperation(C.C_Operation_Alloc())
}

func (n *NSOperation) Init() Operation {
	result := C.C_NSOperation_Init(n.Ptr())
	return MakeOperation(result)
}

func (n *NSOperation) Start() {
	C.C_NSOperation_Start(n.Ptr())
}

func (n *NSOperation) Main() {
	C.C_NSOperation_Main(n.Ptr())
}

func (n *NSOperation) Cancel() {
	C.C_NSOperation_Cancel(n.Ptr())
}

func (n *NSOperation) AddDependency(op Operation) {
	C.C_NSOperation_AddDependency(n.Ptr(), objc.ExtractPtr(op))
}

func (n *NSOperation) RemoveDependency(op Operation) {
	C.C_NSOperation_RemoveDependency(n.Ptr(), objc.ExtractPtr(op))
}

func (n *NSOperation) WaitUntilFinished() {
	C.C_NSOperation_WaitUntilFinished(n.Ptr())
}

func (n *NSOperation) IsCancelled() bool {
	result := C.C_NSOperation_IsCancelled(n.Ptr())
	return bool(result)
}

func (n *NSOperation) IsExecuting() bool {
	result := C.C_NSOperation_IsExecuting(n.Ptr())
	return bool(result)
}

func (n *NSOperation) IsFinished() bool {
	result := C.C_NSOperation_IsFinished(n.Ptr())
	return bool(result)
}

func (n *NSOperation) IsConcurrent() bool {
	result := C.C_NSOperation_IsConcurrent(n.Ptr())
	return bool(result)
}

func (n *NSOperation) IsAsynchronous() bool {
	result := C.C_NSOperation_IsAsynchronous(n.Ptr())
	return bool(result)
}

func (n *NSOperation) IsReady() bool {
	result := C.C_NSOperation_IsReady(n.Ptr())
	return bool(result)
}

func (n *NSOperation) Name() string {
	result := C.C_NSOperation_Name(n.Ptr())
	return MakeString(result).String()
}

func (n *NSOperation) SetName(value string) {
	C.C_NSOperation_SetName(n.Ptr(), NewString(value).Ptr())
}
