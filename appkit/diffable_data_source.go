package appkit

// #include "diffable_data_source.h"
import "C"
import (
	"github.com/hsiafan/cocoa/objc"
	"unsafe"
)

type CollectionViewDiffableDataSource interface {
	objc.Object
	Snapshot() DiffableDataSourceSnapshot
	ApplySnapshot_AnimatingDifferences(snapshot DiffableDataSourceSnapshot, animatingDifferences bool)
}

type NSCollectionViewDiffableDataSource struct {
	objc.NSObject
}

func MakeCollectionViewDiffableDataSource(ptr unsafe.Pointer) NSCollectionViewDiffableDataSource {
	return NSCollectionViewDiffableDataSource{
		NSObject: objc.MakeObject(ptr),
	}
}

func AllocCollectionViewDiffableDataSource() NSCollectionViewDiffableDataSource {
	result_ := C.C_NSCollectionViewDiffableDataSource_AllocCollectionViewDiffableDataSource()
	return MakeCollectionViewDiffableDataSource(result_)
}

func (n NSCollectionViewDiffableDataSource) Autorelease() NSCollectionViewDiffableDataSource {
	result_ := C.C_NSCollectionViewDiffableDataSource_Autorelease(n.Ptr())
	return MakeCollectionViewDiffableDataSource(result_)
}

func (n NSCollectionViewDiffableDataSource) Retain() NSCollectionViewDiffableDataSource {
	result_ := C.C_NSCollectionViewDiffableDataSource_Retain(n.Ptr())
	return MakeCollectionViewDiffableDataSource(result_)
}

func (n NSCollectionViewDiffableDataSource) Snapshot() DiffableDataSourceSnapshot {
	result_ := C.C_NSCollectionViewDiffableDataSource_Snapshot(n.Ptr())
	return MakeDiffableDataSourceSnapshot(result_)
}

func (n NSCollectionViewDiffableDataSource) ApplySnapshot_AnimatingDifferences(snapshot DiffableDataSourceSnapshot, animatingDifferences bool) {
	C.C_NSCollectionViewDiffableDataSource_ApplySnapshot_AnimatingDifferences(n.Ptr(), objc.ExtractPtr(snapshot), C.bool(animatingDifferences))
}

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
	result_ := C.C_NSDiffableDataSourceSnapshot_AllocDiffableDataSourceSnapshot()
	return MakeDiffableDataSourceSnapshot(result_)
}

func (n NSDiffableDataSourceSnapshot) Init() NSDiffableDataSourceSnapshot {
	result_ := C.C_NSDiffableDataSourceSnapshot_Init(n.Ptr())
	return MakeDiffableDataSourceSnapshot(result_)
}

func NewDiffableDataSourceSnapshot() NSDiffableDataSourceSnapshot {
	result_ := C.C_NSDiffableDataSourceSnapshot_NewDiffableDataSourceSnapshot()
	return MakeDiffableDataSourceSnapshot(result_)
}

func (n NSDiffableDataSourceSnapshot) Autorelease() NSDiffableDataSourceSnapshot {
	result_ := C.C_NSDiffableDataSourceSnapshot_Autorelease(n.Ptr())
	return MakeDiffableDataSourceSnapshot(result_)
}

func (n NSDiffableDataSourceSnapshot) Retain() NSDiffableDataSourceSnapshot {
	result_ := C.C_NSDiffableDataSourceSnapshot_Retain(n.Ptr())
	return MakeDiffableDataSourceSnapshot(result_)
}

func (n NSDiffableDataSourceSnapshot) AppendSectionsWithIdentifiers(sectionIdentifiers []objc.Object) {
	var cSectionIdentifiers C.Array
	if len(sectionIdentifiers) > 0 {
		cSectionIdentifiersData := make([]unsafe.Pointer, len(sectionIdentifiers))
		for idx, v := range sectionIdentifiers {
			cSectionIdentifiersData[idx] = objc.ExtractPtr(v)
		}
		cSectionIdentifiers.data = unsafe.Pointer(&cSectionIdentifiersData[0])
		cSectionIdentifiers.len = C.int(len(sectionIdentifiers))
	}
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
