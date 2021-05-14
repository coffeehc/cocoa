package appkit

import "github.com/hsiafan/cocoa/foundation"

// WindowStyleMask specify the style of a window, and can be combined using bitwise OR operator.
type WindowStyleMask uint

const (
	WindowStyleMaskBorderless     WindowStyleMask = 0
	WindowStyleMaskTitled         WindowStyleMask = 1 << 0
	WindowStyleMaskClosable       WindowStyleMask = 1 << 1
	WindowStyleMaskMiniaturizable WindowStyleMask = 1 << 2
	WindowStyleMaskResizable      WindowStyleMask = 1 << 3

	WindowStyleMaskUnifiedTitleAndToolbaWindowStyleMaskr = 1 << 12

	WindowStyleMaskFullScreeWindowStyleMaskn = 1 << 14

	WindowStyleMaskFullSizeContentVieWindowStyleMaskw = 1 << 15

	WindowStyleMaskUtilityWindow                     WindowStyleMask = 1 << 4
	WindowStyleMaskDocModalWindow                    WindowStyleMask = 1 << 6
	WindowStyleMaskNonactivatingPaneWindowStyleMaskl                 = 1 << 7
	WindowStyleMaskHUDWindow                         WindowStyleMask = 1 << 13
)

// BackingStoreType specify how the drawing done in a window is buffered by the window device.
type BackingStoreType uint

const (
	BackingStoreRetained    BackingStoreType = 0
	BackingStoreNonretained BackingStoreType = 1
	BackingStoreBuffered    BackingStoreType = 2
)

// NewPlainWindow create a common window with close/minimize buttons
func NewPlainWindow(rect foundation.Rect) Window {
	return AllocWindow().InitWithContentRect_StyleMask_Backing_Defer(
		rect,
		WindowStyleMaskTitled|WindowStyleMaskClosable|WindowStyleMaskResizable|WindowStyleMaskMiniaturizable,
		BackingStoreBuffered,
		false,
	)
}

type WindowFrameAutosaveName string
type WindowPersistableFrameDescriptor string

type WindowLevel int

type WindowOrderingMode int

type WindowButton uint

type DisplayGamut int

type WindowCollectionBehavior uint

type WindowSharingType uint

type WindowDepth int32

type WindowOcclusionState uint

type WindowToolbarStyle int

type WindowTabbingMode int

type SelectionDirection uint

type WindowAnimationBehavior int

type WindowTitleVisibility int
