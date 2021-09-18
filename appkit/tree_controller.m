#import "tree_controller.h"
#import <Foundation/NSDictionary.h>
#import <Foundation/NSArray.h>
#import <AppKit/NSTreeController.h>

void* C_TreeController_Alloc() {
    return [NSTreeController alloc];
}

void* C_NSTreeController_InitWithContent(void* ptr, void* content) {
    NSTreeController* nSTreeController = (NSTreeController*)ptr;
    NSTreeController* result_ = [nSTreeController initWithContent:(id)content];
    return result_;
}

void* C_NSTreeController_InitWithCoder(void* ptr, void* coder) {
    NSTreeController* nSTreeController = (NSTreeController*)ptr;
    NSTreeController* result_ = [nSTreeController initWithCoder:(NSCoder*)coder];
    return result_;
}

void* C_NSTreeController_Init(void* ptr) {
    NSTreeController* nSTreeController = (NSTreeController*)ptr;
    NSTreeController* result_ = [nSTreeController init];
    return result_;
}

void* C_NSTreeController_AllocTreeController() {
    NSTreeController* result_ = [NSTreeController alloc];
    return result_;
}

void* C_NSTreeController_NewTreeController() {
    NSTreeController* result_ = [NSTreeController new];
    return result_;
}

void* C_NSTreeController_Autorelease(void* ptr) {
    NSTreeController* nSTreeController = (NSTreeController*)ptr;
    NSTreeController* result_ = [nSTreeController autorelease];
    return result_;
}

void* C_NSTreeController_Retain(void* ptr) {
    NSTreeController* nSTreeController = (NSTreeController*)ptr;
    NSTreeController* result_ = [nSTreeController retain];
    return result_;
}

void C_NSTreeController_RearrangeObjects(void* ptr) {
    NSTreeController* nSTreeController = (NSTreeController*)ptr;
    [nSTreeController rearrangeObjects];
}

bool C_NSTreeController_SetSelectionIndexPath(void* ptr, void* indexPath) {
    NSTreeController* nSTreeController = (NSTreeController*)ptr;
    BOOL result_ = [nSTreeController setSelectionIndexPath:(NSIndexPath*)indexPath];
    return result_;
}

bool C_NSTreeController_SetSelectionIndexPaths(void* ptr, Array indexPaths) {
    NSTreeController* nSTreeController = (NSTreeController*)ptr;
    NSMutableArray* objcIndexPaths = [NSMutableArray arrayWithCapacity:indexPaths.len];
    if (indexPaths.len > 0) {
    	void** indexPathsData = (void**)indexPaths.data;
    	for (int i = 0; i < indexPaths.len; i++) {
    		void* p = indexPathsData[i];
    		[objcIndexPaths addObject:(NSIndexPath*)p];
    	}
    }
    BOOL result_ = [nSTreeController setSelectionIndexPaths:objcIndexPaths];
    return result_;
}

bool C_NSTreeController_AddSelectionIndexPaths(void* ptr, Array indexPaths) {
    NSTreeController* nSTreeController = (NSTreeController*)ptr;
    NSMutableArray* objcIndexPaths = [NSMutableArray arrayWithCapacity:indexPaths.len];
    if (indexPaths.len > 0) {
    	void** indexPathsData = (void**)indexPaths.data;
    	for (int i = 0; i < indexPaths.len; i++) {
    		void* p = indexPathsData[i];
    		[objcIndexPaths addObject:(NSIndexPath*)p];
    	}
    }
    BOOL result_ = [nSTreeController addSelectionIndexPaths:objcIndexPaths];
    return result_;
}

bool C_NSTreeController_RemoveSelectionIndexPaths(void* ptr, Array indexPaths) {
    NSTreeController* nSTreeController = (NSTreeController*)ptr;
    NSMutableArray* objcIndexPaths = [NSMutableArray arrayWithCapacity:indexPaths.len];
    if (indexPaths.len > 0) {
    	void** indexPathsData = (void**)indexPaths.data;
    	for (int i = 0; i < indexPaths.len; i++) {
    		void* p = indexPathsData[i];
    		[objcIndexPaths addObject:(NSIndexPath*)p];
    	}
    }
    BOOL result_ = [nSTreeController removeSelectionIndexPaths:objcIndexPaths];
    return result_;
}

