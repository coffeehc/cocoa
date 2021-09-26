package appkit

// #include "view.h"
import "C"
import (
	"github.com/hsiafan/cocoa/coregraphics"
	"github.com/hsiafan/cocoa/foundation"
	"github.com/hsiafan/cocoa/objc"
	"unsafe"
)

type View interface {
	Responder
	PrepareForReuse()
	AddSubview(view View)
	AddSubview_Positioned_RelativeTo(view View, place WindowOrderingMode, otherView View)
	DidAddSubview(subview View)
	RemoveFromSuperview()
	RemoveFromSuperviewWithoutNeedingDisplay()
	ReplaceSubview_With(oldView View, newView View)
	IsDescendantOf(view View) bool
	AncestorSharedWithView(view View) View
	ViewDidMoveToSuperview()
	ViewDidMoveToWindow()
	ViewWillMoveToSuperview(newSuperview View)
	ViewWillMoveToWindow(newWindow Window)
	WillRemoveSubview(subview View)
	SetFrameOrigin(newOrigin foundation.Point)
	SetFrameSize(newSize foundation.Size)
	SetBoundsOrigin(newOrigin foundation.Point)
	SetBoundsSize(newSize foundation.Size)
	UpdateLayer()
	DrawRect(dirtyRect foundation.Rect)
	NeedsToDrawRect(rect foundation.Rect) bool
	BitmapImageRepForCachingDisplayInRect(rect foundation.Rect) BitmapImageRep
	CacheDisplayInRect_ToBitmapImageRep(rect foundation.Rect, bitmapImageRep BitmapImageRep)
	Print(sender objc.Object)
	BeginPageInRect_AtPlacement(rect foundation.Rect, location foundation.Point)
	DataWithEPSInsideRect(rect foundation.Rect) []byte
	DataWithPDFInsideRect(rect foundation.Rect) []byte
	WriteEPSInsideRect_ToPasteboard(rect foundation.Rect, pasteboard Pasteboard)
	WritePDFInsideRect_ToPasteboard(rect foundation.Rect, pasteboard Pasteboard)
	DrawPageBorderWithSize(borderSize foundation.Size)
	RectForPage(page int) foundation.Rect
	LocationOfPrintRect(rect foundation.Rect) foundation.Point
	SetNeedsDisplayInRect(invalidRect foundation.Rect)
	Display()
	DisplayRect(rect foundation.Rect)
	DisplayRectIgnoringOpacity(rect foundation.Rect)
	DisplayRectIgnoringOpacity_InContext(rect foundation.Rect, context GraphicsContext)
	DisplayIfNeeded()
	DisplayIfNeededInRect(rect foundation.Rect)
	DisplayIfNeededIgnoringOpacity()
	DisplayIfNeededInRectIgnoringOpacity(rect foundation.Rect)
	TranslateRectsNeedingDisplayInRect_By(clipRect foundation.Rect, delta foundation.Size)
	ViewWillDraw()
	ConvertPointFromBacking(point foundation.Point) foundation.Point
	ConvertPointToBacking(point foundation.Point) foundation.Point
	ConvertPointFromLayer(point foundation.Point) foundation.Point
	ConvertPointToLayer(point foundation.Point) foundation.Point
	ConvertRectFromBacking(rect foundation.Rect) foundation.Rect
	ConvertRectToBacking(rect foundation.Rect) foundation.Rect
	ConvertRectFromLayer(rect foundation.Rect) foundation.Rect
	ConvertRectToLayer(rect foundation.Rect) foundation.Rect
	ConvertSizeFromBacking(size foundation.Size) foundation.Size
	ConvertSizeToBacking(size foundation.Size) foundation.Size
	ConvertSizeFromLayer(size foundation.Size) foundation.Size
	ConvertSizeToLayer(size foundation.Size) foundation.Size
	ConvertPoint_FromView(point foundation.Point, view View) foundation.Point
	ConvertPoint_ToView(point foundation.Point, view View) foundation.Point
	ConvertSize_FromView(size foundation.Size, view View) foundation.Size
	ConvertSize_ToView(size foundation.Size, view View) foundation.Size
	ConvertRect_FromView(rect foundation.Rect, view View) foundation.Rect
	ConvertRect_ToView(rect foundation.Rect, view View) foundation.Rect
	CenterScanRect(rect foundation.Rect) foundation.Rect
	TranslateOriginToPoint(translation foundation.Point)
	ScaleUnitSquareToSize(newUnitSize foundation.Size)
	RotateByAngle(angle coregraphics.Float)
	ResizeSubviewsWithOldSize(oldSize foundation.Size)
	ResizeWithOldSuperviewSize(oldSize foundation.Size)
	AddConstraint(constraint LayoutConstraint)
	AddConstraints(constraints []LayoutConstraint)
	RemoveConstraint(constraint LayoutConstraint)
	RemoveConstraints(constraints []LayoutConstraint)
	AddLayoutGuide(guide LayoutGuide)
	RemoveLayoutGuide(guide LayoutGuide)
	InvalidateIntrinsicContentSize()
	ContentCompressionResistancePriorityForOrientation(orientation LayoutConstraintOrientation) LayoutPriority
	SetContentCompressionResistancePriority_ForOrientation(priority LayoutPriority, orientation LayoutConstraintOrientation)
	ContentHuggingPriorityForOrientation(orientation LayoutConstraintOrientation) LayoutPriority
	SetContentHuggingPriority_ForOrientation(priority LayoutPriority, orientation LayoutConstraintOrientation)
	AlignmentRectForFrame(frame foundation.Rect) foundation.Rect
	FrameForAlignmentRect(alignmentRect foundation.Rect) foundation.Rect
	Layout()
	LayoutSubtreeIfNeeded()
	UpdateConstraints()
	UpdateConstraintsForSubtreeIfNeeded()
	ConstraintsAffectingLayoutForOrientation(orientation LayoutConstraintOrientation) []LayoutConstraint
	ExerciseAmbiguityInLayout()
	DrawFocusRingMask()
	NoteFocusRingMaskChanged()
	SetKeyboardFocusRingNeedsDisplayInRect(rect foundation.Rect)
	EnterFullScreenMode_WithOptions(screen Screen, options map[ViewFullScreenModeOptionKey]objc.Object) bool
	ExitFullScreenModeWithOptions(options map[ViewFullScreenModeOptionKey]objc.Object)
	ViewDidHide()
	ViewDidUnhide()
	ViewWillStartLiveResize()
	ViewDidEndLiveResize()
	AddGestureRecognizer(gestureRecognizer GestureRecognizer)
	RemoveGestureRecognizer(gestureRecognizer GestureRecognizer)
	AcceptsFirstMouse(event Event) bool
	HitTest(point foundation.Point) View
	Mouse_InRect(point foundation.Point, rect foundation.Rect) bool
	PrepareContentInRect(rect foundation.Rect)
	ScrollPoint(point foundation.Point)
	ScrollRectToVisible(rect foundation.Rect) bool
	Autoscroll(event Event) bool
	AdjustScroll(newVisible foundation.Rect) foundation.Rect
	ScrollClipView_ToPoint(clipView ClipView, point foundation.Point)
	ReflectScrolledClipView(clipView ClipView)
	RegisterForDraggedTypes(newTypes []PasteboardType)
	UnregisterDraggedTypes()
	BeginDraggingSessionWithItems_Event_Source(items []DraggingItem, event Event, source objc.Object) DraggingSession
	ShouldDelayWindowOrderingForEvent(event Event) bool
	RectForSmartMagnificationAtPoint_InRect(location foundation.Point, visibleRect foundation.Rect) foundation.Rect
	ViewDidChangeBackingProperties()
	ViewWithTag(tag int) View
	RemoveAllToolTips()
	RemoveToolTip(tag ToolTipTag)
	RemoveTrackingRect(tag TrackingRectTag)
	AddTrackingArea(trackingArea TrackingArea)
	RemoveTrackingArea(trackingArea TrackingArea)
	UpdateTrackingAreas()
	AddCursorRect_Cursor(rect foundation.Rect, object Cursor)
	RemoveCursorRect_Cursor(rect foundation.Rect, object Cursor)
	DiscardCursorRects()
	ResetCursorRects()
	MenuForEvent(event Event) Menu
	WillOpenMenu_WithEvent(menu Menu, event Event)
	DidCloseMenu_WithEvent(menu Menu, event Event)
	BeginDocument()
	EndDocument()
	EndPage()
	ShowDefinitionForAttributedString_AtPoint(attrString foundation.AttributedString, textBaselineOrigin foundation.Point)
	RulerView_DidAddMarker(ruler RulerView, marker RulerMarker)
	RulerView_DidMoveMarker(ruler RulerView, marker RulerMarker)
	RulerView_DidRemoveMarker(ruler RulerView, marker RulerMarker)
	RulerView_HandleMouseDown(ruler RulerView, event Event)
	RulerView_LocationForPoint(ruler RulerView, point foundation.Point) coregraphics.Float
	RulerView_PointForLocation(ruler RulerView, point coregraphics.Float) foundation.Point
	RulerView_ShouldAddMarker(ruler RulerView, marker RulerMarker) bool
	RulerView_ShouldMoveMarker(ruler RulerView, marker RulerMarker) bool
	RulerView_ShouldRemoveMarker(ruler RulerView, marker RulerMarker) bool
	RulerView_WillAddMarker_AtLocation(ruler RulerView, marker RulerMarker, location coregraphics.Float) coregraphics.Float
	RulerView_WillMoveMarker_ToLocation(ruler RulerView, marker RulerMarker, location coregraphics.Float) coregraphics.Float
	RulerView_WillSetClientView(ruler RulerView, newClient View)
	ViewDidChangeEffectiveAppearance()
	Superview() View
	Subviews() []View
	SetSubviews(value []View)
	Window() Window
	OpaqueAncestor() View
	EnclosingMenuItem() MenuItem
	Frame() foundation.Rect
	SetFrame(value foundation.Rect)
	FrameRotation() coregraphics.Float
	SetFrameRotation(value coregraphics.Float)
	Bounds() foundation.Rect
	SetBounds(value foundation.Rect)
	BoundsRotation() coregraphics.Float
	SetBoundsRotation(value coregraphics.Float)
	WantsLayer() bool
	SetWantsLayer(value bool)
	WantsUpdateLayer() bool
	LayerContentsPlacement() ViewLayerContentsPlacement
	SetLayerContentsPlacement(value ViewLayerContentsPlacement)
	LayerContentsRedrawPolicy() ViewLayerContentsRedrawPolicy
	SetLayerContentsRedrawPolicy(value ViewLayerContentsRedrawPolicy)
	CanDrawSubviewsIntoLayer() bool
	SetCanDrawSubviewsIntoLayer(value bool)
	LayerUsesCoreImageFilters() bool
	SetLayerUsesCoreImageFilters(value bool)
	AlphaValue() coregraphics.Float
	SetAlphaValue(value coregraphics.Float)
	FrameCenterRotation() coregraphics.Float
	SetFrameCenterRotation(value coregraphics.Float)
	Shadow() Shadow
	SetShadow(value Shadow)
	CanDrawConcurrently() bool
	SetCanDrawConcurrently(value bool)
	VisibleRect() foundation.Rect
	WantsDefaultClipping() bool
	PrintJobTitle() string
	PageHeader() foundation.AttributedString
	PageFooter() foundation.AttributedString
	HeightAdjustLimit() coregraphics.Float
	WidthAdjustLimit() coregraphics.Float
	NeedsDisplay() bool
	SetNeedsDisplay(value bool)
	IsOpaque() bool
	IsFlipped() bool
	IsRotatedFromBase() bool
	IsRotatedOrScaledFromBase() bool
	AutoresizesSubviews() bool
	SetAutoresizesSubviews(value bool)
	AutoresizingMask() AutoresizingMaskOptions
	SetAutoresizingMask(value AutoresizingMaskOptions)
	BottomAnchor() LayoutYAxisAnchor
	CenterXAnchor() LayoutXAxisAnchor
	CenterYAnchor() LayoutYAxisAnchor
	FirstBaselineAnchor() LayoutYAxisAnchor
	HeightAnchor() LayoutDimension
	LastBaselineAnchor() LayoutYAxisAnchor
	LeadingAnchor() LayoutXAxisAnchor
	LeftAnchor() LayoutXAxisAnchor
	RightAnchor() LayoutXAxisAnchor
	TopAnchor() LayoutYAxisAnchor
	TrailingAnchor() LayoutXAxisAnchor
	WidthAnchor() LayoutDimension
	Constraints() []LayoutConstraint
	LayoutGuides() []LayoutGuide
	LayoutMarginsGuide() LayoutGuide
	SafeAreaLayoutGuide() LayoutGuide
	FittingSize() foundation.Size
	IntrinsicContentSize() foundation.Size
	AlignmentRectInsets() foundation.EdgeInsets
	BaselineOffsetFromBottom() coregraphics.Float
	FirstBaselineOffsetFromTop() coregraphics.Float
	LastBaselineOffsetFromBottom() coregraphics.Float
	NeedsLayout() bool
	SetNeedsLayout(value bool)
	NeedsUpdateConstraints() bool
	SetNeedsUpdateConstraints(value bool)
	TranslatesAutoresizingMaskIntoConstraints() bool
	SetTranslatesAutoresizingMaskIntoConstraints(value bool)
	HasAmbiguousLayout() bool
	FocusRingType() FocusRingType
	SetFocusRingType(value FocusRingType)
	FocusRingMaskBounds() foundation.Rect
	AllowsVibrancy() bool
	IsInFullScreenMode() bool
	IsHidden() bool
	SetHidden(value bool)
	IsHiddenOrHasHiddenAncestor() bool
	InLiveResize() bool
	PreservesContentDuringLiveResize() bool
	RectPreservedDuringLiveResize() foundation.Rect
	GestureRecognizers() []GestureRecognizer
	SetGestureRecognizers(value []GestureRecognizer)
	MouseDownCanMoveWindow() bool
	InputContext() TextInputContext
	WantsRestingTouches() bool
	SetWantsRestingTouches(value bool)
	CanBecomeKeyView() bool
	NeedsPanelToBecomeKey() bool
	NextKeyView() View
	SetNextKeyView(value View)
	NextValidKeyView() View
	PreviousKeyView() View
	PreviousValidKeyView() View
	PreparedContentRect() foundation.Rect
	SetPreparedContentRect(value foundation.Rect)
	EnclosingScrollView() ScrollView
	RegisteredDraggedTypes() []PasteboardType
	PostsFrameChangedNotifications() bool
	SetPostsFrameChangedNotifications(value bool)
	PostsBoundsChangedNotifications() bool
	SetPostsBoundsChangedNotifications(value bool)
	Tag() int
	ToolTip() string
	SetToolTip(value string)
	TrackingAreas() []TrackingArea
	IsDrawingFindIndicator() bool
	UserInterfaceLayoutDirection() UserInterfaceLayoutDirection
	SetUserInterfaceLayoutDirection(value UserInterfaceLayoutDirection)
	PressureConfiguration() PressureConfiguration
	SetPressureConfiguration(value PressureConfiguration)
	AdditionalSafeAreaInsets() foundation.EdgeInsets
	SetAdditionalSafeAreaInsets(value foundation.EdgeInsets)
	AllowedTouchTypes() TouchTypeMask
	SetAllowedTouchTypes(value TouchTypeMask)
	CandidateListTouchBarItem() CandidateListTouchBarItem
	IsHorizontalContentSizeConstraintActive() bool
	SetHorizontalContentSizeConstraintActive(value bool)
	IsVerticalContentSizeConstraintActive() bool
	SetVerticalContentSizeConstraintActive(value bool)
	SafeAreaInsets() foundation.EdgeInsets
	SafeAreaRect() foundation.Rect
}

