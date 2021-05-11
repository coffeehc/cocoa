package appkit

// #cgo CFLAGS: -x objective-c
// #cgo LDFLAGS: -framework AppKit
// #include "ruler_view.h"
import "C"
import (
	"github.com/hsiafan/cocoa/coregraphics"
	"github.com/hsiafan/cocoa/foundation"
	"github.com/hsiafan/cocoa/objc"
	"unsafe"
)

type RulerView interface {
	View
	AddMarker(marker RulerMarker)
	RemoveMarker(marker RulerMarker)
	TrackMarker_WithMouseEvent(marker RulerMarker, event Event) bool
	MoveRulerlineFromLocation_ToLocation(oldLocation coregraphics.Float, newLocation coregraphics.Float)
	DrawHashMarksAndLabelsInRect(rect foundation.Rect)
	DrawMarkersInRect(rect foundation.Rect)
	InvalidateHashMarks()
	MeasurementUnits() RulerViewUnitName
	SetMeasurementUnits(value RulerViewUnitName)
	ClientView() View
	SetClientView(value View)
	AccessoryView() View
	SetAccessoryView(value View)
	OriginOffset() coregraphics.Float
	SetOriginOffset(value coregraphics.Float)
	ScrollView() ScrollView
	SetScrollView(value ScrollView)
	ReservedThicknessForAccessoryView() coregraphics.Float
	SetReservedThicknessForAccessoryView(value coregraphics.Float)
	ReservedThicknessForMarkers() coregraphics.Float
	SetReservedThicknessForMarkers(value coregraphics.Float)
	RuleThickness() coregraphics.Float
	SetRuleThickness(value coregraphics.Float)
	RequiredThickness() coregraphics.Float
	BaselineLocation() coregraphics.Float
}

type NSRulerView struct {
	NSView
}

func MakeRulerView(ptr unsafe.Pointer) *NSRulerView {
	if ptr == nil {
		return nil
	}
	return &NSRulerView{
		NSView: *MakeView(ptr),
	}
}

func AllocRulerView() *NSRulerView {
	return MakeRulerView(C.C_RulerView_Alloc())
}

func (n *NSRulerView) InitWithCoder(coder foundation.Coder) RulerView {
	result := C.C_NSRulerView_InitWithCoder(n.Ptr(), objc.ExtractPtr(coder))
	return MakeRulerView(result)
}

func (n *NSRulerView) InitWithFrame(frameRect foundation.Rect) RulerView {
	result := C.C_NSRulerView_InitWithFrame(n.Ptr(), *(*C.CGRect)(coregraphics.ToCGRectPointer(coregraphics.Rect(frameRect))))
	return MakeRulerView(result)
}

func (n *NSRulerView) Init() RulerView {
	result := C.C_NSRulerView_Init(n.Ptr())
	return MakeRulerView(result)
}

func (n *NSRulerView) AddMarker(marker RulerMarker) {
	C.C_NSRulerView_AddMarker(n.Ptr(), objc.ExtractPtr(marker))
}

func (n *NSRulerView) RemoveMarker(marker RulerMarker) {
	C.C_NSRulerView_RemoveMarker(n.Ptr(), objc.ExtractPtr(marker))
}

func (n *NSRulerView) TrackMarker_WithMouseEvent(marker RulerMarker, event Event) bool {
	result := C.C_NSRulerView_TrackMarker_WithMouseEvent(n.Ptr(), objc.ExtractPtr(marker), objc.ExtractPtr(event))
	return bool(result)
}

func (n *NSRulerView) MoveRulerlineFromLocation_ToLocation(oldLocation coregraphics.Float, newLocation coregraphics.Float) {
	C.C_NSRulerView_MoveRulerlineFromLocation_ToLocation(n.Ptr(), C.double(float64(oldLocation)), C.double(float64(newLocation)))
}

func (n *NSRulerView) DrawHashMarksAndLabelsInRect(rect foundation.Rect) {
	C.C_NSRulerView_DrawHashMarksAndLabelsInRect(n.Ptr(), *(*C.CGRect)(coregraphics.ToCGRectPointer(coregraphics.Rect(rect))))
}

