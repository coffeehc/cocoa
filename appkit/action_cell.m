#import <Appkit/Appkit.h>
#import "action_cell.h"

void* C_ActionCell_Alloc() {
    return [NSActionCell alloc];
}

void* C_NSActionCell_InitImageCell(void* ptr, void* image) {
    NSActionCell* nSActionCell = (NSActionCell*)ptr;
    NSActionCell* result_ = [nSActionCell initImageCell:(NSImage*)image];
    return result_;
}

void* C_NSActionCell_InitTextCell(void* ptr, void* _string) {
    NSActionCell* nSActionCell = (NSActionCell*)ptr;
    NSActionCell* result_ = [nSActionCell initTextCell:(NSString*)_string];
    return result_;
}

void* C_NSActionCell_Init(void* ptr) {
    NSActionCell* nSActionCell = (NSActionCell*)ptr;
    NSActionCell* result_ = [nSActionCell init];
    return result_;
}

void* C_NSActionCell_InitWithCoder(void* ptr, void* coder) {
    NSActionCell* nSActionCell = (NSActionCell*)ptr;
    NSActionCell* result_ = [nSActionCell initWithCoder:(NSCoder*)coder];
    return result_;
}

void* C_NSActionCell_AllocActionCell() {
    NSActionCell* result_ = [NSActionCell alloc];
    return result_;
}

void* C_NSActionCell_NewActionCell() {
    NSActionCell* result_ = [NSActionCell new];
    return result_;
}

void* C_NSActionCell_Autorelease(void* ptr) {
    NSActionCell* nSActionCell = (NSActionCell*)ptr;
    NSActionCell* result_ = [nSActionCell autorelease];
    return result_;
}

void* C_NSActionCell_Retain(void* ptr) {
    NSActionCell* nSActionCell = (NSActionCell*)ptr;
    NSActionCell* result_ = [nSActionCell retain];
    return result_;
}
