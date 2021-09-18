#import "number_formatter.h"
#import <Foundation/NSDictionary.h>
#import <Foundation/NSArray.h>
#import <Foundation/NSNumberFormatter.h>

void* C_NumberFormatter_Alloc() {
    return [NSNumberFormatter alloc];
}

void* C_NSNumberFormatter_AllocNumberFormatter() {
    NSNumberFormatter* result_ = [NSNumberFormatter alloc];
    return result_;
}

void* C_NSNumberFormatter_Init(void* ptr) {
    NSNumberFormatter* nSNumberFormatter = (NSNumberFormatter*)ptr;
    NSNumberFormatter* result_ = [nSNumberFormatter init];
    return result_;
}

void* C_NSNumberFormatter_NewNumberFormatter() {
    NSNumberFormatter* result_ = [NSNumberFormatter new];
    return result_;
}

void* C_NSNumberFormatter_Autorelease(void* ptr) {
    NSNumberFormatter* nSNumberFormatter = (NSNumberFormatter*)ptr;
    NSNumberFormatter* result_ = [nSNumberFormatter autorelease];
    return result_;
}

void* C_NSNumberFormatter_Retain(void* ptr) {
    NSNumberFormatter* nSNumberFormatter = (NSNumberFormatter*)ptr;
    NSNumberFormatter* result_ = [nSNumberFormatter retain];
    return result_;
}

void C_NSNumberFormatter_NumberFormatter_SetDefaultFormatterBehavior(unsigned int behavior) {
    [NSNumberFormatter setDefaultFormatterBehavior:behavior];
}

unsigned int C_NSNumberFormatter_NumberFormatter_DefaultFormatterBehavior() {
    NSNumberFormatterBehavior result_ = [NSNumberFormatter defaultFormatterBehavior];
    return result_;
}

void* C_NSNumberFormatter_NumberFromString(void* ptr, void* _string) {
    NSNumberFormatter* nSNumberFormatter = (NSNumberFormatter*)ptr;
    NSNumber* result_ = [nSNumberFormatter numberFromString:(NSString*)_string];
    return result_;
}

void* C_NSNumberFormatter_StringFromNumber(void* ptr, void* number) {
    NSNumberFormatter* nSNumberFormatter = (NSNumberFormatter*)ptr;
    NSString* result_ = [nSNumberFormatter stringFromNumber:(NSNumber*)number];
    return result_;
}

void* C_NSNumberFormatter_NumberFormatter_LocalizedStringFromNumber_NumberStyle(void* num, unsigned int nstyle) {
    NSString* result_ = [NSNumberFormatter localizedStringFromNumber:(NSNumber*)num numberStyle:nstyle];
    return result_;
}

unsigned int C_NSNumberFormatter_FormatterBehavior(void* ptr) {
    NSNumberFormatter* nSNumberFormatter = (NSNumberFormatter*)ptr;
    NSNumberFormatterBehavior result_ = [nSNumberFormatter formatterBehavior];
    return result_;
}

void C_NSNumberFormatter_SetFormatterBehavior(void* ptr, unsigned int value) {
    NSNumberFormatter* nSNumberFormatter = (NSNumberFormatter*)ptr;
    [nSNumberFormatter setFormatterBehavior:value];
}

unsigned int C_NSNumberFormatter_NumberStyle(void* ptr) {
    NSNumberFormatter* nSNumberFormatter = (NSNumberFormatter*)ptr;
    NSNumberFormatterStyle result_ = [nSNumberFormatter numberStyle];
    return result_;
}

void C_NSNumberFormatter_SetNumberStyle(void* ptr, unsigned int value) {
    NSNumberFormatter* nSNumberFormatter = (NSNumberFormatter*)ptr;
    [nSNumberFormatter setNumberStyle:value];
}

bool C_NSNumberFormatter_GeneratesDecimalNumbers(void* ptr) {
    NSNumberFormatter* nSNumberFormatter = (NSNumberFormatter*)ptr;
    BOOL result_ = [nSNumberFormatter generatesDecimalNumbers];
    return result_;
}

void C_NSNumberFormatter_SetGeneratesDecimalNumbers(void* ptr, bool value) {
    NSNumberFormatter* nSNumberFormatter = (NSNumberFormatter*)ptr;
    [nSNumberFormatter setGeneratesDecimalNumbers:value];
}

bool C_NSNumberFormatter_LocalizesFormat(void* ptr) {
    NSNumberFormatter* nSNumberFormatter = (NSNumberFormatter*)ptr;
    BOOL result_ = [nSNumberFormatter localizesFormat];
    return result_;
}

void C_NSNumberFormatter_SetLocalizesFormat(void* ptr, bool value) {
    NSNumberFormatter* nSNumberFormatter = (NSNumberFormatter*)ptr;
    [nSNumberFormatter setLocalizesFormat:value];
}

void* C_NSNumberFormatter_Locale(void* ptr) {
    NSNumberFormatter* nSNumberFormatter = (NSNumberFormatter*)ptr;
    NSLocale* result_ = [nSNumberFormatter locale];
    return result_;
}

void C_NSNumberFormatter_SetLocale(void* ptr, void* value) {
    NSNumberFormatter* nSNumberFormatter = (NSNumberFormatter*)ptr;
    [nSNumberFormatter setLocale:(NSLocale*)value];
}

void* C_NSNumberFormatter_RoundingBehavior(void* ptr) {
    NSNumberFormatter* nSNumberFormatter = (NSNumberFormatter*)ptr;
    NSDecimalNumberHandler* result_ = [nSNumberFormatter roundingBehavior];
    return result_;
}

void C_NSNumberFormatter_SetRoundingBehavior(void* ptr, void* value) {
    NSNumberFormatter* nSNumberFormatter = (NSNumberFormatter*)ptr;
    [nSNumberFormatter setRoundingBehavior:(NSDecimalNumberHandler*)value];
}

