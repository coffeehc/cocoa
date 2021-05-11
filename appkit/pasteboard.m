#import <Appkit/Appkit.h>
#import "pasteboard.h"

void* C_Pasteboard_Alloc() {
	return [NSPasteboard alloc];
}

void* C_NSPasteboard_Init(void* ptr) {
	NSPasteboard* nSPasteboard = (NSPasteboard*)ptr;
	NSPasteboard* result = [nSPasteboard init];
	return result;
}

void* C_NSPasteboard_PasteboardByFilteringData_OfType(Array data, void* _type) {
	NSPasteboard* result = [NSPasteboard pasteboardByFilteringData:[[NSData alloc] initWithBytes:(Byte *)data.data length:data.len] ofType:(NSString*)_type];
	return result;
}

void* C_NSPasteboard_PasteboardByFilteringFile(void* filename) {
	NSPasteboard* result = [NSPasteboard pasteboardByFilteringFile:(NSString*)filename];
	return result;
}

void* C_NSPasteboard_PasteboardByFilteringTypesInPasteboard(void* pboard) {
	NSPasteboard* result = [NSPasteboard pasteboardByFilteringTypesInPasteboard:(NSPasteboard*)pboard];
	return result;
}

void* C_NSPasteboard_PasteboardWithName(void* name) {
	NSPasteboard* result = [NSPasteboard pasteboardWithName:(NSString*)name];
	return result;
}

void* C_NSPasteboard_PasteboardWithUniqueName() {
	NSPasteboard* result = [NSPasteboard pasteboardWithUniqueName];
	return result;
}

int C_NSPasteboard_ClearContents(void* ptr) {
	NSPasteboard* nSPasteboard = (NSPasteboard*)ptr;
	int result = [nSPasteboard clearContents];
	return result;
}

bool C_NSPasteboard_SetData_ForType(void* ptr, Array data, void* dataType) {
	NSPasteboard* nSPasteboard = (NSPasteboard*)ptr;
	bool result = [nSPasteboard setData:[[NSData alloc] initWithBytes:(Byte *)data.data length:data.len] forType:(NSString*)dataType];
	return result;
}

bool C_NSPasteboard_SetPropertyList_ForType(void* ptr, void* plist, void* dataType) {
	NSPasteboard* nSPasteboard = (NSPasteboard*)ptr;
	bool result = [nSPasteboard setPropertyList:(id)plist forType:(NSString*)dataType];
	return result;
}

bool C_NSPasteboard_SetString_ForType(void* ptr, void* _string, void* dataType) {
	NSPasteboard* nSPasteboard = (NSPasteboard*)ptr;
	bool result = [nSPasteboard setString:(NSString*)_string forType:(NSString*)dataType];
	return result;
}

unsigned int C_NSPasteboard_IndexOfPasteboardItem(void* ptr, void* pasteboardItem) {
	NSPasteboard* nSPasteboard = (NSPasteboard*)ptr;
	unsigned int result = [nSPasteboard indexOfPasteboardItem:(NSPasteboardItem*)pasteboardItem];
	return result;
}

Array C_NSPasteboard_DataForType(void* ptr, void* dataType) {
	NSPasteboard* nSPasteboard = (NSPasteboard*)ptr;
	NSData* result = [nSPasteboard dataForType:(NSString*)dataType];
	Array resultarray;
	resultarray.data = [result bytes];
	resultarray.len = result.length;
	return resultarray;
}

void* C_NSPasteboard_PropertyListForType(void* ptr, void* dataType) {
	NSPasteboard* nSPasteboard = (NSPasteboard*)ptr;
	id result = [nSPasteboard propertyListForType:(NSString*)dataType];
	return result;
}

void* C_NSPasteboard_StringForType(void* ptr, void* dataType) {
	NSPasteboard* nSPasteboard = (NSPasteboard*)ptr;
	NSString* result = [nSPasteboard stringForType:(NSString*)dataType];
	return result;
}

bool C_NSPasteboard_WriteFileContents(void* ptr, void* filename) {
	NSPasteboard* nSPasteboard = (NSPasteboard*)ptr;
	bool result = [nSPasteboard writeFileContents:(NSString*)filename];
	return result;
}

bool C_NSPasteboard_WriteFileWrapper(void* ptr, void* wrapper) {
	NSPasteboard* nSPasteboard = (NSPasteboard*)ptr;
	bool result = [nSPasteboard writeFileWrapper:(NSFileWrapper*)wrapper];
	return result;
}

void* C_NSPasteboard_ReadFileContentsType_ToFile(void* ptr, void* _type, void* filename) {
	NSPasteboard* nSPasteboard = (NSPasteboard*)ptr;
	NSString* result = [nSPasteboard readFileContentsType:(NSString*)_type toFile:(NSString*)filename];
	return result;
}

void* C_NSPasteboard_ReadFileWrapper(void* ptr) {
	NSPasteboard* nSPasteboard = (NSPasteboard*)ptr;
	NSFileWrapper* result = [nSPasteboard readFileWrapper];
	return result;
}

void* C_NSPasteboard_Name(void* ptr) {
	NSPasteboard* nSPasteboard = (NSPasteboard*)ptr;
	NSString* result = [nSPasteboard name];
	return result;
}

int C_NSPasteboard_ChangeCount(void* ptr) {
	NSPasteboard* nSPasteboard = (NSPasteboard*)ptr;
	int result = [nSPasteboard changeCount];
	return result;
}
