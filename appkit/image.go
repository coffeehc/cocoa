package appkit

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
	extImage
	SetName(_string ImageName) bool
	Name() ImageName
	ImageWithSymbolConfiguration(configuration ImageSymbolConfiguration) Image
	AddRepresentation(imageRep ImageRep)
	AddRepresentations(imageReps []ImageRep)
	RemoveRepresentation(imageRep ImageRep)
	BestRepresentationForRect_Context_Hints(rect foundation.Rect, referenceContext GraphicsContext, hints map[ImageHintKey]objc.Object) ImageRep
	DrawInRect(rect foundation.Rect)
	DrawAtPoint_FromRect_Operation_Fraction(point foundation.Point, fromRect foundation.Rect, op CompositingOperation, delta coregraphics.Float)
	DrawInRect_FromRect_Operation_Fraction(rect foundation.Rect, fromRect foundation.Rect, op CompositingOperation, delta coregraphics.Float)
	DrawInRect_FromRect_Operation_Fraction_RespectFlipped_Hints(dstSpacePortionRect foundation.Rect, srcSpacePortionRect foundation.Rect, op CompositingOperation, requestedAlpha coregraphics.Float, respectContextIsFlipped bool, hints map[ImageHintKey]objc.Object)
	DrawRepresentation_InRect(imageRep ImageRep, rect foundation.Rect) bool
	LockFocus()
	LockFocusFlipped(flipped bool)
	UnlockFocus()
	Recache()
	TIFFRepresentationUsingCompression_Factor(comp TIFFCompression, factor float32) []byte
	CancelIncrementalLoad()
	HitTestRect_WithImageDestinationRect_Context_Hints_Flipped(testRectDestSpace foundation.Rect, imageRectDestSpace foundation.Rect, context GraphicsContext, hints map[ImageHintKey]objc.Object, flipped bool) bool
	LayerContentsForContentsScale(layerContentsScale coregraphics.Float) objc.Object
	RecommendedLayerContentsScale(preferredContentsScale coregraphics.Float) coregraphics.Float
	Delegate() objc.Object
	SetDelegate(value objc.Object)
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

func MakeImage(ptr unsafe.Pointer) NSImage {
	return NSImage{
		NSObject: objc.MakeObject(ptr),
	}
}

func ImageWithSystemSymbolName_AccessibilityDescription(symbolName string, description string) NSImage {
	result_ := C.C_NSImage_ImageWithSystemSymbolName_AccessibilityDescription(foundation.NewString(symbolName).Ptr(), foundation.NewString(description).Ptr())
	return MakeImage(result_)
}

func (n NSImage) InitByReferencingFile(fileName string) NSImage {
	result_ := C.C_NSImage_InitByReferencingFile(n.Ptr(), foundation.NewString(fileName).Ptr())
	return MakeImage(result_)
}

func (n NSImage) InitByReferencingURL(url foundation.URL) NSImage {
	result_ := C.C_NSImage_InitByReferencingURL(n.Ptr(), objc.ExtractPtr(url))
	return MakeImage(result_)
}

func (n NSImage) InitWithContentsOfFile(fileName string) NSImage {
	result_ := C.C_NSImage_InitWithContentsOfFile(n.Ptr(), foundation.NewString(fileName).Ptr())
	return MakeImage(result_)
}

func (n NSImage) InitWithContentsOfURL(url foundation.URL) NSImage {
	result_ := C.C_NSImage_InitWithContentsOfURL(n.Ptr(), objc.ExtractPtr(url))
	return MakeImage(result_)
}

func (n NSImage) InitWithData(data []byte) NSImage {
	result_ := C.C_NSImage_InitWithData(n.Ptr(), foundation.NewData(data).Ptr())
	return MakeImage(result_)
}

func (n NSImage) InitWithDataIgnoringOrientation(data []byte) NSImage {
	result_ := C.C_NSImage_InitWithDataIgnoringOrientation(n.Ptr(), foundation.NewData(data).Ptr())
	return MakeImage(result_)
}

func (n NSImage) InitWithCGImage_Size(cgImage coregraphics.ImageRef, size foundation.Size) NSImage {
	result_ := C.C_NSImage_InitWithCGImage_Size(n.Ptr(), unsafe.Pointer(cgImage), *(*C.CGSize)(coregraphics.ToCGSizePointer(coregraphics.Size(size))))
	return MakeImage(result_)
}

