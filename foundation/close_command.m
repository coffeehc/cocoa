#import <Foundation/Foundation.h>
#import "close_command.h"

void* C_CloseCommand_Alloc() {
	return [NSCloseCommand alloc];
}

void* C_NSCloseCommand_InitWithCommandDescription(void* ptr, void* commandDef) {
	NSCloseCommand* nSCloseCommand = (NSCloseCommand*)ptr;
	NSCloseCommand* result = [nSCloseCommand initWithCommandDescription:(NSScriptCommandDescription*)commandDef];
	return result;
}

void* C_NSCloseCommand_InitWithCoder(void* ptr, void* inCoder) {
	NSCloseCommand* nSCloseCommand = (NSCloseCommand*)ptr;
	NSCloseCommand* result = [nSCloseCommand initWithCoder:(NSCoder*)inCoder];
	return result;
}

unsigned int C_NSCloseCommand_SaveOptions(void* ptr) {
	NSCloseCommand* nSCloseCommand = (NSCloseCommand*)ptr;
	unsigned int result = [nSCloseCommand saveOptions];
	return result;
}
