#import <Appkit/Appkit.h>
#import "storyboard.h"

void* C_Storyboard_Alloc() {
    return [NSStoryboard alloc];
}

void* C_NSStoryboard_Init(void* ptr) {
    NSStoryboard* nSStoryboard = (NSStoryboard*)ptr;
    NSStoryboard* result_ = [nSStoryboard init];
    return result_;
}

void* C_NSStoryboard_StoryboardWithName_Bundle(void* name, void* storyboardBundleOrNil) {
    NSStoryboard* result_ = [NSStoryboard storyboardWithName:(NSString*)name bundle:(NSBundle*)storyboardBundleOrNil];
    return result_;
}

void* C_NSStoryboard_InstantiateInitialController(void* ptr) {
    NSStoryboard* nSStoryboard = (NSStoryboard*)ptr;
    id result_ = [nSStoryboard instantiateInitialController];
    return result_;
}

void* C_NSStoryboard_MainStoryboard() {
    NSStoryboard* result_ = [NSStoryboard mainStoryboard];
    return result_;
}
