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

func AllocViewAnimation() NSViewAnimation {
	return MakeViewAnimation(C.C_ViewAnimation_Alloc())
}

func (n NSViewAnimation) InitWithDuration_AnimationCurve(duration foundation.TimeInterval, animationCurve AnimationCurve) ViewAnimation {
	result_ := C.C_NSViewAnimation_InitWithDuration_AnimationCurve(n.Ptr(), C.double(float64(duration)), C.uint(uint(animationCurve)))
	return MakeViewAnimation(result_)
}

func (n NSViewAnimation) InitWithCoder(coder foundation.Coder) ViewAnimation {
	result_ := C.C_NSViewAnimation_InitWithCoder(n.Ptr(), objc.ExtractPtr(coder))
	return MakeViewAnimation(result_)
}

func (n NSViewAnimation) Init() ViewAnimation {
	result_ := C.C_NSViewAnimation_Init(n.Ptr())
	return MakeViewAnimation(result_)
}
