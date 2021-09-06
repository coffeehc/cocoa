package foundation

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

func MakeAppleEventDescriptor(ptr unsafe.Pointer) NSAppleEventDescriptor {
	return NSAppleEventDescriptor{
		NSObject: objc.MakeObject(ptr),
	}
}

func AllocAppleEventDescriptor() NSAppleEventDescriptor {
	return MakeAppleEventDescriptor(C.C_AppleEventDescriptor_Alloc())
}

func (n NSAppleEventDescriptor) InitListDescriptor() AppleEventDescriptor {
	result_ := C.C_NSAppleEventDescriptor_InitListDescriptor(n.Ptr())
	return MakeAppleEventDescriptor(result_)
}

func (n NSAppleEventDescriptor) InitRecordDescriptor() AppleEventDescriptor {
	result_ := C.C_NSAppleEventDescriptor_InitRecordDescriptor(n.Ptr())
	return MakeAppleEventDescriptor(result_)
}

func (n NSAppleEventDescriptor) Init() AppleEventDescriptor {
	result_ := C.C_NSAppleEventDescriptor_Init(n.Ptr())
	return MakeAppleEventDescriptor(result_)
}

func AppleEventDescriptor_DescriptorWithString(_string string) AppleEventDescriptor {
	result_ := C.C_NSAppleEventDescriptor_AppleEventDescriptor_DescriptorWithString(NewString(_string).Ptr())
	return MakeAppleEventDescriptor(result_)
}

func AppleEventDescriptor_ListDescriptor() AppleEventDescriptor {
	result_ := C.C_NSAppleEventDescriptor_AppleEventDescriptor_ListDescriptor()
	return MakeAppleEventDescriptor(result_)
}

func AppleEventDescriptor_NullDescriptor() AppleEventDescriptor {
	result_ := C.C_NSAppleEventDescriptor_AppleEventDescriptor_NullDescriptor()
	return MakeAppleEventDescriptor(result_)
}

func AppleEventDescriptor_RecordDescriptor() AppleEventDescriptor {
	result_ := C.C_NSAppleEventDescriptor_AppleEventDescriptor_RecordDescriptor()
	return MakeAppleEventDescriptor(result_)
}

func (n NSAppleEventDescriptor) DescriptorAtIndex(index int) AppleEventDescriptor {
	result_ := C.C_NSAppleEventDescriptor_DescriptorAtIndex(n.Ptr(), C.int(index))
	return MakeAppleEventDescriptor(result_)
}

func (n NSAppleEventDescriptor) InsertDescriptor_AtIndex(descriptor AppleEventDescriptor, index int) {
	C.C_NSAppleEventDescriptor_InsertDescriptor_AtIndex(n.Ptr(), objc.ExtractPtr(descriptor), C.int(index))
}

func (n NSAppleEventDescriptor) RemoveDescriptorAtIndex(index int) {
	C.C_NSAppleEventDescriptor_RemoveDescriptorAtIndex(n.Ptr(), C.int(index))
}

func AppleEventDescriptor_DescriptorWithApplicationURL(applicationURL URL) AppleEventDescriptor {
	result_ := C.C_NSAppleEventDescriptor_AppleEventDescriptor_DescriptorWithApplicationURL(objc.ExtractPtr(applicationURL))
	return MakeAppleEventDescriptor(result_)
}

func AppleEventDescriptor_DescriptorWithBundleIdentifier(bundleIdentifier string) AppleEventDescriptor {
	result_ := C.C_NSAppleEventDescriptor_AppleEventDescriptor_DescriptorWithBundleIdentifier(NewString(bundleIdentifier).Ptr())
	return MakeAppleEventDescriptor(result_)
}

func AppleEventDescriptor_DescriptorWithDate(date Date) AppleEventDescriptor {
	result_ := C.C_NSAppleEventDescriptor_AppleEventDescriptor_DescriptorWithDate(objc.ExtractPtr(date))
	return MakeAppleEventDescriptor(result_)
}

func AppleEventDescriptor_DescriptorWithDouble(doubleValue float64) AppleEventDescriptor {
	result_ := C.C_NSAppleEventDescriptor_AppleEventDescriptor_DescriptorWithDouble(C.double(doubleValue))
	return MakeAppleEventDescriptor(result_)
}

func AppleEventDescriptor_DescriptorWithFileURL(fileURL URL) AppleEventDescriptor {
	result_ := C.C_NSAppleEventDescriptor_AppleEventDescriptor_DescriptorWithFileURL(objc.ExtractPtr(fileURL))
	return MakeAppleEventDescriptor(result_)
}

func AppleEventDescriptor_CurrentProcessDescriptor() AppleEventDescriptor {
	result_ := C.C_NSAppleEventDescriptor_AppleEventDescriptor_CurrentProcessDescriptor()
	return MakeAppleEventDescriptor(result_)
}

func (n NSAppleEventDescriptor) Data() []byte {
	result_ := C.C_NSAppleEventDescriptor_Data(n.Ptr())
	return MakeData(result_).ToBytes()
}

func (n NSAppleEventDescriptor) NumberOfItems() int {
	result_ := C.C_NSAppleEventDescriptor_NumberOfItems(n.Ptr())
	return int(result_)
}

func (n NSAppleEventDescriptor) StringValue() string {
	result_ := C.C_NSAppleEventDescriptor_StringValue(n.Ptr())
	return MakeString(result_).String()
}

func (n NSAppleEventDescriptor) DateValue() Date {
	result_ := C.C_NSAppleEventDescriptor_DateValue(n.Ptr())
	return MakeDate(result_)
}

func (n NSAppleEventDescriptor) DoubleValue() float64 {
	result_ := C.C_NSAppleEventDescriptor_DoubleValue(n.Ptr())
	return float64(result_)
}

func (n NSAppleEventDescriptor) FileURLValue() URL {
	result_ := C.C_NSAppleEventDescriptor_FileURLValue(n.Ptr())
	return MakeURL(result_)
}

func (n NSAppleEventDescriptor) IsRecordDescriptor() bool {
	result_ := C.C_NSAppleEventDescriptor_IsRecordDescriptor(n.Ptr())
	return bool(result_)
}
