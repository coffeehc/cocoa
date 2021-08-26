package appkit

// #include "color.h"
import "C"
import (
	"github.com/hsiafan/cocoa/coregraphics"
	"github.com/hsiafan/cocoa/foundation"
	"github.com/hsiafan/cocoa/objc"
	"unsafe"
)

type Color interface {
	objc.Object
	ColorWithSystemEffect(systemEffect ColorSystemEffect) Color
	ColorUsingColorSpace(space ColorSpace) Color
	BlendedColorWithFraction_OfColor(fraction coregraphics.Float, color Color) Color
	ColorWithAlphaComponent(alpha coregraphics.Float) Color
	HighlightWithLevel(val coregraphics.Float) Color
	ShadowWithLevel(val coregraphics.Float) Color
	WriteToPasteboard(pasteBoard Pasteboard)
	ColorUsingType(_type ColorType) Color
	DrawSwatchInRect(rect foundation.Rect)
	Set()
	SetFill()
	SetStroke()
	NumberOfComponents() int
	AlphaComponent() coregraphics.Float
	WhiteComponent() coregraphics.Float
	RedComponent() coregraphics.Float
	GreenComponent() coregraphics.Float
	BlueComponent() coregraphics.Float
	CyanComponent() coregraphics.Float
	MagentaComponent() coregraphics.Float
	YellowComponent() coregraphics.Float
	BlackComponent() coregraphics.Float
	HueComponent() coregraphics.Float
	SaturationComponent() coregraphics.Float
	BrightnessComponent() coregraphics.Float
	CatalogNameComponent() ColorListName
	LocalizedCatalogNameComponent() string
	ColorNameComponent() ColorName
	LocalizedColorNameComponent() string
	Type() ColorType
	ColorSpace() ColorSpace
	CGColor() coregraphics.ColorRef
	PatternImage() Image
}

type NSColor struct {
	objc.NSObject
}

func MakeColor(ptr unsafe.Pointer) NSColor {
	return NSColor{
		NSObject: objc.MakeObject(ptr),
	}
}

func AllocColor() NSColor {
	return MakeColor(C.C_Color_Alloc())
}

func (n NSColor) Init() Color {
	result_ := C.C_NSColor_Init(n.Ptr())
	return MakeColor(result_)
}

func (n NSColor) InitWithCoder(coder foundation.Coder) Color {
	result_ := C.C_NSColor_InitWithCoder(n.Ptr(), objc.ExtractPtr(coder))
	return MakeColor(result_)
}

func (n NSColor) ColorWithSystemEffect(systemEffect ColorSystemEffect) Color {
	result_ := C.C_NSColor_ColorWithSystemEffect(n.Ptr(), C.int(int(systemEffect)))
	return MakeColor(result_)
}

func (n NSColor) ColorUsingColorSpace(space ColorSpace) Color {
	result_ := C.C_NSColor_ColorUsingColorSpace(n.Ptr(), objc.ExtractPtr(space))
	return MakeColor(result_)
}

func (n NSColor) BlendedColorWithFraction_OfColor(fraction coregraphics.Float, color Color) Color {
	result_ := C.C_NSColor_BlendedColorWithFraction_OfColor(n.Ptr(), C.double(float64(fraction)), objc.ExtractPtr(color))
	return MakeColor(result_)
}

func (n NSColor) ColorWithAlphaComponent(alpha coregraphics.Float) Color {
	result_ := C.C_NSColor_ColorWithAlphaComponent(n.Ptr(), C.double(float64(alpha)))
	return MakeColor(result_)
}

func (n NSColor) HighlightWithLevel(val coregraphics.Float) Color {
	result_ := C.C_NSColor_HighlightWithLevel(n.Ptr(), C.double(float64(val)))
	return MakeColor(result_)
}

func (n NSColor) ShadowWithLevel(val coregraphics.Float) Color {
	result_ := C.C_NSColor_ShadowWithLevel(n.Ptr(), C.double(float64(val)))
	return MakeColor(result_)
}

func ColorFromPasteboard(pasteBoard Pasteboard) Color {
	result_ := C.C_NSColor_ColorFromPasteboard(objc.ExtractPtr(pasteBoard))
	return MakeColor(result_)
}

