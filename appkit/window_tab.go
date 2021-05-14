package appkit

// #cgo CFLAGS: -x objective-c
// #cgo LDFLAGS: -framework AppKit
// #include "window_tab.h"
import "C"
import (
	"github.com/hsiafan/cocoa/foundation"
	"github.com/hsiafan/cocoa/objc"
	"unsafe"
)

type WindowTab interface {
	objc.Object
	Title() string
	SetTitle(value string)
	AttributedTitle() foundation.AttributedString
	SetAttributedTitle(value foundation.AttributedString)
	ToolTip() string
	SetToolTip(value string)
	AccessoryView() View
	SetAccessoryView(value View)
}

type NSWindowTab struct {
	objc.NSObject
}

func MakeWindowTab(ptr unsafe.Pointer) *NSWindowTab {
	if ptr == nil {
		return nil
	}
	return &NSWindowTab{
		NSObject: *objc.MakeObject(ptr),
	}
}

func AllocWindowTab() *NSWindowTab {
	return MakeWindowTab(C.C_WindowTab_Alloc())
}

func (n *NSWindowTab) Init() WindowTab {
	result := C.C_NSWindowTab_Init(n.Ptr())
	return MakeWindowTab(result)
}

func (n *NSWindowTab) Title() string {
	result := C.C_NSWindowTab_Title(n.Ptr())
	return foundation.MakeString(result).String()
}

func (n *NSWindowTab) SetTitle(value string) {
	C.C_NSWindowTab_SetTitle(n.Ptr(), foundation.NewString(value).Ptr())
}

func (n *NSWindowTab) AttributedTitle() foundation.AttributedString {
	result := C.C_NSWindowTab_AttributedTitle(n.Ptr())
	return foundation.MakeAttributedString(result)
}

func (n *NSWindowTab) SetAttributedTitle(value foundation.AttributedString) {
	C.C_NSWindowTab_SetAttributedTitle(n.Ptr(), objc.ExtractPtr(value))
}

func (n *NSWindowTab) ToolTip() string {
	result := C.C_NSWindowTab_ToolTip(n.Ptr())
	return foundation.MakeString(result).String()
}

func (n *NSWindowTab) SetToolTip(value string) {
	C.C_NSWindowTab_SetToolTip(n.Ptr(), foundation.NewString(value).Ptr())
}

func (n *NSWindowTab) AccessoryView() View {
	result := C.C_NSWindowTab_AccessoryView(n.Ptr())
	return MakeView(result)
}

func (n *NSWindowTab) SetAccessoryView(value View) {
	C.C_NSWindowTab_SetAccessoryView(n.Ptr(), objc.ExtractPtr(value))
}
