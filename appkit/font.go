package appkit

// #cgo CFLAGS: -x objective-c
// #cgo LDFLAGS: -framework AppKit
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

func MakeFont(ptr unsafe.Pointer) *NSFont {
	if ptr == nil {
		return nil
	}
	return &NSFont{
		NSObject: *objc.MakeObject(ptr),
	}
}

func AllocFont() *NSFont {
	return MakeFont(C.C_Font_Alloc())
}

func (n *NSFont) Init() Font {
	result := C.C_NSFont_Init(n.Ptr())
	return MakeFont(result)
}

func FontWithName_Size(fontName string, fontSize coregraphics.Float) Font {
	result := C.C_NSFont_FontWithName_Size(foundation.NewString(fontName).Ptr(), C.double(float64(fontSize)))
	return MakeFont(result)
}

func FontWithDescriptor_Size(fontDescriptor FontDescriptor, fontSize coregraphics.Float) Font {
	result := C.C_NSFont_FontWithDescriptor_Size(objc.ExtractPtr(fontDescriptor), C.double(float64(fontSize)))
	return MakeFont(result)
}

func FontWithDescriptor_TextTransform(fontDescriptor FontDescriptor, textTransform foundation.AffineTransform) Font {
	result := C.C_NSFont_FontWithDescriptor_TextTransform(objc.ExtractPtr(fontDescriptor), objc.ExtractPtr(textTransform))
	return MakeFont(result)
}

func Font_UserFontOfSize(fontSize coregraphics.Float) Font {
	result := C.C_NSFont_Font_UserFontOfSize(C.double(float64(fontSize)))
	return MakeFont(result)
}

func Font_UserFixedPitchFontOfSize(fontSize coregraphics.Float) Font {
	result := C.C_NSFont_Font_UserFixedPitchFontOfSize(C.double(float64(fontSize)))
	return MakeFont(result)
}

func Font_SystemFontOfSize(fontSize coregraphics.Float) Font {
	result := C.C_NSFont_Font_SystemFontOfSize(C.double(float64(fontSize)))
	return MakeFont(result)
}

func Font_SystemFontOfSize_Weight(fontSize coregraphics.Float, weight FontWeight) Font {
	result := C.C_NSFont_Font_SystemFontOfSize_Weight(C.double(float64(fontSize)), C.double(float64(coregraphics.Float(weight))))
	return MakeFont(result)
}

func Font_BoldSystemFontOfSize(fontSize coregraphics.Float) Font {
	result := C.C_NSFont_Font_BoldSystemFontOfSize(C.double(float64(fontSize)))
	return MakeFont(result)
}

func Font_MonospacedSystemFontOfSize_Weight(fontSize coregraphics.Float, weight FontWeight) Font {
	result := C.C_NSFont_Font_MonospacedSystemFontOfSize_Weight(C.double(float64(fontSize)), C.double(float64(coregraphics.Float(weight))))
	return MakeFont(result)
}

func Font_MonospacedDigitSystemFontOfSize_Weight(fontSize coregraphics.Float, weight FontWeight) Font {
	result := C.C_NSFont_Font_MonospacedDigitSystemFontOfSize_Weight(C.double(float64(fontSize)), C.double(float64(coregraphics.Float(weight))))
	return MakeFont(result)
}

func Font_LabelFontOfSize(fontSize coregraphics.Float) Font {
	result := C.C_NSFont_Font_LabelFontOfSize(C.double(float64(fontSize)))
	return MakeFont(result)
}

func Font_MessageFontOfSize(fontSize coregraphics.Float) Font {
	result := C.C_NSFont_Font_MessageFontOfSize(C.double(float64(fontSize)))
	return MakeFont(result)
}

func Font_MenuBarFontOfSize(fontSize coregraphics.Float) Font {
	result := C.C_NSFont_Font_MenuBarFontOfSize(C.double(float64(fontSize)))
	return MakeFont(result)
}

func Font_MenuFontOfSize(fontSize coregraphics.Float) Font {
	result := C.C_NSFont_Font_MenuFontOfSize(C.double(float64(fontSize)))
	return MakeFont(result)
}

func Font_ControlContentFontOfSize(fontSize coregraphics.Float) Font {
	result := C.C_NSFont_Font_ControlContentFontOfSize(C.double(float64(fontSize)))
	return MakeFont(result)
}

func Font_TitleBarFontOfSize(fontSize coregraphics.Float) Font {
	result := C.C_NSFont_Font_TitleBarFontOfSize(C.double(float64(fontSize)))
	return MakeFont(result)
}

func Font_PaletteFontOfSize(fontSize coregraphics.Float) Font {
	result := C.C_NSFont_Font_PaletteFontOfSize(C.double(float64(fontSize)))
	return MakeFont(result)
}

func Font_ToolTipsFontOfSize(fontSize coregraphics.Float) Font {
	result := C.C_NSFont_Font_ToolTipsFontOfSize(C.double(float64(fontSize)))
	return MakeFont(result)
}

func Font_SystemFontSizeForControlSize(controlSize ControlSize) coregraphics.Float {
	result := C.C_NSFont_Font_SystemFontSizeForControlSize(C.uint(uint(controlSize)))
	return coregraphics.Float(float64(result))
}