type NSView struct {
	NSResponder
}

func MakeView(ptr unsafe.Pointer) NSView {
	return NSView{
		NSResponder: MakeResponder(ptr),
	}
}

func (n NSView) InitWithFrame(frameRect foundation.Rect) NSView {
	result_ := C.C_NSView_InitWithFrame(n.Ptr(), *(*C.CGRect)(coregraphics.ToCGRectPointer(coregraphics.Rect(frameRect))))
	return MakeView(result_)
}

func (n NSView) InitWithCoder(coder foundation.Coder) NSView {
	result_ := C.C_NSView_InitWithCoder(n.Ptr(), objc.ExtractPtr(coder))
	return MakeView(result_)
}

func (n NSView) Init() NSView {
	result_ := C.C_NSView_Init(n.Ptr())
	return MakeView(result_)
}

func AllocView() NSView {
	result_ := C.C_NSView_AllocView()
	return MakeView(result_)
}

func NewView() NSView {
	result_ := C.C_NSView_NewView()
	return MakeView(result_)
}

func (n NSView) Autorelease() NSView {
	result_ := C.C_NSView_Autorelease(n.Ptr())
	return MakeView(result_)
}

func (n NSView) Retain() NSView {
	result_ := C.C_NSView_Retain(n.Ptr())
	return MakeView(result_)
}

func (n NSView) PrepareForReuse() {
	C.C_NSView_PrepareForReuse(n.Ptr())
}

func (n NSView) AddSubview(view View) {
	C.C_NSView_AddSubview(n.Ptr(), objc.ExtractPtr(view))
}

func (n NSView) AddSubview_Positioned_RelativeTo(view View, place WindowOrderingMode, otherView View) {
	C.C_NSView_AddSubview_Positioned_RelativeTo(n.Ptr(), objc.ExtractPtr(view), C.int(int(place)), objc.ExtractPtr(otherView))
}

func (n NSView) DidAddSubview(subview View) {
	C.C_NSView_DidAddSubview(n.Ptr(), objc.ExtractPtr(subview))
}

func (n NSView) RemoveFromSuperview() {
	C.C_NSView_RemoveFromSuperview(n.Ptr())
}

func (n NSView) RemoveFromSuperviewWithoutNeedingDisplay() {
	C.C_NSView_RemoveFromSuperviewWithoutNeedingDisplay(n.Ptr())
}

func (n NSView) ReplaceSubview_With(oldView View, newView View) {
	C.C_NSView_ReplaceSubview_With(n.Ptr(), objc.ExtractPtr(oldView), objc.ExtractPtr(newView))
}

func (n NSView) IsDescendantOf(view View) bool {
	result_ := C.C_NSView_IsDescendantOf(n.Ptr(), objc.ExtractPtr(view))
	return bool(result_)
}

func (n NSView) AncestorSharedWithView(view View) View {
	result_ := C.C_NSView_AncestorSharedWithView(n.Ptr(), objc.ExtractPtr(view))
	return MakeView(result_)
}

func (n NSView) ViewDidMoveToSuperview() {
	C.C_NSView_ViewDidMoveToSuperview(n.Ptr())
}

func (n NSView) ViewDidMoveToWindow() {
	C.C_NSView_ViewDidMoveToWindow(n.Ptr())
}

func (n NSView) ViewWillMoveToSuperview(newSuperview View) {
	C.C_NSView_ViewWillMoveToSuperview(n.Ptr(), objc.ExtractPtr(newSuperview))
}

func (n NSView) ViewWillMoveToWindow(newWindow Window) {
	C.C_NSView_ViewWillMoveToWindow(n.Ptr(), objc.ExtractPtr(newWindow))
}

func (n NSView) WillRemoveSubview(subview View) {
	C.C_NSView_WillRemoveSubview(n.Ptr(), objc.ExtractPtr(subview))
}

func (n NSView) SetFrameOrigin(newOrigin foundation.Point) {
	C.C_NSView_SetFrameOrigin(n.Ptr(), *(*C.CGPoint)(coregraphics.ToCGPointPointer(coregraphics.Point(newOrigin))))
}

func (n NSView) SetFrameSize(newSize foundation.Size) {
	C.C_NSView_SetFrameSize(n.Ptr(), *(*C.CGSize)(coregraphics.ToCGSizePointer(coregraphics.Size(newSize))))
}

func (n NSView) SetBoundsOrigin(newOrigin foundation.Point) {
	C.C_NSView_SetBoundsOrigin(n.Ptr(), *(*C.CGPoint)(coregraphics.ToCGPointPointer(coregraphics.Point(newOrigin))))
}

func (n NSView) SetBoundsSize(newSize foundation.Size) {
	C.C_NSView_SetBoundsSize(n.Ptr(), *(*C.CGSize)(coregraphics.ToCGSizePointer(coregraphics.Size(newSize))))
}

func (n NSView) UpdateLayer() {
	C.C_NSView_UpdateLayer(n.Ptr())
}

func (n NSView) DrawRect(dirtyRect foundation.Rect) {
	C.C_NSView_DrawRect(n.Ptr(), *(*C.CGRect)(coregraphics.ToCGRectPointer(coregraphics.Rect(dirtyRect))))
}

func (n NSView) NeedsToDrawRect(rect foundation.Rect) bool {
	result_ := C.C_NSView_NeedsToDrawRect(n.Ptr(), *(*C.CGRect)(coregraphics.ToCGRectPointer(coregraphics.Rect(rect))))
	return bool(result_)
}

