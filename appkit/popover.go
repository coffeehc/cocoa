package appkit

// #cgo CFLAGS: -x objective-c
// #cgo LDFLAGS: -framework AppKit
// #include "popover.h"
import "C"
import (
	"github.com/hsiafan/cocoa/foundation"
	"github.com/hsiafan/cocoa/objc"
	"unsafe"
)

type Popover interface {
	Responder
	Behavior() PopoverBehavior
	SetBehavior(behavior PopoverBehavior)
	PositioningRect() foundation.Rect
	SetPositioningRect(positioningRect foundation.Rect)
	Animates() bool
	SetAnimates(animates bool)
	ContentSize() foundation.Size
	SetContentSize(contentSize foundation.Size)
	IsShown() bool
	IsDetached() bool
	Appearance() Appearance
	SetAppearance(appearance Appearance)
	EffectiveAppearance() Appearance
	PerformClose(sender objc.Object)
	Close()
	ShowRelativeTo(positioningRect foundation.Rect, positioningView View, preferredEdge foundation.RectEdge)
}

var _ Popover = (*NSPopover)(nil)

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

func (p *NSPopover) Behavior() PopoverBehavior {
	return PopoverBehavior(C.Popover_Behavior(p.Ptr()))
}

func (p *NSPopover) SetBehavior(behavior PopoverBehavior) {
	C.Popover_SetBehavior(p.Ptr(), C.long(behavior))
}

func (p *NSPopover) PositioningRect() foundation.Rect {
	return toRect(C.Popover_PositioningRect(p.Ptr()))
}

func (p *NSPopover) SetPositioningRect(positioningRect foundation.Rect) {
	C.Popover_SetPositioningRect(p.Ptr(), toNSRect(positioningRect))
}

func (p *NSPopover) Animates() bool {
	return bool(C.Popover_Animates(p.Ptr()))
}

func (p *NSPopover) SetAnimates(animates bool) {
	C.Popover_SetAnimates(p.Ptr(), C.bool(animates))
}

func (p *NSPopover) ContentSize() foundation.Size {
	return toSize(C.Popover_ContentSize(p.Ptr()))
}

func (p *NSPopover) SetContentSize(contentSize foundation.Size) {
	C.Popover_SetContentSize(p.Ptr(), toNSSize(contentSize))
}

func (p *NSPopover) IsShown() bool {
	return bool(C.Popover_IsShown(p.Ptr()))
}

func (p *NSPopover) IsDetached() bool {
	return bool(C.Popover_IsDetached(p.Ptr()))
}

func (p *NSPopover) Appearance() Appearance {
	return MakeAppearance(C.Popover_Appearance(p.Ptr()))
}

func (p *NSPopover) SetAppearance(appearance Appearance) {
	C.Popover_SetAppearance(p.Ptr(), toPointer(appearance))
}

func (p *NSPopover) EffectiveAppearance() Appearance {
	return MakeAppearance(C.Popover_EffectiveAppearance(p.Ptr()))
}

func NewPopover() Popover {
	return MakePopover(C.Popover_NewPopover())
}

func (p *NSPopover) PerformClose(sender objc.Object) {
	C.Popover_PerformClose(p.Ptr(), toPointer(sender))
}

func (p *NSPopover) Close() {
	C.Popover_Close(p.Ptr())
}

func (p *NSPopover) ShowRelativeTo(positioningRect foundation.Rect, positioningView View, preferredEdge foundation.RectEdge) {
	C.Popover_ShowRelativeTo(p.Ptr(), toNSRect(positioningRect), toPointer(positioningView), C.long(preferredEdge))
}