func (n NSColor) WriteToPasteboard(pasteBoard Pasteboard) {
	C.C_NSColor_WriteToPasteboard(n.Ptr(), objc.ExtractPtr(pasteBoard))
}

func (n NSColor) ColorUsingType(_type ColorType) Color {
	result_ := C.C_NSColor_ColorUsingType(n.Ptr(), C.int(int(_type)))
	return MakeColor(result_)
}

func (n NSColor) DrawSwatchInRect(rect foundation.Rect) {
	C.C_NSColor_DrawSwatchInRect(n.Ptr(), *(*C.CGRect)(coregraphics.ToCGRectPointer(coregraphics.Rect(rect))))
}

func (n NSColor) Set() {
	C.C_NSColor_Set(n.Ptr())
}

func (n NSColor) SetFill() {
	C.C_NSColor_SetFill(n.Ptr())
}

func (n NSColor) SetStroke() {
	C.C_NSColor_SetStroke(n.Ptr())
}

func ColorNamed(name ColorName) Color {
	result_ := C.C_NSColor_ColorNamed(foundation.NewString(string(name)).Ptr())
	return MakeColor(result_)
}

func ColorNamed_Bundle(name ColorName, bundle foundation.Bundle) Color {
	result_ := C.C_NSColor_ColorNamed_Bundle(foundation.NewString(string(name)).Ptr(), objc.ExtractPtr(bundle))
	return MakeColor(result_)
}

func ColorWithCatalogName_ColorName(listName ColorListName, colorName ColorName) Color {
	result_ := C.C_NSColor_ColorWithCatalogName_ColorName(foundation.NewString(string(listName)).Ptr(), foundation.NewString(string(colorName)).Ptr())
	return MakeColor(result_)
}

func ColorWithSRGBRed_Green_Blue_Alpha(red coregraphics.Float, green coregraphics.Float, blue coregraphics.Float, alpha coregraphics.Float) Color {
	result_ := C.C_NSColor_ColorWithSRGBRed_Green_Blue_Alpha(C.double(float64(red)), C.double(float64(green)), C.double(float64(blue)), C.double(float64(alpha)))
	return MakeColor(result_)
}

func ColorWithDisplayP3Red_Green_Blue_Alpha(red coregraphics.Float, green coregraphics.Float, blue coregraphics.Float, alpha coregraphics.Float) Color {
	result_ := C.C_NSColor_ColorWithDisplayP3Red_Green_Blue_Alpha(C.double(float64(red)), C.double(float64(green)), C.double(float64(blue)), C.double(float64(alpha)))
	return MakeColor(result_)
}

func ColorWithRed_Green_Blue_Alpha(red coregraphics.Float, green coregraphics.Float, blue coregraphics.Float, alpha coregraphics.Float) Color {
	result_ := C.C_NSColor_ColorWithRed_Green_Blue_Alpha(C.double(float64(red)), C.double(float64(green)), C.double(float64(blue)), C.double(float64(alpha)))
	return MakeColor(result_)
}

func ColorWithCalibratedRed_Green_Blue_Alpha(red coregraphics.Float, green coregraphics.Float, blue coregraphics.Float, alpha coregraphics.Float) Color {
	result_ := C.C_NSColor_ColorWithCalibratedRed_Green_Blue_Alpha(C.double(float64(red)), C.double(float64(green)), C.double(float64(blue)), C.double(float64(alpha)))
	return MakeColor(result_)
}

func ColorWithDeviceRed_Green_Blue_Alpha(red coregraphics.Float, green coregraphics.Float, blue coregraphics.Float, alpha coregraphics.Float) Color {
	result_ := C.C_NSColor_ColorWithDeviceRed_Green_Blue_Alpha(C.double(float64(red)), C.double(float64(green)), C.double(float64(blue)), C.double(float64(alpha)))
	return MakeColor(result_)
}

