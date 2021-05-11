package appkit

// #cgo CFLAGS: -x objective-c
// #cgo LDFLAGS: -framework AppKit
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
	RemoveConstraint(constraint LayoutConstraint)
	AddLayoutGuide(guide LayoutGuide)
	RemoveLayoutGuide(guide LayoutGuide)
	InvalidateIntrinsicContentSize()
	AlignmentRectForFrame(frame foundation.Rect) foundation.Rect
	FrameForAlignmentRect(alignmentRect foundation.Rect) foundation.Rect
	Layout()
	LayoutSubtreeIfNeeded()
	UpdateConstraints()
	UpdateConstraintsForSubtreeIfNeeded()
	ExerciseAmbiguityInLayout()
	DrawFocusRingMask()
	NoteFocusRingMaskChanged()
	SetKeyboardFocusRingNeedsDisplayInRect(rect foundation.Rect)
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
	UnregisterDraggedTypes()
	ShouldDelayWindowOrderingForEvent(event Event) bool
	RectForSmartMagnificationAtPoint_InRect(location foundation.Point, visibleRect foundation.Rect) foundation.Rect
	ViewDidChangeBackingProperties()
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
	FocusRingMaskBounds() foundation.Rect
	AllowsVibrancy() bool
	IsInFullScreenMode() bool
	IsHidden() bool
	SetHidden(value bool)
	IsHiddenOrHasHiddenAncestor() bool
	InLiveResize() bool
	PreservesContentDuringLiveResize() bool
	RectPreservedDuringLiveResize() foundation.Rect
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
	PostsFrameChangedNotifications() bool
	SetPostsFrameChangedNotifications(value bool)
	PostsBoundsChangedNotifications() bool
	SetPostsBoundsChangedNotifications(value bool)
	Tag() int
	ToolTip() string
	SetToolTip(value string)
	IsDrawingFindIndicator() bool
	PressureConfiguration() PressureConfiguration
	SetPressureConfiguration(value PressureConfiguration)
	AdditionalSafeAreaInsets() foundation.EdgeInsets
	SetAdditionalSafeAreaInsets(value foundation.EdgeInsets)
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

func MakeView(ptr unsafe.Pointer) *NSView {
	if ptr == nil {
		return nil
	}
	return &NSView{
		NSResponder: *MakeResponder(ptr),
	}
}

func AllocView() *NSView {
	return MakeView(C.C_View_Alloc())
}

func (n *NSView) InitWithFrame(frameRect foundation.Rect) View {
	result := C.C_NSView_InitWithFrame(n.Ptr(), *(*C.CGRect)(coregraphics.ToCGRectPointer(coregraphics.Rect(frameRect))))
	return MakeView(result)
}

func (n *NSView) InitWithCoder(coder foundation.Coder) View {
	result := C.C_NSView_InitWithCoder(n.Ptr(), objc.ExtractPtr(coder))
	return MakeView(result)
}

func (n *NSView) Init() View {
	result := C.C_NSView_Init(n.Ptr())
	return MakeView(result)
}

func (n *NSView) PrepareForReuse() {
	C.C_NSView_PrepareForReuse(n.Ptr())
}

func (n *NSView) AddSubview(view View) {
	C.C_NSView_AddSubview(n.Ptr(), objc.ExtractPtr(view))
}

func (n *NSView) DidAddSubview(subview View) {
	C.C_NSView_DidAddSubview(n.Ptr(), objc.ExtractPtr(subview))
}

func (n *NSView) RemoveFromSuperview() {
	C.C_NSView_RemoveFromSuperview(n.Ptr())
}

func (n *NSView) RemoveFromSuperviewWithoutNeedingDisplay() {
	C.C_NSView_RemoveFromSuperviewWithoutNeedingDisplay(n.Ptr())
}

func (n *NSView) ReplaceSubview_With(oldView View, newView View) {
	C.C_NSView_ReplaceSubview_With(n.Ptr(), objc.ExtractPtr(oldView), objc.ExtractPtr(newView))
}

func (n *NSView) IsDescendantOf(view View) bool {
	result := C.C_NSView_IsDescendantOf(n.Ptr(), objc.ExtractPtr(view))
	return bool(result)
}

func (n *NSView) AncestorSharedWithView(view View) View {
	result := C.C_NSView_AncestorSharedWithView(n.Ptr(), objc.ExtractPtr(view))
	return MakeView(result)
}

func (n *NSView) ViewDidMoveToSuperview() {
	C.C_NSView_ViewDidMoveToSuperview(n.Ptr())
}

func (n *NSView) ViewDidMoveToWindow() {
	C.C_NSView_ViewDidMoveToWindow(n.Ptr())
}

func (n *NSView) ViewWillMoveToSuperview(newSuperview View) {
	C.C_NSView_ViewWillMoveToSuperview(n.Ptr(), objc.ExtractPtr(newSuperview))
}

func (n *NSView) ViewWillMoveToWindow(newWindow Window) {
	C.C_NSView_ViewWillMoveToWindow(n.Ptr(), objc.ExtractPtr(newWindow))
}

func (n *NSView) WillRemoveSubview(subview View) {
	C.C_NSView_WillRemoveSubview(n.Ptr(), objc.ExtractPtr(subview))
}

func (n *NSView) SetFrameOrigin(newOrigin foundation.Point) {
	C.C_NSView_SetFrameOrigin(n.Ptr(), *(*C.CGPoint)(coregraphics.ToCGPointPointer(coregraphics.Point(newOrigin))))
}

func (n *NSView) SetFrameSize(newSize foundation.Size) {
	C.C_NSView_SetFrameSize(n.Ptr(), *(*C.CGSize)(coregraphics.ToCGSizePointer(coregraphics.Size(newSize))))
}

func (n *NSView) SetBoundsOrigin(newOrigin foundation.Point) {
	C.C_NSView_SetBoundsOrigin(n.Ptr(), *(*C.CGPoint)(coregraphics.ToCGPointPointer(coregraphics.Point(newOrigin))))
}

func (n *NSView) SetBoundsSize(newSize foundation.Size) {
	C.C_NSView_SetBoundsSize(n.Ptr(), *(*C.CGSize)(coregraphics.ToCGSizePointer(coregraphics.Size(newSize))))
}

func (n *NSView) UpdateLayer() {
	C.C_NSView_UpdateLayer(n.Ptr())
}

func (n *NSView) DrawRect(dirtyRect foundation.Rect) {
	C.C_NSView_DrawRect(n.Ptr(), *(*C.CGRect)(coregraphics.ToCGRectPointer(coregraphics.Rect(dirtyRect))))
}

func (n *NSView) NeedsToDrawRect(rect foundation.Rect) bool {
	result := C.C_NSView_NeedsToDrawRect(n.Ptr(), *(*C.CGRect)(coregraphics.ToCGRectPointer(coregraphics.Rect(rect))))
	return bool(result)
}

func (n *NSView) BitmapImageRepForCachingDisplayInRect(rect foundation.Rect) BitmapImageRep {
	result := C.C_NSView_BitmapImageRepForCachingDisplayInRect(n.Ptr(), *(*C.CGRect)(coregraphics.ToCGRectPointer(coregraphics.Rect(rect))))
	return MakeBitmapImageRep(result)
}

