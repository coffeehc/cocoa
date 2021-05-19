package appkit

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
	result_ := C.C_NSWindowTab_Init(n.Ptr())
	return MakeWindowTab(result_)
}

func (n *NSWindowTab) Title() string {
	result_ := C.C_NSWindowTab_Title(n.Ptr())
	return foundation.MakeString(result_).String()
}

func (n *NSWindowTab) SetTitle(value string) {
	C.C_NSWindowTab_SetTitle(n.Ptr(), foundation.NewString(value).Ptr())
}

func (n *NSWindowTab) AttributedTitle() foundation.AttributedString {
	result_ := C.C_NSWindowTab_AttributedTitle(n.Ptr())
	return foundation.MakeAttributedString(result_)
}

func (n *NSWindowTab) SetAttributedTitle(value foundation.AttributedString) {
	C.C_NSWindowTab_SetAttributedTitle(n.Ptr(), objc.ExtractPtr(value))
}

func (n *NSWindowTab) ToolTip() string {
	result_ := C.C_NSWindowTab_ToolTip(n.Ptr())
	return foundation.MakeString(result_).String()
}

func (n *NSWindowTab) SetToolTip(value string) {
	C.C_NSWindowTab_SetToolTip(n.Ptr(), foundation.NewString(value).Ptr())
}

func (n *NSWindowTab) AccessoryView() View {
	result_ := C.C_NSWindowTab_AccessoryView(n.Ptr())
	return MakeView(result_)
}

func (n *NSWindowTab) SetAccessoryView(value View) {
	C.C_NSWindowTab_SetAccessoryView(n.Ptr(), objc.ExtractPtr(value))
}