void* C_NSNumberFormatter_RoundingIncrement(void* ptr) {
    NSNumberFormatter* nSNumberFormatter = (NSNumberFormatter*)ptr;
    NSNumber* result_ = [nSNumberFormatter roundingIncrement];
    return result_;
}

void C_NSNumberFormatter_SetRoundingIncrement(void* ptr, void* value) {
    NSNumberFormatter* nSNumberFormatter = (NSNumberFormatter*)ptr;
    [nSNumberFormatter setRoundingIncrement:(NSNumber*)value];
}

unsigned int C_NSNumberFormatter_RoundingMode(void* ptr) {
    NSNumberFormatter* nSNumberFormatter = (NSNumberFormatter*)ptr;
    NSNumberFormatterRoundingMode result_ = [nSNumberFormatter roundingMode];
    return result_;
}

void C_NSNumberFormatter_SetRoundingMode(void* ptr, unsigned int value) {
    NSNumberFormatter* nSNumberFormatter = (NSNumberFormatter*)ptr;
    [nSNumberFormatter setRoundingMode:value];
}

unsigned int C_NSNumberFormatter_MinimumIntegerDigits(void* ptr) {
    NSNumberFormatter* nSNumberFormatter = (NSNumberFormatter*)ptr;
    NSUInteger result_ = [nSNumberFormatter minimumIntegerDigits];
    return result_;
}

void C_NSNumberFormatter_SetMinimumIntegerDigits(void* ptr, unsigned int value) {
    NSNumberFormatter* nSNumberFormatter = (NSNumberFormatter*)ptr;
    [nSNumberFormatter setMinimumIntegerDigits:value];
}

unsigned int C_NSNumberFormatter_MaximumIntegerDigits(void* ptr) {
    NSNumberFormatter* nSNumberFormatter = (NSNumberFormatter*)ptr;
    NSUInteger result_ = [nSNumberFormatter maximumIntegerDigits];
    return result_;
}

void C_NSNumberFormatter_SetMaximumIntegerDigits(void* ptr, unsigned int value) {
    NSNumberFormatter* nSNumberFormatter = (NSNumberFormatter*)ptr;
    [nSNumberFormatter setMaximumIntegerDigits:value];
}

unsigned int C_NSNumberFormatter_MinimumFractionDigits(void* ptr) {
    NSNumberFormatter* nSNumberFormatter = (NSNumberFormatter*)ptr;
    NSUInteger result_ = [nSNumberFormatter minimumFractionDigits];
    return result_;
}

void C_NSNumberFormatter_SetMinimumFractionDigits(void* ptr, unsigned int value) {
    NSNumberFormatter* nSNumberFormatter = (NSNumberFormatter*)ptr;
    [nSNumberFormatter setMinimumFractionDigits:value];
}

unsigned int C_NSNumberFormatter_MaximumFractionDigits(void* ptr) {
    NSNumberFormatter* nSNumberFormatter = (NSNumberFormatter*)ptr;
    NSUInteger result_ = [nSNumberFormatter maximumFractionDigits];
    return result_;
}

void C_NSNumberFormatter_SetMaximumFractionDigits(void* ptr, unsigned int value) {
    NSNumberFormatter* nSNumberFormatter = (NSNumberFormatter*)ptr;
    [nSNumberFormatter setMaximumFractionDigits:value];
}

bool C_NSNumberFormatter_UsesSignificantDigits(void* ptr) {
    NSNumberFormatter* nSNumberFormatter = (NSNumberFormatter*)ptr;
    BOOL result_ = [nSNumberFormatter usesSignificantDigits];
    return result_;
}

void C_NSNumberFormatter_SetUsesSignificantDigits(void* ptr, bool value) {
    NSNumberFormatter* nSNumberFormatter = (NSNumberFormatter*)ptr;
    [nSNumberFormatter setUsesSignificantDigits:value];
}

unsigned int C_NSNumberFormatter_MinimumSignificantDigits(void* ptr) {
    NSNumberFormatter* nSNumberFormatter = (NSNumberFormatter*)ptr;
    NSUInteger result_ = [nSNumberFormatter minimumSignificantDigits];
    return result_;
}

void C_NSNumberFormatter_SetMinimumSignificantDigits(void* ptr, unsigned int value) {
    NSNumberFormatter* nSNumberFormatter = (NSNumberFormatter*)ptr;
    [nSNumberFormatter setMinimumSignificantDigits:value];
}

unsigned int C_NSNumberFormatter_MaximumSignificantDigits(void* ptr) {
    NSNumberFormatter* nSNumberFormatter = (NSNumberFormatter*)ptr;
    NSUInteger result_ = [nSNumberFormatter maximumSignificantDigits];
    return result_;
}

void C_NSNumberFormatter_SetMaximumSignificantDigits(void* ptr, unsigned int value) {
    NSNumberFormatter* nSNumberFormatter = (NSNumberFormatter*)ptr;
    [nSNumberFormatter setMaximumSignificantDigits:value];
}

void* C_NSNumberFormatter_Format(void* ptr) {
    NSNumberFormatter* nSNumberFormatter = (NSNumberFormatter*)ptr;
    NSString* result_ = [nSNumberFormatter format];
    return result_;
}

void C_NSNumberFormatter_SetFormat(void* ptr, void* value) {
    NSNumberFormatter* nSNumberFormatter = (NSNumberFormatter*)ptr;
    [nSNumberFormatter setFormat:(NSString*)value];
}

int C_NSNumberFormatter_FormattingContext(void* ptr) {
    NSNumberFormatter* nSNumberFormatter = (NSNumberFormatter*)ptr;
    NSFormattingContext result_ = [nSNumberFormatter formattingContext];
    return result_;
}