func (n NSImage) InitWithPasteboard(pasteboard Pasteboard) NSImage {
	result_ := C.C_NSImage_InitWithPasteboard(n.Ptr(), objc.ExtractPtr(pasteboard))
	return MakeImage(result_)
}

func (n NSImage) InitWithCoder(coder foundation.Coder) NSImage {
	result_ := C.C_NSImage_InitWithCoder(n.Ptr(), objc.ExtractPtr(coder))
	return MakeImage(result_)
}

func (n NSImage) InitWithSize(size foundation.Size) NSImage {
	result_ := C.C_NSImage_InitWithSize(n.Ptr(), *(*C.CGSize)(coregraphics.ToCGSizePointer(coregraphics.Size(size))))
	return MakeImage(result_)
}

func AllocImage() NSImage {
	result_ := C.C_NSImage_AllocImage()
	return MakeImage(result_)
}

func (n NSImage) Init() NSImage {
	result_ := C.C_NSImage_Init(n.Ptr())
	return MakeImage(result_)
}

func NewImage() NSImage {
	result_ := C.C_NSImage_NewImage()
	return MakeImage(result_)
}

func (n NSImage) Autorelease() NSImage {
	result_ := C.C_NSImage_Autorelease(n.Ptr())
	return MakeImage(result_)
}

func (n NSImage) Retain() NSImage {
	result_ := C.C_NSImage_Retain(n.Ptr())
	return MakeImage(result_)
}

func ImageNamed(name ImageName) Image {
	result_ := C.C_NSImage_ImageNamed(foundation.NewString(string(name)).Ptr())
	return MakeImage(result_)
}

func (n NSImage) SetName(_string ImageName) bool {
	result_ := C.C_NSImage_SetName(n.Ptr(), foundation.NewString(string(_string)).Ptr())
	return bool(result_)
}

func (n NSImage) Name() ImageName {
	result_ := C.C_NSImage_Name(n.Ptr())
	return ImageName(foundation.MakeString(result_).String())
}

func (n NSImage) ImageWithSymbolConfiguration(configuration ImageSymbolConfiguration) Image {
	result_ := C.C_NSImage_ImageWithSymbolConfiguration(n.Ptr(), objc.ExtractPtr(configuration))
	return MakeImage(result_)
}

func Image_CanInitWithPasteboard(pasteboard Pasteboard) bool {
	result_ := C.C_NSImage_Image_CanInitWithPasteboard(objc.ExtractPtr(pasteboard))
	return bool(result_)
}

func (n NSImage) AddRepresentation(imageRep ImageRep) {
	C.C_NSImage_AddRepresentation(n.Ptr(), objc.ExtractPtr(imageRep))
}

func (n NSImage) AddRepresentations(imageReps []ImageRep) {
	var cImageReps C.Array
	if len(imageReps) > 0 {
		cImageRepsData := make([]unsafe.Pointer, len(imageReps))
		for idx, v := range imageReps {
			cImageRepsData[idx] = objc.ExtractPtr(v)
		}
		cImageReps.data = unsafe.Pointer(&cImageRepsData[0])
		cImageReps.len = C.int(len(imageReps))
	}
	C.C_NSImage_AddRepresentations(n.Ptr(), cImageReps)
}

func (n NSImage) RemoveRepresentation(imageRep ImageRep) {
	C.C_NSImage_RemoveRepresentation(n.Ptr(), objc.ExtractPtr(imageRep))
}

func (n NSImage) BestRepresentationForRect_Context_Hints(rect foundation.Rect, referenceContext GraphicsContext, hints map[ImageHintKey]objc.Object) ImageRep {
	var cHints C.Dictionary
	if len(hints) > 0 {
		cHintsKeyData := make([]unsafe.Pointer, len(hints))
		cHintsValueData := make([]unsafe.Pointer, len(hints))
		var idx = 0
		for k, v := range hints {
			cHintsKeyData[idx] = foundation.NewString(string(k)).Ptr()
			cHintsValueData[idx] = objc.ExtractPtr(v)
			idx++
		}
		cHints.key_data = unsafe.Pointer(&cHintsKeyData[0])
		cHints.value_data = unsafe.Pointer(&cHintsValueData[0])
		cHints.len = C.int(len(hints))
	}
	result_ := C.C_NSImage_BestRepresentationForRect_Context_Hints(n.Ptr(), *(*C.CGRect)(coregraphics.ToCGRectPointer(coregraphics.Rect(rect))), objc.ExtractPtr(referenceContext), cHints)
	return MakeImageRep(result_)
}

