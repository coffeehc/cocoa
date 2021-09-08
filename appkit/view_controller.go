package appkit

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
	DismissController(sender objc.Object)
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
	PresentViewController_Animator(viewController ViewController, animator objc.Object)
	DismissViewController(viewController ViewController)
	PresentViewController_AsPopoverRelativeToRect_OfView_PreferredEdge_Behavior(viewController ViewController, positioningRect foundation.Rect, positioningView View, preferredEdge foundation.RectEdge, behavior PopoverBehavior)
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
	ChildViewControllers() []ViewController
	SetChildViewControllers(value []ViewController)
	ParentViewController() ViewController
	PresentedViewControllers() []ViewController
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

func MakeViewController(ptr unsafe.Pointer) NSViewController {
	return NSViewController{
		NSResponder: MakeResponder(ptr),
	}
}

func (n NSViewController) InitWithNibName_Bundle(nibNameOrNil NibName, nibBundleOrNil foundation.Bundle) NSViewController {
	result_ := C.C_NSViewController_InitWithNibName_Bundle(n.Ptr(), foundation.NewString(string(nibNameOrNil)).Ptr(), objc.ExtractPtr(nibBundleOrNil))
	return MakeViewController(result_)
}

func (n NSViewController) InitWithCoder(coder foundation.Coder) NSViewController {
	result_ := C.C_NSViewController_InitWithCoder(n.Ptr(), objc.ExtractPtr(coder))
	return MakeViewController(result_)
}

func (n NSViewController) Init() NSViewController {
	result_ := C.C_NSViewController_Init(n.Ptr())
	return MakeViewController(result_)
}

func AllocViewController() NSViewController {
	result_ := C.C_NSViewController_AllocViewController()
	return MakeViewController(result_)
}

func NewViewController() NSViewController {
	result_ := C.C_NSViewController_NewViewController()
	return MakeViewController(result_)
}

func (n NSViewController) Autorelease() NSViewController {
	result_ := C.C_NSViewController_Autorelease(n.Ptr())
	return MakeViewController(result_)
}

func (n NSViewController) Retain() NSViewController {
	result_ := C.C_NSViewController_Retain(n.Ptr())
	return MakeViewController(result_)
}

func (n NSViewController) LoadView() {
	C.C_NSViewController_LoadView(n.Ptr())
}

func (n NSViewController) CommitEditing() bool {
	result_ := C.C_NSViewController_CommitEditing(n.Ptr())
	return bool(result_)
}

func (n NSViewController) DiscardEditing() {
	C.C_NSViewController_DiscardEditing(n.Ptr())
}

func (n NSViewController) DismissController(sender objc.Object) {
	C.C_NSViewController_DismissController(n.Ptr(), objc.ExtractPtr(sender))
}

func (n NSViewController) ViewDidLoad() {
	C.C_NSViewController_ViewDidLoad(n.Ptr())
}

func (n NSViewController) ViewWillAppear() {
	C.C_NSViewController_ViewWillAppear(n.Ptr())
}

func (n NSViewController) ViewDidAppear() {
	C.C_NSViewController_ViewDidAppear(n.Ptr())
}

func (n NSViewController) ViewWillDisappear() {
	C.C_NSViewController_ViewWillDisappear(n.Ptr())
}

func (n NSViewController) ViewDidDisappear() {
	C.C_NSViewController_ViewDidDisappear(n.Ptr())
}

func (n NSViewController) UpdateViewConstraints() {
	C.C_NSViewController_UpdateViewConstraints(n.Ptr())
}

func (n NSViewController) ViewWillLayout() {
	C.C_NSViewController_ViewWillLayout(n.Ptr())
}

func (n NSViewController) ViewDidLayout() {
	C.C_NSViewController_ViewDidLayout(n.Ptr())
}

func (n NSViewController) AddChildViewController(childViewController ViewController) {
	C.C_NSViewController_AddChildViewController(n.Ptr(), objc.ExtractPtr(childViewController))
}

func (n NSViewController) InsertChildViewController_AtIndex(childViewController ViewController, index int) {
	C.C_NSViewController_InsertChildViewController_AtIndex(n.Ptr(), objc.ExtractPtr(childViewController), C.int(index))
}

func (n NSViewController) RemoveChildViewControllerAtIndex(index int) {
	C.C_NSViewController_RemoveChildViewControllerAtIndex(n.Ptr(), C.int(index))
}

func (n NSViewController) RemoveFromParentViewController() {
	C.C_NSViewController_RemoveFromParentViewController(n.Ptr())
}

func (n NSViewController) PreferredContentSizeDidChangeForViewController(viewController ViewController) {
	C.C_NSViewController_PreferredContentSizeDidChangeForViewController(n.Ptr(), objc.ExtractPtr(viewController))
}

func (n NSViewController) PresentViewController_Animator(viewController ViewController, animator objc.Object) {
	C.C_NSViewController_PresentViewController_Animator(n.Ptr(), objc.ExtractPtr(viewController), objc.ExtractPtr(animator))
}

func (n NSViewController) DismissViewController(viewController ViewController) {
	C.C_NSViewController_DismissViewController(n.Ptr(), objc.ExtractPtr(viewController))
}

func (n NSViewController) PresentViewController_AsPopoverRelativeToRect_OfView_PreferredEdge_Behavior(viewController ViewController, positioningRect foundation.Rect, positioningView View, preferredEdge foundation.RectEdge, behavior PopoverBehavior) {
	C.C_NSViewController_PresentViewController_AsPopoverRelativeToRect_OfView_PreferredEdge_Behavior(n.Ptr(), objc.ExtractPtr(viewController), *(*C.CGRect)(coregraphics.ToCGRectPointer(coregraphics.Rect(positioningRect))), objc.ExtractPtr(positioningView), C.uint(uint(preferredEdge)), C.int(int(behavior)))
}

