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

func AllocPredicate() NSPredicate {
	return MakePredicate(C.C_Predicate_Alloc())
}

func (n NSPredicate) Init() Predicate {
	result_ := C.C_NSPredicate_Init(n.Ptr())
	return MakePredicate(result_)
}

func PredicateWithFormat_ArgumentArray(predicateFormat string, arguments []objc.Object) Predicate {
	cArgumentsData := make([]unsafe.Pointer, len(arguments))
	for idx, v := range arguments {
		cArgumentsData[idx] = objc.ExtractPtr(v)
	}
	cArguments := C.Array{data: unsafe.Pointer(&cArgumentsData[0]), len: C.int(len(arguments))}
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

func (n NSPredicate) AllowEvaluation() {
	C.C_NSPredicate_AllowEvaluation(n.Ptr())
}

func (n NSPredicate) PredicateFormat() string {
	result_ := C.C_NSPredicate_PredicateFormat(n.Ptr())
	return MakeString(result_).String()
}
