package appkit

// #cgo CFLAGS: -x objective-c
// #cgo LDFLAGS: -framework AppKit
// #include "window_delegate.h"
import "C"
import (
	"github.com/hsiafan/cocoa/coregraphics"
	"github.com/hsiafan/cocoa/foundation"
	"github.com/hsiafan/cocoa/objc"
	"unsafe"
)

type WindowDelegate struct {
	Window_WillPositionSheet_UsingRect                                    func(window Window, sheet Window, rect foundation.Rect) foundation.Rect
	WindowWillBeginSheet                                                  func(notification foundation.Notification)
	WindowDidEndSheet                                                     func(notification foundation.Notification)
	WindowWillResize_ToSize                                               func(sender Window, frameSize foundation.Size) foundation.Size
	WindowDidResize                                                       func(notification foundation.Notification)
	WindowWillStartLiveResize                                             func(notification foundation.Notification)
	WindowDidEndLiveResize                                                func(notification foundation.Notification)
	WindowWillMiniaturize                                                 func(notification foundation.Notification)
	WindowDidMiniaturize                                                  func(notification foundation.Notification)
	WindowDidDeminiaturize                                                func(notification foundation.Notification)
	WindowWillUseStandardFrame_DefaultFrame                               func(window Window, newFrame foundation.Rect) foundation.Rect
	WindowShouldZoom_ToFrame                                              func(window Window, newFrame foundation.Rect) bool
	Window_WillUseFullScreenContentSize                                   func(window Window, proposedSize foundation.Size) foundation.Size
	Window_WillUseFullScreenPresentationOptions                           func(window Window, proposedOptions ApplicationPresentationOptions) ApplicationPresentationOptions
	WindowWillEnterFullScreen                                             func(notification foundation.Notification)
	WindowDidEnterFullScreen                                              func(notification foundation.Notification)
	WindowWillExitFullScreen                                              func(notification foundation.Notification)
	WindowDidExitFullScreen                                               func(notification foundation.Notification)
	Window_StartCustomAnimationToEnterFullScreenWithDuration              func(window Window, duration foundation.TimeInterval)
	Window_StartCustomAnimationToEnterFullScreenOnScreen_WithDuration     func(window Window, screen Screen, duration foundation.TimeInterval)
	WindowDidFailToEnterFullScreen                                        func(window Window)
	Window_StartCustomAnimationToExitFullScreenWithDuration               func(window Window, duration foundation.TimeInterval)
	WindowDidFailToExitFullScreen                                         func(window Window)
	WindowWillMove                                                        func(notification foundation.Notification)
	WindowDidMove                                                         func(notification foundation.Notification)
	WindowDidChangeScreen                                                 func(notification foundation.Notification)
	WindowDidChangeScreenProfile                                          func(notification foundation.Notification)
	WindowDidChangeBackingProperties                                      func(notification foundation.Notification)
	WindowShouldClose                                                     func(sender Window) bool
	WindowWillClose                                                       func(notification foundation.Notification)
	WindowDidBecomeKey                                                    func(notification foundation.Notification)
	WindowDidResignKey                                                    func(notification foundation.Notification)
	WindowDidBecomeMain                                                   func(notification foundation.Notification)
	WindowDidResignMain                                                   func(notification foundation.Notification)
	WindowWillReturnFieldEditor_ToObject                                  func(sender Window, client objc.Object) objc.Object
	WindowDidUpdate                                                       func(notification foundation.Notification)
	WindowDidExpose                                                       func(notification foundation.Notification)
	WindowDidChangeOcclusionState                                         func(notification foundation.Notification)
	Window_ShouldDragDocumentWithEvent_From_WithPasteboard                func(window Window, event Event, dragImageLocation foundation.Point, pasteboard Pasteboard) bool
	WindowWillReturnUndoManager                                           func(window Window) foundation.UndoManager
	Window_ShouldPopUpDocumentPathMenu                                    func(window Window, menu Menu) bool
	Window_WillEncodeRestorableState                                      func(window Window, state foundation.Coder)
	Window_DidDecodeRestorableState                                       func(window Window, state foundation.Coder)
	Window_WillResizeForVersionBrowserWithMaxPreferredSize_MaxAllowedSize func(window Window, maxPreferredFrameSize foundation.Size, maxAllowedFrameSize foundation.Size) foundation.Size
	WindowWillEnterVersionBrowser                                         func(notification foundation.Notification)
	WindowDidEnterVersionBrowser                                          func(notification foundation.Notification)
	WindowWillExitVersionBrowser                                          func(notification foundation.Notification)
	WindowDidExitVersionBrowser                                           func(notification foundation.Notification)
}

