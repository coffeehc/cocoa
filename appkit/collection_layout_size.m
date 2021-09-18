#import "collection_layout_size.h"
#import <Foundation/NSDictionary.h>
#import <Foundation/NSArray.h>
#import <AppKit/NSCollectionViewCompositionalLayout.h>

void* C_CollectionLayoutSize_Alloc() {
    return [NSCollectionLayoutSize alloc];
}

void* C_NSCollectionLayoutSize_CollectionLayoutSize_SizeWithWidthDimension_HeightDimension(void* width, void* height) {
    NSCollectionLayoutSize* result_ = [NSCollectionLayoutSize sizeWithWidthDimension:(NSCollectionLayoutDimension*)width heightDimension:(NSCollectionLayoutDimension*)height];
    return result_;
}

void* C_NSCollectionLayoutSize_AllocCollectionLayoutSize() {
    NSCollectionLayoutSize* result_ = [NSCollectionLayoutSize alloc];
    return result_;
}

void* C_NSCollectionLayoutSize_Autorelease(void* ptr) {
    NSCollectionLayoutSize* nSCollectionLayoutSize = (NSCollectionLayoutSize*)ptr;
    NSCollectionLayoutSize* result_ = [nSCollectionLayoutSize autorelease];
    return result_;
}

void* C_NSCollectionLayoutSize_Retain(void* ptr) {
    NSCollectionLayoutSize* nSCollectionLayoutSize = (NSCollectionLayoutSize*)ptr;
    NSCollectionLayoutSize* result_ = [nSCollectionLayoutSize retain];
    return result_;
}

void* C_NSCollectionLayoutSize_WidthDimension(void* ptr) {
    NSCollectionLayoutSize* nSCollectionLayoutSize = (NSCollectionLayoutSize*)ptr;
    NSCollectionLayoutDimension* result_ = [nSCollectionLayoutSize widthDimension];
    return result_;
}

void* C_NSCollectionLayoutSize_HeightDimension(void* ptr) {
    NSCollectionLayoutSize* nSCollectionLayoutSize = (NSCollectionLayoutSize*)ptr;
    NSCollectionLayoutDimension* result_ = [nSCollectionLayoutSize heightDimension];
    return result_;
}
