package appkit

// #include "animatable_property_container.h"
import "C"
import (
	"github.com/hsiafan/cocoa/foundation"
	"github.com/hsiafan/cocoa/objc"
	"unsafe"
)

type AnimatablePropertyContainer struct {
	Animator        func() objc.Object
	AnimationForKey func(key AnimatablePropertyKey) objc.Object
}

func WrapAnimatablePropertyContainer(delegate *AnimatablePropertyContainer) objc.Object {
	id := resources.NextId()
	resources.Store(id, delegate)
	ptr := C.WrapAnimatablePropertyContainer(C.long(id))
	return objc.MakeObject(ptr)
}

//export animatablePropertyContainer_Animator
func animatablePropertyContainer_Animator(id int64) unsafe.Pointer {
	delegate := resources.Get(id).(*AnimatablePropertyContainer)
	result := delegate.Animator()
	return objc.ExtractPtr(result)
}

//export animatablePropertyContainer_AnimationForKey
func animatablePropertyContainer_AnimationForKey(id int64, key unsafe.Pointer) unsafe.Pointer {
	delegate := resources.Get(id).(*AnimatablePropertyContainer)
	result := delegate.AnimationForKey(AnimatablePropertyKey(foundation.MakeString(key).String()))
	return objc.ExtractPtr(result)
}

//export AnimatablePropertyContainer_RespondsTo
func AnimatablePropertyContainer_RespondsTo(id int64, selectorPtr unsafe.Pointer) bool {
	sel := objc.MakeSelector(selectorPtr)
	selName := objc.Sel_GetName(sel)
	delegate := resources.Get(id).(*AnimatablePropertyContainer)
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
func deleteAnimatablePropertyContainer(id int64) {
	resources.Delete(id)
}
