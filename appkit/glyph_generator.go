package appkit

// #include "glyph_generator.h"
import "C"
import (
	"github.com/hsiafan/cocoa/objc"
	"unsafe"
)

type GlyphGenerator interface {
	objc.Object
}

type NSGlyphGenerator struct {
	objc.NSObject
}

func MakeGlyphGenerator(ptr unsafe.Pointer) *NSGlyphGenerator {
	if ptr == nil {
		return nil
	}
	return &NSGlyphGenerator{
		NSObject: *objc.MakeObject(ptr),
	}
}

func AllocGlyphGenerator() *NSGlyphGenerator {
	return MakeGlyphGenerator(C.C_GlyphGenerator_Alloc())
}

func (n *NSGlyphGenerator) Init() GlyphGenerator {
	result_ := C.C_NSGlyphGenerator_Init(n.Ptr())
	return MakeGlyphGenerator(result_)
}

func SharedGlyphGenerator() GlyphGenerator {
	result_ := C.C_NSGlyphGenerator_SharedGlyphGenerator()
	return MakeGlyphGenerator(result_)
}
