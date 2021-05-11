package foundation

// #cgo CFLAGS: -x objective-c
// #cgo LDFLAGS: -framework Foundation
// #include "url.h"
import "C"
import (
	"github.com/hsiafan/cocoa/objc"
	"unsafe"
)

type URL interface {
	objc.Object
	IsFileReferenceURL() bool
	RemoveAllCachedResourceValues()
	RemoveCachedResourceValueForKey(key URLResourceKey)
	SetTemporaryResourceValue_ForKey(value objc.Object, key URLResourceKey)
	FileReferenceURL() URL
	URLByAppendingPathComponent(pathComponent string) URL
	URLByAppendingPathComponent_IsDirectory(pathComponent string, isDirectory bool) URL
	URLByAppendingPathExtension(pathExtension string) URL
	StartAccessingSecurityScopedResource() bool
	StopAccessingSecurityScopedResource()
	DataRepresentation() []byte
	IsFileURL() bool
	AbsoluteString() string
	AbsoluteURL() URL
	BaseURL() URL
	Fragment() string
	Host() string
	LastPathComponent() string
	Password() string
	Path() string
	PathExtension() string
	Port() Number
	Query() string
	RelativePath() string
	RelativeString() string
	ResourceSpecifier() string
	Scheme() string
	StandardizedURL() URL
	User() string
	FilePathURL() URL
	URLByDeletingLastPathComponent() URL
	URLByDeletingPathExtension() URL
	URLByResolvingSymlinksInPath() URL
	URLByStandardizingPath() URL
	HasDirectoryPath() bool
}

type NSURL struct {
	objc.NSObject
}

func MakeURL(ptr unsafe.Pointer) *NSURL {
	if ptr == nil {
		return nil
	}
	return &NSURL{
		NSObject: *objc.MakeObject(ptr),
	}
}

func AllocURL() *NSURL {
	return MakeURL(C.C_URL_Alloc())
}

func (n *NSURL) InitWithString(URLString string) URL {
	result := C.C_NSURL_InitWithString(n.Ptr(), NewString(URLString).Ptr())
	return MakeURL(result)
}

func (n *NSURL) InitWithString_RelativeToURL(URLString string, baseURL URL) URL {
	result := C.C_NSURL_InitWithString_RelativeToURL(n.Ptr(), NewString(URLString).Ptr(), objc.ExtractPtr(baseURL))
	return MakeURL(result)
}

func (n *NSURL) InitFileURLWithPath_IsDirectory(path string, isDir bool) URL {
	result := C.C_NSURL_InitFileURLWithPath_IsDirectory(n.Ptr(), NewString(path).Ptr(), C.bool(isDir))
	return MakeURL(result)
}

func (n *NSURL) InitFileURLWithPath_RelativeToURL(path string, baseURL URL) URL {
	result := C.C_NSURL_InitFileURLWithPath_RelativeToURL(n.Ptr(), NewString(path).Ptr(), objc.ExtractPtr(baseURL))
	return MakeURL(result)
}

func (n *NSURL) InitFileURLWithPath_IsDirectory_RelativeToURL(path string, isDir bool, baseURL URL) URL {
	result := C.C_NSURL_InitFileURLWithPath_IsDirectory_RelativeToURL(n.Ptr(), NewString(path).Ptr(), C.bool(isDir), objc.ExtractPtr(baseURL))
	return MakeURL(result)
}

func (n *NSURL) InitFileURLWithPath(path string) URL {
	result := C.C_NSURL_InitFileURLWithPath(n.Ptr(), NewString(path).Ptr())
	return MakeURL(result)
}

func (n *NSURL) InitAbsoluteURLWithDataRepresentation_RelativeToURL(data []byte, baseURL URL) URL {
	result := C.C_NSURL_InitAbsoluteURLWithDataRepresentation_RelativeToURL(n.Ptr(), C.Array{data: unsafe.Pointer(&data[0]), len: C.int(len(data))}, objc.ExtractPtr(baseURL))
	return MakeURL(result)
}

func (n *NSURL) InitWithDataRepresentation_RelativeToURL(data []byte, baseURL URL) URL {
	result := C.C_NSURL_InitWithDataRepresentation_RelativeToURL(n.Ptr(), C.Array{data: unsafe.Pointer(&data[0]), len: C.int(len(data))}, objc.ExtractPtr(baseURL))
	return MakeURL(result)
}

