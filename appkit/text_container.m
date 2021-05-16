#import <Appkit/Appkit.h>
#import "text_container.h"

void* C_TextContainer_Alloc() {
    return [NSTextContainer alloc];
}

void* C_NSTextContainer_InitWithSize(void* ptr, CGSize size) {
    NSTextContainer* nSTextContainer = (NSTextContainer*)ptr;
    NSTextContainer* result = [nSTextContainer initWithSize:size];
    return result;
}

void* C_NSTextContainer_InitWithCoder(void* ptr, void* coder) {
    NSTextContainer* nSTextContainer = (NSTextContainer*)ptr;
    NSTextContainer* result = [nSTextContainer initWithCoder:(NSCoder*)coder];
    return result;
}

void* C_NSTextContainer_Init(void* ptr) {
    NSTextContainer* nSTextContainer = (NSTextContainer*)ptr;
    NSTextContainer* result = [nSTextContainer init];
    return result;
}

void C_NSTextContainer_ReplaceLayoutManager(void* ptr, void* newLayoutManager) {
    NSTextContainer* nSTextContainer = (NSTextContainer*)ptr;
    [nSTextContainer replaceLayoutManager:(NSLayoutManager*)newLayoutManager];
}

void* C_NSTextContainer_LayoutManager(void* ptr) {
    NSTextContainer* nSTextContainer = (NSTextContainer*)ptr;
    NSLayoutManager* result = [nSTextContainer layoutManager];
    return result;
}

void C_NSTextContainer_SetLayoutManager(void* ptr, void* value) {
    NSTextContainer* nSTextContainer = (NSTextContainer*)ptr;
    [nSTextContainer setLayoutManager:(NSLayoutManager*)value];
}

void* C_NSTextContainer_TextView(void* ptr) {
    NSTextContainer* nSTextContainer = (NSTextContainer*)ptr;
    NSTextView* result = [nSTextContainer textView];
    return result;
}

void C_NSTextContainer_SetTextView(void* ptr, void* value) {
    NSTextContainer* nSTextContainer = (NSTextContainer*)ptr;
    [nSTextContainer setTextView:(NSTextView*)value];
}

CGSize C_NSTextContainer_Size(void* ptr) {
    NSTextContainer* nSTextContainer = (NSTextContainer*)ptr;
    NSSize result = [nSTextContainer size];
    return result;
}

void C_NSTextContainer_SetSize(void* ptr, CGSize value) {
    NSTextContainer* nSTextContainer = (NSTextContainer*)ptr;
    [nSTextContainer setSize:value];
}

Array C_NSTextContainer_ExclusionPaths(void* ptr) {
    NSTextContainer* nSTextContainer = (NSTextContainer*)ptr;
    NSArray* result = [nSTextContainer exclusionPaths];
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

void C_NSTextContainer_SetExclusionPaths(void* ptr, Array value) {
    NSTextContainer* nSTextContainer = (NSTextContainer*)ptr;
    NSMutableArray* objcValue = [[NSMutableArray alloc] init];
    void** valueData = (void**)value.data;
    for (int i = 0; i < value.len; i++) {
    	void* p = valueData[i];
    	[objcValue addObject:(NSBezierPath*)(NSBezierPath*)p];
    }
    [nSTextContainer setExclusionPaths:objcValue];
}

unsigned int C_NSTextContainer_LineBreakMode(void* ptr) {
    NSTextContainer* nSTextContainer = (NSTextContainer*)ptr;
    NSLineBreakMode result = [nSTextContainer lineBreakMode];
    return result;
}

void C_NSTextContainer_SetLineBreakMode(void* ptr, unsigned int value) {
    NSTextContainer* nSTextContainer = (NSTextContainer*)ptr;
    [nSTextContainer setLineBreakMode:value];
}

bool C_NSTextContainer_WidthTracksTextView(void* ptr) {
    NSTextContainer* nSTextContainer = (NSTextContainer*)ptr;
    BOOL result = [nSTextContainer widthTracksTextView];
    return result;
}

void C_NSTextContainer_SetWidthTracksTextView(void* ptr, bool value) {
    NSTextContainer* nSTextContainer = (NSTextContainer*)ptr;
    [nSTextContainer setWidthTracksTextView:value];
}

bool C_NSTextContainer_HeightTracksTextView(void* ptr) {
    NSTextContainer* nSTextContainer = (NSTextContainer*)ptr;
    BOOL result = [nSTextContainer heightTracksTextView];
    return result;
}

void C_NSTextContainer_SetHeightTracksTextView(void* ptr, bool value) {
    NSTextContainer* nSTextContainer = (NSTextContainer*)ptr;
    [nSTextContainer setHeightTracksTextView:value];
}

unsigned int C_NSTextContainer_MaximumNumberOfLines(void* ptr) {
    NSTextContainer* nSTextContainer = (NSTextContainer*)ptr;
    NSUInteger result = [nSTextContainer maximumNumberOfLines];
    return result;
}

void C_NSTextContainer_SetMaximumNumberOfLines(void* ptr, unsigned int value) {
    NSTextContainer* nSTextContainer = (NSTextContainer*)ptr;
    [nSTextContainer setMaximumNumberOfLines:value];
}

double C_NSTextContainer_LineFragmentPadding(void* ptr) {
    NSTextContainer* nSTextContainer = (NSTextContainer*)ptr;
    CGFloat result = [nSTextContainer lineFragmentPadding];
    return result;
}

void C_NSTextContainer_SetLineFragmentPadding(void* ptr, double value) {
    NSTextContainer* nSTextContainer = (NSTextContainer*)ptr;
    [nSTextContainer setLineFragmentPadding:value];
}

bool C_NSTextContainer_IsSimpleRectangularTextContainer(void* ptr) {
    NSTextContainer* nSTextContainer = (NSTextContainer*)ptr;
    BOOL result = [nSTextContainer isSimpleRectangularTextContainer];
    return result;
}
