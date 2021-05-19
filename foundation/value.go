package foundation

// #include "value.h"
import "C"
import (
	"github.com/hsiafan/cocoa/coregraphics"
	"github.com/hsiafan/cocoa/objc"
	"unsafe"
)

type Value interface {
	objc.Object
	IsEqualToValue(value Value) bool
	NonretainedObjectValue() objc.Object
	RangeValue() Range
	PointValue() Point
	SizeValue() Size
	RectValue() Rect
	EdgeInsetsValue() EdgeInsets
}

type NSValue struct {
	objc.NSObject
}

func MakeValue(ptr unsafe.Pointer) *NSValue {
	if ptr == nil {
		return nil
	}
	return &NSValue{
		NSObject: *objc.MakeObject(ptr),
	}
}

func AllocValue() *NSValue {
	return MakeValue(C.C_Value_Alloc())
}

func (n *NSValue) InitWithCoder(coder Coder) Value {
	result_ := C.C_NSValue_InitWithCoder(n.Ptr(), objc.ExtractPtr(coder))
	return MakeValue(result_)
}

func (n *NSValue) Init() Value {
	result_ := C.C_NSValue_Init(n.Ptr())
	return MakeValue(result_)
}

func ValueWithNonretainedObject(anObject objc.Object) Value {
	result_ := C.C_NSValue_ValueWithNonretainedObject(objc.ExtractPtr(anObject))
	return MakeValue(result_)
}

func ValueWithRange(_range Range) Value {
	result_ := C.C_NSValue_ValueWithRange(*(*C.NSRange)(ToNSRangePointer(_range)))
	return MakeValue(result_)
}

func ValueWithPoint(point Point) Value {
	result_ := C.C_NSValue_ValueWithPoint(*(*C.CGPoint)(coregraphics.ToCGPointPointer(coregraphics.Point(point))))
	return MakeValue(result_)
}

func ValueWithSize(size Size) Value {
	result_ := C.C_NSValue_ValueWithSize(*(*C.CGSize)(coregraphics.ToCGSizePointer(coregraphics.Size(size))))
	return MakeValue(result_)
}

func ValueWithRect(rect Rect) Value {
	result_ := C.C_NSValue_ValueWithRect(*(*C.CGRect)(coregraphics.ToCGRectPointer(coregraphics.Rect(rect))))
	return MakeValue(result_)
}

func (n *NSValue) IsEqualToValue(value Value) bool {
	result_ := C.C_NSValue_IsEqualToValue(n.Ptr(), objc.ExtractPtr(value))
	return bool(result_)
}

func ValueWithEdgeInsets(insets EdgeInsets) Value {
	result_ := C.C_NSValue_ValueWithEdgeInsets(*(*C.NSEdgeInsets)(ToNSEdgeInsetsPointer(insets)))
	return MakeValue(result_)
}

func (n *NSValue) NonretainedObjectValue() objc.Object {
	result_ := C.C_NSValue_NonretainedObjectValue(n.Ptr())
	return objc.MakeObject(result_)
}

func (n *NSValue) RangeValue() Range {
	result_ := C.C_NSValue_RangeValue(n.Ptr())
	return FromNSRangePointer(unsafe.Pointer(&result_))
}

func (n *NSValue) PointValue() Point {
	result_ := C.C_NSValue_PointValue(n.Ptr())
	return Point(coregraphics.FromCGPointPointer(unsafe.Pointer(&result_)))
}

func (n *NSValue) SizeValue() Size {
	result_ := C.C_NSValue_SizeValue(n.Ptr())
	return Size(coregraphics.FromCGSizePointer(unsafe.Pointer(&result_)))
}

func (n *NSValue) RectValue() Rect {
	result_ := C.C_NSValue_RectValue(n.Ptr())
	return Rect(coregraphics.FromCGRectPointer(unsafe.Pointer(&result_)))
}

func (n *NSValue) EdgeInsetsValue() EdgeInsets {
	result_ := C.C_NSValue_EdgeInsetsValue(n.Ptr())
	return FromNSEdgeInsetsPointer(unsafe.Pointer(&result_))
}
