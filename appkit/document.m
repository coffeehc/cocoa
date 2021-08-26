#import <Appkit/Appkit.h>
#import "document.h"

void* C_Document_Alloc() {
    return [NSDocument alloc];
}

void* C_NSDocument_Init(void* ptr) {
    NSDocument* nSDocument = (NSDocument*)ptr;
    NSDocument* result_ = [nSDocument init];
    return result_;
}

bool C_NSDocument_Document_CanConcurrentlyReadDocumentsOfType(void* typeName) {
    BOOL result_ = [NSDocument canConcurrentlyReadDocumentsOfType:(NSString*)typeName];
    return result_;
}

bool C_NSDocument_CanAsynchronouslyWriteToURL_OfType_ForSaveOperation(void* ptr, void* url, void* typeName, unsigned int saveOperation) {
    NSDocument* nSDocument = (NSDocument*)ptr;
    BOOL result_ = [nSDocument canAsynchronouslyWriteToURL:(NSURL*)url ofType:(NSString*)typeName forSaveOperation:saveOperation];
    return result_;
}

void C_NSDocument_UnblockUserInteraction(void* ptr) {
    NSDocument* nSDocument = (NSDocument*)ptr;
    [nSDocument unblockUserInteraction];
}

bool C_NSDocument_Document_IsNativeType(void* _type) {
    BOOL result_ = [NSDocument isNativeType:(NSString*)_type];
    return result_;
}

Array C_NSDocument_WritableTypesForSaveOperation(void* ptr, unsigned int saveOperation) {
    NSDocument* nSDocument = (NSDocument*)ptr;
    NSArray* result_ = [nSDocument writableTypesForSaveOperation:saveOperation];
    Array result_Array;
    int result_count = [result_ count];
    if (result_count > 0) {
    	void** result_Data = malloc(result_count * sizeof(void*));
    	for (int i = 0; i < result_count; i++) {
    		 void* p = [result_ objectAtIndex:i];
    		 result_Data[i] = p;
    	}
    	result_Array.data = result_Data;
    	result_Array.len = result_count;
    }
    return result_Array;
}

void* C_NSDocument_FileNameExtensionForType_SaveOperation(void* ptr, void* typeName, unsigned int saveOperation) {
    NSDocument* nSDocument = (NSDocument*)ptr;
    NSString* result_ = [nSDocument fileNameExtensionForType:(NSString*)typeName saveOperation:saveOperation];
    return result_;
}

void C_NSDocument_MakeWindowControllers(void* ptr) {
    NSDocument* nSDocument = (NSDocument*)ptr;
    [nSDocument makeWindowControllers];
}

void C_NSDocument_AddWindowController(void* ptr, void* windowController) {
    NSDocument* nSDocument = (NSDocument*)ptr;
    [nSDocument addWindowController:(NSWindowController*)windowController];
}

void C_NSDocument_RemoveWindowController(void* ptr, void* windowController) {
    NSDocument* nSDocument = (NSDocument*)ptr;
    [nSDocument removeWindowController:(NSWindowController*)windowController];
}

void C_NSDocument_WindowControllerDidLoadNib(void* ptr, void* windowController) {
    NSDocument* nSDocument = (NSDocument*)ptr;
    [nSDocument windowControllerDidLoadNib:(NSWindowController*)windowController];
}

void C_NSDocument_WindowControllerWillLoadNib(void* ptr, void* windowController) {
    NSDocument* nSDocument = (NSDocument*)ptr;
    [nSDocument windowControllerWillLoadNib:(NSWindowController*)windowController];
}

void C_NSDocument_ShowWindows(void* ptr) {
    NSDocument* nSDocument = (NSDocument*)ptr;
    [nSDocument showWindows];
}

void C_NSDocument_SetWindow(void* ptr, void* window) {
    NSDocument* nSDocument = (NSDocument*)ptr;
    [nSDocument setWindow:(NSWindow*)window];
}

