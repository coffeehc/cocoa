#import <Appkit/Appkit.h>
#import "view_controller.h"

void* C_ViewController_Alloc() {
	return [NSViewController alloc];
}

void* C_NSViewController_InitWithNibName_Bundle(void* ptr, void* nibNameOrNil, void* nibBundleOrNil) {
	NSViewController* nSViewController = (NSViewController*)ptr;
	NSViewController* result = [nSViewController initWithNibName:(NSString*)nibNameOrNil bundle:(NSBundle*)nibBundleOrNil];
	return result;
}

void* C_NSViewController_InitWithCoder(void* ptr, void* coder) {
	NSViewController* nSViewController = (NSViewController*)ptr;
	NSViewController* result = [nSViewController initWithCoder:(NSCoder*)coder];
	return result;
}

void* C_NSViewController_Init(void* ptr) {
	NSViewController* nSViewController = (NSViewController*)ptr;
	NSViewController* result = [nSViewController init];
	return result;
}

void C_NSViewController_LoadView(void* ptr) {
	NSViewController* nSViewController = (NSViewController*)ptr;
	[nSViewController loadView];
}

bool C_NSViewController_CommitEditing(void* ptr) {
	NSViewController* nSViewController = (NSViewController*)ptr;
	bool result = [nSViewController commitEditing];
	return result;
}

void C_NSViewController_DiscardEditing(void* ptr) {
	NSViewController* nSViewController = (NSViewController*)ptr;
	[nSViewController discardEditing];
}

void C_NSViewController_ViewDidLoad(void* ptr) {
	NSViewController* nSViewController = (NSViewController*)ptr;
	[nSViewController viewDidLoad];
}

void C_NSViewController_ViewWillAppear(void* ptr) {
	NSViewController* nSViewController = (NSViewController*)ptr;
	[nSViewController viewWillAppear];
}

void C_NSViewController_ViewDidAppear(void* ptr) {
	NSViewController* nSViewController = (NSViewController*)ptr;
	[nSViewController viewDidAppear];
}

void C_NSViewController_ViewWillDisappear(void* ptr) {
	NSViewController* nSViewController = (NSViewController*)ptr;
	[nSViewController viewWillDisappear];
}

void C_NSViewController_ViewDidDisappear(void* ptr) {
	NSViewController* nSViewController = (NSViewController*)ptr;
	[nSViewController viewDidDisappear];
}

void C_NSViewController_UpdateViewConstraints(void* ptr) {
	NSViewController* nSViewController = (NSViewController*)ptr;
	[nSViewController updateViewConstraints];
}

void C_NSViewController_ViewWillLayout(void* ptr) {
	NSViewController* nSViewController = (NSViewController*)ptr;
	[nSViewController viewWillLayout];
}

void C_NSViewController_ViewDidLayout(void* ptr) {
	NSViewController* nSViewController = (NSViewController*)ptr;
	[nSViewController viewDidLayout];
}

void C_NSViewController_AddChildViewController(void* ptr, void* childViewController) {
	NSViewController* nSViewController = (NSViewController*)ptr;
	[nSViewController addChildViewController:(NSViewController*)childViewController];
}

void C_NSViewController_InsertChildViewController_AtIndex(void* ptr, void* childViewController, int index) {
	NSViewController* nSViewController = (NSViewController*)ptr;
	[nSViewController insertChildViewController:(NSViewController*)childViewController atIndex:index];
}

void C_NSViewController_RemoveChildViewControllerAtIndex(void* ptr, int index) {
	NSViewController* nSViewController = (NSViewController*)ptr;
	[nSViewController removeChildViewControllerAtIndex:index];
}

void C_NSViewController_RemoveFromParentViewController(void* ptr) {
	NSViewController* nSViewController = (NSViewController*)ptr;
	[nSViewController removeFromParentViewController];
}

void C_NSViewController_PreferredContentSizeDidChangeForViewController(void* ptr, void* viewController) {
	NSViewController* nSViewController = (NSViewController*)ptr;
	[nSViewController preferredContentSizeDidChangeForViewController:(NSViewController*)viewController];
}

void C_NSViewController_DismissViewController(void* ptr, void* viewController) {
	NSViewController* nSViewController = (NSViewController*)ptr;
	[nSViewController dismissViewController:(NSViewController*)viewController];
}

