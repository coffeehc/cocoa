package foundation

// #include "locale.h"
import "C"
import (
	"github.com/hsiafan/cocoa/objc"
	"unsafe"
)

type Locale interface {
	objc.Object
	LocalizedStringForLocaleIdentifier(localeIdentifier string) string
	LocalizedStringForCountryCode(countryCode string) string
	LocalizedStringForLanguageCode(languageCode string) string
	LocalizedStringForScriptCode(scriptCode string) string
	LocalizedStringForVariantCode(variantCode string) string
	LocalizedStringForCollationIdentifier(collationIdentifier string) string
	LocalizedStringForCollatorIdentifier(collatorIdentifier string) string
	LocalizedStringForCurrencyCode(currencyCode string) string
	LocalizedStringForCalendarIdentifier(calendarIdentifier string) string
	ObjectForKey(key LocaleKey) objc.Object
	DisplayNameForKey_Value(key LocaleKey, value objc.Object) string
	LocaleIdentifier() string
	CountryCode() string
	LanguageCode() string
	ScriptCode() string
	VariantCode() string
	ExemplarCharacterSet() CharacterSet
	CollationIdentifier() string
	CollatorIdentifier() string
	UsesMetricSystem() bool
	DecimalSeparator() string
	GroupingSeparator() string
	CurrencyCode() string
	CurrencySymbol() string
	CalendarIdentifier() string
	QuotationBeginDelimiter() string
	QuotationEndDelimiter() string
	AlternateQuotationBeginDelimiter() string
	AlternateQuotationEndDelimiter() string
}

type NSLocale struct {
	objc.NSObject
}

func MakeLocale(ptr unsafe.Pointer) NSLocale {
	return NSLocale{
		NSObject: objc.MakeObject(ptr),
	}
}

func AllocLocale() NSLocale {
	return MakeLocale(C.C_Locale_Alloc())
}

func (n NSLocale) InitWithLocaleIdentifier(_string string) Locale {
	result_ := C.C_NSLocale_InitWithLocaleIdentifier(n.Ptr(), NewString(_string).Ptr())
	return MakeLocale(result_)
}

func (n NSLocale) InitWithCoder(coder Coder) Locale {
	result_ := C.C_NSLocale_InitWithCoder(n.Ptr(), objc.ExtractPtr(coder))
	return MakeLocale(result_)
}

func LocaleWithLocaleIdentifier(ident string) Locale {
	result_ := C.C_NSLocale_LocaleWithLocaleIdentifier(NewString(ident).Ptr())
	return MakeLocale(result_)
}

func Locale_CanonicalLocaleIdentifierFromString(_string string) string {
	result_ := C.C_NSLocale_Locale_CanonicalLocaleIdentifierFromString(NewString(_string).Ptr())
	return MakeString(result_).String()
}

func Locale_ComponentsFromLocaleIdentifier(_string string) map[string]string {
	result_ := C.C_NSLocale_Locale_ComponentsFromLocaleIdentifier(NewString(_string).Ptr())
	if result_.len > 0 {
		defer C.free(result_.key_data)
		defer C.free(result_.value_data)
	}
	result_KeySlice := (*[1 << 28]unsafe.Pointer)(unsafe.Pointer(result_.key_data))[:result_.len:result_.len]
	result_ValueSlice := (*[1 << 28]unsafe.Pointer)(unsafe.Pointer(result_.value_data))[:result_.len:result_.len]
	var goResult_ = make(map[string]string)
	for idx, k := range result_KeySlice {
		v := result_ValueSlice[idx]
		goResult_[MakeString(k).String()] = MakeString(v).String()
	}
	return goResult_
}

func LocaleIdentifierFromComponents(dict map[string]string) string {
	var cDict C.Dictionary
	if len(dict) > 0 {
		cDictKeyData := make([]unsafe.Pointer, len(dict))
		cDictValueData := make([]unsafe.Pointer, len(dict))
		var idx = 0
		for k, v := range dict {
			cDictKeyData[idx] = NewString(k).Ptr()
			cDictValueData[idx] = NewString(v).Ptr()
			idx++
		}
		cDict.key_data = unsafe.Pointer(&cDictKeyData[0])
		cDict.value_data = unsafe.Pointer(&cDictValueData[0])
		cDict.len = C.int(len(dict))
	}
	result_ := C.C_NSLocale_LocaleIdentifierFromComponents(cDict)
	return MakeString(result_).String()
}

func Locale_CanonicalLanguageIdentifierFromString(_string string) string {
	result_ := C.C_NSLocale_Locale_CanonicalLanguageIdentifierFromString(NewString(_string).Ptr())
	return MakeString(result_).String()
}

func LocaleIdentifierFromWindowsLocaleCode(lcid uint32) string {
	result_ := C.C_NSLocale_LocaleIdentifierFromWindowsLocaleCode(C.uint(lcid))
	return MakeString(result_).String()
}

func Locale_WindowsLocaleCodeFromLocaleIdentifier(localeIdentifier string) uint32 {
	result_ := C.C_NSLocale_Locale_WindowsLocaleCodeFromLocaleIdentifier(NewString(localeIdentifier).Ptr())
	return uint32(result_)
}

