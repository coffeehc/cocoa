package appkit

// #cgo CFLAGS: -x objective-c
// #cgo LDFLAGS: -framework AppKit
// #include "text_container.h"
import "C"
import (
	"github.com/hsiafan/cocoa/foundation"
	"github.com/hsiafan/cocoa/objc"
	"unsafe"
)

type TextContainer interface {
	objc.Object
	Size() foundation.Size
	SetSize(size foundation.Size)
	WidthTracksTextView() bool
	SetWidthTracksTextView(widthTracksTextView bool)
	HeightTracksTextView() bool
	SetHeightTracksTextView(heightTracksTextView bool)
}

var _ TextContainer = (*NSTextContainer)(nil)

type NSTextContainer struct {
	objc.NSObject
}

func MakeTextContainer(ptr unsafe.Pointer) *NSTextContainer {
	if ptr == nil {
		return nil
	}
	return &NSTextContainer{
		NSObject: *objc.MakeObject(ptr),
	}
}

func (t *NSTextContainer) Size() foundation.Size {
	return toSize(C.TextContainer_Size(t.Ptr()))
}

func (t *NSTextContainer) SetSize(size foundation.Size) {
	C.TextContainer_SetSize(t.Ptr(), toNSSize(size))
}

func (t *NSTextContainer) WidthTracksTextView() bool {
	return bool(C.TextContainer_WidthTracksTextView(t.Ptr()))
}

func (t *NSTextContainer) SetWidthTracksTextView(widthTracksTextView bool) {
	C.TextContainer_SetWidthTracksTextView(t.Ptr(), C.bool(widthTracksTextView))
}

func (t *NSTextContainer) HeightTracksTextView() bool {
	return bool(C.TextContainer_HeightTracksTextView(t.Ptr()))
}

func (t *NSTextContainer) SetHeightTracksTextView(heightTracksTextView bool) {
	C.TextContainer_SetHeightTracksTextView(t.Ptr(), C.bool(heightTracksTextView))
}

func NewTextContainer(size foundation.Size) TextContainer {
	return MakeTextContainer(C.TextContainer_NewTextContainer(toNSSize(size)))
}