func (n *NSView) CacheDisplayInRect_ToBitmapImageRep(rect foundation.Rect, bitmapImageRep BitmapImageRep) {
	C.C_NSView_CacheDisplayInRect_ToBitmapImageRep(n.Ptr(), *(*C.CGRect)(coregraphics.ToCGRectPointer(coregraphics.Rect(rect))), objc.ExtractPtr(bitmapImageRep))
}

func (n *NSView) Print(sender objc.Object) {
	C.C_NSView_Print(n.Ptr(), objc.ExtractPtr(sender))
}

func (n *NSView) BeginPageInRect_AtPlacement(rect foundation.Rect, location foundation.Point) {
	C.C_NSView_BeginPageInRect_AtPlacement(n.Ptr(), *(*C.CGRect)(coregraphics.ToCGRectPointer(coregraphics.Rect(rect))), *(*C.CGPoint)(coregraphics.ToCGPointPointer(coregraphics.Point(location))))
}

func (n *NSView) DataWithEPSInsideRect(rect foundation.Rect) []byte {
	result := C.C_NSView_DataWithEPSInsideRect(n.Ptr(), *(*C.CGRect)(coregraphics.ToCGRectPointer(coregraphics.Rect(rect))))
	resultBuffer := (*[1 << 30]byte)(result.data)[:C.int(result.len)]
	goResult := make([]byte, C.int(result.len))
	copy(goResult, resultBuffer)
	C.free(result.data)
	return goResult
}

func (n *NSView) DataWithPDFInsideRect(rect foundation.Rect) []byte {
	result := C.C_NSView_DataWithPDFInsideRect(n.Ptr(), *(*C.CGRect)(coregraphics.ToCGRectPointer(coregraphics.Rect(rect))))
	resultBuffer := (*[1 << 30]byte)(result.data)[:C.int(result.len)]
	goResult := make([]byte, C.int(result.len))
	copy(goResult, resultBuffer)
	C.free(result.data)
	return goResult
}

func (n *NSView) WriteEPSInsideRect_ToPasteboard(rect foundation.Rect, pasteboard Pasteboard) {
	C.C_NSView_WriteEPSInsideRect_ToPasteboard(n.Ptr(), *(*C.CGRect)(coregraphics.ToCGRectPointer(coregraphics.Rect(rect))), objc.ExtractPtr(pasteboard))
}

func (n *NSView) WritePDFInsideRect_ToPasteboard(rect foundation.Rect, pasteboard Pasteboard) {
	C.C_NSView_WritePDFInsideRect_ToPasteboard(n.Ptr(), *(*C.CGRect)(coregraphics.ToCGRectPointer(coregraphics.Rect(rect))), objc.ExtractPtr(pasteboard))
}

func (n *NSView) DrawPageBorderWithSize(borderSize foundation.Size) {
	C.C_NSView_DrawPageBorderWithSize(n.Ptr(), *(*C.CGSize)(coregraphics.ToCGSizePointer(coregraphics.Size(borderSize))))
}

func (n *NSView) RectForPage(page int) foundation.Rect {
	result := C.C_NSView_RectForPage(n.Ptr(), C.int(page))
	return foundation.Rect(coregraphics.FromCGRectPointer(unsafe.Pointer(&result)))
}

func (n *NSView) LocationOfPrintRect(rect foundation.Rect) foundation.Point {
	result := C.C_NSView_LocationOfPrintRect(n.Ptr(), *(*C.CGRect)(coregraphics.ToCGRectPointer(coregraphics.Rect(rect))))
	return foundation.Point(coregraphics.FromCGPointPointer(unsafe.Pointer(&result)))
}

func (n *NSView) SetNeedsDisplayInRect(invalidRect foundation.Rect) {
	C.C_NSView_SetNeedsDisplayInRect(n.Ptr(), *(*C.CGRect)(coregraphics.ToCGRectPointer(coregraphics.Rect(invalidRect))))
}

func (n *NSView) Display() {
	C.C_NSView_Display(n.Ptr())
}

func (n *NSView) DisplayRect(rect foundation.Rect) {
	C.C_NSView_DisplayRect(n.Ptr(), *(*C.CGRect)(coregraphics.ToCGRectPointer(coregraphics.Rect(rect))))
}

func (n *NSView) DisplayRectIgnoringOpacity(rect foundation.Rect) {
	C.C_NSView_DisplayRectIgnoringOpacity(n.Ptr(), *(*C.CGRect)(coregraphics.ToCGRectPointer(coregraphics.Rect(rect))))
}

func (n *NSView) DisplayRectIgnoringOpacity_InContext(rect foundation.Rect, context GraphicsContext) {
	C.C_NSView_DisplayRectIgnoringOpacity_InContext(n.Ptr(), *(*C.CGRect)(coregraphics.ToCGRectPointer(coregraphics.Rect(rect))), objc.ExtractPtr(context))
}

func (n *NSView) DisplayIfNeeded() {
	C.C_NSView_DisplayIfNeeded(n.Ptr())
}

func (n *NSView) DisplayIfNeededInRect(rect foundation.Rect) {
	C.C_NSView_DisplayIfNeededInRect(n.Ptr(), *(*C.CGRect)(coregraphics.ToCGRectPointer(coregraphics.Rect(rect))))
}

func (n *NSView) DisplayIfNeededIgnoringOpacity() {
	C.C_NSView_DisplayIfNeededIgnoringOpacity(n.Ptr())
}

func (n *NSView) DisplayIfNeededInRectIgnoringOpacity(rect foundation.Rect) {
	C.C_NSView_DisplayIfNeededInRectIgnoringOpacity(n.Ptr(), *(*C.CGRect)(coregraphics.ToCGRectPointer(coregraphics.Rect(rect))))
}

func (n *NSView) TranslateRectsNeedingDisplayInRect_By(clipRect foundation.Rect, delta foundation.Size) {
	C.C_NSView_TranslateRectsNeedingDisplayInRect_By(n.Ptr(), *(*C.CGRect)(coregraphics.ToCGRectPointer(coregraphics.Rect(clipRect))), *(*C.CGSize)(coregraphics.ToCGSizePointer(coregraphics.Size(delta))))
}

func (n *NSView) ViewWillDraw() {
	C.C_NSView_ViewWillDraw(n.Ptr())
}

func (n *NSView) ConvertPointFromBacking(point foundation.Point) foundation.Point {
	result := C.C_NSView_ConvertPointFromBacking(n.Ptr(), *(*C.CGPoint)(coregraphics.ToCGPointPointer(coregraphics.Point(point))))
	return foundation.Point(coregraphics.FromCGPointPointer(unsafe.Pointer(&result)))
}

func (n *NSView) ConvertPointToBacking(point foundation.Point) foundation.Point {
	result := C.C_NSView_ConvertPointToBacking(n.Ptr(), *(*C.CGPoint)(coregraphics.ToCGPointPointer(coregraphics.Point(point))))
	return foundation.Point(coregraphics.FromCGPointPointer(unsafe.Pointer(&result)))
}

