package appkit

// #include "tracking_area.h"
import "C"
import (
	"github.com/hsiafan/cocoa/coregraphics"
	"github.com/hsiafan/cocoa/foundation"
	"github.com/hsiafan/cocoa/objc"
	"unsafe"
)

type TrackingArea interface {
	objc.Object
	Options() TrackingAreaOptions
	Owner() objc.Object
	Rect() foundation.Rect
	UserInfo() map[objc.Object]objc.Object
}

type NSTrackingArea struct {
	objc.NSObject
}

func MakeTrackingArea(ptr unsafe.Pointer) NSTrackingArea {
	return NSTrackingArea{
		NSObject: objc.MakeObject(ptr),
	}
}

func AllocTrackingArea() NSTrackingArea {
	return MakeTrackingArea(C.C_TrackingArea_Alloc())
}

func (n NSTrackingArea) InitWithRect_Options_Owner_UserInfo(rect foundation.Rect, options TrackingAreaOptions, owner objc.Object, userInfo map[objc.Object]objc.Object) TrackingArea {
	var cUserInfo C.Dictionary
	if len(userInfo) > 0 {
		cUserInfoKeyData := make([]unsafe.Pointer, len(userInfo))
		cUserInfoValueData := make([]unsafe.Pointer, len(userInfo))
		var idx = 0
		for k, v := range userInfo {
			cUserInfoKeyData[idx] = objc.ExtractPtr(k)
			cUserInfoValueData[idx] = objc.ExtractPtr(v)
			idx++
		}
		cUserInfo.key_data = unsafe.Pointer(&cUserInfoKeyData[0])
		cUserInfo.value_data = unsafe.Pointer(&cUserInfoValueData[0])
		cUserInfo.len = C.int(len(userInfo))
	}
	result_ := C.C_NSTrackingArea_InitWithRect_Options_Owner_UserInfo(n.Ptr(), *(*C.CGRect)(coregraphics.ToCGRectPointer(coregraphics.Rect(rect))), C.uint(uint(options)), objc.ExtractPtr(owner), cUserInfo)
	return MakeTrackingArea(result_)
}

func (n NSTrackingArea) Init() TrackingArea {
	result_ := C.C_NSTrackingArea_Init(n.Ptr())
	return MakeTrackingArea(result_)
}

func (n NSTrackingArea) Options() TrackingAreaOptions {
	result_ := C.C_NSTrackingArea_Options(n.Ptr())
	return TrackingAreaOptions(uint(result_))
}

func (n NSTrackingArea) Owner() objc.Object {
	result_ := C.C_NSTrackingArea_Owner(n.Ptr())
	return objc.MakeObject(result_)
}

func (n NSTrackingArea) Rect() foundation.Rect {
	result_ := C.C_NSTrackingArea_Rect(n.Ptr())
	return foundation.Rect(coregraphics.FromCGRectPointer(unsafe.Pointer(&result_)))
}

func (n NSTrackingArea) UserInfo() map[objc.Object]objc.Object {
	result_ := C.C_NSTrackingArea_UserInfo(n.Ptr())
	if result_.len > 0 {
		defer C.free(result_.key_data)
		defer C.free(result_.value_data)
	}
	result_KeySlice := (*[1 << 28]unsafe.Pointer)(unsafe.Pointer(result_.key_data))[:result_.len:result_.len]
	result_ValueSlice := (*[1 << 28]unsafe.Pointer)(unsafe.Pointer(result_.value_data))[:result_.len:result_.len]
	var goResult_ = make(map[objc.Object]objc.Object)
	for idx, k := range result_KeySlice {
		v := result_ValueSlice[idx]
		goResult_[objc.MakeObject(k)] = objc.MakeObject(v)
	}
	return goResult_
}
