package foundation

// #include "sort_descriptor.h"
import "C"
import (
	"github.com/hsiafan/cocoa/objc"
	"unsafe"
)

type SortDescriptor interface {
	objc.Object
	CompareObject_ToObject(object1 objc.Object, object2 objc.Object) ComparisonResult
	AllowEvaluation()
	Ascending() bool
	Key() string
	Selector() objc.Selector
	ReversedSortDescriptor() objc.Object
}

type NSSortDescriptor struct {
	objc.NSObject
}

func MakeSortDescriptor(ptr unsafe.Pointer) NSSortDescriptor {
	return NSSortDescriptor{
		NSObject: objc.MakeObject(ptr),
	}
}

func SortDescriptorWithKey_Ascending(key string, ascending bool) NSSortDescriptor {
	result_ := C.C_NSSortDescriptor_SortDescriptorWithKey_Ascending(NewString(key).Ptr(), C.bool(ascending))
	return MakeSortDescriptor(result_)
}

func (n NSSortDescriptor) InitWithKey_Ascending(key string, ascending bool) NSSortDescriptor {
	result_ := C.C_NSSortDescriptor_InitWithKey_Ascending(n.Ptr(), NewString(key).Ptr(), C.bool(ascending))
	return MakeSortDescriptor(result_)
}

func SortDescriptorWithKey_Ascending_Selector(key string, ascending bool, selector objc.Selector) NSSortDescriptor {
	result_ := C.C_NSSortDescriptor_SortDescriptorWithKey_Ascending_Selector(NewString(key).Ptr(), C.bool(ascending), unsafe.Pointer(selector))
	return MakeSortDescriptor(result_)
}

func (n NSSortDescriptor) InitWithKey_Ascending_Selector(key string, ascending bool, selector objc.Selector) NSSortDescriptor {
	result_ := C.C_NSSortDescriptor_InitWithKey_Ascending_Selector(n.Ptr(), NewString(key).Ptr(), C.bool(ascending), unsafe.Pointer(selector))
	return MakeSortDescriptor(result_)
}

func (n NSSortDescriptor) InitWithCoder(coder Coder) NSSortDescriptor {
	result_ := C.C_NSSortDescriptor_InitWithCoder(n.Ptr(), objc.ExtractPtr(coder))
	return MakeSortDescriptor(result_)
}

func AllocSortDescriptor() NSSortDescriptor {
	result_ := C.C_NSSortDescriptor_AllocSortDescriptor()
	return MakeSortDescriptor(result_)
}

func (n NSSortDescriptor) Init() NSSortDescriptor {
	result_ := C.C_NSSortDescriptor_Init(n.Ptr())
	return MakeSortDescriptor(result_)
}

func NewSortDescriptor() NSSortDescriptor {
	result_ := C.C_NSSortDescriptor_NewSortDescriptor()
	return MakeSortDescriptor(result_)
}

func (n NSSortDescriptor) Autorelease() NSSortDescriptor {
	result_ := C.C_NSSortDescriptor_Autorelease(n.Ptr())
	return MakeSortDescriptor(result_)
}

func (n NSSortDescriptor) Retain() NSSortDescriptor {
	result_ := C.C_NSSortDescriptor_Retain(n.Ptr())
	return MakeSortDescriptor(result_)
}

func (n NSSortDescriptor) CompareObject_ToObject(object1 objc.Object, object2 objc.Object) ComparisonResult {
	result_ := C.C_NSSortDescriptor_CompareObject_ToObject(n.Ptr(), objc.ExtractPtr(object1), objc.ExtractPtr(object2))
	return ComparisonResult(int(result_))
}

func (n NSSortDescriptor) AllowEvaluation() {
	C.C_NSSortDescriptor_AllowEvaluation(n.Ptr())
}

func (n NSSortDescriptor) Ascending() bool {
	result_ := C.C_NSSortDescriptor_Ascending(n.Ptr())
	return bool(result_)
}

func (n NSSortDescriptor) Key() string {
	result_ := C.C_NSSortDescriptor_Key(n.Ptr())
	return MakeString(result_).String()
}

func (n NSSortDescriptor) Selector() objc.Selector {
	result_ := C.C_NSSortDescriptor_Selector(n.Ptr())
	return objc.Selector(result_)
}

func (n NSSortDescriptor) ReversedSortDescriptor() objc.Object {
	result_ := C.C_NSSortDescriptor_ReversedSortDescriptor(n.Ptr())
	return objc.MakeObject(result_)
}