void C_NSNumberFormatter_SetFormattingContext(void* ptr, int value) {
    NSNumberFormatter* nSNumberFormatter = (NSNumberFormatter*)ptr;
    [nSNumberFormatter setFormattingContext:value];
}

unsigned int C_NSNumberFormatter_FormatWidth(void* ptr) {
    NSNumberFormatter* nSNumberFormatter = (NSNumberFormatter*)ptr;
    NSUInteger result_ = [nSNumberFormatter formatWidth];
    return result_;
}

void C_NSNumberFormatter_SetFormatWidth(void* ptr, unsigned int value) {
    NSNumberFormatter* nSNumberFormatter = (NSNumberFormatter*)ptr;
    [nSNumberFormatter setFormatWidth:value];
}

void* C_NSNumberFormatter_NegativeFormat(void* ptr) {
    NSNumberFormatter* nSNumberFormatter = (NSNumberFormatter*)ptr;
    NSString* result_ = [nSNumberFormatter negativeFormat];
    return result_;
}

void C_NSNumberFormatter_SetNegativeFormat(void* ptr, void* value) {
    NSNumberFormatter* nSNumberFormatter = (NSNumberFormatter*)ptr;
    [nSNumberFormatter setNegativeFormat:(NSString*)value];
}

void* C_NSNumberFormatter_PositiveFormat(void* ptr) {
    NSNumberFormatter* nSNumberFormatter = (NSNumberFormatter*)ptr;
    NSString* result_ = [nSNumberFormatter positiveFormat];
    return result_;
}

void C_NSNumberFormatter_SetPositiveFormat(void* ptr, void* value) {
    NSNumberFormatter* nSNumberFormatter = (NSNumberFormatter*)ptr;
    [nSNumberFormatter setPositiveFormat:(NSString*)value];
}

void* C_NSNumberFormatter_Multiplier(void* ptr) {
    NSNumberFormatter* nSNumberFormatter = (NSNumberFormatter*)ptr;
    NSNumber* result_ = [nSNumberFormatter multiplier];
    return result_;
}

void C_NSNumberFormatter_SetMultiplier(void* ptr, void* value) {
    NSNumberFormatter* nSNumberFormatter = (NSNumberFormatter*)ptr;
    [nSNumberFormatter setMultiplier:(NSNumber*)value];
}

void* C_NSNumberFormatter_PercentSymbol(void* ptr) {
    NSNumberFormatter* nSNumberFormatter = (NSNumberFormatter*)ptr;
    NSString* result_ = [nSNumberFormatter percentSymbol];
    return result_;
}

void C_NSNumberFormatter_SetPercentSymbol(void* ptr, void* value) {
    NSNumberFormatter* nSNumberFormatter = (NSNumberFormatter*)ptr;
    [nSNumberFormatter setPercentSymbol:(NSString*)value];
}

void* C_NSNumberFormatter_PerMillSymbol(void* ptr) {
    NSNumberFormatter* nSNumberFormatter = (NSNumberFormatter*)ptr;
    NSString* result_ = [nSNumberFormatter perMillSymbol];
    return result_;
}

void C_NSNumberFormatter_SetPerMillSymbol(void* ptr, void* value) {
    NSNumberFormatter* nSNumberFormatter = (NSNumberFormatter*)ptr;
    [nSNumberFormatter setPerMillSymbol:(NSString*)value];
}

void* C_NSNumberFormatter_MinusSign(void* ptr) {
    NSNumberFormatter* nSNumberFormatter = (NSNumberFormatter*)ptr;
    NSString* result_ = [nSNumberFormatter minusSign];
    return result_;
}

void C_NSNumberFormatter_SetMinusSign(void* ptr, void* value) {
    NSNumberFormatter* nSNumberFormatter = (NSNumberFormatter*)ptr;
    [nSNumberFormatter setMinusSign:(NSString*)value];
}

void* C_NSNumberFormatter_PlusSign(void* ptr) {
    NSNumberFormatter* nSNumberFormatter = (NSNumberFormatter*)ptr;
    NSString* result_ = [nSNumberFormatter plusSign];
    return result_;
}

void C_NSNumberFormatter_SetPlusSign(void* ptr, void* value) {
    NSNumberFormatter* nSNumberFormatter = (NSNumberFormatter*)ptr;
    [nSNumberFormatter setPlusSign:(NSString*)value];
}

void* C_NSNumberFormatter_ExponentSymbol(void* ptr) {
    NSNumberFormatter* nSNumberFormatter = (NSNumberFormatter*)ptr;
    NSString* result_ = [nSNumberFormatter exponentSymbol];
    return result_;
}

void C_NSNumberFormatter_SetExponentSymbol(void* ptr, void* value) {
    NSNumberFormatter* nSNumberFormatter = (NSNumberFormatter*)ptr;
    [nSNumberFormatter setExponentSymbol:(NSString*)value];
}

void* C_NSNumberFormatter_ZeroSymbol(void* ptr) {
    NSNumberFormatter* nSNumberFormatter = (NSNumberFormatter*)ptr;
    NSString* result_ = [nSNumberFormatter zeroSymbol];
    return result_;
}

void C_NSNumberFormatter_SetZeroSymbol(void* ptr, void* value) {
    NSNumberFormatter* nSNumberFormatter = (NSNumberFormatter*)ptr;
    [nSNumberFormatter setZeroSymbol:(NSString*)value];
}

void* C_NSNumberFormatter_NilSymbol(void* ptr) {
    NSNumberFormatter* nSNumberFormatter = (NSNumberFormatter*)ptr;
    NSString* result_ = [nSNumberFormatter nilSymbol];
    return result_;
}

