#import <Foundation/NSGeometry.h>
#import <Foundation/NSRange.h>
#import <stdbool.h>
#import <stdint.h>
#import <stdlib.h>
#import <utils.h>

void* C_Application_Alloc();

void* C_NSApplication_Init(void* ptr);
void* C_NSApplication_InitWithCoder(void* ptr, void* coder);
void* C_NSApplication_AllocApplication();
void* C_NSApplication_NewApplication();
void* C_NSApplication_Autorelease(void* ptr);
void* C_NSApplication_Retain(void* ptr);
void C_NSApplication_Run(void* ptr);
void C_NSApplication_FinishLaunching(void* ptr);
void C_NSApplication_Stop(void* ptr, void* sender);
void C_NSApplication_SendEvent(void* ptr, void* event);
void C_NSApplication_PostEvent_AtStart(void* ptr, void* event, bool flag);
bool C_NSApplication_SendAction_To_From(void* ptr, void* action, void* target, void* sender);
void* C_NSApplication_TargetForAction(void* ptr, void* action);
void* C_NSApplication_TargetForAction_To_From(void* ptr, void* action, void* target, void* sender);
void C_NSApplication_Terminate(void* ptr, void* sender);
void C_NSApplication_ReplyToApplicationShouldTerminate(void* ptr, bool shouldTerminate);
void C_NSApplication_ActivateIgnoringOtherApps(void* ptr, bool flag);
void C_NSApplication_Deactivate(void* ptr);
void C_NSApplication_DisableRelaunchOnLogin(void* ptr);
void C_NSApplication_EnableRelaunchOnLogin(void* ptr);
void C_NSApplication_RegisterForRemoteNotifications(void* ptr);
void C_NSApplication_UnregisterForRemoteNotifications(void* ptr);
void C_NSApplication_RegisterForRemoteNotificationTypes(void* ptr, unsigned int types);
void C_NSApplication_ToggleTouchBarCustomizationPalette(void* ptr, void* sender);
int C_NSApplication_RequestUserAttention(void* ptr, unsigned int requestType);
void C_NSApplication_CancelUserAttentionRequest(void* ptr, int request);
void C_NSApplication_ReplyToOpenOrPrint(void* ptr, unsigned int reply);
void C_NSApplication_RegisterUserInterfaceItemSearchHandler(void* ptr, void* handler);
void C_NSApplication_UnregisterUserInterfaceItemSearchHandler(void* ptr, void* handler);
void C_NSApplication_ShowHelp(void* ptr, void* sender);
void C_NSApplication_ActivateContextHelpMode(void* ptr, void* sender);
void C_NSApplication_HideOtherApplications(void* ptr, void* sender);
void C_NSApplication_UnhideAllApplications(void* ptr, void* sender);
void C_NSApplication_Application_DetachDrawingThread_ToTarget_WithObject(void* selector, void* target, void* argument);
void C_NSApplication_ReportException(void* ptr, void* exception);
int C_NSApplication_ActivationPolicy(void* ptr);
bool C_NSApplication_SetActivationPolicy(void* ptr, int activationPolicy);
void* C_NSApplication_WindowWithWindowNumber(void* ptr, int windowNum);
void C_NSApplication_MiniaturizeAll(void* ptr, void* sender);
void C_NSApplication_Hide(void* ptr, void* sender);
void C_NSApplication_Unhide(void* ptr, void* sender);
void C_NSApplication_UnhideWithoutActivation(void* ptr);
void C_NSApplication_UpdateWindows(void* ptr);
void C_NSApplication_SetWindowsNeedUpdate(void* ptr, bool needUpdate);
void C_NSApplication_PreventWindowOrdering(void* ptr);
void C_NSApplication_ArrangeInFront(void* ptr, void* sender);
void C_NSApplication_ExtendStateRestoration(void* ptr);
void C_NSApplication_CompleteStateRestoration(void* ptr);
int C_NSApplication_RunModalForWindow(void* ptr, void* window);
void C_NSApplication_StopModal(void* ptr);
void C_NSApplication_StopModalWithCode(void* ptr, int returnCode);
void C_NSApplication_AbortModal(void* ptr);
void* C_NSApplication_BeginModalSessionForWindow(void* ptr, void* window);
int C_NSApplication_RunModalSession(void* ptr, void* session);
void C_NSApplication_OrderFrontColorPanel(void* ptr, void* sender);
void C_NSApplication_OrderFrontStandardAboutPanel(void* ptr, void* sender);
void C_NSApplication_OrderFrontStandardAboutPanelWithOptions(void* ptr, Dictionary optionsDictionary);
void C_NSApplication_OrderFrontCharacterPalette(void* ptr, void* sender);
void C_NSApplication_RunPageLayout(void* ptr, void* sender);
void C_NSApplication_AddWindowsItem_Title_Filename(void* ptr, void* win, void* _string, bool isFilename);
void C_NSApplication_ChangeWindowsItem_Title_Filename(void* ptr, void* win, void* _string, bool isFilename);
void C_NSApplication_RemoveWindowsItem(void* ptr, void* win);
void C_NSApplication_UpdateWindowsItem(void* ptr, void* win);
void C_NSApplication_RegisterServicesMenuSendTypes_ReturnTypes(void* ptr, Array sendTypes, Array returnTypes);
void C_NSApplication_EndModalSession(void* ptr, void* session);
void* C_NSApplication_SharedApplication();
void* C_NSApplication_Delegate(void* ptr);
void C_NSApplication_SetDelegate(void* ptr, void* value);
void* C_NSApplication_CurrentEvent(void* ptr);
bool C_NSApplication_IsRunning(void* ptr);
bool C_NSApplication_IsActive(void* ptr);
unsigned int C_NSApplication_EnabledRemoteNotificationTypes(void* ptr);
bool C_NSApplication_IsRegisteredForRemoteNotifications(void* ptr);
void* C_NSApplication_Appearance(void* ptr);
void C_NSApplication_SetAppearance(void* ptr, void* value);
void* C_NSApplication_EffectiveAppearance(void* ptr);
unsigned int C_NSApplication_CurrentSystemPresentationOptions(void* ptr);
unsigned int C_NSApplication_PresentationOptions(void* ptr);
void C_NSApplication_SetPresentationOptions(void* ptr, unsigned int value);
int C_NSApplication_UserInterfaceLayoutDirection(void* ptr);
void* C_NSApplication_DockTile(void* ptr);
void* C_NSApplication_ApplicationIconImage(void* ptr);
void C_NSApplication_SetApplicationIconImage(void* ptr, void* value);
void* C_NSApplication_HelpMenu(void* ptr);
void C_NSApplication_SetHelpMenu(void* ptr, void* value);
void* C_NSApplication_ServicesProvider(void* ptr);
void C_NSApplication_SetServicesProvider(void* ptr, void* value);
bool C_NSApplication_IsFullKeyboardAccessEnabled(void* ptr);
Array C_NSApplication_OrderedDocuments(void* ptr);
Array C_NSApplication_OrderedWindows(void* ptr);
void* C_NSApplication_KeyWindow(void* ptr);
void* C_NSApplication_MainWindow(void* ptr);
Array C_NSApplication_Windows(void* ptr);
bool C_NSApplication_IsHidden(void* ptr);
unsigned int C_NSApplication_OcclusionState(void* ptr);
void* C_NSApplication_ModalWindow(void* ptr);
void* C_NSApplication_MainMenu(void* ptr);
void C_NSApplication_SetMainMenu(void* ptr, void* value);
bool C_NSApplication_IsAutomaticCustomizeTouchBarMenuItemEnabled(void* ptr);
void C_NSApplication_SetAutomaticCustomizeTouchBarMenuItemEnabled(void* ptr, bool value);
void* C_NSApplication_WindowsMenu(void* ptr);
void C_NSApplication_SetWindowsMenu(void* ptr, void* value);
void* C_NSApplication_ServicesMenu(void* ptr);
void C_NSApplication_SetServicesMenu(void* ptr, void* value);

void* WrapApplicationDelegate(uintptr_t goID);
