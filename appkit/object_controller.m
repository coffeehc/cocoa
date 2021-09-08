#import <Appkit/Appkit.h>
#import "object_controller.h"

void* C_ObjectController_Alloc() {
    return [NSObjectController alloc];
}

void* C_NSObjectController_InitWithContent(void* ptr, void* content) {
    NSObjectController* nSObjectController = (NSObjectController*)ptr;
    NSObjectController* result_ = [nSObjectController initWithContent:(id)content];
    return result_;
}

void* C_NSObjectController_InitWithCoder(void* ptr, void* coder) {
    NSObjectController* nSObjectController = (NSObjectController*)ptr;
    NSObjectController* result_ = [nSObjectController initWithCoder:(NSCoder*)coder];
    return result_;
}

void* C_NSObjectController_Init(void* ptr) {
    NSObjectController* nSObjectController = (NSObjectController*)ptr;
    NSObjectController* result_ = [nSObjectController init];
    return result_;
}

void* C_NSObjectController_AllocObjectController() {
    NSObjectController* result_ = [NSObjectController alloc];
    return result_;
}

void* C_NSObjectController_NewObjectController() {
    NSObjectController* result_ = [NSObjectController new];
    return result_;
}

void* C_NSObjectController_Autorelease(void* ptr) {
    NSObjectController* nSObjectController = (NSObjectController*)ptr;
    NSObjectController* result_ = [nSObjectController autorelease];
    return result_;
}

void* C_NSObjectController_Retain(void* ptr) {
    NSObjectController* nSObjectController = (NSObjectController*)ptr;
    NSObjectController* result_ = [nSObjectController retain];
    return result_;
}

void C_NSObjectController_PrepareContent(void* ptr) {
    NSObjectController* nSObjectController = (NSObjectController*)ptr;
    [nSObjectController prepareContent];
}

void* C_NSObjectController_NewObject(void* ptr) {
    NSObjectController* nSObjectController = (NSObjectController*)ptr;
    id result_ = [nSObjectController newObject];
    return result_;
}

void C_NSObjectController_AddObject(void* ptr, void* object) {
    NSObjectController* nSObjectController = (NSObjectController*)ptr;
    [nSObjectController addObject:(id)object];
}

void C_NSObjectController_RemoveObject(void* ptr, void* object) {
    NSObjectController* nSObjectController = (NSObjectController*)ptr;
    [nSObjectController removeObject:(id)object];
}

void C_NSObjectController_Add(void* ptr, void* sender) {
    NSObjectController* nSObjectController = (NSObjectController*)ptr;
    [nSObjectController add:(id)sender];
}

void C_NSObjectController_Remove(void* ptr, void* sender) {
    NSObjectController* nSObjectController = (NSObjectController*)ptr;
    [nSObjectController remove:(id)sender];
}

void C_NSObjectController_Fetch(void* ptr, void* sender) {
    NSObjectController* nSObjectController = (NSObjectController*)ptr;
    [nSObjectController fetch:(id)sender];
}

bool C_NSObjectController_ValidateUserInterfaceItem(void* ptr, void* item) {
    NSObjectController* nSObjectController = (NSObjectController*)ptr;
    BOOL result_ = [nSObjectController validateUserInterfaceItem:(id)item];
    return result_;
}

void* C_NSObjectController_Content(void* ptr) {
    NSObjectController* nSObjectController = (NSObjectController*)ptr;
    id result_ = [nSObjectController content];
    return result_;
}

void C_NSObjectController_SetContent(void* ptr, void* value) {
    NSObjectController* nSObjectController = (NSObjectController*)ptr;
    [nSObjectController setContent:(id)value];
}

bool C_NSObjectController_AutomaticallyPreparesContent(void* ptr) {
    NSObjectController* nSObjectController = (NSObjectController*)ptr;
    BOOL result_ = [nSObjectController automaticallyPreparesContent];
    return result_;
}

void C_NSObjectController_SetAutomaticallyPreparesContent(void* ptr, bool value) {
    NSObjectController* nSObjectController = (NSObjectController*)ptr;
    [nSObjectController setAutomaticallyPreparesContent:value];
}

bool C_NSObjectController_CanAdd(void* ptr) {
    NSObjectController* nSObjectController = (NSObjectController*)ptr;
    BOOL result_ = [nSObjectController canAdd];
    return result_;
}

bool C_NSObjectController_CanRemove(void* ptr) {
    NSObjectController* nSObjectController = (NSObjectController*)ptr;
    BOOL result_ = [nSObjectController canRemove];
    return result_;
}

bool C_NSObjectController_IsEditable(void* ptr) {
    NSObjectController* nSObjectController = (NSObjectController*)ptr;
    BOOL result_ = [nSObjectController isEditable];
    return result_;
}

void C_NSObjectController_SetEditable(void* ptr, bool value) {
    NSObjectController* nSObjectController = (NSObjectController*)ptr;
    [nSObjectController setEditable:value];
}

void* C_NSObjectController_EntityName(void* ptr) {
    NSObjectController* nSObjectController = (NSObjectController*)ptr;
    NSString* result_ = [nSObjectController entityName];
    return result_;
}

void C_NSObjectController_SetEntityName(void* ptr, void* value) {
    NSObjectController* nSObjectController = (NSObjectController*)ptr;
    [nSObjectController setEntityName:(NSString*)value];
}

bool C_NSObjectController_UsesLazyFetching(void* ptr) {
    NSObjectController* nSObjectController = (NSObjectController*)ptr;
    BOOL result_ = [nSObjectController usesLazyFetching];
    return result_;
}

void C_NSObjectController_SetUsesLazyFetching(void* ptr, bool value) {
    NSObjectController* nSObjectController = (NSObjectController*)ptr;
    [nSObjectController setUsesLazyFetching:value];
}

void* C_NSObjectController_FetchPredicate(void* ptr) {
    NSObjectController* nSObjectController = (NSObjectController*)ptr;
    NSPredicate* result_ = [nSObjectController fetchPredicate];
    return result_;
}

void C_NSObjectController_SetFetchPredicate(void* ptr, void* value) {
    NSObjectController* nSObjectController = (NSObjectController*)ptr;
    [nSObjectController setFetchPredicate:(NSPredicate*)value];
}

Array C_NSObjectController_SelectedObjects(void* ptr) {
    NSObjectController* nSObjectController = (NSObjectController*)ptr;
    NSArray* result_ = [nSObjectController selectedObjects];
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

void* C_NSObjectController_Selection(void* ptr) {
    NSObjectController* nSObjectController = (NSObjectController*)ptr;
    id result_ = [nSObjectController selection];
    return result_;
}
