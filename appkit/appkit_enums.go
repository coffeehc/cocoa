package appkit

import (
	"math"
)

type ApplicationPresentationOptions uint

const (
	ApplicationPresentationDefault                         ApplicationPresentationOptions = 0
	ApplicationPresentationAutoHideDock                    ApplicationPresentationOptions = 1 << 0  // Dock appears when moused to
	ApplicationPresentationHideDock                        ApplicationPresentationOptions = 1 << 1  // Dock is entirely unavailable
	ApplicationPresentationAutoHideMenuBar                 ApplicationPresentationOptions = 1 << 2  // Menu Bar appears when moused to
	ApplicationPresentationHideMenuBar                     ApplicationPresentationOptions = 1 << 3  // Menu Bar is entirely unavailable
	ApplicationPresentationDisableAppleMenu                ApplicationPresentationOptions = 1 << 4  // all Apple menu items are disabled
	ApplicationPresentationDisableProcessSwitching         ApplicationPresentationOptions = 1 << 5  // Cmd+Tab UI is disabled
	ApplicationPresentationDisableForceQuit                ApplicationPresentationOptions = 1 << 6  // Cmd+Opt+Esc panel is disabled
	ApplicationPresentationDisableSessionTermination       ApplicationPresentationOptions = 1 << 7  // PowerKey panel and Restart/Shut Down/Log Out disabled
	ApplicationPresentationDisableHideApplication          ApplicationPresentationOptions = 1 << 8  // Application "Hide" menu item is disabled
	ApplicationPresentationDisableMenuBarTransparency      ApplicationPresentationOptions = 1 << 9  // Menu Bar's transparent appearance is disabled
	ApplicationPresentationFullScreen                      ApplicationPresentationOptions = 1 << 10 // Application is in fullscreen mode
	ApplicationPresentationAutoHideToolbar                 ApplicationPresentationOptions = 1 << 11 // Fullscreen window toolbar is detached from window and hides/shows with autoHidden menuBar.  May be used only when both ApplicationPresentationFullScreen and ApplicationPresentationAutoHideMenuBar are also set
	ApplicationPresentationDisableCursorLocationAssistance ApplicationPresentationOptions = 1 << 12 // "Shake mouse pointer to locate" is disabled
)

type EventGestureAxis int

const (
	EventGestureAxisNone EventGestureAxis = iota
	EventGestureAxisHorizontal
	EventGestureAxisVertical
)

type RemoteNotificationType uint

const (
	RemoteNotificationTypeNone  RemoteNotificationType = 0
	RemoteNotificationTypeBadge RemoteNotificationType = 1 << 0
	RemoteNotificationTypeSound RemoteNotificationType = 1 << 1
	RemoteNotificationTypeAlert RemoteNotificationType = 1 << 2
)

type RequestUserAttentionType uint

const (
	CriticalRequest      RequestUserAttentionType = 0
	InformationalRequest RequestUserAttentionType = 10
)

type ApplicationDelegateReply uint

const (
	ApplicationDelegateReplySuccess ApplicationDelegateReply = 0
	ApplicationDelegateReplyCancel  ApplicationDelegateReply = 1
	ApplicationDelegateReplyFailure ApplicationDelegateReply = 2
)

type ApplicationTerminateReply uint

const (
	TerminateCancel ApplicationTerminateReply = 0
	TerminateNow    ApplicationTerminateReply = 1
	TerminateLater  ApplicationTerminateReply = 2
)

type ApplicationOcclusionState uint

const (
	ApplicationOcclusionStateVisible ApplicationOcclusionState = 1 << 1
)

type AppearanceName string

const (
	AppearanceNameAqua                                  AppearanceName = "NSAppearanceNameAqua"
	AppearanceNameDarkAqua                              AppearanceName = "NSAppearanceNameDarkAqua"
	AppearanceNameVibrantDark                           AppearanceName = "NSAppearanceNameVibrantDark"
	AppearanceNameVibrantLight                          AppearanceName = "NSAppearanceNameVibrantLight"
	AppearanceNameAccessibilityHighContrastAqua         AppearanceName = "NSAppearanceNameAccessibilityAqua"
	AppearanceNameAccessibilityHighContrastDarkAqua     AppearanceName = "NSAppearanceNameAccessibilityDarkAqua"
	AppearanceNameAccessibilityHighContrastVibrantLight AppearanceName = "NSAppearanceNameAccessibilityVibrantLight"
	AppearanceNameAccessibilityHighContrastVibrantDark  AppearanceName = "NSAppearanceNameAccessibilityVibrantDark"
)

type TextAlignment int

const (
	TextAlignmentRight     TextAlignment = 1
	TextAlignmentCenter    TextAlignment = 2
	TextAlignmentJustified TextAlignment = 3
	TextAlignmentNatural   TextAlignment = 4
)

type LineBreakMode int

const (
	LineBreakByWordWrapping     LineBreakMode = iota // Wrap at word boundaries, default
	LineBreakByCharWrapping                          // Wrap at character boundaries
	LineBreakByClipping                              // Simply clip
	LineBreakByTruncatingHead                        // Truncate at head of line: "...wxyz"
	LineBreakByTruncatingTail                        // Truncate at tail of line: "abcd..."
	LineBreakByTruncatingMiddle                      // Truncate middle of line:  "ab...yz"
)

type WritingDirection int

const (
	WritingDirectionNatural     WritingDirection = -1 // Determines direction using the Unicode Bidi Algorithm rules P2 and P3
	WritingDirectionLeftToRight WritingDirection = 0  // Left to right writing direction
	WritingDirectionRightToLeft WritingDirection = 1  // Right to left writing direction
)

type ColorSystemEffect int

const (
	ColorSystemEffectNone ColorSystemEffect = iota
	ColorSystemEffectPressed
	ColorSystemEffectDeepPressed
	ColorSystemEffectDisabled
	ColorSystemEffectRollover
)

type ColorType int

const (
	ColorTypeComponentBased ColorType = iota // Colors with colorSpace, and floating point color components
	ColorTypePattern                         // Colors with patternImage
	ColorTypeCatalog                         // Colors with catalogNameComponent and colorNameComponent; these colors may have device and appearance specific representations
)

type ColorListName string

type ColorName string

type ControlTint uint

const (
	DefaultControlTint  ControlTint = 0
	BlueControlTint     ControlTint = 1
	GraphiteControlTint ControlTint = 6
	ClearControlTint    ControlTint = 7
)

type EventType uint

const (
	EventTypeLeftMouseDown      EventType = 1
	EventTypeLeftMouseUp        EventType = 2
	EventTypeRightMouseDown     EventType = 3
	EventTypeRightMouseUp       EventType = 4
	EventTypeMouseMoved         EventType = 5
	EventTypeLeftMouseDragged   EventType = 6
	EventTypeRightMouseDragged  EventType = 7
	EventTypeMouseEntered       EventType = 8
	EventTypeMouseExited        EventType = 9
	EventTypeKeyDown            EventType = 10
	EventTypeKeyUp              EventType = 11
	EventTypeFlagsChanged       EventType = 12
	EventTypeAppKitDefined      EventType = 13
	EventTypeSystemDefined      EventType = 14
	EventTypeApplicationDefined EventType = 15
	EventTypePeriodic           EventType = 16
	EventTypeCursorUpdate       EventType = 17
	EventTypeScrollWheel        EventType = 22
	EventTypeTabletPoint        EventType = 23
	EventTypeTabletProximity    EventType = 24
	EventTypeOtherMouseDown     EventType = 25
	EventTypeOtherMouseUp       EventType = 26
	EventTypeOtherMouseDragged  EventType = 27
	EventTypeGesture            EventType = 29
	EventTypeMagnify            EventType = 30
	EventTypeSwipe              EventType = 31
	EventTypeRotate             EventType = 18
	EventTypeBeginGesture       EventType = 19
	EventTypeEndGesture         EventType = 20
	EventTypeSmartMagnify       EventType = 32
	EventTypeQuickLook          EventType = 33
	EventTypePressure           EventType = 34
	EventTypeDirectTouch        EventType = 37
	EventTypeChangeMode         EventType = 38
)

type PressureBehavior int

const (
	PressureBehaviorUnknown            PressureBehavior = -1
	PressureBehaviorPrimaryDefault     PressureBehavior = 0
	PressureBehaviorPrimaryClick       PressureBehavior = 1
	PressureBehaviorPrimaryGeneric     PressureBehavior = 2
	PressureBehaviorPrimaryAccelerator PressureBehavior = 3
	PressureBehaviorPrimaryDeepClick   PressureBehavior = 5
	PressureBehaviorPrimaryDeepDrag    PressureBehavior = 6
)

type PointingDeviceType uint

const (
	PointingDeviceTypeUnknown PointingDeviceType = 0
	PointingDeviceTypePen     PointingDeviceType = 1
	PointingDeviceTypeCursor  PointingDeviceType = 2
	PointingDeviceTypeEraser  PointingDeviceType = 3
)

type EventButtonMask uint

const (
	EventButtonMaskPenTip       EventButtonMask = 1
	EventButtonMaskPenLowerSide EventButtonMask = 2
	EventButtonMaskPenUpperSide EventButtonMask = 4
)

type EventPhase uint

const (
	EventPhaseNone       EventPhase = 0 // event not associated with a phase.
	EventPhaseBegan      EventPhase = 0x1 << 0
	EventPhaseStationary EventPhase = 0x1 << 1
	EventPhaseChanged    EventPhase = 0x1 << 2
	EventPhaseEnded      EventPhase = 0x1 << 3
	EventPhaseCancelled  EventPhase = 0x1 << 4
	EventPhaseMayBegin   EventPhase = 0x1 << 5
)

type FontDescriptorAttributeName string

const (
	FontFamilyAttribute          FontDescriptorAttributeName = "NSFontFamilyAttribute"
	FontNameAttribute            FontDescriptorAttributeName = "NSFontNameAttribute"
	FontFaceAttribute            FontDescriptorAttributeName = "NSFontFaceAttribute"
	FontSizeAttribute            FontDescriptorAttributeName = "NSFontSizeAttribute"
	FontVisibleNameAttribute     FontDescriptorAttributeName = "NSFontVisibleNameAttribute"
	FontMatrixAttribute          FontDescriptorAttributeName = "NSFontMatrixAttribute"
	FontVariationAttribute       FontDescriptorAttributeName = "NSCTFontVariationAttribute"
	FontCharacterSetAttribute    FontDescriptorAttributeName = "NSCTFontCharacterSetAttribute"
	FontCascadeListAttribute     FontDescriptorAttributeName = "NSCTFontCascadeListAttribute"
	FontTraitsAttribute          FontDescriptorAttributeName = "NSCTFontTraitsAttribute"
	FontFixedAdvanceAttribute    FontDescriptorAttributeName = "NSCTFontFixedAdvanceAttribute"
	FontFeatureSettingsAttribute FontDescriptorAttributeName = "NSCTFontFeatureSettingsAttribute"
)

type FontDescriptorSymbolicTraits uint32

