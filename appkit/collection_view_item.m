#import <Appkit/Appkit.h>
#import "collection_view_item.h"

void* C_CollectionViewItem_Alloc() {
    return [NSCollectionViewItem alloc];
}

void* C_NSCollectionViewItem_InitWithNibName_Bundle(void* ptr, void* nibNameOrNil, void* nibBundleOrNil) {
    NSCollectionViewItem* nSCollectionViewItem = (NSCollectionViewItem*)ptr;
    NSCollectionViewItem* result_ = [nSCollectionViewItem initWithNibName:(NSString*)nibNameOrNil bundle:(NSBundle*)nibBundleOrNil];
    return result_;
}

void* C_NSCollectionViewItem_InitWithCoder(void* ptr, void* coder) {
    NSCollectionViewItem* nSCollectionViewItem = (NSCollectionViewItem*)ptr;
    NSCollectionViewItem* result_ = [nSCollectionViewItem initWithCoder:(NSCoder*)coder];
    return result_;
}

void* C_NSCollectionViewItem_Init(void* ptr) {
    NSCollectionViewItem* nSCollectionViewItem = (NSCollectionViewItem*)ptr;
    NSCollectionViewItem* result_ = [nSCollectionViewItem init];
    return result_;
}

bool C_NSCollectionViewItem_IsSelected(void* ptr) {
    NSCollectionViewItem* nSCollectionViewItem = (NSCollectionViewItem*)ptr;
    BOOL result_ = [nSCollectionViewItem isSelected];
    return result_;
}

void C_NSCollectionViewItem_SetSelected(void* ptr, bool value) {
    NSCollectionViewItem* nSCollectionViewItem = (NSCollectionViewItem*)ptr;
    [nSCollectionViewItem setSelected:value];
}

int C_NSCollectionViewItem_HighlightState(void* ptr) {
    NSCollectionViewItem* nSCollectionViewItem = (NSCollectionViewItem*)ptr;
    NSCollectionViewItemHighlightState result_ = [nSCollectionViewItem highlightState];
    return result_;
}

void C_NSCollectionViewItem_SetHighlightState(void* ptr, int value) {
    NSCollectionViewItem* nSCollectionViewItem = (NSCollectionViewItem*)ptr;
    [nSCollectionViewItem setHighlightState:value];
}

void* C_NSCollectionViewItem_CollectionView(void* ptr) {
    NSCollectionViewItem* nSCollectionViewItem = (NSCollectionViewItem*)ptr;
    NSCollectionView* result_ = [nSCollectionViewItem collectionView];
    return result_;
}

Array C_NSCollectionViewItem_DraggingImageComponents(void* ptr) {
    NSCollectionViewItem* nSCollectionViewItem = (NSCollectionViewItem*)ptr;
    NSArray* result_ = [nSCollectionViewItem draggingImageComponents];
    Array result_Array;
    int result_count = [result_ count];
    if (result_count > 0) {
    	void** result_Data = malloc(result_count * sizeof(void*));
    	for (int i = 0; i < result_count; i++) {
    		 void* p = [result_ objectAtIndex:i];
    		 result_Data[i] = p;
    	}
    	result_Array.data = result_Data;
    	result_Array.len = result_count;
    }
    return result_Array;
}
