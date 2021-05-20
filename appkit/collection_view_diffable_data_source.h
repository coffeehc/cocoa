#import <stdbool.h>
#import <stdlib.h>
#import <utils.h>
#import <Appkit/Appkit.h>

void* C_CollectionViewDiffableDataSource_Alloc();

void* C_NSCollectionViewDiffableDataSource_Snapshot(void* ptr);
void C_NSCollectionViewDiffableDataSource_ApplySnapshot_AnimatingDifferences(void* ptr, void* snapshot, bool animatingDifferences);
