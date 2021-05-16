#import <Appkit/Appkit.h>
#import "text_storage.h"

void* C_TextStorage_Alloc() {
    return [NSTextStorage alloc];
}

void* C_NSTextStorage_InitWithString(void* ptr, void* str) {
    NSTextStorage* nSTextStorage = (NSTextStorage*)ptr;
    NSTextStorage* result = [nSTextStorage initWithString:(NSString*)str];
    return result;
}

void* C_NSTextStorage_InitWithAttributedString(void* ptr, void* attrStr) {
    NSTextStorage* nSTextStorage = (NSTextStorage*)ptr;
    NSTextStorage* result = [nSTextStorage initWithAttributedString:(NSAttributedString*)attrStr];
    return result;
}

void* C_NSTextStorage_Init(void* ptr) {
    NSTextStorage* nSTextStorage = (NSTextStorage*)ptr;
    NSTextStorage* result = [nSTextStorage init];
    return result;
}

void C_NSTextStorage_AddLayoutManager(void* ptr, void* aLayoutManager) {
    NSTextStorage* nSTextStorage = (NSTextStorage*)ptr;
    [nSTextStorage addLayoutManager:(NSLayoutManager*)aLayoutManager];
}

void C_NSTextStorage_RemoveLayoutManager(void* ptr, void* aLayoutManager) {
    NSTextStorage* nSTextStorage = (NSTextStorage*)ptr;
    [nSTextStorage removeLayoutManager:(NSLayoutManager*)aLayoutManager];
}

void C_NSTextStorage_Edited_Range_ChangeInLength(void* ptr, unsigned int editedMask, NSRange editedRange, int delta) {
    NSTextStorage* nSTextStorage = (NSTextStorage*)ptr;
    [nSTextStorage edited:editedMask range:editedRange changeInLength:delta];
}

void C_NSTextStorage_ProcessEditing(void* ptr) {
    NSTextStorage* nSTextStorage = (NSTextStorage*)ptr;
    [nSTextStorage processEditing];
}

void C_NSTextStorage_InvalidateAttributesInRange(void* ptr, NSRange _range) {
    NSTextStorage* nSTextStorage = (NSTextStorage*)ptr;
    [nSTextStorage invalidateAttributesInRange:_range];
}

void C_NSTextStorage_EnsureAttributesAreFixedInRange(void* ptr, NSRange _range) {
    NSTextStorage* nSTextStorage = (NSTextStorage*)ptr;
    [nSTextStorage ensureAttributesAreFixedInRange:_range];
}

Array C_NSTextStorage_LayoutManagers(void* ptr) {
    NSTextStorage* nSTextStorage = (NSTextStorage*)ptr;
    NSArray* result = [nSTextStorage layoutManagers];
    int resultcount = [result count];
    void** resultData = malloc(resultcount * sizeof(void*));
    for (int i = 0; i < resultcount; i++) {
    	 void* p = [result objectAtIndex:i];
    	 resultData[i] = p;
    }
    Array resultArray;
    resultArray.data = resultData;
    resultArray.len = resultcount;
    return resultArray;
}

bool C_NSTextStorage_FixesAttributesLazily(void* ptr) {
    NSTextStorage* nSTextStorage = (NSTextStorage*)ptr;
    BOOL result = [nSTextStorage fixesAttributesLazily];
    return result;
}

unsigned int C_NSTextStorage_EditedMask(void* ptr) {
    NSTextStorage* nSTextStorage = (NSTextStorage*)ptr;
    NSTextStorageEditActions result = [nSTextStorage editedMask];
    return result;
}

NSRange C_NSTextStorage_EditedRange(void* ptr) {
    NSTextStorage* nSTextStorage = (NSTextStorage*)ptr;
    NSRange result = [nSTextStorage editedRange];
    return result;
}

int C_NSTextStorage_ChangeInLength(void* ptr) {
    NSTextStorage* nSTextStorage = (NSTextStorage*)ptr;
    NSInteger result = [nSTextStorage changeInLength];
    return result;
}

Array C_NSTextStorage_AttributeRuns(void* ptr) {
    NSTextStorage* nSTextStorage = (NSTextStorage*)ptr;
    NSArray* result = [nSTextStorage attributeRuns];
    int resultcount = [result count];
    void** resultData = malloc(resultcount * sizeof(void*));
    for (int i = 0; i < resultcount; i++) {
    	 void* p = [result objectAtIndex:i];
    	 resultData[i] = p;
    }
    Array resultArray;
    resultArray.data = resultData;
    resultArray.len = resultcount;
    return resultArray;
}

