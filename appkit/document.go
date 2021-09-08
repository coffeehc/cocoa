package appkit

// #include "document.h"
import "C"
import (
	"github.com/hsiafan/cocoa/foundation"
	"github.com/hsiafan/cocoa/objc"
	"unsafe"
)

type Document interface {
	objc.Object
	CanAsynchronouslyWriteToURL_OfType_ForSaveOperation(url foundation.URL, typeName string, saveOperation SaveOperationType) bool
	UnblockUserInteraction()
	WritableTypesForSaveOperation(saveOperation SaveOperationType) []string
	FileNameExtensionForType_SaveOperation(typeName string, saveOperation SaveOperationType) string
	MakeWindowControllers()
	AddWindowController(windowController WindowController)
	RemoveWindowController(windowController WindowController)
	WindowControllerDidLoadNib(windowController WindowController)
	WindowControllerWillLoadNib(windowController WindowController)
	ShowWindows()
	SetWindow(window Window)
	DefaultDraftName() string
	EncodeRestorableStateWithCoder_BackgroundQueue(coder foundation.Coder, queue foundation.OperationQueue)
	ScheduleAutosaving()
	BrowseDocumentVersions(sender objc.Object)
	MoveDocumentToUbiquityContainer(sender objc.Object)
	UpdateChangeCountWithToken_ForSaveOperation(changeCountToken objc.Object, saveOperation SaveOperationType)
	UpdateChangeCount(change DocumentChangeType)
	ChangeCountTokenForSaveOperation(saveOperation SaveOperationType) objc.Object
	EncodeRestorableStateWithCoder(coder foundation.Coder)
	RestoreStateWithCoder(coder foundation.Coder)
	InvalidateRestorableState()
	PrepareSavePanel(savePanel SavePanel) bool
	UpdateUserActivityState(activity foundation.UserActivity)
	ValidateUserInterfaceItem(item objc.Object) bool
	PrintDocument(sender objc.Object)
	RunPageLayout(sender objc.Object)
	RevertDocumentToSaved(sender objc.Object)
	SaveDocument(sender objc.Object)
	SaveDocumentAs(sender objc.Object)
	SaveDocumentTo(sender objc.Object)
	Close()
	DuplicateDocument(sender objc.Object)
	RenameDocument(sender objc.Object)
	MoveDocument(sender objc.Object)
	LockDocument(sender objc.Object)
	UnlockDocument(sender objc.Object)
	PreparePageLayout(pageLayout PageLayout) bool
	ShouldChangePrintInfo(newPrintInfo PrintInfo) bool
	SaveDocumentToPDF(sender objc.Object)
	PrepareSharingServicePicker(sharingServicePicker SharingServicePicker)
	HandleCloseScriptCommand(command foundation.CloseCommand) objc.Object
	HandlePrintScriptCommand(command foundation.ScriptCommand) objc.Object
	HandleSaveScriptCommand(command foundation.ScriptCommand) objc.Object
	PresentError(error foundation.Error) bool
	WillPresentError(error foundation.Error) foundation.Error
	WillNotPresentError(error foundation.Error)
	FileURL() foundation.URL
	SetFileURL(value foundation.URL)
	IsEntireFileLoaded() bool
	FileModificationDate() foundation.Date
	SetFileModificationDate(value foundation.Date)
	KeepBackupFile() bool
	IsDraft() bool
	SetDraft(value bool)
	FileType() string
	SetFileType(value string)
	IsDocumentEdited() bool
	IsInViewingMode() bool
	WindowControllers() []WindowController
	WindowNibName() NibName
	WindowForSheet() Window
	DisplayName() string
	AutosavedContentsFileURL() foundation.URL
	SetAutosavedContentsFileURL(value foundation.URL)
	AutosavingFileType() string
	AutosavingIsImplicitlyCancellable() bool
	HasUnautosavedChanges() bool
	BackupFileURL() foundation.URL
	IsBrowsingVersions() bool
	UndoManager() foundation.UndoManager
	SetUndoManager(value foundation.UndoManager)
	HasUndoManager() bool
	SetHasUndoManager(value bool)
	ShouldRunSavePanelWithAccessoryView() bool
	FileTypeFromLastRunSavePanel() string
	FileNameExtensionWasHiddenInLastRunSavePanel() bool
	UserActivity() foundation.UserActivity
	SetUserActivity(value foundation.UserActivity)
	IsLocked() bool
	PrintInfo() PrintInfo
	SetPrintInfo(value PrintInfo)
	PDFPrintOperation() PrintOperation
	AllowsDocumentSharing() bool
	ObjectSpecifier() foundation.ScriptObjectSpecifier
	LastComponentOfFileName() string
	SetLastComponentOfFileName(value string)
}

