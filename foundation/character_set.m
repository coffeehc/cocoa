#import <Foundation/Foundation.h>
#import "character_set.h"

void* C_CharacterSet_Alloc() {
    return [NSCharacterSet alloc];
}

void* C_NSCharacterSet_InitWithCoder(void* ptr, void* coder) {
    NSCharacterSet* nSCharacterSet = (NSCharacterSet*)ptr;
    NSCharacterSet* result = [nSCharacterSet initWithCoder:(NSCoder*)coder];
    return result;
}

void* C_NSCharacterSet_CharacterSetWithCharactersInString(void* aString) {
    NSCharacterSet* result = [NSCharacterSet characterSetWithCharactersInString:(NSString*)aString];
    return result;
}

void* C_NSCharacterSet_CharacterSetWithRange(NSRange aRange) {
    NSCharacterSet* result = [NSCharacterSet characterSetWithRange:aRange];
    return result;
}

void* C_NSCharacterSet_CharacterSetWithBitmapRepresentation(Array data) {
    NSCharacterSet* result = [NSCharacterSet characterSetWithBitmapRepresentation:[[NSData alloc] initWithBytes:(Byte *)data.data length:data.len]];
    return result;
}

void* C_NSCharacterSet_CharacterSetWithContentsOfFile(void* fName) {
    NSCharacterSet* result = [NSCharacterSet characterSetWithContentsOfFile:(NSString*)fName];
    return result;
}

bool C_NSCharacterSet_IsSupersetOfSet(void* ptr, void* theOtherSet) {
    NSCharacterSet* nSCharacterSet = (NSCharacterSet*)ptr;
    BOOL result = [nSCharacterSet isSupersetOfSet:(NSCharacterSet*)theOtherSet];
    return result;
}

void* C_NSCharacterSet_AlphanumericCharacterSet() {
    NSCharacterSet* result = [NSCharacterSet alphanumericCharacterSet];
    return result;
}

void* C_NSCharacterSet_CapitalizedLetterCharacterSet() {
    NSCharacterSet* result = [NSCharacterSet capitalizedLetterCharacterSet];
    return result;
}

void* C_NSCharacterSet_ControlCharacterSet() {
    NSCharacterSet* result = [NSCharacterSet controlCharacterSet];
    return result;
}

void* C_NSCharacterSet_DecimalDigitCharacterSet() {
    NSCharacterSet* result = [NSCharacterSet decimalDigitCharacterSet];
    return result;
}

void* C_NSCharacterSet_DecomposableCharacterSet() {
    NSCharacterSet* result = [NSCharacterSet decomposableCharacterSet];
    return result;
}

void* C_NSCharacterSet_IllegalCharacterSet() {
    NSCharacterSet* result = [NSCharacterSet illegalCharacterSet];
    return result;
}

void* C_NSCharacterSet_LetterCharacterSet() {
    NSCharacterSet* result = [NSCharacterSet letterCharacterSet];
    return result;
}

void* C_NSCharacterSet_LowercaseLetterCharacterSet() {
    NSCharacterSet* result = [NSCharacterSet lowercaseLetterCharacterSet];
    return result;
}

void* C_NSCharacterSet_NewlineCharacterSet() {
    NSCharacterSet* result = [NSCharacterSet newlineCharacterSet];
    return result;
}

void* C_NSCharacterSet_NonBaseCharacterSet() {
    NSCharacterSet* result = [NSCharacterSet nonBaseCharacterSet];
    return result;
}

void* C_NSCharacterSet_PunctuationCharacterSet() {
    NSCharacterSet* result = [NSCharacterSet punctuationCharacterSet];
    return result;
}

void* C_NSCharacterSet_SymbolCharacterSet() {
    NSCharacterSet* result = [NSCharacterSet symbolCharacterSet];
    return result;
}

void* C_NSCharacterSet_UppercaseLetterCharacterSet() {
    NSCharacterSet* result = [NSCharacterSet uppercaseLetterCharacterSet];
    return result;
}

void* C_NSCharacterSet_WhitespaceAndNewlineCharacterSet() {
    NSCharacterSet* result = [NSCharacterSet whitespaceAndNewlineCharacterSet];
    return result;
}

void* C_NSCharacterSet_WhitespaceCharacterSet() {
    NSCharacterSet* result = [NSCharacterSet whitespaceCharacterSet];
    return result;
}

void* C_NSCharacterSet_URLFragmentAllowedCharacterSet() {
    NSCharacterSet* result = [NSCharacterSet URLFragmentAllowedCharacterSet];
    return result;
}

void* C_NSCharacterSet_URLHostAllowedCharacterSet() {
    NSCharacterSet* result = [NSCharacterSet URLHostAllowedCharacterSet];
    return result;
}

void* C_NSCharacterSet_URLPasswordAllowedCharacterSet() {
    NSCharacterSet* result = [NSCharacterSet URLPasswordAllowedCharacterSet];
    return result;
}

void* C_NSCharacterSet_URLPathAllowedCharacterSet() {
    NSCharacterSet* result = [NSCharacterSet URLPathAllowedCharacterSet];
    return result;
}

void* C_NSCharacterSet_URLQueryAllowedCharacterSet() {
    NSCharacterSet* result = [NSCharacterSet URLQueryAllowedCharacterSet];
    return result;
}

void* C_NSCharacterSet_URLUserAllowedCharacterSet() {
    NSCharacterSet* result = [NSCharacterSet URLUserAllowedCharacterSet];
    return result;
}

Array C_NSCharacterSet_BitmapRepresentation(void* ptr) {
    NSCharacterSet* nSCharacterSet = (NSCharacterSet*)ptr;
    NSData* result = [nSCharacterSet bitmapRepresentation];
    Array resultarray;
    resultarray.data = [result bytes];
    resultarray.len = result.length;
    return resultarray;
}

void* C_NSCharacterSet_InvertedSet(void* ptr) {
    NSCharacterSet* nSCharacterSet = (NSCharacterSet*)ptr;
    NSCharacterSet* result = [nSCharacterSet invertedSet];
    return result;
}
