package appkit

// #cgo CFLAGS: -x objective-c
// #cgo LDFLAGS: -framework AppKit
// #include "bezier_path.h"
import "C"
import (
	"github.com/hsiafan/cocoa/coregraphics"
	"github.com/hsiafan/cocoa/foundation"
	"github.com/hsiafan/cocoa/objc"
	"unsafe"
)

type BezierPath interface {
	objc.Object
	MoveToPoint(point foundation.Point)
	LineToPoint(point foundation.Point)
	CurveToPoint_ControlPoint1_ControlPoint2(endPoint foundation.Point, controlPoint1 foundation.Point, controlPoint2 foundation.Point)
	ClosePath()
	RelativeMoveToPoint(point foundation.Point)
	RelativeLineToPoint(point foundation.Point)
	RelativeCurveToPoint_ControlPoint1_ControlPoint2(endPoint foundation.Point, controlPoint1 foundation.Point, controlPoint2 foundation.Point)
	AppendBezierPath(path BezierPath)
	AppendBezierPathWithOvalInRect(rect foundation.Rect)
	AppendBezierPathWithArcFromPoint_ToPoint_Radius(point1 foundation.Point, point2 foundation.Point, radius coregraphics.Float)
	AppendBezierPathWithArcWithCenter_Radius_StartAngle_EndAngle(center foundation.Point, radius coregraphics.Float, startAngle coregraphics.Float, endAngle coregraphics.Float)
	AppendBezierPathWithArcWithCenter_Radius_StartAngle_EndAngle_Clockwise(center foundation.Point, radius coregraphics.Float, startAngle coregraphics.Float, endAngle coregraphics.Float, clockwise bool)
	AppendBezierPathWithRect(rect foundation.Rect)
	AppendBezierPathWithRoundedRect_XRadius_YRadius(rect foundation.Rect, xRadius coregraphics.Float, yRadius coregraphics.Float)
	Stroke()
	Fill()
	AddClip()
	SetClip()
	ContainsPoint(point foundation.Point) bool
	TransformUsingAffineTransform(transform foundation.AffineTransform)
	ElementAtIndex(index int) BezierPathElement
	RemoveAllPoints()
	BezierPathByFlatteningPath() BezierPath
	BezierPathByReversingPath() BezierPath
	WindingRule() WindingRule
	SetWindingRule(value WindingRule)
	LineCapStyle() LineCapStyle
	SetLineCapStyle(value LineCapStyle)
	LineJoinStyle() LineJoinStyle
	SetLineJoinStyle(value LineJoinStyle)
	LineWidth() coregraphics.Float
	SetLineWidth(value coregraphics.Float)
	MiterLimit() coregraphics.Float
	SetMiterLimit(value coregraphics.Float)
	Flatness() coregraphics.Float
	SetFlatness(value coregraphics.Float)
	Bounds() foundation.Rect
	ControlPointBounds() foundation.Rect
	CurrentPoint() foundation.Point
	IsEmpty() bool
	ElementCount() int
}

type NSBezierPath struct {
	objc.NSObject
}

func MakeBezierPath(ptr unsafe.Pointer) *NSBezierPath {
	if ptr == nil {
		return nil
	}
	return &NSBezierPath{
		NSObject: *objc.MakeObject(ptr),
	}
}

func AllocBezierPath() *NSBezierPath {
	return MakeBezierPath(C.C_BezierPath_Alloc())
}

func (n *NSBezierPath) Init() BezierPath {
	result_ := C.C_NSBezierPath_Init(n.Ptr())
	return MakeBezierPath(result_)
}

func BezierPath_() BezierPath {
	result_ := C.C_NSBezierPath_BezierPath_()
	return MakeBezierPath(result_)
}

func BezierPathWithOvalInRect(rect foundation.Rect) BezierPath {
	result_ := C.C_NSBezierPath_BezierPathWithOvalInRect(*(*C.CGRect)(coregraphics.ToCGRectPointer(coregraphics.Rect(rect))))
	return MakeBezierPath(result_)
}

func BezierPathWithRect(rect foundation.Rect) BezierPath {
	result_ := C.C_NSBezierPath_BezierPathWithRect(*(*C.CGRect)(coregraphics.ToCGRectPointer(coregraphics.Rect(rect))))
	return MakeBezierPath(result_)
}

func BezierPathWithRoundedRect_XRadius_YRadius(rect foundation.Rect, xRadius coregraphics.Float, yRadius coregraphics.Float) BezierPath {
	result_ := C.C_NSBezierPath_BezierPathWithRoundedRect_XRadius_YRadius(*(*C.CGRect)(coregraphics.ToCGRectPointer(coregraphics.Rect(rect))), C.double(float64(xRadius)), C.double(float64(yRadius)))
	return MakeBezierPath(result_)
}

