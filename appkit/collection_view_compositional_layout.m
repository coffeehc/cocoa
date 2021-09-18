#import "collection_view_compositional_layout.h"
#import <Foundation/NSDictionary.h>
#import <Foundation/NSArray.h>
#import <AppKit/NSCollectionViewCompositionalLayout.h>

void* C_CollectionViewCompositionalLayout_Alloc() {
    return [NSCollectionViewCompositionalLayout alloc];
}

void* C_NSCollectionViewCompositionalLayout_InitWithSection(void* ptr, void* section) {
    NSCollectionViewCompositionalLayout* nSCollectionViewCompositionalLayout = (NSCollectionViewCompositionalLayout*)ptr;
    NSCollectionViewCompositionalLayout* result_ = [nSCollectionViewCompositionalLayout initWithSection:(NSCollectionLayoutSection*)section];
    return result_;
}

void* C_NSCollectionViewCompositionalLayout_InitWithSection_Configuration(void* ptr, void* section, void* configuration) {
    NSCollectionViewCompositionalLayout* nSCollectionViewCompositionalLayout = (NSCollectionViewCompositionalLayout*)ptr;
    NSCollectionViewCompositionalLayout* result_ = [nSCollectionViewCompositionalLayout initWithSection:(NSCollectionLayoutSection*)section configuration:(NSCollectionViewCompositionalLayoutConfiguration*)configuration];
    return result_;
}

void* C_NSCollectionViewCompositionalLayout_AllocCollectionViewCompositionalLayout() {
    NSCollectionViewCompositionalLayout* result_ = [NSCollectionViewCompositionalLayout alloc];
    return result_;
}

void* C_NSCollectionViewCompositionalLayout_Autorelease(void* ptr) {
    NSCollectionViewCompositionalLayout* nSCollectionViewCompositionalLayout = (NSCollectionViewCompositionalLayout*)ptr;
    NSCollectionViewCompositionalLayout* result_ = [nSCollectionViewCompositionalLayout autorelease];
    return result_;
}

void* C_NSCollectionViewCompositionalLayout_Retain(void* ptr) {
    NSCollectionViewCompositionalLayout* nSCollectionViewCompositionalLayout = (NSCollectionViewCompositionalLayout*)ptr;
    NSCollectionViewCompositionalLayout* result_ = [nSCollectionViewCompositionalLayout retain];
    return result_;
}

void* C_NSCollectionViewCompositionalLayout_Configuration(void* ptr) {
    NSCollectionViewCompositionalLayout* nSCollectionViewCompositionalLayout = (NSCollectionViewCompositionalLayout*)ptr;
    NSCollectionViewCompositionalLayoutConfiguration* result_ = [nSCollectionViewCompositionalLayout configuration];
    return result_;
}

void C_NSCollectionViewCompositionalLayout_SetConfiguration(void* ptr, void* value) {
    NSCollectionViewCompositionalLayout* nSCollectionViewCompositionalLayout = (NSCollectionViewCompositionalLayout*)ptr;
    [nSCollectionViewCompositionalLayout setConfiguration:(NSCollectionViewCompositionalLayoutConfiguration*)value];
}
