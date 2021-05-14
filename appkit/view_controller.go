package appkit

// #cgo CFLAGS: -x objective-c
// #cgo LDFLAGS: -framework AppKit
// #include "view_controller.h"
import "C"
import (
	"github.com/hsiafan/cocoa/coregraphics"
	"github.com/hsiafan/cocoa/foundation"
	"github.com/hsiafan/cocoa/objc"
	"unsafe"
)

type ViewController interface {
	Responder
	LoadView()
	CommitEditing() bool
	DiscardEditing()
	ViewDidLoad()
	ViewWillAppear()
	ViewDidAppear()
	ViewWillDisappear()
	ViewDidDisappear()
	UpdateViewConstraints()
	ViewWillLayout()
	ViewDidLayout()
	AddChildViewController(childViewController ViewController)
	InsertChildViewController_AtIndex(childViewController ViewController, index int)
	RemoveChildViewControllerAtIndex(index int)
	RemoveFromParentViewController()
	PreferredContentSizeDidChangeForViewController(viewController ViewController)
	DismissViewController(viewController ViewController)
	PresentViewControllerAsModalWindow(viewController ViewController)
	PresentViewControllerAsSheet(viewController ViewController)
	ViewWillTransitionToSize(newSize foundation.Size)
	RepresentedObject() objc.Object
	SetRepresentedObject(value objc.Object)
	NibBundle() foundation.Bundle
	NibName() NibName
	Title() string
	SetTitle(value string)
	Storyboard() Storyboard
	IsViewLoaded() bool
	PreferredContentSize() foundation.Size
	SetPreferredContentSize(value foundation.Size)
	ParentViewController() ViewController
	PresentingViewController() ViewController
	ExtensionContext() foundation.ExtensionContext
	PreferredScreenOrigin() foundation.Point
	SetPreferredScreenOrigin(value foundation.Point)
	PreferredMaximumSize() foundation.Size
	PreferredMinimumSize() foundation.Size
}

type NSViewController struct {
	NSResponder
}

func MakeViewController(ptr unsafe.Pointer) *NSViewController {
	if ptr == nil {
		return nil
	}
	return &NSViewController{
		NSResponder: *MakeResponder(ptr),
	}
}

func AllocViewController() *NSViewController {
	return MakeViewController(C.C_ViewController_Alloc())
}

func (n *NSViewController) InitWithNibName_Bundle(nibNameOrNil NibName, nibBundleOrNil foundation.Bundle) ViewController {
	result := C.C_NSViewController_InitWithNibName_Bundle(n.Ptr(), foundation.NewString(string(nibNameOrNil)).Ptr(), objc.ExtractPtr(nibBundleOrNil))
	return MakeViewController(result)
}

func (n *NSViewController) InitWithCoder(coder foundation.Coder) ViewController {
	result := C.C_NSViewController_InitWithCoder(n.Ptr(), objc.ExtractPtr(coder))
	return MakeViewController(result)
}

func (n *NSViewController) Init() ViewController {
	result := C.C_NSViewController_Init(n.Ptr())
	return MakeViewController(result)
}

func (n *NSViewController) LoadView() {
	C.C_NSViewController_LoadView(n.Ptr())
}

func (n *NSViewController) CommitEditing() bool {
	result := C.C_NSViewController_CommitEditing(n.Ptr())
	return bool(result)
}

func (n *NSViewController) DiscardEditing() {
	C.C_NSViewController_DiscardEditing(n.Ptr())
}

func (n *NSViewController) ViewDidLoad() {
	C.C_NSViewController_ViewDidLoad(n.Ptr())
}

func (n *NSViewController) ViewWillAppear() {
	C.C_NSViewController_ViewWillAppear(n.Ptr())
}

func (n *NSViewController) ViewDidAppear() {
	C.C_NSViewController_ViewDidAppear(n.Ptr())
}

func (n *NSViewController) ViewWillDisappear() {
	C.C_NSViewController_ViewWillDisappear(n.Ptr())
}

func (n *NSViewController) ViewDidDisappear() {
	C.C_NSViewController_ViewDidDisappear(n.Ptr())
}

func (n *NSViewController) UpdateViewConstraints() {
	C.C_NSViewController_UpdateViewConstraints(n.Ptr())
}

func (n *NSViewController) ViewWillLayout() {
	C.C_NSViewController_ViewWillLayout(n.Ptr())
}

func (n *NSViewController) ViewDidLayout() {
	C.C_NSViewController_ViewDidLayout(n.Ptr())
}

func (n *NSViewController) AddChildViewController(childViewController ViewController) {
	C.C_NSViewController_AddChildViewController(n.Ptr(), objc.ExtractPtr(childViewController))
}

func (n *NSViewController) InsertChildViewController_AtIndex(childViewController ViewController, index int) {
	C.C_NSViewController_InsertChildViewController_AtIndex(n.Ptr(), objc.ExtractPtr(childViewController), C.int(index))
}

