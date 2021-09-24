package appkit

// #include "application.h"
import "C"
import (
	"github.com/hsiafan/cocoa/foundation"
	"github.com/hsiafan/cocoa/objc"
	"runtime/cgo"
	"unsafe"
)

type Application interface {
	Responder
	Run()
	FinishLaunching()
	Stop(sender objc.Object)
	SendEvent(event Event)
	PostEvent_AtStart(event Event, flag bool)
	SendAction_To_From(action objc.Selector, target objc.Object, sender objc.Object) bool
	TargetForAction(action objc.Selector) objc.Object
	TargetForAction_To_From(action objc.Selector, target objc.Object, sender objc.Object) objc.Object
	Terminate(sender objc.Object)
	ReplyToApplicationShouldTerminate(shouldTerminate bool)
	ActivateIgnoringOtherApps(flag bool)
	Deactivate()
	DisableRelaunchOnLogin()
	EnableRelaunchOnLogin()
	RegisterForRemoteNotifications()
	UnregisterForRemoteNotifications()
	RegisterForRemoteNotificationTypes(types RemoteNotificationType)
	ToggleTouchBarCustomizationPalette(sender objc.Object)
	RequestUserAttention(requestType RequestUserAttentionType) int
	CancelUserAttentionRequest(request int)
	ReplyToOpenOrPrint(reply ApplicationDelegateReply)
	RegisterUserInterfaceItemSearchHandler(handler objc.Object)
	UnregisterUserInterfaceItemSearchHandler(handler objc.Object)
	ShowHelp(sender objc.Object)
	ActivateContextHelpMode(sender objc.Object)
	HideOtherApplications(sender objc.Object)
	UnhideAllApplications(sender objc.Object)
	ReportException(exception foundation.Exception)
	ActivationPolicy() ApplicationActivationPolicy
	SetActivationPolicy(activationPolicy ApplicationActivationPolicy) bool
	WindowWithWindowNumber(windowNum int) Window
	MiniaturizeAll(sender objc.Object)
	Hide(sender objc.Object)
	Unhide(sender objc.Object)
	UnhideWithoutActivation()
	UpdateWindows()
	SetWindowsNeedUpdate(needUpdate bool)
	PreventWindowOrdering()
	ArrangeInFront(sender objc.Object)
	ExtendStateRestoration()
	CompleteStateRestoration()
	RunModalForWindow(window Window) ModalResponse
	StopModal()
	StopModalWithCode(returnCode ModalResponse)
	AbortModal()
	BeginModalSessionForWindow(window Window) ModalSession
	RunModalSession(session ModalSession) ModalResponse
	OrderFrontColorPanel(sender objc.Object)
	OrderFrontStandardAboutPanel(sender objc.Object)
	OrderFrontStandardAboutPanelWithOptions(optionsDictionary map[AboutPanelOptionKey]objc.Object)
	OrderFrontCharacterPalette(sender objc.Object)
	RunPageLayout(sender objc.Object)
	AddWindowsItem_Title_Filename(win Window, _string string, isFilename bool)
	ChangeWindowsItem_Title_Filename(win Window, _string string, isFilename bool)
	RemoveWindowsItem(win Window)
	UpdateWindowsItem(win Window)
	RegisterServicesMenuSendTypes_ReturnTypes(sendTypes []PasteboardType, returnTypes []PasteboardType)
	EndModalSession(session ModalSession)
	Delegate() objc.Object
	SetDelegate(value objc.Object)
	CurrentEvent() Event
	IsRunning() bool
	IsActive() bool
	EnabledRemoteNotificationTypes() RemoteNotificationType
	IsRegisteredForRemoteNotifications() bool
	Appearance() Appearance
	SetAppearance(value Appearance)
	EffectiveAppearance() Appearance
	CurrentSystemPresentationOptions() ApplicationPresentationOptions
	PresentationOptions() ApplicationPresentationOptions
	SetPresentationOptions(value ApplicationPresentationOptions)
	UserInterfaceLayoutDirection() UserInterfaceLayoutDirection
	DockTile() DockTile
	ApplicationIconImage() Image
	SetApplicationIconImage(value Image)
	HelpMenu() Menu
	SetHelpMenu(value Menu)
	ServicesProvider() objc.Object
	SetServicesProvider(value objc.Object)
	IsFullKeyboardAccessEnabled() bool
	OrderedDocuments() []Document
	OrderedWindows() []Window
	KeyWindow() Window
	MainWindow() Window
	Windows() []Window
	IsHidden() bool
	OcclusionState() ApplicationOcclusionState
	ModalWindow() Window
	MainMenu() Menu
	SetMainMenu(value Menu)
	IsAutomaticCustomizeTouchBarMenuItemEnabled() bool
	SetAutomaticCustomizeTouchBarMenuItemEnabled(value bool)
	WindowsMenu() Menu
	SetWindowsMenu(value Menu)
	ServicesMenu() Menu
	SetServicesMenu(value Menu)
}

type NSApplication struct {
	NSResponder
}

func MakeApplication(ptr unsafe.Pointer) NSApplication {
	return NSApplication{
		NSResponder: MakeResponder(ptr),
	}
}

func (n NSApplication) Init() NSApplication {
	result_ := C.C_NSApplication_Init(n.Ptr())
	return MakeApplication(result_)
}

func (n NSApplication) InitWithCoder(coder foundation.Coder) NSApplication {
	result_ := C.C_NSApplication_InitWithCoder(n.Ptr(), objc.ExtractPtr(coder))
	return MakeApplication(result_)
}

func AllocApplication() NSApplication {
	result_ := C.C_NSApplication_AllocApplication()
	return MakeApplication(result_)
}

func NewApplication() NSApplication {
	result_ := C.C_NSApplication_NewApplication()
	return MakeApplication(result_)
}

func (n NSApplication) Autorelease() NSApplication {
	result_ := C.C_NSApplication_Autorelease(n.Ptr())
	return MakeApplication(result_)
}

func (n NSApplication) Retain() NSApplication {
	result_ := C.C_NSApplication_Retain(n.Ptr())
	return MakeApplication(result_)
}

func (n NSApplication) Run() {
	C.C_NSApplication_Run(n.Ptr())
}

func (n NSApplication) FinishLaunching() {
	C.C_NSApplication_FinishLaunching(n.Ptr())
}

func (n NSApplication) Stop(sender objc.Object) {
	C.C_NSApplication_Stop(n.Ptr(), objc.ExtractPtr(sender))
}

func (n NSApplication) SendEvent(event Event) {
	C.C_NSApplication_SendEvent(n.Ptr(), objc.ExtractPtr(event))
}

