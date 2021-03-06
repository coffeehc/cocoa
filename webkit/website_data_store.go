package webkit

// #include "website_data_store.h"
import "C"
import (
	"github.com/hsiafan/cocoa/foundation"
	"github.com/hsiafan/cocoa/objc"
	"unsafe"
)

type WebsiteDataStore interface {
	objc.Object
	IsPersistent() bool
	HttpCookieStore() HTTPCookieStore
}

type WKWebsiteDataStore struct {
	objc.NSObject
}

func MakeWebsiteDataStore(ptr unsafe.Pointer) WKWebsiteDataStore {
	return WKWebsiteDataStore{
		NSObject: objc.MakeObject(ptr),
	}
}

func AllocWebsiteDataStore() WKWebsiteDataStore {
	result_ := C.C_WKWebsiteDataStore_AllocWebsiteDataStore()
	return MakeWebsiteDataStore(result_)
}

func (w WKWebsiteDataStore) Autorelease() WKWebsiteDataStore {
	result_ := C.C_WKWebsiteDataStore_Autorelease(w.Ptr())
	return MakeWebsiteDataStore(result_)
}

func (w WKWebsiteDataStore) Retain() WKWebsiteDataStore {
	result_ := C.C_WKWebsiteDataStore_Retain(w.Ptr())
	return MakeWebsiteDataStore(result_)
}

func WebsiteDataStore_DefaultDataStore() WebsiteDataStore {
	result_ := C.C_WKWebsiteDataStore_WebsiteDataStore_DefaultDataStore()
	return MakeWebsiteDataStore(result_)
}

func WebsiteDataStore_NonPersistentDataStore() WebsiteDataStore {
	result_ := C.C_WKWebsiteDataStore_WebsiteDataStore_NonPersistentDataStore()
	return MakeWebsiteDataStore(result_)
}

func WebsiteDataStore_AllWebsiteDataTypes() foundation.Set {
	result_ := C.C_WKWebsiteDataStore_WebsiteDataStore_AllWebsiteDataTypes()
	return foundation.MakeSet(result_)
}

func (w WKWebsiteDataStore) IsPersistent() bool {
	result_ := C.C_WKWebsiteDataStore_IsPersistent(w.Ptr())
	return bool(result_)
}

func (w WKWebsiteDataStore) HttpCookieStore() HTTPCookieStore {
	result_ := C.C_WKWebsiteDataStore_HttpCookieStore(w.Ptr())
	return MakeHTTPCookieStore(result_)
}
