#import "collection_view_data_source.h"
#import <AppKit/NSCollectionView.h>
#import <Foundation/NSArray.h>
#import <Foundation/NSDictionary.h>
#import <_cgo_export.h>

@interface NSCollectionViewDataSourceAdaptor : NSObject <NSCollectionViewDataSource>
@property (assign) uintptr_t goID;
@end

@implementation NSCollectionViewDataSourceAdaptor


- (NSInteger)numberOfSectionsInCollectionView:(NSCollectionView*)collectionView {
    int result_ = collectionViewDataSource_NumberOfSectionsInCollectionView([self goID], collectionView);
    return result_;
}

- (NSInteger)collectionView:(NSCollectionView*)collectionView numberOfItemsInSection:(NSInteger)section {
    int result_ = collectionViewDataSource_CollectionView_NumberOfItemsInSection([self goID], collectionView, section);
    return result_;
}

- (NSCollectionViewItem*)collectionView:(NSCollectionView*)collectionView itemForRepresentedObjectAtIndexPath:(NSIndexPath*)indexPath {
    void* result_ = collectionViewDataSource_CollectionView_ItemForRepresentedObjectAtIndexPath([self goID], collectionView, indexPath);
    return (NSCollectionViewItem*)result_;
}

- (NSView*)collectionView:(NSCollectionView*)collectionView viewForSupplementaryElementOfKind:(NSCollectionViewSupplementaryElementKind)kind atIndexPath:(NSIndexPath*)indexPath {
    void* result_ = collectionViewDataSource_CollectionView_ViewForSupplementaryElementOfKind_AtIndexPath([self goID], collectionView, kind, indexPath);
    return (NSView*)result_;
}


- (BOOL)respondsToSelector:(SEL)aSelector {
	return CollectionViewDataSource_RespondsTo([self goID], aSelector);
}

- (void)dealloc {
	deleteAppKitHandle([self goID]);
	[super dealloc];
}
@end

void* WrapCollectionViewDataSource(uintptr_t goID) {
    NSCollectionViewDataSourceAdaptor* adaptor = [[NSCollectionViewDataSourceAdaptor alloc] init];
    adaptor.goID = goID;
    return adaptor;
}
