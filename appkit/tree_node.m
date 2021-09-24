#import "tree_node.h"
#import <AppKit/NSTreeNode.h>
#import <Foundation/NSArray.h>
#import <Foundation/NSDictionary.h>

void* C_TreeNode_Alloc() {
    return [NSTreeNode alloc];
}

void* C_NSTreeNode_TreeNodeWithRepresentedObject(void* modelObject) {
    NSTreeNode* result_ = [NSTreeNode treeNodeWithRepresentedObject:(id)modelObject];
    return result_;
}

void* C_NSTreeNode_InitWithRepresentedObject(void* ptr, void* modelObject) {
    NSTreeNode* nSTreeNode = (NSTreeNode*)ptr;
    NSTreeNode* result_ = [nSTreeNode initWithRepresentedObject:(id)modelObject];
    return result_;
}

void* C_NSTreeNode_AllocTreeNode() {
    NSTreeNode* result_ = [NSTreeNode alloc];
    return result_;
}

void* C_NSTreeNode_Init(void* ptr) {
    NSTreeNode* nSTreeNode = (NSTreeNode*)ptr;
    NSTreeNode* result_ = [nSTreeNode init];
    return result_;
}

void* C_NSTreeNode_NewTreeNode() {
    NSTreeNode* result_ = [NSTreeNode new];
    return result_;
}

void* C_NSTreeNode_Autorelease(void* ptr) {
    NSTreeNode* nSTreeNode = (NSTreeNode*)ptr;
    NSTreeNode* result_ = [nSTreeNode autorelease];
    return result_;
}

void* C_NSTreeNode_Retain(void* ptr) {
    NSTreeNode* nSTreeNode = (NSTreeNode*)ptr;
    NSTreeNode* result_ = [nSTreeNode retain];
    return result_;
}

void* C_NSTreeNode_DescendantNodeAtIndexPath(void* ptr, void* indexPath) {
    NSTreeNode* nSTreeNode = (NSTreeNode*)ptr;
    NSTreeNode* result_ = [nSTreeNode descendantNodeAtIndexPath:(NSIndexPath*)indexPath];
    return result_;
}

void C_NSTreeNode_SortWithSortDescriptors_Recursively(void* ptr, Array sortDescriptors, bool recursively) {
    NSTreeNode* nSTreeNode = (NSTreeNode*)ptr;
    NSMutableArray* objcSortDescriptors = [NSMutableArray arrayWithCapacity:sortDescriptors.len];
    if (sortDescriptors.len > 0) {
    	void** sortDescriptorsData = (void**)sortDescriptors.data;
    	for (int i = 0; i < sortDescriptors.len; i++) {
    		void* p = sortDescriptorsData[i];
    		[objcSortDescriptors addObject:(NSSortDescriptor*)p];
    	}
    }
    [nSTreeNode sortWithSortDescriptors:objcSortDescriptors recursively:recursively];
}

void* C_NSTreeNode_RepresentedObject(void* ptr) {
    NSTreeNode* nSTreeNode = (NSTreeNode*)ptr;
    id result_ = [nSTreeNode representedObject];
    return result_;
}

void* C_NSTreeNode_IndexPath(void* ptr) {
    NSTreeNode* nSTreeNode = (NSTreeNode*)ptr;
    NSIndexPath* result_ = [nSTreeNode indexPath];
    return result_;
}

bool C_NSTreeNode_IsLeaf(void* ptr) {
    NSTreeNode* nSTreeNode = (NSTreeNode*)ptr;
    BOOL result_ = [nSTreeNode isLeaf];
    return result_;
}

Array C_NSTreeNode_ChildNodes(void* ptr) {
    NSTreeNode* nSTreeNode = (NSTreeNode*)ptr;
    NSArray* result_ = [nSTreeNode childNodes];
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

void* C_NSTreeNode_ParentNode(void* ptr) {
    NSTreeNode* nSTreeNode = (NSTreeNode*)ptr;
    NSTreeNode* result_ = [nSTreeNode parentNode];
    return result_;
}