void C_NSViewController_PresentViewControllerAsModalWindow(void* ptr, void* viewController) {
	NSViewController* nSViewController = (NSViewController*)ptr;
	[nSViewController presentViewControllerAsModalWindow:(NSViewController*)viewController];
}

void C_NSViewController_PresentViewControllerAsSheet(void* ptr, void* viewController) {
	NSViewController* nSViewController = (NSViewController*)ptr;
	[nSViewController presentViewControllerAsSheet:(NSViewController*)viewController];
}

void C_NSViewController_ViewWillTransitionToSize(void* ptr, CGSize newSize) {
	NSViewController* nSViewController = (NSViewController*)ptr;
	[nSViewController viewWillTransitionToSize:newSize];
}

void* C_NSViewController_RepresentedObject(void* ptr) {
	NSViewController* nSViewController = (NSViewController*)ptr;
	id result = [nSViewController representedObject];
	return result;
}

void C_NSViewController_SetRepresentedObject(void* ptr, void* value) {
	NSViewController* nSViewController = (NSViewController*)ptr;
	[nSViewController setRepresentedObject:(id)value];
}

void* C_NSViewController_NibBundle(void* ptr) {
	NSViewController* nSViewController = (NSViewController*)ptr;
	NSBundle* result = [nSViewController nibBundle];
	return result;
}

void* C_NSViewController_NibName(void* ptr) {
	NSViewController* nSViewController = (NSViewController*)ptr;
	NSString* result = [nSViewController nibName];
	return result;
}

void* C_NSViewController_Title(void* ptr) {
	NSViewController* nSViewController = (NSViewController*)ptr;
	NSString* result = [nSViewController title];
	return result;
}

void C_NSViewController_SetTitle(void* ptr, void* value) {
	NSViewController* nSViewController = (NSViewController*)ptr;
	[nSViewController setTitle:(NSString*)value];
}

void* C_NSViewController_Storyboard(void* ptr) {
	NSViewController* nSViewController = (NSViewController*)ptr;
	NSStoryboard* result = [nSViewController storyboard];
	return result;
}

bool C_NSViewController_IsViewLoaded(void* ptr) {
	NSViewController* nSViewController = (NSViewController*)ptr;
	bool result = [nSViewController isViewLoaded];
	return result;
}

CGSize C_NSViewController_PreferredContentSize(void* ptr) {
	NSViewController* nSViewController = (NSViewController*)ptr;
	CGSize result = [nSViewController preferredContentSize];
	return result;
}

void C_NSViewController_SetPreferredContentSize(void* ptr, CGSize value) {
	NSViewController* nSViewController = (NSViewController*)ptr;
	[nSViewController setPreferredContentSize:value];
}

void* C_NSViewController_ParentViewController(void* ptr) {
	NSViewController* nSViewController = (NSViewController*)ptr;
	NSViewController* result = [nSViewController parentViewController];
	return result;
}

void* C_NSViewController_PresentingViewController(void* ptr) {
	NSViewController* nSViewController = (NSViewController*)ptr;
	NSViewController* result = [nSViewController presentingViewController];
	return result;
}

void* C_NSViewController_ExtensionContext(void* ptr) {
	NSViewController* nSViewController = (NSViewController*)ptr;
	NSExtensionContext* result = [nSViewController extensionContext];
	return result;
}

CGPoint C_NSViewController_PreferredScreenOrigin(void* ptr) {
	NSViewController* nSViewController = (NSViewController*)ptr;
	CGPoint result = [nSViewController preferredScreenOrigin];
	return result;
}

void C_NSViewController_SetPreferredScreenOrigin(void* ptr, CGPoint value) {
	NSViewController* nSViewController = (NSViewController*)ptr;
	[nSViewController setPreferredScreenOrigin:value];
}

CGSize C_NSViewController_PreferredMaximumSize(void* ptr) {
	NSViewController* nSViewController = (NSViewController*)ptr;
	CGSize result = [nSViewController preferredMaximumSize];
	return result;
}

CGSize C_NSViewController_PreferredMinimumSize(void* ptr) {
	NSViewController* nSViewController = (NSViewController*)ptr;
	CGSize result = [nSViewController preferredMinimumSize];
	return result;
}