func (n *NSBezierPath) MoveToPoint(point foundation.Point) {
	C.C_NSBezierPath_MoveToPoint(n.Ptr(), *(*C.CGPoint)(coregraphics.ToCGPointPointer(coregraphics.Point(point))))
}

func (n *NSBezierPath) LineToPoint(point foundation.Point) {
	C.C_NSBezierPath_LineToPoint(n.Ptr(), *(*C.CGPoint)(coregraphics.ToCGPointPointer(coregraphics.Point(point))))
}

func (n *NSBezierPath) CurveToPoint_ControlPoint1_ControlPoint2(endPoint foundation.Point, controlPoint1 foundation.Point, controlPoint2 foundation.Point) {
	C.C_NSBezierPath_CurveToPoint_ControlPoint1_ControlPoint2(n.Ptr(), *(*C.CGPoint)(coregraphics.ToCGPointPointer(coregraphics.Point(endPoint))), *(*C.CGPoint)(coregraphics.ToCGPointPointer(coregraphics.Point(controlPoint1))), *(*C.CGPoint)(coregraphics.ToCGPointPointer(coregraphics.Point(controlPoint2))))
}

func (n *NSBezierPath) ClosePath() {
	C.C_NSBezierPath_ClosePath(n.Ptr())
}

func (n *NSBezierPath) RelativeMoveToPoint(point foundation.Point) {
	C.C_NSBezierPath_RelativeMoveToPoint(n.Ptr(), *(*C.CGPoint)(coregraphics.ToCGPointPointer(coregraphics.Point(point))))
}

func (n *NSBezierPath) RelativeLineToPoint(point foundation.Point) {
	C.C_NSBezierPath_RelativeLineToPoint(n.Ptr(), *(*C.CGPoint)(coregraphics.ToCGPointPointer(coregraphics.Point(point))))
}

func (n *NSBezierPath) RelativeCurveToPoint_ControlPoint1_ControlPoint2(endPoint foundation.Point, controlPoint1 foundation.Point, controlPoint2 foundation.Point) {
	C.C_NSBezierPath_RelativeCurveToPoint_ControlPoint1_ControlPoint2(n.Ptr(), *(*C.CGPoint)(coregraphics.ToCGPointPointer(coregraphics.Point(endPoint))), *(*C.CGPoint)(coregraphics.ToCGPointPointer(coregraphics.Point(controlPoint1))), *(*C.CGPoint)(coregraphics.ToCGPointPointer(coregraphics.Point(controlPoint2))))
}

func (n *NSBezierPath) AppendBezierPath(path BezierPath) {
	C.C_NSBezierPath_AppendBezierPath(n.Ptr(), objc.ExtractPtr(path))
}

func (n *NSBezierPath) AppendBezierPathWithOvalInRect(rect foundation.Rect) {
	C.C_NSBezierPath_AppendBezierPathWithOvalInRect(n.Ptr(), *(*C.CGRect)(coregraphics.ToCGRectPointer(coregraphics.Rect(rect))))
}

func (n *NSBezierPath) AppendBezierPathWithArcFromPoint_ToPoint_Radius(point1 foundation.Point, point2 foundation.Point, radius coregraphics.Float) {
	C.C_NSBezierPath_AppendBezierPathWithArcFromPoint_ToPoint_Radius(n.Ptr(), *(*C.CGPoint)(coregraphics.ToCGPointPointer(coregraphics.Point(point1))), *(*C.CGPoint)(coregraphics.ToCGPointPointer(coregraphics.Point(point2))), C.double(float64(radius)))
}

func (n *NSBezierPath) AppendBezierPathWithArcWithCenter_Radius_StartAngle_EndAngle(center foundation.Point, radius coregraphics.Float, startAngle coregraphics.Float, endAngle coregraphics.Float) {
	C.C_NSBezierPath_AppendBezierPathWithArcWithCenter_Radius_StartAngle_EndAngle(n.Ptr(), *(*C.CGPoint)(coregraphics.ToCGPointPointer(coregraphics.Point(center))), C.double(float64(radius)), C.double(float64(startAngle)), C.double(float64(endAngle)))
}

func (n *NSBezierPath) AppendBezierPathWithArcWithCenter_Radius_StartAngle_EndAngle_Clockwise(center foundation.Point, radius coregraphics.Float, startAngle coregraphics.Float, endAngle coregraphics.Float, clockwise bool) {
	C.C_NSBezierPath_AppendBezierPathWithArcWithCenter_Radius_StartAngle_EndAngle_Clockwise(n.Ptr(), *(*C.CGPoint)(coregraphics.ToCGPointPointer(coregraphics.Point(center))), C.double(float64(radius)), C.double(float64(startAngle)), C.double(float64(endAngle)), C.bool(clockwise))
}

