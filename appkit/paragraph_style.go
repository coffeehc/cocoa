package appkit

// #include "paragraph_style.h"
import "C"
import (
	"github.com/hsiafan/cocoa/coregraphics"
	"github.com/hsiafan/cocoa/foundation"
	"github.com/hsiafan/cocoa/objc"
	"unsafe"
)

type ParagraphStyle interface {
	objc.Object
	Alignment() TextAlignment
	FirstLineHeadIndent() coregraphics.Float
	HeadIndent() coregraphics.Float
	TailIndent() coregraphics.Float
	LineHeightMultiple() coregraphics.Float
	MaximumLineHeight() coregraphics.Float
	MinimumLineHeight() coregraphics.Float
	LineSpacing() coregraphics.Float
	ParagraphSpacing() coregraphics.Float
	ParagraphSpacingBefore() coregraphics.Float
	TabStops() []TextTab
	DefaultTabInterval() coregraphics.Float
	TextBlocks() []TextBlock
	TextLists() []TextList
	LineBreakMode() LineBreakMode
	LineBreakStrategy() LineBreakStrategy
	HyphenationFactor() float32
	TighteningFactorForTruncation() float32
	AllowsDefaultTighteningForTruncation() bool
	HeaderLevel() int
	BaseWritingDirection() WritingDirection
}

type NSParagraphStyle struct {
	objc.NSObject
}

func MakeParagraphStyle(ptr unsafe.Pointer) NSParagraphStyle {
	return NSParagraphStyle{
		NSObject: objc.MakeObject(ptr),
	}
}

func AllocParagraphStyle() NSParagraphStyle {
	result_ := C.C_NSParagraphStyle_AllocParagraphStyle()
	return MakeParagraphStyle(result_)
}

func (n NSParagraphStyle) Init() NSParagraphStyle {
	result_ := C.C_NSParagraphStyle_Init(n.Ptr())
	return MakeParagraphStyle(result_)
}

func NewParagraphStyle() NSParagraphStyle {
	result_ := C.C_NSParagraphStyle_NewParagraphStyle()
	return MakeParagraphStyle(result_)
}

func (n NSParagraphStyle) Autorelease() NSParagraphStyle {
	result_ := C.C_NSParagraphStyle_Autorelease(n.Ptr())
	return MakeParagraphStyle(result_)
}

func (n NSParagraphStyle) Retain() NSParagraphStyle {
	result_ := C.C_NSParagraphStyle_Retain(n.Ptr())
	return MakeParagraphStyle(result_)
}

func ParagraphStyle_DefaultWritingDirectionForLanguage(languageName string) WritingDirection {
	result_ := C.C_NSParagraphStyle_ParagraphStyle_DefaultWritingDirectionForLanguage(foundation.NewString(languageName).Ptr())
	return WritingDirection(int(result_))
}

func DefaultParagraphStyle() ParagraphStyle {
	result_ := C.C_NSParagraphStyle_DefaultParagraphStyle()
	return MakeParagraphStyle(result_)
}

func (n NSParagraphStyle) Alignment() TextAlignment {
	result_ := C.C_NSParagraphStyle_Alignment(n.Ptr())
	return TextAlignment(int(result_))
}

func (n NSParagraphStyle) FirstLineHeadIndent() coregraphics.Float {
	result_ := C.C_NSParagraphStyle_FirstLineHeadIndent(n.Ptr())
	return coregraphics.Float(float64(result_))
}

func (n NSParagraphStyle) HeadIndent() coregraphics.Float {
	result_ := C.C_NSParagraphStyle_HeadIndent(n.Ptr())
	return coregraphics.Float(float64(result_))
}

func (n NSParagraphStyle) TailIndent() coregraphics.Float {
	result_ := C.C_NSParagraphStyle_TailIndent(n.Ptr())
	return coregraphics.Float(float64(result_))
}

func (n NSParagraphStyle) LineHeightMultiple() coregraphics.Float {
	result_ := C.C_NSParagraphStyle_LineHeightMultiple(n.Ptr())
	return coregraphics.Float(float64(result_))
}

func (n NSParagraphStyle) MaximumLineHeight() coregraphics.Float {
	result_ := C.C_NSParagraphStyle_MaximumLineHeight(n.Ptr())
	return coregraphics.Float(float64(result_))
}