func (n NSView) BitmapImageRepForCachingDisplayInRect(rect foundation.Rect) BitmapImageRep {
	result_ := C.C_NSView_BitmapImageRepForCachingDisplayInRect(n.Ptr(), *(*C.CGRect)(coregraphics.ToCGRectPointer(coregraphics.Rect(rect))))
	return MakeBitmapImageRep(result_)
}

func (n NSView) CacheDisplayInRect_ToBitmapImageRep(rect foundation.Rect, bitmapImageRep BitmapImageRep) {
	C.C_NSView_CacheDisplayInRect_ToBitmapImageRep(n.Ptr(), *(*C.CGRect)(coregraphics.ToCGRectPointer(coregraphics.Rect(rect))), objc.ExtractPtr(bitmapImageRep))
}

func (n NSView) Print(sender objc.Object) {
	C.C_NSView_Print(n.Ptr(), objc.ExtractPtr(sender))
}

func (n NSView) BeginPageInRect_AtPlacement(rect foundation.Rect, location foundation.Point) {
	C.C_NSView_BeginPageInRect_AtPlacement(n.Ptr(), *(*C.CGRect)(coregraphics.ToCGRectPointer(coregraphics.Rect(rect))), *(*C.CGPoint)(coregraphics.ToCGPointPointer(coregraphics.Point(location))))
}

func (n NSView) DataWithEPSInsideRect(rect foundation.Rect) []byte {
	result_ := C.C_NSView_DataWithEPSInsideRect(n.Ptr(), *(*C.CGRect)(coregraphics.ToCGRectPointer(coregraphics.Rect(rect))))
	return foundation.MakeData(result_).ToBytes()
}

func (n NSView) DataWithPDFInsideRect(rect foundation.Rect) []byte {
	result_ := C.C_NSView_DataWithPDFInsideRect(n.Ptr(), *(*C.CGRect)(coregraphics.ToCGRectPointer(coregraphics.Rect(rect))))
	return foundation.MakeData(result_).ToBytes()
}

func (n NSView) WriteEPSInsideRect_ToPasteboard(rect foundation.Rect, pasteboard Pasteboard) {
	C.C_NSView_WriteEPSInsideRect_ToPasteboard(n.Ptr(), *(*C.CGRect)(coregraphics.ToCGRectPointer(coregraphics.Rect(rect))), objc.ExtractPtr(pasteboard))
}

func (n NSView) WritePDFInsideRect_ToPasteboard(rect foundation.Rect, pasteboard Pasteboard) {
	C.C_NSView_WritePDFInsideRect_ToPasteboard(n.Ptr(), *(*C.CGRect)(coregraphics.ToCGRectPointer(coregraphics.Rect(rect))), objc.ExtractPtr(pasteboard))
}

func (n NSView) DrawPageBorderWithSize(borderSize foundation.Size) {
	C.C_NSView_DrawPageBorderWithSize(n.Ptr(), *(*C.CGSize)(coregraphics.ToCGSizePointer(coregraphics.Size(borderSize))))
}

func (n NSView) RectForPage(page int) foundation.Rect {
	result_ := C.C_NSView_RectForPage(n.Ptr(), C.int(page))
	return foundation.Rect(coregraphics.FromCGRectPointer(unsafe.Pointer(&result_)))
}

func (n NSView) LocationOfPrintRect(rect foundation.Rect) foundation.Point {
	result_ := C.C_NSView_LocationOfPrintRect(n.Ptr(), *(*C.CGRect)(coregraphics.ToCGRectPointer(coregraphics.Rect(rect))))
	return foundation.Point(coregraphics.FromCGPointPointer(unsafe.Pointer(&result_)))
}

func (n NSView) SetNeedsDisplayInRect(invalidRect foundation.Rect) {
	C.C_NSView_SetNeedsDisplayInRect(n.Ptr(), *(*C.CGRect)(coregraphics.ToCGRectPointer(coregraphics.Rect(invalidRect))))
}

func (n NSView) Display() {
	C.C_NSView_Display(n.Ptr())
}

func (n NSView) DisplayRect(rect foundation.Rect) {
	C.C_NSView_DisplayRect(n.Ptr(), *(*C.CGRect)(coregraphics.ToCGRectPointer(coregraphics.Rect(rect))))
}

func (n NSView) DisplayRectIgnoringOpacity(rect foundation.Rect) {
	C.C_NSView_DisplayRectIgnoringOpacity(n.Ptr(), *(*C.CGRect)(coregraphics.ToCGRectPointer(coregraphics.Rect(rect))))
}

func (n NSView) DisplayRectIgnoringOpacity_InContext(rect foundation.Rect, context GraphicsContext) {
	C.C_NSView_DisplayRectIgnoringOpacity_InContext(n.Ptr(), *(*C.CGRect)(coregraphics.ToCGRectPointer(coregraphics.Rect(rect))), objc.ExtractPtr(context))
}

func (n NSView) DisplayIfNeeded() {
	C.C_NSView_DisplayIfNeeded(n.Ptr())
}

func (n NSView) DisplayIfNeededInRect(rect foundation.Rect) {
	C.C_NSView_DisplayIfNeededInRect(n.Ptr(), *(*C.CGRect)(coregraphics.ToCGRectPointer(coregraphics.Rect(rect))))
}

func (n NSView) DisplayIfNeededIgnoringOpacity() {
	C.C_NSView_DisplayIfNeededIgnoringOpacity(n.Ptr())
}

func (n NSView) DisplayIfNeededInRectIgnoringOpacity(rect foundation.Rect) {
	C.C_NSView_DisplayIfNeededInRectIgnoringOpacity(n.Ptr(), *(*C.CGRect)(coregraphics.ToCGRectPointer(coregraphics.Rect(rect))))
}

func (n NSView) TranslateRectsNeedingDisplayInRect_By(clipRect foundation.Rect, delta foundation.Size) {
	C.C_NSView_TranslateRectsNeedingDisplayInRect_By(n.Ptr(), *(*C.CGRect)(coregraphics.ToCGRectPointer(coregraphics.Rect(clipRect))), *(*C.CGSize)(coregraphics.ToCGSizePointer(coregraphics.Size(delta))))
}

func (n NSView) ViewWillDraw() {
	C.C_NSView_ViewWillDraw(n.Ptr())
}

func (n NSView) ConvertPointFromBacking(point foundation.Point) foundation.Point {
	result_ := C.C_NSView_ConvertPointFromBacking(n.Ptr(), *(*C.CGPoint)(coregraphics.ToCGPointPointer(coregraphics.Point(point))))
	return foundation.Point(coregraphics.FromCGPointPointer(unsafe.Pointer(&result_)))
}

func (n NSView) ConvertPointToBacking(point foundation.Point) foundation.Point {
	result_ := C.C_NSView_ConvertPointToBacking(n.Ptr(), *(*C.CGPoint)(coregraphics.ToCGPointPointer(coregraphics.Point(point))))
	return foundation.Point(coregraphics.FromCGPointPointer(unsafe.Pointer(&result_)))
}

func (n NSView) ConvertPointFromLayer(point foundation.Point) foundation.Point {
	result_ := C.C_NSView_ConvertPointFromLayer(n.Ptr(), *(*C.CGPoint)(coregraphics.ToCGPointPointer(coregraphics.Point(point))))
	return foundation.Point(coregraphics.FromCGPointPointer(unsafe.Pointer(&result_)))
}

func (n NSView) ConvertPointToLayer(point foundation.Point) foundation.Point {
	result_ := C.C_NSView_ConvertPointToLayer(n.Ptr(), *(*C.CGPoint)(coregraphics.ToCGPointPointer(coregraphics.Point(point))))
	return foundation.Point(coregraphics.FromCGPointPointer(unsafe.Pointer(&result_)))
}

func (n NSView) ConvertRectFromBacking(rect foundation.Rect) foundation.Rect {
	result_ := C.C_NSView_ConvertRectFromBacking(n.Ptr(), *(*C.CGRect)(coregraphics.ToCGRectPointer(coregraphics.Rect(rect))))
	return foundation.Rect(coregraphics.FromCGRectPointer(unsafe.Pointer(&result_)))
}

func (n NSView) ConvertRectToBacking(rect foundation.Rect) foundation.Rect {
	result_ := C.C_NSView_ConvertRectToBacking(n.Ptr(), *(*C.CGRect)(coregraphics.ToCGRectPointer(coregraphics.Rect(rect))))
	return foundation.Rect(coregraphics.FromCGRectPointer(unsafe.Pointer(&result_)))
}

func (n NSView) ConvertRectFromLayer(rect foundation.Rect) foundation.Rect {
	result_ := C.C_NSView_ConvertRectFromLayer(n.Ptr(), *(*C.CGRect)(coregraphics.ToCGRectPointer(coregraphics.Rect(rect))))
	return foundation.Rect(coregraphics.FromCGRectPointer(unsafe.Pointer(&result_)))
}

func (n NSView) ConvertRectToLayer(rect foundation.Rect) foundation.Rect {
	result_ := C.C_NSView_ConvertRectToLayer(n.Ptr(), *(*C.CGRect)(coregraphics.ToCGRectPointer(coregraphics.Rect(rect))))
	return foundation.Rect(coregraphics.FromCGRectPointer(unsafe.Pointer(&result_)))
}

func (n NSView) ConvertSizeFromBacking(size foundation.Size) foundation.Size {
	result_ := C.C_NSView_ConvertSizeFromBacking(n.Ptr(), *(*C.CGSize)(coregraphics.ToCGSizePointer(coregraphics.Size(size))))
	return foundation.Size(coregraphics.FromCGSizePointer(unsafe.Pointer(&result_)))
}

func (n NSView) ConvertSizeToBacking(size foundation.Size) foundation.Size {
	result_ := C.C_NSView_ConvertSizeToBacking(n.Ptr(), *(*C.CGSize)(coregraphics.ToCGSizePointer(coregraphics.Size(size))))
	return foundation.Size(coregraphics.FromCGSizePointer(unsafe.Pointer(&result_)))
}

func (n NSView) ConvertSizeFromLayer(size foundation.Size) foundation.Size {
	result_ := C.C_NSView_ConvertSizeFromLayer(n.Ptr(), *(*C.CGSize)(coregraphics.ToCGSizePointer(coregraphics.Size(size))))
	return foundation.Size(coregraphics.FromCGSizePointer(unsafe.Pointer(&result_)))
}

