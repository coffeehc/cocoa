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
	Section() int
	Item() int
	Length() uint
}

type NSIndexPath struct {
	objc.NSObject
}

func MakeIndexPath(ptr unsafe.Pointer) *NSIndexPath {
	if ptr == nil {
		return nil
	}
	return &NSIndexPath{
		NSObject: *objc.MakeObject(ptr),
	}
}

func AllocIndexPath() *NSIndexPath {
	return MakeIndexPath(C.C_IndexPath_Alloc())
}

func (n *NSIndexPath) InitWithIndex(index uint) IndexPath {
	result_ := C.C_NSIndexPath_InitWithIndex(n.Ptr(), C.uint(index))
	return MakeIndexPath(result_)
}

func (n *NSIndexPath) Init() IndexPath {
	result_ := C.C_NSIndexPath_Init(n.Ptr())
	return MakeIndexPath(result_)
}

func IndexPathWithIndex(index uint) IndexPath {
	result_ := C.C_NSIndexPath_IndexPathWithIndex(C.uint(index))
	return MakeIndexPath(result_)
}

func IndexPathForItem_InSection(item int, section int) IndexPath {
	result_ := C.C_NSIndexPath_IndexPathForItem_InSection(C.int(item), C.int(section))
	return MakeIndexPath(result_)
}

func (n *NSIndexPath) IndexPathByAddingIndex(index uint) IndexPath {
	result_ := C.C_NSIndexPath_IndexPathByAddingIndex(n.Ptr(), C.uint(index))
	return MakeIndexPath(result_)
}

func (n *NSIndexPath) IndexPathByRemovingLastIndex() IndexPath {
	result_ := C.C_NSIndexPath_IndexPathByRemovingLastIndex(n.Ptr())
	return MakeIndexPath(result_)
}

func (n *NSIndexPath) Compare(otherObject IndexPath) ComparisonResult {
	result_ := C.C_NSIndexPath_Compare(n.Ptr(), objc.ExtractPtr(otherObject))
	return ComparisonResult(int(result_))
}

func (n *NSIndexPath) IndexAtPosition(position uint) uint {
	result_ := C.C_NSIndexPath_IndexAtPosition(n.Ptr(), C.uint(position))
	return uint(result_)
}

func (n *NSIndexPath) Section() int {
	result_ := C.C_NSIndexPath_Section(n.Ptr())
	return int(result_)
}

func (n *NSIndexPath) Item() int {
	result_ := C.C_NSIndexPath_Item(n.Ptr())
	return int(result_)
}

func (n *NSIndexPath) Length() uint {
	result_ := C.C_NSIndexPath_Length(n.Ptr())
	return uint(result_)
}
