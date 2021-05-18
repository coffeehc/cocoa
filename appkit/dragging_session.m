#import <Appkit/Appkit.h>
#import "dragging_session.h"

void* C_DraggingSession_Alloc() {
    return [NSDraggingSession alloc];
}

void* C_NSDraggingSession_Init(void* ptr) {
    NSDraggingSession* nSDraggingSession = (NSDraggingSession*)ptr;
    NSDraggingSession* result_ = [nSDraggingSession init];
    return result_;
}

void* C_NSDraggingSession_DraggingPasteboard(void* ptr) {
    NSDraggingSession* nSDraggingSession = (NSDraggingSession*)ptr;
    NSPasteboard* result_ = [nSDraggingSession draggingPasteboard];
    return result_;
}

bool C_NSDraggingSession_AnimatesToStartingPositionsOnCancelOrFail(void* ptr) {
    NSDraggingSession* nSDraggingSession = (NSDraggingSession*)ptr;
    BOOL result_ = [nSDraggingSession animatesToStartingPositionsOnCancelOrFail];
    return result_;
}

void C_NSDraggingSession_SetAnimatesToStartingPositionsOnCancelOrFail(void* ptr, bool value) {
    NSDraggingSession* nSDraggingSession = (NSDraggingSession*)ptr;
    [nSDraggingSession setAnimatesToStartingPositionsOnCancelOrFail:value];
}

int C_NSDraggingSession_DraggingFormation(void* ptr) {
    NSDraggingSession* nSDraggingSession = (NSDraggingSession*)ptr;
    NSDraggingFormation result_ = [nSDraggingSession draggingFormation];
    return result_;
}

void C_NSDraggingSession_SetDraggingFormation(void* ptr, int value) {
    NSDraggingSession* nSDraggingSession = (NSDraggingSession*)ptr;
    [nSDraggingSession setDraggingFormation:value];
}

int C_NSDraggingSession_DraggingSequenceNumber(void* ptr) {
    NSDraggingSession* nSDraggingSession = (NSDraggingSession*)ptr;
    NSInteger result_ = [nSDraggingSession draggingSequenceNumber];
    return result_;
}

CGPoint C_NSDraggingSession_DraggingLocation(void* ptr) {
    NSDraggingSession* nSDraggingSession = (NSDraggingSession*)ptr;
    NSPoint result_ = [nSDraggingSession draggingLocation];
    return result_;
}

int C_NSDraggingSession_DraggingLeaderIndex(void* ptr) {
    NSDraggingSession* nSDraggingSession = (NSDraggingSession*)ptr;
    NSInteger result_ = [nSDraggingSession draggingLeaderIndex];
    return result_;
}

void C_NSDraggingSession_SetDraggingLeaderIndex(void* ptr, int value) {
    NSDraggingSession* nSDraggingSession = (NSDraggingSession*)ptr;
    [nSDraggingSession setDraggingLeaderIndex:value];
}
