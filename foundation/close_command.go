package foundation

// #cgo CFLAGS: -x objective-c
// #cgo LDFLAGS: -framework Foundation
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

func MakeCloseCommand(ptr unsafe.Pointer) *NSCloseCommand {
	if ptr == nil {
		return nil
	}
	return &NSCloseCommand{
		NSScriptCommand: *MakeScriptCommand(ptr),
	}
}

func AllocCloseCommand() *NSCloseCommand {
	return MakeCloseCommand(C.C_CloseCommand_Alloc())
}

func (n *NSCloseCommand) InitWithCommandDescription(commandDef ScriptCommandDescription) CloseCommand {
	result := C.C_NSCloseCommand_InitWithCommandDescription(n.Ptr(), objc.ExtractPtr(commandDef))
	return MakeCloseCommand(result)
}

func (n *NSCloseCommand) InitWithCoder(inCoder Coder) CloseCommand {
	result := C.C_NSCloseCommand_InitWithCoder(n.Ptr(), objc.ExtractPtr(inCoder))
	return MakeCloseCommand(result)
}

func (n *NSCloseCommand) SaveOptions() SaveOptions {
	result := C.C_NSCloseCommand_SaveOptions(n.Ptr())
	return SaveOptions(uint(result))
}