func (n NSApplication) PostEvent_AtStart(event Event, flag bool) {
	C.C_NSApplication_PostEvent_AtStart(n.Ptr(), objc.ExtractPtr(event), C.bool(flag))
}

func (n NSApplication) SendAction_To_From(action objc.Selector, target objc.Object, sender objc.Object) bool {
	result_ := C.C_NSApplication_SendAction_To_From(n.Ptr(), unsafe.Pointer(action), objc.ExtractPtr(target), objc.ExtractPtr(sender))
	return bool(result_)
}

func (n NSApplication) TargetForAction(action objc.Selector) objc.Object {
	result_ := C.C_NSApplication_TargetForAction(n.Ptr(), unsafe.Pointer(action))
	return objc.MakeObject(result_)
}

func (n NSApplication) TargetForAction_To_From(action objc.Selector, target objc.Object, sender objc.Object) objc.Object {
	result_ := C.C_NSApplication_TargetForAction_To_From(n.Ptr(), unsafe.Pointer(action), objc.ExtractPtr(target), objc.ExtractPtr(sender))
	return objc.MakeObject(result_)
}

func (n NSApplication) Terminate(sender objc.Object) {
	C.C_NSApplication_Terminate(n.Ptr(), objc.ExtractPtr(sender))
}

func (n NSApplication) ReplyToApplicationShouldTerminate(shouldTerminate bool) {
	C.C_NSApplication_ReplyToApplicationShouldTerminate(n.Ptr(), C.bool(shouldTerminate))
}

func (n NSApplication) ActivateIgnoringOtherApps(flag bool) {
	C.C_NSApplication_ActivateIgnoringOtherApps(n.Ptr(), C.bool(flag))
}

func (n NSApplication) Deactivate() {
	C.C_NSApplication_Deactivate(n.Ptr())
}

func (n NSApplication) DisableRelaunchOnLogin() {
	C.C_NSApplication_DisableRelaunchOnLogin(n.Ptr())
}

func (n NSApplication) EnableRelaunchOnLogin() {
	C.C_NSApplication_EnableRelaunchOnLogin(n.Ptr())
}

func (n NSApplication) RegisterForRemoteNotifications() {
	C.C_NSApplication_RegisterForRemoteNotifications(n.Ptr())
}

func (n NSApplication) UnregisterForRemoteNotifications() {
	C.C_NSApplication_UnregisterForRemoteNotifications(n.Ptr())
}

func (n NSApplication) RegisterForRemoteNotificationTypes(types RemoteNotificationType) {
	C.C_NSApplication_RegisterForRemoteNotificationTypes(n.Ptr(), C.uint(uint(types)))
}

func (n NSApplication) ToggleTouchBarCustomizationPalette(sender objc.Object) {
	C.C_NSApplication_ToggleTouchBarCustomizationPalette(n.Ptr(), objc.ExtractPtr(sender))
}

func (n NSApplication) RequestUserAttention(requestType RequestUserAttentionType) int {
	result_ := C.C_NSApplication_RequestUserAttention(n.Ptr(), C.uint(uint(requestType)))
	return int(result_)
}

func (n NSApplication) CancelUserAttentionRequest(request int) {
	C.C_NSApplication_CancelUserAttentionRequest(n.Ptr(), C.int(request))
}

func (n NSApplication) ReplyToOpenOrPrint(reply ApplicationDelegateReply) {
	C.C_NSApplication_ReplyToOpenOrPrint(n.Ptr(), C.uint(uint(reply)))
}

func (n NSApplication) RegisterUserInterfaceItemSearchHandler(handler objc.Object) {
	C.C_NSApplication_RegisterUserInterfaceItemSearchHandler(n.Ptr(), objc.ExtractPtr(handler))
}

func (n NSApplication) UnregisterUserInterfaceItemSearchHandler(handler objc.Object) {
	C.C_NSApplication_UnregisterUserInterfaceItemSearchHandler(n.Ptr(), objc.ExtractPtr(handler))
}

func (n NSApplication) ShowHelp(sender objc.Object) {
	C.C_NSApplication_ShowHelp(n.Ptr(), objc.ExtractPtr(sender))
}

func (n NSApplication) ActivateContextHelpMode(sender objc.Object) {
	C.C_NSApplication_ActivateContextHelpMode(n.Ptr(), objc.ExtractPtr(sender))
}

func (n NSApplication) HideOtherApplications(sender objc.Object) {
	C.C_NSApplication_HideOtherApplications(n.Ptr(), objc.ExtractPtr(sender))
}

func (n NSApplication) UnhideAllApplications(sender objc.Object) {
	C.C_NSApplication_UnhideAllApplications(n.Ptr(), objc.ExtractPtr(sender))
}

func Application_DetachDrawingThread_ToTarget_WithObject(selector objc.Selector, target objc.Object, argument objc.Object) {
	C.C_NSApplication_Application_DetachDrawingThread_ToTarget_WithObject(unsafe.Pointer(selector), objc.ExtractPtr(target), objc.ExtractPtr(argument))
}

func (n NSApplication) ReportException(exception foundation.Exception) {
	C.C_NSApplication_ReportException(n.Ptr(), objc.ExtractPtr(exception))
}

func (n NSApplication) ActivationPolicy() ApplicationActivationPolicy {
	result_ := C.C_NSApplication_ActivationPolicy(n.Ptr())
	return ApplicationActivationPolicy(int(result_))
}

func (n NSApplication) SetActivationPolicy(activationPolicy ApplicationActivationPolicy) bool {
	result_ := C.C_NSApplication_SetActivationPolicy(n.Ptr(), C.int(int(activationPolicy)))
	return bool(result_)
}

func (n NSApplication) WindowWithWindowNumber(windowNum int) Window {
	result_ := C.C_NSApplication_WindowWithWindowNumber(n.Ptr(), C.int(windowNum))
	return MakeWindow(result_)
}

func (n NSApplication) MiniaturizeAll(sender objc.Object) {
	C.C_NSApplication_MiniaturizeAll(n.Ptr(), objc.ExtractPtr(sender))
}

func (n NSApplication) Hide(sender objc.Object) {
	C.C_NSApplication_Hide(n.Ptr(), objc.ExtractPtr(sender))
}

func (n NSApplication) Unhide(sender objc.Object) {
	C.C_NSApplication_Unhide(n.Ptr(), objc.ExtractPtr(sender))
}

func (n NSApplication) UnhideWithoutActivation() {
	C.C_NSApplication_UnhideWithoutActivation(n.Ptr())
}

