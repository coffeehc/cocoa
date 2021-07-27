package appkit

// #include "application_delegate.h"
import "C"
import (
	"github.com/hsiafan/cocoa/foundation"
	"github.com/hsiafan/cocoa/objc"
	"runtime/cgo"
	"unsafe"
)

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
	Application_OpenURLs                                         func(application Application, urls []foundation.URL)
	Application_OpenFile                                         func(sender Application, filename string) bool
	Application_OpenFileWithoutUI                                func(sender objc.Object, filename string) bool
	Application_OpenTempFile                                     func(sender Application, filename string) bool
	Application_OpenFiles                                        func(sender Application, filenames []string)
	ApplicationOpenUntitledFile                                  func(sender Application) bool
	ApplicationShouldOpenUntitledFile                            func(sender Application) bool
	Application_PrintFile                                        func(sender Application, filename string) bool
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
func applicationDelegate_Application_DidRegisterForRemoteNotificationsWithDeviceToken(hp C.uintptr_t, application unsafe.Pointer, deviceToken C.Array) {
	delegate := cgo.Handle(hp).Value().(*ApplicationDelegate)
	deviceTokenBuffer := (*[1 << 30]byte)(deviceToken.data)[:C.int(deviceToken.len)]
	goDeviceToken := make([]byte, C.int(deviceToken.len))
	copy(goDeviceToken, deviceTokenBuffer)
	C.free(deviceToken.data)
	delegate.Application_DidRegisterForRemoteNotificationsWithDeviceToken(MakeApplication(application), goDeviceToken)
}

//export applicationDelegate_Application_DidFailToRegisterForRemoteNotificationsWithError
func applicationDelegate_Application_DidFailToRegisterForRemoteNotificationsWithError(hp C.uintptr_t, application unsafe.Pointer, error unsafe.Pointer) {
	delegate := cgo.Handle(hp).Value().(*ApplicationDelegate)
	delegate.Application_DidFailToRegisterForRemoteNotificationsWithError(MakeApplication(application), foundation.MakeError(error))
}

//export applicationDelegate_Application_OpenURLs
func applicationDelegate_Application_OpenURLs(hp C.uintptr_t, application unsafe.Pointer, urls C.Array) {
	delegate := cgo.Handle(hp).Value().(*ApplicationDelegate)
	defer C.free(urls.data)
	urlsSlice := (*[1 << 28]unsafe.Pointer)(unsafe.Pointer(urls.data))[:urls.len:urls.len]
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
	defer C.free(filenames.data)
	filenamesSlice := (*[1 << 28]unsafe.Pointer)(unsafe.Pointer(filenames.data))[:filenames.len:filenames.len]
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

//export deleteApplicationDelegate
func deleteApplicationDelegate(hp C.uintptr_t) {
	cgo.Handle(hp).Delete()
}
