#import <stdbool.h>
#import <stdlib.h>
#import <utils.h>
#import <Foundation/Foundation.h>

void* C_PersonNameComponents_Alloc();

void* C_NSPersonNameComponents_Init(void* ptr);
void* C_NSPersonNameComponents_NamePrefix(void* ptr);
void C_NSPersonNameComponents_SetNamePrefix(void* ptr, void* value);
void* C_NSPersonNameComponents_GivenName(void* ptr);
void C_NSPersonNameComponents_SetGivenName(void* ptr, void* value);
void* C_NSPersonNameComponents_MiddleName(void* ptr);
void C_NSPersonNameComponents_SetMiddleName(void* ptr, void* value);
void* C_NSPersonNameComponents_FamilyName(void* ptr);
void C_NSPersonNameComponents_SetFamilyName(void* ptr, void* value);
void* C_NSPersonNameComponents_NameSuffix(void* ptr);
void C_NSPersonNameComponents_SetNameSuffix(void* ptr, void* value);
void* C_NSPersonNameComponents_Nickname(void* ptr);
void C_NSPersonNameComponents_SetNickname(void* ptr, void* value);
void* C_NSPersonNameComponents_PhoneticRepresentation(void* ptr);
void C_NSPersonNameComponents_SetPhoneticRepresentation(void* ptr, void* value);