func (n NSApplication) UpdateWindows() {
	C.C_NSApplication_UpdateWindows(n.Ptr())
}

func (n NSApplication) SetWindowsNeedUpdate(needUpdate bool) {
	C.C_NSApplication_SetWindowsNeedUpdate(n.Ptr(), C.bool(needUpdate))
}

func (n NSApplication) PreventWindowOrdering() {
	C.C_NSApplication_PreventWindowOrdering(n.Ptr())
}

func (n NSApplication) ArrangeInFront(sender objc.Object) {
	C.C_NSApplication_ArrangeInFront(n.Ptr(), objc.ExtractPtr(sender))
}

func (n NSApplication) ExtendStateRestoration() {
	C.C_NSApplication_ExtendStateRestoration(n.Ptr())
}

func (n NSApplication) CompleteStateRestoration() {
	C.C_NSApplication_CompleteStateRestoration(n.Ptr())
}

func (n NSApplication) RunModalForWindow(window Window) ModalResponse {
	result_ := C.C_NSApplication_RunModalForWindow(n.Ptr(), objc.ExtractPtr(window))
	return ModalResponse(int(result_))
}

func (n NSApplication) StopModal() {
	C.C_NSApplication_StopModal(n.Ptr())
}

func (n NSApplication) StopModalWithCode(returnCode ModalResponse) {
	C.C_NSApplication_StopModalWithCode(n.Ptr(), C.int(int(returnCode)))
}

func (n NSApplication) AbortModal() {
	C.C_NSApplication_AbortModal(n.Ptr())
}

func (n NSApplication) BeginModalSessionForWindow(window Window) ModalSession {
	result_ := C.C_NSApplication_BeginModalSessionForWindow(n.Ptr(), objc.ExtractPtr(window))
	return ModalSession(result_)
}

func (n NSApplication) RunModalSession(session ModalSession) ModalResponse {
	result_ := C.C_NSApplication_RunModalSession(n.Ptr(), unsafe.Pointer(session))
	return ModalResponse(int(result_))
}

func (n NSApplication) OrderFrontColorPanel(sender objc.Object) {
	C.C_NSApplication_OrderFrontColorPanel(n.Ptr(), objc.ExtractPtr(sender))
}

func (n NSApplication) OrderFrontStandardAboutPanel(sender objc.Object) {
	C.C_NSApplication_OrderFrontStandardAboutPanel(n.Ptr(), objc.ExtractPtr(sender))
}

func (n NSApplication) OrderFrontStandardAboutPanelWithOptions(optionsDictionary map[AboutPanelOptionKey]objc.Object) {
	var cOptionsDictionary C.Dictionary
	if len(optionsDictionary) > 0 {
		cOptionsDictionaryKeyData := make([]unsafe.Pointer, len(optionsDictionary))
		cOptionsDictionaryValueData := make([]unsafe.Pointer, len(optionsDictionary))
		var idx = 0
		for k, v := range optionsDictionary {
			cOptionsDictionaryKeyData[idx] = foundation.NewString(string(k)).Ptr()
			cOptionsDictionaryValueData[idx] = objc.ExtractPtr(v)
			idx++
		}
		cOptionsDictionary.key_data = unsafe.Pointer(&cOptionsDictionaryKeyData[0])
		cOptionsDictionary.value_data = unsafe.Pointer(&cOptionsDictionaryValueData[0])
		cOptionsDictionary.len = C.int(len(optionsDictionary))
	}
	C.C_NSApplication_OrderFrontStandardAboutPanelWithOptions(n.Ptr(), cOptionsDictionary)
}

func (n NSApplication) OrderFrontCharacterPalette(sender objc.Object) {
	C.C_NSApplication_OrderFrontCharacterPalette(n.Ptr(), objc.ExtractPtr(sender))
}

func (n NSApplication) RunPageLayout(sender objc.Object) {
	C.C_NSApplication_RunPageLayout(n.Ptr(), objc.ExtractPtr(sender))
}

func (n NSApplication) AddWindowsItem_Title_Filename(win Window, _string string, isFilename bool) {
	C.C_NSApplication_AddWindowsItem_Title_Filename(n.Ptr(), objc.ExtractPtr(win), foundation.NewString(_string).Ptr(), C.bool(isFilename))
}

func (n NSApplication) ChangeWindowsItem_Title_Filename(win Window, _string string, isFilename bool) {
	C.C_NSApplication_ChangeWindowsItem_Title_Filename(n.Ptr(), objc.ExtractPtr(win), foundation.NewString(_string).Ptr(), C.bool(isFilename))
}

func (n NSApplication) RemoveWindowsItem(win Window) {
	C.C_NSApplication_RemoveWindowsItem(n.Ptr(), objc.ExtractPtr(win))
}

func (n NSApplication) UpdateWindowsItem(win Window) {
	C.C_NSApplication_UpdateWindowsItem(n.Ptr(), objc.ExtractPtr(win))
}

func (n NSApplication) RegisterServicesMenuSendTypes_ReturnTypes(sendTypes []PasteboardType, returnTypes []PasteboardType) {
	var cSendTypes C.Array
	if len(sendTypes) > 0 {
		cSendTypesData := make([]unsafe.Pointer, len(sendTypes))
		for idx, v := range sendTypes {
			cSendTypesData[idx] = foundation.NewString(string(v)).Ptr()
		}
		cSendTypes.data = unsafe.Pointer(&cSendTypesData[0])
		cSendTypes.len = C.int(len(sendTypes))
	}
	var cReturnTypes C.Array
	if len(returnTypes) > 0 {
		cReturnTypesData := make([]unsafe.Pointer, len(returnTypes))
		for idx, v := range returnTypes {
			cReturnTypesData[idx] = foundation.NewString(string(v)).Ptr()
		}
		cReturnTypes.data = unsafe.Pointer(&cReturnTypesData[0])
		cReturnTypes.len = C.int(len(returnTypes))
	}
	C.C_NSApplication_RegisterServicesMenuSendTypes_ReturnTypes(n.Ptr(), cSendTypes, cReturnTypes)
}

func (n NSApplication) EndModalSession(session ModalSession) {
	C.C_NSApplication_EndModalSession(n.Ptr(), unsafe.Pointer(session))
}

func SharedApplication() Application {
	result_ := C.C_NSApplication_SharedApplication()
	return MakeApplication(result_)
}

func (n NSApplication) Delegate() objc.Object {
	result_ := C.C_NSApplication_Delegate(n.Ptr())
	return objc.MakeObject(result_)
}

func (n NSApplication) SetDelegate(value objc.Object) {
	C.C_NSApplication_SetDelegate(n.Ptr(), objc.ExtractPtr(value))
}

