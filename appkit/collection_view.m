#import <Appkit/Appkit.h>
#import "collection_view.h"

void* C_CollectionView_Alloc() {
    return [NSCollectionView alloc];
}

void* C_NSCollectionView_InitWithFrame(void* ptr, CGRect frameRect) {
    NSCollectionView* nSCollectionView = (NSCollectionView*)ptr;
    NSCollectionView* result_ = [nSCollectionView initWithFrame:frameRect];
    return result_;
}

void* C_NSCollectionView_InitWithCoder(void* ptr, void* coder) {
    NSCollectionView* nSCollectionView = (NSCollectionView*)ptr;
    NSCollectionView* result_ = [nSCollectionView initWithCoder:(NSCoder*)coder];
    return result_;
}

void* C_NSCollectionView_Init(void* ptr) {
    NSCollectionView* nSCollectionView = (NSCollectionView*)ptr;
    NSCollectionView* result_ = [nSCollectionView init];
    return result_;
}

void* C_NSCollectionView_AllocCollectionView() {
    NSCollectionView* result_ = [NSCollectionView alloc];
    return result_;
}

void* C_NSCollectionView_NewCollectionView() {
    NSCollectionView* result_ = [NSCollectionView new];
    return result_;
}

void* C_NSCollectionView_Autorelease(void* ptr) {
    NSCollectionView* nSCollectionView = (NSCollectionView*)ptr;
    NSCollectionView* result_ = [nSCollectionView autorelease];
    return result_;
}

void* C_NSCollectionView_Retain(void* ptr) {
    NSCollectionView* nSCollectionView = (NSCollectionView*)ptr;
    NSCollectionView* result_ = [nSCollectionView retain];
    return result_;
}

void* C_NSCollectionView_MakeItemWithIdentifier_ForIndexPath(void* ptr, void* identifier, void* indexPath) {
    NSCollectionView* nSCollectionView = (NSCollectionView*)ptr;
    NSCollectionViewItem* result_ = [nSCollectionView makeItemWithIdentifier:(NSString*)identifier forIndexPath:(NSIndexPath*)indexPath];
    return result_;
}

void C_NSCollectionView_RegisterNib_ForItemWithIdentifier(void* ptr, void* nib, void* identifier) {
    NSCollectionView* nSCollectionView = (NSCollectionView*)ptr;
    [nSCollectionView registerNib:(NSNib*)nib forItemWithIdentifier:(NSString*)identifier];
}

void* C_NSCollectionView_MakeSupplementaryViewOfKind_WithIdentifier_ForIndexPath(void* ptr, void* elementKind, void* identifier, void* indexPath) {
    NSCollectionView* nSCollectionView = (NSCollectionView*)ptr;
    NSView* result_ = [nSCollectionView makeSupplementaryViewOfKind:(NSString*)elementKind withIdentifier:(NSString*)identifier forIndexPath:(NSIndexPath*)indexPath];
    return result_;
}

void C_NSCollectionView_RegisterNib_ForSupplementaryViewOfKind_WithIdentifier(void* ptr, void* nib, void* kind, void* identifier) {
    NSCollectionView* nSCollectionView = (NSCollectionView*)ptr;
    [nSCollectionView registerNib:(NSNib*)nib forSupplementaryViewOfKind:(NSString*)kind withIdentifier:(NSString*)identifier];
}

void C_NSCollectionView_ReloadData(void* ptr) {
    NSCollectionView* nSCollectionView = (NSCollectionView*)ptr;
    [nSCollectionView reloadData];
}

void C_NSCollectionView_ReloadSections(void* ptr, void* sections) {
    NSCollectionView* nSCollectionView = (NSCollectionView*)ptr;
    [nSCollectionView reloadSections:(NSIndexSet*)sections];
}

void C_NSCollectionView_ReloadItemsAtIndexPaths(void* ptr, void* indexPaths) {
    NSCollectionView* nSCollectionView = (NSCollectionView*)ptr;
    [nSCollectionView reloadItemsAtIndexPaths:(NSSet*)indexPaths];
}

int C_NSCollectionView_NumberOfItemsInSection(void* ptr, int section) {
    NSCollectionView* nSCollectionView = (NSCollectionView*)ptr;
    NSInteger result_ = [nSCollectionView numberOfItemsInSection:section];
    return result_;
}

void C_NSCollectionView_InsertItemsAtIndexPaths(void* ptr, void* indexPaths) {
    NSCollectionView* nSCollectionView = (NSCollectionView*)ptr;
    [nSCollectionView insertItemsAtIndexPaths:(NSSet*)indexPaths];
}

void C_NSCollectionView_MoveItemAtIndexPath_ToIndexPath(void* ptr, void* indexPath, void* newIndexPath) {
    NSCollectionView* nSCollectionView = (NSCollectionView*)ptr;
    [nSCollectionView moveItemAtIndexPath:(NSIndexPath*)indexPath toIndexPath:(NSIndexPath*)newIndexPath];
}

