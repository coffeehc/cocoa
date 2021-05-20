#import <Appkit/Appkit.h>
#import "collection_layout_size.h"

void* C_CollectionLayoutSize_Alloc() {
    return [NSCollectionLayoutSize alloc];
}

void* C_NSCollectionLayoutSize_CollectionLayoutSize_SizeWithWidthDimension_HeightDimension(void* width, void* height) {
    NSCollectionLayoutSize* result_ = [NSCollectionLayoutSize sizeWithWidthDimension:(NSCollectionLayoutDimension*)width heightDimension:(NSCollectionLayoutDimension*)height];
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
