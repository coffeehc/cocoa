#import <Foundation/Foundation.h>
#import "locale.h"

void* C_Locale_Alloc() {
    return [NSLocale alloc];
}

void* C_NSLocale_LocaleWithLocaleIdentifier(void* ident) {
    NSLocale* result_ = [NSLocale localeWithLocaleIdentifier:(NSString*)ident];
    return result_;
}

void* C_NSLocale_InitWithLocaleIdentifier(void* ptr, void* _string) {
    NSLocale* nSLocale = (NSLocale*)ptr;
    NSLocale* result_ = [nSLocale initWithLocaleIdentifier:(NSString*)_string];
    return result_;
}

void* C_NSLocale_InitWithCoder(void* ptr, void* coder) {
    NSLocale* nSLocale = (NSLocale*)ptr;
    NSLocale* result_ = [nSLocale initWithCoder:(NSCoder*)coder];
    return result_;
}

void* C_NSLocale_AllocLocale() {
    NSLocale* result_ = [NSLocale alloc];
    return result_;
}

void* C_NSLocale_Autorelease(void* ptr) {
    NSLocale* nSLocale = (NSLocale*)ptr;
    NSLocale* result_ = [nSLocale autorelease];
    return result_;
}

void* C_NSLocale_Retain(void* ptr) {
    NSLocale* nSLocale = (NSLocale*)ptr;
    NSLocale* result_ = [nSLocale retain];
    return result_;
}

void* C_NSLocale_CanonicalLocaleIdentifierFromString(void* _string) {
    NSString* result_ = [NSLocale canonicalLocaleIdentifierFromString:(NSString*)_string];
    return result_;
}

Dictionary C_NSLocale_ComponentsFromLocaleIdentifier(void* _string) {
    NSDictionary* result_ = [NSLocale componentsFromLocaleIdentifier:(NSString*)_string];
    Dictionary result_Array;
    NSArray * result_Keys = [result_ allKeys];
    int result_Count = [result_Keys count];
    if (result_Count > 0) {
    	void** result_KeyData = malloc(result_Count * sizeof(void*));
    	void** result_ValueData = malloc(result_Count * sizeof(void*));
    	for (int i = 0; i < result_Count; i++) {
    		NSString* kp = [result_Keys objectAtIndex:i];
    		NSString* vp = result_[kp];
    		 result_KeyData[i] = kp;
    		 result_ValueData[i] = vp;
    	}
    	result_Array.key_data = result_KeyData;
    	result_Array.value_data = result_ValueData;
    	result_Array.len = result_Count;
    }
    return result_Array;
}

void* C_NSLocale_LocaleIdentifierFromComponents(Dictionary dict) {
    NSMutableDictionary* objcDict = [NSMutableDictionary dictionaryWithCapacity:dict.len];
    if (dict.len > 0) {
    	void** dictKeyData = (void**)dict.key_data;
    	void** dictValueData = (void**)dict.value_data;
    	for (int i = 0; i < dict.len; i++) {
    		void* kp = dictKeyData[i];
    		void* vp = dictValueData[i];
    		[objcDict setObject:(NSString*)(NSString*)kp forKey:(NSString*)(NSString*)vp];
    	}
    }
    NSString* result_ = [NSLocale localeIdentifierFromComponents:objcDict];
    return result_;
}

void* C_NSLocale_Locale_CanonicalLanguageIdentifierFromString(void* _string) {
    NSString* result_ = [NSLocale canonicalLanguageIdentifierFromString:(NSString*)_string];
    return result_;
}

void* C_NSLocale_LocaleIdentifierFromWindowsLocaleCode(uint32_t lcid) {
    NSString* result_ = [NSLocale localeIdentifierFromWindowsLocaleCode:lcid];
    return result_;
}

uint32_t C_NSLocale_WindowsLocaleCodeFromLocaleIdentifier(void* localeIdentifier) {
    uint32_t result_ = [NSLocale windowsLocaleCodeFromLocaleIdentifier:(NSString*)localeIdentifier];
    return result_;
}

void* C_NSLocale_LocalizedStringForLocaleIdentifier(void* ptr, void* localeIdentifier) {
    NSLocale* nSLocale = (NSLocale*)ptr;
    NSString* result_ = [nSLocale localizedStringForLocaleIdentifier:(NSString*)localeIdentifier];
    return result_;
}

void* C_NSLocale_LocalizedStringForCountryCode(void* ptr, void* countryCode) {
    NSLocale* nSLocale = (NSLocale*)ptr;
    NSString* result_ = [nSLocale localizedStringForCountryCode:(NSString*)countryCode];
    return result_;
}

void* C_NSLocale_LocalizedStringForLanguageCode(void* ptr, void* languageCode) {
    NSLocale* nSLocale = (NSLocale*)ptr;
    NSString* result_ = [nSLocale localizedStringForLanguageCode:(NSString*)languageCode];
    return result_;
}

void* C_NSLocale_LocalizedStringForScriptCode(void* ptr, void* scriptCode) {
    NSLocale* nSLocale = (NSLocale*)ptr;
    NSString* result_ = [nSLocale localizedStringForScriptCode:(NSString*)scriptCode];
    return result_;
}