func WrapWindowDelegate(delegate *WindowDelegate) objc.Object {
	id := resources.NextId()
	resources.Store(id, delegate)
	ptr := C.WrapWindowDelegate(C.long(id))
	return objc.MakeObject(ptr)
}

//export Window_WillPositionSheet_UsingRect
func Window_WillPositionSheet_UsingRect(id int64, window unsafe.Pointer, sheet unsafe.Pointer, rect C.CGRect) C.CGRect {
	delegate := resources.Get(id).(*WindowDelegate)
	result := delegate.Window_WillPositionSheet_UsingRect(MakeWindow(window), MakeWindow(sheet), foundation.Rect(coregraphics.FromCGRectPointer(unsafe.Pointer(&rect))))
	return *(*C.CGRect)(coregraphics.ToCGRectPointer(coregraphics.Rect(result)))
}

//export WindowWillBeginSheet
func WindowWillBeginSheet(id int64, notification unsafe.Pointer) {
	delegate := resources.Get(id).(*WindowDelegate)
	delegate.WindowWillBeginSheet(foundation.MakeNotification(notification))
}

//export WindowDidEndSheet
func WindowDidEndSheet(id int64, notification unsafe.Pointer) {
	delegate := resources.Get(id).(*WindowDelegate)
	delegate.WindowDidEndSheet(foundation.MakeNotification(notification))
}

//export WindowWillResize_ToSize
func WindowWillResize_ToSize(id int64, sender unsafe.Pointer, frameSize C.CGSize) C.CGSize {
	delegate := resources.Get(id).(*WindowDelegate)
	result := delegate.WindowWillResize_ToSize(MakeWindow(sender), foundation.Size(coregraphics.FromCGSizePointer(unsafe.Pointer(&frameSize))))
	return *(*C.CGSize)(coregraphics.ToCGSizePointer(coregraphics.Size(result)))
}

//export WindowDidResize
func WindowDidResize(id int64, notification unsafe.Pointer) {
	delegate := resources.Get(id).(*WindowDelegate)
	delegate.WindowDidResize(foundation.MakeNotification(notification))
}

//export WindowWillStartLiveResize
func WindowWillStartLiveResize(id int64, notification unsafe.Pointer) {
	delegate := resources.Get(id).(*WindowDelegate)
	delegate.WindowWillStartLiveResize(foundation.MakeNotification(notification))
}

//export WindowDidEndLiveResize
func WindowDidEndLiveResize(id int64, notification unsafe.Pointer) {
	delegate := resources.Get(id).(*WindowDelegate)
	delegate.WindowDidEndLiveResize(foundation.MakeNotification(notification))
}

//export WindowWillMiniaturize
func WindowWillMiniaturize(id int64, notification unsafe.Pointer) {
	delegate := resources.Get(id).(*WindowDelegate)
	delegate.WindowWillMiniaturize(foundation.MakeNotification(notification))
}

//export WindowDidMiniaturize
func WindowDidMiniaturize(id int64, notification unsafe.Pointer) {
	delegate := resources.Get(id).(*WindowDelegate)
	delegate.WindowDidMiniaturize(foundation.MakeNotification(notification))
}

//export WindowDidDeminiaturize
func WindowDidDeminiaturize(id int64, notification unsafe.Pointer) {
	delegate := resources.Get(id).(*WindowDelegate)
	delegate.WindowDidDeminiaturize(foundation.MakeNotification(notification))
}

//export WindowWillUseStandardFrame_DefaultFrame
func WindowWillUseStandardFrame_DefaultFrame(id int64, window unsafe.Pointer, newFrame C.CGRect) C.CGRect {
	delegate := resources.Get(id).(*WindowDelegate)
	result := delegate.WindowWillUseStandardFrame_DefaultFrame(MakeWindow(window), foundation.Rect(coregraphics.FromCGRectPointer(unsafe.Pointer(&newFrame))))
	return *(*C.CGRect)(coregraphics.ToCGRectPointer(coregraphics.Rect(result)))
}

//export WindowShouldZoom_ToFrame
func WindowShouldZoom_ToFrame(id int64, window unsafe.Pointer, newFrame C.CGRect) C.bool {
	delegate := resources.Get(id).(*WindowDelegate)
	result := delegate.WindowShouldZoom_ToFrame(MakeWindow(window), foundation.Rect(coregraphics.FromCGRectPointer(unsafe.Pointer(&newFrame))))
	return C.bool(result)
}