const (
	FontDescriptorTraitItalic             FontDescriptorSymbolicTraits = 1 << 0
	FontDescriptorTraitBold               FontDescriptorSymbolicTraits = 1 << 1
	FontDescriptorTraitExpanded           FontDescriptorSymbolicTraits = 1 << 5
	FontDescriptorTraitCondensed          FontDescriptorSymbolicTraits = 1 << 6
	FontDescriptorTraitMonoSpace          FontDescriptorSymbolicTraits = 1 << 10
	FontDescriptorTraitVertical           FontDescriptorSymbolicTraits = 1 << 11
	FontDescriptorTraitUIOptimized        FontDescriptorSymbolicTraits = 1 << 12
	FontDescriptorTraitTightLeading       FontDescriptorSymbolicTraits = 1 << 15
	FontDescriptorTraitLooseLeading       FontDescriptorSymbolicTraits = 1 << 16
	FontDescriptorClassMask               FontDescriptorSymbolicTraits = 0xF0000000
	FontDescriptorClassUnknown            FontDescriptorSymbolicTraits = 0 << 28
	FontDescriptorClassOldStyleSerifs     FontDescriptorSymbolicTraits = 1 << 28
	FontDescriptorClassTransitionalSerifs FontDescriptorSymbolicTraits = 2 << 28
	FontDescriptorClassModernSerifs       FontDescriptorSymbolicTraits = 3 << 28
	FontDescriptorClassClarendonSerifs    FontDescriptorSymbolicTraits = 4 << 28
	FontDescriptorClassSlabSerifs         FontDescriptorSymbolicTraits = 5 << 28
	FontDescriptorClassFreeformSerifs     FontDescriptorSymbolicTraits = 7 << 28
	FontDescriptorClassSansSerif          FontDescriptorSymbolicTraits = 8 << 28
	FontDescriptorClassOrnamentals        FontDescriptorSymbolicTraits = 9 << 28
	FontDescriptorClassScripts            FontDescriptorSymbolicTraits = 10 << 28
	FontDescriptorClassSymbolic           FontDescriptorSymbolicTraits = 12 << 28
)

type FontDescriptorSystemDesign string

const (
	FontDescriptorSystemDesignDefault    FontDescriptorSystemDesign = "NSCTFontUIFontDesignDefault"
	FontDescriptorSystemDesignSerif      FontDescriptorSystemDesign = "NSCTFontUIFontDesignSerif"
	FontDescriptorSystemDesignMonospaced FontDescriptorSystemDesign = "NSCTFontUIFontDesignMonospaced"
	FontDescriptorSystemDesignRounded    FontDescriptorSystemDesign = "NSCTFontUIFontDesignRounded"
)

type CompositingOperation uint

const (
	CompositingOperationClear CompositingOperation = iota
	CompositingOperationCopy
	CompositingOperationSourceOver
	CompositingOperationSourceIn
	CompositingOperationSourceOut
	CompositingOperationSourceAtop
	CompositingOperationDestinationOver
	CompositingOperationDestinationIn
	CompositingOperationDestinationOut
	CompositingOperationDestinationAtop
	CompositingOperationXOR
	CompositingOperationPlusDarker
	CompositingOperationPlusLighter

	CompositingOperationMultiply
	CompositingOperationScreen
	CompositingOperationOverlay
	CompositingOperationDarken
	CompositingOperationLighten
	CompositingOperationColorDodge
	CompositingOperationColorBurn
	CompositingOperationSoftLight
	CompositingOperationHardLight
	CompositingOperationDifference
	CompositingOperationExclusion

	CompositingOperationHue
	CompositingOperationSaturation
	CompositingOperationColor
	CompositingOperationLuminosity
)

type TIFFCompression uint

const (
	TIFFCompressionNone      TIFFCompression = 1
	TIFFCompressionCCITTFAX3 TIFFCompression = 3 /* 1 bps only */
	TIFFCompressionCCITTFAX4 TIFFCompression = 4 /* 1 bps only */
	TIFFCompressionLZW       TIFFCompression = 5
	TIFFCompressionJPEG      TIFFCompression = 6     /* No longer supported for input or output */
	TIFFCompressionNEXT      TIFFCompression = 32766 /* Input only */
	TIFFCompressionPackBits  TIFFCompression = 32773
	TIFFCompressionOldJPEG   TIFFCompression = 32865 /* No longer supported for input or output */
)

type ImageResizingMode int

const (
	ImageResizingModeStretch ImageResizingMode = 0
	ImageResizingModeTile    ImageResizingMode = 1
)

type ImageCacheMode uint

const (
	ImageCacheDefault ImageCacheMode = iota // unspecified. use image rep's default
	ImageCacheAlways                        // always generate a cache when drawing
	ImageCacheBySize                        // cache if cache size is smaller than original data
	ImageCacheNever                         // never cache, always draw direct
)

type UnderlineStyle int

const (
	UnderlineStyleNone              UnderlineStyle = 0x00
	UnderlineStyleSingle            UnderlineStyle = 0x01
	UnderlineStyleThick             UnderlineStyle = 0x02
	UnderlineStyleDouble            UnderlineStyle = 0x09
	UnderlineStylePatternSolid      UnderlineStyle = 0x0000
	UnderlineStylePatternDot        UnderlineStyle = 0x0100
	UnderlineStylePatternDash       UnderlineStyle = 0x0200
	UnderlineStylePatternDashDot    UnderlineStyle = 0x0300
	UnderlineStylePatternDashDotDot UnderlineStyle = 0x0400
	UnderlineStyleByWord            UnderlineStyle = 0x8000
)

type TypesetterBehavior int

type UserInterfaceItemIdentifier string

type FontTraitMask uint

const (
	ItalicFontMask                  FontTraitMask = 0x00000001
	BoldFontMask                    FontTraitMask = 0x00000002
	UnboldFontMask                  FontTraitMask = 0x00000004
	NonStandardCharacterSetFontMask FontTraitMask = 0x00000008
	NarrowFontMask                  FontTraitMask = 0x00000010
	ExpandedFontMask                FontTraitMask = 0x00000020
	CondensedFontMask               FontTraitMask = 0x00000040
	SmallCapsFontMask               FontTraitMask = 0x00000080
	PosterFontMask                  FontTraitMask = 0x00000100
	CompressedFontMask              FontTraitMask = 0x00000200
	FixedPitchFontMask              FontTraitMask = 0x00000400
	UnitalicFontMask                FontTraitMask = 0x01000000
)

type TextStorageEditActions uint

type GlyphProperty int

type LineBreakStrategy uint

const (
	LineBreakStrategyNone               LineBreakStrategy = 0
	LineBreakStrategyPushOut            LineBreakStrategy = 1 << 0
	LineBreakStrategyHangulWordPriority LineBreakStrategy = 1 << 1
	LineBreakStrategyStandard           LineBreakStrategy = 0xFFFF
)

type TextBlockValueType uint

type TextBlockDimension uint

type TextBlockLayer int

type TextBlockVerticalAlignment uint

type TextListOptions uint

type TextListMarkerFormat string

type TouchType int

const (
	TouchTypeDirect   TouchType = iota // A direct touch from a finger (on a screen)
	TouchTypeIndirect                  // An indirect touch (not a screen)
)

type TouchTypeMask uint

const (
	TouchTypeMaskDirect   TouchTypeMask = 1 << TouchTypeDirect   // A direct touch from a finger (on a screen)
	TouchTypeMaskIndirect TouchTypeMask = 1 << TouchTypeIndirect // An indirect touch (not a screen)
)

type TouchPhase uint

const (
	NSTouchPhaseBegan      TouchPhase = 1 << 0
	NSTouchPhaseMoved      TouchPhase = 1 << 1
	NSTouchPhaseStationary TouchPhase = 1 << 2
	NSTouchPhaseEnded      TouchPhase = 1 << 3
	NSTouchPhaseCancelled  TouchPhase = 1 << 4
	NSTouchPhaseTouching   TouchPhase = NSTouchPhaseBegan | NSTouchPhaseMoved | NSTouchPhaseStationary
	NSTouchPhaseAny        TouchPhase = 0xffffffffffffffff
)

type LevelIndicatorPlaceholderVisibility int

const (
	LevelIndicatorPlaceholderVisibilityAutomatic    LevelIndicatorPlaceholderVisibility = 0
	LevelIndicatorPlaceholderVisibilityAlways       LevelIndicatorPlaceholderVisibility = 1
	LevelIndicatorPlaceholderVisibilityWhileEditing LevelIndicatorPlaceholderVisibility = 2
)

type ProgressIndicatorStyle uint

const (
	ProgressIndicatorStyleBar      ProgressIndicatorStyle = 0
	ProgressIndicatorStyleSpinning ProgressIndicatorStyle = 1
)

type ScrollerKnobStyle int

const (
	ScrollerKnobStyleDefault ScrollerKnobStyle = 0 // dark with light border; good against any background
	ScrollerKnobStyleDark    ScrollerKnobStyle = 1 // dark; good against a light background
	ScrollerKnobStyleLight   ScrollerKnobStyle = 2 // light; good against a dark background
)

type ScrollerStyle int

const (
	ScrollerStyleLegacy  ScrollerStyle = 0
	ScrollerStyleOverlay ScrollerStyle = 1
)

type ScrollViewFindBarPosition int

const (
	ScrollViewFindBarPositionAboveHorizontalRuler ScrollViewFindBarPosition = 0
	ScrollViewFindBarPositionAboveContent         ScrollViewFindBarPosition = 1
	ScrollViewFindBarPositionBelowContent         ScrollViewFindBarPosition = 2
)

type ScrollElasticity int

const (
	ScrollElasticityAutomatic ScrollElasticity = 0 // automatically determine whether to allow elasticity on this axis
	ScrollElasticityNone      ScrollElasticity = 1 // disallow scrolling beyond document bounds on this axis
	ScrollElasticityAllowed   ScrollElasticity = 2 // allow content to be scrolled past its bounds on this axis in an elastic fashion
)

type SplitViewAutosaveName string
type StatusItemAutosaveName string

type TextFieldBezelStyle uint

const (
	TextFieldSquareBezel                     TextFieldBezelStyle = 0
	TextFieldRoundedBezelTextFieldBezelStyle TextFieldBezelStyle = 1
)

type SelectionAffinity uint

const (
	SelectionAffinityUpstream   SelectionAffinity = 0
	SelectionAffinityDownstream SelectionAffinity = 1
)

type SelectionGranularity uint

const (
	SelectByCharacter SelectionGranularity = 0
	SelectByWord      SelectionGranularity = 1
	SelectByParagraph SelectionGranularity = 2
)

type TextLayoutOrientation int

const (
	TextLayoutOrientationHorizontal TextLayoutOrientation = 0 // Lines rendered horizontally, grow top to bottom
	TextLayoutOrientationVertical   TextLayoutOrientation = 1 // Lines rendered vertically, grow right to left
)

type UsableScrollerParts uint
type ScrollerPart uint

type BezierPathElement uint
type WindingRule uint
type LineCapStyle uint
type LineJoinStyle uint

type PaperOrientation int
type PrintingPaginationMode uint

type PrinterPaperName string

type PrintJobDispositionValue string

