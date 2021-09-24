#import "collection_view_update_item.h"
#import <AppKit/NSCollectionViewLayout.h>
#import <Foundation/NSArray.h>
#import <Foundation/NSDictionary.h>

void* C_CollectionViewUpdateItem_Alloc() {
    return [NSCollectionViewUpdateItem alloc];
}

void* C_NSCollectionViewUpdateItem_AllocCollectionViewUpdateItem() {
    NSCollectionViewUpdateItem* result_ = [NSCollectionViewUpdateItem alloc];
    return result_;
}

void* C_NSCollectionViewUpdateItem_Init(void* ptr) {
    NSCollectionViewUpdateItem* nSCollectionViewUpdateItem = (NSCollectionViewUpdateItem*)ptr;
    NSCollectionViewUpdateItem* result_ = [nSCollectionViewUpdateItem init];
    return result_;
}

void* C_NSCollectionViewUpdateItem_NewCollectionViewUpdateItem() {
    NSCollectionViewUpdateItem* result_ = [NSCollectionViewUpdateItem new];
    return result_;
}

void* C_NSCollectionViewUpdateItem_Autorelease(void* ptr) {
    NSCollectionViewUpdateItem* nSCollectionViewUpdateItem = (NSCollectionViewUpdateItem*)ptr;
    NSCollectionViewUpdateItem* result_ = [nSCollectionViewUpdateItem autorelease];
    return result_;
}

void* C_NSCollectionViewUpdateItem_Retain(void* ptr) {
    NSCollectionViewUpdateItem* nSCollectionViewUpdateItem = (NSCollectionViewUpdateItem*)ptr;
    NSCollectionViewUpdateItem* result_ = [nSCollectionViewUpdateItem retain];
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
