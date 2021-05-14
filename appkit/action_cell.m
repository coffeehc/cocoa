#import <Appkit/Appkit.h>
#import "action_cell.h"

void* C_ActionCell_Alloc() {
	return [NSActionCell alloc];
}

void* C_NSActionCell_InitImageCell(void* ptr, void* image) {
	NSActionCell* nSActionCell = (NSActionCell*)ptr;
	NSActionCell* result = [nSActionCell initImageCell:(NSImage*)image];
	return result;
}

void* C_NSActionCell_InitTextCell(void* ptr, void* _string) {
	NSActionCell* nSActionCell = (NSActionCell*)ptr;
	NSActionCell* result = [nSActionCell initTextCell:(NSString*)_string];
	return result;
}

void* C_NSActionCell_Init(void* ptr) {
	NSActionCell* nSActionCell = (NSActionCell*)ptr;
	NSActionCell* result = [nSActionCell init];
	return result;
}

void* C_NSActionCell_InitWithCoder(void* ptr, void* coder) {
	NSActionCell* nSActionCell = (NSActionCell*)ptr;
	NSActionCell* result = [nSActionCell initWithCoder:(NSCoder*)coder];
	return result;
}