type PrintRenderingQuality int

type PrintingPageOrder int

type PDFPanelOptions int

type PrintPanelJobStyleHint string

type PrintPanelOptions uint
type HelpAnchorName string

type PrinterTypeName string

type ColorSpaceModel int

const (
	ColorSpaceModelUnknown ColorSpaceModel = iota - 1
	ColorSpaceModelGray
	ColorSpaceModelRGB
	ColorSpaceModelCMYK
	ColorSpaceModelLAB
	ColorSpaceModelDeviceN
	ColorSpaceModelIndexed
	ColorSpaceModelPatterned
)

type ImageInterpolation uint

const (
	NSImageInterpolationDefault ImageInterpolation = 0
	NSImageInterpolationNone    ImageInterpolation = 1
	NSImageInterpolationLow     ImageInterpolation = 2
	NSImageInterpolationMedium  ImageInterpolation = 4
	NSImageInterpolationHigh    ImageInterpolation = 3
)

type ColorRenderingIntent int

const (
	ColorRenderingIntentDefault ColorRenderingIntent = iota
	ColorRenderingIntentAbsoluteColorimetric
	ColorRenderingIntentRelativeColorimetric
	ColorRenderingIntentPerceptual
	ColorRenderingIntentSaturation
)

type ImageLayoutDirection int

const (
	ImageLayoutDirectionUnspecified ImageLayoutDirection = -1
	ImageLayoutDirectionLeftToRight ImageLayoutDirection = 2
	ImageLayoutDirectionRightToLeft ImageLayoutDirection = 3
)

type PasteboardContentsOptions uint

const (
	PasteboardContentsCurrentHostOnly PasteboardContentsOptions = 1 << 0 // Specifies that the pasteboard contents should not be available to other devices
)

type ToolbarDisplayMode uint

const (
	ToolbarDisplayModeDefault ToolbarDisplayMode = iota
	ToolbarDisplayModeIconAndLabel
	ToolbarDisplayModeIconOnly
	ToolbarDisplayModeLabelOnly
)

type ToolbarSizeMode uint

const (
	ToolbarSizeModeDefault ToolbarSizeMode = iota
	ToolbarSizeModeRegular
	ToolbarSizeModeSmall
)

type TrackingAreaOptions uint

const (
	TrackingMouseEnteredAndExited TrackingAreaOptions = 0x01 // owner receives mouseEntered when mouse enters area, and mouseExited when mouse leaves area
	TrackingMouseMoved            TrackingAreaOptions = 0x02 // owner receives mouseMoved while mouse is within area.  Note that mouseMoved events do not contain userInfo
	TrackingCursorUpdate          TrackingAreaOptions = 0x04 // owner receives cursorUpdate when mouse enters area.  Cursor is set appropriately when mouse exits area

	/* When tracking area is active.  You must specify one of the following in the NSTrackingAreaOptions argument of -initWithRect:options:owner:userInfo: */

	TrackingActiveWhenFirstResponder TrackingAreaOptions = 0x10 // owner receives mouseEntered/Exited, mouseMoved, or cursorUpdate when view is first responder
	TrackingActiveInKeyWindow        TrackingAreaOptions = 0x20 // owner receives mouseEntered/Exited, mouseMoved, or cursorUpdate when view is in key window
	TrackingActiveInActiveApp        TrackingAreaOptions = 0x40 // owner receives mouseEntered/Exited, mouseMoved, or cursorUpdate when app is active
	TrackingActiveAlways             TrackingAreaOptions = 0x80 // owner receives mouseEntered/Exited or mouseMoved regardless of activation.  Not supported for NSTrackingCursorUpdate.

	/* Behavior of tracking area.  These values are used in NSTrackingAreaOptions.  You may specify any number of the following in the NSTrackingAreaOptions argument of -initWithRect:options:owner:userInfo: */

	TrackingAssumeInside           TrackingAreaOptions = 0x100 // If set, generate mouseExited event when mouse leaves area (same as assumeInside argument in deprecated addTrackingRect:owner:userData:assumeInside:)
	TrackingInVisibleRect          TrackingAreaOptions = 0x200 // If set, tracking occurs in visibleRect of view and rect is ignored
	TrackingEnabledDuringMouseDrag TrackingAreaOptions = 0x400 // If set, mouseEntered events will be generated as mouse is dragged.  If not set, mouseEntered events will be generated as mouse is moved, and on mouseUp after a drag.  mouseExited events are paired with mouseEntered events so their delivery is affected indirectly.  That is, if a mouseEntered event is generated and the mouse subsequently moves out of the trackingArea, a mouseExited event will be generated whether the mouse is being moved or dragged, independent of this flag.
)

type ViewLayerContentsPlacement int

type ViewLayerContentsRedrawPolicy int

type FocusRingType uint

const (
	FocusRingTypeDefault  FocusRingType = 0
	FocusRingTypeNone     FocusRingType = 1
	FocusRingTypeExterior FocusRingType = 2
)

type WindowNumberListOptions uint

const (
	WindowNumberListAllApplications WindowNumberListOptions = 1 << 0
	WindowNumberListAllSpaces       WindowNumberListOptions = 1 << 4
)

type WindowUserTabbingPreference int

const (
	WindowUserTabbingPreferenceManual WindowUserTabbingPreference = iota
	WindowUserTabbingPreferenceAlways
	WindowUserTabbingPreferenceInFullScreen
)

type ToolbarItemVisibilityPriority int

const (
	ToolbarItemVisibilityPriorityStandard ToolbarItemVisibilityPriority = 0
	ToolbarItemVisibilityPriorityLow      ToolbarItemVisibilityPriority = -1000
	ToolbarItemVisibilityPriorityHigh     ToolbarItemVisibilityPriority = 1000
	ToolbarItemVisibilityPriorityUser     ToolbarItemVisibilityPriority = 2000
)

type TouchBarItemPriority float32

type DragOperation uint

const (
	NSDragOperationNone    DragOperation = 0
	NSDragOperationCopy    DragOperation = 1
	NSDragOperationLink    DragOperation = 2
	NSDragOperationGeneric DragOperation = 4
	NSDragOperationPrivate DragOperation = 8
	NSDragOperationMove    DragOperation = 16
	NSDragOperationDelete  DragOperation = 32
	NSDragOperationEvery   DragOperation = 0xffffffffffffffff
)

type DraggingFormation int

const (
	DraggingFormationDefault DraggingFormation = iota // System determined formation.
	DraggingFormationNone                             // Drag images maintain their set positions relative to each other
	DraggingFormationPile                             // Drag images are placed on top of each other with random rotations
	DraggingFormationList                             // Drag images are laid out vertically, non-overlapping with the left edges aligned
	DraggingFormationStack                            // Drag images are laid out overlapping diagonally
)

type DraggingImageComponentKey string

type CellAttribute uint

const (
	CellDisabled               CellAttribute = 0
	CellState                  CellAttribute = 1
	PushInCell                 CellAttribute = 2
	CellEditable               CellAttribute = 3
	ChangeGrayCell             CellAttribute = 4
	CellHighlighted            CellAttribute = 5
	CellLightsByContents       CellAttribute = 6
	CellLightsByGray           CellAttribute = 7
	ChangeBackgroundCell       CellAttribute = 8
	CellLightsByBackground     CellAttribute = 9
	CellIsBordered             CellAttribute = 10
	CellHasOverlappingImage    CellAttribute = 11
	CellHasImageHorizontal     CellAttribute = 12
	CellHasImageOnLeftOrBottom CellAttribute = 13
	CellChangesContents        CellAttribute = 14
	CellIsInsetButton          CellAttribute = 15
	CellAllowsMixedState       CellAttribute = 16
)

type CellType uint

const (
	NullCellType  CellType = 0
	TextCellType  CellType = 1
	ImageCellType CellType = 2
)

type BackgroundStyle int

const (
	BackgroundStyleNormal BackgroundStyle = iota
	BackgroundStyleEmphasized
	BackgroundStyleRaised
	BackgroundStyleLowered
)

type RulerOrientation uint

const (
	HorizontalRuler RulerOrientation = iota
	VerticalRuler
)

type CellHitResult uint

const (
	CellHitNone             CellHitResult = 0
	CellHitContentArea      CellHitResult = 1 << 0
	CellHitEditableTextArea CellHitResult = 1 << 1
	CellHitTrackableArea    CellHitResult = 1 << 2
)

type BitmapFormat uint

const (
	BitmapFormatAlphaFirst               BitmapFormat = 1 << 0 // 0 means is alpha last (RGBA CMYKA etc.)
	BitmapFormatAlphaNonpremultiplied    BitmapFormat = 1 << 1 // 0 means is premultiplied
	BitmapFormatFloatingPointSamples     BitmapFormat = 1 << 2 // 0 is integer
	BitmapFormatSixteenBitLittleEndian   BitmapFormat = 1 << 8
	BitmapFormatThirtyTwoBitLittleEndian BitmapFormat = 1 << 9
	BitmapFormatSixteenBitBigEndian      BitmapFormat = 1 << 10
	BitmapFormatThirtyTwoBitBigEndian    BitmapFormat = 1 << 11
)

type GestureRecognizerState int

const (
	GestureRecognizerStatePossible   GestureRecognizerState        = iota // the recognizer has not yet recognized its gesture, but may be evaluating events. this is the default state
	GestureRecognizerStateBegan                                           // the recognizer has received events recognized as the gesture. the action method will be called at the next turn of the run loop
	GestureRecognizerStateChanged                                         // the recognizer has received events recognized as a change to the gesture. the action method will be called at the next turn of the run loop
	GestureRecognizerStateEnded                                           // the recognizer has received events recognized as the end of the gesture. the action method will be called at the next turn of the run loop and the recognizer will be reset to GestureRecognizerStatePossible
	GestureRecognizerStateCancelled                                       // the recognizer has received events resulting in the cancellation of the gesture. the action method will be called at the next turn of the run loop. the recognizer will be reset to GestureRecognizerStatePossible
	GestureRecognizerStateFailed                                          // the recognizer has received an event sequence that can not be recognized as the gesture. the action method will not be called and the recognizer will be reset to GestureRecognizerStatePossible
	GestureRecognizerStateRecognized = GestureRecognizerStateEnded        // the recognizer has received events recognized as the gesture. the action method will be called at the next turn of the run loop and the recognizer will be reset to GestureRecognizerStatePossible
)

type SaveOperationType uint

const (
	SaveOperation              SaveOperationType = 0
	SaveAsOperation            SaveOperationType = 1
	SaveToOperation            SaveOperationType = 2
	AutosaveInPlaceOperation   SaveOperationType = 4
	AutosaveElsewhereOperation SaveOperationType = 3
	AutosaveAsOperation        SaveOperationType = 5
)

type DocumentChangeType uint

const (
	ChangeDone              DocumentChangeType = 0
	ChangeUndone            DocumentChangeType = 1
	ChangeRedone            DocumentChangeType = 5
	ChangeCleared           DocumentChangeType = 2
	ChangeReadOtherContents DocumentChangeType = 3
	ChangeAutosaved         DocumentChangeType = 4
	ChangeDiscardable       DocumentChangeType = 256
)

