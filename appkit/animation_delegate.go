package appkit

// #include "animation_delegate.h"
import "C"
import (
	"github.com/hsiafan/cocoa/objc"
	"runtime/cgo"
	"unsafe"
)

type AnimationDelegate struct {
	AnimationDidEnd                func(animation Animation)
	AnimationDidStop               func(animation Animation)
	AnimationShouldStart           func(animation Animation) bool
	Animation_ValueForProgress     func(animation Animation, progress AnimationProgress) float32
	Animation_DidReachProgressMark func(animation Animation, progress AnimationProgress)
}

func (delegate *AnimationDelegate) ToObjc() objc.Object {
	h := cgo.NewHandle(delegate)
	ptr := C.WrapAnimationDelegate(C.uintptr_t(h))
	return objc.MakeObject(ptr)
}

//export animationDelegate_AnimationDidEnd
func animationDelegate_AnimationDidEnd(hp C.uintptr_t, animation unsafe.Pointer) {
	delegate := cgo.Handle(hp).Value().(*AnimationDelegate)
	delegate.AnimationDidEnd(MakeAnimation(animation))
}

//export animationDelegate_AnimationDidStop
func animationDelegate_AnimationDidStop(hp C.uintptr_t, animation unsafe.Pointer) {
	delegate := cgo.Handle(hp).Value().(*AnimationDelegate)
	delegate.AnimationDidStop(MakeAnimation(animation))
}

//export animationDelegate_AnimationShouldStart
func animationDelegate_AnimationShouldStart(hp C.uintptr_t, animation unsafe.Pointer) C.bool {
	delegate := cgo.Handle(hp).Value().(*AnimationDelegate)
	result := delegate.AnimationShouldStart(MakeAnimation(animation))
	return C.bool(result)
}

//export animationDelegate_Animation_ValueForProgress
func animationDelegate_Animation_ValueForProgress(hp C.uintptr_t, animation unsafe.Pointer, progress C.float) C.float {
	delegate := cgo.Handle(hp).Value().(*AnimationDelegate)
	result := delegate.Animation_ValueForProgress(MakeAnimation(animation), AnimationProgress(float32(progress)))
	return C.float(result)
}

//export animationDelegate_Animation_DidReachProgressMark
func animationDelegate_Animation_DidReachProgressMark(hp C.uintptr_t, animation unsafe.Pointer, progress C.float) {
	delegate := cgo.Handle(hp).Value().(*AnimationDelegate)
	delegate.Animation_DidReachProgressMark(MakeAnimation(animation), AnimationProgress(float32(progress)))
}

//export AnimationDelegate_RespondsTo
func AnimationDelegate_RespondsTo(hp C.uintptr_t, selectorPtr unsafe.Pointer) bool {
	sel := objc.Selector(selectorPtr)
	selName := objc.Sel_GetName(sel)
	delegate := cgo.Handle(hp).Value().(*AnimationDelegate)
	switch selName {
	case "animationDidEnd:":
		return delegate.AnimationDidEnd != nil
	case "animationDidStop:":
		return delegate.AnimationDidStop != nil
	case "animationShouldStart:":
		return delegate.AnimationShouldStart != nil
	case "animation:valueForProgress:":
		return delegate.Animation_ValueForProgress != nil
	case "animation:didReachProgressMark:":
		return delegate.Animation_DidReachProgressMark != nil
	default:
		return false
	}
}
