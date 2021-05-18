#import <Appkit/Appkit.h>
#import "text_tab.h"

void* C_TextTab_Alloc() {
    return [NSTextTab alloc];
}

void* C_NSTextTab_Init(void* ptr) {
    NSTextTab* nSTextTab = (NSTextTab*)ptr;
    NSTextTab* result_ = [nSTextTab init];
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
