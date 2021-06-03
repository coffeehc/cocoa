package appkit

// #include "path_component_cell.h"
import "C"
import (
	"github.com/hsiafan/cocoa/foundation"
	"github.com/hsiafan/cocoa/objc"
	"unsafe"
)

type PathComponentCell interface {
	TextFieldCell
	URL() foundation.URL
	SetURL(value foundation.URL)
}

type NSPathComponentCell struct {
	NSTextFieldCell
}

func MakePathComponentCell(ptr unsafe.Pointer) NSPathComponentCell {
	return NSPathComponentCell{
		NSTextFieldCell: MakeTextFieldCell(ptr),
	}
}

func AllocPathComponentCell() NSPathComponentCell {
	return MakePathComponentCell(C.C_PathComponentCell_Alloc())
}

func (n NSPathComponentCell) InitTextCell(_string string) PathComponentCell {
	result_ := C.C_NSPathComponentCell_InitTextCell(n.Ptr(), foundation.NewString(_string).Ptr())
	return MakePathComponentCell(result_)
}

func (n NSPathComponentCell) InitWithCoder(coder foundation.Coder) PathComponentCell {
	result_ := C.C_NSPathComponentCell_InitWithCoder(n.Ptr(), objc.ExtractPtr(coder))
	return MakePathComponentCell(result_)
}

func (n NSPathComponentCell) Init() PathComponentCell {
	result_ := C.C_NSPathComponentCell_Init(n.Ptr())
	return MakePathComponentCell(result_)
}

func (n NSPathComponentCell) URL() foundation.URL {
	result_ := C.C_NSPathComponentCell_URL(n.Ptr())
	return foundation.MakeURL(result_)
}

func (n NSPathComponentCell) SetURL(value foundation.URL) {
	C.C_NSPathComponentCell_SetURL(n.Ptr(), objc.ExtractPtr(value))
}
