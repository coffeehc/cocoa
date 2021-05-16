#import <Appkit/Appkit.h>
#import "application.h"

void* C_Application_Alloc() {
    return [NSApplication alloc];
}

void* C_NSApplication_Init(void* ptr) {
    NSApplication* nSApplication = (NSApplication*)ptr;
    NSApplication* result = [nSApplication init];
    return result;
}

void* C_NSApplication_InitWithCoder(void* ptr, void* coder) {
    NSApplication* nSApplication = (NSApplication*)ptr;
    NSApplication* result = [nSApplication initWithCoder:(NSCoder*)coder];
    return result;
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
    BOOL result = [nSApplication sendAction:(SEL)action to:(id)target from:(id)sender];
    return result;
}

void* C_NSApplication_TargetForAction(void* ptr, void* action) {
    NSApplication* nSApplication = (NSApplication*)ptr;
    id result = [nSApplication targetForAction:(SEL)action];
    return result;
}

void* C_NSApplication_TargetForAction_To_From(void* ptr, void* action, void* target, void* sender) {
    NSApplication* nSApplication = (NSApplication*)ptr;
    id result = [nSApplication targetForAction:(SEL)action to:(id)target from:(id)sender];
    return result;
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
    NSInteger result = [nSApplication requestUserAttention:requestType];
    return result;
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
    NSApplicationActivationPolicy result = [nSApplication activationPolicy];
    return result;
}

bool C_NSApplication_SetActivationPolicy(void* ptr, int activationPolicy) {
    NSApplication* nSApplication = (NSApplication*)ptr;
    BOOL result = [nSApplication setActivationPolicy:activationPolicy];
    return result;
}

