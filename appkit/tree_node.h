#import <Foundation/NSGeometry.h>
#import <Foundation/NSRange.h>
#import <stdbool.h>
#import <stdint.h>
#import <stdlib.h>
#import <utils.h>

void* C_TreeNode_Alloc();

void* C_NSTreeNode_TreeNodeWithRepresentedObject(void* modelObject);
void* C_NSTreeNode_InitWithRepresentedObject(void* ptr, void* modelObject);
void* C_NSTreeNode_AllocTreeNode();
void* C_NSTreeNode_Init(void* ptr);
void* C_NSTreeNode_NewTreeNode();
void* C_NSTreeNode_Autorelease(void* ptr);
void* C_NSTreeNode_Retain(void* ptr);
void* C_NSTreeNode_DescendantNodeAtIndexPath(void* ptr, void* indexPath);
void C_NSTreeNode_SortWithSortDescriptors_Recursively(void* ptr, Array sortDescriptors, bool recursively);
void* C_NSTreeNode_RepresentedObject(void* ptr);
void* C_NSTreeNode_IndexPath(void* ptr);
bool C_NSTreeNode_IsLeaf(void* ptr);
Array C_NSTreeNode_ChildNodes(void* ptr);
void* C_NSTreeNode_ParentNode(void* ptr);
