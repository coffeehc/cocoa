package appkit

// #include "animation_delegate.h"
import "C"
import (
	"github.com/hsiafan/cocoa/objc"
	"unsafe"
)

type AnimationDelegate struct {
	AnimationDidEnd                func(animation Animation)
	AnimationDidStop               func(animation Animation)
	AnimationShouldStart           func(animation Animation) bool
	Animation_ValueForProgress     func(animation Animation, progress AnimationProgress) float32
	Animation_DidReachProgressMark func(animation Animation, progress AnimationProgress)
}

func WrapAnimationDelegate(delegate *AnimationDelegate) objc.Object {
	id := resources.NextId()
	resources.Store(id, delegate)
	ptr := C.WrapAnimationDelegate(C.long(id))
	return objc.MakeObject(ptr)
}

//export animationDelegate_AnimationDidEnd
func animationDelegate_AnimationDidEnd(id int64, animation unsafe.Pointer) {
	delegate := resources.Get(id).(*AnimationDelegate)
	delegate.AnimationDidEnd(MakeAnimation(animation))
}

//export animationDelegate_AnimationDidStop
func animationDelegate_AnimationDidStop(id int64, animation unsafe.Pointer) {
	delegate := resources.Get(id).(*AnimationDelegate)
	delegate.AnimationDidStop(MakeAnimation(animation))
}

//export animationDelegate_AnimationShouldStart
func animationDelegate_AnimationShouldStart(id int64, animation unsafe.Pointer) C.bool {
	delegate := resources.Get(id).(*AnimationDelegate)
	result := delegate.AnimationShouldStart(MakeAnimation(animation))
	return C.bool(result)
}

//export animationDelegate_Animation_ValueForProgress
func animationDelegate_Animation_ValueForProgress(id int64, animation unsafe.Pointer, progress C.float) C.float {
	delegate := resources.Get(id).(*AnimationDelegate)
	result := delegate.Animation_ValueForProgress(MakeAnimation(animation), AnimationProgress(float32(progress)))
	return C.float(result)
}

//export animationDelegate_Animation_DidReachProgressMark
func animationDelegate_Animation_DidReachProgressMark(id int64, animation unsafe.Pointer, progress C.float) {
	delegate := resources.Get(id).(*AnimationDelegate)
	delegate.Animation_DidReachProgressMark(MakeAnimation(animation), AnimationProgress(float32(progress)))
}

//export AnimationDelegate_RespondsTo
func AnimationDelegate_RespondsTo(id int64, selectorPtr unsafe.Pointer) bool {
	sel := objc.Selector(selectorPtr)
	selName := objc.Sel_GetName(sel)
	delegate := resources.Get(id).(*AnimationDelegate)
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

//export deleteAnimationDelegate
func deleteAnimationDelegate(id int64) {
	resources.Delete(id)
}