func (n NSParagraphStyle) MinimumLineHeight() coregraphics.Float {
	result_ := C.C_NSParagraphStyle_MinimumLineHeight(n.Ptr())
	return coregraphics.Float(float64(result_))
}

func (n NSParagraphStyle) LineSpacing() coregraphics.Float {
	result_ := C.C_NSParagraphStyle_LineSpacing(n.Ptr())
	return coregraphics.Float(float64(result_))
}

func (n NSParagraphStyle) ParagraphSpacing() coregraphics.Float {
	result_ := C.C_NSParagraphStyle_ParagraphSpacing(n.Ptr())
	return coregraphics.Float(float64(result_))
}

func (n NSParagraphStyle) ParagraphSpacingBefore() coregraphics.Float {
	result_ := C.C_NSParagraphStyle_ParagraphSpacingBefore(n.Ptr())
	return coregraphics.Float(float64(result_))
}

func (n NSParagraphStyle) TabStops() []TextTab {
	result_ := C.C_NSParagraphStyle_TabStops(n.Ptr())
	if result_.len > 0 {
		defer C.free(result_.data)
	}
	result_Slice := unsafe.Slice((*unsafe.Pointer)(result_.data), int(result_.len))
	var goResult_ = make([]TextTab, len(result_Slice))
	for idx, r := range result_Slice {
		goResult_[idx] = MakeTextTab(r)
	}
	return goResult_
}

func (n NSParagraphStyle) DefaultTabInterval() coregraphics.Float {
	result_ := C.C_NSParagraphStyle_DefaultTabInterval(n.Ptr())
	return coregraphics.Float(float64(result_))
}

func (n NSParagraphStyle) TextBlocks() []TextBlock {
	result_ := C.C_NSParagraphStyle_TextBlocks(n.Ptr())
	if result_.len > 0 {
		defer C.free(result_.data)
	}
	result_Slice := unsafe.Slice((*unsafe.Pointer)(result_.data), int(result_.len))
	var goResult_ = make([]TextBlock, len(result_Slice))
	for idx, r := range result_Slice {
		goResult_[idx] = MakeTextBlock(r)
	}
	return goResult_
}

func (n NSParagraphStyle) TextLists() []TextList {
	result_ := C.C_NSParagraphStyle_TextLists(n.Ptr())
	if result_.len > 0 {
		defer C.free(result_.data)
	}
	result_Slice := unsafe.Slice((*unsafe.Pointer)(result_.data), int(result_.len))
	var goResult_ = make([]TextList, len(result_Slice))
	for idx, r := range result_Slice {
		goResult_[idx] = MakeTextList(r)
	}
	return goResult_
}

func (n NSParagraphStyle) LineBreakMode() LineBreakMode {
	result_ := C.C_NSParagraphStyle_LineBreakMode(n.Ptr())
	return LineBreakMode(uint(result_))
}

func (n NSParagraphStyle) LineBreakStrategy() LineBreakStrategy {
	result_ := C.C_NSParagraphStyle_LineBreakStrategy(n.Ptr())
	return LineBreakStrategy(uint(result_))
}

func (n NSParagraphStyle) HyphenationFactor() float32 {
	result_ := C.C_NSParagraphStyle_HyphenationFactor(n.Ptr())
	return float32(result_)
}

func (n NSParagraphStyle) TighteningFactorForTruncation() float32 {
	result_ := C.C_NSParagraphStyle_TighteningFactorForTruncation(n.Ptr())
	return float32(result_)
}

func (n NSParagraphStyle) AllowsDefaultTighteningForTruncation() bool {
	result_ := C.C_NSParagraphStyle_AllowsDefaultTighteningForTruncation(n.Ptr())
	return bool(result_)
}

func (n NSParagraphStyle) HeaderLevel() int {
	result_ := C.C_NSParagraphStyle_HeaderLevel(n.Ptr())
	return int(result_)
}

func (n NSParagraphStyle) BaseWritingDirection() WritingDirection {
	result_ := C.C_NSParagraphStyle_BaseWritingDirection(n.Ptr())
	return WritingDirection(int(result_))
}

