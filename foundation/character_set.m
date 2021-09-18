#import "character_set.h"
#import <Foundation/NSDictionary.h>
#import <Foundation/NSArray.h>
#import <Foundation/NSCharacterSet.h>
#import <Foundation/NSURL.h>

void* C_CharacterSet_Alloc() {
    return [NSCharacterSet alloc];
}

void* C_NSCharacterSet_InitWithCoder(void* ptr, void* coder) {
    NSCharacterSet* nSCharacterSet = (NSCharacterSet*)ptr;
    NSCharacterSet* result_ = [nSCharacterSet initWithCoder:(NSCoder*)coder];
    return result_;
}

void* C_NSCharacterSet_AllocCharacterSet() {
    NSCharacterSet* result_ = [NSCharacterSet alloc];
    return result_;
}

void* C_NSCharacterSet_Autorelease(void* ptr) {
    NSCharacterSet* nSCharacterSet = (NSCharacterSet*)ptr;
    NSCharacterSet* result_ = [nSCharacterSet autorelease];
    return result_;
}

void* C_NSCharacterSet_Retain(void* ptr) {
    NSCharacterSet* nSCharacterSet = (NSCharacterSet*)ptr;
    NSCharacterSet* result_ = [nSCharacterSet retain];
    return result_;
}

void* C_NSCharacterSet_CharacterSetWithCharactersInString(void* aString) {
    NSCharacterSet* result_ = [NSCharacterSet characterSetWithCharactersInString:(NSString*)aString];
    return result_;
}

void* C_NSCharacterSet_CharacterSetWithRange(NSRange aRange) {
    NSCharacterSet* result_ = [NSCharacterSet characterSetWithRange:aRange];
    return result_;
}

void* C_NSCharacterSet_CharacterSetWithBitmapRepresentation(void* data) {
    NSCharacterSet* result_ = [NSCharacterSet characterSetWithBitmapRepresentation:(NSData*)data];
    return result_;
}

void* C_NSCharacterSet_CharacterSetWithContentsOfFile(void* fName) {
    NSCharacterSet* result_ = [NSCharacterSet characterSetWithContentsOfFile:(NSString*)fName];
    return result_;
}

bool C_NSCharacterSet_IsSupersetOfSet(void* ptr, void* theOtherSet) {
    NSCharacterSet* nSCharacterSet = (NSCharacterSet*)ptr;
    BOOL result_ = [nSCharacterSet isSupersetOfSet:(NSCharacterSet*)theOtherSet];
    return result_;
}

void* C_NSCharacterSet_AlphanumericCharacterSet() {
    NSCharacterSet* result_ = [NSCharacterSet alphanumericCharacterSet];
    return result_;
}

void* C_NSCharacterSet_CapitalizedLetterCharacterSet() {
    NSCharacterSet* result_ = [NSCharacterSet capitalizedLetterCharacterSet];
    return result_;
}

void* C_NSCharacterSet_ControlCharacterSet() {
    NSCharacterSet* result_ = [NSCharacterSet controlCharacterSet];
    return result_;
}

void* C_NSCharacterSet_DecimalDigitCharacterSet() {
    NSCharacterSet* result_ = [NSCharacterSet decimalDigitCharacterSet];
    return result_;
}

void* C_NSCharacterSet_DecomposableCharacterSet() {
    NSCharacterSet* result_ = [NSCharacterSet decomposableCharacterSet];
    return result_;
}

void* C_NSCharacterSet_IllegalCharacterSet() {
    NSCharacterSet* result_ = [NSCharacterSet illegalCharacterSet];
    return result_;
}

void* C_NSCharacterSet_LetterCharacterSet() {
    NSCharacterSet* result_ = [NSCharacterSet letterCharacterSet];
    return result_;
}

void* C_NSCharacterSet_LowercaseLetterCharacterSet() {
    NSCharacterSet* result_ = [NSCharacterSet lowercaseLetterCharacterSet];
    return result_;
}

void* C_NSCharacterSet_NewlineCharacterSet() {
    NSCharacterSet* result_ = [NSCharacterSet newlineCharacterSet];
    return result_;
}

void* C_NSCharacterSet_NonBaseCharacterSet() {
    NSCharacterSet* result_ = [NSCharacterSet nonBaseCharacterSet];
    return result_;
}

void* C_NSCharacterSet_PunctuationCharacterSet() {
    NSCharacterSet* result_ = [NSCharacterSet punctuationCharacterSet];
    return result_;
}

void* C_NSCharacterSet_SymbolCharacterSet() {
    NSCharacterSet* result_ = [NSCharacterSet symbolCharacterSet];
    return result_;
}

void* C_NSCharacterSet_UppercaseLetterCharacterSet() {
    NSCharacterSet* result_ = [NSCharacterSet uppercaseLetterCharacterSet];
    return result_;
}

void* C_NSCharacterSet_WhitespaceAndNewlineCharacterSet() {
    NSCharacterSet* result_ = [NSCharacterSet whitespaceAndNewlineCharacterSet];
    return result_;
}

void* C_NSCharacterSet_WhitespaceCharacterSet() {
    NSCharacterSet* result_ = [NSCharacterSet whitespaceCharacterSet];
    return result_;
}

void* C_NSCharacterSet_URLFragmentAllowedCharacterSet() {
    NSCharacterSet* result_ = [NSCharacterSet URLFragmentAllowedCharacterSet];
    return result_;
}

void* C_NSCharacterSet_URLHostAllowedCharacterSet() {
    NSCharacterSet* result_ = [NSCharacterSet URLHostAllowedCharacterSet];
    return result_;
}

void* C_NSCharacterSet_URLPasswordAllowedCharacterSet() {
    NSCharacterSet* result_ = [NSCharacterSet URLPasswordAllowedCharacterSet];
    return result_;
}

void* C_NSCharacterSet_URLPathAllowedCharacterSet() {
    NSCharacterSet* result_ = [NSCharacterSet URLPathAllowedCharacterSet];
    return result_;
}

void* C_NSCharacterSet_URLQueryAllowedCharacterSet() {
    NSCharacterSet* result_ = [NSCharacterSet URLQueryAllowedCharacterSet];
    return result_;
}

void* C_NSCharacterSet_URLUserAllowedCharacterSet() {
    NSCharacterSet* result_ = [NSCharacterSet URLUserAllowedCharacterSet];
    return result_;
}

void* C_NSCharacterSet_BitmapRepresentation(void* ptr) {
    NSCharacterSet* nSCharacterSet = (NSCharacterSet*)ptr;
    NSData* result_ = [nSCharacterSet bitmapRepresentation];
    return result_;
}

void* C_NSCharacterSet_InvertedSet(void* ptr) {
    NSCharacterSet* nSCharacterSet = (NSCharacterSet*)ptr;
    NSCharacterSet* result_ = [nSCharacterSet invertedSet];
    return result_;
}
