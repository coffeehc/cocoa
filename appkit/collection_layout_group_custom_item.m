#import <Appkit/Appkit.h>
#import "collection_layout_group_custom_item.h"

void* C_CollectionLayoutGroupCustomItem_Alloc() {
    return [NSCollectionLayoutGroupCustomItem alloc];
}

void* C_NSCollectionLayoutGroupCustomItem_CollectionLayoutGroupCustomItem_CustomItemWithFrame(CGRect frame) {
    NSCollectionLayoutGroupCustomItem* result_ = [NSCollectionLayoutGroupCustomItem customItemWithFrame:frame];
    return result_;
}

void* C_NSCollectionLayoutGroupCustomItem_CollectionLayoutGroupCustomItem_CustomItemWithFrame_ZIndex(CGRect frame, int zIndex) {
    NSCollectionLayoutGroupCustomItem* result_ = [NSCollectionLayoutGroupCustomItem customItemWithFrame:frame zIndex:zIndex];
    return result_;
}

CGRect C_NSCollectionLayoutGroupCustomItem_Frame(void* ptr) {
    NSCollectionLayoutGroupCustomItem* nSCollectionLayoutGroupCustomItem = (NSCollectionLayoutGroupCustomItem*)ptr;
    NSRect result_ = [nSCollectionLayoutGroupCustomItem frame];
    return result_;
}

int C_NSCollectionLayoutGroupCustomItem_ZIndex(void* ptr) {
    NSCollectionLayoutGroupCustomItem* nSCollectionLayoutGroupCustomItem = (NSCollectionLayoutGroupCustomItem*)ptr;
    NSInteger result_ = [nSCollectionLayoutGroupCustomItem zIndex];
    return result_;
}
