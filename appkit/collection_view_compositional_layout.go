package appkit

// #include "collection_view_compositional_layout.h"
import "C"
import (
	"github.com/hsiafan/cocoa/coregraphics"
	"github.com/hsiafan/cocoa/foundation"
	"github.com/hsiafan/cocoa/objc"
	"unsafe"
)

type CollectionViewCompositionalLayout interface {
	CollectionViewLayout
	Configuration() CollectionViewCompositionalLayoutConfiguration
	SetConfiguration(value CollectionViewCompositionalLayoutConfiguration)
}

type NSCollectionViewCompositionalLayout struct {
	NSCollectionViewLayout
}

func MakeCollectionViewCompositionalLayout(ptr unsafe.Pointer) NSCollectionViewCompositionalLayout {
	return NSCollectionViewCompositionalLayout{
		NSCollectionViewLayout: MakeCollectionViewLayout(ptr),
	}
}

func (n NSCollectionViewCompositionalLayout) InitWithSection(section CollectionLayoutSection) NSCollectionViewCompositionalLayout {
	result_ := C.C_NSCollectionViewCompositionalLayout_InitWithSection(n.Ptr(), objc.ExtractPtr(section))
	return MakeCollectionViewCompositionalLayout(result_)
}

func (n NSCollectionViewCompositionalLayout) InitWithSection_Configuration(section CollectionLayoutSection, configuration CollectionViewCompositionalLayoutConfiguration) NSCollectionViewCompositionalLayout {
	result_ := C.C_NSCollectionViewCompositionalLayout_InitWithSection_Configuration(n.Ptr(), objc.ExtractPtr(section), objc.ExtractPtr(configuration))
	return MakeCollectionViewCompositionalLayout(result_)
}

func AllocCollectionViewCompositionalLayout() NSCollectionViewCompositionalLayout {
	result_ := C.C_NSCollectionViewCompositionalLayout_AllocCollectionViewCompositionalLayout()
	return MakeCollectionViewCompositionalLayout(result_)
}

func (n NSCollectionViewCompositionalLayout) Autorelease() NSCollectionViewCompositionalLayout {
	result_ := C.C_NSCollectionViewCompositionalLayout_Autorelease(n.Ptr())
	return MakeCollectionViewCompositionalLayout(result_)
}

func (n NSCollectionViewCompositionalLayout) Retain() NSCollectionViewCompositionalLayout {
	result_ := C.C_NSCollectionViewCompositionalLayout_Retain(n.Ptr())
	return MakeCollectionViewCompositionalLayout(result_)
}

func (n NSCollectionViewCompositionalLayout) Configuration() CollectionViewCompositionalLayoutConfiguration {
	result_ := C.C_NSCollectionViewCompositionalLayout_Configuration(n.Ptr())
	return MakeCollectionViewCompositionalLayoutConfiguration(result_)
}

func (n NSCollectionViewCompositionalLayout) SetConfiguration(value CollectionViewCompositionalLayoutConfiguration) {
	C.C_NSCollectionViewCompositionalLayout_SetConfiguration(n.Ptr(), objc.ExtractPtr(value))
}

type CollectionViewCompositionalLayoutConfiguration interface {
	objc.Object
	ScrollDirection() CollectionViewScrollDirection
	SetScrollDirection(value CollectionViewScrollDirection)
	InterSectionSpacing() coregraphics.Float
	SetInterSectionSpacing(value coregraphics.Float)
	BoundarySupplementaryItems() []CollectionLayoutBoundarySupplementaryItem
	SetBoundarySupplementaryItems(value []CollectionLayoutBoundarySupplementaryItem)
}

type NSCollectionViewCompositionalLayoutConfiguration struct {
	objc.NSObject
}

func MakeCollectionViewCompositionalLayoutConfiguration(ptr unsafe.Pointer) NSCollectionViewCompositionalLayoutConfiguration {
	return NSCollectionViewCompositionalLayoutConfiguration{
		NSObject: objc.MakeObject(ptr),
	}
}

func AllocCollectionViewCompositionalLayoutConfiguration() NSCollectionViewCompositionalLayoutConfiguration {
	result_ := C.C_NSCollectionViewCompositionalLayoutConfiguration_AllocCollectionViewCompositionalLayoutConfiguration()
	return MakeCollectionViewCompositionalLayoutConfiguration(result_)
}

func (n NSCollectionViewCompositionalLayoutConfiguration) Init() NSCollectionViewCompositionalLayoutConfiguration {
	result_ := C.C_NSCollectionViewCompositionalLayoutConfiguration_Init(n.Ptr())
	return MakeCollectionViewCompositionalLayoutConfiguration(result_)
}

func NewCollectionViewCompositionalLayoutConfiguration() NSCollectionViewCompositionalLayoutConfiguration {
	result_ := C.C_NSCollectionViewCompositionalLayoutConfiguration_NewCollectionViewCompositionalLayoutConfiguration()
	return MakeCollectionViewCompositionalLayoutConfiguration(result_)
}

func (n NSCollectionViewCompositionalLayoutConfiguration) Autorelease() NSCollectionViewCompositionalLayoutConfiguration {
	result_ := C.C_NSCollectionViewCompositionalLayoutConfiguration_Autorelease(n.Ptr())
	return MakeCollectionViewCompositionalLayoutConfiguration(result_)
}

func (n NSCollectionViewCompositionalLayoutConfiguration) Retain() NSCollectionViewCompositionalLayoutConfiguration {
	result_ := C.C_NSCollectionViewCompositionalLayoutConfiguration_Retain(n.Ptr())
	return MakeCollectionViewCompositionalLayoutConfiguration(result_)
}

func (n NSCollectionViewCompositionalLayoutConfiguration) ScrollDirection() CollectionViewScrollDirection {
	result_ := C.C_NSCollectionViewCompositionalLayoutConfiguration_ScrollDirection(n.Ptr())
	return CollectionViewScrollDirection(int(result_))
}

func (n NSCollectionViewCompositionalLayoutConfiguration) SetScrollDirection(value CollectionViewScrollDirection) {
	C.C_NSCollectionViewCompositionalLayoutConfiguration_SetScrollDirection(n.Ptr(), C.int(int(value)))
}

func (n NSCollectionViewCompositionalLayoutConfiguration) InterSectionSpacing() coregraphics.Float {
	result_ := C.C_NSCollectionViewCompositionalLayoutConfiguration_InterSectionSpacing(n.Ptr())
	return coregraphics.Float(float64(result_))
}

func (n NSCollectionViewCompositionalLayoutConfiguration) SetInterSectionSpacing(value coregraphics.Float) {
	C.C_NSCollectionViewCompositionalLayoutConfiguration_SetInterSectionSpacing(n.Ptr(), C.double(float64(value)))
}

func (n NSCollectionViewCompositionalLayoutConfiguration) BoundarySupplementaryItems() []CollectionLayoutBoundarySupplementaryItem {
	result_ := C.C_NSCollectionViewCompositionalLayoutConfiguration_BoundarySupplementaryItems(n.Ptr())
	if result_.len > 0 {
		defer C.free(result_.data)
	}
	result_Slice := unsafe.Slice((*unsafe.Pointer)(result_.data), int(result_.len))
	var goResult_ = make([]CollectionLayoutBoundarySupplementaryItem, len(result_Slice))
	for idx, r := range result_Slice {
		goResult_[idx] = MakeCollectionLayoutBoundarySupplementaryItem(r)
	}
	return goResult_
}