type TableViewAnimationOptions uint

const (
	TableViewAnimationEffectNone TableViewAnimationOptions = 0x0

	TableViewAnimationEffectFade TableViewAnimationOptions = 0x1 // Fades in new rows.
	TableViewAnimationEffectGap  TableViewAnimationOptions = 0x2 // Creates a gap for newly inserted rows. This is useful for drag and drop animations that animate to a newly opened gap and should be used in -tableView:acceptDrop:row:dropOperation:.

	TableViewAnimationSlideUp    TableViewAnimationOptions = 0x10 // Animates a row in or out by sliding upward.
	TableViewAnimationSlideDown  TableViewAnimationOptions = 0x20 // Animates a row in or out by sliding downward.
	TableViewAnimationSlideLeft  TableViewAnimationOptions = 0x30 // Animates a row in by sliding from the left. Animates a row out by sliding towards the left.
	TableViewAnimationSlideRight TableViewAnimationOptions = 0x40 // Animates a row in by sliding from the right. Animates a row out by sliding towards the right.
)

type TableViewStyle int

const (
	TableViewStyleAutomatic TableViewStyle = iota
	TableViewStyleFullWidth
	TableViewStyleInset
	TableViewStyleSourceList
	TableViewStylePlain
)

type TableViewSelectionHighlightStyle int

const (
	TableViewSelectionHighlightStyleNone    TableViewSelectionHighlightStyle = -1
	TableViewSelectionHighlightStyleRegular TableViewSelectionHighlightStyle = 0
)

type TableViewGridLineStyle uint

const (
	TableViewGridNone                     TableViewGridLineStyle = 0
	TableViewSolidVerticalGridLineMask    TableViewGridLineStyle = 1 << 0
	TableViewSolidHorizontalGridLineMask  TableViewGridLineStyle = 1 << 1
	TableViewDashedHorizontalGridLineMask TableViewGridLineStyle = 1 << 3
)

type TableViewRowSizeStyle int

const (
	TableViewRowSizeStyleDefault TableViewRowSizeStyle = -1
	TableViewRowSizeStyleCustom  TableViewRowSizeStyle = 0
	TableViewRowSizeStyleSmall   TableViewRowSizeStyle = 1
	TableViewRowSizeStyleMedium  TableViewRowSizeStyle = 2
	TableViewRowSizeStyleLarge   TableViewRowSizeStyle = 3
)

type CollectionViewScrollDirection int

const (
	CollectionViewScrollDirectionVertical CollectionViewScrollDirection = iota
	CollectionViewScrollDirectionHorizontal
)

type TableViewDraggingDestinationFeedbackStyle int

const (
	TableViewDraggingDestinationFeedbackStyleNone       TableViewDraggingDestinationFeedbackStyle = -1
	TableViewDraggingDestinationFeedbackStyleRegular    TableViewDraggingDestinationFeedbackStyle = 0
	TableViewDraggingDestinationFeedbackStyleSourceList TableViewDraggingDestinationFeedbackStyle = 1
	TableViewDraggingDestinationFeedbackStyleGap        TableViewDraggingDestinationFeedbackStyle = 2
)

type TableViewDropOperation uint

const (
	øTableViewDropOn TableViewDropOperation = iota
	øTableViewDropAbove
)

type BrowserColumnsAutosaveName string

type BrowserColumnResizingType uint

const (
	BrowserNoColumnResizing   BrowserColumnResizingType = 0
	BrowserAutoColumnResizing BrowserColumnResizingType = 1
	BrowserUserColumnResizing BrowserColumnResizingType = 2
)

type BrowserDropOperation uint

const (
	BrowserDropOn BrowserDropOperation = iota
	BrowserDropAbove
)

type CollectionViewSupplementaryElementKind string

type CollectionViewScrollPosition uint

const (
	CollectionViewScrollPositionNone                  CollectionViewScrollPosition = 0
	CollectionViewScrollPositionTop                   CollectionViewScrollPosition = 1 << 0
	CollectionViewScrollPositionCenteredVertically    CollectionViewScrollPosition = 1 << 1
	CollectionViewScrollPositionBottom                CollectionViewScrollPosition = 1 << 2
	CollectionViewScrollPositionNearestHorizontalEdge CollectionViewScrollPosition = 1 << 9 /* Nearer of TopBottom */
	CollectionViewScrollPositionLeft                  CollectionViewScrollPosition = 1 << 3
	CollectionViewScrollPositionCenteredHorizontally  CollectionViewScrollPosition = 1 << 4
	CollectionViewScrollPositionRight                 CollectionViewScrollPosition = 1 << 5
	CollectionViewScrollPositionLeadingEdge           CollectionViewScrollPosition = 1 << 6 /* Left if LTR Right if RTL */
	CollectionViewScrollPositionTrailingEdge          CollectionViewScrollPosition = 1 << 7 /* Right if LTR Left if RTL */
	CollectionViewScrollPositionNearestVerticalEdge   CollectionViewScrollPosition = 1 << 8 /* Nearer of Leading Trailing */
)

type CollectionViewItemHighlightState int

const (
	CollectionViewItemHighlightNone           CollectionViewItemHighlightState = 0
	CollectionViewItemHighlightForSelection   CollectionViewItemHighlightState = 1
	CollectionViewItemHighlightForDeselection CollectionViewItemHighlightState = 2
	CollectionViewItemHighlightAsDropTarget   CollectionViewItemHighlightState = 3
)

type CollectionViewDropOperation int

const (
	CollectionViewDropOn     CollectionViewDropOperation = 0
	CollectionViewDropBefore CollectionViewDropOperation = 1
)

type CollectionViewDecorationElementKind string

type CollectionElementCategory int

const (
	CollectionElementCategoryItem CollectionElementCategory = iota
	CollectionElementCategorySupplementaryView
	CollectionElementCategoryDecorationView
	CollectionElementCategoryInterItemGap
)

type CollectionViewTransitionLayoutAnimatedKey string

type CollectionUpdateAction int

const (
	CollectionUpdateActionInsert CollectionUpdateAction = iota
	CollectionUpdateActionDelete
	CollectionUpdateActionReload
	CollectionUpdateActionMove
	CollectionUpdateActionNone
)

type PathStyle int

type SearchFieldRecentsAutosaveName string

type TableViewColumnAutoresizingStyle uint

const (
	TableViewNoColumnAutoresizing TableViewColumnAutoresizingStyle = iota
	TableViewUniformColumnAutoresizingStyle
	TableViewSequentialColumnAutoresizingStyle
	TableViewReverseSequentialColumnAutoresizingStyle
	TableViewLastColumnOnlyAutoresizingStyle
	TableViewFirstColumnOnlyAutoresizingStyle
)

type TableViewAutosaveName string

type TableColumnResizingOptions uint

const (
	TableColumnNoResizing       TableColumnResizingOptions = 0      // Disallow any kind of resizing.
	TableColumnAutoresizingMask TableColumnResizingOptions = 1 << 0 // This column can be resized as the table is resized.
	TableColumnUserResizingMask TableColumnResizingOptions = 1 << 1 // The user can resize this column manually.
)

type CollectionLayoutSectionOrthogonalScrollingBehavior int

const (
	CollectionLayoutSectionOrthogonalScrollingBehaviorNone CollectionLayoutSectionOrthogonalScrollingBehavior = iota
	CollectionLayoutSectionOrthogonalScrollingBehaviorContinuous
	CollectionLayoutSectionOrthogonalScrollingBehaviorContinuousGroupLeadingBoundary
	CollectionLayoutSectionOrthogonalScrollingBehaviorPaging
	CollectionLayoutSectionOrthogonalScrollingBehaviorGroupPaging
	CollectionLayoutSectionOrthogonalScrollingBehaviorGroupPagingCentered
)

type RectAlignment int

const (
	RectAlignmentNone RectAlignment = iota
	RectAlignmentTop
	RectAlignmentTopLeading
	RectAlignmentLeading
	RectAlignmentBottomLeading
	RectAlignmentBottom
	RectAlignmentBottomTrailing
	RectAlignmentTrailing
	RectAlignmentTopTrailing
)

type DirectionalRectEdge uint

const (
	DirectionalRectEdgeNone     DirectionalRectEdge = 0
	DirectionalRectEdgeTop      DirectionalRectEdge = 1 << 0
	DirectionalRectEdgeLeading  DirectionalRectEdge = 1 << 1
	DirectionalRectEdgeBottom   DirectionalRectEdge = 1 << 2
	DirectionalRectEdgeTrailing DirectionalRectEdge = 1 << 3
	DirectionalRectEdgeAll      DirectionalRectEdge = DirectionalRectEdgeTop | DirectionalRectEdgeLeading | DirectionalRectEdgeBottom | DirectionalRectEdgeTrailing
)

type MatrixMode uint

const (
	RadioModeMatrix     MatrixMode = 0
	HighlightModeMatrix MatrixMode = 1
	ListModeMatrix      MatrixMode = 2
	TrackModeMatrix     MatrixMode = 3
)

type TableRowActionEdge int

const (
	TableRowActionEdgeLeading  TableRowActionEdge = iota // Action buttons that appear on the leading (or left) edge of an TableRowView
	TableRowActionEdgeTrailing                           // Action buttons that appear on the trailing (or right) edge of an TableRowView
)

type TableViewRowActionStyle int

const (
	TableViewRowActionStyleRegular     TableViewRowActionStyle = iota
	TableViewRowActionStyleDestructive                         // Destructive buttons will do a special animation when being deleted
)

type AnimatablePropertyKey string

const (
	AnimationTriggerOrderIn  AnimatablePropertyKey = "NSAnimationTriggerOrderIn"
	AnimationTriggerOrderOut AnimatablePropertyKey = "NSAnimationTriggerOrderOut"
)

type AnimationProgress float32

type AnimationBlockingMode uint

const (
	AnimationBlocking AnimationBlockingMode = iota
	AnimationNonblocking
	AnimationNonblockingThreaded
)

type AnimationCurve uint

const (
	AnimationEaseInOut AnimationCurve = iota
	AnimationEaseIn
	AnimationEaseOut
	AnimationLinear
)

type DraggingContext int

const (
	DraggingContextOutsideApplication DraggingContext = iota
	DraggingContextWithinApplication
)

type SpringLoadingOptions uint

const (
	SpringLoadingDisabled             SpringLoadingOptions = 0
	SpringLoadingEnabled              SpringLoadingOptions = 1 << 0
	SpringLoadingContinuousActivation SpringLoadingOptions = 1 << 1
	SpringLoadingNoHover              SpringLoadingOptions = 1 << 3
)

type AlertStyle uint

const (
	AlertStyleWarning       AlertStyle = 0
	AlertStyleInformational AlertStyle = 1
	AlertStyleCritical      AlertStyle = 2
)

type ColorPanelMode int