void C_NSNumberFormatter_SetNilSymbol(void* ptr, void* value) {
    NSNumberFormatter* nSNumberFormatter = (NSNumberFormatter*)ptr;
    [nSNumberFormatter setNilSymbol:(NSString*)value];
}

void* C_NSNumberFormatter_NotANumberSymbol(void* ptr) {
    NSNumberFormatter* nSNumberFormatter = (NSNumberFormatter*)ptr;
    NSString* result_ = [nSNumberFormatter notANumberSymbol];
    return result_;
}

void C_NSNumberFormatter_SetNotANumberSymbol(void* ptr, void* value) {
    NSNumberFormatter* nSNumberFormatter = (NSNumberFormatter*)ptr;
    [nSNumberFormatter setNotANumberSymbol:(NSString*)value];
}

void* C_NSNumberFormatter_NegativeInfinitySymbol(void* ptr) {
    NSNumberFormatter* nSNumberFormatter = (NSNumberFormatter*)ptr;
    NSString* result_ = [nSNumberFormatter negativeInfinitySymbol];
    return result_;
}

void C_NSNumberFormatter_SetNegativeInfinitySymbol(void* ptr, void* value) {
    NSNumberFormatter* nSNumberFormatter = (NSNumberFormatter*)ptr;
    [nSNumberFormatter setNegativeInfinitySymbol:(NSString*)value];
}

void* C_NSNumberFormatter_PositiveInfinitySymbol(void* ptr) {
    NSNumberFormatter* nSNumberFormatter = (NSNumberFormatter*)ptr;
    NSString* result_ = [nSNumberFormatter positiveInfinitySymbol];
    return result_;
}

void C_NSNumberFormatter_SetPositiveInfinitySymbol(void* ptr, void* value) {
    NSNumberFormatter* nSNumberFormatter = (NSNumberFormatter*)ptr;
    [nSNumberFormatter setPositiveInfinitySymbol:(NSString*)value];
}

void* C_NSNumberFormatter_CurrencySymbol(void* ptr) {
    NSNumberFormatter* nSNumberFormatter = (NSNumberFormatter*)ptr;
    NSString* result_ = [nSNumberFormatter currencySymbol];
    return result_;
}

void C_NSNumberFormatter_SetCurrencySymbol(void* ptr, void* value) {
    NSNumberFormatter* nSNumberFormatter = (NSNumberFormatter*)ptr;
    [nSNumberFormatter setCurrencySymbol:(NSString*)value];
}

void* C_NSNumberFormatter_CurrencyCode(void* ptr) {
    NSNumberFormatter* nSNumberFormatter = (NSNumberFormatter*)ptr;
    NSString* result_ = [nSNumberFormatter currencyCode];
    return result_;
}

void C_NSNumberFormatter_SetCurrencyCode(void* ptr, void* value) {
    NSNumberFormatter* nSNumberFormatter = (NSNumberFormatter*)ptr;
    [nSNumberFormatter setCurrencyCode:(NSString*)value];
}

void* C_NSNumberFormatter_InternationalCurrencySymbol(void* ptr) {
    NSNumberFormatter* nSNumberFormatter = (NSNumberFormatter*)ptr;
    NSString* result_ = [nSNumberFormatter internationalCurrencySymbol];
    return result_;
}

void C_NSNumberFormatter_SetInternationalCurrencySymbol(void* ptr, void* value) {
    NSNumberFormatter* nSNumberFormatter = (NSNumberFormatter*)ptr;
    [nSNumberFormatter setInternationalCurrencySymbol:(NSString*)value];
}

void* C_NSNumberFormatter_CurrencyGroupingSeparator(void* ptr) {
    NSNumberFormatter* nSNumberFormatter = (NSNumberFormatter*)ptr;
    NSString* result_ = [nSNumberFormatter currencyGroupingSeparator];
    return result_;
}

void C_NSNumberFormatter_SetCurrencyGroupingSeparator(void* ptr, void* value) {
    NSNumberFormatter* nSNumberFormatter = (NSNumberFormatter*)ptr;
    [nSNumberFormatter setCurrencyGroupingSeparator:(NSString*)value];
}

void* C_NSNumberFormatter_PositivePrefix(void* ptr) {
    NSNumberFormatter* nSNumberFormatter = (NSNumberFormatter*)ptr;
    NSString* result_ = [nSNumberFormatter positivePrefix];
    return result_;
}

void C_NSNumberFormatter_SetPositivePrefix(void* ptr, void* value) {
    NSNumberFormatter* nSNumberFormatter = (NSNumberFormatter*)ptr;
    [nSNumberFormatter setPositivePrefix:(NSString*)value];
}

void* C_NSNumberFormatter_PositiveSuffix(void* ptr) {
    NSNumberFormatter* nSNumberFormatter = (NSNumberFormatter*)ptr;
    NSString* result_ = [nSNumberFormatter positiveSuffix];
    return result_;
}

void C_NSNumberFormatter_SetPositiveSuffix(void* ptr, void* value) {
    NSNumberFormatter* nSNumberFormatter = (NSNumberFormatter*)ptr;
    [nSNumberFormatter setPositiveSuffix:(NSString*)value];
}

void* C_NSNumberFormatter_NegativePrefix(void* ptr) {
    NSNumberFormatter* nSNumberFormatter = (NSNumberFormatter*)ptr;
    NSString* result_ = [nSNumberFormatter negativePrefix];
    return result_;
}

void C_NSNumberFormatter_SetNegativePrefix(void* ptr, void* value) {
    NSNumberFormatter* nSNumberFormatter = (NSNumberFormatter*)ptr;
    [nSNumberFormatter setNegativePrefix:(NSString*)value];
}