func (n *NSView) ConvertPointFromLayer(point foundation.Point) foundation.Point {
	result := C.C_NSView_ConvertPointFromLayer(n.Ptr(), *(*C.CGPoint)(coregraphics.ToCGPointPointer(coregraphics.Point(point))))
	return foundation.Point(coregraphics.FromCGPointPointer(unsafe.Pointer(&result)))
}

func (n *NSView) ConvertPointToLayer(point foundation.Point) foundation.Point {
	result := C.C_NSView_ConvertPointToLayer(n.Ptr(), *(*C.CGPoint)(coregraphics.ToCGPointPointer(coregraphics.Point(point))))
	return foundation.Point(coregraphics.FromCGPointPointer(unsafe.Pointer(&result)))
}

func (n *NSView) ConvertRectFromBacking(rect foundation.Rect) foundation.Rect {
	result := C.C_NSView_ConvertRectFromBacking(n.Ptr(), *(*C.CGRect)(coregraphics.ToCGRectPointer(coregraphics.Rect(rect))))
	return foundation.Rect(coregraphics.FromCGRectPointer(unsafe.Pointer(&result)))
}

func (n *NSView) ConvertRectToBacking(rect foundation.Rect) foundation.Rect {
	result := C.C_NSView_ConvertRectToBacking(n.Ptr(), *(*C.CGRect)(coregraphics.ToCGRectPointer(coregraphics.Rect(rect))))
	return foundation.Rect(coregraphics.FromCGRectPointer(unsafe.Pointer(&result)))
}

func (n *NSView) ConvertRectFromLayer(rect foundation.Rect) foundation.Rect {
	result := C.C_NSView_ConvertRectFromLayer(n.Ptr(), *(*C.CGRect)(coregraphics.ToCGRectPointer(coregraphics.Rect(rect))))
	return foundation.Rect(coregraphics.FromCGRectPointer(unsafe.Pointer(&result)))
}

func (n *NSView) ConvertRectToLayer(rect foundation.Rect) foundation.Rect {
	result := C.C_NSView_ConvertRectToLayer(n.Ptr(), *(*C.CGRect)(coregraphics.ToCGRectPointer(coregraphics.Rect(rect))))
	return foundation.Rect(coregraphics.FromCGRectPointer(unsafe.Pointer(&result)))
}

func (n *NSView) ConvertSizeFromBacking(size foundation.Size) foundation.Size {
	result := C.C_NSView_ConvertSizeFromBacking(n.Ptr(), *(*C.CGSize)(coregraphics.ToCGSizePointer(coregraphics.Size(size))))
	return foundation.Size(coregraphics.FromCGSizePointer(unsafe.Pointer(&result)))
}

func (n *NSView) ConvertSizeToBacking(size foundation.Size) foundation.Size {
	result := C.C_NSView_ConvertSizeToBacking(n.Ptr(), *(*C.CGSize)(coregraphics.ToCGSizePointer(coregraphics.Size(size))))
	return foundation.Size(coregraphics.FromCGSizePointer(unsafe.Pointer(&result)))
}

func (n *NSView) ConvertSizeFromLayer(size foundation.Size) foundation.Size {
	result := C.C_NSView_ConvertSizeFromLayer(n.Ptr(), *(*C.CGSize)(coregraphics.ToCGSizePointer(coregraphics.Size(size))))
	return foundation.Size(coregraphics.FromCGSizePointer(unsafe.Pointer(&result)))
}

func (n *NSView) ConvertSizeToLayer(size foundation.Size) foundation.Size {
	result := C.C_NSView_ConvertSizeToLayer(n.Ptr(), *(*C.CGSize)(coregraphics.ToCGSizePointer(coregraphics.Size(size))))
	return foundation.Size(coregraphics.FromCGSizePointer(unsafe.Pointer(&result)))
}

func (n *NSView) ConvertPoint_FromView(point foundation.Point, view View) foundation.Point {
	result := C.C_NSView_ConvertPoint_FromView(n.Ptr(), *(*C.CGPoint)(coregraphics.ToCGPointPointer(coregraphics.Point(point))), objc.ExtractPtr(view))
	return foundation.Point(coregraphics.FromCGPointPointer(unsafe.Pointer(&result)))
}

func (n *NSView) ConvertPoint_ToView(point foundation.Point, view View) foundation.Point {
	result := C.C_NSView_ConvertPoint_ToView(n.Ptr(), *(*C.CGPoint)(coregraphics.ToCGPointPointer(coregraphics.Point(point))), objc.ExtractPtr(view))
	return foundation.Point(coregraphics.FromCGPointPointer(unsafe.Pointer(&result)))
}

func (n *NSView) ConvertSize_FromView(size foundation.Size, view View) foundation.Size {
	result := C.C_NSView_ConvertSize_FromView(n.Ptr(), *(*C.CGSize)(coregraphics.ToCGSizePointer(coregraphics.Size(size))), objc.ExtractPtr(view))
	return foundation.Size(coregraphics.FromCGSizePointer(unsafe.Pointer(&result)))
}

func (n *NSView) ConvertSize_ToView(size foundation.Size, view View) foundation.Size {
	result := C.C_NSView_ConvertSize_ToView(n.Ptr(), *(*C.CGSize)(coregraphics.ToCGSizePointer(coregraphics.Size(size))), objc.ExtractPtr(view))
	return foundation.Size(coregraphics.FromCGSizePointer(unsafe.Pointer(&result)))
}

func (n *NSView) ConvertRect_FromView(rect foundation.Rect, view View) foundation.Rect {
	result := C.C_NSView_ConvertRect_FromView(n.Ptr(), *(*C.CGRect)(coregraphics.ToCGRectPointer(coregraphics.Rect(rect))), objc.ExtractPtr(view))
	return foundation.Rect(coregraphics.FromCGRectPointer(unsafe.Pointer(&result)))
}

func (n *NSView) ConvertRect_ToView(rect foundation.Rect, view View) foundation.Rect {
	result := C.C_NSView_ConvertRect_ToView(n.Ptr(), *(*C.CGRect)(coregraphics.ToCGRectPointer(coregraphics.Rect(rect))), objc.ExtractPtr(view))
	return foundation.Rect(coregraphics.FromCGRectPointer(unsafe.Pointer(&result)))
}

func (n *NSView) CenterScanRect(rect foundation.Rect) foundation.Rect {
	result := C.C_NSView_CenterScanRect(n.Ptr(), *(*C.CGRect)(coregraphics.ToCGRectPointer(coregraphics.Rect(rect))))
	return foundation.Rect(coregraphics.FromCGRectPointer(unsafe.Pointer(&result)))
}

func (n *NSView) TranslateOriginToPoint(translation foundation.Point) {
	C.C_NSView_TranslateOriginToPoint(n.Ptr(), *(*C.CGPoint)(coregraphics.ToCGPointPointer(coregraphics.Point(translation))))
}

func (n *NSView) ScaleUnitSquareToSize(newUnitSize foundation.Size) {
	C.C_NSView_ScaleUnitSquareToSize(n.Ptr(), *(*C.CGSize)(coregraphics.ToCGSizePointer(coregraphics.Size(newUnitSize))))
}

func (n *NSView) RotateByAngle(angle coregraphics.Float) {
	C.C_NSView_RotateByAngle(n.Ptr(), C.double(float64(angle)))
}

