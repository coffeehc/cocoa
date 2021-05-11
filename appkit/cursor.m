#import <Appkit/Appkit.h>
#import "cursor.h"

void* C_Cursor_Alloc() {
	return [NSCursor alloc];
}

void* C_NSCursor_InitWithImage_HotSpot(void* ptr, void* newImage, CGPoint point) {
	NSCursor* nSCursor = (NSCursor*)ptr;
	NSCursor* result = [nSCursor initWithImage:(NSImage*)newImage hotSpot:point];
	return result;
}

void* C_NSCursor_InitWithCoder(void* ptr, void* coder) {
	NSCursor* nSCursor = (NSCursor*)ptr;
	NSCursor* result = [nSCursor initWithCoder:(NSCoder*)coder];
	return result;
}

void* C_NSCursor_Init(void* ptr) {
	NSCursor* nSCursor = (NSCursor*)ptr;
	NSCursor* result = [nSCursor init];
	return result;
}

void C_NSCursor_CursorHide() {
	[NSCursor hide];
}

void C_NSCursor_CursorUnhide() {
	[NSCursor unhide];
}

void C_NSCursor_CursorSetHiddenUntilMouseMoves(bool flag) {
	[NSCursor setHiddenUntilMouseMoves:flag];
}

void C_NSCursor_CursorPop() {
	[NSCursor pop];
}

void C_NSCursor_Pop(void* ptr) {
	NSCursor* nSCursor = (NSCursor*)ptr;
	[nSCursor pop];
}

void C_NSCursor_Push(void* ptr) {
	NSCursor* nSCursor = (NSCursor*)ptr;
	[nSCursor push];
}

void C_NSCursor_Set(void* ptr) {
	NSCursor* nSCursor = (NSCursor*)ptr;
	[nSCursor set];
}

void* C_NSCursor_Image(void* ptr) {
	NSCursor* nSCursor = (NSCursor*)ptr;
	NSImage* result = [nSCursor image];
	return result;
}

CGPoint C_NSCursor_HotSpot(void* ptr) {
	NSCursor* nSCursor = (NSCursor*)ptr;
	CGPoint result = [nSCursor hotSpot];
	return result;
}
