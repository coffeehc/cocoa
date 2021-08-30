package foundation

type CalendarOptions uint

const (
	CalendarWrapComponents                          CalendarOptions = 1 << 0
	CalendarMatchStrictly                           CalendarOptions = 1 << 1
	CalendarSearchBackwards                         CalendarOptions = 1 << 2
	CalendarMatchPreviousTimePreservingSmallerUnits CalendarOptions = 1 << 8
	CalendarMatchNextTimePreservingSmallerUnits     CalendarOptions = 1 << 9
	CalendarMatchNextTime                           CalendarOptions = 1 << 10
	CalendarMatchFirst                              CalendarOptions = 1 << 12
	CalendarMatchLast                               CalendarOptions = 1 << 13
)

type CalendarUnit uint

const (
	CalendarUnitEra               CalendarUnit = 2
	CalendarUnitYear              CalendarUnit = 4
	CalendarUnitMonth             CalendarUnit = 8
	CalendarUnitDay               CalendarUnit = 16
	CalendarUnitHour              CalendarUnit = 32
	CalendarUnitMinute            CalendarUnit = 64
	CalendarUnitSecond            CalendarUnit = 128
	CalendarUnitWeekday           CalendarUnit = 512
	CalendarUnitWeekdayOrdinal    CalendarUnit = 1024
	CalendarUnitQuarter           CalendarUnit = 2048
	CalendarUnitWeekOfMonth       CalendarUnit = 4096
	CalendarUnitWeekOfYear        CalendarUnit = 8192
	CalendarUnitYearForWeekOfYear CalendarUnit = 16384
	CalendarUnitNanosecond        CalendarUnit = 32768
	CalendarUnitCalendar          CalendarUnit = 1048576
	CalendarUnitTimeZone          CalendarUnit = 2097152
)

type TimeZoneNameStyle int

const (
	TimeZoneNameStyleStandard            TimeZoneNameStyle = 0
	TimeZoneNameStyleShortStandard       TimeZoneNameStyle = 1
	TimeZoneNameStyleDaylightSaving      TimeZoneNameStyle = 2
	TimeZoneNameStyleShortDaylightSaving TimeZoneNameStyle = 3
	TimeZoneNameStyleGeneric             TimeZoneNameStyle = 4
	TimeZoneNameStyleShortGeneric        TimeZoneNameStyle = 5
)

type CalendarIdentifier string

const (
	CalendarIdentifierGregorian           CalendarIdentifier = "gregorian"
	CalendarIdentifierBuddhist            CalendarIdentifier = "buddhist"
	CalendarIdentifierChinese             CalendarIdentifier = "chinese"
	CalendarIdentifierCoptic              CalendarIdentifier = "coptic"
	CalendarIdentifierEthiopicAmeteMihret CalendarIdentifier = "ethiopic"
	CalendarIdentifierEthiopicAmeteAlem   CalendarIdentifier = "ethiopic-amete-alem"
	CalendarIdentifierHebrew              CalendarIdentifier = "hebrew"
	CalendarIdentifierISO8601             CalendarIdentifier = "iso8601"
	CalendarIdentifierIndian              CalendarIdentifier = "indian"
	CalendarIdentifierIslamic             CalendarIdentifier = "islamic"
	CalendarIdentifierIslamicCivil        CalendarIdentifier = "islamic-civil"
	CalendarIdentifierJapanese            CalendarIdentifier = "japanese"
	CalendarIdentifierPersian             CalendarIdentifier = "persian"
	CalendarIdentifierRepublicOfChina     CalendarIdentifier = "roc"
	CalendarIdentifierIslamicTabular      CalendarIdentifier = "islamic-tbla"
	CalendarIdentifierIslamicUmmAlQura    CalendarIdentifier = "islamic-umalqura"
)

type StringEncoding uint