func (n NSCollectionViewCompositionalLayoutConfiguration) SetBoundarySupplementaryItems(value []CollectionLayoutBoundarySupplementaryItem) {
	var cValue C.Array
	if len(value) > 0 {
		cValueData := make([]unsafe.Pointer, len(value))
		for idx, v := range value {
			cValueData[idx] = objc.ExtractPtr(v)
		}
		cValue.data = unsafe.Pointer(&cValueData[0])
		cValue.len = C.int(len(value))
	}
	C.C_NSCollectionViewCompositionalLayoutConfiguration_SetBoundarySupplementaryItems(n.Ptr(), cValue)
}

type CollectionLayoutItem interface {
	objc.Object
	LayoutSize() CollectionLayoutSize
	SupplementaryItems() []CollectionLayoutSupplementaryItem
	EdgeSpacing() CollectionLayoutEdgeSpacing
	SetEdgeSpacing(value CollectionLayoutEdgeSpacing)
	ContentInsets() DirectionalEdgeInsets
	SetContentInsets(value DirectionalEdgeInsets)
}

type NSCollectionLayoutItem struct {
	objc.NSObject
}

func MakeCollectionLayoutItem(ptr unsafe.Pointer) NSCollectionLayoutItem {
	return NSCollectionLayoutItem{
		NSObject: objc.MakeObject(ptr),
	}
}

func CollectionLayoutItem_ItemWithLayoutSize(layoutSize CollectionLayoutSize) NSCollectionLayoutItem {
	result_ := C.C_NSCollectionLayoutItem_CollectionLayoutItem_ItemWithLayoutSize(objc.ExtractPtr(layoutSize))
	return MakeCollectionLayoutItem(result_)
}

func CollectionLayoutItem_ItemWithLayoutSize_SupplementaryItems(layoutSize CollectionLayoutSize, supplementaryItems []CollectionLayoutSupplementaryItem) NSCollectionLayoutItem {
	var cSupplementaryItems C.Array
	if len(supplementaryItems) > 0 {
		cSupplementaryItemsData := make([]unsafe.Pointer, len(supplementaryItems))
		for idx, v := range supplementaryItems {
			cSupplementaryItemsData[idx] = objc.ExtractPtr(v)
		}
		cSupplementaryItems.data = unsafe.Pointer(&cSupplementaryItemsData[0])
		cSupplementaryItems.len = C.int(len(supplementaryItems))
	}
	result_ := C.C_NSCollectionLayoutItem_CollectionLayoutItem_ItemWithLayoutSize_SupplementaryItems(objc.ExtractPtr(layoutSize), cSupplementaryItems)
	return MakeCollectionLayoutItem(result_)
}

func AllocCollectionLayoutItem() NSCollectionLayoutItem {
	result_ := C.C_NSCollectionLayoutItem_AllocCollectionLayoutItem()
	return MakeCollectionLayoutItem(result_)
}

func (n NSCollectionLayoutItem) Autorelease() NSCollectionLayoutItem {
	result_ := C.C_NSCollectionLayoutItem_Autorelease(n.Ptr())
	return MakeCollectionLayoutItem(result_)
}

func (n NSCollectionLayoutItem) Retain() NSCollectionLayoutItem {
	result_ := C.C_NSCollectionLayoutItem_Retain(n.Ptr())
	return MakeCollectionLayoutItem(result_)
}

func (n NSCollectionLayoutItem) LayoutSize() CollectionLayoutSize {
	result_ := C.C_NSCollectionLayoutItem_LayoutSize(n.Ptr())
	return MakeCollectionLayoutSize(result_)
}

func (n NSCollectionLayoutItem) SupplementaryItems() []CollectionLayoutSupplementaryItem {
	result_ := C.C_NSCollectionLayoutItem_SupplementaryItems(n.Ptr())
	if result_.len > 0 {
		defer C.free(result_.data)
	}
	result_Slice := unsafe.Slice((*unsafe.Pointer)(result_.data), int(result_.len))
	var goResult_ = make([]CollectionLayoutSupplementaryItem, len(result_Slice))
	for idx, r := range result_Slice {
		goResult_[idx] = MakeCollectionLayoutSupplementaryItem(r)
	}
	return goResult_
}

func (n NSCollectionLayoutItem) EdgeSpacing() CollectionLayoutEdgeSpacing {
	result_ := C.C_NSCollectionLayoutItem_EdgeSpacing(n.Ptr())
	return MakeCollectionLayoutEdgeSpacing(result_)
}

func (n NSCollectionLayoutItem) SetEdgeSpacing(value CollectionLayoutEdgeSpacing) {
	C.C_NSCollectionLayoutItem_SetEdgeSpacing(n.Ptr(), objc.ExtractPtr(value))
}

func (n NSCollectionLayoutItem) ContentInsets() DirectionalEdgeInsets {
	result_ := C.C_NSCollectionLayoutItem_ContentInsets(n.Ptr())
	return FromNSDirectionalEdgeInsetsPointer(unsafe.Pointer(&result_))
}

func (n NSCollectionLayoutItem) SetContentInsets(value DirectionalEdgeInsets) {
	C.C_NSCollectionLayoutItem_SetContentInsets(n.Ptr(), *(*C.NSDirectionalEdgeInsets)(ToNSDirectionalEdgeInsetsPointer(value)))
}

type CollectionLayoutBoundarySupplementaryItem interface {
	CollectionLayoutSupplementaryItem
	PinToVisibleBounds() bool
	SetPinToVisibleBounds(value bool)
	Offset() foundation.Point
	Alignment() RectAlignment
	ExtendsBoundary() bool
	SetExtendsBoundary(value bool)
}

type NSCollectionLayoutBoundarySupplementaryItem struct {
	NSCollectionLayoutSupplementaryItem
}

func MakeCollectionLayoutBoundarySupplementaryItem(ptr unsafe.Pointer) NSCollectionLayoutBoundarySupplementaryItem {
	return NSCollectionLayoutBoundarySupplementaryItem{
		NSCollectionLayoutSupplementaryItem: MakeCollectionLayoutSupplementaryItem(ptr),
	}
}

func CollectionLayoutBoundarySupplementaryItem_BoundarySupplementaryItemWithLayoutSize_ElementKind_Alignment(layoutSize CollectionLayoutSize, elementKind string, alignment RectAlignment) NSCollectionLayoutBoundarySupplementaryItem {
	result_ := C.C_NSCollectionLayoutBoundarySupplementaryItem_CollectionLayoutBoundarySupplementaryItem_BoundarySupplementaryItemWithLayoutSize_ElementKind_Alignment(objc.ExtractPtr(layoutSize), foundation.NewString(elementKind).Ptr(), C.int(int(alignment)))
	return MakeCollectionLayoutBoundarySupplementaryItem(result_)
}

func CollectionLayoutBoundarySupplementaryItem_BoundarySupplementaryItemWithLayoutSize_ElementKind_Alignment_AbsoluteOffset(layoutSize CollectionLayoutSize, elementKind string, alignment RectAlignment, absoluteOffset foundation.Point) NSCollectionLayoutBoundarySupplementaryItem {
	result_ := C.C_NSCollectionLayoutBoundarySupplementaryItem_CollectionLayoutBoundarySupplementaryItem_BoundarySupplementaryItemWithLayoutSize_ElementKind_Alignment_AbsoluteOffset(objc.ExtractPtr(layoutSize), foundation.NewString(elementKind).Ptr(), C.int(int(alignment)), *(*C.CGPoint)(coregraphics.ToCGPointPointer(coregraphics.Point(absoluteOffset))))
	return MakeCollectionLayoutBoundarySupplementaryItem(result_)
}

func CollectionLayoutBoundarySupplementaryItem_SupplementaryItemWithLayoutSize_ElementKind_ContainerAnchor(layoutSize CollectionLayoutSize, elementKind string, containerAnchor CollectionLayoutAnchor) NSCollectionLayoutBoundarySupplementaryItem {
	result_ := C.C_NSCollectionLayoutBoundarySupplementaryItem_CollectionLayoutBoundarySupplementaryItem_SupplementaryItemWithLayoutSize_ElementKind_ContainerAnchor(objc.ExtractPtr(layoutSize), foundation.NewString(elementKind).Ptr(), objc.ExtractPtr(containerAnchor))
	return MakeCollectionLayoutBoundarySupplementaryItem(result_)
}

