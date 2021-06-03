package appkit

// #include "button.h"
import "C"
import (
	"github.com/hsiafan/cocoa/coregraphics"
	"github.com/hsiafan/cocoa/foundation"
	"github.com/hsiafan/cocoa/objc"
	"unsafe"
)

type Button interface {
	Control
	SetButtonType(_type ButtonType)
	SetPeriodicDelay_Interval(delay float32, interval float32)
	CompressWithPrioritizedCompressionOptions(prioritizedOptions []UserInterfaceCompressionOptions)
	MinimumSizeWithPrioritizedCompressionOptions(prioritizedOptions []UserInterfaceCompressionOptions) foundation.Size
	SetNextState()
	Highlight(flag bool)
	ContentTintColor() Color
	SetContentTintColor(value Color)
	HasDestructiveAction() bool
	SetHasDestructiveAction(value bool)
	AlternateTitle() string
	SetAlternateTitle(value string)
	AttributedTitle() foundation.AttributedString
	SetAttributedTitle(value foundation.AttributedString)
	AttributedAlternateTitle() foundation.AttributedString
	SetAttributedAlternateTitle(value foundation.AttributedString)
	Title() string
	SetTitle(value string)
	Sound() Sound
	SetSound(value Sound)
	IsSpringLoaded() bool
	SetSpringLoaded(value bool)
	MaxAcceleratorLevel() int
	SetMaxAcceleratorLevel(value int)
	Image() Image
	SetImage(value Image)
	AlternateImage() Image
	SetAlternateImage(value Image)
	ImagePosition() CellImagePosition
	SetImagePosition(value CellImagePosition)
	IsBordered() bool
	SetBordered(value bool)
	IsTransparent() bool
	SetTransparent(value bool)
	BezelStyle() BezelStyle
	SetBezelStyle(value BezelStyle)
	BezelColor() Color
	SetBezelColor(value Color)
	ShowsBorderOnlyWhileMouseInside() bool
	SetShowsBorderOnlyWhileMouseInside(value bool)
	ImageHugsTitle() bool
	SetImageHugsTitle(value bool)
	ImageScaling() ImageScaling
	SetImageScaling(value ImageScaling)
	ActiveCompressionOptions() UserInterfaceCompressionOptions
	AllowsMixedState() bool
	SetAllowsMixedState(value bool)
	State() ControlStateValue
	SetState(value ControlStateValue)
	KeyEquivalent() string
	SetKeyEquivalent(value string)
	KeyEquivalentModifierMask() EventModifierFlags
	SetKeyEquivalentModifierMask(value EventModifierFlags)
	SymbolConfiguration() ImageSymbolConfiguration
	SetSymbolConfiguration(value ImageSymbolConfiguration)
}

type NSButton struct {
	NSControl
}

func MakeButton(ptr unsafe.Pointer) NSButton {
	return NSButton{
		NSControl: MakeControl(ptr),
	}
}

func AllocButton() NSButton {
	return MakeButton(C.C_Button_Alloc())
}

func (n NSButton) InitWithFrame(frameRect foundation.Rect) Button {
	result_ := C.C_NSButton_InitWithFrame(n.Ptr(), *(*C.CGRect)(coregraphics.ToCGRectPointer(coregraphics.Rect(frameRect))))
	return MakeButton(result_)
}

func (n NSButton) InitWithCoder(coder foundation.Coder) Button {
	result_ := C.C_NSButton_InitWithCoder(n.Ptr(), objc.ExtractPtr(coder))
	return MakeButton(result_)
}

func (n NSButton) Init() Button {
	result_ := C.C_NSButton_Init(n.Ptr())
	return MakeButton(result_)
}

func Button_CheckboxWithTitle_Target_Action(title string, target objc.Object, action objc.Selector) Button {
	result_ := C.C_NSButton_Button_CheckboxWithTitle_Target_Action(foundation.NewString(title).Ptr(), objc.ExtractPtr(target), unsafe.Pointer(action))
	return MakeButton(result_)
}