func (n NSApplication) CurrentEvent() Event {
	result_ := C.C_NSApplication_CurrentEvent(n.Ptr())
	return MakeEvent(result_)
}

func (n NSApplication) IsRunning() bool {
	result_ := C.C_NSApplication_IsRunning(n.Ptr())
	return bool(result_)
}

func (n NSApplication) IsActive() bool {
	result_ := C.C_NSApplication_IsActive(n.Ptr())
	return bool(result_)
}

func (n NSApplication) EnabledRemoteNotificationTypes() RemoteNotificationType {
	result_ := C.C_NSApplication_EnabledRemoteNotificationTypes(n.Ptr())
	return RemoteNotificationType(uint(result_))
}

func (n NSApplication) IsRegisteredForRemoteNotifications() bool {
	result_ := C.C_NSApplication_IsRegisteredForRemoteNotifications(n.Ptr())
	return bool(result_)
}

func (n NSApplication) Appearance() Appearance {
	result_ := C.C_NSApplication_Appearance(n.Ptr())
	return MakeAppearance(result_)
}

func (n NSApplication) SetAppearance(value Appearance) {
	C.C_NSApplication_SetAppearance(n.Ptr(), objc.ExtractPtr(value))
}

func (n NSApplication) EffectiveAppearance() Appearance {
	result_ := C.C_NSApplication_EffectiveAppearance(n.Ptr())
	return MakeAppearance(result_)
}

func (n NSApplication) CurrentSystemPresentationOptions() ApplicationPresentationOptions {
	result_ := C.C_NSApplication_CurrentSystemPresentationOptions(n.Ptr())
	return ApplicationPresentationOptions(uint(result_))
}

func (n NSApplication) PresentationOptions() ApplicationPresentationOptions {
	result_ := C.C_NSApplication_PresentationOptions(n.Ptr())
	return ApplicationPresentationOptions(uint(result_))
}

func (n NSApplication) SetPresentationOptions(value ApplicationPresentationOptions) {
	C.C_NSApplication_SetPresentationOptions(n.Ptr(), C.uint(uint(value)))
}

func (n NSApplication) UserInterfaceLayoutDirection() UserInterfaceLayoutDirection {
	result_ := C.C_NSApplication_UserInterfaceLayoutDirection(n.Ptr())
	return UserInterfaceLayoutDirection(int(result_))
}

func (n NSApplication) DockTile() DockTile {
	result_ := C.C_NSApplication_DockTile(n.Ptr())
	return MakeDockTile(result_)
}

func (n NSApplication) ApplicationIconImage() Image {
	result_ := C.C_NSApplication_ApplicationIconImage(n.Ptr())
	return MakeImage(result_)
}

func (n NSApplication) SetApplicationIconImage(value Image) {
	C.C_NSApplication_SetApplicationIconImage(n.Ptr(), objc.ExtractPtr(value))
}

func (n NSApplication) HelpMenu() Menu {
	result_ := C.C_NSApplication_HelpMenu(n.Ptr())
	return MakeMenu(result_)
}

func (n NSApplication) SetHelpMenu(value Menu) {
	C.C_NSApplication_SetHelpMenu(n.Ptr(), objc.ExtractPtr(value))
}

func (n NSApplication) ServicesProvider() objc.Object {
	result_ := C.C_NSApplication_ServicesProvider(n.Ptr())
	return objc.MakeObject(result_)
}

func (n NSApplication) SetServicesProvider(value objc.Object) {
	C.C_NSApplication_SetServicesProvider(n.Ptr(), objc.ExtractPtr(value))
}

func (n NSApplication) IsFullKeyboardAccessEnabled() bool {
	result_ := C.C_NSApplication_IsFullKeyboardAccessEnabled(n.Ptr())
	return bool(result_)
}

func (n NSApplication) OrderedDocuments() []Document {
	result_ := C.C_NSApplication_OrderedDocuments(n.Ptr())
	if result_.len > 0 {
		defer C.free(result_.data)
	}
	result_Slice := unsafe.Slice((*unsafe.Pointer)(result_.data), int(result_.len))
	var goResult_ = make([]Document, len(result_Slice))
	for idx, r := range result_Slice {
		goResult_[idx] = MakeDocument(r)
	}
	return goResult_
}

func (n NSApplication) OrderedWindows() []Window {
	result_ := C.C_NSApplication_OrderedWindows(n.Ptr())
	if result_.len > 0 {
		defer C.free(result_.data)
	}
	result_Slice := unsafe.Slice((*unsafe.Pointer)(result_.data), int(result_.len))
	var goResult_ = make([]Window, len(result_Slice))
	for idx, r := range result_Slice {
		goResult_[idx] = MakeWindow(r)
	}
	return goResult_
}

func (n NSApplication) KeyWindow() Window {
	result_ := C.C_NSApplication_KeyWindow(n.Ptr())
	return MakeWindow(result_)
}

func (n NSApplication) MainWindow() Window {
	result_ := C.C_NSApplication_MainWindow(n.Ptr())
	return MakeWindow(result_)
}

func (n NSApplication) Windows() []Window {
	result_ := C.C_NSApplication_Windows(n.Ptr())
	if result_.len > 0 {
		defer C.free(result_.data)
	}
	result_Slice := unsafe.Slice((*unsafe.Pointer)(result_.data), int(result_.len))
	var goResult_ = make([]Window, len(result_Slice))
	for idx, r := range result_Slice {
		goResult_[idx] = MakeWindow(r)
	}
	return goResult_
}

func (n NSApplication) IsHidden() bool {
	result_ := C.C_NSApplication_IsHidden(n.Ptr())
	return bool(result_)
}

func (n NSApplication) OcclusionState() ApplicationOcclusionState {
	result_ := C.C_NSApplication_OcclusionState(n.Ptr())
	return ApplicationOcclusionState(uint(result_))
}

func (n NSApplication) ModalWindow() Window {
	result_ := C.C_NSApplication_ModalWindow(n.Ptr())
	return MakeWindow(result_)
}

func (n NSApplication) MainMenu() Menu {
	result_ := C.C_NSApplication_MainMenu(n.Ptr())
	return MakeMenu(result_)
}

func (n NSApplication) SetMainMenu(value Menu) {
	C.C_NSApplication_SetMainMenu(n.Ptr(), objc.ExtractPtr(value))
}

func (n NSApplication) IsAutomaticCustomizeTouchBarMenuItemEnabled() bool {
	result_ := C.C_NSApplication_IsAutomaticCustomizeTouchBarMenuItemEnabled(n.Ptr())
	return bool(result_)
}

