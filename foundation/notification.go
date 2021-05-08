package foundation

// #cgo CFLAGS: -x objective-c
// #cgo LDFLAGS: -framework Foundation
// #include "notification.h"
import "C"
import (
	"github.com/hsiafan/cocoa/objc"
	"unsafe"
)

type Notification interface {
	objc.Object
	Name() NotificationName
	Object() objc.Object
}

type NSNotification struct {
	objc.NSObject
}

func MakeNotification(ptr unsafe.Pointer) *NSNotification {
	if ptr == nil {
		return nil
	}
	return &NSNotification{
		NSObject: *objc.MakeObject(ptr),
	}
}

func AllocNotification() *NSNotification {
	return MakeNotification(C.C_Notification_Alloc())
}

func (n *NSNotification) Init() Notification {
	result := C.C_NSNotification_Init(n.Ptr())
	return MakeNotification(result)
}

func (n *NSNotification) InitWithCoder(coder Coder) Notification {
	result := C.C_NSNotification_InitWithCoder(n.Ptr(), toPointer(coder))
	return MakeNotification(result)
}

func NotificationWithName_Object(aName NotificationName, anObject objc.Object) Notification {
	result := C.C_NSNotification_NotificationWithName_Object(NewString(string(aName)).Ptr(), toPointer(anObject))
	return MakeNotification(result)
}

func (n *NSNotification) Name() NotificationName {
	result := C.C_NSNotification_Name(n.Ptr())
	return NotificationName(MakeString(result).String())
}

func (n *NSNotification) Object() objc.Object {
	result := C.C_NSNotification_Object(n.Ptr())
	return objc.MakeObject(result)
}