func (n *NSBezierPath) AppendBezierPathWithRect(rect foundation.Rect) {
	C.C_NSBezierPath_AppendBezierPathWithRect(n.Ptr(), *(*C.CGRect)(coregraphics.ToCGRectPointer(coregraphics.Rect(rect))))
}

func (n *NSBezierPath) AppendBezierPathWithRoundedRect_XRadius_YRadius(rect foundation.Rect, xRadius coregraphics.Float, yRadius coregraphics.Float) {
	C.C_NSBezierPath_AppendBezierPathWithRoundedRect_XRadius_YRadius(n.Ptr(), *(*C.CGRect)(coregraphics.ToCGRectPointer(coregraphics.Rect(rect))), C.double(float64(xRadius)), C.double(float64(yRadius)))
}

func (n *NSBezierPath) Stroke() {
	C.C_NSBezierPath_Stroke(n.Ptr())
}

func (n *NSBezierPath) Fill() {
	C.C_NSBezierPath_Fill(n.Ptr())
}

func BezierPath_FillRect(rect foundation.Rect) {
	C.C_NSBezierPath_BezierPath_FillRect(*(*C.CGRect)(coregraphics.ToCGRectPointer(coregraphics.Rect(rect))))
}

func BezierPath_StrokeRect(rect foundation.Rect) {
	C.C_NSBezierPath_BezierPath_StrokeRect(*(*C.CGRect)(coregraphics.ToCGRectPointer(coregraphics.Rect(rect))))
}

func BezierPath_StrokeLineFromPoint_ToPoint(point1 foundation.Point, point2 foundation.Point) {
	C.C_NSBezierPath_BezierPath_StrokeLineFromPoint_ToPoint(*(*C.CGPoint)(coregraphics.ToCGPointPointer(coregraphics.Point(point1))), *(*C.CGPoint)(coregraphics.ToCGPointPointer(coregraphics.Point(point2))))
}

func (n *NSBezierPath) AddClip() {
	C.C_NSBezierPath_AddClip(n.Ptr())
}

func (n *NSBezierPath) SetClip() {
	C.C_NSBezierPath_SetClip(n.Ptr())
}

func BezierPath_ClipRect(rect foundation.Rect) {
	C.C_NSBezierPath_BezierPath_ClipRect(*(*C.CGRect)(coregraphics.ToCGRectPointer(coregraphics.Rect(rect))))
}

func (n *NSBezierPath) ContainsPoint(point foundation.Point) bool {
	result_ := C.C_NSBezierPath_ContainsPoint(n.Ptr(), *(*C.CGPoint)(coregraphics.ToCGPointPointer(coregraphics.Point(point))))
	return bool(result_)
}

func (n *NSBezierPath) TransformUsingAffineTransform(transform foundation.AffineTransform) {
	C.C_NSBezierPath_TransformUsingAffineTransform(n.Ptr(), objc.ExtractPtr(transform))
}

func (n *NSBezierPath) ElementAtIndex(index int) BezierPathElement {
	result_ := C.C_NSBezierPath_ElementAtIndex(n.Ptr(), C.int(index))
	return BezierPathElement(uint(result_))
}

func (n *NSBezierPath) RemoveAllPoints() {
	C.C_NSBezierPath_RemoveAllPoints(n.Ptr())
}

func (n *NSBezierPath) BezierPathByFlatteningPath() BezierPath {
	result_ := C.C_NSBezierPath_BezierPathByFlatteningPath(n.Ptr())
	return MakeBezierPath(result_)
}

func (n *NSBezierPath) BezierPathByReversingPath() BezierPath {
	result_ := C.C_NSBezierPath_BezierPathByReversingPath(n.Ptr())
	return MakeBezierPath(result_)
}

func (n *NSBezierPath) WindingRule() WindingRule {
	result_ := C.C_NSBezierPath_WindingRule(n.Ptr())
	return WindingRule(uint(result_))
}

func (n *NSBezierPath) SetWindingRule(value WindingRule) {
	C.C_NSBezierPath_SetWindingRule(n.Ptr(), C.uint(uint(value)))
}

func (n *NSBezierPath) LineCapStyle() LineCapStyle {
	result_ := C.C_NSBezierPath_LineCapStyle(n.Ptr())
	return LineCapStyle(uint(result_))
}

func (n *NSBezierPath) SetLineCapStyle(value LineCapStyle) {
	C.C_NSBezierPath_SetLineCapStyle(n.Ptr(), C.uint(uint(value)))
}

func (n *NSBezierPath) LineJoinStyle() LineJoinStyle {
	result_ := C.C_NSBezierPath_LineJoinStyle(n.Ptr())
	return LineJoinStyle(uint(result_))
}