func (n *NSURL) Init() URL {
	result := C.C_NSURL_Init(n.Ptr())
	return MakeURL(result)
}

func URLWithString(URLString string) URL {
	result := C.C_NSURL_URLWithString(NewString(URLString).Ptr())
	return MakeURL(result)
}

func URLWithString_RelativeToURL(URLString string, baseURL URL) URL {
	result := C.C_NSURL_URLWithString_RelativeToURL(NewString(URLString).Ptr(), objc.ExtractPtr(baseURL))
	return MakeURL(result)
}

func URLFileURLWithPath_IsDirectory(path string, isDir bool) URL {
	result := C.C_NSURL_URLFileURLWithPath_IsDirectory(NewString(path).Ptr(), C.bool(isDir))
	return MakeURL(result)
}

func URLFileURLWithPath_RelativeToURL(path string, baseURL URL) URL {
	result := C.C_NSURL_URLFileURLWithPath_RelativeToURL(NewString(path).Ptr(), objc.ExtractPtr(baseURL))
	return MakeURL(result)
}

func URLFileURLWithPath_IsDirectory_RelativeToURL(path string, isDir bool, baseURL URL) URL {
	result := C.C_NSURL_URLFileURLWithPath_IsDirectory_RelativeToURL(NewString(path).Ptr(), C.bool(isDir), objc.ExtractPtr(baseURL))
	return MakeURL(result)
}

func URLFileURLWithPath(path string) URL {
	result := C.C_NSURL_URLFileURLWithPath(NewString(path).Ptr())
	return MakeURL(result)
}

func URLAbsoluteURLWithDataRepresentation_RelativeToURL(data []byte, baseURL URL) URL {
	result := C.C_NSURL_URLAbsoluteURLWithDataRepresentation_RelativeToURL(C.Array{data: unsafe.Pointer(&data[0]), len: C.int(len(data))}, objc.ExtractPtr(baseURL))
	return MakeURL(result)
}

func URLWithDataRepresentation_RelativeToURL(data []byte, baseURL URL) URL {
	result := C.C_NSURL_URLWithDataRepresentation_RelativeToURL(C.Array{data: unsafe.Pointer(&data[0]), len: C.int(len(data))}, objc.ExtractPtr(baseURL))
	return MakeURL(result)
}

func (n *NSURL) IsFileReferenceURL() bool {
	result := C.C_NSURL_IsFileReferenceURL(n.Ptr())
	return bool(result)
}

func (n *NSURL) RemoveAllCachedResourceValues() {
	C.C_NSURL_RemoveAllCachedResourceValues(n.Ptr())
}

func (n *NSURL) RemoveCachedResourceValueForKey(key URLResourceKey) {
	C.C_NSURL_RemoveCachedResourceValueForKey(n.Ptr(), NewString(string(key)).Ptr())
}

func (n *NSURL) SetTemporaryResourceValue_ForKey(value objc.Object, key URLResourceKey) {
	C.C_NSURL_SetTemporaryResourceValue_ForKey(n.Ptr(), objc.ExtractPtr(value), NewString(string(key)).Ptr())
}

func (n *NSURL) FileReferenceURL() URL {
	result := C.C_NSURL_FileReferenceURL(n.Ptr())
	return MakeURL(result)
}

func (n *NSURL) URLByAppendingPathComponent(pathComponent string) URL {
	result := C.C_NSURL_URLByAppendingPathComponent(n.Ptr(), NewString(pathComponent).Ptr())
	return MakeURL(result)
}

func (n *NSURL) URLByAppendingPathComponent_IsDirectory(pathComponent string, isDirectory bool) URL {
	result := C.C_NSURL_URLByAppendingPathComponent_IsDirectory(n.Ptr(), NewString(pathComponent).Ptr(), C.bool(isDirectory))
	return MakeURL(result)
}

func (n *NSURL) URLByAppendingPathExtension(pathExtension string) URL {
	result := C.C_NSURL_URLByAppendingPathExtension(n.Ptr(), NewString(pathExtension).Ptr())
	return MakeURL(result)
}