func (n *NSView) ResizeSubviewsWithOldSize(oldSize foundation.Size) {
	C.C_NSView_ResizeSubviewsWithOldSize(n.Ptr(), *(*C.CGSize)(coregraphics.ToCGSizePointer(coregraphics.Size(oldSize))))
}

func (n *NSView) ResizeWithOldSuperviewSize(oldSize foundation.Size) {
	C.C_NSView_ResizeWithOldSuperviewSize(n.Ptr(), *(*C.CGSize)(coregraphics.ToCGSizePointer(coregraphics.Size(oldSize))))
}

func (n *NSView) AddConstraint(constraint LayoutConstraint) {
	C.C_NSView_AddConstraint(n.Ptr(), objc.ExtractPtr(constraint))
}

func (n *NSView) RemoveConstraint(constraint LayoutConstraint) {
	C.C_NSView_RemoveConstraint(n.Ptr(), objc.ExtractPtr(constraint))
}

func (n *NSView) AddLayoutGuide(guide LayoutGuide) {
	C.C_NSView_AddLayoutGuide(n.Ptr(), objc.ExtractPtr(guide))
}

func (n *NSView) RemoveLayoutGuide(guide LayoutGuide) {
	C.C_NSView_RemoveLayoutGuide(n.Ptr(), objc.ExtractPtr(guide))
}

func (n *NSView) InvalidateIntrinsicContentSize() {
	C.C_NSView_InvalidateIntrinsicContentSize(n.Ptr())
}

func (n *NSView) AlignmentRectForFrame(frame foundation.Rect) foundation.Rect {
	result := C.C_NSView_AlignmentRectForFrame(n.Ptr(), *(*C.CGRect)(coregraphics.ToCGRectPointer(coregraphics.Rect(frame))))
	return foundation.Rect(coregraphics.FromCGRectPointer(unsafe.Pointer(&result)))
}

func (n *NSView) FrameForAlignmentRect(alignmentRect foundation.Rect) foundation.Rect {
	result := C.C_NSView_FrameForAlignmentRect(n.Ptr(), *(*C.CGRect)(coregraphics.ToCGRectPointer(coregraphics.Rect(alignmentRect))))
	return foundation.Rect(coregraphics.FromCGRectPointer(unsafe.Pointer(&result)))
}

func (n *NSView) Layout() {
	C.C_NSView_Layout(n.Ptr())
}

func (n *NSView) LayoutSubtreeIfNeeded() {
	C.C_NSView_LayoutSubtreeIfNeeded(n.Ptr())
}

func (n *NSView) UpdateConstraints() {
	C.C_NSView_UpdateConstraints(n.Ptr())
}

func (n *NSView) UpdateConstraintsForSubtreeIfNeeded() {
	C.C_NSView_UpdateConstraintsForSubtreeIfNeeded(n.Ptr())
}

func (n *NSView) ExerciseAmbiguityInLayout() {
	C.C_NSView_ExerciseAmbiguityInLayout(n.Ptr())
}

func (n *NSView) DrawFocusRingMask() {
	C.C_NSView_DrawFocusRingMask(n.Ptr())
}

func (n *NSView) NoteFocusRingMaskChanged() {
	C.C_NSView_NoteFocusRingMaskChanged(n.Ptr())
}

func (n *NSView) SetKeyboardFocusRingNeedsDisplayInRect(rect foundation.Rect) {
	C.C_NSView_SetKeyboardFocusRingNeedsDisplayInRect(n.Ptr(), *(*C.CGRect)(coregraphics.ToCGRectPointer(coregraphics.Rect(rect))))
}

func (n *NSView) ViewDidHide() {
	C.C_NSView_ViewDidHide(n.Ptr())
}

func (n *NSView) ViewDidUnhide() {
	C.C_NSView_ViewDidUnhide(n.Ptr())
}

func (n *NSView) ViewWillStartLiveResize() {
	C.C_NSView_ViewWillStartLiveResize(n.Ptr())
}

func (n *NSView) ViewDidEndLiveResize() {
	C.C_NSView_ViewDidEndLiveResize(n.Ptr())
}

func (n *NSView) AddGestureRecognizer(gestureRecognizer GestureRecognizer) {
	C.C_NSView_AddGestureRecognizer(n.Ptr(), objc.ExtractPtr(gestureRecognizer))
}

func (n *NSView) RemoveGestureRecognizer(gestureRecognizer GestureRecognizer) {
	C.C_NSView_RemoveGestureRecognizer(n.Ptr(), objc.ExtractPtr(gestureRecognizer))
}

func (n *NSView) AcceptsFirstMouse(event Event) bool {
	result := C.C_NSView_AcceptsFirstMouse(n.Ptr(), objc.ExtractPtr(event))
	return bool(result)
}

func (n *NSView) HitTest(point foundation.Point) View {
	result := C.C_NSView_HitTest(n.Ptr(), *(*C.CGPoint)(coregraphics.ToCGPointPointer(coregraphics.Point(point))))
	return MakeView(result)
}

func (n *NSView) Mouse_InRect(point foundation.Point, rect foundation.Rect) bool {
	result := C.C_NSView_Mouse_InRect(n.Ptr(), *(*C.CGPoint)(coregraphics.ToCGPointPointer(coregraphics.Point(point))), *(*C.CGRect)(coregraphics.ToCGRectPointer(coregraphics.Rect(rect))))
	return bool(result)
}

func (n *NSView) PrepareContentInRect(rect foundation.Rect) {
	C.C_NSView_PrepareContentInRect(n.Ptr(), *(*C.CGRect)(coregraphics.ToCGRectPointer(coregraphics.Rect(rect))))
}

func (n *NSView) ScrollPoint(point foundation.Point) {
	C.C_NSView_ScrollPoint(n.Ptr(), *(*C.CGPoint)(coregraphics.ToCGPointPointer(coregraphics.Point(point))))
}

func (n *NSView) ScrollRectToVisible(rect foundation.Rect) bool {
	result := C.C_NSView_ScrollRectToVisible(n.Ptr(), *(*C.CGRect)(coregraphics.ToCGRectPointer(coregraphics.Rect(rect))))
	return bool(result)
}

func (n *NSView) Autoscroll(event Event) bool {
	result := C.C_NSView_Autoscroll(n.Ptr(), objc.ExtractPtr(event))
	return bool(result)
}

func (n *NSView) AdjustScroll(newVisible foundation.Rect) foundation.Rect {
	result := C.C_NSView_AdjustScroll(n.Ptr(), *(*C.CGRect)(coregraphics.ToCGRectPointer(coregraphics.Rect(newVisible))))
	return foundation.Rect(coregraphics.FromCGRectPointer(unsafe.Pointer(&result)))
}

func (n *NSView) ScrollClipView_ToPoint(clipView ClipView, point foundation.Point) {
	C.C_NSView_ScrollClipView_ToPoint(n.Ptr(), objc.ExtractPtr(clipView), *(*C.CGPoint)(coregraphics.ToCGPointPointer(coregraphics.Point(point))))
}