func (n *NSBezierPath) SetLineJoinStyle(value LineJoinStyle) {
	C.C_NSBezierPath_SetLineJoinStyle(n.Ptr(), C.uint(uint(value)))
}

func (n *NSBezierPath) LineWidth() coregraphics.Float {
	result_ := C.C_NSBezierPath_LineWidth(n.Ptr())
	return coregraphics.Float(float64(result_))
}

func (n *NSBezierPath) SetLineWidth(value coregraphics.Float) {
	C.C_NSBezierPath_SetLineWidth(n.Ptr(), C.double(float64(value)))
}

func (n *NSBezierPath) MiterLimit() coregraphics.Float {
	result_ := C.C_NSBezierPath_MiterLimit(n.Ptr())
	return coregraphics.Float(float64(result_))
}

func (n *NSBezierPath) SetMiterLimit(value coregraphics.Float) {
	C.C_NSBezierPath_SetMiterLimit(n.Ptr(), C.double(float64(value)))
}

func (n *NSBezierPath) Flatness() coregraphics.Float {
	result_ := C.C_NSBezierPath_Flatness(n.Ptr())
	return coregraphics.Float(float64(result_))
}

func (n *NSBezierPath) SetFlatness(value coregraphics.Float) {
	C.C_NSBezierPath_SetFlatness(n.Ptr(), C.double(float64(value)))
}

func BezierPath_DefaultWindingRule() WindingRule {
	result_ := C.C_NSBezierPath_BezierPath_DefaultWindingRule()
	return WindingRule(uint(result_))
}

func BezierPath_SetDefaultWindingRule(value WindingRule) {
	C.C_NSBezierPath_BezierPath_SetDefaultWindingRule(C.uint(uint(value)))
}

func BezierPath_DefaultLineCapStyle() LineCapStyle {
	result_ := C.C_NSBezierPath_BezierPath_DefaultLineCapStyle()
	return LineCapStyle(uint(result_))
}

func BezierPath_SetDefaultLineCapStyle(value LineCapStyle) {
	C.C_NSBezierPath_BezierPath_SetDefaultLineCapStyle(C.uint(uint(value)))
}

func BezierPath_DefaultLineJoinStyle() LineJoinStyle {
	result_ := C.C_NSBezierPath_BezierPath_DefaultLineJoinStyle()
	return LineJoinStyle(uint(result_))
}

func BezierPath_SetDefaultLineJoinStyle(value LineJoinStyle) {
	C.C_NSBezierPath_BezierPath_SetDefaultLineJoinStyle(C.uint(uint(value)))
}

func BezierPath_DefaultLineWidth() coregraphics.Float {
	result_ := C.C_NSBezierPath_BezierPath_DefaultLineWidth()
	return coregraphics.Float(float64(result_))
}

func BezierPath_SetDefaultLineWidth(value coregraphics.Float) {
	C.C_NSBezierPath_BezierPath_SetDefaultLineWidth(C.double(float64(value)))
}

func BezierPath_DefaultMiterLimit() coregraphics.Float {
	result_ := C.C_NSBezierPath_BezierPath_DefaultMiterLimit()
	return coregraphics.Float(float64(result_))
}

func BezierPath_SetDefaultMiterLimit(value coregraphics.Float) {
	C.C_NSBezierPath_BezierPath_SetDefaultMiterLimit(C.double(float64(value)))
}

func BezierPath_DefaultFlatness() coregraphics.Float {
	result_ := C.C_NSBezierPath_BezierPath_DefaultFlatness()
	return coregraphics.Float(float64(result_))
}

func BezierPath_SetDefaultFlatness(value coregraphics.Float) {
	C.C_NSBezierPath_BezierPath_SetDefaultFlatness(C.double(float64(value)))
}

func (n *NSBezierPath) Bounds() foundation.Rect {
	result_ := C.C_NSBezierPath_Bounds(n.Ptr())
	return foundation.Rect(coregraphics.FromCGRectPointer(unsafe.Pointer(&result_)))
}

func (n *NSBezierPath) ControlPointBounds() foundation.Rect {
	result_ := C.C_NSBezierPath_ControlPointBounds(n.Ptr())
	return foundation.Rect(coregraphics.FromCGRectPointer(unsafe.Pointer(&result_)))
}

func (n *NSBezierPath) CurrentPoint() foundation.Point {
	result_ := C.C_NSBezierPath_CurrentPoint(n.Ptr())
	return foundation.Point(coregraphics.FromCGPointPointer(unsafe.Pointer(&result_)))
}

func (n *NSBezierPath) IsEmpty() bool {
	result_ := C.C_NSBezierPath_IsEmpty(n.Ptr())
	return bool(result_)
}

func (n *NSBezierPath) ElementCount() int {
	result_ := C.C_NSBezierPath_ElementCount(n.Ptr())
	return int(result_)
}
