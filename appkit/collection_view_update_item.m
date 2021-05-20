#import <Appkit/Appkit.h>
#import "collection_view_update_item.h"

void* C_CollectionViewUpdateItem_Alloc() {
    return [NSCollectionViewUpdateItem alloc];
}

void* C_NSCollectionViewUpdateItem_Init(void* ptr) {
    NSCollectionViewUpdateItem* nSCollectionViewUpdateItem = (NSCollectionViewUpdateItem*)ptr;
    NSCollectionViewUpdateItem* result_ = [nSCollectionViewUpdateItem init];
    return result_;
}

void* C_NSCollectionViewUpdateItem_IndexPathBeforeUpdate(void* ptr) {
    NSCollectionViewUpdateItem* nSCollectionViewUpdateItem = (NSCollectionViewUpdateItem*)ptr;
    NSIndexPath* result_ = [nSCollectionViewUpdateItem indexPathBeforeUpdate];
    return result_;
}

void* C_NSCollectionViewUpdateItem_IndexPathAfterUpdate(void* ptr) {
    NSCollectionViewUpdateItem* nSCollectionViewUpdateItem = (NSCollectionViewUpdateItem*)ptr;
    NSIndexPath* result_ = [nSCollectionViewUpdateItem indexPathAfterUpdate];
    return result_;
}

int C_NSCollectionViewUpdateItem_UpdateAction(void* ptr) {
    NSCollectionViewUpdateItem* nSCollectionViewUpdateItem = (NSCollectionViewUpdateItem*)ptr;
    NSCollectionUpdateAction result_ = [nSCollectionViewUpdateItem updateAction];
    return result_;
}