func (n NSImage) DrawInRect(rect foundation.Rect) {
	C.C_NSImage_DrawInRect(n.Ptr(), *(*C.CGRect)(coregraphics.ToCGRectPointer(coregraphics.Rect(rect))))
}

func (n NSImage) DrawAtPoint_FromRect_Operation_Fraction(point foundation.Point, fromRect foundation.Rect, op CompositingOperation, delta coregraphics.Float) {
	C.C_NSImage_DrawAtPoint_FromRect_Operation_Fraction(n.Ptr(), *(*C.CGPoint)(coregraphics.ToCGPointPointer(coregraphics.Point(point))), *(*C.CGRect)(coregraphics.ToCGRectPointer(coregraphics.Rect(fromRect))), C.uint(uint(op)), C.double(float64(delta)))
}

func (n NSImage) DrawInRect_FromRect_Operation_Fraction(rect foundation.Rect, fromRect foundation.Rect, op CompositingOperation, delta coregraphics.Float) {
	C.C_NSImage_DrawInRect_FromRect_Operation_Fraction(n.Ptr(), *(*C.CGRect)(coregraphics.ToCGRectPointer(coregraphics.Rect(rect))), *(*C.CGRect)(coregraphics.ToCGRectPointer(coregraphics.Rect(fromRect))), C.uint(uint(op)), C.double(float64(delta)))
}

func (n NSImage) DrawInRect_FromRect_Operation_Fraction_RespectFlipped_Hints(dstSpacePortionRect foundation.Rect, srcSpacePortionRect foundation.Rect, op CompositingOperation, requestedAlpha coregraphics.Float, respectContextIsFlipped bool, hints map[ImageHintKey]objc.Object) {
	var cHints C.Dictionary
	if len(hints) > 0 {
		cHintsKeyData := make([]unsafe.Pointer, len(hints))
		cHintsValueData := make([]unsafe.Pointer, len(hints))
		var idx = 0
		for k, v := range hints {
			cHintsKeyData[idx] = foundation.NewString(string(k)).Ptr()
			cHintsValueData[idx] = objc.ExtractPtr(v)
			idx++
		}
		cHints.key_data = unsafe.Pointer(&cHintsKeyData[0])
		cHints.value_data = unsafe.Pointer(&cHintsValueData[0])
		cHints.len = C.int(len(hints))
	}
	C.C_NSImage_DrawInRect_FromRect_Operation_Fraction_RespectFlipped_Hints(n.Ptr(), *(*C.CGRect)(coregraphics.ToCGRectPointer(coregraphics.Rect(dstSpacePortionRect))), *(*C.CGRect)(coregraphics.ToCGRectPointer(coregraphics.Rect(srcSpacePortionRect))), C.uint(uint(op)), C.double(float64(requestedAlpha)), C.bool(respectContextIsFlipped), cHints)
}

func (n NSImage) DrawRepresentation_InRect(imageRep ImageRep, rect foundation.Rect) bool {
	result_ := C.C_NSImage_DrawRepresentation_InRect(n.Ptr(), objc.ExtractPtr(imageRep), *(*C.CGRect)(coregraphics.ToCGRectPointer(coregraphics.Rect(rect))))
	return bool(result_)
}

func (n NSImage) LockFocus() {
	C.C_NSImage_LockFocus(n.Ptr())
}

func (n NSImage) LockFocusFlipped(flipped bool) {
	C.C_NSImage_LockFocusFlipped(n.Ptr(), C.bool(flipped))
}

func (n NSImage) UnlockFocus() {
	C.C_NSImage_UnlockFocus(n.Ptr())
}

func (n NSImage) Recache() {
	C.C_NSImage_Recache(n.Ptr())
}

func (n NSImage) TIFFRepresentationUsingCompression_Factor(comp TIFFCompression, factor float32) []byte {
	result_ := C.C_NSImage_TIFFRepresentationUsingCompression_Factor(n.Ptr(), C.uint(uint(comp)), C.float(factor))
	return foundation.MakeData(result_).ToBytes()
}

func (n NSImage) CancelIncrementalLoad() {
	C.C_NSImage_CancelIncrementalLoad(n.Ptr())
}

