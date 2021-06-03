package appkit

// #include "collection_view_diffable_data_source.h"
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
	return MakeCollectionViewDiffableDataSource(C.C_CollectionViewDiffableDataSource_Alloc())
}

func (n NSCollectionViewDiffableDataSource) Snapshot() DiffableDataSourceSnapshot {
	result_ := C.C_NSCollectionViewDiffableDataSource_Snapshot(n.Ptr())
	return MakeDiffableDataSourceSnapshot(result_)
}

func (n NSCollectionViewDiffableDataSource) ApplySnapshot_AnimatingDifferences(snapshot DiffableDataSourceSnapshot, animatingDifferences bool) {
	C.C_NSCollectionViewDiffableDataSource_ApplySnapshot_AnimatingDifferences(n.Ptr(), objc.ExtractPtr(snapshot), C.bool(animatingDifferences))
}