void* C_NSDocument_DefaultDraftName(void* ptr) {
    NSDocument* nSDocument = (NSDocument*)ptr;
    NSString* result_ = [nSDocument defaultDraftName];
    return result_;
}

void C_NSDocument_EncodeRestorableStateWithCoder_BackgroundQueue(void* ptr, void* coder, void* queue) {
    NSDocument* nSDocument = (NSDocument*)ptr;
    [nSDocument encodeRestorableStateWithCoder:(NSCoder*)coder backgroundQueue:(NSOperationQueue*)queue];
}

void C_NSDocument_ScheduleAutosaving(void* ptr) {
    NSDocument* nSDocument = (NSDocument*)ptr;
    [nSDocument scheduleAutosaving];
}

void C_NSDocument_BrowseDocumentVersions(void* ptr, void* sender) {
    NSDocument* nSDocument = (NSDocument*)ptr;
    [nSDocument browseDocumentVersions:(id)sender];
}

void C_NSDocument_MoveDocumentToUbiquityContainer(void* ptr, void* sender) {
    NSDocument* nSDocument = (NSDocument*)ptr;
    [nSDocument moveDocumentToUbiquityContainer:(id)sender];
}

void C_NSDocument_UpdateChangeCountWithToken_ForSaveOperation(void* ptr, void* changeCountToken, unsigned int saveOperation) {
    NSDocument* nSDocument = (NSDocument*)ptr;
    [nSDocument updateChangeCountWithToken:(id)changeCountToken forSaveOperation:saveOperation];
}

void C_NSDocument_UpdateChangeCount(void* ptr, unsigned int change) {
    NSDocument* nSDocument = (NSDocument*)ptr;
    [nSDocument updateChangeCount:change];
}

void* C_NSDocument_ChangeCountTokenForSaveOperation(void* ptr, unsigned int saveOperation) {
    NSDocument* nSDocument = (NSDocument*)ptr;
    id result_ = [nSDocument changeCountTokenForSaveOperation:saveOperation];
    return result_;
}

void C_NSDocument_EncodeRestorableStateWithCoder(void* ptr, void* coder) {
    NSDocument* nSDocument = (NSDocument*)ptr;
    [nSDocument encodeRestorableStateWithCoder:(NSCoder*)coder];
}

void C_NSDocument_RestoreStateWithCoder(void* ptr, void* coder) {
    NSDocument* nSDocument = (NSDocument*)ptr;
    [nSDocument restoreStateWithCoder:(NSCoder*)coder];
}

void C_NSDocument_InvalidateRestorableState(void* ptr) {
    NSDocument* nSDocument = (NSDocument*)ptr;
    [nSDocument invalidateRestorableState];
}

bool C_NSDocument_PrepareSavePanel(void* ptr, void* savePanel) {
    NSDocument* nSDocument = (NSDocument*)ptr;
    BOOL result_ = [nSDocument prepareSavePanel:(NSSavePanel*)savePanel];
    return result_;
}

void C_NSDocument_UpdateUserActivityState(void* ptr, void* activity) {
    NSDocument* nSDocument = (NSDocument*)ptr;
    [nSDocument updateUserActivityState:(NSUserActivity*)activity];
}

bool C_NSDocument_ValidateUserInterfaceItem(void* ptr, void* item) {
    NSDocument* nSDocument = (NSDocument*)ptr;
    BOOL result_ = [nSDocument validateUserInterfaceItem:(id)item];
    return result_;
}

void C_NSDocument_PrintDocument(void* ptr, void* sender) {
    NSDocument* nSDocument = (NSDocument*)ptr;
    [nSDocument printDocument:(id)sender];
}

void C_NSDocument_RunPageLayout(void* ptr, void* sender) {
    NSDocument* nSDocument = (NSDocument*)ptr;
    [nSDocument runPageLayout:(id)sender];
}

void C_NSDocument_RevertDocumentToSaved(void* ptr, void* sender) {
    NSDocument* nSDocument = (NSDocument*)ptr;
    [nSDocument revertDocumentToSaved:(id)sender];
}