func (n NSApplication) SetAutomaticCustomizeTouchBarMenuItemEnabled(value bool) {
	C.C_NSApplication_SetAutomaticCustomizeTouchBarMenuItemEnabled(n.Ptr(), C.bool(value))
}

func (n NSApplication) WindowsMenu() Menu {
	result_ := C.C_NSApplication_WindowsMenu(n.Ptr())
	return MakeMenu(result_)
}

func (n NSApplication) SetWindowsMenu(value Menu) {
	C.C_NSApplication_SetWindowsMenu(n.Ptr(), objc.ExtractPtr(value))
}

func (n NSApplication) ServicesMenu() Menu {
	result_ := C.C_NSApplication_ServicesMenu(n.Ptr())
	return MakeMenu(result_)
}

func (n NSApplication) SetServicesMenu(value Menu) {
	C.C_NSApplication_SetServicesMenu(n.Ptr(), objc.ExtractPtr(value))
}

type ApplicationDelegate struct {
	ApplicationWillFinishLaunching                               func(notification foundation.Notification)
	ApplicationDidFinishLaunching                                func(notification foundation.Notification)
	ApplicationWillBecomeActive                                  func(notification foundation.Notification)
	ApplicationDidBecomeActive                                   func(notification foundation.Notification)
	ApplicationWillResignActive                                  func(notification foundation.Notification)
	ApplicationDidResignActive                                   func(notification foundation.Notification)
	ApplicationShouldTerminate                                   func(sender Application) ApplicationTerminateReply
	ApplicationShouldTerminateAfterLastWindowClosed              func(sender Application) bool
	ApplicationWillTerminate                                     func(notification foundation.Notification)
	ApplicationWillHide                                          func(notification foundation.Notification)
	ApplicationDidHide                                           func(notification foundation.Notification)
	ApplicationWillUnhide                                        func(notification foundation.Notification)
	ApplicationDidUnhide                                         func(notification foundation.Notification)
	ApplicationWillUpdate                                        func(notification foundation.Notification)
	ApplicationDidUpdate                                         func(notification foundation.Notification)
	ApplicationShouldHandleReopen_HasVisibleWindows              func(sender Application, flag bool) bool
	ApplicationDockMenu                                          func(sender Application) Menu
	Application_WillPresentError                                 func(application Application, error foundation.Error) foundation.Error
	ApplicationDidChangeScreenParameters                         func(notification foundation.Notification)
	Application_WillContinueUserActivityWithType                 func(application Application, userActivityType string) bool
	Application_DidFailToContinueUserActivityWithType_Error      func(application Application, userActivityType string, error foundation.Error)
	Application_DidUpdateUserActivity                            func(application Application, userActivity foundation.UserActivity)
	Application_DidRegisterForRemoteNotificationsWithDeviceToken func(application Application, deviceToken []byte)
	Application_DidFailToRegisterForRemoteNotificationsWithError func(application Application, error foundation.Error)
	Application_DidReceiveRemoteNotification                     func(application Application, userInfo map[string]objc.Object)
	Application_OpenURLs                                         func(application Application, urls []foundation.URL)
	Application_OpenFile                                         func(sender Application, filename string) bool
	Application_OpenFileWithoutUI                                func(sender objc.Object, filename string) bool
	Application_OpenTempFile                                     func(sender Application, filename string) bool
	Application_OpenFiles                                        func(sender Application, filenames []string)
	ApplicationOpenUntitledFile                                  func(sender Application) bool
	ApplicationShouldOpenUntitledFile                            func(sender Application) bool
	Application_PrintFile                                        func(sender Application, filename string) bool
	Application_PrintFiles_WithSettings_ShowPrintPanels          func(application Application, fileNames []string, printSettings map[PrintInfoAttributeKey]objc.Object, showPrintPanels bool) ApplicationPrintReply
	Application_DidDecodeRestorableState                         func(app Application, coder foundation.Coder)
	Application_WillEncodeRestorableState                        func(app Application, coder foundation.Coder)
	ApplicationDidChangeOcclusionState                           func(notification foundation.Notification)
	Application_DelegateHandlesKey                               func(sender Application, key string) bool
}

func (delegate *ApplicationDelegate) ToObjc() objc.Object {
	h := cgo.NewHandle(delegate)
	ptr := C.WrapApplicationDelegate(C.uintptr_t(h))
	return objc.MakeObject(ptr)
}

//export applicationDelegate_ApplicationWillFinishLaunching
func applicationDelegate_ApplicationWillFinishLaunching(hp C.uintptr_t, notification unsafe.Pointer) {
	delegate := cgo.Handle(hp).Value().(*ApplicationDelegate)
	delegate.ApplicationWillFinishLaunching(foundation.MakeNotification(notification))
}

//export applicationDelegate_ApplicationDidFinishLaunching
func applicationDelegate_ApplicationDidFinishLaunching(hp C.uintptr_t, notification unsafe.Pointer) {
	delegate := cgo.Handle(hp).Value().(*ApplicationDelegate)
	delegate.ApplicationDidFinishLaunching(foundation.MakeNotification(notification))
}

//export applicationDelegate_ApplicationWillBecomeActive
func applicationDelegate_ApplicationWillBecomeActive(hp C.uintptr_t, notification unsafe.Pointer) {
	delegate := cgo.Handle(hp).Value().(*ApplicationDelegate)
	delegate.ApplicationWillBecomeActive(foundation.MakeNotification(notification))
}

//export applicationDelegate_ApplicationDidBecomeActive
func applicationDelegate_ApplicationDidBecomeActive(hp C.uintptr_t, notification unsafe.Pointer) {
	delegate := cgo.Handle(hp).Value().(*ApplicationDelegate)
	delegate.ApplicationDidBecomeActive(foundation.MakeNotification(notification))
}

//export applicationDelegate_ApplicationWillResignActive
func applicationDelegate_ApplicationWillResignActive(hp C.uintptr_t, notification unsafe.Pointer) {
	delegate := cgo.Handle(hp).Value().(*ApplicationDelegate)
	delegate.ApplicationWillResignActive(foundation.MakeNotification(notification))
}

//export applicationDelegate_ApplicationDidResignActive
func applicationDelegate_ApplicationDidResignActive(hp C.uintptr_t, notification unsafe.Pointer) {
	delegate := cgo.Handle(hp).Value().(*ApplicationDelegate)
	delegate.ApplicationDidResignActive(foundation.MakeNotification(notification))
}

