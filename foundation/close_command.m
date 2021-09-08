#import <Foundation/Foundation.h>
#import "close_command.h"

void* C_CloseCommand_Alloc() {
    return [NSCloseCommand alloc];
}

void* C_NSCloseCommand_InitWithCommandDescription(void* ptr, void* commandDef) {
    NSCloseCommand* nSCloseCommand = (NSCloseCommand*)ptr;
    NSCloseCommand* result_ = [nSCloseCommand initWithCommandDescription:(NSScriptCommandDescription*)commandDef];
    return result_;
}

void* C_NSCloseCommand_InitWithCoder(void* ptr, void* inCoder) {
    NSCloseCommand* nSCloseCommand = (NSCloseCommand*)ptr;
    NSCloseCommand* result_ = [nSCloseCommand initWithCoder:(NSCoder*)inCoder];
    return result_;
}

void* C_NSCloseCommand_AllocCloseCommand() {
    NSCloseCommand* result_ = [NSCloseCommand alloc];
    return result_;
}

void* C_NSCloseCommand_Autorelease(void* ptr) {
    NSCloseCommand* nSCloseCommand = (NSCloseCommand*)ptr;
    NSCloseCommand* result_ = [nSCloseCommand autorelease];
    return result_;
}

void* C_NSCloseCommand_Retain(void* ptr) {
    NSCloseCommand* nSCloseCommand = (NSCloseCommand*)ptr;
    NSCloseCommand* result_ = [nSCloseCommand retain];
    return result_;
}

unsigned int C_NSCloseCommand_SaveOptions(void* ptr) {
    NSCloseCommand* nSCloseCommand = (NSCloseCommand*)ptr;
    NSSaveOptions result_ = [nSCloseCommand saveOptions];
    return result_;
}
