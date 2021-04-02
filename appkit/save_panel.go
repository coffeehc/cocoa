package appkit

// #cgo CFLAGS: -x objective-c
// #cgo LDFLAGS: -framework AppKit
// #include "save_panel.h"
import "C"
import (
	"github.com/hsiafan/cocoa/foundation"
	"github.com/hsiafan/cocoa/uti"
	"unsafe"
)

type SavePanel interface {
	Panel
	URL() foundation.URL
	Prompt() string
	SetPrompt(prompt string)
	Message() string
	SetMessage(message string)
	NameFieldLabel() string
	SetNameFieldLabel(nameFieldLabel string)
	DirectoryURL() foundation.URL
	SetDirectoryURL(directoryURL foundation.URL)
	AccessoryView() View
	SetAccessoryView(accessoryView View)
	ShowsTagField() bool
	SetShowsTagField(showsTagField bool)
	TagNames() []string
	SetTagNames(tagNames []string)
	CanCreateDirectories() bool
	SetCanCreateDirectories(canCreateDirectories bool)
	CanSelectHiddenExtension() bool
	SetCanSelectHiddenExtension(canSelectHiddenExtension bool)
	ShowsHiddenFiles() bool
	SetShowsHiddenFiles(showsHiddenFiles bool)
	IsExtensionHidden() bool
	SetExtensionHidden(extensionHidden bool)
	IsExpanded() bool
	AllowedContentTypes() []uti.UTType
	SetAllowedContentTypes(allowedContentTypes []uti.UTType)
	AllowsOtherFileTypes() bool
	SetAllowsOtherFileTypes(allowsOtherFileTypes bool)
	TreatsFilePackagesAsDirectories() bool
	SetTreatsFilePackagesAsDirectories(treatsFilePackagesAsDirectories bool)
	RunModal() ModalResponse
	ValidateVisibleColumns()

	BeginSheetModalForWindow(window Window, handler func(response ModalResponse))
	BeginWithCompletionHandler(handler func(response ModalResponse))
}

var _ SavePanel = (*NSSavePanel)(nil)

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

func (s *NSSavePanel) URL() foundation.URL {
	return foundation.MakeURL(C.SavePanel_URL(s.Ptr()))
}

func (s *NSSavePanel) Prompt() string {
	return C.GoString(C.SavePanel_Prompt(s.Ptr()))
}

func (s *NSSavePanel) SetPrompt(prompt string) {
	cPrompt := C.CString(prompt)
	defer C.free(unsafe.Pointer(cPrompt))
	C.SavePanel_SetPrompt(s.Ptr(), cPrompt)
}

func (s *NSSavePanel) Message() string {
	return C.GoString(C.SavePanel_Message(s.Ptr()))
}

func (s *NSSavePanel) SetMessage(message string) {
	cMessage := C.CString(message)
	defer C.free(unsafe.Pointer(cMessage))
	C.SavePanel_SetMessage(s.Ptr(), cMessage)
}

func (s *NSSavePanel) NameFieldLabel() string {
	return C.GoString(C.SavePanel_NameFieldLabel(s.Ptr()))
}

func (s *NSSavePanel) SetNameFieldLabel(nameFieldLabel string) {
	cNameFieldLabel := C.CString(nameFieldLabel)
	defer C.free(unsafe.Pointer(cNameFieldLabel))
	C.SavePanel_SetNameFieldLabel(s.Ptr(), cNameFieldLabel)
}

func (s *NSSavePanel) DirectoryURL() foundation.URL {
	return foundation.MakeURL(C.SavePanel_DirectoryURL(s.Ptr()))
}

func (s *NSSavePanel) SetDirectoryURL(directoryURL foundation.URL) {
	C.SavePanel_SetDirectoryURL(s.Ptr(), toPointer(directoryURL))
}

func (s *NSSavePanel) AccessoryView() View {
	return MakeView(C.SavePanel_AccessoryView(s.Ptr()))
}

func (s *NSSavePanel) SetAccessoryView(accessoryView View) {
	C.SavePanel_SetAccessoryView(s.Ptr(), toPointer(accessoryView))
}

func (s *NSSavePanel) ShowsTagField() bool {
	return bool(C.SavePanel_ShowsTagField(s.Ptr()))
}

func (s *NSSavePanel) SetShowsTagField(showsTagField bool) {
	C.SavePanel_SetShowsTagField(s.Ptr(), C.bool(showsTagField))
}

func (s *NSSavePanel) TagNames() []string {
	var cArray C.Array = C.SavePanel_TagNames(s.Ptr())
	defer C.free(cArray.data)
	result := (*[1 << 28]unsafe.Pointer)(unsafe.Pointer(cArray.data))[:cArray.len:cArray.len]
	var goArray = make([]string, len(result))
	for idx, r := range result {
		goArray[idx] = C.GoString((*C.char)(r))
	}
	return goArray
}

func (s *NSSavePanel) SetTagNames(tagNames []string) {
	c_tagNamesData := make([]unsafe.Pointer, len(tagNames))
	for idx, v := range tagNames {
		c_tagNamesData[idx] = unsafe.Pointer(C.CString(v))
	}
	c_tagNames := C.Array{data: unsafe.Pointer(&c_tagNamesData[0]), len: C.int(len(tagNames))}
	defer func() {
		for _, p := range c_tagNamesData {
			C.free(p)
		}
	}()
	C.SavePanel_SetTagNames(s.Ptr(), c_tagNames)
}

