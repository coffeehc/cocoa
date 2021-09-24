#import "application.h"
#import <AppKit/NSApplication.h>
#import <AppKit/NSApplicationScripting.h>
#import <AppKit/NSColorPanel.h>
#import <AppKit/NSHelpManager.h>
#import <AppKit/NSPageLayout.h>
#import <AppKit/NSTouchBar.h>
#import <AppKit/NSUserInterfaceItemSearching.h>
#import <AppKit/NSWindowRestoration.h>
#import <Foundation/NSArray.h>
#import <Foundation/NSDictionary.h>

void* C_Application_Alloc() {
    return [NSApplication alloc];
}

void* C_NSApplication_Init(void* ptr) {
    NSApplication* nSApplication = (NSApplication*)ptr;
    NSApplication* result_ = [nSApplication init];
    return result_;
}

void* C_NSApplication_InitWithCoder(void* ptr, void* coder) {
    NSApplication* nSApplication = (NSApplication*)ptr;
    NSApplication* result_ = [nSApplication initWithCoder:(NSCoder*)coder];
    return result_;
}

void* C_NSApplication_AllocApplication() {
    NSApplication* result_ = [NSApplication alloc];
    return result_;
}

void* C_NSApplication_NewApplication() {
    NSApplication* result_ = [NSApplication new];
    return result_;
}

void* C_NSApplication_Autorelease(void* ptr) {
    NSApplication* nSApplication = (NSApplication*)ptr;
    NSApplication* result_ = [nSApplication autorelease];
    return result_;
}

void* C_NSApplication_Retain(void* ptr) {
    NSApplication* nSApplication = (NSApplication*)ptr;
    NSApplication* result_ = [nSApplication retain];
    return result_;
}

void C_NSApplication_Run(void* ptr) {
    NSApplication* nSApplication = (NSApplication*)ptr;
    [nSApplication run];
}

void C_NSApplication_FinishLaunching(void* ptr) {
    NSApplication* nSApplication = (NSApplication*)ptr;
    [nSApplication finishLaunching];
}

void C_NSApplication_Stop(void* ptr, void* sender) {
    NSApplication* nSApplication = (NSApplication*)ptr;
    [nSApplication stop:(id)sender];
}

void C_NSApplication_SendEvent(void* ptr, void* event) {
    NSApplication* nSApplication = (NSApplication*)ptr;
    [nSApplication sendEvent:(NSEvent*)event];
}

void C_NSApplication_PostEvent_AtStart(void* ptr, void* event, bool flag) {
    NSApplication* nSApplication = (NSApplication*)ptr;
    [nSApplication postEvent:(NSEvent*)event atStart:flag];
}

bool C_NSApplication_SendAction_To_From(void* ptr, void* action, void* target, void* sender) {
    NSApplication* nSApplication = (NSApplication*)ptr;
    BOOL result_ = [nSApplication sendAction:(SEL)action to:(id)target from:(id)sender];
    return result_;
}

void* C_NSApplication_TargetForAction(void* ptr, void* action) {
    NSApplication* nSApplication = (NSApplication*)ptr;
    id result_ = [nSApplication targetForAction:(SEL)action];
    return result_;
}

void* C_NSApplication_TargetForAction_To_From(void* ptr, void* action, void* target, void* sender) {
    NSApplication* nSApplication = (NSApplication*)ptr;
    id result_ = [nSApplication targetForAction:(SEL)action to:(id)target from:(id)sender];
    return result_;
}

void C_NSApplication_Terminate(void* ptr, void* sender) {
    NSApplication* nSApplication = (NSApplication*)ptr;
    [nSApplication terminate:(id)sender];
}

void C_NSApplication_ReplyToApplicationShouldTerminate(void* ptr, bool shouldTerminate) {
    NSApplication* nSApplication = (NSApplication*)ptr;
    [nSApplication replyToApplicationShouldTerminate:shouldTerminate];
}

void C_NSApplication_ActivateIgnoringOtherApps(void* ptr, bool flag) {
    NSApplication* nSApplication = (NSApplication*)ptr;
    [nSApplication activateIgnoringOtherApps:flag];
}