const (
	ColorPanelModeNone          ColorPanelMode = -1
	ColorPanelModeGray          ColorPanelMode = 0
	ColorPanelModeRGB           ColorPanelMode = 1
	ColorPanelModeCMYK          ColorPanelMode = 2
	ColorPanelModeHSB           ColorPanelMode = 3
	ColorPanelModeCustomPalette ColorPanelMode = 4
	ColorPanelModeColorList     ColorPanelMode = 5
	ColorPanelModeWheel         ColorPanelMode = 6
	ColorPanelModeCrayon        ColorPanelMode = 7
)

type ColorPanelOptions uint

const (
	ColorPanelGrayModeMask          ColorPanelOptions = 0x00000001
	ColorPanelRGBModeMask           ColorPanelOptions = 0x00000002
	ColorPanelCMYKModeMask          ColorPanelOptions = 0x00000004
	ColorPanelHSBModeMask           ColorPanelOptions = 0x00000008
	ColorPanelCustomPaletteModeMask ColorPanelOptions = 0x00000010
	ColorPanelColorListModeMask     ColorPanelOptions = 0x00000020
	ColorPanelWheelModeMask         ColorPanelOptions = 0x00000040
	ColorPanelCrayonModeMask        ColorPanelOptions = 0x00000080
	ColorPanelAllModesMask          ColorPanelOptions = 0x0000ffff
)

type FontPanelModeMask uint

const (
	SFontPanelModeMaskFace                FontPanelModeMask = 1 << 0
	SFontPanelModeMaskSize                FontPanelModeMask = 1 << 1
	SFontPanelModeMaskCollection          FontPanelModeMask = 1 << 2
	SFontPanelModeMaskUnderlineEffect     FontPanelModeMask = 1 << 8
	SFontPanelModeMaskStrikethroughEffect FontPanelModeMask = 1 << 9
	SFontPanelModeMaskTextColorEffect     FontPanelModeMask = 1 << 10
	SFontPanelModeMaskDocumentColorEffect FontPanelModeMask = 1 << 11
	SFontPanelModeMaskShadowEffect        FontPanelModeMask = 1 << 12
	SFontPanelModeMaskAllEffects          FontPanelModeMask = 0xFFF00
	SFontPanelModesMaskStandardModes      FontPanelModeMask = 0xFFFF
	SFontPanelModesMaskAllModes           FontPanelModeMask = 0xFFFFFFFF
)

type FontAction uint

const (
	NoFontChangeAction    FontAction = 0
	ViaPanelFontAction    FontAction = 1
	AddTraitFontAction    FontAction = 2
	SizeUpFontAction      FontAction = 3
	SizeDownFontAction    FontAction = 4
	HeavierFontAction     FontAction = 5
	LighterFontAction     FontAction = 6
	RemoveTraitFontAction FontAction = 7
)

type ApplicationActivationPolicy int

const (
	ApplicationActivationPolicyRegular ApplicationActivationPolicy = iota
	ApplicationActivationPolicyAccessory
	ApplicationActivationPolicyProhibited
)

type ModalResponse int

const (
	ModalResponseStop       ModalResponse = -1000
	ModalResponseAbort      ModalResponse = -1001
	ModalResponseContinue   ModalResponse = -1002
	ModalResponseOK         ModalResponse = 1
	ModalResponseCancel     ModalResponse = 0
	AlertFirstButtonReturn  ModalResponse = 1000
	AlertSecondButtonReturn ModalResponse = 1001
	AlertThirdButtonReturn  ModalResponse = 1002
)

// BezelStyle used by the bezelStyle property.
type BezelStyle uint

const (
	BezelStyleRounded           BezelStyle = 1
	BezelStyleRegularSquare     BezelStyle = 2
	BezelStyleDisclosure        BezelStyle = 5
	BezelStyleShadowlessSquare  BezelStyle = 6
	BezelStyleCircular          BezelStyle = 7
	BezelStyleTexturedSquare    BezelStyle = 8
	BezelStyleHelpButton        BezelStyle = 9
	BezelStyleSmallSquare       BezelStyle = 10
	BezelStyleTexturedRounded   BezelStyle = 11
	BezelStyleRoundRect         BezelStyle = 12
	BezelStyleRecessed          BezelStyle = 13
	BezelStyleRoundedDisclosure BezelStyle = 14
	BezelStyleInline            BezelStyle = 15
)

// ButtonType is button types that can be specified using setButtonType.
type ButtonType uint

const (
	ButtonTypeMomentaryLight        ButtonType = 0
	ButtonTypePushOnPushOff         ButtonType = 1
	ButtonTypeToggle                ButtonType = 2
	ButtonTypeSwitch                ButtonType = 3
	ButtonTypeRadio                 ButtonType = 4
	ButtonTypeMomentaryChange       ButtonType = 5
	ButtonTypeOnOff                 ButtonType = 6
	ButtonTypeMomentaryPushIn       ButtonType = 7
	ButtonTypeAccelerator           ButtonType = 8
	ButtonTypeMultiLevelAccelerator ButtonType = 9
)

type CellImagePosition uint

const (
	NoImage       CellImagePosition = 0
	ImageOnly     CellImagePosition = 1
	ImageLeft     CellImagePosition = 2
	ImageRight    CellImagePosition = 3
	ImageBelow    CellImagePosition = 4
	ImageAbove    CellImagePosition = 5
	ImageOverlaps CellImagePosition = 6
	ImageLeading  CellImagePosition = 7
	ImageTrailing CellImagePosition = 8
)

type CellStyleMask uint

const (
	NoCellMask               CellStyleMask = 0
	ContentsCellMask         CellStyleMask = 1
	PushInCellMask           CellStyleMask = 2
	ChangeGrayCellMask       CellStyleMask = 4
	ChangeBackgroundCellMask CellStyleMask = 8
)

type NibName string

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

// EventModifierFlags is flags that represent key states in an event object
type EventModifierFlags uint

const (
	EventModifierFlagCapsLock                   EventModifierFlags = 1 << 16 // Set if Caps Lock key is pressed.
	EventModifierFlagShift                      EventModifierFlags = 1 << 17 // Set if Shift key is pressed.
	EventModifierFlagControl                    EventModifierFlags = 1 << 18 // Set if Control key is pressed.
	EventModifierFlagOption                     EventModifierFlags = 1 << 19 // Set if Option or Alternate key is pressed.
	EventModifierFlagCommand                    EventModifierFlags = 1 << 20 // Set if Command key is pressed.
	EventModifierFlagNumericPad                 EventModifierFlags = 1 << 21 // Set if any key in the numeric keypad is pressed.
	EventModifierFlagHelp                       EventModifierFlags = 1 << 22 // Set if the Help key is pressed.
	EventModifierFlagFunction                   EventModifierFlags = 1 << 23 // Set if any function key is pressed.
	EventModifierFlagDeviceIndependentFlagsMask EventModifierFlags = 0xffff0000
)

type FontTextStyle string

const (
	FontTextStyleLargeTitle  FontTextStyle = "UICTFontTextStyleTitle0"
	FontTextStyleTitle1      FontTextStyle = "UICTFontTextStyleTitle1"
	FontTextStyleTitle2      FontTextStyle = "UICTFontTextStyleTitle2"
	FontTextStyleTitle3      FontTextStyle = "UICTFontTextStyleTitle3"
	FontTextStyleHeadline    FontTextStyle = "UICTFontTextStyleHeadline"
	FontTextStyleSubheadline FontTextStyle = "UICTFontTextStyleSubhead"
	FontTextStyleBody        FontTextStyle = "UICTFontTextStyleBody"
	FontTextStyleCallout     FontTextStyle = "UICTFontTextStyleCallout"
	FontTextStyleFootnote    FontTextStyle = "UICTFontTextStyleFootnote"
	FontTextStyleCaption1    FontTextStyle = "UICTFontTextStyleCaption1"
	FontTextStyleCaption2    FontTextStyle = "UICTFontTextStyleCaption2"
)

type FontWeight float64

const (
	FontWeightUltraLight FontWeight = -0.8
	FontWeightThin       FontWeight = -0.6
	FontWeightLight      FontWeight = -0.4
	FontWeightRegular    FontWeight = 0
	FontWeightMedium     FontWeight = 0.23
	FontWeightSemibold   FontWeight = 0.3
	FontWeightBold       FontWeight = 0.4
	FontWeightHeavy      FontWeight = 0.56
	FontWeightBlack      FontWeight = 0.62
)

type GridRowAlignment int

const (
	GridRowAlignmentInherited GridRowAlignment = iota
	GridRowAlignmentNone
	GridRowAlignmentFirstBaseline
	GridRowAlignmentLastBaseline
)

type GridCellPlacement int

const (
	GridCellPlacementInherited GridCellPlacement = iota
	GridCellPlacementNone
	GridCellPlacementLeading

	GridCellPlacementTrailing

	GridCellPlacementCenter
	GridCellPlacementFill
	GridCellPlacementTop    = GridCellPlacementLeading
	GridCellPlacementBottom = GridCellPlacementTrailing
)

type ImageFrameStyle uint

const (
	ImageFrameNone ImageFrameStyle = iota
	ImageFramePhoto
	ImageFrameGrayBezel
	ImageFrameGroove
	ImageFrameButton
)

type ImageAlignment uint

const (
	ImageAlignCenter ImageAlignment = iota
	ImageAlignTop
	ImageAlignTopLeft
	ImageAlignTopRight
	ImageAlignLeft
	ImageAlignBottom
	ImageAlignBottomLeft
	ImageAlignBottomRight
	ImageAlignRight
)

type ImageSymbolScale int

const (
	ImageSymbolScaleSmall  ImageSymbolScale = 1
	ImageSymbolScaleMedium ImageSymbolScale = 2
	ImageSymbolScaleLarge  ImageSymbolScale = 3
)

type ImageName string

