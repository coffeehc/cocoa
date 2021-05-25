package appkit

// #include "view_controller_presentation_animator.h"
import "C"
import (
	"github.com/hsiafan/cocoa/objc"
	"unsafe"
)

type ViewControllerPresentationAnimator struct {
	AnimatePresentationOfViewController_FromViewController func(viewController ViewController, fromViewController ViewController) // required
	AnimateDismissalOfViewController_FromViewController    func(viewController ViewController, fromViewController ViewController) // required
}

func WrapViewControllerPresentationAnimator(delegate *ViewControllerPresentationAnimator) objc.Object {
	id := resources.NextId()
	resources.Store(id, delegate)
	ptr := C.WrapViewControllerPresentationAnimator(C.long(id))
	return objc.MakeObject(ptr)
}

//export viewControllerPresentationAnimator_AnimatePresentationOfViewController_FromViewController
func viewControllerPresentationAnimator_AnimatePresentationOfViewController_FromViewController(id int64, viewController unsafe.Pointer, fromViewController unsafe.Pointer) {
	delegate := resources.Get(id).(*ViewControllerPresentationAnimator)
	delegate.AnimatePresentationOfViewController_FromViewController(MakeViewController(viewController), MakeViewController(fromViewController))
}

//export viewControllerPresentationAnimator_AnimateDismissalOfViewController_FromViewController
func viewControllerPresentationAnimator_AnimateDismissalOfViewController_FromViewController(id int64, viewController unsafe.Pointer, fromViewController unsafe.Pointer) {
	delegate := resources.Get(id).(*ViewControllerPresentationAnimator)
	delegate.AnimateDismissalOfViewController_FromViewController(MakeViewController(viewController), MakeViewController(fromViewController))
}

//export ViewControllerPresentationAnimator_RespondsTo
func ViewControllerPresentationAnimator_RespondsTo(id int64, selectorPtr unsafe.Pointer) bool {
	sel := objc.MakeSelector(selectorPtr)
	selName := objc.Sel_GetName(sel)
	delegate := resources.Get(id).(*ViewControllerPresentationAnimator)
	switch selName {
	case "animatePresentationOfViewController:fromViewController:":
		return delegate.AnimatePresentationOfViewController_FromViewController != nil
	case "animateDismissalOfViewController:fromViewController:":
		return delegate.AnimateDismissalOfViewController_FromViewController != nil
	default:
		return false
	}
}

//export deleteViewControllerPresentationAnimator
func deleteViewControllerPresentationAnimator(id int64) {
	resources.Delete(id)
}
