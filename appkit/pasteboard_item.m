#import <Appkit/Appkit.h>
#import "pasteboard_item.h"

void* C_PasteboardItem_Alloc() {
	return [NSPasteboardItem alloc];
}

void* C_NSPasteboardItem_Init(void* ptr) {
	NSPasteboardItem* nSPasteboardItem = (NSPasteboardItem*)ptr;
	NSPasteboardItem* result = [nSPasteboardItem init];
	return result;
}

bool C_NSPasteboardItem_SetData_ForType(void* ptr, Array data, void* _type) {
	NSPasteboardItem* nSPasteboardItem = (NSPasteboardItem*)ptr;
	bool result = [nSPasteboardItem setData:[[NSData alloc] initWithBytes:(Byte *)data.data length:data.len] forType:(NSString*)_type];
	return result;
}

bool C_NSPasteboardItem_SetString_ForType(void* ptr, void* _string, void* _type) {
	NSPasteboardItem* nSPasteboardItem = (NSPasteboardItem*)ptr;
	bool result = [nSPasteboardItem setString:(NSString*)_string forType:(NSString*)_type];
	return result;
}

bool C_NSPasteboardItem_SetPropertyList_ForType(void* ptr, void* propertyList, void* _type) {
	NSPasteboardItem* nSPasteboardItem = (NSPasteboardItem*)ptr;
	bool result = [nSPasteboardItem setPropertyList:(id)propertyList forType:(NSString*)_type];
	return result;
}

Array C_NSPasteboardItem_DataForType(void* ptr, void* _type) {
	NSPasteboardItem* nSPasteboardItem = (NSPasteboardItem*)ptr;
	NSData* result = [nSPasteboardItem dataForType:(NSString*)_type];
	Array resultarray;
	resultarray.data = [result bytes];
	resultarray.len = result.length;
	return resultarray;
}

void* C_NSPasteboardItem_StringForType(void* ptr, void* _type) {
	NSPasteboardItem* nSPasteboardItem = (NSPasteboardItem*)ptr;
	NSString* result = [nSPasteboardItem stringForType:(NSString*)_type];
	return result;
}

void* C_NSPasteboardItem_PropertyListForType(void* ptr, void* _type) {
	NSPasteboardItem* nSPasteboardItem = (NSPasteboardItem*)ptr;
	id result = [nSPasteboardItem propertyListForType:(NSString*)_type];
	return result;
}
