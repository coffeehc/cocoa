#import "person_name_components_formatter.h"
#import <Foundation/NSDictionary.h>
#import <Foundation/NSArray.h>
#import <Foundation/NSPersonNameComponentsFormatter.h>

void* C_PersonNameComponentsFormatter_Alloc() {
    return [NSPersonNameComponentsFormatter alloc];
}

void* C_NSPersonNameComponentsFormatter_AllocPersonNameComponentsFormatter() {
    NSPersonNameComponentsFormatter* result_ = [NSPersonNameComponentsFormatter alloc];
    return result_;
}

void* C_NSPersonNameComponentsFormatter_Init(void* ptr) {
    NSPersonNameComponentsFormatter* nSPersonNameComponentsFormatter = (NSPersonNameComponentsFormatter*)ptr;
    NSPersonNameComponentsFormatter* result_ = [nSPersonNameComponentsFormatter init];
    return result_;
}

void* C_NSPersonNameComponentsFormatter_NewPersonNameComponentsFormatter() {
    NSPersonNameComponentsFormatter* result_ = [NSPersonNameComponentsFormatter new];
    return result_;
}

void* C_NSPersonNameComponentsFormatter_Autorelease(void* ptr) {
    NSPersonNameComponentsFormatter* nSPersonNameComponentsFormatter = (NSPersonNameComponentsFormatter*)ptr;
    NSPersonNameComponentsFormatter* result_ = [nSPersonNameComponentsFormatter autorelease];
    return result_;
}

void* C_NSPersonNameComponentsFormatter_Retain(void* ptr) {
    NSPersonNameComponentsFormatter* nSPersonNameComponentsFormatter = (NSPersonNameComponentsFormatter*)ptr;
    NSPersonNameComponentsFormatter* result_ = [nSPersonNameComponentsFormatter retain];
    return result_;
}

void* C_NSPersonNameComponentsFormatter_PersonNameComponentsFormatter_LocalizedStringFromPersonNameComponents_Style_Options(void* components, int nameFormatStyle, unsigned int nameOptions) {
    NSString* result_ = [NSPersonNameComponentsFormatter localizedStringFromPersonNameComponents:(NSPersonNameComponents*)components style:nameFormatStyle options:nameOptions];
    return result_;
}

void* C_NSPersonNameComponentsFormatter_StringFromPersonNameComponents(void* ptr, void* components) {
    NSPersonNameComponentsFormatter* nSPersonNameComponentsFormatter = (NSPersonNameComponentsFormatter*)ptr;
    NSString* result_ = [nSPersonNameComponentsFormatter stringFromPersonNameComponents:(NSPersonNameComponents*)components];
    return result_;
}

void* C_NSPersonNameComponentsFormatter_AnnotatedStringFromPersonNameComponents(void* ptr, void* components) {
    NSPersonNameComponentsFormatter* nSPersonNameComponentsFormatter = (NSPersonNameComponentsFormatter*)ptr;
    NSAttributedString* result_ = [nSPersonNameComponentsFormatter annotatedStringFromPersonNameComponents:(NSPersonNameComponents*)components];
    return result_;
}

void* C_NSPersonNameComponentsFormatter_PersonNameComponentsFromString(void* ptr, void* _string) {
    NSPersonNameComponentsFormatter* nSPersonNameComponentsFormatter = (NSPersonNameComponentsFormatter*)ptr;
    NSPersonNameComponents* result_ = [nSPersonNameComponentsFormatter personNameComponentsFromString:(NSString*)_string];
    return result_;
}

int C_NSPersonNameComponentsFormatter_Style(void* ptr) {
    NSPersonNameComponentsFormatter* nSPersonNameComponentsFormatter = (NSPersonNameComponentsFormatter*)ptr;
    NSPersonNameComponentsFormatterStyle result_ = [nSPersonNameComponentsFormatter style];
    return result_;
}

void C_NSPersonNameComponentsFormatter_SetStyle(void* ptr, int value) {
    NSPersonNameComponentsFormatter* nSPersonNameComponentsFormatter = (NSPersonNameComponentsFormatter*)ptr;
    [nSPersonNameComponentsFormatter setStyle:value];
}

bool C_NSPersonNameComponentsFormatter_IsPhonetic(void* ptr) {
    NSPersonNameComponentsFormatter* nSPersonNameComponentsFormatter = (NSPersonNameComponentsFormatter*)ptr;
    BOOL result_ = [nSPersonNameComponentsFormatter isPhonetic];
    return result_;
}

void C_NSPersonNameComponentsFormatter_SetPhonetic(void* ptr, bool value) {
    NSPersonNameComponentsFormatter* nSPersonNameComponentsFormatter = (NSPersonNameComponentsFormatter*)ptr;
    [nSPersonNameComponentsFormatter setPhonetic:value];
}
