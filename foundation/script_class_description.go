package foundation

// #cgo CFLAGS: -x objective-c
// #cgo LDFLAGS: -framework Foundation
// #include "script_class_description.h"
import "C"
import (
	"github.com/hsiafan/cocoa/objc"
	"unsafe"
)

type ScriptClassDescription interface {
	ClassDescription
	ClassDescriptionForKey(key string) ScriptClassDescription
	IsLocationRequiredToCreateForKey(toManyRelationshipKey string) bool
	HasOrderedToManyRelationshipForKey(key string) bool
	HasPropertyForKey(key string) bool
	HasReadablePropertyForKey(key string) bool
	HasWritablePropertyForKey(key string) bool
	TypeForKey(key string) string
	SelectorForCommand(commandDescription ScriptCommandDescription) *objc.Selector
	SupportsCommand(commandDescription ScriptCommandDescription) bool
	SuperclassDescription() ScriptClassDescription
	ClassName() string
	DefaultSubcontainerAttributeKey() string
	ImplementationClassName() string
	SuiteName() string
}

type NSScriptClassDescription struct {
	NSClassDescription
}

func MakeScriptClassDescription(ptr unsafe.Pointer) *NSScriptClassDescription {
	if ptr == nil {
		return nil
	}
	return &NSScriptClassDescription{
		NSClassDescription: *MakeClassDescription(ptr),
	}
}

func AllocScriptClassDescription() *NSScriptClassDescription {
	return MakeScriptClassDescription(C.C_ScriptClassDescription_Alloc())
}

func (n *NSScriptClassDescription) Init() ScriptClassDescription {
	result_ := C.C_NSScriptClassDescription_Init(n.Ptr())
	return MakeScriptClassDescription(result_)
}

func (n *NSScriptClassDescription) ClassDescriptionForKey(key string) ScriptClassDescription {
	result_ := C.C_NSScriptClassDescription_ClassDescriptionForKey(n.Ptr(), NewString(key).Ptr())
	return MakeScriptClassDescription(result_)
}

func (n *NSScriptClassDescription) IsLocationRequiredToCreateForKey(toManyRelationshipKey string) bool {
	result_ := C.C_NSScriptClassDescription_IsLocationRequiredToCreateForKey(n.Ptr(), NewString(toManyRelationshipKey).Ptr())
	return bool(result_)
}

func (n *NSScriptClassDescription) HasOrderedToManyRelationshipForKey(key string) bool {
	result_ := C.C_NSScriptClassDescription_HasOrderedToManyRelationshipForKey(n.Ptr(), NewString(key).Ptr())
	return bool(result_)
}

func (n *NSScriptClassDescription) HasPropertyForKey(key string) bool {
	result_ := C.C_NSScriptClassDescription_HasPropertyForKey(n.Ptr(), NewString(key).Ptr())
	return bool(result_)
}

func (n *NSScriptClassDescription) HasReadablePropertyForKey(key string) bool {
	result_ := C.C_NSScriptClassDescription_HasReadablePropertyForKey(n.Ptr(), NewString(key).Ptr())
	return bool(result_)
}

func (n *NSScriptClassDescription) HasWritablePropertyForKey(key string) bool {
	result_ := C.C_NSScriptClassDescription_HasWritablePropertyForKey(n.Ptr(), NewString(key).Ptr())
	return bool(result_)
}

func (n *NSScriptClassDescription) TypeForKey(key string) string {
	result_ := C.C_NSScriptClassDescription_TypeForKey(n.Ptr(), NewString(key).Ptr())
	return MakeString(result_).String()
}

func (n *NSScriptClassDescription) SelectorForCommand(commandDescription ScriptCommandDescription) *objc.Selector {
	result_ := C.C_NSScriptClassDescription_SelectorForCommand(n.Ptr(), objc.ExtractPtr(commandDescription))
	return objc.MakeSelector(result_)
}

func (n *NSScriptClassDescription) SupportsCommand(commandDescription ScriptCommandDescription) bool {
	result_ := C.C_NSScriptClassDescription_SupportsCommand(n.Ptr(), objc.ExtractPtr(commandDescription))
	return bool(result_)
}

func (n *NSScriptClassDescription) SuperclassDescription() ScriptClassDescription {
	result_ := C.C_NSScriptClassDescription_SuperclassDescription(n.Ptr())
	return MakeScriptClassDescription(result_)
}

func (n *NSScriptClassDescription) ClassName() string {
	result_ := C.C_NSScriptClassDescription_ClassName(n.Ptr())
	return MakeString(result_).String()
}

func (n *NSScriptClassDescription) DefaultSubcontainerAttributeKey() string {
	result_ := C.C_NSScriptClassDescription_DefaultSubcontainerAttributeKey(n.Ptr())
	return MakeString(result_).String()
}

func (n *NSScriptClassDescription) ImplementationClassName() string {
	result_ := C.C_NSScriptClassDescription_ImplementationClassName(n.Ptr())
	return MakeString(result_).String()
}

func (n *NSScriptClassDescription) SuiteName() string {
	result_ := C.C_NSScriptClassDescription_SuiteName(n.Ptr())
	return MakeString(result_).String()
}
