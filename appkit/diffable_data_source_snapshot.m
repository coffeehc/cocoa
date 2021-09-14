#import <Appkit/Appkit.h>
#import "diffable_data_source_snapshot.h"

void* C_DiffableDataSourceSnapshot_Alloc() {
    return [NSDiffableDataSourceSnapshot alloc];
}

void* C_NSDiffableDataSourceSnapshot_AllocDiffableDataSourceSnapshot() {
    NSDiffableDataSourceSnapshot* result_ = [NSDiffableDataSourceSnapshot alloc];
    return result_;
}

void* C_NSDiffableDataSourceSnapshot_Init(void* ptr) {
    NSDiffableDataSourceSnapshot* nSDiffableDataSourceSnapshot = (NSDiffableDataSourceSnapshot*)ptr;
    NSDiffableDataSourceSnapshot* result_ = [nSDiffableDataSourceSnapshot init];
    return result_;
}

void* C_NSDiffableDataSourceSnapshot_NewDiffableDataSourceSnapshot() {
    NSDiffableDataSourceSnapshot* result_ = [NSDiffableDataSourceSnapshot new];
    return result_;
}

void* C_NSDiffableDataSourceSnapshot_Autorelease(void* ptr) {
    NSDiffableDataSourceSnapshot* nSDiffableDataSourceSnapshot = (NSDiffableDataSourceSnapshot*)ptr;
    NSDiffableDataSourceSnapshot* result_ = [nSDiffableDataSourceSnapshot autorelease];
    return result_;
}

void* C_NSDiffableDataSourceSnapshot_Retain(void* ptr) {
    NSDiffableDataSourceSnapshot* nSDiffableDataSourceSnapshot = (NSDiffableDataSourceSnapshot*)ptr;
    NSDiffableDataSourceSnapshot* result_ = [nSDiffableDataSourceSnapshot retain];
    return result_;
}

void C_NSDiffableDataSourceSnapshot_AppendSectionsWithIdentifiers(void* ptr, Array sectionIdentifiers) {
    NSDiffableDataSourceSnapshot* nSDiffableDataSourceSnapshot = (NSDiffableDataSourceSnapshot*)ptr;
    NSMutableArray* objcSectionIdentifiers = [NSMutableArray arrayWithCapacity:sectionIdentifiers.len];
    if (sectionIdentifiers.len > 0) {
    	void** sectionIdentifiersData = (void**)sectionIdentifiers.data;
    	for (int i = 0; i < sectionIdentifiers.len; i++) {
    		void* p = sectionIdentifiersData[i];
    		[objcSectionIdentifiers addObject:(NSObject*)p];
    	}
    }
    [nSDiffableDataSourceSnapshot appendSectionsWithIdentifiers:objcSectionIdentifiers];
}

void C_NSDiffableDataSourceSnapshot_DeleteAllItems(void* ptr) {
    NSDiffableDataSourceSnapshot* nSDiffableDataSourceSnapshot = (NSDiffableDataSourceSnapshot*)ptr;
    [nSDiffableDataSourceSnapshot deleteAllItems];
}

int C_NSDiffableDataSourceSnapshot_NumberOfItems(void* ptr) {
    NSDiffableDataSourceSnapshot* nSDiffableDataSourceSnapshot = (NSDiffableDataSourceSnapshot*)ptr;
    NSInteger result_ = [nSDiffableDataSourceSnapshot numberOfItems];
    return result_;
}

int C_NSDiffableDataSourceSnapshot_NumberOfSections(void* ptr) {
    NSDiffableDataSourceSnapshot* nSDiffableDataSourceSnapshot = (NSDiffableDataSourceSnapshot*)ptr;
    NSInteger result_ = [nSDiffableDataSourceSnapshot numberOfSections];
    return result_;
}