func (n NSView) ConvertSizeToLayer(size foundation.Size) foundation.Size {
	result_ := C.C_NSView_ConvertSizeToLayer(n.Ptr(), *(*C.CGSize)(coregraphics.ToCGSizePointer(coregraphics.Size(size))))
	return foundation.Size(coregraphics.FromCGSizePointer(unsafe.Pointer(&result_)))
}

func (n NSView) ConvertPoint_FromView(point foundation.Point, view View) foundation.Point {
	result_ := C.C_NSView_ConvertPoint_FromView(n.Ptr(), *(*C.CGPoint)(coregraphics.ToCGPointPointer(coregraphics.Point(point))), objc.ExtractPtr(view))
	return foundation.Point(coregraphics.FromCGPointPointer(unsafe.Pointer(&result_)))
}

func (n NSView) ConvertPoint_ToView(point foundation.Point, view View) foundation.Point {
	result_ := C.C_NSView_ConvertPoint_ToView(n.Ptr(), *(*C.CGPoint)(coregraphics.ToCGPointPointer(coregraphics.Point(point))), objc.ExtractPtr(view))
	return foundation.Point(coregraphics.FromCGPointPointer(unsafe.Pointer(&result_)))
}

func (n NSView) ConvertSize_FromView(size foundation.Size, view View) foundation.Size {
	result_ := C.C_NSView_ConvertSize_FromView(n.Ptr(), *(*C.CGSize)(coregraphics.ToCGSizePointer(coregraphics.Size(size))), objc.ExtractPtr(view))
	return foundation.Size(coregraphics.FromCGSizePointer(unsafe.Pointer(&result_)))
}

func (n NSView) ConvertSize_ToView(size foundation.Size, view View) foundation.Size {
	result_ := C.C_NSView_ConvertSize_ToView(n.Ptr(), *(*C.CGSize)(coregraphics.ToCGSizePointer(coregraphics.Size(size))), objc.ExtractPtr(view))
	return foundation.Size(coregraphics.FromCGSizePointer(unsafe.Pointer(&result_)))
}

func (n NSView) ConvertRect_FromView(rect foundation.Rect, view View) foundation.Rect {
	result_ := C.C_NSView_ConvertRect_FromView(n.Ptr(), *(*C.CGRect)(coregraphics.ToCGRectPointer(coregraphics.Rect(rect))), objc.ExtractPtr(view))
	return foundation.Rect(coregraphics.FromCGRectPointer(unsafe.Pointer(&result_)))
}

func (n NSView) ConvertRect_ToView(rect foundation.Rect, view View) foundation.Rect {
	result_ := C.C_NSView_ConvertRect_ToView(n.Ptr(), *(*C.CGRect)(coregraphics.ToCGRectPointer(coregraphics.Rect(rect))), objc.ExtractPtr(view))
	return foundation.Rect(coregraphics.FromCGRectPointer(unsafe.Pointer(&result_)))
}

func (n NSView) CenterScanRect(rect foundation.Rect) foundation.Rect {
	result_ := C.C_NSView_CenterScanRect(n.Ptr(), *(*C.CGRect)(coregraphics.ToCGRectPointer(coregraphics.Rect(rect))))
	return foundation.Rect(coregraphics.FromCGRectPointer(unsafe.Pointer(&result_)))
}

func (n NSView) TranslateOriginToPoint(translation foundation.Point) {
	C.C_NSView_TranslateOriginToPoint(n.Ptr(), *(*C.CGPoint)(coregraphics.ToCGPointPointer(coregraphics.Point(translation))))
}

func (n NSView) ScaleUnitSquareToSize(newUnitSize foundation.Size) {
	C.C_NSView_ScaleUnitSquareToSize(n.Ptr(), *(*C.CGSize)(coregraphics.ToCGSizePointer(coregraphics.Size(newUnitSize))))
}

func (n NSView) RotateByAngle(angle coregraphics.Float) {
	C.C_NSView_RotateByAngle(n.Ptr(), C.double(float64(angle)))
}

func (n NSView) ResizeSubviewsWithOldSize(oldSize foundation.Size) {
	C.C_NSView_ResizeSubviewsWithOldSize(n.Ptr(), *(*C.CGSize)(coregraphics.ToCGSizePointer(coregraphics.Size(oldSize))))
}

func (n NSView) ResizeWithOldSuperviewSize(oldSize foundation.Size) {
	C.C_NSView_ResizeWithOldSuperviewSize(n.Ptr(), *(*C.CGSize)(coregraphics.ToCGSizePointer(coregraphics.Size(oldSize))))
}

func (n NSView) AddConstraint(constraint LayoutConstraint) {
	C.C_NSView_AddConstraint(n.Ptr(), objc.ExtractPtr(constraint))
}

func (n NSView) AddConstraints(constraints []LayoutConstraint) {
	var cConstraints C.Array
	if len(constraints) > 0 {
		cConstraintsData := make([]unsafe.Pointer, len(constraints))
		for idx, v := range constraints {
			cConstraintsData[idx] = objc.ExtractPtr(v)
		}
		cConstraints.data = unsafe.Pointer(&cConstraintsData[0])
		cConstraints.len = C.int(len(constraints))
	}
	C.C_NSView_AddConstraints(n.Ptr(), cConstraints)
}

func (n NSView) RemoveConstraint(constraint LayoutConstraint) {
	C.C_NSView_RemoveConstraint(n.Ptr(), objc.ExtractPtr(constraint))
}

func (n NSView) RemoveConstraints(constraints []LayoutConstraint) {
	var cConstraints C.Array
	if len(constraints) > 0 {
		cConstraintsData := make([]unsafe.Pointer, len(constraints))
		for idx, v := range constraints {
			cConstraintsData[idx] = objc.ExtractPtr(v)
		}
		cConstraints.data = unsafe.Pointer(&cConstraintsData[0])
		cConstraints.len = C.int(len(constraints))
	}
	C.C_NSView_RemoveConstraints(n.Ptr(), cConstraints)
}

func (n NSView) AddLayoutGuide(guide LayoutGuide) {
	C.C_NSView_AddLayoutGuide(n.Ptr(), objc.ExtractPtr(guide))
}

func (n NSView) RemoveLayoutGuide(guide LayoutGuide) {
	C.C_NSView_RemoveLayoutGuide(n.Ptr(), objc.ExtractPtr(guide))
}

func (n NSView) InvalidateIntrinsicContentSize() {
	C.C_NSView_InvalidateIntrinsicContentSize(n.Ptr())
}

func (n NSView) ContentCompressionResistancePriorityForOrientation(orientation LayoutConstraintOrientation) LayoutPriority {
	result_ := C.C_NSView_ContentCompressionResistancePriorityForOrientation(n.Ptr(), C.int(int(orientation)))
	return LayoutPriority(float32(result_))
}

func (n NSView) SetContentCompressionResistancePriority_ForOrientation(priority LayoutPriority, orientation LayoutConstraintOrientation) {
	C.C_NSView_SetContentCompressionResistancePriority_ForOrientation(n.Ptr(), C.float(float32(priority)), C.int(int(orientation)))
}

func (n NSView) ContentHuggingPriorityForOrientation(orientation LayoutConstraintOrientation) LayoutPriority {
	result_ := C.C_NSView_ContentHuggingPriorityForOrientation(n.Ptr(), C.int(int(orientation)))
	return LayoutPriority(float32(result_))
}

func (n NSView) SetContentHuggingPriority_ForOrientation(priority LayoutPriority, orientation LayoutConstraintOrientation) {
	C.C_NSView_SetContentHuggingPriority_ForOrientation(n.Ptr(), C.float(float32(priority)), C.int(int(orientation)))
}

func (n NSView) AlignmentRectForFrame(frame foundation.Rect) foundation.Rect {
	result_ := C.C_NSView_AlignmentRectForFrame(n.Ptr(), *(*C.CGRect)(coregraphics.ToCGRectPointer(coregraphics.Rect(frame))))
	return foundation.Rect(coregraphics.FromCGRectPointer(unsafe.Pointer(&result_)))
}

func (n NSView) FrameForAlignmentRect(alignmentRect foundation.Rect) foundation.Rect {
	result_ := C.C_NSView_FrameForAlignmentRect(n.Ptr(), *(*C.CGRect)(coregraphics.ToCGRectPointer(coregraphics.Rect(alignmentRect))))
	return foundation.Rect(coregraphics.FromCGRectPointer(unsafe.Pointer(&result_)))
}

func (n NSView) Layout() {
	C.C_NSView_Layout(n.Ptr())
}

func (n NSView) LayoutSubtreeIfNeeded() {
	C.C_NSView_LayoutSubtreeIfNeeded(n.Ptr())
}

func (n NSView) UpdateConstraints() {
	C.C_NSView_UpdateConstraints(n.Ptr())
}

func (n NSView) UpdateConstraintsForSubtreeIfNeeded() {
	C.C_NSView_UpdateConstraintsForSubtreeIfNeeded(n.Ptr())
}

func (n NSView) ConstraintsAffectingLayoutForOrientation(orientation LayoutConstraintOrientation) []LayoutConstraint {
	result_ := C.C_NSView_ConstraintsAffectingLayoutForOrientation(n.Ptr(), C.int(int(orientation)))
	if result_.len > 0 {
		defer C.free(result_.data)
	}
	result_Slice := unsafe.Slice((*unsafe.Pointer)(result_.data), int(result_.len))
	var goResult_ = make([]LayoutConstraint, len(result_Slice))
	for idx, r := range result_Slice {
		goResult_[idx] = MakeLayoutConstraint(r)
	}
	return goResult_
}

func (n NSView) ExerciseAmbiguityInLayout() {
	C.C_NSView_ExerciseAmbiguityInLayout(n.Ptr())
}

func (n NSView) DrawFocusRingMask() {
	C.C_NSView_DrawFocusRingMask(n.Ptr())
}

func (n NSView) NoteFocusRingMaskChanged() {
	C.C_NSView_NoteFocusRingMaskChanged(n.Ptr())
}

func (n NSView) SetKeyboardFocusRingNeedsDisplayInRect(rect foundation.Rect) {
	C.C_NSView_SetKeyboardFocusRingNeedsDisplayInRect(n.Ptr(), *(*C.CGRect)(coregraphics.ToCGRectPointer(coregraphics.Rect(rect))))
}