const (
	ImageNameAddTemplate                             ImageName = "NSAddTemplate"
	ImageNameBluetoothTemplate                       ImageName = "NSBluetoothTemplate"
	ImageNameBonjour                                 ImageName = "NSBonjour"
	ImageNameBookmarksTemplate                       ImageName = "NSBookmarksTemplate"
	ImageNameCaution                                 ImageName = "NSCaution"
	ImageNameComputer                                ImageName = "NSComputer"
	ImageNameEnterFullScreenTemplate                 ImageName = "NSEnterFullScreenTemplate"
	ImageNameExitFullScreenTemplate                  ImageName = "NSExitFullScreenTemplate"
	ImageNameFolder                                  ImageName = "NSFolder"
	ImageNameFolderBurnable                          ImageName = "NSFolderBurnable"
	ImageNameFolderSmart                             ImageName = "NSFolderSmart"
	ImageNameFollowLinkFreestandingTemplate          ImageName = "NSFollowLinkFreestandingTemplate"
	ImageNameHomeTemplate                            ImageName = "NSHomeTemplate"
	ImageNameIChatTheaterTemplate                    ImageName = "NSIChatTheaterTemplate"
	ImageNameLockLockedTemplate                      ImageName = "NSLockLockedTemplate"
	ImageNameLockUnlockedTemplate                    ImageName = "NSLockUnlockedTemplate"
	ImageNameNetwork                                 ImageName = "NSNetwork"
	ImageNamePathTemplate                            ImageName = "NSPathTemplate"
	ImageNameQuickLookTemplate                       ImageName = "NSQuickLookTemplate"
	ImageNameRefreshFreestandingTemplate             ImageName = "NSRefreshFreestandingTemplate"
	ImageNameRefreshTemplate                         ImageName = "NSRefreshTemplate"
	ImageNameRemoveTemplate                          ImageName = "NSRemoveTemplate"
	ImageNameRevealFreestandingTemplate              ImageName = "NSRevealFreestandingTemplate"
	ImageNameShareTemplate                           ImageName = "NSShareTemplate"
	ImageNameSlideshowTemplate                       ImageName = "NSSlideshowTemplate"
	ImageNameStatusAvailable                         ImageName = "NSStatusAvailable"
	ImageNameStatusNone                              ImageName = "NSStatusNone"
	ImageNameStatusPartiallyAvailable                ImageName = "NSStatusPartiallyAvailable"
	ImageNameStatusUnavailable                       ImageName = "NSStatusUnavailable"
	ImageNameStopProgressFreestandingTemplate        ImageName = "NSStopProgressFreestandingTemplate"
	ImageNameStopProgressTemplate                    ImageName = "NSStopProgressTemplate"
	ImageNameTrashEmpty                              ImageName = "NSTrashEmpty"
	ImageNameTrashFull                               ImageName = "NSTrashFull"
	ImageNameActionTemplate                          ImageName = "NSActionTemplate"
	ImageNameSmartBadgeTemplate                      ImageName = "NSSmartBadgeTemplate"
	ImageNameIconViewTemplate                        ImageName = "NSIconViewTemplate"
	ImageNameListViewTemplate                        ImageName = "NSListViewTemplate"
	ImageNameColumnViewTemplate                      ImageName = "NSColumnViewTemplate"
	ImageNameFlowViewTemplate                        ImageName = "NSFlowViewTemplate"
	ImageNameInvalidDataFreestandingTemplate         ImageName = "NSInvalidDataFreestandingTemplate"
	ImageNameGoForwardTemplate                       ImageName = "NSGoForwardTemplate"
	ImageNameGoBackTemplate                          ImageName = "NSGoBackTemplate"
	ImageNameGoRightTemplate                         ImageName = "NSGoRightTemplate"
	ImageNameGoLeftTemplate                          ImageName = "NSGoLeftTemplate"
	ImageNameRightFacingTriangleTemplate             ImageName = "NSRightFacingTriangleTemplate"
	ImageNameLeftFacingTriangleTemplate              ImageName = "NSLeftFacingTriangleTemplate"
	ImageNameMobileMe                                ImageName = "NSMobileMe"
	ImageNameMultipleDocuments                       ImageName = "NSMultipleDocuments"
	ImageNameUserAccounts                            ImageName = "NSUserAccounts"
	ImageNamePreferencesGeneral                      ImageName = "NSPreferencesGeneral"
	ImageNameAdvanced                                ImageName = "NSAdvanced"
	ImageNameInfo                                    ImageName = "NSInfo"
	ImageNameFontPanel                               ImageName = "NSFontPanel"
	ImageNameColorPanel                              ImageName = "NSColorPanel"
	ImageNameUser                                    ImageName = "NSUser"
	ImageNameUserGroup                               ImageName = "NSUserGroup"
	ImageNameEveryone                                ImageName = "NSEveryone"
	ImageNameUserGuest                               ImageName = "NSUserGuest"
	ImageNameMenuOnStateTemplate                     ImageName = "NSMenuOnStateTemplate"
	ImageNameMenuMixedStateTemplate                  ImageName = "NSMenuMixedStateTemplate"
	ImageNameApplicationIcon                         ImageName = "NSApplicationIcon"
	ImageNameTouchBarAddDetailTemplate               ImageName = "NSTouchBarAddDetailTemplate"
	ImageNameTouchBarAddTemplate                     ImageName = "NSTouchBarAddTemplate"
	ImageNameTouchBarAlarmTemplate                   ImageName = "NSTouchBarAlarmTemplate"
	ImageNameTouchBarAudioInputMuteTemplate          ImageName = "NSTouchBarAudioInputMuteTemplate"
	ImageNameTouchBarAudioInputTemplate              ImageName = "NSTouchBarAudioInputTemplate"
	ImageNameTouchBarAudioOutputMuteTemplate         ImageName = "NSTouchBarAudioOutputMuteTemplate"
	ImageNameTouchBarAudioOutputVolumeHighTemplate   ImageName = "NSTouchBarAudioOutputVolumeHighTemplate"
	ImageNameTouchBarAudioOutputVolumeLowTemplate    ImageName = "NSTouchBarAudioOutputVolumeLowTemplate"
	ImageNameTouchBarAudioOutputVolumeMediumTemplate ImageName = "NSTouchBarAudioOutputVolumeMediumTemplate"
	ImageNameTouchBarAudioOutputVolumeOffTemplate    ImageName = "NSTouchBarAudioOutputVolumeOffTemplate"
	ImageNameTouchBarBookmarksTemplate               ImageName = "NSTouchBarBookmarksTemplate"
	ImageNameTouchBarColorPickerFill                 ImageName = "NSTouchBarColorPickerFill"
	ImageNameTouchBarColorPickerFont                 ImageName = "NSTouchBarColorPickerFont"
	ImageNameTouchBarColorPickerStroke               ImageName = "NSTouchBarColorPickerStroke"
	ImageNameTouchBarCommunicationAudioTemplate      ImageName = "NSTouchBarCommunicationAudioTemplate"
	ImageNameTouchBarCommunicationVideoTemplate      ImageName = "NSTouchBarCommunicationVideoTemplate"
	ImageNameTouchBarComposeTemplate                 ImageName = "NSTouchBarComposeTemplate"
	ImageNameTouchBarDeleteTemplate                  ImageName = "NSTouchBarDeleteTemplate"
	ImageNameTouchBarDownloadTemplate                ImageName = "NSTouchBarDownloadTemplate"
	ImageNameTouchBarEnterFullScreenTemplate         ImageName = "NSTouchBarEnterFullScreenTemplate"
	ImageNameTouchBarExitFullScreenTemplate          ImageName = "NSTouchBarExitFullScreenTemplate"
	ImageNameTouchBarFastForwardTemplate             ImageName = "NSTouchBarFastForwardTemplate"
	ImageNameTouchBarFolderCopyToTemplate            ImageName = "NSTouchBarFolderCopyToTemplate"
	ImageNameTouchBarFolderMoveToTemplate            ImageName = "NSTouchBarFolderMoveToTemplate"
	ImageNameTouchBarFolderTemplate                  ImageName = "NSTouchBarFolderTemplate"
	ImageNameTouchBarGetInfoTemplate                 ImageName = "NSTouchBarGetInfoTemplate"
	ImageNameTouchBarGoBackTemplate                  ImageName = "NSTouchBarGoBackTemplate"
	ImageNameTouchBarGoDownTemplate                  ImageName = "NSTouchBarGoDownTemplate"
	ImageNameTouchBarGoForwardTemplate               ImageName = "NSTouchBarGoForwardTemplate"
	ImageNameTouchBarGoUpTemplate                    ImageName = "NSTouchBarGoUpTemplate"
	ImageNameTouchBarHistoryTemplate                 ImageName = "NSTouchBarHistoryTemplate"
	ImageNameTouchBarIconViewTemplate                ImageName = "NSTouchBarIconViewTemplate"
	ImageNameTouchBarListViewTemplate                ImageName = "NSTouchBarListViewTemplate"
	ImageNameTouchBarMailTemplate                    ImageName = "NSTouchBarMailTemplate"
	ImageNameTouchBarNewFolderTemplate               ImageName = "NSTouchBarNewFolderTemplate"
	ImageNameTouchBarNewMessageTemplate              ImageName = "NSTouchBarNewMessageTemplate"
	ImageNameTouchBarOpenInBrowserTemplate           ImageName = "NSTouchBarOpenInBrowserTemplate"
	ImageNameTouchBarPauseTemplate                   ImageName = "NSTouchBarPauseTemplate"
	ImageNameTouchBarPlayPauseTemplate               ImageName = "NSTouchBarPlayPauseTemplate"
	ImageNameTouchBarPlayTemplate                    ImageName = "NSTouchBarPlayTemplate"
	ImageNameTouchBarQuickLookTemplate               ImageName = "NSTouchBarQuickLookTemplate"
	ImageNameTouchBarRecordStartTemplate             ImageName = "NSTouchBarRecordStartTemplate"
	ImageNameTouchBarRecordStopTemplate              ImageName = "NSTouchBarRecordStopTemplate"
	ImageNameTouchBarRefreshTemplate                 ImageName = "NSTouchBarRefreshTemplate"
	ImageNameTouchBarRemoveTemplate                  ImageName = "NSTouchBarRemoveTemplate"
	ImageNameTouchBarRewindTemplate                  ImageName = "NSTouchBarRewindTemplate"
	ImageNameTouchBarRotateLeftTemplate              ImageName = "NSTouchBarRotateLeftTemplate"
	ImageNameTouchBarRotateRightTemplate             ImageName = "NSTouchBarRotateRightTemplate"
	ImageNameTouchBarSearchTemplate                  ImageName = "NSTouchBarSearchTemplate"
	ImageNameTouchBarShareTemplate                   ImageName = "NSTouchBarShareTemplate"
	ImageNameTouchBarSidebarTemplate                 ImageName = "NSTouchBarSidebarTemplate"
	ImageNameTouchBarSkipAhead15SecondsTemplate      ImageName = "NSTouchBarSkipAhead15SecondsTemplate"
	ImageNameTouchBarSkipAhead30SecondsTemplate      ImageName = "NSTouchBarSkipAhead30SecondsTemplate"
	ImageNameTouchBarSkipAheadTemplate               ImageName = "NSTouchBarSkipAheadTemplate"
	ImageNameTouchBarSkipBack15SecondsTemplate       ImageName = "NSTouchBarSkipBack15SecondsTemplate"
	ImageNameTouchBarSkipBack30SecondsTemplate       ImageName = "NSTouchBarSkipBack30SecondsTemplate"
	ImageNameTouchBarSkipBackTemplate                ImageName = "NSTouchBarSkipBackTemplate"
	ImageNameTouchBarSkipToEndTemplate               ImageName = "NSTouchBarSkipToEndTemplate"
	ImageNameTouchBarSkipToStartTemplate             ImageName = "NSTouchBarSkipToStartTemplate"
	ImageNameTouchBarSlideshowTemplate               ImageName = "NSTouchBarSlideshowTemplate"
	ImageNameTouchBarTagIconTemplate                 ImageName = "NSTouchBarTagIconTemplate"
	ImageNameTouchBarTextBoldTemplate                ImageName = "NSTouchBarTextBoldTemplate"
	ImageNameTouchBarTextBoxTemplate                 ImageName = "NSTouchBarTextBoxTemplate"
	ImageNameTouchBarTextCenterAlignTemplate         ImageName = "NSTouchBarTextCenterAlignTemplate"
	ImageNameTouchBarTextItalicTemplate              ImageName = "NSTouchBarTextItalicTemplate"
	ImageNameTouchBarTextJustifiedAlignTemplate      ImageName = "NSTouchBarTextJustifiedAlignTemplate"
	ImageNameTouchBarTextLeftAlignTemplate           ImageName = "NSTouchBarTextLeftAlignTemplate"
	ImageNameTouchBarTextListTemplate                ImageName = "NSTouchBarTextListTemplate"
	ImageNameTouchBarTextRightAlignTemplate          ImageName = "NSTouchBarTextRightAlignTemplate"
	ImageNameTouchBarTextStrikethroughTemplate       ImageName = "NSTouchBarTextStrikethroughTemplate"
	ImageNameTouchBarTextUnderlineTemplate           ImageName = "NSTouchBarTextUnderlineTemplate"
	ImageNameTouchBarUserAddTemplate                 ImageName = "NSTouchBarUserAddTemplate"
	ImageNameTouchBarUserGroupTemplate               ImageName = "NSTouchBarUserGroupTemplate"
	ImageNameTouchBarUserTemplate                    ImageName = "NSTouchBarUserTemplate"
	ImageNameTouchBarVolumeDownTemplate              ImageName = "NSTouchBarVolumeDownTemplate"
	ImageNameTouchBarVolumeUpTemplate                ImageName = "NSTouchBarVolumeUpTemplate"
	ImageNameTouchBarPlayheadTemplate                ImageName = "NSTouchBarPlayheadTemplate"
)

