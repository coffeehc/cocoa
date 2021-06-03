package appkit

// #include "dock_tile.h"
import "C"
import (
	"github.com/hsiafan/cocoa/coregraphics"
	"github.com/hsiafan/cocoa/foundation"
	"github.com/hsiafan/cocoa/objc"
	"unsafe"
)

type DockTile interface {
	objc.Object
	Display()
	ContentView() View
	SetContentView(value View)
	Size() foundation.Size
	Owner() objc.Object
	ShowsApplicationBadge() bool
	SetShowsApplicationBadge(value bool)
	BadgeLabel() string
	SetBadgeLabel(value string)
}

type NSDockTile struct {
	objc.NSObject
}

func MakeDockTile(ptr unsafe.Pointer) NSDockTile {
	return NSDockTile{
		NSObject: objc.MakeObject(ptr),
	}
}

func AllocDockTile() NSDockTile {
	return MakeDockTile(C.C_DockTile_Alloc())
}

func (n NSDockTile) Init() DockTile {
	result_ := C.C_NSDockTile_Init(n.Ptr())
	return MakeDockTile(result_)
}

func (n NSDockTile) Display() {
	C.C_NSDockTile_Display(n.Ptr())
}

func (n NSDockTile) ContentView() View {
	result_ := C.C_NSDockTile_ContentView(n.Ptr())
	return MakeView(result_)
}

func (n NSDockTile) SetContentView(value View) {
	C.C_NSDockTile_SetContentView(n.Ptr(), objc.ExtractPtr(value))
}

func (n NSDockTile) Size() foundation.Size {
	result_ := C.C_NSDockTile_Size(n.Ptr())
	return foundation.Size(coregraphics.FromCGSizePointer(unsafe.Pointer(&result_)))
}

func (n NSDockTile) Owner() objc.Object {
	result_ := C.C_NSDockTile_Owner(n.Ptr())
	return objc.MakeObject(result_)
}

func (n NSDockTile) ShowsApplicationBadge() bool {
	result_ := C.C_NSDockTile_ShowsApplicationBadge(n.Ptr())
	return bool(result_)
}

func (n NSDockTile) SetShowsApplicationBadge(value bool) {
	C.C_NSDockTile_SetShowsApplicationBadge(n.Ptr(), C.bool(value))
}

func (n NSDockTile) BadgeLabel() string {
	result_ := C.C_NSDockTile_BadgeLabel(n.Ptr())
	return foundation.MakeString(result_).String()
}

func (n NSDockTile) SetBadgeLabel(value string) {
	C.C_NSDockTile_SetBadgeLabel(n.Ptr(), foundation.NewString(value).Ptr())
}
