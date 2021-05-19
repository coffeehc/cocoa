package appkit

// #include "popover.h"
import "C"
import (
	"github.com/hsiafan/cocoa/coregraphics"
	"github.com/hsiafan/cocoa/foundation"
	"github.com/hsiafan/cocoa/objc"
	"unsafe"
)

type Popover interface {
	Responder
	ShowRelativeToRect_OfView_PreferredEdge(positioningRect foundation.Rect, positioningView View, preferredEdge foundation.RectEdge)
	PerformClose(sender objc.Object)
	Close()
	Behavior() PopoverBehavior
	SetBehavior(value PopoverBehavior)
	PositioningRect() foundation.Rect
	SetPositioningRect(value foundation.Rect)
	Appearance() Appearance
	SetAppearance(value Appearance)
	EffectiveAppearance() Appearance
	Animates() bool
	SetAnimates(value bool)
	ContentSize() foundation.Size
	SetContentSize(value foundation.Size)
	IsShown() bool
	IsDetached() bool
}

type NSPopover struct {
	NSResponder
}

func MakePopover(ptr unsafe.Pointer) *NSPopover {
	if ptr == nil {
		return nil
	}
	return &NSPopover{
		NSResponder: *MakeResponder(ptr),
	}
}

func AllocPopover() *NSPopover {
	return MakePopover(C.C_Popover_Alloc())
}

func (n *NSPopover) Init() Popover {
	result_ := C.C_NSPopover_Init(n.Ptr())
	return MakePopover(result_)
}

func (n *NSPopover) InitWithCoder(coder foundation.Coder) Popover {
	result_ := C.C_NSPopover_InitWithCoder(n.Ptr(), objc.ExtractPtr(coder))
	return MakePopover(result_)
}

func (n *NSPopover) ShowRelativeToRect_OfView_PreferredEdge(positioningRect foundation.Rect, positioningView View, preferredEdge foundation.RectEdge) {
	C.C_NSPopover_ShowRelativeToRect_OfView_PreferredEdge(n.Ptr(), *(*C.CGRect)(coregraphics.ToCGRectPointer(coregraphics.Rect(positioningRect))), objc.ExtractPtr(positioningView), C.uint(uint(preferredEdge)))
}

func (n *NSPopover) PerformClose(sender objc.Object) {
	C.C_NSPopover_PerformClose(n.Ptr(), objc.ExtractPtr(sender))
}

func (n *NSPopover) Close() {
	C.C_NSPopover_Close(n.Ptr())
}

func (n *NSPopover) Behavior() PopoverBehavior {
	result_ := C.C_NSPopover_Behavior(n.Ptr())
	return PopoverBehavior(int(result_))
}

func (n *NSPopover) SetBehavior(value PopoverBehavior) {
	C.C_NSPopover_SetBehavior(n.Ptr(), C.int(int(value)))
}

func (n *NSPopover) PositioningRect() foundation.Rect {
	result_ := C.C_NSPopover_PositioningRect(n.Ptr())
	return foundation.Rect(coregraphics.FromCGRectPointer(unsafe.Pointer(&result_)))
}

func (n *NSPopover) SetPositioningRect(value foundation.Rect) {
	C.C_NSPopover_SetPositioningRect(n.Ptr(), *(*C.CGRect)(coregraphics.ToCGRectPointer(coregraphics.Rect(value))))
}

func (n *NSPopover) Appearance() Appearance {
	result_ := C.C_NSPopover_Appearance(n.Ptr())
	return MakeAppearance(result_)
}

func (n *NSPopover) SetAppearance(value Appearance) {
	C.C_NSPopover_SetAppearance(n.Ptr(), objc.ExtractPtr(value))
}

func (n *NSPopover) EffectiveAppearance() Appearance {
	result_ := C.C_NSPopover_EffectiveAppearance(n.Ptr())
	return MakeAppearance(result_)
}

func (n *NSPopover) Animates() bool {
	result_ := C.C_NSPopover_Animates(n.Ptr())
	return bool(result_)
}

func (n *NSPopover) SetAnimates(value bool) {
	C.C_NSPopover_SetAnimates(n.Ptr(), C.bool(value))
}

func (n *NSPopover) ContentSize() foundation.Size {
	result_ := C.C_NSPopover_ContentSize(n.Ptr())
	return foundation.Size(coregraphics.FromCGSizePointer(unsafe.Pointer(&result_)))
}

func (n *NSPopover) SetContentSize(value foundation.Size) {
	C.C_NSPopover_SetContentSize(n.Ptr(), *(*C.CGSize)(coregraphics.ToCGSizePointer(coregraphics.Size(value))))
}

func (n *NSPopover) IsShown() bool {
	result_ := C.C_NSPopover_IsShown(n.Ptr())
	return bool(result_)
}

func (n *NSPopover) IsDetached() bool {
	result_ := C.C_NSPopover_IsDetached(n.Ptr())
	return bool(result_)
}