func ButtonWithImage_Target_Action(image Image, target objc.Object, action objc.Selector) Button {
	result_ := C.C_NSButton_ButtonWithImage_Target_Action(objc.ExtractPtr(image), objc.ExtractPtr(target), unsafe.Pointer(action))
	return MakeButton(result_)
}

func Button_RadioButtonWithTitle_Target_Action(title string, target objc.Object, action objc.Selector) Button {
	result_ := C.C_NSButton_Button_RadioButtonWithTitle_Target_Action(foundation.NewString(title).Ptr(), objc.ExtractPtr(target), unsafe.Pointer(action))
	return MakeButton(result_)
}

func ButtonWithTitle_Image_Target_Action(title string, image Image, target objc.Object, action objc.Selector) Button {
	result_ := C.C_NSButton_ButtonWithTitle_Image_Target_Action(foundation.NewString(title).Ptr(), objc.ExtractPtr(image), objc.ExtractPtr(target), unsafe.Pointer(action))
	return MakeButton(result_)
}

func ButtonWithTitle_Target_Action(title string, target objc.Object, action objc.Selector) Button {
	result_ := C.C_NSButton_ButtonWithTitle_Target_Action(foundation.NewString(title).Ptr(), objc.ExtractPtr(target), unsafe.Pointer(action))
	return MakeButton(result_)
}

func (n NSButton) SetButtonType(_type ButtonType) {
	C.C_NSButton_SetButtonType(n.Ptr(), C.uint(uint(_type)))
}

func (n NSButton) SetPeriodicDelay_Interval(delay float32, interval float32) {
	C.C_NSButton_SetPeriodicDelay_Interval(n.Ptr(), C.float(delay), C.float(interval))
}

func (n NSButton) CompressWithPrioritizedCompressionOptions(prioritizedOptions []UserInterfaceCompressionOptions) {
	cPrioritizedOptionsData := make([]unsafe.Pointer, len(prioritizedOptions))
	for idx, v := range prioritizedOptions {
		cPrioritizedOptionsData[idx] = objc.ExtractPtr(v)
	}
	cPrioritizedOptions := C.Array{data: unsafe.Pointer(&cPrioritizedOptionsData[0]), len: C.int(len(prioritizedOptions))}
	C.C_NSButton_CompressWithPrioritizedCompressionOptions(n.Ptr(), cPrioritizedOptions)
}

func (n NSButton) MinimumSizeWithPrioritizedCompressionOptions(prioritizedOptions []UserInterfaceCompressionOptions) foundation.Size {
	cPrioritizedOptionsData := make([]unsafe.Pointer, len(prioritizedOptions))
	for idx, v := range prioritizedOptions {
		cPrioritizedOptionsData[idx] = objc.ExtractPtr(v)
	}
	cPrioritizedOptions := C.Array{data: unsafe.Pointer(&cPrioritizedOptionsData[0]), len: C.int(len(prioritizedOptions))}
	result_ := C.C_NSButton_MinimumSizeWithPrioritizedCompressionOptions(n.Ptr(), cPrioritizedOptions)
	return foundation.Size(coregraphics.FromCGSizePointer(unsafe.Pointer(&result_)))
}

func (n NSButton) SetNextState() {
	C.C_NSButton_SetNextState(n.Ptr())
}

func (n NSButton) Highlight(flag bool) {
	C.C_NSButton_Highlight(n.Ptr(), C.bool(flag))
}

func (n NSButton) ContentTintColor() Color {
	result_ := C.C_NSButton_ContentTintColor(n.Ptr())
	return MakeColor(result_)
}

func (n NSButton) SetContentTintColor(value Color) {
	C.C_NSButton_SetContentTintColor(n.Ptr(), objc.ExtractPtr(value))
}

func (n NSButton) HasDestructiveAction() bool {
	result_ := C.C_NSButton_HasDestructiveAction(n.Ptr())
	return bool(result_)
}

