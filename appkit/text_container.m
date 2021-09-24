#import "text_container.h"
#import <AppKit/NSTextContainer.h>
#import <Foundation/NSArray.h>
#import <Foundation/NSDictionary.h>

void* C_TextContainer_Alloc() {
    return [NSTextContainer alloc];
}

void* C_NSTextContainer_InitWithSize(void* ptr, CGSize size) {
    NSTextContainer* nSTextContainer = (NSTextContainer*)ptr;
    NSTextContainer* result_ = [nSTextContainer initWithSize:size];
    return result_;
}

void* C_NSTextContainer_InitWithCoder(void* ptr, void* coder) {
    NSTextContainer* nSTextContainer = (NSTextContainer*)ptr;
    NSTextContainer* result_ = [nSTextContainer initWithCoder:(NSCoder*)coder];
    return result_;
}

void* C_NSTextContainer_AllocTextContainer() {
    NSTextContainer* result_ = [NSTextContainer alloc];
    return result_;
}

void* C_NSTextContainer_Init(void* ptr) {
    NSTextContainer* nSTextContainer = (NSTextContainer*)ptr;
    NSTextContainer* result_ = [nSTextContainer init];
    return result_;
}

void* C_NSTextContainer_NewTextContainer() {
    NSTextContainer* result_ = [NSTextContainer new];
    return result_;
}

void* C_NSTextContainer_Autorelease(void* ptr) {
    NSTextContainer* nSTextContainer = (NSTextContainer*)ptr;
    NSTextContainer* result_ = [nSTextContainer autorelease];
    return result_;
}

void* C_NSTextContainer_Retain(void* ptr) {
    NSTextContainer* nSTextContainer = (NSTextContainer*)ptr;
    NSTextContainer* result_ = [nSTextContainer retain];
    return result_;
}

void C_NSTextContainer_ReplaceLayoutManager(void* ptr, void* newLayoutManager) {
    NSTextContainer* nSTextContainer = (NSTextContainer*)ptr;
    [nSTextContainer replaceLayoutManager:(NSLayoutManager*)newLayoutManager];
}

void* C_NSTextContainer_LayoutManager(void* ptr) {
    NSTextContainer* nSTextContainer = (NSTextContainer*)ptr;
    NSLayoutManager* result_ = [nSTextContainer layoutManager];
    return result_;
}

void C_NSTextContainer_SetLayoutManager(void* ptr, void* value) {
    NSTextContainer* nSTextContainer = (NSTextContainer*)ptr;
    [nSTextContainer setLayoutManager:(NSLayoutManager*)value];
}

void* C_NSTextContainer_TextView(void* ptr) {
    NSTextContainer* nSTextContainer = (NSTextContainer*)ptr;
    NSTextView* result_ = [nSTextContainer textView];
    return result_;
}

void C_NSTextContainer_SetTextView(void* ptr, void* value) {
    NSTextContainer* nSTextContainer = (NSTextContainer*)ptr;
    [nSTextContainer setTextView:(NSTextView*)value];
}

CGSize C_NSTextContainer_Size(void* ptr) {
    NSTextContainer* nSTextContainer = (NSTextContainer*)ptr;
    NSSize result_ = [nSTextContainer size];
    return result_;
}

void C_NSTextContainer_SetSize(void* ptr, CGSize value) {
    NSTextContainer* nSTextContainer = (NSTextContainer*)ptr;
    [nSTextContainer setSize:value];
}

Array C_NSTextContainer_ExclusionPaths(void* ptr) {
    NSTextContainer* nSTextContainer = (NSTextContainer*)ptr;
    NSArray* result_ = [nSTextContainer exclusionPaths];
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

void C_NSTextContainer_SetExclusionPaths(void* ptr, Array value) {
    NSTextContainer* nSTextContainer = (NSTextContainer*)ptr;
    NSMutableArray* objcValue = [NSMutableArray arrayWithCapacity:value.len];
    if (value.len > 0) {
    	void** valueData = (void**)value.data;
    	for (int i = 0; i < value.len; i++) {
    		void* p = valueData[i];
    		[objcValue addObject:(NSBezierPath*)p];
    	}
    }
    [nSTextContainer setExclusionPaths:objcValue];
}

unsigned int C_NSTextContainer_LineBreakMode(void* ptr) {
    NSTextContainer* nSTextContainer = (NSTextContainer*)ptr;
    NSLineBreakMode result_ = [nSTextContainer lineBreakMode];
    return result_;
}

void C_NSTextContainer_SetLineBreakMode(void* ptr, unsigned int value) {
    NSTextContainer* nSTextContainer = (NSTextContainer*)ptr;
    [nSTextContainer setLineBreakMode:value];
}

bool C_NSTextContainer_WidthTracksTextView(void* ptr) {
    NSTextContainer* nSTextContainer = (NSTextContainer*)ptr;
    BOOL result_ = [nSTextContainer widthTracksTextView];
    return result_;
}

void C_NSTextContainer_SetWidthTracksTextView(void* ptr, bool value) {
    NSTextContainer* nSTextContainer = (NSTextContainer*)ptr;
    [nSTextContainer setWidthTracksTextView:value];
}

bool C_NSTextContainer_HeightTracksTextView(void* ptr) {
    NSTextContainer* nSTextContainer = (NSTextContainer*)ptr;
    BOOL result_ = [nSTextContainer heightTracksTextView];
    return result_;
}

void C_NSTextContainer_SetHeightTracksTextView(void* ptr, bool value) {
    NSTextContainer* nSTextContainer = (NSTextContainer*)ptr;
    [nSTextContainer setHeightTracksTextView:value];
}

unsigned int C_NSTextContainer_MaximumNumberOfLines(void* ptr) {
    NSTextContainer* nSTextContainer = (NSTextContainer*)ptr;
    NSUInteger result_ = [nSTextContainer maximumNumberOfLines];
    return result_;
}

void C_NSTextContainer_SetMaximumNumberOfLines(void* ptr, unsigned int value) {
    NSTextContainer* nSTextContainer = (NSTextContainer*)ptr;
    [nSTextContainer setMaximumNumberOfLines:value];
}

double C_NSTextContainer_LineFragmentPadding(void* ptr) {
    NSTextContainer* nSTextContainer = (NSTextContainer*)ptr;
    CGFloat result_ = [nSTextContainer lineFragmentPadding];
    return result_;
}

void C_NSTextContainer_SetLineFragmentPadding(void* ptr, double value) {
    NSTextContainer* nSTextContainer = (NSTextContainer*)ptr;
    [nSTextContainer setLineFragmentPadding:value];
}

bool C_NSTextContainer_IsSimpleRectangularTextContainer(void* ptr) {
    NSTextContainer* nSTextContainer = (NSTextContainer*)ptr;
    BOOL result_ = [nSTextContainer isSimpleRectangularTextContainer];
    return result_;
}