func CollectionLayoutBoundarySupplementaryItem_SupplementaryItemWithLayoutSize_ElementKind_ContainerAnchor_ItemAnchor(layoutSize CollectionLayoutSize, elementKind string, containerAnchor CollectionLayoutAnchor, itemAnchor CollectionLayoutAnchor) NSCollectionLayoutBoundarySupplementaryItem {
	result_ := C.C_NSCollectionLayoutBoundarySupplementaryItem_CollectionLayoutBoundarySupplementaryItem_SupplementaryItemWithLayoutSize_ElementKind_ContainerAnchor_ItemAnchor(objc.ExtractPtr(layoutSize), foundation.NewString(elementKind).Ptr(), objc.ExtractPtr(containerAnchor), objc.ExtractPtr(itemAnchor))
	return MakeCollectionLayoutBoundarySupplementaryItem(result_)
}

func CollectionLayoutBoundarySupplementaryItem_ItemWithLayoutSize(layoutSize CollectionLayoutSize) NSCollectionLayoutBoundarySupplementaryItem {
	result_ := C.C_NSCollectionLayoutBoundarySupplementaryItem_CollectionLayoutBoundarySupplementaryItem_ItemWithLayoutSize(objc.ExtractPtr(layoutSize))
	return MakeCollectionLayoutBoundarySupplementaryItem(result_)
}

func CollectionLayoutBoundarySupplementaryItem_ItemWithLayoutSize_SupplementaryItems(layoutSize CollectionLayoutSize, supplementaryItems []CollectionLayoutSupplementaryItem) NSCollectionLayoutBoundarySupplementaryItem {
	var cSupplementaryItems C.Array
	if len(supplementaryItems) > 0 {
		cSupplementaryItemsData := make([]unsafe.Pointer, len(supplementaryItems))
		for idx, v := range supplementaryItems {
			cSupplementaryItemsData[idx] = objc.ExtractPtr(v)
		}
		cSupplementaryItems.data = unsafe.Pointer(&cSupplementaryItemsData[0])
		cSupplementaryItems.len = C.int(len(supplementaryItems))
	}
	result_ := C.C_NSCollectionLayoutBoundarySupplementaryItem_CollectionLayoutBoundarySupplementaryItem_ItemWithLayoutSize_SupplementaryItems(objc.ExtractPtr(layoutSize), cSupplementaryItems)
	return MakeCollectionLayoutBoundarySupplementaryItem(result_)
}

func AllocCollectionLayoutBoundarySupplementaryItem() NSCollectionLayoutBoundarySupplementaryItem {
	result_ := C.C_NSCollectionLayoutBoundarySupplementaryItem_AllocCollectionLayoutBoundarySupplementaryItem()
	return MakeCollectionLayoutBoundarySupplementaryItem(result_)
}

func (n NSCollectionLayoutBoundarySupplementaryItem) Autorelease() NSCollectionLayoutBoundarySupplementaryItem {
	result_ := C.C_NSCollectionLayoutBoundarySupplementaryItem_Autorelease(n.Ptr())
	return MakeCollectionLayoutBoundarySupplementaryItem(result_)
}

func (n NSCollectionLayoutBoundarySupplementaryItem) Retain() NSCollectionLayoutBoundarySupplementaryItem {
	result_ := C.C_NSCollectionLayoutBoundarySupplementaryItem_Retain(n.Ptr())
	return MakeCollectionLayoutBoundarySupplementaryItem(result_)
}

func (n NSCollectionLayoutBoundarySupplementaryItem) PinToVisibleBounds() bool {
	result_ := C.C_NSCollectionLayoutBoundarySupplementaryItem_PinToVisibleBounds(n.Ptr())
	return bool(result_)
}

func (n NSCollectionLayoutBoundarySupplementaryItem) SetPinToVisibleBounds(value bool) {
	C.C_NSCollectionLayoutBoundarySupplementaryItem_SetPinToVisibleBounds(n.Ptr(), C.bool(value))
}

func (n NSCollectionLayoutBoundarySupplementaryItem) Offset() foundation.Point {
	result_ := C.C_NSCollectionLayoutBoundarySupplementaryItem_Offset(n.Ptr())
	return foundation.Point(coregraphics.FromCGPointPointer(unsafe.Pointer(&result_)))
}

func (n NSCollectionLayoutBoundarySupplementaryItem) Alignment() RectAlignment {
	result_ := C.C_NSCollectionLayoutBoundarySupplementaryItem_Alignment(n.Ptr())
	return RectAlignment(int(result_))
}

func (n NSCollectionLayoutBoundarySupplementaryItem) ExtendsBoundary() bool {
	result_ := C.C_NSCollectionLayoutBoundarySupplementaryItem_ExtendsBoundary(n.Ptr())
	return bool(result_)
}

func (n NSCollectionLayoutBoundarySupplementaryItem) SetExtendsBoundary(value bool) {
	C.C_NSCollectionLayoutBoundarySupplementaryItem_SetExtendsBoundary(n.Ptr(), C.bool(value))
}

type CollectionLayoutSpacing interface {
	objc.Object
	Spacing() coregraphics.Float
	IsFixedSpacing() bool
	IsFlexibleSpacing() bool
}

type NSCollectionLayoutSpacing struct {
	objc.NSObject
}

func MakeCollectionLayoutSpacing(ptr unsafe.Pointer) NSCollectionLayoutSpacing {
	return NSCollectionLayoutSpacing{
		NSObject: objc.MakeObject(ptr),
	}
}

func CollectionLayoutSpacing_FixedSpacing(fixedSpacing coregraphics.Float) NSCollectionLayoutSpacing {
	result_ := C.C_NSCollectionLayoutSpacing_CollectionLayoutSpacing_FixedSpacing(C.double(float64(fixedSpacing)))
	return MakeCollectionLayoutSpacing(result_)
}

func CollectionLayoutSpacing_FlexibleSpacing(flexibleSpacing coregraphics.Float) NSCollectionLayoutSpacing {
	result_ := C.C_NSCollectionLayoutSpacing_CollectionLayoutSpacing_FlexibleSpacing(C.double(float64(flexibleSpacing)))
	return MakeCollectionLayoutSpacing(result_)
}

func AllocCollectionLayoutSpacing() NSCollectionLayoutSpacing {
	result_ := C.C_NSCollectionLayoutSpacing_AllocCollectionLayoutSpacing()
	return MakeCollectionLayoutSpacing(result_)
}

func (n NSCollectionLayoutSpacing) Autorelease() NSCollectionLayoutSpacing {
	result_ := C.C_NSCollectionLayoutSpacing_Autorelease(n.Ptr())
	return MakeCollectionLayoutSpacing(result_)
}

func (n NSCollectionLayoutSpacing) Retain() NSCollectionLayoutSpacing {
	result_ := C.C_NSCollectionLayoutSpacing_Retain(n.Ptr())
	return MakeCollectionLayoutSpacing(result_)
}

func (n NSCollectionLayoutSpacing) Spacing() coregraphics.Float {
	result_ := C.C_NSCollectionLayoutSpacing_Spacing(n.Ptr())
	return coregraphics.Float(float64(result_))
}

func (n NSCollectionLayoutSpacing) IsFixedSpacing() bool {
	result_ := C.C_NSCollectionLayoutSpacing_IsFixedSpacing(n.Ptr())
	return bool(result_)
}

