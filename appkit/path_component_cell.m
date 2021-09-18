#import "path_component_cell.h"
#import <Foundation/NSDictionary.h>
#import <Foundation/NSArray.h>
#import <AppKit/NSPathComponentCell.h>

void* C_PathComponentCell_Alloc() {
    return [NSPathComponentCell alloc];
}

void* C_NSPathComponentCell_InitTextCell(void* ptr, void* _string) {
    NSPathComponentCell* nSPathComponentCell = (NSPathComponentCell*)ptr;
    NSPathComponentCell* result_ = [nSPathComponentCell initTextCell:(NSString*)_string];
    return result_;
}

void* C_NSPathComponentCell_InitWithCoder(void* ptr, void* coder) {
    NSPathComponentCell* nSPathComponentCell = (NSPathComponentCell*)ptr;
    NSPathComponentCell* result_ = [nSPathComponentCell initWithCoder:(NSCoder*)coder];
    return result_;
}

void* C_NSPathComponentCell_Init(void* ptr) {
    NSPathComponentCell* nSPathComponentCell = (NSPathComponentCell*)ptr;
    NSPathComponentCell* result_ = [nSPathComponentCell init];
    return result_;
}

void* C_NSPathComponentCell_AllocPathComponentCell() {
    NSPathComponentCell* result_ = [NSPathComponentCell alloc];
    return result_;
}

void* C_NSPathComponentCell_NewPathComponentCell() {
    NSPathComponentCell* result_ = [NSPathComponentCell new];
    return result_;
}

void* C_NSPathComponentCell_Autorelease(void* ptr) {
    NSPathComponentCell* nSPathComponentCell = (NSPathComponentCell*)ptr;
    NSPathComponentCell* result_ = [nSPathComponentCell autorelease];
    return result_;
}

void* C_NSPathComponentCell_Retain(void* ptr) {
    NSPathComponentCell* nSPathComponentCell = (NSPathComponentCell*)ptr;
    NSPathComponentCell* result_ = [nSPathComponentCell retain];
    return result_;
}

void* C_NSPathComponentCell_URL(void* ptr) {
    NSPathComponentCell* nSPathComponentCell = (NSPathComponentCell*)ptr;
    NSURL* result_ = [nSPathComponentCell URL];
    return result_;
}

void C_NSPathComponentCell_SetURL(void* ptr, void* value) {
    NSPathComponentCell* nSPathComponentCell = (NSPathComponentCell*)ptr;
    [nSPathComponentCell setURL:(NSURL*)value];
}
