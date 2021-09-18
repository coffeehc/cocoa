#import <stdbool.h>
#import <stdlib.h>
#import <utils.h>
#import <Foundation/NSGeometry.h>
#import <Foundation/NSRange.h>

void* C_CollectionViewDiffableDataSource_Alloc();

void* C_NSCollectionViewDiffableDataSource_AllocCollectionViewDiffableDataSource();
void* C_NSCollectionViewDiffableDataSource_Autorelease(void* ptr);
void* C_NSCollectionViewDiffableDataSource_Retain(void* ptr);
void* C_NSCollectionViewDiffableDataSource_Snapshot(void* ptr);
void C_NSCollectionViewDiffableDataSource_ApplySnapshot_AnimatingDifferences(void* ptr, void* snapshot, bool animatingDifferences);
