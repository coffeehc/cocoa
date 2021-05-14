package foundation

// #cgo CFLAGS: -x objective-c
// #cgo LDFLAGS: -framework Foundation
// #include "script_command_description.h"
import "C"
import (
	"github.com/hsiafan/cocoa/objc"
	"unsafe"
)

type ScriptCommandDescription interface {
	objc.Object
	IsOptionalArgumentWithName(argumentName string) bool
	TypeForArgumentWithName(argumentName string) string
	CreateCommandInstance() ScriptCommand
	CommandClassName() string
	CommandName() string
	SuiteName() string
	ReturnType() string
}

type NSScriptCommandDescription struct {
	objc.NSObject
}

func MakeScriptCommandDescription(ptr unsafe.Pointer) *NSScriptCommandDescription {
	if ptr == nil {
		return nil
	}
	return &NSScriptCommandDescription{
		NSObject: *objc.MakeObject(ptr),
	}
}

func AllocScriptCommandDescription() *NSScriptCommandDescription {
	return MakeScriptCommandDescription(C.C_ScriptCommandDescription_Alloc())
}

func (n *NSScriptCommandDescription) InitWithCoder(inCoder Coder) ScriptCommandDescription {
	result := C.C_NSScriptCommandDescription_InitWithCoder(n.Ptr(), objc.ExtractPtr(inCoder))
	return MakeScriptCommandDescription(result)
}

func (n *NSScriptCommandDescription) IsOptionalArgumentWithName(argumentName string) bool {
	result := C.C_NSScriptCommandDescription_IsOptionalArgumentWithName(n.Ptr(), NewString(argumentName).Ptr())
	return bool(result)
}

func (n *NSScriptCommandDescription) TypeForArgumentWithName(argumentName string) string {
	result := C.C_NSScriptCommandDescription_TypeForArgumentWithName(n.Ptr(), NewString(argumentName).Ptr())
	return MakeString(result).String()
}

func (n *NSScriptCommandDescription) CreateCommandInstance() ScriptCommand {
	result := C.C_NSScriptCommandDescription_CreateCommandInstance(n.Ptr())
	return MakeScriptCommand(result)
}

func (n *NSScriptCommandDescription) CommandClassName() string {
	result := C.C_NSScriptCommandDescription_CommandClassName(n.Ptr())
	return MakeString(result).String()
}

func (n *NSScriptCommandDescription) CommandName() string {
	result := C.C_NSScriptCommandDescription_CommandName(n.Ptr())
	return MakeString(result).String()
}

func (n *NSScriptCommandDescription) SuiteName() string {
	result := C.C_NSScriptCommandDescription_SuiteName(n.Ptr())
	return MakeString(result).String()
}

func (n *NSScriptCommandDescription) ReturnType() string {
	result := C.C_NSScriptCommandDescription_ReturnType(n.Ptr())
	return MakeString(result).String()
}
