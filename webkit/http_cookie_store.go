package webkit

// #cgo CFLAGS: -x objective-c
// #cgo LDFLAGS: -framework WebKit
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

func MakeHTTPCookieStore(ptr unsafe.Pointer) *WKHTTPCookieStore {
	if ptr == nil {
		return nil
	}
	return &WKHTTPCookieStore{
		NSObject: *objc.MakeObject(ptr),
	}
}

func AllocHTTPCookieStore() *WKHTTPCookieStore {
	return MakeHTTPCookieStore(C.C_HTTPCookieStore_Alloc())
}

func (w *WKHTTPCookieStore) AddObserver(observer objc.Object) {
	C.C_WKHTTPCookieStore_AddObserver(w.Ptr(), objc.ExtractPtr(observer))
}

func (w *WKHTTPCookieStore) RemoveObserver(observer objc.Object) {
	C.C_WKHTTPCookieStore_RemoveObserver(w.Ptr(), objc.ExtractPtr(observer))
}