//export Window_WillUseFullScreenContentSize
func Window_WillUseFullScreenContentSize(id int64, window unsafe.Pointer, proposedSize C.CGSize) C.CGSize {
	delegate := resources.Get(id).(*WindowDelegate)
	result := delegate.Window_WillUseFullScreenContentSize(MakeWindow(window), foundation.Size(coregraphics.FromCGSizePointer(unsafe.Pointer(&proposedSize))))
	return *(*C.CGSize)(coregraphics.ToCGSizePointer(coregraphics.Size(result)))
}

//export Window_WillUseFullScreenPresentationOptions
func Window_WillUseFullScreenPresentationOptions(id int64, window unsafe.Pointer, proposedOptions C.uint) C.uint {
	delegate := resources.Get(id).(*WindowDelegate)
	result := delegate.Window_WillUseFullScreenPresentationOptions(MakeWindow(window), ApplicationPresentationOptions(uint(proposedOptions)))
	return C.uint(uint(result))
}

//export WindowWillEnterFullScreen
func WindowWillEnterFullScreen(id int64, notification unsafe.Pointer) {
	delegate := resources.Get(id).(*WindowDelegate)
	delegate.WindowWillEnterFullScreen(foundation.MakeNotification(notification))
}

//export WindowDidEnterFullScreen
func WindowDidEnterFullScreen(id int64, notification unsafe.Pointer) {
	delegate := resources.Get(id).(*WindowDelegate)
	delegate.WindowDidEnterFullScreen(foundation.MakeNotification(notification))
}

//export WindowWillExitFullScreen
func WindowWillExitFullScreen(id int64, notification unsafe.Pointer) {
	delegate := resources.Get(id).(*WindowDelegate)
	delegate.WindowWillExitFullScreen(foundation.MakeNotification(notification))
}

//export WindowDidExitFullScreen
func WindowDidExitFullScreen(id int64, notification unsafe.Pointer) {
	delegate := resources.Get(id).(*WindowDelegate)
	delegate.WindowDidExitFullScreen(foundation.MakeNotification(notification))
}

//export Window_StartCustomAnimationToEnterFullScreenWithDuration
func Window_StartCustomAnimationToEnterFullScreenWithDuration(id int64, window unsafe.Pointer, duration C.double) {
	delegate := resources.Get(id).(*WindowDelegate)
	delegate.Window_StartCustomAnimationToEnterFullScreenWithDuration(MakeWindow(window), foundation.TimeInterval(float64(duration)))
}

//export Window_StartCustomAnimationToEnterFullScreenOnScreen_WithDuration
func Window_StartCustomAnimationToEnterFullScreenOnScreen_WithDuration(id int64, window unsafe.Pointer, screen unsafe.Pointer, duration C.double) {
	delegate := resources.Get(id).(*WindowDelegate)
	delegate.Window_StartCustomAnimationToEnterFullScreenOnScreen_WithDuration(MakeWindow(window), MakeScreen(screen), foundation.TimeInterval(float64(duration)))
}

//export WindowDidFailToEnterFullScreen
func WindowDidFailToEnterFullScreen(id int64, window unsafe.Pointer) {
	delegate := resources.Get(id).(*WindowDelegate)
	delegate.WindowDidFailToEnterFullScreen(MakeWindow(window))
}

//export Window_StartCustomAnimationToExitFullScreenWithDuration
func Window_StartCustomAnimationToExitFullScreenWithDuration(id int64, window unsafe.Pointer, duration C.double) {
	delegate := resources.Get(id).(*WindowDelegate)
	delegate.Window_StartCustomAnimationToExitFullScreenWithDuration(MakeWindow(window), foundation.TimeInterval(float64(duration)))
}

//export WindowDidFailToExitFullScreen
func WindowDidFailToExitFullScreen(id int64, window unsafe.Pointer) {
	delegate := resources.Get(id).(*WindowDelegate)
	delegate.WindowDidFailToExitFullScreen(MakeWindow(window))
}

//export WindowWillMove
func WindowWillMove(id int64, notification unsafe.Pointer) {
	delegate := resources.Get(id).(*WindowDelegate)
	delegate.WindowWillMove(foundation.MakeNotification(notification))
}

//export WindowDidMove
func WindowDidMove(id int64, notification unsafe.Pointer) {
	delegate := resources.Get(id).(*WindowDelegate)
	delegate.WindowDidMove(foundation.MakeNotification(notification))
}

