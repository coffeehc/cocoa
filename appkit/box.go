package appkit

// #cgo CFLAGS: -x objective-c
// #cgo LDFLAGS: -framework AppKit
// #include "box.h"
import "C"
import (
	"github.com/hsiafan/cocoa/foundation"
	"github.com/hsiafan/cocoa/objc"
	"unsafe"
)

type Box interface {
	View
	BorderRect() foundation.Rect
	BoxType() BoxType
	SetBoxType(boxType BoxType)
	IsTransparent() bool
	SetTransparent(transparent bool)
	Title() string
	SetTitle(title string)
	TitleFont() Font
	SetTitleFont(titleFont Font)
	TitlePosition() TitlePosition
	SetTitlePosition(titlePosition TitlePosition)
	TitleCell() objc.Object
	TitleRect() foundation.Rect
	BorderColor() Color
	SetBorderColor(borderColor Color)
	BorderWidth() float64
	SetBorderWidth(borderWidth float64)
	CornerRadius() float64
	SetCornerRadius(cornerRadius float64)
	FillColor() Color
	SetFillColor(fillColor Color)
	ContentView() View
	SetContentView(contentView View)
	ContentViewMargins() foundation.Size
	SetContentViewMargins(contentViewMargins foundation.Size)
	SetFrameFromContentFrame(contentFrame foundation.Rect)
	SizeToFit()
}

var _ Box = (*NSBox)(nil)

type NSBox struct {
	NSView
}

func MakeBox(ptr unsafe.Pointer) *NSBox {
	if ptr == nil {
		return nil
	}
	return &NSBox{
		NSView: *MakeView(ptr),
	}
}

func (b *NSBox) BorderRect() foundation.Rect {
	return toRect(C.Box_BorderRect(b.Ptr()))
}

func (b *NSBox) BoxType() BoxType {
	return BoxType(C.Box_BoxType(b.Ptr()))
}

func (b *NSBox) SetBoxType(boxType BoxType) {
	C.Box_SetBoxType(b.Ptr(), C.ulong(boxType))
}

func (b *NSBox) IsTransparent() bool {
	return bool(C.Box_IsTransparent(b.Ptr()))
}

func (b *NSBox) SetTransparent(transparent bool) {
	C.Box_SetTransparent(b.Ptr(), C.bool(transparent))
}

func (b *NSBox) Title() string {
	return C.GoString(C.Box_Title(b.Ptr()))
}

func (b *NSBox) SetTitle(title string) {
	cTitle := C.CString(title)
	defer C.free(unsafe.Pointer(cTitle))
	C.Box_SetTitle(b.Ptr(), cTitle)
}

func (b *NSBox) TitleFont() Font {
	return MakeFont(C.Box_TitleFont(b.Ptr()))
}

func (b *NSBox) SetTitleFont(titleFont Font) {
	C.Box_SetTitleFont(b.Ptr(), toPointer(titleFont))
}

func (b *NSBox) TitlePosition() TitlePosition {
	return TitlePosition(C.Box_TitlePosition(b.Ptr()))
}

func (b *NSBox) SetTitlePosition(titlePosition TitlePosition) {
	C.Box_SetTitlePosition(b.Ptr(), C.ulong(titlePosition))
}

func (b *NSBox) TitleCell() objc.Object {
	return objc.MakeObject(C.Box_TitleCell(b.Ptr()))
}

func (b *NSBox) TitleRect() foundation.Rect {
	return toRect(C.Box_TitleRect(b.Ptr()))
}

func (b *NSBox) BorderColor() Color {
	return MakeColor(C.Box_BorderColor(b.Ptr()))
}

func (b *NSBox) SetBorderColor(borderColor Color) {
	C.Box_SetBorderColor(b.Ptr(), toPointer(borderColor))
}

func (b *NSBox) BorderWidth() float64 {
	return float64(C.Box_BorderWidth(b.Ptr()))
}

func (b *NSBox) SetBorderWidth(borderWidth float64) {
	C.Box_SetBorderWidth(b.Ptr(), C.double(borderWidth))
}

func (b *NSBox) CornerRadius() float64 {
	return float64(C.Box_CornerRadius(b.Ptr()))
}

func (b *NSBox) SetCornerRadius(cornerRadius float64) {
	C.Box_SetCornerRadius(b.Ptr(), C.double(cornerRadius))
}

func (b *NSBox) FillColor() Color {
	return MakeColor(C.Box_FillColor(b.Ptr()))
}

func (b *NSBox) SetFillColor(fillColor Color) {
	C.Box_SetFillColor(b.Ptr(), toPointer(fillColor))
}

func (b *NSBox) ContentView() View {
	return MakeView(C.Box_ContentView(b.Ptr()))
}

func (b *NSBox) SetContentView(contentView View) {
	C.Box_SetContentView(b.Ptr(), toPointer(contentView))
}

func (b *NSBox) ContentViewMargins() foundation.Size {
	return toSize(C.Box_ContentViewMargins(b.Ptr()))
}

func (b *NSBox) SetContentViewMargins(contentViewMargins foundation.Size) {
	C.Box_SetContentViewMargins(b.Ptr(), toNSSize(contentViewMargins))
}

func NewBox(frame foundation.Rect) Box {
	return MakeBox(C.Box_NewBox(*(*C.NSRect)(foundation.ToNSRectPointer(frame))))
}

func (b *NSBox) SetFrameFromContentFrame(contentFrame foundation.Rect) {
	C.Box_SetFrameFromContentFrame(b.Ptr(), toNSRect(contentFrame))
}

func (b *NSBox) SizeToFit() {
	C.Box_SizeToFit(b.Ptr())
}
