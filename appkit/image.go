package appkit

// #cgo CFLAGS: -x objective-c
// #cgo LDFLAGS: -framework AppKit
// #include "image.h"
import "C"
import (
	"github.com/hsiafan/cocoa/foundation"
	"github.com/hsiafan/cocoa/objc"
	"unsafe"
)

type Image interface {
	objc.Object
	Size() foundation.Size
	SetSize(size foundation.Size)
	IsTemplate() bool
	SetTemplate(template bool)
	AccessibilityDescription() string
	SetAccessibilityDescription(accessibilityDescription string)
	MatchesOnlyOnBestFittingAxis() bool
	SetMatchesOnlyOnBestFittingAxis(matchesOnlyOnBestFittingAxis bool)
	Name() ImageName
	SetName(name ImageName)
	ImageWithSymbolConfiguration(configuration ImageSymbolConfiguration)
	InitByReferencingFile(fileName string) Image
	InitByReferencingURL(url foundation.URL) Image
	InitWithContentsOfFile(fileName string) Image
	InitWithContentsOfURL(fileName foundation.URL) Image
	InitWithData(data []byte) Image
	InitWithDataIgnoringOrientation(data []byte) Image
	InitWithSize(szie foundation.Size) Image
}

var _ Image = (*NSImage)(nil)

type NSImage struct {
	objc.NSObject
}

func MakeImage(ptr unsafe.Pointer) *NSImage {
	if ptr == nil {
		return nil
	}
	return &NSImage{
		NSObject: *objc.MakeObject(ptr),
	}
}

func (i *NSImage) Size() foundation.Size {
	return toSize(C.Image_Size(i.Ptr()))
}

func (i *NSImage) SetSize(size foundation.Size) {
	C.Image_SetSize(i.Ptr(), toNSSize(size))
}

func (i *NSImage) IsTemplate() bool {
	return bool(C.Image_IsTemplate(i.Ptr()))
}

func (i *NSImage) SetTemplate(template bool) {
	C.Image_SetTemplate(i.Ptr(), C.bool(template))
}

func (i *NSImage) AccessibilityDescription() string {
	return C.GoString(C.Image_AccessibilityDescription(i.Ptr()))
}

func (i *NSImage) SetAccessibilityDescription(accessibilityDescription string) {
	cAccessibilityDescription := C.CString(accessibilityDescription)
	defer C.free(unsafe.Pointer(cAccessibilityDescription))
	C.Image_SetAccessibilityDescription(i.Ptr(), cAccessibilityDescription)
}

func (i *NSImage) MatchesOnlyOnBestFittingAxis() bool {
	return bool(C.Image_MatchesOnlyOnBestFittingAxis(i.Ptr()))
}

func (i *NSImage) SetMatchesOnlyOnBestFittingAxis(matchesOnlyOnBestFittingAxis bool) {
	C.Image_SetMatchesOnlyOnBestFittingAxis(i.Ptr(), C.bool(matchesOnlyOnBestFittingAxis))
}

func ImageTypes() []string {
	var cArray C.Array = C.Image_ImageTypes()
	defer C.free(cArray.data)
	result := (*[1 << 28]unsafe.Pointer)(unsafe.Pointer(cArray.data))[:cArray.len:cArray.len]
	var goArray = make([]string, len(result))
	for idx, r := range result {
		goArray[idx] = C.GoString((*C.char)(r))
	}
	return goArray
}

func ImageUnfilteredTypes() []string {
	var cArray C.Array = C.Image_ImageUnfilteredTypes()
	defer C.free(cArray.data)
	result := (*[1 << 28]unsafe.Pointer)(unsafe.Pointer(cArray.data))[:cArray.len:cArray.len]
	var goArray = make([]string, len(result))
	for idx, r := range result {
		goArray[idx] = C.GoString((*C.char)(r))
	}
	return goArray
}

func ImageNamed(name ImageName) Image {
	cName := C.CString(string(name))
	defer C.free(unsafe.Pointer(cName))
	return MakeImage(C.Image_ImageNamed(cName))
}

func (i *NSImage) Name() ImageName {
	return ImageName(C.GoString(C.Image_Name(i.Ptr())))
}

func (i *NSImage) SetName(name ImageName) {
	cName := C.CString(string(name))
	defer C.free(unsafe.Pointer(cName))
	C.Image_SetName(i.Ptr(), cName)
}

func ImageWithSystemSymbolName(symbolName string, description string) Image {
	cSymbolName := C.CString(symbolName)
	defer C.free(unsafe.Pointer(cSymbolName))
	cDescription := C.CString(description)
	defer C.free(unsafe.Pointer(cDescription))
	return MakeImage(C.Image_ImageWithSystemSymbolName(cSymbolName, cDescription))
}

func (i *NSImage) ImageWithSymbolConfiguration(configuration ImageSymbolConfiguration) {
	C.Image_ImageWithSymbolConfiguration(i.Ptr(), toPointer(configuration))
}

func (i *NSImage) InitByReferencingFile(fileName string) Image {
	cFileName := C.CString(fileName)
	defer C.free(unsafe.Pointer(cFileName))
	return MakeImage(C.Image_InitByReferencingFile(i.Ptr(), cFileName))
}

func (i *NSImage) InitByReferencingURL(url foundation.URL) Image {
	return MakeImage(C.Image_InitByReferencingURL(i.Ptr(), toPointer(url)))
}

func (i *NSImage) InitWithContentsOfFile(fileName string) Image {
	cFileName := C.CString(fileName)
	defer C.free(unsafe.Pointer(cFileName))
	return MakeImage(C.Image_InitWithContentsOfFile(i.Ptr(), cFileName))
}

func (i *NSImage) InitWithContentsOfURL(fileName foundation.URL) Image {
	return MakeImage(C.Image_InitWithContentsOfURL(i.Ptr(), toPointer(fileName)))
}

func (i *NSImage) InitWithData(data []byte) Image {
	return MakeImage(C.Image_InitWithData(i.Ptr(), C.Array{data: unsafe.Pointer(&data[0]), len: C.int(len(data))}))
}

func (i *NSImage) InitWithDataIgnoringOrientation(data []byte) Image {
	return MakeImage(C.Image_InitWithDataIgnoringOrientation(i.Ptr(), C.Array{data: unsafe.Pointer(&data[0]), len: C.int(len(data))}))
}

func (i *NSImage) InitWithSize(szie foundation.Size) Image {
	return MakeImage(C.Image_InitWithSize(i.Ptr(), toNSSize(szie)))
}
