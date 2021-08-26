package appkit

// #include "view_controller_presentation_animator.h"
import "C"
import (
	"github.com/hsiafan/cocoa/objc"
	"runtime/cgo"
	"unsafe"
)

type ViewControllerPresentationAnimator struct {
	AnimatePresentationOfViewController_FromViewController func(viewController ViewController, fromViewController ViewController) // required
	AnimateDismissalOfViewController_FromViewController    func(viewController ViewController, fromViewController ViewController) // required
}

func (delegate *ViewControllerPresentationAnimator) ToObjc() objc.Object {
	h := cgo.NewHandle(delegate)
	ptr := C.WrapViewControllerPresentationAnimator(C.uintptr_t(h))
	return objc.MakeObject(ptr)
}

//export viewControllerPresentationAnimator_AnimatePresentationOfViewController_FromViewController
func viewControllerPresentationAnimator_AnimatePresentationOfViewController_FromViewController(hp C.uintptr_t, viewController unsafe.Pointer, fromViewController unsafe.Pointer) {
	delegate := cgo.Handle(hp).Value().(*ViewControllerPresentationAnimator)
	delegate.AnimatePresentationOfViewController_FromViewController(MakeViewController(viewController), MakeViewController(fromViewController))
}

//export viewControllerPresentationAnimator_AnimateDismissalOfViewController_FromViewController
func viewControllerPresentationAnimator_AnimateDismissalOfViewController_FromViewController(hp C.uintptr_t, viewController unsafe.Pointer, fromViewController unsafe.Pointer) {
	delegate := cgo.Handle(hp).Value().(*ViewControllerPresentationAnimator)
	delegate.AnimateDismissalOfViewController_FromViewController(MakeViewController(viewController), MakeViewController(fromViewController))
}

//export ViewControllerPresentationAnimator_RespondsTo
func ViewControllerPresentationAnimator_RespondsTo(hp C.uintptr_t, selectorPtr unsafe.Pointer) bool {
	sel := objc.Selector(selectorPtr)
	selName := objc.Sel_GetName(sel)
	delegate := cgo.Handle(hp).Value().(*ViewControllerPresentationAnimator)
	switch selName {
	case "animatePresentationOfViewController:fromViewController:":
		return delegate.AnimatePresentationOfViewController_FromViewController != nil
	case "animateDismissalOfViewController:fromViewController:":
		return delegate.AnimateDismissalOfViewController_FromViewController != nil
	default:
		return false
	}
}
