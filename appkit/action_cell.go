package appkit

// #cgo CFLAGS: -x objective-c
// #cgo LDFLAGS: -framework AppKit
// #include "action_cell.h"
import "C"
import (
	"github.com/hsiafan/cocoa/foundation"
	"github.com/hsiafan/cocoa/objc"
	"unsafe"
)

type ActionCell interface {
	Cell
}

type NSActionCell struct {
	NSCell
}

func MakeActionCell(ptr unsafe.Pointer) *NSActionCell {
	if ptr == nil {
		return nil
	}
	return &NSActionCell{
		NSCell: *MakeCell(ptr),
	}
}

func AllocActionCell() *NSActionCell {
	return MakeActionCell(C.C_ActionCell_Alloc())
}

func (n *NSActionCell) InitImageCell(image Image) ActionCell {
	result := C.C_NSActionCell_InitImageCell(n.Ptr(), objc.ExtractPtr(image))
	return MakeActionCell(result)
}

func (n *NSActionCell) InitTextCell(_string string) ActionCell {
	result := C.C_NSActionCell_InitTextCell(n.Ptr(), foundation.NewString(_string).Ptr())
	return MakeActionCell(result)
}

func (n *NSActionCell) Init() ActionCell {
	result := C.C_NSActionCell_Init(n.Ptr())
	return MakeActionCell(result)
}

func (n *NSActionCell) InitWithCoder(coder foundation.Coder) ActionCell {
	result := C.C_NSActionCell_InitWithCoder(n.Ptr(), objc.ExtractPtr(coder))
	return MakeActionCell(result)
}
