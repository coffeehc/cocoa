#import <Foundation/Foundation.h>
#import "person_name_components.h"

void* C_PersonNameComponents_Alloc() {
    return [NSPersonNameComponents alloc];
}

void* C_NSPersonNameComponents_Init(void* ptr) {
    NSPersonNameComponents* nSPersonNameComponents = (NSPersonNameComponents*)ptr;
    NSPersonNameComponents* result_ = [nSPersonNameComponents init];
    return result_;
}

void* C_NSPersonNameComponents_NamePrefix(void* ptr) {
    NSPersonNameComponents* nSPersonNameComponents = (NSPersonNameComponents*)ptr;
    NSString* result_ = [nSPersonNameComponents namePrefix];
    return result_;
}

void C_NSPersonNameComponents_SetNamePrefix(void* ptr, void* value) {
    NSPersonNameComponents* nSPersonNameComponents = (NSPersonNameComponents*)ptr;
    [nSPersonNameComponents setNamePrefix:(NSString*)value];
}

void* C_NSPersonNameComponents_GivenName(void* ptr) {
    NSPersonNameComponents* nSPersonNameComponents = (NSPersonNameComponents*)ptr;
    NSString* result_ = [nSPersonNameComponents givenName];
    return result_;
}

void C_NSPersonNameComponents_SetGivenName(void* ptr, void* value) {
    NSPersonNameComponents* nSPersonNameComponents = (NSPersonNameComponents*)ptr;
    [nSPersonNameComponents setGivenName:(NSString*)value];
}

void* C_NSPersonNameComponents_MiddleName(void* ptr) {
    NSPersonNameComponents* nSPersonNameComponents = (NSPersonNameComponents*)ptr;
    NSString* result_ = [nSPersonNameComponents middleName];
    return result_;
}

void C_NSPersonNameComponents_SetMiddleName(void* ptr, void* value) {
    NSPersonNameComponents* nSPersonNameComponents = (NSPersonNameComponents*)ptr;
    [nSPersonNameComponents setMiddleName:(NSString*)value];
}

void* C_NSPersonNameComponents_FamilyName(void* ptr) {
    NSPersonNameComponents* nSPersonNameComponents = (NSPersonNameComponents*)ptr;
    NSString* result_ = [nSPersonNameComponents familyName];
    return result_;
}

void C_NSPersonNameComponents_SetFamilyName(void* ptr, void* value) {
    NSPersonNameComponents* nSPersonNameComponents = (NSPersonNameComponents*)ptr;
    [nSPersonNameComponents setFamilyName:(NSString*)value];
}

void* C_NSPersonNameComponents_NameSuffix(void* ptr) {
    NSPersonNameComponents* nSPersonNameComponents = (NSPersonNameComponents*)ptr;
    NSString* result_ = [nSPersonNameComponents nameSuffix];
    return result_;
}

void C_NSPersonNameComponents_SetNameSuffix(void* ptr, void* value) {
    NSPersonNameComponents* nSPersonNameComponents = (NSPersonNameComponents*)ptr;
    [nSPersonNameComponents setNameSuffix:(NSString*)value];
}

void* C_NSPersonNameComponents_Nickname(void* ptr) {
    NSPersonNameComponents* nSPersonNameComponents = (NSPersonNameComponents*)ptr;
    NSString* result_ = [nSPersonNameComponents nickname];
    return result_;
}

void C_NSPersonNameComponents_SetNickname(void* ptr, void* value) {
    NSPersonNameComponents* nSPersonNameComponents = (NSPersonNameComponents*)ptr;
    [nSPersonNameComponents setNickname:(NSString*)value];
}

void* C_NSPersonNameComponents_PhoneticRepresentation(void* ptr) {
    NSPersonNameComponents* nSPersonNameComponents = (NSPersonNameComponents*)ptr;
    NSPersonNameComponents* result_ = [nSPersonNameComponents phoneticRepresentation];
    return result_;
}

void C_NSPersonNameComponents_SetPhoneticRepresentation(void* ptr, void* value) {
    NSPersonNameComponents* nSPersonNameComponents = (NSPersonNameComponents*)ptr;
    [nSPersonNameComponents setPhoneticRepresentation:(NSPersonNameComponents*)value];
}