func ColorWithCalibratedHue_Saturation_Brightness_Alpha(hue coregraphics.Float, saturation coregraphics.Float, brightness coregraphics.Float, alpha coregraphics.Float) Color {
	result_ := C.C_NSColor_ColorWithCalibratedHue_Saturation_Brightness_Alpha(C.double(float64(hue)), C.double(float64(saturation)), C.double(float64(brightness)), C.double(float64(alpha)))
	return MakeColor(result_)
}

func ColorWithDeviceHue_Saturation_Brightness_Alpha(hue coregraphics.Float, saturation coregraphics.Float, brightness coregraphics.Float, alpha coregraphics.Float) Color {
	result_ := C.C_NSColor_ColorWithDeviceHue_Saturation_Brightness_Alpha(C.double(float64(hue)), C.double(float64(saturation)), C.double(float64(brightness)), C.double(float64(alpha)))
	return MakeColor(result_)
}

func ColorWithHue_Saturation_Brightness_Alpha(hue coregraphics.Float, saturation coregraphics.Float, brightness coregraphics.Float, alpha coregraphics.Float) Color {
	result_ := C.C_NSColor_ColorWithHue_Saturation_Brightness_Alpha(C.double(float64(hue)), C.double(float64(saturation)), C.double(float64(brightness)), C.double(float64(alpha)))
	return MakeColor(result_)
}

func ColorWithColorSpace_Hue_Saturation_Brightness_Alpha(space ColorSpace, hue coregraphics.Float, saturation coregraphics.Float, brightness coregraphics.Float, alpha coregraphics.Float) Color {
	result_ := C.C_NSColor_ColorWithColorSpace_Hue_Saturation_Brightness_Alpha(objc.ExtractPtr(space), C.double(float64(hue)), C.double(float64(saturation)), C.double(float64(brightness)), C.double(float64(alpha)))
	return MakeColor(result_)
}

func ColorWithDeviceCyan_Magenta_Yellow_Black_Alpha(cyan coregraphics.Float, magenta coregraphics.Float, yellow coregraphics.Float, black coregraphics.Float, alpha coregraphics.Float) Color {
	result_ := C.C_NSColor_ColorWithDeviceCyan_Magenta_Yellow_Black_Alpha(C.double(float64(cyan)), C.double(float64(magenta)), C.double(float64(yellow)), C.double(float64(black)), C.double(float64(alpha)))
	return MakeColor(result_)
}

func ColorWithWhite_Alpha(white coregraphics.Float, alpha coregraphics.Float) Color {
	result_ := C.C_NSColor_ColorWithWhite_Alpha(C.double(float64(white)), C.double(float64(alpha)))
	return MakeColor(result_)
}

func ColorWithCalibratedWhite_Alpha(white coregraphics.Float, alpha coregraphics.Float) Color {
	result_ := C.C_NSColor_ColorWithCalibratedWhite_Alpha(C.double(float64(white)), C.double(float64(alpha)))
	return MakeColor(result_)
}

func ColorWithDeviceWhite_Alpha(white coregraphics.Float, alpha coregraphics.Float) Color {
	result_ := C.C_NSColor_ColorWithDeviceWhite_Alpha(C.double(float64(white)), C.double(float64(alpha)))
	return MakeColor(result_)
}

func ColorWithGenericGamma22White_Alpha(white coregraphics.Float, alpha coregraphics.Float) Color {
	result_ := C.C_NSColor_ColorWithGenericGamma22White_Alpha(C.double(float64(white)), C.double(float64(alpha)))
	return MakeColor(result_)
}

func ColorWithPatternImage(image Image) Color {
	result_ := C.C_NSColor_ColorWithPatternImage(objc.ExtractPtr(image))
	return MakeColor(result_)
}

func ColorWithCGColor(cgColor coregraphics.ColorRef) Color {
	result_ := C.C_NSColor_ColorWithCGColor(unsafe.Pointer(cgColor))
	return MakeColor(result_)
}

func Color_IgnoresAlpha() bool {
	result_ := C.C_NSColor_Color_IgnoresAlpha()
	return bool(result_)
}

func Color_SetIgnoresAlpha(value bool) {
	C.C_NSColor_Color_SetIgnoresAlpha(C.bool(value))
}

func (n NSColor) NumberOfComponents() int {
	result_ := C.C_NSColor_NumberOfComponents(n.Ptr())
	return int(result_)
}

