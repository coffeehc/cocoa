#import "paragraph_style.h"
#import <AppKit/NSParagraphStyle.h>
#import <Foundation/NSArray.h>
#import <Foundation/NSDictionary.h>

void* C_ParagraphStyle_Alloc() {
    return [NSParagraphStyle alloc];
}

void* C_NSParagraphStyle_AllocParagraphStyle() {
    NSParagraphStyle* result_ = [NSParagraphStyle alloc];
    return result_;
}

void* C_NSParagraphStyle_Init(void* ptr) {
    NSParagraphStyle* nSParagraphStyle = (NSParagraphStyle*)ptr;
    NSParagraphStyle* result_ = [nSParagraphStyle init];
    return result_;
}

void* C_NSParagraphStyle_NewParagraphStyle() {
    NSParagraphStyle* result_ = [NSParagraphStyle new];
    return result_;
}

void* C_NSParagraphStyle_Autorelease(void* ptr) {
    NSParagraphStyle* nSParagraphStyle = (NSParagraphStyle*)ptr;
    NSParagraphStyle* result_ = [nSParagraphStyle autorelease];
    return result_;
}

void* C_NSParagraphStyle_Retain(void* ptr) {
    NSParagraphStyle* nSParagraphStyle = (NSParagraphStyle*)ptr;
    NSParagraphStyle* result_ = [nSParagraphStyle retain];
    return result_;
}

int C_NSParagraphStyle_ParagraphStyle_DefaultWritingDirectionForLanguage(void* languageName) {
    NSWritingDirection result_ = [NSParagraphStyle defaultWritingDirectionForLanguage:(NSString*)languageName];
    return result_;
}

void* C_NSParagraphStyle_DefaultParagraphStyle() {
    NSParagraphStyle* result_ = [NSParagraphStyle defaultParagraphStyle];
    return result_;
}

int C_NSParagraphStyle_Alignment(void* ptr) {
    NSParagraphStyle* nSParagraphStyle = (NSParagraphStyle*)ptr;
    NSTextAlignment result_ = [nSParagraphStyle alignment];
    return result_;
}

double C_NSParagraphStyle_FirstLineHeadIndent(void* ptr) {
    NSParagraphStyle* nSParagraphStyle = (NSParagraphStyle*)ptr;
    CGFloat result_ = [nSParagraphStyle firstLineHeadIndent];
    return result_;
}

double C_NSParagraphStyle_HeadIndent(void* ptr) {
    NSParagraphStyle* nSParagraphStyle = (NSParagraphStyle*)ptr;
    CGFloat result_ = [nSParagraphStyle headIndent];
    return result_;
}

double C_NSParagraphStyle_TailIndent(void* ptr) {
    NSParagraphStyle* nSParagraphStyle = (NSParagraphStyle*)ptr;
    CGFloat result_ = [nSParagraphStyle tailIndent];
    return result_;
}

double C_NSParagraphStyle_LineHeightMultiple(void* ptr) {
    NSParagraphStyle* nSParagraphStyle = (NSParagraphStyle*)ptr;
    CGFloat result_ = [nSParagraphStyle lineHeightMultiple];
    return result_;
}

double C_NSParagraphStyle_MaximumLineHeight(void* ptr) {
    NSParagraphStyle* nSParagraphStyle = (NSParagraphStyle*)ptr;
    CGFloat result_ = [nSParagraphStyle maximumLineHeight];
    return result_;
}

double C_NSParagraphStyle_MinimumLineHeight(void* ptr) {
    NSParagraphStyle* nSParagraphStyle = (NSParagraphStyle*)ptr;
    CGFloat result_ = [nSParagraphStyle minimumLineHeight];
    return result_;
}

double C_NSParagraphStyle_LineSpacing(void* ptr) {
    NSParagraphStyle* nSParagraphStyle = (NSParagraphStyle*)ptr;
    CGFloat result_ = [nSParagraphStyle lineSpacing];
    return result_;
}

