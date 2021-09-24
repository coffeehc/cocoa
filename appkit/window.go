package appkit

// #include "window.h"
import "C"
import (
	"github.com/hsiafan/cocoa/coregraphics"
	"github.com/hsiafan/cocoa/foundation"
	"github.com/hsiafan/cocoa/objc"
	"runtime/cgo"
	"unsafe"
)

type Window interface {
	Responder
	ToggleFullScreen(sender objc.Object)
	InvalidateShadow()
	AutorecalculatesContentBorderThicknessForEdge(edge foundation.RectEdge) bool
	SetAutorecalculatesContentBorderThickness_ForEdge(flag bool, edge foundation.RectEdge)
	ContentBorderThicknessForEdge(edge foundation.RectEdge) coregraphics.Float
	SetContentBorderThickness_ForEdge(thickness coregraphics.Float, edge foundation.RectEdge)
	ContentRectForFrameRect(frameRect foundation.Rect) foundation.Rect
	FrameRectForContentRect(contentRect foundation.Rect) foundation.Rect
	EndSheet(sheetWindow Window)
	EndSheet_ReturnCode(sheetWindow Window, returnCode ModalResponse)
	SetFrameOrigin(point foundation.Point)
	SetFrameTopLeftPoint(point foundation.Point)
	ConstrainFrameRect_ToScreen(frameRect foundation.Rect, screen Screen) foundation.Rect
	CascadeTopLeftFromPoint(topLeftPoint foundation.Point) foundation.Point
	SetFrame_Display(frameRect foundation.Rect, flag bool)
	SetFrame_Display_Animate(frameRect foundation.Rect, displayFlag bool, animateFlag bool)
	AnimationResizeTime(newFrame foundation.Rect) foundation.TimeInterval
	PerformZoom(sender objc.Object)
	Zoom(sender objc.Object)
	SetContentSize(size foundation.Size)
	OrderOut(sender objc.Object)
	OrderBack(sender objc.Object)
	OrderFront(sender objc.Object)
	OrderFrontRegardless()
	OrderWindow_RelativeTo(place WindowOrderingMode, otherWin int)
	SetFrameUsingName(name WindowFrameAutosaveName) bool
	SetFrameUsingName_Force(name WindowFrameAutosaveName, force bool) bool
	SaveFrameUsingName(name WindowFrameAutosaveName)
	SetFrameFromString(_string WindowPersistableFrameDescriptor)
	MakeKeyWindow()
	MakeKeyAndOrderFront(sender objc.Object)
	BecomeKeyWindow()
	ResignKeyWindow()
	MakeMainWindow()
	BecomeMainWindow()
	ResignMainWindow()
	ToggleToolbarShown(sender objc.Object)
	RunToolbarCustomizationPalette(sender objc.Object)
	AddChildWindow_Ordered(childWin Window, place WindowOrderingMode)
	RemoveChildWindow(childWin Window)
	EnableKeyEquivalentForDefaultButtonCell()
	DisableKeyEquivalentForDefaultButtonCell()
	FieldEditor_ForObject(createFlag bool, object objc.Object) Text
	EndEditingFor(object objc.Object)
	EnableCursorRects()
	DisableCursorRects()
	DiscardCursorRects()
	InvalidateCursorRectsForView(view View)
	ResetCursorRects()
	StandardWindowButton(b WindowButton) Button
	AddTitlebarAccessoryViewController(childViewController TitlebarAccessoryViewController)
	InsertTitlebarAccessoryViewController_AtIndex(childViewController TitlebarAccessoryViewController, index int)
	RemoveTitlebarAccessoryViewControllerAtIndex(index int)
	AddTabbedWindow_Ordered(window Window, ordered WindowOrderingMode)
	SelectNextTab(sender objc.Object)
	SelectPreviousTab(sender objc.Object)
	MoveTabToNewWindow(sender objc.Object)
	ToggleTabBar(sender objc.Object)
	ToggleTabOverview(sender objc.Object)
	PostEvent_AtStart(event Event, flag bool)
	SendEvent(event Event)
	MakeFirstResponder(responder Responder) bool
	SelectKeyViewPrecedingView(view View)
	SelectKeyViewFollowingView(view View)
	SelectPreviousKeyView(sender objc.Object)
	SelectNextKeyView(sender objc.Object)
	RecalculateKeyViewLoop()
	PerformWindowDragWithEvent(event Event)
	DisableSnapshotRestoration()
	EnableSnapshotRestoration()
	Display()
	DisplayIfNeeded()
	DisableScreenUpdatesUntilFlush()
	Update()
	DragImage_At_Offset_Event_Pasteboard_Source_SlideBack(image Image, baseLocation foundation.Point, initialOffset foundation.Size, event Event, pboard Pasteboard, sourceObj objc.Object, slideFlag bool)
	RegisterForDraggedTypes(newTypes []PasteboardType)
	UnregisterDraggedTypes()
	ConvertRectFromBacking(rect foundation.Rect) foundation.Rect
	ConvertRectToBacking(rect foundation.Rect) foundation.Rect
	ConvertRectToScreen(rect foundation.Rect) foundation.Rect
	ConvertRectFromScreen(rect foundation.Rect) foundation.Rect
	SetTitleWithRepresentedFilename(filename string)
	Center()
	PerformClose(sender objc.Object)
	Close()
	PerformMiniaturize(sender objc.Object)
	Miniaturize(sender objc.Object)
	Deminiaturize(sender objc.Object)
	Print(sender objc.Object)
	DataWithEPSInsideRect(rect foundation.Rect) []byte
	DataWithPDFInsideRect(rect foundation.Rect) []byte
	UpdateConstraintsIfNeeded()
	LayoutIfNeeded()
	VisualizeConstraints(constraints []LayoutConstraint)
	AnchorAttributeForOrientation(orientation LayoutConstraintOrientation) LayoutAttribute
	SetAnchorAttribute_ForOrientation(attr LayoutAttribute, orientation LayoutConstraintOrientation)
	SetIsMiniaturized(flag bool)
	SetIsVisible(flag bool)
	SetIsZoomed(flag bool)
	HandleCloseScriptCommand(command foundation.CloseCommand) objc.Object
	HandlePrintScriptCommand(command foundation.ScriptCommand) objc.Object
	HandleSaveScriptCommand(command foundation.ScriptCommand) objc.Object
	CanRepresentDisplayGamut(displayGamut DisplayGamut) bool
	ConvertPointFromScreen(point foundation.Point) foundation.Point
	ConvertPointToScreen(point foundation.Point) foundation.Point
	ConvertPointFromBacking(point foundation.Point) foundation.Point
	ConvertPointToBacking(point foundation.Point) foundation.Point
	MergeAllWindows(sender objc.Object)
	SetDynamicDepthLimit(flag bool)
	SetFrameAutosaveName(name WindowFrameAutosaveName) bool
	Delegate() objc.Object
	SetDelegate(value objc.Object)
	ContentViewController() ViewController
	SetContentViewController(value ViewController)
	ContentView() View
	SetContentView(value View)
	WorksWhenModal() bool
	AlphaValue() coregraphics.Float
	SetAlphaValue(value coregraphics.Float)
	BackgroundColor() Color
	SetBackgroundColor(value Color)
	ColorSpace() ColorSpace
	SetColorSpace(value ColorSpace)
	CanHide() bool
	SetCanHide(value bool)
	IsOnActiveSpace() bool
	HidesOnDeactivate() bool
	SetHidesOnDeactivate(value bool)
	CollectionBehavior() WindowCollectionBehavior
	SetCollectionBehavior(value WindowCollectionBehavior)
	IsOpaque() bool
	SetOpaque(value bool)
	HasShadow() bool
	SetHasShadow(value bool)
	PreventsApplicationTerminationWhenModal() bool
	SetPreventsApplicationTerminationWhenModal(value bool)
	WindowNumber() int
	DeviceDescription() map[DeviceDescriptionKey]objc.Object
	CanBecomeVisibleWithoutLogin() bool
	SetCanBecomeVisibleWithoutLogin(value bool)
	SharingType() WindowSharingType
	SetSharingType(value WindowSharingType)
	BackingType() BackingStoreType
	SetBackingType(value BackingStoreType)
	DepthLimit() WindowDepth
	SetDepthLimit(value WindowDepth)
	HasDynamicDepthLimit() bool
	WindowController() WindowController
	SetWindowController(value WindowController)
	AttachedSheet() Window
	IsSheet() bool
	SheetParent() Window
	Sheets() []Window
	Frame() foundation.Rect
	AspectRatio() foundation.Size
	SetAspectRatio(value foundation.Size)
	MinSize() foundation.Size
	SetMinSize(value foundation.Size)
	MaxSize() foundation.Size
	SetMaxSize(value foundation.Size)
	IsZoomed() bool
	ResizeFlags() EventModifierFlags
	ResizeIncrements() foundation.Size
	SetResizeIncrements(value foundation.Size)
	PreservesContentDuringLiveResize() bool
	SetPreservesContentDuringLiveResize(value bool)
	InLiveResize() bool
	ContentAspectRatio() foundation.Size
	SetContentAspectRatio(value foundation.Size)
	ContentMinSize() foundation.Size
	SetContentMinSize(value foundation.Size)
	ContentMaxSize() foundation.Size
	SetContentMaxSize(value foundation.Size)
	ContentResizeIncrements() foundation.Size
	SetContentResizeIncrements(value foundation.Size)
	ContentLayoutGuide() objc.Object
	ContentLayoutRect() foundation.Rect
	MaxFullScreenContentSize() foundation.Size
	SetMaxFullScreenContentSize(value foundation.Size)
	MinFullScreenContentSize() foundation.Size
	SetMinFullScreenContentSize(value foundation.Size)
	Level() WindowLevel
	SetLevel(value WindowLevel)
	IsVisible() bool
	OcclusionState() WindowOcclusionState
	FrameAutosaveName() WindowFrameAutosaveName
	StringWithSavedFrame() WindowPersistableFrameDescriptor
	IsKeyWindow() bool
	CanBecomeKeyWindow() bool
	IsMainWindow() bool
	CanBecomeMainWindow() bool
	Toolbar() Toolbar
	SetToolbar(value Toolbar)
	ChildWindows() []Window
	ParentWindow() Window
	SetParentWindow(value Window)
	DefaultButtonCell() ButtonCell
	SetDefaultButtonCell(value ButtonCell)
	IsExcludedFromWindowsMenu() bool
	SetExcludedFromWindowsMenu(value bool)
	AreCursorRectsEnabled() bool
	ShowsToolbarButton() bool
	SetShowsToolbarButton(value bool)
	TitlebarAppearsTransparent() bool
	SetTitlebarAppearsTransparent(value bool)
	ToolbarStyle() WindowToolbarStyle
	SetToolbarStyle(value WindowToolbarStyle)
	TitlebarSeparatorStyle() TitlebarSeparatorStyle
	SetTitlebarSeparatorStyle(value TitlebarSeparatorStyle)
	TitlebarAccessoryViewControllers() []TitlebarAccessoryViewController
	SetTitlebarAccessoryViewControllers(value []TitlebarAccessoryViewController)
	Tab() WindowTab
	TabbingIdentifier() WindowTabbingIdentifier
	SetTabbingIdentifier(value WindowTabbingIdentifier)
	TabbingMode() WindowTabbingMode
	SetTabbingMode(value WindowTabbingMode)
	TabbedWindows() []Window
	TabGroup() WindowTabGroup
	AllowsToolTipsWhenApplicationIsInactive() bool
	SetAllowsToolTipsWhenApplicationIsInactive(value bool)
	CurrentEvent() Event
	InitialFirstResponder() View
	SetInitialFirstResponder(value View)
	FirstResponder() Responder
	KeyViewSelectionDirection() SelectionDirection
	AutorecalculatesKeyViewLoop() bool
	SetAutorecalculatesKeyViewLoop(value bool)
	AcceptsMouseMovedEvents() bool
	SetAcceptsMouseMovedEvents(value bool)
	IgnoresMouseEvents() bool
	SetIgnoresMouseEvents(value bool)
	MouseLocationOutsideOfEventStream() foundation.Point
	IsRestorable() bool
	SetRestorable(value bool)
	ViewsNeedDisplay() bool
	SetViewsNeedDisplay(value bool)
	AllowsConcurrentViewDrawing() bool
	SetAllowsConcurrentViewDrawing(value bool)
	AnimationBehavior() WindowAnimationBehavior
	SetAnimationBehavior(value WindowAnimationBehavior)
	IsDocumentEdited() bool
	SetDocumentEdited(value bool)
	BackingScaleFactor() coregraphics.Float
	Title() string
	SetTitle(value string)
	Subtitle() string
	SetSubtitle(value string)
	TitleVisibility() WindowTitleVisibility
	SetTitleVisibility(value WindowTitleVisibility)
	RepresentedFilename() string
	SetRepresentedFilename(value string)
	RepresentedURL() foundation.URL
	SetRepresentedURL(value foundation.URL)
	Screen() Screen
	DeepestScreen() Screen
	DisplaysWhenScreenProfileChanges() bool
	SetDisplaysWhenScreenProfileChanges(value bool)
	IsMovableByWindowBackground() bool
	SetMovableByWindowBackground(value bool)
	IsMovable() bool
	SetMovable(value bool)
	IsReleasedWhenClosed() bool
	SetReleasedWhenClosed(value bool)
	IsMiniaturized() bool
	MiniwindowImage() Image
	SetMiniwindowImage(value Image)
	MiniwindowTitle() string
	SetMiniwindowTitle(value string)
	DockTile() DockTile
	HasCloseBox() bool
	HasTitleBar() bool
	OrderedIndex() int
	SetOrderedIndex(value int)
	AppearanceSource() objc.Object
	SetAppearanceSource(value objc.Object)
	IsFloatingPanel() bool
	IsMiniaturizable() bool
	IsModalPanel() bool
	IsResizable() bool
	StyleMask() WindowStyleMask
	SetStyleMask(value WindowStyleMask)
	WindowTitlebarLayoutDirection() UserInterfaceLayoutDirection
	IsZoomable() bool
}

