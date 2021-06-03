package foundation

// #include "script_command.h"
import "C"
import (
	"github.com/hsiafan/cocoa/objc"
	"unsafe"
)

type ScriptCommand interface {
	objc.Object
	ExecuteCommand() objc.Object
	PerformDefaultImplementation() objc.Object
	SuspendExecution()
	ResumeExecutionWithResult(result objc.Object)
	AppleEvent() AppleEventDescriptor
	EvaluatedReceivers() objc.Object
	ReceiversSpecifier() ScriptObjectSpecifier
	SetReceiversSpecifier(value ScriptObjectSpecifier)
	DirectParameter() objc.Object
	SetDirectParameter(value objc.Object)
	CommandDescription() ScriptCommandDescription
	ScriptErrorExpectedTypeDescriptor() AppleEventDescriptor
	SetScriptErrorExpectedTypeDescriptor(value AppleEventDescriptor)
	ScriptErrorNumber() int
	SetScriptErrorNumber(value int)
	ScriptErrorOffendingObjectDescriptor() AppleEventDescriptor
	SetScriptErrorOffendingObjectDescriptor(value AppleEventDescriptor)
	ScriptErrorString() string
	SetScriptErrorString(value string)
	IsWellFormed() bool
}

type NSScriptCommand struct {
	objc.NSObject
}

func MakeScriptCommand(ptr unsafe.Pointer) NSScriptCommand {
	return NSScriptCommand{
		NSObject: objc.MakeObject(ptr),
	}
}

func AllocScriptCommand() NSScriptCommand {
	return MakeScriptCommand(C.C_ScriptCommand_Alloc())
}

func (n NSScriptCommand) InitWithCommandDescription(commandDef ScriptCommandDescription) ScriptCommand {
	result_ := C.C_NSScriptCommand_InitWithCommandDescription(n.Ptr(), objc.ExtractPtr(commandDef))
	return MakeScriptCommand(result_)
}

func (n NSScriptCommand) InitWithCoder(inCoder Coder) ScriptCommand {
	result_ := C.C_NSScriptCommand_InitWithCoder(n.Ptr(), objc.ExtractPtr(inCoder))
	return MakeScriptCommand(result_)
}

func ScriptCommand_CurrentCommand() ScriptCommand {
	result_ := C.C_NSScriptCommand_ScriptCommand_CurrentCommand()
	return MakeScriptCommand(result_)
}

func (n NSScriptCommand) ExecuteCommand() objc.Object {
	result_ := C.C_NSScriptCommand_ExecuteCommand(n.Ptr())
	return objc.MakeObject(result_)
}

func (n NSScriptCommand) PerformDefaultImplementation() objc.Object {
	result_ := C.C_NSScriptCommand_PerformDefaultImplementation(n.Ptr())
	return objc.MakeObject(result_)
}

func (n NSScriptCommand) SuspendExecution() {
	C.C_NSScriptCommand_SuspendExecution(n.Ptr())
}

func (n NSScriptCommand) ResumeExecutionWithResult(result objc.Object) {
	C.C_NSScriptCommand_ResumeExecutionWithResult(n.Ptr(), objc.ExtractPtr(result))
}

func (n NSScriptCommand) AppleEvent() AppleEventDescriptor {
	result_ := C.C_NSScriptCommand_AppleEvent(n.Ptr())
	return MakeAppleEventDescriptor(result_)
}

func (n NSScriptCommand) EvaluatedReceivers() objc.Object {
	result_ := C.C_NSScriptCommand_EvaluatedReceivers(n.Ptr())
	return objc.MakeObject(result_)
}

func (n NSScriptCommand) ReceiversSpecifier() ScriptObjectSpecifier {
	result_ := C.C_NSScriptCommand_ReceiversSpecifier(n.Ptr())
	return MakeScriptObjectSpecifier(result_)
}

func (n NSScriptCommand) SetReceiversSpecifier(value ScriptObjectSpecifier) {
	C.C_NSScriptCommand_SetReceiversSpecifier(n.Ptr(), objc.ExtractPtr(value))
}

func (n NSScriptCommand) DirectParameter() objc.Object {
	result_ := C.C_NSScriptCommand_DirectParameter(n.Ptr())
	return objc.MakeObject(result_)
}

func (n NSScriptCommand) SetDirectParameter(value objc.Object) {
	C.C_NSScriptCommand_SetDirectParameter(n.Ptr(), objc.ExtractPtr(value))
}

func (n NSScriptCommand) CommandDescription() ScriptCommandDescription {
	result_ := C.C_NSScriptCommand_CommandDescription(n.Ptr())
	return MakeScriptCommandDescription(result_)
}

func (n NSScriptCommand) ScriptErrorExpectedTypeDescriptor() AppleEventDescriptor {
	result_ := C.C_NSScriptCommand_ScriptErrorExpectedTypeDescriptor(n.Ptr())
	return MakeAppleEventDescriptor(result_)
}

func (n NSScriptCommand) SetScriptErrorExpectedTypeDescriptor(value AppleEventDescriptor) {
	C.C_NSScriptCommand_SetScriptErrorExpectedTypeDescriptor(n.Ptr(), objc.ExtractPtr(value))
}

func (n NSScriptCommand) ScriptErrorNumber() int {
	result_ := C.C_NSScriptCommand_ScriptErrorNumber(n.Ptr())
	return int(result_)
}

func (n NSScriptCommand) SetScriptErrorNumber(value int) {
	C.C_NSScriptCommand_SetScriptErrorNumber(n.Ptr(), C.int(value))
}

func (n NSScriptCommand) ScriptErrorOffendingObjectDescriptor() AppleEventDescriptor {
	result_ := C.C_NSScriptCommand_ScriptErrorOffendingObjectDescriptor(n.Ptr())
	return MakeAppleEventDescriptor(result_)
}

func (n NSScriptCommand) SetScriptErrorOffendingObjectDescriptor(value AppleEventDescriptor) {
	C.C_NSScriptCommand_SetScriptErrorOffendingObjectDescriptor(n.Ptr(), objc.ExtractPtr(value))
}

func (n NSScriptCommand) ScriptErrorString() string {
	result_ := C.C_NSScriptCommand_ScriptErrorString(n.Ptr())
	return MakeString(result_).String()
}

func (n NSScriptCommand) SetScriptErrorString(value string) {
	C.C_NSScriptCommand_SetScriptErrorString(n.Ptr(), NewString(value).Ptr())
}

func (n NSScriptCommand) IsWellFormed() bool {
	result_ := C.C_NSScriptCommand_IsWellFormed(n.Ptr())
	return bool(result_)
}
