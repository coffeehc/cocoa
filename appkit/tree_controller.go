package appkit

// #include "tree_controller.h"
import "C"
import (
	"github.com/hsiafan/cocoa/foundation"
	"github.com/hsiafan/cocoa/objc"
	"unsafe"
)

type TreeController interface {
	ObjectController
	RearrangeObjects()
	SetSelectionIndexPath(indexPath foundation.IndexPath) bool
	SetSelectionIndexPaths(indexPaths []foundation.IndexPath) bool
	AddSelectionIndexPaths(indexPaths []foundation.IndexPath) bool
	RemoveSelectionIndexPaths(indexPaths []foundation.IndexPath) bool
	AddChild(sender objc.Object)
	Insert(sender objc.Object)
	InsertChild(sender objc.Object)
	InsertObject_AtArrangedObjectIndexPath(object objc.Object, indexPath foundation.IndexPath)
	InsertObjects_AtArrangedObjectIndexPaths(objects []objc.Object, indexPaths []foundation.IndexPath)
	RemoveObjectAtArrangedObjectIndexPath(indexPath foundation.IndexPath)
	RemoveObjectsAtArrangedObjectIndexPaths(indexPaths []foundation.IndexPath)
	MoveNode_ToIndexPath(node TreeNode, indexPath foundation.IndexPath)
	MoveNodes_ToIndexPath(nodes []TreeNode, startingIndexPath foundation.IndexPath)
	ChildrenKeyPathForNode(node TreeNode) string
	CountKeyPathForNode(node TreeNode) string
	LeafKeyPathForNode(node TreeNode) string
	SortDescriptors() []foundation.SortDescriptor
	SetSortDescriptors(value []foundation.SortDescriptor)
	ArrangedObjects() TreeNode
	SelectionIndexPath() foundation.IndexPath
	SelectionIndexPaths() []foundation.IndexPath
	SelectedNodes() []TreeNode
	SelectsInsertedObjects() bool
	SetSelectsInsertedObjects(value bool)
	AvoidsEmptySelection() bool
	SetAvoidsEmptySelection(value bool)
	PreservesSelection() bool
	SetPreservesSelection(value bool)
	AlwaysUsesMultipleValuesMarker() bool
	SetAlwaysUsesMultipleValuesMarker(value bool)
	CanAddChild() bool
	CanInsert() bool
	CanInsertChild() bool
	ChildrenKeyPath() string
	SetChildrenKeyPath(value string)
	CountKeyPath() string
	SetCountKeyPath(value string)
	LeafKeyPath() string
	SetLeafKeyPath(value string)
}

type NSTreeController struct {
	NSObjectController
}

func MakeTreeController(ptr unsafe.Pointer) NSTreeController {
	return NSTreeController{
		NSObjectController: MakeObjectController(ptr),
	}
}

func AllocTreeController() NSTreeController {
	return MakeTreeController(C.C_TreeController_Alloc())
}

func (n NSTreeController) InitWithContent(content objc.Object) TreeController {
	result_ := C.C_NSTreeController_InitWithContent(n.Ptr(), objc.ExtractPtr(content))
	return MakeTreeController(result_)
}

func (n NSTreeController) InitWithCoder(coder foundation.Coder) TreeController {
	result_ := C.C_NSTreeController_InitWithCoder(n.Ptr(), objc.ExtractPtr(coder))
	return MakeTreeController(result_)
}

func (n NSTreeController) Init() TreeController {
	result_ := C.C_NSTreeController_Init(n.Ptr())
	return MakeTreeController(result_)
}

func (n NSTreeController) RearrangeObjects() {
	C.C_NSTreeController_RearrangeObjects(n.Ptr())
}

func (n NSTreeController) SetSelectionIndexPath(indexPath foundation.IndexPath) bool {
	result_ := C.C_NSTreeController_SetSelectionIndexPath(n.Ptr(), objc.ExtractPtr(indexPath))
	return bool(result_)
}

