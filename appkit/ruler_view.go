package appkit

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
	Markers() []RulerMarker
	SetMarkers(value []RulerMarker)
	ScrollView() ScrollView
	SetScrollView(value ScrollView)
	Orientation() RulerOrientation
	SetOrientation(value RulerOrientation)
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

func MakeRulerView(ptr unsafe.Pointer) NSRulerView {
	return NSRulerView{
		NSView: MakeView(ptr),
	}
}

func (n NSRulerView) InitWithScrollView_Orientation(scrollView ScrollView, orientation RulerOrientation) NSRulerView {
	result_ := C.C_NSRulerView_InitWithScrollView_Orientation(n.Ptr(), objc.ExtractPtr(scrollView), C.uint(uint(orientation)))
	return MakeRulerView(result_)
}

func (n NSRulerView) InitWithCoder(coder foundation.Coder) NSRulerView {
	result_ := C.C_NSRulerView_InitWithCoder(n.Ptr(), objc.ExtractPtr(coder))
	return MakeRulerView(result_)
}

func (n NSRulerView) InitWithFrame(frameRect foundation.Rect) NSRulerView {
	result_ := C.C_NSRulerView_InitWithFrame(n.Ptr(), *(*C.CGRect)(coregraphics.ToCGRectPointer(coregraphics.Rect(frameRect))))
	return MakeRulerView(result_)
}

func (n NSRulerView) Init() NSRulerView {
	result_ := C.C_NSRulerView_Init(n.Ptr())
	return MakeRulerView(result_)
}

func AllocRulerView() NSRulerView {
	result_ := C.C_NSRulerView_AllocRulerView()
	return MakeRulerView(result_)
}

func NewRulerView() NSRulerView {
	result_ := C.C_NSRulerView_NewRulerView()
	return MakeRulerView(result_)
}

func (n NSRulerView) Autorelease() NSRulerView {
	result_ := C.C_NSRulerView_Autorelease(n.Ptr())
	return MakeRulerView(result_)
}

func (n NSRulerView) Retain() NSRulerView {
	result_ := C.C_NSRulerView_Retain(n.Ptr())
	return MakeRulerView(result_)
}

func RulerView_RegisterUnitWithName_Abbreviation_UnitToPointsConversionFactor_StepUpCycle_StepDownCycle(unitName RulerViewUnitName, abbreviation string, conversionFactor coregraphics.Float, stepUpCycle []foundation.Number, stepDownCycle []foundation.Number) {
	var cStepUpCycle C.Array
	if len(stepUpCycle) > 0 {
		cStepUpCycleData := make([]unsafe.Pointer, len(stepUpCycle))
		for idx, v := range stepUpCycle {
			cStepUpCycleData[idx] = objc.ExtractPtr(v)
		}
		cStepUpCycle.data = unsafe.Pointer(&cStepUpCycleData[0])
		cStepUpCycle.len = C.int(len(stepUpCycle))
	}
	var cStepDownCycle C.Array
	if len(stepDownCycle) > 0 {
		cStepDownCycleData := make([]unsafe.Pointer, len(stepDownCycle))
		for idx, v := range stepDownCycle {
			cStepDownCycleData[idx] = objc.ExtractPtr(v)
		}
		cStepDownCycle.data = unsafe.Pointer(&cStepDownCycleData[0])
		cStepDownCycle.len = C.int(len(stepDownCycle))
	}
	C.C_NSRulerView_RulerView_RegisterUnitWithName_Abbreviation_UnitToPointsConversionFactor_StepUpCycle_StepDownCycle(foundation.NewString(string(unitName)).Ptr(), foundation.NewString(abbreviation).Ptr(), C.double(float64(conversionFactor)), cStepUpCycle, cStepDownCycle)
}

func (n NSRulerView) AddMarker(marker RulerMarker) {
	C.C_NSRulerView_AddMarker(n.Ptr(), objc.ExtractPtr(marker))
}

func (n NSRulerView) RemoveMarker(marker RulerMarker) {
	C.C_NSRulerView_RemoveMarker(n.Ptr(), objc.ExtractPtr(marker))
}

func (n NSRulerView) TrackMarker_WithMouseEvent(marker RulerMarker, event Event) bool {
	result_ := C.C_NSRulerView_TrackMarker_WithMouseEvent(n.Ptr(), objc.ExtractPtr(marker), objc.ExtractPtr(event))
	return bool(result_)
}

func (n NSRulerView) MoveRulerlineFromLocation_ToLocation(oldLocation coregraphics.Float, newLocation coregraphics.Float) {
	C.C_NSRulerView_MoveRulerlineFromLocation_ToLocation(n.Ptr(), C.double(float64(oldLocation)), C.double(float64(newLocation)))
}

func (n NSRulerView) DrawHashMarksAndLabelsInRect(rect foundation.Rect) {
	C.C_NSRulerView_DrawHashMarksAndLabelsInRect(n.Ptr(), *(*C.CGRect)(coregraphics.ToCGRectPointer(coregraphics.Rect(rect))))
}