void* C_NSApplication_WindowWithWindowNumber(void* ptr, int windowNum) {
    NSApplication* nSApplication = (NSApplication*)ptr;
    NSWindow* result = [nSApplication windowWithWindowNumber:windowNum];
    return result;
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
    NSModalResponse result = [nSApplication runModalForWindow:(NSWindow*)window];
    return result;
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

NSModalSession C_NSApplication_BeginModalSessionForWindow(void* ptr, void* window) {
    NSApplication* nSApplication = (NSApplication*)ptr;
    NSModalSession result = [nSApplication beginModalSessionForWindow:(NSWindow*)window];
    return result;
}

int C_NSApplication_RunModalSession(void* ptr, NSModalSession session) {
    NSApplication* nSApplication = (NSApplication*)ptr;
    NSModalResponse result = [nSApplication runModalSession:session];
    return result;
}

void C_NSApplication_OrderFrontColorPanel(void* ptr, void* sender) {
    NSApplication* nSApplication = (NSApplication*)ptr;
    [nSApplication orderFrontColorPanel:(id)sender];
}

void C_NSApplication_OrderFrontStandardAboutPanel(void* ptr, void* sender) {
    NSApplication* nSApplication = (NSApplication*)ptr;
    [nSApplication orderFrontStandardAboutPanel:(id)sender];
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

void C_NSApplication_EndModalSession(void* ptr, NSModalSession session) {
    NSApplication* nSApplication = (NSApplication*)ptr;
    [nSApplication endModalSession:session];
}

void* C_NSApplication_SharedApplication() {
    NSApplication* result = [NSApplication sharedApplication];
    return result;
}

void* C_NSApplication_Delegate(void* ptr) {
    NSApplication* nSApplication = (NSApplication*)ptr;
    id result = [nSApplication delegate];
    return result;
}

void C_NSApplication_SetDelegate(void* ptr, void* value) {
    NSApplication* nSApplication = (NSApplication*)ptr;
    [nSApplication setDelegate:(id)value];
}

void* C_NSApplication_CurrentEvent(void* ptr) {
    NSApplication* nSApplication = (NSApplication*)ptr;
    NSEvent* result = [nSApplication currentEvent];
    return result;
}

bool C_NSApplication_IsRunning(void* ptr) {
    NSApplication* nSApplication = (NSApplication*)ptr;
    BOOL result = [nSApplication isRunning];
    return result;
}

bool C_NSApplication_IsActive(void* ptr) {
    NSApplication* nSApplication = (NSApplication*)ptr;
    BOOL result = [nSApplication isActive];
    return result;
}

unsigned int C_NSApplication_EnabledRemoteNotificationTypes(void* ptr) {
    NSApplication* nSApplication = (NSApplication*)ptr;
    NSRemoteNotificationType result = [nSApplication enabledRemoteNotificationTypes];
    return result;
}

bool C_NSApplication_IsRegisteredForRemoteNotifications(void* ptr) {
    NSApplication* nSApplication = (NSApplication*)ptr;
    BOOL result = [nSApplication isRegisteredForRemoteNotifications];
    return result;
}

void* C_NSApplication_Appearance(void* ptr) {
    NSApplication* nSApplication = (NSApplication*)ptr;
    NSAppearance* result = [nSApplication appearance];
    return result;
}

void C_NSApplication_SetAppearance(void* ptr, void* value) {
    NSApplication* nSApplication = (NSApplication*)ptr;
    [nSApplication setAppearance:(NSAppearance*)value];
}

void* C_NSApplication_EffectiveAppearance(void* ptr) {
    NSApplication* nSApplication = (NSApplication*)ptr;
    NSAppearance* result = [nSApplication effectiveAppearance];
    return result;
}

unsigned int C_NSApplication_CurrentSystemPresentationOptions(void* ptr) {
    NSApplication* nSApplication = (NSApplication*)ptr;
    NSApplicationPresentationOptions result = [nSApplication currentSystemPresentationOptions];
    return result;
}

unsigned int C_NSApplication_PresentationOptions(void* ptr) {
    NSApplication* nSApplication = (NSApplication*)ptr;
    NSApplicationPresentationOptions result = [nSApplication presentationOptions];
    return result;
}

void C_NSApplication_SetPresentationOptions(void* ptr, unsigned int value) {
    NSApplication* nSApplication = (NSApplication*)ptr;
    [nSApplication setPresentationOptions:value];
}

int C_NSApplication_UserInterfaceLayoutDirection(void* ptr) {
    NSApplication* nSApplication = (NSApplication*)ptr;
    NSUserInterfaceLayoutDirection result = [nSApplication userInterfaceLayoutDirection];
    return result;
}

void* C_NSApplication_DockTile(void* ptr) {
    NSApplication* nSApplication = (NSApplication*)ptr;
    NSDockTile* result = [nSApplication dockTile];
    return result;
}

void* C_NSApplication_ApplicationIconImage(void* ptr) {
    NSApplication* nSApplication = (NSApplication*)ptr;
    NSImage* result = [nSApplication applicationIconImage];
    return result;
}

void C_NSApplication_SetApplicationIconImage(void* ptr, void* value) {
    NSApplication* nSApplication = (NSApplication*)ptr;
    [nSApplication setApplicationIconImage:(NSImage*)value];
}

void* C_NSApplication_HelpMenu(void* ptr) {
    NSApplication* nSApplication = (NSApplication*)ptr;
    NSMenu* result = [nSApplication helpMenu];
    return result;
}

void C_NSApplication_SetHelpMenu(void* ptr, void* value) {
    NSApplication* nSApplication = (NSApplication*)ptr;
    [nSApplication setHelpMenu:(NSMenu*)value];
}

void* C_NSApplication_ServicesProvider(void* ptr) {
    NSApplication* nSApplication = (NSApplication*)ptr;
    id result = [nSApplication servicesProvider];
    return result;
}

void C_NSApplication_SetServicesProvider(void* ptr, void* value) {
    NSApplication* nSApplication = (NSApplication*)ptr;
    [nSApplication setServicesProvider:(id)value];
}

bool C_NSApplication_IsFullKeyboardAccessEnabled(void* ptr) {
    NSApplication* nSApplication = (NSApplication*)ptr;
    BOOL result = [nSApplication isFullKeyboardAccessEnabled];
    return result;
}

void* C_NSApplication_KeyWindow(void* ptr) {
    NSApplication* nSApplication = (NSApplication*)ptr;
    NSWindow* result = [nSApplication keyWindow];
    return result;
}

void* C_NSApplication_MainWindow(void* ptr) {
    NSApplication* nSApplication = (NSApplication*)ptr;
    NSWindow* result = [nSApplication mainWindow];
    return result;
}

bool C_NSApplication_IsHidden(void* ptr) {
    NSApplication* nSApplication = (NSApplication*)ptr;
    BOOL result = [nSApplication isHidden];
    return result;
}

unsigned int C_NSApplication_OcclusionState(void* ptr) {
    NSApplication* nSApplication = (NSApplication*)ptr;
    NSApplicationOcclusionState result = [nSApplication occlusionState];
    return result;
}

void* C_NSApplication_ModalWindow(void* ptr) {
    NSApplication* nSApplication = (NSApplication*)ptr;
    NSWindow* result = [nSApplication modalWindow];
    return result;
}

void* C_NSApplication_MainMenu(void* ptr) {
    NSApplication* nSApplication = (NSApplication*)ptr;
    NSMenu* result = [nSApplication mainMenu];
    return result;
}

void C_NSApplication_SetMainMenu(void* ptr, void* value) {
    NSApplication* nSApplication = (NSApplication*)ptr;
    [nSApplication setMainMenu:(NSMenu*)value];
}

bool C_NSApplication_IsAutomaticCustomizeTouchBarMenuItemEnabled(void* ptr) {
    NSApplication* nSApplication = (NSApplication*)ptr;
    BOOL result = [nSApplication isAutomaticCustomizeTouchBarMenuItemEnabled];
    return result;
}

void C_NSApplication_SetAutomaticCustomizeTouchBarMenuItemEnabled(void* ptr, bool value) {
    NSApplication* nSApplication = (NSApplication*)ptr;
    [nSApplication setAutomaticCustomizeTouchBarMenuItemEnabled:value];
}

void* C_NSApplication_WindowsMenu(void* ptr) {
    NSApplication* nSApplication = (NSApplication*)ptr;
    NSMenu* result = [nSApplication windowsMenu];
    return result;
}

void C_NSApplication_SetWindowsMenu(void* ptr, void* value) {
    NSApplication* nSApplication = (NSApplication*)ptr;
    [nSApplication setWindowsMenu:(NSMenu*)value];
}

void* C_NSApplication_ServicesMenu(void* ptr) {
    NSApplication* nSApplication = (NSApplication*)ptr;
    NSMenu* result = [nSApplication servicesMenu];
    return result;
}

void C_NSApplication_SetServicesMenu(void* ptr, void* value) {
    NSApplication* nSApplication = (NSApplication*)ptr;
    [nSApplication setServicesMenu:(NSMenu*)value];
}
