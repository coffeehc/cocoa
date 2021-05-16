package appkit

// #cgo CFLAGS: -x objective-c
// #cgo LDFLAGS: -framework AppKit
// #include "image.h"
import "C"
import (
	"github.com/hsiafan/cocoa/coregraphics"
	"github.com/hsiafan/cocoa/foundation"
	"github.com/hsiafan/cocoa/objc"
	"unsafe"
)

type Image interface {
	objc.Object
	SetName(_string ImageName) bool
	Name() ImageName
	ImageWithSymbolConfiguration(configuration ImageSymbolConfiguration) Image
	AddRepresentation(imageRep ImageRep)
	AddRepresentations(imageReps []ImageRep)
	RemoveRepresentation(imageRep ImageRep)
	DrawInRect(rect foundation.Rect)
	DrawAtPoint_FromRect_Operation_Fraction(point foundation.Point, fromRect foundation.Rect, op CompositingOperation, delta coregraphics.Float)
	DrawInRect_FromRect_Operation_Fraction(rect foundation.Rect, fromRect foundation.Rect, op CompositingOperation, delta coregraphics.Float)
	DrawRepresentation_InRect(imageRep ImageRep, rect foundation.Rect) bool
	LockFocus()
	LockFocusFlipped(flipped bool)
	UnlockFocus()
	Recache()
	TIFFRepresentationUsingCompression_Factor(comp TIFFCompression, factor float32) []byte
	CancelIncrementalLoad()
	LayerContentsForContentsScale(layerContentsScale coregraphics.Float) objc.Object
	RecommendedLayerContentsScale(preferredContentsScale coregraphics.Float) coregraphics.Float
	Size() foundation.Size
	SetSize(value foundation.Size)
	IsTemplate() bool
	SetTemplate(value bool)
	Representations() []ImageRep
	PrefersColorMatch() bool
	SetPrefersColorMatch(value bool)
	UsesEPSOnResolutionMismatch() bool
	SetUsesEPSOnResolutionMismatch(value bool)
	MatchesOnMultipleResolution() bool
	SetMatchesOnMultipleResolution(value bool)
	IsValid() bool
	BackgroundColor() Color
	SetBackgroundColor(value Color)
	CapInsets() foundation.EdgeInsets
	SetCapInsets(value foundation.EdgeInsets)
	ResizingMode() ImageResizingMode
	SetResizingMode(value ImageResizingMode)
	AlignmentRect() foundation.Rect
	SetAlignmentRect(value foundation.Rect)
	CacheMode() ImageCacheMode
	SetCacheMode(value ImageCacheMode)
	TIFFRepresentation() []byte
	AccessibilityDescription() string
	SetAccessibilityDescription(value string)
	MatchesOnlyOnBestFittingAxis() bool
	SetMatchesOnlyOnBestFittingAxis(value bool)
}

type NSImage struct {
	objc.NSObject
}

func MakeImage(ptr unsafe.Pointer) *NSImage {
	if ptr == nil {
		return nil
	}
	return &NSImage{
		NSObject: *objc.MakeObject(ptr),
	}
}

func AllocImage() *NSImage {
	return MakeImage(C.C_Image_Alloc())
}

func (n *NSImage) InitByReferencingFile(fileName string) Image {
	result := C.C_NSImage_InitByReferencingFile(n.Ptr(), foundation.NewString(fileName).Ptr())
	return MakeImage(result)
}

func (n *NSImage) InitByReferencingURL(url foundation.URL) Image {
	result := C.C_NSImage_InitByReferencingURL(n.Ptr(), objc.ExtractPtr(url))
	return MakeImage(result)
}

func (n *NSImage) InitWithContentsOfFile(fileName string) Image {
	result := C.C_NSImage_InitWithContentsOfFile(n.Ptr(), foundation.NewString(fileName).Ptr())
	return MakeImage(result)
}

func (n *NSImage) InitWithContentsOfURL(url foundation.URL) Image {
	result := C.C_NSImage_InitWithContentsOfURL(n.Ptr(), objc.ExtractPtr(url))
	return MakeImage(result)
}

func (n *NSImage) InitWithData(data []byte) Image {
	result := C.C_NSImage_InitWithData(n.Ptr(), C.Array{data: unsafe.Pointer(&data[0]), len: C.int(len(data))})
	return MakeImage(result)
}

func (n *NSImage) InitWithDataIgnoringOrientation(data []byte) Image {
	result := C.C_NSImage_InitWithDataIgnoringOrientation(n.Ptr(), C.Array{data: unsafe.Pointer(&data[0]), len: C.int(len(data))})
	return MakeImage(result)
}

