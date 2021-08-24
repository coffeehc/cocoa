package webkit

// #include "snapshot_configuration.h"
import "C"
import (
	"github.com/hsiafan/cocoa/coregraphics"
	"github.com/hsiafan/cocoa/foundation"
	"github.com/hsiafan/cocoa/objc"
	"unsafe"
)

type SnapshotConfiguration interface {
	objc.Object
	Rect() coregraphics.Rect
	SetRect(value coregraphics.Rect)
	SnapshotWidth() foundation.Number
	SetSnapshotWidth(value foundation.Number)
	AfterScreenUpdates() bool
	SetAfterScreenUpdates(value bool)
}

type WKSnapshotConfiguration struct {
	objc.NSObject
}

func MakeSnapshotConfiguration(ptr unsafe.Pointer) WKSnapshotConfiguration {
	return WKSnapshotConfiguration{
		NSObject: objc.MakeObject(ptr),
	}
}

func AllocSnapshotConfiguration() WKSnapshotConfiguration {
	return MakeSnapshotConfiguration(C.C_SnapshotConfiguration_Alloc())
}

func (w WKSnapshotConfiguration) Init() SnapshotConfiguration {
	result_ := C.C_WKSnapshotConfiguration_Init(w.Ptr())
	return MakeSnapshotConfiguration(result_)
}

func (w WKSnapshotConfiguration) Rect() coregraphics.Rect {
	result_ := C.C_WKSnapshotConfiguration_Rect(w.Ptr())
	return coregraphics.FromCGRectPointer(unsafe.Pointer(&result_))
}

func (w WKSnapshotConfiguration) SetRect(value coregraphics.Rect) {
	C.C_WKSnapshotConfiguration_SetRect(w.Ptr(), *(*C.CGRect)(coregraphics.ToCGRectPointer(value)))
}

func (w WKSnapshotConfiguration) SnapshotWidth() foundation.Number {
	result_ := C.C_WKSnapshotConfiguration_SnapshotWidth(w.Ptr())
	return foundation.MakeNumber(result_)
}

func (w WKSnapshotConfiguration) SetSnapshotWidth(value foundation.Number) {
	C.C_WKSnapshotConfiguration_SetSnapshotWidth(w.Ptr(), objc.ExtractPtr(value))
}

func (w WKSnapshotConfiguration) AfterScreenUpdates() bool {
	result_ := C.C_WKSnapshotConfiguration_AfterScreenUpdates(w.Ptr())
	return bool(result_)
}

func (w WKSnapshotConfiguration) SetAfterScreenUpdates(value bool) {
	C.C_WKSnapshotConfiguration_SetAfterScreenUpdates(w.Ptr(), C.bool(value))
}