func (n NSLocale) LocalizedStringForLocaleIdentifier(localeIdentifier string) string {
	result_ := C.C_NSLocale_LocalizedStringForLocaleIdentifier(n.Ptr(), NewString(localeIdentifier).Ptr())
	return MakeString(result_).String()
}

func (n NSLocale) LocalizedStringForCountryCode(countryCode string) string {
	result_ := C.C_NSLocale_LocalizedStringForCountryCode(n.Ptr(), NewString(countryCode).Ptr())
	return MakeString(result_).String()
}

func (n NSLocale) LocalizedStringForLanguageCode(languageCode string) string {
	result_ := C.C_NSLocale_LocalizedStringForLanguageCode(n.Ptr(), NewString(languageCode).Ptr())
	return MakeString(result_).String()
}

func (n NSLocale) LocalizedStringForScriptCode(scriptCode string) string {
	result_ := C.C_NSLocale_LocalizedStringForScriptCode(n.Ptr(), NewString(scriptCode).Ptr())
	return MakeString(result_).String()
}

func (n NSLocale) LocalizedStringForVariantCode(variantCode string) string {
	result_ := C.C_NSLocale_LocalizedStringForVariantCode(n.Ptr(), NewString(variantCode).Ptr())
	return MakeString(result_).String()
}

func (n NSLocale) LocalizedStringForCollationIdentifier(collationIdentifier string) string {
	result_ := C.C_NSLocale_LocalizedStringForCollationIdentifier(n.Ptr(), NewString(collationIdentifier).Ptr())
	return MakeString(result_).String()
}

func (n NSLocale) LocalizedStringForCollatorIdentifier(collatorIdentifier string) string {
	result_ := C.C_NSLocale_LocalizedStringForCollatorIdentifier(n.Ptr(), NewString(collatorIdentifier).Ptr())
	return MakeString(result_).String()
}

func (n NSLocale) LocalizedStringForCurrencyCode(currencyCode string) string {
	result_ := C.C_NSLocale_LocalizedStringForCurrencyCode(n.Ptr(), NewString(currencyCode).Ptr())
	return MakeString(result_).String()
}

func (n NSLocale) LocalizedStringForCalendarIdentifier(calendarIdentifier string) string {
	result_ := C.C_NSLocale_LocalizedStringForCalendarIdentifier(n.Ptr(), NewString(calendarIdentifier).Ptr())
	return MakeString(result_).String()
}

func (n NSLocale) ObjectForKey(key LocaleKey) objc.Object {
	result_ := C.C_NSLocale_ObjectForKey(n.Ptr(), NewString(string(key)).Ptr())
	return objc.MakeObject(result_)
}

func (n NSLocale) DisplayNameForKey_Value(key LocaleKey, value objc.Object) string {
	result_ := C.C_NSLocale_DisplayNameForKey_Value(n.Ptr(), NewString(string(key)).Ptr(), objc.ExtractPtr(value))
	return MakeString(result_).String()
}

func Locale_CharacterDirectionForLanguage(isoLangCode string) LocaleLanguageDirection {
	result_ := C.C_NSLocale_Locale_CharacterDirectionForLanguage(NewString(isoLangCode).Ptr())
	return LocaleLanguageDirection(uint(result_))
}

func Locale_LineDirectionForLanguage(isoLangCode string) LocaleLanguageDirection {
	result_ := C.C_NSLocale_Locale_LineDirectionForLanguage(NewString(isoLangCode).Ptr())
	return LocaleLanguageDirection(uint(result_))
}

func AutoupdatingCurrentLocale() Locale {
	result_ := C.C_NSLocale_AutoupdatingCurrentLocale()
	return MakeLocale(result_)
}

func CurrentLocale() Locale {
	result_ := C.C_NSLocale_CurrentLocale()
	return MakeLocale(result_)
}

func SystemLocale() Locale {
	result_ := C.C_NSLocale_SystemLocale()
	return MakeLocale(result_)
}

func Locale_AvailableLocaleIdentifiers() []string {
	result_ := C.C_NSLocale_Locale_AvailableLocaleIdentifiers()
	if result_.len > 0 {
		defer C.free(result_.data)
	}
	result_Slice := (*[1 << 28]unsafe.Pointer)(unsafe.Pointer(result_.data))[:result_.len:result_.len]
	var goResult_ = make([]string, len(result_Slice))
	for idx, r := range result_Slice {
		goResult_[idx] = MakeString(r).String()
	}
	return goResult_
}

func Locale_ISOCountryCodes() []string {
	result_ := C.C_NSLocale_Locale_ISOCountryCodes()
	if result_.len > 0 {
		defer C.free(result_.data)
	}
	result_Slice := (*[1 << 28]unsafe.Pointer)(unsafe.Pointer(result_.data))[:result_.len:result_.len]
	var goResult_ = make([]string, len(result_Slice))
	for idx, r := range result_Slice {
		goResult_[idx] = MakeString(r).String()
	}
	return goResult_
}

