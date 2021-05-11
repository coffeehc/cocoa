#import <stdbool.h>
#import <stdlib.h>
#import <utils.h>
#import <Appkit/Appkit.h>

void* C_Cursor_Alloc();

void* C_NSCursor_InitWithImage_HotSpot(void* ptr, void* newImage, CGPoint point);
void* C_NSCursor_InitWithCoder(void* ptr, void* coder);
void* C_NSCursor_Init(void* ptr);
void C_NSCursor_CursorHide();
void C_NSCursor_CursorUnhide();
void C_NSCursor_CursorSetHiddenUntilMouseMoves(bool flag);
void C_NSCursor_CursorPop();
void C_NSCursor_Pop(void* ptr);
void C_NSCursor_Push(void* ptr);
void C_NSCursor_Set(void* ptr);
void* C_NSCursor_Image(void* ptr);
CGPoint C_NSCursor_HotSpot(void* ptr);