func (n NSCollectionLayoutSpacing) IsFlexibleSpacing() bool {
	result_ := C.C_NSCollectionLayoutSpacing_IsFlexibleSpacing(n.Ptr())
	return bool(result_)
}

type CollectionLayoutSection interface {
	objc.Object
	OrthogonalScrollingBehavior() CollectionLayoutSectionOrthogonalScrollingBehavior
	SetOrthogonalScrollingBehavior(value CollectionLayoutSectionOrthogonalScrollingBehavior)
	InterGroupSpacing() coregraphics.Float
	SetInterGroupSpacing(value coregraphics.Float)
	ContentInsets() DirectionalEdgeInsets
	SetContentInsets(value DirectionalEdgeInsets)
	SupplementariesFollowContentInsets() bool
	SetSupplementariesFollowContentInsets(value bool)
	BoundarySupplementaryItems() []CollectionLayoutBoundarySupplementaryItem
	SetBoundarySupplementaryItems(value []CollectionLayoutBoundarySupplementaryItem)
	DecorationItems() []CollectionLayoutDecorationItem
	SetDecorationItems(value []CollectionLayoutDecorationItem)
}

type NSCollectionLayoutSection struct {
	objc.NSObject
}

func MakeCollectionLayoutSection(ptr unsafe.Pointer) NSCollectionLayoutSection {
	return NSCollectionLayoutSection{
		NSObject: objc.MakeObject(ptr),
	}
}

func CollectionLayoutSection_SectionWithGroup(group CollectionLayoutGroup) NSCollectionLayoutSection {
	result_ := C.C_NSCollectionLayoutSection_CollectionLayoutSection_SectionWithGroup(objc.ExtractPtr(group))
	return MakeCollectionLayoutSection(result_)
}

func AllocCollectionLayoutSection() NSCollectionLayoutSection {
	result_ := C.C_NSCollectionLayoutSection_AllocCollectionLayoutSection()
	return MakeCollectionLayoutSection(result_)
}

func (n NSCollectionLayoutSection) Autorelease() NSCollectionLayoutSection {
	result_ := C.C_NSCollectionLayoutSection_Autorelease(n.Ptr())
	return MakeCollectionLayoutSection(result_)
}

func (n NSCollectionLayoutSection) Retain() NSCollectionLayoutSection {
	result_ := C.C_NSCollectionLayoutSection_Retain(n.Ptr())
	return MakeCollectionLayoutSection(result_)
}

func (n NSCollectionLayoutSection) OrthogonalScrollingBehavior() CollectionLayoutSectionOrthogonalScrollingBehavior {
	result_ := C.C_NSCollectionLayoutSection_OrthogonalScrollingBehavior(n.Ptr())
	return CollectionLayoutSectionOrthogonalScrollingBehavior(int(result_))
}

func (n NSCollectionLayoutSection) SetOrthogonalScrollingBehavior(value CollectionLayoutSectionOrthogonalScrollingBehavior) {
	C.C_NSCollectionLayoutSection_SetOrthogonalScrollingBehavior(n.Ptr(), C.int(int(value)))
}

func (n NSCollectionLayoutSection) InterGroupSpacing() coregraphics.Float {
	result_ := C.C_NSCollectionLayoutSection_InterGroupSpacing(n.Ptr())
	return coregraphics.Float(float64(result_))
}

func (n NSCollectionLayoutSection) SetInterGroupSpacing(value coregraphics.Float) {
	C.C_NSCollectionLayoutSection_SetInterGroupSpacing(n.Ptr(), C.double(float64(value)))
}

func (n NSCollectionLayoutSection) ContentInsets() DirectionalEdgeInsets {
	result_ := C.C_NSCollectionLayoutSection_ContentInsets(n.Ptr())
	return FromNSDirectionalEdgeInsetsPointer(unsafe.Pointer(&result_))
}

func (n NSCollectionLayoutSection) SetContentInsets(value DirectionalEdgeInsets) {
	C.C_NSCollectionLayoutSection_SetContentInsets(n.Ptr(), *(*C.NSDirectionalEdgeInsets)(ToNSDirectionalEdgeInsetsPointer(value)))
}

func (n NSCollectionLayoutSection) SupplementariesFollowContentInsets() bool {
	result_ := C.C_NSCollectionLayoutSection_SupplementariesFollowContentInsets(n.Ptr())
	return bool(result_)
}

func (n NSCollectionLayoutSection) SetSupplementariesFollowContentInsets(value bool) {
	C.C_NSCollectionLayoutSection_SetSupplementariesFollowContentInsets(n.Ptr(), C.bool(value))
}

func (n NSCollectionLayoutSection) BoundarySupplementaryItems() []CollectionLayoutBoundarySupplementaryItem {
	result_ := C.C_NSCollectionLayoutSection_BoundarySupplementaryItems(n.Ptr())
	if result_.len > 0 {
		defer C.free(result_.data)
	}
	result_Slice := unsafe.Slice((*unsafe.Pointer)(result_.data), int(result_.len))
	var goResult_ = make([]CollectionLayoutBoundarySupplementaryItem, len(result_Slice))
	for idx, r := range result_Slice {
		goResult_[idx] = MakeCollectionLayoutBoundarySupplementaryItem(r)
	}
	return goResult_
}

func (n NSCollectionLayoutSection) SetBoundarySupplementaryItems(value []CollectionLayoutBoundarySupplementaryItem) {
	var cValue C.Array
	if len(value) > 0 {
		cValueData := make([]unsafe.Pointer, len(value))
		for idx, v := range value {
			cValueData[idx] = objc.ExtractPtr(v)
		}
		cValue.data = unsafe.Pointer(&cValueData[0])
		cValue.len = C.int(len(value))
	}
	C.C_NSCollectionLayoutSection_SetBoundarySupplementaryItems(n.Ptr(), cValue)
}

func (n NSCollectionLayoutSection) DecorationItems() []CollectionLayoutDecorationItem {
	result_ := C.C_NSCollectionLayoutSection_DecorationItems(n.Ptr())
	if result_.len > 0 {
		defer C.free(result_.data)
	}
	result_Slice := unsafe.Slice((*unsafe.Pointer)(result_.data), int(result_.len))
	var goResult_ = make([]CollectionLayoutDecorationItem, len(result_Slice))
	for idx, r := range result_Slice {
		goResult_[idx] = MakeCollectionLayoutDecorationItem(r)
	}
	return goResult_
}

func (n NSCollectionLayoutSection) SetDecorationItems(value []CollectionLayoutDecorationItem) {
	var cValue C.Array
	if len(value) > 0 {
		cValueData := make([]unsafe.Pointer, len(value))
		for idx, v := range value {
			cValueData[idx] = objc.ExtractPtr(v)
		}
		cValue.data = unsafe.Pointer(&cValueData[0])
		cValue.len = C.int(len(value))
	}
	C.C_NSCollectionLayoutSection_SetDecorationItems(n.Ptr(), cValue)
}

type CollectionLayoutGroupCustomItem interface {
	objc.Object
	Frame() foundation.Rect
	ZIndex() int
}

type NSCollectionLayoutGroupCustomItem struct {
	objc.NSObject
}

func MakeCollectionLayoutGroupCustomItem(ptr unsafe.Pointer) NSCollectionLayoutGroupCustomItem {
	return NSCollectionLayoutGroupCustomItem{
		NSObject: objc.MakeObject(ptr),
	}
}

func CollectionLayoutGroupCustomItem_CustomItemWithFrame(frame foundation.Rect) NSCollectionLayoutGroupCustomItem {
	result_ := C.C_NSCollectionLayoutGroupCustomItem_CollectionLayoutGroupCustomItem_CustomItemWithFrame(*(*C.CGRect)(coregraphics.ToCGRectPointer(coregraphics.Rect(frame))))
	return MakeCollectionLayoutGroupCustomItem(result_)
}

