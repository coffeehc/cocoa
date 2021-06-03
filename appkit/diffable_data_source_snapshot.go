package appkit

// #include "diffable_data_source_snapshot.h"
import "C"
import (
	"github.com/hsiafan/cocoa/objc"
	"unsafe"
)

type DiffableDataSourceSnapshot interface {
	objc.Object
	AppendSectionsWithIdentifiers(sectionIdentifiers []objc.Object)
	DeleteAllItems()
	NumberOfItems() int
	NumberOfSections() int
}

type NSDiffableDataSourceSnapshot struct {
	objc.NSObject
}

func MakeDiffableDataSourceSnapshot(ptr unsafe.Pointer) NSDiffableDataSourceSnapshot {
	return NSDiffableDataSourceSnapshot{
		NSObject: objc.MakeObject(ptr),
	}
}

func AllocDiffableDataSourceSnapshot() NSDiffableDataSourceSnapshot {
	return MakeDiffableDataSourceSnapshot(C.C_DiffableDataSourceSnapshot_Alloc())
}

func (n NSDiffableDataSourceSnapshot) Init() DiffableDataSourceSnapshot {
	result_ := C.C_NSDiffableDataSourceSnapshot_Init(n.Ptr())
	return MakeDiffableDataSourceSnapshot(result_)
}

func (n NSDiffableDataSourceSnapshot) AppendSectionsWithIdentifiers(sectionIdentifiers []objc.Object) {
	cSectionIdentifiersData := make([]unsafe.Pointer, len(sectionIdentifiers))
	for idx, v := range sectionIdentifiers {
		cSectionIdentifiersData[idx] = objc.ExtractPtr(v)
	}
	cSectionIdentifiers := C.Array{data: unsafe.Pointer(&cSectionIdentifiersData[0]), len: C.int(len(sectionIdentifiers))}
	C.C_NSDiffableDataSourceSnapshot_AppendSectionsWithIdentifiers(n.Ptr(), cSectionIdentifiers)
}

func (n NSDiffableDataSourceSnapshot) DeleteAllItems() {
	C.C_NSDiffableDataSourceSnapshot_DeleteAllItems(n.Ptr())
}

func (n NSDiffableDataSourceSnapshot) NumberOfItems() int {
	result_ := C.C_NSDiffableDataSourceSnapshot_NumberOfItems(n.Ptr())
	return int(result_)
}

func (n NSDiffableDataSourceSnapshot) NumberOfSections() int {
	result_ := C.C_NSDiffableDataSourceSnapshot_NumberOfSections(n.Ptr())
	return int(result_)
}
