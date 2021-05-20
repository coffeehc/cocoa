#import <Appkit/Appkit.h>
#import "collection_layout_anchor.h"

void* C_CollectionLayoutAnchor_Alloc() {
    return [NSCollectionLayoutAnchor alloc];
}

void* C_NSCollectionLayoutAnchor_CollectionLayoutAnchor_LayoutAnchorWithEdges(unsigned int edges) {
    NSCollectionLayoutAnchor* result_ = [NSCollectionLayoutAnchor layoutAnchorWithEdges:edges];
    return result_;
}

void* C_NSCollectionLayoutAnchor_CollectionLayoutAnchor_LayoutAnchorWithEdges_AbsoluteOffset(unsigned int edges, CGPoint absoluteOffset) {
    NSCollectionLayoutAnchor* result_ = [NSCollectionLayoutAnchor layoutAnchorWithEdges:edges absoluteOffset:absoluteOffset];
    return result_;
}

void* C_NSCollectionLayoutAnchor_CollectionLayoutAnchor_LayoutAnchorWithEdges_FractionalOffset(unsigned int edges, CGPoint fractionalOffset) {
    NSCollectionLayoutAnchor* result_ = [NSCollectionLayoutAnchor layoutAnchorWithEdges:edges fractionalOffset:fractionalOffset];
    return result_;
}

unsigned int C_NSCollectionLayoutAnchor_Edges(void* ptr) {
    NSCollectionLayoutAnchor* nSCollectionLayoutAnchor = (NSCollectionLayoutAnchor*)ptr;
    NSDirectionalRectEdge result_ = [nSCollectionLayoutAnchor edges];
    return result_;
}

CGPoint C_NSCollectionLayoutAnchor_Offset(void* ptr) {
    NSCollectionLayoutAnchor* nSCollectionLayoutAnchor = (NSCollectionLayoutAnchor*)ptr;
    NSPoint result_ = [nSCollectionLayoutAnchor offset];
    return result_;
}

bool C_NSCollectionLayoutAnchor_IsAbsoluteOffset(void* ptr) {
    NSCollectionLayoutAnchor* nSCollectionLayoutAnchor = (NSCollectionLayoutAnchor*)ptr;
    BOOL result_ = [nSCollectionLayoutAnchor isAbsoluteOffset];
    return result_;
}

bool C_NSCollectionLayoutAnchor_IsFractionalOffset(void* ptr) {
    NSCollectionLayoutAnchor* nSCollectionLayoutAnchor = (NSCollectionLayoutAnchor*)ptr;
    BOOL result_ = [nSCollectionLayoutAnchor isFractionalOffset];
    return result_;
}
