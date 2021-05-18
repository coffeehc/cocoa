package foundation

// #cgo CFLAGS: -x objective-c
// #cgo LDFLAGS: -framework Foundation
// #include "coder.h"
import "C"
import (
	"github.com/hsiafan/cocoa/coregraphics"
	"github.com/hsiafan/cocoa/objc"
	"unsafe"
)

type Coder interface {
	objc.Object
	ContainsValueForKey(key string) bool
	EncodeBool_ForKey(value bool, key string)
	EncodeBycopyObject(anObject objc.Object)
	EncodeByrefObject(anObject objc.Object)
	EncodeConditionalObject(object objc.Object)
	EncodeConditionalObject_ForKey(object objc.Object, key string)
	EncodeDataObject(data []byte)
	EncodeDouble_ForKey(value float64, key string)
	EncodeFloat_ForKey(value float32, key string)
	EncodeInteger_ForKey(value int, key string)
	EncodeInt32_ForKey(value int32, key string)
	EncodeInt64_ForKey(value int64, key string)
	EncodeObject(object objc.Object)
	EncodeObject_ForKey(object objc.Object, key string)
	EncodePoint(point Point)
	EncodePoint_ForKey(point Point, key string)
	EncodePropertyList(aPropertyList objc.Object)
	EncodeRect(rect Rect)
	EncodeRect_ForKey(rect Rect, key string)
	EncodeRootObject(rootObject objc.Object)
	EncodeSize(size Size)
	EncodeSize_ForKey(size Size, key string)
	DecodeBoolForKey(key string) bool
	DecodeDataObject() []byte
	DecodeDoubleForKey(key string) float64
	DecodeFloatForKey(key string) float32
	DecodeIntegerForKey(key string) int
	DecodeInt32ForKey(key string) int32
	DecodeInt64ForKey(key string) int64
	DecodeObject() objc.Object
	DecodeObjectForKey(key string) objc.Object
	DecodePoint() Point
	DecodePointForKey(key string) Point
	DecodePropertyList() objc.Object
	DecodeRect() Rect
	DecodeRectForKey(key string) Rect
	DecodeSize() Size
	DecodeSizeForKey(key string) Size
	DecodeObjectOfClasses_ForKey(classes Set, key string) objc.Object
	DecodePropertyListForKey(key string) objc.Object
	FailWithError(error Error)
	VersionForClassName(className string) int
	DecodeArrayOfObjectsOfClasses_ForKey(classes Set, key string) []objc.Object
	AllowsKeyedCoding() bool
	DecodingFailurePolicy() DecodingFailurePolicy
	RequiresSecureCoding() bool
	AllowedClasses() Set
	Error() Error
}

type NSCoder struct {
	objc.NSObject
}

func MakeCoder(ptr unsafe.Pointer) *NSCoder {
	if ptr == nil {
		return nil
	}
	return &NSCoder{
		NSObject: *objc.MakeObject(ptr),
	}
}

func AllocCoder() *NSCoder {
	return MakeCoder(C.C_Coder_Alloc())
}

func (n *NSCoder) Init() Coder {
	result_ := C.C_NSCoder_Init(n.Ptr())
	return MakeCoder(result_)
}

func (n *NSCoder) ContainsValueForKey(key string) bool {
	result_ := C.C_NSCoder_ContainsValueForKey(n.Ptr(), NewString(key).Ptr())
	return bool(result_)
}

func (n *NSCoder) EncodeBool_ForKey(value bool, key string) {
	C.C_NSCoder_EncodeBool_ForKey(n.Ptr(), C.bool(value), NewString(key).Ptr())
}

func (n *NSCoder) EncodeBycopyObject(anObject objc.Object) {
	C.C_NSCoder_EncodeBycopyObject(n.Ptr(), objc.ExtractPtr(anObject))
}

func (n *NSCoder) EncodeByrefObject(anObject objc.Object) {
	C.C_NSCoder_EncodeByrefObject(n.Ptr(), objc.ExtractPtr(anObject))
}

func (n *NSCoder) EncodeConditionalObject(object objc.Object) {
	C.C_NSCoder_EncodeConditionalObject(n.Ptr(), objc.ExtractPtr(object))
}

func (n *NSCoder) EncodeConditionalObject_ForKey(object objc.Object, key string) {
	C.C_NSCoder_EncodeConditionalObject_ForKey(n.Ptr(), objc.ExtractPtr(object), NewString(key).Ptr())
}

func (n *NSCoder) EncodeDataObject(data []byte) {
	C.C_NSCoder_EncodeDataObject(n.Ptr(), C.Array{data: unsafe.Pointer(&data[0]), len: C.int(len(data))})
}

func (n *NSCoder) EncodeDouble_ForKey(value float64, key string) {
	C.C_NSCoder_EncodeDouble_ForKey(n.Ptr(), C.double(value), NewString(key).Ptr())
}