func (n NSView) EnterFullScreenMode_WithOptions(screen Screen, options map[ViewFullScreenModeOptionKey]objc.Object) bool {
	var cOptions C.Dictionary
	if len(options) > 0 {
		cOptionsKeyData := make([]unsafe.Pointer, len(options))
		cOptionsValueData := make([]unsafe.Pointer, len(options))
		var idx = 0
		for k, v := range options {
			cOptionsKeyData[idx] = foundation.NewString(string(k)).Ptr()
			cOptionsValueData[idx] = objc.ExtractPtr(v)
			idx++
		}
		cOptions.key_data = unsafe.Pointer(&cOptionsKeyData[0])
		cOptions.value_data = unsafe.Pointer(&cOptionsValueData[0])
		cOptions.len = C.int(len(options))
	}
	result_ := C.C_NSView_EnterFullScreenMode_WithOptions(n.Ptr(), objc.ExtractPtr(screen), cOptions)
	return bool(result_)
}

func (n NSView) ExitFullScreenModeWithOptions(options map[ViewFullScreenModeOptionKey]objc.Object) {
	var cOptions C.Dictionary
	if len(options) > 0 {
		cOptionsKeyData := make([]unsafe.Pointer, len(options))
		cOptionsValueData := make([]unsafe.Pointer, len(options))
		var idx = 0
		for k, v := range options {
			cOptionsKeyData[idx] = foundation.NewString(string(k)).Ptr()
			cOptionsValueData[idx] = objc.ExtractPtr(v)
			idx++
		}
		cOptions.key_data = unsafe.Pointer(&cOptionsKeyData[0])
		cOptions.value_data = unsafe.Pointer(&cOptionsValueData[0])
		cOptions.len = C.int(len(options))
	}
	C.C_NSView_ExitFullScreenModeWithOptions(n.Ptr(), cOptions)
}

func (n NSView) ViewDidHide() {
	C.C_NSView_ViewDidHide(n.Ptr())
}

func (n NSView) ViewDidUnhide() {
	C.C_NSView_ViewDidUnhide(n.Ptr())
}

func (n NSView) ViewWillStartLiveResize() {
	C.C_NSView_ViewWillStartLiveResize(n.Ptr())
}

func (n NSView) ViewDidEndLiveResize() {
	C.C_NSView_ViewDidEndLiveResize(n.Ptr())
}

func (n NSView) AddGestureRecognizer(gestureRecognizer GestureRecognizer) {
	C.C_NSView_AddGestureRecognizer(n.Ptr(), objc.ExtractPtr(gestureRecognizer))
}

func (n NSView) RemoveGestureRecognizer(gestureRecognizer GestureRecognizer) {
	C.C_NSView_RemoveGestureRecognizer(n.Ptr(), objc.ExtractPtr(gestureRecognizer))
}

func (n NSView) AcceptsFirstMouse(event Event) bool {
	result_ := C.C_NSView_AcceptsFirstMouse(n.Ptr(), objc.ExtractPtr(event))
	return bool(result_)
}

func (n NSView) HitTest(point foundation.Point) View {
	result_ := C.C_NSView_HitTest(n.Ptr(), *(*C.CGPoint)(coregraphics.ToCGPointPointer(coregraphics.Point(point))))
	return MakeView(result_)
}

func (n NSView) Mouse_InRect(point foundation.Point, rect foundation.Rect) bool {
	result_ := C.C_NSView_Mouse_InRect(n.Ptr(), *(*C.CGPoint)(coregraphics.ToCGPointPointer(coregraphics.Point(point))), *(*C.CGRect)(coregraphics.ToCGRectPointer(coregraphics.Rect(rect))))
	return bool(result_)
}

func (n NSView) PrepareContentInRect(rect foundation.Rect) {
	C.C_NSView_PrepareContentInRect(n.Ptr(), *(*C.CGRect)(coregraphics.ToCGRectPointer(coregraphics.Rect(rect))))
}

func (n NSView) ScrollPoint(point foundation.Point) {
	C.C_NSView_ScrollPoint(n.Ptr(), *(*C.CGPoint)(coregraphics.ToCGPointPointer(coregraphics.Point(point))))
}

func (n NSView) ScrollRectToVisible(rect foundation.Rect) bool {
	result_ := C.C_NSView_ScrollRectToVisible(n.Ptr(), *(*C.CGRect)(coregraphics.ToCGRectPointer(coregraphics.Rect(rect))))
	return bool(result_)
}

func (n NSView) Autoscroll(event Event) bool {
	result_ := C.C_NSView_Autoscroll(n.Ptr(), objc.ExtractPtr(event))
	return bool(result_)
}

func (n NSView) AdjustScroll(newVisible foundation.Rect) foundation.Rect {
	result_ := C.C_NSView_AdjustScroll(n.Ptr(), *(*C.CGRect)(coregraphics.ToCGRectPointer(coregraphics.Rect(newVisible))))
	return foundation.Rect(coregraphics.FromCGRectPointer(unsafe.Pointer(&result_)))
}

func (n NSView) ScrollClipView_ToPoint(clipView ClipView, point foundation.Point) {
	C.C_NSView_ScrollClipView_ToPoint(n.Ptr(), objc.ExtractPtr(clipView), *(*C.CGPoint)(coregraphics.ToCGPointPointer(coregraphics.Point(point))))
}

func (n NSView) ReflectScrolledClipView(clipView ClipView) {
	C.C_NSView_ReflectScrolledClipView(n.Ptr(), objc.ExtractPtr(clipView))
}

func (n NSView) RegisterForDraggedTypes(newTypes []PasteboardType) {
	var cNewTypes C.Array
	if len(newTypes) > 0 {
		cNewTypesData := make([]unsafe.Pointer, len(newTypes))
		for idx, v := range newTypes {
			cNewTypesData[idx] = foundation.NewString(string(v)).Ptr()
		}
		cNewTypes.data = unsafe.Pointer(&cNewTypesData[0])
		cNewTypes.len = C.int(len(newTypes))
	}
	C.C_NSView_RegisterForDraggedTypes(n.Ptr(), cNewTypes)
}

func (n NSView) UnregisterDraggedTypes() {
	C.C_NSView_UnregisterDraggedTypes(n.Ptr())
}

func (n NSView) BeginDraggingSessionWithItems_Event_Source(items []DraggingItem, event Event, source objc.Object) DraggingSession {
	var cItems C.Array
	if len(items) > 0 {
		cItemsData := make([]unsafe.Pointer, len(items))
		for idx, v := range items {
			cItemsData[idx] = objc.ExtractPtr(v)
		}
		cItems.data = unsafe.Pointer(&cItemsData[0])
		cItems.len = C.int(len(items))
	}
	result_ := C.C_NSView_BeginDraggingSessionWithItems_Event_Source(n.Ptr(), cItems, objc.ExtractPtr(event), objc.ExtractPtr(source))
	return MakeDraggingSession(result_)
}

func (n NSView) ShouldDelayWindowOrderingForEvent(event Event) bool {
	result_ := C.C_NSView_ShouldDelayWindowOrderingForEvent(n.Ptr(), objc.ExtractPtr(event))
	return bool(result_)
}

func (n NSView) RectForSmartMagnificationAtPoint_InRect(location foundation.Point, visibleRect foundation.Rect) foundation.Rect {
	result_ := C.C_NSView_RectForSmartMagnificationAtPoint_InRect(n.Ptr(), *(*C.CGPoint)(coregraphics.ToCGPointPointer(coregraphics.Point(location))), *(*C.CGRect)(coregraphics.ToCGRectPointer(coregraphics.Rect(visibleRect))))
	return foundation.Rect(coregraphics.FromCGRectPointer(unsafe.Pointer(&result_)))
}

func (n NSView) ViewDidChangeBackingProperties() {
	C.C_NSView_ViewDidChangeBackingProperties(n.Ptr())
}

func (n NSView) ViewWithTag(tag int) View {
	result_ := C.C_NSView_ViewWithTag(n.Ptr(), C.int(tag))
	return MakeView(result_)
}

func (n NSView) RemoveAllToolTips() {
	C.C_NSView_RemoveAllToolTips(n.Ptr())
}

func (n NSView) RemoveToolTip(tag ToolTipTag) {
	C.C_NSView_RemoveToolTip(n.Ptr(), C.int(int(tag)))
}

func (n NSView) RemoveTrackingRect(tag TrackingRectTag) {
	C.C_NSView_RemoveTrackingRect(n.Ptr(), C.int(int(tag)))
}

func (n NSView) AddTrackingArea(trackingArea TrackingArea) {
	C.C_NSView_AddTrackingArea(n.Ptr(), objc.ExtractPtr(trackingArea))
}

func (n NSView) RemoveTrackingArea(trackingArea TrackingArea) {
	C.C_NSView_RemoveTrackingArea(n.Ptr(), objc.ExtractPtr(trackingArea))
}

func (n NSView) UpdateTrackingAreas() {
	C.C_NSView_UpdateTrackingAreas(n.Ptr())
}

func (n NSView) AddCursorRect_Cursor(rect foundation.Rect, object Cursor) {
	C.C_NSView_AddCursorRect_Cursor(n.Ptr(), *(*C.CGRect)(coregraphics.ToCGRectPointer(coregraphics.Rect(rect))), objc.ExtractPtr(object))
}

func (n NSView) RemoveCursorRect_Cursor(rect foundation.Rect, object Cursor) {
	C.C_NSView_RemoveCursorRect_Cursor(n.Ptr(), *(*C.CGRect)(coregraphics.ToCGRectPointer(coregraphics.Rect(rect))), objc.ExtractPtr(object))
}

func (n NSView) DiscardCursorRects() {
	C.C_NSView_DiscardCursorRects(n.Ptr())
}

func (n NSView) ResetCursorRects() {
	C.C_NSView_ResetCursorRects(n.Ptr())
}

func (n NSView) MenuForEvent(event Event) Menu {
	result_ := C.C_NSView_MenuForEvent(n.Ptr(), objc.ExtractPtr(event))
	return MakeMenu(result_)
}

func (n NSView) WillOpenMenu_WithEvent(menu Menu, event Event) {
	C.C_NSView_WillOpenMenu_WithEvent(n.Ptr(), objc.ExtractPtr(menu), objc.ExtractPtr(event))
}

func (n NSView) DidCloseMenu_WithEvent(menu Menu, event Event) {
	C.C_NSView_DidCloseMenu_WithEvent(n.Ptr(), objc.ExtractPtr(menu), objc.ExtractPtr(event))
}

func (n NSView) BeginDocument() {
	C.C_NSView_BeginDocument(n.Ptr())
}

func (n NSView) EndDocument() {
	C.C_NSView_EndDocument(n.Ptr())
}

func (n NSView) EndPage() {
	C.C_NSView_EndPage(n.Ptr())
}

