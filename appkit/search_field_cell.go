package appkit

// #include "search_field_cell.h"
import "C"
import (
	"github.com/hsiafan/cocoa/coregraphics"
	"github.com/hsiafan/cocoa/foundation"
	"github.com/hsiafan/cocoa/objc"
	"unsafe"
)

type SearchFieldCell interface {
	TextFieldCell
	ResetSearchButtonCell()
	ResetCancelButtonCell()
	SearchTextRectForBounds(rect foundation.Rect) foundation.Rect
	SearchButtonRectForBounds(rect foundation.Rect) foundation.Rect
	CancelButtonRectForBounds(rect foundation.Rect) foundation.Rect
	SearchButtonCell() ButtonCell
	SetSearchButtonCell(value ButtonCell)
	CancelButtonCell() ButtonCell
	SetCancelButtonCell(value ButtonCell)
	SearchMenuTemplate() Menu
	SetSearchMenuTemplate(value Menu)
	SendsWholeSearchString() bool
	SetSendsWholeSearchString(value bool)
	SendsSearchStringImmediately() bool
	SetSendsSearchStringImmediately(value bool)
	MaximumRecents() int
	SetMaximumRecents(value int)
	RecentSearches() []string
	SetRecentSearches(value []string)
	RecentsAutosaveName() SearchFieldRecentsAutosaveName
	SetRecentsAutosaveName(value SearchFieldRecentsAutosaveName)
}

type NSSearchFieldCell struct {
	NSTextFieldCell
}

func MakeSearchFieldCell(ptr unsafe.Pointer) NSSearchFieldCell {
	return NSSearchFieldCell{
		NSTextFieldCell: MakeTextFieldCell(ptr),
	}
}

func AllocSearchFieldCell() NSSearchFieldCell {
	return MakeSearchFieldCell(C.C_SearchFieldCell_Alloc())
}

func (n NSSearchFieldCell) InitWithCoder(coder foundation.Coder) SearchFieldCell {
	result_ := C.C_NSSearchFieldCell_InitWithCoder(n.Ptr(), objc.ExtractPtr(coder))
	return MakeSearchFieldCell(result_)
}

func (n NSSearchFieldCell) InitTextCell(_string string) SearchFieldCell {
	result_ := C.C_NSSearchFieldCell_InitTextCell(n.Ptr(), foundation.NewString(_string).Ptr())
	return MakeSearchFieldCell(result_)
}

func (n NSSearchFieldCell) Init() SearchFieldCell {
	result_ := C.C_NSSearchFieldCell_Init(n.Ptr())
	return MakeSearchFieldCell(result_)
}

func (n NSSearchFieldCell) ResetSearchButtonCell() {
	C.C_NSSearchFieldCell_ResetSearchButtonCell(n.Ptr())
}

func (n NSSearchFieldCell) ResetCancelButtonCell() {
	C.C_NSSearchFieldCell_ResetCancelButtonCell(n.Ptr())
}

func (n NSSearchFieldCell) SearchTextRectForBounds(rect foundation.Rect) foundation.Rect {
	result_ := C.C_NSSearchFieldCell_SearchTextRectForBounds(n.Ptr(), *(*C.CGRect)(coregraphics.ToCGRectPointer(coregraphics.Rect(rect))))
	return foundation.Rect(coregraphics.FromCGRectPointer(unsafe.Pointer(&result_)))
}

func (n NSSearchFieldCell) SearchButtonRectForBounds(rect foundation.Rect) foundation.Rect {
	result_ := C.C_NSSearchFieldCell_SearchButtonRectForBounds(n.Ptr(), *(*C.CGRect)(coregraphics.ToCGRectPointer(coregraphics.Rect(rect))))
	return foundation.Rect(coregraphics.FromCGRectPointer(unsafe.Pointer(&result_)))
}

func (n NSSearchFieldCell) CancelButtonRectForBounds(rect foundation.Rect) foundation.Rect {
	result_ := C.C_NSSearchFieldCell_CancelButtonRectForBounds(n.Ptr(), *(*C.CGRect)(coregraphics.ToCGRectPointer(coregraphics.Rect(rect))))
	return foundation.Rect(coregraphics.FromCGRectPointer(unsafe.Pointer(&result_)))
}