func (n *NSCoder) EncodeFloat_ForKey(value float32, key string) {
	C.C_NSCoder_EncodeFloat_ForKey(n.Ptr(), C.float(value), NewString(key).Ptr())
}

func (n *NSCoder) EncodeInteger_ForKey(value int, key string) {
	C.C_NSCoder_EncodeInteger_ForKey(n.Ptr(), C.int(value), NewString(key).Ptr())
}

func (n *NSCoder) EncodeInt32_ForKey(value int32, key string) {
	C.C_NSCoder_EncodeInt32_ForKey(n.Ptr(), C.int(value), NewString(key).Ptr())
}

func (n *NSCoder) EncodeInt64_ForKey(value int64, key string) {
	C.C_NSCoder_EncodeInt64_ForKey(n.Ptr(), C.long(value), NewString(key).Ptr())
}

func (n *NSCoder) EncodeObject(object objc.Object) {
	C.C_NSCoder_EncodeObject(n.Ptr(), objc.ExtractPtr(object))
}

func (n *NSCoder) EncodeObject_ForKey(object objc.Object, key string) {
	C.C_NSCoder_EncodeObject_ForKey(n.Ptr(), objc.ExtractPtr(object), NewString(key).Ptr())
}

func (n *NSCoder) EncodePoint(point Point) {
	C.C_NSCoder_EncodePoint(n.Ptr(), *(*C.CGPoint)(coregraphics.ToCGPointPointer(coregraphics.Point(point))))
}

func (n *NSCoder) EncodePoint_ForKey(point Point, key string) {
	C.C_NSCoder_EncodePoint_ForKey(n.Ptr(), *(*C.CGPoint)(coregraphics.ToCGPointPointer(coregraphics.Point(point))), NewString(key).Ptr())
}

func (n *NSCoder) EncodePropertyList(aPropertyList objc.Object) {
	C.C_NSCoder_EncodePropertyList(n.Ptr(), objc.ExtractPtr(aPropertyList))
}

func (n *NSCoder) EncodeRect(rect Rect) {
	C.C_NSCoder_EncodeRect(n.Ptr(), *(*C.CGRect)(coregraphics.ToCGRectPointer(coregraphics.Rect(rect))))
}

func (n *NSCoder) EncodeRect_ForKey(rect Rect, key string) {
	C.C_NSCoder_EncodeRect_ForKey(n.Ptr(), *(*C.CGRect)(coregraphics.ToCGRectPointer(coregraphics.Rect(rect))), NewString(key).Ptr())
}

func (n *NSCoder) EncodeRootObject(rootObject objc.Object) {
	C.C_NSCoder_EncodeRootObject(n.Ptr(), objc.ExtractPtr(rootObject))
}

func (n *NSCoder) EncodeSize(size Size) {
	C.C_NSCoder_EncodeSize(n.Ptr(), *(*C.CGSize)(coregraphics.ToCGSizePointer(coregraphics.Size(size))))
}

func (n *NSCoder) EncodeSize_ForKey(size Size, key string) {
	C.C_NSCoder_EncodeSize_ForKey(n.Ptr(), *(*C.CGSize)(coregraphics.ToCGSizePointer(coregraphics.Size(size))), NewString(key).Ptr())
}

func (n *NSCoder) DecodeBoolForKey(key string) bool {
	result_ := C.C_NSCoder_DecodeBoolForKey(n.Ptr(), NewString(key).Ptr())
	return bool(result_)
}

func (n *NSCoder) DecodeDataObject() []byte {
	result_ := C.C_NSCoder_DecodeDataObject(n.Ptr())
	result_Buffer := (*[1 << 30]byte)(result_.data)[:C.int(result_.len)]
	goResult_ := make([]byte, C.int(result_.len))
	copy(goResult_, result_Buffer)
	C.free(result_.data)
	return goResult_
}

func (n *NSCoder) DecodeDoubleForKey(key string) float64 {
	result_ := C.C_NSCoder_DecodeDoubleForKey(n.Ptr(), NewString(key).Ptr())
	return float64(result_)
}

func (n *NSCoder) DecodeFloatForKey(key string) float32 {
	result_ := C.C_NSCoder_DecodeFloatForKey(n.Ptr(), NewString(key).Ptr())
	return float32(result_)
}

func (n *NSCoder) DecodeIntegerForKey(key string) int {
	result_ := C.C_NSCoder_DecodeIntegerForKey(n.Ptr(), NewString(key).Ptr())
	return int(result_)
}

func (n *NSCoder) DecodeInt32ForKey(key string) int32 {
	result_ := C.C_NSCoder_DecodeInt32ForKey(n.Ptr(), NewString(key).Ptr())
	return int32(result_)
}