void C_NSApplication_Deactivate(void* ptr) {
    NSApplication* nSApplication = (NSApplication*)ptr;
    [nSApplication deactivate];
}

void C_NSApplication_DisableRelaunchOnLogin(void* ptr) {
    NSApplication* nSApplication = (NSApplication*)ptr;
    [nSApplication disableRelaunchOnLogin];
}

void C_NSApplication_EnableRelaunchOnLogin(void* ptr) {
    NSApplication* nSApplication = (NSApplication*)ptr;
    [nSApplication enableRelaunchOnLogin];
}

void C_NSApplication_RegisterForRemoteNotifications(void* ptr) {
    NSApplication* nSApplication = (NSApplication*)ptr;
    [nSApplication registerForRemoteNotifications];
}

void C_NSApplication_UnregisterForRemoteNotifications(void* ptr) {
    NSApplication* nSApplication = (NSApplication*)ptr;
    [nSApplication unregisterForRemoteNotifications];
}

void C_NSApplication_RegisterForRemoteNotificationTypes(void* ptr, unsigned int types) {
    NSApplication* nSApplication = (NSApplication*)ptr;
    [nSApplication registerForRemoteNotificationTypes:types];
}

void C_NSApplication_ToggleTouchBarCustomizationPalette(void* ptr, void* sender) {
    NSApplication* nSApplication = (NSApplication*)ptr;
    [nSApplication toggleTouchBarCustomizationPalette:(id)sender];
}

int C_NSApplication_RequestUserAttention(void* ptr, unsigned int requestType) {
    NSApplication* nSApplication = (NSApplication*)ptr;
    NSInteger result_ = [nSApplication requestUserAttention:requestType];
    return result_;
}

void C_NSApplication_CancelUserAttentionRequest(void* ptr, int request) {
    NSApplication* nSApplication = (NSApplication*)ptr;
    [nSApplication cancelUserAttentionRequest:request];
}

void C_NSApplication_ReplyToOpenOrPrint(void* ptr, unsigned int reply) {
    NSApplication* nSApplication = (NSApplication*)ptr;
    [nSApplication replyToOpenOrPrint:reply];
}

void C_NSApplication_RegisterUserInterfaceItemSearchHandler(void* ptr, void* handler) {
    NSApplication* nSApplication = (NSApplication*)ptr;
    [nSApplication registerUserInterfaceItemSearchHandler:(id)handler];
}

void C_NSApplication_UnregisterUserInterfaceItemSearchHandler(void* ptr, void* handler) {
    NSApplication* nSApplication = (NSApplication*)ptr;
    [nSApplication unregisterUserInterfaceItemSearchHandler:(id)handler];
}

void C_NSApplication_ShowHelp(void* ptr, void* sender) {
    NSApplication* nSApplication = (NSApplication*)ptr;
    [nSApplication showHelp:(id)sender];
}

void C_NSApplication_ActivateContextHelpMode(void* ptr, void* sender) {
    NSApplication* nSApplication = (NSApplication*)ptr;
    [nSApplication activateContextHelpMode:(id)sender];
}

void C_NSApplication_HideOtherApplications(void* ptr, void* sender) {
    NSApplication* nSApplication = (NSApplication*)ptr;
    [nSApplication hideOtherApplications:(id)sender];
}

void C_NSApplication_UnhideAllApplications(void* ptr, void* sender) {
    NSApplication* nSApplication = (NSApplication*)ptr;
    [nSApplication unhideAllApplications:(id)sender];
}

void C_NSApplication_Application_DetachDrawingThread_ToTarget_WithObject(void* selector, void* target, void* argument) {
    [NSApplication detachDrawingThread:(SEL)selector toTarget:(id)target withObject:(id)argument];
}

void C_NSApplication_ReportException(void* ptr, void* exception) {
    NSApplication* nSApplication = (NSApplication*)ptr;
    [nSApplication reportException:(NSException*)exception];
}

int C_NSApplication_ActivationPolicy(void* ptr) {
    NSApplication* nSApplication = (NSApplication*)ptr;
    NSApplicationActivationPolicy result_ = [nSApplication activationPolicy];
    return result_;
}

