#import <Foundation/NSGeometry.h>
#import <Foundation/NSRange.h>
#import <stdbool.h>
#import <stdint.h>
#import <stdlib.h>
#import <utils.h>

void* C_TreeController_Alloc();

void* C_NSTreeController_InitWithContent(void* ptr, void* content);
void* C_NSTreeController_InitWithCoder(void* ptr, void* coder);
void* C_NSTreeController_Init(void* ptr);
void* C_NSTreeController_AllocTreeController();
void* C_NSTreeController_NewTreeController();
void* C_NSTreeController_Autorelease(void* ptr);
void* C_NSTreeController_Retain(void* ptr);
void C_NSTreeController_RearrangeObjects(void* ptr);
bool C_NSTreeController_SetSelectionIndexPath(void* ptr, void* indexPath);
bool C_NSTreeController_SetSelectionIndexPaths(void* ptr, Array indexPaths);
bool C_NSTreeController_AddSelectionIndexPaths(void* ptr, Array indexPaths);
bool C_NSTreeController_RemoveSelectionIndexPaths(void* ptr, Array indexPaths);
void C_NSTreeController_AddChild(void* ptr, void* sender);
void C_NSTreeController_Insert(void* ptr, void* sender);
void C_NSTreeController_InsertChild(void* ptr, void* sender);
void C_NSTreeController_InsertObject_AtArrangedObjectIndexPath(void* ptr, void* object, void* indexPath);
void C_NSTreeController_InsertObjects_AtArrangedObjectIndexPaths(void* ptr, Array objects, Array indexPaths);
void C_NSTreeController_RemoveObjectAtArrangedObjectIndexPath(void* ptr, void* indexPath);
void C_NSTreeController_RemoveObjectsAtArrangedObjectIndexPaths(void* ptr, Array indexPaths);
void C_NSTreeController_MoveNode_ToIndexPath(void* ptr, void* node, void* indexPath);
void C_NSTreeController_MoveNodes_ToIndexPath(void* ptr, Array nodes, void* startingIndexPath);
void* C_NSTreeController_ChildrenKeyPathForNode(void* ptr, void* node);
void* C_NSTreeController_CountKeyPathForNode(void* ptr, void* node);
void* C_NSTreeController_LeafKeyPathForNode(void* ptr, void* node);
Array C_NSTreeController_SortDescriptors(void* ptr);
void C_NSTreeController_SetSortDescriptors(void* ptr, Array value);
void* C_NSTreeController_ArrangedObjects(void* ptr);
void* C_NSTreeController_SelectionIndexPath(void* ptr);
Array C_NSTreeController_SelectionIndexPaths(void* ptr);
Array C_NSTreeController_SelectedNodes(void* ptr);
bool C_NSTreeController_SelectsInsertedObjects(void* ptr);
void C_NSTreeController_SetSelectsInsertedObjects(void* ptr, bool value);
bool C_NSTreeController_AvoidsEmptySelection(void* ptr);
void C_NSTreeController_SetAvoidsEmptySelection(void* ptr, bool value);
bool C_NSTreeController_PreservesSelection(void* ptr);
void C_NSTreeController_SetPreservesSelection(void* ptr, bool value);
bool C_NSTreeController_AlwaysUsesMultipleValuesMarker(void* ptr);
void C_NSTreeController_SetAlwaysUsesMultipleValuesMarker(void* ptr, bool value);
bool C_NSTreeController_CanAddChild(void* ptr);
bool C_NSTreeController_CanInsert(void* ptr);
bool C_NSTreeController_CanInsertChild(void* ptr);
void* C_NSTreeController_ChildrenKeyPath(void* ptr);
void C_NSTreeController_SetChildrenKeyPath(void* ptr, void* value);
void* C_NSTreeController_CountKeyPath(void* ptr);
void C_NSTreeController_SetCountKeyPath(void* ptr, void* value);
void* C_NSTreeController_LeafKeyPath(void* ptr);
void C_NSTreeController_SetLeafKeyPath(void* ptr, void* value);
