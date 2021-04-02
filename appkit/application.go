package appkit

// #cgo CFLAGS: -x objective-c
// #cgo LDFLAGS: -framework AppKit
// #include "application.h"
import "C"
import (
	"github.com/hsiafan/cocoa/foundation"
	"github.com/hsiafan/cocoa/objc"
	"unsafe"
)

type Application interface {
	Responder
	CurrentEvent() Event
	IsRunning() bool
	IsActive() bool
	MainMenu() Menu
	SetMainMenu(mainMenu Menu)
	WindowsMenu() Menu
	SetWindowsMenu(windowsMenu Menu)
	ServicesMenu() Menu
	SetServicesMenu(servicesMenu Menu)
	MainWindow() Window
	Windows() []Window
	IsHidden() bool
	Run()
	FinishLaunching()
	Stop(sender objc.Object)
	Terminate(sender objc.Object)
	SendEvent(event Event)
	ActivateIgnoringOtherApps(flag bool)
	Deactivate()
	SetActivationPolicy(activationPolicy ApplicationActivationPolicy)
	WindowWithWindowNumber(windowNum int)
	MiniaturizeAll(sender objc.Object)
	Hide(sender objc.Object)
	Unhide(sender objc.Object)
	UnhideWithoutActivation()
	UpdateWindows()
	SetWindowsNeedUpdate(needUpdate bool)
	ArrangeInFront(sender objc.Object)
	PreventWindowOrdering()
	ApplicationWillFinishLaunching(callback func(notification foundation.Notification))
	_applicationWillFinishLaunching() func(notification foundation.Notification)
	ApplicationDidFinishLaunching(callback func(notification foundation.Notification))
	_applicationDidFinishLaunching() func(notification foundation.Notification)
	ApplicationShouldTerminateAfterLastWindowClosed(callback func(theApplication Application) bool)
	_applicationShouldTerminateAfterLastWindowClosed() func(theApplication Application) bool
}

var _ Application = (*NSApplication)(nil)

type NSApplication struct {
	NSResponder
	applicationWillFinishLaunching                  func(notification foundation.Notification)
	applicationDidFinishLaunching                   func(notification foundation.Notification)
	applicationShouldTerminateAfterLastWindowClosed func(theApplication Application) bool
}

func MakeApplication(ptr unsafe.Pointer) *NSApplication {
	if ptr == nil {
		return nil
	}
	return &NSApplication{
		NSResponder: *MakeResponder(ptr),
	}
}
func (a *NSApplication) setDelegate() {
	id := resources.Register(a)
	C.Application_SetDelegate(a.Ptr(), C.long(id))
}

func (a *NSApplication) CurrentEvent() Event {
	return MakeEvent(C.Application_CurrentEvent(a.Ptr()))
}

func (a *NSApplication) IsRunning() bool {
	return bool(C.Application_IsRunning(a.Ptr()))
}

func (a *NSApplication) IsActive() bool {
	return bool(C.Application_IsActive(a.Ptr()))
}

func (a *NSApplication) MainMenu() Menu {
	return MakeMenu(C.Application_MainMenu(a.Ptr()))
}

func (a *NSApplication) SetMainMenu(mainMenu Menu) {
	C.Application_SetMainMenu(a.Ptr(), toPointer(mainMenu))
}

func (a *NSApplication) WindowsMenu() Menu {
	return MakeMenu(C.Application_WindowsMenu(a.Ptr()))
}

func (a *NSApplication) SetWindowsMenu(windowsMenu Menu) {
	C.Application_SetWindowsMenu(a.Ptr(), toPointer(windowsMenu))
}

func (a *NSApplication) ServicesMenu() Menu {
	return MakeMenu(C.Application_ServicesMenu(a.Ptr()))
}

func (a *NSApplication) SetServicesMenu(servicesMenu Menu) {
	C.Application_SetServicesMenu(a.Ptr(), toPointer(servicesMenu))
}

func (a *NSApplication) MainWindow() Window {
	return MakeWindow(C.Application_MainWindow(a.Ptr()))
}

func (a *NSApplication) Windows() []Window {
	var cArray C.Array = C.Application_Windows(a.Ptr())
	defer C.free(cArray.data)
	result := (*[1 << 28]unsafe.Pointer)(unsafe.Pointer(cArray.data))[:cArray.len:cArray.len]
	var goArray = make([]Window, len(result))
	for idx, r := range result {
		goArray[idx] = MakeWindow(r)
	}
	return goArray
}

func (a *NSApplication) IsHidden() bool {
	return bool(C.Application_IsHidden(a.Ptr()))
}

