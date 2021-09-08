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

func MakeGlyphGenerator(ptr unsafe.Pointer) NSGlyphGenerator {
	return NSGlyphGenerator{
		NSObject: objc.MakeObject(ptr),
	}
}

func AllocGlyphGenerator() NSGlyphGenerator {
	result_ := C.C_NSGlyphGenerator_AllocGlyphGenerator()
	return MakeGlyphGenerator(result_)
}

func (n NSGlyphGenerator) Init() NSGlyphGenerator {
	result_ := C.C_NSGlyphGenerator_Init(n.Ptr())
	return MakeGlyphGenerator(result_)
}

func NewGlyphGenerator() NSGlyphGenerator {
	result_ := C.C_NSGlyphGenerator_NewGlyphGenerator()
	return MakeGlyphGenerator(result_)
}

func (n NSGlyphGenerator) Autorelease() NSGlyphGenerator {
	result_ := C.C_NSGlyphGenerator_Autorelease(n.Ptr())
	return MakeGlyphGenerator(result_)
}

func (n NSGlyphGenerator) Retain() NSGlyphGenerator {
	result_ := C.C_NSGlyphGenerator_Retain(n.Ptr())
	return MakeGlyphGenerator(result_)
}

func SharedGlyphGenerator() GlyphGenerator {
	result_ := C.C_NSGlyphGenerator_SharedGlyphGenerator()
	return MakeGlyphGenerator(result_)
}
