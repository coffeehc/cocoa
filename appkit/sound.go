package appkit

// #include "sound.h"
import "C"
import (
	"github.com/hsiafan/cocoa/foundation"
	"github.com/hsiafan/cocoa/objc"
	"unsafe"
)

type Sound interface {
	objc.Object
	SetName(_string SoundName) bool
	Pause() bool
	Play() bool
	Resume() bool
	Stop() bool
	WriteToPasteboard(pasteboard Pasteboard)
	Delegate() objc.Object
	SetDelegate(value objc.Object)
	Name() SoundName
	Volume() float32
	SetVolume(value float32)
	CurrentTime() foundation.TimeInterval
	SetCurrentTime(value foundation.TimeInterval)
	Loops() bool
	SetLoops(value bool)
	PlaybackDeviceIdentifier() SoundPlaybackDeviceIdentifier
	SetPlaybackDeviceIdentifier(value SoundPlaybackDeviceIdentifier)
	Duration() foundation.TimeInterval
	IsPlaying() bool
}

type NSSound struct {
	objc.NSObject
}

func MakeSound(ptr unsafe.Pointer) NSSound {
	return NSSound{
		NSObject: objc.MakeObject(ptr),
	}
}

func (n NSSound) InitWithContentsOfFile_ByReference(path string, byRef bool) NSSound {
	result_ := C.C_NSSound_InitWithContentsOfFile_ByReference(n.Ptr(), foundation.NewString(path).Ptr(), C.bool(byRef))
	return MakeSound(result_)
}

func (n NSSound) InitWithContentsOfURL_ByReference(url foundation.URL, byRef bool) NSSound {
	result_ := C.C_NSSound_InitWithContentsOfURL_ByReference(n.Ptr(), objc.ExtractPtr(url), C.bool(byRef))
	return MakeSound(result_)
}

func (n NSSound) InitWithData(data []byte) NSSound {
	result_ := C.C_NSSound_InitWithData(n.Ptr(), foundation.NewData(data).Ptr())
	return MakeSound(result_)
}

func (n NSSound) InitWithPasteboard(pasteboard Pasteboard) NSSound {
	result_ := C.C_NSSound_InitWithPasteboard(n.Ptr(), objc.ExtractPtr(pasteboard))
	return MakeSound(result_)
}

func AllocSound() NSSound {
	result_ := C.C_NSSound_AllocSound()
	return MakeSound(result_)
}

func (n NSSound) Init() NSSound {
	result_ := C.C_NSSound_Init(n.Ptr())
	return MakeSound(result_)
}

func NewSound() NSSound {
	result_ := C.C_NSSound_NewSound()
	return MakeSound(result_)
}

func (n NSSound) Autorelease() NSSound {
	result_ := C.C_NSSound_Autorelease(n.Ptr())
	return MakeSound(result_)
}

func (n NSSound) Retain() NSSound {
	result_ := C.C_NSSound_Retain(n.Ptr())
	return MakeSound(result_)
}

func Sound_CanInitWithPasteboard(pasteboard Pasteboard) bool {
	result_ := C.C_NSSound_Sound_CanInitWithPasteboard(objc.ExtractPtr(pasteboard))
	return bool(result_)
}

func (n NSSound) SetName(_string SoundName) bool {
	result_ := C.C_NSSound_SetName(n.Ptr(), foundation.NewString(string(_string)).Ptr())
	return bool(result_)
}

func SoundNamed(name SoundName) Sound {
	result_ := C.C_NSSound_SoundNamed(foundation.NewString(string(name)).Ptr())
	return MakeSound(result_)
}

func (n NSSound) Pause() bool {
	result_ := C.C_NSSound_Pause(n.Ptr())
	return bool(result_)
}

func (n NSSound) Play() bool {
	result_ := C.C_NSSound_Play(n.Ptr())
	return bool(result_)
}

func (n NSSound) Resume() bool {
	result_ := C.C_NSSound_Resume(n.Ptr())
	return bool(result_)
}

func (n NSSound) Stop() bool {
	result_ := C.C_NSSound_Stop(n.Ptr())
	return bool(result_)
}

func (n NSSound) WriteToPasteboard(pasteboard Pasteboard) {
	C.C_NSSound_WriteToPasteboard(n.Ptr(), objc.ExtractPtr(pasteboard))
}

func (n NSSound) Delegate() objc.Object {
	result_ := C.C_NSSound_Delegate(n.Ptr())
	return objc.MakeObject(result_)
}

func (n NSSound) SetDelegate(value objc.Object) {
	C.C_NSSound_SetDelegate(n.Ptr(), objc.ExtractPtr(value))
}

func (n NSSound) Name() SoundName {
	result_ := C.C_NSSound_Name(n.Ptr())
	return SoundName(foundation.MakeString(result_).String())
}

func (n NSSound) Volume() float32 {
	result_ := C.C_NSSound_Volume(n.Ptr())
	return float32(result_)
}

func (n NSSound) SetVolume(value float32) {
	C.C_NSSound_SetVolume(n.Ptr(), C.float(value))
}

func (n NSSound) CurrentTime() foundation.TimeInterval {
	result_ := C.C_NSSound_CurrentTime(n.Ptr())
	return foundation.TimeInterval(float64(result_))
}

func (n NSSound) SetCurrentTime(value foundation.TimeInterval) {
	C.C_NSSound_SetCurrentTime(n.Ptr(), C.double(float64(value)))
}

func (n NSSound) Loops() bool {
	result_ := C.C_NSSound_Loops(n.Ptr())
	return bool(result_)
}

func (n NSSound) SetLoops(value bool) {
	C.C_NSSound_SetLoops(n.Ptr(), C.bool(value))
}

func (n NSSound) PlaybackDeviceIdentifier() SoundPlaybackDeviceIdentifier {
	result_ := C.C_NSSound_PlaybackDeviceIdentifier(n.Ptr())
	return SoundPlaybackDeviceIdentifier(foundation.MakeString(result_).String())
}

func (n NSSound) SetPlaybackDeviceIdentifier(value SoundPlaybackDeviceIdentifier) {
	C.C_NSSound_SetPlaybackDeviceIdentifier(n.Ptr(), foundation.NewString(string(value)).Ptr())
}

func SoundUnfilteredTypes() []string {
	result_ := C.C_NSSound_SoundUnfilteredTypes()
	if result_.len > 0 {
		defer C.free(result_.data)
	}
	result_Slice := unsafe.Slice((*unsafe.Pointer)(result_.data), int(result_.len))
	var goResult_ = make([]string, len(result_Slice))
	for idx, r := range result_Slice {
		goResult_[idx] = foundation.MakeString(r).String()
	}
	return goResult_
}

func (n NSSound) Duration() foundation.TimeInterval {
	result_ := C.C_NSSound_Duration(n.Ptr())
	return foundation.TimeInterval(float64(result_))
}

func (n NSSound) IsPlaying() bool {
	result_ := C.C_NSSound_IsPlaying(n.Ptr())
	return bool(result_)
}
