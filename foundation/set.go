package foundation

// #include "set.h"
import "C"
import (
	"github.com/hsiafan/cocoa/objc"
	"unsafe"
)

type Set interface {
	objc.Object
	SetByAddingObjectsFromSet(other Set) Set
	MakeObjectsPerformSelector(aSelector objc.Selector)
	MakeObjectsPerformSelector_WithObject(aSelector objc.Selector, argument objc.Object)
	IsSubsetOfSet(otherSet Set) bool
	IntersectsSet(otherSet Set) bool
	IsEqualToSet(otherSet Set) bool
	ValueForKey(key string) objc.Object
	SetValue_ForKey(value objc.Object, key string)
	RemoveObserver_ForKeyPath(observer objc.Object, keyPath string)
	DescriptionWithLocale(locale objc.Object) string
	Count() uint
	Description() string
}

type NSSet struct {
	objc.NSObject
}

func MakeSet(ptr unsafe.Pointer) NSSet {
	return NSSet{
		NSObject: objc.MakeObject(ptr),
	}
}

func AllocSet() NSSet {
	return MakeSet(C.C_Set_Alloc())
}

func (n NSSet) InitWithSet(set Set) Set {
	result_ := C.C_NSSet_InitWithSet(n.Ptr(), objc.ExtractPtr(set))
	return MakeSet(result_)
}

func (n NSSet) InitWithSet_CopyItems(set Set, flag bool) Set {
	result_ := C.C_NSSet_InitWithSet_CopyItems(n.Ptr(), objc.ExtractPtr(set), C.bool(flag))
	return MakeSet(result_)
}

func (n NSSet) Init() Set {
	result_ := C.C_NSSet_Init(n.Ptr())
	return MakeSet(result_)
}

func (n NSSet) InitWithCoder(coder Coder) Set {
	result_ := C.C_NSSet_InitWithCoder(n.Ptr(), objc.ExtractPtr(coder))
	return MakeSet(result_)
}

func Set_() Set {
	result_ := C.C_NSSet_Set_()
	return MakeSet(result_)
}

func SetWithSet(set Set) Set {
	result_ := C.C_NSSet_SetWithSet(objc.ExtractPtr(set))
	return MakeSet(result_)
}

func (n NSSet) SetByAddingObjectsFromSet(other Set) Set {
	result_ := C.C_NSSet_SetByAddingObjectsFromSet(n.Ptr(), objc.ExtractPtr(other))
	return MakeSet(result_)
}

func (n NSSet) MakeObjectsPerformSelector(aSelector objc.Selector) {
	C.C_NSSet_MakeObjectsPerformSelector(n.Ptr(), unsafe.Pointer(aSelector))
}

func (n NSSet) MakeObjectsPerformSelector_WithObject(aSelector objc.Selector, argument objc.Object) {
	C.C_NSSet_MakeObjectsPerformSelector_WithObject(n.Ptr(), unsafe.Pointer(aSelector), objc.ExtractPtr(argument))
}

func (n NSSet) IsSubsetOfSet(otherSet Set) bool {
	result_ := C.C_NSSet_IsSubsetOfSet(n.Ptr(), objc.ExtractPtr(otherSet))
	return bool(result_)
}

func (n NSSet) IntersectsSet(otherSet Set) bool {
	result_ := C.C_NSSet_IntersectsSet(n.Ptr(), objc.ExtractPtr(otherSet))
	return bool(result_)
}

func (n NSSet) IsEqualToSet(otherSet Set) bool {
	result_ := C.C_NSSet_IsEqualToSet(n.Ptr(), objc.ExtractPtr(otherSet))
	return bool(result_)
}

func (n NSSet) ValueForKey(key string) objc.Object {
	result_ := C.C_NSSet_ValueForKey(n.Ptr(), NewString(key).Ptr())
	return objc.MakeObject(result_)
}

func (n NSSet) SetValue_ForKey(value objc.Object, key string) {
	C.C_NSSet_SetValue_ForKey(n.Ptr(), objc.ExtractPtr(value), NewString(key).Ptr())
}

func (n NSSet) RemoveObserver_ForKeyPath(observer objc.Object, keyPath string) {
	C.C_NSSet_RemoveObserver_ForKeyPath(n.Ptr(), objc.ExtractPtr(observer), NewString(keyPath).Ptr())
}

func (n NSSet) DescriptionWithLocale(locale objc.Object) string {
	result_ := C.C_NSSet_DescriptionWithLocale(n.Ptr(), objc.ExtractPtr(locale))
	return MakeString(result_).String()
}

func (n NSSet) Count() uint {
	result_ := C.C_NSSet_Count(n.Ptr())
	return uint(result_)
}

func (n NSSet) Description() string {
	result_ := C.C_NSSet_Description(n.Ptr())
	return MakeString(result_).String()
}
