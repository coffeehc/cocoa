package foundation

// #cgo CFLAGS: -x objective-c
// #cgo LDFLAGS: -framework Foundation
// #include "script_object_specifier.h"
import "C"
import (
	"github.com/hsiafan/cocoa/objc"
	"unsafe"
)

type ScriptObjectSpecifier interface {
	objc.Object
	ObjectsByEvaluatingWithContainers(containers objc.Object) objc.Object
	ObjectsByEvaluatingSpecifier() objc.Object
	ContainerClassDescription() ScriptClassDescription
	SetContainerClassDescription(value ScriptClassDescription)
	ContainerIsObjectBeingTested() bool
	SetContainerIsObjectBeingTested(value bool)
	ContainerIsRangeContainerObject() bool
	SetContainerIsRangeContainerObject(value bool)
	ContainerSpecifier() ScriptObjectSpecifier
	SetContainerSpecifier(value ScriptObjectSpecifier)
	ChildSpecifier() ScriptObjectSpecifier
	SetChildSpecifier(value ScriptObjectSpecifier)
	Key() string
	SetKey(value string)
	KeyClassDescription() ScriptClassDescription
	EvaluationErrorSpecifier() ScriptObjectSpecifier
	EvaluationErrorNumber() int
	SetEvaluationErrorNumber(value int)
	Descriptor() AppleEventDescriptor
}

type NSScriptObjectSpecifier struct {
	objc.NSObject
}

func MakeScriptObjectSpecifier(ptr unsafe.Pointer) *NSScriptObjectSpecifier {
	if ptr == nil {
		return nil
	}
	return &NSScriptObjectSpecifier{
		NSObject: *objc.MakeObject(ptr),
	}
}

func AllocScriptObjectSpecifier() *NSScriptObjectSpecifier {
	return MakeScriptObjectSpecifier(C.C_ScriptObjectSpecifier_Alloc())
}

func (n *NSScriptObjectSpecifier) InitWithContainerClassDescription_ContainerSpecifier_Key(classDesc ScriptClassDescription, container ScriptObjectSpecifier, property string) ScriptObjectSpecifier {
	result := C.C_NSScriptObjectSpecifier_InitWithContainerClassDescription_ContainerSpecifier_Key(n.Ptr(), objc.ExtractPtr(classDesc), objc.ExtractPtr(container), NewString(property).Ptr())
	return MakeScriptObjectSpecifier(result)
}

func (n *NSScriptObjectSpecifier) InitWithContainerSpecifier_Key(container ScriptObjectSpecifier, property string) ScriptObjectSpecifier {
	result := C.C_NSScriptObjectSpecifier_InitWithContainerSpecifier_Key(n.Ptr(), objc.ExtractPtr(container), NewString(property).Ptr())
	return MakeScriptObjectSpecifier(result)
}

func (n *NSScriptObjectSpecifier) InitWithCoder(inCoder Coder) ScriptObjectSpecifier {
	result := C.C_NSScriptObjectSpecifier_InitWithCoder(n.Ptr(), objc.ExtractPtr(inCoder))
	return MakeScriptObjectSpecifier(result)
}

func (n *NSScriptObjectSpecifier) Init() ScriptObjectSpecifier {
	result := C.C_NSScriptObjectSpecifier_Init(n.Ptr())
	return MakeScriptObjectSpecifier(result)
}

func ScriptObjectSpecifierObjectSpecifierWithDescriptor(descriptor AppleEventDescriptor) ScriptObjectSpecifier {
	result := C.C_NSScriptObjectSpecifier_ScriptObjectSpecifierObjectSpecifierWithDescriptor(objc.ExtractPtr(descriptor))
	return MakeScriptObjectSpecifier(result)
}

func (n *NSScriptObjectSpecifier) ObjectsByEvaluatingWithContainers(containers objc.Object) objc.Object {
	result := C.C_NSScriptObjectSpecifier_ObjectsByEvaluatingWithContainers(n.Ptr(), objc.ExtractPtr(containers))
	return objc.MakeObject(result)
}

