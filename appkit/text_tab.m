#import <Appkit/Appkit.h>
#import "text_tab.h"

void* C_TextTab_Alloc() {
    return [NSTextTab alloc];
}

void* C_NSTextTab_Init(void* ptr) {
    NSTextTab* nSTextTab = (NSTextTab*)ptr;
    NSTextTab* result = [nSTextTab init];
    return result;
}

void* C_NSTextTab_TextTab_ColumnTerminatorsForLocale(void* aLocale) {
    NSCharacterSet* result = [NSTextTab columnTerminatorsForLocale:(NSLocale*)aLocale];
    return result;
}

double C_NSTextTab_Location(void* ptr) {
    NSTextTab* nSTextTab = (NSTextTab*)ptr;
    CGFloat result = [nSTextTab location];
    return result;
}

int C_NSTextTab_Alignment(void* ptr) {
    NSTextTab* nSTextTab = (NSTextTab*)ptr;
    NSTextAlignment result = [nSTextTab alignment];
    return result;
}
