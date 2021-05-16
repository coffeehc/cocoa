package appkit

// #cgo CFLAGS: -x objective-c
// #cgo LDFLAGS: -framework AppKit
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

func MakeColor(ptr unsafe.Pointer) *NSColor {
	if ptr == nil {
		return nil
	}
	return &NSColor{
		NSObject: *objc.MakeObject(ptr),
	}
}

func AllocColor() *NSColor {
	return MakeColor(C.C_Color_Alloc())
}

func (n *NSColor) Init() Color {
	result := C.C_NSColor_Init(n.Ptr())
	return MakeColor(result)
}

func (n *NSColor) InitWithCoder(coder foundation.Coder) Color {
	result := C.C_NSColor_InitWithCoder(n.Ptr(), objc.ExtractPtr(coder))
	return MakeColor(result)
}

func (n *NSColor) ColorWithSystemEffect(systemEffect ColorSystemEffect) Color {
	result := C.C_NSColor_ColorWithSystemEffect(n.Ptr(), C.int(int(systemEffect)))
	return MakeColor(result)
}

func (n *NSColor) ColorUsingColorSpace(space ColorSpace) Color {
	result := C.C_NSColor_ColorUsingColorSpace(n.Ptr(), objc.ExtractPtr(space))
	return MakeColor(result)
}

func (n *NSColor) BlendedColorWithFraction_OfColor(fraction coregraphics.Float, color Color) Color {
	result := C.C_NSColor_BlendedColorWithFraction_OfColor(n.Ptr(), C.double(float64(fraction)), objc.ExtractPtr(color))
	return MakeColor(result)
}

func (n *NSColor) ColorWithAlphaComponent(alpha coregraphics.Float) Color {
	result := C.C_NSColor_ColorWithAlphaComponent(n.Ptr(), C.double(float64(alpha)))
	return MakeColor(result)
}

func (n *NSColor) HighlightWithLevel(val coregraphics.Float) Color {
	result := C.C_NSColor_HighlightWithLevel(n.Ptr(), C.double(float64(val)))
	return MakeColor(result)
}

func (n *NSColor) ShadowWithLevel(val coregraphics.Float) Color {
	result := C.C_NSColor_ShadowWithLevel(n.Ptr(), C.double(float64(val)))
	return MakeColor(result)
}

func ColorFromPasteboard(pasteBoard Pasteboard) Color {
	result := C.C_NSColor_ColorFromPasteboard(objc.ExtractPtr(pasteBoard))
	return MakeColor(result)
}

func (n *NSColor) WriteToPasteboard(pasteBoard Pasteboard) {
	C.C_NSColor_WriteToPasteboard(n.Ptr(), objc.ExtractPtr(pasteBoard))
}

func (n *NSColor) ColorUsingType(_type ColorType) Color {
	result := C.C_NSColor_ColorUsingType(n.Ptr(), C.int(int(_type)))
	return MakeColor(result)
}

func (n *NSColor) DrawSwatchInRect(rect foundation.Rect) {
	C.C_NSColor_DrawSwatchInRect(n.Ptr(), *(*C.CGRect)(coregraphics.ToCGRectPointer(coregraphics.Rect(rect))))
}

func (n *NSColor) Set() {
	C.C_NSColor_Set(n.Ptr())
}

func (n *NSColor) SetFill() {
	C.C_NSColor_SetFill(n.Ptr())
}

func (n *NSColor) SetStroke() {
	C.C_NSColor_SetStroke(n.Ptr())
}

func ColorNamed(name ColorName) Color {
	result := C.C_NSColor_ColorNamed(foundation.NewString(string(name)).Ptr())
	return MakeColor(result)
}

func ColorNamed_Bundle(name ColorName, bundle foundation.Bundle) Color {
	result := C.C_NSColor_ColorNamed_Bundle(foundation.NewString(string(name)).Ptr(), objc.ExtractPtr(bundle))
	return MakeColor(result)
}

func ColorWithCatalogName_ColorName(listName ColorListName, colorName ColorName) Color {
	result := C.C_NSColor_ColorWithCatalogName_ColorName(foundation.NewString(string(listName)).Ptr(), foundation.NewString(string(colorName)).Ptr())
	return MakeColor(result)
}

func ColorWithSRGBRed_Green_Blue_Alpha(red coregraphics.Float, green coregraphics.Float, blue coregraphics.Float, alpha coregraphics.Float) Color {
	result := C.C_NSColor_ColorWithSRGBRed_Green_Blue_Alpha(C.double(float64(red)), C.double(float64(green)), C.double(float64(blue)), C.double(float64(alpha)))
	return MakeColor(result)
}

