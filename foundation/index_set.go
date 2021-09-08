package foundation

// #include "index_set.h"
import "C"
import (
	"github.com/hsiafan/cocoa/objc"
	"unsafe"
)

type IndexSet interface {
	objc.Object
	ContainsIndex(value uint) bool
	ContainsIndexes(indexSet IndexSet) bool
	ContainsIndexesInRange(_range Range) bool
	IntersectsIndexesInRange(_range Range) bool
	CountOfIndexesInRange(_range Range) uint
	IsEqualToIndexSet(indexSet IndexSet) bool
	IndexLessThanIndex(value uint) uint
	IndexLessThanOrEqualToIndex(value uint) uint
	IndexGreaterThanOrEqualToIndex(value uint) uint
	IndexGreaterThanIndex(value uint) uint
	Count() uint
	FirstIndex() uint
	LastIndex() uint
}

type NSIndexSet struct {
	objc.NSObject
}

func MakeIndexSet(ptr unsafe.Pointer) NSIndexSet {
	return NSIndexSet{
		NSObject: objc.MakeObject(ptr),
	}
}

func IndexSet_() NSIndexSet {
	result_ := C.C_NSIndexSet_IndexSet_()
	return MakeIndexSet(result_)
}

func IndexSetWithIndex(value uint) NSIndexSet {
	result_ := C.C_NSIndexSet_IndexSetWithIndex(C.uint(value))
	return MakeIndexSet(result_)
}

func IndexSetWithIndexesInRange(_range Range) NSIndexSet {
	result_ := C.C_NSIndexSet_IndexSetWithIndexesInRange(*(*C.NSRange)(ToNSRangePointer(_range)))
	return MakeIndexSet(result_)
}

func (n NSIndexSet) InitWithIndex(value uint) NSIndexSet {
	result_ := C.C_NSIndexSet_InitWithIndex(n.Ptr(), C.uint(value))
	return MakeIndexSet(result_)
}

func (n NSIndexSet) InitWithIndexesInRange(_range Range) NSIndexSet {
	result_ := C.C_NSIndexSet_InitWithIndexesInRange(n.Ptr(), *(*C.NSRange)(ToNSRangePointer(_range)))
	return MakeIndexSet(result_)
}

func (n NSIndexSet) InitWithIndexSet(indexSet IndexSet) NSIndexSet {
	result_ := C.C_NSIndexSet_InitWithIndexSet(n.Ptr(), objc.ExtractPtr(indexSet))
	return MakeIndexSet(result_)
}

func AllocIndexSet() NSIndexSet {
	result_ := C.C_NSIndexSet_AllocIndexSet()
	return MakeIndexSet(result_)
}

func (n NSIndexSet) Init() NSIndexSet {
	result_ := C.C_NSIndexSet_Init(n.Ptr())
	return MakeIndexSet(result_)
}

func NewIndexSet() NSIndexSet {
	result_ := C.C_NSIndexSet_NewIndexSet()
	return MakeIndexSet(result_)
}

func (n NSIndexSet) Autorelease() NSIndexSet {
	result_ := C.C_NSIndexSet_Autorelease(n.Ptr())
	return MakeIndexSet(result_)
}

func (n NSIndexSet) Retain() NSIndexSet {
	result_ := C.C_NSIndexSet_Retain(n.Ptr())
	return MakeIndexSet(result_)
}

func (n NSIndexSet) ContainsIndex(value uint) bool {
	result_ := C.C_NSIndexSet_ContainsIndex(n.Ptr(), C.uint(value))
	return bool(result_)
}

func (n NSIndexSet) ContainsIndexes(indexSet IndexSet) bool {
	result_ := C.C_NSIndexSet_ContainsIndexes(n.Ptr(), objc.ExtractPtr(indexSet))
	return bool(result_)
}

func (n NSIndexSet) ContainsIndexesInRange(_range Range) bool {
	result_ := C.C_NSIndexSet_ContainsIndexesInRange(n.Ptr(), *(*C.NSRange)(ToNSRangePointer(_range)))
	return bool(result_)
}

func (n NSIndexSet) IntersectsIndexesInRange(_range Range) bool {
	result_ := C.C_NSIndexSet_IntersectsIndexesInRange(n.Ptr(), *(*C.NSRange)(ToNSRangePointer(_range)))
	return bool(result_)
}

func (n NSIndexSet) CountOfIndexesInRange(_range Range) uint {
	result_ := C.C_NSIndexSet_CountOfIndexesInRange(n.Ptr(), *(*C.NSRange)(ToNSRangePointer(_range)))
	return uint(result_)
}

func (n NSIndexSet) IsEqualToIndexSet(indexSet IndexSet) bool {
	result_ := C.C_NSIndexSet_IsEqualToIndexSet(n.Ptr(), objc.ExtractPtr(indexSet))
	return bool(result_)
}

func (n NSIndexSet) IndexLessThanIndex(value uint) uint {
	result_ := C.C_NSIndexSet_IndexLessThanIndex(n.Ptr(), C.uint(value))
	return uint(result_)
}

func (n NSIndexSet) IndexLessThanOrEqualToIndex(value uint) uint {
	result_ := C.C_NSIndexSet_IndexLessThanOrEqualToIndex(n.Ptr(), C.uint(value))
	return uint(result_)
}

func (n NSIndexSet) IndexGreaterThanOrEqualToIndex(value uint) uint {
	result_ := C.C_NSIndexSet_IndexGreaterThanOrEqualToIndex(n.Ptr(), C.uint(value))
	return uint(result_)
}

func (n NSIndexSet) IndexGreaterThanIndex(value uint) uint {
	result_ := C.C_NSIndexSet_IndexGreaterThanIndex(n.Ptr(), C.uint(value))
	return uint(result_)
}

func (n NSIndexSet) Count() uint {
	result_ := C.C_NSIndexSet_Count(n.Ptr())
	return uint(result_)
}

func (n NSIndexSet) FirstIndex() uint {
	result_ := C.C_NSIndexSet_FirstIndex(n.Ptr())
	return uint(result_)
}

func (n NSIndexSet) LastIndex() uint {
	result_ := C.C_NSIndexSet_LastIndex(n.Ptr())
	return uint(result_)
}