func (n *NSImage) InitWithCGImage_Size(cgImage coregraphics.ImageRef, size foundation.Size) Image {
	result := C.C_NSImage_InitWithCGImage_Size(n.Ptr(), *(*C.CGImageRef)(coregraphics.ToCGImageRefPointer(cgImage)), *(*C.CGSize)(coregraphics.ToCGSizePointer(coregraphics.Size(size))))
	return MakeImage(result)
}

func (n *NSImage) InitWithPasteboard(pasteboard Pasteboard) Image {
	result := C.C_NSImage_InitWithPasteboard(n.Ptr(), objc.ExtractPtr(pasteboard))
	return MakeImage(result)
}

func (n *NSImage) InitWithCoder(coder foundation.Coder) Image {
	result := C.C_NSImage_InitWithCoder(n.Ptr(), objc.ExtractPtr(coder))
	return MakeImage(result)
}

func (n *NSImage) InitWithSize(size foundation.Size) Image {
	result := C.C_NSImage_InitWithSize(n.Ptr(), *(*C.CGSize)(coregraphics.ToCGSizePointer(coregraphics.Size(size))))
	return MakeImage(result)
}

func (n *NSImage) Init() Image {
	result := C.C_NSImage_Init(n.Ptr())
	return MakeImage(result)
}

func ImageNamed(name ImageName) Image {
	result := C.C_NSImage_ImageNamed(foundation.NewString(string(name)).Ptr())
	return MakeImage(result)
}

func (n *NSImage) SetName(_string ImageName) bool {
	result := C.C_NSImage_SetName(n.Ptr(), foundation.NewString(string(_string)).Ptr())
	return bool(result)
}

func (n *NSImage) Name() ImageName {
	result := C.C_NSImage_Name(n.Ptr())
	return ImageName(foundation.MakeString(result).String())
}

func ImageWithSystemSymbolName_AccessibilityDescription(symbolName string, description string) Image {
	result := C.C_NSImage_ImageWithSystemSymbolName_AccessibilityDescription(foundation.NewString(symbolName).Ptr(), foundation.NewString(description).Ptr())
	return MakeImage(result)
}

func (n *NSImage) ImageWithSymbolConfiguration(configuration ImageSymbolConfiguration) Image {
	result := C.C_NSImage_ImageWithSymbolConfiguration(n.Ptr(), objc.ExtractPtr(configuration))
	return MakeImage(result)
}

func Image_CanInitWithPasteboard(pasteboard Pasteboard) bool {
	result := C.C_NSImage_Image_CanInitWithPasteboard(objc.ExtractPtr(pasteboard))
	return bool(result)
}

func (n *NSImage) AddRepresentation(imageRep ImageRep) {
	C.C_NSImage_AddRepresentation(n.Ptr(), objc.ExtractPtr(imageRep))
}

func (n *NSImage) AddRepresentations(imageReps []ImageRep) {
	cImageRepsData := make([]unsafe.Pointer, len(imageReps))
	for idx, v := range imageReps {
		cImageRepsData[idx] = objc.ExtractPtr(v)
	}
	cImageReps := C.Array{data: unsafe.Pointer(&cImageRepsData[0]), len: C.int(len(imageReps))}
	C.C_NSImage_AddRepresentations(n.Ptr(), cImageReps)
}

func (n *NSImage) RemoveRepresentation(imageRep ImageRep) {
	C.C_NSImage_RemoveRepresentation(n.Ptr(), objc.ExtractPtr(imageRep))
}

func (n *NSImage) DrawInRect(rect foundation.Rect) {
	C.C_NSImage_DrawInRect(n.Ptr(), *(*C.CGRect)(coregraphics.ToCGRectPointer(coregraphics.Rect(rect))))
}

func (n *NSImage) DrawAtPoint_FromRect_Operation_Fraction(point foundation.Point, fromRect foundation.Rect, op CompositingOperation, delta coregraphics.Float) {
	C.C_NSImage_DrawAtPoint_FromRect_Operation_Fraction(n.Ptr(), *(*C.CGPoint)(coregraphics.ToCGPointPointer(coregraphics.Point(point))), *(*C.CGRect)(coregraphics.ToCGRectPointer(coregraphics.Rect(fromRect))), C.uint(uint(op)), C.double(float64(delta)))
}

