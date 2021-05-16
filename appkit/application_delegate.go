package appkit

// #cgo CFLAGS: -x objective-c
// #cgo LDFLAGS: -framework AppKit
// #include "application_delegate.h"
import "C"
import (
	"github.com/hsiafan/cocoa/foundation"
	"github.com/hsiafan/cocoa/objc"
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
	Application_OpenFile                                         func(sender Application, filename string) bool
	Application_OpenFileWithoutUI                                func(sender objc.Object, filename string) bool
	Application_OpenTempFile                                     func(sender Application, filename string) bool
	ApplicationOpenUntitledFile                                  func(sender Application) bool
	ApplicationShouldOpenUntitledFile                            func(sender Application) bool
	Application_PrintFile                                        func(sender Application, filename string) bool
	Application_DidDecodeRestorableState                         func(app Application, coder foundation.Coder)
	Application_WillEncodeRestorableState                        func(app Application, coder foundation.Coder)
	ApplicationDidChangeOcclusionState                           func(notification foundation.Notification)
	Application_DelegateHandlesKey                               func(sender Application, key string) bool
}

func WrapApplicationDelegate(delegate *ApplicationDelegate) objc.Object {
	id := resources.NextId()
	resources.Store(id, delegate)
	ptr := C.WrapApplicationDelegate(C.long(id))
	return objc.MakeObject(ptr)
}

//export ApplicationWillFinishLaunching
func ApplicationWillFinishLaunching(id int64, notification unsafe.Pointer) {
	delegate := resources.Get(id).(*ApplicationDelegate)
	delegate.ApplicationWillFinishLaunching(foundation.MakeNotification(notification))
}

//export ApplicationDidFinishLaunching
func ApplicationDidFinishLaunching(id int64, notification unsafe.Pointer) {
	delegate := resources.Get(id).(*ApplicationDelegate)
	delegate.ApplicationDidFinishLaunching(foundation.MakeNotification(notification))
}

//export ApplicationWillBecomeActive
func ApplicationWillBecomeActive(id int64, notification unsafe.Pointer) {
	delegate := resources.Get(id).(*ApplicationDelegate)
	delegate.ApplicationWillBecomeActive(foundation.MakeNotification(notification))
}

//export ApplicationDidBecomeActive
func ApplicationDidBecomeActive(id int64, notification unsafe.Pointer) {
	delegate := resources.Get(id).(*ApplicationDelegate)
	delegate.ApplicationDidBecomeActive(foundation.MakeNotification(notification))
}

//export ApplicationWillResignActive
func ApplicationWillResignActive(id int64, notification unsafe.Pointer) {
	delegate := resources.Get(id).(*ApplicationDelegate)
	delegate.ApplicationWillResignActive(foundation.MakeNotification(notification))
}

//export ApplicationDidResignActive
func ApplicationDidResignActive(id int64, notification unsafe.Pointer) {
	delegate := resources.Get(id).(*ApplicationDelegate)
	delegate.ApplicationDidResignActive(foundation.MakeNotification(notification))
}

//export ApplicationShouldTerminate
func ApplicationShouldTerminate(id int64, sender unsafe.Pointer) C.uint {
	delegate := resources.Get(id).(*ApplicationDelegate)
	result := delegate.ApplicationShouldTerminate(MakeApplication(sender))
	return C.uint(uint(result))
}

//export ApplicationShouldTerminateAfterLastWindowClosed
func ApplicationShouldTerminateAfterLastWindowClosed(id int64, sender unsafe.Pointer) C.bool {
	delegate := resources.Get(id).(*ApplicationDelegate)
	result := delegate.ApplicationShouldTerminateAfterLastWindowClosed(MakeApplication(sender))
	return C.bool(result)
}

//export ApplicationWillTerminate
func ApplicationWillTerminate(id int64, notification unsafe.Pointer) {
	delegate := resources.Get(id).(*ApplicationDelegate)
	delegate.ApplicationWillTerminate(foundation.MakeNotification(notification))
}

//export ApplicationWillHide
func ApplicationWillHide(id int64, notification unsafe.Pointer) {
	delegate := resources.Get(id).(*ApplicationDelegate)
	delegate.ApplicationWillHide(foundation.MakeNotification(notification))
}

//export ApplicationDidHide
func ApplicationDidHide(id int64, notification unsafe.Pointer) {
	delegate := resources.Get(id).(*ApplicationDelegate)
	delegate.ApplicationDidHide(foundation.MakeNotification(notification))
}

