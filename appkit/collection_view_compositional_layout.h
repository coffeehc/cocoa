#import <AppKit/NSCollectionViewCompositionalLayout.h>
#import <Foundation/NSGeometry.h>
#import <Foundation/NSRange.h>
#import <stdbool.h>
#import <stdint.h>
#import <stdlib.h>
#import <utils.h>

void* C_CollectionViewCompositionalLayout_Alloc();

void* C_NSCollectionViewCompositionalLayout_InitWithSection(void* ptr, void* section);
void* C_NSCollectionViewCompositionalLayout_InitWithSection_Configuration(void* ptr, void* section, void* configuration);
void* C_NSCollectionViewCompositionalLayout_AllocCollectionViewCompositionalLayout();
void* C_NSCollectionViewCompositionalLayout_Autorelease(void* ptr);
void* C_NSCollectionViewCompositionalLayout_Retain(void* ptr);
void* C_NSCollectionViewCompositionalLayout_Configuration(void* ptr);
void C_NSCollectionViewCompositionalLayout_SetConfiguration(void* ptr, void* value);

void* C_CollectionViewCompositionalLayoutConfiguration_Alloc();

void* C_NSCollectionViewCompositionalLayoutConfiguration_AllocCollectionViewCompositionalLayoutConfiguration();
void* C_NSCollectionViewCompositionalLayoutConfiguration_Init(void* ptr);
void* C_NSCollectionViewCompositionalLayoutConfiguration_NewCollectionViewCompositionalLayoutConfiguration();
void* C_NSCollectionViewCompositionalLayoutConfiguration_Autorelease(void* ptr);
void* C_NSCollectionViewCompositionalLayoutConfiguration_Retain(void* ptr);
int C_NSCollectionViewCompositionalLayoutConfiguration_ScrollDirection(void* ptr);
void C_NSCollectionViewCompositionalLayoutConfiguration_SetScrollDirection(void* ptr, int value);
double C_NSCollectionViewCompositionalLayoutConfiguration_InterSectionSpacing(void* ptr);
void C_NSCollectionViewCompositionalLayoutConfiguration_SetInterSectionSpacing(void* ptr, double value);
Array C_NSCollectionViewCompositionalLayoutConfiguration_BoundarySupplementaryItems(void* ptr);
void C_NSCollectionViewCompositionalLayoutConfiguration_SetBoundarySupplementaryItems(void* ptr, Array value);

void* C_CollectionLayoutItem_Alloc();

void* C_NSCollectionLayoutItem_CollectionLayoutItem_ItemWithLayoutSize(void* layoutSize);
void* C_NSCollectionLayoutItem_CollectionLayoutItem_ItemWithLayoutSize_SupplementaryItems(void* layoutSize, Array supplementaryItems);
void* C_NSCollectionLayoutItem_AllocCollectionLayoutItem();
void* C_NSCollectionLayoutItem_Autorelease(void* ptr);
void* C_NSCollectionLayoutItem_Retain(void* ptr);
void* C_NSCollectionLayoutItem_LayoutSize(void* ptr);
Array C_NSCollectionLayoutItem_SupplementaryItems(void* ptr);
void* C_NSCollectionLayoutItem_EdgeSpacing(void* ptr);
void C_NSCollectionLayoutItem_SetEdgeSpacing(void* ptr, void* value);
NSDirectionalEdgeInsets C_NSCollectionLayoutItem_ContentInsets(void* ptr);
void C_NSCollectionLayoutItem_SetContentInsets(void* ptr, NSDirectionalEdgeInsets value);

void* C_CollectionLayoutBoundarySupplementaryItem_Alloc();