func (n *NSImage) DrawInRect_FromRect_Operation_Fraction(rect foundation.Rect, fromRect foundation.Rect, op CompositingOperation, delta coregraphics.Float) {
	C.C_NSImage_DrawInRect_FromRect_Operation_Fraction(n.Ptr(), *(*C.CGRect)(coregraphics.ToCGRectPointer(coregraphics.Rect(rect))), *(*C.CGRect)(coregraphics.ToCGRectPointer(coregraphics.Rect(fromRect))), C.uint(uint(op)), C.double(float64(delta)))
}

func (n *NSImage) DrawRepresentation_InRect(imageRep ImageRep, rect foundation.Rect) bool {
	result := C.C_NSImage_DrawRepresentation_InRect(n.Ptr(), objc.ExtractPtr(imageRep), *(*C.CGRect)(coregraphics.ToCGRectPointer(coregraphics.Rect(rect))))
	return bool(result)
}

func (n *NSImage) LockFocus() {
	C.C_NSImage_LockFocus(n.Ptr())
}

func (n *NSImage) LockFocusFlipped(flipped bool) {
	C.C_NSImage_LockFocusFlipped(n.Ptr(), C.bool(flipped))
}

func (n *NSImage) UnlockFocus() {
	C.C_NSImage_UnlockFocus(n.Ptr())
}

func (n *NSImage) Recache() {
	C.C_NSImage_Recache(n.Ptr())
}

func (n *NSImage) TIFFRepresentationUsingCompression_Factor(comp TIFFCompression, factor float32) []byte {
	result := C.C_NSImage_TIFFRepresentationUsingCompression_Factor(n.Ptr(), C.uint(uint(comp)), C.float(factor))
	resultBuffer := (*[1 << 30]byte)(result.data)[:C.int(result.len)]
	goResult := make([]byte, C.int(result.len))
	copy(goResult, resultBuffer)
	C.free(result.data)
	return goResult
}

func (n *NSImage) CancelIncrementalLoad() {
	C.C_NSImage_CancelIncrementalLoad(n.Ptr())
}

func (n *NSImage) LayerContentsForContentsScale(layerContentsScale coregraphics.Float) objc.Object {
	result := C.C_NSImage_LayerContentsForContentsScale(n.Ptr(), C.double(float64(layerContentsScale)))
	return objc.MakeObject(result)
}

func (n *NSImage) RecommendedLayerContentsScale(preferredContentsScale coregraphics.Float) coregraphics.Float {
	result := C.C_NSImage_RecommendedLayerContentsScale(n.Ptr(), C.double(float64(preferredContentsScale)))
	return coregraphics.Float(float64(result))
}

func (n *NSImage) Size() foundation.Size {
	result := C.C_NSImage_Size(n.Ptr())
	return foundation.Size(coregraphics.FromCGSizePointer(unsafe.Pointer(&result)))
}

func (n *NSImage) SetSize(value foundation.Size) {
	C.C_NSImage_SetSize(n.Ptr(), *(*C.CGSize)(coregraphics.ToCGSizePointer(coregraphics.Size(value))))
}

func (n *NSImage) IsTemplate() bool {
	result := C.C_NSImage_IsTemplate(n.Ptr())
	return bool(result)
}

func (n *NSImage) SetTemplate(value bool) {
	C.C_NSImage_SetTemplate(n.Ptr(), C.bool(value))
}

func ImageTypes() []string {
	result := C.C_NSImage_ImageTypes()
	defer C.free(result.data)
	resultSlice := (*[1 << 28]unsafe.Pointer)(unsafe.Pointer(result.data))[:result.len:result.len]
	var goResult = make([]string, len(resultSlice))
	for idx, r := range resultSlice {
		goResult[idx] = foundation.MakeString(r).String()
	}
	return goResult
}

func ImageUnfilteredTypes() []string {
	result := C.C_NSImage_ImageUnfilteredTypes()
	defer C.free(result.data)
	resultSlice := (*[1 << 28]unsafe.Pointer)(unsafe.Pointer(result.data))[:result.len:result.len]
	var goResult = make([]string, len(resultSlice))
	for idx, r := range resultSlice {
		goResult[idx] = foundation.MakeString(r).String()
	}
	return goResult
}

func (n *NSImage) Representations() []ImageRep {
	result := C.C_NSImage_Representations(n.Ptr())
	defer C.free(result.data)
	resultSlice := (*[1 << 28]unsafe.Pointer)(unsafe.Pointer(result.data))[:result.len:result.len]
	var goResult = make([]ImageRep, len(resultSlice))
	for idx, r := range resultSlice {
		goResult[idx] = MakeImageRep(r)
	}
	return goResult
}