void C_NSDocument_SaveDocument(void* ptr, void* sender) {
    NSDocument* nSDocument = (NSDocument*)ptr;
    [nSDocument saveDocument:(id)sender];
}

void C_NSDocument_SaveDocumentAs(void* ptr, void* sender) {
    NSDocument* nSDocument = (NSDocument*)ptr;
    [nSDocument saveDocumentAs:(id)sender];
}

void C_NSDocument_SaveDocumentTo(void* ptr, void* sender) {
    NSDocument* nSDocument = (NSDocument*)ptr;
    [nSDocument saveDocumentTo:(id)sender];
}

void C_NSDocument_Close(void* ptr) {
    NSDocument* nSDocument = (NSDocument*)ptr;
    [nSDocument close];
}

void C_NSDocument_DuplicateDocument(void* ptr, void* sender) {
    NSDocument* nSDocument = (NSDocument*)ptr;
    [nSDocument duplicateDocument:(id)sender];
}

void C_NSDocument_RenameDocument(void* ptr, void* sender) {
    NSDocument* nSDocument = (NSDocument*)ptr;
    [nSDocument renameDocument:(id)sender];
}

void C_NSDocument_MoveDocument(void* ptr, void* sender) {
    NSDocument* nSDocument = (NSDocument*)ptr;
    [nSDocument moveDocument:(id)sender];
}

void C_NSDocument_LockDocument(void* ptr, void* sender) {
    NSDocument* nSDocument = (NSDocument*)ptr;
    [nSDocument lockDocument:(id)sender];
}

void C_NSDocument_UnlockDocument(void* ptr, void* sender) {
    NSDocument* nSDocument = (NSDocument*)ptr;
    [nSDocument unlockDocument:(id)sender];
}

bool C_NSDocument_PreparePageLayout(void* ptr, void* pageLayout) {
    NSDocument* nSDocument = (NSDocument*)ptr;
    BOOL result_ = [nSDocument preparePageLayout:(NSPageLayout*)pageLayout];
    return result_;
}

bool C_NSDocument_ShouldChangePrintInfo(void* ptr, void* newPrintInfo) {
    NSDocument* nSDocument = (NSDocument*)ptr;
    BOOL result_ = [nSDocument shouldChangePrintInfo:(NSPrintInfo*)newPrintInfo];
    return result_;
}

void C_NSDocument_SaveDocumentToPDF(void* ptr, void* sender) {
    NSDocument* nSDocument = (NSDocument*)ptr;
    [nSDocument saveDocumentToPDF:(id)sender];
}

void C_NSDocument_PrepareSharingServicePicker(void* ptr, void* sharingServicePicker) {
    NSDocument* nSDocument = (NSDocument*)ptr;
    [nSDocument prepareSharingServicePicker:(NSSharingServicePicker*)sharingServicePicker];
}

void* C_NSDocument_HandleCloseScriptCommand(void* ptr, void* command) {
    NSDocument* nSDocument = (NSDocument*)ptr;
    id result_ = [nSDocument handleCloseScriptCommand:(NSCloseCommand*)command];
    return result_;
}

void* C_NSDocument_HandlePrintScriptCommand(void* ptr, void* command) {
    NSDocument* nSDocument = (NSDocument*)ptr;
    id result_ = [nSDocument handlePrintScriptCommand:(NSScriptCommand*)command];
    return result_;
}

void* C_NSDocument_HandleSaveScriptCommand(void* ptr, void* command) {
    NSDocument* nSDocument = (NSDocument*)ptr;
    id result_ = [nSDocument handleSaveScriptCommand:(NSScriptCommand*)command];
    return result_;
}

bool C_NSDocument_PresentError(void* ptr, void* error) {
    NSDocument* nSDocument = (NSDocument*)ptr;
    BOOL result_ = [nSDocument presentError:(NSError*)error];
    return result_;
}