func Locale_ISOLanguageCodes() []string {
	result_ := C.C_NSLocale_Locale_ISOLanguageCodes()
	if result_.len > 0 {
		defer C.free(result_.data)
	}
	result_Slice := (*[1 << 28]unsafe.Pointer)(unsafe.Pointer(result_.data))[:result_.len:result_.len]
	var goResult_ = make([]string, len(result_Slice))
	for idx, r := range result_Slice {
		goResult_[idx] = MakeString(r).String()
	}
	return goResult_
}

func Locale_ISOCurrencyCodes() []string {
	result_ := C.C_NSLocale_Locale_ISOCurrencyCodes()
	if result_.len > 0 {
		defer C.free(result_.data)
	}
	result_Slice := (*[1 << 28]unsafe.Pointer)(unsafe.Pointer(result_.data))[:result_.len:result_.len]
	var goResult_ = make([]string, len(result_Slice))
	for idx, r := range result_Slice {
		goResult_[idx] = MakeString(r).String()
	}
	return goResult_
}

func Locale_CommonISOCurrencyCodes() []string {
	result_ := C.C_NSLocale_Locale_CommonISOCurrencyCodes()
	if result_.len > 0 {
		defer C.free(result_.data)
	}
	result_Slice := (*[1 << 28]unsafe.Pointer)(unsafe.Pointer(result_.data))[:result_.len:result_.len]
	var goResult_ = make([]string, len(result_Slice))
	for idx, r := range result_Slice {
		goResult_[idx] = MakeString(r).String()
	}
	return goResult_
}

func (n NSLocale) LocaleIdentifier() string {
	result_ := C.C_NSLocale_LocaleIdentifier(n.Ptr())
	return MakeString(result_).String()
}

func (n NSLocale) CountryCode() string {
	result_ := C.C_NSLocale_CountryCode(n.Ptr())
	return MakeString(result_).String()
}

func (n NSLocale) LanguageCode() string {
	result_ := C.C_NSLocale_LanguageCode(n.Ptr())
	return MakeString(result_).String()
}

func (n NSLocale) ScriptCode() string {
	result_ := C.C_NSLocale_ScriptCode(n.Ptr())
	return MakeString(result_).String()
}

func (n NSLocale) VariantCode() string {
	result_ := C.C_NSLocale_VariantCode(n.Ptr())
	return MakeString(result_).String()
}

func (n NSLocale) ExemplarCharacterSet() CharacterSet {
	result_ := C.C_NSLocale_ExemplarCharacterSet(n.Ptr())
	return MakeCharacterSet(result_)
}

func (n NSLocale) CollationIdentifier() string {
	result_ := C.C_NSLocale_CollationIdentifier(n.Ptr())
	return MakeString(result_).String()
}

func (n NSLocale) CollatorIdentifier() string {
	result_ := C.C_NSLocale_CollatorIdentifier(n.Ptr())
	return MakeString(result_).String()
}

func (n NSLocale) UsesMetricSystem() bool {
	result_ := C.C_NSLocale_UsesMetricSystem(n.Ptr())
	return bool(result_)
}

func (n NSLocale) DecimalSeparator() string {
	result_ := C.C_NSLocale_DecimalSeparator(n.Ptr())
	return MakeString(result_).String()
}

func (n NSLocale) GroupingSeparator() string {
	result_ := C.C_NSLocale_GroupingSeparator(n.Ptr())
	return MakeString(result_).String()
}

func (n NSLocale) CurrencyCode() string {
	result_ := C.C_NSLocale_CurrencyCode(n.Ptr())
	return MakeString(result_).String()
}

func (n NSLocale) CurrencySymbol() string {
	result_ := C.C_NSLocale_CurrencySymbol(n.Ptr())
	return MakeString(result_).String()
}

func (n NSLocale) CalendarIdentifier() string {
	result_ := C.C_NSLocale_CalendarIdentifier(n.Ptr())
	return MakeString(result_).String()
}

func (n NSLocale) QuotationBeginDelimiter() string {
	result_ := C.C_NSLocale_QuotationBeginDelimiter(n.Ptr())
	return MakeString(result_).String()
}

func (n NSLocale) QuotationEndDelimiter() string {
	result_ := C.C_NSLocale_QuotationEndDelimiter(n.Ptr())
	return MakeString(result_).String()
}

func (n NSLocale) AlternateQuotationBeginDelimiter() string {
	result_ := C.C_NSLocale_AlternateQuotationBeginDelimiter(n.Ptr())
	return MakeString(result_).String()
}

func (n NSLocale) AlternateQuotationEndDelimiter() string {
	result_ := C.C_NSLocale_AlternateQuotationEndDelimiter(n.Ptr())
	return MakeString(result_).String()
}

func Locale_PreferredLanguages() []string {
	result_ := C.C_NSLocale_Locale_PreferredLanguages()
	if result_.len > 0 {
		defer C.free(result_.data)
	}
	result_Slice := (*[1 << 28]unsafe.Pointer)(unsafe.Pointer(result_.data))[:result_.len:result_.len]
	var goResult_ = make([]string, len(result_Slice))
	for idx, r := range result_Slice {
		goResult_[idx] = MakeString(r).String()
	}
	return goResult_
}