func (n *NSImage) PrefersColorMatch() bool {
	result := C.C_NSImage_PrefersColorMatch(n.Ptr())
	return bool(result)
}

func (n *NSImage) SetPrefersColorMatch(value bool) {
	C.C_NSImage_SetPrefersColorMatch(n.Ptr(), C.bool(value))
}

func (n *NSImage) UsesEPSOnResolutionMismatch() bool {
	result := C.C_NSImage_UsesEPSOnResolutionMismatch(n.Ptr())
	return bool(result)
}

func (n *NSImage) SetUsesEPSOnResolutionMismatch(value bool) {
	C.C_NSImage_SetUsesEPSOnResolutionMismatch(n.Ptr(), C.bool(value))
}

func (n *NSImage) MatchesOnMultipleResolution() bool {
	result := C.C_NSImage_MatchesOnMultipleResolution(n.Ptr())
	return bool(result)
}

func (n *NSImage) SetMatchesOnMultipleResolution(value bool) {
	C.C_NSImage_SetMatchesOnMultipleResolution(n.Ptr(), C.bool(value))
}

func (n *NSImage) IsValid() bool {
	result := C.C_NSImage_IsValid(n.Ptr())
	return bool(result)
}

func (n *NSImage) BackgroundColor() Color {
	result := C.C_NSImage_BackgroundColor(n.Ptr())
	return MakeColor(result)
}

func (n *NSImage) SetBackgroundColor(value Color) {
	C.C_NSImage_SetBackgroundColor(n.Ptr(), objc.ExtractPtr(value))
}

func (n *NSImage) CapInsets() foundation.EdgeInsets {
	result := C.C_NSImage_CapInsets(n.Ptr())
	return foundation.FromNSEdgeInsetsPointer(unsafe.Pointer(&result))
}

func (n *NSImage) SetCapInsets(value foundation.EdgeInsets) {
	C.C_NSImage_SetCapInsets(n.Ptr(), *(*C.NSEdgeInsets)(foundation.ToNSEdgeInsetsPointer(value)))
}

func (n *NSImage) ResizingMode() ImageResizingMode {
	result := C.C_NSImage_ResizingMode(n.Ptr())
	return ImageResizingMode(int(result))
}

func (n *NSImage) SetResizingMode(value ImageResizingMode) {
	C.C_NSImage_SetResizingMode(n.Ptr(), C.int(int(value)))
}

func (n *NSImage) AlignmentRect() foundation.Rect {
	result := C.C_NSImage_AlignmentRect(n.Ptr())
	return foundation.Rect(coregraphics.FromCGRectPointer(unsafe.Pointer(&result)))
}

func (n *NSImage) SetAlignmentRect(value foundation.Rect) {
	C.C_NSImage_SetAlignmentRect(n.Ptr(), *(*C.CGRect)(coregraphics.ToCGRectPointer(coregraphics.Rect(value))))
}

func (n *NSImage) CacheMode() ImageCacheMode {
	result := C.C_NSImage_CacheMode(n.Ptr())
	return ImageCacheMode(uint(result))
}

func (n *NSImage) SetCacheMode(value ImageCacheMode) {
	C.C_NSImage_SetCacheMode(n.Ptr(), C.uint(uint(value)))
}

func (n *NSImage) TIFFRepresentation() []byte {
	result := C.C_NSImage_TIFFRepresentation(n.Ptr())
	resultBuffer := (*[1 << 30]byte)(result.data)[:C.int(result.len)]
	goResult := make([]byte, C.int(result.len))
	copy(goResult, resultBuffer)
	C.free(result.data)
	return goResult
}

func (n *NSImage) AccessibilityDescription() string {
	result := C.C_NSImage_AccessibilityDescription(n.Ptr())
	return foundation.MakeString(result).String()
}

func (n *NSImage) SetAccessibilityDescription(value string) {
	C.C_NSImage_SetAccessibilityDescription(n.Ptr(), foundation.NewString(value).Ptr())
}

func (n *NSImage) MatchesOnlyOnBestFittingAxis() bool {
	result := C.C_NSImage_MatchesOnlyOnBestFittingAxis(n.Ptr())
	return bool(result)
}

func (n *NSImage) SetMatchesOnlyOnBestFittingAxis(value bool) {
	C.C_NSImage_SetMatchesOnlyOnBestFittingAxis(n.Ptr(), C.bool(value))
}