const (
	ASCIIStringEncoding         StringEncoding = 1
	NEXTSTEPStringEncoding      StringEncoding = 2
	JapaneseEUCStringEncoding   StringEncoding = 3
	UTF8StringEncoding          StringEncoding = 4
	ISOLatin1StringEncoding     StringEncoding = 5
	SymbolStringEncoding        StringEncoding = 6
	NonLossyASCIIStringEncoding StringEncoding = 7
	ShiftJISStringEncoding      StringEncoding = 8
	ISOLatin2StringEncoding     StringEncoding = 9
	UnicodeStringEncoding       StringEncoding = 10
	WindowsCP1251StringEncoding StringEncoding = 11
	WindowsCP1252StringEncoding StringEncoding = 12
	WindowsCP1253StringEncoding StringEncoding = 13
	WindowsCP1254StringEncoding StringEncoding = 14
	WindowsCP1250StringEncoding StringEncoding = 15
	ISO2022JPStringEncoding     StringEncoding = 21
	MacOSRomanStringEncoding    StringEncoding = 30

	UTF16StringEncoding StringEncoding = UnicodeStringEncoding

	UTF16BigEndianStringEncoding    StringEncoding = 0x90000100
	UTF16LittleEndianStringEncoding StringEncoding = 0x94000100

	UTF32StringEncoding             StringEncoding = 0x8c000100
	UTF32BigEndianStringEncoding    StringEncoding = 0x98000100
	UTF32LittleEndianStringEncoding StringEncoding = 0x9c000100
)

type LocaleKey string

const (
	LocaleIdentifier                          LocaleKey = "kCFLocaleIdentifierKey"
	LocaleLanguageCode                        LocaleKey = "kCFLocaleLanguageCodeKey"
	LocaleCountryCode                         LocaleKey = "kCFLocaleCountryCodeKey"
	LocaleScriptCode                          LocaleKey = "kCFLocaleScriptCodeKey"
	LocaleVariantCode                         LocaleKey = "kCFLocaleVariantCodeKey"
	LocaleExemplarCharacterSet                LocaleKey = "kCFLocaleExemplarCharacterSetKey"
	LocaleCalendar                            LocaleKey = "kCFLocaleCalendarKey"
	LocaleCollationIdentifier                 LocaleKey = "collation"
	LocaleUsesMetricSystem                    LocaleKey = "kCFLocaleUsesMetricSystemKey"
	LocaleMeasurementSystem                   LocaleKey = "kCFLocaleMeasurementSystemKey"
	LocaleDecimalSeparator                    LocaleKey = "kCFLocaleDecimalSeparatorKey"
	LocaleGroupingSeparator                   LocaleKey = "kCFLocaleGroupingSeparatorKey"
	LocaleCurrencySymbol                      LocaleKey = "kCFLocaleCurrencySymbolKey"
	LocaleCurrencyCode                        LocaleKey = "currency"
	LocaleCollatorIdentifier                  LocaleKey = "kCFLocaleCollatorIdentifierKey"
	LocaleQuotationBeginDelimiterKey          LocaleKey = "kCFLocaleQuotationBeginDelimiterKey"
	LocaleQuotationEndDelimiterKey            LocaleKey = "kCFLocaleQuotationEndDelimiterKey"
	LocaleAlternateQuotationBeginDelimiterKey LocaleKey = "kCFLocaleAlternateQuotationBeginDelimiterKey"
	LocaleAlternateQuotationEndDelimiterKey   LocaleKey = "kCFLocaleAlternateQuotationEndDelimiterKey"
)

type LocaleLanguageDirection uint

const (
	LocaleLanguageDirectionUnknown     LocaleLanguageDirection = 0
	LocaleLanguageDirectionLeftToRight LocaleLanguageDirection = 1
	LocaleLanguageDirectionRightToLeft LocaleLanguageDirection = 2
	LocaleLanguageDirectionTopToBottom LocaleLanguageDirection = 3
	LocaleLanguageDirectionBottomToTop LocaleLanguageDirection = 4
)

type AttributedStringKey string

type KeyValueChange uint

const (
	KeyValueChangeSetting     KeyValueChange = 1
	KeyValueChangeInsertion   KeyValueChange = 2
	KeyValueChangeRemoval     KeyValueChange = 3
	KeyValueChangeReplacement KeyValueChange = 4
)

type KeyValueSetMutationKind uint

const (
	KeyValueUnionSetMutation     KeyValueSetMutationKind = 1
	KeyValueMinusSetMutation     KeyValueSetMutationKind = 2
	KeyValueIntersectSetMutation KeyValueSetMutationKind = 3
	KeyValueSetSetMutation       KeyValueSetMutationKind = 4
)

type MatchingOptions uint