func (n *NSURL) StartAccessingSecurityScopedResource() bool {
	result := C.C_NSURL_StartAccessingSecurityScopedResource(n.Ptr())
	return bool(result)
}

func (n *NSURL) StopAccessingSecurityScopedResource() {
	C.C_NSURL_StopAccessingSecurityScopedResource(n.Ptr())
}

func (n *NSURL) DataRepresentation() []byte {
	result := C.C_NSURL_DataRepresentation(n.Ptr())
	resultBuffer := (*[1 << 30]byte)(result.data)[:C.int(result.len)]
	goResult := make([]byte, C.int(result.len))
	copy(goResult, resultBuffer)
	C.free(result.data)
	return goResult
}

func (n *NSURL) IsFileURL() bool {
	result := C.C_NSURL_IsFileURL(n.Ptr())
	return bool(result)
}

func (n *NSURL) AbsoluteString() string {
	result := C.C_NSURL_AbsoluteString(n.Ptr())
	return MakeString(result).String()
}

func (n *NSURL) AbsoluteURL() URL {
	result := C.C_NSURL_AbsoluteURL(n.Ptr())
	return MakeURL(result)
}

func (n *NSURL) BaseURL() URL {
	result := C.C_NSURL_BaseURL(n.Ptr())
	return MakeURL(result)
}

func (n *NSURL) Fragment() string {
	result := C.C_NSURL_Fragment(n.Ptr())
	return MakeString(result).String()
}

func (n *NSURL) Host() string {
	result := C.C_NSURL_Host(n.Ptr())
	return MakeString(result).String()
}

func (n *NSURL) LastPathComponent() string {
	result := C.C_NSURL_LastPathComponent(n.Ptr())
	return MakeString(result).String()
}

func (n *NSURL) Password() string {
	result := C.C_NSURL_Password(n.Ptr())
	return MakeString(result).String()
}

func (n *NSURL) Path() string {
	result := C.C_NSURL_Path(n.Ptr())
	return MakeString(result).String()
}

func (n *NSURL) PathExtension() string {
	result := C.C_NSURL_PathExtension(n.Ptr())
	return MakeString(result).String()
}

func (n *NSURL) Port() Number {
	result := C.C_NSURL_Port(n.Ptr())
	return MakeNumber(result)
}

func (n *NSURL) Query() string {
	result := C.C_NSURL_Query(n.Ptr())
	return MakeString(result).String()
}

func (n *NSURL) RelativePath() string {
	result := C.C_NSURL_RelativePath(n.Ptr())
	return MakeString(result).String()
}

func (n *NSURL) RelativeString() string {
	result := C.C_NSURL_RelativeString(n.Ptr())
	return MakeString(result).String()
}

func (n *NSURL) ResourceSpecifier() string {
	result := C.C_NSURL_ResourceSpecifier(n.Ptr())
	return MakeString(result).String()
}

func (n *NSURL) Scheme() string {
	result := C.C_NSURL_Scheme(n.Ptr())
	return MakeString(result).String()
}

func (n *NSURL) StandardizedURL() URL {
	result := C.C_NSURL_StandardizedURL(n.Ptr())
	return MakeURL(result)
}

func (n *NSURL) User() string {
	result := C.C_NSURL_User(n.Ptr())
	return MakeString(result).String()
}

func (n *NSURL) FilePathURL() URL {
	result := C.C_NSURL_FilePathURL(n.Ptr())
	return MakeURL(result)
}

func (n *NSURL) URLByDeletingLastPathComponent() URL {
	result := C.C_NSURL_URLByDeletingLastPathComponent(n.Ptr())
	return MakeURL(result)
}

func (n *NSURL) URLByDeletingPathExtension() URL {
	result := C.C_NSURL_URLByDeletingPathExtension(n.Ptr())
	return MakeURL(result)
}

func (n *NSURL) URLByResolvingSymlinksInPath() URL {
	result := C.C_NSURL_URLByResolvingSymlinksInPath(n.Ptr())
	return MakeURL(result)
}

func (n *NSURL) URLByStandardizingPath() URL {
	result := C.C_NSURL_URLByStandardizingPath(n.Ptr())
	return MakeURL(result)
}

func (n *NSURL) HasDirectoryPath() bool {
	result := C.C_NSURL_HasDirectoryPath(n.Ptr())
	return bool(result)
}
