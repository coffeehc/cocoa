package appkit

// #include "touch_bar.h"
import "C"
import (
	"github.com/hsiafan/cocoa/foundation"
	"github.com/hsiafan/cocoa/objc"
	"unsafe"
)

type TouchBar interface {
	objc.Object
	ItemForIdentifier(identifier TouchBarItemIdentifier) TouchBarItem
	Delegate() objc.Object
	SetDelegate(value objc.Object)
	TemplateItems() foundation.Set
	SetTemplateItems(value foundation.Set)
	DefaultItemIdentifiers() []TouchBarItemIdentifier
	SetDefaultItemIdentifiers(value []TouchBarItemIdentifier)
	PrincipalItemIdentifier() TouchBarItemIdentifier
	SetPrincipalItemIdentifier(value TouchBarItemIdentifier)
	EscapeKeyReplacementItemIdentifier() TouchBarItemIdentifier
	SetEscapeKeyReplacementItemIdentifier(value TouchBarItemIdentifier)
	IsVisible() bool
	ItemIdentifiers() []TouchBarItemIdentifier
	CustomizationIdentifier() TouchBarCustomizationIdentifier
	SetCustomizationIdentifier(value TouchBarCustomizationIdentifier)
	CustomizationAllowedItemIdentifiers() []TouchBarItemIdentifier
	SetCustomizationAllowedItemIdentifiers(value []TouchBarItemIdentifier)
	CustomizationRequiredItemIdentifiers() []TouchBarItemIdentifier
	SetCustomizationRequiredItemIdentifiers(value []TouchBarItemIdentifier)
}

type NSTouchBar struct {
	objc.NSObject
}

func MakeTouchBar(ptr unsafe.Pointer) NSTouchBar {
	return NSTouchBar{
		NSObject: objc.MakeObject(ptr),
	}
}

func AllocTouchBar() NSTouchBar {
	return MakeTouchBar(C.C_TouchBar_Alloc())
}

func (n NSTouchBar) Init() TouchBar {
	result_ := C.C_NSTouchBar_Init(n.Ptr())
	return MakeTouchBar(result_)
}

func (n NSTouchBar) InitWithCoder(coder foundation.Coder) TouchBar {
	result_ := C.C_NSTouchBar_InitWithCoder(n.Ptr(), objc.ExtractPtr(coder))
	return MakeTouchBar(result_)
}

func (n NSTouchBar) ItemForIdentifier(identifier TouchBarItemIdentifier) TouchBarItem {
	result_ := C.C_NSTouchBar_ItemForIdentifier(n.Ptr(), foundation.NewString(string(identifier)).Ptr())
	return MakeTouchBarItem(result_)
}

func (n NSTouchBar) Delegate() objc.Object {
	result_ := C.C_NSTouchBar_Delegate(n.Ptr())
	return objc.MakeObject(result_)
}

func (n NSTouchBar) SetDelegate(value objc.Object) {
	C.C_NSTouchBar_SetDelegate(n.Ptr(), objc.ExtractPtr(value))
}

func (n NSTouchBar) TemplateItems() foundation.Set {
	result_ := C.C_NSTouchBar_TemplateItems(n.Ptr())
	return foundation.MakeSet(result_)
}

func (n NSTouchBar) SetTemplateItems(value foundation.Set) {
	C.C_NSTouchBar_SetTemplateItems(n.Ptr(), objc.ExtractPtr(value))
}

func (n NSTouchBar) DefaultItemIdentifiers() []TouchBarItemIdentifier {
	result_ := C.C_NSTouchBar_DefaultItemIdentifiers(n.Ptr())
	if result_.len > 0 {
		defer C.free(result_.data)
	}
	result_Slice := (*[1 << 28]unsafe.Pointer)(unsafe.Pointer(result_.data))[:result_.len:result_.len]
	var goResult_ = make([]TouchBarItemIdentifier, len(result_Slice))
	for idx, r := range result_Slice {
		goResult_[idx] = TouchBarItemIdentifier(foundation.MakeString(r).String())
	}
	return goResult_
}

func (n NSTouchBar) SetDefaultItemIdentifiers(value []TouchBarItemIdentifier) {
	var cValue C.Array
	if len(value) > 0 {
		cValueData := make([]unsafe.Pointer, len(value))
		for idx, v := range value {
			cValueData[idx] = foundation.NewString(string(v)).Ptr()
		}
		cValue.data = unsafe.Pointer(&cValueData[0])
		cValue.len = C.int(len(value))
	}
	C.C_NSTouchBar_SetDefaultItemIdentifiers(n.Ptr(), cValue)
}

func (n NSTouchBar) PrincipalItemIdentifier() TouchBarItemIdentifier {
	result_ := C.C_NSTouchBar_PrincipalItemIdentifier(n.Ptr())
	return TouchBarItemIdentifier(foundation.MakeString(result_).String())
}