//export WindowDidChangeScreen
func WindowDidChangeScreen(id int64, notification unsafe.Pointer) {
	delegate := resources.Get(id).(*WindowDelegate)
	delegate.WindowDidChangeScreen(foundation.MakeNotification(notification))
}

//export WindowDidChangeScreenProfile
func WindowDidChangeScreenProfile(id int64, notification unsafe.Pointer) {
	delegate := resources.Get(id).(*WindowDelegate)
	delegate.WindowDidChangeScreenProfile(foundation.MakeNotification(notification))
}

//export WindowDidChangeBackingProperties
func WindowDidChangeBackingProperties(id int64, notification unsafe.Pointer) {
	delegate := resources.Get(id).(*WindowDelegate)
	delegate.WindowDidChangeBackingProperties(foundation.MakeNotification(notification))
}

//export WindowShouldClose
func WindowShouldClose(id int64, sender unsafe.Pointer) C.bool {
	delegate := resources.Get(id).(*WindowDelegate)
	result := delegate.WindowShouldClose(MakeWindow(sender))
	return C.bool(result)
}

//export WindowWillClose
func WindowWillClose(id int64, notification unsafe.Pointer) {
	delegate := resources.Get(id).(*WindowDelegate)
	delegate.WindowWillClose(foundation.MakeNotification(notification))
}

//export WindowDidBecomeKey
func WindowDidBecomeKey(id int64, notification unsafe.Pointer) {
	delegate := resources.Get(id).(*WindowDelegate)
	delegate.WindowDidBecomeKey(foundation.MakeNotification(notification))
}

//export WindowDidResignKey
func WindowDidResignKey(id int64, notification unsafe.Pointer) {
	delegate := resources.Get(id).(*WindowDelegate)
	delegate.WindowDidResignKey(foundation.MakeNotification(notification))
}

//export WindowDidBecomeMain
func WindowDidBecomeMain(id int64, notification unsafe.Pointer) {
	delegate := resources.Get(id).(*WindowDelegate)
	delegate.WindowDidBecomeMain(foundation.MakeNotification(notification))
}

//export WindowDidResignMain
func WindowDidResignMain(id int64, notification unsafe.Pointer) {
	delegate := resources.Get(id).(*WindowDelegate)
	delegate.WindowDidResignMain(foundation.MakeNotification(notification))
}

//export WindowWillReturnFieldEditor_ToObject
func WindowWillReturnFieldEditor_ToObject(id int64, sender unsafe.Pointer, client unsafe.Pointer) unsafe.Pointer {
	delegate := resources.Get(id).(*WindowDelegate)
	result := delegate.WindowWillReturnFieldEditor_ToObject(MakeWindow(sender), objc.MakeObject(client))
	return objc.ExtractPtr(result)
}

//export WindowDidUpdate
func WindowDidUpdate(id int64, notification unsafe.Pointer) {
	delegate := resources.Get(id).(*WindowDelegate)
	delegate.WindowDidUpdate(foundation.MakeNotification(notification))
}

//export WindowDidExpose
func WindowDidExpose(id int64, notification unsafe.Pointer) {
	delegate := resources.Get(id).(*WindowDelegate)
	delegate.WindowDidExpose(foundation.MakeNotification(notification))
}

//export WindowDidChangeOcclusionState
func WindowDidChangeOcclusionState(id int64, notification unsafe.Pointer) {
	delegate := resources.Get(id).(*WindowDelegate)
	delegate.WindowDidChangeOcclusionState(foundation.MakeNotification(notification))
}

//export Window_ShouldDragDocumentWithEvent_From_WithPasteboard
func Window_ShouldDragDocumentWithEvent_From_WithPasteboard(id int64, window unsafe.Pointer, event unsafe.Pointer, dragImageLocation C.CGPoint, pasteboard unsafe.Pointer) C.bool {
	delegate := resources.Get(id).(*WindowDelegate)
	result := delegate.Window_ShouldDragDocumentWithEvent_From_WithPasteboard(MakeWindow(window), MakeEvent(event), foundation.Point(coregraphics.FromCGPointPointer(unsafe.Pointer(&dragImageLocation))), MakePasteboard(pasteboard))
	return C.bool(result)
}

//export WindowWillReturnUndoManager
func WindowWillReturnUndoManager(id int64, window unsafe.Pointer) unsafe.Pointer {
	delegate := resources.Get(id).(*WindowDelegate)
	result := delegate.WindowWillReturnUndoManager(MakeWindow(window))
	return objc.ExtractPtr(result)
}

