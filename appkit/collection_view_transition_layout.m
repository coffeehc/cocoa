#import <Appkit/Appkit.h>
#import "collection_view_transition_layout.h"

void* C_CollectionViewTransitionLayout_Alloc() {
    return [NSCollectionViewTransitionLayout alloc];
}

void* C_NSCollectionViewTransitionLayout_InitWithCurrentLayout_NextLayout(void* ptr, void* currentLayout, void* newLayout) {
    NSCollectionViewTransitionLayout* nSCollectionViewTransitionLayout = (NSCollectionViewTransitionLayout*)ptr;
    NSCollectionViewTransitionLayout* result_ = [nSCollectionViewTransitionLayout initWithCurrentLayout:(NSCollectionViewLayout*)currentLayout nextLayout:(NSCollectionViewLayout*)newLayout];
    return result_;
}

void* C_NSCollectionViewTransitionLayout_AllocCollectionViewTransitionLayout() {
    NSCollectionViewTransitionLayout* result_ = [NSCollectionViewTransitionLayout alloc];
    return result_;
}

void* C_NSCollectionViewTransitionLayout_Init(void* ptr) {
    NSCollectionViewTransitionLayout* nSCollectionViewTransitionLayout = (NSCollectionViewTransitionLayout*)ptr;
    NSCollectionViewTransitionLayout* result_ = [nSCollectionViewTransitionLayout init];
    return result_;
}

void* C_NSCollectionViewTransitionLayout_NewCollectionViewTransitionLayout() {
    NSCollectionViewTransitionLayout* result_ = [NSCollectionViewTransitionLayout new];
    return result_;
}

void* C_NSCollectionViewTransitionLayout_Autorelease(void* ptr) {
    NSCollectionViewTransitionLayout* nSCollectionViewTransitionLayout = (NSCollectionViewTransitionLayout*)ptr;
    NSCollectionViewTransitionLayout* result_ = [nSCollectionViewTransitionLayout autorelease];
    return result_;
}

void* C_NSCollectionViewTransitionLayout_Retain(void* ptr) {
    NSCollectionViewTransitionLayout* nSCollectionViewTransitionLayout = (NSCollectionViewTransitionLayout*)ptr;
    NSCollectionViewTransitionLayout* result_ = [nSCollectionViewTransitionLayout retain];
    return result_;
}

void C_NSCollectionViewTransitionLayout_UpdateValue_ForAnimatedKey(void* ptr, double value, void* key) {
    NSCollectionViewTransitionLayout* nSCollectionViewTransitionLayout = (NSCollectionViewTransitionLayout*)ptr;
    [nSCollectionViewTransitionLayout updateValue:value forAnimatedKey:(NSString*)key];
}

double C_NSCollectionViewTransitionLayout_ValueForAnimatedKey(void* ptr, void* key) {
    NSCollectionViewTransitionLayout* nSCollectionViewTransitionLayout = (NSCollectionViewTransitionLayout*)ptr;
    CGFloat result_ = [nSCollectionViewTransitionLayout valueForAnimatedKey:(NSString*)key];
    return result_;
}

double C_NSCollectionViewTransitionLayout_TransitionProgress(void* ptr) {
    NSCollectionViewTransitionLayout* nSCollectionViewTransitionLayout = (NSCollectionViewTransitionLayout*)ptr;
    CGFloat result_ = [nSCollectionViewTransitionLayout transitionProgress];
    return result_;
}

void C_NSCollectionViewTransitionLayout_SetTransitionProgress(void* ptr, double value) {
    NSCollectionViewTransitionLayout* nSCollectionViewTransitionLayout = (NSCollectionViewTransitionLayout*)ptr;
    [nSCollectionViewTransitionLayout setTransitionProgress:value];
}

void* C_NSCollectionViewTransitionLayout_CurrentLayout(void* ptr) {
    NSCollectionViewTransitionLayout* nSCollectionViewTransitionLayout = (NSCollectionViewTransitionLayout*)ptr;
    NSCollectionViewLayout* result_ = [nSCollectionViewTransitionLayout currentLayout];
    return result_;
}

void* C_NSCollectionViewTransitionLayout_NextLayout(void* ptr) {
    NSCollectionViewTransitionLayout* nSCollectionViewTransitionLayout = (NSCollectionViewTransitionLayout*)ptr;
    NSCollectionViewLayout* result_ = [nSCollectionViewTransitionLayout nextLayout];
    return result_;
}