const (
	MatchingReportProgress         MatchingOptions = 1 << 0
	MatchingReportCompletion       MatchingOptions = 1 << 1
	MatchingAnchored               MatchingOptions = 1 << 2
	MatchingWithTransparentBounds  MatchingOptions = 1 << 3
	MatchingWithoutAnchoringBounds MatchingOptions = 1 << 4
)

type MatchingFlags uint

const (
	MatchingProgress      MatchingFlags = 1 << 0
	MatchingCompleted     MatchingFlags = 1 << 1
	MatchingHitEnd        MatchingFlags = 1 << 2
	MatchingRequiredEnd   MatchingFlags = 1 << 3
	MatchingInternalError MatchingFlags = 1 << 4
)

type RegularExpressionOptions uint

const (
	RegularExpressionCaseInsensitive            RegularExpressionOptions = 1 << 0
	RegularExpressionAllowCommentsAndWhitespace RegularExpressionOptions = 1 << 1
	RegularExpressionIgnoreMetacharacters       RegularExpressionOptions = 1 << 2
	RegularExpressionDotMatchesLineSeparators   RegularExpressionOptions = 1 << 3
	RegularExpressionAnchorsMatchLines          RegularExpressionOptions = 1 << 4
	RegularExpressionUseUnixLineSeparators      RegularExpressionOptions = 1 << 5
	RegularExpressionUseUnicodeWordBoundaries   RegularExpressionOptions = 1 << 6
)

type DecodingFailurePolicy int

const (
	DecodingFailurePolicyRaiseException    DecodingFailurePolicy = 0
	DecodingFailurePolicySetErrorAndReturn DecodingFailurePolicy = 1
)

type QualityOfService int

const (
	QualityOfServiceUserInteractive QualityOfService = 0x21
	QualityOfServiceUserInitiated   QualityOfService = 0x19
	QualityOfServiceUtility         QualityOfService = 0x11
	QualityOfServiceBackground      QualityOfService = 0x09
	QualityOfServiceDefault         QualityOfService = -1
)

type OperationQueuePriority int

const (
	OperationQueuePriorityVeryLow  OperationQueuePriority = -8
	OperationQueuePriorityLow      OperationQueuePriority = -4
	OperationQueuePriorityNormal   OperationQueuePriority = 0
	OperationQueuePriorityHigh     OperationQueuePriority = 4
	OperationQueuePriorityVeryHigh OperationQueuePriority = 8
)

type RunLoopMode string

type URLRequestNetworkServiceType uint

type NumberFormatterBehavior uint
type NumberFormatterStyle uint
type NumberFormatterRoundingMode uint
type FormattingContext int
type DateComponentsFormatterUnitsStyle int
type DateComponentsFormatterZeroFormattingBehavior uint
type DateFormatterStyle int
type DateFormatterBehavior uint
type DateIntervalFormatterStyle uint
type MeasurementFormatterUnitOptions uint
type FormattingUnitStyle int
type ByteCountFormatterCountStyle int
type ByteCountFormatterUnits uint
type PersonNameComponentsFormatterStyle int
type PersonNameComponentsFormatterOptions uint
type NumberFormatterPadPosition uint
type ExceptionName string
type ErrorDomain string
type NotificationName string

type ComparisonResult int

const (
	OrderedAscending  ComparisonResult = -1
	OrderedSame       ComparisonResult = 0
	OrderedDescending ComparisonResult = 1
)

type URLRequestCachePolicy uint

const (
	URLRequestUseProtocolCachePolicy                URLRequestCachePolicy = 0
	URLRequestReloadIgnoringLocalCacheData          URLRequestCachePolicy = 1
	URLRequestReloadIgnoringLocalAndRemoteCacheData URLRequestCachePolicy = 4
	URLRequestReloadIgnoringCacheData               URLRequestCachePolicy = URLRequestReloadIgnoringLocalCacheData
	URLRequestReturnCacheDataElseLoad               URLRequestCachePolicy = 2
	URLRequestReturnCacheDataDontLoad               URLRequestCachePolicy = 3
	URLRequestReloadRevalidatingCacheData           URLRequestCachePolicy = 5
)

type URLResourceKey string

type URLBookmarkFileCreationOptions uint

type TimeInterval float64

