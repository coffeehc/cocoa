package appkit

// #include "font_manager.h"
import "C"
import (
	"github.com/hsiafan/cocoa/coregraphics"
	"github.com/hsiafan/cocoa/foundation"
	"github.com/hsiafan/cocoa/objc"
	"unsafe"
)

type FontManager interface {
	objc.Object
	AvailableFontNamesWithTraits(someTraits FontTraitMask) []string
	SetSelectedFont_IsMultiple(fontObj Font, flag bool)
	SendAction() bool
	LocalizedNameForFamily_Face(family string, faceKey string) string
	AddFontTrait(sender objc.Object)
	RemoveFontTrait(sender objc.Object)
	ModifyFont(sender objc.Object)
	ModifyFontViaPanel(sender objc.Object)
	OrderFrontStylesPanel(sender objc.Object)
	OrderFrontFontPanel(sender objc.Object)
	ConvertFont(fontObj Font) Font
	ConvertFont_ToFace(fontObj Font, typeface string) Font
	ConvertFont_ToFamily(fontObj Font, family string) Font
	ConvertFont_ToHaveTrait(fontObj Font, trait FontTraitMask) Font
	ConvertFont_ToNotHaveTrait(fontObj Font, trait FontTraitMask) Font
	ConvertFont_ToSize(fontObj Font, size coregraphics.Float) Font
	ConvertWeight_OfFont(upFlag bool, fontObj Font) Font
	ConvertFontTraits(traits FontTraitMask) FontTraitMask
	FontWithFamily_Traits_Weight_Size(family string, traits FontTraitMask, weight int, size coregraphics.Float) Font
	TraitsOfFont(fontObj Font) FontTraitMask
	FontNamed_HasTraits(fName string, someTraits FontTraitMask) bool
	WeightOfFont(fontObj Font) int
	FontPanel(create bool) FontPanel
	SetFontMenu(newMenu Menu)
	FontMenu(create bool) Menu
	SetSelectedAttributes_IsMultiple(attributes map[string]objc.Object, flag bool)
	ConvertAttributes(attributes map[string]objc.Object) map[string]objc.Object
	AvailableFonts() []string
	AvailableFontFamilies() []string
	SelectedFont() Font
	IsMultiple() bool
	CurrentFontAction() FontAction
	IsEnabled() bool
	SetEnabled(value bool)
	Action() objc.Selector
	SetAction(value objc.Selector)
	Target() objc.Object
	SetTarget(value objc.Object)
}

type NSFontManager struct {
	objc.NSObject
}

func MakeFontManager(ptr unsafe.Pointer) NSFontManager {
	return NSFontManager{
		NSObject: objc.MakeObject(ptr),
	}
}

func AllocFontManager() NSFontManager {
	return MakeFontManager(C.C_FontManager_Alloc())
}

func (n NSFontManager) Init() FontManager {
	result_ := C.C_NSFontManager_Init(n.Ptr())
	return MakeFontManager(result_)
}

