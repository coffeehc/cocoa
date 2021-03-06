package foundation

// #include "progress.h"
import "C"
import (
	"github.com/hsiafan/cocoa/objc"
	"unsafe"
)

type Progress interface {
	objc.Object
	BecomeCurrentWithPendingUnitCount(unitCount int64)
	AddChild_WithPendingUnitCount(child Progress, inUnitCount int64)
	ResignCurrent()
	Cancel()
	Pause()
	Resume()
	SetUserInfoObject_ForKey(objectOrNil objc.Object, key ProgressUserInfoKey)
	Publish()
	Unpublish()
	TotalUnitCount() int64
	SetTotalUnitCount(value int64)
	CompletedUnitCount() int64
	SetCompletedUnitCount(value int64)
	LocalizedDescription() string
	SetLocalizedDescription(value string)
	LocalizedAdditionalDescription() string
	SetLocalizedAdditionalDescription(value string)
	FractionCompleted() float64
	IsCancellable() bool
	SetCancellable(value bool)
	IsCancelled() bool
	IsPausable() bool
	SetPausable(value bool)
	IsPaused() bool
	IsIndeterminate() bool
	Kind() ProgressKind
	SetKind(value ProgressKind)
	UserInfo() map[ProgressUserInfoKey]objc.Object
	FileOperationKind() ProgressFileOperationKind
	SetFileOperationKind(value ProgressFileOperationKind)
	FileURL() URL
	SetFileURL(value URL)
	IsFinished() bool
	IsOld() bool
	EstimatedTimeRemaining() Number
	SetEstimatedTimeRemaining(value Number)
	FileCompletedCount() Number
	SetFileCompletedCount(value Number)
	FileTotalCount() Number
	SetFileTotalCount(value Number)
	Throughput() Number
	SetThroughput(value Number)
}

type NSProgress struct {
	objc.NSObject
}

func MakeProgress(ptr unsafe.Pointer) NSProgress {
	return NSProgress{
		NSObject: objc.MakeObject(ptr),
	}
}

func (n NSProgress) InitWithParent_UserInfo(parentProgressOrNil Progress, userInfoOrNil map[ProgressUserInfoKey]objc.Object) NSProgress {
	var cUserInfoOrNil C.Dictionary
	if len(userInfoOrNil) > 0 {
		cUserInfoOrNilKeyData := make([]unsafe.Pointer, len(userInfoOrNil))
		cUserInfoOrNilValueData := make([]unsafe.Pointer, len(userInfoOrNil))
		var idx = 0
		for k, v := range userInfoOrNil {
			cUserInfoOrNilKeyData[idx] = NewString(string(k)).Ptr()
			cUserInfoOrNilValueData[idx] = objc.ExtractPtr(v)
			idx++
		}
		cUserInfoOrNil.key_data = unsafe.Pointer(&cUserInfoOrNilKeyData[0])
		cUserInfoOrNil.value_data = unsafe.Pointer(&cUserInfoOrNilValueData[0])
		cUserInfoOrNil.len = C.int(len(userInfoOrNil))
	}
	result_ := C.C_NSProgress_InitWithParent_UserInfo(n.Ptr(), objc.ExtractPtr(parentProgressOrNil), cUserInfoOrNil)
	return MakeProgress(result_)
}

func AllocProgress() NSProgress {
	result_ := C.C_NSProgress_AllocProgress()
	return MakeProgress(result_)
}

func (n NSProgress) Init() NSProgress {
	result_ := C.C_NSProgress_Init(n.Ptr())
	return MakeProgress(result_)
}

func NewProgress() NSProgress {
	result_ := C.C_NSProgress_NewProgress()
	return MakeProgress(result_)
}

func (n NSProgress) Autorelease() NSProgress {
	result_ := C.C_NSProgress_Autorelease(n.Ptr())
	return MakeProgress(result_)
}

func (n NSProgress) Retain() NSProgress {
	result_ := C.C_NSProgress_Retain(n.Ptr())
	return MakeProgress(result_)
}

func DiscreteProgressWithTotalUnitCount(unitCount int64) Progress {
	result_ := C.C_NSProgress_DiscreteProgressWithTotalUnitCount(C.long(unitCount))
	return MakeProgress(result_)
}