type NSWindow struct {
	NSResponder
}

func MakeWindow(ptr unsafe.Pointer) NSWindow {
	return NSWindow{
		NSResponder: MakeResponder(ptr),
	}
}

func WindowWithContentViewController(contentViewController ViewController) NSWindow {
	result_ := C.C_NSWindow_WindowWithContentViewController(objc.ExtractPtr(contentViewController))
	return MakeWindow(result_)
}

func (n NSWindow) InitWithContentRect_StyleMask_Backing_Defer(contentRect foundation.Rect, style WindowStyleMask, backingStoreType BackingStoreType, flag bool) NSWindow {
	result_ := C.C_NSWindow_InitWithContentRect_StyleMask_Backing_Defer(n.Ptr(), *(*C.CGRect)(coregraphics.ToCGRectPointer(coregraphics.Rect(contentRect))), C.uint(uint(style)), C.uint(uint(backingStoreType)), C.bool(flag))
	return MakeWindow(result_)
}

func (n NSWindow) InitWithContentRect_StyleMask_Backing_Defer_Screen(contentRect foundation.Rect, style WindowStyleMask, backingStoreType BackingStoreType, flag bool, screen Screen) NSWindow {
	result_ := C.C_NSWindow_InitWithContentRect_StyleMask_Backing_Defer_Screen(n.Ptr(), *(*C.CGRect)(coregraphics.ToCGRectPointer(coregraphics.Rect(contentRect))), C.uint(uint(style)), C.uint(uint(backingStoreType)), C.bool(flag), objc.ExtractPtr(screen))
	return MakeWindow(result_)
}

func (n NSWindow) Init() NSWindow {
	result_ := C.C_NSWindow_Init(n.Ptr())
	return MakeWindow(result_)
}

func AllocWindow() NSWindow {
	result_ := C.C_NSWindow_AllocWindow()
	return MakeWindow(result_)
}

func NewWindow() NSWindow {
	result_ := C.C_NSWindow_NewWindow()
	return MakeWindow(result_)
}

func (n NSWindow) Autorelease() NSWindow {
	result_ := C.C_NSWindow_Autorelease(n.Ptr())
	return MakeWindow(result_)
}

func (n NSWindow) Retain() NSWindow {
	result_ := C.C_NSWindow_Retain(n.Ptr())
	return MakeWindow(result_)
}

func (n NSWindow) ToggleFullScreen(sender objc.Object) {
	C.C_NSWindow_ToggleFullScreen(n.Ptr(), objc.ExtractPtr(sender))
}

func (n NSWindow) InvalidateShadow() {
	C.C_NSWindow_InvalidateShadow(n.Ptr())
}

func (n NSWindow) AutorecalculatesContentBorderThicknessForEdge(edge foundation.RectEdge) bool {
	result_ := C.C_NSWindow_AutorecalculatesContentBorderThicknessForEdge(n.Ptr(), C.uint(uint(edge)))
	return bool(result_)
}

func (n NSWindow) SetAutorecalculatesContentBorderThickness_ForEdge(flag bool, edge foundation.RectEdge) {
	C.C_NSWindow_SetAutorecalculatesContentBorderThickness_ForEdge(n.Ptr(), C.bool(flag), C.uint(uint(edge)))
}

func (n NSWindow) ContentBorderThicknessForEdge(edge foundation.RectEdge) coregraphics.Float {
	result_ := C.C_NSWindow_ContentBorderThicknessForEdge(n.Ptr(), C.uint(uint(edge)))
	return coregraphics.Float(float64(result_))
}

func (n NSWindow) SetContentBorderThickness_ForEdge(thickness coregraphics.Float, edge foundation.RectEdge) {
	C.C_NSWindow_SetContentBorderThickness_ForEdge(n.Ptr(), C.double(float64(thickness)), C.uint(uint(edge)))
}

func WindowNumbersWithOptions(options WindowNumberListOptions) []foundation.Number {
	result_ := C.C_NSWindow_WindowNumbersWithOptions(C.uint(uint(options)))
	if result_.len > 0 {
		defer C.free(result_.data)
	}
	result_Slice := unsafe.Slice((*unsafe.Pointer)(result_.data), int(result_.len))
	var goResult_ = make([]foundation.Number, len(result_Slice))
	for idx, r := range result_Slice {
		goResult_[idx] = foundation.MakeNumber(r)
	}
	return goResult_
}

func Window_ContentRectForFrameRect_StyleMask(fRect foundation.Rect, style WindowStyleMask) foundation.Rect {
	result_ := C.C_NSWindow_Window_ContentRectForFrameRect_StyleMask(*(*C.CGRect)(coregraphics.ToCGRectPointer(coregraphics.Rect(fRect))), C.uint(uint(style)))
	return foundation.Rect(coregraphics.FromCGRectPointer(unsafe.Pointer(&result_)))
}

func Window_FrameRectForContentRect_StyleMask(cRect foundation.Rect, style WindowStyleMask) foundation.Rect {
	result_ := C.C_NSWindow_Window_FrameRectForContentRect_StyleMask(*(*C.CGRect)(coregraphics.ToCGRectPointer(coregraphics.Rect(cRect))), C.uint(uint(style)))
	return foundation.Rect(coregraphics.FromCGRectPointer(unsafe.Pointer(&result_)))
}

func Window_MinFrameWidthWithTitle_StyleMask(title string, style WindowStyleMask) coregraphics.Float {
	result_ := C.C_NSWindow_Window_MinFrameWidthWithTitle_StyleMask(foundation.NewString(title).Ptr(), C.uint(uint(style)))
	return coregraphics.Float(float64(result_))
}

func (n NSWindow) ContentRectForFrameRect(frameRect foundation.Rect) foundation.Rect {
	result_ := C.C_NSWindow_ContentRectForFrameRect(n.Ptr(), *(*C.CGRect)(coregraphics.ToCGRectPointer(coregraphics.Rect(frameRect))))
	return foundation.Rect(coregraphics.FromCGRectPointer(unsafe.Pointer(&result_)))
}

func (n NSWindow) FrameRectForContentRect(contentRect foundation.Rect) foundation.Rect {
	result_ := C.C_NSWindow_FrameRectForContentRect(n.Ptr(), *(*C.CGRect)(coregraphics.ToCGRectPointer(coregraphics.Rect(contentRect))))
	return foundation.Rect(coregraphics.FromCGRectPointer(unsafe.Pointer(&result_)))
}

func (n NSWindow) EndSheet(sheetWindow Window) {
	C.C_NSWindow_EndSheet(n.Ptr(), objc.ExtractPtr(sheetWindow))
}

func (n NSWindow) EndSheet_ReturnCode(sheetWindow Window, returnCode ModalResponse) {
	C.C_NSWindow_EndSheet_ReturnCode(n.Ptr(), objc.ExtractPtr(sheetWindow), C.int(int(returnCode)))
}

func (n NSWindow) SetFrameOrigin(point foundation.Point) {
	C.C_NSWindow_SetFrameOrigin(n.Ptr(), *(*C.CGPoint)(coregraphics.ToCGPointPointer(coregraphics.Point(point))))
}

func (n NSWindow) SetFrameTopLeftPoint(point foundation.Point) {
	C.C_NSWindow_SetFrameTopLeftPoint(n.Ptr(), *(*C.CGPoint)(coregraphics.ToCGPointPointer(coregraphics.Point(point))))
}

func (n NSWindow) ConstrainFrameRect_ToScreen(frameRect foundation.Rect, screen Screen) foundation.Rect {
	result_ := C.C_NSWindow_ConstrainFrameRect_ToScreen(n.Ptr(), *(*C.CGRect)(coregraphics.ToCGRectPointer(coregraphics.Rect(frameRect))), objc.ExtractPtr(screen))
	return foundation.Rect(coregraphics.FromCGRectPointer(unsafe.Pointer(&result_)))
}

func (n NSWindow) CascadeTopLeftFromPoint(topLeftPoint foundation.Point) foundation.Point {
	result_ := C.C_NSWindow_CascadeTopLeftFromPoint(n.Ptr(), *(*C.CGPoint)(coregraphics.ToCGPointPointer(coregraphics.Point(topLeftPoint))))
	return foundation.Point(coregraphics.FromCGPointPointer(unsafe.Pointer(&result_)))
}

func (n NSWindow) SetFrame_Display(frameRect foundation.Rect, flag bool) {
	C.C_NSWindow_SetFrame_Display(n.Ptr(), *(*C.CGRect)(coregraphics.ToCGRectPointer(coregraphics.Rect(frameRect))), C.bool(flag))
}

func (n NSWindow) SetFrame_Display_Animate(frameRect foundation.Rect, displayFlag bool, animateFlag bool) {
	C.C_NSWindow_SetFrame_Display_Animate(n.Ptr(), *(*C.CGRect)(coregraphics.ToCGRectPointer(coregraphics.Rect(frameRect))), C.bool(displayFlag), C.bool(animateFlag))
}

func (n NSWindow) AnimationResizeTime(newFrame foundation.Rect) foundation.TimeInterval {
	result_ := C.C_NSWindow_AnimationResizeTime(n.Ptr(), *(*C.CGRect)(coregraphics.ToCGRectPointer(coregraphics.Rect(newFrame))))
	return foundation.TimeInterval(float64(result_))
}

func (n NSWindow) PerformZoom(sender objc.Object) {
	C.C_NSWindow_PerformZoom(n.Ptr(), objc.ExtractPtr(sender))
}

func (n NSWindow) Zoom(sender objc.Object) {
	C.C_NSWindow_Zoom(n.Ptr(), objc.ExtractPtr(sender))
}

func (n NSWindow) SetContentSize(size foundation.Size) {
	C.C_NSWindow_SetContentSize(n.Ptr(), *(*C.CGSize)(coregraphics.ToCGSizePointer(coregraphics.Size(size))))
}

func (n NSWindow) OrderOut(sender objc.Object) {
	C.C_NSWindow_OrderOut(n.Ptr(), objc.ExtractPtr(sender))
}

func (n NSWindow) OrderBack(sender objc.Object) {
	C.C_NSWindow_OrderBack(n.Ptr(), objc.ExtractPtr(sender))
}

func (n NSWindow) OrderFront(sender objc.Object) {
	C.C_NSWindow_OrderFront(n.Ptr(), objc.ExtractPtr(sender))
}

func (n NSWindow) OrderFrontRegardless() {
	C.C_NSWindow_OrderFrontRegardless(n.Ptr())
}

func (n NSWindow) OrderWindow_RelativeTo(place WindowOrderingMode, otherWin int) {
	C.C_NSWindow_OrderWindow_RelativeTo(n.Ptr(), C.int(int(place)), C.int(otherWin))
}

func Window_RemoveFrameUsingName(name WindowFrameAutosaveName) {
	C.C_NSWindow_Window_RemoveFrameUsingName(foundation.NewString(string(name)).Ptr())
}

func (n NSWindow) SetFrameUsingName(name WindowFrameAutosaveName) bool {
	result_ := C.C_NSWindow_SetFrameUsingName(n.Ptr(), foundation.NewString(string(name)).Ptr())
	return bool(result_)
}

func (n NSWindow) SetFrameUsingName_Force(name WindowFrameAutosaveName, force bool) bool {
	result_ := C.C_NSWindow_SetFrameUsingName_Force(n.Ptr(), foundation.NewString(string(name)).Ptr(), C.bool(force))
	return bool(result_)
}

func (n NSWindow) SaveFrameUsingName(name WindowFrameAutosaveName) {
	C.C_NSWindow_SaveFrameUsingName(n.Ptr(), foundation.NewString(string(name)).Ptr())
}

func (n NSWindow) SetFrameFromString(_string WindowPersistableFrameDescriptor) {
	C.C_NSWindow_SetFrameFromString(n.Ptr(), foundation.NewString(string(_string)).Ptr())
}

func (n NSWindow) MakeKeyWindow() {
	C.C_NSWindow_MakeKeyWindow(n.Ptr())
}

func (n NSWindow) MakeKeyAndOrderFront(sender objc.Object) {
	C.C_NSWindow_MakeKeyAndOrderFront(n.Ptr(), objc.ExtractPtr(sender))
}

func (n NSWindow) BecomeKeyWindow() {
	C.C_NSWindow_BecomeKeyWindow(n.Ptr())
}

func (n NSWindow) ResignKeyWindow() {
	C.C_NSWindow_ResignKeyWindow(n.Ptr())
}

func (n NSWindow) MakeMainWindow() {
	C.C_NSWindow_MakeMainWindow(n.Ptr())
}

func (n NSWindow) BecomeMainWindow() {
	C.C_NSWindow_BecomeMainWindow(n.Ptr())
}

func (n NSWindow) ResignMainWindow() {
	C.C_NSWindow_ResignMainWindow(n.Ptr())
}

func (n NSWindow) ToggleToolbarShown(sender objc.Object) {
	C.C_NSWindow_ToggleToolbarShown(n.Ptr(), objc.ExtractPtr(sender))
}

func (n NSWindow) RunToolbarCustomizationPalette(sender objc.Object) {
	C.C_NSWindow_RunToolbarCustomizationPalette(n.Ptr(), objc.ExtractPtr(sender))
}

func (n NSWindow) AddChildWindow_Ordered(childWin Window, place WindowOrderingMode) {
	C.C_NSWindow_AddChildWindow_Ordered(n.Ptr(), objc.ExtractPtr(childWin), C.int(int(place)))
}

func (n NSWindow) RemoveChildWindow(childWin Window) {
	C.C_NSWindow_RemoveChildWindow(n.Ptr(), objc.ExtractPtr(childWin))
}

func (n NSWindow) EnableKeyEquivalentForDefaultButtonCell() {
	C.C_NSWindow_EnableKeyEquivalentForDefaultButtonCell(n.Ptr())
}