func (n NSTreeController) SetSelectionIndexPaths(indexPaths []foundation.IndexPath) bool {
	cIndexPathsData := make([]unsafe.Pointer, len(indexPaths))
	for idx, v := range indexPaths {
		cIndexPathsData[idx] = objc.ExtractPtr(v)
	}
	cIndexPaths := C.Array{data: unsafe.Pointer(&cIndexPathsData[0]), len: C.int(len(indexPaths))}
	result_ := C.C_NSTreeController_SetSelectionIndexPaths(n.Ptr(), cIndexPaths)
	return bool(result_)
}

func (n NSTreeController) AddSelectionIndexPaths(indexPaths []foundation.IndexPath) bool {
	cIndexPathsData := make([]unsafe.Pointer, len(indexPaths))
	for idx, v := range indexPaths {
		cIndexPathsData[idx] = objc.ExtractPtr(v)
	}
	cIndexPaths := C.Array{data: unsafe.Pointer(&cIndexPathsData[0]), len: C.int(len(indexPaths))}
	result_ := C.C_NSTreeController_AddSelectionIndexPaths(n.Ptr(), cIndexPaths)
	return bool(result_)
}

func (n NSTreeController) RemoveSelectionIndexPaths(indexPaths []foundation.IndexPath) bool {
	cIndexPathsData := make([]unsafe.Pointer, len(indexPaths))
	for idx, v := range indexPaths {
		cIndexPathsData[idx] = objc.ExtractPtr(v)
	}
	cIndexPaths := C.Array{data: unsafe.Pointer(&cIndexPathsData[0]), len: C.int(len(indexPaths))}
	result_ := C.C_NSTreeController_RemoveSelectionIndexPaths(n.Ptr(), cIndexPaths)
	return bool(result_)
}

func (n NSTreeController) AddChild(sender objc.Object) {
	C.C_NSTreeController_AddChild(n.Ptr(), objc.ExtractPtr(sender))
}

func (n NSTreeController) Insert(sender objc.Object) {
	C.C_NSTreeController_Insert(n.Ptr(), objc.ExtractPtr(sender))
}

func (n NSTreeController) InsertChild(sender objc.Object) {
	C.C_NSTreeController_InsertChild(n.Ptr(), objc.ExtractPtr(sender))
}

func (n NSTreeController) InsertObject_AtArrangedObjectIndexPath(object objc.Object, indexPath foundation.IndexPath) {
	C.C_NSTreeController_InsertObject_AtArrangedObjectIndexPath(n.Ptr(), objc.ExtractPtr(object), objc.ExtractPtr(indexPath))
}

func (n NSTreeController) InsertObjects_AtArrangedObjectIndexPaths(objects []objc.Object, indexPaths []foundation.IndexPath) {
	cObjectsData := make([]unsafe.Pointer, len(objects))
	for idx, v := range objects {
		cObjectsData[idx] = objc.ExtractPtr(v)
	}
	cObjects := C.Array{data: unsafe.Pointer(&cObjectsData[0]), len: C.int(len(objects))}
	cIndexPathsData := make([]unsafe.Pointer, len(indexPaths))
	for idx, v := range indexPaths {
		cIndexPathsData[idx] = objc.ExtractPtr(v)
	}
	cIndexPaths := C.Array{data: unsafe.Pointer(&cIndexPathsData[0]), len: C.int(len(indexPaths))}
	C.C_NSTreeController_InsertObjects_AtArrangedObjectIndexPaths(n.Ptr(), cObjects, cIndexPaths)
}

func (n NSTreeController) RemoveObjectAtArrangedObjectIndexPath(indexPath foundation.IndexPath) {
	C.C_NSTreeController_RemoveObjectAtArrangedObjectIndexPath(n.Ptr(), objc.ExtractPtr(indexPath))
}

