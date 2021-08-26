#import <Appkit/Appkit.h>
#import "paragraph_style.h"

void* C_ParagraphStyle_Alloc() {
    return [NSParagraphStyle alloc];
}

void* C_NSParagraphStyle_Init(void* ptr) {
    NSParagraphStyle* nSParagraphStyle = (NSParagraphStyle*)ptr;
    NSParagraphStyle* result_ = [nSParagraphStyle init];
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
