package foundation

// #cgo CFLAGS: -x objective-c
// #cgo LDFLAGS: -framework Foundation
// #include "class_description.h"
import "C"
import (
	"github.com/hsiafan/cocoa/objc"
	"unsafe"
)

type ClassDescription interface {
	objc.Object
	InverseForRelationshipKey(relationshipKey string) string
}

type NSClassDescription struct {
	objc.NSObject
}

func MakeClassDescription(ptr unsafe.Pointer) *NSClassDescription {
	if ptr == nil {
		return nil
	}
	return &NSClassDescription{
		NSObject: *objc.MakeObject(ptr),
	}
}

func AllocClassDescription() *NSClassDescription {
	return MakeClassDescription(C.C_ClassDescription_Alloc())
}

func (n *NSClassDescription) Init() ClassDescription {
	result := C.C_NSClassDescription_Init(n.Ptr())
	return MakeClassDescription(result)
}

func ClassDescriptionInvalidateClassDescriptionCache() {
	C.C_NSClassDescription_ClassDescriptionInvalidateClassDescriptionCache()
}

func (n *NSClassDescription) InverseForRelationshipKey(relationshipKey string) string {
	result := C.C_NSClassDescription_InverseForRelationshipKey(n.Ptr(), NewString(relationshipKey).Ptr())
	return MakeString(result).String()
}