func (n *NSRulerView) DrawMarkersInRect(rect foundation.Rect) {
	C.C_NSRulerView_DrawMarkersInRect(n.Ptr(), *(*C.CGRect)(coregraphics.ToCGRectPointer(coregraphics.Rect(rect))))
}

func (n *NSRulerView) InvalidateHashMarks() {
	C.C_NSRulerView_InvalidateHashMarks(n.Ptr())
}

func (n *NSRulerView) MeasurementUnits() RulerViewUnitName {
	result := C.C_NSRulerView_MeasurementUnits(n.Ptr())
	return RulerViewUnitName(foundation.MakeString(result).String())
}

func (n *NSRulerView) SetMeasurementUnits(value RulerViewUnitName) {
	C.C_NSRulerView_SetMeasurementUnits(n.Ptr(), foundation.NewString(string(value)).Ptr())
}

func (n *NSRulerView) ClientView() View {
	result := C.C_NSRulerView_ClientView(n.Ptr())
	return MakeView(result)
}

func (n *NSRulerView) SetClientView(value View) {
	C.C_NSRulerView_SetClientView(n.Ptr(), objc.ExtractPtr(value))
}

func (n *NSRulerView) AccessoryView() View {
	result := C.C_NSRulerView_AccessoryView(n.Ptr())
	return MakeView(result)
}

func (n *NSRulerView) SetAccessoryView(value View) {
	C.C_NSRulerView_SetAccessoryView(n.Ptr(), objc.ExtractPtr(value))
}

func (n *NSRulerView) OriginOffset() coregraphics.Float {
	result := C.C_NSRulerView_OriginOffset(n.Ptr())
	return coregraphics.Float(float64(result))
}

func (n *NSRulerView) SetOriginOffset(value coregraphics.Float) {
	C.C_NSRulerView_SetOriginOffset(n.Ptr(), C.double(float64(value)))
}

func (n *NSRulerView) ScrollView() ScrollView {
	result := C.C_NSRulerView_ScrollView(n.Ptr())
	return MakeScrollView(result)
}

func (n *NSRulerView) SetScrollView(value ScrollView) {
	C.C_NSRulerView_SetScrollView(n.Ptr(), objc.ExtractPtr(value))
}

func (n *NSRulerView) ReservedThicknessForAccessoryView() coregraphics.Float {
	result := C.C_NSRulerView_ReservedThicknessForAccessoryView(n.Ptr())
	return coregraphics.Float(float64(result))
}

func (n *NSRulerView) SetReservedThicknessForAccessoryView(value coregraphics.Float) {
	C.C_NSRulerView_SetReservedThicknessForAccessoryView(n.Ptr(), C.double(float64(value)))
}

func (n *NSRulerView) ReservedThicknessForMarkers() coregraphics.Float {
	result := C.C_NSRulerView_ReservedThicknessForMarkers(n.Ptr())
	return coregraphics.Float(float64(result))
}

func (n *NSRulerView) SetReservedThicknessForMarkers(value coregraphics.Float) {
	C.C_NSRulerView_SetReservedThicknessForMarkers(n.Ptr(), C.double(float64(value)))
}

func (n *NSRulerView) RuleThickness() coregraphics.Float {
	result := C.C_NSRulerView_RuleThickness(n.Ptr())
	return coregraphics.Float(float64(result))
}

func (n *NSRulerView) SetRuleThickness(value coregraphics.Float) {
	C.C_NSRulerView_SetRuleThickness(n.Ptr(), C.double(float64(value)))
}

func (n *NSRulerView) RequiredThickness() coregraphics.Float {
	result := C.C_NSRulerView_RequiredThickness(n.Ptr())
	return coregraphics.Float(float64(result))
}

func (n *NSRulerView) BaselineLocation() coregraphics.Float {
	result := C.C_NSRulerView_BaselineLocation(n.Ptr())
	return coregraphics.Float(float64(result))
}