type MutableParagraphStyle interface {
	ParagraphStyle
	SetParagraphStyle(obj ParagraphStyle)
	AddTabStop(anObject TextTab)
	RemoveTabStop(anObject TextTab)
	SetAlignment(value TextAlignment)
	SetFirstLineHeadIndent(value coregraphics.Float)
	SetHeadIndent(value coregraphics.Float)
	SetTailIndent(value coregraphics.Float)
	SetLineHeightMultiple(value coregraphics.Float)
	SetMaximumLineHeight(value coregraphics.Float)
	SetMinimumLineHeight(value coregraphics.Float)
	SetLineSpacing(value coregraphics.Float)
	SetParagraphSpacing(value coregraphics.Float)
	SetParagraphSpacingBefore(value coregraphics.Float)
	SetBaseWritingDirection(value WritingDirection)
	SetTabStops(value []TextTab)
	SetDefaultTabInterval(value coregraphics.Float)
	SetTextBlocks(value []TextBlock)
	SetTextLists(value []TextList)
	SetLineBreakMode(value LineBreakMode)
	SetLineBreakStrategy(value LineBreakStrategy)
	SetHyphenationFactor(value float32)
	SetTighteningFactorForTruncation(value float32)
	SetAllowsDefaultTighteningForTruncation(value bool)
	SetHeaderLevel(value int)
}

type NSMutableParagraphStyle struct {
	NSParagraphStyle
}

func MakeMutableParagraphStyle(ptr unsafe.Pointer) NSMutableParagraphStyle {
	return NSMutableParagraphStyle{
		NSParagraphStyle: MakeParagraphStyle(ptr),
	}
}

func AllocMutableParagraphStyle() NSMutableParagraphStyle {
	result_ := C.C_NSMutableParagraphStyle_AllocMutableParagraphStyle()
	return MakeMutableParagraphStyle(result_)
}

func (n NSMutableParagraphStyle) Init() NSMutableParagraphStyle {
	result_ := C.C_NSMutableParagraphStyle_Init(n.Ptr())
	return MakeMutableParagraphStyle(result_)
}

func NewMutableParagraphStyle() NSMutableParagraphStyle {
	result_ := C.C_NSMutableParagraphStyle_NewMutableParagraphStyle()
	return MakeMutableParagraphStyle(result_)
}

func (n NSMutableParagraphStyle) Autorelease() NSMutableParagraphStyle {
	result_ := C.C_NSMutableParagraphStyle_Autorelease(n.Ptr())
	return MakeMutableParagraphStyle(result_)
}

func (n NSMutableParagraphStyle) Retain() NSMutableParagraphStyle {
	result_ := C.C_NSMutableParagraphStyle_Retain(n.Ptr())
	return MakeMutableParagraphStyle(result_)
}

func (n NSMutableParagraphStyle) SetParagraphStyle(obj ParagraphStyle) {
	C.C_NSMutableParagraphStyle_SetParagraphStyle(n.Ptr(), objc.ExtractPtr(obj))
}

func (n NSMutableParagraphStyle) AddTabStop(anObject TextTab) {
	C.C_NSMutableParagraphStyle_AddTabStop(n.Ptr(), objc.ExtractPtr(anObject))
}

func (n NSMutableParagraphStyle) RemoveTabStop(anObject TextTab) {
	C.C_NSMutableParagraphStyle_RemoveTabStop(n.Ptr(), objc.ExtractPtr(anObject))
}

func (n NSMutableParagraphStyle) SetAlignment(value TextAlignment) {
	C.C_NSMutableParagraphStyle_SetAlignment(n.Ptr(), C.int(int(value)))
}

func (n NSMutableParagraphStyle) SetFirstLineHeadIndent(value coregraphics.Float) {
	C.C_NSMutableParagraphStyle_SetFirstLineHeadIndent(n.Ptr(), C.double(float64(value)))
}

func (n NSMutableParagraphStyle) SetHeadIndent(value coregraphics.Float) {
	C.C_NSMutableParagraphStyle_SetHeadIndent(n.Ptr(), C.double(float64(value)))
}

func (n NSMutableParagraphStyle) SetTailIndent(value coregraphics.Float) {
	C.C_NSMutableParagraphStyle_SetTailIndent(n.Ptr(), C.double(float64(value)))
}

