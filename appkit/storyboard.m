#import <Appkit/Appkit.h>
#import "storyboard.h"

void* C_Storyboard_Alloc() {
	return [NSStoryboard alloc];
}

void* C_NSStoryboard_Init(void* ptr) {
	NSStoryboard* nSStoryboard = (NSStoryboard*)ptr;
	NSStoryboard* result = [nSStoryboard init];
	return result;
}

void* C_NSStoryboard_StoryboardWithName_Bundle(void* name, void* storyboardBundleOrNil) {
	NSStoryboard* result = [NSStoryboard storyboardWithName:(NSString*)name bundle:(NSBundle*)storyboardBundleOrNil];
	return result;
}

void* C_NSStoryboard_InstantiateInitialController(void* ptr) {
	NSStoryboard* nSStoryboard = (NSStoryboard*)ptr;
	id result = [nSStoryboard instantiateInitialController];
	return result;
}
