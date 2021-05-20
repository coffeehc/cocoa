#import <Appkit/Appkit.h>
#import "collection_view_grid_layout.h"

void* C_CollectionViewGridLayout_Alloc() {
    return [NSCollectionViewGridLayout alloc];
}

void* C_NSCollectionViewGridLayout_Init(void* ptr) {
    NSCollectionViewGridLayout* nSCollectionViewGridLayout = (NSCollectionViewGridLayout*)ptr;
    NSCollectionViewGridLayout* result_ = [nSCollectionViewGridLayout init];
    return result_;
}

unsigned int C_NSCollectionViewGridLayout_MaximumNumberOfRows(void* ptr) {
    NSCollectionViewGridLayout* nSCollectionViewGridLayout = (NSCollectionViewGridLayout*)ptr;
    NSUInteger result_ = [nSCollectionViewGridLayout maximumNumberOfRows];
    return result_;
}

void C_NSCollectionViewGridLayout_SetMaximumNumberOfRows(void* ptr, unsigned int value) {
    NSCollectionViewGridLayout* nSCollectionViewGridLayout = (NSCollectionViewGridLayout*)ptr;
    [nSCollectionViewGridLayout setMaximumNumberOfRows:value];
}

unsigned int C_NSCollectionViewGridLayout_MaximumNumberOfColumns(void* ptr) {
    NSCollectionViewGridLayout* nSCollectionViewGridLayout = (NSCollectionViewGridLayout*)ptr;
    NSUInteger result_ = [nSCollectionViewGridLayout maximumNumberOfColumns];
    return result_;
}

void C_NSCollectionViewGridLayout_SetMaximumNumberOfColumns(void* ptr, unsigned int value) {
    NSCollectionViewGridLayout* nSCollectionViewGridLayout = (NSCollectionViewGridLayout*)ptr;
    [nSCollectionViewGridLayout setMaximumNumberOfColumns:value];
}

CGSize C_NSCollectionViewGridLayout_MinimumItemSize(void* ptr) {
    NSCollectionViewGridLayout* nSCollectionViewGridLayout = (NSCollectionViewGridLayout*)ptr;
    NSSize result_ = [nSCollectionViewGridLayout minimumItemSize];
    return result_;
}

void C_NSCollectionViewGridLayout_SetMinimumItemSize(void* ptr, CGSize value) {
    NSCollectionViewGridLayout* nSCollectionViewGridLayout = (NSCollectionViewGridLayout*)ptr;
    [nSCollectionViewGridLayout setMinimumItemSize:value];
}

CGSize C_NSCollectionViewGridLayout_MaximumItemSize(void* ptr) {
    NSCollectionViewGridLayout* nSCollectionViewGridLayout = (NSCollectionViewGridLayout*)ptr;
    NSSize result_ = [nSCollectionViewGridLayout maximumItemSize];
    return result_;
}

void C_NSCollectionViewGridLayout_SetMaximumItemSize(void* ptr, CGSize value) {
    NSCollectionViewGridLayout* nSCollectionViewGridLayout = (NSCollectionViewGridLayout*)ptr;
    [nSCollectionViewGridLayout setMaximumItemSize:value];
}

double C_NSCollectionViewGridLayout_MinimumInteritemSpacing(void* ptr) {
    NSCollectionViewGridLayout* nSCollectionViewGridLayout = (NSCollectionViewGridLayout*)ptr;
    CGFloat result_ = [nSCollectionViewGridLayout minimumInteritemSpacing];
    return result_;
}

void C_NSCollectionViewGridLayout_SetMinimumInteritemSpacing(void* ptr, double value) {
    NSCollectionViewGridLayout* nSCollectionViewGridLayout = (NSCollectionViewGridLayout*)ptr;
    [nSCollectionViewGridLayout setMinimumInteritemSpacing:value];
}

double C_NSCollectionViewGridLayout_MinimumLineSpacing(void* ptr) {
    NSCollectionViewGridLayout* nSCollectionViewGridLayout = (NSCollectionViewGridLayout*)ptr;
    CGFloat result_ = [nSCollectionViewGridLayout minimumLineSpacing];
    return result_;
}

void C_NSCollectionViewGridLayout_SetMinimumLineSpacing(void* ptr, double value) {
    NSCollectionViewGridLayout* nSCollectionViewGridLayout = (NSCollectionViewGridLayout*)ptr;
    [nSCollectionViewGridLayout setMinimumLineSpacing:value];
}

NSEdgeInsets C_NSCollectionViewGridLayout_Margins(void* ptr) {
    NSCollectionViewGridLayout* nSCollectionViewGridLayout = (NSCollectionViewGridLayout*)ptr;
    NSEdgeInsets result_ = [nSCollectionViewGridLayout margins];
    return result_;
}

void C_NSCollectionViewGridLayout_SetMargins(void* ptr, NSEdgeInsets value) {
    NSCollectionViewGridLayout* nSCollectionViewGridLayout = (NSCollectionViewGridLayout*)ptr;
    [nSCollectionViewGridLayout setMargins:value];
}

Array C_NSCollectionViewGridLayout_BackgroundColors(void* ptr) {
    NSCollectionViewGridLayout* nSCollectionViewGridLayout = (NSCollectionViewGridLayout*)ptr;
    NSArray* result_ = [nSCollectionViewGridLayout backgroundColors];
    int result_count = [result_ count];
    void** result_Data = malloc(result_count * sizeof(void*));
    for (int i = 0; i < result_count; i++) {
    	 void* p = [result_ objectAtIndex:i];
    	 result_Data[i] = p;
    }
    Array result_Array;
    result_Array.data = result_Data;
    result_Array.len = result_count;
    return result_Array;
}

void C_NSCollectionViewGridLayout_SetBackgroundColors(void* ptr, Array value) {
    NSCollectionViewGridLayout* nSCollectionViewGridLayout = (NSCollectionViewGridLayout*)ptr;
    NSMutableArray* objcValue = [[NSMutableArray alloc] init];
    void** valueData = (void**)value.data;
    for (int i = 0; i < value.len; i++) {
    	void* p = valueData[i];
    	[objcValue addObject:(NSColor*)(NSColor*)p];
    }
    [nSCollectionViewGridLayout setBackgroundColors:objcValue];
}