func (n NSMutableParagraphStyle) SetLineHeightMultiple(value coregraphics.Float) {
	C.C_NSMutableParagraphStyle_SetLineHeightMultiple(n.Ptr(), C.double(float64(value)))
}

func (n NSMutableParagraphStyle) SetMaximumLineHeight(value coregraphics.Float) {
	C.C_NSMutableParagraphStyle_SetMaximumLineHeight(n.Ptr(), C.double(float64(value)))
}

func (n NSMutableParagraphStyle) SetMinimumLineHeight(value coregraphics.Float) {
	C.C_NSMutableParagraphStyle_SetMinimumLineHeight(n.Ptr(), C.double(float64(value)))
}

func (n NSMutableParagraphStyle) SetLineSpacing(value coregraphics.Float) {
	C.C_NSMutableParagraphStyle_SetLineSpacing(n.Ptr(), C.double(float64(value)))
}

func (n NSMutableParagraphStyle) SetParagraphSpacing(value coregraphics.Float) {
	C.C_NSMutableParagraphStyle_SetParagraphSpacing(n.Ptr(), C.double(float64(value)))
}

func (n NSMutableParagraphStyle) SetParagraphSpacingBefore(value coregraphics.Float) {
	C.C_NSMutableParagraphStyle_SetParagraphSpacingBefore(n.Ptr(), C.double(float64(value)))
}

func (n NSMutableParagraphStyle) SetBaseWritingDirection(value WritingDirection) {
	C.C_NSMutableParagraphStyle_SetBaseWritingDirection(n.Ptr(), C.int(int(value)))
}

func (n NSMutableParagraphStyle) SetTabStops(value []TextTab) {
	var cValue C.Array
	if len(value) > 0 {
		cValueData := make([]unsafe.Pointer, len(value))
		for idx, v := range value {
			cValueData[idx] = objc.ExtractPtr(v)
		}
		cValue.data = unsafe.Pointer(&cValueData[0])
		cValue.len = C.int(len(value))
	}
	C.C_NSMutableParagraphStyle_SetTabStops(n.Ptr(), cValue)
}

func (n NSMutableParagraphStyle) SetDefaultTabInterval(value coregraphics.Float) {
	C.C_NSMutableParagraphStyle_SetDefaultTabInterval(n.Ptr(), C.double(float64(value)))
}

func (n NSMutableParagraphStyle) SetTextBlocks(value []TextBlock) {
	var cValue C.Array
	if len(value) > 0 {
		cValueData := make([]unsafe.Pointer, len(value))
		for idx, v := range value {
			cValueData[idx] = objc.ExtractPtr(v)
		}
		cValue.data = unsafe.Pointer(&cValueData[0])
		cValue.len = C.int(len(value))
	}
	C.C_NSMutableParagraphStyle_SetTextBlocks(n.Ptr(), cValue)
}

func (n NSMutableParagraphStyle) SetTextLists(value []TextList) {
	var cValue C.Array
	if len(value) > 0 {
		cValueData := make([]unsafe.Pointer, len(value))
		for idx, v := range value {
			cValueData[idx] = objc.ExtractPtr(v)
		}
		cValue.data = unsafe.Pointer(&cValueData[0])
		cValue.len = C.int(len(value))
	}
	C.C_NSMutableParagraphStyle_SetTextLists(n.Ptr(), cValue)
}

func (n NSMutableParagraphStyle) SetLineBreakMode(value LineBreakMode) {
	C.C_NSMutableParagraphStyle_SetLineBreakMode(n.Ptr(), C.uint(uint(value)))
}

func (n NSMutableParagraphStyle) SetLineBreakStrategy(value LineBreakStrategy) {
	C.C_NSMutableParagraphStyle_SetLineBreakStrategy(n.Ptr(), C.uint(uint(value)))
}

func (n NSMutableParagraphStyle) SetHyphenationFactor(value float32) {
	C.C_NSMutableParagraphStyle_SetHyphenationFactor(n.Ptr(), C.float(value))
}

func (n NSMutableParagraphStyle) SetTighteningFactorForTruncation(value float32) {
	C.C_NSMutableParagraphStyle_SetTighteningFactorForTruncation(n.Ptr(), C.float(value))
}

