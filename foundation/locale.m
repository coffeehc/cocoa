#import <Foundation/Foundation.h>
#import "locale.h"

void* C_Locale_Alloc() {
    return [NSLocale alloc];
}

void* C_NSLocale_InitWithLocaleIdentifier(void* ptr, void* _string) {
    NSLocale* nSLocale = (NSLocale*)ptr;
    NSLocale* result = [nSLocale initWithLocaleIdentifier:(NSString*)_string];
    return result;
}

void* C_NSLocale_InitWithCoder(void* ptr, void* coder) {
    NSLocale* nSLocale = (NSLocale*)ptr;
    NSLocale* result = [nSLocale initWithCoder:(NSCoder*)coder];
    return result;
}

void* C_NSLocale_LocaleWithLocaleIdentifier(void* ident) {
    NSLocale* result = [NSLocale localeWithLocaleIdentifier:(NSString*)ident];
    return result;
}

void* C_NSLocale_Locale_CanonicalLocaleIdentifierFromString(void* _string) {
    NSString* result = [NSLocale canonicalLocaleIdentifierFromString:(NSString*)_string];
    return result;
}

void* C_NSLocale_Locale_CanonicalLanguageIdentifierFromString(void* _string) {
    NSString* result = [NSLocale canonicalLanguageIdentifierFromString:(NSString*)_string];
    return result;
}

void* C_NSLocale_LocaleIdentifierFromWindowsLocaleCode(uint32_t lcid) {
    NSString* result = [NSLocale localeIdentifierFromWindowsLocaleCode:lcid];
    return result;
}

uint32_t C_NSLocale_Locale_WindowsLocaleCodeFromLocaleIdentifier(void* localeIdentifier) {
    uint32_t result = [NSLocale windowsLocaleCodeFromLocaleIdentifier:(NSString*)localeIdentifier];
    return result;
}

void* C_NSLocale_LocalizedStringForLocaleIdentifier(void* ptr, void* localeIdentifier) {
    NSLocale* nSLocale = (NSLocale*)ptr;
    NSString* result = [nSLocale localizedStringForLocaleIdentifier:(NSString*)localeIdentifier];
    return result;
}

void* C_NSLocale_LocalizedStringForCountryCode(void* ptr, void* countryCode) {
    NSLocale* nSLocale = (NSLocale*)ptr;
    NSString* result = [nSLocale localizedStringForCountryCode:(NSString*)countryCode];
    return result;
}

void* C_NSLocale_LocalizedStringForLanguageCode(void* ptr, void* languageCode) {
    NSLocale* nSLocale = (NSLocale*)ptr;
    NSString* result = [nSLocale localizedStringForLanguageCode:(NSString*)languageCode];
    return result;
}

void* C_NSLocale_LocalizedStringForScriptCode(void* ptr, void* scriptCode) {
    NSLocale* nSLocale = (NSLocale*)ptr;
    NSString* result = [nSLocale localizedStringForScriptCode:(NSString*)scriptCode];
    return result;
}

void* C_NSLocale_LocalizedStringForVariantCode(void* ptr, void* variantCode) {
    NSLocale* nSLocale = (NSLocale*)ptr;
    NSString* result = [nSLocale localizedStringForVariantCode:(NSString*)variantCode];
    return result;
}

void* C_NSLocale_LocalizedStringForCollationIdentifier(void* ptr, void* collationIdentifier) {
    NSLocale* nSLocale = (NSLocale*)ptr;
    NSString* result = [nSLocale localizedStringForCollationIdentifier:(NSString*)collationIdentifier];
    return result;
}

void* C_NSLocale_LocalizedStringForCollatorIdentifier(void* ptr, void* collatorIdentifier) {
    NSLocale* nSLocale = (NSLocale*)ptr;
    NSString* result = [nSLocale localizedStringForCollatorIdentifier:(NSString*)collatorIdentifier];
    return result;
}

void* C_NSLocale_LocalizedStringForCurrencyCode(void* ptr, void* currencyCode) {
    NSLocale* nSLocale = (NSLocale*)ptr;
    NSString* result = [nSLocale localizedStringForCurrencyCode:(NSString*)currencyCode];
    return result;
}

