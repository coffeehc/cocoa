#import <stdbool.h>
#import <stdlib.h>
#import <utils.h>
#import <Appkit/Appkit.h>

void* C_DraggingSession_Alloc();

void* C_NSDraggingSession_AllocDraggingSession();
void* C_NSDraggingSession_Init(void* ptr);
void* C_NSDraggingSession_NewDraggingSession();
void* C_NSDraggingSession_Autorelease(void* ptr);
void* C_NSDraggingSession_Retain(void* ptr);
void* C_NSDraggingSession_DraggingPasteboard(void* ptr);
bool C_NSDraggingSession_AnimatesToStartingPositionsOnCancelOrFail(void* ptr);
void C_NSDraggingSession_SetAnimatesToStartingPositionsOnCancelOrFail(void* ptr, bool value);
int C_NSDraggingSession_DraggingFormation(void* ptr);
void C_NSDraggingSession_SetDraggingFormation(void* ptr, int value);
int C_NSDraggingSession_DraggingSequenceNumber(void* ptr);
CGPoint C_NSDraggingSession_DraggingLocation(void* ptr);
int C_NSDraggingSession_DraggingLeaderIndex(void* ptr);
void C_NSDraggingSession_SetDraggingLeaderIndex(void* ptr, int value);