type ProgressUserInfoKey string

const (
	ProgressEstimatedTimeRemainingKey         ProgressUserInfoKey = "NSProgressEstimatedTimeRemainingKey"
	ProgressThroughputKey                     ProgressUserInfoKey = "NSProgressThroughputKey"
	ProgressFileAnimationImageKey             ProgressUserInfoKey = "NSProgressFlyToImageKey"
	ProgressFileAnimationImageOriginalRectKey ProgressUserInfoKey = "NSProgressFileAnimationImageOriginalRectKey"
	ProgressFileCompletedCountKey             ProgressUserInfoKey = "NSProgressFileCompletedCountKey"
	ProgressFileIconKey                       ProgressUserInfoKey = "NSProgressFileIconKey"
	ProgressFileOperationKindKey              ProgressUserInfoKey = "NSProgressFileOperationKindKey"
	ProgressFileTotalCountKey                 ProgressUserInfoKey = "NSProgressFileTotalCountKey"
	ProgressFileURLKey                        ProgressUserInfoKey = "NSProgressFileURLKey"
)

type ProgressKind string

const (
	ProgressKindFile ProgressKind = "NSProgressKindFile"
)

type ProgressFileOperationKind string

const (
	ProgressFileOperationKindCopying                       ProgressFileOperationKind = "NSProgressFileOperationKindCopying"
	ProgressFileOperationKindDecompressingAfterDownloading ProgressFileOperationKind = "NSProgressFileOperationKindDecompressingAfterDownloading"
	ProgressFileOperationKindDownloading                   ProgressFileOperationKind = "NSProgressFileOperationKindDownloading"
	ProgressFileOperationKindReceiving                     ProgressFileOperationKind = "NSProgressFileOperationKindReceiving"
	ProgressFileOperationKindUploading                     ProgressFileOperationKind = "NSProgressFileOperationKindUploading"
)

type SaveOptions uint

const (
	SaveOptionsYes SaveOptions = 0
	SaveOptionsNo  SaveOptions = 1
	SaveOptionsAsk SaveOptions = 2
)

type ErrorUserInfoKey string

const (
	URLErrorKey                         ErrorUserInfoKey = "NSURL"
	FilePathErrorKey                    ErrorUserInfoKey = "NSFilePath"
	HelpAnchorErrorKey                  ErrorUserInfoKey = "NSHelpAnchor"
	LocalizedDescriptionKey             ErrorUserInfoKey = "NSLocalizedDescription"
	LocalizedFailureErrorKey            ErrorUserInfoKey = "NSLocalizedFailure"
	LocalizedFailureReasonErrorKey      ErrorUserInfoKey = "NSLocalizedFailureReason"
	LocalizedRecoveryOptionsErrorKey    ErrorUserInfoKey = "NSLocalizedRecoveryOptions"
	LocalizedRecoverySuggestionErrorKey ErrorUserInfoKey = "NSLocalizedRecoverySuggestion"
	RecoveryAttempterErrorKey           ErrorUserInfoKey = "NSRecoveryAttempter"
	StringEncodingErrorKey              ErrorUserInfoKey = "NSStringEncoding"
	UnderlyingErrorKey                  ErrorUserInfoKey = "NSUnderlyingError"
	DebugDescriptionErrorKey            ErrorUserInfoKey = "NSDebugDescription"
	MultipleUnderlyingErrorsKey         ErrorUserInfoKey = "NSMultipleUnderlyingErrorsKey"
)

type TextCheckingKey string

const (
	TextCheckingAirlineKey      TextCheckingKey = "Airline"
	TextCheckingCityKey         TextCheckingKey = "City"
	TextCheckingCountryKey      TextCheckingKey = "Country"
	TextCheckingFlightKey       TextCheckingKey = "Flight"
	TextCheckingJobTitleKey     TextCheckingKey = "JobTitle"
	TextCheckingNameKey         TextCheckingKey = "Name"
	TextCheckingOrganizationKey TextCheckingKey = "Organization"
	TextCheckingPhoneKey        TextCheckingKey = "Phone"
	TextCheckingStateKey        TextCheckingKey = "State"
	TextCheckingStreetKey       TextCheckingKey = "Street"
	TextCheckingZIPKey          TextCheckingKey = "ZIP"
)
