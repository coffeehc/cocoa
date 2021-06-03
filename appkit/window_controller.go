package appkit

// #include "window_controller.h"
import "C"
import (
	"github.com/hsiafan/cocoa/foundation"
	"github.com/hsiafan/cocoa/objc"
	"unsafe"
)

type WindowController interface {
	Responder
	LoadWindow()
	ShowWindow(sender objc.Object)
	WindowDidLoad()
	WindowWillLoad()
	SetDocumentEdited(dirtyFlag bool)
	Close()
	SynchronizeWindowTitleWithDocumentName()
	WindowTitleForDocumentDisplayName(displayName string) string
	DismissController(sender objc.Object)
	IsWindowLoaded() bool
	Window() Window
	SetWindow(value Window)
	Document() objc.Object
	SetDocument(value objc.Object)
	ShouldCloseDocument() bool
	SetShouldCloseDocument(value bool)
	Owner() objc.Object
	Storyboard() Storyboard
	WindowNibName() NibName
	WindowNibPath() string
	ShouldCascadeWindows() bool
	SetShouldCascadeWindows(value bool)
	WindowFrameAutosaveName() WindowFrameAutosaveName
	SetWindowFrameAutosaveName(value WindowFrameAutosaveName)
	ContentViewController() ViewController
	SetContentViewController(value ViewController)
}

type NSWindowController struct {
	NSResponder
}

func MakeWindowController(ptr unsafe.Pointer) NSWindowController {
	return NSWindowController{
		NSResponder: MakeResponder(ptr),
	}
}

func AllocWindowController() NSWindowController {
	return MakeWindowController(C.C_WindowController_Alloc())
}

func (n NSWindowController) InitWithWindow(window Window) WindowController {
	result_ := C.C_NSWindowController_InitWithWindow(n.Ptr(), objc.ExtractPtr(window))
	return MakeWindowController(result_)
}

func (n NSWindowController) InitWithWindowNibName(windowNibName NibName) WindowController {
	result_ := C.C_NSWindowController_InitWithWindowNibName(n.Ptr(), foundation.NewString(string(windowNibName)).Ptr())
	return MakeWindowController(result_)
}

func (n NSWindowController) InitWithWindowNibName_Owner(windowNibName NibName, owner objc.Object) WindowController {
	result_ := C.C_NSWindowController_InitWithWindowNibName_Owner(n.Ptr(), foundation.NewString(string(windowNibName)).Ptr(), objc.ExtractPtr(owner))
	return MakeWindowController(result_)
}

func (n NSWindowController) InitWithWindowNibPath_Owner(windowNibPath string, owner objc.Object) WindowController {
	result_ := C.C_NSWindowController_InitWithWindowNibPath_Owner(n.Ptr(), foundation.NewString(windowNibPath).Ptr(), objc.ExtractPtr(owner))
	return MakeWindowController(result_)
}

func (n NSWindowController) InitWithCoder(coder foundation.Coder) WindowController {
	result_ := C.C_NSWindowController_InitWithCoder(n.Ptr(), objc.ExtractPtr(coder))
	return MakeWindowController(result_)
}

func (n NSWindowController) Init() WindowController {
	result_ := C.C_NSWindowController_Init(n.Ptr())
	return MakeWindowController(result_)
}

func (n NSWindowController) LoadWindow() {
	C.C_NSWindowController_LoadWindow(n.Ptr())
}

func (n NSWindowController) ShowWindow(sender objc.Object) {
	C.C_NSWindowController_ShowWindow(n.Ptr(), objc.ExtractPtr(sender))
}

func (n NSWindowController) WindowDidLoad() {
	C.C_NSWindowController_WindowDidLoad(n.Ptr())
}

func (n NSWindowController) WindowWillLoad() {
	C.C_NSWindowController_WindowWillLoad(n.Ptr())
}

func (n NSWindowController) SetDocumentEdited(dirtyFlag bool) {
	C.C_NSWindowController_SetDocumentEdited(n.Ptr(), C.bool(dirtyFlag))
}

