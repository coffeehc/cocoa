package foundation

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
	PathComponents() []string
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

func MakeURL(ptr unsafe.Pointer) NSURL {
	return NSURL{
		NSObject: objc.MakeObject(ptr),
	}
}

func AllocURL() NSURL {
	return MakeURL(C.C_URL_Alloc())
}

func (n NSURL) InitWithString(URLString string) URL {
	result_ := C.C_NSURL_InitWithString(n.Ptr(), NewString(URLString).Ptr())
	return MakeURL(result_)
}

func (n NSURL) InitWithString_RelativeToURL(URLString string, baseURL URL) URL {
	result_ := C.C_NSURL_InitWithString_RelativeToURL(n.Ptr(), NewString(URLString).Ptr(), objc.ExtractPtr(baseURL))
	return MakeURL(result_)
}

func (n NSURL) InitFileURLWithPath_IsDirectory(path string, isDir bool) URL {
	result_ := C.C_NSURL_InitFileURLWithPath_IsDirectory(n.Ptr(), NewString(path).Ptr(), C.bool(isDir))
	return MakeURL(result_)
}

func (n NSURL) InitFileURLWithPath_RelativeToURL(path string, baseURL URL) URL {
	result_ := C.C_NSURL_InitFileURLWithPath_RelativeToURL(n.Ptr(), NewString(path).Ptr(), objc.ExtractPtr(baseURL))
	return MakeURL(result_)
}

func (n NSURL) InitFileURLWithPath_IsDirectory_RelativeToURL(path string, isDir bool, baseURL URL) URL {
	result_ := C.C_NSURL_InitFileURLWithPath_IsDirectory_RelativeToURL(n.Ptr(), NewString(path).Ptr(), C.bool(isDir), objc.ExtractPtr(baseURL))
	return MakeURL(result_)
}

func (n NSURL) InitFileURLWithPath(path string) URL {
	result_ := C.C_NSURL_InitFileURLWithPath(n.Ptr(), NewString(path).Ptr())
	return MakeURL(result_)
}

func (n NSURL) InitAbsoluteURLWithDataRepresentation_RelativeToURL(data []byte, baseURL URL) URL {
	result_ := C.C_NSURL_InitAbsoluteURLWithDataRepresentation_RelativeToURL(n.Ptr(), C.Array{data: unsafe.Pointer(&data[0]), len: C.int(len(data))}, objc.ExtractPtr(baseURL))
	return MakeURL(result_)
}

func (n NSURL) InitWithDataRepresentation_RelativeToURL(data []byte, baseURL URL) URL {
	result_ := C.C_NSURL_InitWithDataRepresentation_RelativeToURL(n.Ptr(), C.Array{data: unsafe.Pointer(&data[0]), len: C.int(len(data))}, objc.ExtractPtr(baseURL))
	return MakeURL(result_)
}

func (n NSURL) InitWithScheme_Host_Path(scheme string, host string, path string) URL {
	result_ := C.C_NSURL_InitWithScheme_Host_Path(n.Ptr(), NewString(scheme).Ptr(), NewString(host).Ptr(), NewString(path).Ptr())
	return MakeURL(result_)
}

func (n NSURL) Init() URL {
	result_ := C.C_NSURL_Init(n.Ptr())
	return MakeURL(result_)
}

func URLWithString(URLString string) URL {
	result_ := C.C_NSURL_URLWithString(NewString(URLString).Ptr())
	return MakeURL(result_)
}

func URLWithString_RelativeToURL(URLString string, baseURL URL) URL {
	result_ := C.C_NSURL_URLWithString_RelativeToURL(NewString(URLString).Ptr(), objc.ExtractPtr(baseURL))
	return MakeURL(result_)
}

func FileURLWithPath_IsDirectory(path string, isDir bool) URL {
	result_ := C.C_NSURL_FileURLWithPath_IsDirectory(NewString(path).Ptr(), C.bool(isDir))
	return MakeURL(result_)
}

func FileURLWithPath_RelativeToURL(path string, baseURL URL) URL {
	result_ := C.C_NSURL_FileURLWithPath_RelativeToURL(NewString(path).Ptr(), objc.ExtractPtr(baseURL))
	return MakeURL(result_)
}

func FileURLWithPath_IsDirectory_RelativeToURL(path string, isDir bool, baseURL URL) URL {
	result_ := C.C_NSURL_FileURLWithPath_IsDirectory_RelativeToURL(NewString(path).Ptr(), C.bool(isDir), objc.ExtractPtr(baseURL))
	return MakeURL(result_)
}

func FileURLWithPath(path string) URL {
	result_ := C.C_NSURL_FileURLWithPath(NewString(path).Ptr())
	return MakeURL(result_)
}