func (n NSImage) HitTestRect_WithImageDestinationRect_Context_Hints_Flipped(testRectDestSpace foundation.Rect, imageRectDestSpace foundation.Rect, context GraphicsContext, hints map[ImageHintKey]objc.Object, flipped bool) bool {
	var cHints C.Dictionary
	if len(hints) > 0 {
		cHintsKeyData := make([]unsafe.Pointer, len(hints))
		cHintsValueData := make([]unsafe.Pointer, len(hints))
		var idx = 0
		for k, v := range hints {
			cHintsKeyData[idx] = foundation.NewString(string(k)).Ptr()
			cHintsValueData[idx] = objc.ExtractPtr(v)
			idx++
		}
		cHints.key_data = unsafe.Pointer(&cHintsKeyData[0])
		cHints.value_data = unsafe.Pointer(&cHintsValueData[0])
		cHints.len = C.int(len(hints))
	}
	result_ := C.C_NSImage_HitTestRect_WithImageDestinationRect_Context_Hints_Flipped(n.Ptr(), *(*C.CGRect)(coregraphics.ToCGRectPointer(coregraphics.Rect(testRectDestSpace))), *(*C.CGRect)(coregraphics.ToCGRectPointer(coregraphics.Rect(imageRectDestSpace))), objc.ExtractPtr(context), cHints, C.bool(flipped))
	return bool(result_)
}

func (n NSImage) LayerContentsForContentsScale(layerContentsScale coregraphics.Float) objc.Object {
	result_ := C.C_NSImage_LayerContentsForContentsScale(n.Ptr(), C.double(float64(layerContentsScale)))
	return objc.MakeObject(result_)
}

func (n NSImage) RecommendedLayerContentsScale(preferredContentsScale coregraphics.Float) coregraphics.Float {
	result_ := C.C_NSImage_RecommendedLayerContentsScale(n.Ptr(), C.double(float64(preferredContentsScale)))
	return coregraphics.Float(float64(result_))
}

func (n NSImage) Delegate() objc.Object {
	result_ := C.C_NSImage_Delegate(n.Ptr())
	return objc.MakeObject(result_)
}

func (n NSImage) SetDelegate(value objc.Object) {
	C.C_NSImage_SetDelegate(n.Ptr(), objc.ExtractPtr(value))
}

func (n NSImage) Size() foundation.Size {
	result_ := C.C_NSImage_Size(n.Ptr())
	return foundation.Size(coregraphics.FromCGSizePointer(unsafe.Pointer(&result_)))
}

func (n NSImage) SetSize(value foundation.Size) {
	C.C_NSImage_SetSize(n.Ptr(), *(*C.CGSize)(coregraphics.ToCGSizePointer(coregraphics.Size(value))))
}

func (n NSImage) IsTemplate() bool {
	result_ := C.C_NSImage_IsTemplate(n.Ptr())
	return bool(result_)
}

func (n NSImage) SetTemplate(value bool) {
	C.C_NSImage_SetTemplate(n.Ptr(), C.bool(value))
}

func ImageTypes() []string {
	result_ := C.C_NSImage_ImageTypes()
	if result_.len > 0 {
		defer C.free(result_.data)
	}
	result_Slice := unsafe.Slice((*unsafe.Pointer)(result_.data), int(result_.len))
	var goResult_ = make([]string, len(result_Slice))
	for idx, r := range result_Slice {
		goResult_[idx] = foundation.MakeString(r).String()
	}
	return goResult_
}

func ImageUnfilteredTypes() []string {
	result_ := C.C_NSImage_ImageUnfilteredTypes()
	if result_.len > 0 {
		defer C.free(result_.data)
	}
	result_Slice := unsafe.Slice((*unsafe.Pointer)(result_.data), int(result_.len))
	var goResult_ = make([]string, len(result_Slice))
	for idx, r := range result_Slice {
		goResult_[idx] = foundation.MakeString(r).String()
	}
	return goResult_
}

func (n NSImage) Representations() []ImageRep {
	result_ := C.C_NSImage_Representations(n.Ptr())
	if result_.len > 0 {
		defer C.free(result_.data)
	}
	result_Slice := unsafe.Slice((*unsafe.Pointer)(result_.data), int(result_.len))
	var goResult_ = make([]ImageRep, len(result_Slice))
	for idx, r := range result_Slice {
		goResult_[idx] = MakeImageRep(r)
	}
	return goResult_
}