func (n NSWindow) DisableKeyEquivalentForDefaultButtonCell() {
	C.C_NSWindow_DisableKeyEquivalentForDefaultButtonCell(n.Ptr())
}

func (n NSWindow) FieldEditor_ForObject(createFlag bool, object objc.Object) Text {
	result_ := C.C_NSWindow_FieldEditor_ForObject(n.Ptr(), C.bool(createFlag), objc.ExtractPtr(object))
	return MakeText(result_)
}

func (n NSWindow) EndEditingFor(object objc.Object) {
	C.C_NSWindow_EndEditingFor(n.Ptr(), objc.ExtractPtr(object))
}

func (n NSWindow) EnableCursorRects() {
	C.C_NSWindow_EnableCursorRects(n.Ptr())
}

func (n NSWindow) DisableCursorRects() {
	C.C_NSWindow_DisableCursorRects(n.Ptr())
}

func (n NSWindow) DiscardCursorRects() {
	C.C_NSWindow_DiscardCursorRects(n.Ptr())
}

func (n NSWindow) InvalidateCursorRectsForView(view View) {
	C.C_NSWindow_InvalidateCursorRectsForView(n.Ptr(), objc.ExtractPtr(view))
}

func (n NSWindow) ResetCursorRects() {
	C.C_NSWindow_ResetCursorRects(n.Ptr())
}

func StandardWindowButton_ForStyleMask(b WindowButton, styleMask WindowStyleMask) Button {
	result_ := C.C_NSWindow_StandardWindowButton_ForStyleMask(C.uint(uint(b)), C.uint(uint(styleMask)))
	return MakeButton(result_)
}

func (n NSWindow) StandardWindowButton(b WindowButton) Button {
	result_ := C.C_NSWindow_StandardWindowButton(n.Ptr(), C.uint(uint(b)))
	return MakeButton(result_)
}

func (n NSWindow) AddTitlebarAccessoryViewController(childViewController TitlebarAccessoryViewController) {
	C.C_NSWindow_AddTitlebarAccessoryViewController(n.Ptr(), objc.ExtractPtr(childViewController))
}

func (n NSWindow) InsertTitlebarAccessoryViewController_AtIndex(childViewController TitlebarAccessoryViewController, index int) {
	C.C_NSWindow_InsertTitlebarAccessoryViewController_AtIndex(n.Ptr(), objc.ExtractPtr(childViewController), C.int(index))
}

func (n NSWindow) RemoveTitlebarAccessoryViewControllerAtIndex(index int) {
	C.C_NSWindow_RemoveTitlebarAccessoryViewControllerAtIndex(n.Ptr(), C.int(index))
}

func (n NSWindow) AddTabbedWindow_Ordered(window Window, ordered WindowOrderingMode) {
	C.C_NSWindow_AddTabbedWindow_Ordered(n.Ptr(), objc.ExtractPtr(window), C.int(int(ordered)))
}

func (n NSWindow) SelectNextTab(sender objc.Object) {
	C.C_NSWindow_SelectNextTab(n.Ptr(), objc.ExtractPtr(sender))
}

func (n NSWindow) SelectPreviousTab(sender objc.Object) {
	C.C_NSWindow_SelectPreviousTab(n.Ptr(), objc.ExtractPtr(sender))
}

func (n NSWindow) MoveTabToNewWindow(sender objc.Object) {
	C.C_NSWindow_MoveTabToNewWindow(n.Ptr(), objc.ExtractPtr(sender))
}

func (n NSWindow) ToggleTabBar(sender objc.Object) {
	C.C_NSWindow_ToggleTabBar(n.Ptr(), objc.ExtractPtr(sender))
}

func (n NSWindow) ToggleTabOverview(sender objc.Object) {
	C.C_NSWindow_ToggleTabOverview(n.Ptr(), objc.ExtractPtr(sender))
}

func (n NSWindow) PostEvent_AtStart(event Event, flag bool) {
	C.C_NSWindow_PostEvent_AtStart(n.Ptr(), objc.ExtractPtr(event), C.bool(flag))
}

func (n NSWindow) SendEvent(event Event) {
	C.C_NSWindow_SendEvent(n.Ptr(), objc.ExtractPtr(event))
}

func (n NSWindow) MakeFirstResponder(responder Responder) bool {
	result_ := C.C_NSWindow_MakeFirstResponder(n.Ptr(), objc.ExtractPtr(responder))
	return bool(result_)
}

func (n NSWindow) SelectKeyViewPrecedingView(view View) {
	C.C_NSWindow_SelectKeyViewPrecedingView(n.Ptr(), objc.ExtractPtr(view))
}

func (n NSWindow) SelectKeyViewFollowingView(view View) {
	C.C_NSWindow_SelectKeyViewFollowingView(n.Ptr(), objc.ExtractPtr(view))
}

func (n NSWindow) SelectPreviousKeyView(sender objc.Object) {
	C.C_NSWindow_SelectPreviousKeyView(n.Ptr(), objc.ExtractPtr(sender))
}

func (n NSWindow) SelectNextKeyView(sender objc.Object) {
	C.C_NSWindow_SelectNextKeyView(n.Ptr(), objc.ExtractPtr(sender))
}

func (n NSWindow) RecalculateKeyViewLoop() {
	C.C_NSWindow_RecalculateKeyViewLoop(n.Ptr())
}

func WindowNumberAtPoint_BelowWindowWithWindowNumber(point foundation.Point, windowNumber int) int {
	result_ := C.C_NSWindow_WindowNumberAtPoint_BelowWindowWithWindowNumber(*(*C.CGPoint)(coregraphics.ToCGPointPointer(coregraphics.Point(point))), C.int(windowNumber))
	return int(result_)
}

func (n NSWindow) PerformWindowDragWithEvent(event Event) {
	C.C_NSWindow_PerformWindowDragWithEvent(n.Ptr(), objc.ExtractPtr(event))
}

func (n NSWindow) DisableSnapshotRestoration() {
	C.C_NSWindow_DisableSnapshotRestoration(n.Ptr())
}

func (n NSWindow) EnableSnapshotRestoration() {
	C.C_NSWindow_EnableSnapshotRestoration(n.Ptr())
}

func (n NSWindow) Display() {
	C.C_NSWindow_Display(n.Ptr())
}

func (n NSWindow) DisplayIfNeeded() {
	C.C_NSWindow_DisplayIfNeeded(n.Ptr())
}

func (n NSWindow) DisableScreenUpdatesUntilFlush() {
	C.C_NSWindow_DisableScreenUpdatesUntilFlush(n.Ptr())
}

func (n NSWindow) Update() {
	C.C_NSWindow_Update(n.Ptr())
}

func (n NSWindow) DragImage_At_Offset_Event_Pasteboard_Source_SlideBack(image Image, baseLocation foundation.Point, initialOffset foundation.Size, event Event, pboard Pasteboard, sourceObj objc.Object, slideFlag bool) {
	C.C_NSWindow_DragImage_At_Offset_Event_Pasteboard_Source_SlideBack(n.Ptr(), objc.ExtractPtr(image), *(*C.CGPoint)(coregraphics.ToCGPointPointer(coregraphics.Point(baseLocation))), *(*C.CGSize)(coregraphics.ToCGSizePointer(coregraphics.Size(initialOffset))), objc.ExtractPtr(event), objc.ExtractPtr(pboard), objc.ExtractPtr(sourceObj), C.bool(slideFlag))
}

func (n NSWindow) RegisterForDraggedTypes(newTypes []PasteboardType) {
	var cNewTypes C.Array
	if len(newTypes) > 0 {
		cNewTypesData := make([]unsafe.Pointer, len(newTypes))
		for idx, v := range newTypes {
			cNewTypesData[idx] = foundation.NewString(string(v)).Ptr()
		}
		cNewTypes.data = unsafe.Pointer(&cNewTypesData[0])
		cNewTypes.len = C.int(len(newTypes))
	}
	C.C_NSWindow_RegisterForDraggedTypes(n.Ptr(), cNewTypes)
}

func (n NSWindow) UnregisterDraggedTypes() {
	C.C_NSWindow_UnregisterDraggedTypes(n.Ptr())
}

func (n NSWindow) ConvertRectFromBacking(rect foundation.Rect) foundation.Rect {
	result_ := C.C_NSWindow_ConvertRectFromBacking(n.Ptr(), *(*C.CGRect)(coregraphics.ToCGRectPointer(coregraphics.Rect(rect))))
	return foundation.Rect(coregraphics.FromCGRectPointer(unsafe.Pointer(&result_)))
}

func (n NSWindow) ConvertRectToBacking(rect foundation.Rect) foundation.Rect {
	result_ := C.C_NSWindow_ConvertRectToBacking(n.Ptr(), *(*C.CGRect)(coregraphics.ToCGRectPointer(coregraphics.Rect(rect))))
	return foundation.Rect(coregraphics.FromCGRectPointer(unsafe.Pointer(&result_)))
}

func (n NSWindow) ConvertRectToScreen(rect foundation.Rect) foundation.Rect {
	result_ := C.C_NSWindow_ConvertRectToScreen(n.Ptr(), *(*C.CGRect)(coregraphics.ToCGRectPointer(coregraphics.Rect(rect))))
	return foundation.Rect(coregraphics.FromCGRectPointer(unsafe.Pointer(&result_)))
}

func (n NSWindow) ConvertRectFromScreen(rect foundation.Rect) foundation.Rect {
	result_ := C.C_NSWindow_ConvertRectFromScreen(n.Ptr(), *(*C.CGRect)(coregraphics.ToCGRectPointer(coregraphics.Rect(rect))))
	return foundation.Rect(coregraphics.FromCGRectPointer(unsafe.Pointer(&result_)))
}

func (n NSWindow) SetTitleWithRepresentedFilename(filename string) {
	C.C_NSWindow_SetTitleWithRepresentedFilename(n.Ptr(), foundation.NewString(filename).Ptr())
}

func (n NSWindow) Center() {
	C.C_NSWindow_Center(n.Ptr())
}

func (n NSWindow) PerformClose(sender objc.Object) {
	C.C_NSWindow_PerformClose(n.Ptr(), objc.ExtractPtr(sender))
}

func (n NSWindow) Close() {
	C.C_NSWindow_Close(n.Ptr())
}

func (n NSWindow) PerformMiniaturize(sender objc.Object) {
	C.C_NSWindow_PerformMiniaturize(n.Ptr(), objc.ExtractPtr(sender))
}

func (n NSWindow) Miniaturize(sender objc.Object) {
	C.C_NSWindow_Miniaturize(n.Ptr(), objc.ExtractPtr(sender))
}

func (n NSWindow) Deminiaturize(sender objc.Object) {
	C.C_NSWindow_Deminiaturize(n.Ptr(), objc.ExtractPtr(sender))
}

func (n NSWindow) Print(sender objc.Object) {
	C.C_NSWindow_Print(n.Ptr(), objc.ExtractPtr(sender))
}

func (n NSWindow) DataWithEPSInsideRect(rect foundation.Rect) []byte {
	result_ := C.C_NSWindow_DataWithEPSInsideRect(n.Ptr(), *(*C.CGRect)(coregraphics.ToCGRectPointer(coregraphics.Rect(rect))))
	return foundation.MakeData(result_).ToBytes()
}

func (n NSWindow) DataWithPDFInsideRect(rect foundation.Rect) []byte {
	result_ := C.C_NSWindow_DataWithPDFInsideRect(n.Ptr(), *(*C.CGRect)(coregraphics.ToCGRectPointer(coregraphics.Rect(rect))))
	return foundation.MakeData(result_).ToBytes()
}

func (n NSWindow) UpdateConstraintsIfNeeded() {
	C.C_NSWindow_UpdateConstraintsIfNeeded(n.Ptr())
}

func (n NSWindow) LayoutIfNeeded() {
	C.C_NSWindow_LayoutIfNeeded(n.Ptr())
}

func (n NSWindow) VisualizeConstraints(constraints []LayoutConstraint) {
	var cConstraints C.Array
	if len(constraints) > 0 {
		cConstraintsData := make([]unsafe.Pointer, len(constraints))
		for idx, v := range constraints {
			cConstraintsData[idx] = objc.ExtractPtr(v)
		}
		cConstraints.data = unsafe.Pointer(&cConstraintsData[0])
		cConstraints.len = C.int(len(constraints))
	}
	C.C_NSWindow_VisualizeConstraints(n.Ptr(), cConstraints)
}

func (n NSWindow) AnchorAttributeForOrientation(orientation LayoutConstraintOrientation) LayoutAttribute {
	result_ := C.C_NSWindow_AnchorAttributeForOrientation(n.Ptr(), C.int(int(orientation)))
	return LayoutAttribute(int(result_))
}

func (n NSWindow) SetAnchorAttribute_ForOrientation(attr LayoutAttribute, orientation LayoutConstraintOrientation) {
	C.C_NSWindow_SetAnchorAttribute_ForOrientation(n.Ptr(), C.int(int(attr)), C.int(int(orientation)))
}

func (n NSWindow) SetIsMiniaturized(flag bool) {
	C.C_NSWindow_SetIsMiniaturized(n.Ptr(), C.bool(flag))
}

func (n NSWindow) SetIsVisible(flag bool) {
	C.C_NSWindow_SetIsVisible(n.Ptr(), C.bool(flag))
}

func (n NSWindow) SetIsZoomed(flag bool) {
	C.C_NSWindow_SetIsZoomed(n.Ptr(), C.bool(flag))
}

func (n NSWindow) HandleCloseScriptCommand(command foundation.CloseCommand) objc.Object {
	result_ := C.C_NSWindow_HandleCloseScriptCommand(n.Ptr(), objc.ExtractPtr(command))
	return objc.MakeObject(result_)
}

func (n NSWindow) HandlePrintScriptCommand(command foundation.ScriptCommand) objc.Object {
	result_ := C.C_NSWindow_HandlePrintScriptCommand(n.Ptr(), objc.ExtractPtr(command))
	return objc.MakeObject(result_)
}

