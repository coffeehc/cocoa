#import <Appkit/Appkit.h>
#import "pasteboard.h"

void* C_Pasteboard_Alloc() {
    return [NSPasteboard alloc];
}

void* C_NSPasteboard_Init(void* ptr) {
    NSPasteboard* nSPasteboard = (NSPasteboard*)ptr;
    NSPasteboard* result_ = [nSPasteboard init];
    return result_;
}

void* C_NSPasteboard_PasteboardByFilteringData_OfType(Array data, void* _type) {
    NSPasteboard* result_ = [NSPasteboard pasteboardByFilteringData:[[NSData alloc] initWithBytes:(Byte *)data.data length:data.len] ofType:(NSString*)_type];
    return result_;
}

void* C_NSPasteboard_PasteboardByFilteringFile(void* filename) {
    NSPasteboard* result_ = [NSPasteboard pasteboardByFilteringFile:(NSString*)filename];
    return result_;
}

void* C_NSPasteboard_PasteboardByFilteringTypesInPasteboard(void* pboard) {
    NSPasteboard* result_ = [NSPasteboard pasteboardByFilteringTypesInPasteboard:(NSPasteboard*)pboard];
    return result_;
}

void* C_NSPasteboard_PasteboardWithName(void* name) {
    NSPasteboard* result_ = [NSPasteboard pasteboardWithName:(NSString*)name];
    return result_;
}

void* C_NSPasteboard_PasteboardWithUniqueName() {
    NSPasteboard* result_ = [NSPasteboard pasteboardWithUniqueName];
    return result_;
}

int C_NSPasteboard_ClearContents(void* ptr) {
    NSPasteboard* nSPasteboard = (NSPasteboard*)ptr;
    NSInteger result_ = [nSPasteboard clearContents];
    return result_;
}

bool C_NSPasteboard_SetData_ForType(void* ptr, Array data, void* dataType) {
    NSPasteboard* nSPasteboard = (NSPasteboard*)ptr;
    BOOL result_ = [nSPasteboard setData:[[NSData alloc] initWithBytes:(Byte *)data.data length:data.len] forType:(NSString*)dataType];
    return result_;
}

bool C_NSPasteboard_SetPropertyList_ForType(void* ptr, void* plist, void* dataType) {
    NSPasteboard* nSPasteboard = (NSPasteboard*)ptr;
    BOOL result_ = [nSPasteboard setPropertyList:(id)plist forType:(NSString*)dataType];
    return result_;
}

bool C_NSPasteboard_SetString_ForType(void* ptr, void* _string, void* dataType) {
    NSPasteboard* nSPasteboard = (NSPasteboard*)ptr;
    BOOL result_ = [nSPasteboard setString:(NSString*)_string forType:(NSString*)dataType];
    return result_;
}

unsigned int C_NSPasteboard_IndexOfPasteboardItem(void* ptr, void* pasteboardItem) {
    NSPasteboard* nSPasteboard = (NSPasteboard*)ptr;
    NSUInteger result_ = [nSPasteboard indexOfPasteboardItem:(NSPasteboardItem*)pasteboardItem];
    return result_;
}

Array C_NSPasteboard_DataForType(void* ptr, void* dataType) {
    NSPasteboard* nSPasteboard = (NSPasteboard*)ptr;
    NSData* result_ = [nSPasteboard dataForType:(NSString*)dataType];
    Array result_array;
    result_array.data = [result_ bytes];
    result_array.len = result_.length;
    return result_array;
}

void* C_NSPasteboard_PropertyListForType(void* ptr, void* dataType) {
    NSPasteboard* nSPasteboard = (NSPasteboard*)ptr;
    id result_ = [nSPasteboard propertyListForType:(NSString*)dataType];
    return result_;
}

void* C_NSPasteboard_StringForType(void* ptr, void* dataType) {
    NSPasteboard* nSPasteboard = (NSPasteboard*)ptr;
    NSString* result_ = [nSPasteboard stringForType:(NSString*)dataType];
    return result_;
}

void* C_NSPasteboard_AvailableTypeFromArray(void* ptr, Array types) {
    NSPasteboard* nSPasteboard = (NSPasteboard*)ptr;
    NSMutableArray* objcTypes = [[NSMutableArray alloc] init];
    void** typesData = (void**)types.data;
    for (int i = 0; i < types.len; i++) {
    	void* p = typesData[i];
    	[objcTypes addObject:(NSPasteboardType)(NSString*)p];
    }
    NSPasteboardType result_ = [nSPasteboard availableTypeFromArray:objcTypes];
    return result_;
}

bool C_NSPasteboard_CanReadItemWithDataConformingToTypes(void* ptr, Array types) {
    NSPasteboard* nSPasteboard = (NSPasteboard*)ptr;
    NSMutableArray* objcTypes = [[NSMutableArray alloc] init];
    void** typesData = (void**)types.data;
    for (int i = 0; i < types.len; i++) {
    	void* p = typesData[i];
    	[objcTypes addObject:(NSString*)(NSString*)p];
    }
    BOOL result_ = [nSPasteboard canReadItemWithDataConformingToTypes:objcTypes];
    return result_;
}