void* C_NSLocale_LocalizedStringForCalendarIdentifier(void* ptr, void* calendarIdentifier) {
    NSLocale* nSLocale = (NSLocale*)ptr;
    NSString* result = [nSLocale localizedStringForCalendarIdentifier:(NSString*)calendarIdentifier];
    return result;
}

void* C_NSLocale_ObjectForKey(void* ptr, void* key) {
    NSLocale* nSLocale = (NSLocale*)ptr;
    id result = [nSLocale objectForKey:(NSString*)key];
    return result;
}

void* C_NSLocale_DisplayNameForKey_Value(void* ptr, void* key, void* value) {
    NSLocale* nSLocale = (NSLocale*)ptr;
    NSString* result = [nSLocale displayNameForKey:(NSString*)key value:(id)value];
    return result;
}

unsigned int C_NSLocale_Locale_CharacterDirectionForLanguage(void* isoLangCode) {
    NSLocaleLanguageDirection result = [NSLocale characterDirectionForLanguage:(NSString*)isoLangCode];
    return result;
}

unsigned int C_NSLocale_Locale_LineDirectionForLanguage(void* isoLangCode) {
    NSLocaleLanguageDirection result = [NSLocale lineDirectionForLanguage:(NSString*)isoLangCode];
    return result;
}

void* C_NSLocale_AutoupdatingCurrentLocale() {
    NSLocale* result = [NSLocale autoupdatingCurrentLocale];
    return result;
}

void* C_NSLocale_CurrentLocale() {
    NSLocale* result = [NSLocale currentLocale];
    return result;
}

void* C_NSLocale_SystemLocale() {
    NSLocale* result = [NSLocale systemLocale];
    return result;
}

Array C_NSLocale_Locale_AvailableLocaleIdentifiers() {
    NSArray* result = [NSLocale availableLocaleIdentifiers];
    int resultcount = [result count];
    void** resultData = malloc(resultcount * sizeof(void*));
    for (int i = 0; i < resultcount; i++) {
    	 void* p = [result objectAtIndex:i];
    	 resultData[i] = p;
    }
    Array resultArray;
    resultArray.data = resultData;
    resultArray.len = resultcount;
    return resultArray;
}

Array C_NSLocale_Locale_ISOCountryCodes() {
    NSArray* result = [NSLocale ISOCountryCodes];
    int resultcount = [result count];
    void** resultData = malloc(resultcount * sizeof(void*));
    for (int i = 0; i < resultcount; i++) {
    	 void* p = [result objectAtIndex:i];
    	 resultData[i] = p;
    }
    Array resultArray;
    resultArray.data = resultData;
    resultArray.len = resultcount;
    return resultArray;
}

Array C_NSLocale_Locale_ISOLanguageCodes() {
    NSArray* result = [NSLocale ISOLanguageCodes];
    int resultcount = [result count];
    void** resultData = malloc(resultcount * sizeof(void*));
    for (int i = 0; i < resultcount; i++) {
    	 void* p = [result objectAtIndex:i];
    	 resultData[i] = p;
    }
    Array resultArray;
    resultArray.data = resultData;
    resultArray.len = resultcount;
    return resultArray;
}

Array C_NSLocale_Locale_ISOCurrencyCodes() {
    NSArray* result = [NSLocale ISOCurrencyCodes];
    int resultcount = [result count];
    void** resultData = malloc(resultcount * sizeof(void*));
    for (int i = 0; i < resultcount; i++) {
    	 void* p = [result objectAtIndex:i];
    	 resultData[i] = p;
    }
    Array resultArray;
    resultArray.data = resultData;
    resultArray.len = resultcount;
    return resultArray;
}

Array C_NSLocale_Locale_CommonISOCurrencyCodes() {
    NSArray* result = [NSLocale commonISOCurrencyCodes];
    int resultcount = [result count];
    void** resultData = malloc(resultcount * sizeof(void*));
    for (int i = 0; i < resultcount; i++) {
    	 void* p = [result objectAtIndex:i];
    	 resultData[i] = p;
    }
    Array resultArray;
    resultArray.data = resultData;
    resultArray.len = resultcount;
    return resultArray;
}