void* C_NSCollectionLayoutBoundarySupplementaryItem_CollectionLayoutBoundarySupplementaryItem_BoundarySupplementaryItemWithLayoutSize_ElementKind_Alignment(void* layoutSize, void* elementKind, int alignment);
void* C_NSCollectionLayoutBoundarySupplementaryItem_CollectionLayoutBoundarySupplementaryItem_BoundarySupplementaryItemWithLayoutSize_ElementKind_Alignment_AbsoluteOffset(void* layoutSize, void* elementKind, int alignment, CGPoint absoluteOffset);
void* C_NSCollectionLayoutBoundarySupplementaryItem_CollectionLayoutBoundarySupplementaryItem_SupplementaryItemWithLayoutSize_ElementKind_ContainerAnchor(void* layoutSize, void* elementKind, void* containerAnchor);
void* C_NSCollectionLayoutBoundarySupplementaryItem_CollectionLayoutBoundarySupplementaryItem_SupplementaryItemWithLayoutSize_ElementKind_ContainerAnchor_ItemAnchor(void* layoutSize, void* elementKind, void* containerAnchor, void* itemAnchor);
void* C_NSCollectionLayoutBoundarySupplementaryItem_CollectionLayoutBoundarySupplementaryItem_ItemWithLayoutSize(void* layoutSize);
void* C_NSCollectionLayoutBoundarySupplementaryItem_CollectionLayoutBoundarySupplementaryItem_ItemWithLayoutSize_SupplementaryItems(void* layoutSize, Array supplementaryItems);
void* C_NSCollectionLayoutBoundarySupplementaryItem_AllocCollectionLayoutBoundarySupplementaryItem();
void* C_NSCollectionLayoutBoundarySupplementaryItem_Autorelease(void* ptr);
void* C_NSCollectionLayoutBoundarySupplementaryItem_Retain(void* ptr);
bool C_NSCollectionLayoutBoundarySupplementaryItem_PinToVisibleBounds(void* ptr);
void C_NSCollectionLayoutBoundarySupplementaryItem_SetPinToVisibleBounds(void* ptr, bool value);
CGPoint C_NSCollectionLayoutBoundarySupplementaryItem_Offset(void* ptr);
int C_NSCollectionLayoutBoundarySupplementaryItem_Alignment(void* ptr);
bool C_NSCollectionLayoutBoundarySupplementaryItem_ExtendsBoundary(void* ptr);
void C_NSCollectionLayoutBoundarySupplementaryItem_SetExtendsBoundary(void* ptr, bool value);

void* C_CollectionLayoutSpacing_Alloc();

void* C_NSCollectionLayoutSpacing_CollectionLayoutSpacing_FixedSpacing(double fixedSpacing);
void* C_NSCollectionLayoutSpacing_CollectionLayoutSpacing_FlexibleSpacing(double flexibleSpacing);
void* C_NSCollectionLayoutSpacing_AllocCollectionLayoutSpacing();
void* C_NSCollectionLayoutSpacing_Autorelease(void* ptr);
void* C_NSCollectionLayoutSpacing_Retain(void* ptr);
double C_NSCollectionLayoutSpacing_Spacing(void* ptr);
bool C_NSCollectionLayoutSpacing_IsFixedSpacing(void* ptr);
bool C_NSCollectionLayoutSpacing_IsFlexibleSpacing(void* ptr);

void* C_CollectionLayoutSection_Alloc();

void* C_NSCollectionLayoutSection_CollectionLayoutSection_SectionWithGroup(void* group);
void* C_NSCollectionLayoutSection_AllocCollectionLayoutSection();
void* C_NSCollectionLayoutSection_Autorelease(void* ptr);
void* C_NSCollectionLayoutSection_Retain(void* ptr);
int C_NSCollectionLayoutSection_OrthogonalScrollingBehavior(void* ptr);
void C_NSCollectionLayoutSection_SetOrthogonalScrollingBehavior(void* ptr, int value);
double C_NSCollectionLayoutSection_InterGroupSpacing(void* ptr);
void C_NSCollectionLayoutSection_SetInterGroupSpacing(void* ptr, double value);
NSDirectionalEdgeInsets C_NSCollectionLayoutSection_ContentInsets(void* ptr);
void C_NSCollectionLayoutSection_SetContentInsets(void* ptr, NSDirectionalEdgeInsets value);
bool C_NSCollectionLayoutSection_SupplementariesFollowContentInsets(void* ptr);
void C_NSCollectionLayoutSection_SetSupplementariesFollowContentInsets(void* ptr, bool value);
Array C_NSCollectionLayoutSection_BoundarySupplementaryItems(void* ptr);
void C_NSCollectionLayoutSection_SetBoundarySupplementaryItems(void* ptr, Array value);
Array C_NSCollectionLayoutSection_DecorationItems(void* ptr);
void C_NSCollectionLayoutSection_SetDecorationItems(void* ptr, Array value);

void* C_CollectionLayoutGroupCustomItem_Alloc();

