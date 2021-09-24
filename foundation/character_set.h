#import <Foundation/NSGeometry.h>
#import <Foundation/NSRange.h>
#import <stdbool.h>
#import <stdint.h>
#import <stdlib.h>
#import <utils.h>

void* C_CharacterSet_Alloc();

void* C_NSCharacterSet_InitWithCoder(void* ptr, void* coder);
void* C_NSCharacterSet_AllocCharacterSet();
void* C_NSCharacterSet_Autorelease(void* ptr);
void* C_NSCharacterSet_Retain(void* ptr);
void* C_NSCharacterSet_CharacterSetWithCharactersInString(void* aString);
void* C_NSCharacterSet_CharacterSetWithRange(NSRange aRange);
void* C_NSCharacterSet_CharacterSetWithBitmapRepresentation(void* data);
void* C_NSCharacterSet_CharacterSetWithContentsOfFile(void* fName);
bool C_NSCharacterSet_IsSupersetOfSet(void* ptr, void* theOtherSet);
void* C_NSCharacterSet_AlphanumericCharacterSet();
void* C_NSCharacterSet_CapitalizedLetterCharacterSet();
void* C_NSCharacterSet_ControlCharacterSet();
void* C_NSCharacterSet_DecimalDigitCharacterSet();
void* C_NSCharacterSet_DecomposableCharacterSet();
void* C_NSCharacterSet_IllegalCharacterSet();
void* C_NSCharacterSet_LetterCharacterSet();
void* C_NSCharacterSet_LowercaseLetterCharacterSet();
void* C_NSCharacterSet_NewlineCharacterSet();
void* C_NSCharacterSet_NonBaseCharacterSet();
void* C_NSCharacterSet_PunctuationCharacterSet();
void* C_NSCharacterSet_SymbolCharacterSet();
void* C_NSCharacterSet_UppercaseLetterCharacterSet();
void* C_NSCharacterSet_WhitespaceAndNewlineCharacterSet();
void* C_NSCharacterSet_WhitespaceCharacterSet();
void* C_NSCharacterSet_URLFragmentAllowedCharacterSet();
void* C_NSCharacterSet_URLHostAllowedCharacterSet();
void* C_NSCharacterSet_URLPasswordAllowedCharacterSet();
void* C_NSCharacterSet_URLPathAllowedCharacterSet();
void* C_NSCharacterSet_URLQueryAllowedCharacterSet();
void* C_NSCharacterSet_URLUserAllowedCharacterSet();
void* C_NSCharacterSet_BitmapRepresentation(void* ptr);
void* C_NSCharacterSet_InvertedSet(void* ptr);