type UserInterfaceLayoutDirection int

const (
	UserInterfaceLayoutDirectionLeftToRight UserInterfaceLayoutDirection = 0
	UserInterfaceLayoutDirectionRightToLeft UserInterfaceLayoutDirection = 1
)

type UserInterfaceLayoutOrientation int

const (
	UserInterfaceLayoutOrientationHorizontal UserInterfaceLayoutOrientation = 0
	UserInterfaceLayoutOrientationVertical   UserInterfaceLayoutOrientation = 1
)

type LayoutConstraintOrientation int

const (
	LayoutConstraintOrientationHorizontal LayoutConstraintOrientation = 0
	LayoutConstraintOrientationVertical   LayoutConstraintOrientation = 1
)

type LayoutAttribute int

const (
	LayoutAttributeNotAnAttribute LayoutAttribute = iota
	LayoutAttributeLeft
	LayoutAttributeRight
	LayoutAttributeTop
	LayoutAttributeBottom
	LayoutAttributeLeading
	LayoutAttributeTrailing
	LayoutAttributeWidth
	LayoutAttributeHeight
	LayoutAttributeCenterX
	LayoutAttributeCenterY
	LayoutAttributeLastBaseline
	LayoutAttributeFirstBaseline
	LayoutAttributeLeftMargin
	LayoutAttributeRightMargin
	LayoutAttributeTopMargin
	LayoutAttributeBottomMargin
	LayoutAttributeLeadingMargin
	LayoutAttributeTrailingMargin
	LayoutAttributeCenterXWithinMargins
	LayoutAttributeCenterYWithinMargins
	// ios end

	LayoutAttributeBaseline LayoutAttribute = LayoutAttributeLastBaseline
)

// LayoutPriority used to indicate the relative importance of constraints allowing Auto Layout to make appropriate tradeoffs when satisfying the constraints of the system as a whole.
type LayoutPriority float32

const (
	LayoutPriorityRequired                   LayoutPriority = 1000 // a required constraint.  Do not exceed this.
	LayoutPriorityDefaultHigh                LayoutPriority = 750  // this is the priority level with which a button resists compressing its content.  Note that it is higher than LayoutPriorityWindowSizeStayPut.  Thus dragging to resize a window will not make buttons clip.  Rather the window frame is constrained.
	LayoutPriorityDragThatCanResizeWindow    LayoutPriority = 510  // This is the appropriate priority level for a drag that may end up resizing the window.  This needn't be a drag whose explicit purpose is to resize the window. The user might be dragging around window contents and it might be desirable that the window get bigger to accommodate.
	LayoutPriorityWindowSizeStayPut          LayoutPriority = 500  // This is the priority level at which the window prefers to stay the same size.  It's generally not appropriate to make a constraint at exactly this priority. You want to be higher or lower.
	LayoutPriorityDragThatCannotResizeWindow LayoutPriority = 490  // This is the priority level at which a split view divider say is dragged.  It won't resize the window.
	LayoutPriorityDefaultLow                 LayoutPriority = 250  // this is the priority level at which a button hugs its contents horizontally.
	LayoutPriorityFittingSizeCompression     LayoutPriority = 50   // When you issue -[NSView fittingSize] the smallest size that is large enough for the view's contents is computed.  This is the priority level with which the view wants to be as small as possible in that computation.  It's quite low.  It is generally not appropriate to make a constraint at exactly this priority.  You want to be higher or lower.
)

type LayoutRelation int

const (
	LayoutRelationLessThanOrEqual    LayoutRelation = -1
	LayoutRelationEqual              LayoutRelation = 0
	LayoutRelationGreaterThanOrEqual LayoutRelation = 1
)

type LayoutFormatOptions uint

const (
	LayoutFormatAlignAllLeft               LayoutFormatOptions = 1 << LayoutAttributeLeft
	LayoutFormatAlignAllRight              LayoutFormatOptions = 1 << LayoutAttributeRight
	LayoutFormatAlignAllTop                LayoutFormatOptions = 1 << LayoutAttributeTop
	LayoutFormatAlignAllBottom             LayoutFormatOptions = 1 << LayoutAttributeBottom
	LayoutFormatAlignAllLeading            LayoutFormatOptions = 1 << LayoutAttributeLeading
	LayoutFormatAlignAllTrailing           LayoutFormatOptions = 1 << LayoutAttributeTrailing
	LayoutFormatAlignAllCenterX            LayoutFormatOptions = 1 << LayoutAttributeCenterX
	LayoutFormatAlignAllCenterY            LayoutFormatOptions = 1 << LayoutAttributeCenterY
	LayoutFormatAlignAllLastBaseline       LayoutFormatOptions = 1 << LayoutAttributeLastBaseline
	LayoutFormatAlignAllFirstBaseline      LayoutFormatOptions = 1 << LayoutAttributeFirstBaseline
	LayoutFormatAlignAllBaseline           LayoutFormatOptions = LayoutFormatAlignAllLastBaseline
	LayoutFormatAlignmentMask              LayoutFormatOptions = 0xFFFF
	LayoutFormatDirectionLeadingToTrailing LayoutFormatOptions = 0 << 16 // default
	LayoutFormatDirectionLeftToRight       LayoutFormatOptions = 1 << 16
	LayoutFormatDirectionRightToLeft       LayoutFormatOptions = 2 << 16
	LayoutFormatDirectionMask              LayoutFormatOptions = 0x3 << 16
	LayoutFormatSpacingEdgeToEdge          LayoutFormatOptions = 0 << 19 // default
	LayoutFormatSpacingBaselineToBaseline  LayoutFormatOptions = 1 << 19
	LayoutFormatSpacingMask                LayoutFormatOptions = 0x1 << 19
)

type TickMarkPosition uint

const (
	TickMarkPositionBelow    TickMarkPosition = 0
	TickMarkPositionAbove    TickMarkPosition = 1
	TickMarkPositionLeading  TickMarkPosition = TickMarkPositionAbove
	TickMarkPositionTrailing TickMarkPosition = TickMarkPositionBelow
)

type LevelIndicatorStyle uint

const (
	LevelIndicatorStyleRelevancy LevelIndicatorStyle = iota
	LevelIndicatorStyleContinuousCapacity
	LevelIndicatorStyleDiscreteCapacity
	LevelIndicatorStyleRating
)

type PasteboardName string

type PasteboardType string

type RulerViewUnitName string

type SliderType uint

const (
	SliderTypeLinear   SliderType = 0
	SliderTypeCircular SliderType = 1
)

type SoundName string
type SoundPlaybackDeviceIdentifier string

// SplitViewDividerStyle constants that specify the style of the split view’s dividers.
type SplitViewDividerStyle int

const (
	SplitViewDividerStyleThick        SplitViewDividerStyle = 1
	SplitViewDividerStyleThin         SplitViewDividerStyle = 2
	SplitViewDividerStylePaneSplitter SplitViewDividerStyle = 3
)

// StackViewGravity is the gravity areas available in a stack view.
type StackViewGravity int

const (
	StackViewGravityTop      StackViewGravity = 1
	StackViewGravityLeading  StackViewGravity = 1
	StackViewGravityCenter   StackViewGravity = 2
	StackViewGravityBottom   StackViewGravity = 3
	StackViewGravityTrailing StackViewGravity = 3
)

const StackViewSpacingUseDefault = math.MaxFloat32

type StackViewDistribution int

const (
	StackViewDistributionGravityAreas StackViewDistribution = iota - 1
	StackViewDistributionFill
	StackViewDistributionFillEqually
	StackViewDistributionFillProportionally
	StackViewDistributionEqualSpacing
	StackViewDistributionEqualCentering
)

type StackViewVisibilityPriority float32

const (
	StackViewVisibilityPriorityMustHold              StackViewVisibilityPriority = 1000
	StackViewVisibilityPriorityDetachOnlyIfNecessary StackViewVisibilityPriority = 900
	StackViewVisibilityPriorityNotVisible            StackViewVisibilityPriority = 0
)

type StatusItemBehavior uint

const (
	StatusItemBehaviorRemovalAllowed       StatusItemBehavior = 1
	StatusItemBehaviorTerminationOnRemoval StatusItemBehavior = 1 << 2
)
const (
	SquareStatusItemLength   float64 = -2.0
	VariableStatusItemLength float64 = -1.0
)

type StoryboardSceneIdentifier string

type StoryboardName string

// TabState describe the current display state of a tab
type TabState uint

const (
	SelectedTab   TabState = 0
	BackgroundTab TabState = 1
	PressedTab    TabState = 2
)

type TabViewType uint

const (
	TopTabsBezelBorder    TabViewType = 0
	LeftTabsBezelBorder   TabViewType = 1
	BottomTabsBezelBorder TabViewType = 2
	RightTabsBezelBorder  TabViewType = 3
	NoTabsBezelBorder     TabViewType = 4
	NoTabsLineBorder      TabViewType = 5
	NoTabsNoBorder        TabViewType = 6
)

type TabPosition uint

const (
	TabPositionNone   TabPosition = 0
	TabPositionTop    TabPosition = 1
	TabPositionLeft   TabPosition = 2
	TabPositionBottom TabPosition = 3
	TabPositionRight  TabPosition = 4
)

type TabViewBorderType uint

const (
	TabViewBorderTypeNone  TabViewBorderType = 0
	TabViewBorderTypeLine  TabViewBorderType = 1
	TabViewBorderTypeBezel TabViewBorderType = 2
)

type TextInputSourceIdentifier string

type ToolbarItemIdentifier string
type ToolbarIdentifier string
type TitlebarSeparatorStyle int

type TouchBarItemIdentifier string
type TouchBarCustomizationIdentifier string

// BorderType specify the type of a view’s border.
type BorderType uint

