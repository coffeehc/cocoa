#import <Appkit/Appkit.h>
#import "paragraph_style.h"

void* C_ParagraphStyle_Alloc() {
    return [NSParagraphStyle alloc];
}

void* C_NSParagraphStyle_Init(void* ptr) {
    NSParagraphStyle* nSParagraphStyle = (NSParagraphStyle*)ptr;
    NSParagraphStyle* result = [nSParagraphStyle init];
    return result;
}

int C_NSParagraphStyle_ParagraphStyle_DefaultWritingDirectionForLanguage(void* languageName) {
    NSWritingDirection result = [NSParagraphStyle defaultWritingDirectionForLanguage:(NSString*)languageName];
    return result;
}

void* C_NSParagraphStyle_DefaultParagraphStyle() {
    NSParagraphStyle* result = [NSParagraphStyle defaultParagraphStyle];
    return result;
}

int C_NSParagraphStyle_Alignment(void* ptr) {
    NSParagraphStyle* nSParagraphStyle = (NSParagraphStyle*)ptr;
    NSTextAlignment result = [nSParagraphStyle alignment];
    return result;
}

double C_NSParagraphStyle_FirstLineHeadIndent(void* ptr) {
    NSParagraphStyle* nSParagraphStyle = (NSParagraphStyle*)ptr;
    CGFloat result = [nSParagraphStyle firstLineHeadIndent];
    return result;
}

double C_NSParagraphStyle_HeadIndent(void* ptr) {
    NSParagraphStyle* nSParagraphStyle = (NSParagraphStyle*)ptr;
    CGFloat result = [nSParagraphStyle headIndent];
    return result;
}

double C_NSParagraphStyle_TailIndent(void* ptr) {
    NSParagraphStyle* nSParagraphStyle = (NSParagraphStyle*)ptr;
    CGFloat result = [nSParagraphStyle tailIndent];
    return result;
}

double C_NSParagraphStyle_LineHeightMultiple(void* ptr) {
    NSParagraphStyle* nSParagraphStyle = (NSParagraphStyle*)ptr;
    CGFloat result = [nSParagraphStyle lineHeightMultiple];
    return result;
}

double C_NSParagraphStyle_MaximumLineHeight(void* ptr) {
    NSParagraphStyle* nSParagraphStyle = (NSParagraphStyle*)ptr;
    CGFloat result = [nSParagraphStyle maximumLineHeight];
    return result;
}

double C_NSParagraphStyle_MinimumLineHeight(void* ptr) {
    NSParagraphStyle* nSParagraphStyle = (NSParagraphStyle*)ptr;
    CGFloat result = [nSParagraphStyle minimumLineHeight];
    return result;
}

double C_NSParagraphStyle_LineSpacing(void* ptr) {
    NSParagraphStyle* nSParagraphStyle = (NSParagraphStyle*)ptr;
    CGFloat result = [nSParagraphStyle lineSpacing];
    return result;
}

double C_NSParagraphStyle_ParagraphSpacing(void* ptr) {
    NSParagraphStyle* nSParagraphStyle = (NSParagraphStyle*)ptr;
    CGFloat result = [nSParagraphStyle paragraphSpacing];
    return result;
}

double C_NSParagraphStyle_ParagraphSpacingBefore(void* ptr) {
    NSParagraphStyle* nSParagraphStyle = (NSParagraphStyle*)ptr;
    CGFloat result = [nSParagraphStyle paragraphSpacingBefore];
    return result;
}

Array C_NSParagraphStyle_TabStops(void* ptr) {
    NSParagraphStyle* nSParagraphStyle = (NSParagraphStyle*)ptr;
    NSArray* result = [nSParagraphStyle tabStops];
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

double C_NSParagraphStyle_DefaultTabInterval(void* ptr) {
    NSParagraphStyle* nSParagraphStyle = (NSParagraphStyle*)ptr;
    CGFloat result = [nSParagraphStyle defaultTabInterval];
    return result;
}

Array C_NSParagraphStyle_TextBlocks(void* ptr) {
    NSParagraphStyle* nSParagraphStyle = (NSParagraphStyle*)ptr;
    NSArray* result = [nSParagraphStyle textBlocks];
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

Array C_NSParagraphStyle_TextLists(void* ptr) {
    NSParagraphStyle* nSParagraphStyle = (NSParagraphStyle*)ptr;
    NSArray* result = [nSParagraphStyle textLists];
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

int C_NSParagraphStyle_LineBreakMode(void* ptr) {
    NSParagraphStyle* nSParagraphStyle = (NSParagraphStyle*)ptr;
    NSLineBreakMode result = [nSParagraphStyle lineBreakMode];
    return result;
}

unsigned int C_NSParagraphStyle_LineBreakStrategy(void* ptr) {
    NSParagraphStyle* nSParagraphStyle = (NSParagraphStyle*)ptr;
    NSLineBreakStrategy result = [nSParagraphStyle lineBreakStrategy];
    return result;
}

float C_NSParagraphStyle_HyphenationFactor(void* ptr) {
    NSParagraphStyle* nSParagraphStyle = (NSParagraphStyle*)ptr;
    float result = [nSParagraphStyle hyphenationFactor];
    return result;
}

float C_NSParagraphStyle_TighteningFactorForTruncation(void* ptr) {
    NSParagraphStyle* nSParagraphStyle = (NSParagraphStyle*)ptr;
    float result = [nSParagraphStyle tighteningFactorForTruncation];
    return result;
}

bool C_NSParagraphStyle_AllowsDefaultTighteningForTruncation(void* ptr) {
    NSParagraphStyle* nSParagraphStyle = (NSParagraphStyle*)ptr;
    BOOL result = [nSParagraphStyle allowsDefaultTighteningForTruncation];
    return result;
}

int C_NSParagraphStyle_HeaderLevel(void* ptr) {
    NSParagraphStyle* nSParagraphStyle = (NSParagraphStyle*)ptr;
    NSInteger result = [nSParagraphStyle headerLevel];
    return result;
}

int C_NSParagraphStyle_BaseWritingDirection(void* ptr) {
    NSParagraphStyle* nSParagraphStyle = (NSParagraphStyle*)ptr;
    NSWritingDirection result = [nSParagraphStyle baseWritingDirection];
    return result;
}
