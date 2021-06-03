package appkit

// #include "text_storage.h"
import "C"
import (
	"github.com/hsiafan/cocoa/foundation"
	"github.com/hsiafan/cocoa/objc"
	"unsafe"
)

type TextStorage interface {
	foundation.MutableAttributedString
	AddLayoutManager(aLayoutManager LayoutManager)
	RemoveLayoutManager(aLayoutManager LayoutManager)
	Edited_Range_ChangeInLength(editedMask TextStorageEditActions, editedRange foundation.Range, delta int)
	ProcessEditing()
	InvalidateAttributesInRange(_range foundation.Range)
	EnsureAttributesAreFixedInRange(_range foundation.Range)
	Delegate() objc.Object
	SetDelegate(value objc.Object)
	LayoutManagers() []LayoutManager
	FixesAttributesLazily() bool
	EditedMask() TextStorageEditActions
	EditedRange() foundation.Range
	ChangeInLength() int
	AttributeRuns() []TextStorage
	SetAttributeRuns(value []TextStorage)
	Paragraphs() []TextStorage
	SetParagraphs(value []TextStorage)
	Words() []TextStorage
	SetWords(value []TextStorage)
	Characters() []TextStorage
	SetCharacters(value []TextStorage)
	Font() Font
	SetFont(value Font)
	ForegroundColor() Color
	SetForegroundColor(value Color)
}

type NSTextStorage struct {
	foundation.NSMutableAttributedString
}

func MakeTextStorage(ptr unsafe.Pointer) NSTextStorage {
	return NSTextStorage{
		NSMutableAttributedString: foundation.MakeMutableAttributedString(ptr),
	}
}

func AllocTextStorage() NSTextStorage {
	return MakeTextStorage(C.C_TextStorage_Alloc())
}

func (n NSTextStorage) InitWithString(str string) TextStorage {
	result_ := C.C_NSTextStorage_InitWithString(n.Ptr(), foundation.NewString(str).Ptr())
	return MakeTextStorage(result_)
}

func (n NSTextStorage) InitWithAttributedString(attrStr foundation.AttributedString) TextStorage {
	result_ := C.C_NSTextStorage_InitWithAttributedString(n.Ptr(), objc.ExtractPtr(attrStr))
	return MakeTextStorage(result_)
}

func (n NSTextStorage) Init() TextStorage {
	result_ := C.C_NSTextStorage_Init(n.Ptr())
	return MakeTextStorage(result_)
}

func (n NSTextStorage) AddLayoutManager(aLayoutManager LayoutManager) {
	C.C_NSTextStorage_AddLayoutManager(n.Ptr(), objc.ExtractPtr(aLayoutManager))
}

func (n NSTextStorage) RemoveLayoutManager(aLayoutManager LayoutManager) {
	C.C_NSTextStorage_RemoveLayoutManager(n.Ptr(), objc.ExtractPtr(aLayoutManager))
}

func (n NSTextStorage) Edited_Range_ChangeInLength(editedMask TextStorageEditActions, editedRange foundation.Range, delta int) {
	C.C_NSTextStorage_Edited_Range_ChangeInLength(n.Ptr(), C.uint(uint(editedMask)), *(*C.NSRange)(foundation.ToNSRangePointer(editedRange)), C.int(delta))
}

func (n NSTextStorage) ProcessEditing() {
	C.C_NSTextStorage_ProcessEditing(n.Ptr())
}

func (n NSTextStorage) InvalidateAttributesInRange(_range foundation.Range) {
	C.C_NSTextStorage_InvalidateAttributesInRange(n.Ptr(), *(*C.NSRange)(foundation.ToNSRangePointer(_range)))
}

func (n NSTextStorage) EnsureAttributesAreFixedInRange(_range foundation.Range) {
	C.C_NSTextStorage_EnsureAttributesAreFixedInRange(n.Ptr(), *(*C.NSRange)(foundation.ToNSRangePointer(_range)))
}

func (n NSTextStorage) Delegate() objc.Object {
	result_ := C.C_NSTextStorage_Delegate(n.Ptr())
	return objc.MakeObject(result_)
}

func (n NSTextStorage) SetDelegate(value objc.Object) {
	C.C_NSTextStorage_SetDelegate(n.Ptr(), objc.ExtractPtr(value))
}

func (n NSTextStorage) LayoutManagers() []LayoutManager {
	result_ := C.C_NSTextStorage_LayoutManagers(n.Ptr())
	defer C.free(result_.data)
	result_Slice := (*[1 << 28]unsafe.Pointer)(unsafe.Pointer(result_.data))[:result_.len:result_.len]
	var goResult_ = make([]LayoutManager, len(result_Slice))
	for idx, r := range result_Slice {
		goResult_[idx] = MakeLayoutManager(r)
	}
	return goResult_
}

func (n NSTextStorage) FixesAttributesLazily() bool {
	result_ := C.C_NSTextStorage_FixesAttributesLazily(n.Ptr())
	return bool(result_)
}