func (n NSView) ShowDefinitionForAttributedString_AtPoint(attrString foundation.AttributedString, textBaselineOrigin foundation.Point) {
	C.C_NSView_ShowDefinitionForAttributedString_AtPoint(n.Ptr(), objc.ExtractPtr(attrString), *(*C.CGPoint)(coregraphics.ToCGPointPointer(coregraphics.Point(textBaselineOrigin))))
}

func (n NSView) RulerView_DidAddMarker(ruler RulerView, marker RulerMarker) {
	C.C_NSView_RulerView_DidAddMarker(n.Ptr(), objc.ExtractPtr(ruler), objc.ExtractPtr(marker))
}

func (n NSView) RulerView_DidMoveMarker(ruler RulerView, marker RulerMarker) {
	C.C_NSView_RulerView_DidMoveMarker(n.Ptr(), objc.ExtractPtr(ruler), objc.ExtractPtr(marker))
}

func (n NSView) RulerView_DidRemoveMarker(ruler RulerView, marker RulerMarker) {
	C.C_NSView_RulerView_DidRemoveMarker(n.Ptr(), objc.ExtractPtr(ruler), objc.ExtractPtr(marker))
}

func (n NSView) RulerView_HandleMouseDown(ruler RulerView, event Event) {
	C.C_NSView_RulerView_HandleMouseDown(n.Ptr(), objc.ExtractPtr(ruler), objc.ExtractPtr(event))
}

func (n NSView) RulerView_LocationForPoint(ruler RulerView, point foundation.Point) coregraphics.Float {
	result_ := C.C_NSView_RulerView_LocationForPoint(n.Ptr(), objc.ExtractPtr(ruler), *(*C.CGPoint)(coregraphics.ToCGPointPointer(coregraphics.Point(point))))
	return coregraphics.Float(float64(result_))
}

func (n NSView) RulerView_PointForLocation(ruler RulerView, point coregraphics.Float) foundation.Point {
	result_ := C.C_NSView_RulerView_PointForLocation(n.Ptr(), objc.ExtractPtr(ruler), C.double(float64(point)))
	return foundation.Point(coregraphics.FromCGPointPointer(unsafe.Pointer(&result_)))
}

func (n NSView) RulerView_ShouldAddMarker(ruler RulerView, marker RulerMarker) bool {
	result_ := C.C_NSView_RulerView_ShouldAddMarker(n.Ptr(), objc.ExtractPtr(ruler), objc.ExtractPtr(marker))
	return bool(result_)
}

func (n NSView) RulerView_ShouldMoveMarker(ruler RulerView, marker RulerMarker) bool {
	result_ := C.C_NSView_RulerView_ShouldMoveMarker(n.Ptr(), objc.ExtractPtr(ruler), objc.ExtractPtr(marker))
	return bool(result_)
}

func (n NSView) RulerView_ShouldRemoveMarker(ruler RulerView, marker RulerMarker) bool {
	result_ := C.C_NSView_RulerView_ShouldRemoveMarker(n.Ptr(), objc.ExtractPtr(ruler), objc.ExtractPtr(marker))
	return bool(result_)
}

func (n NSView) RulerView_WillAddMarker_AtLocation(ruler RulerView, marker RulerMarker, location coregraphics.Float) coregraphics.Float {
	result_ := C.C_NSView_RulerView_WillAddMarker_AtLocation(n.Ptr(), objc.ExtractPtr(ruler), objc.ExtractPtr(marker), C.double(float64(location)))
	return coregraphics.Float(float64(result_))
}

func (n NSView) RulerView_WillMoveMarker_ToLocation(ruler RulerView, marker RulerMarker, location coregraphics.Float) coregraphics.Float {
	result_ := C.C_NSView_RulerView_WillMoveMarker_ToLocation(n.Ptr(), objc.ExtractPtr(ruler), objc.ExtractPtr(marker), C.double(float64(location)))
	return coregraphics.Float(float64(result_))
}

func (n NSView) RulerView_WillSetClientView(ruler RulerView, newClient View) {
	C.C_NSView_RulerView_WillSetClientView(n.Ptr(), objc.ExtractPtr(ruler), objc.ExtractPtr(newClient))
}

func (n NSView) ViewDidChangeEffectiveAppearance() {
	C.C_NSView_ViewDidChangeEffectiveAppearance(n.Ptr())
}

func (n NSView) Superview() View {
	result_ := C.C_NSView_Superview(n.Ptr())
	return MakeView(result_)
}

func (n NSView) Subviews() []View {
	result_ := C.C_NSView_Subviews(n.Ptr())
	if result_.len > 0 {
		defer C.free(result_.data)
	}
	result_Slice := unsafe.Slice((*unsafe.Pointer)(result_.data), int(result_.len))
	var goResult_ = make([]View, len(result_Slice))
	for idx, r := range result_Slice {
		goResult_[idx] = MakeView(r)
	}
	return goResult_
}

func (n NSView) SetSubviews(value []View) {
	var cValue C.Array
	if len(value) > 0 {
		cValueData := make([]unsafe.Pointer, len(value))
		for idx, v := range value {
			cValueData[idx] = objc.ExtractPtr(v)
		}
		cValue.data = unsafe.Pointer(&cValueData[0])
		cValue.len = C.int(len(value))
	}
	C.C_NSView_SetSubviews(n.Ptr(), cValue)
}

func (n NSView) Window() Window {
	result_ := C.C_NSView_Window(n.Ptr())
	return MakeWindow(result_)
}

func (n NSView) OpaqueAncestor() View {
	result_ := C.C_NSView_OpaqueAncestor(n.Ptr())
	return MakeView(result_)
}

func (n NSView) EnclosingMenuItem() MenuItem {
	result_ := C.C_NSView_EnclosingMenuItem(n.Ptr())
	return MakeMenuItem(result_)
}

func (n NSView) Frame() foundation.Rect {
	result_ := C.C_NSView_Frame(n.Ptr())
	return foundation.Rect(coregraphics.FromCGRectPointer(unsafe.Pointer(&result_)))
}

func (n NSView) SetFrame(value foundation.Rect) {
	C.C_NSView_SetFrame(n.Ptr(), *(*C.CGRect)(coregraphics.ToCGRectPointer(coregraphics.Rect(value))))
}

func (n NSView) FrameRotation() coregraphics.Float {
	result_ := C.C_NSView_FrameRotation(n.Ptr())
	return coregraphics.Float(float64(result_))
}

func (n NSView) SetFrameRotation(value coregraphics.Float) {
	C.C_NSView_SetFrameRotation(n.Ptr(), C.double(float64(value)))
}

func (n NSView) Bounds() foundation.Rect {
	result_ := C.C_NSView_Bounds(n.Ptr())
	return foundation.Rect(coregraphics.FromCGRectPointer(unsafe.Pointer(&result_)))
}

func (n NSView) SetBounds(value foundation.Rect) {
	C.C_NSView_SetBounds(n.Ptr(), *(*C.CGRect)(coregraphics.ToCGRectPointer(coregraphics.Rect(value))))
}

func (n NSView) BoundsRotation() coregraphics.Float {
	result_ := C.C_NSView_BoundsRotation(n.Ptr())
	return coregraphics.Float(float64(result_))
}

func (n NSView) SetBoundsRotation(value coregraphics.Float) {
	C.C_NSView_SetBoundsRotation(n.Ptr(), C.double(float64(value)))
}

func (n NSView) WantsLayer() bool {
	result_ := C.C_NSView_WantsLayer(n.Ptr())
	return bool(result_)
}

func (n NSView) SetWantsLayer(value bool) {
	C.C_NSView_SetWantsLayer(n.Ptr(), C.bool(value))
}

func (n NSView) WantsUpdateLayer() bool {
	result_ := C.C_NSView_WantsUpdateLayer(n.Ptr())
	return bool(result_)
}

func (n NSView) LayerContentsPlacement() ViewLayerContentsPlacement {
	result_ := C.C_NSView_LayerContentsPlacement(n.Ptr())
	return ViewLayerContentsPlacement(int(result_))
}

func (n NSView) SetLayerContentsPlacement(value ViewLayerContentsPlacement) {
	C.C_NSView_SetLayerContentsPlacement(n.Ptr(), C.int(int(value)))
}

func (n NSView) LayerContentsRedrawPolicy() ViewLayerContentsRedrawPolicy {
	result_ := C.C_NSView_LayerContentsRedrawPolicy(n.Ptr())
	return ViewLayerContentsRedrawPolicy(int(result_))
}

func (n NSView) SetLayerContentsRedrawPolicy(value ViewLayerContentsRedrawPolicy) {
	C.C_NSView_SetLayerContentsRedrawPolicy(n.Ptr(), C.int(int(value)))
}

func (n NSView) CanDrawSubviewsIntoLayer() bool {
	result_ := C.C_NSView_CanDrawSubviewsIntoLayer(n.Ptr())
	return bool(result_)
}

func (n NSView) SetCanDrawSubviewsIntoLayer(value bool) {
	C.C_NSView_SetCanDrawSubviewsIntoLayer(n.Ptr(), C.bool(value))
}

func (n NSView) LayerUsesCoreImageFilters() bool {
	result_ := C.C_NSView_LayerUsesCoreImageFilters(n.Ptr())
	return bool(result_)
}

func (n NSView) SetLayerUsesCoreImageFilters(value bool) {
	C.C_NSView_SetLayerUsesCoreImageFilters(n.Ptr(), C.bool(value))
}

func (n NSView) AlphaValue() coregraphics.Float {
	result_ := C.C_NSView_AlphaValue(n.Ptr())
	return coregraphics.Float(float64(result_))
}

func (n NSView) SetAlphaValue(value coregraphics.Float) {
	C.C_NSView_SetAlphaValue(n.Ptr(), C.double(float64(value)))
}

func (n NSView) FrameCenterRotation() coregraphics.Float {
	result_ := C.C_NSView_FrameCenterRotation(n.Ptr())
	return coregraphics.Float(float64(result_))
}

func (n NSView) SetFrameCenterRotation(value coregraphics.Float) {
	C.C_NSView_SetFrameCenterRotation(n.Ptr(), C.double(float64(value)))
}

func (n NSView) Shadow() Shadow {
	result_ := C.C_NSView_Shadow(n.Ptr())
	return MakeShadow(result_)
}

func (n NSView) SetShadow(value Shadow) {
	C.C_NSView_SetShadow(n.Ptr(), objc.ExtractPtr(value))
}

func (n NSView) CanDrawConcurrently() bool {
	result_ := C.C_NSView_CanDrawConcurrently(n.Ptr())
	return bool(result_)
}

func (n NSView) SetCanDrawConcurrently(value bool) {
	C.C_NSView_SetCanDrawConcurrently(n.Ptr(), C.bool(value))
}

