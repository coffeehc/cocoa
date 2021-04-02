package appkit

// #cgo CFLAGS: -x objective-c
// #cgo LDFLAGS: -framework AppKit
// #include "split_view.h"
import "C"
import (
	"github.com/hsiafan/cocoa/foundation"
	"unsafe"
)

type SplitView interface {
	View
	ArrangesAllSubviews() bool
	SetArrangesAllSubviews(arrangesAllSubviews bool)
	IsVertical() bool
	SetVertical(vertical bool)
	DividerColor() Color
	DividerThickness() float64
	DividerStyle() SplitViewDividerStyle
	SetDividerStyle(dividerStyle SplitViewDividerStyle)
	MinPossiblePositionOfDividerAtIndex(dividerIndex int) float64
	MaxPossiblePositionOfDividerAtIndex(dividerIndex int) float64
	SetPosition(position float64, dividerIndex int)
}

var _ SplitView = (*NSSplitView)(nil)

type NSSplitView struct {
	NSView
}

func MakeSplitView(ptr unsafe.Pointer) *NSSplitView {
	if ptr == nil {
		return nil
	}
	return &NSSplitView{
		NSView: *MakeView(ptr),
	}
}

func (s *NSSplitView) ArrangesAllSubviews() bool {
	return bool(C.SplitView_ArrangesAllSubviews(s.Ptr()))
}

func (s *NSSplitView) SetArrangesAllSubviews(arrangesAllSubviews bool) {
	C.SplitView_SetArrangesAllSubviews(s.Ptr(), C.bool(arrangesAllSubviews))
}

func (s *NSSplitView) IsVertical() bool {
	return bool(C.SplitView_IsVertical(s.Ptr()))
}

func (s *NSSplitView) SetVertical(vertical bool) {
	C.SplitView_SetVertical(s.Ptr(), C.bool(vertical))
}

func (s *NSSplitView) DividerColor() Color {
	return MakeColor(C.SplitView_DividerColor(s.Ptr()))
}

func (s *NSSplitView) DividerThickness() float64 {
	return float64(C.SplitView_DividerThickness(s.Ptr()))
}

func (s *NSSplitView) DividerStyle() SplitViewDividerStyle {
	return SplitViewDividerStyle(C.SplitView_DividerStyle(s.Ptr()))
}

func (s *NSSplitView) SetDividerStyle(dividerStyle SplitViewDividerStyle) {
	C.SplitView_SetDividerStyle(s.Ptr(), C.long(dividerStyle))
}

func NewSplitView(frame foundation.Rect) SplitView {
	return MakeSplitView(C.SplitView_NewSplitView(toNSRect(frame)))
}

func (s *NSSplitView) MinPossiblePositionOfDividerAtIndex(dividerIndex int) float64 {
	return float64(C.SplitView_MinPossiblePositionOfDividerAtIndex(s.Ptr(), C.long(dividerIndex)))
}

func (s *NSSplitView) MaxPossiblePositionOfDividerAtIndex(dividerIndex int) float64 {
	return float64(C.SplitView_MaxPossiblePositionOfDividerAtIndex(s.Ptr(), C.long(dividerIndex)))
}

func (s *NSSplitView) SetPosition(position float64, dividerIndex int) {
	C.SplitView_SetPosition(s.Ptr(), C.double(position), C.long(dividerIndex))
}
