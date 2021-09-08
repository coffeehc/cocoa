package appkit

// #include "view_animation.h"
import "C"
import (
	"github.com/hsiafan/cocoa/foundation"
	"github.com/hsiafan/cocoa/objc"
	"unsafe"
)

type ViewAnimation interface {
	Animation
}

type NSViewAnimation struct {
	NSAnimation
}

func MakeViewAnimation(ptr unsafe.Pointer) NSViewAnimation {
	return NSViewAnimation{
		NSAnimation: MakeAnimation(ptr),
	}
}

func (n NSViewAnimation) InitWithDuration_AnimationCurve(duration foundation.TimeInterval, animationCurve AnimationCurve) NSViewAnimation {
	result_ := C.C_NSViewAnimation_InitWithDuration_AnimationCurve(n.Ptr(), C.double(float64(duration)), C.uint(uint(animationCurve)))
	return MakeViewAnimation(result_)
}

func (n NSViewAnimation) InitWithCoder(coder foundation.Coder) NSViewAnimation {
	result_ := C.C_NSViewAnimation_InitWithCoder(n.Ptr(), objc.ExtractPtr(coder))
	return MakeViewAnimation(result_)
}

func AllocViewAnimation() NSViewAnimation {
	result_ := C.C_NSViewAnimation_AllocViewAnimation()
	return MakeViewAnimation(result_)
}

func (n NSViewAnimation) Init() NSViewAnimation {
	result_ := C.C_NSViewAnimation_Init(n.Ptr())
	return MakeViewAnimation(result_)
}

func NewViewAnimation() NSViewAnimation {
	result_ := C.C_NSViewAnimation_NewViewAnimation()
	return MakeViewAnimation(result_)
}

func (n NSViewAnimation) Autorelease() NSViewAnimation {
	result_ := C.C_NSViewAnimation_Autorelease(n.Ptr())
	return MakeViewAnimation(result_)
}

func (n NSViewAnimation) Retain() NSViewAnimation {
	result_ := C.C_NSViewAnimation_Retain(n.Ptr())
	return MakeViewAnimation(result_)
}