//export applicationDelegate_ApplicationShouldTerminate
func applicationDelegate_ApplicationShouldTerminate(hp C.uintptr_t, sender unsafe.Pointer) C.uint {
	delegate := cgo.Handle(hp).Value().(*ApplicationDelegate)
	result := delegate.ApplicationShouldTerminate(MakeApplication(sender))
	return C.uint(uint(result))
}

//export applicationDelegate_ApplicationShouldTerminateAfterLastWindowClosed
func applicationDelegate_ApplicationShouldTerminateAfterLastWindowClosed(hp C.uintptr_t, sender unsafe.Pointer) C.bool {
	delegate := cgo.Handle(hp).Value().(*ApplicationDelegate)
	result := delegate.ApplicationShouldTerminateAfterLastWindowClosed(MakeApplication(sender))
	return C.bool(result)
}

//export applicationDelegate_ApplicationWillTerminate
func applicationDelegate_ApplicationWillTerminate(hp C.uintptr_t, notification unsafe.Pointer) {
	delegate := cgo.Handle(hp).Value().(*ApplicationDelegate)
	delegate.ApplicationWillTerminate(foundation.MakeNotification(notification))
}

//export applicationDelegate_ApplicationWillHide
func applicationDelegate_ApplicationWillHide(hp C.uintptr_t, notification unsafe.Pointer) {
	delegate := cgo.Handle(hp).Value().(*ApplicationDelegate)
	delegate.ApplicationWillHide(foundation.MakeNotification(notification))
}

//export applicationDelegate_ApplicationDidHide
func applicationDelegate_ApplicationDidHide(hp C.uintptr_t, notification unsafe.Pointer) {
	delegate := cgo.Handle(hp).Value().(*ApplicationDelegate)
	delegate.ApplicationDidHide(foundation.MakeNotification(notification))
}

//export applicationDelegate_ApplicationWillUnhide
func applicationDelegate_ApplicationWillUnhide(hp C.uintptr_t, notification unsafe.Pointer) {
	delegate := cgo.Handle(hp).Value().(*ApplicationDelegate)
	delegate.ApplicationWillUnhide(foundation.MakeNotification(notification))
}

//export applicationDelegate_ApplicationDidUnhide
func applicationDelegate_ApplicationDidUnhide(hp C.uintptr_t, notification unsafe.Pointer) {
	delegate := cgo.Handle(hp).Value().(*ApplicationDelegate)
	delegate.ApplicationDidUnhide(foundation.MakeNotification(notification))
}

//export applicationDelegate_ApplicationWillUpdate
func applicationDelegate_ApplicationWillUpdate(hp C.uintptr_t, notification unsafe.Pointer) {
	delegate := cgo.Handle(hp).Value().(*ApplicationDelegate)
	delegate.ApplicationWillUpdate(foundation.MakeNotification(notification))
}

//export applicationDelegate_ApplicationDidUpdate
func applicationDelegate_ApplicationDidUpdate(hp C.uintptr_t, notification unsafe.Pointer) {
	delegate := cgo.Handle(hp).Value().(*ApplicationDelegate)
	delegate.ApplicationDidUpdate(foundation.MakeNotification(notification))
}

//export applicationDelegate_ApplicationShouldHandleReopen_HasVisibleWindows
func applicationDelegate_ApplicationShouldHandleReopen_HasVisibleWindows(hp C.uintptr_t, sender unsafe.Pointer, flag C.bool) C.bool {
	delegate := cgo.Handle(hp).Value().(*ApplicationDelegate)
	result := delegate.ApplicationShouldHandleReopen_HasVisibleWindows(MakeApplication(sender), bool(flag))
	return C.bool(result)
}

//export applicationDelegate_ApplicationDockMenu
func applicationDelegate_ApplicationDockMenu(hp C.uintptr_t, sender unsafe.Pointer) unsafe.Pointer {
	delegate := cgo.Handle(hp).Value().(*ApplicationDelegate)
	result := delegate.ApplicationDockMenu(MakeApplication(sender))
	return objc.ExtractPtr(result)
}

//export applicationDelegate_Application_WillPresentError
func applicationDelegate_Application_WillPresentError(hp C.uintptr_t, application unsafe.Pointer, error unsafe.Pointer) unsafe.Pointer {
	delegate := cgo.Handle(hp).Value().(*ApplicationDelegate)
	result := delegate.Application_WillPresentError(MakeApplication(application), foundation.MakeError(error))
	return objc.ExtractPtr(result)
}

//export applicationDelegate_ApplicationDidChangeScreenParameters
func applicationDelegate_ApplicationDidChangeScreenParameters(hp C.uintptr_t, notification unsafe.Pointer) {
	delegate := cgo.Handle(hp).Value().(*ApplicationDelegate)
	delegate.ApplicationDidChangeScreenParameters(foundation.MakeNotification(notification))
}

//export applicationDelegate_Application_WillContinueUserActivityWithType
func applicationDelegate_Application_WillContinueUserActivityWithType(hp C.uintptr_t, application unsafe.Pointer, userActivityType unsafe.Pointer) C.bool {
	delegate := cgo.Handle(hp).Value().(*ApplicationDelegate)
	result := delegate.Application_WillContinueUserActivityWithType(MakeApplication(application), foundation.MakeString(userActivityType).String())
	return C.bool(result)
}

//export applicationDelegate_Application_DidFailToContinueUserActivityWithType_Error
func applicationDelegate_Application_DidFailToContinueUserActivityWithType_Error(hp C.uintptr_t, application unsafe.Pointer, userActivityType unsafe.Pointer, error unsafe.Pointer) {
	delegate := cgo.Handle(hp).Value().(*ApplicationDelegate)
	delegate.Application_DidFailToContinueUserActivityWithType_Error(MakeApplication(application), foundation.MakeString(userActivityType).String(), foundation.MakeError(error))
}

//export applicationDelegate_Application_DidUpdateUserActivity
func applicationDelegate_Application_DidUpdateUserActivity(hp C.uintptr_t, application unsafe.Pointer, userActivity unsafe.Pointer) {
	delegate := cgo.Handle(hp).Value().(*ApplicationDelegate)
	delegate.Application_DidUpdateUserActivity(MakeApplication(application), foundation.MakeUserActivity(userActivity))
}

//export applicationDelegate_Application_DidRegisterForRemoteNotificationsWithDeviceToken
func applicationDelegate_Application_DidRegisterForRemoteNotificationsWithDeviceToken(hp C.uintptr_t, application unsafe.Pointer, deviceToken unsafe.Pointer) {
	delegate := cgo.Handle(hp).Value().(*ApplicationDelegate)
	delegate.Application_DidRegisterForRemoteNotificationsWithDeviceToken(MakeApplication(application), foundation.MakeData(deviceToken).ToBytes())
}