func (n NSWindow) HandleSaveScriptCommand(command foundation.ScriptCommand) objc.Object {
	result_ := C.C_NSWindow_HandleSaveScriptCommand(n.Ptr(), objc.ExtractPtr(command))
	return objc.MakeObject(result_)
}

func (n NSWindow) CanRepresentDisplayGamut(displayGamut DisplayGamut) bool {
	result_ := C.C_NSWindow_CanRepresentDisplayGamut(n.Ptr(), C.int(int(displayGamut)))
	return bool(result_)
}

func (n NSWindow) ConvertPointFromScreen(point foundation.Point) foundation.Point {
	result_ := C.C_NSWindow_ConvertPointFromScreen(n.Ptr(), *(*C.CGPoint)(coregraphics.ToCGPointPointer(coregraphics.Point(point))))
	return foundation.Point(coregraphics.FromCGPointPointer(unsafe.Pointer(&result_)))
}

func (n NSWindow) ConvertPointToScreen(point foundation.Point) foundation.Point {
	result_ := C.C_NSWindow_ConvertPointToScreen(n.Ptr(), *(*C.CGPoint)(coregraphics.ToCGPointPointer(coregraphics.Point(point))))
	return foundation.Point(coregraphics.FromCGPointPointer(unsafe.Pointer(&result_)))
}

func (n NSWindow) ConvertPointFromBacking(point foundation.Point) foundation.Point {
	result_ := C.C_NSWindow_ConvertPointFromBacking(n.Ptr(), *(*C.CGPoint)(coregraphics.ToCGPointPointer(coregraphics.Point(point))))
	return foundation.Point(coregraphics.FromCGPointPointer(unsafe.Pointer(&result_)))
}

func (n NSWindow) ConvertPointToBacking(point foundation.Point) foundation.Point {
	result_ := C.C_NSWindow_ConvertPointToBacking(n.Ptr(), *(*C.CGPoint)(coregraphics.ToCGPointPointer(coregraphics.Point(point))))
	return foundation.Point(coregraphics.FromCGPointPointer(unsafe.Pointer(&result_)))
}

func (n NSWindow) MergeAllWindows(sender objc.Object) {
	C.C_NSWindow_MergeAllWindows(n.Ptr(), objc.ExtractPtr(sender))
}

func (n NSWindow) SetDynamicDepthLimit(flag bool) {
	C.C_NSWindow_SetDynamicDepthLimit(n.Ptr(), C.bool(flag))
}

func (n NSWindow) SetFrameAutosaveName(name WindowFrameAutosaveName) bool {
	result_ := C.C_NSWindow_SetFrameAutosaveName(n.Ptr(), foundation.NewString(string(name)).Ptr())
	return bool(result_)
}

func (n NSWindow) Delegate() objc.Object {
	result_ := C.C_NSWindow_Delegate(n.Ptr())
	return objc.MakeObject(result_)
}

func (n NSWindow) SetDelegate(value objc.Object) {
	C.C_NSWindow_SetDelegate(n.Ptr(), objc.ExtractPtr(value))
}

func (n NSWindow) ContentViewController() ViewController {
	result_ := C.C_NSWindow_ContentViewController(n.Ptr())
	return MakeViewController(result_)
}

func (n NSWindow) SetContentViewController(value ViewController) {
	C.C_NSWindow_SetContentViewController(n.Ptr(), objc.ExtractPtr(value))
}

func (n NSWindow) ContentView() View {
	result_ := C.C_NSWindow_ContentView(n.Ptr())
	return MakeView(result_)
}

func (n NSWindow) SetContentView(value View) {
	C.C_NSWindow_SetContentView(n.Ptr(), objc.ExtractPtr(value))
}

func (n NSWindow) WorksWhenModal() bool {
	result_ := C.C_NSWindow_WorksWhenModal(n.Ptr())
	return bool(result_)
}

func (n NSWindow) AlphaValue() coregraphics.Float {
	result_ := C.C_NSWindow_AlphaValue(n.Ptr())
	return coregraphics.Float(float64(result_))
}

func (n NSWindow) SetAlphaValue(value coregraphics.Float) {
	C.C_NSWindow_SetAlphaValue(n.Ptr(), C.double(float64(value)))
}

func (n NSWindow) BackgroundColor() Color {
	result_ := C.C_NSWindow_BackgroundColor(n.Ptr())
	return MakeColor(result_)
}

func (n NSWindow) SetBackgroundColor(value Color) {
	C.C_NSWindow_SetBackgroundColor(n.Ptr(), objc.ExtractPtr(value))
}

func (n NSWindow) ColorSpace() ColorSpace {
	result_ := C.C_NSWindow_ColorSpace(n.Ptr())
	return MakeColorSpace(result_)
}

func (n NSWindow) SetColorSpace(value ColorSpace) {
	C.C_NSWindow_SetColorSpace(n.Ptr(), objc.ExtractPtr(value))
}

func (n NSWindow) CanHide() bool {
	result_ := C.C_NSWindow_CanHide(n.Ptr())
	return bool(result_)
}

func (n NSWindow) SetCanHide(value bool) {
	C.C_NSWindow_SetCanHide(n.Ptr(), C.bool(value))
}

func (n NSWindow) IsOnActiveSpace() bool {
	result_ := C.C_NSWindow_IsOnActiveSpace(n.Ptr())
	return bool(result_)
}

func (n NSWindow) HidesOnDeactivate() bool {
	result_ := C.C_NSWindow_HidesOnDeactivate(n.Ptr())
	return bool(result_)
}

func (n NSWindow) SetHidesOnDeactivate(value bool) {
	C.C_NSWindow_SetHidesOnDeactivate(n.Ptr(), C.bool(value))
}

func (n NSWindow) CollectionBehavior() WindowCollectionBehavior {
	result_ := C.C_NSWindow_CollectionBehavior(n.Ptr())
	return WindowCollectionBehavior(uint(result_))
}

func (n NSWindow) SetCollectionBehavior(value WindowCollectionBehavior) {
	C.C_NSWindow_SetCollectionBehavior(n.Ptr(), C.uint(uint(value)))
}

func (n NSWindow) IsOpaque() bool {
	result_ := C.C_NSWindow_IsOpaque(n.Ptr())
	return bool(result_)
}

func (n NSWindow) SetOpaque(value bool) {
	C.C_NSWindow_SetOpaque(n.Ptr(), C.bool(value))
}

func (n NSWindow) HasShadow() bool {
	result_ := C.C_NSWindow_HasShadow(n.Ptr())
	return bool(result_)
}

func (n NSWindow) SetHasShadow(value bool) {
	C.C_NSWindow_SetHasShadow(n.Ptr(), C.bool(value))
}

func (n NSWindow) PreventsApplicationTerminationWhenModal() bool {
	result_ := C.C_NSWindow_PreventsApplicationTerminationWhenModal(n.Ptr())
	return bool(result_)
}

func (n NSWindow) SetPreventsApplicationTerminationWhenModal(value bool) {
	C.C_NSWindow_SetPreventsApplicationTerminationWhenModal(n.Ptr(), C.bool(value))
}

func Window_DefaultDepthLimit() WindowDepth {
	result_ := C.C_NSWindow_Window_DefaultDepthLimit()
	return WindowDepth(int32(result_))
}

func (n NSWindow) WindowNumber() int {
	result_ := C.C_NSWindow_WindowNumber(n.Ptr())
	return int(result_)
}

func (n NSWindow) DeviceDescription() map[DeviceDescriptionKey]objc.Object {
	result_ := C.C_NSWindow_DeviceDescription(n.Ptr())
	if result_.len > 0 {
		defer C.free(result_.key_data)
		defer C.free(result_.value_data)
	}
	result_KeySlice := unsafe.Slice((*unsafe.Pointer)(result_.key_data), int(result_.len))
	result_ValueSlice := unsafe.Slice((*unsafe.Pointer)(result_.value_data), int(result_.len))
	var goResult_ = make(map[DeviceDescriptionKey]objc.Object)
	for idx, k := range result_KeySlice {
		v := result_ValueSlice[idx]
		goResult_[DeviceDescriptionKey(foundation.MakeString(k).String())] = objc.MakeObject(v)
	}
	return goResult_
}

func (n NSWindow) CanBecomeVisibleWithoutLogin() bool {
	result_ := C.C_NSWindow_CanBecomeVisibleWithoutLogin(n.Ptr())
	return bool(result_)
}

func (n NSWindow) SetCanBecomeVisibleWithoutLogin(value bool) {
	C.C_NSWindow_SetCanBecomeVisibleWithoutLogin(n.Ptr(), C.bool(value))
}

func (n NSWindow) SharingType() WindowSharingType {
	result_ := C.C_NSWindow_SharingType(n.Ptr())
	return WindowSharingType(uint(result_))
}

func (n NSWindow) SetSharingType(value WindowSharingType) {
	C.C_NSWindow_SetSharingType(n.Ptr(), C.uint(uint(value)))
}

func (n NSWindow) BackingType() BackingStoreType {
	result_ := C.C_NSWindow_BackingType(n.Ptr())
	return BackingStoreType(uint(result_))
}

func (n NSWindow) SetBackingType(value BackingStoreType) {
	C.C_NSWindow_SetBackingType(n.Ptr(), C.uint(uint(value)))
}

func (n NSWindow) DepthLimit() WindowDepth {
	result_ := C.C_NSWindow_DepthLimit(n.Ptr())
	return WindowDepth(int32(result_))
}

func (n NSWindow) SetDepthLimit(value WindowDepth) {
	C.C_NSWindow_SetDepthLimit(n.Ptr(), C.int(int32(value)))
}

func (n NSWindow) HasDynamicDepthLimit() bool {
	result_ := C.C_NSWindow_HasDynamicDepthLimit(n.Ptr())
	return bool(result_)
}

func (n NSWindow) WindowController() WindowController {
	result_ := C.C_NSWindow_WindowController(n.Ptr())
	return MakeWindowController(result_)
}

func (n NSWindow) SetWindowController(value WindowController) {
	C.C_NSWindow_SetWindowController(n.Ptr(), objc.ExtractPtr(value))
}

func (n NSWindow) AttachedSheet() Window {
	result_ := C.C_NSWindow_AttachedSheet(n.Ptr())
	return MakeWindow(result_)
}

func (n NSWindow) IsSheet() bool {
	result_ := C.C_NSWindow_IsSheet(n.Ptr())
	return bool(result_)
}

func (n NSWindow) SheetParent() Window {
	result_ := C.C_NSWindow_SheetParent(n.Ptr())
	return MakeWindow(result_)
}

func (n NSWindow) Sheets() []Window {
	result_ := C.C_NSWindow_Sheets(n.Ptr())
	if result_.len > 0 {
		defer C.free(result_.data)
	}
	result_Slice := unsafe.Slice((*unsafe.Pointer)(result_.data), int(result_.len))
	var goResult_ = make([]Window, len(result_Slice))
	for idx, r := range result_Slice {
		goResult_[idx] = MakeWindow(r)
	}
	return goResult_
}

func (n NSWindow) Frame() foundation.Rect {
	result_ := C.C_NSWindow_Frame(n.Ptr())
	return foundation.Rect(coregraphics.FromCGRectPointer(unsafe.Pointer(&result_)))
}

func (n NSWindow) AspectRatio() foundation.Size {
	result_ := C.C_NSWindow_AspectRatio(n.Ptr())
	return foundation.Size(coregraphics.FromCGSizePointer(unsafe.Pointer(&result_)))
}

func (n NSWindow) SetAspectRatio(value foundation.Size) {
	C.C_NSWindow_SetAspectRatio(n.Ptr(), *(*C.CGSize)(coregraphics.ToCGSizePointer(coregraphics.Size(value))))
}

func (n NSWindow) MinSize() foundation.Size {
	result_ := C.C_NSWindow_MinSize(n.Ptr())
	return foundation.Size(coregraphics.FromCGSizePointer(unsafe.Pointer(&result_)))
}

func (n NSWindow) SetMinSize(value foundation.Size) {
	C.C_NSWindow_SetMinSize(n.Ptr(), *(*C.CGSize)(coregraphics.ToCGSizePointer(coregraphics.Size(value))))
}

func (n NSWindow) MaxSize() foundation.Size {
	result_ := C.C_NSWindow_MaxSize(n.Ptr())
	return foundation.Size(coregraphics.FromCGSizePointer(unsafe.Pointer(&result_)))
}

func (n NSWindow) SetMaxSize(value foundation.Size) {
	C.C_NSWindow_SetMaxSize(n.Ptr(), *(*C.CGSize)(coregraphics.ToCGSizePointer(coregraphics.Size(value))))
}

func (n NSWindow) IsZoomed() bool {
	result_ := C.C_NSWindow_IsZoomed(n.Ptr())
	return bool(result_)
}

func (n NSWindow) ResizeFlags() EventModifierFlags {
	result_ := C.C_NSWindow_ResizeFlags(n.Ptr())
	return EventModifierFlags(uint(result_))
}

func (n NSWindow) ResizeIncrements() foundation.Size {
	result_ := C.C_NSWindow_ResizeIncrements(n.Ptr())
	return foundation.Size(coregraphics.FromCGSizePointer(unsafe.Pointer(&result_)))
}

func (n NSWindow) SetResizeIncrements(value foundation.Size) {
	C.C_NSWindow_SetResizeIncrements(n.Ptr(), *(*C.CGSize)(coregraphics.ToCGSizePointer(coregraphics.Size(value))))
}

func (n NSWindow) PreservesContentDuringLiveResize() bool {
	result_ := C.C_NSWindow_PreservesContentDuringLiveResize(n.Ptr())
	return bool(result_)
}

func (n NSWindow) SetPreservesContentDuringLiveResize(value bool) {
	C.C_NSWindow_SetPreservesContentDuringLiveResize(n.Ptr(), C.bool(value))
}