void C_NSCollectionView_DeleteItemsAtIndexPaths(void* ptr, void* indexPaths) {
    NSCollectionView* nSCollectionView = (NSCollectionView*)ptr;
    [nSCollectionView deleteItemsAtIndexPaths:(NSSet*)indexPaths];
}

void C_NSCollectionView_InsertSections(void* ptr, void* sections) {
    NSCollectionView* nSCollectionView = (NSCollectionView*)ptr;
    [nSCollectionView insertSections:(NSIndexSet*)sections];
}

void C_NSCollectionView_MoveSection_ToSection(void* ptr, int section, int newSection) {
    NSCollectionView* nSCollectionView = (NSCollectionView*)ptr;
    [nSCollectionView moveSection:section toSection:newSection];
}

void C_NSCollectionView_DeleteSections(void* ptr, void* sections) {
    NSCollectionView* nSCollectionView = (NSCollectionView*)ptr;
    [nSCollectionView deleteSections:(NSIndexSet*)sections];
}

void C_NSCollectionView_ToggleSectionCollapse(void* ptr, void* sender) {
    NSCollectionView* nSCollectionView = (NSCollectionView*)ptr;
    [nSCollectionView toggleSectionCollapse:(id)sender];
}

void C_NSCollectionView_SelectAll(void* ptr, void* sender) {
    NSCollectionView* nSCollectionView = (NSCollectionView*)ptr;
    [nSCollectionView selectAll:(id)sender];
}

void C_NSCollectionView_DeselectAll(void* ptr, void* sender) {
    NSCollectionView* nSCollectionView = (NSCollectionView*)ptr;
    [nSCollectionView deselectAll:(id)sender];
}

void C_NSCollectionView_SelectItemsAtIndexPaths_ScrollPosition(void* ptr, void* indexPaths, unsigned int scrollPosition) {
    NSCollectionView* nSCollectionView = (NSCollectionView*)ptr;
    [nSCollectionView selectItemsAtIndexPaths:(NSSet*)indexPaths scrollPosition:scrollPosition];
}

void C_NSCollectionView_DeselectItemsAtIndexPaths(void* ptr, void* indexPaths) {
    NSCollectionView* nSCollectionView = (NSCollectionView*)ptr;
    [nSCollectionView deselectItemsAtIndexPaths:(NSSet*)indexPaths];
}

