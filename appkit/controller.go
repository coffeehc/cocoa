package appkit

// #include "controller.h"
import "C"
import (
	"github.com/hsiafan/cocoa/foundation"
	"github.com/hsiafan/cocoa/objc"
	"unsafe"
)

type Controller interface {
	objc.Object
	ObjectDidBeginEditing(editor objc.Object)
	ObjectDidEndEditing(editor objc.Object)
	CommitEditing() bool
	DiscardEditing()
	IsEditing() bool
}

type NSController struct {
	objc.NSObject
}

func MakeController(ptr unsafe.Pointer) NSController {
	return NSController{
		NSObject: objc.MakeObject(ptr),
	}
}

func AllocController() NSController {
	return MakeController(C.C_Controller_Alloc())
}

func (n NSController) Init() Controller {
	result_ := C.C_NSController_Init(n.Ptr())
	return MakeController(result_)
}

func (n NSController) InitWithCoder(coder foundation.Coder) Controller {
	result_ := C.C_NSController_InitWithCoder(n.Ptr(), objc.ExtractPtr(coder))
	return MakeController(result_)
}

func (n NSController) ObjectDidBeginEditing(editor objc.Object) {
	C.C_NSController_ObjectDidBeginEditing(n.Ptr(), objc.ExtractPtr(editor))
}

func (n NSController) ObjectDidEndEditing(editor objc.Object) {
	C.C_NSController_ObjectDidEndEditing(n.Ptr(), objc.ExtractPtr(editor))
}

func (n NSController) CommitEditing() bool {
	result_ := C.C_NSController_CommitEditing(n.Ptr())
	return bool(result_)
}

func (n NSController) DiscardEditing() {
	C.C_NSController_DiscardEditing(n.Ptr())
}

func (n NSController) IsEditing() bool {
	result_ := C.C_NSController_IsEditing(n.Ptr())
	return bool(result_)
}