func (n NSViewController) PresentViewControllerAsModalWindow(viewController ViewController) {
	C.C_NSViewController_PresentViewControllerAsModalWindow(n.Ptr(), objc.ExtractPtr(viewController))
}

func (n NSViewController) PresentViewControllerAsSheet(viewController ViewController) {
	C.C_NSViewController_PresentViewControllerAsSheet(n.Ptr(), objc.ExtractPtr(viewController))
}

func (n NSViewController) ViewWillTransitionToSize(newSize foundation.Size) {
	C.C_NSViewController_ViewWillTransitionToSize(n.Ptr(), *(*C.CGSize)(coregraphics.ToCGSizePointer(coregraphics.Size(newSize))))
}

func (n NSViewController) RepresentedObject() objc.Object {
	result_ := C.C_NSViewController_RepresentedObject(n.Ptr())
	return objc.MakeObject(result_)
}

func (n NSViewController) SetRepresentedObject(value objc.Object) {
	C.C_NSViewController_SetRepresentedObject(n.Ptr(), objc.ExtractPtr(value))
}

func (n NSViewController) NibBundle() foundation.Bundle {
	result_ := C.C_NSViewController_NibBundle(n.Ptr())
	return foundation.MakeBundle(result_)
}

func (n NSViewController) NibName() NibName {
	result_ := C.C_NSViewController_NibName(n.Ptr())
	return NibName(foundation.MakeString(result_).String())
}

func (n NSViewController) Title() string {
	result_ := C.C_NSViewController_Title(n.Ptr())
	return foundation.MakeString(result_).String()
}

func (n NSViewController) SetTitle(value string) {
	C.C_NSViewController_SetTitle(n.Ptr(), foundation.NewString(value).Ptr())
}

func (n NSViewController) Storyboard() Storyboard {
	result_ := C.C_NSViewController_Storyboard(n.Ptr())
	return MakeStoryboard(result_)
}

func (n NSViewController) IsViewLoaded() bool {
	result_ := C.C_NSViewController_IsViewLoaded(n.Ptr())
	return bool(result_)
}

func (n NSViewController) PreferredContentSize() foundation.Size {
	result_ := C.C_NSViewController_PreferredContentSize(n.Ptr())
	return foundation.Size(coregraphics.FromCGSizePointer(unsafe.Pointer(&result_)))
}

func (n NSViewController) SetPreferredContentSize(value foundation.Size) {
	C.C_NSViewController_SetPreferredContentSize(n.Ptr(), *(*C.CGSize)(coregraphics.ToCGSizePointer(coregraphics.Size(value))))
}

func (n NSViewController) ChildViewControllers() []ViewController {
	result_ := C.C_NSViewController_ChildViewControllers(n.Ptr())
	if result_.len > 0 {
		defer C.free(result_.data)
	}
	result_Slice := unsafe.Slice((*unsafe.Pointer)(result_.data), int(result_.len))
	var goResult_ = make([]ViewController, len(result_Slice))
	for idx, r := range result_Slice {
		goResult_[idx] = MakeViewController(r)
	}
	return goResult_
}

func (n NSViewController) SetChildViewControllers(value []ViewController) {
	var cValue C.Array
	if len(value) > 0 {
		cValueData := make([]unsafe.Pointer, len(value))
		for idx, v := range value {
			cValueData[idx] = objc.ExtractPtr(v)
		}
		cValue.data = unsafe.Pointer(&cValueData[0])
		cValue.len = C.int(len(value))
	}
	C.C_NSViewController_SetChildViewControllers(n.Ptr(), cValue)
}

func (n NSViewController) ParentViewController() ViewController {
	result_ := C.C_NSViewController_ParentViewController(n.Ptr())
	return MakeViewController(result_)
}

func (n NSViewController) PresentedViewControllers() []ViewController {
	result_ := C.C_NSViewController_PresentedViewControllers(n.Ptr())
	if result_.len > 0 {
		defer C.free(result_.data)
	}
	result_Slice := unsafe.Slice((*unsafe.Pointer)(result_.data), int(result_.len))
	var goResult_ = make([]ViewController, len(result_Slice))
	for idx, r := range result_Slice {
		goResult_[idx] = MakeViewController(r)
	}
	return goResult_
}

func (n NSViewController) PresentingViewController() ViewController {
	result_ := C.C_NSViewController_PresentingViewController(n.Ptr())
	return MakeViewController(result_)
}

func (n NSViewController) ExtensionContext() foundation.ExtensionContext {
	result_ := C.C_NSViewController_ExtensionContext(n.Ptr())
	return foundation.MakeExtensionContext(result_)
}

func (n NSViewController) PreferredScreenOrigin() foundation.Point {
	result_ := C.C_NSViewController_PreferredScreenOrigin(n.Ptr())
	return foundation.Point(coregraphics.FromCGPointPointer(unsafe.Pointer(&result_)))
}

func (n NSViewController) SetPreferredScreenOrigin(value foundation.Point) {
	C.C_NSViewController_SetPreferredScreenOrigin(n.Ptr(), *(*C.CGPoint)(coregraphics.ToCGPointPointer(coregraphics.Point(value))))
}

func (n NSViewController) PreferredMaximumSize() foundation.Size {
	result_ := C.C_NSViewController_PreferredMaximumSize(n.Ptr())
	return foundation.Size(coregraphics.FromCGSizePointer(unsafe.Pointer(&result_)))
}

func (n NSViewController) PreferredMinimumSize() foundation.Size {
	result_ := C.C_NSViewController_PreferredMinimumSize(n.Ptr())
	return foundation.Size(coregraphics.FromCGSizePointer(unsafe.Pointer(&result_)))
}