bool C_NSApplication_SetActivationPolicy(void* ptr, int activationPolicy) {
    NSApplication* nSApplication = (NSApplication*)ptr;
    BOOL result_ = [nSApplication setActivationPolicy:activationPolicy];
    return result_;
}

void* C_NSApplication_WindowWithWindowNumber(void* ptr, int windowNum) {
    NSApplication* nSApplication = (NSApplication*)ptr;
    NSWindow* result_ = [nSApplication windowWithWindowNumber:windowNum];
    return result_;
}

void C_NSApplication_MiniaturizeAll(void* ptr, void* sender) {
    NSApplication* nSApplication = (NSApplication*)ptr;
    [nSApplication miniaturizeAll:(id)sender];
}

void C_NSApplication_Hide(void* ptr, void* sender) {
    NSApplication* nSApplication = (NSApplication*)ptr;
    [nSApplication hide:(id)sender];
}

void C_NSApplication_Unhide(void* ptr, void* sender) {
    NSApplication* nSApplication = (NSApplication*)ptr;
    [nSApplication unhide:(id)sender];
}

void C_NSApplication_UnhideWithoutActivation(void* ptr) {
    NSApplication* nSApplication = (NSApplication*)ptr;
    [nSApplication unhideWithoutActivation];
}

void C_NSApplication_UpdateWindows(void* ptr) {
    NSApplication* nSApplication = (NSApplication*)ptr;
    [nSApplication updateWindows];
}

void C_NSApplication_SetWindowsNeedUpdate(void* ptr, bool needUpdate) {
    NSApplication* nSApplication = (NSApplication*)ptr;
    [nSApplication setWindowsNeedUpdate:needUpdate];
}

void C_NSApplication_PreventWindowOrdering(void* ptr) {
    NSApplication* nSApplication = (NSApplication*)ptr;
    [nSApplication preventWindowOrdering];
}

void C_NSApplication_ArrangeInFront(void* ptr, void* sender) {
    NSApplication* nSApplication = (NSApplication*)ptr;
    [nSApplication arrangeInFront:(id)sender];
}

void C_NSApplication_ExtendStateRestoration(void* ptr) {
    NSApplication* nSApplication = (NSApplication*)ptr;
    [nSApplication extendStateRestoration];
}

void C_NSApplication_CompleteStateRestoration(void* ptr) {
    NSApplication* nSApplication = (NSApplication*)ptr;
    [nSApplication completeStateRestoration];
}

int C_NSApplication_RunModalForWindow(void* ptr, void* window) {
    NSApplication* nSApplication = (NSApplication*)ptr;
    NSModalResponse result_ = [nSApplication runModalForWindow:(NSWindow*)window];
    return result_;
}

void C_NSApplication_StopModal(void* ptr) {
    NSApplication* nSApplication = (NSApplication*)ptr;
    [nSApplication stopModal];
}

void C_NSApplication_StopModalWithCode(void* ptr, int returnCode) {
    NSApplication* nSApplication = (NSApplication*)ptr;
    [nSApplication stopModalWithCode:returnCode];
}

void C_NSApplication_AbortModal(void* ptr) {
    NSApplication* nSApplication = (NSApplication*)ptr;
    [nSApplication abortModal];
}

void* C_NSApplication_BeginModalSessionForWindow(void* ptr, void* window) {
    NSApplication* nSApplication = (NSApplication*)ptr;
    NSModalSession result_ = [nSApplication beginModalSessionForWindow:(NSWindow*)window];
    return result_;
}

int C_NSApplication_RunModalSession(void* ptr, void* session) {
    NSApplication* nSApplication = (NSApplication*)ptr;
    NSModalResponse result_ = [nSApplication runModalSession:(NSModalSession)session];
    return result_;
}

void C_NSApplication_OrderFrontColorPanel(void* ptr, void* sender) {
    NSApplication* nSApplication = (NSApplication*)ptr;
    [nSApplication orderFrontColorPanel:(id)sender];
}

