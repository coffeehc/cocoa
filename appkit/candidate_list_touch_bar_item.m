#import "candidate_list_touch_bar_item.h"
#import <Foundation/NSDictionary.h>
#import <Foundation/NSArray.h>
#import <AppKit/NSCandidateListTouchBarItem.h>

void* C_CandidateListTouchBarItem_Alloc() {
    return [NSCandidateListTouchBarItem alloc];
}

void* C_NSCandidateListTouchBarItem_InitWithIdentifier(void* ptr, void* identifier) {
    NSCandidateListTouchBarItem* nSCandidateListTouchBarItem = (NSCandidateListTouchBarItem*)ptr;
    NSCandidateListTouchBarItem* result_ = [nSCandidateListTouchBarItem initWithIdentifier:(NSString*)identifier];
    return result_;
}

void* C_NSCandidateListTouchBarItem_InitWithCoder(void* ptr, void* coder) {
    NSCandidateListTouchBarItem* nSCandidateListTouchBarItem = (NSCandidateListTouchBarItem*)ptr;
    NSCandidateListTouchBarItem* result_ = [nSCandidateListTouchBarItem initWithCoder:(NSCoder*)coder];
    return result_;
}

void* C_NSCandidateListTouchBarItem_AllocCandidateListTouchBarItem() {
    NSCandidateListTouchBarItem* result_ = [NSCandidateListTouchBarItem alloc];
    return result_;
}

void* C_NSCandidateListTouchBarItem_Autorelease(void* ptr) {
    NSCandidateListTouchBarItem* nSCandidateListTouchBarItem = (NSCandidateListTouchBarItem*)ptr;
    NSCandidateListTouchBarItem* result_ = [nSCandidateListTouchBarItem autorelease];
    return result_;
}

void* C_NSCandidateListTouchBarItem_Retain(void* ptr) {
    NSCandidateListTouchBarItem* nSCandidateListTouchBarItem = (NSCandidateListTouchBarItem*)ptr;
    NSCandidateListTouchBarItem* result_ = [nSCandidateListTouchBarItem retain];
    return result_;
}
