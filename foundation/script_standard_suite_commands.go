package foundation

// #include "script_standard_suite_commands.h"
import "C"
import (
	"github.com/hsiafan/cocoa/objc"
	"unsafe"
)

type CloseCommand interface {
	ScriptCommand
	SaveOptions() SaveOptions
}

type NSCloseCommand struct {
	NSScriptCommand
}

func MakeCloseCommand(ptr unsafe.Pointer) NSCloseCommand {
	return NSCloseCommand{
		NSScriptCommand: MakeScriptCommand(ptr),
	}
}

func (n NSCloseCommand) InitWithCommandDescription(commandDef ScriptCommandDescription) NSCloseCommand {
	result_ := C.C_NSCloseCommand_InitWithCommandDescription(n.Ptr(), objc.ExtractPtr(commandDef))
	return MakeCloseCommand(result_)
}

func (n NSCloseCommand) InitWithCoder(inCoder Coder) NSCloseCommand {
	result_ := C.C_NSCloseCommand_InitWithCoder(n.Ptr(), objc.ExtractPtr(inCoder))
	return MakeCloseCommand(result_)
}

func AllocCloseCommand() NSCloseCommand {
	result_ := C.C_NSCloseCommand_AllocCloseCommand()
	return MakeCloseCommand(result_)
}

func (n NSCloseCommand) Autorelease() NSCloseCommand {
	result_ := C.C_NSCloseCommand_Autorelease(n.Ptr())
	return MakeCloseCommand(result_)
}

func (n NSCloseCommand) Retain() NSCloseCommand {
	result_ := C.C_NSCloseCommand_Retain(n.Ptr())
	return MakeCloseCommand(result_)
}

func (n NSCloseCommand) SaveOptions() SaveOptions {
	result_ := C.C_NSCloseCommand_SaveOptions(n.Ptr())
	return SaveOptions(uint(result_))
}
