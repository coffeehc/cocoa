package appkit

// #cgo CFLAGS: -x objective-c
// #cgo LDFLAGS: -framework AppKit
// #include "scroll_view.h"
import "C"
import (
	"github.com/hsiafan/cocoa/foundation"
	"unsafe"
)

type ScrollView interface {
	View
	HasVerticalScroller() bool
	SetHasVerticalScroller(hasVerticalScroller bool)
	HasHorizontalScroller() bool
	SetHasHorizontalScroller(hasHorizontalScroller bool)
	DocumentView() View
	SetDocumentView(documentView View)
	BorderType() BorderType
	SetBorderType(borderType BorderType)
	ContentSize() foundation.Size
}

var _ ScrollView = (*NSScrollView)(nil)

type NSScrollView struct {
	NSView
}

func MakeScrollView(ptr unsafe.Pointer) *NSScrollView {
	if ptr == nil {
		return nil
	}
	return &NSScrollView{
		NSView: *MakeView(ptr),
	}
}

func (s *NSScrollView) HasVerticalScroller() bool {
	return bool(C.ScrollView_HasVerticalScroller(s.Ptr()))
}

func (s *NSScrollView) SetHasVerticalScroller(hasVerticalScroller bool) {
	C.ScrollView_SetHasVerticalScroller(s.Ptr(), C.bool(hasVerticalScroller))
}

func (s *NSScrollView) HasHorizontalScroller() bool {
	return bool(C.ScrollView_HasHorizontalScroller(s.Ptr()))
}

func (s *NSScrollView) SetHasHorizontalScroller(hasHorizontalScroller bool) {
	C.ScrollView_SetHasHorizontalScroller(s.Ptr(), C.bool(hasHorizontalScroller))
}

func (s *NSScrollView) DocumentView() View {
	return MakeView(C.ScrollView_DocumentView(s.Ptr()))
}

func (s *NSScrollView) SetDocumentView(documentView View) {
	C.ScrollView_SetDocumentView(s.Ptr(), toPointer(documentView))
}

func (s *NSScrollView) BorderType() BorderType {
	return BorderType(C.ScrollView_BorderType(s.Ptr()))
}

func (s *NSScrollView) SetBorderType(borderType BorderType) {
	C.ScrollView_SetBorderType(s.Ptr(), C.ulong(borderType))
}

func (s *NSScrollView) ContentSize() foundation.Size {
	return toSize(C.ScrollView_ContentSize(s.Ptr()))
}

func NewScrollView(frame foundation.Rect) ScrollView {
	return MakeScrollView(C.ScrollView_NewScrollView(toNSRect(frame)))
}