func CollectionLayoutGroupCustomItem_CustomItemWithFrame_ZIndex(frame foundation.Rect, zIndex int) NSCollectionLayoutGroupCustomItem {
	result_ := C.C_NSCollectionLayoutGroupCustomItem_CollectionLayoutGroupCustomItem_CustomItemWithFrame_ZIndex(*(*C.CGRect)(coregraphics.ToCGRectPointer(coregraphics.Rect(frame))), C.int(zIndex))
	return MakeCollectionLayoutGroupCustomItem(result_)
}

func AllocCollectionLayoutGroupCustomItem() NSCollectionLayoutGroupCustomItem {
	result_ := C.C_NSCollectionLayoutGroupCustomItem_AllocCollectionLayoutGroupCustomItem()
	return MakeCollectionLayoutGroupCustomItem(result_)
}

func (n NSCollectionLayoutGroupCustomItem) Autorelease() NSCollectionLayoutGroupCustomItem {
	result_ := C.C_NSCollectionLayoutGroupCustomItem_Autorelease(n.Ptr())
	return MakeCollectionLayoutGroupCustomItem(result_)
}

func (n NSCollectionLayoutGroupCustomItem) Retain() NSCollectionLayoutGroupCustomItem {
	result_ := C.C_NSCollectionLayoutGroupCustomItem_Retain(n.Ptr())
	return MakeCollectionLayoutGroupCustomItem(result_)
}

func (n NSCollectionLayoutGroupCustomItem) Frame() foundation.Rect {
	result_ := C.C_NSCollectionLayoutGroupCustomItem_Frame(n.Ptr())
	return foundation.Rect(coregraphics.FromCGRectPointer(unsafe.Pointer(&result_)))
}

func (n NSCollectionLayoutGroupCustomItem) ZIndex() int {
	result_ := C.C_NSCollectionLayoutGroupCustomItem_ZIndex(n.Ptr())
	return int(result_)
}

type CollectionLayoutSupplementaryItem interface {
	CollectionLayoutItem
	ItemAnchor() CollectionLayoutAnchor
	ContainerAnchor() CollectionLayoutAnchor
	ElementKind() string
	ZIndex() int
	SetZIndex(value int)
}

type NSCollectionLayoutSupplementaryItem struct {
	NSCollectionLayoutItem
}

func MakeCollectionLayoutSupplementaryItem(ptr unsafe.Pointer) NSCollectionLayoutSupplementaryItem {
	return NSCollectionLayoutSupplementaryItem{
		NSCollectionLayoutItem: MakeCollectionLayoutItem(ptr),
	}
}

func CollectionLayoutSupplementaryItem_SupplementaryItemWithLayoutSize_ElementKind_ContainerAnchor(layoutSize CollectionLayoutSize, elementKind string, containerAnchor CollectionLayoutAnchor) NSCollectionLayoutSupplementaryItem {
	result_ := C.C_NSCollectionLayoutSupplementaryItem_CollectionLayoutSupplementaryItem_SupplementaryItemWithLayoutSize_ElementKind_ContainerAnchor(objc.ExtractPtr(layoutSize), foundation.NewString(elementKind).Ptr(), objc.ExtractPtr(containerAnchor))
	return MakeCollectionLayoutSupplementaryItem(result_)
}

func CollectionLayoutSupplementaryItem_SupplementaryItemWithLayoutSize_ElementKind_ContainerAnchor_ItemAnchor(layoutSize CollectionLayoutSize, elementKind string, containerAnchor CollectionLayoutAnchor, itemAnchor CollectionLayoutAnchor) NSCollectionLayoutSupplementaryItem {
	result_ := C.C_NSCollectionLayoutSupplementaryItem_CollectionLayoutSupplementaryItem_SupplementaryItemWithLayoutSize_ElementKind_ContainerAnchor_ItemAnchor(objc.ExtractPtr(layoutSize), foundation.NewString(elementKind).Ptr(), objc.ExtractPtr(containerAnchor), objc.ExtractPtr(itemAnchor))
	return MakeCollectionLayoutSupplementaryItem(result_)
}

func CollectionLayoutSupplementaryItem_ItemWithLayoutSize(layoutSize CollectionLayoutSize) NSCollectionLayoutSupplementaryItem {
	result_ := C.C_NSCollectionLayoutSupplementaryItem_CollectionLayoutSupplementaryItem_ItemWithLayoutSize(objc.ExtractPtr(layoutSize))
	return MakeCollectionLayoutSupplementaryItem(result_)
}

func CollectionLayoutSupplementaryItem_ItemWithLayoutSize_SupplementaryItems(layoutSize CollectionLayoutSize, supplementaryItems []CollectionLayoutSupplementaryItem) NSCollectionLayoutSupplementaryItem {
	var cSupplementaryItems C.Array
	if len(supplementaryItems) > 0 {
		cSupplementaryItemsData := make([]unsafe.Pointer, len(supplementaryItems))
		for idx, v := range supplementaryItems {
			cSupplementaryItemsData[idx] = objc.ExtractPtr(v)
		}
		cSupplementaryItems.data = unsafe.Pointer(&cSupplementaryItemsData[0])
		cSupplementaryItems.len = C.int(len(supplementaryItems))
	}
	result_ := C.C_NSCollectionLayoutSupplementaryItem_CollectionLayoutSupplementaryItem_ItemWithLayoutSize_SupplementaryItems(objc.ExtractPtr(layoutSize), cSupplementaryItems)
	return MakeCollectionLayoutSupplementaryItem(result_)
}

func AllocCollectionLayoutSupplementaryItem() NSCollectionLayoutSupplementaryItem {
	result_ := C.C_NSCollectionLayoutSupplementaryItem_AllocCollectionLayoutSupplementaryItem()
	return MakeCollectionLayoutSupplementaryItem(result_)
}

func (n NSCollectionLayoutSupplementaryItem) Autorelease() NSCollectionLayoutSupplementaryItem {
	result_ := C.C_NSCollectionLayoutSupplementaryItem_Autorelease(n.Ptr())
	return MakeCollectionLayoutSupplementaryItem(result_)
}

func (n NSCollectionLayoutSupplementaryItem) Retain() NSCollectionLayoutSupplementaryItem {
	result_ := C.C_NSCollectionLayoutSupplementaryItem_Retain(n.Ptr())
	return MakeCollectionLayoutSupplementaryItem(result_)
}

func (n NSCollectionLayoutSupplementaryItem) ItemAnchor() CollectionLayoutAnchor {
	result_ := C.C_NSCollectionLayoutSupplementaryItem_ItemAnchor(n.Ptr())
	return MakeCollectionLayoutAnchor(result_)
}

func (n NSCollectionLayoutSupplementaryItem) ContainerAnchor() CollectionLayoutAnchor {
	result_ := C.C_NSCollectionLayoutSupplementaryItem_ContainerAnchor(n.Ptr())
	return MakeCollectionLayoutAnchor(result_)
}

func (n NSCollectionLayoutSupplementaryItem) ElementKind() string {
	result_ := C.C_NSCollectionLayoutSupplementaryItem_ElementKind(n.Ptr())
	return foundation.MakeString(result_).String()
}

func (n NSCollectionLayoutSupplementaryItem) ZIndex() int {
	result_ := C.C_NSCollectionLayoutSupplementaryItem_ZIndex(n.Ptr())
	return int(result_)
}

func (n NSCollectionLayoutSupplementaryItem) SetZIndex(value int) {
	C.C_NSCollectionLayoutSupplementaryItem_SetZIndex(n.Ptr(), C.int(value))
}

type CollectionLayoutSize interface {
	objc.Object
	WidthDimension() CollectionLayoutDimension
	HeightDimension() CollectionLayoutDimension
}

