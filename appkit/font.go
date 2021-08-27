package appkit

// #include "font.h"
import "C"
import (
	"github.com/hsiafan/cocoa/coregraphics"
	"github.com/hsiafan/cocoa/foundation"
	"github.com/hsiafan/cocoa/objc"
	"unsafe"
)

type Font interface {
	objc.Object
	Set()
	SetInContext(graphicsContext GraphicsContext)
	FontWithSize(fontSize coregraphics.Float) Font
	PointSize() coregraphics.Float
	CoveredCharacterSet() foundation.CharacterSet
	FontDescriptor() FontDescriptor
	IsFixedPitch() bool
	MostCompatibleStringEncoding() foundation.StringEncoding
	NumberOfGlyphs() uint
	DisplayName() string
	FamilyName() string
	FontName() string
	IsVertical() bool
	VerticalFont() Font
	Ascender() coregraphics.Float
	Descender() coregraphics.Float
	CapHeight() coregraphics.Float
	Leading() coregraphics.Float
	XHeight() coregraphics.Float
	ItalicAngle() coregraphics.Float
	UnderlinePosition() coregraphics.Float
	UnderlineThickness() coregraphics.Float
	BoundingRectForFont() foundation.Rect
	MaximumAdvancement() foundation.Size
	TextTransform() foundation.AffineTransform
}

type NSFont struct {
	objc.NSObject
}

func MakeFont(ptr unsafe.Pointer) NSFont {
	return NSFont{
		NSObject: objc.MakeObject(ptr),
	}
}

func AllocFont() NSFont {
	return MakeFont(C.C_Font_Alloc())
}

func (n NSFont) Init() Font {
	result_ := C.C_NSFont_Init(n.Ptr())
	return MakeFont(result_)
}

func FontWithName_Size(fontName string, fontSize coregraphics.Float) Font {
	result_ := C.C_NSFont_FontWithName_Size(foundation.NewString(fontName).Ptr(), C.double(float64(fontSize)))
	return MakeFont(result_)
}

func FontWithDescriptor_Size(fontDescriptor FontDescriptor, fontSize coregraphics.Float) Font {
	result_ := C.C_NSFont_FontWithDescriptor_Size(objc.ExtractPtr(fontDescriptor), C.double(float64(fontSize)))
	return MakeFont(result_)
}

func FontWithDescriptor_TextTransform(fontDescriptor FontDescriptor, textTransform foundation.AffineTransform) Font {
	result_ := C.C_NSFont_FontWithDescriptor_TextTransform(objc.ExtractPtr(fontDescriptor), objc.ExtractPtr(textTransform))
	return MakeFont(result_)
}

func UserFontOfSize(fontSize coregraphics.Float) Font {
	result_ := C.C_NSFont_UserFontOfSize(C.double(float64(fontSize)))
	return MakeFont(result_)
}

func UserFixedPitchFontOfSize(fontSize coregraphics.Float) Font {
	result_ := C.C_NSFont_UserFixedPitchFontOfSize(C.double(float64(fontSize)))
	return MakeFont(result_)
}

func PreferredFontForTextStyle_Options(style FontTextStyle, options map[FontTextStyleOptionKey]objc.Object) Font {
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
	result_ := C.C_NSFont_PreferredFontForTextStyle_Options(foundation.NewString(string(style)).Ptr(), cOptions)
	return MakeFont(result_)
}

func SystemFontOfSize(fontSize coregraphics.Float) Font {
	result_ := C.C_NSFont_SystemFontOfSize(C.double(float64(fontSize)))
	return MakeFont(result_)
}

func SystemFontOfSize_Weight(fontSize coregraphics.Float, weight FontWeight) Font {
	result_ := C.C_NSFont_SystemFontOfSize_Weight(C.double(float64(fontSize)), C.double(float64(coregraphics.Float(weight))))
	return MakeFont(result_)
}

func BoldSystemFontOfSize(fontSize coregraphics.Float) Font {
	result_ := C.C_NSFont_BoldSystemFontOfSize(C.double(float64(fontSize)))
	return MakeFont(result_)
}