func (n NSTreeController) RemoveObjectsAtArrangedObjectIndexPaths(indexPaths []foundation.IndexPath) {
	cIndexPathsData := make([]unsafe.Pointer, len(indexPaths))
	for idx, v := range indexPaths {
		cIndexPathsData[idx] = objc.ExtractPtr(v)
	}
	cIndexPaths := C.Array{data: unsafe.Pointer(&cIndexPathsData[0]), len: C.int(len(indexPaths))}
	C.C_NSTreeController_RemoveObjectsAtArrangedObjectIndexPaths(n.Ptr(), cIndexPaths)
}

func (n NSTreeController) MoveNode_ToIndexPath(node TreeNode, indexPath foundation.IndexPath) {
	C.C_NSTreeController_MoveNode_ToIndexPath(n.Ptr(), objc.ExtractPtr(node), objc.ExtractPtr(indexPath))
}

func (n NSTreeController) MoveNodes_ToIndexPath(nodes []TreeNode, startingIndexPath foundation.IndexPath) {
	cNodesData := make([]unsafe.Pointer, len(nodes))
	for idx, v := range nodes {
		cNodesData[idx] = objc.ExtractPtr(v)
	}
	cNodes := C.Array{data: unsafe.Pointer(&cNodesData[0]), len: C.int(len(nodes))}
	C.C_NSTreeController_MoveNodes_ToIndexPath(n.Ptr(), cNodes, objc.ExtractPtr(startingIndexPath))
}

func (n NSTreeController) ChildrenKeyPathForNode(node TreeNode) string {
	result_ := C.C_NSTreeController_ChildrenKeyPathForNode(n.Ptr(), objc.ExtractPtr(node))
	return foundation.MakeString(result_).String()
}

func (n NSTreeController) CountKeyPathForNode(node TreeNode) string {
	result_ := C.C_NSTreeController_CountKeyPathForNode(n.Ptr(), objc.ExtractPtr(node))
	return foundation.MakeString(result_).String()
}

func (n NSTreeController) LeafKeyPathForNode(node TreeNode) string {
	result_ := C.C_NSTreeController_LeafKeyPathForNode(n.Ptr(), objc.ExtractPtr(node))
	return foundation.MakeString(result_).String()
}

func (n NSTreeController) SortDescriptors() []foundation.SortDescriptor {
	result_ := C.C_NSTreeController_SortDescriptors(n.Ptr())
	defer C.free(result_.data)
	result_Slice := (*[1 << 28]unsafe.Pointer)(unsafe.Pointer(result_.data))[:result_.len:result_.len]
	var goResult_ = make([]foundation.SortDescriptor, len(result_Slice))
	for idx, r := range result_Slice {
		goResult_[idx] = foundation.MakeSortDescriptor(r)
	}
	return goResult_
}

func (n NSTreeController) SetSortDescriptors(value []foundation.SortDescriptor) {
	cValueData := make([]unsafe.Pointer, len(value))
	for idx, v := range value {
		cValueData[idx] = objc.ExtractPtr(v)
	}
	cValue := C.Array{data: unsafe.Pointer(&cValueData[0]), len: C.int(len(value))}
	C.C_NSTreeController_SetSortDescriptors(n.Ptr(), cValue)
}

func (n NSTreeController) ArrangedObjects() TreeNode {
	result_ := C.C_NSTreeController_ArrangedObjects(n.Ptr())
	return MakeTreeNode(result_)
}

func (n NSTreeController) SelectionIndexPath() foundation.IndexPath {
	result_ := C.C_NSTreeController_SelectionIndexPath(n.Ptr())
	return foundation.MakeIndexPath(result_)
}

func (n NSTreeController) SelectionIndexPaths() []foundation.IndexPath {
	result_ := C.C_NSTreeController_SelectionIndexPaths(n.Ptr())
	defer C.free(result_.data)
	result_Slice := (*[1 << 28]unsafe.Pointer)(unsafe.Pointer(result_.data))[:result_.len:result_.len]
	var goResult_ = make([]foundation.IndexPath, len(result_Slice))
	for idx, r := range result_Slice {
		goResult_[idx] = foundation.MakeIndexPath(r)
	}
	return goResult_
}