type NSCollectionLayoutSize struct {
	objc.NSObject
}

func MakeCollectionLayoutSize(ptr unsafe.Pointer) NSCollectionLayoutSize {
	return NSCollectionLayoutSize{
		NSObject: objc.MakeObject(ptr),
	}
}

func CollectionLayoutSize_SizeWithWidthDimension_HeightDimension(width CollectionLayoutDimension, height CollectionLayoutDimension) NSCollectionLayoutSize {
	result_ := C.C_NSCollectionLayoutSize_CollectionLayoutSize_SizeWithWidthDimension_HeightDimension(objc.ExtractPtr(width), objc.ExtractPtr(height))
	return MakeCollectionLayoutSize(result_)
}

func AllocCollectionLayoutSize() NSCollectionLayoutSize {
	result_ := C.C_NSCollectionLayoutSize_AllocCollectionLayoutSize()
	return MakeCollectionLayoutSize(result_)
}

func (n NSCollectionLayoutSize) Autorelease() NSCollectionLayoutSize {
	result_ := C.C_NSCollectionLayoutSize_Autorelease(n.Ptr())
	return MakeCollectionLayoutSize(result_)
}

func (n NSCollectionLayoutSize) Retain() NSCollectionLayoutSize {
	result_ := C.C_NSCollectionLayoutSize_Retain(n.Ptr())
	return MakeCollectionLayoutSize(result_)
}

func (n NSCollectionLayoutSize) WidthDimension() CollectionLayoutDimension {
	result_ := C.C_NSCollectionLayoutSize_WidthDimension(n.Ptr())
	return MakeCollectionLayoutDimension(result_)
}

func (n NSCollectionLayoutSize) HeightDimension() CollectionLayoutDimension {
	result_ := C.C_NSCollectionLayoutSize_HeightDimension(n.Ptr())
	return MakeCollectionLayoutDimension(result_)
}

type CollectionLayoutEdgeSpacing interface {
	objc.Object
	Leading() CollectionLayoutSpacing
	Top() CollectionLayoutSpacing
	Trailing() CollectionLayoutSpacing
	Bottom() CollectionLayoutSpacing
}

type NSCollectionLayoutEdgeSpacing struct {
	objc.NSObject
}

func MakeCollectionLayoutEdgeSpacing(ptr unsafe.Pointer) NSCollectionLayoutEdgeSpacing {
	return NSCollectionLayoutEdgeSpacing{
		NSObject: objc.MakeObject(ptr),
	}
}

func CollectionLayoutEdgeSpacing_SpacingForLeading_Top_Trailing_Bottom(leading CollectionLayoutSpacing, top CollectionLayoutSpacing, trailing CollectionLayoutSpacing, bottom CollectionLayoutSpacing) NSCollectionLayoutEdgeSpacing {
	result_ := C.C_NSCollectionLayoutEdgeSpacing_CollectionLayoutEdgeSpacing_SpacingForLeading_Top_Trailing_Bottom(objc.ExtractPtr(leading), objc.ExtractPtr(top), objc.ExtractPtr(trailing), objc.ExtractPtr(bottom))
	return MakeCollectionLayoutEdgeSpacing(result_)
}

func AllocCollectionLayoutEdgeSpacing() NSCollectionLayoutEdgeSpacing {
	result_ := C.C_NSCollectionLayoutEdgeSpacing_AllocCollectionLayoutEdgeSpacing()
	return MakeCollectionLayoutEdgeSpacing(result_)
}

func (n NSCollectionLayoutEdgeSpacing) Autorelease() NSCollectionLayoutEdgeSpacing {
	result_ := C.C_NSCollectionLayoutEdgeSpacing_Autorelease(n.Ptr())
	return MakeCollectionLayoutEdgeSpacing(result_)
}

func (n NSCollectionLayoutEdgeSpacing) Retain() NSCollectionLayoutEdgeSpacing {
	result_ := C.C_NSCollectionLayoutEdgeSpacing_Retain(n.Ptr())
	return MakeCollectionLayoutEdgeSpacing(result_)
}

func (n NSCollectionLayoutEdgeSpacing) Leading() CollectionLayoutSpacing {
	result_ := C.C_NSCollectionLayoutEdgeSpacing_Leading(n.Ptr())
	return MakeCollectionLayoutSpacing(result_)
}

func (n NSCollectionLayoutEdgeSpacing) Top() CollectionLayoutSpacing {
	result_ := C.C_NSCollectionLayoutEdgeSpacing_Top(n.Ptr())
	return MakeCollectionLayoutSpacing(result_)
}

func (n NSCollectionLayoutEdgeSpacing) Trailing() CollectionLayoutSpacing {
	result_ := C.C_NSCollectionLayoutEdgeSpacing_Trailing(n.Ptr())
	return MakeCollectionLayoutSpacing(result_)
}

func (n NSCollectionLayoutEdgeSpacing) Bottom() CollectionLayoutSpacing {
	result_ := C.C_NSCollectionLayoutEdgeSpacing_Bottom(n.Ptr())
	return MakeCollectionLayoutSpacing(result_)
}

type CollectionLayoutAnchor interface {
	objc.Object
	Edges() DirectionalRectEdge
	Offset() foundation.Point
	IsAbsoluteOffset() bool
	IsFractionalOffset() bool
}

type NSCollectionLayoutAnchor struct {
	objc.NSObject
}

func MakeCollectionLayoutAnchor(ptr unsafe.Pointer) NSCollectionLayoutAnchor {
	return NSCollectionLayoutAnchor{
		NSObject: objc.MakeObject(ptr),
	}
}

func CollectionLayoutAnchor_LayoutAnchorWithEdges(edges DirectionalRectEdge) NSCollectionLayoutAnchor {
	result_ := C.C_NSCollectionLayoutAnchor_CollectionLayoutAnchor_LayoutAnchorWithEdges(C.uint(uint(edges)))
	return MakeCollectionLayoutAnchor(result_)
}

func CollectionLayoutAnchor_LayoutAnchorWithEdges_AbsoluteOffset(edges DirectionalRectEdge, absoluteOffset foundation.Point) NSCollectionLayoutAnchor {
	result_ := C.C_NSCollectionLayoutAnchor_CollectionLayoutAnchor_LayoutAnchorWithEdges_AbsoluteOffset(C.uint(uint(edges)), *(*C.CGPoint)(coregraphics.ToCGPointPointer(coregraphics.Point(absoluteOffset))))
	return MakeCollectionLayoutAnchor(result_)
}

func CollectionLayoutAnchor_LayoutAnchorWithEdges_FractionalOffset(edges DirectionalRectEdge, fractionalOffset foundation.Point) NSCollectionLayoutAnchor {
	result_ := C.C_NSCollectionLayoutAnchor_CollectionLayoutAnchor_LayoutAnchorWithEdges_FractionalOffset(C.uint(uint(edges)), *(*C.CGPoint)(coregraphics.ToCGPointPointer(coregraphics.Point(fractionalOffset))))
	return MakeCollectionLayoutAnchor(result_)
}

func AllocCollectionLayoutAnchor() NSCollectionLayoutAnchor {
	result_ := C.C_NSCollectionLayoutAnchor_AllocCollectionLayoutAnchor()
	return MakeCollectionLayoutAnchor(result_)
}

func (n NSCollectionLayoutAnchor) Autorelease() NSCollectionLayoutAnchor {
	result_ := C.C_NSCollectionLayoutAnchor_Autorelease(n.Ptr())
	return MakeCollectionLayoutAnchor(result_)
}

func (n NSCollectionLayoutAnchor) Retain() NSCollectionLayoutAnchor {
	result_ := C.C_NSCollectionLayoutAnchor_Retain(n.Ptr())
	return MakeCollectionLayoutAnchor(result_)
}

