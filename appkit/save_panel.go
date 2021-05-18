package appkit

// #cgo CFLAGS: -x objective-c
// #cgo LDFLAGS: -framework AppKit
// #include "save_panel.h"
import "C"
import (
	"github.com/hsiafan/cocoa/coregraphics"
	"github.com/hsiafan/cocoa/foundation"
	"github.com/hsiafan/cocoa/objc"
	"github.com/hsiafan/cocoa/uti"
	"unsafe"
)

type SavePanel interface {
	Panel
	RunModal() ModalResponse
	ValidateVisibleColumns()
	Ok(sender objc.Object)
	Cancel(sender objc.Object)
	URL() foundation.URL
	Prompt() string
	SetPrompt(value string)
	Message() string
	SetMessage(value string)
	NameFieldLabel() string
	SetNameFieldLabel(value string)
	NameFieldStringValue() string
	SetNameFieldStringValue(value string)
	DirectoryURL() foundation.URL
	SetDirectoryURL(value foundation.URL)
	AccessoryView() View
	SetAccessoryView(value View)
	ShowsTagField() bool
	SetShowsTagField(value bool)
	TagNames() []string
	SetTagNames(value []string)
	CanCreateDirectories() bool
	SetCanCreateDirectories(value bool)
	CanSelectHiddenExtension() bool
	SetCanSelectHiddenExtension(value bool)
	ShowsHiddenFiles() bool
	SetShowsHiddenFiles(value bool)
	IsExtensionHidden() bool
	SetExtensionHidden(value bool)
	IsExpanded() bool
	AllowedContentTypes() []uti.UTType
	SetAllowedContentTypes(value []uti.UTType)
	AllowsOtherFileTypes() bool
	SetAllowsOtherFileTypes(value bool)
	TreatsFilePackagesAsDirectories() bool
	SetTreatsFilePackagesAsDirectories(value bool)
}

type NSSavePanel struct {
	NSPanel
}

func MakeSavePanel(ptr unsafe.Pointer) *NSSavePanel {
	if ptr == nil {
		return nil
	}
	return &NSSavePanel{
		NSPanel: *MakePanel(ptr),
	}
}

func AllocSavePanel() *NSSavePanel {
	return MakeSavePanel(C.C_SavePanel_Alloc())
}

func (n *NSSavePanel) InitWithContentRect_StyleMask_Backing_Defer(contentRect foundation.Rect, style WindowStyleMask, backingStoreType BackingStoreType, flag bool) SavePanel {
	result_ := C.C_NSSavePanel_InitWithContentRect_StyleMask_Backing_Defer(n.Ptr(), *(*C.CGRect)(coregraphics.ToCGRectPointer(coregraphics.Rect(contentRect))), C.uint(uint(style)), C.uint(uint(backingStoreType)), C.bool(flag))
	return MakeSavePanel(result_)
}

func (n *NSSavePanel) InitWithContentRect_StyleMask_Backing_Defer_Screen(contentRect foundation.Rect, style WindowStyleMask, backingStoreType BackingStoreType, flag bool, screen Screen) SavePanel {
	result_ := C.C_NSSavePanel_InitWithContentRect_StyleMask_Backing_Defer_Screen(n.Ptr(), *(*C.CGRect)(coregraphics.ToCGRectPointer(coregraphics.Rect(contentRect))), C.uint(uint(style)), C.uint(uint(backingStoreType)), C.bool(flag), objc.ExtractPtr(screen))
	return MakeSavePanel(result_)
}

func (n *NSSavePanel) Init() SavePanel {
	result_ := C.C_NSSavePanel_Init(n.Ptr())
	return MakeSavePanel(result_)
}

func SavePanel_() SavePanel {
	result_ := C.C_NSSavePanel_SavePanel_()
	return MakeSavePanel(result_)
}

func (n *NSSavePanel) RunModal() ModalResponse {
	result_ := C.C_NSSavePanel_RunModal(n.Ptr())
	return ModalResponse(int(result_))
}

func (n *NSSavePanel) ValidateVisibleColumns() {
	C.C_NSSavePanel_ValidateVisibleColumns(n.Ptr())
}

func (n *NSSavePanel) Ok(sender objc.Object) {
	C.C_NSSavePanel_Ok(n.Ptr(), objc.ExtractPtr(sender))
}

func (n *NSSavePanel) Cancel(sender objc.Object) {
	C.C_NSSavePanel_Cancel(n.Ptr(), objc.ExtractPtr(sender))
}

func (n *NSSavePanel) URL() foundation.URL {
	result_ := C.C_NSSavePanel_URL(n.Ptr())
	return foundation.MakeURL(result_)
}

func (n *NSSavePanel) Prompt() string {
	result_ := C.C_NSSavePanel_Prompt(n.Ptr())
	return foundation.MakeString(result_).String()
}

func (n *NSSavePanel) SetPrompt(value string) {
	C.C_NSSavePanel_SetPrompt(n.Ptr(), foundation.NewString(value).Ptr())
}

func (n *NSSavePanel) Message() string {
	result_ := C.C_NSSavePanel_Message(n.Ptr())
	return foundation.MakeString(result_).String()
}

func (n *NSSavePanel) SetMessage(value string) {
	C.C_NSSavePanel_SetMessage(n.Ptr(), foundation.NewString(value).Ptr())
}

func (n *NSSavePanel) NameFieldLabel() string {
	result_ := C.C_NSSavePanel_NameFieldLabel(n.Ptr())
	return foundation.MakeString(result_).String()
}

func (n *NSSavePanel) SetNameFieldLabel(value string) {
	C.C_NSSavePanel_SetNameFieldLabel(n.Ptr(), foundation.NewString(value).Ptr())
}

