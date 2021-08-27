package appkit

// #include "animation.h"
import "C"
import (
	"github.com/hsiafan/cocoa/foundation"
	"github.com/hsiafan/cocoa/objc"
	"unsafe"
)

type Animation interface {
	objc.Object
	StartAnimation()
	StopAnimation()
	AddProgressMark(progressMark AnimationProgress)
	RemoveProgressMark(progressMark AnimationProgress)
	StartWhenAnimation_ReachesProgress(animation Animation, startProgress AnimationProgress)
	StopWhenAnimation_ReachesProgress(animation Animation, stopProgress AnimationProgress)
	ClearStartAnimation()
	ClearStopAnimation()
	AnimationBlockingMode() AnimationBlockingMode
	SetAnimationBlockingMode(value AnimationBlockingMode)
	RunLoopModesForAnimating() []foundation.RunLoopMode
	AnimationCurve() AnimationCurve
	SetAnimationCurve(value AnimationCurve)
	Duration() foundation.TimeInterval
	SetDuration(value foundation.TimeInterval)
	FrameRate() float32
	SetFrameRate(value float32)
	Delegate() objc.Object
	SetDelegate(value objc.Object)
	IsAnimating() bool
	CurrentProgress() AnimationProgress
	SetCurrentProgress(value AnimationProgress)
	CurrentValue() float32
	ProgressMarks() []foundation.Number
	SetProgressMarks(value []foundation.Number)
}

type NSAnimation struct {
	objc.NSObject
}

func MakeAnimation(ptr unsafe.Pointer) NSAnimation {
	return NSAnimation{
		NSObject: objc.MakeObject(ptr),
	}
}

func AllocAnimation() NSAnimation {
	return MakeAnimation(C.C_Animation_Alloc())
}

func (n NSAnimation) InitWithDuration_AnimationCurve(duration foundation.TimeInterval, animationCurve AnimationCurve) Animation {
	result_ := C.C_NSAnimation_InitWithDuration_AnimationCurve(n.Ptr(), C.double(float64(duration)), C.uint(uint(animationCurve)))
	return MakeAnimation(result_)
}

func (n NSAnimation) InitWithCoder(coder foundation.Coder) Animation {
	result_ := C.C_NSAnimation_InitWithCoder(n.Ptr(), objc.ExtractPtr(coder))
	return MakeAnimation(result_)
}

func (n NSAnimation) Init() Animation {
	result_ := C.C_NSAnimation_Init(n.Ptr())
	return MakeAnimation(result_)
}

func (n NSAnimation) StartAnimation() {
	C.C_NSAnimation_StartAnimation(n.Ptr())
}

func (n NSAnimation) StopAnimation() {
	C.C_NSAnimation_StopAnimation(n.Ptr())
}

func (n NSAnimation) AddProgressMark(progressMark AnimationProgress) {
	C.C_NSAnimation_AddProgressMark(n.Ptr(), C.float(float32(progressMark)))
}

func (n NSAnimation) RemoveProgressMark(progressMark AnimationProgress) {
	C.C_NSAnimation_RemoveProgressMark(n.Ptr(), C.float(float32(progressMark)))
}

func (n NSAnimation) StartWhenAnimation_ReachesProgress(animation Animation, startProgress AnimationProgress) {
	C.C_NSAnimation_StartWhenAnimation_ReachesProgress(n.Ptr(), objc.ExtractPtr(animation), C.float(float32(startProgress)))
}

func (n NSAnimation) StopWhenAnimation_ReachesProgress(animation Animation, stopProgress AnimationProgress) {
	C.C_NSAnimation_StopWhenAnimation_ReachesProgress(n.Ptr(), objc.ExtractPtr(animation), C.float(float32(stopProgress)))
}

func (n NSAnimation) ClearStartAnimation() {
	C.C_NSAnimation_ClearStartAnimation(n.Ptr())
}

func (n NSAnimation) ClearStopAnimation() {
	C.C_NSAnimation_ClearStopAnimation(n.Ptr())
}