void C_NSTreeController_AddChild(void* ptr, void* sender) {
    NSTreeController* nSTreeController = (NSTreeController*)ptr;
    [nSTreeController addChild:(id)sender];
}

void C_NSTreeController_Insert(void* ptr, void* sender) {
    NSTreeController* nSTreeController = (NSTreeController*)ptr;
    [nSTreeController insert:(id)sender];
}

void C_NSTreeController_InsertChild(void* ptr, void* sender) {
    NSTreeController* nSTreeController = (NSTreeController*)ptr;
    [nSTreeController insertChild:(id)sender];
}

void C_NSTreeController_InsertObject_AtArrangedObjectIndexPath(void* ptr, void* object, void* indexPath) {
    NSTreeController* nSTreeController = (NSTreeController*)ptr;
    [nSTreeController insertObject:(id)object atArrangedObjectIndexPath:(NSIndexPath*)indexPath];
}

void C_NSTreeController_InsertObjects_AtArrangedObjectIndexPaths(void* ptr, Array objects, Array indexPaths) {
    NSTreeController* nSTreeController = (NSTreeController*)ptr;
    NSMutableArray* objcObjects = [NSMutableArray arrayWithCapacity:objects.len];
    if (objects.len > 0) {
    	void** objectsData = (void**)objects.data;
    	for (int i = 0; i < objects.len; i++) {
    		void* p = objectsData[i];
    		[objcObjects addObject:(NSObject*)p];
    	}
    }
    NSMutableArray* objcIndexPaths = [NSMutableArray arrayWithCapacity:indexPaths.len];
    if (indexPaths.len > 0) {
    	void** indexPathsData = (void**)indexPaths.data;
    	for (int i = 0; i < indexPaths.len; i++) {
    		void* p = indexPathsData[i];
    		[objcIndexPaths addObject:(NSIndexPath*)p];
    	}
    }
    [nSTreeController insertObjects:objcObjects atArrangedObjectIndexPaths:objcIndexPaths];
}

void C_NSTreeController_RemoveObjectAtArrangedObjectIndexPath(void* ptr, void* indexPath) {
    NSTreeController* nSTreeController = (NSTreeController*)ptr;
    [nSTreeController removeObjectAtArrangedObjectIndexPath:(NSIndexPath*)indexPath];
}

void C_NSTreeController_RemoveObjectsAtArrangedObjectIndexPaths(void* ptr, Array indexPaths) {
    NSTreeController* nSTreeController = (NSTreeController*)ptr;
    NSMutableArray* objcIndexPaths = [NSMutableArray arrayWithCapacity:indexPaths.len];
    if (indexPaths.len > 0) {
    	void** indexPathsData = (void**)indexPaths.data;
    	for (int i = 0; i < indexPaths.len; i++) {
    		void* p = indexPathsData[i];
    		[objcIndexPaths addObject:(NSIndexPath*)p];
    	}
    }
    [nSTreeController removeObjectsAtArrangedObjectIndexPaths:objcIndexPaths];
}

void C_NSTreeController_MoveNode_ToIndexPath(void* ptr, void* node, void* indexPath) {
    NSTreeController* nSTreeController = (NSTreeController*)ptr;
    [nSTreeController moveNode:(NSTreeNode*)node toIndexPath:(NSIndexPath*)indexPath];
}

void C_NSTreeController_MoveNodes_ToIndexPath(void* ptr, Array nodes, void* startingIndexPath) {
    NSTreeController* nSTreeController = (NSTreeController*)ptr;
    NSMutableArray* objcNodes = [NSMutableArray arrayWithCapacity:nodes.len];
    if (nodes.len > 0) {
    	void** nodesData = (void**)nodes.data;
    	for (int i = 0; i < nodes.len; i++) {
    		void* p = nodesData[i];
    		[objcNodes addObject:(NSTreeNode*)p];
    	}
    }
    [nSTreeController moveNodes:objcNodes toIndexPath:(NSIndexPath*)startingIndexPath];
}