Array C_NSPasteboard_Pasteboard_TypesFilterableTo(void* _type) {
    NSArray* result_ = [NSPasteboard typesFilterableTo:(NSString*)_type];
    int result_count = [result_ count];
    void** result_Data = malloc(result_count * sizeof(void*));
    for (int i = 0; i < result_count; i++) {
    	 void* p = [result_ objectAtIndex:i];
    	 result_Data[i] = p;
    }
    Array result_Array;
    result_Array.data = result_Data;
    result_Array.len = result_count;
    return result_Array;
}

int C_NSPasteboard_PrepareForNewContentsWithOptions(void* ptr, unsigned int options) {
    NSPasteboard* nSPasteboard = (NSPasteboard*)ptr;
    NSInteger result_ = [nSPasteboard prepareForNewContentsWithOptions:options];
    return result_;
}

int C_NSPasteboard_DeclareTypes_Owner(void* ptr, Array newTypes, void* newOwner) {
    NSPasteboard* nSPasteboard = (NSPasteboard*)ptr;
    NSMutableArray* objcNewTypes = [[NSMutableArray alloc] init];
    void** newTypesData = (void**)newTypes.data;
    for (int i = 0; i < newTypes.len; i++) {
    	void* p = newTypesData[i];
    	[objcNewTypes addObject:(NSPasteboardType)(NSString*)p];
    }
    NSInteger result_ = [nSPasteboard declareTypes:objcNewTypes owner:(id)newOwner];
    return result_;
}

int C_NSPasteboard_AddTypes_Owner(void* ptr, Array newTypes, void* newOwner) {
    NSPasteboard* nSPasteboard = (NSPasteboard*)ptr;
    NSMutableArray* objcNewTypes = [[NSMutableArray alloc] init];
    void** newTypesData = (void**)newTypes.data;
    for (int i = 0; i < newTypes.len; i++) {
    	void* p = newTypesData[i];
    	[objcNewTypes addObject:(NSPasteboardType)(NSString*)p];
    }
    NSInteger result_ = [nSPasteboard addTypes:objcNewTypes owner:(id)newOwner];
    return result_;
}

bool C_NSPasteboard_WriteFileContents(void* ptr, void* filename) {
    NSPasteboard* nSPasteboard = (NSPasteboard*)ptr;
    BOOL result_ = [nSPasteboard writeFileContents:(NSString*)filename];
    return result_;
}

bool C_NSPasteboard_WriteFileWrapper(void* ptr, void* wrapper) {
    NSPasteboard* nSPasteboard = (NSPasteboard*)ptr;
    BOOL result_ = [nSPasteboard writeFileWrapper:(NSFileWrapper*)wrapper];
    return result_;
}

void* C_NSPasteboard_ReadFileContentsType_ToFile(void* ptr, void* _type, void* filename) {
    NSPasteboard* nSPasteboard = (NSPasteboard*)ptr;
    NSString* result_ = [nSPasteboard readFileContentsType:(NSString*)_type toFile:(NSString*)filename];
    return result_;
}

void* C_NSPasteboard_ReadFileWrapper(void* ptr) {
    NSPasteboard* nSPasteboard = (NSPasteboard*)ptr;
    NSFileWrapper* result_ = [nSPasteboard readFileWrapper];
    return result_;
}

void* C_NSPasteboard_GeneralPasteboard() {
    NSPasteboard* result_ = [NSPasteboard generalPasteboard];
    return result_;
}

Array C_NSPasteboard_PasteboardItems(void* ptr) {
    NSPasteboard* nSPasteboard = (NSPasteboard*)ptr;
    NSArray* result_ = [nSPasteboard pasteboardItems];
    int result_count = [result_ count];
    void** result_Data = malloc(result_count * sizeof(void*));
    for (int i = 0; i < result_count; i++) {
    	 void* p = [result_ objectAtIndex:i];
    	 result_Data[i] = p;
    }
    Array result_Array;
    result_Array.data = result_Data;
    result_Array.len = result_count;
    return result_Array;
}

Array C_NSPasteboard_Types(void* ptr) {
    NSPasteboard* nSPasteboard = (NSPasteboard*)ptr;
    NSArray* result_ = [nSPasteboard types];
    int result_count = [result_ count];
    void** result_Data = malloc(result_count * sizeof(void*));
    for (int i = 0; i < result_count; i++) {
    	 void* p = [result_ objectAtIndex:i];
    	 result_Data[i] = p;
    }
    Array result_Array;
    result_Array.data = result_Data;
    result_Array.len = result_count;
    return result_Array;
}

void* C_NSPasteboard_Name(void* ptr) {
    NSPasteboard* nSPasteboard = (NSPasteboard*)ptr;
    NSPasteboardName result_ = [nSPasteboard name];
    return result_;
}

int C_NSPasteboard_ChangeCount(void* ptr) {
    NSPasteboard* nSPasteboard = (NSPasteboard*)ptr;
    NSInteger result_ = [nSPasteboard changeCount];
    return result_;
}