//export ApplicationWillUnhide
func ApplicationWillUnhide(id int64, notification unsafe.Pointer) {
	delegate := resources.Get(id).(*ApplicationDelegate)
	delegate.ApplicationWillUnhide(foundation.MakeNotification(notification))
}

//export ApplicationDidUnhide
func ApplicationDidUnhide(id int64, notification unsafe.Pointer) {
	delegate := resources.Get(id).(*ApplicationDelegate)
	delegate.ApplicationDidUnhide(foundation.MakeNotification(notification))
}

//export ApplicationWillUpdate
func ApplicationWillUpdate(id int64, notification unsafe.Pointer) {
	delegate := resources.Get(id).(*ApplicationDelegate)
	delegate.ApplicationWillUpdate(foundation.MakeNotification(notification))
}

//export ApplicationDidUpdate
func ApplicationDidUpdate(id int64, notification unsafe.Pointer) {
	delegate := resources.Get(id).(*ApplicationDelegate)
	delegate.ApplicationDidUpdate(foundation.MakeNotification(notification))
}

//export ApplicationShouldHandleReopen_HasVisibleWindows
func ApplicationShouldHandleReopen_HasVisibleWindows(id int64, sender unsafe.Pointer, flag C.bool) C.bool {
	delegate := resources.Get(id).(*ApplicationDelegate)
	result := delegate.ApplicationShouldHandleReopen_HasVisibleWindows(MakeApplication(sender), bool(flag))
	return C.bool(result)
}

//export ApplicationDockMenu
func ApplicationDockMenu(id int64, sender unsafe.Pointer) unsafe.Pointer {
	delegate := resources.Get(id).(*ApplicationDelegate)
	result := delegate.ApplicationDockMenu(MakeApplication(sender))
	return objc.ExtractPtr(result)
}

//export Application_WillPresentError
func Application_WillPresentError(id int64, application unsafe.Pointer, error unsafe.Pointer) unsafe.Pointer {
	delegate := resources.Get(id).(*ApplicationDelegate)
	result := delegate.Application_WillPresentError(MakeApplication(application), foundation.MakeError(error))
	return objc.ExtractPtr(result)
}

//export ApplicationDidChangeScreenParameters
func ApplicationDidChangeScreenParameters(id int64, notification unsafe.Pointer) {
	delegate := resources.Get(id).(*ApplicationDelegate)
	delegate.ApplicationDidChangeScreenParameters(foundation.MakeNotification(notification))
}

//export Application_WillContinueUserActivityWithType
func Application_WillContinueUserActivityWithType(id int64, application unsafe.Pointer, userActivityType unsafe.Pointer) C.bool {
	delegate := resources.Get(id).(*ApplicationDelegate)
	result := delegate.Application_WillContinueUserActivityWithType(MakeApplication(application), foundation.MakeString(userActivityType).String())
	return C.bool(result)
}

//export Application_DidFailToContinueUserActivityWithType_Error
func Application_DidFailToContinueUserActivityWithType_Error(id int64, application unsafe.Pointer, userActivityType unsafe.Pointer, error unsafe.Pointer) {
	delegate := resources.Get(id).(*ApplicationDelegate)
	delegate.Application_DidFailToContinueUserActivityWithType_Error(MakeApplication(application), foundation.MakeString(userActivityType).String(), foundation.MakeError(error))
}

//export Application_DidUpdateUserActivity
func Application_DidUpdateUserActivity(id int64, application unsafe.Pointer, userActivity unsafe.Pointer) {
	delegate := resources.Get(id).(*ApplicationDelegate)
	delegate.Application_DidUpdateUserActivity(MakeApplication(application), foundation.MakeUserActivity(userActivity))
}

//export Application_DidRegisterForRemoteNotificationsWithDeviceToken
func Application_DidRegisterForRemoteNotificationsWithDeviceToken(id int64, application unsafe.Pointer, deviceToken C.Array) {
	delegate := resources.Get(id).(*ApplicationDelegate)
	deviceTokenBuffer := (*[1 << 30]byte)(deviceToken.data)[:C.int(deviceToken.len)]
	goDeviceToken := make([]byte, C.int(deviceToken.len))
	copy(goDeviceToken, deviceTokenBuffer)
	C.free(deviceToken.data)
	delegate.Application_DidRegisterForRemoteNotificationsWithDeviceToken(MakeApplication(application), goDeviceToken)
}