type NSDocument struct {
	objc.NSObject
}

func MakeDocument(ptr unsafe.Pointer) NSDocument {
	return NSDocument{
		NSObject: objc.MakeObject(ptr),
	}
}

func (n NSDocument) Init() NSDocument {
	result_ := C.C_NSDocument_Init(n.Ptr())
	return MakeDocument(result_)
}

func AllocDocument() NSDocument {
	result_ := C.C_NSDocument_AllocDocument()
	return MakeDocument(result_)
}

func NewDocument() NSDocument {
	result_ := C.C_NSDocument_NewDocument()
	return MakeDocument(result_)
}

func (n NSDocument) Autorelease() NSDocument {
	result_ := C.C_NSDocument_Autorelease(n.Ptr())
	return MakeDocument(result_)
}

func (n NSDocument) Retain() NSDocument {
	result_ := C.C_NSDocument_Retain(n.Ptr())
	return MakeDocument(result_)
}

func CanConcurrentlyReadDocumentsOfType(typeName string) bool {
	result_ := C.C_NSDocument_CanConcurrentlyReadDocumentsOfType(foundation.NewString(typeName).Ptr())
	return bool(result_)
}

func (n NSDocument) CanAsynchronouslyWriteToURL_OfType_ForSaveOperation(url foundation.URL, typeName string, saveOperation SaveOperationType) bool {
	result_ := C.C_NSDocument_CanAsynchronouslyWriteToURL_OfType_ForSaveOperation(n.Ptr(), objc.ExtractPtr(url), foundation.NewString(typeName).Ptr(), C.uint(uint(saveOperation)))
	return bool(result_)
}

func (n NSDocument) UnblockUserInteraction() {
	C.C_NSDocument_UnblockUserInteraction(n.Ptr())
}

func Document_IsNativeType(_type string) bool {
	result_ := C.C_NSDocument_Document_IsNativeType(foundation.NewString(_type).Ptr())
	return bool(result_)
}