func ProgressWithTotalUnitCount(unitCount int64) Progress {
	result_ := C.C_NSProgress_ProgressWithTotalUnitCount(C.long(unitCount))
	return MakeProgress(result_)
}

func ProgressWithTotalUnitCount_Parent_PendingUnitCount(unitCount int64, parent Progress, portionOfParentTotalUnitCount int64) Progress {
	result_ := C.C_NSProgress_ProgressWithTotalUnitCount_Parent_PendingUnitCount(C.long(unitCount), objc.ExtractPtr(parent), C.long(portionOfParentTotalUnitCount))
	return MakeProgress(result_)
}

func CurrentProgress() Progress {
	result_ := C.C_NSProgress_CurrentProgress()
	return MakeProgress(result_)
}

func (n NSProgress) BecomeCurrentWithPendingUnitCount(unitCount int64) {
	C.C_NSProgress_BecomeCurrentWithPendingUnitCount(n.Ptr(), C.long(unitCount))
}

func (n NSProgress) AddChild_WithPendingUnitCount(child Progress, inUnitCount int64) {
	C.C_NSProgress_AddChild_WithPendingUnitCount(n.Ptr(), objc.ExtractPtr(child), C.long(inUnitCount))
}

func (n NSProgress) ResignCurrent() {
	C.C_NSProgress_ResignCurrent(n.Ptr())
}

func (n NSProgress) Cancel() {
	C.C_NSProgress_Cancel(n.Ptr())
}

func (n NSProgress) Pause() {
	C.C_NSProgress_Pause(n.Ptr())
}

func (n NSProgress) Resume() {
	C.C_NSProgress_Resume(n.Ptr())
}

func (n NSProgress) SetUserInfoObject_ForKey(objectOrNil objc.Object, key ProgressUserInfoKey) {
	C.C_NSProgress_SetUserInfoObject_ForKey(n.Ptr(), objc.ExtractPtr(objectOrNil), NewString(string(key)).Ptr())
}

func (n NSProgress) Publish() {
	C.C_NSProgress_Publish(n.Ptr())
}

func (n NSProgress) Unpublish() {
	C.C_NSProgress_Unpublish(n.Ptr())
}

func Progress_RemoveSubscriber(subscriber objc.Object) {
	C.C_NSProgress_Progress_RemoveSubscriber(objc.ExtractPtr(subscriber))
}

func (n NSProgress) TotalUnitCount() int64 {
	result_ := C.C_NSProgress_TotalUnitCount(n.Ptr())
	return int64(result_)
}

func (n NSProgress) SetTotalUnitCount(value int64) {
	C.C_NSProgress_SetTotalUnitCount(n.Ptr(), C.long(value))
}

func (n NSProgress) CompletedUnitCount() int64 {
	result_ := C.C_NSProgress_CompletedUnitCount(n.Ptr())
	return int64(result_)
}

func (n NSProgress) SetCompletedUnitCount(value int64) {
	C.C_NSProgress_SetCompletedUnitCount(n.Ptr(), C.long(value))
}

func (n NSProgress) LocalizedDescription() string {
	result_ := C.C_NSProgress_LocalizedDescription(n.Ptr())
	return MakeString(result_).String()
}

func (n NSProgress) SetLocalizedDescription(value string) {
	C.C_NSProgress_SetLocalizedDescription(n.Ptr(), NewString(value).Ptr())
}

func (n NSProgress) LocalizedAdditionalDescription() string {
	result_ := C.C_NSProgress_LocalizedAdditionalDescription(n.Ptr())
	return MakeString(result_).String()
}

func (n NSProgress) SetLocalizedAdditionalDescription(value string) {
	C.C_NSProgress_SetLocalizedAdditionalDescription(n.Ptr(), NewString(value).Ptr())
}

func (n NSProgress) FractionCompleted() float64 {
	result_ := C.C_NSProgress_FractionCompleted(n.Ptr())
	return float64(result_)
}

func (n NSProgress) IsCancellable() bool {
	result_ := C.C_NSProgress_IsCancellable(n.Ptr())
	return bool(result_)
}

func (n NSProgress) SetCancellable(value bool) {
	C.C_NSProgress_SetCancellable(n.Ptr(), C.bool(value))
}