func (n NSSearchFieldCell) SearchButtonCell() ButtonCell {
	result_ := C.C_NSSearchFieldCell_SearchButtonCell(n.Ptr())
	return MakeButtonCell(result_)
}

func (n NSSearchFieldCell) SetSearchButtonCell(value ButtonCell) {
	C.C_NSSearchFieldCell_SetSearchButtonCell(n.Ptr(), objc.ExtractPtr(value))
}

func (n NSSearchFieldCell) CancelButtonCell() ButtonCell {
	result_ := C.C_NSSearchFieldCell_CancelButtonCell(n.Ptr())
	return MakeButtonCell(result_)
}

func (n NSSearchFieldCell) SetCancelButtonCell(value ButtonCell) {
	C.C_NSSearchFieldCell_SetCancelButtonCell(n.Ptr(), objc.ExtractPtr(value))
}

func (n NSSearchFieldCell) SearchMenuTemplate() Menu {
	result_ := C.C_NSSearchFieldCell_SearchMenuTemplate(n.Ptr())
	return MakeMenu(result_)
}

func (n NSSearchFieldCell) SetSearchMenuTemplate(value Menu) {
	C.C_NSSearchFieldCell_SetSearchMenuTemplate(n.Ptr(), objc.ExtractPtr(value))
}

func (n NSSearchFieldCell) SendsWholeSearchString() bool {
	result_ := C.C_NSSearchFieldCell_SendsWholeSearchString(n.Ptr())
	return bool(result_)
}

func (n NSSearchFieldCell) SetSendsWholeSearchString(value bool) {
	C.C_NSSearchFieldCell_SetSendsWholeSearchString(n.Ptr(), C.bool(value))
}

func (n NSSearchFieldCell) SendsSearchStringImmediately() bool {
	result_ := C.C_NSSearchFieldCell_SendsSearchStringImmediately(n.Ptr())
	return bool(result_)
}

func (n NSSearchFieldCell) SetSendsSearchStringImmediately(value bool) {
	C.C_NSSearchFieldCell_SetSendsSearchStringImmediately(n.Ptr(), C.bool(value))
}

func (n NSSearchFieldCell) MaximumRecents() int {
	result_ := C.C_NSSearchFieldCell_MaximumRecents(n.Ptr())
	return int(result_)
}

func (n NSSearchFieldCell) SetMaximumRecents(value int) {
	C.C_NSSearchFieldCell_SetMaximumRecents(n.Ptr(), C.int(value))
}

func (n NSSearchFieldCell) RecentSearches() []string {
	result_ := C.C_NSSearchFieldCell_RecentSearches(n.Ptr())
	if result_.len > 0 {
		defer C.free(result_.data)
	}
	result_Slice := (*[1 << 28]unsafe.Pointer)(unsafe.Pointer(result_.data))[:result_.len:result_.len]
	var goResult_ = make([]string, len(result_Slice))
	for idx, r := range result_Slice {
		goResult_[idx] = foundation.MakeString(r).String()
	}
	return goResult_
}

func (n NSSearchFieldCell) SetRecentSearches(value []string) {
	var cValue C.Array
	if len(value) > 0 {
		cValueData := make([]unsafe.Pointer, len(value))
		for idx, v := range value {
			cValueData[idx] = foundation.NewString(v).Ptr()
		}
		cValue.data = unsafe.Pointer(&cValueData[0])
		cValue.len = C.int(len(value))
	}
	C.C_NSSearchFieldCell_SetRecentSearches(n.Ptr(), cValue)
}

func (n NSSearchFieldCell) RecentsAutosaveName() SearchFieldRecentsAutosaveName {
	result_ := C.C_NSSearchFieldCell_RecentsAutosaveName(n.Ptr())
	return SearchFieldRecentsAutosaveName(foundation.MakeString(result_).String())
}

func (n NSSearchFieldCell) SetRecentsAutosaveName(value SearchFieldRecentsAutosaveName) {
	C.C_NSSearchFieldCell_SetRecentsAutosaveName(n.Ptr(), foundation.NewString(string(value)).Ptr())
}