void* C_NSLocale_LocalizedStringForVariantCode(void* ptr, void* variantCode) {
    NSLocale* nSLocale = (NSLocale*)ptr;
    NSString* result_ = [nSLocale localizedStringForVariantCode:(NSString*)variantCode];
    return result_;
}

void* C_NSLocale_LocalizedStringForCollationIdentifier(void* ptr, void* collationIdentifier) {
    NSLocale* nSLocale = (NSLocale*)ptr;
    NSString* result_ = [nSLocale localizedStringForCollationIdentifier:(NSString*)collationIdentifier];
    return result_;
}

void* C_NSLocale_LocalizedStringForCollatorIdentifier(void* ptr, void* collatorIdentifier) {
    NSLocale* nSLocale = (NSLocale*)ptr;
    NSString* result_ = [nSLocale localizedStringForCollatorIdentifier:(NSString*)collatorIdentifier];
    return result_;
}

void* C_NSLocale_LocalizedStringForCurrencyCode(void* ptr, void* currencyCode) {
    NSLocale* nSLocale = (NSLocale*)ptr;
    NSString* result_ = [nSLocale localizedStringForCurrencyCode:(NSString*)currencyCode];
    return result_;
}

void* C_NSLocale_LocalizedStringForCalendarIdentifier(void* ptr, void* calendarIdentifier) {
    NSLocale* nSLocale = (NSLocale*)ptr;
    NSString* result_ = [nSLocale localizedStringForCalendarIdentifier:(NSString*)calendarIdentifier];
    return result_;
}

void* C_NSLocale_ObjectForKey(void* ptr, void* key) {
    NSLocale* nSLocale = (NSLocale*)ptr;
    id result_ = [nSLocale objectForKey:(NSString*)key];
    return result_;
}

void* C_NSLocale_DisplayNameForKey_Value(void* ptr, void* key, void* value) {
    NSLocale* nSLocale = (NSLocale*)ptr;
    NSString* result_ = [nSLocale displayNameForKey:(NSString*)key value:(id)value];
    return result_;
}

unsigned int C_NSLocale_Locale_CharacterDirectionForLanguage(void* isoLangCode) {
    NSLocaleLanguageDirection result_ = [NSLocale characterDirectionForLanguage:(NSString*)isoLangCode];
    return result_;
}

unsigned int C_NSLocale_Locale_LineDirectionForLanguage(void* isoLangCode) {
    NSLocaleLanguageDirection result_ = [NSLocale lineDirectionForLanguage:(NSString*)isoLangCode];
    return result_;
}

void* C_NSLocale_AutoupdatingCurrentLocale() {
    NSLocale* result_ = [NSLocale autoupdatingCurrentLocale];
    return result_;
}

void* C_NSLocale_CurrentLocale() {
    NSLocale* result_ = [NSLocale currentLocale];
    return result_;
}

void* C_NSLocale_SystemLocale() {
    NSLocale* result_ = [NSLocale systemLocale];
    return result_;
}

Array C_NSLocale_AvailableLocaleIdentifiers() {
    NSArray* result_ = [NSLocale availableLocaleIdentifiers];
    Array result_Array;
    int result_count = [result_ count];
    if (result_count > 0) {
    	void** result_Data = malloc(result_count * sizeof(void*));
    	for (int i = 0; i < result_count; i++) {
    		 void* p = [result_ objectAtIndex:i];
    		 result_Data[i] = p;
    	}
    	result_Array.data = result_Data;
    	result_Array.len = result_count;
    }
    return result_Array;
}

Array C_NSLocale_Locale_ISOCountryCodes() {
    NSArray* result_ = [NSLocale ISOCountryCodes];
    Array result_Array;
    int result_count = [result_ count];
    if (result_count > 0) {
    	void** result_Data = malloc(result_count * sizeof(void*));
    	for (int i = 0; i < result_count; i++) {
    		 void* p = [result_ objectAtIndex:i];
    		 result_Data[i] = p;
    	}
    	result_Array.data = result_Data;
    	result_Array.len = result_count;
    }
    return result_Array;
}

Array C_NSLocale_Locale_ISOLanguageCodes() {
    NSArray* result_ = [NSLocale ISOLanguageCodes];
    Array result_Array;
    int result_count = [result_ count];
    if (result_count > 0) {
    	void** result_Data = malloc(result_count * sizeof(void*));
    	for (int i = 0; i < result_count; i++) {
    		 void* p = [result_ objectAtIndex:i];
    		 result_Data[i] = p;
    	}
    	result_Array.data = result_Data;
    	result_Array.len = result_count;
    }
    return result_Array;
}

Array C_NSLocale_Locale_ISOCurrencyCodes() {
    NSArray* result_ = [NSLocale ISOCurrencyCodes];
    Array result_Array;
    int result_count = [result_ count];
    if (result_count > 0) {
    	void** result_Data = malloc(result_count * sizeof(void*));
    	for (int i = 0; i < result_count; i++) {
    		 void* p = [result_ objectAtIndex:i];
    		 result_Data[i] = p;
    	}
    	result_Array.data = result_Data;
    	result_Array.len = result_count;
    }
    return result_Array;
}

