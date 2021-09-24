#import "collection_view_layout_attributes.h"
#import <AppKit/NSCollectionViewLayout.h>
#import <Foundation/NSArray.h>
#import <Foundation/NSDictionary.h>

void* C_CollectionViewLayoutAttributes_Alloc() {
    return [NSCollectionViewLayoutAttributes alloc];
}

void* C_NSCollectionViewLayoutAttributes_CollectionViewLayoutAttributes_LayoutAttributesForItemWithIndexPath(void* indexPath) {
    NSCollectionViewLayoutAttributes* result_ = [NSCollectionViewLayoutAttributes layoutAttributesForItemWithIndexPath:(NSIndexPath*)indexPath];
    return result_;
}

void* C_NSCollectionViewLayoutAttributes_CollectionViewLayoutAttributes_LayoutAttributesForSupplementaryViewOfKind_WithIndexPath(void* elementKind, void* indexPath) {
    NSCollectionViewLayoutAttributes* result_ = [NSCollectionViewLayoutAttributes layoutAttributesForSupplementaryViewOfKind:(NSString*)elementKind withIndexPath:(NSIndexPath*)indexPath];
    return result_;
}

void* C_NSCollectionViewLayoutAttributes_CollectionViewLayoutAttributes_LayoutAttributesForDecorationViewOfKind_WithIndexPath(void* decorationViewKind, void* indexPath) {
    NSCollectionViewLayoutAttributes* result_ = [NSCollectionViewLayoutAttributes layoutAttributesForDecorationViewOfKind:(NSString*)decorationViewKind withIndexPath:(NSIndexPath*)indexPath];
    return result_;
}

void* C_NSCollectionViewLayoutAttributes_CollectionViewLayoutAttributes_LayoutAttributesForInterItemGapBeforeIndexPath(void* indexPath) {
    NSCollectionViewLayoutAttributes* result_ = [NSCollectionViewLayoutAttributes layoutAttributesForInterItemGapBeforeIndexPath:(NSIndexPath*)indexPath];
    return result_;
}

void* C_NSCollectionViewLayoutAttributes_AllocCollectionViewLayoutAttributes() {
    NSCollectionViewLayoutAttributes* result_ = [NSCollectionViewLayoutAttributes alloc];
    return result_;
}

void* C_NSCollectionViewLayoutAttributes_Init(void* ptr) {
    NSCollectionViewLayoutAttributes* nSCollectionViewLayoutAttributes = (NSCollectionViewLayoutAttributes*)ptr;
    NSCollectionViewLayoutAttributes* result_ = [nSCollectionViewLayoutAttributes init];
    return result_;
}

void* C_NSCollectionViewLayoutAttributes_NewCollectionViewLayoutAttributes() {
    NSCollectionViewLayoutAttributes* result_ = [NSCollectionViewLayoutAttributes new];
    return result_;
}

void* C_NSCollectionViewLayoutAttributes_Autorelease(void* ptr) {
    NSCollectionViewLayoutAttributes* nSCollectionViewLayoutAttributes = (NSCollectionViewLayoutAttributes*)ptr;
    NSCollectionViewLayoutAttributes* result_ = [nSCollectionViewLayoutAttributes autorelease];
    return result_;
}

void* C_NSCollectionViewLayoutAttributes_Retain(void* ptr) {
    NSCollectionViewLayoutAttributes* nSCollectionViewLayoutAttributes = (NSCollectionViewLayoutAttributes*)ptr;
    NSCollectionViewLayoutAttributes* result_ = [nSCollectionViewLayoutAttributes retain];
    return result_;
}

int C_NSCollectionViewLayoutAttributes_RepresentedElementCategory(void* ptr) {
    NSCollectionViewLayoutAttributes* nSCollectionViewLayoutAttributes = (NSCollectionViewLayoutAttributes*)ptr;
    NSCollectionElementCategory result_ = [nSCollectionViewLayoutAttributes representedElementCategory];
    return result_;
}

void* C_NSCollectionViewLayoutAttributes_IndexPath(void* ptr) {
    NSCollectionViewLayoutAttributes* nSCollectionViewLayoutAttributes = (NSCollectionViewLayoutAttributes*)ptr;
    NSIndexPath* result_ = [nSCollectionViewLayoutAttributes indexPath];
    return result_;
}

