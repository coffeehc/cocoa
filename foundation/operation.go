package foundation

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
	Dependencies() []Operation
	QualityOfService() QualityOfService
	SetQualityOfService(value QualityOfService)
	QueuePriority() OperationQueuePriority
	SetQueuePriority(value OperationQueuePriority)
}

type NSOperation struct {
	objc.NSObject
}

func MakeOperation(ptr unsafe.Pointer) NSOperation {
	return NSOperation{
		NSObject: objc.MakeObject(ptr),
	}
}

func AllocOperation() NSOperation {
	return MakeOperation(C.C_Operation_Alloc())
}

func (n NSOperation) Init() Operation {
	result_ := C.C_NSOperation_Init(n.Ptr())
	return MakeOperation(result_)
}

func (n NSOperation) Start() {
	C.C_NSOperation_Start(n.Ptr())
}

func (n NSOperation) Main() {
	C.C_NSOperation_Main(n.Ptr())
}

func (n NSOperation) Cancel() {
	C.C_NSOperation_Cancel(n.Ptr())
}

func (n NSOperation) AddDependency(op Operation) {
	C.C_NSOperation_AddDependency(n.Ptr(), objc.ExtractPtr(op))
}

func (n NSOperation) RemoveDependency(op Operation) {
	C.C_NSOperation_RemoveDependency(n.Ptr(), objc.ExtractPtr(op))
}

func (n NSOperation) WaitUntilFinished() {
	C.C_NSOperation_WaitUntilFinished(n.Ptr())
}

func (n NSOperation) IsCancelled() bool {
	result_ := C.C_NSOperation_IsCancelled(n.Ptr())
	return bool(result_)
}

func (n NSOperation) IsExecuting() bool {
	result_ := C.C_NSOperation_IsExecuting(n.Ptr())
	return bool(result_)
}

func (n NSOperation) IsFinished() bool {
	result_ := C.C_NSOperation_IsFinished(n.Ptr())
	return bool(result_)
}

func (n NSOperation) IsConcurrent() bool {
	result_ := C.C_NSOperation_IsConcurrent(n.Ptr())
	return bool(result_)
}

func (n NSOperation) IsAsynchronous() bool {
	result_ := C.C_NSOperation_IsAsynchronous(n.Ptr())
	return bool(result_)
}

func (n NSOperation) IsReady() bool {
	result_ := C.C_NSOperation_IsReady(n.Ptr())
	return bool(result_)
}

func (n NSOperation) Name() string {
	result_ := C.C_NSOperation_Name(n.Ptr())
	return MakeString(result_).String()
}

func (n NSOperation) SetName(value string) {
	C.C_NSOperation_SetName(n.Ptr(), NewString(value).Ptr())
}

func (n NSOperation) Dependencies() []Operation {
	result_ := C.C_NSOperation_Dependencies(n.Ptr())
	if result_.len > 0 {
		defer C.free(result_.data)
	}
	result_Slice := (*[1 << 28]unsafe.Pointer)(unsafe.Pointer(result_.data))[:result_.len:result_.len]
	var goResult_ = make([]Operation, len(result_Slice))
	for idx, r := range result_Slice {
		goResult_[idx] = MakeOperation(r)
	}
	return goResult_
}

func (n NSOperation) QualityOfService() QualityOfService {
	result_ := C.C_NSOperation_QualityOfService(n.Ptr())
	return QualityOfService(int(result_))
}

func (n NSOperation) SetQualityOfService(value QualityOfService) {
	C.C_NSOperation_SetQualityOfService(n.Ptr(), C.int(int(value)))
}

func (n NSOperation) QueuePriority() OperationQueuePriority {
	result_ := C.C_NSOperation_QueuePriority(n.Ptr())
	return OperationQueuePriority(int(result_))
}

func (n NSOperation) SetQueuePriority(value OperationQueuePriority) {
	C.C_NSOperation_SetQueuePriority(n.Ptr(), C.int(int(value)))
}