func (n NSColor) AlphaComponent() coregraphics.Float {
	result_ := C.C_NSColor_AlphaComponent(n.Ptr())
	return coregraphics.Float(float64(result_))
}

func (n NSColor) WhiteComponent() coregraphics.Float {
	result_ := C.C_NSColor_WhiteComponent(n.Ptr())
	return coregraphics.Float(float64(result_))
}

func (n NSColor) RedComponent() coregraphics.Float {
	result_ := C.C_NSColor_RedComponent(n.Ptr())
	return coregraphics.Float(float64(result_))
}

func (n NSColor) GreenComponent() coregraphics.Float {
	result_ := C.C_NSColor_GreenComponent(n.Ptr())
	return coregraphics.Float(float64(result_))
}

func (n NSColor) BlueComponent() coregraphics.Float {
	result_ := C.C_NSColor_BlueComponent(n.Ptr())
	return coregraphics.Float(float64(result_))
}

func (n NSColor) CyanComponent() coregraphics.Float {
	result_ := C.C_NSColor_CyanComponent(n.Ptr())
	return coregraphics.Float(float64(result_))
}

func (n NSColor) MagentaComponent() coregraphics.Float {
	result_ := C.C_NSColor_MagentaComponent(n.Ptr())
	return coregraphics.Float(float64(result_))
}

func (n NSColor) YellowComponent() coregraphics.Float {
	result_ := C.C_NSColor_YellowComponent(n.Ptr())
	return coregraphics.Float(float64(result_))
}

func (n NSColor) BlackComponent() coregraphics.Float {
	result_ := C.C_NSColor_BlackComponent(n.Ptr())
	return coregraphics.Float(float64(result_))
}

func (n NSColor) HueComponent() coregraphics.Float {
	result_ := C.C_NSColor_HueComponent(n.Ptr())
	return coregraphics.Float(float64(result_))
}

func (n NSColor) SaturationComponent() coregraphics.Float {
	result_ := C.C_NSColor_SaturationComponent(n.Ptr())
	return coregraphics.Float(float64(result_))
}

func (n NSColor) BrightnessComponent() coregraphics.Float {
	result_ := C.C_NSColor_BrightnessComponent(n.Ptr())
	return coregraphics.Float(float64(result_))
}

func (n NSColor) CatalogNameComponent() ColorListName {
	result_ := C.C_NSColor_CatalogNameComponent(n.Ptr())
	return ColorListName(foundation.MakeString(result_).String())
}

func (n NSColor) LocalizedCatalogNameComponent() string {
	result_ := C.C_NSColor_LocalizedCatalogNameComponent(n.Ptr())
	return foundation.MakeString(result_).String()
}

func (n NSColor) ColorNameComponent() ColorName {
	result_ := C.C_NSColor_ColorNameComponent(n.Ptr())
	return ColorName(foundation.MakeString(result_).String())
}

func (n NSColor) LocalizedColorNameComponent() string {
	result_ := C.C_NSColor_LocalizedColorNameComponent(n.Ptr())
	return foundation.MakeString(result_).String()
}

func (n NSColor) Type() ColorType {
	result_ := C.C_NSColor_Type(n.Ptr())
	return ColorType(int(result_))
}

func (n NSColor) ColorSpace() ColorSpace {
	result_ := C.C_NSColor_ColorSpace(n.Ptr())
	return MakeColorSpace(result_)
}

func (n NSColor) CGColor() coregraphics.ColorRef {
	result_ := C.C_NSColor_CGColor(n.Ptr())
	return coregraphics.ColorRef(result_)
}

func LabelColor() Color {
	result_ := C.C_NSColor_LabelColor()
	return MakeColor(result_)
}

func SecondaryLabelColor() Color {
	result_ := C.C_NSColor_SecondaryLabelColor()
	return MakeColor(result_)
}

func TertiaryLabelColor() Color {
	result_ := C.C_NSColor_TertiaryLabelColor()
	return MakeColor(result_)
}

func QuaternaryLabelColor() Color {
	result_ := C.C_NSColor_QuaternaryLabelColor()
	return MakeColor(result_)
}