void* C_NSTreeController_ChildrenKeyPathForNode(void* ptr, void* node) {
    NSTreeController* nSTreeController = (NSTreeController*)ptr;
    NSString* result_ = [nSTreeController childrenKeyPathForNode:(NSTreeNode*)node];
    return result_;
}

void* C_NSTreeController_CountKeyPathForNode(void* ptr, void* node) {
    NSTreeController* nSTreeController = (NSTreeController*)ptr;
    NSString* result_ = [nSTreeController countKeyPathForNode:(NSTreeNode*)node];
    return result_;
}

void* C_NSTreeController_LeafKeyPathForNode(void* ptr, void* node) {
    NSTreeController* nSTreeController = (NSTreeController*)ptr;
    NSString* result_ = [nSTreeController leafKeyPathForNode:(NSTreeNode*)node];
    return result_;
}

Array C_NSTreeController_SortDescriptors(void* ptr) {
    NSTreeController* nSTreeController = (NSTreeController*)ptr;
    NSArray* result_ = [nSTreeController sortDescriptors];
    Array result_Array;
    int result_count = [result_ count];
    if (result_count > 0) {
    	void** result_Data = malloc(result_count * sizeof(void*));
    	for (int i = 0; i < result_count; i++) {
    		 void* p = [result_ objectAtIndex:i];
    		 result_Data[i] = p;
    	}
    	result_Array.data = result_Data;
    	result_Array.len = result_count;
    }
    return result_Array;
}

void C_NSTreeController_SetSortDescriptors(void* ptr, Array value) {
    NSTreeController* nSTreeController = (NSTreeController*)ptr;
    NSMutableArray* objcValue = [NSMutableArray arrayWithCapacity:value.len];
    if (value.len > 0) {
    	void** valueData = (void**)value.data;
    	for (int i = 0; i < value.len; i++) {
    		void* p = valueData[i];
    		[objcValue addObject:(NSSortDescriptor*)p];
    	}
    }
    [nSTreeController setSortDescriptors:objcValue];
}

void* C_NSTreeController_ArrangedObjects(void* ptr) {
    NSTreeController* nSTreeController = (NSTreeController*)ptr;
    NSTreeNode* result_ = [nSTreeController arrangedObjects];
    return result_;
}

void* C_NSTreeController_SelectionIndexPath(void* ptr) {
    NSTreeController* nSTreeController = (NSTreeController*)ptr;
    NSIndexPath* result_ = [nSTreeController selectionIndexPath];
    return result_;
}

Array C_NSTreeController_SelectionIndexPaths(void* ptr) {
    NSTreeController* nSTreeController = (NSTreeController*)ptr;
    NSArray* result_ = [nSTreeController selectionIndexPaths];
    Array result_Array;
    int result_count = [result_ count];
    if (result_count > 0) {
    	void** result_Data = malloc(result_count * sizeof(void*));
    	for (int i = 0; i < result_count; i++) {
    		 void* p = [result_ objectAtIndex:i];
    		 result_Data[i] = p;
    	}
    	result_Array.data = result_Data;
    	result_Array.len = result_count;
    }
    return result_Array;
}

Array C_NSTreeController_SelectedNodes(void* ptr) {
    NSTreeController* nSTreeController = (NSTreeController*)ptr;
    NSArray* result_ = [nSTreeController selectedNodes];
    Array result_Array;
    int result_count = [result_ count];
    if (result_count > 0) {
    	void** result_Data = malloc(result_count * sizeof(void*));
    	for (int i = 0; i < result_count; i++) {
    		 void* p = [result_ objectAtIndex:i];
    		 result_Data[i] = p;
    	}
    	result_Array.data = result_Data;
    	result_Array.len = result_count;
    }
    return result_Array;
}

