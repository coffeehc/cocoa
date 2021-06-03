package foundation

// #include "close_command.h"
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

func AllocCloseCommand() NSCloseCommand {
	return MakeCloseCommand(C.C_CloseCommand_Alloc())
}

func (n NSCloseCommand) InitWithCommandDescription(commandDef ScriptCommandDescription) CloseCommand {
	result_ := C.C_NSCloseCommand_InitWithCommandDescription(n.Ptr(), objc.ExtractPtr(commandDef))
	return MakeCloseCommand(result_)
}

func (n NSCloseCommand) InitWithCoder(inCoder Coder) CloseCommand {
	result_ := C.C_NSCloseCommand_InitWithCoder(n.Ptr(), objc.ExtractPtr(inCoder))
	return MakeCloseCommand(result_)
}

func (n NSCloseCommand) SaveOptions() SaveOptions {
	result_ := C.C_NSCloseCommand_SaveOptions(n.Ptr())
	return SaveOptions(uint(result_))
}
