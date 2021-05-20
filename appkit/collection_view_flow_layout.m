#import <Appkit/Appkit.h>
#import "collection_view_flow_layout.h"

void* C_CollectionViewFlowLayout_Alloc() {
    return [NSCollectionViewFlowLayout alloc];
}

void* C_NSCollectionViewFlowLayout_Init(void* ptr) {
    NSCollectionViewFlowLayout* nSCollectionViewFlowLayout = (NSCollectionViewFlowLayout*)ptr;
    NSCollectionViewFlowLayout* result_ = [nSCollectionViewFlowLayout init];
    return result_;
}

void C_NSCollectionViewFlowLayout_CollapseSectionAtIndex(void* ptr, unsigned int sectionIndex) {
    NSCollectionViewFlowLayout* nSCollectionViewFlowLayout = (NSCollectionViewFlowLayout*)ptr;
    [nSCollectionViewFlowLayout collapseSectionAtIndex:sectionIndex];
}

void C_NSCollectionViewFlowLayout_ExpandSectionAtIndex(void* ptr, unsigned int sectionIndex) {
    NSCollectionViewFlowLayout* nSCollectionViewFlowLayout = (NSCollectionViewFlowLayout*)ptr;
    [nSCollectionViewFlowLayout expandSectionAtIndex:sectionIndex];
}

bool C_NSCollectionViewFlowLayout_SectionAtIndexIsCollapsed(void* ptr, unsigned int sectionIndex) {
    NSCollectionViewFlowLayout* nSCollectionViewFlowLayout = (NSCollectionViewFlowLayout*)ptr;
    BOOL result_ = [nSCollectionViewFlowLayout sectionAtIndexIsCollapsed:sectionIndex];
    return result_;
}

int C_NSCollectionViewFlowLayout_ScrollDirection(void* ptr) {
    NSCollectionViewFlowLayout* nSCollectionViewFlowLayout = (NSCollectionViewFlowLayout*)ptr;
    NSCollectionViewScrollDirection result_ = [nSCollectionViewFlowLayout scrollDirection];
    return result_;
}

void C_NSCollectionViewFlowLayout_SetScrollDirection(void* ptr, int value) {
    NSCollectionViewFlowLayout* nSCollectionViewFlowLayout = (NSCollectionViewFlowLayout*)ptr;
    [nSCollectionViewFlowLayout setScrollDirection:value];
}

double C_NSCollectionViewFlowLayout_MinimumLineSpacing(void* ptr) {
    NSCollectionViewFlowLayout* nSCollectionViewFlowLayout = (NSCollectionViewFlowLayout*)ptr;
    CGFloat result_ = [nSCollectionViewFlowLayout minimumLineSpacing];
    return result_;
}

void C_NSCollectionViewFlowLayout_SetMinimumLineSpacing(void* ptr, double value) {
    NSCollectionViewFlowLayout* nSCollectionViewFlowLayout = (NSCollectionViewFlowLayout*)ptr;
    [nSCollectionViewFlowLayout setMinimumLineSpacing:value];
}

double C_NSCollectionViewFlowLayout_MinimumInteritemSpacing(void* ptr) {
    NSCollectionViewFlowLayout* nSCollectionViewFlowLayout = (NSCollectionViewFlowLayout*)ptr;
    CGFloat result_ = [nSCollectionViewFlowLayout minimumInteritemSpacing];
    return result_;
}

void C_NSCollectionViewFlowLayout_SetMinimumInteritemSpacing(void* ptr, double value) {
    NSCollectionViewFlowLayout* nSCollectionViewFlowLayout = (NSCollectionViewFlowLayout*)ptr;
    [nSCollectionViewFlowLayout setMinimumInteritemSpacing:value];
}

CGSize C_NSCollectionViewFlowLayout_EstimatedItemSize(void* ptr) {
    NSCollectionViewFlowLayout* nSCollectionViewFlowLayout = (NSCollectionViewFlowLayout*)ptr;
    NSSize result_ = [nSCollectionViewFlowLayout estimatedItemSize];
    return result_;
}

void C_NSCollectionViewFlowLayout_SetEstimatedItemSize(void* ptr, CGSize value) {
    NSCollectionViewFlowLayout* nSCollectionViewFlowLayout = (NSCollectionViewFlowLayout*)ptr;
    [nSCollectionViewFlowLayout setEstimatedItemSize:value];
}