void* C_NSLocale_LocaleIdentifier(void* ptr) {
    NSLocale* nSLocale = (NSLocale*)ptr;
    NSString* result = [nSLocale localeIdentifier];
    return result;
}

void* C_NSLocale_CountryCode(void* ptr) {
    NSLocale* nSLocale = (NSLocale*)ptr;
    NSString* result = [nSLocale countryCode];
    return result;
}

void* C_NSLocale_LanguageCode(void* ptr) {
    NSLocale* nSLocale = (NSLocale*)ptr;
    NSString* result = [nSLocale languageCode];
    return result;
}

void* C_NSLocale_ScriptCode(void* ptr) {
    NSLocale* nSLocale = (NSLocale*)ptr;
    NSString* result = [nSLocale scriptCode];
    return result;
}

void* C_NSLocale_VariantCode(void* ptr) {
    NSLocale* nSLocale = (NSLocale*)ptr;
    NSString* result = [nSLocale variantCode];
    return result;
}

void* C_NSLocale_ExemplarCharacterSet(void* ptr) {
    NSLocale* nSLocale = (NSLocale*)ptr;
    NSCharacterSet* result = [nSLocale exemplarCharacterSet];
    return result;
}

void* C_NSLocale_CollationIdentifier(void* ptr) {
    NSLocale* nSLocale = (NSLocale*)ptr;
    NSString* result = [nSLocale collationIdentifier];
    return result;
}

void* C_NSLocale_CollatorIdentifier(void* ptr) {
    NSLocale* nSLocale = (NSLocale*)ptr;
    NSString* result = [nSLocale collatorIdentifier];
    return result;
}

bool C_NSLocale_UsesMetricSystem(void* ptr) {
    NSLocale* nSLocale = (NSLocale*)ptr;
    BOOL result = [nSLocale usesMetricSystem];
    return result;
}

void* C_NSLocale_DecimalSeparator(void* ptr) {
    NSLocale* nSLocale = (NSLocale*)ptr;
    NSString* result = [nSLocale decimalSeparator];
    return result;
}

void* C_NSLocale_GroupingSeparator(void* ptr) {
    NSLocale* nSLocale = (NSLocale*)ptr;
    NSString* result = [nSLocale groupingSeparator];
    return result;
}

void* C_NSLocale_CurrencyCode(void* ptr) {
    NSLocale* nSLocale = (NSLocale*)ptr;
    NSString* result = [nSLocale currencyCode];
    return result;
}

void* C_NSLocale_CurrencySymbol(void* ptr) {
    NSLocale* nSLocale = (NSLocale*)ptr;
    NSString* result = [nSLocale currencySymbol];
    return result;
}

void* C_NSLocale_CalendarIdentifier(void* ptr) {
    NSLocale* nSLocale = (NSLocale*)ptr;
    NSString* result = [nSLocale calendarIdentifier];
    return result;
}

void* C_NSLocale_QuotationBeginDelimiter(void* ptr) {
    NSLocale* nSLocale = (NSLocale*)ptr;
    NSString* result = [nSLocale quotationBeginDelimiter];
    return result;
}

void* C_NSLocale_QuotationEndDelimiter(void* ptr) {
    NSLocale* nSLocale = (NSLocale*)ptr;
    NSString* result = [nSLocale quotationEndDelimiter];
    return result;
}

void* C_NSLocale_AlternateQuotationBeginDelimiter(void* ptr) {
    NSLocale* nSLocale = (NSLocale*)ptr;
    NSString* result = [nSLocale alternateQuotationBeginDelimiter];
    return result;
}

void* C_NSLocale_AlternateQuotationEndDelimiter(void* ptr) {
    NSLocale* nSLocale = (NSLocale*)ptr;
    NSString* result = [nSLocale alternateQuotationEndDelimiter];
    return result;
}

Array C_NSLocale_Locale_PreferredLanguages() {
    NSArray* result = [NSLocale preferredLanguages];
    int resultcount = [result count];
    void** resultData = malloc(resultcount * sizeof(void*));
    for (int i = 0; i < resultcount; i++) {
    	 void* p = [result objectAtIndex:i];
    	 resultData[i] = p;
    }
    Array resultArray;
    resultArray.data = resultData;
    resultArray.len = resultcount;
    return resultArray;
}
