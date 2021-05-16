package foundation

// #cgo CFLAGS: -x objective-c
// #cgo LDFLAGS: -framework Foundation
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

func MakeLocale(ptr unsafe.Pointer) *NSLocale {
	if ptr == nil {
		return nil
	}
	return &NSLocale{
		NSObject: *objc.MakeObject(ptr),
	}
}

func AllocLocale() *NSLocale {
	return MakeLocale(C.C_Locale_Alloc())
}

func (n *NSLocale) InitWithLocaleIdentifier(_string string) Locale {
	result := C.C_NSLocale_InitWithLocaleIdentifier(n.Ptr(), NewString(_string).Ptr())
	return MakeLocale(result)
}

func (n *NSLocale) InitWithCoder(coder Coder) Locale {
	result := C.C_NSLocale_InitWithCoder(n.Ptr(), objc.ExtractPtr(coder))
	return MakeLocale(result)
}

func LocaleWithLocaleIdentifier(ident string) Locale {
	result := C.C_NSLocale_LocaleWithLocaleIdentifier(NewString(ident).Ptr())
	return MakeLocale(result)
}

func Locale_CanonicalLocaleIdentifierFromString(_string string) string {
	result := C.C_NSLocale_Locale_CanonicalLocaleIdentifierFromString(NewString(_string).Ptr())
	return MakeString(result).String()
}

func Locale_CanonicalLanguageIdentifierFromString(_string string) string {
	result := C.C_NSLocale_Locale_CanonicalLanguageIdentifierFromString(NewString(_string).Ptr())
	return MakeString(result).String()
}

func LocaleIdentifierFromWindowsLocaleCode(lcid uint32) string {
	result := C.C_NSLocale_LocaleIdentifierFromWindowsLocaleCode(C.uint(lcid))
	return MakeString(result).String()
}

func Locale_WindowsLocaleCodeFromLocaleIdentifier(localeIdentifier string) uint32 {
	result := C.C_NSLocale_Locale_WindowsLocaleCodeFromLocaleIdentifier(NewString(localeIdentifier).Ptr())
	return uint32(result)
}

func (n *NSLocale) LocalizedStringForLocaleIdentifier(localeIdentifier string) string {
	result := C.C_NSLocale_LocalizedStringForLocaleIdentifier(n.Ptr(), NewString(localeIdentifier).Ptr())
	return MakeString(result).String()
}

func (n *NSLocale) LocalizedStringForCountryCode(countryCode string) string {
	result := C.C_NSLocale_LocalizedStringForCountryCode(n.Ptr(), NewString(countryCode).Ptr())
	return MakeString(result).String()
}

func (n *NSLocale) LocalizedStringForLanguageCode(languageCode string) string {
	result := C.C_NSLocale_LocalizedStringForLanguageCode(n.Ptr(), NewString(languageCode).Ptr())
	return MakeString(result).String()
}

func (n *NSLocale) LocalizedStringForScriptCode(scriptCode string) string {
	result := C.C_NSLocale_LocalizedStringForScriptCode(n.Ptr(), NewString(scriptCode).Ptr())
	return MakeString(result).String()
}

func (n *NSLocale) LocalizedStringForVariantCode(variantCode string) string {
	result := C.C_NSLocale_LocalizedStringForVariantCode(n.Ptr(), NewString(variantCode).Ptr())
	return MakeString(result).String()
}

func (n *NSLocale) LocalizedStringForCollationIdentifier(collationIdentifier string) string {
	result := C.C_NSLocale_LocalizedStringForCollationIdentifier(n.Ptr(), NewString(collationIdentifier).Ptr())
	return MakeString(result).String()
}

func (n *NSLocale) LocalizedStringForCollatorIdentifier(collatorIdentifier string) string {
	result := C.C_NSLocale_LocalizedStringForCollatorIdentifier(n.Ptr(), NewString(collatorIdentifier).Ptr())
	return MakeString(result).String()
}

