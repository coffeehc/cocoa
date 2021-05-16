package appkit

// #cgo CFLAGS: -x objective-c
// #cgo LDFLAGS: -framework AppKit
// #include "image_view.h"
import "C"
import (
	"github.com/hsiafan/cocoa/coregraphics"
	"github.com/hsiafan/cocoa/foundation"
	"github.com/hsiafan/cocoa/objc"
	"unsafe"
)

type ImageView interface {
	Control
	Image() Image
	SetImage(value Image)
	ImageFrameStyle() ImageFrameStyle
	SetImageFrameStyle(value ImageFrameStyle)
	ImageAlignment() ImageAlignment
	SetImageAlignment(value ImageAlignment)
	ImageScaling() ImageScaling
	SetImageScaling(value ImageScaling)
	Animates() bool
	SetAnimates(value bool)
	IsEditable() bool
	SetEditable(value bool)
	AllowsCutCopyPaste() bool
	SetAllowsCutCopyPaste(value bool)
	ContentTintColor() Color
	SetContentTintColor(value Color)
	SymbolConfiguration() ImageSymbolConfiguration
	SetSymbolConfiguration(value ImageSymbolConfiguration)
}

type NSImageView struct {
	NSControl
}

func MakeImageView(ptr unsafe.Pointer) *NSImageView {
	if ptr == nil {
		return nil
	}
	return &NSImageView{
		NSControl: *MakeControl(ptr),
	}
}

func AllocImageView() *NSImageView {
	return MakeImageView(C.C_ImageView_Alloc())
}

func (n *NSImageView) InitWithFrame(frameRect foundation.Rect) ImageView {
	result := C.C_NSImageView_InitWithFrame(n.Ptr(), *(*C.CGRect)(coregraphics.ToCGRectPointer(coregraphics.Rect(frameRect))))
	return MakeImageView(result)
}

func (n *NSImageView) InitWithCoder(coder foundation.Coder) ImageView {
	result := C.C_NSImageView_InitWithCoder(n.Ptr(), objc.ExtractPtr(coder))
	return MakeImageView(result)
}

func (n *NSImageView) Init() ImageView {
	result := C.C_NSImageView_Init(n.Ptr())
	return MakeImageView(result)
}

func ImageViewWithImage(image Image) ImageView {
	result := C.C_NSImageView_ImageViewWithImage(objc.ExtractPtr(image))
	return MakeImageView(result)
}

func (n *NSImageView) Image() Image {
	result := C.C_NSImageView_Image(n.Ptr())
	return MakeImage(result)
}

func (n *NSImageView) SetImage(value Image) {
	C.C_NSImageView_SetImage(n.Ptr(), objc.ExtractPtr(value))
}

func (n *NSImageView) ImageFrameStyle() ImageFrameStyle {
	result := C.C_NSImageView_ImageFrameStyle(n.Ptr())
	return ImageFrameStyle(uint(result))
}

func (n *NSImageView) SetImageFrameStyle(value ImageFrameStyle) {
	C.C_NSImageView_SetImageFrameStyle(n.Ptr(), C.uint(uint(value)))
}

func (n *NSImageView) ImageAlignment() ImageAlignment {
	result := C.C_NSImageView_ImageAlignment(n.Ptr())
	return ImageAlignment(uint(result))
}

func (n *NSImageView) SetImageAlignment(value ImageAlignment) {
	C.C_NSImageView_SetImageAlignment(n.Ptr(), C.uint(uint(value)))
}

func (n *NSImageView) ImageScaling() ImageScaling {
	result := C.C_NSImageView_ImageScaling(n.Ptr())
	return ImageScaling(uint(result))
}

func (n *NSImageView) SetImageScaling(value ImageScaling) {
	C.C_NSImageView_SetImageScaling(n.Ptr(), C.uint(uint(value)))
}

func (n *NSImageView) Animates() bool {
	result := C.C_NSImageView_Animates(n.Ptr())
	return bool(result)
}

func (n *NSImageView) SetAnimates(value bool) {
	C.C_NSImageView_SetAnimates(n.Ptr(), C.bool(value))
}

func (n *NSImageView) IsEditable() bool {
	result := C.C_NSImageView_IsEditable(n.Ptr())
	return bool(result)
}

func (n *NSImageView) SetEditable(value bool) {
	C.C_NSImageView_SetEditable(n.Ptr(), C.bool(value))
}

func (n *NSImageView) AllowsCutCopyPaste() bool {
	result := C.C_NSImageView_AllowsCutCopyPaste(n.Ptr())
	return bool(result)
}

func (n *NSImageView) SetAllowsCutCopyPaste(value bool) {
	C.C_NSImageView_SetAllowsCutCopyPaste(n.Ptr(), C.bool(value))
}

func (n *NSImageView) ContentTintColor() Color {
	result := C.C_NSImageView_ContentTintColor(n.Ptr())
	return MakeColor(result)
}

func (n *NSImageView) SetContentTintColor(value Color) {
	C.C_NSImageView_SetContentTintColor(n.Ptr(), objc.ExtractPtr(value))
}

func (n *NSImageView) SymbolConfiguration() ImageSymbolConfiguration {
	result := C.C_NSImageView_SymbolConfiguration(n.Ptr())
	return MakeImageSymbolConfiguration(result)
}

func (n *NSImageView) SetSymbolConfiguration(value ImageSymbolConfiguration) {
	C.C_NSImageView_SetSymbolConfiguration(n.Ptr(), objc.ExtractPtr(value))
}