func ColorWithDisplayP3Red_Green_Blue_Alpha(red coregraphics.Float, green coregraphics.Float, blue coregraphics.Float, alpha coregraphics.Float) Color {
	result := C.C_NSColor_ColorWithDisplayP3Red_Green_Blue_Alpha(C.double(float64(red)), C.double(float64(green)), C.double(float64(blue)), C.double(float64(alpha)))
	return MakeColor(result)
}

func ColorWithRed_Green_Blue_Alpha(red coregraphics.Float, green coregraphics.Float, blue coregraphics.Float, alpha coregraphics.Float) Color {
	result := C.C_NSColor_ColorWithRed_Green_Blue_Alpha(C.double(float64(red)), C.double(float64(green)), C.double(float64(blue)), C.double(float64(alpha)))
	return MakeColor(result)
}

func ColorWithCalibratedRed_Green_Blue_Alpha(red coregraphics.Float, green coregraphics.Float, blue coregraphics.Float, alpha coregraphics.Float) Color {
	result := C.C_NSColor_ColorWithCalibratedRed_Green_Blue_Alpha(C.double(float64(red)), C.double(float64(green)), C.double(float64(blue)), C.double(float64(alpha)))
	return MakeColor(result)
}

func ColorWithDeviceRed_Green_Blue_Alpha(red coregraphics.Float, green coregraphics.Float, blue coregraphics.Float, alpha coregraphics.Float) Color {
	result := C.C_NSColor_ColorWithDeviceRed_Green_Blue_Alpha(C.double(float64(red)), C.double(float64(green)), C.double(float64(blue)), C.double(float64(alpha)))
	return MakeColor(result)
}

func ColorWithCalibratedHue_Saturation_Brightness_Alpha(hue coregraphics.Float, saturation coregraphics.Float, brightness coregraphics.Float, alpha coregraphics.Float) Color {
	result := C.C_NSColor_ColorWithCalibratedHue_Saturation_Brightness_Alpha(C.double(float64(hue)), C.double(float64(saturation)), C.double(float64(brightness)), C.double(float64(alpha)))
	return MakeColor(result)
}

func ColorWithDeviceHue_Saturation_Brightness_Alpha(hue coregraphics.Float, saturation coregraphics.Float, brightness coregraphics.Float, alpha coregraphics.Float) Color {
	result := C.C_NSColor_ColorWithDeviceHue_Saturation_Brightness_Alpha(C.double(float64(hue)), C.double(float64(saturation)), C.double(float64(brightness)), C.double(float64(alpha)))
	return MakeColor(result)
}

func ColorWithHue_Saturation_Brightness_Alpha(hue coregraphics.Float, saturation coregraphics.Float, brightness coregraphics.Float, alpha coregraphics.Float) Color {
	result := C.C_NSColor_ColorWithHue_Saturation_Brightness_Alpha(C.double(float64(hue)), C.double(float64(saturation)), C.double(float64(brightness)), C.double(float64(alpha)))
	return MakeColor(result)
}

func ColorWithColorSpace_Hue_Saturation_Brightness_Alpha(space ColorSpace, hue coregraphics.Float, saturation coregraphics.Float, brightness coregraphics.Float, alpha coregraphics.Float) Color {
	result := C.C_NSColor_ColorWithColorSpace_Hue_Saturation_Brightness_Alpha(objc.ExtractPtr(space), C.double(float64(hue)), C.double(float64(saturation)), C.double(float64(brightness)), C.double(float64(alpha)))
	return MakeColor(result)
}

func ColorWithDeviceCyan_Magenta_Yellow_Black_Alpha(cyan coregraphics.Float, magenta coregraphics.Float, yellow coregraphics.Float, black coregraphics.Float, alpha coregraphics.Float) Color {
	result := C.C_NSColor_ColorWithDeviceCyan_Magenta_Yellow_Black_Alpha(C.double(float64(cyan)), C.double(float64(magenta)), C.double(float64(yellow)), C.double(float64(black)), C.double(float64(alpha)))
	return MakeColor(result)
}

func ColorWithWhite_Alpha(white coregraphics.Float, alpha coregraphics.Float) Color {
	result := C.C_NSColor_ColorWithWhite_Alpha(C.double(float64(white)), C.double(float64(alpha)))
	return MakeColor(result)
}