func (n NSRulerView) DrawMarkersInRect(rect foundation.Rect) {
	C.C_NSRulerView_DrawMarkersInRect(n.Ptr(), *(*C.CGRect)(coregraphics.ToCGRectPointer(coregraphics.Rect(rect))))
}

func (n NSRulerView) InvalidateHashMarks() {
	C.C_NSRulerView_InvalidateHashMarks(n.Ptr())
}

func (n NSRulerView) MeasurementUnits() RulerViewUnitName {
	result_ := C.C_NSRulerView_MeasurementUnits(n.Ptr())
	return RulerViewUnitName(foundation.MakeString(result_).String())
}

func (n NSRulerView) SetMeasurementUnits(value RulerViewUnitName) {
	C.C_NSRulerView_SetMeasurementUnits(n.Ptr(), foundation.NewString(string(value)).Ptr())
}

func (n NSRulerView) ClientView() View {
	result_ := C.C_NSRulerView_ClientView(n.Ptr())
	return MakeView(result_)
}

func (n NSRulerView) SetClientView(value View) {
	C.C_NSRulerView_SetClientView(n.Ptr(), objc.ExtractPtr(value))
}

func (n NSRulerView) AccessoryView() View {
	result_ := C.C_NSRulerView_AccessoryView(n.Ptr())
	return MakeView(result_)
}

func (n NSRulerView) SetAccessoryView(value View) {
	C.C_NSRulerView_SetAccessoryView(n.Ptr(), objc.ExtractPtr(value))
}

func (n NSRulerView) OriginOffset() coregraphics.Float {
	result_ := C.C_NSRulerView_OriginOffset(n.Ptr())
	return coregraphics.Float(float64(result_))
}

func (n NSRulerView) SetOriginOffset(value coregraphics.Float) {
	C.C_NSRulerView_SetOriginOffset(n.Ptr(), C.double(float64(value)))
}

func (n NSRulerView) Markers() []RulerMarker {
	result_ := C.C_NSRulerView_Markers(n.Ptr())
	if result_.len > 0 {
		defer C.free(result_.data)
	}
	result_Slice := unsafe.Slice((*unsafe.Pointer)(result_.data), int(result_.len))
	var goResult_ = make([]RulerMarker, len(result_Slice))
	for idx, r := range result_Slice {
		goResult_[idx] = MakeRulerMarker(r)
	}
	return goResult_
}

func (n NSRulerView) SetMarkers(value []RulerMarker) {
	var cValue C.Array
	if len(value) > 0 {
		cValueData := make([]unsafe.Pointer, len(value))
		for idx, v := range value {
			cValueData[idx] = objc.ExtractPtr(v)
		}
		cValue.data = unsafe.Pointer(&cValueData[0])
		cValue.len = C.int(len(value))
	}
	C.C_NSRulerView_SetMarkers(n.Ptr(), cValue)
}

func (n NSRulerView) ScrollView() ScrollView {
	result_ := C.C_NSRulerView_ScrollView(n.Ptr())
	return MakeScrollView(result_)
}

func (n NSRulerView) SetScrollView(value ScrollView) {
	C.C_NSRulerView_SetScrollView(n.Ptr(), objc.ExtractPtr(value))
}

func (n NSRulerView) Orientation() RulerOrientation {
	result_ := C.C_NSRulerView_Orientation(n.Ptr())
	return RulerOrientation(uint(result_))
}

func (n NSRulerView) SetOrientation(value RulerOrientation) {
	C.C_NSRulerView_SetOrientation(n.Ptr(), C.uint(uint(value)))
}

func (n NSRulerView) ReservedThicknessForAccessoryView() coregraphics.Float {
	result_ := C.C_NSRulerView_ReservedThicknessForAccessoryView(n.Ptr())
	return coregraphics.Float(float64(result_))
}

func (n NSRulerView) SetReservedThicknessForAccessoryView(value coregraphics.Float) {
	C.C_NSRulerView_SetReservedThicknessForAccessoryView(n.Ptr(), C.double(float64(value)))
}

func (n NSRulerView) ReservedThicknessForMarkers() coregraphics.Float {
	result_ := C.C_NSRulerView_ReservedThicknessForMarkers(n.Ptr())
	return coregraphics.Float(float64(result_))
}

func (n NSRulerView) SetReservedThicknessForMarkers(value coregraphics.Float) {
	C.C_NSRulerView_SetReservedThicknessForMarkers(n.Ptr(), C.double(float64(value)))
}

func (n NSRulerView) RuleThickness() coregraphics.Float {
	result_ := C.C_NSRulerView_RuleThickness(n.Ptr())
	return coregraphics.Float(float64(result_))
}

func (n NSRulerView) SetRuleThickness(value coregraphics.Float) {
	C.C_NSRulerView_SetRuleThickness(n.Ptr(), C.double(float64(value)))
}

func (n NSRulerView) RequiredThickness() coregraphics.Float {
	result_ := C.C_NSRulerView_RequiredThickness(n.Ptr())
	return coregraphics.Float(float64(result_))
}

func (n NSRulerView) BaselineLocation() coregraphics.Float {
	result_ := C.C_NSRulerView_BaselineLocation(n.Ptr())
	return coregraphics.Float(float64(result_))
}