func (n *NSCoder) DecodeInt64ForKey(key string) int64 {
	result_ := C.C_NSCoder_DecodeInt64ForKey(n.Ptr(), NewString(key).Ptr())
	return int64(result_)
}

func (n *NSCoder) DecodeObject() objc.Object {
	result_ := C.C_NSCoder_DecodeObject(n.Ptr())
	return objc.MakeObject(result_)
}

func (n *NSCoder) DecodeObjectForKey(key string) objc.Object {
	result_ := C.C_NSCoder_DecodeObjectForKey(n.Ptr(), NewString(key).Ptr())
	return objc.MakeObject(result_)
}

func (n *NSCoder) DecodePoint() Point {
	result_ := C.C_NSCoder_DecodePoint(n.Ptr())
	return Point(coregraphics.FromCGPointPointer(unsafe.Pointer(&result_)))
}

func (n *NSCoder) DecodePointForKey(key string) Point {
	result_ := C.C_NSCoder_DecodePointForKey(n.Ptr(), NewString(key).Ptr())
	return Point(coregraphics.FromCGPointPointer(unsafe.Pointer(&result_)))
}

func (n *NSCoder) DecodePropertyList() objc.Object {
	result_ := C.C_NSCoder_DecodePropertyList(n.Ptr())
	return objc.MakeObject(result_)
}

func (n *NSCoder) DecodeRect() Rect {
	result_ := C.C_NSCoder_DecodeRect(n.Ptr())
	return Rect(coregraphics.FromCGRectPointer(unsafe.Pointer(&result_)))
}

func (n *NSCoder) DecodeRectForKey(key string) Rect {
	result_ := C.C_NSCoder_DecodeRectForKey(n.Ptr(), NewString(key).Ptr())
	return Rect(coregraphics.FromCGRectPointer(unsafe.Pointer(&result_)))
}

func (n *NSCoder) DecodeSize() Size {
	result_ := C.C_NSCoder_DecodeSize(n.Ptr())
	return Size(coregraphics.FromCGSizePointer(unsafe.Pointer(&result_)))
}

func (n *NSCoder) DecodeSizeForKey(key string) Size {
	result_ := C.C_NSCoder_DecodeSizeForKey(n.Ptr(), NewString(key).Ptr())
	return Size(coregraphics.FromCGSizePointer(unsafe.Pointer(&result_)))
}

func (n *NSCoder) DecodeObjectOfClasses_ForKey(classes Set, key string) objc.Object {
	result_ := C.C_NSCoder_DecodeObjectOfClasses_ForKey(n.Ptr(), objc.ExtractPtr(classes), NewString(key).Ptr())
	return objc.MakeObject(result_)
}

func (n *NSCoder) DecodePropertyListForKey(key string) objc.Object {
	result_ := C.C_NSCoder_DecodePropertyListForKey(n.Ptr(), NewString(key).Ptr())
	return objc.MakeObject(result_)
}

func (n *NSCoder) FailWithError(error Error) {
	C.C_NSCoder_FailWithError(n.Ptr(), objc.ExtractPtr(error))
}

func (n *NSCoder) VersionForClassName(className string) int {
	result_ := C.C_NSCoder_VersionForClassName(n.Ptr(), NewString(className).Ptr())
	return int(result_)
}

func (n *NSCoder) DecodeArrayOfObjectsOfClasses_ForKey(classes Set, key string) []objc.Object {
	result_ := C.C_NSCoder_DecodeArrayOfObjectsOfClasses_ForKey(n.Ptr(), objc.ExtractPtr(classes), NewString(key).Ptr())
	defer C.free(result_.data)
	result_Slice := (*[1 << 28]unsafe.Pointer)(unsafe.Pointer(result_.data))[:result_.len:result_.len]
	var goResult_ = make([]objc.Object, len(result_Slice))
	for idx, r := range result_Slice {
		goResult_[idx] = objc.MakeObject(r)
	}
	return goResult_
}

func (n *NSCoder) AllowsKeyedCoding() bool {
	result_ := C.C_NSCoder_AllowsKeyedCoding(n.Ptr())
	return bool(result_)
}

func (n *NSCoder) DecodingFailurePolicy() DecodingFailurePolicy {
	result_ := C.C_NSCoder_DecodingFailurePolicy(n.Ptr())
	return DecodingFailurePolicy(int(result_))
}

func (n *NSCoder) RequiresSecureCoding() bool {
	result_ := C.C_NSCoder_RequiresSecureCoding(n.Ptr())
	return bool(result_)
}

func (n *NSCoder) AllowedClasses() Set {
	result_ := C.C_NSCoder_AllowedClasses(n.Ptr())
	return MakeSet(result_)
}

func (n *NSCoder) Error() Error {
	result_ := C.C_NSCoder_Error(n.Ptr())
	return MakeError(result_)
}
