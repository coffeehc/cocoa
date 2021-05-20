package appkit

// #include "path_control_item.h"
import "C"
import (
	"github.com/hsiafan/cocoa/foundation"
	"github.com/hsiafan/cocoa/objc"
	"unsafe"
)

type PathControlItem interface {
	objc.Object
	URL() foundation.URL
	AttributedTitle() foundation.AttributedString
	SetAttributedTitle(value foundation.AttributedString)
	Image() Image
	SetImage(value Image)
	Title() string
	SetTitle(value string)
}

type NSPathControlItem struct {
	objc.NSObject
}

func MakePathControlItem(ptr unsafe.Pointer) *NSPathControlItem {
	if ptr == nil {
		return nil
	}
	return &NSPathControlItem{
		NSObject: *objc.MakeObject(ptr),
	}
}

func AllocPathControlItem() *NSPathControlItem {
	return MakePathControlItem(C.C_PathControlItem_Alloc())
}

func (n *NSPathControlItem) Init() PathControlItem {
	result_ := C.C_NSPathControlItem_Init(n.Ptr())
	return MakePathControlItem(result_)
}

func (n *NSPathControlItem) URL() foundation.URL {
	result_ := C.C_NSPathControlItem_URL(n.Ptr())
	return foundation.MakeURL(result_)
}

func (n *NSPathControlItem) AttributedTitle() foundation.AttributedString {
	result_ := C.C_NSPathControlItem_AttributedTitle(n.Ptr())
	return foundation.MakeAttributedString(result_)
}

func (n *NSPathControlItem) SetAttributedTitle(value foundation.AttributedString) {
	C.C_NSPathControlItem_SetAttributedTitle(n.Ptr(), objc.ExtractPtr(value))
}

func (n *NSPathControlItem) Image() Image {
	result_ := C.C_NSPathControlItem_Image(n.Ptr())
	return MakeImage(result_)
}

func (n *NSPathControlItem) SetImage(value Image) {
	C.C_NSPathControlItem_SetImage(n.Ptr(), objc.ExtractPtr(value))
}

func (n *NSPathControlItem) Title() string {
	result_ := C.C_NSPathControlItem_Title(n.Ptr())
	return foundation.MakeString(result_).String()
}

func (n *NSPathControlItem) SetTitle(value string) {
	C.C_NSPathControlItem_SetTitle(n.Ptr(), foundation.NewString(value).Ptr())
}