Array C_NSLocale_Locale_CommonISOCurrencyCodes() {
    NSArray* result_ = [NSLocale commonISOCurrencyCodes];
    Array result_Array;
    int result_count = [result_ count];
    if (result_count > 0) {
    	void** result_Data = malloc(result_count * sizeof(void*));
    	for (int i = 0; i < result_count; i++) {
    		 void* p = [result_ objectAtIndex:i];
    		 result_Data[i] = p;
    	}
    	result_Array.data = result_Data;
    	result_Array.len = result_count;
    }
    return result_Array;
}

void* C_NSLocale_LocaleIdentifier(void* ptr) {
    NSLocale* nSLocale = (NSLocale*)ptr;
    NSString* result_ = [nSLocale localeIdentifier];
    return result_;
}

void* C_NSLocale_CountryCode(void* ptr) {
    NSLocale* nSLocale = (NSLocale*)ptr;
    NSString* result_ = [nSLocale countryCode];
    return result_;
}

void* C_NSLocale_LanguageCode(void* ptr) {
    NSLocale* nSLocale = (NSLocale*)ptr;
    NSString* result_ = [nSLocale languageCode];
    return result_;
}

void* C_NSLocale_ScriptCode(void* ptr) {
    NSLocale* nSLocale = (NSLocale*)ptr;
    NSString* result_ = [nSLocale scriptCode];
    return result_;
}

void* C_NSLocale_VariantCode(void* ptr) {
    NSLocale* nSLocale = (NSLocale*)ptr;
    NSString* result_ = [nSLocale variantCode];
    return result_;
}

void* C_NSLocale_ExemplarCharacterSet(void* ptr) {
    NSLocale* nSLocale = (NSLocale*)ptr;
    NSCharacterSet* result_ = [nSLocale exemplarCharacterSet];
    return result_;
}

void* C_NSLocale_CollationIdentifier(void* ptr) {
    NSLocale* nSLocale = (NSLocale*)ptr;
    NSString* result_ = [nSLocale collationIdentifier];
    return result_;
}

void* C_NSLocale_CollatorIdentifier(void* ptr) {
    NSLocale* nSLocale = (NSLocale*)ptr;
    NSString* result_ = [nSLocale collatorIdentifier];
    return result_;
}

bool C_NSLocale_UsesMetricSystem(void* ptr) {
    NSLocale* nSLocale = (NSLocale*)ptr;
    BOOL result_ = [nSLocale usesMetricSystem];
    return result_;
}

void* C_NSLocale_DecimalSeparator(void* ptr) {
    NSLocale* nSLocale = (NSLocale*)ptr;
    NSString* result_ = [nSLocale decimalSeparator];
    return result_;
}

void* C_NSLocale_GroupingSeparator(void* ptr) {
    NSLocale* nSLocale = (NSLocale*)ptr;
    NSString* result_ = [nSLocale groupingSeparator];
    return result_;
}

void* C_NSLocale_CurrencyCode(void* ptr) {
    NSLocale* nSLocale = (NSLocale*)ptr;
    NSString* result_ = [nSLocale currencyCode];
    return result_;
}

void* C_NSLocale_CurrencySymbol(void* ptr) {
    NSLocale* nSLocale = (NSLocale*)ptr;
    NSString* result_ = [nSLocale currencySymbol];
    return result_;
}

void* C_NSLocale_CalendarIdentifier(void* ptr) {
    NSLocale* nSLocale = (NSLocale*)ptr;
    NSString* result_ = [nSLocale calendarIdentifier];
    return result_;
}

void* C_NSLocale_QuotationBeginDelimiter(void* ptr) {
    NSLocale* nSLocale = (NSLocale*)ptr;
    NSString* result_ = [nSLocale quotationBeginDelimiter];
    return result_;
}

void* C_NSLocale_QuotationEndDelimiter(void* ptr) {
    NSLocale* nSLocale = (NSLocale*)ptr;
    NSString* result_ = [nSLocale quotationEndDelimiter];
    return result_;
}

void* C_NSLocale_AlternateQuotationBeginDelimiter(void* ptr) {
    NSLocale* nSLocale = (NSLocale*)ptr;
    NSString* result_ = [nSLocale alternateQuotationBeginDelimiter];
    return result_;
}

void* C_NSLocale_AlternateQuotationEndDelimiter(void* ptr) {
    NSLocale* nSLocale = (NSLocale*)ptr;
    NSString* result_ = [nSLocale alternateQuotationEndDelimiter];
    return result_;
}

Array C_NSLocale_Locale_PreferredLanguages() {
    NSArray* result_ = [NSLocale preferredLanguages];
    Array result_Array;
    int result_count = [result_ count];
    if (result_count > 0) {
    	void** result_Data = malloc(result_count * sizeof(void*));
    	for (int i = 0; i < result_count; i++) {
    		 void* p = [result_ objectAtIndex:i];
    		 result_Data[i] = p;
    	}
    	result_Array.data = result_Data;
    	result_Array.len = result_count;
    }
    return result_Array;
}
