package appkit

type FontTextStyle string

var FontTextStyleLargeTitle = FontTextStyle("UICTFontTextStyleTitle0")
var FontTextStyleTitle1 = FontTextStyle("UICTFontTextStyleTitle1")
var FontTextStyleTitle2 = FontTextStyle("UICTFontTextStyleTitle2")
var FontTextStyleTitle3 = FontTextStyle("UICTFontTextStyleTitle3")
var FontTextStyleHeadline = FontTextStyle("UICTFontTextStyleHeadline")
var FontTextStyleSubheadline = FontTextStyle("UICTFontTextStyleSubhead")
var FontTextStyleBody = FontTextStyle("UICTFontTextStyleBody")
var FontTextStyleCallout = FontTextStyle("UICTFontTextStyleCallout")
var FontTextStyleFootnote = FontTextStyle("UICTFontTextStyleFootnote")
var FontTextStyleCaption1 = FontTextStyle("UICTFontTextStyleCaption1")
var FontTextStyleCaption2 = FontTextStyle("UICTFontTextStyleCaption2")

type FontWeight float64

const (
	FontWeightUltraLight = FontWeight(-0.8)
	FontWeightThin       = FontWeight(-0.6)
	FontWeightLight      = FontWeight(-0.4)
	FontWeightRegular    = FontWeight(0)
	FontWeightMedium     = FontWeight(0.23)
	FontWeightSemibold   = FontWeight(0.3)
	FontWeightBold       = FontWeight(0.4)
	FontWeightHeavy      = FontWeight(0.56)
	FontWeightBlack      = FontWeight(0.62)
)