bool C_NSTreeController_SelectsInsertedObjects(void* ptr) {
    NSTreeController* nSTreeController = (NSTreeController*)ptr;
    BOOL result_ = [nSTreeController selectsInsertedObjects];
    return result_;
}

void C_NSTreeController_SetSelectsInsertedObjects(void* ptr, bool value) {
    NSTreeController* nSTreeController = (NSTreeController*)ptr;
    [nSTreeController setSelectsInsertedObjects:value];
}

bool C_NSTreeController_AvoidsEmptySelection(void* ptr) {
    NSTreeController* nSTreeController = (NSTreeController*)ptr;
    BOOL result_ = [nSTreeController avoidsEmptySelection];
    return result_;
}

void C_NSTreeController_SetAvoidsEmptySelection(void* ptr, bool value) {
    NSTreeController* nSTreeController = (NSTreeController*)ptr;
    [nSTreeController setAvoidsEmptySelection:value];
}

bool C_NSTreeController_PreservesSelection(void* ptr) {
    NSTreeController* nSTreeController = (NSTreeController*)ptr;
    BOOL result_ = [nSTreeController preservesSelection];
    return result_;
}

void C_NSTreeController_SetPreservesSelection(void* ptr, bool value) {
    NSTreeController* nSTreeController = (NSTreeController*)ptr;
    [nSTreeController setPreservesSelection:value];
}

bool C_NSTreeController_AlwaysUsesMultipleValuesMarker(void* ptr) {
    NSTreeController* nSTreeController = (NSTreeController*)ptr;
    BOOL result_ = [nSTreeController alwaysUsesMultipleValuesMarker];
    return result_;
}

void C_NSTreeController_SetAlwaysUsesMultipleValuesMarker(void* ptr, bool value) {
    NSTreeController* nSTreeController = (NSTreeController*)ptr;
    [nSTreeController setAlwaysUsesMultipleValuesMarker:value];
}

bool C_NSTreeController_CanAddChild(void* ptr) {
    NSTreeController* nSTreeController = (NSTreeController*)ptr;
    BOOL result_ = [nSTreeController canAddChild];
    return result_;
}

bool C_NSTreeController_CanInsert(void* ptr) {
    NSTreeController* nSTreeController = (NSTreeController*)ptr;
    BOOL result_ = [nSTreeController canInsert];
    return result_;
}

bool C_NSTreeController_CanInsertChild(void* ptr) {
    NSTreeController* nSTreeController = (NSTreeController*)ptr;
    BOOL result_ = [nSTreeController canInsertChild];
    return result_;
}

void* C_NSTreeController_ChildrenKeyPath(void* ptr) {
    NSTreeController* nSTreeController = (NSTreeController*)ptr;
    NSString* result_ = [nSTreeController childrenKeyPath];
    return result_;
}

void C_NSTreeController_SetChildrenKeyPath(void* ptr, void* value) {
    NSTreeController* nSTreeController = (NSTreeController*)ptr;
    [nSTreeController setChildrenKeyPath:(NSString*)value];
}

void* C_NSTreeController_CountKeyPath(void* ptr) {
    NSTreeController* nSTreeController = (NSTreeController*)ptr;
    NSString* result_ = [nSTreeController countKeyPath];
    return result_;
}

void C_NSTreeController_SetCountKeyPath(void* ptr, void* value) {
    NSTreeController* nSTreeController = (NSTreeController*)ptr;
    [nSTreeController setCountKeyPath:(NSString*)value];
}

void* C_NSTreeController_LeafKeyPath(void* ptr) {
    NSTreeController* nSTreeController = (NSTreeController*)ptr;
    NSString* result_ = [nSTreeController leafKeyPath];
    return result_;
}

void C_NSTreeController_SetLeafKeyPath(void* ptr, void* value) {
    NSTreeController* nSTreeController = (NSTreeController*)ptr;
    [nSTreeController setLeafKeyPath:(NSString*)value];
}
