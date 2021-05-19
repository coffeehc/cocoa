package appkit

// #include "user_interface_compression_options.h"
import "C"
import (
	"github.com/hsiafan/cocoa/foundation"
	"github.com/hsiafan/cocoa/objc"
	"unsafe"
)

type UserInterfaceCompressionOptions interface {
	objc.Object
	ContainsOptions(options UserInterfaceCompressionOptions) bool
	IntersectsOptions(options UserInterfaceCompressionOptions) bool
	OptionsByAddingOptions(options UserInterfaceCompressionOptions) UserInterfaceCompressionOptions
	OptionsByRemovingOptions(options UserInterfaceCompressionOptions) UserInterfaceCompressionOptions
	IsEmpty() bool
}

type NSUserInterfaceCompressionOptions struct {
	objc.NSObject
}

func MakeUserInterfaceCompressionOptions(ptr unsafe.Pointer) *NSUserInterfaceCompressionOptions {
	if ptr == nil {
		return nil
	}
	return &NSUserInterfaceCompressionOptions{
		NSObject: *objc.MakeObject(ptr),
	}
}

func AllocUserInterfaceCompressionOptions() *NSUserInterfaceCompressionOptions {
	return MakeUserInterfaceCompressionOptions(C.C_UserInterfaceCompressionOptions_Alloc())
}

func (n *NSUserInterfaceCompressionOptions) Init() UserInterfaceCompressionOptions {
	result_ := C.C_NSUserInterfaceCompressionOptions_Init(n.Ptr())
	return MakeUserInterfaceCompressionOptions(result_)
}

func (n *NSUserInterfaceCompressionOptions) InitWithCompressionOptions(options foundation.Set) UserInterfaceCompressionOptions {
	result_ := C.C_NSUserInterfaceCompressionOptions_InitWithCompressionOptions(n.Ptr(), objc.ExtractPtr(options))
	return MakeUserInterfaceCompressionOptions(result_)
}

func (n *NSUserInterfaceCompressionOptions) InitWithIdentifier(identifier string) UserInterfaceCompressionOptions {
	result_ := C.C_NSUserInterfaceCompressionOptions_InitWithIdentifier(n.Ptr(), foundation.NewString(identifier).Ptr())
	return MakeUserInterfaceCompressionOptions(result_)
}

func (n *NSUserInterfaceCompressionOptions) InitWithCoder(coder foundation.Coder) UserInterfaceCompressionOptions {
	result_ := C.C_NSUserInterfaceCompressionOptions_InitWithCoder(n.Ptr(), objc.ExtractPtr(coder))
	return MakeUserInterfaceCompressionOptions(result_)
}

func (n *NSUserInterfaceCompressionOptions) ContainsOptions(options UserInterfaceCompressionOptions) bool {
	result_ := C.C_NSUserInterfaceCompressionOptions_ContainsOptions(n.Ptr(), objc.ExtractPtr(options))
	return bool(result_)
}

func (n *NSUserInterfaceCompressionOptions) IntersectsOptions(options UserInterfaceCompressionOptions) bool {
	result_ := C.C_NSUserInterfaceCompressionOptions_IntersectsOptions(n.Ptr(), objc.ExtractPtr(options))
	return bool(result_)
}

func (n *NSUserInterfaceCompressionOptions) OptionsByAddingOptions(options UserInterfaceCompressionOptions) UserInterfaceCompressionOptions {
	result_ := C.C_NSUserInterfaceCompressionOptions_OptionsByAddingOptions(n.Ptr(), objc.ExtractPtr(options))
	return MakeUserInterfaceCompressionOptions(result_)
}

func (n *NSUserInterfaceCompressionOptions) OptionsByRemovingOptions(options UserInterfaceCompressionOptions) UserInterfaceCompressionOptions {
	result_ := C.C_NSUserInterfaceCompressionOptions_OptionsByRemovingOptions(n.Ptr(), objc.ExtractPtr(options))
	return MakeUserInterfaceCompressionOptions(result_)
}

func UserInterfaceCompressionOptions_HideImagesOption() UserInterfaceCompressionOptions {
	result_ := C.C_NSUserInterfaceCompressionOptions_UserInterfaceCompressionOptions_HideImagesOption()
	return MakeUserInterfaceCompressionOptions(result_)
}

func UserInterfaceCompressionOptions_HideTextOption() UserInterfaceCompressionOptions {
	result_ := C.C_NSUserInterfaceCompressionOptions_UserInterfaceCompressionOptions_HideTextOption()
	return MakeUserInterfaceCompressionOptions(result_)
}

func UserInterfaceCompressionOptions_ReduceMetricsOption() UserInterfaceCompressionOptions {
	result_ := C.C_NSUserInterfaceCompressionOptions_UserInterfaceCompressionOptions_ReduceMetricsOption()
	return MakeUserInterfaceCompressionOptions(result_)
}

func UserInterfaceCompressionOptions_BreakEqualWidthsOption() UserInterfaceCompressionOptions {
	result_ := C.C_NSUserInterfaceCompressionOptions_UserInterfaceCompressionOptions_BreakEqualWidthsOption()
	return MakeUserInterfaceCompressionOptions(result_)
}

func UserInterfaceCompressionOptions_StandardOptions() UserInterfaceCompressionOptions {
	result_ := C.C_NSUserInterfaceCompressionOptions_UserInterfaceCompressionOptions_StandardOptions()
	return MakeUserInterfaceCompressionOptions(result_)
}

func (n *NSUserInterfaceCompressionOptions) IsEmpty() bool {
	result_ := C.C_NSUserInterfaceCompressionOptions_IsEmpty(n.Ptr())
	return bool(result_)
}
