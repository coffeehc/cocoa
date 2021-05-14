#import <Appkit/Appkit.h>
#import "titlebar_accessory_view_controller.h"

void* C_TitlebarAccessoryViewController_Alloc() {
	return [NSTitlebarAccessoryViewController alloc];
}

void* C_NSTitlebarAccessoryViewController_InitWithNibName_Bundle(void* ptr, void* nibNameOrNil, void* nibBundleOrNil) {
	NSTitlebarAccessoryViewController* nSTitlebarAccessoryViewController = (NSTitlebarAccessoryViewController*)ptr;
	NSTitlebarAccessoryViewController* result = [nSTitlebarAccessoryViewController initWithNibName:(NSString*)nibNameOrNil bundle:(NSBundle*)nibBundleOrNil];
	return result;
}

void* C_NSTitlebarAccessoryViewController_InitWithCoder(void* ptr, void* coder) {
	NSTitlebarAccessoryViewController* nSTitlebarAccessoryViewController = (NSTitlebarAccessoryViewController*)ptr;
	NSTitlebarAccessoryViewController* result = [nSTitlebarAccessoryViewController initWithCoder:(NSCoder*)coder];
	return result;
}

void* C_NSTitlebarAccessoryViewController_Init(void* ptr) {
	NSTitlebarAccessoryViewController* nSTitlebarAccessoryViewController = (NSTitlebarAccessoryViewController*)ptr;
	NSTitlebarAccessoryViewController* result = [nSTitlebarAccessoryViewController init];
	return result;
}

double C_NSTitlebarAccessoryViewController_FullScreenMinHeight(void* ptr) {
	NSTitlebarAccessoryViewController* nSTitlebarAccessoryViewController = (NSTitlebarAccessoryViewController*)ptr;
	double result = [nSTitlebarAccessoryViewController fullScreenMinHeight];
	return result;
}

void C_NSTitlebarAccessoryViewController_SetFullScreenMinHeight(void* ptr, double value) {
	NSTitlebarAccessoryViewController* nSTitlebarAccessoryViewController = (NSTitlebarAccessoryViewController*)ptr;
	[nSTitlebarAccessoryViewController setFullScreenMinHeight:value];
}

int C_NSTitlebarAccessoryViewController_LayoutAttribute(void* ptr) {
	NSTitlebarAccessoryViewController* nSTitlebarAccessoryViewController = (NSTitlebarAccessoryViewController*)ptr;
	int result = [nSTitlebarAccessoryViewController layoutAttribute];
	return result;
}

void C_NSTitlebarAccessoryViewController_SetLayoutAttribute(void* ptr, int value) {
	NSTitlebarAccessoryViewController* nSTitlebarAccessoryViewController = (NSTitlebarAccessoryViewController*)ptr;
	[nSTitlebarAccessoryViewController setLayoutAttribute:value];
}

bool C_NSTitlebarAccessoryViewController_AutomaticallyAdjustsSize(void* ptr) {
	NSTitlebarAccessoryViewController* nSTitlebarAccessoryViewController = (NSTitlebarAccessoryViewController*)ptr;
	bool result = [nSTitlebarAccessoryViewController automaticallyAdjustsSize];
	return result;
}

void C_NSTitlebarAccessoryViewController_SetAutomaticallyAdjustsSize(void* ptr, bool value) {
	NSTitlebarAccessoryViewController* nSTitlebarAccessoryViewController = (NSTitlebarAccessoryViewController*)ptr;
	[nSTitlebarAccessoryViewController setAutomaticallyAdjustsSize:value];
}

bool C_NSTitlebarAccessoryViewController_IsHidden(void* ptr) {
	NSTitlebarAccessoryViewController* nSTitlebarAccessoryViewController = (NSTitlebarAccessoryViewController*)ptr;
	bool result = [nSTitlebarAccessoryViewController isHidden];
	return result;
}

void C_NSTitlebarAccessoryViewController_SetHidden(void* ptr, bool value) {
	NSTitlebarAccessoryViewController* nSTitlebarAccessoryViewController = (NSTitlebarAccessoryViewController*)ptr;
	[nSTitlebarAccessoryViewController setHidden:value];
}