void* C_NSNumberFormatter_NegativeSuffix(void* ptr) {
    NSNumberFormatter* nSNumberFormatter = (NSNumberFormatter*)ptr;
    NSString* result_ = [nSNumberFormatter negativeSuffix];
    return result_;
}

void C_NSNumberFormatter_SetNegativeSuffix(void* ptr, void* value) {
    NSNumberFormatter* nSNumberFormatter = (NSNumberFormatter*)ptr;
    [nSNumberFormatter setNegativeSuffix:(NSString*)value];
}

Dictionary C_NSNumberFormatter_TextAttributesForNegativeValues(void* ptr) {
    NSNumberFormatter* nSNumberFormatter = (NSNumberFormatter*)ptr;
    NSDictionary* result_ = [nSNumberFormatter textAttributesForNegativeValues];
    Dictionary result_Array;
    NSArray * result_Keys = [result_ allKeys];
    int result_Count = [result_Keys count];
    if (result_Count > 0) {
    	void** result_KeyData = malloc(result_Count * sizeof(void*));
    	void** result_ValueData = malloc(result_Count * sizeof(void*));
    	for (int i = 0; i < result_Count; i++) {
    		NSString* kp = [result_Keys objectAtIndex:i];
    		id vp = result_[kp];
    		 result_KeyData[i] = kp;
    		 result_ValueData[i] = vp;
    	}
    	result_Array.key_data = result_KeyData;
    	result_Array.value_data = result_ValueData;
    	result_Array.len = result_Count;
    }
    return result_Array;
}

void C_NSNumberFormatter_SetTextAttributesForNegativeValues(void* ptr, Dictionary value) {
    NSNumberFormatter* nSNumberFormatter = (NSNumberFormatter*)ptr;
    NSMutableDictionary* objcValue = [NSMutableDictionary dictionaryWithCapacity:value.len];
    if (value.len > 0) {
    	void** valueKeyData = (void**)value.key_data;
    	void** valueValueData = (void**)value.value_data;
    	for (int i = 0; i < value.len; i++) {
    		void* kp = valueKeyData[i];
    		void* vp = valueValueData[i];
    		[objcValue setObject:(NSString*)kp forKey:(id)vp];
    	}
    }
    [nSNumberFormatter setTextAttributesForNegativeValues:objcValue];
}

Dictionary C_NSNumberFormatter_TextAttributesForPositiveValues(void* ptr) {
    NSNumberFormatter* nSNumberFormatter = (NSNumberFormatter*)ptr;
    NSDictionary* result_ = [nSNumberFormatter textAttributesForPositiveValues];
    Dictionary result_Array;
    NSArray * result_Keys = [result_ allKeys];
    int result_Count = [result_Keys count];
    if (result_Count > 0) {
    	void** result_KeyData = malloc(result_Count * sizeof(void*));
    	void** result_ValueData = malloc(result_Count * sizeof(void*));
    	for (int i = 0; i < result_Count; i++) {
    		NSString* kp = [result_Keys objectAtIndex:i];
    		id vp = result_[kp];
    		 result_KeyData[i] = kp;
    		 result_ValueData[i] = vp;
    	}
    	result_Array.key_data = result_KeyData;
    	result_Array.value_data = result_ValueData;
    	result_Array.len = result_Count;
    }
    return result_Array;
}

void C_NSNumberFormatter_SetTextAttributesForPositiveValues(void* ptr, Dictionary value) {
    NSNumberFormatter* nSNumberFormatter = (NSNumberFormatter*)ptr;
    NSMutableDictionary* objcValue = [NSMutableDictionary dictionaryWithCapacity:value.len];
    if (value.len > 0) {
    	void** valueKeyData = (void**)value.key_data;
    	void** valueValueData = (void**)value.value_data;
    	for (int i = 0; i < value.len; i++) {
    		void* kp = valueKeyData[i];
    		void* vp = valueValueData[i];
    		[objcValue setObject:(NSString*)kp forKey:(id)vp];
    	}
    }
    [nSNumberFormatter setTextAttributesForPositiveValues:objcValue];
}

void* C_NSNumberFormatter_AttributedStringForZero(void* ptr) {
    NSNumberFormatter* nSNumberFormatter = (NSNumberFormatter*)ptr;
    NSAttributedString* result_ = [nSNumberFormatter attributedStringForZero];
    return result_;
}

void C_NSNumberFormatter_SetAttributedStringForZero(void* ptr, void* value) {
    NSNumberFormatter* nSNumberFormatter = (NSNumberFormatter*)ptr;
    [nSNumberFormatter setAttributedStringForZero:(NSAttributedString*)value];
}

Dictionary C_NSNumberFormatter_TextAttributesForZero(void* ptr) {
    NSNumberFormatter* nSNumberFormatter = (NSNumberFormatter*)ptr;
    NSDictionary* result_ = [nSNumberFormatter textAttributesForZero];
    Dictionary result_Array;
    NSArray * result_Keys = [result_ allKeys];
    int result_Count = [result_Keys count];
    if (result_Count > 0) {
    	void** result_KeyData = malloc(result_Count * sizeof(void*));
    	void** result_ValueData = malloc(result_Count * sizeof(void*));
    	for (int i = 0; i < result_Count; i++) {
    		NSString* kp = [result_Keys objectAtIndex:i];
    		id vp = result_[kp];
    		 result_KeyData[i] = kp;
    		 result_ValueData[i] = vp;
    	}
    	result_Array.key_data = result_KeyData;
    	result_Array.value_data = result_ValueData;
    	result_Array.len = result_Count;
    }
    return result_Array;
}