func (n NSWindow) InLiveResize() bool {
	result_ := C.C_NSWindow_InLiveResize(n.Ptr())
	return bool(result_)
}

func (n NSWindow) ContentAspectRatio() foundation.Size {
	result_ := C.C_NSWindow_ContentAspectRatio(n.Ptr())
	return foundation.Size(coregraphics.FromCGSizePointer(unsafe.Pointer(&result_)))
}

func (n NSWindow) SetContentAspectRatio(value foundation.Size) {
	C.C_NSWindow_SetContentAspectRatio(n.Ptr(), *(*C.CGSize)(coregraphics.ToCGSizePointer(coregraphics.Size(value))))
}

func (n NSWindow) ContentMinSize() foundation.Size {
	result_ := C.C_NSWindow_ContentMinSize(n.Ptr())
	return foundation.Size(coregraphics.FromCGSizePointer(unsafe.Pointer(&result_)))
}

func (n NSWindow) SetContentMinSize(value foundation.Size) {
	C.C_NSWindow_SetContentMinSize(n.Ptr(), *(*C.CGSize)(coregraphics.ToCGSizePointer(coregraphics.Size(value))))
}

func (n NSWindow) ContentMaxSize() foundation.Size {
	result_ := C.C_NSWindow_ContentMaxSize(n.Ptr())
	return foundation.Size(coregraphics.FromCGSizePointer(unsafe.Pointer(&result_)))
}

func (n NSWindow) SetContentMaxSize(value foundation.Size) {
	C.C_NSWindow_SetContentMaxSize(n.Ptr(), *(*C.CGSize)(coregraphics.ToCGSizePointer(coregraphics.Size(value))))
}

func (n NSWindow) ContentResizeIncrements() foundation.Size {
	result_ := C.C_NSWindow_ContentResizeIncrements(n.Ptr())
	return foundation.Size(coregraphics.FromCGSizePointer(unsafe.Pointer(&result_)))
}

func (n NSWindow) SetContentResizeIncrements(value foundation.Size) {
	C.C_NSWindow_SetContentResizeIncrements(n.Ptr(), *(*C.CGSize)(coregraphics.ToCGSizePointer(coregraphics.Size(value))))
}

func (n NSWindow) ContentLayoutGuide() objc.Object {
	result_ := C.C_NSWindow_ContentLayoutGuide(n.Ptr())
	return objc.MakeObject(result_)
}

func (n NSWindow) ContentLayoutRect() foundation.Rect {
	result_ := C.C_NSWindow_ContentLayoutRect(n.Ptr())
	return foundation.Rect(coregraphics.FromCGRectPointer(unsafe.Pointer(&result_)))
}

func (n NSWindow) MaxFullScreenContentSize() foundation.Size {
	result_ := C.C_NSWindow_MaxFullScreenContentSize(n.Ptr())
	return foundation.Size(coregraphics.FromCGSizePointer(unsafe.Pointer(&result_)))
}

func (n NSWindow) SetMaxFullScreenContentSize(value foundation.Size) {
	C.C_NSWindow_SetMaxFullScreenContentSize(n.Ptr(), *(*C.CGSize)(coregraphics.ToCGSizePointer(coregraphics.Size(value))))
}

func (n NSWindow) MinFullScreenContentSize() foundation.Size {
	result_ := C.C_NSWindow_MinFullScreenContentSize(n.Ptr())
	return foundation.Size(coregraphics.FromCGSizePointer(unsafe.Pointer(&result_)))
}

func (n NSWindow) SetMinFullScreenContentSize(value foundation.Size) {
	C.C_NSWindow_SetMinFullScreenContentSize(n.Ptr(), *(*C.CGSize)(coregraphics.ToCGSizePointer(coregraphics.Size(value))))
}

func (n NSWindow) Level() WindowLevel {
	result_ := C.C_NSWindow_Level(n.Ptr())
	return WindowLevel(int(result_))
}

func (n NSWindow) SetLevel(value WindowLevel) {
	C.C_NSWindow_SetLevel(n.Ptr(), C.int(int(value)))
}

func (n NSWindow) IsVisible() bool {
	result_ := C.C_NSWindow_IsVisible(n.Ptr())
	return bool(result_)
}

func (n NSWindow) OcclusionState() WindowOcclusionState {
	result_ := C.C_NSWindow_OcclusionState(n.Ptr())
	return WindowOcclusionState(uint(result_))
}

func (n NSWindow) FrameAutosaveName() WindowFrameAutosaveName {
	result_ := C.C_NSWindow_FrameAutosaveName(n.Ptr())
	return WindowFrameAutosaveName(foundation.MakeString(result_).String())
}

func (n NSWindow) StringWithSavedFrame() WindowPersistableFrameDescriptor {
	result_ := C.C_NSWindow_StringWithSavedFrame(n.Ptr())
	return WindowPersistableFrameDescriptor(foundation.MakeString(result_).String())
}

func (n NSWindow) IsKeyWindow() bool {
	result_ := C.C_NSWindow_IsKeyWindow(n.Ptr())
	return bool(result_)
}

func (n NSWindow) CanBecomeKeyWindow() bool {
	result_ := C.C_NSWindow_CanBecomeKeyWindow(n.Ptr())
	return bool(result_)
}

func (n NSWindow) IsMainWindow() bool {
	result_ := C.C_NSWindow_IsMainWindow(n.Ptr())
	return bool(result_)
}

func (n NSWindow) CanBecomeMainWindow() bool {
	result_ := C.C_NSWindow_CanBecomeMainWindow(n.Ptr())
	return bool(result_)
}

func (n NSWindow) Toolbar() Toolbar {
	result_ := C.C_NSWindow_Toolbar(n.Ptr())
	return MakeToolbar(result_)
}

func (n NSWindow) SetToolbar(value Toolbar) {
	C.C_NSWindow_SetToolbar(n.Ptr(), objc.ExtractPtr(value))
}

func (n NSWindow) ChildWindows() []Window {
	result_ := C.C_NSWindow_ChildWindows(n.Ptr())
	if result_.len > 0 {
		defer C.free(result_.data)
	}
	result_Slice := unsafe.Slice((*unsafe.Pointer)(result_.data), int(result_.len))
	var goResult_ = make([]Window, len(result_Slice))
	for idx, r := range result_Slice {
		goResult_[idx] = MakeWindow(r)
	}
	return goResult_
}

func (n NSWindow) ParentWindow() Window {
	result_ := C.C_NSWindow_ParentWindow(n.Ptr())
	return MakeWindow(result_)
}

func (n NSWindow) SetParentWindow(value Window) {
	C.C_NSWindow_SetParentWindow(n.Ptr(), objc.ExtractPtr(value))
}

func (n NSWindow) DefaultButtonCell() ButtonCell {
	result_ := C.C_NSWindow_DefaultButtonCell(n.Ptr())
	return MakeButtonCell(result_)
}

func (n NSWindow) SetDefaultButtonCell(value ButtonCell) {
	C.C_NSWindow_SetDefaultButtonCell(n.Ptr(), objc.ExtractPtr(value))
}

func (n NSWindow) IsExcludedFromWindowsMenu() bool {
	result_ := C.C_NSWindow_IsExcludedFromWindowsMenu(n.Ptr())
	return bool(result_)
}

func (n NSWindow) SetExcludedFromWindowsMenu(value bool) {
	C.C_NSWindow_SetExcludedFromWindowsMenu(n.Ptr(), C.bool(value))
}

func (n NSWindow) AreCursorRectsEnabled() bool {
	result_ := C.C_NSWindow_AreCursorRectsEnabled(n.Ptr())
	return bool(result_)
}

func (n NSWindow) ShowsToolbarButton() bool {
	result_ := C.C_NSWindow_ShowsToolbarButton(n.Ptr())
	return bool(result_)
}

func (n NSWindow) SetShowsToolbarButton(value bool) {
	C.C_NSWindow_SetShowsToolbarButton(n.Ptr(), C.bool(value))
}

func (n NSWindow) TitlebarAppearsTransparent() bool {
	result_ := C.C_NSWindow_TitlebarAppearsTransparent(n.Ptr())
	return bool(result_)
}

func (n NSWindow) SetTitlebarAppearsTransparent(value bool) {
	C.C_NSWindow_SetTitlebarAppearsTransparent(n.Ptr(), C.bool(value))
}

func (n NSWindow) ToolbarStyle() WindowToolbarStyle {
	result_ := C.C_NSWindow_ToolbarStyle(n.Ptr())
	return WindowToolbarStyle(int(result_))
}

func (n NSWindow) SetToolbarStyle(value WindowToolbarStyle) {
	C.C_NSWindow_SetToolbarStyle(n.Ptr(), C.int(int(value)))
}

func (n NSWindow) TitlebarSeparatorStyle() TitlebarSeparatorStyle {
	result_ := C.C_NSWindow_TitlebarSeparatorStyle(n.Ptr())
	return TitlebarSeparatorStyle(int(result_))
}

func (n NSWindow) SetTitlebarSeparatorStyle(value TitlebarSeparatorStyle) {
	C.C_NSWindow_SetTitlebarSeparatorStyle(n.Ptr(), C.int(int(value)))
}

func (n NSWindow) TitlebarAccessoryViewControllers() []TitlebarAccessoryViewController {
	result_ := C.C_NSWindow_TitlebarAccessoryViewControllers(n.Ptr())
	if result_.len > 0 {
		defer C.free(result_.data)
	}
	result_Slice := unsafe.Slice((*unsafe.Pointer)(result_.data), int(result_.len))
	var goResult_ = make([]TitlebarAccessoryViewController, len(result_Slice))
	for idx, r := range result_Slice {
		goResult_[idx] = MakeTitlebarAccessoryViewController(r)
	}
	return goResult_
}

func (n NSWindow) SetTitlebarAccessoryViewControllers(value []TitlebarAccessoryViewController) {
	var cValue C.Array
	if len(value) > 0 {
		cValueData := make([]unsafe.Pointer, len(value))
		for idx, v := range value {
			cValueData[idx] = objc.ExtractPtr(v)
		}
		cValue.data = unsafe.Pointer(&cValueData[0])
		cValue.len = C.int(len(value))
	}
	C.C_NSWindow_SetTitlebarAccessoryViewControllers(n.Ptr(), cValue)
}

func AllowsAutomaticWindowTabbing() bool {
	result_ := C.C_NSWindow_AllowsAutomaticWindowTabbing()
	return bool(result_)
}

func SetAllowsAutomaticWindowTabbing(value bool) {
	C.C_NSWindow_SetAllowsAutomaticWindowTabbing(C.bool(value))
}

func Window_UserTabbingPreference() WindowUserTabbingPreference {
	result_ := C.C_NSWindow_Window_UserTabbingPreference()
	return WindowUserTabbingPreference(int(result_))
}

func (n NSWindow) Tab() WindowTab {
	result_ := C.C_NSWindow_Tab(n.Ptr())
	return MakeWindowTab(result_)
}

func (n NSWindow) TabbingIdentifier() WindowTabbingIdentifier {
	result_ := C.C_NSWindow_TabbingIdentifier(n.Ptr())
	return WindowTabbingIdentifier(foundation.MakeString(result_).String())
}

func (n NSWindow) SetTabbingIdentifier(value WindowTabbingIdentifier) {
	C.C_NSWindow_SetTabbingIdentifier(n.Ptr(), foundation.NewString(string(value)).Ptr())
}

func (n NSWindow) TabbingMode() WindowTabbingMode {
	result_ := C.C_NSWindow_TabbingMode(n.Ptr())
	return WindowTabbingMode(int(result_))
}

func (n NSWindow) SetTabbingMode(value WindowTabbingMode) {
	C.C_NSWindow_SetTabbingMode(n.Ptr(), C.int(int(value)))
}

func (n NSWindow) TabbedWindows() []Window {
	result_ := C.C_NSWindow_TabbedWindows(n.Ptr())
	if result_.len > 0 {
		defer C.free(result_.data)
	}
	result_Slice := unsafe.Slice((*unsafe.Pointer)(result_.data), int(result_.len))
	var goResult_ = make([]Window, len(result_Slice))
	for idx, r := range result_Slice {
		goResult_[idx] = MakeWindow(r)
	}
	return goResult_
}

func (n NSWindow) TabGroup() WindowTabGroup {
	result_ := C.C_NSWindow_TabGroup(n.Ptr())
	return MakeWindowTabGroup(result_)
}

func (n NSWindow) AllowsToolTipsWhenApplicationIsInactive() bool {
	result_ := C.C_NSWindow_AllowsToolTipsWhenApplicationIsInactive(n.Ptr())
	return bool(result_)
}

func (n NSWindow) SetAllowsToolTipsWhenApplicationIsInactive(value bool) {
	C.C_NSWindow_SetAllowsToolTipsWhenApplicationIsInactive(n.Ptr(), C.bool(value))
}

func (n NSWindow) CurrentEvent() Event {
	result_ := C.C_NSWindow_CurrentEvent(n.Ptr())
	return MakeEvent(result_)
}

func (n NSWindow) InitialFirstResponder() View {
	result_ := C.C_NSWindow_InitialFirstResponder(n.Ptr())
	return MakeView(result_)
}

func (n NSWindow) SetInitialFirstResponder(value View) {
	C.C_NSWindow_SetInitialFirstResponder(n.Ptr(), objc.ExtractPtr(value))
}

func (n NSWindow) FirstResponder() Responder {
	result_ := C.C_NSWindow_FirstResponder(n.Ptr())
	return MakeResponder(result_)
}

func (n NSWindow) KeyViewSelectionDirection() SelectionDirection {
	result_ := C.C_NSWindow_KeyViewSelectionDirection(n.Ptr())
	return SelectionDirection(uint(result_))
}

func (n NSWindow) AutorecalculatesKeyViewLoop() bool {
	result_ := C.C_NSWindow_AutorecalculatesKeyViewLoop(n.Ptr())
	return bool(result_)
}

func (n NSWindow) SetAutorecalculatesKeyViewLoop(value bool) {
	C.C_NSWindow_SetAutorecalculatesKeyViewLoop(n.Ptr(), C.bool(value))
}

