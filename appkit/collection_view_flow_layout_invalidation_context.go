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

func MakeCollectionViewFlowLayoutInvalidationContext(ptr unsafe.Pointer) NSCollectionViewFlowLayoutInvalidationContext {
	return NSCollectionViewFlowLayoutInvalidationContext{
		NSCollectionViewLayoutInvalidationContext: MakeCollectionViewLayoutInvalidationContext(ptr),
	}
}

func AllocCollectionViewFlowLayoutInvalidationContext() NSCollectionViewFlowLayoutInvalidationContext {
	result_ := C.C_NSCollectionViewFlowLayoutInvalidationContext_AllocCollectionViewFlowLayoutInvalidationContext()
	return MakeCollectionViewFlowLayoutInvalidationContext(result_)
}

func (n NSCollectionViewFlowLayoutInvalidationContext) Init() NSCollectionViewFlowLayoutInvalidationContext {
	result_ := C.C_NSCollectionViewFlowLayoutInvalidationContext_Init(n.Ptr())
	return MakeCollectionViewFlowLayoutInvalidationContext(result_)
}

func NewCollectionViewFlowLayoutInvalidationContext() NSCollectionViewFlowLayoutInvalidationContext {
	result_ := C.C_NSCollectionViewFlowLayoutInvalidationContext_NewCollectionViewFlowLayoutInvalidationContext()
	return MakeCollectionViewFlowLayoutInvalidationContext(result_)
}

func (n NSCollectionViewFlowLayoutInvalidationContext) Autorelease() NSCollectionViewFlowLayoutInvalidationContext {
	result_ := C.C_NSCollectionViewFlowLayoutInvalidationContext_Autorelease(n.Ptr())
	return MakeCollectionViewFlowLayoutInvalidationContext(result_)
}

func (n NSCollectionViewFlowLayoutInvalidationContext) Retain() NSCollectionViewFlowLayoutInvalidationContext {
	result_ := C.C_NSCollectionViewFlowLayoutInvalidationContext_Retain(n.Ptr())
	return MakeCollectionViewFlowLayoutInvalidationContext(result_)
}

func (n NSCollectionViewFlowLayoutInvalidationContext) InvalidateFlowLayoutAttributes() bool {
	result_ := C.C_NSCollectionViewFlowLayoutInvalidationContext_InvalidateFlowLayoutAttributes(n.Ptr())
	return bool(result_)
}

func (n NSCollectionViewFlowLayoutInvalidationContext) SetInvalidateFlowLayoutAttributes(value bool) {
	C.C_NSCollectionViewFlowLayoutInvalidationContext_SetInvalidateFlowLayoutAttributes(n.Ptr(), C.bool(value))
}

func (n NSCollectionViewFlowLayoutInvalidationContext) InvalidateFlowLayoutDelegateMetrics() bool {
	result_ := C.C_NSCollectionViewFlowLayoutInvalidationContext_InvalidateFlowLayoutDelegateMetrics(n.Ptr())
	return bool(result_)
}

func (n NSCollectionViewFlowLayoutInvalidationContext) SetInvalidateFlowLayoutDelegateMetrics(value bool) {
	C.C_NSCollectionViewFlowLayoutInvalidationContext_SetInvalidateFlowLayoutDelegateMetrics(n.Ptr(), C.bool(value))
}