void C_NSNumberFormatter_SetTextAttributesForZero(void* ptr, Dictionary value) {
    NSNumberFormatter* nSNumberFormatter = (NSNumberFormatter*)ptr;
    NSMutableDictionary* objcValue = [NSMutableDictionary dictionaryWithCapacity:value.len];
    if (value.len > 0) {
    	void** valueKeyData = (void**)value.key_data;
    	void** valueValueData = (void**)value.value_data;
    	for (int i = 0; i < value.len; i++) {
    		void* kp = valueKeyData[i];
    		void* vp = valueValueData[i];
    		[objcValue setObject:(NSString*)kp forKey:(id)vp];
    	}
    }
    [nSNumberFormatter setTextAttributesForZero:objcValue];
}

void* C_NSNumberFormatter_AttributedStringForNil(void* ptr) {
    NSNumberFormatter* nSNumberFormatter = (NSNumberFormatter*)ptr;
    NSAttributedString* result_ = [nSNumberFormatter attributedStringForNil];
    return result_;
}

void C_NSNumberFormatter_SetAttributedStringForNil(void* ptr, void* value) {
    NSNumberFormatter* nSNumberFormatter = (NSNumberFormatter*)ptr;
    [nSNumberFormatter setAttributedStringForNil:(NSAttributedString*)value];
}

Dictionary C_NSNumberFormatter_TextAttributesForNil(void* ptr) {
    NSNumberFormatter* nSNumberFormatter = (NSNumberFormatter*)ptr;
    NSDictionary* result_ = [nSNumberFormatter textAttributesForNil];
    Dictionary result_Array;
    NSArray * result_Keys = [result_ allKeys];
    int result_Count = [result_Keys count];
    if (result_Count > 0) {
    	void** result_KeyData = malloc(result_Count * sizeof(void*));
    	void** result_ValueData = malloc(result_Count * sizeof(void*));
    	for (int i = 0; i < result_Count; i++) {
    		NSString* kp = [result_Keys objectAtIndex:i];
    		id vp = result_[kp];
    		 result_KeyData[i] = kp;
    		 result_ValueData[i] = vp;
    	}
    	result_Array.key_data = result_KeyData;
    	result_Array.value_data = result_ValueData;
    	result_Array.len = result_Count;
    }
    return result_Array;
}

void C_NSNumberFormatter_SetTextAttributesForNil(void* ptr, Dictionary value) {
    NSNumberFormatter* nSNumberFormatter = (NSNumberFormatter*)ptr;
    NSMutableDictionary* objcValue = [NSMutableDictionary dictionaryWithCapacity:value.len];
    if (value.len > 0) {
    	void** valueKeyData = (void**)value.key_data;
    	void** valueValueData = (void**)value.value_data;
    	for (int i = 0; i < value.len; i++) {
    		void* kp = valueKeyData[i];
    		void* vp = valueValueData[i];
    		[objcValue setObject:(NSString*)kp forKey:(id)vp];
    	}
    }
    [nSNumberFormatter setTextAttributesForNil:objcValue];
}

void* C_NSNumberFormatter_AttributedStringForNotANumber(void* ptr) {
    NSNumberFormatter* nSNumberFormatter = (NSNumberFormatter*)ptr;
    NSAttributedString* result_ = [nSNumberFormatter attributedStringForNotANumber];
    return result_;
}

void C_NSNumberFormatter_SetAttributedStringForNotANumber(void* ptr, void* value) {
    NSNumberFormatter* nSNumberFormatter = (NSNumberFormatter*)ptr;
    [nSNumberFormatter setAttributedStringForNotANumber:(NSAttributedString*)value];
}

Dictionary C_NSNumberFormatter_TextAttributesForNotANumber(void* ptr) {
    NSNumberFormatter* nSNumberFormatter = (NSNumberFormatter*)ptr;
    NSDictionary* result_ = [nSNumberFormatter textAttributesForNotANumber];
    Dictionary result_Array;
    NSArray * result_Keys = [result_ allKeys];
    int result_Count = [result_Keys count];
    if (result_Count > 0) {
    	void** result_KeyData = malloc(result_Count * sizeof(void*));
    	void** result_ValueData = malloc(result_Count * sizeof(void*));
    	for (int i = 0; i < result_Count; i++) {
    		NSString* kp = [result_Keys objectAtIndex:i];
    		id vp = result_[kp];
    		 result_KeyData[i] = kp;
    		 result_ValueData[i] = vp;
    	}
    	result_Array.key_data = result_KeyData;
    	result_Array.value_data = result_ValueData;
    	result_Array.len = result_Count;
    }
    return result_Array;
}

void C_NSNumberFormatter_SetTextAttributesForNotANumber(void* ptr, Dictionary value) {
    NSNumberFormatter* nSNumberFormatter = (NSNumberFormatter*)ptr;
    NSMutableDictionary* objcValue = [NSMutableDictionary dictionaryWithCapacity:value.len];
    if (value.len > 0) {
    	void** valueKeyData = (void**)value.key_data;
    	void** valueValueData = (void**)value.value_data;
    	for (int i = 0; i < value.len; i++) {
    		void* kp = valueKeyData[i];
    		void* vp = valueValueData[i];
    		[objcValue setObject:(NSString*)kp forKey:(id)vp];
    	}
    }
    [nSNumberFormatter setTextAttributesForNotANumber:objcValue];
}

Dictionary C_NSNumberFormatter_TextAttributesForPositiveInfinity(void* ptr) {
    NSNumberFormatter* nSNumberFormatter = (NSNumberFormatter*)ptr;
    NSDictionary* result_ = [nSNumberFormatter textAttributesForPositiveInfinity];
    Dictionary result_Array;
    NSArray * result_Keys = [result_ allKeys];
    int result_Count = [result_Keys count];
    if (result_Count > 0) {
    	void** result_KeyData = malloc(result_Count * sizeof(void*));
    	void** result_ValueData = malloc(result_Count * sizeof(void*));
    	for (int i = 0; i < result_Count; i++) {
    		NSString* kp = [result_Keys objectAtIndex:i];
    		id vp = result_[kp];
    		 result_KeyData[i] = kp;
    		 result_ValueData[i] = vp;
    	}
    	result_Array.key_data = result_KeyData;
    	result_Array.value_data = result_ValueData;
    	result_Array.len = result_Count;
    }
    return result_Array;
}