func ColorWithCalibratedWhite_Alpha(white coregraphics.Float, alpha coregraphics.Float) Color {
	result := C.C_NSColor_ColorWithCalibratedWhite_Alpha(C.double(float64(white)), C.double(float64(alpha)))
	return MakeColor(result)
}

func ColorWithDeviceWhite_Alpha(white coregraphics.Float, alpha coregraphics.Float) Color {
	result := C.C_NSColor_ColorWithDeviceWhite_Alpha(C.double(float64(white)), C.double(float64(alpha)))
	return MakeColor(result)
}

func ColorWithGenericGamma22White_Alpha(white coregraphics.Float, alpha coregraphics.Float) Color {
	result := C.C_NSColor_ColorWithGenericGamma22White_Alpha(C.double(float64(white)), C.double(float64(alpha)))
	return MakeColor(result)
}

func ColorWithPatternImage(image Image) Color {
	result := C.C_NSColor_ColorWithPatternImage(objc.ExtractPtr(image))
	return MakeColor(result)
}

func ColorWithCGColor(cgColor coregraphics.ColorRef) Color {
	result := C.C_NSColor_ColorWithCGColor(*(*C.CGColorRef)(coregraphics.ToCGColorRefPointer(cgColor)))
	return MakeColor(result)
}

func Color_IgnoresAlpha() bool {
	result := C.C_NSColor_Color_IgnoresAlpha()
	return bool(result)
}

func Color_SetIgnoresAlpha(value bool) {
	C.C_NSColor_Color_SetIgnoresAlpha(C.bool(value))
}

func (n *NSColor) NumberOfComponents() int {
	result := C.C_NSColor_NumberOfComponents(n.Ptr())
	return int(result)
}

func (n *NSColor) AlphaComponent() coregraphics.Float {
	result := C.C_NSColor_AlphaComponent(n.Ptr())
	return coregraphics.Float(float64(result))
}

func (n *NSColor) WhiteComponent() coregraphics.Float {
	result := C.C_NSColor_WhiteComponent(n.Ptr())
	return coregraphics.Float(float64(result))
}

func (n *NSColor) RedComponent() coregraphics.Float {
	result := C.C_NSColor_RedComponent(n.Ptr())
	return coregraphics.Float(float64(result))
}

func (n *NSColor) GreenComponent() coregraphics.Float {
	result := C.C_NSColor_GreenComponent(n.Ptr())
	return coregraphics.Float(float64(result))
}

func (n *NSColor) BlueComponent() coregraphics.Float {
	result := C.C_NSColor_BlueComponent(n.Ptr())
	return coregraphics.Float(float64(result))
}

func (n *NSColor) CyanComponent() coregraphics.Float {
	result := C.C_NSColor_CyanComponent(n.Ptr())
	return coregraphics.Float(float64(result))
}

func (n *NSColor) MagentaComponent() coregraphics.Float {
	result := C.C_NSColor_MagentaComponent(n.Ptr())
	return coregraphics.Float(float64(result))
}

func (n *NSColor) YellowComponent() coregraphics.Float {
	result := C.C_NSColor_YellowComponent(n.Ptr())
	return coregraphics.Float(float64(result))
}

func (n *NSColor) BlackComponent() coregraphics.Float {
	result := C.C_NSColor_BlackComponent(n.Ptr())
	return coregraphics.Float(float64(result))
}

func (n *NSColor) HueComponent() coregraphics.Float {
	result := C.C_NSColor_HueComponent(n.Ptr())
	return coregraphics.Float(float64(result))
}

func (n *NSColor) SaturationComponent() coregraphics.Float {
	result := C.C_NSColor_SaturationComponent(n.Ptr())
	return coregraphics.Float(float64(result))
}

func (n *NSColor) BrightnessComponent() coregraphics.Float {
	result := C.C_NSColor_BrightnessComponent(n.Ptr())
	return coregraphics.Float(float64(result))
}

func (n *NSColor) CatalogNameComponent() ColorListName {
	result := C.C_NSColor_CatalogNameComponent(n.Ptr())
	return ColorListName(foundation.MakeString(result).String())
}

func (n *NSColor) LocalizedCatalogNameComponent() string {
	result := C.C_NSColor_LocalizedCatalogNameComponent(n.Ptr())
	return foundation.MakeString(result).String()
}

func (n *NSColor) ColorNameComponent() ColorName {
	result := C.C_NSColor_ColorNameComponent(n.Ptr())
	return ColorName(foundation.MakeString(result).String())
}

