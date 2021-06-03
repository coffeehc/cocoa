package appkit

// #include "table_view_row_action.h"
import "C"
import (
	"github.com/hsiafan/cocoa/foundation"
	"github.com/hsiafan/cocoa/objc"
	"unsafe"
)

type TableViewRowAction interface {
	objc.Object
	Style() TableViewRowActionStyle
	Title() string
	SetTitle(value string)
	BackgroundColor() Color
	SetBackgroundColor(value Color)
	Image() Image
	SetImage(value Image)
}

type NSTableViewRowAction struct {
	objc.NSObject
}

func MakeTableViewRowAction(ptr unsafe.Pointer) NSTableViewRowAction {
	return NSTableViewRowAction{
		NSObject: objc.MakeObject(ptr),
	}
}

func AllocTableViewRowAction() NSTableViewRowAction {
	return MakeTableViewRowAction(C.C_TableViewRowAction_Alloc())
}

func (n NSTableViewRowAction) Init() TableViewRowAction {
	result_ := C.C_NSTableViewRowAction_Init(n.Ptr())
	return MakeTableViewRowAction(result_)
}

func (n NSTableViewRowAction) Style() TableViewRowActionStyle {
	result_ := C.C_NSTableViewRowAction_Style(n.Ptr())
	return TableViewRowActionStyle(int(result_))
}

func (n NSTableViewRowAction) Title() string {
	result_ := C.C_NSTableViewRowAction_Title(n.Ptr())
	return foundation.MakeString(result_).String()
}

func (n NSTableViewRowAction) SetTitle(value string) {
	C.C_NSTableViewRowAction_SetTitle(n.Ptr(), foundation.NewString(value).Ptr())
}

func (n NSTableViewRowAction) BackgroundColor() Color {
	result_ := C.C_NSTableViewRowAction_BackgroundColor(n.Ptr())
	return MakeColor(result_)
}

func (n NSTableViewRowAction) SetBackgroundColor(value Color) {
	C.C_NSTableViewRowAction_SetBackgroundColor(n.Ptr(), objc.ExtractPtr(value))
}

func (n NSTableViewRowAction) Image() Image {
	result_ := C.C_NSTableViewRowAction_Image(n.Ptr())
	return MakeImage(result_)
}

func (n NSTableViewRowAction) SetImage(value Image) {
	C.C_NSTableViewRowAction_SetImage(n.Ptr(), objc.ExtractPtr(value))
}
