#import <Appkit/Appkit.h>
#import "collection_layout_spacing.h"

void* C_CollectionLayoutSpacing_Alloc() {
    return [NSCollectionLayoutSpacing alloc];
}

void* C_NSCollectionLayoutSpacing_CollectionLayoutSpacing_FixedSpacing(double fixedSpacing) {
    NSCollectionLayoutSpacing* result_ = [NSCollectionLayoutSpacing fixedSpacing:fixedSpacing];
    return result_;
}

void* C_NSCollectionLayoutSpacing_CollectionLayoutSpacing_FlexibleSpacing(double flexibleSpacing) {
    NSCollectionLayoutSpacing* result_ = [NSCollectionLayoutSpacing flexibleSpacing:flexibleSpacing];
    return result_;
}

void* C_NSCollectionLayoutSpacing_AllocCollectionLayoutSpacing() {
    NSCollectionLayoutSpacing* result_ = [NSCollectionLayoutSpacing alloc];
    return result_;
}

void* C_NSCollectionLayoutSpacing_Autorelease(void* ptr) {
    NSCollectionLayoutSpacing* nSCollectionLayoutSpacing = (NSCollectionLayoutSpacing*)ptr;
    NSCollectionLayoutSpacing* result_ = [nSCollectionLayoutSpacing autorelease];
    return result_;
}

void* C_NSCollectionLayoutSpacing_Retain(void* ptr) {
    NSCollectionLayoutSpacing* nSCollectionLayoutSpacing = (NSCollectionLayoutSpacing*)ptr;
    NSCollectionLayoutSpacing* result_ = [nSCollectionLayoutSpacing retain];
    return result_;
}

double C_NSCollectionLayoutSpacing_Spacing(void* ptr) {
    NSCollectionLayoutSpacing* nSCollectionLayoutSpacing = (NSCollectionLayoutSpacing*)ptr;
    CGFloat result_ = [nSCollectionLayoutSpacing spacing];
    return result_;
}

bool C_NSCollectionLayoutSpacing_IsFixedSpacing(void* ptr) {
    NSCollectionLayoutSpacing* nSCollectionLayoutSpacing = (NSCollectionLayoutSpacing*)ptr;
    BOOL result_ = [nSCollectionLayoutSpacing isFixedSpacing];
    return result_;
}

bool C_NSCollectionLayoutSpacing_IsFlexibleSpacing(void* ptr) {
    NSCollectionLayoutSpacing* nSCollectionLayoutSpacing = (NSCollectionLayoutSpacing*)ptr;
    BOOL result_ = [nSCollectionLayoutSpacing isFlexibleSpacing];
    return result_;
}
