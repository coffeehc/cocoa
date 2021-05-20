package appkit

// #include "object_controller.h"
import "C"
import (
	"github.com/hsiafan/cocoa/foundation"
	"github.com/hsiafan/cocoa/objc"
	"unsafe"
)

type ObjectController interface {
	Controller
	PrepareContent()
	NewObject() objc.Object
	AddObject(object objc.Object)
	RemoveObject(object objc.Object)
	Add(sender objc.Object)
	Remove(sender objc.Object)
	Fetch(sender objc.Object)
	ValidateUserInterfaceItem(item objc.Object) bool
	Content() objc.Object
	SetContent(value objc.Object)
	AutomaticallyPreparesContent() bool
	SetAutomaticallyPreparesContent(value bool)
	CanAdd() bool
	CanRemove() bool
	IsEditable() bool
	SetEditable(value bool)
	EntityName() string
	SetEntityName(value string)
	UsesLazyFetching() bool
	SetUsesLazyFetching(value bool)
	FetchPredicate() foundation.Predicate
	SetFetchPredicate(value foundation.Predicate)
	SelectedObjects() []objc.Object
	Selection() objc.Object
}

type NSObjectController struct {
	NSController
}

func MakeObjectController(ptr unsafe.Pointer) *NSObjectController {
	if ptr == nil {
		return nil
	}
	return &NSObjectController{
		NSController: *MakeController(ptr),
	}
}

func AllocObjectController() *NSObjectController {
	return MakeObjectController(C.C_ObjectController_Alloc())
}

func (n *NSObjectController) InitWithContent(content objc.Object) ObjectController {
	result_ := C.C_NSObjectController_InitWithContent(n.Ptr(), objc.ExtractPtr(content))
	return MakeObjectController(result_)
}

func (n *NSObjectController) InitWithCoder(coder foundation.Coder) ObjectController {
	result_ := C.C_NSObjectController_InitWithCoder(n.Ptr(), objc.ExtractPtr(coder))
	return MakeObjectController(result_)
}

func (n *NSObjectController) Init() ObjectController {
	result_ := C.C_NSObjectController_Init(n.Ptr())
	return MakeObjectController(result_)
}

func (n *NSObjectController) PrepareContent() {
	C.C_NSObjectController_PrepareContent(n.Ptr())
}

func (n *NSObjectController) NewObject() objc.Object {
	result_ := C.C_NSObjectController_NewObject(n.Ptr())
	return objc.MakeObject(result_)
}

func (n *NSObjectController) AddObject(object objc.Object) {
	C.C_NSObjectController_AddObject(n.Ptr(), objc.ExtractPtr(object))
}

func (n *NSObjectController) RemoveObject(object objc.Object) {
	C.C_NSObjectController_RemoveObject(n.Ptr(), objc.ExtractPtr(object))
}

func (n *NSObjectController) Add(sender objc.Object) {
	C.C_NSObjectController_Add(n.Ptr(), objc.ExtractPtr(sender))
}

func (n *NSObjectController) Remove(sender objc.Object) {
	C.C_NSObjectController_Remove(n.Ptr(), objc.ExtractPtr(sender))
}

func (n *NSObjectController) Fetch(sender objc.Object) {
	C.C_NSObjectController_Fetch(n.Ptr(), objc.ExtractPtr(sender))
}

func (n *NSObjectController) ValidateUserInterfaceItem(item objc.Object) bool {
	result_ := C.C_NSObjectController_ValidateUserInterfaceItem(n.Ptr(), objc.ExtractPtr(item))
	return bool(result_)
}

func (n *NSObjectController) Content() objc.Object {
	result_ := C.C_NSObjectController_Content(n.Ptr())
	return objc.MakeObject(result_)
}

func (n *NSObjectController) SetContent(value objc.Object) {
	C.C_NSObjectController_SetContent(n.Ptr(), objc.ExtractPtr(value))
}

func (n *NSObjectController) AutomaticallyPreparesContent() bool {
	result_ := C.C_NSObjectController_AutomaticallyPreparesContent(n.Ptr())
	return bool(result_)
}

func (n *NSObjectController) SetAutomaticallyPreparesContent(value bool) {
	C.C_NSObjectController_SetAutomaticallyPreparesContent(n.Ptr(), C.bool(value))
}

func (n *NSObjectController) CanAdd() bool {
	result_ := C.C_NSObjectController_CanAdd(n.Ptr())
	return bool(result_)
}

func (n *NSObjectController) CanRemove() bool {
	result_ := C.C_NSObjectController_CanRemove(n.Ptr())
	return bool(result_)
}

func (n *NSObjectController) IsEditable() bool {
	result_ := C.C_NSObjectController_IsEditable(n.Ptr())
	return bool(result_)
}

func (n *NSObjectController) SetEditable(value bool) {
	C.C_NSObjectController_SetEditable(n.Ptr(), C.bool(value))
}

func (n *NSObjectController) EntityName() string {
	result_ := C.C_NSObjectController_EntityName(n.Ptr())
	return foundation.MakeString(result_).String()
}

func (n *NSObjectController) SetEntityName(value string) {
	C.C_NSObjectController_SetEntityName(n.Ptr(), foundation.NewString(value).Ptr())
}

func (n *NSObjectController) UsesLazyFetching() bool {
	result_ := C.C_NSObjectController_UsesLazyFetching(n.Ptr())
	return bool(result_)
}

func (n *NSObjectController) SetUsesLazyFetching(value bool) {
	C.C_NSObjectController_SetUsesLazyFetching(n.Ptr(), C.bool(value))
}

func (n *NSObjectController) FetchPredicate() foundation.Predicate {
	result_ := C.C_NSObjectController_FetchPredicate(n.Ptr())
	return foundation.MakePredicate(result_)
}

func (n *NSObjectController) SetFetchPredicate(value foundation.Predicate) {
	C.C_NSObjectController_SetFetchPredicate(n.Ptr(), objc.ExtractPtr(value))
}

func (n *NSObjectController) SelectedObjects() []objc.Object {
	result_ := C.C_NSObjectController_SelectedObjects(n.Ptr())
	defer C.free(result_.data)
	result_Slice := (*[1 << 28]unsafe.Pointer)(unsafe.Pointer(result_.data))[:result_.len:result_.len]
	var goResult_ = make([]objc.Object, len(result_Slice))
	for idx, r := range result_Slice {
		goResult_[idx] = objc.MakeObject(r)
	}
	return goResult_
}

func (n *NSObjectController) Selection() objc.Object {
	result_ := C.C_NSObjectController_Selection(n.Ptr())
	return objc.MakeObject(result_)
}
