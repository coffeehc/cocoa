#import <stdbool.h>
#import <stdlib.h>
#import <utils.h>
#import <Foundation/Foundation.h>

void* C_PersonNameComponentsFormatter_Alloc();

void* C_NSPersonNameComponentsFormatter_AllocPersonNameComponentsFormatter();
void* C_NSPersonNameComponentsFormatter_Init(void* ptr);
void* C_NSPersonNameComponentsFormatter_NewPersonNameComponentsFormatter();
void* C_NSPersonNameComponentsFormatter_Autorelease(void* ptr);
void* C_NSPersonNameComponentsFormatter_Retain(void* ptr);
void* C_NSPersonNameComponentsFormatter_PersonNameComponentsFormatter_LocalizedStringFromPersonNameComponents_Style_Options(void* components, int nameFormatStyle, unsigned int nameOptions);
void* C_NSPersonNameComponentsFormatter_StringFromPersonNameComponents(void* ptr, void* components);
void* C_NSPersonNameComponentsFormatter_AnnotatedStringFromPersonNameComponents(void* ptr, void* components);
void* C_NSPersonNameComponentsFormatter_PersonNameComponentsFromString(void* ptr, void* _string);
int C_NSPersonNameComponentsFormatter_Style(void* ptr);
void C_NSPersonNameComponentsFormatter_SetStyle(void* ptr, int value);
bool C_NSPersonNameComponentsFormatter_IsPhonetic(void* ptr);
void C_NSPersonNameComponentsFormatter_SetPhonetic(void* ptr, bool value);
