#import <Appkit/Appkit.h>
#import "collection_layout_dimension.h"

void* C_CollectionLayoutDimension_Alloc() {
    return [NSCollectionLayoutDimension alloc];
}

void* C_NSCollectionLayoutDimension_CollectionLayoutDimension_AbsoluteDimension(double absoluteDimension) {
    NSCollectionLayoutDimension* result_ = [NSCollectionLayoutDimension absoluteDimension:absoluteDimension];
    return result_;
}

void* C_NSCollectionLayoutDimension_CollectionLayoutDimension_EstimatedDimension(double estimatedDimension) {
    NSCollectionLayoutDimension* result_ = [NSCollectionLayoutDimension estimatedDimension:estimatedDimension];
    return result_;
}

void* C_NSCollectionLayoutDimension_CollectionLayoutDimension_FractionalHeightDimension(double fractionalHeight) {
    NSCollectionLayoutDimension* result_ = [NSCollectionLayoutDimension fractionalHeightDimension:fractionalHeight];
    return result_;
}

void* C_NSCollectionLayoutDimension_CollectionLayoutDimension_FractionalWidthDimension(double fractionalWidth) {
    NSCollectionLayoutDimension* result_ = [NSCollectionLayoutDimension fractionalWidthDimension:fractionalWidth];
    return result_;
}

void* C_NSCollectionLayoutDimension_AllocCollectionLayoutDimension() {
    NSCollectionLayoutDimension* result_ = [NSCollectionLayoutDimension alloc];
    return result_;
}

void* C_NSCollectionLayoutDimension_Autorelease(void* ptr) {
    NSCollectionLayoutDimension* nSCollectionLayoutDimension = (NSCollectionLayoutDimension*)ptr;
    NSCollectionLayoutDimension* result_ = [nSCollectionLayoutDimension autorelease];
    return result_;
}

void* C_NSCollectionLayoutDimension_Retain(void* ptr) {
    NSCollectionLayoutDimension* nSCollectionLayoutDimension = (NSCollectionLayoutDimension*)ptr;
    NSCollectionLayoutDimension* result_ = [nSCollectionLayoutDimension retain];
    return result_;
}

double C_NSCollectionLayoutDimension_Dimension(void* ptr) {
    NSCollectionLayoutDimension* nSCollectionLayoutDimension = (NSCollectionLayoutDimension*)ptr;
    CGFloat result_ = [nSCollectionLayoutDimension dimension];
    return result_;
}

bool C_NSCollectionLayoutDimension_IsAbsolute(void* ptr) {
    NSCollectionLayoutDimension* nSCollectionLayoutDimension = (NSCollectionLayoutDimension*)ptr;
    BOOL result_ = [nSCollectionLayoutDimension isAbsolute];
    return result_;
}

bool C_NSCollectionLayoutDimension_IsEstimated(void* ptr) {
    NSCollectionLayoutDimension* nSCollectionLayoutDimension = (NSCollectionLayoutDimension*)ptr;
    BOOL result_ = [nSCollectionLayoutDimension isEstimated];
    return result_;
}

bool C_NSCollectionLayoutDimension_IsFractionalHeight(void* ptr) {
    NSCollectionLayoutDimension* nSCollectionLayoutDimension = (NSCollectionLayoutDimension*)ptr;
    BOOL result_ = [nSCollectionLayoutDimension isFractionalHeight];
    return result_;
}

bool C_NSCollectionLayoutDimension_IsFractionalWidth(void* ptr) {
    NSCollectionLayoutDimension* nSCollectionLayoutDimension = (NSCollectionLayoutDimension*)ptr;
    BOOL result_ = [nSCollectionLayoutDimension isFractionalWidth];
    return result_;
}