func (n *NSLocale) LocalizedStringForCurrencyCode(currencyCode string) string {
	result := C.C_NSLocale_LocalizedStringForCurrencyCode(n.Ptr(), NewString(currencyCode).Ptr())
	return MakeString(result).String()
}

func (n *NSLocale) LocalizedStringForCalendarIdentifier(calendarIdentifier string) string {
	result := C.C_NSLocale_LocalizedStringForCalendarIdentifier(n.Ptr(), NewString(calendarIdentifier).Ptr())
	return MakeString(result).String()
}

func (n *NSLocale) ObjectForKey(key LocaleKey) objc.Object {
	result := C.C_NSLocale_ObjectForKey(n.Ptr(), NewString(string(key)).Ptr())
	return objc.MakeObject(result)
}

func (n *NSLocale) DisplayNameForKey_Value(key LocaleKey, value objc.Object) string {
	result := C.C_NSLocale_DisplayNameForKey_Value(n.Ptr(), NewString(string(key)).Ptr(), objc.ExtractPtr(value))
	return MakeString(result).String()
}

func Locale_CharacterDirectionForLanguage(isoLangCode string) LocaleLanguageDirection {
	result := C.C_NSLocale_Locale_CharacterDirectionForLanguage(NewString(isoLangCode).Ptr())
	return LocaleLanguageDirection(uint(result))
}

func Locale_LineDirectionForLanguage(isoLangCode string) LocaleLanguageDirection {
	result := C.C_NSLocale_Locale_LineDirectionForLanguage(NewString(isoLangCode).Ptr())
	return LocaleLanguageDirection(uint(result))
}

func AutoupdatingCurrentLocale() Locale {
	result := C.C_NSLocale_AutoupdatingCurrentLocale()
	return MakeLocale(result)
}

func CurrentLocale() Locale {
	result := C.C_NSLocale_CurrentLocale()
	return MakeLocale(result)
}

func SystemLocale() Locale {
	result := C.C_NSLocale_SystemLocale()
	return MakeLocale(result)
}

func Locale_AvailableLocaleIdentifiers() []string {
	result := C.C_NSLocale_Locale_AvailableLocaleIdentifiers()
	defer C.free(result.data)
	resultSlice := (*[1 << 28]unsafe.Pointer)(unsafe.Pointer(result.data))[:result.len:result.len]
	var goResult = make([]string, len(resultSlice))
	for idx, r := range resultSlice {
		goResult[idx] = MakeString(r).String()
	}
	return goResult
}

func Locale_ISOCountryCodes() []string {
	result := C.C_NSLocale_Locale_ISOCountryCodes()
	defer C.free(result.data)
	resultSlice := (*[1 << 28]unsafe.Pointer)(unsafe.Pointer(result.data))[:result.len:result.len]
	var goResult = make([]string, len(resultSlice))
	for idx, r := range resultSlice {
		goResult[idx] = MakeString(r).String()
	}
	return goResult
}

func Locale_ISOLanguageCodes() []string {
	result := C.C_NSLocale_Locale_ISOLanguageCodes()
	defer C.free(result.data)
	resultSlice := (*[1 << 28]unsafe.Pointer)(unsafe.Pointer(result.data))[:result.len:result.len]
	var goResult = make([]string, len(resultSlice))
	for idx, r := range resultSlice {
		goResult[idx] = MakeString(r).String()
	}
	return goResult
}

func Locale_ISOCurrencyCodes() []string {
	result := C.C_NSLocale_Locale_ISOCurrencyCodes()
	defer C.free(result.data)
	resultSlice := (*[1 << 28]unsafe.Pointer)(unsafe.Pointer(result.data))[:result.len:result.len]
	var goResult = make([]string, len(resultSlice))
	for idx, r := range resultSlice {
		goResult[idx] = MakeString(r).String()
	}
	return goResult
}

