#import <Appkit/Appkit.h>
#import "mutable_paragraph_style.h"

void* C_MutableParagraphStyle_Alloc() {
    return [NSMutableParagraphStyle alloc];
}

void* C_NSMutableParagraphStyle_Init(void* ptr) {
    NSMutableParagraphStyle* nSMutableParagraphStyle = (NSMutableParagraphStyle*)ptr;
    NSMutableParagraphStyle* result_ = [nSMutableParagraphStyle init];
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
    NSMutableArray* objcValue = [[NSMutableArray alloc] init];
    if (value.len > 0) {
    	void** valueData = (void**)value.data;
    	for (int i = 0; i < value.len; i++) {
    		void* p = valueData[i];
    		[objcValue addObject:(NSTextTab*)(NSTextTab*)p];
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
    NSMutableArray* objcValue = [[NSMutableArray alloc] init];
    if (value.len > 0) {
    	void** valueData = (void**)value.data;
    	for (int i = 0; i < value.len; i++) {
    		void* p = valueData[i];
    		[objcValue addObject:(NSTextBlock*)(NSTextBlock*)p];
    	}
    }
    [nSMutableParagraphStyle setTextBlocks:objcValue];
}

void C_NSMutableParagraphStyle_SetTextLists(void* ptr, Array value) {
    NSMutableParagraphStyle* nSMutableParagraphStyle = (NSMutableParagraphStyle*)ptr;
    NSMutableArray* objcValue = [[NSMutableArray alloc] init];
    if (value.len > 0) {
    	void** valueData = (void**)value.data;
    	for (int i = 0; i < value.len; i++) {
    		void* p = valueData[i];
    		[objcValue addObject:(NSTextList*)(NSTextList*)p];
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