func (n *NSFont) Set() {
	C.C_NSFont_Set(n.Ptr())
}

func (n *NSFont) SetInContext(graphicsContext GraphicsContext) {
	C.C_NSFont_SetInContext(n.Ptr(), objc.ExtractPtr(graphicsContext))
}

func SetUserFont(font Font) {
	C.C_NSFont_SetUserFont(objc.ExtractPtr(font))
}

func SetUserFixedPitchFont(font Font) {
	C.C_NSFont_SetUserFixedPitchFont(objc.ExtractPtr(font))
}

func (n *NSFont) FontWithSize(fontSize coregraphics.Float) Font {
	result := C.C_NSFont_FontWithSize(n.Ptr(), C.double(float64(fontSize)))
	return MakeFont(result)
}

func Font_SystemFontSize() coregraphics.Float {
	result := C.C_NSFont_Font_SystemFontSize()
	return coregraphics.Float(float64(result))
}

func Font_SmallSystemFontSize() coregraphics.Float {
	result := C.C_NSFont_Font_SmallSystemFontSize()
	return coregraphics.Float(float64(result))
}

func Font_LabelFontSize() coregraphics.Float {
	result := C.C_NSFont_Font_LabelFontSize()
	return coregraphics.Float(float64(result))
}

func (n *NSFont) PointSize() coregraphics.Float {
	result := C.C_NSFont_PointSize(n.Ptr())
	return coregraphics.Float(float64(result))
}

func (n *NSFont) CoveredCharacterSet() foundation.CharacterSet {
	result := C.C_NSFont_CoveredCharacterSet(n.Ptr())
	return foundation.MakeCharacterSet(result)
}

func (n *NSFont) FontDescriptor() FontDescriptor {
	result := C.C_NSFont_FontDescriptor(n.Ptr())
	return MakeFontDescriptor(result)
}

func (n *NSFont) IsFixedPitch() bool {
	result := C.C_NSFont_IsFixedPitch(n.Ptr())
	return bool(result)
}

func (n *NSFont) MostCompatibleStringEncoding() foundation.StringEncoding {
	result := C.C_NSFont_MostCompatibleStringEncoding(n.Ptr())
	return foundation.StringEncoding(uint(result))
}

func (n *NSFont) NumberOfGlyphs() uint {
	result := C.C_NSFont_NumberOfGlyphs(n.Ptr())
	return uint(result)
}

func (n *NSFont) DisplayName() string {
	result := C.C_NSFont_DisplayName(n.Ptr())
	return foundation.MakeString(result).String()
}

func (n *NSFont) FamilyName() string {
	result := C.C_NSFont_FamilyName(n.Ptr())
	return foundation.MakeString(result).String()
}

func (n *NSFont) FontName() string {
	result := C.C_NSFont_FontName(n.Ptr())
	return foundation.MakeString(result).String()
}

func (n *NSFont) IsVertical() bool {
	result := C.C_NSFont_IsVertical(n.Ptr())
	return bool(result)
}

func (n *NSFont) VerticalFont() Font {
	result := C.C_NSFont_VerticalFont(n.Ptr())
	return MakeFont(result)
}

func (n *NSFont) Ascender() coregraphics.Float {
	result := C.C_NSFont_Ascender(n.Ptr())
	return coregraphics.Float(float64(result))
}

func (n *NSFont) Descender() coregraphics.Float {
	result := C.C_NSFont_Descender(n.Ptr())
	return coregraphics.Float(float64(result))
}

func (n *NSFont) CapHeight() coregraphics.Float {
	result := C.C_NSFont_CapHeight(n.Ptr())
	return coregraphics.Float(float64(result))
}

func (n *NSFont) Leading() coregraphics.Float {
	result := C.C_NSFont_Leading(n.Ptr())
	return coregraphics.Float(float64(result))
}

func (n *NSFont) XHeight() coregraphics.Float {
	result := C.C_NSFont_XHeight(n.Ptr())
	return coregraphics.Float(float64(result))
}

func (n *NSFont) ItalicAngle() coregraphics.Float {
	result := C.C_NSFont_ItalicAngle(n.Ptr())
	return coregraphics.Float(float64(result))
}

func (n *NSFont) UnderlinePosition() coregraphics.Float {
	result := C.C_NSFont_UnderlinePosition(n.Ptr())
	return coregraphics.Float(float64(result))
}

func (n *NSFont) UnderlineThickness() coregraphics.Float {
	result := C.C_NSFont_UnderlineThickness(n.Ptr())
	return coregraphics.Float(float64(result))
}

func (n *NSFont) BoundingRectForFont() foundation.Rect {
	result := C.C_NSFont_BoundingRectForFont(n.Ptr())
	return foundation.Rect(coregraphics.FromCGRectPointer(unsafe.Pointer(&result)))
}

func (n *NSFont) MaximumAdvancement() foundation.Size {
	result := C.C_NSFont_MaximumAdvancement(n.Ptr())
	return foundation.Size(coregraphics.FromCGSizePointer(unsafe.Pointer(&result)))
}

func (n *NSFont) TextTransform() foundation.AffineTransform {
	result := C.C_NSFont_TextTransform(n.Ptr())
	return foundation.MakeAffineTransform(result)
}
