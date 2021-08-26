package appkit

// #include "tree_node.h"
import "C"
import (
	"github.com/hsiafan/cocoa/foundation"
	"github.com/hsiafan/cocoa/objc"
	"unsafe"
)

type TreeNode interface {
	objc.Object
	DescendantNodeAtIndexPath(indexPath foundation.IndexPath) TreeNode
	SortWithSortDescriptors_Recursively(sortDescriptors []foundation.SortDescriptor, recursively bool)
	RepresentedObject() objc.Object
	IndexPath() foundation.IndexPath
	IsLeaf() bool
	ChildNodes() []TreeNode
	ParentNode() TreeNode
}

type NSTreeNode struct {
	objc.NSObject
}

func MakeTreeNode(ptr unsafe.Pointer) NSTreeNode {
	return NSTreeNode{
		NSObject: objc.MakeObject(ptr),
	}
}

func AllocTreeNode() NSTreeNode {
	return MakeTreeNode(C.C_TreeNode_Alloc())
}

func (n NSTreeNode) InitWithRepresentedObject(modelObject objc.Object) TreeNode {
	result_ := C.C_NSTreeNode_InitWithRepresentedObject(n.Ptr(), objc.ExtractPtr(modelObject))
	return MakeTreeNode(result_)
}

func (n NSTreeNode) Init() TreeNode {
	result_ := C.C_NSTreeNode_Init(n.Ptr())
	return MakeTreeNode(result_)
}

func TreeNodeWithRepresentedObject(modelObject objc.Object) TreeNode {
	result_ := C.C_NSTreeNode_TreeNodeWithRepresentedObject(objc.ExtractPtr(modelObject))
	return MakeTreeNode(result_)
}

func (n NSTreeNode) DescendantNodeAtIndexPath(indexPath foundation.IndexPath) TreeNode {
	result_ := C.C_NSTreeNode_DescendantNodeAtIndexPath(n.Ptr(), objc.ExtractPtr(indexPath))
	return MakeTreeNode(result_)
}

func (n NSTreeNode) SortWithSortDescriptors_Recursively(sortDescriptors []foundation.SortDescriptor, recursively bool) {
	var cSortDescriptors C.Array
	if len(sortDescriptors) > 0 {
		cSortDescriptorsData := make([]unsafe.Pointer, len(sortDescriptors))
		for idx, v := range sortDescriptors {
			cSortDescriptorsData[idx] = objc.ExtractPtr(v)
		}
		cSortDescriptors.data = unsafe.Pointer(&cSortDescriptorsData[0])
		cSortDescriptors.len = C.int(len(sortDescriptors))
	}
	C.C_NSTreeNode_SortWithSortDescriptors_Recursively(n.Ptr(), cSortDescriptors, C.bool(recursively))
}

func (n NSTreeNode) RepresentedObject() objc.Object {
	result_ := C.C_NSTreeNode_RepresentedObject(n.Ptr())
	return objc.MakeObject(result_)
}

func (n NSTreeNode) IndexPath() foundation.IndexPath {
	result_ := C.C_NSTreeNode_IndexPath(n.Ptr())
	return foundation.MakeIndexPath(result_)
}

func (n NSTreeNode) IsLeaf() bool {
	result_ := C.C_NSTreeNode_IsLeaf(n.Ptr())
	return bool(result_)
}

func (n NSTreeNode) ChildNodes() []TreeNode {
	result_ := C.C_NSTreeNode_ChildNodes(n.Ptr())
	if result_.len > 0 {
		defer C.free(result_.data)
	}
	result_Slice := (*[1 << 28]unsafe.Pointer)(unsafe.Pointer(result_.data))[:result_.len:result_.len]
	var goResult_ = make([]TreeNode, len(result_Slice))
	for idx, r := range result_Slice {
		goResult_[idx] = MakeTreeNode(r)
	}
	return goResult_
}

func (n NSTreeNode) ParentNode() TreeNode {
	result_ := C.C_NSTreeNode_ParentNode(n.Ptr())
	return MakeTreeNode(result_)
}
