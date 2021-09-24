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
	result_ := C.C_NSOperation_AllocOperation()
	return MakeOperation(result_)
}

func (n NSOperation) Init() NSOperation {
	result_ := C.C_NSOperation_Init(n.Ptr())
	return MakeOperation(result_)
}

func NewOperation() NSOperation {
	result_ := C.C_NSOperation_NewOperation()
	return MakeOperation(result_)
}

func (n NSOperation) Autorelease() NSOperation {
	result_ := C.C_NSOperation_Autorelease(n.Ptr())
	return MakeOperation(result_)
}

func (n NSOperation) Retain() NSOperation {
	result_ := C.C_NSOperation_Retain(n.Ptr())
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
	result_Slice := unsafe.Slice((*unsafe.Pointer)(result_.data), int(result_.len))
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
	result_ := C.C_NSOperationQueue_AllocOperationQueue()
	return MakeOperationQueue(result_)
}

func (n NSOperationQueue) Init() NSOperationQueue {
	result_ := C.C_NSOperationQueue_Init(n.Ptr())
	return MakeOperationQueue(result_)
}

func NewOperationQueue() NSOperationQueue {
	result_ := C.C_NSOperationQueue_NewOperationQueue()
	return MakeOperationQueue(result_)
}

func (n NSOperationQueue) Autorelease() NSOperationQueue {
	result_ := C.C_NSOperationQueue_Autorelease(n.Ptr())
	return MakeOperationQueue(result_)
}

func (n NSOperationQueue) Retain() NSOperationQueue {
	result_ := C.C_NSOperationQueue_Retain(n.Ptr())
	return MakeOperationQueue(result_)
}

func (n NSOperationQueue) AddOperation(op Operation) {
	C.C_NSOperationQueue_AddOperation(n.Ptr(), objc.ExtractPtr(op))
}

func (n NSOperationQueue) AddOperations_WaitUntilFinished(ops []Operation, wait bool) {
	var cOps C.Array
	if len(ops) > 0 {
		cOpsData := make([]unsafe.Pointer, len(ops))
		for idx, v := range ops {
			cOpsData[idx] = objc.ExtractPtr(v)
		}
		cOps.data = unsafe.Pointer(&cOpsData[0])
		cOps.len = C.int(len(ops))
	}
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
