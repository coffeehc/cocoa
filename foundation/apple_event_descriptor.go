package foundation

// #cgo CFLAGS: -x objective-c
// #cgo LDFLAGS: -framework Foundation
// #include "apple_event_descriptor.h"
import "C"
import (
	"github.com/hsiafan/cocoa/objc"
	"unsafe"
)

type AppleEventDescriptor interface {
	objc.Object
	DescriptorAtIndex(index int) AppleEventDescriptor
	InsertDescriptor_AtIndex(descriptor AppleEventDescriptor, index int)
	RemoveDescriptorAtIndex(index int)
	Data() []byte
	NumberOfItems() int
	StringValue() string
	DateValue() Date
	DoubleValue() float64
	FileURLValue() URL
	IsRecordDescriptor() bool
}

type NSAppleEventDescriptor struct {
	objc.NSObject
}

func MakeAppleEventDescriptor(ptr unsafe.Pointer) *NSAppleEventDescriptor {
	if ptr == nil {
		return nil
	}
	return &NSAppleEventDescriptor{
		NSObject: *objc.MakeObject(ptr),
	}
}

func AllocAppleEventDescriptor() *NSAppleEventDescriptor {
	return MakeAppleEventDescriptor(C.C_AppleEventDescriptor_Alloc())
}

func (n *NSAppleEventDescriptor) InitListDescriptor() AppleEventDescriptor {
	result := C.C_NSAppleEventDescriptor_InitListDescriptor(n.Ptr())
	return MakeAppleEventDescriptor(result)
}

func (n *NSAppleEventDescriptor) InitRecordDescriptor() AppleEventDescriptor {
	result := C.C_NSAppleEventDescriptor_InitRecordDescriptor(n.Ptr())
	return MakeAppleEventDescriptor(result)
}

func (n *NSAppleEventDescriptor) Init() AppleEventDescriptor {
	result := C.C_NSAppleEventDescriptor_Init(n.Ptr())
	return MakeAppleEventDescriptor(result)
}

func AppleEventDescriptorDescriptorWithString(_string string) AppleEventDescriptor {
	result := C.C_NSAppleEventDescriptor_AppleEventDescriptorDescriptorWithString(NewString(_string).Ptr())
	return MakeAppleEventDescriptor(result)
}

func AppleEventDescriptorListDescriptor() AppleEventDescriptor {
	result := C.C_NSAppleEventDescriptor_AppleEventDescriptorListDescriptor()
	return MakeAppleEventDescriptor(result)
}

func AppleEventDescriptorNullDescriptor() AppleEventDescriptor {
	result := C.C_NSAppleEventDescriptor_AppleEventDescriptorNullDescriptor()
	return MakeAppleEventDescriptor(result)
}

func AppleEventDescriptorRecordDescriptor() AppleEventDescriptor {
	result := C.C_NSAppleEventDescriptor_AppleEventDescriptorRecordDescriptor()
	return MakeAppleEventDescriptor(result)
}

func (n *NSAppleEventDescriptor) DescriptorAtIndex(index int) AppleEventDescriptor {
	result := C.C_NSAppleEventDescriptor_DescriptorAtIndex(n.Ptr(), C.int(index))
	return MakeAppleEventDescriptor(result)
}

func (n *NSAppleEventDescriptor) InsertDescriptor_AtIndex(descriptor AppleEventDescriptor, index int) {
	C.C_NSAppleEventDescriptor_InsertDescriptor_AtIndex(n.Ptr(), objc.ExtractPtr(descriptor), C.int(index))
}

func (n *NSAppleEventDescriptor) RemoveDescriptorAtIndex(index int) {
	C.C_NSAppleEventDescriptor_RemoveDescriptorAtIndex(n.Ptr(), C.int(index))
}

func AppleEventDescriptorDescriptorWithApplicationURL(applicationURL URL) AppleEventDescriptor {
	result := C.C_NSAppleEventDescriptor_AppleEventDescriptorDescriptorWithApplicationURL(objc.ExtractPtr(applicationURL))
	return MakeAppleEventDescriptor(result)
}

func AppleEventDescriptorDescriptorWithBundleIdentifier(bundleIdentifier string) AppleEventDescriptor {
	result := C.C_NSAppleEventDescriptor_AppleEventDescriptorDescriptorWithBundleIdentifier(NewString(bundleIdentifier).Ptr())
	return MakeAppleEventDescriptor(result)
}

func AppleEventDescriptorDescriptorWithDate(date Date) AppleEventDescriptor {
	result := C.C_NSAppleEventDescriptor_AppleEventDescriptorDescriptorWithDate(objc.ExtractPtr(date))
	return MakeAppleEventDescriptor(result)
}

func AppleEventDescriptorDescriptorWithDouble(doubleValue float64) AppleEventDescriptor {
	result := C.C_NSAppleEventDescriptor_AppleEventDescriptorDescriptorWithDouble(C.double(doubleValue))
	return MakeAppleEventDescriptor(result)
}

func AppleEventDescriptorDescriptorWithFileURL(fileURL URL) AppleEventDescriptor {
	result := C.C_NSAppleEventDescriptor_AppleEventDescriptorDescriptorWithFileURL(objc.ExtractPtr(fileURL))
	return MakeAppleEventDescriptor(result)
}

func AppleEventDescriptorCurrentProcessDescriptor() AppleEventDescriptor {
	result := C.C_NSAppleEventDescriptor_AppleEventDescriptorCurrentProcessDescriptor()
	return MakeAppleEventDescriptor(result)
}

func (n *NSAppleEventDescriptor) Data() []byte {
	result := C.C_NSAppleEventDescriptor_Data(n.Ptr())
	resultBuffer := (*[1 << 30]byte)(result.data)[:C.int(result.len)]
	goResult := make([]byte, C.int(result.len))
	copy(goResult, resultBuffer)
	C.free(result.data)
	return goResult
}

func (n *NSAppleEventDescriptor) NumberOfItems() int {
	result := C.C_NSAppleEventDescriptor_NumberOfItems(n.Ptr())
	return int(result)
}

func (n *NSAppleEventDescriptor) StringValue() string {
	result := C.C_NSAppleEventDescriptor_StringValue(n.Ptr())
	return MakeString(result).String()
}

func (n *NSAppleEventDescriptor) DateValue() Date {
	result := C.C_NSAppleEventDescriptor_DateValue(n.Ptr())
	return MakeDate(result)
}

func (n *NSAppleEventDescriptor) DoubleValue() float64 {
	result := C.C_NSAppleEventDescriptor_DoubleValue(n.Ptr())
	return float64(result)
}

func (n *NSAppleEventDescriptor) FileURLValue() URL {
	result := C.C_NSAppleEventDescriptor_FileURLValue(n.Ptr())
	return MakeURL(result)
}

func (n *NSAppleEventDescriptor) IsRecordDescriptor() bool {
	result := C.C_NSAppleEventDescriptor_IsRecordDescriptor(n.Ptr())
	return bool(result)
}
