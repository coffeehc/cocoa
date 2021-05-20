package appkit

// #include "browser_cell.h"
import "C"
import (
	"github.com/hsiafan/cocoa/foundation"
	"github.com/hsiafan/cocoa/objc"
	"unsafe"
)

type BrowserCell interface {
	Cell
	Reset()
	Set()
	HighlightColorInView(controlView View) Color
	AlternateImage() Image
	SetAlternateImage(value Image)
	IsLeaf() bool
	SetLeaf(value bool)
	IsLoaded() bool
	SetLoaded(value bool)
}

type NSBrowserCell struct {
	NSCell
}

func MakeBrowserCell(ptr unsafe.Pointer) *NSBrowserCell {
	if ptr == nil {
		return nil
	}
	return &NSBrowserCell{
		NSCell: *MakeCell(ptr),
	}
}

func AllocBrowserCell() *NSBrowserCell {
	return MakeBrowserCell(C.C_BrowserCell_Alloc())
}

func (n *NSBrowserCell) InitWithCoder(coder foundation.Coder) BrowserCell {
	result_ := C.C_NSBrowserCell_InitWithCoder(n.Ptr(), objc.ExtractPtr(coder))
	return MakeBrowserCell(result_)
}

func (n *NSBrowserCell) InitImageCell(image Image) BrowserCell {
	result_ := C.C_NSBrowserCell_InitImageCell(n.Ptr(), objc.ExtractPtr(image))
	return MakeBrowserCell(result_)
}

func (n *NSBrowserCell) InitTextCell(_string string) BrowserCell {
	result_ := C.C_NSBrowserCell_InitTextCell(n.Ptr(), foundation.NewString(_string).Ptr())
	return MakeBrowserCell(result_)
}

func (n *NSBrowserCell) Init() BrowserCell {
	result_ := C.C_NSBrowserCell_Init(n.Ptr())
	return MakeBrowserCell(result_)
}

func (n *NSBrowserCell) Reset() {
	C.C_NSBrowserCell_Reset(n.Ptr())
}

func (n *NSBrowserCell) Set() {
	C.C_NSBrowserCell_Set(n.Ptr())
}

func (n *NSBrowserCell) HighlightColorInView(controlView View) Color {
	result_ := C.C_NSBrowserCell_HighlightColorInView(n.Ptr(), objc.ExtractPtr(controlView))
	return MakeColor(result_)
}

func BrowserCell_BranchImage() Image {
	result_ := C.C_NSBrowserCell_BrowserCell_BranchImage()
	return MakeImage(result_)
}

func BrowserCell_HighlightedBranchImage() Image {
	result_ := C.C_NSBrowserCell_BrowserCell_HighlightedBranchImage()
	return MakeImage(result_)
}

func (n *NSBrowserCell) AlternateImage() Image {
	result_ := C.C_NSBrowserCell_AlternateImage(n.Ptr())
	return MakeImage(result_)
}

func (n *NSBrowserCell) SetAlternateImage(value Image) {
	C.C_NSBrowserCell_SetAlternateImage(n.Ptr(), objc.ExtractPtr(value))
}

func (n *NSBrowserCell) IsLeaf() bool {
	result_ := C.C_NSBrowserCell_IsLeaf(n.Ptr())
	return bool(result_)
}

func (n *NSBrowserCell) SetLeaf(value bool) {
	C.C_NSBrowserCell_SetLeaf(n.Ptr(), C.bool(value))
}

func (n *NSBrowserCell) IsLoaded() bool {
	result_ := C.C_NSBrowserCell_IsLoaded(n.Ptr())
	return bool(result_)
}

func (n *NSBrowserCell) SetLoaded(value bool) {
	C.C_NSBrowserCell_SetLoaded(n.Ptr(), C.bool(value))
}