void C_NSApplication_OrderFrontStandardAboutPanel(void* ptr, void* sender) {
    NSApplication* nSApplication = (NSApplication*)ptr;
    [nSApplication orderFrontStandardAboutPanel:(id)sender];
}

void C_NSApplication_OrderFrontStandardAboutPanelWithOptions(void* ptr, Dictionary optionsDictionary) {
    NSApplication* nSApplication = (NSApplication*)ptr;
    NSMutableDictionary* objcOptionsDictionary = [NSMutableDictionary dictionaryWithCapacity:optionsDictionary.len];
    if (optionsDictionary.len > 0) {
    	void** optionsDictionaryKeyData = (void**)optionsDictionary.key_data;
    	void** optionsDictionaryValueData = (void**)optionsDictionary.value_data;
    	for (int i = 0; i < optionsDictionary.len; i++) {
    		void* kp = optionsDictionaryKeyData[i];
    		void* vp = optionsDictionaryValueData[i];
    		[objcOptionsDictionary setObject:(NSString*)kp forKey:(id)vp];
    	}
    }
    [nSApplication orderFrontStandardAboutPanelWithOptions:objcOptionsDictionary];
}

void C_NSApplication_OrderFrontCharacterPalette(void* ptr, void* sender) {
    NSApplication* nSApplication = (NSApplication*)ptr;
    [nSApplication orderFrontCharacterPalette:(id)sender];
}

void C_NSApplication_RunPageLayout(void* ptr, void* sender) {
    NSApplication* nSApplication = (NSApplication*)ptr;
    [nSApplication runPageLayout:(id)sender];
}

void C_NSApplication_AddWindowsItem_Title_Filename(void* ptr, void* win, void* _string, bool isFilename) {
    NSApplication* nSApplication = (NSApplication*)ptr;
    [nSApplication addWindowsItem:(NSWindow*)win title:(NSString*)_string filename:isFilename];
}

void C_NSApplication_ChangeWindowsItem_Title_Filename(void* ptr, void* win, void* _string, bool isFilename) {
    NSApplication* nSApplication = (NSApplication*)ptr;
    [nSApplication changeWindowsItem:(NSWindow*)win title:(NSString*)_string filename:isFilename];
}

void C_NSApplication_RemoveWindowsItem(void* ptr, void* win) {
    NSApplication* nSApplication = (NSApplication*)ptr;
    [nSApplication removeWindowsItem:(NSWindow*)win];
}

void C_NSApplication_UpdateWindowsItem(void* ptr, void* win) {
    NSApplication* nSApplication = (NSApplication*)ptr;
    [nSApplication updateWindowsItem:(NSWindow*)win];
}

void C_NSApplication_RegisterServicesMenuSendTypes_ReturnTypes(void* ptr, Array sendTypes, Array returnTypes) {
    NSApplication* nSApplication = (NSApplication*)ptr;
    NSMutableArray* objcSendTypes = [NSMutableArray arrayWithCapacity:sendTypes.len];
    if (sendTypes.len > 0) {
    	void** sendTypesData = (void**)sendTypes.data;
    	for (int i = 0; i < sendTypes.len; i++) {
    		void* p = sendTypesData[i];
    		[objcSendTypes addObject:(NSString*)p];
    	}
    }
    NSMutableArray* objcReturnTypes = [NSMutableArray arrayWithCapacity:returnTypes.len];
    if (returnTypes.len > 0) {
    	void** returnTypesData = (void**)returnTypes.data;
    	for (int i = 0; i < returnTypes.len; i++) {
    		void* p = returnTypesData[i];
    		[objcReturnTypes addObject:(NSString*)p];
    	}
    }
    [nSApplication registerServicesMenuSendTypes:objcSendTypes returnTypes:objcReturnTypes];
}

void C_NSApplication_EndModalSession(void* ptr, void* session) {
    NSApplication* nSApplication = (NSApplication*)ptr;
    [nSApplication endModalSession:(NSModalSession)session];
}

void* C_NSApplication_SharedApplication() {
    NSApplication* result_ = [NSApplication sharedApplication];
    return result_;
}

void* C_NSApplication_Delegate(void* ptr) {
    NSApplication* nSApplication = (NSApplication*)ptr;
    id result_ = [nSApplication delegate];
    return result_;
}

