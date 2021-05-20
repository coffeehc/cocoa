#import <Appkit/Appkit.h>
#import "collection_view_diffable_data_source.h"

void* C_CollectionViewDiffableDataSource_Alloc() {
    return [NSCollectionViewDiffableDataSource alloc];
}

void* C_NSCollectionViewDiffableDataSource_Snapshot(void* ptr) {
    NSCollectionViewDiffableDataSource* nSCollectionViewDiffableDataSource = (NSCollectionViewDiffableDataSource*)ptr;
    NSDiffableDataSourceSnapshot* result_ = [nSCollectionViewDiffableDataSource snapshot];
    return result_;
}

void C_NSCollectionViewDiffableDataSource_ApplySnapshot_AnimatingDifferences(void* ptr, void* snapshot, bool animatingDifferences) {
    NSCollectionViewDiffableDataSource* nSCollectionViewDiffableDataSource = (NSCollectionViewDiffableDataSource*)ptr;
    [nSCollectionViewDiffableDataSource applySnapshot:(NSDiffableDataSourceSnapshot*)snapshot animatingDifferences:animatingDifferences];
}