//export applicationDelegate_Application_DidFailToRegisterForRemoteNotificationsWithError
func applicationDelegate_Application_DidFailToRegisterForRemoteNotificationsWithError(hp C.uintptr_t, application unsafe.Pointer, error unsafe.Pointer) {
	delegate := cgo.Handle(hp).Value().(*ApplicationDelegate)
	delegate.Application_DidFailToRegisterForRemoteNotificationsWithError(MakeApplication(application), foundation.MakeError(error))
}

//export applicationDelegate_Application_DidReceiveRemoteNotification
func applicationDelegate_Application_DidReceiveRemoteNotification(hp C.uintptr_t, application unsafe.Pointer, userInfo C.Dictionary) {
	delegate := cgo.Handle(hp).Value().(*ApplicationDelegate)
	if userInfo.len > 0 {
		defer C.free(userInfo.key_data)
		defer C.free(userInfo.value_data)
	}
	userInfoKeySlice := unsafe.Slice((*unsafe.Pointer)(userInfo.key_data), int(userInfo.len))
	userInfoValueSlice := unsafe.Slice((*unsafe.Pointer)(userInfo.value_data), int(userInfo.len))
	var goUserInfo = make(map[string]objc.Object)
	for idx, k := range userInfoKeySlice {
		v := userInfoValueSlice[idx]
		goUserInfo[foundation.MakeString(k).String()] = objc.MakeObject(v)
	}
	delegate.Application_DidReceiveRemoteNotification(MakeApplication(application), goUserInfo)
}

//export applicationDelegate_Application_OpenURLs
func applicationDelegate_Application_OpenURLs(hp C.uintptr_t, application unsafe.Pointer, urls C.Array) {
	delegate := cgo.Handle(hp).Value().(*ApplicationDelegate)
	if urls.len > 0 {
		defer C.free(urls.data)
	}
	urlsSlice := unsafe.Slice((*unsafe.Pointer)(urls.data), int(urls.len))
	var goUrls = make([]foundation.URL, len(urlsSlice))
	for idx, r := range urlsSlice {
		goUrls[idx] = foundation.MakeURL(r)
	}
	delegate.Application_OpenURLs(MakeApplication(application), goUrls)
}

//export applicationDelegate_Application_OpenFile
func applicationDelegate_Application_OpenFile(hp C.uintptr_t, sender unsafe.Pointer, filename unsafe.Pointer) C.bool {
	delegate := cgo.Handle(hp).Value().(*ApplicationDelegate)
	result := delegate.Application_OpenFile(MakeApplication(sender), foundation.MakeString(filename).String())
	return C.bool(result)
}

//export applicationDelegate_Application_OpenFileWithoutUI
func applicationDelegate_Application_OpenFileWithoutUI(hp C.uintptr_t, sender unsafe.Pointer, filename unsafe.Pointer) C.bool {
	delegate := cgo.Handle(hp).Value().(*ApplicationDelegate)
	result := delegate.Application_OpenFileWithoutUI(objc.MakeObject(sender), foundation.MakeString(filename).String())
	return C.bool(result)
}

//export applicationDelegate_Application_OpenTempFile
func applicationDelegate_Application_OpenTempFile(hp C.uintptr_t, sender unsafe.Pointer, filename unsafe.Pointer) C.bool {
	delegate := cgo.Handle(hp).Value().(*ApplicationDelegate)
	result := delegate.Application_OpenTempFile(MakeApplication(sender), foundation.MakeString(filename).String())
	return C.bool(result)
}

//export applicationDelegate_Application_OpenFiles
func applicationDelegate_Application_OpenFiles(hp C.uintptr_t, sender unsafe.Pointer, filenames C.Array) {
	delegate := cgo.Handle(hp).Value().(*ApplicationDelegate)
	if filenames.len > 0 {
		defer C.free(filenames.data)
	}
	filenamesSlice := unsafe.Slice((*unsafe.Pointer)(filenames.data), int(filenames.len))
	var goFilenames = make([]string, len(filenamesSlice))
	for idx, r := range filenamesSlice {
		goFilenames[idx] = foundation.MakeString(r).String()
	}
	delegate.Application_OpenFiles(MakeApplication(sender), goFilenames)
}

//export applicationDelegate_ApplicationOpenUntitledFile
func applicationDelegate_ApplicationOpenUntitledFile(hp C.uintptr_t, sender unsafe.Pointer) C.bool {
	delegate := cgo.Handle(hp).Value().(*ApplicationDelegate)
	result := delegate.ApplicationOpenUntitledFile(MakeApplication(sender))
	return C.bool(result)
}

//export applicationDelegate_ApplicationShouldOpenUntitledFile
func applicationDelegate_ApplicationShouldOpenUntitledFile(hp C.uintptr_t, sender unsafe.Pointer) C.bool {
	delegate := cgo.Handle(hp).Value().(*ApplicationDelegate)
	result := delegate.ApplicationShouldOpenUntitledFile(MakeApplication(sender))
	return C.bool(result)
}

//export applicationDelegate_Application_PrintFile
func applicationDelegate_Application_PrintFile(hp C.uintptr_t, sender unsafe.Pointer, filename unsafe.Pointer) C.bool {
	delegate := cgo.Handle(hp).Value().(*ApplicationDelegate)
	result := delegate.Application_PrintFile(MakeApplication(sender), foundation.MakeString(filename).String())
	return C.bool(result)
}

//export applicationDelegate_Application_PrintFiles_WithSettings_ShowPrintPanels
func applicationDelegate_Application_PrintFiles_WithSettings_ShowPrintPanels(hp C.uintptr_t, application unsafe.Pointer, fileNames C.Array, printSettings C.Dictionary, showPrintPanels C.bool) C.uint {
	delegate := cgo.Handle(hp).Value().(*ApplicationDelegate)
	if fileNames.len > 0 {
		defer C.free(fileNames.data)
	}
	fileNamesSlice := unsafe.Slice((*unsafe.Pointer)(fileNames.data), int(fileNames.len))
	var goFileNames = make([]string, len(fileNamesSlice))
	for idx, r := range fileNamesSlice {
		goFileNames[idx] = foundation.MakeString(r).String()
	}
	if printSettings.len > 0 {
		defer C.free(printSettings.key_data)
		defer C.free(printSettings.value_data)
	}
	printSettingsKeySlice := unsafe.Slice((*unsafe.Pointer)(printSettings.key_data), int(printSettings.len))
	printSettingsValueSlice := unsafe.Slice((*unsafe.Pointer)(printSettings.value_data), int(printSettings.len))
	var goPrintSettings = make(map[PrintInfoAttributeKey]objc.Object)
	for idx, k := range printSettingsKeySlice {
		v := printSettingsValueSlice[idx]
		goPrintSettings[PrintInfoAttributeKey(foundation.MakeString(k).String())] = objc.MakeObject(v)
	}
	result := delegate.Application_PrintFiles_WithSettings_ShowPrintPanels(MakeApplication(application), goFileNames, goPrintSettings, bool(showPrintPanels))
	return C.uint(uint(result))
}