func (n NSWindow) AcceptsMouseMovedEvents() bool {
	result_ := C.C_NSWindow_AcceptsMouseMovedEvents(n.Ptr())
	return bool(result_)
}

func (n NSWindow) SetAcceptsMouseMovedEvents(value bool) {
	C.C_NSWindow_SetAcceptsMouseMovedEvents(n.Ptr(), C.bool(value))
}

func (n NSWindow) IgnoresMouseEvents() bool {
	result_ := C.C_NSWindow_IgnoresMouseEvents(n.Ptr())
	return bool(result_)
}

func (n NSWindow) SetIgnoresMouseEvents(value bool) {
	C.C_NSWindow_SetIgnoresMouseEvents(n.Ptr(), C.bool(value))
}

func (n NSWindow) MouseLocationOutsideOfEventStream() foundation.Point {
	result_ := C.C_NSWindow_MouseLocationOutsideOfEventStream(n.Ptr())
	return foundation.Point(coregraphics.FromCGPointPointer(unsafe.Pointer(&result_)))
}

func (n NSWindow) IsRestorable() bool {
	result_ := C.C_NSWindow_IsRestorable(n.Ptr())
	return bool(result_)
}

func (n NSWindow) SetRestorable(value bool) {
	C.C_NSWindow_SetRestorable(n.Ptr(), C.bool(value))
}

func (n NSWindow) ViewsNeedDisplay() bool {
	result_ := C.C_NSWindow_ViewsNeedDisplay(n.Ptr())
	return bool(result_)
}

func (n NSWindow) SetViewsNeedDisplay(value bool) {
	C.C_NSWindow_SetViewsNeedDisplay(n.Ptr(), C.bool(value))
}

func (n NSWindow) AllowsConcurrentViewDrawing() bool {
	result_ := C.C_NSWindow_AllowsConcurrentViewDrawing(n.Ptr())
	return bool(result_)
}

func (n NSWindow) SetAllowsConcurrentViewDrawing(value bool) {
	C.C_NSWindow_SetAllowsConcurrentViewDrawing(n.Ptr(), C.bool(value))
}

func (n NSWindow) AnimationBehavior() WindowAnimationBehavior {
	result_ := C.C_NSWindow_AnimationBehavior(n.Ptr())
	return WindowAnimationBehavior(int(result_))
}

func (n NSWindow) SetAnimationBehavior(value WindowAnimationBehavior) {
	C.C_NSWindow_SetAnimationBehavior(n.Ptr(), C.int(int(value)))
}

func (n NSWindow) IsDocumentEdited() bool {
	result_ := C.C_NSWindow_IsDocumentEdited(n.Ptr())
	return bool(result_)
}

func (n NSWindow) SetDocumentEdited(value bool) {
	C.C_NSWindow_SetDocumentEdited(n.Ptr(), C.bool(value))
}

func (n NSWindow) BackingScaleFactor() coregraphics.Float {
	result_ := C.C_NSWindow_BackingScaleFactor(n.Ptr())
	return coregraphics.Float(float64(result_))
}

func (n NSWindow) Title() string {
	result_ := C.C_NSWindow_Title(n.Ptr())
	return foundation.MakeString(result_).String()
}

func (n NSWindow) SetTitle(value string) {
	C.C_NSWindow_SetTitle(n.Ptr(), foundation.NewString(value).Ptr())
}

func (n NSWindow) Subtitle() string {
	result_ := C.C_NSWindow_Subtitle(n.Ptr())
	return foundation.MakeString(result_).String()
}

func (n NSWindow) SetSubtitle(value string) {
	C.C_NSWindow_SetSubtitle(n.Ptr(), foundation.NewString(value).Ptr())
}

func (n NSWindow) TitleVisibility() WindowTitleVisibility {
	result_ := C.C_NSWindow_TitleVisibility(n.Ptr())
	return WindowTitleVisibility(int(result_))
}

func (n NSWindow) SetTitleVisibility(value WindowTitleVisibility) {
	C.C_NSWindow_SetTitleVisibility(n.Ptr(), C.int(int(value)))
}

func (n NSWindow) RepresentedFilename() string {
	result_ := C.C_NSWindow_RepresentedFilename(n.Ptr())
	return foundation.MakeString(result_).String()
}

func (n NSWindow) SetRepresentedFilename(value string) {
	C.C_NSWindow_SetRepresentedFilename(n.Ptr(), foundation.NewString(value).Ptr())
}

func (n NSWindow) RepresentedURL() foundation.URL {
	result_ := C.C_NSWindow_RepresentedURL(n.Ptr())
	return foundation.MakeURL(result_)
}

func (n NSWindow) SetRepresentedURL(value foundation.URL) {
	C.C_NSWindow_SetRepresentedURL(n.Ptr(), objc.ExtractPtr(value))
}

func (n NSWindow) Screen() Screen {
	result_ := C.C_NSWindow_Screen(n.Ptr())
	return MakeScreen(result_)
}

func (n NSWindow) DeepestScreen() Screen {
	result_ := C.C_NSWindow_DeepestScreen(n.Ptr())
	return MakeScreen(result_)
}

func (n NSWindow) DisplaysWhenScreenProfileChanges() bool {
	result_ := C.C_NSWindow_DisplaysWhenScreenProfileChanges(n.Ptr())
	return bool(result_)
}

func (n NSWindow) SetDisplaysWhenScreenProfileChanges(value bool) {
	C.C_NSWindow_SetDisplaysWhenScreenProfileChanges(n.Ptr(), C.bool(value))
}

func (n NSWindow) IsMovableByWindowBackground() bool {
	result_ := C.C_NSWindow_IsMovableByWindowBackground(n.Ptr())
	return bool(result_)
}

func (n NSWindow) SetMovableByWindowBackground(value bool) {
	C.C_NSWindow_SetMovableByWindowBackground(n.Ptr(), C.bool(value))
}

func (n NSWindow) IsMovable() bool {
	result_ := C.C_NSWindow_IsMovable(n.Ptr())
	return bool(result_)
}

func (n NSWindow) SetMovable(value bool) {
	C.C_NSWindow_SetMovable(n.Ptr(), C.bool(value))
}

func (n NSWindow) IsReleasedWhenClosed() bool {
	result_ := C.C_NSWindow_IsReleasedWhenClosed(n.Ptr())
	return bool(result_)
}

func (n NSWindow) SetReleasedWhenClosed(value bool) {
	C.C_NSWindow_SetReleasedWhenClosed(n.Ptr(), C.bool(value))
}

func (n NSWindow) IsMiniaturized() bool {
	result_ := C.C_NSWindow_IsMiniaturized(n.Ptr())
	return bool(result_)
}

func (n NSWindow) MiniwindowImage() Image {
	result_ := C.C_NSWindow_MiniwindowImage(n.Ptr())
	return MakeImage(result_)
}

func (n NSWindow) SetMiniwindowImage(value Image) {
	C.C_NSWindow_SetMiniwindowImage(n.Ptr(), objc.ExtractPtr(value))
}

func (n NSWindow) MiniwindowTitle() string {
	result_ := C.C_NSWindow_MiniwindowTitle(n.Ptr())
	return foundation.MakeString(result_).String()
}

func (n NSWindow) SetMiniwindowTitle(value string) {
	C.C_NSWindow_SetMiniwindowTitle(n.Ptr(), foundation.NewString(value).Ptr())
}

func (n NSWindow) DockTile() DockTile {
	result_ := C.C_NSWindow_DockTile(n.Ptr())
	return MakeDockTile(result_)
}

func (n NSWindow) HasCloseBox() bool {
	result_ := C.C_NSWindow_HasCloseBox(n.Ptr())
	return bool(result_)
}

func (n NSWindow) HasTitleBar() bool {
	result_ := C.C_NSWindow_HasTitleBar(n.Ptr())
	return bool(result_)
}

func (n NSWindow) OrderedIndex() int {
	result_ := C.C_NSWindow_OrderedIndex(n.Ptr())
	return int(result_)
}

func (n NSWindow) SetOrderedIndex(value int) {
	C.C_NSWindow_SetOrderedIndex(n.Ptr(), C.int(value))
}

func (n NSWindow) AppearanceSource() objc.Object {
	result_ := C.C_NSWindow_AppearanceSource(n.Ptr())
	return objc.MakeObject(result_)
}

func (n NSWindow) SetAppearanceSource(value objc.Object) {
	C.C_NSWindow_SetAppearanceSource(n.Ptr(), objc.ExtractPtr(value))
}

func (n NSWindow) IsFloatingPanel() bool {
	result_ := C.C_NSWindow_IsFloatingPanel(n.Ptr())
	return bool(result_)
}

func (n NSWindow) IsMiniaturizable() bool {
	result_ := C.C_NSWindow_IsMiniaturizable(n.Ptr())
	return bool(result_)
}

func (n NSWindow) IsModalPanel() bool {
	result_ := C.C_NSWindow_IsModalPanel(n.Ptr())
	return bool(result_)
}

func (n NSWindow) IsResizable() bool {
	result_ := C.C_NSWindow_IsResizable(n.Ptr())
	return bool(result_)
}

func (n NSWindow) StyleMask() WindowStyleMask {
	result_ := C.C_NSWindow_StyleMask(n.Ptr())
	return WindowStyleMask(uint(result_))
}

func (n NSWindow) SetStyleMask(value WindowStyleMask) {
	C.C_NSWindow_SetStyleMask(n.Ptr(), C.uint(uint(value)))
}

func (n NSWindow) WindowTitlebarLayoutDirection() UserInterfaceLayoutDirection {
	result_ := C.C_NSWindow_WindowTitlebarLayoutDirection(n.Ptr())
	return UserInterfaceLayoutDirection(int(result_))
}

func (n NSWindow) IsZoomable() bool {
	result_ := C.C_NSWindow_IsZoomable(n.Ptr())
	return bool(result_)
}

type WindowDelegate struct {
	Window_WillPositionSheet_UsingRect                                    func(window Window, sheet Window, rect foundation.Rect) foundation.Rect
	WindowWillBeginSheet                                                  func(notification foundation.Notification)
	WindowDidEndSheet                                                     func(notification foundation.Notification)
	WindowWillResize_ToSize                                               func(sender Window, frameSize foundation.Size) foundation.Size
	WindowDidResize                                                       func(notification foundation.Notification)
	WindowWillStartLiveResize                                             func(notification foundation.Notification)
	WindowDidEndLiveResize                                                func(notification foundation.Notification)
	WindowWillMiniaturize                                                 func(notification foundation.Notification)
	WindowDidMiniaturize                                                  func(notification foundation.Notification)
	WindowDidDeminiaturize                                                func(notification foundation.Notification)
	WindowWillUseStandardFrame_DefaultFrame                               func(window Window, newFrame foundation.Rect) foundation.Rect
	WindowShouldZoom_ToFrame                                              func(window Window, newFrame foundation.Rect) bool
	Window_WillUseFullScreenContentSize                                   func(window Window, proposedSize foundation.Size) foundation.Size
	Window_WillUseFullScreenPresentationOptions                           func(window Window, proposedOptions ApplicationPresentationOptions) ApplicationPresentationOptions
	WindowWillEnterFullScreen                                             func(notification foundation.Notification)
	WindowDidEnterFullScreen                                              func(notification foundation.Notification)
	WindowWillExitFullScreen                                              func(notification foundation.Notification)
	WindowDidExitFullScreen                                               func(notification foundation.Notification)
	CustomWindowsToEnterFullScreenForWindow                               func(window Window) []Window
	CustomWindowsToEnterFullScreenForWindow_OnScreen                      func(window Window, screen Screen) []Window
	Window_StartCustomAnimationToEnterFullScreenWithDuration              func(window Window, duration foundation.TimeInterval)
	Window_StartCustomAnimationToEnterFullScreenOnScreen_WithDuration     func(window Window, screen Screen, duration foundation.TimeInterval)
	WindowDidFailToEnterFullScreen                                        func(window Window)
	CustomWindowsToExitFullScreenForWindow                                func(window Window) []Window
	Window_StartCustomAnimationToExitFullScreenWithDuration               func(window Window, duration foundation.TimeInterval)
	WindowDidFailToExitFullScreen                                         func(window Window)
	WindowWillMove                                                        func(notification foundation.Notification)
	WindowDidMove                                                         func(notification foundation.Notification)
	WindowDidChangeScreen                                                 func(notification foundation.Notification)
	WindowDidChangeScreenProfile                                          func(notification foundation.Notification)
	WindowDidChangeBackingProperties                                      func(notification foundation.Notification)
	WindowShouldClose                                                     func(sender Window) bool
	WindowWillClose                                                       func(notification foundation.Notification)
	WindowDidBecomeKey                                                    func(notification foundation.Notification)
	WindowDidResignKey                                                    func(notification foundation.Notification)
	WindowDidBecomeMain                                                   func(notification foundation.Notification)
	WindowDidResignMain                                                   func(notification foundation.Notification)
	WindowWillReturnFieldEditor_ToObject                                  func(sender Window, client objc.Object) objc.Object
	WindowDidUpdate                                                       func(notification foundation.Notification)
	WindowDidExpose                                                       func(notification foundation.Notification)
	WindowDidChangeOcclusionState                                         func(notification foundation.Notification)
	Window_ShouldDragDocumentWithEvent_From_WithPasteboard                func(window Window, event Event, dragImageLocation foundation.Point, pasteboard Pasteboard) bool
	WindowWillReturnUndoManager                                           func(window Window) foundation.UndoManager
	Window_ShouldPopUpDocumentPathMenu                                    func(window Window, menu Menu) bool
	Window_WillEncodeRestorableState                                      func(window Window, state foundation.Coder)
	Window_DidDecodeRestorableState                                       func(window Window, state foundation.Coder)
	Window_WillResizeForVersionBrowserWithMaxPreferredSize_MaxAllowedSize func(window Window, maxPreferredFrameSize foundation.Size, maxAllowedFrameSize foundation.Size) foundation.Size
	WindowWillEnterVersionBrowser                                         func(notification foundation.Notification)
	WindowDidEnterVersionBrowser                                          func(notification foundation.Notification)
	WindowWillExitVersionBrowser                                          func(notification foundation.Notification)
	WindowDidExitVersionBrowser                                           func(notification foundation.Notification)
}

