#import <Foundation/NSGeometry.h>
#import <stdbool.h>
#import <stdlib.h>
#import <utils.h>

NSEdgeInsets ClipView_ContentInsets(void* ptr);
void ClipView_SetContentInsets(void* ptr, NSEdgeInsets contentInsets);
void* ClipView_DocumentView(void* ptr);
void ClipView_SetDocumentView(void* ptr, void* documentView);
bool ClipView_AutomaticallyAdjustsContentInsets(void* ptr);
void ClipView_SetAutomaticallyAdjustsContentInsets(void* ptr, bool automaticallyAdjustsContentInsets);
NSRect ClipView_DocumentRect(void* ptr);
NSRect ClipView_DocumentVisibleRect(void* ptr);
bool ClipView_DrawsBackground(void* ptr);
void ClipView_SetDrawsBackground(void* ptr, bool drawsBackground);
void* ClipView_BackgroundColor(void* ptr);
void ClipView_SetBackgroundColor(void* ptr, void* backgroundColor);

void ClipView_ScrollToPoint(void* ptr, NSPoint newOrigin);
void ClipView_Autoscroll(void* ptr, void* event);
NSRect ClipView_ConstrainBoundsRect(void* ptr, NSRect proposedBounds);
