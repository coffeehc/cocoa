#import <stdbool.h>
#import <stdlib.h>
#import <utils.h>
#import <Foundation/Foundation.h>

void* C_Orthography_Alloc();

void* C_NSOrthography_DefaultOrthographyForLanguage(void* language);
void* C_NSOrthography_InitWithCoder(void* ptr, void* coder);
void* C_NSOrthography_AllocOrthography();
void* C_NSOrthography_Init(void* ptr);
void* C_NSOrthography_NewOrthography();
void* C_NSOrthography_Autorelease(void* ptr);
void* C_NSOrthography_Retain(void* ptr);
void* C_NSOrthography_DominantLanguageForScript(void* ptr, void* script);
Array C_NSOrthography_LanguagesForScript(void* ptr, void* script);
void* C_NSOrthography_DominantLanguage(void* ptr);
void* C_NSOrthography_DominantScript(void* ptr);
Array C_NSOrthography_AllScripts(void* ptr);
Array C_NSOrthography_AllLanguages(void* ptr);