void C_NSNumberFormatter_SetTextAttributesForPositiveInfinity(void* ptr, Dictionary value) {
    NSNumberFormatter* nSNumberFormatter = (NSNumberFormatter*)ptr;
    NSMutableDictionary* objcValue = [NSMutableDictionary dictionaryWithCapacity:value.len];
    if (value.len > 0) {
    	void** valueKeyData = (void**)value.key_data;
    	void** valueValueData = (void**)value.value_data;
    	for (int i = 0; i < value.len; i++) {
    		void* kp = valueKeyData[i];
    		void* vp = valueValueData[i];
    		[objcValue setObject:(NSString*)kp forKey:(id)vp];
    	}
    }
    [nSNumberFormatter setTextAttributesForPositiveInfinity:objcValue];
}

Dictionary C_NSNumberFormatter_TextAttributesForNegativeInfinity(void* ptr) {
    NSNumberFormatter* nSNumberFormatter = (NSNumberFormatter*)ptr;
    NSDictionary* result_ = [nSNumberFormatter textAttributesForNegativeInfinity];
    Dictionary result_Array;
    NSArray * result_Keys = [result_ allKeys];
    int result_Count = [result_Keys count];
    if (result_Count > 0) {
    	void** result_KeyData = malloc(result_Count * sizeof(void*));
    	void** result_ValueData = malloc(result_Count * sizeof(void*));
    	for (int i = 0; i < result_Count; i++) {
    		NSString* kp = [result_Keys objectAtIndex:i];
    		id vp = result_[kp];
    		 result_KeyData[i] = kp;
    		 result_ValueData[i] = vp;
    	}
    	result_Array.key_data = result_KeyData;
    	result_Array.value_data = result_ValueData;
    	result_Array.len = result_Count;
    }
    return result_Array;
}

void C_NSNumberFormatter_SetTextAttributesForNegativeInfinity(void* ptr, Dictionary value) {
    NSNumberFormatter* nSNumberFormatter = (NSNumberFormatter*)ptr;
    NSMutableDictionary* objcValue = [NSMutableDictionary dictionaryWithCapacity:value.len];
    if (value.len > 0) {
    	void** valueKeyData = (void**)value.key_data;
    	void** valueValueData = (void**)value.value_data;
    	for (int i = 0; i < value.len; i++) {
    		void* kp = valueKeyData[i];
    		void* vp = valueValueData[i];
    		[objcValue setObject:(NSString*)kp forKey:(id)vp];
    	}
    }
    [nSNumberFormatter setTextAttributesForNegativeInfinity:objcValue];
}

void* C_NSNumberFormatter_GroupingSeparator(void* ptr) {
    NSNumberFormatter* nSNumberFormatter = (NSNumberFormatter*)ptr;
    NSString* result_ = [nSNumberFormatter groupingSeparator];
    return result_;
}

void C_NSNumberFormatter_SetGroupingSeparator(void* ptr, void* value) {
    NSNumberFormatter* nSNumberFormatter = (NSNumberFormatter*)ptr;
    [nSNumberFormatter setGroupingSeparator:(NSString*)value];
}

bool C_NSNumberFormatter_UsesGroupingSeparator(void* ptr) {
    NSNumberFormatter* nSNumberFormatter = (NSNumberFormatter*)ptr;
    BOOL result_ = [nSNumberFormatter usesGroupingSeparator];
    return result_;
}

void C_NSNumberFormatter_SetUsesGroupingSeparator(void* ptr, bool value) {
    NSNumberFormatter* nSNumberFormatter = (NSNumberFormatter*)ptr;
    [nSNumberFormatter setUsesGroupingSeparator:value];
}

void* C_NSNumberFormatter_ThousandSeparator(void* ptr) {
    NSNumberFormatter* nSNumberFormatter = (NSNumberFormatter*)ptr;
    NSString* result_ = [nSNumberFormatter thousandSeparator];
    return result_;
}

void C_NSNumberFormatter_SetThousandSeparator(void* ptr, void* value) {
    NSNumberFormatter* nSNumberFormatter = (NSNumberFormatter*)ptr;
    [nSNumberFormatter setThousandSeparator:(NSString*)value];
}

bool C_NSNumberFormatter_HasThousandSeparators(void* ptr) {
    NSNumberFormatter* nSNumberFormatter = (NSNumberFormatter*)ptr;
    BOOL result_ = [nSNumberFormatter hasThousandSeparators];
    return result_;
}

void C_NSNumberFormatter_SetHasThousandSeparators(void* ptr, bool value) {
    NSNumberFormatter* nSNumberFormatter = (NSNumberFormatter*)ptr;
    [nSNumberFormatter setHasThousandSeparators:value];
}

void* C_NSNumberFormatter_DecimalSeparator(void* ptr) {
    NSNumberFormatter* nSNumberFormatter = (NSNumberFormatter*)ptr;
    NSString* result_ = [nSNumberFormatter decimalSeparator];
    return result_;
}

void C_NSNumberFormatter_SetDecimalSeparator(void* ptr, void* value) {
    NSNumberFormatter* nSNumberFormatter = (NSNumberFormatter*)ptr;
    [nSNumberFormatter setDecimalSeparator:(NSString*)value];
}