func (n *NSView) ReflectScrolledClipView(clipView ClipView) {
	C.C_NSView_ReflectScrolledClipView(n.Ptr(), objc.ExtractPtr(clipView))
}

func (n *NSView) UnregisterDraggedTypes() {
	C.C_NSView_UnregisterDraggedTypes(n.Ptr())
}

func (n *NSView) ShouldDelayWindowOrderingForEvent(event Event) bool {
	result := C.C_NSView_ShouldDelayWindowOrderingForEvent(n.Ptr(), objc.ExtractPtr(event))
	return bool(result)
}

func (n *NSView) RectForSmartMagnificationAtPoint_InRect(location foundation.Point, visibleRect foundation.Rect) foundation.Rect {
	result := C.C_NSView_RectForSmartMagnificationAtPoint_InRect(n.Ptr(), *(*C.CGPoint)(coregraphics.ToCGPointPointer(coregraphics.Point(location))), *(*C.CGRect)(coregraphics.ToCGRectPointer(coregraphics.Rect(visibleRect))))
	return foundation.Rect(coregraphics.FromCGRectPointer(unsafe.Pointer(&result)))
}

func (n *NSView) ViewDidChangeBackingProperties() {
	C.C_NSView_ViewDidChangeBackingProperties(n.Ptr())
}

func (n *NSView) RemoveAllToolTips() {
	C.C_NSView_RemoveAllToolTips(n.Ptr())
}

func (n *NSView) RemoveToolTip(tag ToolTipTag) {
	C.C_NSView_RemoveToolTip(n.Ptr(), C.int(int(tag)))
}

func (n *NSView) RemoveTrackingRect(tag TrackingRectTag) {
	C.C_NSView_RemoveTrackingRect(n.Ptr(), C.int(int(tag)))
}

func (n *NSView) AddTrackingArea(trackingArea TrackingArea) {
	C.C_NSView_AddTrackingArea(n.Ptr(), objc.ExtractPtr(trackingArea))
}

func (n *NSView) RemoveTrackingArea(trackingArea TrackingArea) {
	C.C_NSView_RemoveTrackingArea(n.Ptr(), objc.ExtractPtr(trackingArea))
}

func (n *NSView) UpdateTrackingAreas() {
	C.C_NSView_UpdateTrackingAreas(n.Ptr())
}

func (n *NSView) AddCursorRect_Cursor(rect foundation.Rect, object Cursor) {
	C.C_NSView_AddCursorRect_Cursor(n.Ptr(), *(*C.CGRect)(coregraphics.ToCGRectPointer(coregraphics.Rect(rect))), objc.ExtractPtr(object))
}

func (n *NSView) RemoveCursorRect_Cursor(rect foundation.Rect, object Cursor) {
	C.C_NSView_RemoveCursorRect_Cursor(n.Ptr(), *(*C.CGRect)(coregraphics.ToCGRectPointer(coregraphics.Rect(rect))), objc.ExtractPtr(object))
}

func (n *NSView) DiscardCursorRects() {
	C.C_NSView_DiscardCursorRects(n.Ptr())
}

func (n *NSView) ResetCursorRects() {
	C.C_NSView_ResetCursorRects(n.Ptr())
}

func (n *NSView) MenuForEvent(event Event) Menu {
	result := C.C_NSView_MenuForEvent(n.Ptr(), objc.ExtractPtr(event))
	return MakeMenu(result)
}

func (n *NSView) WillOpenMenu_WithEvent(menu Menu, event Event) {
	C.C_NSView_WillOpenMenu_WithEvent(n.Ptr(), objc.ExtractPtr(menu), objc.ExtractPtr(event))
}

func (n *NSView) DidCloseMenu_WithEvent(menu Menu, event Event) {
	C.C_NSView_DidCloseMenu_WithEvent(n.Ptr(), objc.ExtractPtr(menu), objc.ExtractPtr(event))
}

func (n *NSView) BeginDocument() {
	C.C_NSView_BeginDocument(n.Ptr())
}

func (n *NSView) EndDocument() {
	C.C_NSView_EndDocument(n.Ptr())
}

func (n *NSView) EndPage() {
	C.C_NSView_EndPage(n.Ptr())
}

func (n *NSView) ShowDefinitionForAttributedString_AtPoint(attrString foundation.AttributedString, textBaselineOrigin foundation.Point) {
	C.C_NSView_ShowDefinitionForAttributedString_AtPoint(n.Ptr(), objc.ExtractPtr(attrString), *(*C.CGPoint)(coregraphics.ToCGPointPointer(coregraphics.Point(textBaselineOrigin))))
}

func (n *NSView) RulerView_DidAddMarker(ruler RulerView, marker RulerMarker) {
	C.C_NSView_RulerView_DidAddMarker(n.Ptr(), objc.ExtractPtr(ruler), objc.ExtractPtr(marker))
}

func (n *NSView) RulerView_DidMoveMarker(ruler RulerView, marker RulerMarker) {
	C.C_NSView_RulerView_DidMoveMarker(n.Ptr(), objc.ExtractPtr(ruler), objc.ExtractPtr(marker))
}

func (n *NSView) RulerView_DidRemoveMarker(ruler RulerView, marker RulerMarker) {
	C.C_NSView_RulerView_DidRemoveMarker(n.Ptr(), objc.ExtractPtr(ruler), objc.ExtractPtr(marker))
}

func (n *NSView) RulerView_HandleMouseDown(ruler RulerView, event Event) {
	C.C_NSView_RulerView_HandleMouseDown(n.Ptr(), objc.ExtractPtr(ruler), objc.ExtractPtr(event))
}

func (n *NSView) RulerView_LocationForPoint(ruler RulerView, point foundation.Point) coregraphics.Float {
	result := C.C_NSView_RulerView_LocationForPoint(n.Ptr(), objc.ExtractPtr(ruler), *(*C.CGPoint)(coregraphics.ToCGPointPointer(coregraphics.Point(point))))
	return coregraphics.Float(float64(result))
}

func (n *NSView) RulerView_PointForLocation(ruler RulerView, point coregraphics.Float) foundation.Point {
	result := C.C_NSView_RulerView_PointForLocation(n.Ptr(), objc.ExtractPtr(ruler), C.double(float64(point)))
	return foundation.Point(coregraphics.FromCGPointPointer(unsafe.Pointer(&result)))
}

func (n *NSView) RulerView_ShouldAddMarker(ruler RulerView, marker RulerMarker) bool {
	result := C.C_NSView_RulerView_ShouldAddMarker(n.Ptr(), objc.ExtractPtr(ruler), objc.ExtractPtr(marker))
	return bool(result)
}

func (n *NSView) RulerView_ShouldMoveMarker(ruler RulerView, marker RulerMarker) bool {
	result := C.C_NSView_RulerView_ShouldMoveMarker(n.Ptr(), objc.ExtractPtr(ruler), objc.ExtractPtr(marker))
	return bool(result)
}

func (n *NSView) RulerView_ShouldRemoveMarker(ruler RulerView, marker RulerMarker) bool {
	result := C.C_NSView_RulerView_ShouldRemoveMarker(n.Ptr(), objc.ExtractPtr(ruler), objc.ExtractPtr(marker))
	return bool(result)
}

