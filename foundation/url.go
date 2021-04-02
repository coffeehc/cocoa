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
	IsEqual(anObject objc.Object) bool
	IsFileReferenceURL() bool
}

var _ URL = (*NSURL)(nil)

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

func (u *NSURL) IsFileURL() bool {
	return bool(C.URL_IsFileURL(u.Ptr()))
}

func (u *NSURL) AbsoluteString() string {
	return C.GoString(C.URL_AbsoluteString(u.Ptr()))
}

func (u *NSURL) AbsoluteURL() URL {
	return MakeURL(C.URL_AbsoluteURL(u.Ptr()))
}

func (u *NSURL) BaseURL() URL {
	return MakeURL(C.URL_BaseURL(u.Ptr()))
}

func (u *NSURL) Fragment() string {
	return C.GoString(C.URL_Fragment(u.Ptr()))
}

func (u *NSURL) Host() string {
	return C.GoString(C.URL_Host(u.Ptr()))
}

func (u *NSURL) LastPathComponent() string {
	return C.GoString(C.URL_LastPathComponent(u.Ptr()))
}

func (u *NSURL) Password() string {
	return C.GoString(C.URL_Password(u.Ptr()))
}

func (u *NSURL) Path() string {
	return C.GoString(C.URL_Path(u.Ptr()))
}

func (u *NSURL) PathComponents() []string {
	var cArray C.Array = C.URL_PathComponents(u.Ptr())
	defer C.free(cArray.data)
	result := (*[1 << 28]unsafe.Pointer)(unsafe.Pointer(cArray.data))[:cArray.len:cArray.len]
	var goArray = make([]string, len(result))
	for idx, r := range result {
		goArray[idx] = C.GoString((*C.char)(r))
	}
	return goArray
}

func (u *NSURL) PathExtension() string {
	return C.GoString(C.URL_PathExtension(u.Ptr()))
}

func (u *NSURL) Port() Number {
	return MakeNumber(C.URL_Port(u.Ptr()))
}

func (u *NSURL) Query() string {
	return C.GoString(C.URL_Query(u.Ptr()))
}

func (u *NSURL) RelativePath() string {
	return C.GoString(C.URL_RelativePath(u.Ptr()))
}

func (u *NSURL) RelativeString() string {
	return C.GoString(C.URL_RelativeString(u.Ptr()))
}

func (u *NSURL) ResourceSpecifier() string {
	return C.GoString(C.URL_ResourceSpecifier(u.Ptr()))
}

func (u *NSURL) Scheme() string {
	return C.GoString(C.URL_Scheme(u.Ptr()))
}

func (u *NSURL) StandardizedURL() URL {
	return MakeURL(C.URL_StandardizedURL(u.Ptr()))
}

func (u *NSURL) User() string {
	return C.GoString(C.URL_User(u.Ptr()))
}

func (u *NSURL) FilePathURL() URL {
	return MakeURL(C.URL_FilePathURL(u.Ptr()))
}

func URLWithString(URLString string) URL {
	cURLString := C.CString(URLString)
	defer C.free(unsafe.Pointer(cURLString))
	return MakeURL(C.URL_URLWithString(cURLString))
}

func URLWithStringRelativeToURL(URLString string, baseURL URL) URL {
	cURLString := C.CString(URLString)
	defer C.free(unsafe.Pointer(cURLString))
	return MakeURL(C.URL_URLWithStringRelativeToURL(cURLString, toPointer(baseURL)))
}

func FileURLWithPath(path string) URL {
	cPath := C.CString(path)
	defer C.free(unsafe.Pointer(cPath))
	return MakeURL(C.URL_FileURLWithPath(cPath))
}

func FileURLWithPath2(path string, isDir bool) URL {
	cPath := C.CString(path)
	defer C.free(unsafe.Pointer(cPath))
	return MakeURL(C.URL_FileURLWithPath2(cPath, C.bool(isDir)))
}

func FileURLWithPathRelativeToURL(path string, baseURL URL) URL {
	cPath := C.CString(path)
	defer C.free(unsafe.Pointer(cPath))
	return MakeURL(C.URL_FileURLWithPathRelativeToURL(cPath, toPointer(baseURL)))
}

func FileURLWithPathRelativeToURL2(path string, isDir bool, baseURL URL) URL {
	cPath := C.CString(path)
	defer C.free(unsafe.Pointer(cPath))
	return MakeURL(C.URL_FileURLWithPathRelativeToURL2(cPath, C.bool(isDir), toPointer(baseURL)))
}

func FileURLWithPathComponents(components []string) URL {
	c_componentsData := make([]unsafe.Pointer, len(components))
	for idx, v := range components {
		c_componentsData[idx] = unsafe.Pointer(C.CString(v))
	}
	c_components := C.Array{data: unsafe.Pointer(&c_componentsData[0]), len: C.int(len(components))}
	defer func() {
		for _, p := range c_componentsData {
			C.free(p)
		}
	}()
	return MakeURL(C.URL_FileURLWithPathComponents(c_components))
}

func (u *NSURL) IsEqual(anObject objc.Object) bool {
	return bool(C.URL_IsEqual(u.Ptr(), toPointer(anObject)))
}

func (u *NSURL) IsFileReferenceURL() bool {
	return bool(C.URL_IsFileReferenceURL(u.Ptr()))
}