func (s *NSSavePanel) CanCreateDirectories() bool {
	return bool(C.SavePanel_CanCreateDirectories(s.Ptr()))
}

func (s *NSSavePanel) SetCanCreateDirectories(canCreateDirectories bool) {
	C.SavePanel_SetCanCreateDirectories(s.Ptr(), C.bool(canCreateDirectories))
}

func (s *NSSavePanel) CanSelectHiddenExtension() bool {
	return bool(C.SavePanel_CanSelectHiddenExtension(s.Ptr()))
}

func (s *NSSavePanel) SetCanSelectHiddenExtension(canSelectHiddenExtension bool) {
	C.SavePanel_SetCanSelectHiddenExtension(s.Ptr(), C.bool(canSelectHiddenExtension))
}

func (s *NSSavePanel) ShowsHiddenFiles() bool {
	return bool(C.SavePanel_ShowsHiddenFiles(s.Ptr()))
}

func (s *NSSavePanel) SetShowsHiddenFiles(showsHiddenFiles bool) {
	C.SavePanel_SetShowsHiddenFiles(s.Ptr(), C.bool(showsHiddenFiles))
}

func (s *NSSavePanel) IsExtensionHidden() bool {
	return bool(C.SavePanel_IsExtensionHidden(s.Ptr()))
}

func (s *NSSavePanel) SetExtensionHidden(extensionHidden bool) {
	C.SavePanel_SetExtensionHidden(s.Ptr(), C.bool(extensionHidden))
}

func (s *NSSavePanel) IsExpanded() bool {
	return bool(C.SavePanel_IsExpanded(s.Ptr()))
}

func (s *NSSavePanel) AllowedContentTypes() []uti.UTType {
	var cArray C.Array = C.SavePanel_AllowedContentTypes(s.Ptr())
	defer C.free(cArray.data)
	result := (*[1 << 28]unsafe.Pointer)(unsafe.Pointer(cArray.data))[:cArray.len:cArray.len]
	var goArray = make([]uti.UTType, len(result))
	for idx, r := range result {
		goArray[idx] = uti.MakeUTType(r)
	}
	return goArray
}

func (s *NSSavePanel) SetAllowedContentTypes(allowedContentTypes []uti.UTType) {
	c_allowedContentTypesData := make([]unsafe.Pointer, len(allowedContentTypes))
	for idx, v := range allowedContentTypes {
		c_allowedContentTypesData[idx] = v.Ptr()
	}
	c_allowedContentTypes := C.Array{data: unsafe.Pointer(&c_allowedContentTypesData[0]), len: C.int(len(allowedContentTypes))}
	C.SavePanel_SetAllowedContentTypes(s.Ptr(), c_allowedContentTypes)
}

func (s *NSSavePanel) AllowsOtherFileTypes() bool {
	return bool(C.SavePanel_AllowsOtherFileTypes(s.Ptr()))
}

func (s *NSSavePanel) SetAllowsOtherFileTypes(allowsOtherFileTypes bool) {
	C.SavePanel_SetAllowsOtherFileTypes(s.Ptr(), C.bool(allowsOtherFileTypes))
}

func (s *NSSavePanel) TreatsFilePackagesAsDirectories() bool {
	return bool(C.SavePanel_TreatsFilePackagesAsDirectories(s.Ptr()))
}

func (s *NSSavePanel) SetTreatsFilePackagesAsDirectories(treatsFilePackagesAsDirectories bool) {
	C.SavePanel_SetTreatsFilePackagesAsDirectories(s.Ptr(), C.bool(treatsFilePackagesAsDirectories))
}

func NewSavePanel() SavePanel {
	return MakeSavePanel(C.SavePanel_NewSavePanel())
}

func (s *NSSavePanel) RunModal() ModalResponse {
	return ModalResponse(C.SavePanel_RunModal(s.Ptr()))
}

func (s *NSSavePanel) ValidateVisibleColumns() {
	C.SavePanel_ValidateVisibleColumns(s.Ptr())
}

var completionHandlers = foundation.NewResourceRegistry()

func (s *NSSavePanel) BeginSheetModalForWindow(window Window, handler func(response ModalResponse)) {
	id := completionHandlers.NextId()
	completionHandlers.Store(id, handler)
	C.SavePanel_BeginSheetModalForWindow(s.Ptr(), window.Ptr(), C.long(id))
}

func (s *NSSavePanel) BeginWithCompletionHandler(handler func(response ModalResponse)) {
	id := completionHandlers.NextId()
	completionHandlers.Store(id, handler)
	C.SavePanel_BeginWithCompletionHandler(s.Ptr(), C.long(id))
}

//export SavePanel_CallCompletionHandler
func SavePanel_CallCompletionHandler(id int64, response int) {
	defer completionHandlers.Delete(id)
	handler := completionHandlers.Get(id).(func(response ModalResponse))
	handler(ModalResponse(response))
}
