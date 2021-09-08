package appkit

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

func MakeActionCell(ptr unsafe.Pointer) NSActionCell {
	return NSActionCell{
		NSCell: MakeCell(ptr),
	}
}

func (n NSActionCell) InitImageCell(image Image) NSActionCell {
	result_ := C.C_NSActionCell_InitImageCell(n.Ptr(), objc.ExtractPtr(image))
	return MakeActionCell(result_)
}

func (n NSActionCell) InitTextCell(_string string) NSActionCell {
	result_ := C.C_NSActionCell_InitTextCell(n.Ptr(), foundation.NewString(_string).Ptr())
	return MakeActionCell(result_)
}

func (n NSActionCell) Init() NSActionCell {
	result_ := C.C_NSActionCell_Init(n.Ptr())
	return MakeActionCell(result_)
}

func (n NSActionCell) InitWithCoder(coder foundation.Coder) NSActionCell {
	result_ := C.C_NSActionCell_InitWithCoder(n.Ptr(), objc.ExtractPtr(coder))
	return MakeActionCell(result_)
}

func AllocActionCell() NSActionCell {
	result_ := C.C_NSActionCell_AllocActionCell()
	return MakeActionCell(result_)
}

func NewActionCell() NSActionCell {
	result_ := C.C_NSActionCell_NewActionCell()
	return MakeActionCell(result_)
}

func (n NSActionCell) Autorelease() NSActionCell {
	result_ := C.C_NSActionCell_Autorelease(n.Ptr())
	return MakeActionCell(result_)
}

func (n NSActionCell) Retain() NSActionCell {
	result_ := C.C_NSActionCell_Retain(n.Ptr())
	return MakeActionCell(result_)
}