func Locale_CommonISOCurrencyCodes() []string {
	result := C.C_NSLocale_Locale_CommonISOCurrencyCodes()
	defer C.free(result.data)
	resultSlice := (*[1 << 28]unsafe.Pointer)(unsafe.Pointer(result.data))[:result.len:result.len]
	var goResult = make([]string, len(resultSlice))
	for idx, r := range resultSlice {
		goResult[idx] = MakeString(r).String()
	}
	return goResult
}

func (n *NSLocale) LocaleIdentifier() string {
	result := C.C_NSLocale_LocaleIdentifier(n.Ptr())
	return MakeString(result).String()
}

func (n *NSLocale) CountryCode() string {
	result := C.C_NSLocale_CountryCode(n.Ptr())
	return MakeString(result).String()
}

func (n *NSLocale) LanguageCode() string {
	result := C.C_NSLocale_LanguageCode(n.Ptr())
	return MakeString(result).String()
}

func (n *NSLocale) ScriptCode() string {
	result := C.C_NSLocale_ScriptCode(n.Ptr())
	return MakeString(result).String()
}

func (n *NSLocale) VariantCode() string {
	result := C.C_NSLocale_VariantCode(n.Ptr())
	return MakeString(result).String()
}

func (n *NSLocale) ExemplarCharacterSet() CharacterSet {
	result := C.C_NSLocale_ExemplarCharacterSet(n.Ptr())
	return MakeCharacterSet(result)
}

func (n *NSLocale) CollationIdentifier() string {
	result := C.C_NSLocale_CollationIdentifier(n.Ptr())
	return MakeString(result).String()
}

func (n *NSLocale) CollatorIdentifier() string {
	result := C.C_NSLocale_CollatorIdentifier(n.Ptr())
	return MakeString(result).String()
}

func (n *NSLocale) UsesMetricSystem() bool {
	result := C.C_NSLocale_UsesMetricSystem(n.Ptr())
	return bool(result)
}

func (n *NSLocale) DecimalSeparator() string {
	result := C.C_NSLocale_DecimalSeparator(n.Ptr())
	return MakeString(result).String()
}

func (n *NSLocale) GroupingSeparator() string {
	result := C.C_NSLocale_GroupingSeparator(n.Ptr())
	return MakeString(result).String()
}

func (n *NSLocale) CurrencyCode() string {
	result := C.C_NSLocale_CurrencyCode(n.Ptr())
	return MakeString(result).String()
}

func (n *NSLocale) CurrencySymbol() string {
	result := C.C_NSLocale_CurrencySymbol(n.Ptr())
	return MakeString(result).String()
}

func (n *NSLocale) CalendarIdentifier() string {
	result := C.C_NSLocale_CalendarIdentifier(n.Ptr())
	return MakeString(result).String()
}

func (n *NSLocale) QuotationBeginDelimiter() string {
	result := C.C_NSLocale_QuotationBeginDelimiter(n.Ptr())
	return MakeString(result).String()
}

func (n *NSLocale) QuotationEndDelimiter() string {
	result := C.C_NSLocale_QuotationEndDelimiter(n.Ptr())
	return MakeString(result).String()
}

func (n *NSLocale) AlternateQuotationBeginDelimiter() string {
	result := C.C_NSLocale_AlternateQuotationBeginDelimiter(n.Ptr())
	return MakeString(result).String()
}

func (n *NSLocale) AlternateQuotationEndDelimiter() string {
	result := C.C_NSLocale_AlternateQuotationEndDelimiter(n.Ptr())
	return MakeString(result).String()
}

func Locale_PreferredLanguages() []string {
	result := C.C_NSLocale_Locale_PreferredLanguages()
	defer C.free(result.data)
	resultSlice := (*[1 << 28]unsafe.Pointer)(unsafe.Pointer(result.data))[:result.len:result.len]
	var goResult = make([]string, len(resultSlice))
	for idx, r := range resultSlice {
		goResult[idx] = MakeString(r).String()
	}
	return goResult
}
