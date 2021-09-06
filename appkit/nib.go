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

func AllocNib() NSNib {
	return MakeNib(C.C_Nib_Alloc())
}

func (n NSNib) InitWithNibNamed_Bundle(nibName NibName, bundle foundation.Bundle) Nib {
	result_ := C.C_NSNib_InitWithNibNamed_Bundle(n.Ptr(), foundation.NewString(string(nibName)).Ptr(), objc.ExtractPtr(bundle))
	return MakeNib(result_)
}

func (n NSNib) InitWithNibData_Bundle(nibData []byte, bundle foundation.Bundle) Nib {
	result_ := C.C_NSNib_InitWithNibData_Bundle(n.Ptr(), foundation.NewData(nibData).Ptr(), objc.ExtractPtr(bundle))
	return MakeNib(result_)
}

func (n NSNib) Init() Nib {
	result_ := C.C_NSNib_Init(n.Ptr())
	return MakeNib(result_)
}