func (n *NSView) RulerView_WillAddMarker_AtLocation(ruler RulerView, marker RulerMarker, location coregraphics.Float) coregraphics.Float {
	result := C.C_NSView_RulerView_WillAddMarker_AtLocation(n.Ptr(), objc.ExtractPtr(ruler), objc.ExtractPtr(marker), C.double(float64(location)))
	return coregraphics.Float(float64(result))
}

func (n *NSView) RulerView_WillMoveMarker_ToLocation(ruler RulerView, marker RulerMarker, location coregraphics.Float) coregraphics.Float {
	result := C.C_NSView_RulerView_WillMoveMarker_ToLocation(n.Ptr(), objc.ExtractPtr(ruler), objc.ExtractPtr(marker), C.double(float64(location)))
	return coregraphics.Float(float64(result))
}

func (n *NSView) RulerView_WillSetClientView(ruler RulerView, newClient View) {
	C.C_NSView_RulerView_WillSetClientView(n.Ptr(), objc.ExtractPtr(ruler), objc.ExtractPtr(newClient))
}

func (n *NSView) ViewDidChangeEffectiveAppearance() {
	C.C_NSView_ViewDidChangeEffectiveAppearance(n.Ptr())
}

func (n *NSView) Superview() View {
	result := C.C_NSView_Superview(n.Ptr())
	return MakeView(result)
}

func (n *NSView) Window() Window {
	result := C.C_NSView_Window(n.Ptr())
	return MakeWindow(result)
}

func (n *NSView) OpaqueAncestor() View {
	result := C.C_NSView_OpaqueAncestor(n.Ptr())
	return MakeView(result)
}

func (n *NSView) EnclosingMenuItem() MenuItem {
	result := C.C_NSView_EnclosingMenuItem(n.Ptr())
	return MakeMenuItem(result)
}

func (n *NSView) Frame() foundation.Rect {
	result := C.C_NSView_Frame(n.Ptr())
	return foundation.Rect(coregraphics.FromCGRectPointer(unsafe.Pointer(&result)))
}

func (n *NSView) SetFrame(value foundation.Rect) {
	C.C_NSView_SetFrame(n.Ptr(), *(*C.CGRect)(coregraphics.ToCGRectPointer(coregraphics.Rect(value))))
}

func (n *NSView) FrameRotation() coregraphics.Float {
	result := C.C_NSView_FrameRotation(n.Ptr())
	return coregraphics.Float(float64(result))
}

func (n *NSView) SetFrameRotation(value coregraphics.Float) {
	C.C_NSView_SetFrameRotation(n.Ptr(), C.double(float64(value)))
}

func (n *NSView) Bounds() foundation.Rect {
	result := C.C_NSView_Bounds(n.Ptr())
	return foundation.Rect(coregraphics.FromCGRectPointer(unsafe.Pointer(&result)))
}

func (n *NSView) SetBounds(value foundation.Rect) {
	C.C_NSView_SetBounds(n.Ptr(), *(*C.CGRect)(coregraphics.ToCGRectPointer(coregraphics.Rect(value))))
}

func (n *NSView) BoundsRotation() coregraphics.Float {
	result := C.C_NSView_BoundsRotation(n.Ptr())
	return coregraphics.Float(float64(result))
}

func (n *NSView) SetBoundsRotation(value coregraphics.Float) {
	C.C_NSView_SetBoundsRotation(n.Ptr(), C.double(float64(value)))
}

func (n *NSView) WantsLayer() bool {
	result := C.C_NSView_WantsLayer(n.Ptr())
	return bool(result)
}

func (n *NSView) SetWantsLayer(value bool) {
	C.C_NSView_SetWantsLayer(n.Ptr(), C.bool(value))
}

func (n *NSView) WantsUpdateLayer() bool {
	result := C.C_NSView_WantsUpdateLayer(n.Ptr())
	return bool(result)
}

func (n *NSView) CanDrawSubviewsIntoLayer() bool {
	result := C.C_NSView_CanDrawSubviewsIntoLayer(n.Ptr())
	return bool(result)
}

func (n *NSView) SetCanDrawSubviewsIntoLayer(value bool) {
	C.C_NSView_SetCanDrawSubviewsIntoLayer(n.Ptr(), C.bool(value))
}

func (n *NSView) LayerUsesCoreImageFilters() bool {
	result := C.C_NSView_LayerUsesCoreImageFilters(n.Ptr())
	return bool(result)
}

func (n *NSView) SetLayerUsesCoreImageFilters(value bool) {
	C.C_NSView_SetLayerUsesCoreImageFilters(n.Ptr(), C.bool(value))
}

func (n *NSView) AlphaValue() coregraphics.Float {
	result := C.C_NSView_AlphaValue(n.Ptr())
	return coregraphics.Float(float64(result))
}

func (n *NSView) SetAlphaValue(value coregraphics.Float) {
	C.C_NSView_SetAlphaValue(n.Ptr(), C.double(float64(value)))
}

func (n *NSView) FrameCenterRotation() coregraphics.Float {
	result := C.C_NSView_FrameCenterRotation(n.Ptr())
	return coregraphics.Float(float64(result))
}

func (n *NSView) SetFrameCenterRotation(value coregraphics.Float) {
	C.C_NSView_SetFrameCenterRotation(n.Ptr(), C.double(float64(value)))
}

func (n *NSView) Shadow() Shadow {
	result := C.C_NSView_Shadow(n.Ptr())
	return MakeShadow(result)
}

func (n *NSView) SetShadow(value Shadow) {
	C.C_NSView_SetShadow(n.Ptr(), objc.ExtractPtr(value))
}

func (n *NSView) CanDrawConcurrently() bool {
	result := C.C_NSView_CanDrawConcurrently(n.Ptr())
	return bool(result)
}

func (n *NSView) SetCanDrawConcurrently(value bool) {
	C.C_NSView_SetCanDrawConcurrently(n.Ptr(), C.bool(value))
}

func (n *NSView) VisibleRect() foundation.Rect {
	result := C.C_NSView_VisibleRect(n.Ptr())
	return foundation.Rect(coregraphics.FromCGRectPointer(unsafe.Pointer(&result)))
}

func (n *NSView) WantsDefaultClipping() bool {
	result := C.C_NSView_WantsDefaultClipping(n.Ptr())
	return bool(result)
}

func (n *NSView) PrintJobTitle() string {
	result := C.C_NSView_PrintJobTitle(n.Ptr())
	return foundation.MakeString(result).String()
}

func (n *NSView) PageHeader() foundation.AttributedString {
	result := C.C_NSView_PageHeader(n.Ptr())
	return foundation.MakeAttributedString(result)
}

func (n *NSView) PageFooter() foundation.AttributedString {
	result := C.C_NSView_PageFooter(n.Ptr())
	return foundation.MakeAttributedString(result)
}

func (n *NSView) HeightAdjustLimit() coregraphics.Float {
	result := C.C_NSView_HeightAdjustLimit(n.Ptr())
	return coregraphics.Float(float64(result))
}

func (n *NSView) WidthAdjustLimit() coregraphics.Float {
	result := C.C_NSView_WidthAdjustLimit(n.Ptr())
	return coregraphics.Float(float64(result))
}

