package appkit

// #cgo CFLAGS: -x objective-c
// #cgo LDFLAGS: -framework AppKit
// #include "tab_view_item.h"
import "C"
import (
	"github.com/hsiafan/cocoa/objc"
	"unsafe"
)

type TabViewItem interface {
	objc.Object
	Label() string
	SetLabel(label string)
	ToolTip() string
	SetToolTip(toolTip string)
	TabState() TabState
	Identifier() objc.Object
	SetIdentifier(identifier objc.Object)
	Color() Color
	SetColor(color Color)
	View() View
	SetView(view View)
	InitialFirstResponder() View
	SetInitialFirstResponder(initialFirstResponder View)
	TabView() TabView
}

var _ TabViewItem = (*NSTabViewItem)(nil)

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

func (t *NSTabViewItem) Label() string {
	return C.GoString(C.TabViewItem_Label(t.Ptr()))
}

func (t *NSTabViewItem) SetLabel(label string) {
	cLabel := C.CString(label)
	defer C.free(unsafe.Pointer(cLabel))
	C.TabViewItem_SetLabel(t.Ptr(), cLabel)
}

func (t *NSTabViewItem) ToolTip() string {
	return C.GoString(C.TabViewItem_ToolTip(t.Ptr()))
}

func (t *NSTabViewItem) SetToolTip(toolTip string) {
	cToolTip := C.CString(toolTip)
	defer C.free(unsafe.Pointer(cToolTip))
	C.TabViewItem_SetToolTip(t.Ptr(), cToolTip)
}

func (t *NSTabViewItem) TabState() TabState {
	return TabState(C.TabViewItem_TabState(t.Ptr()))
}

func (t *NSTabViewItem) Identifier() objc.Object {
	return objc.MakeObject(C.TabViewItem_Identifier(t.Ptr()))
}

func (t *NSTabViewItem) SetIdentifier(identifier objc.Object) {
	C.TabViewItem_SetIdentifier(t.Ptr(), toPointer(identifier))
}

func (t *NSTabViewItem) Color() Color {
	return MakeColor(C.TabViewItem_Color(t.Ptr()))
}

func (t *NSTabViewItem) SetColor(color Color) {
	C.TabViewItem_SetColor(t.Ptr(), toPointer(color))
}

func (t *NSTabViewItem) View() View {
	return MakeView(C.TabViewItem_View(t.Ptr()))
}

func (t *NSTabViewItem) SetView(view View) {
	C.TabViewItem_SetView(t.Ptr(), toPointer(view))
}

func (t *NSTabViewItem) InitialFirstResponder() View {
	return MakeView(C.TabViewItem_InitialFirstResponder(t.Ptr()))
}

func (t *NSTabViewItem) SetInitialFirstResponder(initialFirstResponder View) {
	C.TabViewItem_SetInitialFirstResponder(t.Ptr(), toPointer(initialFirstResponder))
}

func (t *NSTabViewItem) TabView() TabView {
	return MakeTabView(C.TabViewItem_TabView(t.Ptr()))
}

func NewTabViewItem(identifier objc.Object) TabViewItem {
	return MakeTabViewItem(C.TabViewItem_NewTabViewItem(toPointer(identifier)))
}
