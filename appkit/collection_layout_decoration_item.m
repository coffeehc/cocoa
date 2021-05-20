#import <Appkit/Appkit.h>
#import "collection_layout_decoration_item.h"

void* C_CollectionLayoutDecorationItem_Alloc() {
    return [NSCollectionLayoutDecorationItem alloc];
}

void* C_NSCollectionLayoutDecorationItem_CollectionLayoutDecorationItem_BackgroundDecorationItemWithElementKind(void* elementKind) {
    NSCollectionLayoutDecorationItem* result_ = [NSCollectionLayoutDecorationItem backgroundDecorationItemWithElementKind:(NSString*)elementKind];
    return result_;
}

void* C_NSCollectionLayoutDecorationItem_ElementKind(void* ptr) {
    NSCollectionLayoutDecorationItem* nSCollectionLayoutDecorationItem = (NSCollectionLayoutDecorationItem*)ptr;
    NSString* result_ = [nSCollectionLayoutDecorationItem elementKind];
    return result_;
}

int C_NSCollectionLayoutDecorationItem_ZIndex(void* ptr) {
    NSCollectionLayoutDecorationItem* nSCollectionLayoutDecorationItem = (NSCollectionLayoutDecorationItem*)ptr;
    NSInteger result_ = [nSCollectionLayoutDecorationItem zIndex];
    return result_;
}

void C_NSCollectionLayoutDecorationItem_SetZIndex(void* ptr, int value) {
    NSCollectionLayoutDecorationItem* nSCollectionLayoutDecorationItem = (NSCollectionLayoutDecorationItem*)ptr;
    [nSCollectionLayoutDecorationItem setZIndex:value];
}