func (n NSView) VisibleRect() foundation.Rect {
	result_ := C.C_NSView_VisibleRect(n.Ptr())
	return foundation.Rect(coregraphics.FromCGRectPointer(unsafe.Pointer(&result_)))
}

func (n NSView) WantsDefaultClipping() bool {
	result_ := C.C_NSView_WantsDefaultClipping(n.Ptr())
	return bool(result_)
}

func (n NSView) PrintJobTitle() string {
	result_ := C.C_NSView_PrintJobTitle(n.Ptr())
	return foundation.MakeString(result_).String()
}

func (n NSView) PageHeader() foundation.AttributedString {
	result_ := C.C_NSView_PageHeader(n.Ptr())
	return foundation.MakeAttributedString(result_)
}

func (n NSView) PageFooter() foundation.AttributedString {
	result_ := C.C_NSView_PageFooter(n.Ptr())
	return foundation.MakeAttributedString(result_)
}

func (n NSView) HeightAdjustLimit() coregraphics.Float {
	result_ := C.C_NSView_HeightAdjustLimit(n.Ptr())
	return coregraphics.Float(float64(result_))
}

func (n NSView) WidthAdjustLimit() coregraphics.Float {
	result_ := C.C_NSView_WidthAdjustLimit(n.Ptr())
	return coregraphics.Float(float64(result_))
}

func (n NSView) NeedsDisplay() bool {
	result_ := C.C_NSView_NeedsDisplay(n.Ptr())
	return bool(result_)
}

func (n NSView) SetNeedsDisplay(value bool) {
	C.C_NSView_SetNeedsDisplay(n.Ptr(), C.bool(value))
}

func (n NSView) IsOpaque() bool {
	result_ := C.C_NSView_IsOpaque(n.Ptr())
	return bool(result_)
}

func (n NSView) IsFlipped() bool {
	result_ := C.C_NSView_IsFlipped(n.Ptr())
	return bool(result_)
}

func (n NSView) IsRotatedFromBase() bool {
	result_ := C.C_NSView_IsRotatedFromBase(n.Ptr())
	return bool(result_)
}

func (n NSView) IsRotatedOrScaledFromBase() bool {
	result_ := C.C_NSView_IsRotatedOrScaledFromBase(n.Ptr())
	return bool(result_)
}

func (n NSView) AutoresizesSubviews() bool {
	result_ := C.C_NSView_AutoresizesSubviews(n.Ptr())
	return bool(result_)
}

func (n NSView) SetAutoresizesSubviews(value bool) {
	C.C_NSView_SetAutoresizesSubviews(n.Ptr(), C.bool(value))
}

func (n NSView) AutoresizingMask() AutoresizingMaskOptions {
	result_ := C.C_NSView_AutoresizingMask(n.Ptr())
	return AutoresizingMaskOptions(uint(result_))
}

func (n NSView) SetAutoresizingMask(value AutoresizingMaskOptions) {
	C.C_NSView_SetAutoresizingMask(n.Ptr(), C.uint(uint(value)))
}

func (n NSView) BottomAnchor() LayoutYAxisAnchor {
	result_ := C.C_NSView_BottomAnchor(n.Ptr())
	return MakeLayoutYAxisAnchor(result_)
}

func (n NSView) CenterXAnchor() LayoutXAxisAnchor {
	result_ := C.C_NSView_CenterXAnchor(n.Ptr())
	return MakeLayoutXAxisAnchor(result_)
}

func (n NSView) CenterYAnchor() LayoutYAxisAnchor {
	result_ := C.C_NSView_CenterYAnchor(n.Ptr())
	return MakeLayoutYAxisAnchor(result_)
}

func (n NSView) FirstBaselineAnchor() LayoutYAxisAnchor {
	result_ := C.C_NSView_FirstBaselineAnchor(n.Ptr())
	return MakeLayoutYAxisAnchor(result_)
}

func (n NSView) HeightAnchor() LayoutDimension {
	result_ := C.C_NSView_HeightAnchor(n.Ptr())
	return MakeLayoutDimension(result_)
}

func (n NSView) LastBaselineAnchor() LayoutYAxisAnchor {
	result_ := C.C_NSView_LastBaselineAnchor(n.Ptr())
	return MakeLayoutYAxisAnchor(result_)
}

func (n NSView) LeadingAnchor() LayoutXAxisAnchor {
	result_ := C.C_NSView_LeadingAnchor(n.Ptr())
	return MakeLayoutXAxisAnchor(result_)
}

func (n NSView) LeftAnchor() LayoutXAxisAnchor {
	result_ := C.C_NSView_LeftAnchor(n.Ptr())
	return MakeLayoutXAxisAnchor(result_)
}

func (n NSView) RightAnchor() LayoutXAxisAnchor {
	result_ := C.C_NSView_RightAnchor(n.Ptr())
	return MakeLayoutXAxisAnchor(result_)
}

func (n NSView) TopAnchor() LayoutYAxisAnchor {
	result_ := C.C_NSView_TopAnchor(n.Ptr())
	return MakeLayoutYAxisAnchor(result_)
}

func (n NSView) TrailingAnchor() LayoutXAxisAnchor {
	result_ := C.C_NSView_TrailingAnchor(n.Ptr())
	return MakeLayoutXAxisAnchor(result_)
}

func (n NSView) WidthAnchor() LayoutDimension {
	result_ := C.C_NSView_WidthAnchor(n.Ptr())
	return MakeLayoutDimension(result_)
}

func (n NSView) Constraints() []LayoutConstraint {
	result_ := C.C_NSView_Constraints(n.Ptr())
	if result_.len > 0 {
		defer C.free(result_.data)
	}
	result_Slice := unsafe.Slice((*unsafe.Pointer)(result_.data), int(result_.len))
	var goResult_ = make([]LayoutConstraint, len(result_Slice))
	for idx, r := range result_Slice {
		goResult_[idx] = MakeLayoutConstraint(r)
	}
	return goResult_
}

func (n NSView) LayoutGuides() []LayoutGuide {
	result_ := C.C_NSView_LayoutGuides(n.Ptr())
	if result_.len > 0 {
		defer C.free(result_.data)
	}
	result_Slice := unsafe.Slice((*unsafe.Pointer)(result_.data), int(result_.len))
	var goResult_ = make([]LayoutGuide, len(result_Slice))
	for idx, r := range result_Slice {
		goResult_[idx] = MakeLayoutGuide(r)
	}
	return goResult_
}

func (n NSView) LayoutMarginsGuide() LayoutGuide {
	result_ := C.C_NSView_LayoutMarginsGuide(n.Ptr())
	return MakeLayoutGuide(result_)
}

func (n NSView) SafeAreaLayoutGuide() LayoutGuide {
	result_ := C.C_NSView_SafeAreaLayoutGuide(n.Ptr())
	return MakeLayoutGuide(result_)
}

func (n NSView) FittingSize() foundation.Size {
	result_ := C.C_NSView_FittingSize(n.Ptr())
	return foundation.Size(coregraphics.FromCGSizePointer(unsafe.Pointer(&result_)))
}

func (n NSView) IntrinsicContentSize() foundation.Size {
	result_ := C.C_NSView_IntrinsicContentSize(n.Ptr())
	return foundation.Size(coregraphics.FromCGSizePointer(unsafe.Pointer(&result_)))
}

func (n NSView) AlignmentRectInsets() foundation.EdgeInsets {
	result_ := C.C_NSView_AlignmentRectInsets(n.Ptr())
	return foundation.FromNSEdgeInsetsPointer(unsafe.Pointer(&result_))
}

func (n NSView) BaselineOffsetFromBottom() coregraphics.Float {
	result_ := C.C_NSView_BaselineOffsetFromBottom(n.Ptr())
	return coregraphics.Float(float64(result_))
}

func (n NSView) FirstBaselineOffsetFromTop() coregraphics.Float {
	result_ := C.C_NSView_FirstBaselineOffsetFromTop(n.Ptr())
	return coregraphics.Float(float64(result_))
}

func (n NSView) LastBaselineOffsetFromBottom() coregraphics.Float {
	result_ := C.C_NSView_LastBaselineOffsetFromBottom(n.Ptr())
	return coregraphics.Float(float64(result_))
}

func (n NSView) NeedsLayout() bool {
	result_ := C.C_NSView_NeedsLayout(n.Ptr())
	return bool(result_)
}

func (n NSView) SetNeedsLayout(value bool) {
	C.C_NSView_SetNeedsLayout(n.Ptr(), C.bool(value))
}

func (n NSView) NeedsUpdateConstraints() bool {
	result_ := C.C_NSView_NeedsUpdateConstraints(n.Ptr())
	return bool(result_)
}

func (n NSView) SetNeedsUpdateConstraints(value bool) {
	C.C_NSView_SetNeedsUpdateConstraints(n.Ptr(), C.bool(value))
}

func View_RequiresConstraintBasedLayout() bool {
	result_ := C.C_NSView_View_RequiresConstraintBasedLayout()
	return bool(result_)
}

func (n NSView) TranslatesAutoresizingMaskIntoConstraints() bool {
	result_ := C.C_NSView_TranslatesAutoresizingMaskIntoConstraints(n.Ptr())
	return bool(result_)
}

func (n NSView) SetTranslatesAutoresizingMaskIntoConstraints(value bool) {
	C.C_NSView_SetTranslatesAutoresizingMaskIntoConstraints(n.Ptr(), C.bool(value))
}

func (n NSView) HasAmbiguousLayout() bool {
	result_ := C.C_NSView_HasAmbiguousLayout(n.Ptr())
	return bool(result_)
}

func FocusView() View {
	result_ := C.C_NSView_FocusView()
	return MakeView(result_)
}

func (n NSView) FocusRingType() FocusRingType {
	result_ := C.C_NSView_FocusRingType(n.Ptr())
	return FocusRingType(uint(result_))
}

func (n NSView) SetFocusRingType(value FocusRingType) {
	C.C_NSView_SetFocusRingType(n.Ptr(), C.uint(uint(value)))
}

func (n NSView) FocusRingMaskBounds() foundation.Rect {
	result_ := C.C_NSView_FocusRingMaskBounds(n.Ptr())
	return foundation.Rect(coregraphics.FromCGRectPointer(unsafe.Pointer(&result_)))
}

func View_DefaultFocusRingType() FocusRingType {
	result_ := C.C_NSView_View_DefaultFocusRingType()
	return FocusRingType(uint(result_))
}

func (n NSView) AllowsVibrancy() bool {
	result_ := C.C_NSView_AllowsVibrancy(n.Ptr())
	return bool(result_)
}

