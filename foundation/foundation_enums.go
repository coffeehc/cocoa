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

type TimeZoneNameStyle int

type CalendarIdentifier string

type StringEncoding uint

type LocaleKey string

type LocaleLanguageDirection uint

type AttributedStringKey string

type TextBlockLayer int

type KeyValueChange uint
type KeyValueSetMutationKind uint

type MatchingOptions uint

type RegularExpressionOptions uint

type DecodingFailurePolicy int

type QualityOfService int

type OperationQueuePriority int

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