func (n NSButton) SetHasDestructiveAction(value bool) {
	C.C_NSButton_SetHasDestructiveAction(n.Ptr(), C.bool(value))
}

func (n NSButton) AlternateTitle() string {
	result_ := C.C_NSButton_AlternateTitle(n.Ptr())
	return foundation.MakeString(result_).String()
}

func (n NSButton) SetAlternateTitle(value string) {
	C.C_NSButton_SetAlternateTitle(n.Ptr(), foundation.NewString(value).Ptr())
}

func (n NSButton) AttributedTitle() foundation.AttributedString {
	result_ := C.C_NSButton_AttributedTitle(n.Ptr())
	return foundation.MakeAttributedString(result_)
}

func (n NSButton) SetAttributedTitle(value foundation.AttributedString) {
	C.C_NSButton_SetAttributedTitle(n.Ptr(), objc.ExtractPtr(value))
}

func (n NSButton) AttributedAlternateTitle() foundation.AttributedString {
	result_ := C.C_NSButton_AttributedAlternateTitle(n.Ptr())
	return foundation.MakeAttributedString(result_)
}

func (n NSButton) SetAttributedAlternateTitle(value foundation.AttributedString) {
	C.C_NSButton_SetAttributedAlternateTitle(n.Ptr(), objc.ExtractPtr(value))
}

func (n NSButton) Title() string {
	result_ := C.C_NSButton_Title(n.Ptr())
	return foundation.MakeString(result_).String()
}

func (n NSButton) SetTitle(value string) {
	C.C_NSButton_SetTitle(n.Ptr(), foundation.NewString(value).Ptr())
}

func (n NSButton) Sound() Sound {
	result_ := C.C_NSButton_Sound(n.Ptr())
	return MakeSound(result_)
}

func (n NSButton) SetSound(value Sound) {
	C.C_NSButton_SetSound(n.Ptr(), objc.ExtractPtr(value))
}

func (n NSButton) IsSpringLoaded() bool {
	result_ := C.C_NSButton_IsSpringLoaded(n.Ptr())
	return bool(result_)
}

func (n NSButton) SetSpringLoaded(value bool) {
	C.C_NSButton_SetSpringLoaded(n.Ptr(), C.bool(value))
}

func (n NSButton) MaxAcceleratorLevel() int {
	result_ := C.C_NSButton_MaxAcceleratorLevel(n.Ptr())
	return int(result_)
}

func (n NSButton) SetMaxAcceleratorLevel(value int) {
	C.C_NSButton_SetMaxAcceleratorLevel(n.Ptr(), C.int(value))
}

func (n NSButton) Image() Image {
	result_ := C.C_NSButton_Image(n.Ptr())
	return MakeImage(result_)
}

func (n NSButton) SetImage(value Image) {
	C.C_NSButton_SetImage(n.Ptr(), objc.ExtractPtr(value))
}

func (n NSButton) AlternateImage() Image {
	result_ := C.C_NSButton_AlternateImage(n.Ptr())
	return MakeImage(result_)
}

func (n NSButton) SetAlternateImage(value Image) {
	C.C_NSButton_SetAlternateImage(n.Ptr(), objc.ExtractPtr(value))
}

func (n NSButton) ImagePosition() CellImagePosition {
	result_ := C.C_NSButton_ImagePosition(n.Ptr())
	return CellImagePosition(uint(result_))
}

func (n NSButton) SetImagePosition(value CellImagePosition) {
	C.C_NSButton_SetImagePosition(n.Ptr(), C.uint(uint(value)))
}

func (n NSButton) IsBordered() bool {
	result_ := C.C_NSButton_IsBordered(n.Ptr())
	return bool(result_)
}

func (n NSButton) SetBordered(value bool) {
	C.C_NSButton_SetBordered(n.Ptr(), C.bool(value))
}

func (n NSButton) IsTransparent() bool {
	result_ := C.C_NSButton_IsTransparent(n.Ptr())
	return bool(result_)
}

