#import <Appkit/Appkit.h>
#import "pasteboard_item.h"

void* C_PasteboardItem_Alloc() {
    return [NSPasteboardItem alloc];
}

void* C_NSPasteboardItem_Init(void* ptr) {
    NSPasteboardItem* nSPasteboardItem = (NSPasteboardItem*)ptr;
    NSPasteboardItem* result_ = [nSPasteboardItem init];
    return result_;
}

void* C_NSPasteboardItem_AvailableTypeFromArray(void* ptr, Array types) {
    NSPasteboardItem* nSPasteboardItem = (NSPasteboardItem*)ptr;
    NSMutableArray* objcTypes = [[NSMutableArray alloc] init];
    if (types.len > 0) {
    	void** typesData = (void**)types.data;
    	for (int i = 0; i < types.len; i++) {
    		void* p = typesData[i];
    		[objcTypes addObject:(NSPasteboardType)(NSString*)p];
    	}
    }
    NSPasteboardType result_ = [nSPasteboardItem availableTypeFromArray:objcTypes];
    return result_;
}

bool C_NSPasteboardItem_SetDataProvider_ForTypes(void* ptr, void* dataProvider, Array types) {
    NSPasteboardItem* nSPasteboardItem = (NSPasteboardItem*)ptr;
    NSMutableArray* objcTypes = [[NSMutableArray alloc] init];
    if (types.len > 0) {
    	void** typesData = (void**)types.data;
    	for (int i = 0; i < types.len; i++) {
    		void* p = typesData[i];
    		[objcTypes addObject:(NSPasteboardType)(NSString*)p];
    	}
    }
    BOOL result_ = [nSPasteboardItem setDataProvider:(id)dataProvider forTypes:objcTypes];
    return result_;
}

bool C_NSPasteboardItem_SetData_ForType(void* ptr, void* data, void* _type) {
    NSPasteboardItem* nSPasteboardItem = (NSPasteboardItem*)ptr;
    BOOL result_ = [nSPasteboardItem setData:(NSData*)data forType:(NSString*)_type];
    return result_;
}

bool C_NSPasteboardItem_SetString_ForType(void* ptr, void* _string, void* _type) {
    NSPasteboardItem* nSPasteboardItem = (NSPasteboardItem*)ptr;
    BOOL result_ = [nSPasteboardItem setString:(NSString*)_string forType:(NSString*)_type];
    return result_;
}

bool C_NSPasteboardItem_SetPropertyList_ForType(void* ptr, void* propertyList, void* _type) {
    NSPasteboardItem* nSPasteboardItem = (NSPasteboardItem*)ptr;
    BOOL result_ = [nSPasteboardItem setPropertyList:(id)propertyList forType:(NSString*)_type];
    return result_;
}

void* C_NSPasteboardItem_DataForType(void* ptr, void* _type) {
    NSPasteboardItem* nSPasteboardItem = (NSPasteboardItem*)ptr;
    NSData* result_ = [nSPasteboardItem dataForType:(NSString*)_type];
    return result_;
}

void* C_NSPasteboardItem_StringForType(void* ptr, void* _type) {
    NSPasteboardItem* nSPasteboardItem = (NSPasteboardItem*)ptr;
    NSString* result_ = [nSPasteboardItem stringForType:(NSString*)_type];
    return result_;
}

void* C_NSPasteboardItem_PropertyListForType(void* ptr, void* _type) {
    NSPasteboardItem* nSPasteboardItem = (NSPasteboardItem*)ptr;
    id result_ = [nSPasteboardItem propertyListForType:(NSString*)_type];
    return result_;
}

Array C_NSPasteboardItem_Types(void* ptr) {
    NSPasteboardItem* nSPasteboardItem = (NSPasteboardItem*)ptr;
    NSArray* result_ = [nSPasteboardItem types];
    Array result_Array;
    int result_count = [result_ count];
    if (result_count > 0) {
    	void** result_Data = malloc(result_count * sizeof(void*));
    	for (int i = 0; i < result_count; i++) {
    		 void* p = [result_ objectAtIndex:i];
    		 result_Data[i] = p;
    	}
    	result_Array.data = result_Data;
    	result_Array.len = result_count;
    }
    return result_Array;
}