void* C_NSDocument_WillPresentError(void* ptr, void* error) {
    NSDocument* nSDocument = (NSDocument*)ptr;
    NSError* result_ = [nSDocument willPresentError:(NSError*)error];
    return result_;
}

void C_NSDocument_WillNotPresentError(void* ptr, void* error) {
    NSDocument* nSDocument = (NSDocument*)ptr;
    [nSDocument willNotPresentError:(NSError*)error];
}

void* C_NSDocument_FileURL(void* ptr) {
    NSDocument* nSDocument = (NSDocument*)ptr;
    NSURL* result_ = [nSDocument fileURL];
    return result_;
}

void C_NSDocument_SetFileURL(void* ptr, void* value) {
    NSDocument* nSDocument = (NSDocument*)ptr;
    [nSDocument setFileURL:(NSURL*)value];
}

bool C_NSDocument_IsEntireFileLoaded(void* ptr) {
    NSDocument* nSDocument = (NSDocument*)ptr;
    BOOL result_ = [nSDocument isEntireFileLoaded];
    return result_;
}

void* C_NSDocument_FileModificationDate(void* ptr) {
    NSDocument* nSDocument = (NSDocument*)ptr;
    NSDate* result_ = [nSDocument fileModificationDate];
    return result_;
}

void C_NSDocument_SetFileModificationDate(void* ptr, void* value) {
    NSDocument* nSDocument = (NSDocument*)ptr;
    [nSDocument setFileModificationDate:(NSDate*)value];
}

bool C_NSDocument_KeepBackupFile(void* ptr) {
    NSDocument* nSDocument = (NSDocument*)ptr;
    BOOL result_ = [nSDocument keepBackupFile];
    return result_;
}

bool C_NSDocument_IsDraft(void* ptr) {
    NSDocument* nSDocument = (NSDocument*)ptr;
    BOOL result_ = [nSDocument isDraft];
    return result_;
}

void C_NSDocument_SetDraft(void* ptr, bool value) {
    NSDocument* nSDocument = (NSDocument*)ptr;
    [nSDocument setDraft:value];
}

void* C_NSDocument_FileType(void* ptr) {
    NSDocument* nSDocument = (NSDocument*)ptr;
    NSString* result_ = [nSDocument fileType];
    return result_;
}

void C_NSDocument_SetFileType(void* ptr, void* value) {
    NSDocument* nSDocument = (NSDocument*)ptr;
    [nSDocument setFileType:(NSString*)value];
}

bool C_NSDocument_IsDocumentEdited(void* ptr) {
    NSDocument* nSDocument = (NSDocument*)ptr;
    BOOL result_ = [nSDocument isDocumentEdited];
    return result_;
}

bool C_NSDocument_IsInViewingMode(void* ptr) {
    NSDocument* nSDocument = (NSDocument*)ptr;
    BOOL result_ = [nSDocument isInViewingMode];
    return result_;
}

Array C_NSDocument_Document_ReadableTypes() {
    NSArray* result_ = [NSDocument readableTypes];
    Array result_Array;
    int result_count = [result_ count];
    if (result_count > 0) {
    	void** result_Data = malloc(result_count * sizeof(void*));
    	for (int i = 0; i < result_count; i++) {
    		 void* p = [result_ objectAtIndex:i];
    		 result_Data[i] = p;
    	}
    	result_Array.data = result_Data;
    	result_Array.len = result_count;
    }
    return result_Array;
}

Array C_NSDocument_Document_WritableTypes() {
    NSArray* result_ = [NSDocument writableTypes];
    Array result_Array;
    int result_count = [result_ count];
    if (result_count > 0) {
    	void** result_Data = malloc(result_count * sizeof(void*));
    	for (int i = 0; i < result_count; i++) {
    		 void* p = [result_ objectAtIndex:i];
    		 result_Data[i] = p;
    	}
    	result_Array.data = result_Data;
    	result_Array.len = result_count;
    }
    return result_Array;
}

