package foundation

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
	ArgumentNames() []string
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
	result_ := C.C_NSScriptCommandDescription_InitWithCoder(n.Ptr(), objc.ExtractPtr(inCoder))
	return MakeScriptCommandDescription(result_)
}

func (n *NSScriptCommandDescription) IsOptionalArgumentWithName(argumentName string) bool {
	result_ := C.C_NSScriptCommandDescription_IsOptionalArgumentWithName(n.Ptr(), NewString(argumentName).Ptr())
	return bool(result_)
}

func (n *NSScriptCommandDescription) TypeForArgumentWithName(argumentName string) string {
	result_ := C.C_NSScriptCommandDescription_TypeForArgumentWithName(n.Ptr(), NewString(argumentName).Ptr())
	return MakeString(result_).String()
}

func (n *NSScriptCommandDescription) CreateCommandInstance() ScriptCommand {
	result_ := C.C_NSScriptCommandDescription_CreateCommandInstance(n.Ptr())
	return MakeScriptCommand(result_)
}

func (n *NSScriptCommandDescription) CommandClassName() string {
	result_ := C.C_NSScriptCommandDescription_CommandClassName(n.Ptr())
	return MakeString(result_).String()
}

func (n *NSScriptCommandDescription) CommandName() string {
	result_ := C.C_NSScriptCommandDescription_CommandName(n.Ptr())
	return MakeString(result_).String()
}

func (n *NSScriptCommandDescription) SuiteName() string {
	result_ := C.C_NSScriptCommandDescription_SuiteName(n.Ptr())
	return MakeString(result_).String()
}

func (n *NSScriptCommandDescription) ArgumentNames() []string {
	result_ := C.C_NSScriptCommandDescription_ArgumentNames(n.Ptr())
	defer C.free(result_.data)
	result_Slice := (*[1 << 28]unsafe.Pointer)(unsafe.Pointer(result_.data))[:result_.len:result_.len]
	var goResult_ = make([]string, len(result_Slice))
	for idx, r := range result_Slice {
		goResult_[idx] = MakeString(r).String()
	}
	return goResult_
}

func (n *NSScriptCommandDescription) ReturnType() string {
	result_ := C.C_NSScriptCommandDescription_ReturnType(n.Ptr())
	return MakeString(result_).String()
}