func (n NSMutableParagraphStyle) SetAllowsDefaultTighteningForTruncation(value bool) {
	C.C_NSMutableParagraphStyle_SetAllowsDefaultTighteningForTruncation(n.Ptr(), C.bool(value))
}

func (n NSMutableParagraphStyle) SetHeaderLevel(value int) {
	C.C_NSMutableParagraphStyle_SetHeaderLevel(n.Ptr(), C.int(value))
}

type TextTab interface {
	objc.Object
	Location() coregraphics.Float
	Alignment() TextAlignment
	Options() map[TextTabOptionKey]objc.Object
}

type NSTextTab struct {
	objc.NSObject
}

func MakeTextTab(ptr unsafe.Pointer) NSTextTab {
	return NSTextTab{
		NSObject: objc.MakeObject(ptr),
	}
}

func (n NSTextTab) InitWithTextAlignment_Location_Options(alignment TextAlignment, loc coregraphics.Float, options map[TextTabOptionKey]objc.Object) NSTextTab {
	var cOptions C.Dictionary
	if len(options) > 0 {
		cOptionsKeyData := make([]unsafe.Pointer, len(options))
		cOptionsValueData := make([]unsafe.Pointer, len(options))
		var idx = 0
		for k, v := range options {
			cOptionsKeyData[idx] = foundation.NewString(string(k)).Ptr()
			cOptionsValueData[idx] = objc.ExtractPtr(v)
			idx++
		}
		cOptions.key_data = unsafe.Pointer(&cOptionsKeyData[0])
		cOptions.value_data = unsafe.Pointer(&cOptionsValueData[0])
		cOptions.len = C.int(len(options))
	}
	result_ := C.C_NSTextTab_InitWithTextAlignment_Location_Options(n.Ptr(), C.int(int(alignment)), C.double(float64(loc)), cOptions)
	return MakeTextTab(result_)
}

func AllocTextTab() NSTextTab {
	result_ := C.C_NSTextTab_AllocTextTab()
	return MakeTextTab(result_)
}

func (n NSTextTab) Init() NSTextTab {
	result_ := C.C_NSTextTab_Init(n.Ptr())
	return MakeTextTab(result_)
}

func NewTextTab() NSTextTab {
	result_ := C.C_NSTextTab_NewTextTab()
	return MakeTextTab(result_)
}

func (n NSTextTab) Autorelease() NSTextTab {
	result_ := C.C_NSTextTab_Autorelease(n.Ptr())
	return MakeTextTab(result_)
}

func (n NSTextTab) Retain() NSTextTab {
	result_ := C.C_NSTextTab_Retain(n.Ptr())
	return MakeTextTab(result_)
}

func TextTab_ColumnTerminatorsForLocale(aLocale foundation.Locale) foundation.CharacterSet {
	result_ := C.C_NSTextTab_TextTab_ColumnTerminatorsForLocale(objc.ExtractPtr(aLocale))
	return foundation.MakeCharacterSet(result_)
}

func (n NSTextTab) Location() coregraphics.Float {
	result_ := C.C_NSTextTab_Location(n.Ptr())
	return coregraphics.Float(float64(result_))
}

func (n NSTextTab) Alignment() TextAlignment {
	result_ := C.C_NSTextTab_Alignment(n.Ptr())
	return TextAlignment(int(result_))
}

func (n NSTextTab) Options() map[TextTabOptionKey]objc.Object {
	result_ := C.C_NSTextTab_Options(n.Ptr())
	if result_.len > 0 {
		defer C.free(result_.key_data)
		defer C.free(result_.value_data)
	}
	result_KeySlice := unsafe.Slice((*unsafe.Pointer)(result_.key_data), int(result_.len))
	result_ValueSlice := unsafe.Slice((*unsafe.Pointer)(result_.value_data), int(result_.len))
	var goResult_ = make(map[TextTabOptionKey]objc.Object)
	for idx, k := range result_KeySlice {
		v := result_ValueSlice[idx]
		goResult_[TextTabOptionKey(foundation.MakeString(k).String())] = objc.MakeObject(v)
	}
	return goResult_
}
