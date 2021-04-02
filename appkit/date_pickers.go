package appkit

type DatePickerStyle uint

const (
	DatePickerStyleTextFieldAndStepper DatePickerStyle = 0
	DatePickerStyleClockAndCalendar    DatePickerStyle = 1
	DatePickerStyleTextField           DatePickerStyle = 2
)

type DatePickerElementFlags uint

const (
	DatePickerElementFlagHourMinute       DatePickerElementFlags = 0x000c
	DatePickerElementFlagHourMinuteSecond DatePickerElementFlags = 0x000e
	DatePickerElementFlagTimeZone         DatePickerElementFlags = 0x0010
	DatePickerElementFlagYearMonth        DatePickerElementFlags = 0x00c0
	DatePickerElementFlagYearMonthDay     DatePickerElementFlags = 0x00e0
	DatePickerElementFlagEra              DatePickerElementFlags = 0x0100
)

type DatePickerMode uint

const (
	DatePickerModeSingle DatePickerMode = 0
	DatePickerModeRange  DatePickerMode = 1
)