void* C_NSCollectionLayoutGroupCustomItem_CollectionLayoutGroupCustomItem_CustomItemWithFrame(CGRect frame);
void* C_NSCollectionLayoutGroupCustomItem_CollectionLayoutGroupCustomItem_CustomItemWithFrame_ZIndex(CGRect frame, int zIndex);
void* C_NSCollectionLayoutGroupCustomItem_AllocCollectionLayoutGroupCustomItem();
void* C_NSCollectionLayoutGroupCustomItem_Autorelease(void* ptr);
void* C_NSCollectionLayoutGroupCustomItem_Retain(void* ptr);
CGRect C_NSCollectionLayoutGroupCustomItem_Frame(void* ptr);
int C_NSCollectionLayoutGroupCustomItem_ZIndex(void* ptr);

void* C_CollectionLayoutSupplementaryItem_Alloc();

void* C_NSCollectionLayoutSupplementaryItem_CollectionLayoutSupplementaryItem_SupplementaryItemWithLayoutSize_ElementKind_ContainerAnchor(void* layoutSize, void* elementKind, void* containerAnchor);
void* C_NSCollectionLayoutSupplementaryItem_CollectionLayoutSupplementaryItem_SupplementaryItemWithLayoutSize_ElementKind_ContainerAnchor_ItemAnchor(void* layoutSize, void* elementKind, void* containerAnchor, void* itemAnchor);
void* C_NSCollectionLayoutSupplementaryItem_CollectionLayoutSupplementaryItem_ItemWithLayoutSize(void* layoutSize);
void* C_NSCollectionLayoutSupplementaryItem_CollectionLayoutSupplementaryItem_ItemWithLayoutSize_SupplementaryItems(void* layoutSize, Array supplementaryItems);
void* C_NSCollectionLayoutSupplementaryItem_AllocCollectionLayoutSupplementaryItem();
void* C_NSCollectionLayoutSupplementaryItem_Autorelease(void* ptr);
void* C_NSCollectionLayoutSupplementaryItem_Retain(void* ptr);
void* C_NSCollectionLayoutSupplementaryItem_ItemAnchor(void* ptr);
void* C_NSCollectionLayoutSupplementaryItem_ContainerAnchor(void* ptr);
void* C_NSCollectionLayoutSupplementaryItem_ElementKind(void* ptr);
int C_NSCollectionLayoutSupplementaryItem_ZIndex(void* ptr);
void C_NSCollectionLayoutSupplementaryItem_SetZIndex(void* ptr, int value);

void* C_CollectionLayoutSize_Alloc();

void* C_NSCollectionLayoutSize_CollectionLayoutSize_SizeWithWidthDimension_HeightDimension(void* width, void* height);
void* C_NSCollectionLayoutSize_AllocCollectionLayoutSize();
void* C_NSCollectionLayoutSize_Autorelease(void* ptr);
void* C_NSCollectionLayoutSize_Retain(void* ptr);
void* C_NSCollectionLayoutSize_WidthDimension(void* ptr);
void* C_NSCollectionLayoutSize_HeightDimension(void* ptr);

void* C_CollectionLayoutEdgeSpacing_Alloc();

void* C_NSCollectionLayoutEdgeSpacing_CollectionLayoutEdgeSpacing_SpacingForLeading_Top_Trailing_Bottom(void* leading, void* top, void* trailing, void* bottom);
void* C_NSCollectionLayoutEdgeSpacing_AllocCollectionLayoutEdgeSpacing();
void* C_NSCollectionLayoutEdgeSpacing_Autorelease(void* ptr);
void* C_NSCollectionLayoutEdgeSpacing_Retain(void* ptr);
void* C_NSCollectionLayoutEdgeSpacing_Leading(void* ptr);
void* C_NSCollectionLayoutEdgeSpacing_Top(void* ptr);
void* C_NSCollectionLayoutEdgeSpacing_Trailing(void* ptr);
void* C_NSCollectionLayoutEdgeSpacing_Bottom(void* ptr);

void* C_CollectionLayoutAnchor_Alloc();

void* C_NSCollectionLayoutAnchor_CollectionLayoutAnchor_LayoutAnchorWithEdges(unsigned int edges);
void* C_NSCollectionLayoutAnchor_CollectionLayoutAnchor_LayoutAnchorWithEdges_AbsoluteOffset(unsigned int edges, CGPoint absoluteOffset);
void* C_NSCollectionLayoutAnchor_CollectionLayoutAnchor_LayoutAnchorWithEdges_FractionalOffset(unsigned int edges, CGPoint fractionalOffset);
void* C_NSCollectionLayoutAnchor_AllocCollectionLayoutAnchor();
void* C_NSCollectionLayoutAnchor_Autorelease(void* ptr);
void* C_NSCollectionLayoutAnchor_Retain(void* ptr);
unsigned int C_NSCollectionLayoutAnchor_Edges(void* ptr);
CGPoint C_NSCollectionLayoutAnchor_Offset(void* ptr);
bool C_NSCollectionLayoutAnchor_IsAbsoluteOffset(void* ptr);
bool C_NSCollectionLayoutAnchor_IsFractionalOffset(void* ptr);