func (n NSImage) PrefersColorMatch() bool {
	result_ := C.C_NSImage_PrefersColorMatch(n.Ptr())
	return bool(result_)
}

func (n NSImage) SetPrefersColorMatch(value bool) {
	C.C_NSImage_SetPrefersColorMatch(n.Ptr(), C.bool(value))
}

func (n NSImage) UsesEPSOnResolutionMismatch() bool {
	result_ := C.C_NSImage_UsesEPSOnResolutionMismatch(n.Ptr())
	return bool(result_)
}

func (n NSImage) SetUsesEPSOnResolutionMismatch(value bool) {
	C.C_NSImage_SetUsesEPSOnResolutionMismatch(n.Ptr(), C.bool(value))
}

func (n NSImage) MatchesOnMultipleResolution() bool {
	result_ := C.C_NSImage_MatchesOnMultipleResolution(n.Ptr())
	return bool(result_)
}

func (n NSImage) SetMatchesOnMultipleResolution(value bool) {
	C.C_NSImage_SetMatchesOnMultipleResolution(n.Ptr(), C.bool(value))
}

func (n NSImage) IsValid() bool {
	result_ := C.C_NSImage_IsValid(n.Ptr())
	return bool(result_)
}

func (n NSImage) BackgroundColor() Color {
	result_ := C.C_NSImage_BackgroundColor(n.Ptr())
	return MakeColor(result_)
}

func (n NSImage) SetBackgroundColor(value Color) {
	C.C_NSImage_SetBackgroundColor(n.Ptr(), objc.ExtractPtr(value))
}

func (n NSImage) CapInsets() foundation.EdgeInsets {
	result_ := C.C_NSImage_CapInsets(n.Ptr())
	return foundation.FromNSEdgeInsetsPointer(unsafe.Pointer(&result_))
}

func (n NSImage) SetCapInsets(value foundation.EdgeInsets) {
	C.C_NSImage_SetCapInsets(n.Ptr(), *(*C.NSEdgeInsets)(foundation.ToNSEdgeInsetsPointer(value)))
}

func (n NSImage) ResizingMode() ImageResizingMode {
	result_ := C.C_NSImage_ResizingMode(n.Ptr())
	return ImageResizingMode(int(result_))
}

func (n NSImage) SetResizingMode(value ImageResizingMode) {
	C.C_NSImage_SetResizingMode(n.Ptr(), C.int(int(value)))
}

func (n NSImage) AlignmentRect() foundation.Rect {
	result_ := C.C_NSImage_AlignmentRect(n.Ptr())
	return foundation.Rect(coregraphics.FromCGRectPointer(unsafe.Pointer(&result_)))
}

func (n NSImage) SetAlignmentRect(value foundation.Rect) {
	C.C_NSImage_SetAlignmentRect(n.Ptr(), *(*C.CGRect)(coregraphics.ToCGRectPointer(coregraphics.Rect(value))))
}

func (n NSImage) CacheMode() ImageCacheMode {
	result_ := C.C_NSImage_CacheMode(n.Ptr())
	return ImageCacheMode(uint(result_))
}

func (n NSImage) SetCacheMode(value ImageCacheMode) {
	C.C_NSImage_SetCacheMode(n.Ptr(), C.uint(uint(value)))
}

func (n NSImage) TIFFRepresentation() []byte {
	result_ := C.C_NSImage_TIFFRepresentation(n.Ptr())
	return foundation.MakeData(result_).ToBytes()
}

func (n NSImage) AccessibilityDescription() string {
	result_ := C.C_NSImage_AccessibilityDescription(n.Ptr())
	return foundation.MakeString(result_).String()
}

func (n NSImage) SetAccessibilityDescription(value string) {
	C.C_NSImage_SetAccessibilityDescription(n.Ptr(), foundation.NewString(value).Ptr())
}

func (n NSImage) MatchesOnlyOnBestFittingAxis() bool {
	result_ := C.C_NSImage_MatchesOnlyOnBestFittingAxis(n.Ptr())
	return bool(result_)
}

func (n NSImage) SetMatchesOnlyOnBestFittingAxis(value bool) {
	C.C_NSImage_SetMatchesOnlyOnBestFittingAxis(n.Ptr(), C.bool(value))
}