func TextColor() Color {
	result_ := C.C_NSColor_TextColor()
	return MakeColor(result_)
}

func PlaceholderTextColor() Color {
	result_ := C.C_NSColor_PlaceholderTextColor()
	return MakeColor(result_)
}

func SelectedTextColor() Color {
	result_ := C.C_NSColor_SelectedTextColor()
	return MakeColor(result_)
}

func TextBackgroundColor() Color {
	result_ := C.C_NSColor_TextBackgroundColor()
	return MakeColor(result_)
}

func SelectedTextBackgroundColor() Color {
	result_ := C.C_NSColor_SelectedTextBackgroundColor()
	return MakeColor(result_)
}

func KeyboardFocusIndicatorColor() Color {
	result_ := C.C_NSColor_KeyboardFocusIndicatorColor()
	return MakeColor(result_)
}

func UnemphasizedSelectedTextColor() Color {
	result_ := C.C_NSColor_UnemphasizedSelectedTextColor()
	return MakeColor(result_)
}

func UnemphasizedSelectedTextBackgroundColor() Color {
	result_ := C.C_NSColor_UnemphasizedSelectedTextBackgroundColor()
	return MakeColor(result_)
}

func LinkColor() Color {
	result_ := C.C_NSColor_LinkColor()
	return MakeColor(result_)
}

func SeparatorColor() Color {
	result_ := C.C_NSColor_SeparatorColor()
	return MakeColor(result_)
}

func SelectedContentBackgroundColor() Color {
	result_ := C.C_NSColor_SelectedContentBackgroundColor()
	return MakeColor(result_)
}

func UnemphasizedSelectedContentBackgroundColor() Color {
	result_ := C.C_NSColor_UnemphasizedSelectedContentBackgroundColor()
	return MakeColor(result_)
}

func SelectedMenuItemTextColor() Color {
	result_ := C.C_NSColor_SelectedMenuItemTextColor()
	return MakeColor(result_)
}

func GridColor() Color {
	result_ := C.C_NSColor_GridColor()
	return MakeColor(result_)
}

func HeaderTextColor() Color {
	result_ := C.C_NSColor_HeaderTextColor()
	return MakeColor(result_)
}

func Color_AlternatingContentBackgroundColors() []Color {
	result_ := C.C_NSColor_Color_AlternatingContentBackgroundColors()
	if result_.len > 0 {
		defer C.free(result_.data)
	}
	result_Slice := (*[1 << 28]unsafe.Pointer)(unsafe.Pointer(result_.data))[:result_.len:result_.len]
	var goResult_ = make([]Color, len(result_Slice))
	for idx, r := range result_Slice {
		goResult_[idx] = MakeColor(r)
	}
	return goResult_
}

func ControlAccentColor() Color {
	result_ := C.C_NSColor_ControlAccentColor()
	return MakeColor(result_)
}

func ControlColor() Color {
	result_ := C.C_NSColor_ControlColor()
	return MakeColor(result_)
}

func ControlBackgroundColor() Color {
	result_ := C.C_NSColor_ControlBackgroundColor()
	return MakeColor(result_)
}

func ControlTextColor() Color {
	result_ := C.C_NSColor_ControlTextColor()
	return MakeColor(result_)
}

func DisabledControlTextColor() Color {
	result_ := C.C_NSColor_DisabledControlTextColor()
	return MakeColor(result_)
}

func Color_CurrentControlTint() ControlTint {
	result_ := C.C_NSColor_Color_CurrentControlTint()
	return ControlTint(uint(result_))
}

func SelectedControlColor() Color {
	result_ := C.C_NSColor_SelectedControlColor()
	return MakeColor(result_)
}

func SelectedControlTextColor() Color {
	result_ := C.C_NSColor_SelectedControlTextColor()
	return MakeColor(result_)
}

func AlternateSelectedControlTextColor() Color {
	result_ := C.C_NSColor_AlternateSelectedControlTextColor()
	return MakeColor(result_)
}

func ScrubberTexturedBackgroundColor() Color {
	result_ := C.C_NSColor_ScrubberTexturedBackgroundColor()
	return MakeColor(result_)
}

