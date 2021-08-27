#import <stdbool.h>
#import <stdlib.h>
#import <utils.h>
#import <Appkit/Appkit.h>

void* C_Cursor_Alloc();

void* C_NSCursor_InitWithImage_HotSpot(void* ptr, void* newImage, CGPoint point);
void* C_NSCursor_InitWithCoder(void* ptr, void* coder);
void* C_NSCursor_Init(void* ptr);
void C_NSCursor_Cursor_Hide();
void C_NSCursor_Cursor_Unhide();
void C_NSCursor_Cursor_SetHiddenUntilMouseMoves(bool flag);
void C_NSCursor_Cursor_Pop();
void C_NSCursor_Push(void* ptr);
void C_NSCursor_Set(void* ptr);
void* C_NSCursor_Image(void* ptr);
CGPoint C_NSCursor_HotSpot(void* ptr);
void* C_NSCursor_CurrentCursor();
void* C_NSCursor_CurrentSystemCursor();
void* C_NSCursor_ArrowCursor();
void* C_NSCursor_ContextualMenuCursor();
void* C_NSCursor_ClosedHandCursor();
void* C_NSCursor_CrosshairCursor();
void* C_NSCursor_DisappearingItemCursor();
void* C_NSCursor_DragCopyCursor();
void* C_NSCursor_DragLinkCursor();
void* C_NSCursor_IBeamCursor();
void* C_NSCursor_OpenHandCursor();
void* C_NSCursor_OperationNotAllowedCursor();
void* C_NSCursor_PointingHandCursor();
void* C_NSCursor_ResizeDownCursor();
void* C_NSCursor_ResizeLeftCursor();
void* C_NSCursor_ResizeLeftRightCursor();
void* C_NSCursor_ResizeRightCursor();
void* C_NSCursor_ResizeUpCursor();
void* C_NSCursor_ResizeUpDownCursor();
void* C_NSCursor_IBeamCursorForVerticalLayout();
