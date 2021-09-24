#import <Foundation/NSGeometry.h>
#import <Foundation/NSRange.h>
#import <stdbool.h>
#import <stdint.h>
#import <stdlib.h>
#import <utils.h>

void* C_Storyboard_Alloc();

void* C_NSStoryboard_StoryboardWithName_Bundle(void* name, void* storyboardBundleOrNil);
void* C_NSStoryboard_AllocStoryboard();
void* C_NSStoryboard_Init(void* ptr);
void* C_NSStoryboard_NewStoryboard();
void* C_NSStoryboard_Autorelease(void* ptr);
void* C_NSStoryboard_Retain(void* ptr);
void* C_NSStoryboard_InstantiateInitialController(void* ptr);
void* C_NSStoryboard_MainStoryboard();
