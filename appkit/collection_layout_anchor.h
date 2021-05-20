#import <stdbool.h>
#import <stdlib.h>
#import <utils.h>
#import <Appkit/Appkit.h>

void* C_CollectionLayoutAnchor_Alloc();

void* C_NSCollectionLayoutAnchor_CollectionLayoutAnchor_LayoutAnchorWithEdges(unsigned int edges);
void* C_NSCollectionLayoutAnchor_CollectionLayoutAnchor_LayoutAnchorWithEdges_AbsoluteOffset(unsigned int edges, CGPoint absoluteOffset);
void* C_NSCollectionLayoutAnchor_CollectionLayoutAnchor_LayoutAnchorWithEdges_FractionalOffset(unsigned int edges, CGPoint fractionalOffset);
unsigned int C_NSCollectionLayoutAnchor_Edges(void* ptr);
CGPoint C_NSCollectionLayoutAnchor_Offset(void* ptr);
bool C_NSCollectionLayoutAnchor_IsAbsoluteOffset(void* ptr);
bool C_NSCollectionLayoutAnchor_IsFractionalOffset(void* ptr);