//export Window_ShouldPopUpDocumentPathMenu
func Window_ShouldPopUpDocumentPathMenu(id int64, window unsafe.Pointer, menu unsafe.Pointer) C.bool {
	delegate := resources.Get(id).(*WindowDelegate)
	result := delegate.Window_ShouldPopUpDocumentPathMenu(MakeWindow(window), MakeMenu(menu))
	return C.bool(result)
}

//export Window_WillEncodeRestorableState
func Window_WillEncodeRestorableState(id int64, window unsafe.Pointer, state unsafe.Pointer) {
	delegate := resources.Get(id).(*WindowDelegate)
	delegate.Window_WillEncodeRestorableState(MakeWindow(window), foundation.MakeCoder(state))
}

//export Window_DidDecodeRestorableState
func Window_DidDecodeRestorableState(id int64, window unsafe.Pointer, state unsafe.Pointer) {
	delegate := resources.Get(id).(*WindowDelegate)
	delegate.Window_DidDecodeRestorableState(MakeWindow(window), foundation.MakeCoder(state))
}

//export Window_WillResizeForVersionBrowserWithMaxPreferredSize_MaxAllowedSize
func Window_WillResizeForVersionBrowserWithMaxPreferredSize_MaxAllowedSize(id int64, window unsafe.Pointer, maxPreferredFrameSize C.CGSize, maxAllowedFrameSize C.CGSize) C.CGSize {
	delegate := resources.Get(id).(*WindowDelegate)
	result := delegate.Window_WillResizeForVersionBrowserWithMaxPreferredSize_MaxAllowedSize(MakeWindow(window), foundation.Size(coregraphics.FromCGSizePointer(unsafe.Pointer(&maxPreferredFrameSize))), foundation.Size(coregraphics.FromCGSizePointer(unsafe.Pointer(&maxAllowedFrameSize))))
	return *(*C.CGSize)(coregraphics.ToCGSizePointer(coregraphics.Size(result)))
}

//export WindowWillEnterVersionBrowser
func WindowWillEnterVersionBrowser(id int64, notification unsafe.Pointer) {
	delegate := resources.Get(id).(*WindowDelegate)
	delegate.WindowWillEnterVersionBrowser(foundation.MakeNotification(notification))
}

//export WindowDidEnterVersionBrowser
func WindowDidEnterVersionBrowser(id int64, notification unsafe.Pointer) {
	delegate := resources.Get(id).(*WindowDelegate)
	delegate.WindowDidEnterVersionBrowser(foundation.MakeNotification(notification))
}

//export WindowWillExitVersionBrowser
func WindowWillExitVersionBrowser(id int64, notification unsafe.Pointer) {
	delegate := resources.Get(id).(*WindowDelegate)
	delegate.WindowWillExitVersionBrowser(foundation.MakeNotification(notification))
}

//export WindowDidExitVersionBrowser
func WindowDidExitVersionBrowser(id int64, notification unsafe.Pointer) {
	delegate := resources.Get(id).(*WindowDelegate)
	delegate.WindowDidExitVersionBrowser(foundation.MakeNotification(notification))
}