bool C_NSNumberFormatter_AlwaysShowsDecimalSeparator(void* ptr) {
    NSNumberFormatter* nSNumberFormatter = (NSNumberFormatter*)ptr;
    BOOL result_ = [nSNumberFormatter alwaysShowsDecimalSeparator];
    return result_;
}

void C_NSNumberFormatter_SetAlwaysShowsDecimalSeparator(void* ptr, bool value) {
    NSNumberFormatter* nSNumberFormatter = (NSNumberFormatter*)ptr;
    [nSNumberFormatter setAlwaysShowsDecimalSeparator:value];
}

void* C_NSNumberFormatter_CurrencyDecimalSeparator(void* ptr) {
    NSNumberFormatter* nSNumberFormatter = (NSNumberFormatter*)ptr;
    NSString* result_ = [nSNumberFormatter currencyDecimalSeparator];
    return result_;
}

void C_NSNumberFormatter_SetCurrencyDecimalSeparator(void* ptr, void* value) {
    NSNumberFormatter* nSNumberFormatter = (NSNumberFormatter*)ptr;
    [nSNumberFormatter setCurrencyDecimalSeparator:(NSString*)value];
}

unsigned int C_NSNumberFormatter_GroupingSize(void* ptr) {
    NSNumberFormatter* nSNumberFormatter = (NSNumberFormatter*)ptr;
    NSUInteger result_ = [nSNumberFormatter groupingSize];
    return result_;
}

void C_NSNumberFormatter_SetGroupingSize(void* ptr, unsigned int value) {
    NSNumberFormatter* nSNumberFormatter = (NSNumberFormatter*)ptr;
    [nSNumberFormatter setGroupingSize:value];
}

unsigned int C_NSNumberFormatter_SecondaryGroupingSize(void* ptr) {
    NSNumberFormatter* nSNumberFormatter = (NSNumberFormatter*)ptr;
    NSUInteger result_ = [nSNumberFormatter secondaryGroupingSize];
    return result_;
}

void C_NSNumberFormatter_SetSecondaryGroupingSize(void* ptr, unsigned int value) {
    NSNumberFormatter* nSNumberFormatter = (NSNumberFormatter*)ptr;
    [nSNumberFormatter setSecondaryGroupingSize:value];
}

void* C_NSNumberFormatter_PaddingCharacter(void* ptr) {
    NSNumberFormatter* nSNumberFormatter = (NSNumberFormatter*)ptr;
    NSString* result_ = [nSNumberFormatter paddingCharacter];
    return result_;
}

void C_NSNumberFormatter_SetPaddingCharacter(void* ptr, void* value) {
    NSNumberFormatter* nSNumberFormatter = (NSNumberFormatter*)ptr;
    [nSNumberFormatter setPaddingCharacter:(NSString*)value];
}

unsigned int C_NSNumberFormatter_PaddingPosition(void* ptr) {
    NSNumberFormatter* nSNumberFormatter = (NSNumberFormatter*)ptr;
    NSNumberFormatterPadPosition result_ = [nSNumberFormatter paddingPosition];
    return result_;
}

void C_NSNumberFormatter_SetPaddingPosition(void* ptr, unsigned int value) {
    NSNumberFormatter* nSNumberFormatter = (NSNumberFormatter*)ptr;
    [nSNumberFormatter setPaddingPosition:value];
}

bool C_NSNumberFormatter_AllowsFloats(void* ptr) {
    NSNumberFormatter* nSNumberFormatter = (NSNumberFormatter*)ptr;
    BOOL result_ = [nSNumberFormatter allowsFloats];
    return result_;
}

void C_NSNumberFormatter_SetAllowsFloats(void* ptr, bool value) {
    NSNumberFormatter* nSNumberFormatter = (NSNumberFormatter*)ptr;
    [nSNumberFormatter setAllowsFloats:value];
}

void* C_NSNumberFormatter_Minimum(void* ptr) {
    NSNumberFormatter* nSNumberFormatter = (NSNumberFormatter*)ptr;
    NSNumber* result_ = [nSNumberFormatter minimum];
    return result_;
}

void C_NSNumberFormatter_SetMinimum(void* ptr, void* value) {
    NSNumberFormatter* nSNumberFormatter = (NSNumberFormatter*)ptr;
    [nSNumberFormatter setMinimum:(NSNumber*)value];
}

void* C_NSNumberFormatter_Maximum(void* ptr) {
    NSNumberFormatter* nSNumberFormatter = (NSNumberFormatter*)ptr;
    NSNumber* result_ = [nSNumberFormatter maximum];
    return result_;
}

void C_NSNumberFormatter_SetMaximum(void* ptr, void* value) {
    NSNumberFormatter* nSNumberFormatter = (NSNumberFormatter*)ptr;
    [nSNumberFormatter setMaximum:(NSNumber*)value];
}

bool C_NSNumberFormatter_IsLenient(void* ptr) {
    NSNumberFormatter* nSNumberFormatter = (NSNumberFormatter*)ptr;
    BOOL result_ = [nSNumberFormatter isLenient];
    return result_;
}

void C_NSNumberFormatter_SetLenient(void* ptr, bool value) {
    NSNumberFormatter* nSNumberFormatter = (NSNumberFormatter*)ptr;
    [nSNumberFormatter setLenient:value];
}

bool C_NSNumberFormatter_IsPartialStringValidationEnabled(void* ptr) {
    NSNumberFormatter* nSNumberFormatter = (NSNumberFormatter*)ptr;
    BOOL result_ = [nSNumberFormatter isPartialStringValidationEnabled];
    return result_;
}

void C_NSNumberFormatter_SetPartialStringValidationEnabled(void* ptr, bool value) {
    NSNumberFormatter* nSNumberFormatter = (NSNumberFormatter*)ptr;
    [nSNumberFormatter setPartialStringValidationEnabled:value];
}
