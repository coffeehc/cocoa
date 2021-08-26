package foundation

// #include "error.h"
import "C"
import (
	"github.com/hsiafan/cocoa/objc"
	"unsafe"
)

type Error interface {
	objc.Object
	AttemptRecoveryFromError_OptionIndex(error Error, recoveryOptionIndex uint) bool
	Code() int
	Domain() ErrorDomain
	UserInfo() map[ErrorUserInfoKey]objc.Object
	LocalizedDescription() string
	LocalizedRecoveryOptions() []string
	LocalizedRecoverySuggestion() string
	LocalizedFailureReason() string
	RecoveryAttempter() objc.Object
	HelpAnchor() string
	UnderlyingErrors() []Error
}

type NSError struct {
	objc.NSObject
}

func MakeError(ptr unsafe.Pointer) NSError {
	return NSError{
		NSObject: objc.MakeObject(ptr),
	}
}

func AllocError() NSError {
	return MakeError(C.C_Error_Alloc())
}

func (n NSError) InitWithDomain_Code_UserInfo(domain ErrorDomain, code int, dict map[ErrorUserInfoKey]objc.Object) Error {
	var cDict C.Dictionary
	if len(dict) > 0 {
		cDictKeyData := make([]unsafe.Pointer, len(dict))
		cDictValueData := make([]unsafe.Pointer, len(dict))
		var idx = 0
		for k, v := range dict {
			cDictKeyData[idx] = NewString(string(k)).Ptr()
			cDictValueData[idx] = objc.ExtractPtr(v)
			idx++
		}
		cDict.key_data = unsafe.Pointer(&cDictKeyData[0])
		cDict.value_data = unsafe.Pointer(&cDictValueData[0])
		cDict.len = C.int(len(dict))
	}
	result_ := C.C_NSError_InitWithDomain_Code_UserInfo(n.Ptr(), NewString(string(domain)).Ptr(), C.int(code), cDict)
	return MakeError(result_)
}

func (n NSError) Init() Error {
	result_ := C.C_NSError_Init(n.Ptr())
	return MakeError(result_)
}

func ErrorWithDomain_Code_UserInfo(domain ErrorDomain, code int, dict map[ErrorUserInfoKey]objc.Object) Error {
	var cDict C.Dictionary
	if len(dict) > 0 {
		cDictKeyData := make([]unsafe.Pointer, len(dict))
		cDictValueData := make([]unsafe.Pointer, len(dict))
		var idx = 0
		for k, v := range dict {
			cDictKeyData[idx] = NewString(string(k)).Ptr()
			cDictValueData[idx] = objc.ExtractPtr(v)
			idx++
		}
		cDict.key_data = unsafe.Pointer(&cDictKeyData[0])
		cDict.value_data = unsafe.Pointer(&cDictValueData[0])
		cDict.len = C.int(len(dict))
	}
	result_ := C.C_NSError_ErrorWithDomain_Code_UserInfo(NewString(string(domain)).Ptr(), C.int(code), cDict)
	return MakeError(result_)
}

func (n NSError) AttemptRecoveryFromError_OptionIndex(error Error, recoveryOptionIndex uint) bool {
	result_ := C.C_NSError_AttemptRecoveryFromError_OptionIndex(n.Ptr(), objc.ExtractPtr(error), C.uint(recoveryOptionIndex))
	return bool(result_)
}

func (n NSError) Code() int {
	result_ := C.C_NSError_Code(n.Ptr())
	return int(result_)
}

func (n NSError) Domain() ErrorDomain {
	result_ := C.C_NSError_Domain(n.Ptr())
	return ErrorDomain(MakeString(result_).String())
}

func (n NSError) UserInfo() map[ErrorUserInfoKey]objc.Object {
	result_ := C.C_NSError_UserInfo(n.Ptr())
	if result_.len > 0 {
		defer C.free(result_.key_data)
		defer C.free(result_.value_data)
	}
	result_KeySlice := (*[1 << 28]unsafe.Pointer)(unsafe.Pointer(result_.key_data))[:result_.len:result_.len]
	result_ValueSlice := (*[1 << 28]unsafe.Pointer)(unsafe.Pointer(result_.value_data))[:result_.len:result_.len]
	var goResult_ = make(map[ErrorUserInfoKey]objc.Object)
	for idx, k := range result_KeySlice {
		v := result_ValueSlice[idx]
		goResult_[ErrorUserInfoKey(MakeString(k).String())] = objc.MakeObject(v)
	}
	return goResult_
}

func (n NSError) LocalizedDescription() string {
	result_ := C.C_NSError_LocalizedDescription(n.Ptr())
	return MakeString(result_).String()
}

func (n NSError) LocalizedRecoveryOptions() []string {
	result_ := C.C_NSError_LocalizedRecoveryOptions(n.Ptr())
	if result_.len > 0 {
		defer C.free(result_.data)
	}
	result_Slice := (*[1 << 28]unsafe.Pointer)(unsafe.Pointer(result_.data))[:result_.len:result_.len]
	var goResult_ = make([]string, len(result_Slice))
	for idx, r := range result_Slice {
		goResult_[idx] = MakeString(r).String()
	}
	return goResult_
}

func (n NSError) LocalizedRecoverySuggestion() string {
	result_ := C.C_NSError_LocalizedRecoverySuggestion(n.Ptr())
	return MakeString(result_).String()
}

func (n NSError) LocalizedFailureReason() string {
	result_ := C.C_NSError_LocalizedFailureReason(n.Ptr())
	return MakeString(result_).String()
}

func (n NSError) RecoveryAttempter() objc.Object {
	result_ := C.C_NSError_RecoveryAttempter(n.Ptr())
	return objc.MakeObject(result_)
}

func (n NSError) HelpAnchor() string {
	result_ := C.C_NSError_HelpAnchor(n.Ptr())
	return MakeString(result_).String()
}

func (n NSError) UnderlyingErrors() []Error {
	result_ := C.C_NSError_UnderlyingErrors(n.Ptr())
	if result_.len > 0 {
		defer C.free(result_.data)
	}
	result_Slice := (*[1 << 28]unsafe.Pointer)(unsafe.Pointer(result_.data))[:result_.len:result_.len]
	var goResult_ = make([]Error, len(result_Slice))
	for idx, r := range result_Slice {
		goResult_[idx] = MakeError(r)
	}
	return goResult_
}
