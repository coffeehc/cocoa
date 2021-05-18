package appkit

// #cgo CFLAGS: -x objective-c
// #cgo LDFLAGS: -framework AppKit
// #include "box.h"
import "C"
import (
	"github.com/hsiafan/cocoa/coregraphics"
	"github.com/hsiafan/cocoa/foundation"
	"github.com/hsiafan/cocoa/objc"
	"unsafe"
)

type Box interface {
	View
	SetFrameFromContentFrame(contentFrame foundation.Rect)
	SizeToFit()
	BorderRect() foundation.Rect
	BoxType() BoxType
	SetBoxType(value BoxType)
	IsTransparent() bool
	SetTransparent(value bool)
	Title() string
	SetTitle(value string)
	TitleFont() Font
	SetTitleFont(value Font)
	TitlePosition() TitlePosition
	SetTitlePosition(value TitlePosition)
	TitleCell() objc.Object
	TitleRect() foundation.Rect
	BorderColor() Color
	SetBorderColor(value Color)
	BorderWidth() coregraphics.Float
	SetBorderWidth(value coregraphics.Float)
	CornerRadius() coregraphics.Float
	SetCornerRadius(value coregraphics.Float)
	FillColor() Color
	SetFillColor(value Color)
	ContentView() View
	SetContentView(value View)
	ContentViewMargins() foundation.Size
	SetContentViewMargins(value foundation.Size)
}

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

func AllocBox() *NSBox {
	return MakeBox(C.C_Box_Alloc())
}

func (n *NSBox) InitWithFrame(frameRect foundation.Rect) Box {
	result_ := C.C_NSBox_InitWithFrame(n.Ptr(), *(*C.CGRect)(coregraphics.ToCGRectPointer(coregraphics.Rect(frameRect))))
	return MakeBox(result_)
}

func (n *NSBox) InitWithCoder(coder foundation.Coder) Box {
	result_ := C.C_NSBox_InitWithCoder(n.Ptr(), objc.ExtractPtr(coder))
	return MakeBox(result_)
}

func (n *NSBox) Init() Box {
	result_ := C.C_NSBox_Init(n.Ptr())
	return MakeBox(result_)
}

func (n *NSBox) SetFrameFromContentFrame(contentFrame foundation.Rect) {
	C.C_NSBox_SetFrameFromContentFrame(n.Ptr(), *(*C.CGRect)(coregraphics.ToCGRectPointer(coregraphics.Rect(contentFrame))))
}

func (n *NSBox) SizeToFit() {
	C.C_NSBox_SizeToFit(n.Ptr())
}

func (n *NSBox) BorderRect() foundation.Rect {
	result_ := C.C_NSBox_BorderRect(n.Ptr())
	return foundation.Rect(coregraphics.FromCGRectPointer(unsafe.Pointer(&result_)))
}

func (n *NSBox) BoxType() BoxType {
	result_ := C.C_NSBox_BoxType(n.Ptr())
	return BoxType(uint(result_))
}

func (n *NSBox) SetBoxType(value BoxType) {
	C.C_NSBox_SetBoxType(n.Ptr(), C.uint(uint(value)))
}

func (n *NSBox) IsTransparent() bool {
	result_ := C.C_NSBox_IsTransparent(n.Ptr())
	return bool(result_)
}

func (n *NSBox) SetTransparent(value bool) {
	C.C_NSBox_SetTransparent(n.Ptr(), C.bool(value))
}

func (n *NSBox) Title() string {
	result_ := C.C_NSBox_Title(n.Ptr())
	return foundation.MakeString(result_).String()
}

func (n *NSBox) SetTitle(value string) {
	C.C_NSBox_SetTitle(n.Ptr(), foundation.NewString(value).Ptr())
}

func (n *NSBox) TitleFont() Font {
	result_ := C.C_NSBox_TitleFont(n.Ptr())
	return MakeFont(result_)
}

func (n *NSBox) SetTitleFont(value Font) {
	C.C_NSBox_SetTitleFont(n.Ptr(), objc.ExtractPtr(value))
}

func (n *NSBox) TitlePosition() TitlePosition {
	result_ := C.C_NSBox_TitlePosition(n.Ptr())
	return TitlePosition(uint(result_))
}

func (n *NSBox) SetTitlePosition(value TitlePosition) {
	C.C_NSBox_SetTitlePosition(n.Ptr(), C.uint(uint(value)))
}

func (n *NSBox) TitleCell() objc.Object {
	result_ := C.C_NSBox_TitleCell(n.Ptr())
	return objc.MakeObject(result_)
}

func (n *NSBox) TitleRect() foundation.Rect {
	result_ := C.C_NSBox_TitleRect(n.Ptr())
	return foundation.Rect(coregraphics.FromCGRectPointer(unsafe.Pointer(&result_)))
}

func (n *NSBox) BorderColor() Color {
	result_ := C.C_NSBox_BorderColor(n.Ptr())
	return MakeColor(result_)
}

func (n *NSBox) SetBorderColor(value Color) {
	C.C_NSBox_SetBorderColor(n.Ptr(), objc.ExtractPtr(value))
}

func (n *NSBox) BorderWidth() coregraphics.Float {
	result_ := C.C_NSBox_BorderWidth(n.Ptr())
	return coregraphics.Float(float64(result_))
}

func (n *NSBox) SetBorderWidth(value coregraphics.Float) {
	C.C_NSBox_SetBorderWidth(n.Ptr(), C.double(float64(value)))
}

func (n *NSBox) CornerRadius() coregraphics.Float {
	result_ := C.C_NSBox_CornerRadius(n.Ptr())
	return coregraphics.Float(float64(result_))
}

func (n *NSBox) SetCornerRadius(value coregraphics.Float) {
	C.C_NSBox_SetCornerRadius(n.Ptr(), C.double(float64(value)))
}

func (n *NSBox) FillColor() Color {
	result_ := C.C_NSBox_FillColor(n.Ptr())
	return MakeColor(result_)
}

func (n *NSBox) SetFillColor(value Color) {
	C.C_NSBox_SetFillColor(n.Ptr(), objc.ExtractPtr(value))
}

func (n *NSBox) ContentView() View {
	result_ := C.C_NSBox_ContentView(n.Ptr())
	return MakeView(result_)
}

func (n *NSBox) SetContentView(value View) {
	C.C_NSBox_SetContentView(n.Ptr(), objc.ExtractPtr(value))
}

func (n *NSBox) ContentViewMargins() foundation.Size {
	result_ := C.C_NSBox_ContentViewMargins(n.Ptr())
	return foundation.Size(coregraphics.FromCGSizePointer(unsafe.Pointer(&result_)))
}

func (n *NSBox) SetContentViewMargins(value foundation.Size) {
	C.C_NSBox_SetContentViewMargins(n.Ptr(), *(*C.CGSize)(coregraphics.ToCGSizePointer(coregraphics.Size(value))))
}