func (n NSFontManager) AvailableFontNamesWithTraits(someTraits FontTraitMask) []string {
	result_ := C.C_NSFontManager_AvailableFontNamesWithTraits(n.Ptr(), C.uint(uint(someTraits)))
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

func (n NSFontManager) SetSelectedFont_IsMultiple(fontObj Font, flag bool) {
	C.C_NSFontManager_SetSelectedFont_IsMultiple(n.Ptr(), objc.ExtractPtr(fontObj), C.bool(flag))
}

func (n NSFontManager) SendAction() bool {
	result_ := C.C_NSFontManager_SendAction(n.Ptr())
	return bool(result_)
}

func (n NSFontManager) LocalizedNameForFamily_Face(family string, faceKey string) string {
	result_ := C.C_NSFontManager_LocalizedNameForFamily_Face(n.Ptr(), foundation.NewString(family).Ptr(), foundation.NewString(faceKey).Ptr())
	return foundation.MakeString(result_).String()
}

func (n NSFontManager) AddFontTrait(sender objc.Object) {
	C.C_NSFontManager_AddFontTrait(n.Ptr(), objc.ExtractPtr(sender))
}

func (n NSFontManager) RemoveFontTrait(sender objc.Object) {
	C.C_NSFontManager_RemoveFontTrait(n.Ptr(), objc.ExtractPtr(sender))
}

func (n NSFontManager) ModifyFont(sender objc.Object) {
	C.C_NSFontManager_ModifyFont(n.Ptr(), objc.ExtractPtr(sender))
}

func (n NSFontManager) ModifyFontViaPanel(sender objc.Object) {
	C.C_NSFontManager_ModifyFontViaPanel(n.Ptr(), objc.ExtractPtr(sender))
}

func (n NSFontManager) OrderFrontStylesPanel(sender objc.Object) {
	C.C_NSFontManager_OrderFrontStylesPanel(n.Ptr(), objc.ExtractPtr(sender))
}

func (n NSFontManager) OrderFrontFontPanel(sender objc.Object) {
	C.C_NSFontManager_OrderFrontFontPanel(n.Ptr(), objc.ExtractPtr(sender))
}

func (n NSFontManager) ConvertFont(fontObj Font) Font {
	result_ := C.C_NSFontManager_ConvertFont(n.Ptr(), objc.ExtractPtr(fontObj))
	return MakeFont(result_)
}

func (n NSFontManager) ConvertFont_ToFace(fontObj Font, typeface string) Font {
	result_ := C.C_NSFontManager_ConvertFont_ToFace(n.Ptr(), objc.ExtractPtr(fontObj), foundation.NewString(typeface).Ptr())
	return MakeFont(result_)
}

func (n NSFontManager) ConvertFont_ToFamily(fontObj Font, family string) Font {
	result_ := C.C_NSFontManager_ConvertFont_ToFamily(n.Ptr(), objc.ExtractPtr(fontObj), foundation.NewString(family).Ptr())
	return MakeFont(result_)
}

func (n NSFontManager) ConvertFont_ToHaveTrait(fontObj Font, trait FontTraitMask) Font {
	result_ := C.C_NSFontManager_ConvertFont_ToHaveTrait(n.Ptr(), objc.ExtractPtr(fontObj), C.uint(uint(trait)))
	return MakeFont(result_)
}

func (n NSFontManager) ConvertFont_ToNotHaveTrait(fontObj Font, trait FontTraitMask) Font {
	result_ := C.C_NSFontManager_ConvertFont_ToNotHaveTrait(n.Ptr(), objc.ExtractPtr(fontObj), C.uint(uint(trait)))
	return MakeFont(result_)
}

func (n NSFontManager) ConvertFont_ToSize(fontObj Font, size coregraphics.Float) Font {
	result_ := C.C_NSFontManager_ConvertFont_ToSize(n.Ptr(), objc.ExtractPtr(fontObj), C.double(float64(size)))
	return MakeFont(result_)
}

func (n NSFontManager) ConvertWeight_OfFont(upFlag bool, fontObj Font) Font {
	result_ := C.C_NSFontManager_ConvertWeight_OfFont(n.Ptr(), C.bool(upFlag), objc.ExtractPtr(fontObj))
	return MakeFont(result_)
}

func (n NSFontManager) ConvertFontTraits(traits FontTraitMask) FontTraitMask {
	result_ := C.C_NSFontManager_ConvertFontTraits(n.Ptr(), C.uint(uint(traits)))
	return FontTraitMask(uint(result_))
}

func (n NSFontManager) FontWithFamily_Traits_Weight_Size(family string, traits FontTraitMask, weight int, size coregraphics.Float) Font {
	result_ := C.C_NSFontManager_FontWithFamily_Traits_Weight_Size(n.Ptr(), foundation.NewString(family).Ptr(), C.uint(uint(traits)), C.int(weight), C.double(float64(size)))
	return MakeFont(result_)
}

func (n NSFontManager) TraitsOfFont(fontObj Font) FontTraitMask {
	result_ := C.C_NSFontManager_TraitsOfFont(n.Ptr(), objc.ExtractPtr(fontObj))
	return FontTraitMask(uint(result_))
}

func (n NSFontManager) FontNamed_HasTraits(fName string, someTraits FontTraitMask) bool {
	result_ := C.C_NSFontManager_FontNamed_HasTraits(n.Ptr(), foundation.NewString(fName).Ptr(), C.uint(uint(someTraits)))
	return bool(result_)
}

func (n NSFontManager) WeightOfFont(fontObj Font) int {
	result_ := C.C_NSFontManager_WeightOfFont(n.Ptr(), objc.ExtractPtr(fontObj))
	return int(result_)
}

func (n NSFontManager) FontPanel(create bool) FontPanel {
	result_ := C.C_NSFontManager_FontPanel(n.Ptr(), C.bool(create))
	return MakeFontPanel(result_)
}

func (n NSFontManager) SetFontMenu(newMenu Menu) {
	C.C_NSFontManager_SetFontMenu(n.Ptr(), objc.ExtractPtr(newMenu))
}

func (n NSFontManager) FontMenu(create bool) Menu {
	result_ := C.C_NSFontManager_FontMenu(n.Ptr(), C.bool(create))
	return MakeMenu(result_)
}

func (n NSFontManager) SetSelectedAttributes_IsMultiple(attributes map[string]objc.Object, flag bool) {
	var cAttributes C.Dictionary
	if len(attributes) > 0 {
		cAttributesKeyData := make([]unsafe.Pointer, len(attributes))
		cAttributesValueData := make([]unsafe.Pointer, len(attributes))
		var idx = 0
		for k, v := range attributes {
			cAttributesKeyData[idx] = foundation.NewString(k).Ptr()
			cAttributesValueData[idx] = objc.ExtractPtr(v)
			idx++
		}
		cAttributes.key_data = unsafe.Pointer(&cAttributesKeyData[0])
		cAttributes.value_data = unsafe.Pointer(&cAttributesValueData[0])
		cAttributes.len = C.int(len(attributes))
	}
	C.C_NSFontManager_SetSelectedAttributes_IsMultiple(n.Ptr(), cAttributes, C.bool(flag))
}

func (n NSFontManager) ConvertAttributes(attributes map[string]objc.Object) map[string]objc.Object {
	var cAttributes C.Dictionary
	if len(attributes) > 0 {
		cAttributesKeyData := make([]unsafe.Pointer, len(attributes))
		cAttributesValueData := make([]unsafe.Pointer, len(attributes))
		var idx = 0
		for k, v := range attributes {
			cAttributesKeyData[idx] = foundation.NewString(k).Ptr()
			cAttributesValueData[idx] = objc.ExtractPtr(v)
			idx++
		}
		cAttributes.key_data = unsafe.Pointer(&cAttributesKeyData[0])
		cAttributes.value_data = unsafe.Pointer(&cAttributesValueData[0])
		cAttributes.len = C.int(len(attributes))
	}
	result_ := C.C_NSFontManager_ConvertAttributes(n.Ptr(), cAttributes)
	if result_.len > 0 {
		defer C.free(result_.key_data)
		defer C.free(result_.value_data)
	}
	result_KeySlice := unsafe.Slice((*unsafe.Pointer)(result_.key_data), int(result_.len))
	result_ValueSlice := unsafe.Slice((*unsafe.Pointer)(result_.value_data), int(result_.len))
	var goResult_ = make(map[string]objc.Object)
	for idx, k := range result_KeySlice {
		v := result_ValueSlice[idx]
		goResult_[foundation.MakeString(k).String()] = objc.MakeObject(v)
	}
	return goResult_
}

func SharedFontManager() FontManager {
	result_ := C.C_NSFontManager_SharedFontManager()
	return MakeFontManager(result_)
}

func (n NSFontManager) AvailableFonts() []string {
	result_ := C.C_NSFontManager_AvailableFonts(n.Ptr())
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

func (n NSFontManager) AvailableFontFamilies() []string {
	result_ := C.C_NSFontManager_AvailableFontFamilies(n.Ptr())
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

func (n NSFontManager) SelectedFont() Font {
	result_ := C.C_NSFontManager_SelectedFont(n.Ptr())
	return MakeFont(result_)
}

func (n NSFontManager) IsMultiple() bool {
	result_ := C.C_NSFontManager_IsMultiple(n.Ptr())
	return bool(result_)
}

func (n NSFontManager) CurrentFontAction() FontAction {
	result_ := C.C_NSFontManager_CurrentFontAction(n.Ptr())
	return FontAction(uint(result_))
}

func (n NSFontManager) IsEnabled() bool {
	result_ := C.C_NSFontManager_IsEnabled(n.Ptr())
	return bool(result_)
}

func (n NSFontManager) SetEnabled(value bool) {
	C.C_NSFontManager_SetEnabled(n.Ptr(), C.bool(value))
}

func (n NSFontManager) Action() objc.Selector {
	result_ := C.C_NSFontManager_Action(n.Ptr())
	return objc.Selector(result_)
}

func (n NSFontManager) SetAction(value objc.Selector) {
	C.C_NSFontManager_SetAction(n.Ptr(), unsafe.Pointer(value))
}

func (n NSFontManager) Target() objc.Object {
	result_ := C.C_NSFontManager_Target(n.Ptr())
	return objc.MakeObject(result_)
}

func (n NSFontManager) SetTarget(value objc.Object) {
	C.C_NSFontManager_SetTarget(n.Ptr(), objc.ExtractPtr(value))
}
