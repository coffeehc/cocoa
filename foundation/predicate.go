package foundation

// #include "predicate.h"
import "C"
import (
	"github.com/hsiafan/cocoa/objc"
	"unsafe"
)

type Predicate interface {
	objc.Object
	EvaluateWithObject(object objc.Object) bool
	EvaluateWithObject_SubstitutionVariables(object objc.Object, bindings map[string]objc.Object) bool
	AllowEvaluation()
	PredicateFormat() string
}

type NSPredicate struct {
	objc.NSObject
}

func MakePredicate(ptr unsafe.Pointer) NSPredicate {
	return NSPredicate{
		NSObject: objc.MakeObject(ptr),
	}
}

func (n NSPredicate) PredicateWithSubstitutionVariables(variables map[string]objc.Object) NSPredicate {
	var cVariables C.Dictionary
	if len(variables) > 0 {
		cVariablesKeyData := make([]unsafe.Pointer, len(variables))
		cVariablesValueData := make([]unsafe.Pointer, len(variables))
		var idx = 0
		for k, v := range variables {
			cVariablesKeyData[idx] = NewString(k).Ptr()
			cVariablesValueData[idx] = objc.ExtractPtr(v)
			idx++
		}
		cVariables.key_data = unsafe.Pointer(&cVariablesKeyData[0])
		cVariables.value_data = unsafe.Pointer(&cVariablesValueData[0])
		cVariables.len = C.int(len(variables))
	}
	result_ := C.C_NSPredicate_PredicateWithSubstitutionVariables(n.Ptr(), cVariables)
	return MakePredicate(result_)
}

func AllocPredicate() NSPredicate {
	result_ := C.C_NSPredicate_AllocPredicate()
	return MakePredicate(result_)
}

func (n NSPredicate) Init() NSPredicate {
	result_ := C.C_NSPredicate_Init(n.Ptr())
	return MakePredicate(result_)
}

func NewPredicate() NSPredicate {
	result_ := C.C_NSPredicate_NewPredicate()
	return MakePredicate(result_)
}

func (n NSPredicate) Autorelease() NSPredicate {
	result_ := C.C_NSPredicate_Autorelease(n.Ptr())
	return MakePredicate(result_)
}

func (n NSPredicate) Retain() NSPredicate {
	result_ := C.C_NSPredicate_Retain(n.Ptr())
	return MakePredicate(result_)
}

func PredicateWithFormat_ArgumentArray(predicateFormat string, arguments []objc.Object) Predicate {
	var cArguments C.Array
	if len(arguments) > 0 {
		cArgumentsData := make([]unsafe.Pointer, len(arguments))
		for idx, v := range arguments {
			cArgumentsData[idx] = objc.ExtractPtr(v)
		}
		cArguments.data = unsafe.Pointer(&cArgumentsData[0])
		cArguments.len = C.int(len(arguments))
	}
	result_ := C.C_NSPredicate_PredicateWithFormat_ArgumentArray(NewString(predicateFormat).Ptr(), cArguments)
	return MakePredicate(result_)
}

func PredicateWithValue(value bool) Predicate {
	result_ := C.C_NSPredicate_PredicateWithValue(C.bool(value))
	return MakePredicate(result_)
}

func PredicateFromMetadataQueryString(queryString string) Predicate {
	result_ := C.C_NSPredicate_PredicateFromMetadataQueryString(NewString(queryString).Ptr())
	return MakePredicate(result_)
}

func (n NSPredicate) EvaluateWithObject(object objc.Object) bool {
	result_ := C.C_NSPredicate_EvaluateWithObject(n.Ptr(), objc.ExtractPtr(object))
	return bool(result_)
}

func (n NSPredicate) EvaluateWithObject_SubstitutionVariables(object objc.Object, bindings map[string]objc.Object) bool {
	var cBindings C.Dictionary
	if len(bindings) > 0 {
		cBindingsKeyData := make([]unsafe.Pointer, len(bindings))
		cBindingsValueData := make([]unsafe.Pointer, len(bindings))
		var idx = 0
		for k, v := range bindings {
			cBindingsKeyData[idx] = NewString(k).Ptr()
			cBindingsValueData[idx] = objc.ExtractPtr(v)
			idx++
		}
		cBindings.key_data = unsafe.Pointer(&cBindingsKeyData[0])
		cBindings.value_data = unsafe.Pointer(&cBindingsValueData[0])
		cBindings.len = C.int(len(bindings))
	}
	result_ := C.C_NSPredicate_EvaluateWithObject_SubstitutionVariables(n.Ptr(), objc.ExtractPtr(object), cBindings)
	return bool(result_)
}

func (n NSPredicate) AllowEvaluation() {
	C.C_NSPredicate_AllowEvaluation(n.Ptr())
}

func (n NSPredicate) PredicateFormat() string {
	result_ := C.C_NSPredicate_PredicateFormat(n.Ptr())
	return MakeString(result_).String()
}
