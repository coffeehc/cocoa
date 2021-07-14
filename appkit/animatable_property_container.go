package appkit

// #include "animatable_property_container.h"
import "C"
import (
	"github.com/hsiafan/cocoa/foundation"
	"github.com/hsiafan/cocoa/objc"
	"runtime/cgo"
	"unsafe"
)

type AnimatablePropertyContainer struct {
	Animator        func() objc.Object                          // required
	AnimationForKey func(key AnimatablePropertyKey) objc.Object // required
}

func WrapAnimatablePropertyContainer(delegate *AnimatablePropertyContainer) objc.Object {
	h := cgo.NewHandle(delegate)
	ptr := C.WrapAnimatablePropertyContainer(C.uintptr_t(h))
	return objc.MakeObject(ptr)
}

//export animatablePropertyContainer_Animator
func animatablePropertyContainer_Animator(hp C.uintptr_t) unsafe.Pointer {
	delegate := cgo.Handle(hp).Value().(*AnimatablePropertyContainer)
	result := delegate.Animator()
	return objc.ExtractPtr(result)
}

//export animatablePropertyContainer_AnimationForKey
func animatablePropertyContainer_AnimationForKey(hp C.uintptr_t, key unsafe.Pointer) unsafe.Pointer {
	delegate := cgo.Handle(hp).Value().(*AnimatablePropertyContainer)
	result := delegate.AnimationForKey(AnimatablePropertyKey(foundation.MakeString(key).String()))
	return objc.ExtractPtr(result)
}

//export AnimatablePropertyContainer_RespondsTo
func AnimatablePropertyContainer_RespondsTo(hp C.uintptr_t, selectorPtr unsafe.Pointer) bool {
	sel := objc.Selector(selectorPtr)
	selName := objc.Sel_GetName(sel)
	delegate := cgo.Handle(hp).Value().(*AnimatablePropertyContainer)
	switch selName {
	case "animator:":
		return delegate.Animator != nil
	case "animationForKey:":
		return delegate.AnimationForKey != nil
	default:
		return false
	}
}

//export deleteAnimatablePropertyContainer
func deleteAnimatablePropertyContainer(hp C.uintptr_t) {
	cgo.Handle(hp).Delete()
}
