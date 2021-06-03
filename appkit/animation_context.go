package appkit

// #include "animation_context.h"
import "C"
import (
	"github.com/hsiafan/cocoa/foundation"
	"github.com/hsiafan/cocoa/objc"
	"unsafe"
)

type AnimationContext interface {
	objc.Object
	Duration() foundation.TimeInterval
	SetDuration(value foundation.TimeInterval)
	AllowsImplicitAnimation() bool
	SetAllowsImplicitAnimation(value bool)
}

type NSAnimationContext struct {
	objc.NSObject
}

func MakeAnimationContext(ptr unsafe.Pointer) NSAnimationContext {
	return NSAnimationContext{
		NSObject: objc.MakeObject(ptr),
	}
}

func AllocAnimationContext() NSAnimationContext {
	return MakeAnimationContext(C.C_AnimationContext_Alloc())
}

func (n NSAnimationContext) Init() AnimationContext {
	result_ := C.C_NSAnimationContext_Init(n.Ptr())
	return MakeAnimationContext(result_)
}

func AnimationContext_BeginGrouping() {
	C.C_NSAnimationContext_AnimationContext_BeginGrouping()
}

func AnimationContext_EndGrouping() {
	C.C_NSAnimationContext_AnimationContext_EndGrouping()
}

func AnimationContext_CurrentContext() AnimationContext {
	result_ := C.C_NSAnimationContext_AnimationContext_CurrentContext()
	return MakeAnimationContext(result_)
}

func (n NSAnimationContext) Duration() foundation.TimeInterval {
	result_ := C.C_NSAnimationContext_Duration(n.Ptr())
	return foundation.TimeInterval(float64(result_))
}

func (n NSAnimationContext) SetDuration(value foundation.TimeInterval) {
	C.C_NSAnimationContext_SetDuration(n.Ptr(), C.double(float64(value)))
}

func (n NSAnimationContext) AllowsImplicitAnimation() bool {
	result_ := C.C_NSAnimationContext_AllowsImplicitAnimation(n.Ptr())
	return bool(result_)
}

func (n NSAnimationContext) SetAllowsImplicitAnimation(value bool) {
	C.C_NSAnimationContext_SetAllowsImplicitAnimation(n.Ptr(), C.bool(value))
}