func (n NSAnimation) AnimationBlockingMode() AnimationBlockingMode {
	result_ := C.C_NSAnimation_AnimationBlockingMode(n.Ptr())
	return AnimationBlockingMode(uint(result_))
}

func (n NSAnimation) SetAnimationBlockingMode(value AnimationBlockingMode) {
	C.C_NSAnimation_SetAnimationBlockingMode(n.Ptr(), C.uint(uint(value)))
}

func (n NSAnimation) RunLoopModesForAnimating() []foundation.RunLoopMode {
	result_ := C.C_NSAnimation_RunLoopModesForAnimating(n.Ptr())
	if result_.len > 0 {
		defer C.free(result_.data)
	}
	result_Slice := unsafe.Slice((*unsafe.Pointer)(result_.data), int(result_.len))
	var goResult_ = make([]foundation.RunLoopMode, len(result_Slice))
	for idx, r := range result_Slice {
		goResult_[idx] = foundation.RunLoopMode(foundation.MakeString(r).String())
	}
	return goResult_
}

func (n NSAnimation) AnimationCurve() AnimationCurve {
	result_ := C.C_NSAnimation_AnimationCurve(n.Ptr())
	return AnimationCurve(uint(result_))
}

func (n NSAnimation) SetAnimationCurve(value AnimationCurve) {
	C.C_NSAnimation_SetAnimationCurve(n.Ptr(), C.uint(uint(value)))
}

func (n NSAnimation) Duration() foundation.TimeInterval {
	result_ := C.C_NSAnimation_Duration(n.Ptr())
	return foundation.TimeInterval(float64(result_))
}

func (n NSAnimation) SetDuration(value foundation.TimeInterval) {
	C.C_NSAnimation_SetDuration(n.Ptr(), C.double(float64(value)))
}

func (n NSAnimation) FrameRate() float32 {
	result_ := C.C_NSAnimation_FrameRate(n.Ptr())
	return float32(result_)
}

func (n NSAnimation) SetFrameRate(value float32) {
	C.C_NSAnimation_SetFrameRate(n.Ptr(), C.float(value))
}

func (n NSAnimation) Delegate() objc.Object {
	result_ := C.C_NSAnimation_Delegate(n.Ptr())
	return objc.MakeObject(result_)
}

func (n NSAnimation) SetDelegate(value objc.Object) {
	C.C_NSAnimation_SetDelegate(n.Ptr(), objc.ExtractPtr(value))
}

func (n NSAnimation) IsAnimating() bool {
	result_ := C.C_NSAnimation_IsAnimating(n.Ptr())
	return bool(result_)
}

func (n NSAnimation) CurrentProgress() AnimationProgress {
	result_ := C.C_NSAnimation_CurrentProgress(n.Ptr())
	return AnimationProgress(float32(result_))
}

func (n NSAnimation) SetCurrentProgress(value AnimationProgress) {
	C.C_NSAnimation_SetCurrentProgress(n.Ptr(), C.float(float32(value)))
}

func (n NSAnimation) CurrentValue() float32 {
	result_ := C.C_NSAnimation_CurrentValue(n.Ptr())
	return float32(result_)
}

func (n NSAnimation) ProgressMarks() []foundation.Number {
	result_ := C.C_NSAnimation_ProgressMarks(n.Ptr())
	if result_.len > 0 {
		defer C.free(result_.data)
	}
	result_Slice := unsafe.Slice((*unsafe.Pointer)(result_.data), int(result_.len))
	var goResult_ = make([]foundation.Number, len(result_Slice))
	for idx, r := range result_Slice {
		goResult_[idx] = foundation.MakeNumber(r)
	}
	return goResult_
}

func (n NSAnimation) SetProgressMarks(value []foundation.Number) {
	var cValue C.Array
	if len(value) > 0 {
		cValueData := make([]unsafe.Pointer, len(value))
		for idx, v := range value {
			cValueData[idx] = objc.ExtractPtr(v)
		}
		cValue.data = unsafe.Pointer(&cValueData[0])
		cValue.len = C.int(len(value))
	}
	C.C_NSAnimation_SetProgressMarks(n.Ptr(), cValue)
}