func (n *NSView) NeedsDisplay() bool {
	result := C.C_NSView_NeedsDisplay(n.Ptr())
	return bool(result)
}

func (n *NSView) SetNeedsDisplay(value bool) {
	C.C_NSView_SetNeedsDisplay(n.Ptr(), C.bool(value))
}

func (n *NSView) IsOpaque() bool {
	result := C.C_NSView_IsOpaque(n.Ptr())
	return bool(result)
}

func (n *NSView) IsFlipped() bool {
	result := C.C_NSView_IsFlipped(n.Ptr())
	return bool(result)
}

func (n *NSView) IsRotatedFromBase() bool {
	result := C.C_NSView_IsRotatedFromBase(n.Ptr())
	return bool(result)
}

func (n *NSView) IsRotatedOrScaledFromBase() bool {
	result := C.C_NSView_IsRotatedOrScaledFromBase(n.Ptr())
	return bool(result)
}

func (n *NSView) AutoresizesSubviews() bool {
	result := C.C_NSView_AutoresizesSubviews(n.Ptr())
	return bool(result)
}

func (n *NSView) SetAutoresizesSubviews(value bool) {
	C.C_NSView_SetAutoresizesSubviews(n.Ptr(), C.bool(value))
}

func (n *NSView) BottomAnchor() LayoutYAxisAnchor {
	result := C.C_NSView_BottomAnchor(n.Ptr())
	return MakeLayoutYAxisAnchor(result)
}

func (n *NSView) CenterXAnchor() LayoutXAxisAnchor {
	result := C.C_NSView_CenterXAnchor(n.Ptr())
	return MakeLayoutXAxisAnchor(result)
}

func (n *NSView) CenterYAnchor() LayoutYAxisAnchor {
	result := C.C_NSView_CenterYAnchor(n.Ptr())
	return MakeLayoutYAxisAnchor(result)
}

func (n *NSView) FirstBaselineAnchor() LayoutYAxisAnchor {
	result := C.C_NSView_FirstBaselineAnchor(n.Ptr())
	return MakeLayoutYAxisAnchor(result)
}

func (n *NSView) HeightAnchor() LayoutDimension {
	result := C.C_NSView_HeightAnchor(n.Ptr())
	return MakeLayoutDimension(result)
}

func (n *NSView) LastBaselineAnchor() LayoutYAxisAnchor {
	result := C.C_NSView_LastBaselineAnchor(n.Ptr())
	return MakeLayoutYAxisAnchor(result)
}

func (n *NSView) LeadingAnchor() LayoutXAxisAnchor {
	result := C.C_NSView_LeadingAnchor(n.Ptr())
	return MakeLayoutXAxisAnchor(result)
}

func (n *NSView) LeftAnchor() LayoutXAxisAnchor {
	result := C.C_NSView_LeftAnchor(n.Ptr())
	return MakeLayoutXAxisAnchor(result)
}

func (n *NSView) RightAnchor() LayoutXAxisAnchor {
	result := C.C_NSView_RightAnchor(n.Ptr())
	return MakeLayoutXAxisAnchor(result)
}

func (n *NSView) TopAnchor() LayoutYAxisAnchor {
	result := C.C_NSView_TopAnchor(n.Ptr())
	return MakeLayoutYAxisAnchor(result)
}

func (n *NSView) TrailingAnchor() LayoutXAxisAnchor {
	result := C.C_NSView_TrailingAnchor(n.Ptr())
	return MakeLayoutXAxisAnchor(result)
}

func (n *NSView) WidthAnchor() LayoutDimension {
	result := C.C_NSView_WidthAnchor(n.Ptr())
	return MakeLayoutDimension(result)
}

func (n *NSView) LayoutMarginsGuide() LayoutGuide {
	result := C.C_NSView_LayoutMarginsGuide(n.Ptr())
	return MakeLayoutGuide(result)
}

func (n *NSView) SafeAreaLayoutGuide() LayoutGuide {
	result := C.C_NSView_SafeAreaLayoutGuide(n.Ptr())
	return MakeLayoutGuide(result)
}

func (n *NSView) FittingSize() foundation.Size {
	result := C.C_NSView_FittingSize(n.Ptr())
	return foundation.Size(coregraphics.FromCGSizePointer(unsafe.Pointer(&result)))
}

func (n *NSView) IntrinsicContentSize() foundation.Size {
	result := C.C_NSView_IntrinsicContentSize(n.Ptr())
	return foundation.Size(coregraphics.FromCGSizePointer(unsafe.Pointer(&result)))
}

func (n *NSView) AlignmentRectInsets() foundation.EdgeInsets {
	result := C.C_NSView_AlignmentRectInsets(n.Ptr())
	return foundation.FromNSEdgeInsetsPointer(unsafe.Pointer(&result))
}

func (n *NSView) BaselineOffsetFromBottom() coregraphics.Float {
	result := C.C_NSView_BaselineOffsetFromBottom(n.Ptr())
	return coregraphics.Float(float64(result))
}

func (n *NSView) FirstBaselineOffsetFromTop() coregraphics.Float {
	result := C.C_NSView_FirstBaselineOffsetFromTop(n.Ptr())
	return coregraphics.Float(float64(result))
}

func (n *NSView) LastBaselineOffsetFromBottom() coregraphics.Float {
	result := C.C_NSView_LastBaselineOffsetFromBottom(n.Ptr())
	return coregraphics.Float(float64(result))
}

func (n *NSView) NeedsLayout() bool {
	result := C.C_NSView_NeedsLayout(n.Ptr())
	return bool(result)
}

func (n *NSView) SetNeedsLayout(value bool) {
	C.C_NSView_SetNeedsLayout(n.Ptr(), C.bool(value))
}

func (n *NSView) NeedsUpdateConstraints() bool {
	result := C.C_NSView_NeedsUpdateConstraints(n.Ptr())
	return bool(result)
}

func (n *NSView) SetNeedsUpdateConstraints(value bool) {
	C.C_NSView_SetNeedsUpdateConstraints(n.Ptr(), C.bool(value))
}

func (n *NSView) TranslatesAutoresizingMaskIntoConstraints() bool {
	result := C.C_NSView_TranslatesAutoresizingMaskIntoConstraints(n.Ptr())
	return bool(result)
}

func (n *NSView) SetTranslatesAutoresizingMaskIntoConstraints(value bool) {
	C.C_NSView_SetTranslatesAutoresizingMaskIntoConstraints(n.Ptr(), C.bool(value))
}

func (n *NSView) HasAmbiguousLayout() bool {
	result := C.C_NSView_HasAmbiguousLayout(n.Ptr())
	return bool(result)
}

func (n *NSView) FocusRingMaskBounds() foundation.Rect {
	result := C.C_NSView_FocusRingMaskBounds(n.Ptr())
	return foundation.Rect(coregraphics.FromCGRectPointer(unsafe.Pointer(&result)))
}

func (n *NSView) AllowsVibrancy() bool {
	result := C.C_NSView_AllowsVibrancy(n.Ptr())
	return bool(result)
}

func (n *NSView) IsInFullScreenMode() bool {
	result := C.C_NSView_IsInFullScreenMode(n.Ptr())
	return bool(result)
}