double C_NSParagraphStyle_ParagraphSpacing(void* ptr) {
    NSParagraphStyle* nSParagraphStyle = (NSParagraphStyle*)ptr;
    CGFloat result_ = [nSParagraphStyle paragraphSpacing];
    return result_;
}

double C_NSParagraphStyle_ParagraphSpacingBefore(void* ptr) {
    NSParagraphStyle* nSParagraphStyle = (NSParagraphStyle*)ptr;
    CGFloat result_ = [nSParagraphStyle paragraphSpacingBefore];
    return result_;
}

Array C_NSParagraphStyle_TabStops(void* ptr) {
    NSParagraphStyle* nSParagraphStyle = (NSParagraphStyle*)ptr;
    NSArray* result_ = [nSParagraphStyle tabStops];
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

double C_NSParagraphStyle_DefaultTabInterval(void* ptr) {
    NSParagraphStyle* nSParagraphStyle = (NSParagraphStyle*)ptr;
    CGFloat result_ = [nSParagraphStyle defaultTabInterval];
    return result_;
}

Array C_NSParagraphStyle_TextBlocks(void* ptr) {
    NSParagraphStyle* nSParagraphStyle = (NSParagraphStyle*)ptr;
    NSArray* result_ = [nSParagraphStyle textBlocks];
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

Array C_NSParagraphStyle_TextLists(void* ptr) {
    NSParagraphStyle* nSParagraphStyle = (NSParagraphStyle*)ptr;
    NSArray* result_ = [nSParagraphStyle textLists];
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

unsigned int C_NSParagraphStyle_LineBreakMode(void* ptr) {
    NSParagraphStyle* nSParagraphStyle = (NSParagraphStyle*)ptr;
    NSLineBreakMode result_ = [nSParagraphStyle lineBreakMode];
    return result_;
}

unsigned int C_NSParagraphStyle_LineBreakStrategy(void* ptr) {
    NSParagraphStyle* nSParagraphStyle = (NSParagraphStyle*)ptr;
    NSLineBreakStrategy result_ = [nSParagraphStyle lineBreakStrategy];
    return result_;
}

float C_NSParagraphStyle_HyphenationFactor(void* ptr) {
    NSParagraphStyle* nSParagraphStyle = (NSParagraphStyle*)ptr;
    float result_ = [nSParagraphStyle hyphenationFactor];
    return result_;
}

float C_NSParagraphStyle_TighteningFactorForTruncation(void* ptr) {
    NSParagraphStyle* nSParagraphStyle = (NSParagraphStyle*)ptr;
    float result_ = [nSParagraphStyle tighteningFactorForTruncation];
    return result_;
}

bool C_NSParagraphStyle_AllowsDefaultTighteningForTruncation(void* ptr) {
    NSParagraphStyle* nSParagraphStyle = (NSParagraphStyle*)ptr;
    BOOL result_ = [nSParagraphStyle allowsDefaultTighteningForTruncation];
    return result_;
}

int C_NSParagraphStyle_HeaderLevel(void* ptr) {
    NSParagraphStyle* nSParagraphStyle = (NSParagraphStyle*)ptr;
    NSInteger result_ = [nSParagraphStyle headerLevel];
    return result_;
}

int C_NSParagraphStyle_BaseWritingDirection(void* ptr) {
    NSParagraphStyle* nSParagraphStyle = (NSParagraphStyle*)ptr;
    NSWritingDirection result_ = [nSParagraphStyle baseWritingDirection];
    return result_;
}

void* C_MutableParagraphStyle_Alloc() {
    return [NSMutableParagraphStyle alloc];
}

void* C_NSMutableParagraphStyle_AllocMutableParagraphStyle() {
    NSMutableParagraphStyle* result_ = [NSMutableParagraphStyle alloc];
    return result_;
}

void* C_NSMutableParagraphStyle_Init(void* ptr) {
    NSMutableParagraphStyle* nSMutableParagraphStyle = (NSMutableParagraphStyle*)ptr;
    NSMutableParagraphStyle* result_ = [nSMutableParagraphStyle init];
    return result_;
}

void* C_NSMutableParagraphStyle_NewMutableParagraphStyle() {
    NSMutableParagraphStyle* result_ = [NSMutableParagraphStyle new];
    return result_;
}

void* C_NSMutableParagraphStyle_Autorelease(void* ptr) {
    NSMutableParagraphStyle* nSMutableParagraphStyle = (NSMutableParagraphStyle*)ptr;
    NSMutableParagraphStyle* result_ = [nSMutableParagraphStyle autorelease];
    return result_;
}

void* C_NSMutableParagraphStyle_Retain(void* ptr) {
    NSMutableParagraphStyle* nSMutableParagraphStyle = (NSMutableParagraphStyle*)ptr;
    NSMutableParagraphStyle* result_ = [nSMutableParagraphStyle retain];
    return result_;
}

void C_NSMutableParagraphStyle_SetParagraphStyle(void* ptr, void* obj) {
    NSMutableParagraphStyle* nSMutableParagraphStyle = (NSMutableParagraphStyle*)ptr;
    [nSMutableParagraphStyle setParagraphStyle:(NSParagraphStyle*)obj];
}

void C_NSMutableParagraphStyle_AddTabStop(void* ptr, void* anObject) {
    NSMutableParagraphStyle* nSMutableParagraphStyle = (NSMutableParagraphStyle*)ptr;
    [nSMutableParagraphStyle addTabStop:(NSTextTab*)anObject];
}

void C_NSMutableParagraphStyle_RemoveTabStop(void* ptr, void* anObject) {
    NSMutableParagraphStyle* nSMutableParagraphStyle = (NSMutableParagraphStyle*)ptr;
    [nSMutableParagraphStyle removeTabStop:(NSTextTab*)anObject];
}

void C_NSMutableParagraphStyle_SetAlignment(void* ptr, int value) {
    NSMutableParagraphStyle* nSMutableParagraphStyle = (NSMutableParagraphStyle*)ptr;
    [nSMutableParagraphStyle setAlignment:value];
}

void C_NSMutableParagraphStyle_SetFirstLineHeadIndent(void* ptr, double value) {
    NSMutableParagraphStyle* nSMutableParagraphStyle = (NSMutableParagraphStyle*)ptr;
    [nSMutableParagraphStyle setFirstLineHeadIndent:value];
}

void C_NSMutableParagraphStyle_SetHeadIndent(void* ptr, double value) {
    NSMutableParagraphStyle* nSMutableParagraphStyle = (NSMutableParagraphStyle*)ptr;
    [nSMutableParagraphStyle setHeadIndent:value];
}

void C_NSMutableParagraphStyle_SetTailIndent(void* ptr, double value) {
    NSMutableParagraphStyle* nSMutableParagraphStyle = (NSMutableParagraphStyle*)ptr;
    [nSMutableParagraphStyle setTailIndent:value];
}

void C_NSMutableParagraphStyle_SetLineHeightMultiple(void* ptr, double value) {
    NSMutableParagraphStyle* nSMutableParagraphStyle = (NSMutableParagraphStyle*)ptr;
    [nSMutableParagraphStyle setLineHeightMultiple:value];
}

void C_NSMutableParagraphStyle_SetMaximumLineHeight(void* ptr, double value) {
    NSMutableParagraphStyle* nSMutableParagraphStyle = (NSMutableParagraphStyle*)ptr;
    [nSMutableParagraphStyle setMaximumLineHeight:value];
}

void C_NSMutableParagraphStyle_SetMinimumLineHeight(void* ptr, double value) {
    NSMutableParagraphStyle* nSMutableParagraphStyle = (NSMutableParagraphStyle*)ptr;
    [nSMutableParagraphStyle setMinimumLineHeight:value];
}

void C_NSMutableParagraphStyle_SetLineSpacing(void* ptr, double value) {
    NSMutableParagraphStyle* nSMutableParagraphStyle = (NSMutableParagraphStyle*)ptr;
    [nSMutableParagraphStyle setLineSpacing:value];
}

void C_NSMutableParagraphStyle_SetParagraphSpacing(void* ptr, double value) {
    NSMutableParagraphStyle* nSMutableParagraphStyle = (NSMutableParagraphStyle*)ptr;
    [nSMutableParagraphStyle setParagraphSpacing:value];
}

void C_NSMutableParagraphStyle_SetParagraphSpacingBefore(void* ptr, double value) {
    NSMutableParagraphStyle* nSMutableParagraphStyle = (NSMutableParagraphStyle*)ptr;
    [nSMutableParagraphStyle setParagraphSpacingBefore:value];
}

void C_NSMutableParagraphStyle_SetBaseWritingDirection(void* ptr, int value) {
    NSMutableParagraphStyle* nSMutableParagraphStyle = (NSMutableParagraphStyle*)ptr;
    [nSMutableParagraphStyle setBaseWritingDirection:value];
}

void C_NSMutableParagraphStyle_SetTabStops(void* ptr, Array value) {
    NSMutableParagraphStyle* nSMutableParagraphStyle = (NSMutableParagraphStyle*)ptr;
    NSMutableArray* objcValue = [NSMutableArray arrayWithCapacity:value.len];
    if (value.len > 0) {
    	void** valueData = (void**)value.data;
    	for (int i = 0; i < value.len; i++) {
    		void* p = valueData[i];
    		[objcValue addObject:(NSTextTab*)p];
    	}
    }
    [nSMutableParagraphStyle setTabStops:objcValue];
}

void C_NSMutableParagraphStyle_SetDefaultTabInterval(void* ptr, double value) {
    NSMutableParagraphStyle* nSMutableParagraphStyle = (NSMutableParagraphStyle*)ptr;
    [nSMutableParagraphStyle setDefaultTabInterval:value];
}

void C_NSMutableParagraphStyle_SetTextBlocks(void* ptr, Array value) {
    NSMutableParagraphStyle* nSMutableParagraphStyle = (NSMutableParagraphStyle*)ptr;
    NSMutableArray* objcValue = [NSMutableArray arrayWithCapacity:value.len];
    if (value.len > 0) {
    	void** valueData = (void**)value.data;
    	for (int i = 0; i < value.len; i++) {
    		void* p = valueData[i];
    		[objcValue addObject:(NSTextBlock*)p];
    	}
    }
    [nSMutableParagraphStyle setTextBlocks:objcValue];
}

void C_NSMutableParagraphStyle_SetTextLists(void* ptr, Array value) {
    NSMutableParagraphStyle* nSMutableParagraphStyle = (NSMutableParagraphStyle*)ptr;
    NSMutableArray* objcValue = [NSMutableArray arrayWithCapacity:value.len];
    if (value.len > 0) {
    	void** valueData = (void**)value.data;
    	for (int i = 0; i < value.len; i++) {
    		void* p = valueData[i];
    		[objcValue addObject:(NSTextList*)p];
    	}
    }
    [nSMutableParagraphStyle setTextLists:objcValue];
}

void C_NSMutableParagraphStyle_SetLineBreakMode(void* ptr, unsigned int value) {
    NSMutableParagraphStyle* nSMutableParagraphStyle = (NSMutableParagraphStyle*)ptr;
    [nSMutableParagraphStyle setLineBreakMode:value];
}

void C_NSMutableParagraphStyle_SetLineBreakStrategy(void* ptr, unsigned int value) {
    NSMutableParagraphStyle* nSMutableParagraphStyle = (NSMutableParagraphStyle*)ptr;
    [nSMutableParagraphStyle setLineBreakStrategy:value];
}

void C_NSMutableParagraphStyle_SetHyphenationFactor(void* ptr, float value) {
    NSMutableParagraphStyle* nSMutableParagraphStyle = (NSMutableParagraphStyle*)ptr;
    [nSMutableParagraphStyle setHyphenationFactor:value];
}

void C_NSMutableParagraphStyle_SetTighteningFactorForTruncation(void* ptr, float value) {
    NSMutableParagraphStyle* nSMutableParagraphStyle = (NSMutableParagraphStyle*)ptr;
    [nSMutableParagraphStyle setTighteningFactorForTruncation:value];
}

void C_NSMutableParagraphStyle_SetAllowsDefaultTighteningForTruncation(void* ptr, bool value) {
    NSMutableParagraphStyle* nSMutableParagraphStyle = (NSMutableParagraphStyle*)ptr;
    [nSMutableParagraphStyle setAllowsDefaultTighteningForTruncation:value];
}

void C_NSMutableParagraphStyle_SetHeaderLevel(void* ptr, int value) {
    NSMutableParagraphStyle* nSMutableParagraphStyle = (NSMutableParagraphStyle*)ptr;
    [nSMutableParagraphStyle setHeaderLevel:value];
}

void* C_TextTab_Alloc() {
    return [NSTextTab alloc];
}

void* C_NSTextTab_InitWithTextAlignment_Location_Options(void* ptr, int alignment, double loc, Dictionary options) {
    NSTextTab* nSTextTab = (NSTextTab*)ptr;
    NSMutableDictionary* objcOptions = [NSMutableDictionary dictionaryWithCapacity:options.len];
    if (options.len > 0) {
    	void** optionsKeyData = (void**)options.key_data;
    	void** optionsValueData = (void**)options.value_data;
    	for (int i = 0; i < options.len; i++) {
    		void* kp = optionsKeyData[i];
    		void* vp = optionsValueData[i];
    		[objcOptions setObject:(id)vp forKey:(id<NSCopying>)(NSString*)kp];
    	}
    }
    NSTextTab* result_ = [nSTextTab initWithTextAlignment:alignment location:loc options:objcOptions];
    return result_;
}

void* C_NSTextTab_AllocTextTab() {
    NSTextTab* result_ = [NSTextTab alloc];
    return result_;
}

void* C_NSTextTab_Init(void* ptr) {
    NSTextTab* nSTextTab = (NSTextTab*)ptr;
    NSTextTab* result_ = [nSTextTab init];
    return result_;
}

void* C_NSTextTab_NewTextTab() {
    NSTextTab* result_ = [NSTextTab new];
    return result_;
}

void* C_NSTextTab_Autorelease(void* ptr) {
    NSTextTab* nSTextTab = (NSTextTab*)ptr;
    NSTextTab* result_ = [nSTextTab autorelease];
    return result_;
}

void* C_NSTextTab_Retain(void* ptr) {
    NSTextTab* nSTextTab = (NSTextTab*)ptr;
    NSTextTab* result_ = [nSTextTab retain];
    return result_;
}

void* C_NSTextTab_TextTab_ColumnTerminatorsForLocale(void* aLocale) {
    NSCharacterSet* result_ = [NSTextTab columnTerminatorsForLocale:(NSLocale*)aLocale];
    return result_;
}

double C_NSTextTab_Location(void* ptr) {
    NSTextTab* nSTextTab = (NSTextTab*)ptr;
    CGFloat result_ = [nSTextTab location];
    return result_;
}

int C_NSTextTab_Alignment(void* ptr) {
    NSTextTab* nSTextTab = (NSTextTab*)ptr;
    NSTextAlignment result_ = [nSTextTab alignment];
    return result_;
}

Dictionary C_NSTextTab_Options(void* ptr) {
    NSTextTab* nSTextTab = (NSTextTab*)ptr;
    NSDictionary* result_ = [nSTextTab options];
    Dictionary result_Array;
    NSArray * result_Keys = [result_ allKeys];
    int result_Count = [result_Keys count];
    if (result_Count > 0) {
    	void** result_KeyData = malloc(result_Count * sizeof(void*));
    	void** result_ValueData = malloc(result_Count * sizeof(void*));
    	for (int i = 0; i < result_Count; i++) {
    		NSTextTabOptionKey kp = [result_Keys objectAtIndex:i];
    		id vp = result_[kp];
    		 result_KeyData[i] = kp;
    		 result_ValueData[i] = vp;
    	}
    	result_Array.key_data = result_KeyData;
    	result_Array.value_data = result_ValueData;
    	result_Array.len = result_Count;
    }
    return result_Array;
}