func (n *NSViewController) RemoveChildViewControllerAtIndex(index int) {
	C.C_NSViewController_RemoveChildViewControllerAtIndex(n.Ptr(), C.int(index))
}

func (n *NSViewController) RemoveFromParentViewController() {
	C.C_NSViewController_RemoveFromParentViewController(n.Ptr())
}

func (n *NSViewController) PreferredContentSizeDidChangeForViewController(viewController ViewController) {
	C.C_NSViewController_PreferredContentSizeDidChangeForViewController(n.Ptr(), objc.ExtractPtr(viewController))
}

func (n *NSViewController) DismissViewController(viewController ViewController) {
	C.C_NSViewController_DismissViewController(n.Ptr(), objc.ExtractPtr(viewController))
}

func (n *NSViewController) PresentViewControllerAsModalWindow(viewController ViewController) {
	C.C_NSViewController_PresentViewControllerAsModalWindow(n.Ptr(), objc.ExtractPtr(viewController))
}

func (n *NSViewController) PresentViewControllerAsSheet(viewController ViewController) {
	C.C_NSViewController_PresentViewControllerAsSheet(n.Ptr(), objc.ExtractPtr(viewController))
}

func (n *NSViewController) ViewWillTransitionToSize(newSize foundation.Size) {
	C.C_NSViewController_ViewWillTransitionToSize(n.Ptr(), *(*C.CGSize)(coregraphics.ToCGSizePointer(coregraphics.Size(newSize))))
}

func (n *NSViewController) RepresentedObject() objc.Object {
	result := C.C_NSViewController_RepresentedObject(n.Ptr())
	return objc.MakeObject(result)
}

func (n *NSViewController) SetRepresentedObject(value objc.Object) {
	C.C_NSViewController_SetRepresentedObject(n.Ptr(), objc.ExtractPtr(value))
}

func (n *NSViewController) NibBundle() foundation.Bundle {
	result := C.C_NSViewController_NibBundle(n.Ptr())
	return foundation.MakeBundle(result)
}

func (n *NSViewController) NibName() NibName {
	result := C.C_NSViewController_NibName(n.Ptr())
	return NibName(foundation.MakeString(result).String())
}

func (n *NSViewController) Title() string {
	result := C.C_NSViewController_Title(n.Ptr())
	return foundation.MakeString(result).String()
}

func (n *NSViewController) SetTitle(value string) {
	C.C_NSViewController_SetTitle(n.Ptr(), foundation.NewString(value).Ptr())
}

func (n *NSViewController) Storyboard() Storyboard {
	result := C.C_NSViewController_Storyboard(n.Ptr())
	return MakeStoryboard(result)
}

func (n *NSViewController) IsViewLoaded() bool {
	result := C.C_NSViewController_IsViewLoaded(n.Ptr())
	return bool(result)
}

func (n *NSViewController) PreferredContentSize() foundation.Size {
	result := C.C_NSViewController_PreferredContentSize(n.Ptr())
	return foundation.Size(coregraphics.FromCGSizePointer(unsafe.Pointer(&result)))
}

func (n *NSViewController) SetPreferredContentSize(value foundation.Size) {
	C.C_NSViewController_SetPreferredContentSize(n.Ptr(), *(*C.CGSize)(coregraphics.ToCGSizePointer(coregraphics.Size(value))))
}

func (n *NSViewController) ParentViewController() ViewController {
	result := C.C_NSViewController_ParentViewController(n.Ptr())
	return MakeViewController(result)
}

func (n *NSViewController) PresentingViewController() ViewController {
	result := C.C_NSViewController_PresentingViewController(n.Ptr())
	return MakeViewController(result)
}

func (n *NSViewController) ExtensionContext() foundation.ExtensionContext {
	result := C.C_NSViewController_ExtensionContext(n.Ptr())
	return foundation.MakeExtensionContext(result)
}

func (n *NSViewController) PreferredScreenOrigin() foundation.Point {
	result := C.C_NSViewController_PreferredScreenOrigin(n.Ptr())
	return foundation.Point(coregraphics.FromCGPointPointer(unsafe.Pointer(&result)))
}

func (n *NSViewController) SetPreferredScreenOrigin(value foundation.Point) {
	C.C_NSViewController_SetPreferredScreenOrigin(n.Ptr(), *(*C.CGPoint)(coregraphics.ToCGPointPointer(coregraphics.Point(value))))
}

func (n *NSViewController) PreferredMaximumSize() foundation.Size {
	result := C.C_NSViewController_PreferredMaximumSize(n.Ptr())
	return foundation.Size(coregraphics.FromCGSizePointer(unsafe.Pointer(&result)))
}

func (n *NSViewController) PreferredMinimumSize() foundation.Size {
	result := C.C_NSViewController_PreferredMinimumSize(n.Ptr())
	return foundation.Size(coregraphics.FromCGSizePointer(unsafe.Pointer(&result)))
}
