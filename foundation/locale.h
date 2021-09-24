#import <Foundation/NSGeometry.h>
#import <Foundation/NSRange.h>
#import <stdbool.h>
#import <stdint.h>
#import <stdlib.h>
#import <utils.h>

void* C_Locale_Alloc();

void* C_NSLocale_LocaleWithLocaleIdentifier(void* ident);
void* C_NSLocale_InitWithLocaleIdentifier(void* ptr, void* _string);
void* C_NSLocale_InitWithCoder(void* ptr, void* coder);
void* C_NSLocale_AllocLocale();
void* C_NSLocale_Autorelease(void* ptr);
void* C_NSLocale_Retain(void* ptr);
void* C_NSLocale_CanonicalLocaleIdentifierFromString(void* _string);
Dictionary C_NSLocale_ComponentsFromLocaleIdentifier(void* _string);
void* C_NSLocale_LocaleIdentifierFromComponents(Dictionary dict);
void* C_NSLocale_Locale_CanonicalLanguageIdentifierFromString(void* _string);
void* C_NSLocale_LocaleIdentifierFromWindowsLocaleCode(uint32_t lcid);
uint32_t C_NSLocale_WindowsLocaleCodeFromLocaleIdentifier(void* localeIdentifier);
void* C_NSLocale_LocalizedStringForLocaleIdentifier(void* ptr, void* localeIdentifier);
void* C_NSLocale_LocalizedStringForCountryCode(void* ptr, void* countryCode);
void* C_NSLocale_LocalizedStringForLanguageCode(void* ptr, void* languageCode);
void* C_NSLocale_LocalizedStringForScriptCode(void* ptr, void* scriptCode);
void* C_NSLocale_LocalizedStringForVariantCode(void* ptr, void* variantCode);
void* C_NSLocale_LocalizedStringForCollationIdentifier(void* ptr, void* collationIdentifier);
void* C_NSLocale_LocalizedStringForCollatorIdentifier(void* ptr, void* collatorIdentifier);
void* C_NSLocale_LocalizedStringForCurrencyCode(void* ptr, void* currencyCode);
void* C_NSLocale_LocalizedStringForCalendarIdentifier(void* ptr, void* calendarIdentifier);
void* C_NSLocale_ObjectForKey(void* ptr, void* key);
void* C_NSLocale_DisplayNameForKey_Value(void* ptr, void* key, void* value);
unsigned int C_NSLocale_Locale_CharacterDirectionForLanguage(void* isoLangCode);
unsigned int C_NSLocale_Locale_LineDirectionForLanguage(void* isoLangCode);
void* C_NSLocale_AutoupdatingCurrentLocale();
void* C_NSLocale_CurrentLocale();
void* C_NSLocale_SystemLocale();
Array C_NSLocale_AvailableLocaleIdentifiers();
Array C_NSLocale_Locale_ISOCountryCodes();
Array C_NSLocale_Locale_ISOLanguageCodes();
Array C_NSLocale_Locale_ISOCurrencyCodes();
Array C_NSLocale_Locale_CommonISOCurrencyCodes();
void* C_NSLocale_LocaleIdentifier(void* ptr);
void* C_NSLocale_CountryCode(void* ptr);
void* C_NSLocale_LanguageCode(void* ptr);
void* C_NSLocale_ScriptCode(void* ptr);
void* C_NSLocale_VariantCode(void* ptr);
void* C_NSLocale_ExemplarCharacterSet(void* ptr);
void* C_NSLocale_CollationIdentifier(void* ptr);
void* C_NSLocale_CollatorIdentifier(void* ptr);
bool C_NSLocale_UsesMetricSystem(void* ptr);
void* C_NSLocale_DecimalSeparator(void* ptr);
void* C_NSLocale_GroupingSeparator(void* ptr);
void* C_NSLocale_CurrencyCode(void* ptr);
void* C_NSLocale_CurrencySymbol(void* ptr);
void* C_NSLocale_CalendarIdentifier(void* ptr);
void* C_NSLocale_QuotationBeginDelimiter(void* ptr);
void* C_NSLocale_QuotationEndDelimiter(void* ptr);
void* C_NSLocale_AlternateQuotationBeginDelimiter(void* ptr);
void* C_NSLocale_AlternateQuotationEndDelimiter(void* ptr);
Array C_NSLocale_Locale_PreferredLanguages();