func (n NSTreeController) SelectedNodes() []TreeNode {
	result_ := C.C_NSTreeController_SelectedNodes(n.Ptr())
	defer C.free(result_.data)
	result_Slice := (*[1 << 28]unsafe.Pointer)(unsafe.Pointer(result_.data))[:result_.len:result_.len]
	var goResult_ = make([]TreeNode, len(result_Slice))
	for idx, r := range result_Slice {
		goResult_[idx] = MakeTreeNode(r)
	}
	return goResult_
}

func (n NSTreeController) SelectsInsertedObjects() bool {
	result_ := C.C_NSTreeController_SelectsInsertedObjects(n.Ptr())
	return bool(result_)
}

func (n NSTreeController) SetSelectsInsertedObjects(value bool) {
	C.C_NSTreeController_SetSelectsInsertedObjects(n.Ptr(), C.bool(value))
}

func (n NSTreeController) AvoidsEmptySelection() bool {
	result_ := C.C_NSTreeController_AvoidsEmptySelection(n.Ptr())
	return bool(result_)
}

func (n NSTreeController) SetAvoidsEmptySelection(value bool) {
	C.C_NSTreeController_SetAvoidsEmptySelection(n.Ptr(), C.bool(value))
}

func (n NSTreeController) PreservesSelection() bool {
	result_ := C.C_NSTreeController_PreservesSelection(n.Ptr())
	return bool(result_)
}

func (n NSTreeController) SetPreservesSelection(value bool) {
	C.C_NSTreeController_SetPreservesSelection(n.Ptr(), C.bool(value))
}

func (n NSTreeController) AlwaysUsesMultipleValuesMarker() bool {
	result_ := C.C_NSTreeController_AlwaysUsesMultipleValuesMarker(n.Ptr())
	return bool(result_)
}

func (n NSTreeController) SetAlwaysUsesMultipleValuesMarker(value bool) {
	C.C_NSTreeController_SetAlwaysUsesMultipleValuesMarker(n.Ptr(), C.bool(value))
}

func (n NSTreeController) CanAddChild() bool {
	result_ := C.C_NSTreeController_CanAddChild(n.Ptr())
	return bool(result_)
}

func (n NSTreeController) CanInsert() bool {
	result_ := C.C_NSTreeController_CanInsert(n.Ptr())
	return bool(result_)
}

func (n NSTreeController) CanInsertChild() bool {
	result_ := C.C_NSTreeController_CanInsertChild(n.Ptr())
	return bool(result_)
}

func (n NSTreeController) ChildrenKeyPath() string {
	result_ := C.C_NSTreeController_ChildrenKeyPath(n.Ptr())
	return foundation.MakeString(result_).String()
}

func (n NSTreeController) SetChildrenKeyPath(value string) {
	C.C_NSTreeController_SetChildrenKeyPath(n.Ptr(), foundation.NewString(value).Ptr())
}

func (n NSTreeController) CountKeyPath() string {
	result_ := C.C_NSTreeController_CountKeyPath(n.Ptr())
	return foundation.MakeString(result_).String()
}

func (n NSTreeController) SetCountKeyPath(value string) {
	C.C_NSTreeController_SetCountKeyPath(n.Ptr(), foundation.NewString(value).Ptr())
}

func (n NSTreeController) LeafKeyPath() string {
	result_ := C.C_NSTreeController_LeafKeyPath(n.Ptr())
	return foundation.MakeString(result_).String()
}

func (n NSTreeController) SetLeafKeyPath(value string) {
	C.C_NSTreeController_SetLeafKeyPath(n.Ptr(), foundation.NewString(value).Ptr())
}