func (n *NSColor) LocalizedColorNameComponent() string {
	result := C.C_NSColor_LocalizedColorNameComponent(n.Ptr())
	return foundation.MakeString(result).String()
}

func (n *NSColor) Type() ColorType {
	result := C.C_NSColor_Type(n.Ptr())
	return ColorType(int(result))
}

func (n *NSColor) ColorSpace() ColorSpace {
	result := C.C_NSColor_ColorSpace(n.Ptr())
	return MakeColorSpace(result)
}

func (n *NSColor) CGColor() coregraphics.ColorRef {
	result := C.C_NSColor_CGColor(n.Ptr())
	return coregraphics.FromCGColorRefPointer(unsafe.Pointer(&result))
}

func LabelColor() Color {
	result := C.C_NSColor_LabelColor()
	return MakeColor(result)
}

func SecondaryLabelColor() Color {
	result := C.C_NSColor_SecondaryLabelColor()
	return MakeColor(result)
}

func TertiaryLabelColor() Color {
	result := C.C_NSColor_TertiaryLabelColor()
	return MakeColor(result)
}

func QuaternaryLabelColor() Color {
	result := C.C_NSColor_QuaternaryLabelColor()
	return MakeColor(result)
}

func TextColor() Color {
	result := C.C_NSColor_TextColor()
	return MakeColor(result)
}

func PlaceholderTextColor() Color {
	result := C.C_NSColor_PlaceholderTextColor()
	return MakeColor(result)
}

func SelectedTextColor() Color {
	result := C.C_NSColor_SelectedTextColor()
	return MakeColor(result)
}

func TextBackgroundColor() Color {
	result := C.C_NSColor_TextBackgroundColor()
	return MakeColor(result)
}

func SelectedTextBackgroundColor() Color {
	result := C.C_NSColor_SelectedTextBackgroundColor()
	return MakeColor(result)
}

func KeyboardFocusIndicatorColor() Color {
	result := C.C_NSColor_KeyboardFocusIndicatorColor()
	return MakeColor(result)
}

func UnemphasizedSelectedTextColor() Color {
	result := C.C_NSColor_UnemphasizedSelectedTextColor()
	return MakeColor(result)
}

func UnemphasizedSelectedTextBackgroundColor() Color {
	result := C.C_NSColor_UnemphasizedSelectedTextBackgroundColor()
	return MakeColor(result)
}

func LinkColor() Color {
	result := C.C_NSColor_LinkColor()
	return MakeColor(result)
}

func SeparatorColor() Color {
	result := C.C_NSColor_SeparatorColor()
	return MakeColor(result)
}

func SelectedContentBackgroundColor() Color {
	result := C.C_NSColor_SelectedContentBackgroundColor()
	return MakeColor(result)
}

func UnemphasizedSelectedContentBackgroundColor() Color {
	result := C.C_NSColor_UnemphasizedSelectedContentBackgroundColor()
	return MakeColor(result)
}

func SelectedMenuItemTextColor() Color {
	result := C.C_NSColor_SelectedMenuItemTextColor()
	return MakeColor(result)
}

func GridColor() Color {
	result := C.C_NSColor_GridColor()
	return MakeColor(result)
}

func HeaderTextColor() Color {
	result := C.C_NSColor_HeaderTextColor()
	return MakeColor(result)
}

func Color_AlternatingContentBackgroundColors() []Color {
	result := C.C_NSColor_Color_AlternatingContentBackgroundColors()
	defer C.free(result.data)
	resultSlice := (*[1 << 28]unsafe.Pointer)(unsafe.Pointer(result.data))[:result.len:result.len]
	var goResult = make([]Color, len(resultSlice))
	for idx, r := range resultSlice {
		goResult[idx] = MakeColor(r)
	}
	return goResult
}

func ControlAccentColor() Color {
	result := C.C_NSColor_ControlAccentColor()
	return MakeColor(result)
}

func ControlColor() Color {
	result := C.C_NSColor_ControlColor()
	return MakeColor(result)
}

func ControlBackgroundColor() Color {
	result := C.C_NSColor_ControlBackgroundColor()
	return MakeColor(result)
}

func ControlTextColor() Color {
	result := C.C_NSColor_ControlTextColor()
	return MakeColor(result)
}

func DisabledControlTextColor() Color {
	result := C.C_NSColor_DisabledControlTextColor()
	return MakeColor(result)
}

func Color_CurrentControlTint() ControlTint {
	result := C.C_NSColor_Color_CurrentControlTint()
	return ControlTint(uint(result))
}

