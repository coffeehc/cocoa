package foundation

// #cgo CFLAGS: -x objective-c
// #cgo LDFLAGS: -framework Cocoa
// #import "notification.h"
import "C"
import (
	"github.com/hsiafan/cocoa/objc"
	"unsafe"
)

// Notification wrap cocoa NSNotification
type Notification interface {
	objc.Object
	// Name return the notification message name
	Name() string
	// Object return the sender of this notification
	Object() objc.Object
}

var _ Notification = (*NSNotification)(nil)

// MakeObject create new Notification, from native pointer of NSNotification, and the sender object
func MakeNotification(ptr unsafe.Pointer) *NSNotification {
	if ptr == nil {
		return nil
	}
	return &NSNotification{
		NSObject: *objc.MakeObject(ptr),
	}
}

type NSNotification struct {
	objc.NSObject
}

func (n *NSNotification) Name() string {
	cstr := C.Notification_Name(n.Ptr())
	return C.GoString(cstr)
}

func (n *NSNotification) Object() objc.Object {
	return objc.MakeObject(unsafe.Pointer(C.Notification_Object(n.Ptr())))
}
