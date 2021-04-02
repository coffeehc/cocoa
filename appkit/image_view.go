package appkit

// #cgo CFLAGS: -x objective-c
// #cgo LDFLAGS: -framework AppKit
// #include "image_view.h"
import "C"
import (
	"unsafe"
)

type ImageView interface {
	Control
	Image() Image
	SetImage(image Image)
	ImageFrameStyle() ImageFrameStyle
	SetImageFrameStyle(imageFrameStyle ImageFrameStyle)
	ImageAlignment() ImageAlignment
	SetImageAlignment(imageAlignment ImageAlignment)
	ImageScaling() ImageScaling
	SetImageScaling(imageScaling ImageScaling)
	Animates() bool
	SetAnimates(animates bool)
	IsEditable() bool
	SetEditable(editable bool)
	AllowsCutCopyPaste() bool
	SetAllowsCutCopyPaste(allowsCutCopyPaste bool)
	ContentTintColor() Color
	SetContentTintColor(contentTintColor Color)
	SymbolConfiguration() ImageSymbolConfiguration
	SetSymbolConfiguration(symbolConfiguration ImageSymbolConfiguration)
}

var _ ImageView = (*NSImageView)(nil)

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

func (i *NSImageView) Image() Image {
	return MakeImage(C.ImageView_Image(i.Ptr()))
}

func (i *NSImageView) SetImage(image Image) {
	C.ImageView_SetImage(i.Ptr(), toPointer(image))
}

func (i *NSImageView) ImageFrameStyle() ImageFrameStyle {
	return ImageFrameStyle(C.ImageView_ImageFrameStyle(i.Ptr()))
}

func (i *NSImageView) SetImageFrameStyle(imageFrameStyle ImageFrameStyle) {
	C.ImageView_SetImageFrameStyle(i.Ptr(), C.ulong(imageFrameStyle))
}

func (i *NSImageView) ImageAlignment() ImageAlignment {
	return ImageAlignment(C.ImageView_ImageAlignment(i.Ptr()))
}

func (i *NSImageView) SetImageAlignment(imageAlignment ImageAlignment) {
	C.ImageView_SetImageAlignment(i.Ptr(), C.ulong(imageAlignment))
}

func (i *NSImageView) ImageScaling() ImageScaling {
	return ImageScaling(C.ImageView_ImageScaling(i.Ptr()))
}

func (i *NSImageView) SetImageScaling(imageScaling ImageScaling) {
	C.ImageView_SetImageScaling(i.Ptr(), C.ulong(imageScaling))
}

func (i *NSImageView) Animates() bool {
	return bool(C.ImageView_Animates(i.Ptr()))
}

func (i *NSImageView) SetAnimates(animates bool) {
	C.ImageView_SetAnimates(i.Ptr(), C.bool(animates))
}

func (i *NSImageView) IsEditable() bool {
	return bool(C.ImageView_IsEditable(i.Ptr()))
}

func (i *NSImageView) SetEditable(editable bool) {
	C.ImageView_SetEditable(i.Ptr(), C.bool(editable))
}

func (i *NSImageView) AllowsCutCopyPaste() bool {
	return bool(C.ImageView_AllowsCutCopyPaste(i.Ptr()))
}

func (i *NSImageView) SetAllowsCutCopyPaste(allowsCutCopyPaste bool) {
	C.ImageView_SetAllowsCutCopyPaste(i.Ptr(), C.bool(allowsCutCopyPaste))
}

func (i *NSImageView) ContentTintColor() Color {
	return MakeColor(C.ImageView_ContentTintColor(i.Ptr()))
}

func (i *NSImageView) SetContentTintColor(contentTintColor Color) {
	C.ImageView_SetContentTintColor(i.Ptr(), toPointer(contentTintColor))
}

func (i *NSImageView) SymbolConfiguration() ImageSymbolConfiguration {
	return MakeImageSymbolConfiguration(C.ImageView_SymbolConfiguration(i.Ptr()))
}

func (i *NSImageView) SetSymbolConfiguration(symbolConfiguration ImageSymbolConfiguration) {
	C.ImageView_SetSymbolConfiguration(i.Ptr(), toPointer(symbolConfiguration))
}

func ImageViewWithImage(image Image) {
	C.ImageView_ImageViewWithImage(toPointer(image))
}
