package appkit

// #include "collection_view_grid_layout.h"
import "C"
import (
	"github.com/hsiafan/cocoa/coregraphics"
	"github.com/hsiafan/cocoa/foundation"
	"github.com/hsiafan/cocoa/objc"
	"unsafe"
)

type CollectionViewGridLayout interface {
	CollectionViewLayout
	MaximumNumberOfRows() uint
	SetMaximumNumberOfRows(value uint)
	MaximumNumberOfColumns() uint
	SetMaximumNumberOfColumns(value uint)
	MinimumItemSize() foundation.Size
	SetMinimumItemSize(value foundation.Size)
	MaximumItemSize() foundation.Size
	SetMaximumItemSize(value foundation.Size)
	MinimumInteritemSpacing() coregraphics.Float
	SetMinimumInteritemSpacing(value coregraphics.Float)
	MinimumLineSpacing() coregraphics.Float
	SetMinimumLineSpacing(value coregraphics.Float)
	Margins() foundation.EdgeInsets
	SetMargins(value foundation.EdgeInsets)
	BackgroundColors() []Color
	SetBackgroundColors(value []Color)
}

type NSCollectionViewGridLayout struct {
	NSCollectionViewLayout
}

func MakeCollectionViewGridLayout(ptr unsafe.Pointer) NSCollectionViewGridLayout {
	return NSCollectionViewGridLayout{
		NSCollectionViewLayout: MakeCollectionViewLayout(ptr),
	}
}

func AllocCollectionViewGridLayout() NSCollectionViewGridLayout {
	return MakeCollectionViewGridLayout(C.C_CollectionViewGridLayout_Alloc())
}

func (n NSCollectionViewGridLayout) Init() CollectionViewGridLayout {
	result_ := C.C_NSCollectionViewGridLayout_Init(n.Ptr())
	return MakeCollectionViewGridLayout(result_)
}

func (n NSCollectionViewGridLayout) MaximumNumberOfRows() uint {
	result_ := C.C_NSCollectionViewGridLayout_MaximumNumberOfRows(n.Ptr())
	return uint(result_)
}

func (n NSCollectionViewGridLayout) SetMaximumNumberOfRows(value uint) {
	C.C_NSCollectionViewGridLayout_SetMaximumNumberOfRows(n.Ptr(), C.uint(value))
}

func (n NSCollectionViewGridLayout) MaximumNumberOfColumns() uint {
	result_ := C.C_NSCollectionViewGridLayout_MaximumNumberOfColumns(n.Ptr())
	return uint(result_)
}

func (n NSCollectionViewGridLayout) SetMaximumNumberOfColumns(value uint) {
	C.C_NSCollectionViewGridLayout_SetMaximumNumberOfColumns(n.Ptr(), C.uint(value))
}

func (n NSCollectionViewGridLayout) MinimumItemSize() foundation.Size {
	result_ := C.C_NSCollectionViewGridLayout_MinimumItemSize(n.Ptr())
	return foundation.Size(coregraphics.FromCGSizePointer(unsafe.Pointer(&result_)))
}

func (n NSCollectionViewGridLayout) SetMinimumItemSize(value foundation.Size) {
	C.C_NSCollectionViewGridLayout_SetMinimumItemSize(n.Ptr(), *(*C.CGSize)(coregraphics.ToCGSizePointer(coregraphics.Size(value))))
}

func (n NSCollectionViewGridLayout) MaximumItemSize() foundation.Size {
	result_ := C.C_NSCollectionViewGridLayout_MaximumItemSize(n.Ptr())
	return foundation.Size(coregraphics.FromCGSizePointer(unsafe.Pointer(&result_)))
}

func (n NSCollectionViewGridLayout) SetMaximumItemSize(value foundation.Size) {
	C.C_NSCollectionViewGridLayout_SetMaximumItemSize(n.Ptr(), *(*C.CGSize)(coregraphics.ToCGSizePointer(coregraphics.Size(value))))
}

func (n NSCollectionViewGridLayout) MinimumInteritemSpacing() coregraphics.Float {
	result_ := C.C_NSCollectionViewGridLayout_MinimumInteritemSpacing(n.Ptr())
	return coregraphics.Float(float64(result_))
}

func (n NSCollectionViewGridLayout) SetMinimumInteritemSpacing(value coregraphics.Float) {
	C.C_NSCollectionViewGridLayout_SetMinimumInteritemSpacing(n.Ptr(), C.double(float64(value)))
}

func (n NSCollectionViewGridLayout) MinimumLineSpacing() coregraphics.Float {
	result_ := C.C_NSCollectionViewGridLayout_MinimumLineSpacing(n.Ptr())
	return coregraphics.Float(float64(result_))
}

func (n NSCollectionViewGridLayout) SetMinimumLineSpacing(value coregraphics.Float) {
	C.C_NSCollectionViewGridLayout_SetMinimumLineSpacing(n.Ptr(), C.double(float64(value)))
}

func (n NSCollectionViewGridLayout) Margins() foundation.EdgeInsets {
	result_ := C.C_NSCollectionViewGridLayout_Margins(n.Ptr())
	return foundation.FromNSEdgeInsetsPointer(unsafe.Pointer(&result_))
}

func (n NSCollectionViewGridLayout) SetMargins(value foundation.EdgeInsets) {
	C.C_NSCollectionViewGridLayout_SetMargins(n.Ptr(), *(*C.NSEdgeInsets)(foundation.ToNSEdgeInsetsPointer(value)))
}

func (n NSCollectionViewGridLayout) BackgroundColors() []Color {
	result_ := C.C_NSCollectionViewGridLayout_BackgroundColors(n.Ptr())
	defer C.free(result_.data)
	result_Slice := (*[1 << 28]unsafe.Pointer)(unsafe.Pointer(result_.data))[:result_.len:result_.len]
	var goResult_ = make([]Color, len(result_Slice))
	for idx, r := range result_Slice {
		goResult_[idx] = MakeColor(r)
	}
	return goResult_
}

func (n NSCollectionViewGridLayout) SetBackgroundColors(value []Color) {
	cValueData := make([]unsafe.Pointer, len(value))
	for idx, v := range value {
		cValueData[idx] = objc.ExtractPtr(v)
	}
	cValue := C.Array{data: unsafe.Pointer(&cValueData[0]), len: C.int(len(value))}
	C.C_NSCollectionViewGridLayout_SetBackgroundColors(n.Ptr(), cValue)
}
