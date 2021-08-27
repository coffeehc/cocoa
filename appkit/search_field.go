package appkit

// #include "search_field.h"
import "C"
import (
	"github.com/hsiafan/cocoa/coregraphics"
	"github.com/hsiafan/cocoa/foundation"
	"github.com/hsiafan/cocoa/objc"
	"unsafe"
)

type SearchField interface {
	TextField
	SearchMenuTemplate() Menu
	SetSearchMenuTemplate(value Menu)
	SendsSearchStringImmediately() bool
	SetSendsSearchStringImmediately(value bool)
	SendsWholeSearchString() bool
	SetSendsWholeSearchString(value bool)
	RecentSearches() []string
	SetRecentSearches(value []string)
	MaximumRecents() int
	SetMaximumRecents(value int)
	RecentsAutosaveName() SearchFieldRecentsAutosaveName
	SetRecentsAutosaveName(value SearchFieldRecentsAutosaveName)
	CancelButtonBounds() foundation.Rect
	SearchButtonBounds() foundation.Rect
	SearchTextBounds() foundation.Rect
}

type NSSearchField struct {
	NSTextField
}

func MakeSearchField(ptr unsafe.Pointer) NSSearchField {
	return NSSearchField{
		NSTextField: MakeTextField(ptr),
	}
}

func AllocSearchField() NSSearchField {
	return MakeSearchField(C.C_SearchField_Alloc())
}

func (n NSSearchField) InitWithFrame(frameRect foundation.Rect) SearchField {
	result_ := C.C_NSSearchField_InitWithFrame(n.Ptr(), *(*C.CGRect)(coregraphics.ToCGRectPointer(coregraphics.Rect(frameRect))))
	return MakeSearchField(result_)
}

func (n NSSearchField) InitWithCoder(coder foundation.Coder) SearchField {
	result_ := C.C_NSSearchField_InitWithCoder(n.Ptr(), objc.ExtractPtr(coder))
	return MakeSearchField(result_)
}

func (n NSSearchField) Init() SearchField {
	result_ := C.C_NSSearchField_Init(n.Ptr())
	return MakeSearchField(result_)
}

func (n NSSearchField) SearchMenuTemplate() Menu {
	result_ := C.C_NSSearchField_SearchMenuTemplate(n.Ptr())
	return MakeMenu(result_)
}

func (n NSSearchField) SetSearchMenuTemplate(value Menu) {
	C.C_NSSearchField_SetSearchMenuTemplate(n.Ptr(), objc.ExtractPtr(value))
}

func (n NSSearchField) SendsSearchStringImmediately() bool {
	result_ := C.C_NSSearchField_SendsSearchStringImmediately(n.Ptr())
	return bool(result_)
}

func (n NSSearchField) SetSendsSearchStringImmediately(value bool) {
	C.C_NSSearchField_SetSendsSearchStringImmediately(n.Ptr(), C.bool(value))
}

func (n NSSearchField) SendsWholeSearchString() bool {
	result_ := C.C_NSSearchField_SendsWholeSearchString(n.Ptr())
	return bool(result_)
}

func (n NSSearchField) SetSendsWholeSearchString(value bool) {
	C.C_NSSearchField_SetSendsWholeSearchString(n.Ptr(), C.bool(value))
}

func (n NSSearchField) RecentSearches() []string {
	result_ := C.C_NSSearchField_RecentSearches(n.Ptr())
	if result_.len > 0 {
		defer C.free(result_.data)
	}
	result_Slice := unsafe.Slice((*unsafe.Pointer)(result_.data), int(result_.len))
	var goResult_ = make([]string, len(result_Slice))
	for idx, r := range result_Slice {
		goResult_[idx] = foundation.MakeString(r).String()
	}
	return goResult_
}

func (n NSSearchField) SetRecentSearches(value []string) {
	var cValue C.Array
	if len(value) > 0 {
		cValueData := make([]unsafe.Pointer, len(value))
		for idx, v := range value {
			cValueData[idx] = foundation.NewString(v).Ptr()
		}
		cValue.data = unsafe.Pointer(&cValueData[0])
		cValue.len = C.int(len(value))
	}
	C.C_NSSearchField_SetRecentSearches(n.Ptr(), cValue)
}

func (n NSSearchField) MaximumRecents() int {
	result_ := C.C_NSSearchField_MaximumRecents(n.Ptr())
	return int(result_)
}

func (n NSSearchField) SetMaximumRecents(value int) {
	C.C_NSSearchField_SetMaximumRecents(n.Ptr(), C.int(value))
}

func (n NSSearchField) RecentsAutosaveName() SearchFieldRecentsAutosaveName {
	result_ := C.C_NSSearchField_RecentsAutosaveName(n.Ptr())
	return SearchFieldRecentsAutosaveName(foundation.MakeString(result_).String())
}

func (n NSSearchField) SetRecentsAutosaveName(value SearchFieldRecentsAutosaveName) {
	C.C_NSSearchField_SetRecentsAutosaveName(n.Ptr(), foundation.NewString(string(value)).Ptr())
}

func (n NSSearchField) CancelButtonBounds() foundation.Rect {
	result_ := C.C_NSSearchField_CancelButtonBounds(n.Ptr())
	return foundation.Rect(coregraphics.FromCGRectPointer(unsafe.Pointer(&result_)))
}

func (n NSSearchField) SearchButtonBounds() foundation.Rect {
	result_ := C.C_NSSearchField_SearchButtonBounds(n.Ptr())
	return foundation.Rect(coregraphics.FromCGRectPointer(unsafe.Pointer(&result_)))
}

func (n NSSearchField) SearchTextBounds() foundation.Rect {
	result_ := C.C_NSSearchField_SearchTextBounds(n.Ptr())
	return foundation.Rect(coregraphics.FromCGRectPointer(unsafe.Pointer(&result_)))
}
