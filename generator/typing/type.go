package typing

import (
	"github.com/hsiafan/cocoa/generator/base"
)

// Objective-c module
type Module struct {
	Name    string // module name
	Package string // go package for this module
	Header  string // objc header file
}

var Foundation = Module{"Foundation", "foundation", "Foundation/Foundation.h"}
var AppKit = Module{"AppKit", "appkit", "Appkit/Appkit.h"}
var UniformTypeIdentifiers = Module{"UniformTypeIdentifiers", "uti", "UniformTypeIdentifiers/UniformTypeIdentifiers.h"}
var WebKit = Module{"WebKit", "webkit", "WebKit/WebKit.h"}
var FileProvider = Module{"FileProvider", "fileprovider", "FileProvider/FileProvider.h"}

func FindModule(module string) Module {
	switch module {
	case "Foundation":
		fallthrough
	case "foundation":
		return Foundation
	case "AppKit":
		fallthrough
	case "appkit":
		return AppKit
	case "UniformTypeIdentifiers":
		fallthrough
	case "uniformuypeidentifiers":
		return UniformTypeIdentifiers
	case "WebKit":
		fallthrough
	case "webkit":
		return WebKit
	case "fileprovider":
		fallthrough
	case "File Provider":
		fallthrough
	case "FileProvider":
		return FileProvider
	default:
		panic("unknown module " + module)
	}
}

// #import <objc/runtime.h>
// LDFLAGS=-ObjC
var ObjCRuntime = Module{"ObjC", "objc", "objc/runtime.h"}

// interface for all type
type Type interface {
	// Go type name
	GoName(currentModule *Module) string
	// C type name
	CName() string
	// Cgo type name
	CgoName() string
	// Objective-c type name
	ObjcName() string

	// Go type to cgo type, used by Go call c via cgo
	GoToCgoCode(currentModule *Module, param string) (convertCodes []string, result string)

	// C type to objc type, used by c code call objc method
	CToObjcCode(param string) (convertCodes []string, result string)

	// objc type to c type
	ObjcToCCode(param string) (convertCodes []string, result string)

	// cgo type to go type
	CgoToGoCode(currentModule *Module, param string) (convertCodes []string, result string)

	// go imports for this type
	GoImports() base.StringSet
}