func (delegate *WindowDelegate) ToObjc() objc.Object {
	h := cgo.NewHandle(delegate)
	ptr := C.WrapWindowDelegate(C.uintptr_t(h))
	return objc.MakeObject(ptr)
}

//export windowDelegate_Window_WillPositionSheet_UsingRect
func windowDelegate_Window_WillPositionSheet_UsingRect(hp C.uintptr_t, window unsafe.Pointer, sheet unsafe.Pointer, rect C.CGRect) C.CGRect {
	delegate := cgo.Handle(hp).Value().(*WindowDelegate)
	result := delegate.Window_WillPositionSheet_UsingRect(MakeWindow(window), MakeWindow(sheet), foundation.Rect(coregraphics.FromCGRectPointer(unsafe.Pointer(&rect))))
	return *(*C.CGRect)(coregraphics.ToCGRectPointer(coregraphics.Rect(result)))
}

//export windowDelegate_WindowWillBeginSheet
func windowDelegate_WindowWillBeginSheet(hp C.uintptr_t, notification unsafe.Pointer) {
	delegate := cgo.Handle(hp).Value().(*WindowDelegate)
	delegate.WindowWillBeginSheet(foundation.MakeNotification(notification))
}

//export windowDelegate_WindowDidEndSheet
func windowDelegate_WindowDidEndSheet(hp C.uintptr_t, notification unsafe.Pointer) {
	delegate := cgo.Handle(hp).Value().(*WindowDelegate)
	delegate.WindowDidEndSheet(foundation.MakeNotification(notification))
}

//export windowDelegate_WindowWillResize_ToSize
func windowDelegate_WindowWillResize_ToSize(hp C.uintptr_t, sender unsafe.Pointer, frameSize C.CGSize) C.CGSize {
	delegate := cgo.Handle(hp).Value().(*WindowDelegate)
	result := delegate.WindowWillResize_ToSize(MakeWindow(sender), foundation.Size(coregraphics.FromCGSizePointer(unsafe.Pointer(&frameSize))))
	return *(*C.CGSize)(coregraphics.ToCGSizePointer(coregraphics.Size(result)))
}

//export windowDelegate_WindowDidResize
func windowDelegate_WindowDidResize(hp C.uintptr_t, notification unsafe.Pointer) {
	delegate := cgo.Handle(hp).Value().(*WindowDelegate)
	delegate.WindowDidResize(foundation.MakeNotification(notification))
}

//export windowDelegate_WindowWillStartLiveResize
func windowDelegate_WindowWillStartLiveResize(hp C.uintptr_t, notification unsafe.Pointer) {
	delegate := cgo.Handle(hp).Value().(*WindowDelegate)
	delegate.WindowWillStartLiveResize(foundation.MakeNotification(notification))
}

//export windowDelegate_WindowDidEndLiveResize
func windowDelegate_WindowDidEndLiveResize(hp C.uintptr_t, notification unsafe.Pointer) {
	delegate := cgo.Handle(hp).Value().(*WindowDelegate)
	delegate.WindowDidEndLiveResize(foundation.MakeNotification(notification))
}

//export windowDelegate_WindowWillMiniaturize
func windowDelegate_WindowWillMiniaturize(hp C.uintptr_t, notification unsafe.Pointer) {
	delegate := cgo.Handle(hp).Value().(*WindowDelegate)
	delegate.WindowWillMiniaturize(foundation.MakeNotification(notification))
}

//export windowDelegate_WindowDidMiniaturize
func windowDelegate_WindowDidMiniaturize(hp C.uintptr_t, notification unsafe.Pointer) {
	delegate := cgo.Handle(hp).Value().(*WindowDelegate)
	delegate.WindowDidMiniaturize(foundation.MakeNotification(notification))
}

//export windowDelegate_WindowDidDeminiaturize
func windowDelegate_WindowDidDeminiaturize(hp C.uintptr_t, notification unsafe.Pointer) {
	delegate := cgo.Handle(hp).Value().(*WindowDelegate)
	delegate.WindowDidDeminiaturize(foundation.MakeNotification(notification))
}

//export windowDelegate_WindowWillUseStandardFrame_DefaultFrame
func windowDelegate_WindowWillUseStandardFrame_DefaultFrame(hp C.uintptr_t, window unsafe.Pointer, newFrame C.CGRect) C.CGRect {
	delegate := cgo.Handle(hp).Value().(*WindowDelegate)
	result := delegate.WindowWillUseStandardFrame_DefaultFrame(MakeWindow(window), foundation.Rect(coregraphics.FromCGRectPointer(unsafe.Pointer(&newFrame))))
	return *(*C.CGRect)(coregraphics.ToCGRectPointer(coregraphics.Rect(result)))
}

//export windowDelegate_WindowShouldZoom_ToFrame
func windowDelegate_WindowShouldZoom_ToFrame(hp C.uintptr_t, window unsafe.Pointer, newFrame C.CGRect) C.bool {
	delegate := cgo.Handle(hp).Value().(*WindowDelegate)
	result := delegate.WindowShouldZoom_ToFrame(MakeWindow(window), foundation.Rect(coregraphics.FromCGRectPointer(unsafe.Pointer(&newFrame))))
	return C.bool(result)
}

//export windowDelegate_Window_WillUseFullScreenContentSize
func windowDelegate_Window_WillUseFullScreenContentSize(hp C.uintptr_t, window unsafe.Pointer, proposedSize C.CGSize) C.CGSize {
	delegate := cgo.Handle(hp).Value().(*WindowDelegate)
	result := delegate.Window_WillUseFullScreenContentSize(MakeWindow(window), foundation.Size(coregraphics.FromCGSizePointer(unsafe.Pointer(&proposedSize))))
	return *(*C.CGSize)(coregraphics.ToCGSizePointer(coregraphics.Size(result)))
}

//export windowDelegate_Window_WillUseFullScreenPresentationOptions
func windowDelegate_Window_WillUseFullScreenPresentationOptions(hp C.uintptr_t, window unsafe.Pointer, proposedOptions C.uint) C.uint {
	delegate := cgo.Handle(hp).Value().(*WindowDelegate)
	result := delegate.Window_WillUseFullScreenPresentationOptions(MakeWindow(window), ApplicationPresentationOptions(uint(proposedOptions)))
	return C.uint(uint(result))
}

//export windowDelegate_WindowWillEnterFullScreen
func windowDelegate_WindowWillEnterFullScreen(hp C.uintptr_t, notification unsafe.Pointer) {
	delegate := cgo.Handle(hp).Value().(*WindowDelegate)
	delegate.WindowWillEnterFullScreen(foundation.MakeNotification(notification))
}

//export windowDelegate_WindowDidEnterFullScreen
func windowDelegate_WindowDidEnterFullScreen(hp C.uintptr_t, notification unsafe.Pointer) {
	delegate := cgo.Handle(hp).Value().(*WindowDelegate)
	delegate.WindowDidEnterFullScreen(foundation.MakeNotification(notification))
}

//export windowDelegate_WindowWillExitFullScreen
func windowDelegate_WindowWillExitFullScreen(hp C.uintptr_t, notification unsafe.Pointer) {
	delegate := cgo.Handle(hp).Value().(*WindowDelegate)
	delegate.WindowWillExitFullScreen(foundation.MakeNotification(notification))
}

//export windowDelegate_WindowDidExitFullScreen
func windowDelegate_WindowDidExitFullScreen(hp C.uintptr_t, notification unsafe.Pointer) {
	delegate := cgo.Handle(hp).Value().(*WindowDelegate)
	delegate.WindowDidExitFullScreen(foundation.MakeNotification(notification))
}

//export windowDelegate_CustomWindowsToEnterFullScreenForWindow
func windowDelegate_CustomWindowsToEnterFullScreenForWindow(hp C.uintptr_t, window unsafe.Pointer) C.Array {
	delegate := cgo.Handle(hp).Value().(*WindowDelegate)
	result := delegate.CustomWindowsToEnterFullScreenForWindow(MakeWindow(window))
	var cResult C.Array
	if len(result) > 0 {
		cResultData := make([]unsafe.Pointer, len(result))
		for idx, v := range result {
			cResultData[idx] = objc.ExtractPtr(v)
		}
		cResult.data = unsafe.Pointer(&cResultData[0])
		cResult.len = C.int(len(result))
	}
	return cResult
}

//export windowDelegate_CustomWindowsToEnterFullScreenForWindow_OnScreen
func windowDelegate_CustomWindowsToEnterFullScreenForWindow_OnScreen(hp C.uintptr_t, window unsafe.Pointer, screen unsafe.Pointer) C.Array {
	delegate := cgo.Handle(hp).Value().(*WindowDelegate)
	result := delegate.CustomWindowsToEnterFullScreenForWindow_OnScreen(MakeWindow(window), MakeScreen(screen))
	var cResult C.Array
	if len(result) > 0 {
		cResultData := make([]unsafe.Pointer, len(result))
		for idx, v := range result {
			cResultData[idx] = objc.ExtractPtr(v)
		}
		cResult.data = unsafe.Pointer(&cResultData[0])
		cResult.len = C.int(len(result))
	}
	return cResult
}

//export windowDelegate_Window_StartCustomAnimationToEnterFullScreenWithDuration
func windowDelegate_Window_StartCustomAnimationToEnterFullScreenWithDuration(hp C.uintptr_t, window unsafe.Pointer, duration C.double) {
	delegate := cgo.Handle(hp).Value().(*WindowDelegate)
	delegate.Window_StartCustomAnimationToEnterFullScreenWithDuration(MakeWindow(window), foundation.TimeInterval(float64(duration)))
}

//export windowDelegate_Window_StartCustomAnimationToEnterFullScreenOnScreen_WithDuration
func windowDelegate_Window_StartCustomAnimationToEnterFullScreenOnScreen_WithDuration(hp C.uintptr_t, window unsafe.Pointer, screen unsafe.Pointer, duration C.double) {
	delegate := cgo.Handle(hp).Value().(*WindowDelegate)
	delegate.Window_StartCustomAnimationToEnterFullScreenOnScreen_WithDuration(MakeWindow(window), MakeScreen(screen), foundation.TimeInterval(float64(duration)))
}

//export windowDelegate_WindowDidFailToEnterFullScreen
func windowDelegate_WindowDidFailToEnterFullScreen(hp C.uintptr_t, window unsafe.Pointer) {
	delegate := cgo.Handle(hp).Value().(*WindowDelegate)
	delegate.WindowDidFailToEnterFullScreen(MakeWindow(window))
}

//export windowDelegate_CustomWindowsToExitFullScreenForWindow
func windowDelegate_CustomWindowsToExitFullScreenForWindow(hp C.uintptr_t, window unsafe.Pointer) C.Array {
	delegate := cgo.Handle(hp).Value().(*WindowDelegate)
	result := delegate.CustomWindowsToExitFullScreenForWindow(MakeWindow(window))
	var cResult C.Array
	if len(result) > 0 {
		cResultData := make([]unsafe.Pointer, len(result))
		for idx, v := range result {
			cResultData[idx] = objc.ExtractPtr(v)
		}
		cResult.data = unsafe.Pointer(&cResultData[0])
		cResult.len = C.int(len(result))
	}
	return cResult
}

//export windowDelegate_Window_StartCustomAnimationToExitFullScreenWithDuration
func windowDelegate_Window_StartCustomAnimationToExitFullScreenWithDuration(hp C.uintptr_t, window unsafe.Pointer, duration C.double) {
	delegate := cgo.Handle(hp).Value().(*WindowDelegate)
	delegate.Window_StartCustomAnimationToExitFullScreenWithDuration(MakeWindow(window), foundation.TimeInterval(float64(duration)))
}

//export windowDelegate_WindowDidFailToExitFullScreen
func windowDelegate_WindowDidFailToExitFullScreen(hp C.uintptr_t, window unsafe.Pointer) {
	delegate := cgo.Handle(hp).Value().(*WindowDelegate)
	delegate.WindowDidFailToExitFullScreen(MakeWindow(window))
}

//export windowDelegate_WindowWillMove
func windowDelegate_WindowWillMove(hp C.uintptr_t, notification unsafe.Pointer) {
	delegate := cgo.Handle(hp).Value().(*WindowDelegate)
	delegate.WindowWillMove(foundation.MakeNotification(notification))
}

//export windowDelegate_WindowDidMove
func windowDelegate_WindowDidMove(hp C.uintptr_t, notification unsafe.Pointer) {
	delegate := cgo.Handle(hp).Value().(*WindowDelegate)
	delegate.WindowDidMove(foundation.MakeNotification(notification))
}

//export windowDelegate_WindowDidChangeScreen
func windowDelegate_WindowDidChangeScreen(hp C.uintptr_t, notification unsafe.Pointer) {
	delegate := cgo.Handle(hp).Value().(*WindowDelegate)
	delegate.WindowDidChangeScreen(foundation.MakeNotification(notification))
}

//export windowDelegate_WindowDidChangeScreenProfile
func windowDelegate_WindowDidChangeScreenProfile(hp C.uintptr_t, notification unsafe.Pointer) {
	delegate := cgo.Handle(hp).Value().(*WindowDelegate)
	delegate.WindowDidChangeScreenProfile(foundation.MakeNotification(notification))
}

//export windowDelegate_WindowDidChangeBackingProperties
func windowDelegate_WindowDidChangeBackingProperties(hp C.uintptr_t, notification unsafe.Pointer) {
	delegate := cgo.Handle(hp).Value().(*WindowDelegate)
	delegate.WindowDidChangeBackingProperties(foundation.MakeNotification(notification))
}

