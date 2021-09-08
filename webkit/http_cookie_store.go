package webkit

// #include "http_cookie_store.h"
import "C"
import (
	"github.com/hsiafan/cocoa/objc"
	"unsafe"
)

type HTTPCookieStore interface {
	objc.Object
	AddObserver(observer objc.Object)
	RemoveObserver(observer objc.Object)
}

type WKHTTPCookieStore struct {
	objc.NSObject
}

func MakeHTTPCookieStore(ptr unsafe.Pointer) WKHTTPCookieStore {
	return WKHTTPCookieStore{
		NSObject: objc.MakeObject(ptr),
	}
}

func AllocHTTPCookieStore() WKHTTPCookieStore {
	result_ := C.C_WKHTTPCookieStore_AllocHTTPCookieStore()
	return MakeHTTPCookieStore(result_)
}

func (w WKHTTPCookieStore) Autorelease() WKHTTPCookieStore {
	result_ := C.C_WKHTTPCookieStore_Autorelease(w.Ptr())
	return MakeHTTPCookieStore(result_)
}

func (w WKHTTPCookieStore) Retain() WKHTTPCookieStore {
	result_ := C.C_WKHTTPCookieStore_Retain(w.Ptr())
	return MakeHTTPCookieStore(result_)
}

func (w WKHTTPCookieStore) AddObserver(observer objc.Object) {
	C.C_WKHTTPCookieStore_AddObserver(w.Ptr(), objc.ExtractPtr(observer))
}

func (w WKHTTPCookieStore) RemoveObserver(observer objc.Object) {
	C.C_WKHTTPCookieStore_RemoveObserver(w.Ptr(), objc.ExtractPtr(observer))
}