func (n NSTouchBar) SetPrincipalItemIdentifier(value TouchBarItemIdentifier) {
	C.C_NSTouchBar_SetPrincipalItemIdentifier(n.Ptr(), foundation.NewString(string(value)).Ptr())
}

func (n NSTouchBar) EscapeKeyReplacementItemIdentifier() TouchBarItemIdentifier {
	result_ := C.C_NSTouchBar_EscapeKeyReplacementItemIdentifier(n.Ptr())
	return TouchBarItemIdentifier(foundation.MakeString(result_).String())
}

func (n NSTouchBar) SetEscapeKeyReplacementItemIdentifier(value TouchBarItemIdentifier) {
	C.C_NSTouchBar_SetEscapeKeyReplacementItemIdentifier(n.Ptr(), foundation.NewString(string(value)).Ptr())
}

func (n NSTouchBar) IsVisible() bool {
	result_ := C.C_NSTouchBar_IsVisible(n.Ptr())
	return bool(result_)
}

func (n NSTouchBar) ItemIdentifiers() []TouchBarItemIdentifier {
	result_ := C.C_NSTouchBar_ItemIdentifiers(n.Ptr())
	if result_.len > 0 {
		defer C.free(result_.data)
	}
	result_Slice := (*[1 << 28]unsafe.Pointer)(unsafe.Pointer(result_.data))[:result_.len:result_.len]
	var goResult_ = make([]TouchBarItemIdentifier, len(result_Slice))
	for idx, r := range result_Slice {
		goResult_[idx] = TouchBarItemIdentifier(foundation.MakeString(r).String())
	}
	return goResult_
}

func (n NSTouchBar) CustomizationIdentifier() TouchBarCustomizationIdentifier {
	result_ := C.C_NSTouchBar_CustomizationIdentifier(n.Ptr())
	return TouchBarCustomizationIdentifier(foundation.MakeString(result_).String())
}

func (n NSTouchBar) SetCustomizationIdentifier(value TouchBarCustomizationIdentifier) {
	C.C_NSTouchBar_SetCustomizationIdentifier(n.Ptr(), foundation.NewString(string(value)).Ptr())
}

func (n NSTouchBar) CustomizationAllowedItemIdentifiers() []TouchBarItemIdentifier {
	result_ := C.C_NSTouchBar_CustomizationAllowedItemIdentifiers(n.Ptr())
	if result_.len > 0 {
		defer C.free(result_.data)
	}
	result_Slice := (*[1 << 28]unsafe.Pointer)(unsafe.Pointer(result_.data))[:result_.len:result_.len]
	var goResult_ = make([]TouchBarItemIdentifier, len(result_Slice))
	for idx, r := range result_Slice {
		goResult_[idx] = TouchBarItemIdentifier(foundation.MakeString(r).String())
	}
	return goResult_
}

func (n NSTouchBar) SetCustomizationAllowedItemIdentifiers(value []TouchBarItemIdentifier) {
	var cValue C.Array
	if len(value) > 0 {
		cValueData := make([]unsafe.Pointer, len(value))
		for idx, v := range value {
			cValueData[idx] = foundation.NewString(string(v)).Ptr()
		}
		cValue.data = unsafe.Pointer(&cValueData[0])
		cValue.len = C.int(len(value))
	}
	C.C_NSTouchBar_SetCustomizationAllowedItemIdentifiers(n.Ptr(), cValue)
}

func (n NSTouchBar) CustomizationRequiredItemIdentifiers() []TouchBarItemIdentifier {
	result_ := C.C_NSTouchBar_CustomizationRequiredItemIdentifiers(n.Ptr())
	if result_.len > 0 {
		defer C.free(result_.data)
	}
	result_Slice := (*[1 << 28]unsafe.Pointer)(unsafe.Pointer(result_.data))[:result_.len:result_.len]
	var goResult_ = make([]TouchBarItemIdentifier, len(result_Slice))
	for idx, r := range result_Slice {
		goResult_[idx] = TouchBarItemIdentifier(foundation.MakeString(r).String())
	}
	return goResult_
}

func (n NSTouchBar) SetCustomizationRequiredItemIdentifiers(value []TouchBarItemIdentifier) {
	var cValue C.Array
	if len(value) > 0 {
		cValueData := make([]unsafe.Pointer, len(value))
		for idx, v := range value {
			cValueData[idx] = foundation.NewString(string(v)).Ptr()
		}
		cValue.data = unsafe.Pointer(&cValueData[0])
		cValue.len = C.int(len(value))
	}
	C.C_NSTouchBar_SetCustomizationRequiredItemIdentifiers(n.Ptr(), cValue)
}

func AutomaticCustomizeTouchBarMenuItemEnabled() bool {
	result_ := C.C_NSTouchBar_AutomaticCustomizeTouchBarMenuItemEnabled()
	return bool(result_)
}

func SetAutomaticCustomizeTouchBarMenuItemEnabled(value bool) {
	C.C_NSTouchBar_SetAutomaticCustomizeTouchBarMenuItemEnabled(C.bool(value))
}
