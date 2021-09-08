package foundation

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

func MakeNotification(ptr unsafe.Pointer) NSNotification {
	return NSNotification{
		NSObject: objc.MakeObject(ptr),
	}
}

func (n NSNotification) Init() NSNotification {
	result_ := C.C_NSNotification_Init(n.Ptr())
	return MakeNotification(result_)
}

func (n NSNotification) InitWithCoder(coder Coder) NSNotification {
	result_ := C.C_NSNotification_InitWithCoder(n.Ptr(), objc.ExtractPtr(coder))
	return MakeNotification(result_)
}

func NotificationWithName_Object(aName NotificationName, anObject objc.Object) NSNotification {
	result_ := C.C_NSNotification_NotificationWithName_Object(NewString(string(aName)).Ptr(), objc.ExtractPtr(anObject))
	return MakeNotification(result_)
}

func AllocNotification() NSNotification {
	result_ := C.C_NSNotification_AllocNotification()
	return MakeNotification(result_)
}

func NewNotification() NSNotification {
	result_ := C.C_NSNotification_NewNotification()
	return MakeNotification(result_)
}

func (n NSNotification) Autorelease() NSNotification {
	result_ := C.C_NSNotification_Autorelease(n.Ptr())
	return MakeNotification(result_)
}

func (n NSNotification) Retain() NSNotification {
	result_ := C.C_NSNotification_Retain(n.Ptr())
	return MakeNotification(result_)
}

func (n NSNotification) Name() NotificationName {
	result_ := C.C_NSNotification_Name(n.Ptr())
	return NotificationName(MakeString(result_).String())
}

func (n NSNotification) Object() objc.Object {
	result_ := C.C_NSNotification_Object(n.Ptr())
	return objc.MakeObject(result_)
}