func FileURLWithPathComponents(components []string) URL {
	var cComponents C.Array
	if len(components) > 0 {
		cComponentsData := make([]unsafe.Pointer, len(components))
		for idx, v := range components {
			cComponentsData[idx] = NewString(v).Ptr()
		}
		cComponents.data = unsafe.Pointer(&cComponentsData[0])
		cComponents.len = C.int(len(components))
	}
	result_ := C.C_NSURL_FileURLWithPathComponents(cComponents)
	return MakeURL(result_)
}

func AbsoluteURLWithDataRepresentation_RelativeToURL(data []byte, baseURL URL) URL {
	result_ := C.C_NSURL_AbsoluteURLWithDataRepresentation_RelativeToURL(C.Array{data: unsafe.Pointer(&data[0]), len: C.int(len(data))}, objc.ExtractPtr(baseURL))
	return MakeURL(result_)
}

func URLWithDataRepresentation_RelativeToURL(data []byte, baseURL URL) URL {
	result_ := C.C_NSURL_URLWithDataRepresentation_RelativeToURL(C.Array{data: unsafe.Pointer(&data[0]), len: C.int(len(data))}, objc.ExtractPtr(baseURL))
	return MakeURL(result_)
}

func (n NSURL) IsFileReferenceURL() bool {
	result_ := C.C_NSURL_IsFileReferenceURL(n.Ptr())
	return bool(result_)
}

func (n NSURL) RemoveAllCachedResourceValues() {
	C.C_NSURL_RemoveAllCachedResourceValues(n.Ptr())
}

func (n NSURL) RemoveCachedResourceValueForKey(key URLResourceKey) {
	C.C_NSURL_RemoveCachedResourceValueForKey(n.Ptr(), NewString(string(key)).Ptr())
}

func (n NSURL) SetTemporaryResourceValue_ForKey(value objc.Object, key URLResourceKey) {
	C.C_NSURL_SetTemporaryResourceValue_ForKey(n.Ptr(), objc.ExtractPtr(value), NewString(string(key)).Ptr())
}

func (n NSURL) FileReferenceURL() URL {
	result_ := C.C_NSURL_FileReferenceURL(n.Ptr())
	return MakeURL(result_)
}

func (n NSURL) URLByAppendingPathComponent(pathComponent string) URL {
	result_ := C.C_NSURL_URLByAppendingPathComponent(n.Ptr(), NewString(pathComponent).Ptr())
	return MakeURL(result_)
}

func (n NSURL) URLByAppendingPathComponent_IsDirectory(pathComponent string, isDirectory bool) URL {
	result_ := C.C_NSURL_URLByAppendingPathComponent_IsDirectory(n.Ptr(), NewString(pathComponent).Ptr(), C.bool(isDirectory))
	return MakeURL(result_)
}

func (n NSURL) URLByAppendingPathExtension(pathExtension string) URL {
	result_ := C.C_NSURL_URLByAppendingPathExtension(n.Ptr(), NewString(pathExtension).Ptr())
	return MakeURL(result_)
}

func URL_ResourceValuesForKeys_FromBookmarkData(keys []URLResourceKey, bookmarkData []byte) map[URLResourceKey]objc.Object {
	var cKeys C.Array
	if len(keys) > 0 {
		cKeysData := make([]unsafe.Pointer, len(keys))
		for idx, v := range keys {
			cKeysData[idx] = NewString(string(v)).Ptr()
		}
		cKeys.data = unsafe.Pointer(&cKeysData[0])
		cKeys.len = C.int(len(keys))
	}
	result_ := C.C_NSURL_URL_ResourceValuesForKeys_FromBookmarkData(cKeys, C.Array{data: unsafe.Pointer(&bookmarkData[0]), len: C.int(len(bookmarkData))})
	if result_.len > 0 {
		defer C.free(result_.key_data)
		defer C.free(result_.value_data)
	}
	result_KeySlice := unsafe.Slice((*unsafe.Pointer)(result_.key_data), int(result_.len))
	result_ValueSlice := unsafe.Slice((*unsafe.Pointer)(result_.value_data), int(result_.len))
	var goResult_ = make(map[URLResourceKey]objc.Object)
	for idx, k := range result_KeySlice {
		v := result_ValueSlice[idx]
		goResult_[URLResourceKey(MakeString(k).String())] = objc.MakeObject(v)
	}
	return goResult_
}

func (n NSURL) StartAccessingSecurityScopedResource() bool {
	result_ := C.C_NSURL_StartAccessingSecurityScopedResource(n.Ptr())
	return bool(result_)
}

func (n NSURL) StopAccessingSecurityScopedResource() {
	C.C_NSURL_StopAccessingSecurityScopedResource(n.Ptr())
}

