package appkit

// #include "collection_view_transition_layout.h"
import "C"
import (
	"github.com/hsiafan/cocoa/coregraphics"
	"github.com/hsiafan/cocoa/foundation"
	"github.com/hsiafan/cocoa/objc"
	"unsafe"
)

type CollectionViewTransitionLayout interface {
	CollectionViewLayout
	UpdateValue_ForAnimatedKey(value coregraphics.Float, key CollectionViewTransitionLayoutAnimatedKey)
	ValueForAnimatedKey(key CollectionViewTransitionLayoutAnimatedKey) coregraphics.Float
	TransitionProgress() coregraphics.Float
	SetTransitionProgress(value coregraphics.Float)
	CurrentLayout() CollectionViewLayout
	NextLayout() CollectionViewLayout
}

type NSCollectionViewTransitionLayout struct {
	NSCollectionViewLayout
}

func MakeCollectionViewTransitionLayout(ptr unsafe.Pointer) NSCollectionViewTransitionLayout {
	return NSCollectionViewTransitionLayout{
		NSCollectionViewLayout: MakeCollectionViewLayout(ptr),
	}
}

func (n NSCollectionViewTransitionLayout) InitWithCurrentLayout_NextLayout(currentLayout CollectionViewLayout, newLayout CollectionViewLayout) NSCollectionViewTransitionLayout {
	result_ := C.C_NSCollectionViewTransitionLayout_InitWithCurrentLayout_NextLayout(n.Ptr(), objc.ExtractPtr(currentLayout), objc.ExtractPtr(newLayout))
	return MakeCollectionViewTransitionLayout(result_)
}

func AllocCollectionViewTransitionLayout() NSCollectionViewTransitionLayout {
	result_ := C.C_NSCollectionViewTransitionLayout_AllocCollectionViewTransitionLayout()
	return MakeCollectionViewTransitionLayout(result_)
}

func (n NSCollectionViewTransitionLayout) Init() NSCollectionViewTransitionLayout {
	result_ := C.C_NSCollectionViewTransitionLayout_Init(n.Ptr())
	return MakeCollectionViewTransitionLayout(result_)
}

func NewCollectionViewTransitionLayout() NSCollectionViewTransitionLayout {
	result_ := C.C_NSCollectionViewTransitionLayout_NewCollectionViewTransitionLayout()
	return MakeCollectionViewTransitionLayout(result_)
}

func (n NSCollectionViewTransitionLayout) Autorelease() NSCollectionViewTransitionLayout {
	result_ := C.C_NSCollectionViewTransitionLayout_Autorelease(n.Ptr())
	return MakeCollectionViewTransitionLayout(result_)
}

func (n NSCollectionViewTransitionLayout) Retain() NSCollectionViewTransitionLayout {
	result_ := C.C_NSCollectionViewTransitionLayout_Retain(n.Ptr())
	return MakeCollectionViewTransitionLayout(result_)
}

func (n NSCollectionViewTransitionLayout) UpdateValue_ForAnimatedKey(value coregraphics.Float, key CollectionViewTransitionLayoutAnimatedKey) {
	C.C_NSCollectionViewTransitionLayout_UpdateValue_ForAnimatedKey(n.Ptr(), C.double(float64(value)), foundation.NewString(string(key)).Ptr())
}

func (n NSCollectionViewTransitionLayout) ValueForAnimatedKey(key CollectionViewTransitionLayoutAnimatedKey) coregraphics.Float {
	result_ := C.C_NSCollectionViewTransitionLayout_ValueForAnimatedKey(n.Ptr(), foundation.NewString(string(key)).Ptr())
	return coregraphics.Float(float64(result_))
}

func (n NSCollectionViewTransitionLayout) TransitionProgress() coregraphics.Float {
	result_ := C.C_NSCollectionViewTransitionLayout_TransitionProgress(n.Ptr())
	return coregraphics.Float(float64(result_))
}

func (n NSCollectionViewTransitionLayout) SetTransitionProgress(value coregraphics.Float) {
	C.C_NSCollectionViewTransitionLayout_SetTransitionProgress(n.Ptr(), C.double(float64(value)))
}

func (n NSCollectionViewTransitionLayout) CurrentLayout() CollectionViewLayout {
	result_ := C.C_NSCollectionViewTransitionLayout_CurrentLayout(n.Ptr())
	return MakeCollectionViewLayout(result_)
}

func (n NSCollectionViewTransitionLayout) NextLayout() CollectionViewLayout {
	result_ := C.C_NSCollectionViewTransitionLayout_NextLayout(n.Ptr())
	return MakeCollectionViewLayout(result_)
}