void* C_CollectionLayoutDimension_Alloc();

void* C_NSCollectionLayoutDimension_CollectionLayoutDimension_AbsoluteDimension(double absoluteDimension);
void* C_NSCollectionLayoutDimension_CollectionLayoutDimension_EstimatedDimension(double estimatedDimension);
void* C_NSCollectionLayoutDimension_CollectionLayoutDimension_FractionalHeightDimension(double fractionalHeight);
void* C_NSCollectionLayoutDimension_CollectionLayoutDimension_FractionalWidthDimension(double fractionalWidth);
void* C_NSCollectionLayoutDimension_AllocCollectionLayoutDimension();
void* C_NSCollectionLayoutDimension_Autorelease(void* ptr);
void* C_NSCollectionLayoutDimension_Retain(void* ptr);
double C_NSCollectionLayoutDimension_Dimension(void* ptr);
bool C_NSCollectionLayoutDimension_IsAbsolute(void* ptr);
bool C_NSCollectionLayoutDimension_IsEstimated(void* ptr);
bool C_NSCollectionLayoutDimension_IsFractionalHeight(void* ptr);
bool C_NSCollectionLayoutDimension_IsFractionalWidth(void* ptr);

void* C_CollectionLayoutGroup_Alloc();

void* C_NSCollectionLayoutGroup_CollectionLayoutGroup_HorizontalGroupWithLayoutSize_Subitems(void* layoutSize, Array subitems);
void* C_NSCollectionLayoutGroup_CollectionLayoutGroup_HorizontalGroupWithLayoutSize_Subitem_Count(void* layoutSize, void* subitem, int count);
void* C_NSCollectionLayoutGroup_CollectionLayoutGroup_VerticalGroupWithLayoutSize_Subitems(void* layoutSize, Array subitems);
void* C_NSCollectionLayoutGroup_CollectionLayoutGroup_VerticalGroupWithLayoutSize_Subitem_Count(void* layoutSize, void* subitem, int count);
void* C_NSCollectionLayoutGroup_CollectionLayoutGroup_ItemWithLayoutSize(void* layoutSize);
void* C_NSCollectionLayoutGroup_CollectionLayoutGroup_ItemWithLayoutSize_SupplementaryItems(void* layoutSize, Array supplementaryItems);
void* C_NSCollectionLayoutGroup_AllocCollectionLayoutGroup();
void* C_NSCollectionLayoutGroup_Autorelease(void* ptr);
void* C_NSCollectionLayoutGroup_Retain(void* ptr);
void* C_NSCollectionLayoutGroup_VisualDescription(void* ptr);
Array C_NSCollectionLayoutGroup_Subitems(void* ptr);
void C_NSCollectionLayoutGroup_SetSupplementaryItems(void* ptr, Array value);
void* C_NSCollectionLayoutGroup_InterItemSpacing(void* ptr);
void C_NSCollectionLayoutGroup_SetInterItemSpacing(void* ptr, void* value);

void* C_CollectionLayoutDecorationItem_Alloc();

void* C_NSCollectionLayoutDecorationItem_CollectionLayoutDecorationItem_BackgroundDecorationItemWithElementKind(void* elementKind);
void* C_NSCollectionLayoutDecorationItem_CollectionLayoutDecorationItem_ItemWithLayoutSize(void* layoutSize);
void* C_NSCollectionLayoutDecorationItem_CollectionLayoutDecorationItem_ItemWithLayoutSize_SupplementaryItems(void* layoutSize, Array supplementaryItems);
void* C_NSCollectionLayoutDecorationItem_AllocCollectionLayoutDecorationItem();
void* C_NSCollectionLayoutDecorationItem_Autorelease(void* ptr);
void* C_NSCollectionLayoutDecorationItem_Retain(void* ptr);
void* C_NSCollectionLayoutDecorationItem_ElementKind(void* ptr);
int C_NSCollectionLayoutDecorationItem_ZIndex(void* ptr);
void C_NSCollectionLayoutDecorationItem_SetZIndex(void* ptr, int value);