func (n NSButton) SetTransparent(value bool) {
	C.C_NSButton_SetTransparent(n.Ptr(), C.bool(value))
}

func (n NSButton) BezelStyle() BezelStyle {
	result_ := C.C_NSButton_BezelStyle(n.Ptr())
	return BezelStyle(uint(result_))
}

func (n NSButton) SetBezelStyle(value BezelStyle) {
	C.C_NSButton_SetBezelStyle(n.Ptr(), C.uint(uint(value)))
}

func (n NSButton) BezelColor() Color {
	result_ := C.C_NSButton_BezelColor(n.Ptr())
	return MakeColor(result_)
}

func (n NSButton) SetBezelColor(value Color) {
	C.C_NSButton_SetBezelColor(n.Ptr(), objc.ExtractPtr(value))
}

func (n NSButton) ShowsBorderOnlyWhileMouseInside() bool {
	result_ := C.C_NSButton_ShowsBorderOnlyWhileMouseInside(n.Ptr())
	return bool(result_)
}

func (n NSButton) SetShowsBorderOnlyWhileMouseInside(value bool) {
	C.C_NSButton_SetShowsBorderOnlyWhileMouseInside(n.Ptr(), C.bool(value))
}

func (n NSButton) ImageHugsTitle() bool {
	result_ := C.C_NSButton_ImageHugsTitle(n.Ptr())
	return bool(result_)
}

func (n NSButton) SetImageHugsTitle(value bool) {
	C.C_NSButton_SetImageHugsTitle(n.Ptr(), C.bool(value))
}

func (n NSButton) ImageScaling() ImageScaling {
	result_ := C.C_NSButton_ImageScaling(n.Ptr())
	return ImageScaling(uint(result_))
}

func (n NSButton) SetImageScaling(value ImageScaling) {
	C.C_NSButton_SetImageScaling(n.Ptr(), C.uint(uint(value)))
}

func (n NSButton) ActiveCompressionOptions() UserInterfaceCompressionOptions {
	result_ := C.C_NSButton_ActiveCompressionOptions(n.Ptr())
	return MakeUserInterfaceCompressionOptions(result_)
}

func (n NSButton) AllowsMixedState() bool {
	result_ := C.C_NSButton_AllowsMixedState(n.Ptr())
	return bool(result_)
}

func (n NSButton) SetAllowsMixedState(value bool) {
	C.C_NSButton_SetAllowsMixedState(n.Ptr(), C.bool(value))
}

func (n NSButton) State() ControlStateValue {
	result_ := C.C_NSButton_State(n.Ptr())
	return ControlStateValue(int(result_))
}

func (n NSButton) SetState(value ControlStateValue) {
	C.C_NSButton_SetState(n.Ptr(), C.int(int(value)))
}

func (n NSButton) KeyEquivalent() string {
	result_ := C.C_NSButton_KeyEquivalent(n.Ptr())
	return foundation.MakeString(result_).String()
}

func (n NSButton) SetKeyEquivalent(value string) {
	C.C_NSButton_SetKeyEquivalent(n.Ptr(), foundation.NewString(value).Ptr())
}

func (n NSButton) KeyEquivalentModifierMask() EventModifierFlags {
	result_ := C.C_NSButton_KeyEquivalentModifierMask(n.Ptr())
	return EventModifierFlags(uint(result_))
}

func (n NSButton) SetKeyEquivalentModifierMask(value EventModifierFlags) {
	C.C_NSButton_SetKeyEquivalentModifierMask(n.Ptr(), C.uint(uint(value)))
}

func (n NSButton) SymbolConfiguration() ImageSymbolConfiguration {
	result_ := C.C_NSButton_SymbolConfiguration(n.Ptr())
	return MakeImageSymbolConfiguration(result_)
}

func (n NSButton) SetSymbolConfiguration(value ImageSymbolConfiguration) {
	C.C_NSButton_SetSymbolConfiguration(n.Ptr(), objc.ExtractPtr(value))
}
