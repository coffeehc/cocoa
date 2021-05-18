#import <Appkit/Appkit.h>
#import "candidate_list_touch_bar_item.h"

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