func (n NSTextStorage) EditedMask() TextStorageEditActions {
	result_ := C.C_NSTextStorage_EditedMask(n.Ptr())
	return TextStorageEditActions(uint(result_))
}

func (n NSTextStorage) EditedRange() foundation.Range {
	result_ := C.C_NSTextStorage_EditedRange(n.Ptr())
	return foundation.FromNSRangePointer(unsafe.Pointer(&result_))
}

func (n NSTextStorage) ChangeInLength() int {
	result_ := C.C_NSTextStorage_ChangeInLength(n.Ptr())
	return int(result_)
}

func (n NSTextStorage) AttributeRuns() []TextStorage {
	result_ := C.C_NSTextStorage_AttributeRuns(n.Ptr())
	defer C.free(result_.data)
	result_Slice := (*[1 << 28]unsafe.Pointer)(unsafe.Pointer(result_.data))[:result_.len:result_.len]
	var goResult_ = make([]TextStorage, len(result_Slice))
	for idx, r := range result_Slice {
		goResult_[idx] = MakeTextStorage(r)
	}
	return goResult_
}

func (n NSTextStorage) SetAttributeRuns(value []TextStorage) {
	cValueData := make([]unsafe.Pointer, len(value))
	for idx, v := range value {
		cValueData[idx] = objc.ExtractPtr(v)
	}
	cValue := C.Array{data: unsafe.Pointer(&cValueData[0]), len: C.int(len(value))}
	C.C_NSTextStorage_SetAttributeRuns(n.Ptr(), cValue)
}

func (n NSTextStorage) Paragraphs() []TextStorage {
	result_ := C.C_NSTextStorage_Paragraphs(n.Ptr())
	defer C.free(result_.data)
	result_Slice := (*[1 << 28]unsafe.Pointer)(unsafe.Pointer(result_.data))[:result_.len:result_.len]
	var goResult_ = make([]TextStorage, len(result_Slice))
	for idx, r := range result_Slice {
		goResult_[idx] = MakeTextStorage(r)
	}
	return goResult_
}

func (n NSTextStorage) SetParagraphs(value []TextStorage) {
	cValueData := make([]unsafe.Pointer, len(value))
	for idx, v := range value {
		cValueData[idx] = objc.ExtractPtr(v)
	}
	cValue := C.Array{data: unsafe.Pointer(&cValueData[0]), len: C.int(len(value))}
	C.C_NSTextStorage_SetParagraphs(n.Ptr(), cValue)
}

func (n NSTextStorage) Words() []TextStorage {
	result_ := C.C_NSTextStorage_Words(n.Ptr())
	defer C.free(result_.data)
	result_Slice := (*[1 << 28]unsafe.Pointer)(unsafe.Pointer(result_.data))[:result_.len:result_.len]
	var goResult_ = make([]TextStorage, len(result_Slice))
	for idx, r := range result_Slice {
		goResult_[idx] = MakeTextStorage(r)
	}
	return goResult_
}

func (n NSTextStorage) SetWords(value []TextStorage) {
	cValueData := make([]unsafe.Pointer, len(value))
	for idx, v := range value {
		cValueData[idx] = objc.ExtractPtr(v)
	}
	cValue := C.Array{data: unsafe.Pointer(&cValueData[0]), len: C.int(len(value))}
	C.C_NSTextStorage_SetWords(n.Ptr(), cValue)
}

func (n NSTextStorage) Characters() []TextStorage {
	result_ := C.C_NSTextStorage_Characters(n.Ptr())
	defer C.free(result_.data)
	result_Slice := (*[1 << 28]unsafe.Pointer)(unsafe.Pointer(result_.data))[:result_.len:result_.len]
	var goResult_ = make([]TextStorage, len(result_Slice))
	for idx, r := range result_Slice {
		goResult_[idx] = MakeTextStorage(r)
	}
	return goResult_
}

func (n NSTextStorage) SetCharacters(value []TextStorage) {
	cValueData := make([]unsafe.Pointer, len(value))
	for idx, v := range value {
		cValueData[idx] = objc.ExtractPtr(v)
	}
	cValue := C.Array{data: unsafe.Pointer(&cValueData[0]), len: C.int(len(value))}
	C.C_NSTextStorage_SetCharacters(n.Ptr(), cValue)
}

func (n NSTextStorage) Font() Font {
	result_ := C.C_NSTextStorage_Font(n.Ptr())
	return MakeFont(result_)
}

func (n NSTextStorage) SetFont(value Font) {
	C.C_NSTextStorage_SetFont(n.Ptr(), objc.ExtractPtr(value))
}

func (n NSTextStorage) ForegroundColor() Color {
	result_ := C.C_NSTextStorage_ForegroundColor(n.Ptr())
	return MakeColor(result_)
}

func (n NSTextStorage) SetForegroundColor(value Color) {
	C.C_NSTextStorage_SetForegroundColor(n.Ptr(), objc.ExtractPtr(value))
}