const (
	NoBorder     BorderType = 0 // No border
	LineBorder   BorderType = 1 // A black line border around the view
	BezelBorder  BorderType = 2 // A concave border that makes the view look sunken
	GrooveBorder BorderType = 3 // A thin border that looks etched around the image
)

type AutoresizingMaskOptions uint

const (
	ViewNotSizable    AutoresizingMaskOptions = 0
	ViewMinXMargin    AutoresizingMaskOptions = 1
	ViewWidthSizable  AutoresizingMaskOptions = 2
	ViewMaxXMargin    AutoresizingMaskOptions = 4
	ViewMinYMargin    AutoresizingMaskOptions = 8
	ViewHeightSizable AutoresizingMaskOptions = 16
	ViewMaxYMargin    AutoresizingMaskOptions = 32
)

type ToolTipTag int
type TrackingRectTag int

type WindowTabbingIdentifier string

// WindowStyleMask specify the style of a window and can be combined using bitwise OR operator.
type WindowStyleMask uint

const (
	WindowStyleMaskBorderless                            WindowStyleMask = 0
	WindowStyleMaskTitled                                WindowStyleMask = 1 << 0
	WindowStyleMaskClosable                              WindowStyleMask = 1 << 1
	WindowStyleMaskMiniaturizable                        WindowStyleMask = 1 << 2
	WindowStyleMaskResizable                             WindowStyleMask = 1 << 3
	WindowStyleMaskUnifiedTitleAndToolbaWindowStyleMaskr WindowStyleMask = 1 << 12
	WindowStyleMaskFullScreeWindowStyleMaskn             WindowStyleMask = 1 << 14
	WindowStyleMaskFullSizeContentVieWindowStyleMaskw    WindowStyleMask = 1 << 15
	WindowStyleMaskUtilityWindow                         WindowStyleMask = 1 << 4
	WindowStyleMaskDocModalWindow                        WindowStyleMask = 1 << 6
	WindowStyleMaskNonactivatingPaneWindowStyleMaskl     WindowStyleMask = 1 << 7
	WindowStyleMaskHUDWindow                             WindowStyleMask = 1 << 13
)

// BackingStoreType specify how the drawing done in a window is buffered by the window device.
type BackingStoreType uint

const (
	BackingStoreRetained    BackingStoreType = 0
	BackingStoreNonretained BackingStoreType = 1
	BackingStoreBuffered    BackingStoreType = 2
)

type WindowFrameAutosaveName string
type WindowPersistableFrameDescriptor string

type WindowLevel int

const (
	NormalWindowLevel      WindowLevel = 0
	FloatingWindowLevel    WindowLevel = 3
	SubmenuWindowLevel     WindowLevel = 3
	TornOffMenuWindowLevel WindowLevel = 3
	MainMenuWindowLevel    WindowLevel = 24
	StatusWindowLevel      WindowLevel = 25
	ModalPanelWindowLevel  WindowLevel = 8
	PopUpMenuWindowLevel   WindowLevel = 101
	ScreenSaverWindowLevel WindowLevel = 1000
)

type WindowOrderingMode int

const (
	WindowAbove WindowOrderingMode = 1
	WindowBelow WindowOrderingMode = -1
	WindowOut   WindowOrderingMode = 0
)

type WindowButton uint

const (
	WindowCloseButton WindowButton = iota
	WindowMiniaturizeButton
	WindowZoomButton
	WindowToolbarButton
	WindowDocumentIconButton
	WindowDocumentVersionsButton
)

type DisplayGamut int

const (
	DisplayGamutSRGB DisplayGamut = 1
	DisplayGamutP3   DisplayGamut = 2
)

type WindowCollectionBehavior uint

const (
	WindowCollectionBehaviorDefault                   WindowCollectionBehavior = 0
	WindowCollectionBehaviorCanJoinAllSpaces          WindowCollectionBehavior = 1 << 0
	WindowCollectionBehaviorMoveToActiveSpace         WindowCollectionBehavior = 1 << 1
	WindowCollectionBehaviorManaged                   WindowCollectionBehavior = 1 << 2  // participates in spaces exposé.  Default behavior if windowLevel == NSNormalWindowLevel
	WindowCollectionBehaviorTransient                 WindowCollectionBehavior = 1 << 3  // floats in spaces hidden by exposé.  Default behavior if windowLevel != NSNormalWindowLevel
	WindowCollectionBehaviorStationary                WindowCollectionBehavior = 1 << 4  // unaffected by exposé.  Stays visible and stationary like desktop window
	WindowCollectionBehaviorParticipatesInCycle       WindowCollectionBehavior = 1 << 5  // default behavior if windowLevel == NSNormalWindowLevel
	WindowCollectionBehaviorIgnoresCycle              WindowCollectionBehavior = 1 << 6  // default behavior if windowLevel != NSNormalWindowLevel
	WindowCollectionBehaviorFullScreenPrimary         WindowCollectionBehavior = 1 << 7  // the frontmost window with this collection behavior will be the fullscreen window.
	WindowCollectionBehaviorFullScreenAuxiliary       WindowCollectionBehavior = 1 << 8  // windows with this collection behavior can be shown with the fullscreen window.
	WindowCollectionBehaviorFullScreenNone            WindowCollectionBehavior = 1 << 9  // The window can not be made fullscreen when this bit is set
	WindowCollectionBehaviorFullScreenAllowsTiling    WindowCollectionBehavior = 1 << 11 // This window can be a full screen tile window. It does not have to have FullScreenPrimary set.
	WindowCollectionBehaviorFullScreenDisallowsTiling WindowCollectionBehavior = 1 << 12 // This window can NOT be made a full screen tile window; it still may be allowed to be a regular FullScreenPrimary window.
)

type WindowSharingType uint

const (
	WindowSharingNone      WindowSharingType = 0 // Window contents may not be read by another process
	WindowSharingReadOnly  WindowSharingType = 1 // Window contents may be read but not modified by another process
	WindowSharingReadWrite WindowSharingType = 2 // Window contents may be read or modified by another process
)

type WindowDepth int32

const (
	WindowDepthTwentyfourBitRGB            WindowDepth = 0x208
	WindowDepthSixtyfourBitRGB             WindowDepth = 0x210
	WindowDepthOnehundredtwentyeightBitRGB WindowDepth = 0x220
)

type WindowOcclusionState uint

const (
	WindowOcclusionStateVisible WindowOcclusionState = 1 << 1
)

type WindowToolbarStyle int

const (
	WindowToolbarStyleAutomatic WindowTabbingMode = iota
	WindowToolbarStyleExpanded
	WindowToolbarStylePreference
	WindowToolbarStyleUnified
)

type WindowTabbingMode int

const (
	WindowTabbingModeAutomatic  WindowTabbingMode = iota // The system automatically prefers to tab this window when appropriate
	WindowTabbingModePreferred                           // The window explicitly should prefer to tab when shown
	WindowTabbingModeDisallowed                          // The window explicitly should not prefer to tab when shown
)

type SelectionDirection uint

const (
	DirectSelection SelectionDirection = iota
	SelectingNext
	SelectingPrevious
)

type WindowAnimationBehavior int

const (
	WindowAnimationBehaviorDefault        WindowAnimationBehavior = 0 // let AppKit infer animation behavior for this window
	WindowAnimationBehaviorNone           WindowAnimationBehavior = 2 // suppress inferred animations (don't animate)
	WindowAnimationBehaviorDocumentWindow WindowAnimationBehavior = 3
	WindowAnimationBehaviorUtilityWindow  WindowAnimationBehavior = 4
	WindowAnimationBehaviorAlertPanel     WindowAnimationBehavior = 5
)

type WindowTitleVisibility int

const (
	WindowTitleVisible WindowTitleVisibility = 0
	WindowTitleHidden  WindowTitleVisibility = 1
)

type MenuProperties uint

const (
	MenuPropertyItemTitle                    MenuProperties = 1 << 0 // the menu item's title
	MenuPropertyItemAttributedTitle          MenuProperties = 1 << 1 // the menu item's attributed title
	MenuPropertyItemKeyEquivalent            MenuProperties = 1 << 2 // the menu item's key equivalent
	MenuPropertyItemImage                    MenuProperties = 1 << 3 // the menu item's image
	MenuPropertyItemEnabled                  MenuProperties = 1 << 4 // whether the menu item is enabled or disabled
	MenuPropertyItemAccessibilityDescription MenuProperties = 1 << 5 // the menu item's accessibility description
)

type BitmapImageFileType uint

const (
	BitmapImageFileTypeTIFF BitmapImageFileType = iota
	BitmapImageFileTypeBMP
	BitmapImageFileTypeGIF
	BitmapImageFileTypeJPEG
	BitmapImageFileTypePNG
	BitmapImageFileTypeJPEG2000
)

type BitmapImageRepPropertyKey string

const (
	ImageColorSyncProfileData    BitmapImageRepPropertyKey = "NSImageColorSyncProfileData"
	ImageCompressionFactor       BitmapImageRepPropertyKey = "NSImageCompressionFactor"
	ImageCompressionMethod       BitmapImageRepPropertyKey = "NSImageCompressionMethod"
	ImageCurrentFrame            BitmapImageRepPropertyKey = "NSImageCurrentFrame"
	ImageCurrentFrameDuration    BitmapImageRepPropertyKey = "NSImageCurrentFrameDuration"
	ImageDitherTransparency      BitmapImageRepPropertyKey = "NSImageDitherTransparency"
	ImageEXIFData                BitmapImageRepPropertyKey = "NSImageEXIFData"
	ImageFallbackBackgroundColor BitmapImageRepPropertyKey = "NSImageFallbackBackgroundColor"
	ImageFrameCount              BitmapImageRepPropertyKey = "NSImageFrameCount"
	ImageGamma                   BitmapImageRepPropertyKey = "NSImageGamma"
	ImageInterlaced              BitmapImageRepPropertyKey = "NSImageInterlaced"
	ImageLoopCount               BitmapImageRepPropertyKey = "NSImageLoopCount"
	ImageProgressive             BitmapImageRepPropertyKey = "NSImageProgressive"
	ImageRGBColorTable           BitmapImageRepPropertyKey = "NSImageRGBColorTable"
)

type ColorSpaceName string

const (
	CalibratedRGBColorSpaceName   ColorSpaceName = "NSCalibratedRGBColorSpace"
	CalibratedWhiteColorSpaceName ColorSpaceName = "NSCalibratedWhiteColorSpace"
	CustomColorSpaceName          ColorSpaceName = "NSCustomColorSpace"
	DeviceCMYKColorSpaceName      ColorSpaceName = "NSDeviceCMYKColorSpace"
	DeviceRGBColorSpaceName       ColorSpaceName = "NSDeviceRGBColorSpace"
	DeviceWhiteColorSpaceName     ColorSpaceName = "NSDeviceWhiteColorSpace"
	NamedColorSpaceName           ColorSpaceName = "NSNamedColorSpace"
	PatternColorSpaceName         ColorSpaceName = "NSPatternColorSpace"
)

type AboutPanelOptionKey string
type PrintInfoAttributeKey string
type ImageHintKey string
type GraphicsContextAttributeKey string