func MonospacedSystemFontOfSize_Weight(fontSize coregraphics.Float, weight FontWeight) Font {
	result_ := C.C_NSFont_MonospacedSystemFontOfSize_Weight(C.double(float64(fontSize)), C.double(float64(coregraphics.Float(weight))))
	return MakeFont(result_)
}

func MonospacedDigitSystemFontOfSize_Weight(fontSize coregraphics.Float, weight FontWeight) Font {
	result_ := C.C_NSFont_MonospacedDigitSystemFontOfSize_Weight(C.double(float64(fontSize)), C.double(float64(coregraphics.Float(weight))))
	return MakeFont(result_)
}

func LabelFontOfSize(fontSize coregraphics.Float) Font {
	result_ := C.C_NSFont_LabelFontOfSize(C.double(float64(fontSize)))
	return MakeFont(result_)
}

func MessageFontOfSize(fontSize coregraphics.Float) Font {
	result_ := C.C_NSFont_MessageFontOfSize(C.double(float64(fontSize)))
	return MakeFont(result_)
}

func MenuBarFontOfSize(fontSize coregraphics.Float) Font {
	result_ := C.C_NSFont_MenuBarFontOfSize(C.double(float64(fontSize)))
	return MakeFont(result_)
}

func MenuFontOfSize(fontSize coregraphics.Float) Font {
	result_ := C.C_NSFont_MenuFontOfSize(C.double(float64(fontSize)))
	return MakeFont(result_)
}

func ControlContentFontOfSize(fontSize coregraphics.Float) Font {
	result_ := C.C_NSFont_ControlContentFontOfSize(C.double(float64(fontSize)))
	return MakeFont(result_)
}

func TitleBarFontOfSize(fontSize coregraphics.Float) Font {
	result_ := C.C_NSFont_TitleBarFontOfSize(C.double(float64(fontSize)))
	return MakeFont(result_)
}

func PaletteFontOfSize(fontSize coregraphics.Float) Font {
	result_ := C.C_NSFont_PaletteFontOfSize(C.double(float64(fontSize)))
	return MakeFont(result_)
}

func ToolTipsFontOfSize(fontSize coregraphics.Float) Font {
	result_ := C.C_NSFont_ToolTipsFontOfSize(C.double(float64(fontSize)))
	return MakeFont(result_)
}

func SystemFontSizeForControlSize(controlSize ControlSize) coregraphics.Float {
	result_ := C.C_NSFont_SystemFontSizeForControlSize(C.uint(uint(controlSize)))
	return coregraphics.Float(float64(result_))
}

func (n NSFont) Set() {
	C.C_NSFont_Set(n.Ptr())
}

func (n NSFont) SetInContext(graphicsContext GraphicsContext) {
	C.C_NSFont_SetInContext(n.Ptr(), objc.ExtractPtr(graphicsContext))
}

func SetUserFont(font Font) {
	C.C_NSFont_SetUserFont(objc.ExtractPtr(font))
}

func SetUserFixedPitchFont(font Font) {
	C.C_NSFont_SetUserFixedPitchFont(objc.ExtractPtr(font))
}

func (n NSFont) FontWithSize(fontSize coregraphics.Float) Font {
	result_ := C.C_NSFont_FontWithSize(n.Ptr(), C.double(float64(fontSize)))
	return MakeFont(result_)
}

func SystemFontSize() coregraphics.Float {
	result_ := C.C_NSFont_SystemFontSize()
	return coregraphics.Float(float64(result_))
}

func SmallSystemFontSize() coregraphics.Float {
	result_ := C.C_NSFont_SmallSystemFontSize()
	return coregraphics.Float(float64(result_))
}

func LabelFontSize() coregraphics.Float {
	result_ := C.C_NSFont_LabelFontSize()
	return coregraphics.Float(float64(result_))
}