func (n NSView) IsInFullScreenMode() bool {
	result_ := C.C_NSView_IsInFullScreenMode(n.Ptr())
	return bool(result_)
}

func (n NSView) IsHidden() bool {
	result_ := C.C_NSView_IsHidden(n.Ptr())
	return bool(result_)
}

func (n NSView) SetHidden(value bool) {
	C.C_NSView_SetHidden(n.Ptr(), C.bool(value))
}

func (n NSView) IsHiddenOrHasHiddenAncestor() bool {
	result_ := C.C_NSView_IsHiddenOrHasHiddenAncestor(n.Ptr())
	return bool(result_)
}

func (n NSView) InLiveResize() bool {
	result_ := C.C_NSView_InLiveResize(n.Ptr())
	return bool(result_)
}

func (n NSView) PreservesContentDuringLiveResize() bool {
	result_ := C.C_NSView_PreservesContentDuringLiveResize(n.Ptr())
	return bool(result_)
}

func (n NSView) RectPreservedDuringLiveResize() foundation.Rect {
	result_ := C.C_NSView_RectPreservedDuringLiveResize(n.Ptr())
	return foundation.Rect(coregraphics.FromCGRectPointer(unsafe.Pointer(&result_)))
}

func (n NSView) GestureRecognizers() []GestureRecognizer {
	result_ := C.C_NSView_GestureRecognizers(n.Ptr())
	if result_.len > 0 {
		defer C.free(result_.data)
	}
	result_Slice := unsafe.Slice((*unsafe.Pointer)(result_.data), int(result_.len))
	var goResult_ = make([]GestureRecognizer, len(result_Slice))
	for idx, r := range result_Slice {
		goResult_[idx] = MakeGestureRecognizer(r)
	}
	return goResult_
}

func (n NSView) SetGestureRecognizers(value []GestureRecognizer) {
	var cValue C.Array
	if len(value) > 0 {
		cValueData := make([]unsafe.Pointer, len(value))
		for idx, v := range value {
			cValueData[idx] = objc.ExtractPtr(v)
		}
		cValue.data = unsafe.Pointer(&cValueData[0])
		cValue.len = C.int(len(value))
	}
	C.C_NSView_SetGestureRecognizers(n.Ptr(), cValue)
}

func (n NSView) MouseDownCanMoveWindow() bool {
	result_ := C.C_NSView_MouseDownCanMoveWindow(n.Ptr())
	return bool(result_)
}

func (n NSView) InputContext() TextInputContext {
	result_ := C.C_NSView_InputContext(n.Ptr())
	return MakeTextInputContext(result_)
}

func (n NSView) WantsRestingTouches() bool {
	result_ := C.C_NSView_WantsRestingTouches(n.Ptr())
	return bool(result_)
}

func (n NSView) SetWantsRestingTouches(value bool) {
	C.C_NSView_SetWantsRestingTouches(n.Ptr(), C.bool(value))
}

func (n NSView) CanBecomeKeyView() bool {
	result_ := C.C_NSView_CanBecomeKeyView(n.Ptr())
	return bool(result_)
}

func (n NSView) NeedsPanelToBecomeKey() bool {
	result_ := C.C_NSView_NeedsPanelToBecomeKey(n.Ptr())
	return bool(result_)
}

func (n NSView) NextKeyView() View {
	result_ := C.C_NSView_NextKeyView(n.Ptr())
	return MakeView(result_)
}

func (n NSView) SetNextKeyView(value View) {
	C.C_NSView_SetNextKeyView(n.Ptr(), objc.ExtractPtr(value))
}

func (n NSView) NextValidKeyView() View {
	result_ := C.C_NSView_NextValidKeyView(n.Ptr())
	return MakeView(result_)
}

func (n NSView) PreviousKeyView() View {
	result_ := C.C_NSView_PreviousKeyView(n.Ptr())
	return MakeView(result_)
}

func (n NSView) PreviousValidKeyView() View {
	result_ := C.C_NSView_PreviousValidKeyView(n.Ptr())
	return MakeView(result_)
}

func (n NSView) PreparedContentRect() foundation.Rect {
	result_ := C.C_NSView_PreparedContentRect(n.Ptr())
	return foundation.Rect(coregraphics.FromCGRectPointer(unsafe.Pointer(&result_)))
}

func (n NSView) SetPreparedContentRect(value foundation.Rect) {
	C.C_NSView_SetPreparedContentRect(n.Ptr(), *(*C.CGRect)(coregraphics.ToCGRectPointer(coregraphics.Rect(value))))
}

func (n NSView) EnclosingScrollView() ScrollView {
	result_ := C.C_NSView_EnclosingScrollView(n.Ptr())
	return MakeScrollView(result_)
}

func (n NSView) RegisteredDraggedTypes() []PasteboardType {
	result_ := C.C_NSView_RegisteredDraggedTypes(n.Ptr())
	if result_.len > 0 {
		defer C.free(result_.data)
	}
	result_Slice := unsafe.Slice((*unsafe.Pointer)(result_.data), int(result_.len))
	var goResult_ = make([]PasteboardType, len(result_Slice))
	for idx, r := range result_Slice {
		goResult_[idx] = PasteboardType(foundation.MakeString(r).String())
	}
	return goResult_
}

func (n NSView) PostsFrameChangedNotifications() bool {
	result_ := C.C_NSView_PostsFrameChangedNotifications(n.Ptr())
	return bool(result_)
}

func (n NSView) SetPostsFrameChangedNotifications(value bool) {
	C.C_NSView_SetPostsFrameChangedNotifications(n.Ptr(), C.bool(value))
}

func (n NSView) PostsBoundsChangedNotifications() bool {
	result_ := C.C_NSView_PostsBoundsChangedNotifications(n.Ptr())
	return bool(result_)
}

func (n NSView) SetPostsBoundsChangedNotifications(value bool) {
	C.C_NSView_SetPostsBoundsChangedNotifications(n.Ptr(), C.bool(value))
}

func (n NSView) Tag() int {
	result_ := C.C_NSView_Tag(n.Ptr())
	return int(result_)
}

func (n NSView) ToolTip() string {
	result_ := C.C_NSView_ToolTip(n.Ptr())
	return foundation.MakeString(result_).String()
}

func (n NSView) SetToolTip(value string) {
	C.C_NSView_SetToolTip(n.Ptr(), foundation.NewString(value).Ptr())
}

func (n NSView) TrackingAreas() []TrackingArea {
	result_ := C.C_NSView_TrackingAreas(n.Ptr())
	if result_.len > 0 {
		defer C.free(result_.data)
	}
	result_Slice := unsafe.Slice((*unsafe.Pointer)(result_.data), int(result_.len))
	var goResult_ = make([]TrackingArea, len(result_Slice))
	for idx, r := range result_Slice {
		goResult_[idx] = MakeTrackingArea(r)
	}
	return goResult_
}

func View_DefaultMenu() Menu {
	result_ := C.C_NSView_View_DefaultMenu()
	return MakeMenu(result_)
}

func (n NSView) IsDrawingFindIndicator() bool {
	result_ := C.C_NSView_IsDrawingFindIndicator(n.Ptr())
	return bool(result_)
}

func (n NSView) UserInterfaceLayoutDirection() UserInterfaceLayoutDirection {
	result_ := C.C_NSView_UserInterfaceLayoutDirection(n.Ptr())
	return UserInterfaceLayoutDirection(int(result_))
}

func (n NSView) SetUserInterfaceLayoutDirection(value UserInterfaceLayoutDirection) {
	C.C_NSView_SetUserInterfaceLayoutDirection(n.Ptr(), C.int(int(value)))
}

func (n NSView) PressureConfiguration() PressureConfiguration {
	result_ := C.C_NSView_PressureConfiguration(n.Ptr())
	return MakePressureConfiguration(result_)
}

func (n NSView) SetPressureConfiguration(value PressureConfiguration) {
	C.C_NSView_SetPressureConfiguration(n.Ptr(), objc.ExtractPtr(value))
}

func (n NSView) AdditionalSafeAreaInsets() foundation.EdgeInsets {
	result_ := C.C_NSView_AdditionalSafeAreaInsets(n.Ptr())
	return foundation.FromNSEdgeInsetsPointer(unsafe.Pointer(&result_))
}

func (n NSView) SetAdditionalSafeAreaInsets(value foundation.EdgeInsets) {
	C.C_NSView_SetAdditionalSafeAreaInsets(n.Ptr(), *(*C.NSEdgeInsets)(foundation.ToNSEdgeInsetsPointer(value)))
}

func (n NSView) AllowedTouchTypes() TouchTypeMask {
	result_ := C.C_NSView_AllowedTouchTypes(n.Ptr())
	return TouchTypeMask(uint(result_))
}

func (n NSView) SetAllowedTouchTypes(value TouchTypeMask) {
	C.C_NSView_SetAllowedTouchTypes(n.Ptr(), C.uint(uint(value)))
}

func (n NSView) CandidateListTouchBarItem() CandidateListTouchBarItem {
	result_ := C.C_NSView_CandidateListTouchBarItem(n.Ptr())
	return MakeCandidateListTouchBarItem(result_)
}

func (n NSView) IsHorizontalContentSizeConstraintActive() bool {
	result_ := C.C_NSView_IsHorizontalContentSizeConstraintActive(n.Ptr())
	return bool(result_)
}

func (n NSView) SetHorizontalContentSizeConstraintActive(value bool) {
	C.C_NSView_SetHorizontalContentSizeConstraintActive(n.Ptr(), C.bool(value))
}

func (n NSView) IsVerticalContentSizeConstraintActive() bool {
	result_ := C.C_NSView_IsVerticalContentSizeConstraintActive(n.Ptr())
	return bool(result_)
}

func (n NSView) SetVerticalContentSizeConstraintActive(value bool) {
	C.C_NSView_SetVerticalContentSizeConstraintActive(n.Ptr(), C.bool(value))
}

func (n NSView) SafeAreaInsets() foundation.EdgeInsets {
	result_ := C.C_NSView_SafeAreaInsets(n.Ptr())
	return foundation.FromNSEdgeInsetsPointer(unsafe.Pointer(&result_))
}

func (n NSView) SafeAreaRect() foundation.Rect {
	result_ := C.C_NSView_SafeAreaRect(n.Ptr())
	return foundation.Rect(coregraphics.FromCGRectPointer(unsafe.Pointer(&result_)))
}

func View_IsCompatibleWithResponsiveScrolling() bool {
	result_ := C.C_NSView_View_IsCompatibleWithResponsiveScrolling()
	return bool(result_)
}