func (n NSCollectionLayoutAnchor) Edges() DirectionalRectEdge {
	result_ := C.C_NSCollectionLayoutAnchor_Edges(n.Ptr())
	return DirectionalRectEdge(uint(result_))
}

func (n NSCollectionLayoutAnchor) Offset() foundation.Point {
	result_ := C.C_NSCollectionLayoutAnchor_Offset(n.Ptr())
	return foundation.Point(coregraphics.FromCGPointPointer(unsafe.Pointer(&result_)))
}

func (n NSCollectionLayoutAnchor) IsAbsoluteOffset() bool {
	result_ := C.C_NSCollectionLayoutAnchor_IsAbsoluteOffset(n.Ptr())
	return bool(result_)
}

func (n NSCollectionLayoutAnchor) IsFractionalOffset() bool {
	result_ := C.C_NSCollectionLayoutAnchor_IsFractionalOffset(n.Ptr())
	return bool(result_)
}

type CollectionLayoutDimension interface {
	objc.Object
	Dimension() coregraphics.Float
	IsAbsolute() bool
	IsEstimated() bool
	IsFractionalHeight() bool
	IsFractionalWidth() bool
}

type NSCollectionLayoutDimension struct {
	objc.NSObject
}

func MakeCollectionLayoutDimension(ptr unsafe.Pointer) NSCollectionLayoutDimension {
	return NSCollectionLayoutDimension{
		NSObject: objc.MakeObject(ptr),
	}
}

func CollectionLayoutDimension_AbsoluteDimension(absoluteDimension coregraphics.Float) NSCollectionLayoutDimension {
	result_ := C.C_NSCollectionLayoutDimension_CollectionLayoutDimension_AbsoluteDimension(C.double(float64(absoluteDimension)))
	return MakeCollectionLayoutDimension(result_)
}

func CollectionLayoutDimension_EstimatedDimension(estimatedDimension coregraphics.Float) NSCollectionLayoutDimension {
	result_ := C.C_NSCollectionLayoutDimension_CollectionLayoutDimension_EstimatedDimension(C.double(float64(estimatedDimension)))
	return MakeCollectionLayoutDimension(result_)
}

func CollectionLayoutDimension_FractionalHeightDimension(fractionalHeight coregraphics.Float) NSCollectionLayoutDimension {
	result_ := C.C_NSCollectionLayoutDimension_CollectionLayoutDimension_FractionalHeightDimension(C.double(float64(fractionalHeight)))
	return MakeCollectionLayoutDimension(result_)
}

func CollectionLayoutDimension_FractionalWidthDimension(fractionalWidth coregraphics.Float) NSCollectionLayoutDimension {
	result_ := C.C_NSCollectionLayoutDimension_CollectionLayoutDimension_FractionalWidthDimension(C.double(float64(fractionalWidth)))
	return MakeCollectionLayoutDimension(result_)
}

func AllocCollectionLayoutDimension() NSCollectionLayoutDimension {
	result_ := C.C_NSCollectionLayoutDimension_AllocCollectionLayoutDimension()
	return MakeCollectionLayoutDimension(result_)
}

func (n NSCollectionLayoutDimension) Autorelease() NSCollectionLayoutDimension {
	result_ := C.C_NSCollectionLayoutDimension_Autorelease(n.Ptr())
	return MakeCollectionLayoutDimension(result_)
}

func (n NSCollectionLayoutDimension) Retain() NSCollectionLayoutDimension {
	result_ := C.C_NSCollectionLayoutDimension_Retain(n.Ptr())
	return MakeCollectionLayoutDimension(result_)
}

func (n NSCollectionLayoutDimension) Dimension() coregraphics.Float {
	result_ := C.C_NSCollectionLayoutDimension_Dimension(n.Ptr())
	return coregraphics.Float(float64(result_))
}

func (n NSCollectionLayoutDimension) IsAbsolute() bool {
	result_ := C.C_NSCollectionLayoutDimension_IsAbsolute(n.Ptr())
	return bool(result_)
}

func (n NSCollectionLayoutDimension) IsEstimated() bool {
	result_ := C.C_NSCollectionLayoutDimension_IsEstimated(n.Ptr())
	return bool(result_)
}

func (n NSCollectionLayoutDimension) IsFractionalHeight() bool {
	result_ := C.C_NSCollectionLayoutDimension_IsFractionalHeight(n.Ptr())
	return bool(result_)
}

func (n NSCollectionLayoutDimension) IsFractionalWidth() bool {
	result_ := C.C_NSCollectionLayoutDimension_IsFractionalWidth(n.Ptr())
	return bool(result_)
}

type CollectionLayoutGroup interface {
	CollectionLayoutItem
	VisualDescription() string
	Subitems() []CollectionLayoutItem
	SetSupplementaryItems(value []CollectionLayoutSupplementaryItem)
	InterItemSpacing() CollectionLayoutSpacing
	SetInterItemSpacing(value CollectionLayoutSpacing)
}

type NSCollectionLayoutGroup struct {
	NSCollectionLayoutItem
}

func MakeCollectionLayoutGroup(ptr unsafe.Pointer) NSCollectionLayoutGroup {
	return NSCollectionLayoutGroup{
		NSCollectionLayoutItem: MakeCollectionLayoutItem(ptr),
	}
}

func CollectionLayoutGroup_HorizontalGroupWithLayoutSize_Subitems(layoutSize CollectionLayoutSize, subitems []CollectionLayoutItem) NSCollectionLayoutGroup {
	var cSubitems C.Array
	if len(subitems) > 0 {
		cSubitemsData := make([]unsafe.Pointer, len(subitems))
		for idx, v := range subitems {
			cSubitemsData[idx] = objc.ExtractPtr(v)
		}
		cSubitems.data = unsafe.Pointer(&cSubitemsData[0])
		cSubitems.len = C.int(len(subitems))
	}
	result_ := C.C_NSCollectionLayoutGroup_CollectionLayoutGroup_HorizontalGroupWithLayoutSize_Subitems(objc.ExtractPtr(layoutSize), cSubitems)
	return MakeCollectionLayoutGroup(result_)
}

func CollectionLayoutGroup_HorizontalGroupWithLayoutSize_Subitem_Count(layoutSize CollectionLayoutSize, subitem CollectionLayoutItem, count int) NSCollectionLayoutGroup {
	result_ := C.C_NSCollectionLayoutGroup_CollectionLayoutGroup_HorizontalGroupWithLayoutSize_Subitem_Count(objc.ExtractPtr(layoutSize), objc.ExtractPtr(subitem), C.int(count))
	return MakeCollectionLayoutGroup(result_)
}

func CollectionLayoutGroup_VerticalGroupWithLayoutSize_Subitems(layoutSize CollectionLayoutSize, subitems []CollectionLayoutItem) NSCollectionLayoutGroup {
	var cSubitems C.Array
	if len(subitems) > 0 {
		cSubitemsData := make([]unsafe.Pointer, len(subitems))
		for idx, v := range subitems {
			cSubitemsData[idx] = objc.ExtractPtr(v)
		}
		cSubitems.data = unsafe.Pointer(&cSubitemsData[0])
		cSubitems.len = C.int(len(subitems))
	}
	result_ := C.C_NSCollectionLayoutGroup_CollectionLayoutGroup_VerticalGroupWithLayoutSize_Subitems(objc.ExtractPtr(layoutSize), cSubitems)
	return MakeCollectionLayoutGroup(result_)
}

func CollectionLayoutGroup_VerticalGroupWithLayoutSize_Subitem_Count(layoutSize CollectionLayoutSize, subitem CollectionLayoutItem, count int) NSCollectionLayoutGroup {
	result_ := C.C_NSCollectionLayoutGroup_CollectionLayoutGroup_VerticalGroupWithLayoutSize_Subitem_Count(objc.ExtractPtr(layoutSize), objc.ExtractPtr(subitem), C.int(count))
	return MakeCollectionLayoutGroup(result_)
}

