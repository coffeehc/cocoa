package appkit

// #include "nib.h"
import "C"
import (
	"github.com/hsiafan/cocoa/foundation"
	"github.com/hsiafan/cocoa/objc"
	"unsafe"
)

type Nib interface {
	objc.Object
}

type NSNib struct {
	objc.NSObject
}

func MakeNib(ptr unsafe.Pointer) NSNib {
	return NSNib{
		NSObject: objc.MakeObject(ptr),
	}
}

func (n NSNib) InitWithNibNamed_Bundle(nibName NibName, bundle foundation.Bundle) NSNib {
	result_ := C.C_NSNib_InitWithNibNamed_Bundle(n.Ptr(), foundation.NewString(string(nibName)).Ptr(), objc.ExtractPtr(bundle))
	return MakeNib(result_)
}

func (n NSNib) InitWithNibData_Bundle(nibData []byte, bundle foundation.Bundle) NSNib {
	result_ := C.C_NSNib_InitWithNibData_Bundle(n.Ptr(), foundation.NewData(nibData).Ptr(), objc.ExtractPtr(bundle))
	return MakeNib(result_)
}

func AllocNib() NSNib {
	result_ := C.C_NSNib_AllocNib()
	return MakeNib(result_)
}

func (n NSNib) Init() NSNib {
	result_ := C.C_NSNib_Init(n.Ptr())
	return MakeNib(result_)
}

func NewNib() NSNib {
	result_ := C.C_NSNib_NewNib()
	return MakeNib(result_)
}

func (n NSNib) Autorelease() NSNib {
	result_ := C.C_NSNib_Autorelease(n.Ptr())
	return MakeNib(result_)
}

func (n NSNib) Retain() NSNib {
	result_ := C.C_NSNib_Retain(n.Ptr())
	return MakeNib(result_)
}