void C_NSCollectionViewLayoutAttributes_SetIndexPath(void* ptr, void* value) {
    NSCollectionViewLayoutAttributes* nSCollectionViewLayoutAttributes = (NSCollectionViewLayoutAttributes*)ptr;
    [nSCollectionViewLayoutAttributes setIndexPath:(NSIndexPath*)value];
}

void* C_NSCollectionViewLayoutAttributes_RepresentedElementKind(void* ptr) {
    NSCollectionViewLayoutAttributes* nSCollectionViewLayoutAttributes = (NSCollectionViewLayoutAttributes*)ptr;
    NSString* result_ = [nSCollectionViewLayoutAttributes representedElementKind];
    return result_;
}

CGRect C_NSCollectionViewLayoutAttributes_Frame(void* ptr) {
    NSCollectionViewLayoutAttributes* nSCollectionViewLayoutAttributes = (NSCollectionViewLayoutAttributes*)ptr;
    NSRect result_ = [nSCollectionViewLayoutAttributes frame];
    return result_;
}

void C_NSCollectionViewLayoutAttributes_SetFrame(void* ptr, CGRect value) {
    NSCollectionViewLayoutAttributes* nSCollectionViewLayoutAttributes = (NSCollectionViewLayoutAttributes*)ptr;
    [nSCollectionViewLayoutAttributes setFrame:value];
}

CGSize C_NSCollectionViewLayoutAttributes_Size(void* ptr) {
    NSCollectionViewLayoutAttributes* nSCollectionViewLayoutAttributes = (NSCollectionViewLayoutAttributes*)ptr;
    NSSize result_ = [nSCollectionViewLayoutAttributes size];
    return result_;
}

void C_NSCollectionViewLayoutAttributes_SetSize(void* ptr, CGSize value) {
    NSCollectionViewLayoutAttributes* nSCollectionViewLayoutAttributes = (NSCollectionViewLayoutAttributes*)ptr;
    [nSCollectionViewLayoutAttributes setSize:value];
}

double C_NSCollectionViewLayoutAttributes_Alpha(void* ptr) {
    NSCollectionViewLayoutAttributes* nSCollectionViewLayoutAttributes = (NSCollectionViewLayoutAttributes*)ptr;
    CGFloat result_ = [nSCollectionViewLayoutAttributes alpha];
    return result_;
}

void C_NSCollectionViewLayoutAttributes_SetAlpha(void* ptr, double value) {
    NSCollectionViewLayoutAttributes* nSCollectionViewLayoutAttributes = (NSCollectionViewLayoutAttributes*)ptr;
    [nSCollectionViewLayoutAttributes setAlpha:value];
}

bool C_NSCollectionViewLayoutAttributes_IsHidden(void* ptr) {
    NSCollectionViewLayoutAttributes* nSCollectionViewLayoutAttributes = (NSCollectionViewLayoutAttributes*)ptr;
    BOOL result_ = [nSCollectionViewLayoutAttributes isHidden];
    return result_;
}

void C_NSCollectionViewLayoutAttributes_SetHidden(void* ptr, bool value) {
    NSCollectionViewLayoutAttributes* nSCollectionViewLayoutAttributes = (NSCollectionViewLayoutAttributes*)ptr;
    [nSCollectionViewLayoutAttributes setHidden:value];
}

int C_NSCollectionViewLayoutAttributes_ZIndex(void* ptr) {
    NSCollectionViewLayoutAttributes* nSCollectionViewLayoutAttributes = (NSCollectionViewLayoutAttributes*)ptr;
    NSInteger result_ = [nSCollectionViewLayoutAttributes zIndex];
    return result_;
}

void C_NSCollectionViewLayoutAttributes_SetZIndex(void* ptr, int value) {
    NSCollectionViewLayoutAttributes* nSCollectionViewLayoutAttributes = (NSCollectionViewLayoutAttributes*)ptr;
    [nSCollectionViewLayoutAttributes setZIndex:value];
}