func (n NSFont) PointSize() coregraphics.Float {
	result_ := C.C_NSFont_PointSize(n.Ptr())
	return coregraphics.Float(float64(result_))
}

func (n NSFont) CoveredCharacterSet() foundation.CharacterSet {
	result_ := C.C_NSFont_CoveredCharacterSet(n.Ptr())
	return foundation.MakeCharacterSet(result_)
}

func (n NSFont) FontDescriptor() FontDescriptor {
	result_ := C.C_NSFont_FontDescriptor(n.Ptr())
	return MakeFontDescriptor(result_)
}

func (n NSFont) IsFixedPitch() bool {
	result_ := C.C_NSFont_IsFixedPitch(n.Ptr())
	return bool(result_)
}

func (n NSFont) MostCompatibleStringEncoding() foundation.StringEncoding {
	result_ := C.C_NSFont_MostCompatibleStringEncoding(n.Ptr())
	return foundation.StringEncoding(uint(result_))
}

func (n NSFont) NumberOfGlyphs() uint {
	result_ := C.C_NSFont_NumberOfGlyphs(n.Ptr())
	return uint(result_)
}

func (n NSFont) DisplayName() string {
	result_ := C.C_NSFont_DisplayName(n.Ptr())
	return foundation.MakeString(result_).String()
}

func (n NSFont) FamilyName() string {
	result_ := C.C_NSFont_FamilyName(n.Ptr())
	return foundation.MakeString(result_).String()
}

func (n NSFont) FontName() string {
	result_ := C.C_NSFont_FontName(n.Ptr())
	return foundation.MakeString(result_).String()
}

func (n NSFont) IsVertical() bool {
	result_ := C.C_NSFont_IsVertical(n.Ptr())
	return bool(result_)
}

func (n NSFont) VerticalFont() Font {
	result_ := C.C_NSFont_VerticalFont(n.Ptr())
	return MakeFont(result_)
}

func (n NSFont) Ascender() coregraphics.Float {
	result_ := C.C_NSFont_Ascender(n.Ptr())
	return coregraphics.Float(float64(result_))
}

func (n NSFont) Descender() coregraphics.Float {
	result_ := C.C_NSFont_Descender(n.Ptr())
	return coregraphics.Float(float64(result_))
}

func (n NSFont) CapHeight() coregraphics.Float {
	result_ := C.C_NSFont_CapHeight(n.Ptr())
	return coregraphics.Float(float64(result_))
}

func (n NSFont) Leading() coregraphics.Float {
	result_ := C.C_NSFont_Leading(n.Ptr())
	return coregraphics.Float(float64(result_))
}

func (n NSFont) XHeight() coregraphics.Float {
	result_ := C.C_NSFont_XHeight(n.Ptr())
	return coregraphics.Float(float64(result_))
}

func (n NSFont) ItalicAngle() coregraphics.Float {
	result_ := C.C_NSFont_ItalicAngle(n.Ptr())
	return coregraphics.Float(float64(result_))
}

func (n NSFont) UnderlinePosition() coregraphics.Float {
	result_ := C.C_NSFont_UnderlinePosition(n.Ptr())
	return coregraphics.Float(float64(result_))
}

func (n NSFont) UnderlineThickness() coregraphics.Float {
	result_ := C.C_NSFont_UnderlineThickness(n.Ptr())
	return coregraphics.Float(float64(result_))
}

func (n NSFont) BoundingRectForFont() foundation.Rect {
	result_ := C.C_NSFont_BoundingRectForFont(n.Ptr())
	return foundation.Rect(coregraphics.FromCGRectPointer(unsafe.Pointer(&result_)))
}

func (n NSFont) MaximumAdvancement() foundation.Size {
	result_ := C.C_NSFont_MaximumAdvancement(n.Ptr())
	return foundation.Size(coregraphics.FromCGSizePointer(unsafe.Pointer(&result_)))
}

func (n NSFont) TextTransform() foundation.AffineTransform {
	result_ := C.C_NSFont_TextTransform(n.Ptr())
	return foundation.MakeAffineTransform(result_)
}