func (n *NSView) IsHidden() bool {
	result := C.C_NSView_IsHidden(n.Ptr())
	return bool(result)
}

func (n *NSView) SetHidden(value bool) {
	C.C_NSView_SetHidden(n.Ptr(), C.bool(value))
}

func (n *NSView) IsHiddenOrHasHiddenAncestor() bool {
	result := C.C_NSView_IsHiddenOrHasHiddenAncestor(n.Ptr())
	return bool(result)
}

func (n *NSView) InLiveResize() bool {
	result := C.C_NSView_InLiveResize(n.Ptr())
	return bool(result)
}

func (n *NSView) PreservesContentDuringLiveResize() bool {
	result := C.C_NSView_PreservesContentDuringLiveResize(n.Ptr())
	return bool(result)
}

func (n *NSView) RectPreservedDuringLiveResize() foundation.Rect {
	result := C.C_NSView_RectPreservedDuringLiveResize(n.Ptr())
	return foundation.Rect(coregraphics.FromCGRectPointer(unsafe.Pointer(&result)))
}

func (n *NSView) MouseDownCanMoveWindow() bool {
	result := C.C_NSView_MouseDownCanMoveWindow(n.Ptr())
	return bool(result)
}

func (n *NSView) InputContext() TextInputContext {
	result := C.C_NSView_InputContext(n.Ptr())
	return MakeTextInputContext(result)
}

func (n *NSView) WantsRestingTouches() bool {
	result := C.C_NSView_WantsRestingTouches(n.Ptr())
	return bool(result)
}

func (n *NSView) SetWantsRestingTouches(value bool) {
	C.C_NSView_SetWantsRestingTouches(n.Ptr(), C.bool(value))
}

func (n *NSView) CanBecomeKeyView() bool {
	result := C.C_NSView_CanBecomeKeyView(n.Ptr())
	return bool(result)
}

func (n *NSView) NeedsPanelToBecomeKey() bool {
	result := C.C_NSView_NeedsPanelToBecomeKey(n.Ptr())
	return bool(result)
}

func (n *NSView) NextKeyView() View {
	result := C.C_NSView_NextKeyView(n.Ptr())
	return MakeView(result)
}

func (n *NSView) SetNextKeyView(value View) {
	C.C_NSView_SetNextKeyView(n.Ptr(), objc.ExtractPtr(value))
}

func (n *NSView) NextValidKeyView() View {
	result := C.C_NSView_NextValidKeyView(n.Ptr())
	return MakeView(result)
}

func (n *NSView) PreviousKeyView() View {
	result := C.C_NSView_PreviousKeyView(n.Ptr())
	return MakeView(result)
}

func (n *NSView) PreviousValidKeyView() View {
	result := C.C_NSView_PreviousValidKeyView(n.Ptr())
	return MakeView(result)
}

func (n *NSView) PreparedContentRect() foundation.Rect {
	result := C.C_NSView_PreparedContentRect(n.Ptr())
	return foundation.Rect(coregraphics.FromCGRectPointer(unsafe.Pointer(&result)))
}

func (n *NSView) SetPreparedContentRect(value foundation.Rect) {
	C.C_NSView_SetPreparedContentRect(n.Ptr(), *(*C.CGRect)(coregraphics.ToCGRectPointer(coregraphics.Rect(value))))
}

func (n *NSView) EnclosingScrollView() ScrollView {
	result := C.C_NSView_EnclosingScrollView(n.Ptr())
	return MakeScrollView(result)
}

func (n *NSView) PostsFrameChangedNotifications() bool {
	result := C.C_NSView_PostsFrameChangedNotifications(n.Ptr())
	return bool(result)
}

func (n *NSView) SetPostsFrameChangedNotifications(value bool) {
	C.C_NSView_SetPostsFrameChangedNotifications(n.Ptr(), C.bool(value))
}

func (n *NSView) PostsBoundsChangedNotifications() bool {
	result := C.C_NSView_PostsBoundsChangedNotifications(n.Ptr())
	return bool(result)
}

func (n *NSView) SetPostsBoundsChangedNotifications(value bool) {
	C.C_NSView_SetPostsBoundsChangedNotifications(n.Ptr(), C.bool(value))
}

func (n *NSView) Tag() int {
	result := C.C_NSView_Tag(n.Ptr())
	return int(result)
}

func (n *NSView) ToolTip() string {
	result := C.C_NSView_ToolTip(n.Ptr())
	return foundation.MakeString(result).String()
}

func (n *NSView) SetToolTip(value string) {
	C.C_NSView_SetToolTip(n.Ptr(), foundation.NewString(value).Ptr())
}

func (n *NSView) IsDrawingFindIndicator() bool {
	result := C.C_NSView_IsDrawingFindIndicator(n.Ptr())
	return bool(result)
}

func (n *NSView) PressureConfiguration() PressureConfiguration {
	result := C.C_NSView_PressureConfiguration(n.Ptr())
	return MakePressureConfiguration(result)
}

func (n *NSView) SetPressureConfiguration(value PressureConfiguration) {
	C.C_NSView_SetPressureConfiguration(n.Ptr(), objc.ExtractPtr(value))
}

func (n *NSView) AdditionalSafeAreaInsets() foundation.EdgeInsets {
	result := C.C_NSView_AdditionalSafeAreaInsets(n.Ptr())
	return foundation.FromNSEdgeInsetsPointer(unsafe.Pointer(&result))
}

func (n *NSView) SetAdditionalSafeAreaInsets(value foundation.EdgeInsets) {
	C.C_NSView_SetAdditionalSafeAreaInsets(n.Ptr(), *(*C.NSEdgeInsets)(foundation.ToNSEdgeInsetsPointer(value)))
}

func (n *NSView) CandidateListTouchBarItem() CandidateListTouchBarItem {
	result := C.C_NSView_CandidateListTouchBarItem(n.Ptr())
	return MakeCandidateListTouchBarItem(result)
}

func (n *NSView) IsHorizontalContentSizeConstraintActive() bool {
	result := C.C_NSView_IsHorizontalContentSizeConstraintActive(n.Ptr())
	return bool(result)
}

func (n *NSView) SetHorizontalContentSizeConstraintActive(value bool) {
	C.C_NSView_SetHorizontalContentSizeConstraintActive(n.Ptr(), C.bool(value))
}

func (n *NSView) IsVerticalContentSizeConstraintActive() bool {
	result := C.C_NSView_IsVerticalContentSizeConstraintActive(n.Ptr())
	return bool(result)
}

func (n *NSView) SetVerticalContentSizeConstraintActive(value bool) {
	C.C_NSView_SetVerticalContentSizeConstraintActive(n.Ptr(), C.bool(value))
}

func (n *NSView) SafeAreaInsets() foundation.EdgeInsets {
	result := C.C_NSView_SafeAreaInsets(n.Ptr())
	return foundation.FromNSEdgeInsetsPointer(unsafe.Pointer(&result))
}

func (n *NSView) SafeAreaRect() foundation.Rect {
	result := C.C_NSView_SafeAreaRect(n.Ptr())
	return foundation.Rect(coregraphics.FromCGRectPointer(unsafe.Pointer(&result)))
}