func (n NSProgress) IsCancelled() bool {
	result_ := C.C_NSProgress_IsCancelled(n.Ptr())
	return bool(result_)
}

func (n NSProgress) IsPausable() bool {
	result_ := C.C_NSProgress_IsPausable(n.Ptr())
	return bool(result_)
}

func (n NSProgress) SetPausable(value bool) {
	C.C_NSProgress_SetPausable(n.Ptr(), C.bool(value))
}

func (n NSProgress) IsPaused() bool {
	result_ := C.C_NSProgress_IsPaused(n.Ptr())
	return bool(result_)
}

func (n NSProgress) IsIndeterminate() bool {
	result_ := C.C_NSProgress_IsIndeterminate(n.Ptr())
	return bool(result_)
}

func (n NSProgress) Kind() ProgressKind {
	result_ := C.C_NSProgress_Kind(n.Ptr())
	return ProgressKind(MakeString(result_).String())
}

func (n NSProgress) SetKind(value ProgressKind) {
	C.C_NSProgress_SetKind(n.Ptr(), NewString(string(value)).Ptr())
}

func (n NSProgress) UserInfo() map[ProgressUserInfoKey]objc.Object {
	result_ := C.C_NSProgress_UserInfo(n.Ptr())
	if result_.len > 0 {
		defer C.free(result_.key_data)
		defer C.free(result_.value_data)
	}
	result_KeySlice := unsafe.Slice((*unsafe.Pointer)(result_.key_data), int(result_.len))
	result_ValueSlice := unsafe.Slice((*unsafe.Pointer)(result_.value_data), int(result_.len))
	var goResult_ = make(map[ProgressUserInfoKey]objc.Object)
	for idx, k := range result_KeySlice {
		v := result_ValueSlice[idx]
		goResult_[ProgressUserInfoKey(MakeString(k).String())] = objc.MakeObject(v)
	}
	return goResult_
}

func (n NSProgress) FileOperationKind() ProgressFileOperationKind {
	result_ := C.C_NSProgress_FileOperationKind(n.Ptr())
	return ProgressFileOperationKind(MakeString(result_).String())
}

func (n NSProgress) SetFileOperationKind(value ProgressFileOperationKind) {
	C.C_NSProgress_SetFileOperationKind(n.Ptr(), NewString(string(value)).Ptr())
}

func (n NSProgress) FileURL() URL {
	result_ := C.C_NSProgress_FileURL(n.Ptr())
	return MakeURL(result_)
}

func (n NSProgress) SetFileURL(value URL) {
	C.C_NSProgress_SetFileURL(n.Ptr(), objc.ExtractPtr(value))
}

func (n NSProgress) IsFinished() bool {
	result_ := C.C_NSProgress_IsFinished(n.Ptr())
	return bool(result_)
}

func (n NSProgress) IsOld() bool {
	result_ := C.C_NSProgress_IsOld(n.Ptr())
	return bool(result_)
}

func (n NSProgress) EstimatedTimeRemaining() Number {
	result_ := C.C_NSProgress_EstimatedTimeRemaining(n.Ptr())
	return MakeNumber(result_)
}

func (n NSProgress) SetEstimatedTimeRemaining(value Number) {
	C.C_NSProgress_SetEstimatedTimeRemaining(n.Ptr(), objc.ExtractPtr(value))
}

func (n NSProgress) FileCompletedCount() Number {
	result_ := C.C_NSProgress_FileCompletedCount(n.Ptr())
	return MakeNumber(result_)
}

func (n NSProgress) SetFileCompletedCount(value Number) {
	C.C_NSProgress_SetFileCompletedCount(n.Ptr(), objc.ExtractPtr(value))
}

func (n NSProgress) FileTotalCount() Number {
	result_ := C.C_NSProgress_FileTotalCount(n.Ptr())
	return MakeNumber(result_)
}

func (n NSProgress) SetFileTotalCount(value Number) {
	C.C_NSProgress_SetFileTotalCount(n.Ptr(), objc.ExtractPtr(value))
}

func (n NSProgress) Throughput() Number {
	result_ := C.C_NSProgress_Throughput(n.Ptr())
	return MakeNumber(result_)
}

func (n NSProgress) SetThroughput(value Number) {
	C.C_NSProgress_SetThroughput(n.Ptr(), objc.ExtractPtr(value))
}