func WindowBackgroundColor() Color {
	result_ := C.C_NSColor_WindowBackgroundColor()
	return MakeColor(result_)
}

func WindowFrameTextColor() Color {
	result_ := C.C_NSColor_WindowFrameTextColor()
	return MakeColor(result_)
}

func UnderPageBackgroundColor() Color {
	result_ := C.C_NSColor_UnderPageBackgroundColor()
	return MakeColor(result_)
}

func FindHighlightColor() Color {
	result_ := C.C_NSColor_FindHighlightColor()
	return MakeColor(result_)
}

func HighlightColor() Color {
	result_ := C.C_NSColor_HighlightColor()
	return MakeColor(result_)
}

func ShadowColor() Color {
	result_ := C.C_NSColor_ShadowColor()
	return MakeColor(result_)
}

func SystemBlueColor() Color {
	result_ := C.C_NSColor_SystemBlueColor()
	return MakeColor(result_)
}

func SystemBrownColor() Color {
	result_ := C.C_NSColor_SystemBrownColor()
	return MakeColor(result_)
}

func SystemGrayColor() Color {
	result_ := C.C_NSColor_SystemGrayColor()
	return MakeColor(result_)
}

func SystemGreenColor() Color {
	result_ := C.C_NSColor_SystemGreenColor()
	return MakeColor(result_)
}

func SystemIndigoColor() Color {
	result_ := C.C_NSColor_SystemIndigoColor()
	return MakeColor(result_)
}

func SystemOrangeColor() Color {
	result_ := C.C_NSColor_SystemOrangeColor()
	return MakeColor(result_)
}

func SystemPinkColor() Color {
	result_ := C.C_NSColor_SystemPinkColor()
	return MakeColor(result_)
}

func SystemPurpleColor() Color {
	result_ := C.C_NSColor_SystemPurpleColor()
	return MakeColor(result_)
}

func SystemRedColor() Color {
	result_ := C.C_NSColor_SystemRedColor()
	return MakeColor(result_)
}

func SystemTealColor() Color {
	result_ := C.C_NSColor_SystemTealColor()
	return MakeColor(result_)
}

func SystemYellowColor() Color {
	result_ := C.C_NSColor_SystemYellowColor()
	return MakeColor(result_)
}

func ClearColor() Color {
	result_ := C.C_NSColor_ClearColor()
	return MakeColor(result_)
}

func BlackColor() Color {
	result_ := C.C_NSColor_BlackColor()
	return MakeColor(result_)
}

func BlueColor() Color {
	result_ := C.C_NSColor_BlueColor()
	return MakeColor(result_)
}

func BrownColor() Color {
	result_ := C.C_NSColor_BrownColor()
	return MakeColor(result_)
}

func CyanColor() Color {
	result_ := C.C_NSColor_CyanColor()
	return MakeColor(result_)
}

func DarkGrayColor() Color {
	result_ := C.C_NSColor_DarkGrayColor()
	return MakeColor(result_)
}

func GrayColor() Color {
	result_ := C.C_NSColor_GrayColor()
	return MakeColor(result_)
}

func GreenColor() Color {
	result_ := C.C_NSColor_GreenColor()
	return MakeColor(result_)
}

func LightGrayColor() Color {
	result_ := C.C_NSColor_LightGrayColor()
	return MakeColor(result_)
}

func MagentaColor() Color {
	result_ := C.C_NSColor_MagentaColor()
	return MakeColor(result_)
}

func OrangeColor() Color {
	result_ := C.C_NSColor_OrangeColor()
	return MakeColor(result_)
}

func PurpleColor() Color {
	result_ := C.C_NSColor_PurpleColor()
	return MakeColor(result_)
}

func RedColor() Color {
	result_ := C.C_NSColor_RedColor()
	return MakeColor(result_)
}

func WhiteColor() Color {
	result_ := C.C_NSColor_WhiteColor()
	return MakeColor(result_)
}

func YellowColor() Color {
	result_ := C.C_NSColor_YellowColor()
	return MakeColor(result_)
}

func (n NSColor) PatternImage() Image {
	result_ := C.C_NSColor_PatternImage(n.Ptr())
	return MakeImage(result_)
}