func (n NSDocument) WritableTypesForSaveOperation(saveOperation SaveOperationType) []string {
	result_ := C.C_NSDocument_WritableTypesForSaveOperation(n.Ptr(), C.uint(uint(saveOperation)))
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

func (n NSDocument) FileNameExtensionForType_SaveOperation(typeName string, saveOperation SaveOperationType) string {
	result_ := C.C_NSDocument_FileNameExtensionForType_SaveOperation(n.Ptr(), foundation.NewString(typeName).Ptr(), C.uint(uint(saveOperation)))
	return foundation.MakeString(result_).String()
}

func (n NSDocument) MakeWindowControllers() {
	C.C_NSDocument_MakeWindowControllers(n.Ptr())
}

func (n NSDocument) AddWindowController(windowController WindowController) {
	C.C_NSDocument_AddWindowController(n.Ptr(), objc.ExtractPtr(windowController))
}

func (n NSDocument) RemoveWindowController(windowController WindowController) {
	C.C_NSDocument_RemoveWindowController(n.Ptr(), objc.ExtractPtr(windowController))
}

func (n NSDocument) WindowControllerDidLoadNib(windowController WindowController) {
	C.C_NSDocument_WindowControllerDidLoadNib(n.Ptr(), objc.ExtractPtr(windowController))
}

func (n NSDocument) WindowControllerWillLoadNib(windowController WindowController) {
	C.C_NSDocument_WindowControllerWillLoadNib(n.Ptr(), objc.ExtractPtr(windowController))
}

func (n NSDocument) ShowWindows() {
	C.C_NSDocument_ShowWindows(n.Ptr())
}

func (n NSDocument) SetWindow(window Window) {
	C.C_NSDocument_SetWindow(n.Ptr(), objc.ExtractPtr(window))
}

func (n NSDocument) DefaultDraftName() string {
	result_ := C.C_NSDocument_DefaultDraftName(n.Ptr())
	return foundation.MakeString(result_).String()
}

func (n NSDocument) EncodeRestorableStateWithCoder_BackgroundQueue(coder foundation.Coder, queue foundation.OperationQueue) {
	C.C_NSDocument_EncodeRestorableStateWithCoder_BackgroundQueue(n.Ptr(), objc.ExtractPtr(coder), objc.ExtractPtr(queue))
}

func (n NSDocument) ScheduleAutosaving() {
	C.C_NSDocument_ScheduleAutosaving(n.Ptr())
}

func (n NSDocument) BrowseDocumentVersions(sender objc.Object) {
	C.C_NSDocument_BrowseDocumentVersions(n.Ptr(), objc.ExtractPtr(sender))
}

func (n NSDocument) MoveDocumentToUbiquityContainer(sender objc.Object) {
	C.C_NSDocument_MoveDocumentToUbiquityContainer(n.Ptr(), objc.ExtractPtr(sender))
}

func (n NSDocument) UpdateChangeCountWithToken_ForSaveOperation(changeCountToken objc.Object, saveOperation SaveOperationType) {
	C.C_NSDocument_UpdateChangeCountWithToken_ForSaveOperation(n.Ptr(), objc.ExtractPtr(changeCountToken), C.uint(uint(saveOperation)))
}

func (n NSDocument) UpdateChangeCount(change DocumentChangeType) {
	C.C_NSDocument_UpdateChangeCount(n.Ptr(), C.uint(uint(change)))
}

func (n NSDocument) ChangeCountTokenForSaveOperation(saveOperation SaveOperationType) objc.Object {
	result_ := C.C_NSDocument_ChangeCountTokenForSaveOperation(n.Ptr(), C.uint(uint(saveOperation)))
	return objc.MakeObject(result_)
}

func (n NSDocument) EncodeRestorableStateWithCoder(coder foundation.Coder) {
	C.C_NSDocument_EncodeRestorableStateWithCoder(n.Ptr(), objc.ExtractPtr(coder))
}

func (n NSDocument) RestoreStateWithCoder(coder foundation.Coder) {
	C.C_NSDocument_RestoreStateWithCoder(n.Ptr(), objc.ExtractPtr(coder))
}

func (n NSDocument) InvalidateRestorableState() {
	C.C_NSDocument_InvalidateRestorableState(n.Ptr())
}

func (n NSDocument) PrepareSavePanel(savePanel SavePanel) bool {
	result_ := C.C_NSDocument_PrepareSavePanel(n.Ptr(), objc.ExtractPtr(savePanel))
	return bool(result_)
}

func (n NSDocument) UpdateUserActivityState(activity foundation.UserActivity) {
	C.C_NSDocument_UpdateUserActivityState(n.Ptr(), objc.ExtractPtr(activity))
}

func (n NSDocument) ValidateUserInterfaceItem(item objc.Object) bool {
	result_ := C.C_NSDocument_ValidateUserInterfaceItem(n.Ptr(), objc.ExtractPtr(item))
	return bool(result_)
}

func (n NSDocument) PrintDocument(sender objc.Object) {
	C.C_NSDocument_PrintDocument(n.Ptr(), objc.ExtractPtr(sender))
}

func (n NSDocument) RunPageLayout(sender objc.Object) {
	C.C_NSDocument_RunPageLayout(n.Ptr(), objc.ExtractPtr(sender))
}

func (n NSDocument) RevertDocumentToSaved(sender objc.Object) {
	C.C_NSDocument_RevertDocumentToSaved(n.Ptr(), objc.ExtractPtr(sender))
}

func (n NSDocument) SaveDocument(sender objc.Object) {
	C.C_NSDocument_SaveDocument(n.Ptr(), objc.ExtractPtr(sender))
}

func (n NSDocument) SaveDocumentAs(sender objc.Object) {
	C.C_NSDocument_SaveDocumentAs(n.Ptr(), objc.ExtractPtr(sender))
}

func (n NSDocument) SaveDocumentTo(sender objc.Object) {
	C.C_NSDocument_SaveDocumentTo(n.Ptr(), objc.ExtractPtr(sender))
}

func (n NSDocument) Close() {
	C.C_NSDocument_Close(n.Ptr())
}

func (n NSDocument) DuplicateDocument(sender objc.Object) {
	C.C_NSDocument_DuplicateDocument(n.Ptr(), objc.ExtractPtr(sender))
}

func (n NSDocument) RenameDocument(sender objc.Object) {
	C.C_NSDocument_RenameDocument(n.Ptr(), objc.ExtractPtr(sender))
}

func (n NSDocument) MoveDocument(sender objc.Object) {
	C.C_NSDocument_MoveDocument(n.Ptr(), objc.ExtractPtr(sender))
}

func (n NSDocument) LockDocument(sender objc.Object) {
	C.C_NSDocument_LockDocument(n.Ptr(), objc.ExtractPtr(sender))
}

func (n NSDocument) UnlockDocument(sender objc.Object) {
	C.C_NSDocument_UnlockDocument(n.Ptr(), objc.ExtractPtr(sender))
}

func (n NSDocument) PreparePageLayout(pageLayout PageLayout) bool {
	result_ := C.C_NSDocument_PreparePageLayout(n.Ptr(), objc.ExtractPtr(pageLayout))
	return bool(result_)
}

func (n NSDocument) ShouldChangePrintInfo(newPrintInfo PrintInfo) bool {
	result_ := C.C_NSDocument_ShouldChangePrintInfo(n.Ptr(), objc.ExtractPtr(newPrintInfo))
	return bool(result_)
}

func (n NSDocument) SaveDocumentToPDF(sender objc.Object) {
	C.C_NSDocument_SaveDocumentToPDF(n.Ptr(), objc.ExtractPtr(sender))
}

func (n NSDocument) PrepareSharingServicePicker(sharingServicePicker SharingServicePicker) {
	C.C_NSDocument_PrepareSharingServicePicker(n.Ptr(), objc.ExtractPtr(sharingServicePicker))
}

func (n NSDocument) HandleCloseScriptCommand(command foundation.CloseCommand) objc.Object {
	result_ := C.C_NSDocument_HandleCloseScriptCommand(n.Ptr(), objc.ExtractPtr(command))
	return objc.MakeObject(result_)
}

func (n NSDocument) HandlePrintScriptCommand(command foundation.ScriptCommand) objc.Object {
	result_ := C.C_NSDocument_HandlePrintScriptCommand(n.Ptr(), objc.ExtractPtr(command))
	return objc.MakeObject(result_)
}

func (n NSDocument) HandleSaveScriptCommand(command foundation.ScriptCommand) objc.Object {
	result_ := C.C_NSDocument_HandleSaveScriptCommand(n.Ptr(), objc.ExtractPtr(command))
	return objc.MakeObject(result_)
}

func (n NSDocument) PresentError(error foundation.Error) bool {
	result_ := C.C_NSDocument_PresentError(n.Ptr(), objc.ExtractPtr(error))
	return bool(result_)
}

func (n NSDocument) WillPresentError(error foundation.Error) foundation.Error {
	result_ := C.C_NSDocument_WillPresentError(n.Ptr(), objc.ExtractPtr(error))
	return foundation.MakeError(result_)
}

func (n NSDocument) WillNotPresentError(error foundation.Error) {
	C.C_NSDocument_WillNotPresentError(n.Ptr(), objc.ExtractPtr(error))
}

func (n NSDocument) FileURL() foundation.URL {
	result_ := C.C_NSDocument_FileURL(n.Ptr())
	return foundation.MakeURL(result_)
}

func (n NSDocument) SetFileURL(value foundation.URL) {
	C.C_NSDocument_SetFileURL(n.Ptr(), objc.ExtractPtr(value))
}

func (n NSDocument) IsEntireFileLoaded() bool {
	result_ := C.C_NSDocument_IsEntireFileLoaded(n.Ptr())
	return bool(result_)
}

func (n NSDocument) FileModificationDate() foundation.Date {
	result_ := C.C_NSDocument_FileModificationDate(n.Ptr())
	return foundation.MakeDate(result_)
}

func (n NSDocument) SetFileModificationDate(value foundation.Date) {
	C.C_NSDocument_SetFileModificationDate(n.Ptr(), objc.ExtractPtr(value))
}

func (n NSDocument) KeepBackupFile() bool {
	result_ := C.C_NSDocument_KeepBackupFile(n.Ptr())
	return bool(result_)
}

func (n NSDocument) IsDraft() bool {
	result_ := C.C_NSDocument_IsDraft(n.Ptr())
	return bool(result_)
}

func (n NSDocument) SetDraft(value bool) {
	C.C_NSDocument_SetDraft(n.Ptr(), C.bool(value))
}

func (n NSDocument) FileType() string {
	result_ := C.C_NSDocument_FileType(n.Ptr())
	return foundation.MakeString(result_).String()
}

func (n NSDocument) SetFileType(value string) {
	C.C_NSDocument_SetFileType(n.Ptr(), foundation.NewString(value).Ptr())
}

func (n NSDocument) IsDocumentEdited() bool {
	result_ := C.C_NSDocument_IsDocumentEdited(n.Ptr())
	return bool(result_)
}

func (n NSDocument) IsInViewingMode() bool {
	result_ := C.C_NSDocument_IsInViewingMode(n.Ptr())
	return bool(result_)
}

func Document_ReadableTypes() []string {
	result_ := C.C_NSDocument_Document_ReadableTypes()
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

func Document_WritableTypes() []string {
	result_ := C.C_NSDocument_Document_WritableTypes()
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

func (n NSDocument) WindowControllers() []WindowController {
	result_ := C.C_NSDocument_WindowControllers(n.Ptr())
	if result_.len > 0 {
		defer C.free(result_.data)
	}
	result_Slice := unsafe.Slice((*unsafe.Pointer)(result_.data), int(result_.len))
	var goResult_ = make([]WindowController, len(result_Slice))
	for idx, r := range result_Slice {
		goResult_[idx] = MakeWindowController(r)
	}
	return goResult_
}

func (n NSDocument) WindowNibName() NibName {
	result_ := C.C_NSDocument_WindowNibName(n.Ptr())
	return NibName(foundation.MakeString(result_).String())
}

func (n NSDocument) WindowForSheet() Window {
	result_ := C.C_NSDocument_WindowForSheet(n.Ptr())
	return MakeWindow(result_)
}

func (n NSDocument) DisplayName() string {
	result_ := C.C_NSDocument_DisplayName(n.Ptr())
	return foundation.MakeString(result_).String()
}

func Document_AutosavesInPlace() bool {
	result_ := C.C_NSDocument_Document_AutosavesInPlace()
	return bool(result_)
}

func Document_AutosavesDrafts() bool {
	result_ := C.C_NSDocument_Document_AutosavesDrafts()
	return bool(result_)
}

func Document_PreservesVersions() bool {
	result_ := C.C_NSDocument_Document_PreservesVersions()
	return bool(result_)
}

func (n NSDocument) AutosavedContentsFileURL() foundation.URL {
	result_ := C.C_NSDocument_AutosavedContentsFileURL(n.Ptr())
	return foundation.MakeURL(result_)
}

func (n NSDocument) SetAutosavedContentsFileURL(value foundation.URL) {
	C.C_NSDocument_SetAutosavedContentsFileURL(n.Ptr(), objc.ExtractPtr(value))
}

func (n NSDocument) AutosavingFileType() string {
	result_ := C.C_NSDocument_AutosavingFileType(n.Ptr())
	return foundation.MakeString(result_).String()
}

func (n NSDocument) AutosavingIsImplicitlyCancellable() bool {
	result_ := C.C_NSDocument_AutosavingIsImplicitlyCancellable(n.Ptr())
	return bool(result_)
}

func (n NSDocument) HasUnautosavedChanges() bool {
	result_ := C.C_NSDocument_HasUnautosavedChanges(n.Ptr())
	return bool(result_)
}

func (n NSDocument) BackupFileURL() foundation.URL {
	result_ := C.C_NSDocument_BackupFileURL(n.Ptr())
	return foundation.MakeURL(result_)
}

func (n NSDocument) IsBrowsingVersions() bool {
	result_ := C.C_NSDocument_IsBrowsingVersions(n.Ptr())
	return bool(result_)
}

func Document_UsesUbiquitousStorage() bool {
	result_ := C.C_NSDocument_Document_UsesUbiquitousStorage()
	return bool(result_)
}

func (n NSDocument) UndoManager() foundation.UndoManager {
	result_ := C.C_NSDocument_UndoManager(n.Ptr())
	return foundation.MakeUndoManager(result_)
}

func (n NSDocument) SetUndoManager(value foundation.UndoManager) {
	C.C_NSDocument_SetUndoManager(n.Ptr(), objc.ExtractPtr(value))
}

func (n NSDocument) HasUndoManager() bool {
	result_ := C.C_NSDocument_HasUndoManager(n.Ptr())
	return bool(result_)
}

func (n NSDocument) SetHasUndoManager(value bool) {
	C.C_NSDocument_SetHasUndoManager(n.Ptr(), C.bool(value))
}

func Document_RestorableStateKeyPaths() []string {
	result_ := C.C_NSDocument_Document_RestorableStateKeyPaths()
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

func (n NSDocument) ShouldRunSavePanelWithAccessoryView() bool {
	result_ := C.C_NSDocument_ShouldRunSavePanelWithAccessoryView(n.Ptr())
	return bool(result_)
}

func (n NSDocument) FileTypeFromLastRunSavePanel() string {
	result_ := C.C_NSDocument_FileTypeFromLastRunSavePanel(n.Ptr())
	return foundation.MakeString(result_).String()
}

func (n NSDocument) FileNameExtensionWasHiddenInLastRunSavePanel() bool {
	result_ := C.C_NSDocument_FileNameExtensionWasHiddenInLastRunSavePanel(n.Ptr())
	return bool(result_)
}

func (n NSDocument) UserActivity() foundation.UserActivity {
	result_ := C.C_NSDocument_UserActivity(n.Ptr())
	return foundation.MakeUserActivity(result_)
}

func (n NSDocument) SetUserActivity(value foundation.UserActivity) {
	C.C_NSDocument_SetUserActivity(n.Ptr(), objc.ExtractPtr(value))
}

func (n NSDocument) IsLocked() bool {
	result_ := C.C_NSDocument_IsLocked(n.Ptr())
	return bool(result_)
}

func (n NSDocument) PrintInfo() PrintInfo {
	result_ := C.C_NSDocument_PrintInfo(n.Ptr())
	return MakePrintInfo(result_)
}

func (n NSDocument) SetPrintInfo(value PrintInfo) {
	C.C_NSDocument_SetPrintInfo(n.Ptr(), objc.ExtractPtr(value))
}

func (n NSDocument) PDFPrintOperation() PrintOperation {
	result_ := C.C_NSDocument_PDFPrintOperation(n.Ptr())
	return MakePrintOperation(result_)
}

func (n NSDocument) AllowsDocumentSharing() bool {
	result_ := C.C_NSDocument_AllowsDocumentSharing(n.Ptr())
	return bool(result_)
}

func (n NSDocument) ObjectSpecifier() foundation.ScriptObjectSpecifier {
	result_ := C.C_NSDocument_ObjectSpecifier(n.Ptr())
	return foundation.MakeScriptObjectSpecifier(result_)
}

func (n NSDocument) LastComponentOfFileName() string {
	result_ := C.C_NSDocument_LastComponentOfFileName(n.Ptr())
	return foundation.MakeString(result_).String()
}

func (n NSDocument) SetLastComponentOfFileName(value string) {
	C.C_NSDocument_SetLastComponentOfFileName(n.Ptr(), foundation.NewString(value).Ptr())
}