void C_NSApplication_SetDelegate(void* ptr, void* value) {
    NSApplication* nSApplication = (NSApplication*)ptr;
    [nSApplication setDelegate:(id)value];
}

void* C_NSApplication_CurrentEvent(void* ptr) {
    NSApplication* nSApplication = (NSApplication*)ptr;
    NSEvent* result_ = [nSApplication currentEvent];
    return result_;
}

bool C_NSApplication_IsRunning(void* ptr) {
    NSApplication* nSApplication = (NSApplication*)ptr;
    BOOL result_ = [nSApplication isRunning];
    return result_;
}

bool C_NSApplication_IsActive(void* ptr) {
    NSApplication* nSApplication = (NSApplication*)ptr;
    BOOL result_ = [nSApplication isActive];
    return result_;
}

unsigned int C_NSApplication_EnabledRemoteNotificationTypes(void* ptr) {
    NSApplication* nSApplication = (NSApplication*)ptr;
    NSRemoteNotificationType result_ = [nSApplication enabledRemoteNotificationTypes];
    return result_;
}

bool C_NSApplication_IsRegisteredForRemoteNotifications(void* ptr) {
    NSApplication* nSApplication = (NSApplication*)ptr;
    BOOL result_ = [nSApplication isRegisteredForRemoteNotifications];
    return result_;
}

void* C_NSApplication_Appearance(void* ptr) {
    NSApplication* nSApplication = (NSApplication*)ptr;
    NSAppearance* result_ = [nSApplication appearance];
    return result_;
}

void C_NSApplication_SetAppearance(void* ptr, void* value) {
    NSApplication* nSApplication = (NSApplication*)ptr;
    [nSApplication setAppearance:(NSAppearance*)value];
}

void* C_NSApplication_EffectiveAppearance(void* ptr) {
    NSApplication* nSApplication = (NSApplication*)ptr;
    NSAppearance* result_ = [nSApplication effectiveAppearance];
    return result_;
}

unsigned int C_NSApplication_CurrentSystemPresentationOptions(void* ptr) {
    NSApplication* nSApplication = (NSApplication*)ptr;
    NSApplicationPresentationOptions result_ = [nSApplication currentSystemPresentationOptions];
    return result_;
}

unsigned int C_NSApplication_PresentationOptions(void* ptr) {
    NSApplication* nSApplication = (NSApplication*)ptr;
    NSApplicationPresentationOptions result_ = [nSApplication presentationOptions];
    return result_;
}

void C_NSApplication_SetPresentationOptions(void* ptr, unsigned int value) {
    NSApplication* nSApplication = (NSApplication*)ptr;
    [nSApplication setPresentationOptions:value];
}

int C_NSApplication_UserInterfaceLayoutDirection(void* ptr) {
    NSApplication* nSApplication = (NSApplication*)ptr;
    NSUserInterfaceLayoutDirection result_ = [nSApplication userInterfaceLayoutDirection];
    return result_;
}

void* C_NSApplication_DockTile(void* ptr) {
    NSApplication* nSApplication = (NSApplication*)ptr;
    NSDockTile* result_ = [nSApplication dockTile];
    return result_;
}

void* C_NSApplication_ApplicationIconImage(void* ptr) {
    NSApplication* nSApplication = (NSApplication*)ptr;
    NSImage* result_ = [nSApplication applicationIconImage];
    return result_;
}

void C_NSApplication_SetApplicationIconImage(void* ptr, void* value) {
    NSApplication* nSApplication = (NSApplication*)ptr;
    [nSApplication setApplicationIconImage:(NSImage*)value];
}

void* C_NSApplication_HelpMenu(void* ptr) {
    NSApplication* nSApplication = (NSApplication*)ptr;
    NSMenu* result_ = [nSApplication helpMenu];
    return result_;
}

void C_NSApplication_SetHelpMenu(void* ptr, void* value) {
    NSApplication* nSApplication = (NSApplication*)ptr;
    [nSApplication setHelpMenu:(NSMenu*)value];
}