CGSize C_NSCollectionViewFlowLayout_ItemSize(void* ptr) {
    NSCollectionViewFlowLayout* nSCollectionViewFlowLayout = (NSCollectionViewFlowLayout*)ptr;
    NSSize result_ = [nSCollectionViewFlowLayout itemSize];
    return result_;
}

void C_NSCollectionViewFlowLayout_SetItemSize(void* ptr, CGSize value) {
    NSCollectionViewFlowLayout* nSCollectionViewFlowLayout = (NSCollectionViewFlowLayout*)ptr;
    [nSCollectionViewFlowLayout setItemSize:value];
}

NSEdgeInsets C_NSCollectionViewFlowLayout_SectionInset(void* ptr) {
    NSCollectionViewFlowLayout* nSCollectionViewFlowLayout = (NSCollectionViewFlowLayout*)ptr;
    NSEdgeInsets result_ = [nSCollectionViewFlowLayout sectionInset];
    return result_;
}

void C_NSCollectionViewFlowLayout_SetSectionInset(void* ptr, NSEdgeInsets value) {
    NSCollectionViewFlowLayout* nSCollectionViewFlowLayout = (NSCollectionViewFlowLayout*)ptr;
    [nSCollectionViewFlowLayout setSectionInset:value];
}

CGSize C_NSCollectionViewFlowLayout_HeaderReferenceSize(void* ptr) {
    NSCollectionViewFlowLayout* nSCollectionViewFlowLayout = (NSCollectionViewFlowLayout*)ptr;
    NSSize result_ = [nSCollectionViewFlowLayout headerReferenceSize];
    return result_;
}

void C_NSCollectionViewFlowLayout_SetHeaderReferenceSize(void* ptr, CGSize value) {
    NSCollectionViewFlowLayout* nSCollectionViewFlowLayout = (NSCollectionViewFlowLayout*)ptr;
    [nSCollectionViewFlowLayout setHeaderReferenceSize:value];
}

CGSize C_NSCollectionViewFlowLayout_FooterReferenceSize(void* ptr) {
    NSCollectionViewFlowLayout* nSCollectionViewFlowLayout = (NSCollectionViewFlowLayout*)ptr;
    NSSize result_ = [nSCollectionViewFlowLayout footerReferenceSize];
    return result_;
}

void C_NSCollectionViewFlowLayout_SetFooterReferenceSize(void* ptr, CGSize value) {
    NSCollectionViewFlowLayout* nSCollectionViewFlowLayout = (NSCollectionViewFlowLayout*)ptr;
    [nSCollectionViewFlowLayout setFooterReferenceSize:value];
}

bool C_NSCollectionViewFlowLayout_SectionFootersPinToVisibleBounds(void* ptr) {
    NSCollectionViewFlowLayout* nSCollectionViewFlowLayout = (NSCollectionViewFlowLayout*)ptr;
    BOOL result_ = [nSCollectionViewFlowLayout sectionFootersPinToVisibleBounds];
    return result_;
}

void C_NSCollectionViewFlowLayout_SetSectionFootersPinToVisibleBounds(void* ptr, bool value) {
    NSCollectionViewFlowLayout* nSCollectionViewFlowLayout = (NSCollectionViewFlowLayout*)ptr;
    [nSCollectionViewFlowLayout setSectionFootersPinToVisibleBounds:value];
}

bool C_NSCollectionViewFlowLayout_SectionHeadersPinToVisibleBounds(void* ptr) {
    NSCollectionViewFlowLayout* nSCollectionViewFlowLayout = (NSCollectionViewFlowLayout*)ptr;
    BOOL result_ = [nSCollectionViewFlowLayout sectionHeadersPinToVisibleBounds];
    return result_;
}

void C_NSCollectionViewFlowLayout_SetSectionHeadersPinToVisibleBounds(void* ptr, bool value) {
    NSCollectionViewFlowLayout* nSCollectionViewFlowLayout = (NSCollectionViewFlowLayout*)ptr;
    [nSCollectionViewFlowLayout setSectionHeadersPinToVisibleBounds:value];
}
