#import <Foundation/NSGeometry.h>
#import <Foundation/NSRange.h>
#import <stdbool.h>
#import <stdint.h>
#import <stdlib.h>
#import <utils.h>

void* C_ClipView_Alloc();

void* C_NSClipView_InitWithFrame(void* ptr, CGRect frameRect);
void* C_NSClipView_InitWithCoder(void* ptr, void* coder);
void* C_NSClipView_Init(void* ptr);
void* C_NSClipView_AllocClipView();
void* C_NSClipView_NewClipView();
void* C_NSClipView_Autorelease(void* ptr);
void* C_NSClipView_Retain(void* ptr);
void C_NSClipView_ScrollToPoint(void* ptr, CGPoint newOrigin);
CGRect C_NSClipView_ConstrainBoundsRect(void* ptr, CGRect proposedBounds);
void C_NSClipView_ViewBoundsChanged(void* ptr, void* notification);
void C_NSClipView_ViewFrameChanged(void* ptr, void* notification);
void* C_NSClipView_DocumentView(void* ptr);
void C_NSClipView_SetDocumentView(void* ptr, void* value);
NSEdgeInsets C_NSClipView_ContentInsets(void* ptr);
void C_NSClipView_SetContentInsets(void* ptr, NSEdgeInsets value);
bool C_NSClipView_AutomaticallyAdjustsContentInsets(void* ptr);
void C_NSClipView_SetAutomaticallyAdjustsContentInsets(void* ptr, bool value);
CGRect C_NSClipView_DocumentRect(void* ptr);
CGRect C_NSClipView_DocumentVisibleRect(void* ptr);
void* C_NSClipView_DocumentCursor(void* ptr);
void C_NSClipView_SetDocumentCursor(void* ptr, void* value);
bool C_NSClipView_DrawsBackground(void* ptr);
void C_NSClipView_SetDrawsBackground(void* ptr, bool value);
void* C_NSClipView_BackgroundColor(void* ptr);
void C_NSClipView_SetBackgroundColor(void* ptr, void* value);