//export windowDelegate_WindowShouldClose
func windowDelegate_WindowShouldClose(hp C.uintptr_t, sender unsafe.Pointer) C.bool {
	delegate := cgo.Handle(hp).Value().(*WindowDelegate)
	result := delegate.WindowShouldClose(MakeWindow(sender))
	return C.bool(result)
}

//export windowDelegate_WindowWillClose
func windowDelegate_WindowWillClose(hp C.uintptr_t, notification unsafe.Pointer) {
	delegate := cgo.Handle(hp).Value().(*WindowDelegate)
	delegate.WindowWillClose(foundation.MakeNotification(notification))
}

//export windowDelegate_WindowDidBecomeKey
func windowDelegate_WindowDidBecomeKey(hp C.uintptr_t, notification unsafe.Pointer) {
	delegate := cgo.Handle(hp).Value().(*WindowDelegate)
	delegate.WindowDidBecomeKey(foundation.MakeNotification(notification))
}

//export windowDelegate_WindowDidResignKey
func windowDelegate_WindowDidResignKey(hp C.uintptr_t, notification unsafe.Pointer) {
	delegate := cgo.Handle(hp).Value().(*WindowDelegate)
	delegate.WindowDidResignKey(foundation.MakeNotification(notification))
}

//export windowDelegate_WindowDidBecomeMain
func windowDelegate_WindowDidBecomeMain(hp C.uintptr_t, notification unsafe.Pointer) {
	delegate := cgo.Handle(hp).Value().(*WindowDelegate)
	delegate.WindowDidBecomeMain(foundation.MakeNotification(notification))
}

//export windowDelegate_WindowDidResignMain
func windowDelegate_WindowDidResignMain(hp C.uintptr_t, notification unsafe.Pointer) {
	delegate := cgo.Handle(hp).Value().(*WindowDelegate)
	delegate.WindowDidResignMain(foundation.MakeNotification(notification))
}

//export windowDelegate_WindowWillReturnFieldEditor_ToObject
func windowDelegate_WindowWillReturnFieldEditor_ToObject(hp C.uintptr_t, sender unsafe.Pointer, client unsafe.Pointer) unsafe.Pointer {
	delegate := cgo.Handle(hp).Value().(*WindowDelegate)
	result := delegate.WindowWillReturnFieldEditor_ToObject(MakeWindow(sender), objc.MakeObject(client))
	return objc.ExtractPtr(result)
}

//export windowDelegate_WindowDidUpdate
func windowDelegate_WindowDidUpdate(hp C.uintptr_t, notification unsafe.Pointer) {
	delegate := cgo.Handle(hp).Value().(*WindowDelegate)
	delegate.WindowDidUpdate(foundation.MakeNotification(notification))
}

//export windowDelegate_WindowDidExpose
func windowDelegate_WindowDidExpose(hp C.uintptr_t, notification unsafe.Pointer) {
	delegate := cgo.Handle(hp).Value().(*WindowDelegate)
	delegate.WindowDidExpose(foundation.MakeNotification(notification))
}

//export windowDelegate_WindowDidChangeOcclusionState
func windowDelegate_WindowDidChangeOcclusionState(hp C.uintptr_t, notification unsafe.Pointer) {
	delegate := cgo.Handle(hp).Value().(*WindowDelegate)
	delegate.WindowDidChangeOcclusionState(foundation.MakeNotification(notification))
}

//export windowDelegate_Window_ShouldDragDocumentWithEvent_From_WithPasteboard
func windowDelegate_Window_ShouldDragDocumentWithEvent_From_WithPasteboard(hp C.uintptr_t, window unsafe.Pointer, event unsafe.Pointer, dragImageLocation C.CGPoint, pasteboard unsafe.Pointer) C.bool {
	delegate := cgo.Handle(hp).Value().(*WindowDelegate)
	result := delegate.Window_ShouldDragDocumentWithEvent_From_WithPasteboard(MakeWindow(window), MakeEvent(event), foundation.Point(coregraphics.FromCGPointPointer(unsafe.Pointer(&dragImageLocation))), MakePasteboard(pasteboard))
	return C.bool(result)
}

//export windowDelegate_WindowWillReturnUndoManager
func windowDelegate_WindowWillReturnUndoManager(hp C.uintptr_t, window unsafe.Pointer) unsafe.Pointer {
	delegate := cgo.Handle(hp).Value().(*WindowDelegate)
	result := delegate.WindowWillReturnUndoManager(MakeWindow(window))
	return objc.ExtractPtr(result)
}

//export windowDelegate_Window_ShouldPopUpDocumentPathMenu
func windowDelegate_Window_ShouldPopUpDocumentPathMenu(hp C.uintptr_t, window unsafe.Pointer, menu unsafe.Pointer) C.bool {
	delegate := cgo.Handle(hp).Value().(*WindowDelegate)
	result := delegate.Window_ShouldPopUpDocumentPathMenu(MakeWindow(window), MakeMenu(menu))
	return C.bool(result)
}

//export windowDelegate_Window_WillEncodeRestorableState
func windowDelegate_Window_WillEncodeRestorableState(hp C.uintptr_t, window unsafe.Pointer, state unsafe.Pointer) {
	delegate := cgo.Handle(hp).Value().(*WindowDelegate)
	delegate.Window_WillEncodeRestorableState(MakeWindow(window), foundation.MakeCoder(state))
}

//export windowDelegate_Window_DidDecodeRestorableState
func windowDelegate_Window_DidDecodeRestorableState(hp C.uintptr_t, window unsafe.Pointer, state unsafe.Pointer) {
	delegate := cgo.Handle(hp).Value().(*WindowDelegate)
	delegate.Window_DidDecodeRestorableState(MakeWindow(window), foundation.MakeCoder(state))
}

//export windowDelegate_Window_WillResizeForVersionBrowserWithMaxPreferredSize_MaxAllowedSize
func windowDelegate_Window_WillResizeForVersionBrowserWithMaxPreferredSize_MaxAllowedSize(hp C.uintptr_t, window unsafe.Pointer, maxPreferredFrameSize C.CGSize, maxAllowedFrameSize C.CGSize) C.CGSize {
	delegate := cgo.Handle(hp).Value().(*WindowDelegate)
	result := delegate.Window_WillResizeForVersionBrowserWithMaxPreferredSize_MaxAllowedSize(MakeWindow(window), foundation.Size(coregraphics.FromCGSizePointer(unsafe.Pointer(&maxPreferredFrameSize))), foundation.Size(coregraphics.FromCGSizePointer(unsafe.Pointer(&maxAllowedFrameSize))))
	return *(*C.CGSize)(coregraphics.ToCGSizePointer(coregraphics.Size(result)))
}

//export windowDelegate_WindowWillEnterVersionBrowser
func windowDelegate_WindowWillEnterVersionBrowser(hp C.uintptr_t, notification unsafe.Pointer) {
	delegate := cgo.Handle(hp).Value().(*WindowDelegate)
	delegate.WindowWillEnterVersionBrowser(foundation.MakeNotification(notification))
}

//export windowDelegate_WindowDidEnterVersionBrowser
func windowDelegate_WindowDidEnterVersionBrowser(hp C.uintptr_t, notification unsafe.Pointer) {
	delegate := cgo.Handle(hp).Value().(*WindowDelegate)
	delegate.WindowDidEnterVersionBrowser(foundation.MakeNotification(notification))
}

//export windowDelegate_WindowWillExitVersionBrowser
func windowDelegate_WindowWillExitVersionBrowser(hp C.uintptr_t, notification unsafe.Pointer) {
	delegate := cgo.Handle(hp).Value().(*WindowDelegate)
	delegate.WindowWillExitVersionBrowser(foundation.MakeNotification(notification))
}

//export windowDelegate_WindowDidExitVersionBrowser
func windowDelegate_WindowDidExitVersionBrowser(hp C.uintptr_t, notification unsafe.Pointer) {
	delegate := cgo.Handle(hp).Value().(*WindowDelegate)
	delegate.WindowDidExitVersionBrowser(foundation.MakeNotification(notification))
}

//export WindowDelegate_RespondsTo
func WindowDelegate_RespondsTo(hp C.uintptr_t, selectorPtr unsafe.Pointer) bool {
	sel := objc.Selector(selectorPtr)
	selName := objc.Sel_GetName(sel)
	delegate := cgo.Handle(hp).Value().(*WindowDelegate)
	switch selName {
	case "window:willPositionSheet:usingRect:":
		return delegate.Window_WillPositionSheet_UsingRect != nil
	case "windowWillBeginSheet:":
		return delegate.WindowWillBeginSheet != nil
	case "windowDidEndSheet:":
		return delegate.WindowDidEndSheet != nil
	case "windowWillResize:toSize:":
		return delegate.WindowWillResize_ToSize != nil
	case "windowDidResize:":
		return delegate.WindowDidResize != nil
	case "windowWillStartLiveResize:":
		return delegate.WindowWillStartLiveResize != nil
	case "windowDidEndLiveResize:":
		return delegate.WindowDidEndLiveResize != nil
	case "windowWillMiniaturize:":
		return delegate.WindowWillMiniaturize != nil
	case "windowDidMiniaturize:":
		return delegate.WindowDidMiniaturize != nil
	case "windowDidDeminiaturize:":
		return delegate.WindowDidDeminiaturize != nil
	case "windowWillUseStandardFrame:defaultFrame:":
		return delegate.WindowWillUseStandardFrame_DefaultFrame != nil
	case "windowShouldZoom:toFrame:":
		return delegate.WindowShouldZoom_ToFrame != nil
	case "window:willUseFullScreenContentSize:":
		return delegate.Window_WillUseFullScreenContentSize != nil
	case "window:willUseFullScreenPresentationOptions:":
		return delegate.Window_WillUseFullScreenPresentationOptions != nil
	case "windowWillEnterFullScreen:":
		return delegate.WindowWillEnterFullScreen != nil
	case "windowDidEnterFullScreen:":
		return delegate.WindowDidEnterFullScreen != nil
	case "windowWillExitFullScreen:":
		return delegate.WindowWillExitFullScreen != nil
	case "windowDidExitFullScreen:":
		return delegate.WindowDidExitFullScreen != nil
	case "customWindowsToEnterFullScreenForWindow:":
		return delegate.CustomWindowsToEnterFullScreenForWindow != nil
	case "customWindowsToEnterFullScreenForWindow:onScreen:":
		return delegate.CustomWindowsToEnterFullScreenForWindow_OnScreen != nil
	case "window:startCustomAnimationToEnterFullScreenWithDuration:":
		return delegate.Window_StartCustomAnimationToEnterFullScreenWithDuration != nil
	case "window:startCustomAnimationToEnterFullScreenOnScreen:withDuration:":
		return delegate.Window_StartCustomAnimationToEnterFullScreenOnScreen_WithDuration != nil
	case "windowDidFailToEnterFullScreen:":
		return delegate.WindowDidFailToEnterFullScreen != nil
	case "customWindowsToExitFullScreenForWindow:":
		return delegate.CustomWindowsToExitFullScreenForWindow != nil
	case "window:startCustomAnimationToExitFullScreenWithDuration:":
		return delegate.Window_StartCustomAnimationToExitFullScreenWithDuration != nil
	case "windowDidFailToExitFullScreen:":
		return delegate.WindowDidFailToExitFullScreen != nil
	case "windowWillMove:":
		return delegate.WindowWillMove != nil
	case "windowDidMove:":
		return delegate.WindowDidMove != nil
	case "windowDidChangeScreen:":
		return delegate.WindowDidChangeScreen != nil
	case "windowDidChangeScreenProfile:":
		return delegate.WindowDidChangeScreenProfile != nil
	case "windowDidChangeBackingProperties:":
		return delegate.WindowDidChangeBackingProperties != nil
	case "windowShouldClose:":
		return delegate.WindowShouldClose != nil
	case "windowWillClose:":
		return delegate.WindowWillClose != nil
	case "windowDidBecomeKey:":
		return delegate.WindowDidBecomeKey != nil
	case "windowDidResignKey:":
		return delegate.WindowDidResignKey != nil
	case "windowDidBecomeMain:":
		return delegate.WindowDidBecomeMain != nil
	case "windowDidResignMain:":
		return delegate.WindowDidResignMain != nil
	case "windowWillReturnFieldEditor:toObject:":
		return delegate.WindowWillReturnFieldEditor_ToObject != nil
	case "windowDidUpdate:":
		return delegate.WindowDidUpdate != nil
	case "windowDidExpose:":
		return delegate.WindowDidExpose != nil
	case "windowDidChangeOcclusionState:":
		return delegate.WindowDidChangeOcclusionState != nil
	case "window:shouldDragDocumentWithEvent:from:withPasteboard:":
		return delegate.Window_ShouldDragDocumentWithEvent_From_WithPasteboard != nil
	case "windowWillReturnUndoManager:":
		return delegate.WindowWillReturnUndoManager != nil
	case "window:shouldPopUpDocumentPathMenu:":
		return delegate.Window_ShouldPopUpDocumentPathMenu != nil
	case "window:willEncodeRestorableState:":
		return delegate.Window_WillEncodeRestorableState != nil
	case "window:didDecodeRestorableState:":
		return delegate.Window_DidDecodeRestorableState != nil
	case "window:willResizeForVersionBrowserWithMaxPreferredSize:maxAllowedSize:":
		return delegate.Window_WillResizeForVersionBrowserWithMaxPreferredSize_MaxAllowedSize != nil
	case "windowWillEnterVersionBrowser:":
		return delegate.WindowWillEnterVersionBrowser != nil
	case "windowDidEnterVersionBrowser:":
		return delegate.WindowDidEnterVersionBrowser != nil
	case "windowWillExitVersionBrowser:":
		return delegate.WindowWillExitVersionBrowser != nil
	case "windowDidExitVersionBrowser:":
		return delegate.WindowDidExitVersionBrowser != nil
	default:
		return false
	}
}