func (n NSWindowController) Close() {
	C.C_NSWindowController_Close(n.Ptr())
}

func (n NSWindowController) SynchronizeWindowTitleWithDocumentName() {
	C.C_NSWindowController_SynchronizeWindowTitleWithDocumentName(n.Ptr())
}

func (n NSWindowController) WindowTitleForDocumentDisplayName(displayName string) string {
	result_ := C.C_NSWindowController_WindowTitleForDocumentDisplayName(n.Ptr(), foundation.NewString(displayName).Ptr())
	return foundation.MakeString(result_).String()
}

func (n NSWindowController) DismissController(sender objc.Object) {
	C.C_NSWindowController_DismissController(n.Ptr(), objc.ExtractPtr(sender))
}

func (n NSWindowController) IsWindowLoaded() bool {
	result_ := C.C_NSWindowController_IsWindowLoaded(n.Ptr())
	return bool(result_)
}

func (n NSWindowController) Window() Window {
	result_ := C.C_NSWindowController_Window(n.Ptr())
	return MakeWindow(result_)
}

func (n NSWindowController) SetWindow(value Window) {
	C.C_NSWindowController_SetWindow(n.Ptr(), objc.ExtractPtr(value))
}

func (n NSWindowController) Document() objc.Object {
	result_ := C.C_NSWindowController_Document(n.Ptr())
	return objc.MakeObject(result_)
}

func (n NSWindowController) SetDocument(value objc.Object) {
	C.C_NSWindowController_SetDocument(n.Ptr(), objc.ExtractPtr(value))
}

func (n NSWindowController) ShouldCloseDocument() bool {
	result_ := C.C_NSWindowController_ShouldCloseDocument(n.Ptr())
	return bool(result_)
}

func (n NSWindowController) SetShouldCloseDocument(value bool) {
	C.C_NSWindowController_SetShouldCloseDocument(n.Ptr(), C.bool(value))
}

func (n NSWindowController) Owner() objc.Object {
	result_ := C.C_NSWindowController_Owner(n.Ptr())
	return objc.MakeObject(result_)
}

func (n NSWindowController) Storyboard() Storyboard {
	result_ := C.C_NSWindowController_Storyboard(n.Ptr())
	return MakeStoryboard(result_)
}

func (n NSWindowController) WindowNibName() NibName {
	result_ := C.C_NSWindowController_WindowNibName(n.Ptr())
	return NibName(foundation.MakeString(result_).String())
}

func (n NSWindowController) WindowNibPath() string {
	result_ := C.C_NSWindowController_WindowNibPath(n.Ptr())
	return foundation.MakeString(result_).String()
}

func (n NSWindowController) ShouldCascadeWindows() bool {
	result_ := C.C_NSWindowController_ShouldCascadeWindows(n.Ptr())
	return bool(result_)
}

func (n NSWindowController) SetShouldCascadeWindows(value bool) {
	C.C_NSWindowController_SetShouldCascadeWindows(n.Ptr(), C.bool(value))
}

func (n NSWindowController) WindowFrameAutosaveName() WindowFrameAutosaveName {
	result_ := C.C_NSWindowController_WindowFrameAutosaveName(n.Ptr())
	return WindowFrameAutosaveName(foundation.MakeString(result_).String())
}

func (n NSWindowController) SetWindowFrameAutosaveName(value WindowFrameAutosaveName) {
	C.C_NSWindowController_SetWindowFrameAutosaveName(n.Ptr(), foundation.NewString(string(value)).Ptr())
}

func (n NSWindowController) ContentViewController() ViewController {
	result_ := C.C_NSWindowController_ContentViewController(n.Ptr())
	return MakeViewController(result_)
}

func (n NSWindowController) SetContentViewController(value ViewController) {
	C.C_NSWindowController_SetContentViewController(n.Ptr(), objc.ExtractPtr(value))
}