func CollectionLayoutGroup_ItemWithLayoutSize(layoutSize CollectionLayoutSize) NSCollectionLayoutGroup {
	result_ := C.C_NSCollectionLayoutGroup_CollectionLayoutGroup_ItemWithLayoutSize(objc.ExtractPtr(layoutSize))
	return MakeCollectionLayoutGroup(result_)
}

func CollectionLayoutGroup_ItemWithLayoutSize_SupplementaryItems(layoutSize CollectionLayoutSize, supplementaryItems []CollectionLayoutSupplementaryItem) NSCollectionLayoutGroup {
	var cSupplementaryItems C.Array
	if len(supplementaryItems) > 0 {
		cSupplementaryItemsData := make([]unsafe.Pointer, len(supplementaryItems))
		for idx, v := range supplementaryItems {
			cSupplementaryItemsData[idx] = objc.ExtractPtr(v)
		}
		cSupplementaryItems.data = unsafe.Pointer(&cSupplementaryItemsData[0])
		cSupplementaryItems.len = C.int(len(supplementaryItems))
	}
	result_ := C.C_NSCollectionLayoutGroup_CollectionLayoutGroup_ItemWithLayoutSize_SupplementaryItems(objc.ExtractPtr(layoutSize), cSupplementaryItems)
	return MakeCollectionLayoutGroup(result_)
}

func AllocCollectionLayoutGroup() NSCollectionLayoutGroup {
	result_ := C.C_NSCollectionLayoutGroup_AllocCollectionLayoutGroup()
	return MakeCollectionLayoutGroup(result_)
}

func (n NSCollectionLayoutGroup) Autorelease() NSCollectionLayoutGroup {
	result_ := C.C_NSCollectionLayoutGroup_Autorelease(n.Ptr())
	return MakeCollectionLayoutGroup(result_)
}

func (n NSCollectionLayoutGroup) Retain() NSCollectionLayoutGroup {
	result_ := C.C_NSCollectionLayoutGroup_Retain(n.Ptr())
	return MakeCollectionLayoutGroup(result_)
}

func (n NSCollectionLayoutGroup) VisualDescription() string {
	result_ := C.C_NSCollectionLayoutGroup_VisualDescription(n.Ptr())
	return foundation.MakeString(result_).String()
}

func (n NSCollectionLayoutGroup) Subitems() []CollectionLayoutItem {
	result_ := C.C_NSCollectionLayoutGroup_Subitems(n.Ptr())
	if result_.len > 0 {
		defer C.free(result_.data)
	}
	result_Slice := unsafe.Slice((*unsafe.Pointer)(result_.data), int(result_.len))
	var goResult_ = make([]CollectionLayoutItem, len(result_Slice))
	for idx, r := range result_Slice {
		goResult_[idx] = MakeCollectionLayoutItem(r)
	}
	return goResult_
}

func (n NSCollectionLayoutGroup) SetSupplementaryItems(value []CollectionLayoutSupplementaryItem) {
	var cValue C.Array
	if len(value) > 0 {
		cValueData := make([]unsafe.Pointer, len(value))
		for idx, v := range value {
			cValueData[idx] = objc.ExtractPtr(v)
		}
		cValue.data = unsafe.Pointer(&cValueData[0])
		cValue.len = C.int(len(value))
	}
	C.C_NSCollectionLayoutGroup_SetSupplementaryItems(n.Ptr(), cValue)
}

func (n NSCollectionLayoutGroup) InterItemSpacing() CollectionLayoutSpacing {
	result_ := C.C_NSCollectionLayoutGroup_InterItemSpacing(n.Ptr())
	return MakeCollectionLayoutSpacing(result_)
}

func (n NSCollectionLayoutGroup) SetInterItemSpacing(value CollectionLayoutSpacing) {
	C.C_NSCollectionLayoutGroup_SetInterItemSpacing(n.Ptr(), objc.ExtractPtr(value))
}

type CollectionLayoutDecorationItem interface {
	CollectionLayoutItem
	ElementKind() string
	ZIndex() int
	SetZIndex(value int)
}

type NSCollectionLayoutDecorationItem struct {
	NSCollectionLayoutItem
}

func MakeCollectionLayoutDecorationItem(ptr unsafe.Pointer) NSCollectionLayoutDecorationItem {
	return NSCollectionLayoutDecorationItem{
		NSCollectionLayoutItem: MakeCollectionLayoutItem(ptr),
	}
}

func CollectionLayoutDecorationItem_BackgroundDecorationItemWithElementKind(elementKind string) NSCollectionLayoutDecorationItem {
	result_ := C.C_NSCollectionLayoutDecorationItem_CollectionLayoutDecorationItem_BackgroundDecorationItemWithElementKind(foundation.NewString(elementKind).Ptr())
	return MakeCollectionLayoutDecorationItem(result_)
}

func CollectionLayoutDecorationItem_ItemWithLayoutSize(layoutSize CollectionLayoutSize) NSCollectionLayoutDecorationItem {
	result_ := C.C_NSCollectionLayoutDecorationItem_CollectionLayoutDecorationItem_ItemWithLayoutSize(objc.ExtractPtr(layoutSize))
	return MakeCollectionLayoutDecorationItem(result_)
}

func CollectionLayoutDecorationItem_ItemWithLayoutSize_SupplementaryItems(layoutSize CollectionLayoutSize, supplementaryItems []CollectionLayoutSupplementaryItem) NSCollectionLayoutDecorationItem {
	var cSupplementaryItems C.Array
	if len(supplementaryItems) > 0 {
		cSupplementaryItemsData := make([]unsafe.Pointer, len(supplementaryItems))
		for idx, v := range supplementaryItems {
			cSupplementaryItemsData[idx] = objc.ExtractPtr(v)
		}
		cSupplementaryItems.data = unsafe.Pointer(&cSupplementaryItemsData[0])
		cSupplementaryItems.len = C.int(len(supplementaryItems))
	}
	result_ := C.C_NSCollectionLayoutDecorationItem_CollectionLayoutDecorationItem_ItemWithLayoutSize_SupplementaryItems(objc.ExtractPtr(layoutSize), cSupplementaryItems)
	return MakeCollectionLayoutDecorationItem(result_)
}

func AllocCollectionLayoutDecorationItem() NSCollectionLayoutDecorationItem {
	result_ := C.C_NSCollectionLayoutDecorationItem_AllocCollectionLayoutDecorationItem()
	return MakeCollectionLayoutDecorationItem(result_)
}

func (n NSCollectionLayoutDecorationItem) Autorelease() NSCollectionLayoutDecorationItem {
	result_ := C.C_NSCollectionLayoutDecorationItem_Autorelease(n.Ptr())
	return MakeCollectionLayoutDecorationItem(result_)
}

func (n NSCollectionLayoutDecorationItem) Retain() NSCollectionLayoutDecorationItem {
	result_ := C.C_NSCollectionLayoutDecorationItem_Retain(n.Ptr())
	return MakeCollectionLayoutDecorationItem(result_)
}

func (n NSCollectionLayoutDecorationItem) ElementKind() string {
	result_ := C.C_NSCollectionLayoutDecorationItem_ElementKind(n.Ptr())
	return foundation.MakeString(result_).String()
}

func (n NSCollectionLayoutDecorationItem) ZIndex() int {
	result_ := C.C_NSCollectionLayoutDecorationItem_ZIndex(n.Ptr())
	return int(result_)
}

func (n NSCollectionLayoutDecorationItem) SetZIndex(value int) {
	C.C_NSCollectionLayoutDecorationItem_SetZIndex(n.Ptr(), C.int(value))
}
