package foundation

// #include "affine_transform.h"
import "C"
import (
	"github.com/hsiafan/cocoa/coregraphics"
	"github.com/hsiafan/cocoa/objc"
	"unsafe"
)

type AffineTransform interface {
	objc.Object
	RotateByDegrees(angle coregraphics.Float)
	RotateByRadians(angle coregraphics.Float)
	ScaleBy(scale coregraphics.Float)
	ScaleXBy_YBy(scaleX coregraphics.Float, scaleY coregraphics.Float)
	TranslateXBy_YBy(deltaX coregraphics.Float, deltaY coregraphics.Float)
	AppendTransform(transform AffineTransform)
	PrependTransform(transform AffineTransform)
	Invert()
	TransformPoint(aPoint Point) Point
	TransformSize(aSize Size) Size
	Set()
	Concat()
	TransformStruct() AffineTransformStruct
	SetTransformStruct(value AffineTransformStruct)
}

type NSAffineTransform struct {
	objc.NSObject
}

func MakeAffineTransform(ptr unsafe.Pointer) NSAffineTransform {
	return NSAffineTransform{
		NSObject: objc.MakeObject(ptr),
	}
}

func (n NSAffineTransform) InitWithTransform(transform AffineTransform) NSAffineTransform {
	result_ := C.C_NSAffineTransform_InitWithTransform(n.Ptr(), objc.ExtractPtr(transform))
	return MakeAffineTransform(result_)
}

func AllocAffineTransform() NSAffineTransform {
	result_ := C.C_NSAffineTransform_AllocAffineTransform()
	return MakeAffineTransform(result_)
}

func (n NSAffineTransform) Autorelease() NSAffineTransform {
	result_ := C.C_NSAffineTransform_Autorelease(n.Ptr())
	return MakeAffineTransform(result_)
}

func (n NSAffineTransform) Retain() NSAffineTransform {
	result_ := C.C_NSAffineTransform_Retain(n.Ptr())
	return MakeAffineTransform(result_)
}

func AffineTransform_Transform() AffineTransform {
	result_ := C.C_NSAffineTransform_AffineTransform_Transform()
	return MakeAffineTransform(result_)
}

func (n NSAffineTransform) RotateByDegrees(angle coregraphics.Float) {
	C.C_NSAffineTransform_RotateByDegrees(n.Ptr(), C.double(float64(angle)))
}

func (n NSAffineTransform) RotateByRadians(angle coregraphics.Float) {
	C.C_NSAffineTransform_RotateByRadians(n.Ptr(), C.double(float64(angle)))
}

func (n NSAffineTransform) ScaleBy(scale coregraphics.Float) {
	C.C_NSAffineTransform_ScaleBy(n.Ptr(), C.double(float64(scale)))
}

func (n NSAffineTransform) ScaleXBy_YBy(scaleX coregraphics.Float, scaleY coregraphics.Float) {
	C.C_NSAffineTransform_ScaleXBy_YBy(n.Ptr(), C.double(float64(scaleX)), C.double(float64(scaleY)))
}

func (n NSAffineTransform) TranslateXBy_YBy(deltaX coregraphics.Float, deltaY coregraphics.Float) {
	C.C_NSAffineTransform_TranslateXBy_YBy(n.Ptr(), C.double(float64(deltaX)), C.double(float64(deltaY)))
}

func (n NSAffineTransform) AppendTransform(transform AffineTransform) {
	C.C_NSAffineTransform_AppendTransform(n.Ptr(), objc.ExtractPtr(transform))
}

func (n NSAffineTransform) PrependTransform(transform AffineTransform) {
	C.C_NSAffineTransform_PrependTransform(n.Ptr(), objc.ExtractPtr(transform))
}

func (n NSAffineTransform) Invert() {
	C.C_NSAffineTransform_Invert(n.Ptr())
}

func (n NSAffineTransform) TransformPoint(aPoint Point) Point {
	result_ := C.C_NSAffineTransform_TransformPoint(n.Ptr(), *(*C.CGPoint)(coregraphics.ToCGPointPointer(coregraphics.Point(aPoint))))
	return Point(coregraphics.FromCGPointPointer(unsafe.Pointer(&result_)))
}

func (n NSAffineTransform) TransformSize(aSize Size) Size {
	result_ := C.C_NSAffineTransform_TransformSize(n.Ptr(), *(*C.CGSize)(coregraphics.ToCGSizePointer(coregraphics.Size(aSize))))
	return Size(coregraphics.FromCGSizePointer(unsafe.Pointer(&result_)))
}

func (n NSAffineTransform) Set() {
	C.C_NSAffineTransform_Set(n.Ptr())
}

func (n NSAffineTransform) Concat() {
	C.C_NSAffineTransform_Concat(n.Ptr())
}

func (n NSAffineTransform) TransformStruct() AffineTransformStruct {
	result_ := C.C_NSAffineTransform_TransformStruct(n.Ptr())
	return FromNSAffineTransformStructPointer(unsafe.Pointer(&result_))
}

func (n NSAffineTransform) SetTransformStruct(value AffineTransformStruct) {
	C.C_NSAffineTransform_SetTransformStruct(n.Ptr(), *(*C.NSAffineTransformStruct)(ToNSAffineTransformStructPointer(value)))
}
