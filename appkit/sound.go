package appkit

// #cgo CFLAGS: -x objective-c
// #cgo LDFLAGS: -framework AppKit
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

func MakeSound(ptr unsafe.Pointer) *NSSound {
	if ptr == nil {
		return nil
	}
	return &NSSound{
		NSObject: *objc.MakeObject(ptr),
	}
}

func AllocSound() *NSSound {
	return MakeSound(C.C_Sound_Alloc())
}

func (n *NSSound) InitWithContentsOfFile_ByReference(path string, byRef bool) Sound {
	result := C.C_NSSound_InitWithContentsOfFile_ByReference(n.Ptr(), foundation.NewString(path).Ptr(), C.bool(byRef))
	return MakeSound(result)
}

func (n *NSSound) InitWithContentsOfURL_ByReference(url foundation.URL, byRef bool) Sound {
	result := C.C_NSSound_InitWithContentsOfURL_ByReference(n.Ptr(), objc.ExtractPtr(url), C.bool(byRef))
	return MakeSound(result)
}

func (n *NSSound) InitWithData(data []byte) Sound {
	result := C.C_NSSound_InitWithData(n.Ptr(), C.Array{data: unsafe.Pointer(&data[0]), len: C.int(len(data))})
	return MakeSound(result)
}

func (n *NSSound) InitWithPasteboard(pasteboard Pasteboard) Sound {
	result := C.C_NSSound_InitWithPasteboard(n.Ptr(), objc.ExtractPtr(pasteboard))
	return MakeSound(result)
}

func (n *NSSound) Init() Sound {
	result := C.C_NSSound_Init(n.Ptr())
	return MakeSound(result)
}

func SoundCanInitWithPasteboard(pasteboard Pasteboard) bool {
	result := C.C_NSSound_SoundCanInitWithPasteboard(objc.ExtractPtr(pasteboard))
	return bool(result)
}

func (n *NSSound) SetName(_string SoundName) bool {
	result := C.C_NSSound_SetName(n.Ptr(), foundation.NewString(string(_string)).Ptr())
	return bool(result)
}

func SoundNamed(name SoundName) Sound {
	result := C.C_NSSound_SoundNamed(foundation.NewString(string(name)).Ptr())
	return MakeSound(result)
}

func (n *NSSound) Pause() bool {
	result := C.C_NSSound_Pause(n.Ptr())
	return bool(result)
}

func (n *NSSound) Play() bool {
	result := C.C_NSSound_Play(n.Ptr())
	return bool(result)
}

func (n *NSSound) Resume() bool {
	result := C.C_NSSound_Resume(n.Ptr())
	return bool(result)
}

func (n *NSSound) Stop() bool {
	result := C.C_NSSound_Stop(n.Ptr())
	return bool(result)
}

func (n *NSSound) WriteToPasteboard(pasteboard Pasteboard) {
	C.C_NSSound_WriteToPasteboard(n.Ptr(), objc.ExtractPtr(pasteboard))
}

func (n *NSSound) Name() SoundName {
	result := C.C_NSSound_Name(n.Ptr())
	return SoundName(foundation.MakeString(result).String())
}

func (n *NSSound) Volume() float32 {
	result := C.C_NSSound_Volume(n.Ptr())
	return float32(result)
}

func (n *NSSound) SetVolume(value float32) {
	C.C_NSSound_SetVolume(n.Ptr(), C.float(value))
}

func (n *NSSound) CurrentTime() foundation.TimeInterval {
	result := C.C_NSSound_CurrentTime(n.Ptr())
	return foundation.TimeInterval(float64(result))
}

func (n *NSSound) SetCurrentTime(value foundation.TimeInterval) {
	C.C_NSSound_SetCurrentTime(n.Ptr(), C.double(float64(value)))
}

func (n *NSSound) Loops() bool {
	result := C.C_NSSound_Loops(n.Ptr())
	return bool(result)
}

func (n *NSSound) SetLoops(value bool) {
	C.C_NSSound_SetLoops(n.Ptr(), C.bool(value))
}

func (n *NSSound) PlaybackDeviceIdentifier() SoundPlaybackDeviceIdentifier {
	result := C.C_NSSound_PlaybackDeviceIdentifier(n.Ptr())
	return SoundPlaybackDeviceIdentifier(foundation.MakeString(result).String())
}

func (n *NSSound) SetPlaybackDeviceIdentifier(value SoundPlaybackDeviceIdentifier) {
	C.C_NSSound_SetPlaybackDeviceIdentifier(n.Ptr(), foundation.NewString(string(value)).Ptr())
}

func (n *NSSound) Duration() foundation.TimeInterval {
	result := C.C_NSSound_Duration(n.Ptr())
	return foundation.TimeInterval(float64(result))
}

func (n *NSSound) IsPlaying() bool {
	result := C.C_NSSound_IsPlaying(n.Ptr())
	return bool(result)
}