//export applicationDelegate_Application_DidDecodeRestorableState
func applicationDelegate_Application_DidDecodeRestorableState(hp C.uintptr_t, app unsafe.Pointer, coder unsafe.Pointer) {
	delegate := cgo.Handle(hp).Value().(*ApplicationDelegate)
	delegate.Application_DidDecodeRestorableState(MakeApplication(app), foundation.MakeCoder(coder))
}

//export applicationDelegate_Application_WillEncodeRestorableState
func applicationDelegate_Application_WillEncodeRestorableState(hp C.uintptr_t, app unsafe.Pointer, coder unsafe.Pointer) {
	delegate := cgo.Handle(hp).Value().(*ApplicationDelegate)
	delegate.Application_WillEncodeRestorableState(MakeApplication(app), foundation.MakeCoder(coder))
}

//export applicationDelegate_ApplicationDidChangeOcclusionState
func applicationDelegate_ApplicationDidChangeOcclusionState(hp C.uintptr_t, notification unsafe.Pointer) {
	delegate := cgo.Handle(hp).Value().(*ApplicationDelegate)
	delegate.ApplicationDidChangeOcclusionState(foundation.MakeNotification(notification))
}

//export applicationDelegate_Application_DelegateHandlesKey
func applicationDelegate_Application_DelegateHandlesKey(hp C.uintptr_t, sender unsafe.Pointer, key unsafe.Pointer) C.bool {
	delegate := cgo.Handle(hp).Value().(*ApplicationDelegate)
	result := delegate.Application_DelegateHandlesKey(MakeApplication(sender), foundation.MakeString(key).String())
	return C.bool(result)
}

//export ApplicationDelegate_RespondsTo
func ApplicationDelegate_RespondsTo(hp C.uintptr_t, selectorPtr unsafe.Pointer) bool {
	sel := objc.Selector(selectorPtr)
	selName := objc.Sel_GetName(sel)
	delegate := cgo.Handle(hp).Value().(*ApplicationDelegate)
	switch selName {
	case "applicationWillFinishLaunching:":
		return delegate.ApplicationWillFinishLaunching != nil
	case "applicationDidFinishLaunching:":
		return delegate.ApplicationDidFinishLaunching != nil
	case "applicationWillBecomeActive:":
		return delegate.ApplicationWillBecomeActive != nil
	case "applicationDidBecomeActive:":
		return delegate.ApplicationDidBecomeActive != nil
	case "applicationWillResignActive:":
		return delegate.ApplicationWillResignActive != nil
	case "applicationDidResignActive:":
		return delegate.ApplicationDidResignActive != nil
	case "applicationShouldTerminate:":
		return delegate.ApplicationShouldTerminate != nil
	case "applicationShouldTerminateAfterLastWindowClosed:":
		return delegate.ApplicationShouldTerminateAfterLastWindowClosed != nil
	case "applicationWillTerminate:":
		return delegate.ApplicationWillTerminate != nil
	case "applicationWillHide:":
		return delegate.ApplicationWillHide != nil
	case "applicationDidHide:":
		return delegate.ApplicationDidHide != nil
	case "applicationWillUnhide:":
		return delegate.ApplicationWillUnhide != nil
	case "applicationDidUnhide:":
		return delegate.ApplicationDidUnhide != nil
	case "applicationWillUpdate:":
		return delegate.ApplicationWillUpdate != nil
	case "applicationDidUpdate:":
		return delegate.ApplicationDidUpdate != nil
	case "applicationShouldHandleReopen:hasVisibleWindows:":
		return delegate.ApplicationShouldHandleReopen_HasVisibleWindows != nil
	case "applicationDockMenu:":
		return delegate.ApplicationDockMenu != nil
	case "application:willPresentError:":
		return delegate.Application_WillPresentError != nil
	case "applicationDidChangeScreenParameters:":
		return delegate.ApplicationDidChangeScreenParameters != nil
	case "application:willContinueUserActivityWithType:":
		return delegate.Application_WillContinueUserActivityWithType != nil
	case "application:didFailToContinueUserActivityWithType:error:":
		return delegate.Application_DidFailToContinueUserActivityWithType_Error != nil
	case "application:didUpdateUserActivity:":
		return delegate.Application_DidUpdateUserActivity != nil
	case "application:didRegisterForRemoteNotificationsWithDeviceToken:":
		return delegate.Application_DidRegisterForRemoteNotificationsWithDeviceToken != nil
	case "application:didFailToRegisterForRemoteNotificationsWithError:":
		return delegate.Application_DidFailToRegisterForRemoteNotificationsWithError != nil
	case "application:didReceiveRemoteNotification:":
		return delegate.Application_DidReceiveRemoteNotification != nil
	case "application:openURLs:":
		return delegate.Application_OpenURLs != nil
	case "application:openFile:":
		return delegate.Application_OpenFile != nil
	case "application:openFileWithoutUI:":
		return delegate.Application_OpenFileWithoutUI != nil
	case "application:openTempFile:":
		return delegate.Application_OpenTempFile != nil
	case "application:openFiles:":
		return delegate.Application_OpenFiles != nil
	case "applicationOpenUntitledFile:":
		return delegate.ApplicationOpenUntitledFile != nil
	case "applicationShouldOpenUntitledFile:":
		return delegate.ApplicationShouldOpenUntitledFile != nil
	case "application:printFile:":
		return delegate.Application_PrintFile != nil
	case "application:printFiles:withSettings:showPrintPanels:":
		return delegate.Application_PrintFiles_WithSettings_ShowPrintPanels != nil
	case "application:didDecodeRestorableState:":
		return delegate.Application_DidDecodeRestorableState != nil
	case "application:willEncodeRestorableState:":
		return delegate.Application_WillEncodeRestorableState != nil
	case "applicationDidChangeOcclusionState:":
		return delegate.ApplicationDidChangeOcclusionState != nil
	case "application:delegateHandlesKey:":
		return delegate.Application_DelegateHandlesKey != nil
	default:
		return false
	}
}
