#import <Foundation/NSGeometry.h>
#import <Foundation/NSRange.h>
#import <stdbool.h>
#import <stdint.h>
#import <stdlib.h>
#import <utils.h>

void* C_CollectionViewTransitionLayout_Alloc();

void* C_NSCollectionViewTransitionLayout_InitWithCurrentLayout_NextLayout(void* ptr, void* currentLayout, void* newLayout);
void* C_NSCollectionViewTransitionLayout_AllocCollectionViewTransitionLayout();
void* C_NSCollectionViewTransitionLayout_Init(void* ptr);
void* C_NSCollectionViewTransitionLayout_NewCollectionViewTransitionLayout();
void* C_NSCollectionViewTransitionLayout_Autorelease(void* ptr);
void* C_NSCollectionViewTransitionLayout_Retain(void* ptr);
void C_NSCollectionViewTransitionLayout_UpdateValue_ForAnimatedKey(void* ptr, double value, void* key);
double C_NSCollectionViewTransitionLayout_ValueForAnimatedKey(void* ptr, void* key);
double C_NSCollectionViewTransitionLayout_TransitionProgress(void* ptr);
void C_NSCollectionViewTransitionLayout_SetTransitionProgress(void* ptr, double value);
void* C_NSCollectionViewTransitionLayout_CurrentLayout(void* ptr);
void* C_NSCollectionViewTransitionLayout_NextLayout(void* ptr);