func (n NSURL) DataRepresentation() []byte {
	result_ := C.C_NSURL_DataRepresentation(n.Ptr())
	var goResult_ []byte
	if result_.len > 0 {
		result_Buffer := unsafe.Slice((*byte)(result_.data), int(result_.len))
		goResult_ = make([]byte, C.int(result_.len))
		copy(goResult_, result_Buffer)
	}
	return goResult_
}

func (n NSURL) IsFileURL() bool {
	result_ := C.C_NSURL_IsFileURL(n.Ptr())
	return bool(result_)
}

func (n NSURL) AbsoluteString() string {
	result_ := C.C_NSURL_AbsoluteString(n.Ptr())
	return MakeString(result_).String()
}

func (n NSURL) AbsoluteURL() URL {
	result_ := C.C_NSURL_AbsoluteURL(n.Ptr())
	return MakeURL(result_)
}

func (n NSURL) BaseURL() URL {
	result_ := C.C_NSURL_BaseURL(n.Ptr())
	return MakeURL(result_)
}

func (n NSURL) Fragment() string {
	result_ := C.C_NSURL_Fragment(n.Ptr())
	return MakeString(result_).String()
}

func (n NSURL) Host() string {
	result_ := C.C_NSURL_Host(n.Ptr())
	return MakeString(result_).String()
}

func (n NSURL) LastPathComponent() string {
	result_ := C.C_NSURL_LastPathComponent(n.Ptr())
	return MakeString(result_).String()
}

func (n NSURL) Password() string {
	result_ := C.C_NSURL_Password(n.Ptr())
	return MakeString(result_).String()
}

func (n NSURL) Path() string {
	result_ := C.C_NSURL_Path(n.Ptr())
	return MakeString(result_).String()
}

func (n NSURL) PathComponents() []string {
	result_ := C.C_NSURL_PathComponents(n.Ptr())
	if result_.len > 0 {
		defer C.free(result_.data)
	}
	result_Slice := unsafe.Slice((*unsafe.Pointer)(result_.data), int(result_.len))
	var goResult_ = make([]string, len(result_Slice))
	for idx, r := range result_Slice {
		goResult_[idx] = MakeString(r).String()
	}
	return goResult_
}

func (n NSURL) PathExtension() string {
	result_ := C.C_NSURL_PathExtension(n.Ptr())
	return MakeString(result_).String()
}

func (n NSURL) Port() Number {
	result_ := C.C_NSURL_Port(n.Ptr())
	return MakeNumber(result_)
}

func (n NSURL) Query() string {
	result_ := C.C_NSURL_Query(n.Ptr())
	return MakeString(result_).String()
}

func (n NSURL) RelativePath() string {
	result_ := C.C_NSURL_RelativePath(n.Ptr())
	return MakeString(result_).String()
}

func (n NSURL) RelativeString() string {
	result_ := C.C_NSURL_RelativeString(n.Ptr())
	return MakeString(result_).String()
}

func (n NSURL) ResourceSpecifier() string {
	result_ := C.C_NSURL_ResourceSpecifier(n.Ptr())
	return MakeString(result_).String()
}

func (n NSURL) Scheme() string {
	result_ := C.C_NSURL_Scheme(n.Ptr())
	return MakeString(result_).String()
}

func (n NSURL) StandardizedURL() URL {
	result_ := C.C_NSURL_StandardizedURL(n.Ptr())
	return MakeURL(result_)
}

func (n NSURL) User() string {
	result_ := C.C_NSURL_User(n.Ptr())
	return MakeString(result_).String()
}

func (n NSURL) FilePathURL() URL {
	result_ := C.C_NSURL_FilePathURL(n.Ptr())
	return MakeURL(result_)
}

func (n NSURL) URLByDeletingLastPathComponent() URL {
	result_ := C.C_NSURL_URLByDeletingLastPathComponent(n.Ptr())
	return MakeURL(result_)
}

func (n NSURL) URLByDeletingPathExtension() URL {
	result_ := C.C_NSURL_URLByDeletingPathExtension(n.Ptr())
	return MakeURL(result_)
}

func (n NSURL) URLByResolvingSymlinksInPath() URL {
	result_ := C.C_NSURL_URLByResolvingSymlinksInPath(n.Ptr())
	return MakeURL(result_)
}

func (n NSURL) URLByStandardizingPath() URL {
	result_ := C.C_NSURL_URLByStandardizingPath(n.Ptr())
	return MakeURL(result_)
}

func (n NSURL) HasDirectoryPath() bool {
	result_ := C.C_NSURL_HasDirectoryPath(n.Ptr())
	return bool(result_)
}