//export Application_DidFailToRegisterForRemoteNotificationsWithError
func Application_DidFailToRegisterForRemoteNotificationsWithError(id int64, application unsafe.Pointer, error unsafe.Pointer) {
	delegate := resources.Get(id).(*ApplicationDelegate)
	delegate.Application_DidFailToRegisterForRemoteNotificationsWithError(MakeApplication(application), foundation.MakeError(error))
}

//export Application_OpenFile
func Application_OpenFile(id int64, sender unsafe.Pointer, filename unsafe.Pointer) C.bool {
	delegate := resources.Get(id).(*ApplicationDelegate)
	result := delegate.Application_OpenFile(MakeApplication(sender), foundation.MakeString(filename).String())
	return C.bool(result)
}

//export Application_OpenFileWithoutUI
func Application_OpenFileWithoutUI(id int64, sender unsafe.Pointer, filename unsafe.Pointer) C.bool {
	delegate := resources.Get(id).(*ApplicationDelegate)
	result := delegate.Application_OpenFileWithoutUI(objc.MakeObject(sender), foundation.MakeString(filename).String())
	return C.bool(result)
}

//export Application_OpenTempFile
func Application_OpenTempFile(id int64, sender unsafe.Pointer, filename unsafe.Pointer) C.bool {
	delegate := resources.Get(id).(*ApplicationDelegate)
	result := delegate.Application_OpenTempFile(MakeApplication(sender), foundation.MakeString(filename).String())
	return C.bool(result)
}

//export ApplicationOpenUntitledFile
func ApplicationOpenUntitledFile(id int64, sender unsafe.Pointer) C.bool {
	delegate := resources.Get(id).(*ApplicationDelegate)
	result := delegate.ApplicationOpenUntitledFile(MakeApplication(sender))
	return C.bool(result)
}

//export ApplicationShouldOpenUntitledFile
func ApplicationShouldOpenUntitledFile(id int64, sender unsafe.Pointer) C.bool {
	delegate := resources.Get(id).(*ApplicationDelegate)
	result := delegate.ApplicationShouldOpenUntitledFile(MakeApplication(sender))
	return C.bool(result)
}

//export Application_PrintFile
func Application_PrintFile(id int64, sender unsafe.Pointer, filename unsafe.Pointer) C.bool {
	delegate := resources.Get(id).(*ApplicationDelegate)
	result := delegate.Application_PrintFile(MakeApplication(sender), foundation.MakeString(filename).String())
	return C.bool(result)
}

//export Application_DidDecodeRestorableState
func Application_DidDecodeRestorableState(id int64, app unsafe.Pointer, coder unsafe.Pointer) {
	delegate := resources.Get(id).(*ApplicationDelegate)
	delegate.Application_DidDecodeRestorableState(MakeApplication(app), foundation.MakeCoder(coder))
}

//export Application_WillEncodeRestorableState
func Application_WillEncodeRestorableState(id int64, app unsafe.Pointer, coder unsafe.Pointer) {
	delegate := resources.Get(id).(*ApplicationDelegate)
	delegate.Application_WillEncodeRestorableState(MakeApplication(app), foundation.MakeCoder(coder))
}

//export ApplicationDidChangeOcclusionState
func ApplicationDidChangeOcclusionState(id int64, notification unsafe.Pointer) {
	delegate := resources.Get(id).(*ApplicationDelegate)
	delegate.ApplicationDidChangeOcclusionState(foundation.MakeNotification(notification))
}

//export Application_DelegateHandlesKey
func Application_DelegateHandlesKey(id int64, sender unsafe.Pointer, key unsafe.Pointer) C.bool {
	delegate := resources.Get(id).(*ApplicationDelegate)
	result := delegate.Application_DelegateHandlesKey(MakeApplication(sender), foundation.MakeString(key).String())
	return C.bool(result)
}

//export ApplicationDelegate_RespondsTo
func ApplicationDelegate_RespondsTo(id int64, selectorPtr unsafe.Pointer) bool {
	sel := objc.MakeSelector(selectorPtr)
	selName := objc.Sel_GetName(sel)
	delegate := resources.Get(id).(*ApplicationDelegate)
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
	case "application:openFile:":
		return delegate.Application_OpenFile != nil
	case "application:openFileWithoutUI:":
		return delegate.Application_OpenFileWithoutUI != nil
	case "application:openTempFile:":
		return delegate.Application_OpenTempFile != nil
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
func deleteApplicationDelegate(id int64) {
	resources.Delete(id)
}