void C_NSTextStorage_SetAttributeRuns(void* ptr, Array value) {
    NSTextStorage* nSTextStorage = (NSTextStorage*)ptr;
    NSMutableArray* objcValue = [[NSMutableArray alloc] init];
    void** valueData = (void**)value.data;
    for (int i = 0; i < value.len; i++) {
    	void* p = valueData[i];
    	[objcValue addObject:(NSTextStorage*)(NSTextStorage*)p];
    }
    [nSTextStorage setAttributeRuns:objcValue];
}

Array C_NSTextStorage_Paragraphs(void* ptr) {
    NSTextStorage* nSTextStorage = (NSTextStorage*)ptr;
    NSArray* result = [nSTextStorage paragraphs];
    int resultcount = [result count];
    void** resultData = malloc(resultcount * sizeof(void*));
    for (int i = 0; i < resultcount; i++) {
    	 void* p = [result objectAtIndex:i];
    	 resultData[i] = p;
    }
    Array resultArray;
    resultArray.data = resultData;
    resultArray.len = resultcount;
    return resultArray;
}

void C_NSTextStorage_SetParagraphs(void* ptr, Array value) {
    NSTextStorage* nSTextStorage = (NSTextStorage*)ptr;
    NSMutableArray* objcValue = [[NSMutableArray alloc] init];
    void** valueData = (void**)value.data;
    for (int i = 0; i < value.len; i++) {
    	void* p = valueData[i];
    	[objcValue addObject:(NSTextStorage*)(NSTextStorage*)p];
    }
    [nSTextStorage setParagraphs:objcValue];
}

Array C_NSTextStorage_Words(void* ptr) {
    NSTextStorage* nSTextStorage = (NSTextStorage*)ptr;
    NSArray* result = [nSTextStorage words];
    int resultcount = [result count];
    void** resultData = malloc(resultcount * sizeof(void*));
    for (int i = 0; i < resultcount; i++) {
    	 void* p = [result objectAtIndex:i];
    	 resultData[i] = p;
    }
    Array resultArray;
    resultArray.data = resultData;
    resultArray.len = resultcount;
    return resultArray;
}

void C_NSTextStorage_SetWords(void* ptr, Array value) {
    NSTextStorage* nSTextStorage = (NSTextStorage*)ptr;
    NSMutableArray* objcValue = [[NSMutableArray alloc] init];
    void** valueData = (void**)value.data;
    for (int i = 0; i < value.len; i++) {
    	void* p = valueData[i];
    	[objcValue addObject:(NSTextStorage*)(NSTextStorage*)p];
    }
    [nSTextStorage setWords:objcValue];
}

Array C_NSTextStorage_Characters(void* ptr) {
    NSTextStorage* nSTextStorage = (NSTextStorage*)ptr;
    NSArray* result = [nSTextStorage characters];
    int resultcount = [result count];
    void** resultData = malloc(resultcount * sizeof(void*));
    for (int i = 0; i < resultcount; i++) {
    	 void* p = [result objectAtIndex:i];
    	 resultData[i] = p;
    }
    Array resultArray;
    resultArray.data = resultData;
    resultArray.len = resultcount;
    return resultArray;
}

void C_NSTextStorage_SetCharacters(void* ptr, Array value) {
    NSTextStorage* nSTextStorage = (NSTextStorage*)ptr;
    NSMutableArray* objcValue = [[NSMutableArray alloc] init];
    void** valueData = (void**)value.data;
    for (int i = 0; i < value.len; i++) {
    	void* p = valueData[i];
    	[objcValue addObject:(NSTextStorage*)(NSTextStorage*)p];
    }
    [nSTextStorage setCharacters:objcValue];
}

void* C_NSTextStorage_Font(void* ptr) {
    NSTextStorage* nSTextStorage = (NSTextStorage*)ptr;
    NSFont* result = [nSTextStorage font];
    return result;
}

void C_NSTextStorage_SetFont(void* ptr, void* value) {
    NSTextStorage* nSTextStorage = (NSTextStorage*)ptr;
    [nSTextStorage setFont:(NSFont*)value];
}

void* C_NSTextStorage_ForegroundColor(void* ptr) {
    NSTextStorage* nSTextStorage = (NSTextStorage*)ptr;
    NSColor* result = [nSTextStorage foregroundColor];
    return result;
}

void C_NSTextStorage_SetForegroundColor(void* ptr, void* value) {
    NSTextStorage* nSTextStorage = (NSTextStorage*)ptr;
    [nSTextStorage setForegroundColor:(NSColor*)value];
}
