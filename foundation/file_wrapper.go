package foundation

// #cgo CFLAGS: -x objective-c
// #cgo LDFLAGS: -framework Foundation
// #include "file_wrapper.h"
import "C"
import (
	"github.com/hsiafan/cocoa/objc"
	"unsafe"
)

type FileWrapper interface {
	objc.Object
	AddFileWrapper(child FileWrapper) string
	RemoveFileWrapper(child FileWrapper)
	AddRegularFileWithContents_PreferredFilename(data []byte, fileName string) string
	KeyForFileWrapper(child FileWrapper) string
	MatchesContentsOfURL(url URL) bool
	IsRegularFile() bool
	IsDirectory() bool
	IsSymbolicLink() bool
	SymbolicLinkDestinationURL() URL
	SerializedRepresentation() []byte
	Filename() string
	SetFilename(value string)
	PreferredFilename() string
	SetPreferredFilename(value string)
	RegularFileContents() []byte
}

type NSFileWrapper struct {
	objc.NSObject
}

func MakeFileWrapper(ptr unsafe.Pointer) *NSFileWrapper {
	if ptr == nil {
		return nil
	}
	return &NSFileWrapper{
		NSObject: *objc.MakeObject(ptr),
	}
}

func AllocFileWrapper() *NSFileWrapper {
	return MakeFileWrapper(C.C_FileWrapper_Alloc())
}

func (n *NSFileWrapper) InitRegularFileWithContents(contents []byte) FileWrapper {
	result := C.C_NSFileWrapper_InitRegularFileWithContents(n.Ptr(), C.Array{data: unsafe.Pointer(&contents[0]), len: C.int(len(contents))})
	return MakeFileWrapper(result)
}

func (n *NSFileWrapper) InitSymbolicLinkWithDestinationURL(url URL) FileWrapper {
	result := C.C_NSFileWrapper_InitSymbolicLinkWithDestinationURL(n.Ptr(), objc.ExtractPtr(url))
	return MakeFileWrapper(result)
}

func (n *NSFileWrapper) InitWithSerializedRepresentation(serializeRepresentation []byte) FileWrapper {
	result := C.C_NSFileWrapper_InitWithSerializedRepresentation(n.Ptr(), C.Array{data: unsafe.Pointer(&serializeRepresentation[0]), len: C.int(len(serializeRepresentation))})
	return MakeFileWrapper(result)
}

func (n *NSFileWrapper) InitWithCoder(inCoder Coder) FileWrapper {
	result := C.C_NSFileWrapper_InitWithCoder(n.Ptr(), objc.ExtractPtr(inCoder))
	return MakeFileWrapper(result)
}

func (n *NSFileWrapper) Init() FileWrapper {
	result := C.C_NSFileWrapper_Init(n.Ptr())
	return MakeFileWrapper(result)
}

func (n *NSFileWrapper) AddFileWrapper(child FileWrapper) string {
	result := C.C_NSFileWrapper_AddFileWrapper(n.Ptr(), objc.ExtractPtr(child))
	return MakeString(result).String()
}

func (n *NSFileWrapper) RemoveFileWrapper(child FileWrapper) {
	C.C_NSFileWrapper_RemoveFileWrapper(n.Ptr(), objc.ExtractPtr(child))
}

func (n *NSFileWrapper) AddRegularFileWithContents_PreferredFilename(data []byte, fileName string) string {
	result := C.C_NSFileWrapper_AddRegularFileWithContents_PreferredFilename(n.Ptr(), C.Array{data: unsafe.Pointer(&data[0]), len: C.int(len(data))}, NewString(fileName).Ptr())
	return MakeString(result).String()
}

func (n *NSFileWrapper) KeyForFileWrapper(child FileWrapper) string {
	result := C.C_NSFileWrapper_KeyForFileWrapper(n.Ptr(), objc.ExtractPtr(child))
	return MakeString(result).String()
}

func (n *NSFileWrapper) MatchesContentsOfURL(url URL) bool {
	result := C.C_NSFileWrapper_MatchesContentsOfURL(n.Ptr(), objc.ExtractPtr(url))
	return bool(result)
}

func (n *NSFileWrapper) IsRegularFile() bool {
	result := C.C_NSFileWrapper_IsRegularFile(n.Ptr())
	return bool(result)
}

func (n *NSFileWrapper) IsDirectory() bool {
	result := C.C_NSFileWrapper_IsDirectory(n.Ptr())
	return bool(result)
}

func (n *NSFileWrapper) IsSymbolicLink() bool {
	result := C.C_NSFileWrapper_IsSymbolicLink(n.Ptr())
	return bool(result)
}

func (n *NSFileWrapper) SymbolicLinkDestinationURL() URL {
	result := C.C_NSFileWrapper_SymbolicLinkDestinationURL(n.Ptr())
	return MakeURL(result)
}

func (n *NSFileWrapper) SerializedRepresentation() []byte {
	result := C.C_NSFileWrapper_SerializedRepresentation(n.Ptr())
	resultBuffer := (*[1 << 30]byte)(result.data)[:C.int(result.len)]
	goResult := make([]byte, C.int(result.len))
	copy(goResult, resultBuffer)
	C.free(result.data)
	return goResult
}

func (n *NSFileWrapper) Filename() string {
	result := C.C_NSFileWrapper_Filename(n.Ptr())
	return MakeString(result).String()
}

func (n *NSFileWrapper) SetFilename(value string) {
	C.C_NSFileWrapper_SetFilename(n.Ptr(), NewString(value).Ptr())
}

func (n *NSFileWrapper) PreferredFilename() string {
	result := C.C_NSFileWrapper_PreferredFilename(n.Ptr())
	return MakeString(result).String()
}

func (n *NSFileWrapper) SetPreferredFilename(value string) {
	C.C_NSFileWrapper_SetPreferredFilename(n.Ptr(), NewString(value).Ptr())
}

func (n *NSFileWrapper) RegularFileContents() []byte {
	result := C.C_NSFileWrapper_RegularFileContents(n.Ptr())
	resultBuffer := (*[1 << 30]byte)(result.data)[:C.int(result.len)]
	goResult := make([]byte, C.int(result.len))
	copy(goResult, resultBuffer)
	C.free(result.data)
	return goResult
}
