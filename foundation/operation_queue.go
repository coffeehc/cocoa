package foundation

// #include "operation_queue.h"
import "C"
import (
	"github.com/hsiafan/cocoa/objc"
	"unsafe"
)

type OperationQueue interface {
	objc.Object
	AddOperation(op Operation)
	AddOperations_WaitUntilFinished(ops []Operation, wait bool)
	CancelAllOperations()
	WaitUntilAllOperationsAreFinished()
	QualityOfService() QualityOfService
	SetQualityOfService(value QualityOfService)
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

func MakeOperationQueue(ptr unsafe.Pointer) NSOperationQueue {
	return NSOperationQueue{
		NSObject: objc.MakeObject(ptr),
	}
}

func AllocOperationQueue() NSOperationQueue {
	return MakeOperationQueue(C.C_OperationQueue_Alloc())
}

func (n NSOperationQueue) Init() OperationQueue {
	result_ := C.C_NSOperationQueue_Init(n.Ptr())
	return MakeOperationQueue(result_)
}

func (n NSOperationQueue) AddOperation(op Operation) {
	C.C_NSOperationQueue_AddOperation(n.Ptr(), objc.ExtractPtr(op))
}

func (n NSOperationQueue) AddOperations_WaitUntilFinished(ops []Operation, wait bool) {
	cOpsData := make([]unsafe.Pointer, len(ops))
	for idx, v := range ops {
		cOpsData[idx] = objc.ExtractPtr(v)
	}
	cOps := C.Array{data: unsafe.Pointer(&cOpsData[0]), len: C.int(len(ops))}
	C.C_NSOperationQueue_AddOperations_WaitUntilFinished(n.Ptr(), cOps, C.bool(wait))
}

func (n NSOperationQueue) CancelAllOperations() {
	C.C_NSOperationQueue_CancelAllOperations(n.Ptr())
}

func (n NSOperationQueue) WaitUntilAllOperationsAreFinished() {
	C.C_NSOperationQueue_WaitUntilAllOperationsAreFinished(n.Ptr())
}

func OperationQueue_MainQueue() OperationQueue {
	result_ := C.C_NSOperationQueue_OperationQueue_MainQueue()
	return MakeOperationQueue(result_)
}

func OperationQueue_CurrentQueue() OperationQueue {
	result_ := C.C_NSOperationQueue_OperationQueue_CurrentQueue()
	return MakeOperationQueue(result_)
}

func (n NSOperationQueue) QualityOfService() QualityOfService {
	result_ := C.C_NSOperationQueue_QualityOfService(n.Ptr())
	return QualityOfService(int(result_))
}

func (n NSOperationQueue) SetQualityOfService(value QualityOfService) {
	C.C_NSOperationQueue_SetQualityOfService(n.Ptr(), C.int(int(value)))
}

func (n NSOperationQueue) MaxConcurrentOperationCount() int {
	result_ := C.C_NSOperationQueue_MaxConcurrentOperationCount(n.Ptr())
	return int(result_)
}

func (n NSOperationQueue) SetMaxConcurrentOperationCount(value int) {
	C.C_NSOperationQueue_SetMaxConcurrentOperationCount(n.Ptr(), C.int(value))
}

func (n NSOperationQueue) IsSuspended() bool {
	result_ := C.C_NSOperationQueue_IsSuspended(n.Ptr())
	return bool(result_)
}

func (n NSOperationQueue) SetSuspended(value bool) {
	C.C_NSOperationQueue_SetSuspended(n.Ptr(), C.bool(value))
}

func (n NSOperationQueue) Name() string {
	result_ := C.C_NSOperationQueue_Name(n.Ptr())
	return MakeString(result_).String()
}

func (n NSOperationQueue) SetName(value string) {
	C.C_NSOperationQueue_SetName(n.Ptr(), NewString(value).Ptr())
}

func (n NSOperationQueue) Progress() Progress {
	result_ := C.C_NSOperationQueue_Progress(n.Ptr())
	return MakeProgress(result_)
}
