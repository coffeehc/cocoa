package foundation

// #include "class_description.h"
import "C"
import (
	"github.com/hsiafan/cocoa/objc"
	"unsafe"
)

type ClassDescription interface {
	objc.Object
	InverseForRelationshipKey(relationshipKey string) string
	AttributeKeys() []string
	ToManyRelationshipKeys() []string
	ToOneRelationshipKeys() []string
}

type NSClassDescription struct {
	objc.NSObject
}

func MakeClassDescription(ptr unsafe.Pointer) NSClassDescription {
	return NSClassDescription{
		NSObject: objc.MakeObject(ptr),
	}
}

func AllocClassDescription() NSClassDescription {
	return MakeClassDescription(C.C_ClassDescription_Alloc())
}

func (n NSClassDescription) Init() ClassDescription {
	result_ := C.C_NSClassDescription_Init(n.Ptr())
	return MakeClassDescription(result_)
}

func InvalidateClassDescriptionCache() {
	C.C_NSClassDescription_InvalidateClassDescriptionCache()
}

func (n NSClassDescription) InverseForRelationshipKey(relationshipKey string) string {
	result_ := C.C_NSClassDescription_InverseForRelationshipKey(n.Ptr(), NewString(relationshipKey).Ptr())
	return MakeString(result_).String()
}

func (n NSClassDescription) AttributeKeys() []string {
	result_ := C.C_NSClassDescription_AttributeKeys(n.Ptr())
	if result_.len > 0 {
		defer C.free(result_.data)
	}
	result_Slice := unsafe.Slice((*unsafe.Pointer)(result_.data), int(result_.len))
	var goResult_ = make([]string, len(result_Slice))
	for idx, r := range result_Slice {
		goResult_[idx] = MakeString(r).String()
	}
	return goResult_
}

func (n NSClassDescription) ToManyRelationshipKeys() []string {
	result_ := C.C_NSClassDescription_ToManyRelationshipKeys(n.Ptr())
	if result_.len > 0 {
		defer C.free(result_.data)
	}
	result_Slice := unsafe.Slice((*unsafe.Pointer)(result_.data), int(result_.len))
	var goResult_ = make([]string, len(result_Slice))
	for idx, r := range result_Slice {
		goResult_[idx] = MakeString(r).String()
	}
	return goResult_
}

func (n NSClassDescription) ToOneRelationshipKeys() []string {
	result_ := C.C_NSClassDescription_ToOneRelationshipKeys(n.Ptr())
	if result_.len > 0 {
		defer C.free(result_.data)
	}
	result_Slice := unsafe.Slice((*unsafe.Pointer)(result_.data), int(result_.len))
	var goResult_ = make([]string, len(result_Slice))
	for idx, r := range result_Slice {
		goResult_[idx] = MakeString(r).String()
	}
	return goResult_
}
