package appkit

// #include "animation.h"
import "C"
import (
	"github.com/hsiafan/cocoa/foundation"
	"github.com/hsiafan/cocoa/objc"
	"runtime/cgo"
	"unsafe"
)

type AnimatablePropertyContainer interface {
	Animator() objc.Object
	AnimationForKey(key AnimatablePropertyKey) objc.Object
	Animations() map[AnimatablePropertyKey]objc.Object
	SetAnimations(value map[AnimatablePropertyKey]objc.Object)
}

func AnimatablePropertyContainerToObjc(protocol AnimatablePropertyContainer) objc.Object {
	h := cgo.NewHandle(protocol)
	ptr := C.WrapAnimatablePropertyContainer(C.uintptr_t(h))
	return objc.MakeObject(ptr)
}

//export animatablePropertyContainer_Animator
func animatablePropertyContainer_Animator(hp C.uintptr_t) unsafe.Pointer {
	protocol := cgo.Handle(hp).Value().(AnimatablePropertyContainer)
	result := protocol.Animator()
	return objc.ExtractPtr(result)
}

//export animatablePropertyContainer_AnimationForKey
func animatablePropertyContainer_AnimationForKey(hp C.uintptr_t, key unsafe.Pointer) unsafe.Pointer {
	protocol := cgo.Handle(hp).Value().(AnimatablePropertyContainer)
	result := protocol.AnimationForKey(AnimatablePropertyKey(foundation.MakeString(key).String()))
	return objc.ExtractPtr(result)
}

//export animatablePropertyContainer_SetAnimations
func animatablePropertyContainer_SetAnimations(hp C.uintptr_t, value C.Dictionary) {
	protocol := cgo.Handle(hp).Value().(AnimatablePropertyContainer)
	if value.len > 0 {
		defer C.free(value.key_data)
		defer C.free(value.value_data)
	}
	valueKeySlice := unsafe.Slice((*unsafe.Pointer)(value.key_data), int(value.len))
	valueValueSlice := unsafe.Slice((*unsafe.Pointer)(value.value_data), int(value.len))
	var goValue = make(map[AnimatablePropertyKey]objc.Object)
	for idx, k := range valueKeySlice {
		v := valueValueSlice[idx]
		goValue[AnimatablePropertyKey(foundation.MakeString(k).String())] = objc.MakeObject(v)
	}
	protocol.SetAnimations(goValue)
}

//export animatablePropertyContainer_Animations
func animatablePropertyContainer_Animations(hp C.uintptr_t) C.Dictionary {
	protocol := cgo.Handle(hp).Value().(AnimatablePropertyContainer)
	result := protocol.Animations()
	var cResult C.Dictionary
	if len(result) > 0 {
		cResultKeyData := make([]unsafe.Pointer, len(result))
		cResultValueData := make([]unsafe.Pointer, len(result))
		var idx = 0
		for k, v := range result {
			cResultKeyData[idx] = foundation.NewString(string(k)).Ptr()
			cResultValueData[idx] = objc.ExtractPtr(v)
			idx++
		}
		cResult.key_data = unsafe.Pointer(&cResultKeyData[0])
		cResult.value_data = unsafe.Pointer(&cResultValueData[0])
		cResult.len = C.int(len(result))
	}
	return cResult
}

//export AnimatablePropertyContainer_RespondsTo
func AnimatablePropertyContainer_RespondsTo(hp C.uintptr_t, selectorPtr unsafe.Pointer) bool {
	sel := objc.Selector(selectorPtr)
	selName := objc.Sel_GetName(sel)
	protocol := cgo.Handle(hp).Value().(AnimatablePropertyContainer)
	_ = protocol
	switch selName {
	case "animator":
		return true
	case "animationForKey:":
		return true
	case "setAnimations:":
		fallthrough
	case "animations":
		return true
	default:
		return false
	}
}