//export WindowDelegate_RespondsTo
func WindowDelegate_RespondsTo(id int64, selectorPtr unsafe.Pointer) bool {
	sel := objc.MakeSelector(selectorPtr)
	selName := objc.Sel_GetName(sel)
	delegate := resources.Get(id).(*WindowDelegate)
	switch selName {
	case "window:willPositionSheet:usingRect:":
		return delegate.Window_WillPositionSheet_UsingRect != nil
	case "windowWillBeginSheet:":
		return delegate.WindowWillBeginSheet != nil
	case "windowDidEndSheet:":
		return delegate.WindowDidEndSheet != nil
	case "windowWillResize:toSize:":
		return delegate.WindowWillResize_ToSize != nil
	case "windowDidResize:":
		return delegate.WindowDidResize != nil
	case "windowWillStartLiveResize:":
		return delegate.WindowWillStartLiveResize != nil
	case "windowDidEndLiveResize:":
		return delegate.WindowDidEndLiveResize != nil
	case "windowWillMiniaturize:":
		return delegate.WindowWillMiniaturize != nil
	case "windowDidMiniaturize:":
		return delegate.WindowDidMiniaturize != nil
	case "windowDidDeminiaturize:":
		return delegate.WindowDidDeminiaturize != nil
	case "windowWillUseStandardFrame:defaultFrame:":
		return delegate.WindowWillUseStandardFrame_DefaultFrame != nil
	case "windowShouldZoom:toFrame:":
		return delegate.WindowShouldZoom_ToFrame != nil
	case "window:willUseFullScreenContentSize:":
		return delegate.Window_WillUseFullScreenContentSize != nil
	case "window:willUseFullScreenPresentationOptions:":
		return delegate.Window_WillUseFullScreenPresentationOptions != nil
	case "windowWillEnterFullScreen:":
		return delegate.WindowWillEnterFullScreen != nil
	case "windowDidEnterFullScreen:":
		return delegate.WindowDidEnterFullScreen != nil
	case "windowWillExitFullScreen:":
		return delegate.WindowWillExitFullScreen != nil
	case "windowDidExitFullScreen:":
		return delegate.WindowDidExitFullScreen != nil
	case "window:startCustomAnimationToEnterFullScreenWithDuration:":
		return delegate.Window_StartCustomAnimationToEnterFullScreenWithDuration != nil
	case "window:startCustomAnimationToEnterFullScreenOnScreen:withDuration:":
		return delegate.Window_StartCustomAnimationToEnterFullScreenOnScreen_WithDuration != nil
	case "windowDidFailToEnterFullScreen:":
		return delegate.WindowDidFailToEnterFullScreen != nil
	case "window:startCustomAnimationToExitFullScreenWithDuration:":
		return delegate.Window_StartCustomAnimationToExitFullScreenWithDuration != nil
	case "windowDidFailToExitFullScreen:":
		return delegate.WindowDidFailToExitFullScreen != nil
	case "windowWillMove:":
		return delegate.WindowWillMove != nil
	case "windowDidMove:":
		return delegate.WindowDidMove != nil
	case "windowDidChangeScreen:":
		return delegate.WindowDidChangeScreen != nil
	case "windowDidChangeScreenProfile:":
		return delegate.WindowDidChangeScreenProfile != nil
	case "windowDidChangeBackingProperties:":
		return delegate.WindowDidChangeBackingProperties != nil
	case "windowShouldClose:":
		return delegate.WindowShouldClose != nil
	case "windowWillClose:":
		return delegate.WindowWillClose != nil
	case "windowDidBecomeKey:":
		return delegate.WindowDidBecomeKey != nil
	case "windowDidResignKey:":
		return delegate.WindowDidResignKey != nil
	case "windowDidBecomeMain:":
		return delegate.WindowDidBecomeMain != nil
	case "windowDidResignMain:":
		return delegate.WindowDidResignMain != nil
	case "windowWillReturnFieldEditor:toObject:":
		return delegate.WindowWillReturnFieldEditor_ToObject != nil
	case "windowDidUpdate:":
		return delegate.WindowDidUpdate != nil
	case "windowDidExpose:":
		return delegate.WindowDidExpose != nil
	case "windowDidChangeOcclusionState:":
		return delegate.WindowDidChangeOcclusionState != nil
	case "window:shouldDragDocumentWithEvent:from:withPasteboard:":
		return delegate.Window_ShouldDragDocumentWithEvent_From_WithPasteboard != nil
	case "windowWillReturnUndoManager:":
		return delegate.WindowWillReturnUndoManager != nil
	case "window:shouldPopUpDocumentPathMenu:":
		return delegate.Window_ShouldPopUpDocumentPathMenu != nil
	case "window:willEncodeRestorableState:":
		return delegate.Window_WillEncodeRestorableState != nil
	case "window:didDecodeRestorableState:":
		return delegate.Window_DidDecodeRestorableState != nil
	case "window:willResizeForVersionBrowserWithMaxPreferredSize:maxAllowedSize:":
		return delegate.Window_WillResizeForVersionBrowserWithMaxPreferredSize_MaxAllowedSize != nil
	case "windowWillEnterVersionBrowser:":
		return delegate.WindowWillEnterVersionBrowser != nil
	case "windowDidEnterVersionBrowser:":
		return delegate.WindowDidEnterVersionBrowser != nil
	case "windowWillExitVersionBrowser:":
		return delegate.WindowWillExitVersionBrowser != nil
	case "windowDidExitVersionBrowser:":
		return delegate.WindowDidExitVersionBrowser != nil
	default:
		return false
	}
}

//export deleteWindowDelegate
func deleteWindowDelegate(id int64) {
	resources.Delete(id)
}