package foundation

// #include "index_path.h"
import "C"
import (
	"github.com/hsiafan/cocoa/objc"
	"unsafe"
)

type IndexPath interface {
	objc.Object
	IndexPathByAddingIndex(index uint) IndexPath
	IndexPathByRemovingLastIndex() IndexPath
	Compare(otherObject IndexPath) ComparisonResult
	IndexAtPosition(position uint) uint
	Length() uint
}

type NSIndexPath struct {
	objc.NSObject
}

func MakeIndexPath(ptr unsafe.Pointer) NSIndexPath {
	return NSIndexPath{
		NSObject: objc.MakeObject(ptr),
	}
}

func IndexPathWithIndex(index uint) NSIndexPath {
	result_ := C.C_NSIndexPath_IndexPathWithIndex(C.uint(index))
	return MakeIndexPath(result_)
}

func (n NSIndexPath) InitWithIndex(index uint) NSIndexPath {
	result_ := C.C_NSIndexPath_InitWithIndex(n.Ptr(), C.uint(index))
	return MakeIndexPath(result_)
}

func AllocIndexPath() NSIndexPath {
	result_ := C.C_NSIndexPath_AllocIndexPath()
	return MakeIndexPath(result_)
}

func (n NSIndexPath) Init() NSIndexPath {
	result_ := C.C_NSIndexPath_Init(n.Ptr())
	return MakeIndexPath(result_)
}

func NewIndexPath() NSIndexPath {
	result_ := C.C_NSIndexPath_NewIndexPath()
	return MakeIndexPath(result_)
}

func (n NSIndexPath) Autorelease() NSIndexPath {
	result_ := C.C_NSIndexPath_Autorelease(n.Ptr())
	return MakeIndexPath(result_)
}

func (n NSIndexPath) Retain() NSIndexPath {
	result_ := C.C_NSIndexPath_Retain(n.Ptr())
	return MakeIndexPath(result_)
}

func (n NSIndexPath) IndexPathByAddingIndex(index uint) IndexPath {
	result_ := C.C_NSIndexPath_IndexPathByAddingIndex(n.Ptr(), C.uint(index))
	return MakeIndexPath(result_)
}

func (n NSIndexPath) IndexPathByRemovingLastIndex() IndexPath {
	result_ := C.C_NSIndexPath_IndexPathByRemovingLastIndex(n.Ptr())
	return MakeIndexPath(result_)
}

func (n NSIndexPath) Compare(otherObject IndexPath) ComparisonResult {
	result_ := C.C_NSIndexPath_Compare(n.Ptr(), objc.ExtractPtr(otherObject))
	return ComparisonResult(int(result_))
}

func (n NSIndexPath) IndexAtPosition(position uint) uint {
	result_ := C.C_NSIndexPath_IndexAtPosition(n.Ptr(), C.uint(position))
	return uint(result_)
}

func (n NSIndexPath) Length() uint {
	result_ := C.C_NSIndexPath_Length(n.Ptr())
	return uint(result_)
}