func (n *NSSavePanel) NameFieldStringValue() string {
	result_ := C.C_NSSavePanel_NameFieldStringValue(n.Ptr())
	return foundation.MakeString(result_).String()
}

func (n *NSSavePanel) SetNameFieldStringValue(value string) {
	C.C_NSSavePanel_SetNameFieldStringValue(n.Ptr(), foundation.NewString(value).Ptr())
}

func (n *NSSavePanel) DirectoryURL() foundation.URL {
	result_ := C.C_NSSavePanel_DirectoryURL(n.Ptr())
	return foundation.MakeURL(result_)
}

func (n *NSSavePanel) SetDirectoryURL(value foundation.URL) {
	C.C_NSSavePanel_SetDirectoryURL(n.Ptr(), objc.ExtractPtr(value))
}

func (n *NSSavePanel) AccessoryView() View {
	result_ := C.C_NSSavePanel_AccessoryView(n.Ptr())
	return MakeView(result_)
}

func (n *NSSavePanel) SetAccessoryView(value View) {
	C.C_NSSavePanel_SetAccessoryView(n.Ptr(), objc.ExtractPtr(value))
}

func (n *NSSavePanel) ShowsTagField() bool {
	result_ := C.C_NSSavePanel_ShowsTagField(n.Ptr())
	return bool(result_)
}

func (n *NSSavePanel) SetShowsTagField(value bool) {
	C.C_NSSavePanel_SetShowsTagField(n.Ptr(), C.bool(value))
}

func (n *NSSavePanel) TagNames() []string {
	result_ := C.C_NSSavePanel_TagNames(n.Ptr())
	defer C.free(result_.data)
	result_Slice := (*[1 << 28]unsafe.Pointer)(unsafe.Pointer(result_.data))[:result_.len:result_.len]
	var goResult_ = make([]string, len(result_Slice))
	for idx, r := range result_Slice {
		goResult_[idx] = foundation.MakeString(r).String()
	}
	return goResult_
}

func (n *NSSavePanel) SetTagNames(value []string) {
	cValueData := make([]unsafe.Pointer, len(value))
	for idx, v := range value {
		cValueData[idx] = foundation.NewString(v).Ptr()
	}
	cValue := C.Array{data: unsafe.Pointer(&cValueData[0]), len: C.int(len(value))}
	C.C_NSSavePanel_SetTagNames(n.Ptr(), cValue)
}

func (n *NSSavePanel) CanCreateDirectories() bool {
	result_ := C.C_NSSavePanel_CanCreateDirectories(n.Ptr())
	return bool(result_)
}

func (n *NSSavePanel) SetCanCreateDirectories(value bool) {
	C.C_NSSavePanel_SetCanCreateDirectories(n.Ptr(), C.bool(value))
}

func (n *NSSavePanel) CanSelectHiddenExtension() bool {
	result_ := C.C_NSSavePanel_CanSelectHiddenExtension(n.Ptr())
	return bool(result_)
}

func (n *NSSavePanel) SetCanSelectHiddenExtension(value bool) {
	C.C_NSSavePanel_SetCanSelectHiddenExtension(n.Ptr(), C.bool(value))
}

func (n *NSSavePanel) ShowsHiddenFiles() bool {
	result_ := C.C_NSSavePanel_ShowsHiddenFiles(n.Ptr())
	return bool(result_)
}

func (n *NSSavePanel) SetShowsHiddenFiles(value bool) {
	C.C_NSSavePanel_SetShowsHiddenFiles(n.Ptr(), C.bool(value))
}

func (n *NSSavePanel) IsExtensionHidden() bool {
	result_ := C.C_NSSavePanel_IsExtensionHidden(n.Ptr())
	return bool(result_)
}

func (n *NSSavePanel) SetExtensionHidden(value bool) {
	C.C_NSSavePanel_SetExtensionHidden(n.Ptr(), C.bool(value))
}

func (n *NSSavePanel) IsExpanded() bool {
	result_ := C.C_NSSavePanel_IsExpanded(n.Ptr())
	return bool(result_)
}

func (n *NSSavePanel) AllowedContentTypes() []uti.UTType {
	result_ := C.C_NSSavePanel_AllowedContentTypes(n.Ptr())
	defer C.free(result_.data)
	result_Slice := (*[1 << 28]unsafe.Pointer)(unsafe.Pointer(result_.data))[:result_.len:result_.len]
	var goResult_ = make([]uti.UTType, len(result_Slice))
	for idx, r := range result_Slice {
		goResult_[idx] = uti.MakeUTType(r)
	}
	return goResult_
}

func (n *NSSavePanel) SetAllowedContentTypes(value []uti.UTType) {
	cValueData := make([]unsafe.Pointer, len(value))
	for idx, v := range value {
		cValueData[idx] = objc.ExtractPtr(v)
	}
	cValue := C.Array{data: unsafe.Pointer(&cValueData[0]), len: C.int(len(value))}
	C.C_NSSavePanel_SetAllowedContentTypes(n.Ptr(), cValue)
}

func (n *NSSavePanel) AllowsOtherFileTypes() bool {
	result_ := C.C_NSSavePanel_AllowsOtherFileTypes(n.Ptr())
	return bool(result_)
}

func (n *NSSavePanel) SetAllowsOtherFileTypes(value bool) {
	C.C_NSSavePanel_SetAllowsOtherFileTypes(n.Ptr(), C.bool(value))
}

func (n *NSSavePanel) TreatsFilePackagesAsDirectories() bool {
	result_ := C.C_NSSavePanel_TreatsFilePackagesAsDirectories(n.Ptr())
	return bool(result_)
}

func (n *NSSavePanel) SetTreatsFilePackagesAsDirectories(value bool) {
	C.C_NSSavePanel_SetTreatsFilePackagesAsDirectories(n.Ptr(), C.bool(value))
}