Array C_NSCollectionView_VisibleItems(void* ptr) {
    NSCollectionView* nSCollectionView = (NSCollectionView*)ptr;
    NSArray* result_ = [nSCollectionView visibleItems];
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

void* C_NSCollectionView_IndexPathsForVisibleItems(void* ptr) {
    NSCollectionView* nSCollectionView = (NSCollectionView*)ptr;
    NSSet* result_ = [nSCollectionView indexPathsForVisibleItems];
    return result_;
}

void* C_NSCollectionView_IndexPathsForVisibleSupplementaryElementsOfKind(void* ptr, void* elementKind) {
    NSCollectionView* nSCollectionView = (NSCollectionView*)ptr;
    NSSet* result_ = [nSCollectionView indexPathsForVisibleSupplementaryElementsOfKind:(NSString*)elementKind];
    return result_;
}

void* C_NSCollectionView_IndexPathForItem(void* ptr, void* item) {
    NSCollectionView* nSCollectionView = (NSCollectionView*)ptr;
    NSIndexPath* result_ = [nSCollectionView indexPathForItem:(NSCollectionViewItem*)item];
    return result_;
}

void* C_NSCollectionView_IndexPathForItemAtPoint(void* ptr, CGPoint point) {
    NSCollectionView* nSCollectionView = (NSCollectionView*)ptr;
    NSIndexPath* result_ = [nSCollectionView indexPathForItemAtPoint:point];
    return result_;
}

void* C_NSCollectionView_ItemAtIndexPath(void* ptr, void* indexPath) {
    NSCollectionView* nSCollectionView = (NSCollectionView*)ptr;
    NSCollectionViewItem* result_ = [nSCollectionView itemAtIndexPath:(NSIndexPath*)indexPath];
    return result_;
}

void* C_NSCollectionView_SupplementaryViewForElementKind_AtIndexPath(void* ptr, void* elementKind, void* indexPath) {
    NSCollectionView* nSCollectionView = (NSCollectionView*)ptr;
    NSView* result_ = [nSCollectionView supplementaryViewForElementKind:(NSString*)elementKind atIndexPath:(NSIndexPath*)indexPath];
    return result_;
}

void C_NSCollectionView_ScrollToItemsAtIndexPaths_ScrollPosition(void* ptr, void* indexPaths, unsigned int scrollPosition) {
    NSCollectionView* nSCollectionView = (NSCollectionView*)ptr;
    [nSCollectionView scrollToItemsAtIndexPaths:(NSSet*)indexPaths scrollPosition:scrollPosition];
}

void* C_NSCollectionView_LayoutAttributesForItemAtIndexPath(void* ptr, void* indexPath) {
    NSCollectionView* nSCollectionView = (NSCollectionView*)ptr;
    NSCollectionViewLayoutAttributes* result_ = [nSCollectionView layoutAttributesForItemAtIndexPath:(NSIndexPath*)indexPath];
    return result_;
}

void* C_NSCollectionView_LayoutAttributesForSupplementaryElementOfKind_AtIndexPath(void* ptr, void* kind, void* indexPath) {
    NSCollectionView* nSCollectionView = (NSCollectionView*)ptr;
    NSCollectionViewLayoutAttributes* result_ = [nSCollectionView layoutAttributesForSupplementaryElementOfKind:(NSString*)kind atIndexPath:(NSIndexPath*)indexPath];
    return result_;
}

void* C_NSCollectionView_ItemAtIndex(void* ptr, unsigned int index) {
    NSCollectionView* nSCollectionView = (NSCollectionView*)ptr;
    NSCollectionViewItem* result_ = [nSCollectionView itemAtIndex:index];
    return result_;
}

CGRect C_NSCollectionView_FrameForItemAtIndex(void* ptr, unsigned int index) {
    NSCollectionView* nSCollectionView = (NSCollectionView*)ptr;
    NSRect result_ = [nSCollectionView frameForItemAtIndex:index];
    return result_;
}

CGRect C_NSCollectionView_FrameForItemAtIndex_WithNumberOfItems(void* ptr, unsigned int index, unsigned int numberOfItems) {
    NSCollectionView* nSCollectionView = (NSCollectionView*)ptr;
    NSRect result_ = [nSCollectionView frameForItemAtIndex:index withNumberOfItems:numberOfItems];
    return result_;
}

void C_NSCollectionView_SetDraggingSourceOperationMask_ForLocal(void* ptr, unsigned int dragOperationMask, bool localDestination) {
    NSCollectionView* nSCollectionView = (NSCollectionView*)ptr;
    [nSCollectionView setDraggingSourceOperationMask:dragOperationMask forLocal:localDestination];
}

void* C_NSCollectionView_DataSource(void* ptr) {
    NSCollectionView* nSCollectionView = (NSCollectionView*)ptr;
    id result_ = [nSCollectionView dataSource];
    return result_;
}

void C_NSCollectionView_SetDataSource(void* ptr, void* value) {
    NSCollectionView* nSCollectionView = (NSCollectionView*)ptr;
    [nSCollectionView setDataSource:(id)value];
}

void* C_NSCollectionView_Delegate(void* ptr) {
    NSCollectionView* nSCollectionView = (NSCollectionView*)ptr;
    id result_ = [nSCollectionView delegate];
    return result_;
}

void C_NSCollectionView_SetDelegate(void* ptr, void* value) {
    NSCollectionView* nSCollectionView = (NSCollectionView*)ptr;
    [nSCollectionView setDelegate:(id)value];
}

Array C_NSCollectionView_Content(void* ptr) {
    NSCollectionView* nSCollectionView = (NSCollectionView*)ptr;
    NSArray* result_ = [nSCollectionView content];
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

void C_NSCollectionView_SetContent(void* ptr, Array value) {
    NSCollectionView* nSCollectionView = (NSCollectionView*)ptr;
    NSMutableArray* objcValue = [NSMutableArray arrayWithCapacity:value.len];
    if (value.len > 0) {
    	void** valueData = (void**)value.data;
    	for (int i = 0; i < value.len; i++) {
    		void* p = valueData[i];
    		[objcValue addObject:(id)(id)p];
    	}
    }
    [nSCollectionView setContent:objcValue];
}

void* C_NSCollectionView_BackgroundView(void* ptr) {
    NSCollectionView* nSCollectionView = (NSCollectionView*)ptr;
    NSView* result_ = [nSCollectionView backgroundView];
    return result_;
}

void C_NSCollectionView_SetBackgroundView(void* ptr, void* value) {
    NSCollectionView* nSCollectionView = (NSCollectionView*)ptr;
    [nSCollectionView setBackgroundView:(NSView*)value];
}

Array C_NSCollectionView_BackgroundColors(void* ptr) {
    NSCollectionView* nSCollectionView = (NSCollectionView*)ptr;
    NSArray* result_ = [nSCollectionView backgroundColors];
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

void C_NSCollectionView_SetBackgroundColors(void* ptr, Array value) {
    NSCollectionView* nSCollectionView = (NSCollectionView*)ptr;
    NSMutableArray* objcValue = [NSMutableArray arrayWithCapacity:value.len];
    if (value.len > 0) {
    	void** valueData = (void**)value.data;
    	for (int i = 0; i < value.len; i++) {
    		void* p = valueData[i];
    		[objcValue addObject:(NSColor*)(NSColor*)p];
    	}
    }
    [nSCollectionView setBackgroundColors:objcValue];
}

bool C_NSCollectionView_BackgroundViewScrollsWithContent(void* ptr) {
    NSCollectionView* nSCollectionView = (NSCollectionView*)ptr;
    BOOL result_ = [nSCollectionView backgroundViewScrollsWithContent];
    return result_;
}

void C_NSCollectionView_SetBackgroundViewScrollsWithContent(void* ptr, bool value) {
    NSCollectionView* nSCollectionView = (NSCollectionView*)ptr;
    [nSCollectionView setBackgroundViewScrollsWithContent:value];
}

void* C_NSCollectionView_CollectionViewLayout(void* ptr) {
    NSCollectionView* nSCollectionView = (NSCollectionView*)ptr;
    NSCollectionViewLayout* result_ = [nSCollectionView collectionViewLayout];
    return result_;
}

void C_NSCollectionView_SetCollectionViewLayout(void* ptr, void* value) {
    NSCollectionView* nSCollectionView = (NSCollectionView*)ptr;
    [nSCollectionView setCollectionViewLayout:(NSCollectionViewLayout*)value];
}

void* C_NSCollectionView_PrefetchDataSource(void* ptr) {
    NSCollectionView* nSCollectionView = (NSCollectionView*)ptr;
    id result_ = [nSCollectionView prefetchDataSource];
    return result_;
}

void C_NSCollectionView_SetPrefetchDataSource(void* ptr, void* value) {
    NSCollectionView* nSCollectionView = (NSCollectionView*)ptr;
    [nSCollectionView setPrefetchDataSource:(id)value];
}

int C_NSCollectionView_NumberOfSections(void* ptr) {
    NSCollectionView* nSCollectionView = (NSCollectionView*)ptr;
    NSInteger result_ = [nSCollectionView numberOfSections];
    return result_;
}

bool C_NSCollectionView_IsSelectable(void* ptr) {
    NSCollectionView* nSCollectionView = (NSCollectionView*)ptr;
    BOOL result_ = [nSCollectionView isSelectable];
    return result_;
}

void C_NSCollectionView_SetSelectable(void* ptr, bool value) {
    NSCollectionView* nSCollectionView = (NSCollectionView*)ptr;
    [nSCollectionView setSelectable:value];
}

bool C_NSCollectionView_AllowsMultipleSelection(void* ptr) {
    NSCollectionView* nSCollectionView = (NSCollectionView*)ptr;
    BOOL result_ = [nSCollectionView allowsMultipleSelection];
    return result_;
}

void C_NSCollectionView_SetAllowsMultipleSelection(void* ptr, bool value) {
    NSCollectionView* nSCollectionView = (NSCollectionView*)ptr;
    [nSCollectionView setAllowsMultipleSelection:value];
}

bool C_NSCollectionView_AllowsEmptySelection(void* ptr) {
    NSCollectionView* nSCollectionView = (NSCollectionView*)ptr;
    BOOL result_ = [nSCollectionView allowsEmptySelection];
    return result_;
}

void C_NSCollectionView_SetAllowsEmptySelection(void* ptr, bool value) {
    NSCollectionView* nSCollectionView = (NSCollectionView*)ptr;
    [nSCollectionView setAllowsEmptySelection:value];
}

void* C_NSCollectionView_SelectionIndexPaths(void* ptr) {
    NSCollectionView* nSCollectionView = (NSCollectionView*)ptr;
    NSSet* result_ = [nSCollectionView selectionIndexPaths];
    return result_;
}

void C_NSCollectionView_SetSelectionIndexPaths(void* ptr, void* value) {
    NSCollectionView* nSCollectionView = (NSCollectionView*)ptr;
    [nSCollectionView setSelectionIndexPaths:(NSSet*)value];
}

bool C_NSCollectionView_IsFirstResponder(void* ptr) {
    NSCollectionView* nSCollectionView = (NSCollectionView*)ptr;
    BOOL result_ = [nSCollectionView isFirstResponder];
    return result_;
}

void* C_NSCollectionView_SelectionIndexes(void* ptr) {
    NSCollectionView* nSCollectionView = (NSCollectionView*)ptr;
    NSIndexSet* result_ = [nSCollectionView selectionIndexes];
    return result_;
}

void C_NSCollectionView_SetSelectionIndexes(void* ptr, void* value) {
    NSCollectionView* nSCollectionView = (NSCollectionView*)ptr;
    [nSCollectionView setSelectionIndexes:(NSIndexSet*)value];
}