func SelectedControlColor() Color {
	result := C.C_NSColor_SelectedControlColor()
	return MakeColor(result)
}

func SelectedControlTextColor() Color {
	result := C.C_NSColor_SelectedControlTextColor()
	return MakeColor(result)
}

func AlternateSelectedControlTextColor() Color {
	result := C.C_NSColor_AlternateSelectedControlTextColor()
	return MakeColor(result)
}

func ScrubberTexturedBackgroundColor() Color {
	result := C.C_NSColor_ScrubberTexturedBackgroundColor()
	return MakeColor(result)
}

func WindowBackgroundColor() Color {
	result := C.C_NSColor_WindowBackgroundColor()
	return MakeColor(result)
}

func WindowFrameTextColor() Color {
	result := C.C_NSColor_WindowFrameTextColor()
	return MakeColor(result)
}

func UnderPageBackgroundColor() Color {
	result := C.C_NSColor_UnderPageBackgroundColor()
	return MakeColor(result)
}

func FindHighlightColor() Color {
	result := C.C_NSColor_FindHighlightColor()
	return MakeColor(result)
}

func HighlightColor() Color {
	result := C.C_NSColor_HighlightColor()
	return MakeColor(result)
}

func ShadowColor() Color {
	result := C.C_NSColor_ShadowColor()
	return MakeColor(result)
}

func SystemBlueColor() Color {
	result := C.C_NSColor_SystemBlueColor()
	return MakeColor(result)
}

func SystemBrownColor() Color {
	result := C.C_NSColor_SystemBrownColor()
	return MakeColor(result)
}

func SystemGrayColor() Color {
	result := C.C_NSColor_SystemGrayColor()
	return MakeColor(result)
}

func SystemGreenColor() Color {
	result := C.C_NSColor_SystemGreenColor()
	return MakeColor(result)
}

func SystemIndigoColor() Color {
	result := C.C_NSColor_SystemIndigoColor()
	return MakeColor(result)
}

func SystemOrangeColor() Color {
	result := C.C_NSColor_SystemOrangeColor()
	return MakeColor(result)
}

func SystemPinkColor() Color {
	result := C.C_NSColor_SystemPinkColor()
	return MakeColor(result)
}

func SystemPurpleColor() Color {
	result := C.C_NSColor_SystemPurpleColor()
	return MakeColor(result)
}

func SystemRedColor() Color {
	result := C.C_NSColor_SystemRedColor()
	return MakeColor(result)
}

func SystemTealColor() Color {
	result := C.C_NSColor_SystemTealColor()
	return MakeColor(result)
}

func SystemYellowColor() Color {
	result := C.C_NSColor_SystemYellowColor()
	return MakeColor(result)
}

func ClearColor() Color {
	result := C.C_NSColor_ClearColor()
	return MakeColor(result)
}

func BlackColor() Color {
	result := C.C_NSColor_BlackColor()
	return MakeColor(result)
}

func BlueColor() Color {
	result := C.C_NSColor_BlueColor()
	return MakeColor(result)
}

func BrownColor() Color {
	result := C.C_NSColor_BrownColor()
	return MakeColor(result)
}

func CyanColor() Color {
	result := C.C_NSColor_CyanColor()
	return MakeColor(result)
}

func DarkGrayColor() Color {
	result := C.C_NSColor_DarkGrayColor()
	return MakeColor(result)
}

func GrayColor() Color {
	result := C.C_NSColor_GrayColor()
	return MakeColor(result)
}

func GreenColor() Color {
	result := C.C_NSColor_GreenColor()
	return MakeColor(result)
}

func LightGrayColor() Color {
	result := C.C_NSColor_LightGrayColor()
	return MakeColor(result)
}

func MagentaColor() Color {
	result := C.C_NSColor_MagentaColor()
	return MakeColor(result)
}

func OrangeColor() Color {
	result := C.C_NSColor_OrangeColor()
	return MakeColor(result)
}

func PurpleColor() Color {
	result := C.C_NSColor_PurpleColor()
	return MakeColor(result)
}

func RedColor() Color {
	result := C.C_NSColor_RedColor()
	return MakeColor(result)
}

func WhiteColor() Color {
	result := C.C_NSColor_WhiteColor()
	return MakeColor(result)
}

func YellowColor() Color {
	result := C.C_NSColor_YellowColor()
	return MakeColor(result)
}

func (n *NSColor) PatternImage() Image {
	result := C.C_NSColor_PatternImage(n.Ptr())
	return MakeImage(result)
}