func (a *NSApplication) Run() {
	C.Application_Run(a.Ptr())
}

func (a *NSApplication) FinishLaunching() {
	C.Application_FinishLaunching(a.Ptr())
}

func (a *NSApplication) Stop(sender objc.Object) {
	C.Application_Stop(a.Ptr(), toPointer(sender))
}

func (a *NSApplication) Terminate(sender objc.Object) {
	C.Application_Terminate(a.Ptr(), toPointer(sender))
}

func (a *NSApplication) SendEvent(event Event) {
	C.Application_SendEvent(a.Ptr(), toPointer(event))
}

func (a *NSApplication) ActivateIgnoringOtherApps(flag bool) {
	C.Application_ActivateIgnoringOtherApps(a.Ptr(), C.bool(flag))
}

func (a *NSApplication) Deactivate() {
	C.Application_Deactivate(a.Ptr())
}

func SharedApplication() Application {
	return MakeApplication(C.Application_SharedApplication())
}

func (a *NSApplication) SetActivationPolicy(activationPolicy ApplicationActivationPolicy) {
	C.Application_SetActivationPolicy(a.Ptr(), C.long(activationPolicy))
}

func (a *NSApplication) WindowWithWindowNumber(windowNum int) {
	C.Application_WindowWithWindowNumber(a.Ptr(), C.long(windowNum))
}

func (a *NSApplication) MiniaturizeAll(sender objc.Object) {
	C.Application_MiniaturizeAll(a.Ptr(), toPointer(sender))
}

func (a *NSApplication) Hide(sender objc.Object) {
	C.Application_Hide(a.Ptr(), toPointer(sender))
}

func (a *NSApplication) Unhide(sender objc.Object) {
	C.Application_Unhide(a.Ptr(), toPointer(sender))
}

func (a *NSApplication) UnhideWithoutActivation() {
	C.Application_UnhideWithoutActivation(a.Ptr())
}

func (a *NSApplication) UpdateWindows() {
	C.Application_UpdateWindows(a.Ptr())
}

func (a *NSApplication) SetWindowsNeedUpdate(needUpdate bool) {
	C.Application_SetWindowsNeedUpdate(a.Ptr(), C.bool(needUpdate))
}

func (a *NSApplication) ArrangeInFront(sender objc.Object) {
	C.Application_ArrangeInFront(a.Ptr(), toPointer(sender))
}

func (a *NSApplication) PreventWindowOrdering() {
	C.Application_PreventWindowOrdering(a.Ptr())
}

func (a *NSApplication) ApplicationWillFinishLaunching(callback func(notification foundation.Notification)) {
	a.applicationWillFinishLaunching = callback
}

func (a *NSApplication) _applicationWillFinishLaunching() func(notification foundation.Notification) {
	return a.applicationWillFinishLaunching
}

func (a *NSApplication) ApplicationDidFinishLaunching(callback func(notification foundation.Notification)) {
	a.applicationDidFinishLaunching = callback
}

func (a *NSApplication) _applicationDidFinishLaunching() func(notification foundation.Notification) {
	return a.applicationDidFinishLaunching
}

func (a *NSApplication) ApplicationShouldTerminateAfterLastWindowClosed(callback func(theApplication Application) bool) {
	a.applicationShouldTerminateAfterLastWindowClosed = callback
}

func (a *NSApplication) _applicationShouldTerminateAfterLastWindowClosed() func(theApplication Application) bool {
	return a.applicationShouldTerminateAfterLastWindowClosed
}

//export Application_Delegate_ApplicationWillFinishLaunching
func Application_Delegate_ApplicationWillFinishLaunching(id int64, notification unsafe.Pointer) {
	application := resources.Get(id).(Application)
	callback := application._applicationWillFinishLaunching()
	if callback != nil {
		callback(foundation.MakeNotification(notification))
	}
}

//export Application_Delegate_ApplicationDidFinishLaunching
func Application_Delegate_ApplicationDidFinishLaunching(id int64, notification unsafe.Pointer) {
	application := resources.Get(id).(Application)
	callback := application._applicationDidFinishLaunching()
	if callback != nil {
		callback(foundation.MakeNotification(notification))
	}
}

//export Application_Delegate_ApplicationShouldTerminateAfterLastWindowClosed
func Application_Delegate_ApplicationShouldTerminateAfterLastWindowClosed(id int64, theApplication unsafe.Pointer) bool {
	application := resources.Get(id).(Application)
	callback := application._applicationShouldTerminateAfterLastWindowClosed()
	if callback != nil {
		return callback(MakeApplication(theApplication))
	}
	return true
}