Array C_NSDocument_WindowControllers(void* ptr) {
    NSDocument* nSDocument = (NSDocument*)ptr;
    NSArray* result_ = [nSDocument windowControllers];
    Array result_Array;
    int result_count = [result_ count];
    if (result_count > 0) {
    	void** result_Data = malloc(result_count * sizeof(void*));
    	for (int i = 0; i < result_count; i++) {
    		 void* p = [result_ objectAtIndex:i];
    		 result_Data[i] = p;
    	}
    	result_Array.data = result_Data;
    	result_Array.len = result_count;
    }
    return result_Array;
}

void* C_NSDocument_WindowNibName(void* ptr) {
    NSDocument* nSDocument = (NSDocument*)ptr;
    NSNibName result_ = [nSDocument windowNibName];
    return result_;
}

void* C_NSDocument_WindowForSheet(void* ptr) {
    NSDocument* nSDocument = (NSDocument*)ptr;
    NSWindow* result_ = [nSDocument windowForSheet];
    return result_;
}

void* C_NSDocument_DisplayName(void* ptr) {
    NSDocument* nSDocument = (NSDocument*)ptr;
    NSString* result_ = [nSDocument displayName];
    return result_;
}

bool C_NSDocument_Document_AutosavesInPlace() {
    BOOL result_ = [NSDocument autosavesInPlace];
    return result_;
}

bool C_NSDocument_Document_AutosavesDrafts() {
    BOOL result_ = [NSDocument autosavesDrafts];
    return result_;
}

bool C_NSDocument_Document_PreservesVersions() {
    BOOL result_ = [NSDocument preservesVersions];
    return result_;
}

void* C_NSDocument_AutosavedContentsFileURL(void* ptr) {
    NSDocument* nSDocument = (NSDocument*)ptr;
    NSURL* result_ = [nSDocument autosavedContentsFileURL];
    return result_;
}

void C_NSDocument_SetAutosavedContentsFileURL(void* ptr, void* value) {
    NSDocument* nSDocument = (NSDocument*)ptr;
    [nSDocument setAutosavedContentsFileURL:(NSURL*)value];
}

void* C_NSDocument_AutosavingFileType(void* ptr) {
    NSDocument* nSDocument = (NSDocument*)ptr;
    NSString* result_ = [nSDocument autosavingFileType];
    return result_;
}

bool C_NSDocument_AutosavingIsImplicitlyCancellable(void* ptr) {
    NSDocument* nSDocument = (NSDocument*)ptr;
    BOOL result_ = [nSDocument autosavingIsImplicitlyCancellable];
    return result_;
}

bool C_NSDocument_HasUnautosavedChanges(void* ptr) {
    NSDocument* nSDocument = (NSDocument*)ptr;
    BOOL result_ = [nSDocument hasUnautosavedChanges];
    return result_;
}

void* C_NSDocument_BackupFileURL(void* ptr) {
    NSDocument* nSDocument = (NSDocument*)ptr;
    NSURL* result_ = [nSDocument backupFileURL];
    return result_;
}

bool C_NSDocument_IsBrowsingVersions(void* ptr) {
    NSDocument* nSDocument = (NSDocument*)ptr;
    BOOL result_ = [nSDocument isBrowsingVersions];
    return result_;
}

bool C_NSDocument_Document_UsesUbiquitousStorage() {
    BOOL result_ = [NSDocument usesUbiquitousStorage];
    return result_;
}

void* C_NSDocument_UndoManager(void* ptr) {
    NSDocument* nSDocument = (NSDocument*)ptr;
    NSUndoManager* result_ = [nSDocument undoManager];
    return result_;
}

void C_NSDocument_SetUndoManager(void* ptr, void* value) {
    NSDocument* nSDocument = (NSDocument*)ptr;
    [nSDocument setUndoManager:(NSUndoManager*)value];
}

