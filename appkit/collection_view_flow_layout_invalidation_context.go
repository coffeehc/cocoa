package appkit

// #include "collection_view_flow_layout_invalidation_context.h"
import "C"
import (
	"unsafe"
)

type CollectionViewFlowLayoutInvalidationContext interface {
	CollectionViewLayoutInvalidationContext
	InvalidateFlowLayoutAttributes() bool
	SetInvalidateFlowLayoutAttributes(value bool)
	InvalidateFlowLayoutDelegateMetrics() bool
	SetInvalidateFlowLayoutDelegateMetrics(value bool)
}

type NSCollectionViewFlowLayoutInvalidationContext struct {
	NSCollectionViewLayoutInvalidationContext
}

func MakeCollectionViewFlowLayoutInvalidationContext(ptr unsafe.Pointer) *NSCollectionViewFlowLayoutInvalidationContext {
	if ptr == nil {
		return nil
	}
	return &NSCollectionViewFlowLayoutInvalidationContext{
		NSCollectionViewLayoutInvalidationContext: *MakeCollectionViewLayoutInvalidationContext(ptr),
	}
}

func AllocCollectionViewFlowLayoutInvalidationContext() *NSCollectionViewFlowLayoutInvalidationContext {
	return MakeCollectionViewFlowLayoutInvalidationContext(C.C_CollectionViewFlowLayoutInvalidationContext_Alloc())
}

func (n *NSCollectionViewFlowLayoutInvalidationContext) Init() CollectionViewFlowLayoutInvalidationContext {
	result_ := C.C_NSCollectionViewFlowLayoutInvalidationContext_Init(n.Ptr())
	return MakeCollectionViewFlowLayoutInvalidationContext(result_)
}

func (n *NSCollectionViewFlowLayoutInvalidationContext) InvalidateFlowLayoutAttributes() bool {
	result_ := C.C_NSCollectionViewFlowLayoutInvalidationContext_InvalidateFlowLayoutAttributes(n.Ptr())
	return bool(result_)
}

func (n *NSCollectionViewFlowLayoutInvalidationContext) SetInvalidateFlowLayoutAttributes(value bool) {
	C.C_NSCollectionViewFlowLayoutInvalidationContext_SetInvalidateFlowLayoutAttributes(n.Ptr(), C.bool(value))
}

func (n *NSCollectionViewFlowLayoutInvalidationContext) InvalidateFlowLayoutDelegateMetrics() bool {
	result_ := C.C_NSCollectionViewFlowLayoutInvalidationContext_InvalidateFlowLayoutDelegateMetrics(n.Ptr())
	return bool(result_)
}

func (n *NSCollectionViewFlowLayoutInvalidationContext) SetInvalidateFlowLayoutDelegateMetrics(value bool) {
	C.C_NSCollectionViewFlowLayoutInvalidationContext_SetInvalidateFlowLayoutDelegateMetrics(n.Ptr(), C.bool(value))
}