func (n *NSScriptObjectSpecifier) ObjectsByEvaluatingSpecifier() objc.Object {
	result := C.C_NSScriptObjectSpecifier_ObjectsByEvaluatingSpecifier(n.Ptr())
	return objc.MakeObject(result)
}

func (n *NSScriptObjectSpecifier) ContainerClassDescription() ScriptClassDescription {
	result := C.C_NSScriptObjectSpecifier_ContainerClassDescription(n.Ptr())
	return MakeScriptClassDescription(result)
}

func (n *NSScriptObjectSpecifier) SetContainerClassDescription(value ScriptClassDescription) {
	C.C_NSScriptObjectSpecifier_SetContainerClassDescription(n.Ptr(), objc.ExtractPtr(value))
}

func (n *NSScriptObjectSpecifier) ContainerIsObjectBeingTested() bool {
	result := C.C_NSScriptObjectSpecifier_ContainerIsObjectBeingTested(n.Ptr())
	return bool(result)
}

func (n *NSScriptObjectSpecifier) SetContainerIsObjectBeingTested(value bool) {
	C.C_NSScriptObjectSpecifier_SetContainerIsObjectBeingTested(n.Ptr(), C.bool(value))
}

func (n *NSScriptObjectSpecifier) ContainerIsRangeContainerObject() bool {
	result := C.C_NSScriptObjectSpecifier_ContainerIsRangeContainerObject(n.Ptr())
	return bool(result)
}

func (n *NSScriptObjectSpecifier) SetContainerIsRangeContainerObject(value bool) {
	C.C_NSScriptObjectSpecifier_SetContainerIsRangeContainerObject(n.Ptr(), C.bool(value))
}

func (n *NSScriptObjectSpecifier) ContainerSpecifier() ScriptObjectSpecifier {
	result := C.C_NSScriptObjectSpecifier_ContainerSpecifier(n.Ptr())
	return MakeScriptObjectSpecifier(result)
}

func (n *NSScriptObjectSpecifier) SetContainerSpecifier(value ScriptObjectSpecifier) {
	C.C_NSScriptObjectSpecifier_SetContainerSpecifier(n.Ptr(), objc.ExtractPtr(value))
}

func (n *NSScriptObjectSpecifier) ChildSpecifier() ScriptObjectSpecifier {
	result := C.C_NSScriptObjectSpecifier_ChildSpecifier(n.Ptr())
	return MakeScriptObjectSpecifier(result)
}

func (n *NSScriptObjectSpecifier) SetChildSpecifier(value ScriptObjectSpecifier) {
	C.C_NSScriptObjectSpecifier_SetChildSpecifier(n.Ptr(), objc.ExtractPtr(value))
}

func (n *NSScriptObjectSpecifier) Key() string {
	result := C.C_NSScriptObjectSpecifier_Key(n.Ptr())
	return MakeString(result).String()
}

func (n *NSScriptObjectSpecifier) SetKey(value string) {
	C.C_NSScriptObjectSpecifier_SetKey(n.Ptr(), NewString(value).Ptr())
}

func (n *NSScriptObjectSpecifier) KeyClassDescription() ScriptClassDescription {
	result := C.C_NSScriptObjectSpecifier_KeyClassDescription(n.Ptr())
	return MakeScriptClassDescription(result)
}

func (n *NSScriptObjectSpecifier) EvaluationErrorSpecifier() ScriptObjectSpecifier {
	result := C.C_NSScriptObjectSpecifier_EvaluationErrorSpecifier(n.Ptr())
	return MakeScriptObjectSpecifier(result)
}

func (n *NSScriptObjectSpecifier) EvaluationErrorNumber() int {
	result := C.C_NSScriptObjectSpecifier_EvaluationErrorNumber(n.Ptr())
	return int(result)
}

func (n *NSScriptObjectSpecifier) SetEvaluationErrorNumber(value int) {
	C.C_NSScriptObjectSpecifier_SetEvaluationErrorNumber(n.Ptr(), C.int(value))
}

func (n *NSScriptObjectSpecifier) Descriptor() AppleEventDescriptor {
	result := C.C_NSScriptObjectSpecifier_Descriptor(n.Ptr())
	return MakeAppleEventDescriptor(result)
}
