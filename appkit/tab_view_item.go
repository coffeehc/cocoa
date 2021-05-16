package appkit

// #cgo CFLAGS: -x objective-c
// #cgo LDFLAGS: -framework AppKit
// #include "tab_view_item.h"
import "C"
import (
	"github.com/hsiafan/cocoa/coregraphics"
	"github.com/hsiafan/cocoa/foundation"
	"github.com/hsiafan/cocoa/objc"
	"unsafe"
)

type TabViewItem interface {
	objc.Object
	DrawLabel_InRect(shouldTruncateLabel bool, labelRect foundation.Rect)
	SizeOfLabel(computeMin bool) foundation.Size
	Label() string
	SetLabel(value string)
	TabState() TabState
	Identifier() objc.Object
	SetIdentifier(value objc.Object)
	Color() Color
	SetColor(value Color)
	View() View
	SetView(value View)
	InitialFirstResponder() View
	SetInitialFirstResponder(value View)
	TabView() TabView
	ToolTip() string
	SetToolTip(value string)
	Image() Image
	SetImage(value Image)
	ViewController() ViewController
	SetViewController(value ViewController)
}

type NSTabViewItem struct {
	objc.NSObject
}

func MakeTabViewItem(ptr unsafe.Pointer) *NSTabViewItem {
	if ptr == nil {
		return nil
	}
	return &NSTabViewItem{
		NSObject: *objc.MakeObject(ptr),
	}
}

func AllocTabViewItem() *NSTabViewItem {
	return MakeTabViewItem(C.C_TabViewItem_Alloc())
}

func (n *NSTabViewItem) InitWithIdentifier(identifier objc.Object) TabViewItem {
	result := C.C_NSTabViewItem_InitWithIdentifier(n.Ptr(), objc.ExtractPtr(identifier))
	return MakeTabViewItem(result)
}

func (n *NSTabViewItem) Init() TabViewItem {
	result := C.C_NSTabViewItem_Init(n.Ptr())
	return MakeTabViewItem(result)
}

func (n *NSTabViewItem) DrawLabel_InRect(shouldTruncateLabel bool, labelRect foundation.Rect) {
	C.C_NSTabViewItem_DrawLabel_InRect(n.Ptr(), C.bool(shouldTruncateLabel), *(*C.CGRect)(coregraphics.ToCGRectPointer(coregraphics.Rect(labelRect))))
}

func (n *NSTabViewItem) SizeOfLabel(computeMin bool) foundation.Size {
	result := C.C_NSTabViewItem_SizeOfLabel(n.Ptr(), C.bool(computeMin))
	return foundation.Size(coregraphics.FromCGSizePointer(unsafe.Pointer(&result)))
}

func TabViewItemWithViewController(viewController ViewController) TabViewItem {
	result := C.C_NSTabViewItem_TabViewItemWithViewController(objc.ExtractPtr(viewController))
	return MakeTabViewItem(result)
}

func (n *NSTabViewItem) Label() string {
	result := C.C_NSTabViewItem_Label(n.Ptr())
	return foundation.MakeString(result).String()
}

func (n *NSTabViewItem) SetLabel(value string) {
	C.C_NSTabViewItem_SetLabel(n.Ptr(), foundation.NewString(value).Ptr())
}

func (n *NSTabViewItem) TabState() TabState {
	result := C.C_NSTabViewItem_TabState(n.Ptr())
	return TabState(uint(result))
}

func (n *NSTabViewItem) Identifier() objc.Object {
	result := C.C_NSTabViewItem_Identifier(n.Ptr())
	return objc.MakeObject(result)
}

func (n *NSTabViewItem) SetIdentifier(value objc.Object) {
	C.C_NSTabViewItem_SetIdentifier(n.Ptr(), objc.ExtractPtr(value))
}

func (n *NSTabViewItem) Color() Color {
	result := C.C_NSTabViewItem_Color(n.Ptr())
	return MakeColor(result)
}

func (n *NSTabViewItem) SetColor(value Color) {
	C.C_NSTabViewItem_SetColor(n.Ptr(), objc.ExtractPtr(value))
}

func (n *NSTabViewItem) View() View {
	result := C.C_NSTabViewItem_View(n.Ptr())
	return MakeView(result)
}

func (n *NSTabViewItem) SetView(value View) {
	C.C_NSTabViewItem_SetView(n.Ptr(), objc.ExtractPtr(value))
}

func (n *NSTabViewItem) InitialFirstResponder() View {
	result := C.C_NSTabViewItem_InitialFirstResponder(n.Ptr())
	return MakeView(result)
}

func (n *NSTabViewItem) SetInitialFirstResponder(value View) {
	C.C_NSTabViewItem_SetInitialFirstResponder(n.Ptr(), objc.ExtractPtr(value))
}

func (n *NSTabViewItem) TabView() TabView {
	result := C.C_NSTabViewItem_TabView(n.Ptr())
	return MakeTabView(result)
}

func (n *NSTabViewItem) ToolTip() string {
	result := C.C_NSTabViewItem_ToolTip(n.Ptr())
	return foundation.MakeString(result).String()
}

func (n *NSTabViewItem) SetToolTip(value string) {
	C.C_NSTabViewItem_SetToolTip(n.Ptr(), foundation.NewString(value).Ptr())
}

func (n *NSTabViewItem) Image() Image {
	result := C.C_NSTabViewItem_Image(n.Ptr())
	return MakeImage(result)
}

func (n *NSTabViewItem) SetImage(value Image) {
	C.C_NSTabViewItem_SetImage(n.Ptr(), objc.ExtractPtr(value))
}

func (n *NSTabViewItem) ViewController() ViewController {
	result := C.C_NSTabViewItem_ViewController(n.Ptr())
	return MakeViewController(result)
}

func (n *NSTabViewItem) SetViewController(value ViewController) {
	C.C_NSTabViewItem_SetViewController(n.Ptr(), objc.ExtractPtr(value))
}