void* C_NSApplication_ServicesProvider(void* ptr) {
    NSApplication* nSApplication = (NSApplication*)ptr;
    id result_ = [nSApplication servicesProvider];
    return result_;
}

void C_NSApplication_SetServicesProvider(void* ptr, void* value) {
    NSApplication* nSApplication = (NSApplication*)ptr;
    [nSApplication setServicesProvider:(id)value];
}

bool C_NSApplication_IsFullKeyboardAccessEnabled(void* ptr) {
    NSApplication* nSApplication = (NSApplication*)ptr;
    BOOL result_ = [nSApplication isFullKeyboardAccessEnabled];
    return result_;
}

Array C_NSApplication_OrderedDocuments(void* ptr) {
    NSApplication* nSApplication = (NSApplication*)ptr;
    NSArray* result_ = [nSApplication orderedDocuments];
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

Array C_NSApplication_OrderedWindows(void* ptr) {
    NSApplication* nSApplication = (NSApplication*)ptr;
    NSArray* result_ = [nSApplication orderedWindows];
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

void* C_NSApplication_KeyWindow(void* ptr) {
    NSApplication* nSApplication = (NSApplication*)ptr;
    NSWindow* result_ = [nSApplication keyWindow];
    return result_;
}

void* C_NSApplication_MainWindow(void* ptr) {
    NSApplication* nSApplication = (NSApplication*)ptr;
    NSWindow* result_ = [nSApplication mainWindow];
    return result_;
}

Array C_NSApplication_Windows(void* ptr) {
    NSApplication* nSApplication = (NSApplication*)ptr;
    NSArray* result_ = [nSApplication windows];
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

bool C_NSApplication_IsHidden(void* ptr) {
    NSApplication* nSApplication = (NSApplication*)ptr;
    BOOL result_ = [nSApplication isHidden];
    return result_;
}

unsigned int C_NSApplication_OcclusionState(void* ptr) {
    NSApplication* nSApplication = (NSApplication*)ptr;
    NSApplicationOcclusionState result_ = [nSApplication occlusionState];
    return result_;
}

void* C_NSApplication_ModalWindow(void* ptr) {
    NSApplication* nSApplication = (NSApplication*)ptr;
    NSWindow* result_ = [nSApplication modalWindow];
    return result_;
}

void* C_NSApplication_MainMenu(void* ptr) {
    NSApplication* nSApplication = (NSApplication*)ptr;
    NSMenu* result_ = [nSApplication mainMenu];
    return result_;
}

void C_NSApplication_SetMainMenu(void* ptr, void* value) {
    NSApplication* nSApplication = (NSApplication*)ptr;
    [nSApplication setMainMenu:(NSMenu*)value];
}

bool C_NSApplication_IsAutomaticCustomizeTouchBarMenuItemEnabled(void* ptr) {
    NSApplication* nSApplication = (NSApplication*)ptr;
    BOOL result_ = [nSApplication isAutomaticCustomizeTouchBarMenuItemEnabled];
    return result_;
}

void C_NSApplication_SetAutomaticCustomizeTouchBarMenuItemEnabled(void* ptr, bool value) {
    NSApplication* nSApplication = (NSApplication*)ptr;
    [nSApplication setAutomaticCustomizeTouchBarMenuItemEnabled:value];
}

void* C_NSApplication_WindowsMenu(void* ptr) {
    NSApplication* nSApplication = (NSApplication*)ptr;
    NSMenu* result_ = [nSApplication windowsMenu];
    return result_;
}

void C_NSApplication_SetWindowsMenu(void* ptr, void* value) {
    NSApplication* nSApplication = (NSApplication*)ptr;
    [nSApplication setWindowsMenu:(NSMenu*)value];
}

void* C_NSApplication_ServicesMenu(void* ptr) {
    NSApplication* nSApplication = (NSApplication*)ptr;
    NSMenu* result_ = [nSApplication servicesMenu];
    return result_;
}

void C_NSApplication_SetServicesMenu(void* ptr, void* value) {
    NSApplication* nSApplication = (NSApplication*)ptr;
    [nSApplication setServicesMenu:(NSMenu*)value];
}