bool C_NSDocument_HasUndoManager(void* ptr) {
    NSDocument* nSDocument = (NSDocument*)ptr;
    BOOL result_ = [nSDocument hasUndoManager];
    return result_;
}

void C_NSDocument_SetHasUndoManager(void* ptr, bool value) {
    NSDocument* nSDocument = (NSDocument*)ptr;
    [nSDocument setHasUndoManager:value];
}

Array C_NSDocument_Document_RestorableStateKeyPaths() {
    NSArray* result_ = [NSDocument restorableStateKeyPaths];
    Array result_Array;
    int result_count = [result_ count];
    if (result_count > 0) {
    	void** result_Data = malloc(result_count * sizeof(void*));
    	for (int i = 0; i < result_count; i++) {
    		 void* p = [result_ objectAtIndex:i];
    		 result_Data[i] = p;
    	}
    	result_Array.data = result_Data;
    	result_Array.len = result_count;
    }
    return result_Array;
}

bool C_NSDocument_ShouldRunSavePanelWithAccessoryView(void* ptr) {
    NSDocument* nSDocument = (NSDocument*)ptr;
    BOOL result_ = [nSDocument shouldRunSavePanelWithAccessoryView];
    return result_;
}

void* C_NSDocument_FileTypeFromLastRunSavePanel(void* ptr) {
    NSDocument* nSDocument = (NSDocument*)ptr;
    NSString* result_ = [nSDocument fileTypeFromLastRunSavePanel];
    return result_;
}

bool C_NSDocument_FileNameExtensionWasHiddenInLastRunSavePanel(void* ptr) {
    NSDocument* nSDocument = (NSDocument*)ptr;
    BOOL result_ = [nSDocument fileNameExtensionWasHiddenInLastRunSavePanel];
    return result_;
}

void* C_NSDocument_UserActivity(void* ptr) {
    NSDocument* nSDocument = (NSDocument*)ptr;
    NSUserActivity* result_ = [nSDocument userActivity];
    return result_;
}

void C_NSDocument_SetUserActivity(void* ptr, void* value) {
    NSDocument* nSDocument = (NSDocument*)ptr;
    [nSDocument setUserActivity:(NSUserActivity*)value];
}

bool C_NSDocument_IsLocked(void* ptr) {
    NSDocument* nSDocument = (NSDocument*)ptr;
    BOOL result_ = [nSDocument isLocked];
    return result_;
}

void* C_NSDocument_PrintInfo(void* ptr) {
    NSDocument* nSDocument = (NSDocument*)ptr;
    NSPrintInfo* result_ = [nSDocument printInfo];
    return result_;
}

void C_NSDocument_SetPrintInfo(void* ptr, void* value) {
    NSDocument* nSDocument = (NSDocument*)ptr;
    [nSDocument setPrintInfo:(NSPrintInfo*)value];
}

void* C_NSDocument_PDFPrintOperation(void* ptr) {
    NSDocument* nSDocument = (NSDocument*)ptr;
    NSPrintOperation* result_ = [nSDocument PDFPrintOperation];
    return result_;
}

bool C_NSDocument_AllowsDocumentSharing(void* ptr) {
    NSDocument* nSDocument = (NSDocument*)ptr;
    BOOL result_ = [nSDocument allowsDocumentSharing];
    return result_;
}

void* C_NSDocument_ObjectSpecifier(void* ptr) {
    NSDocument* nSDocument = (NSDocument*)ptr;
    NSScriptObjectSpecifier* result_ = [nSDocument objectSpecifier];
    return result_;
}

void* C_NSDocument_LastComponentOfFileName(void* ptr) {
    NSDocument* nSDocument = (NSDocument*)ptr;
    NSString* result_ = [nSDocument lastComponentOfFileName];
    return result_;
}

void C_NSDocument_SetLastComponentOfFileName(void* ptr, void* value) {
    NSDocument* nSDocument = (NSDocument*)ptr;
    [nSDocument setLastComponentOfFileName:(NSString*)value];
}
