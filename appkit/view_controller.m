#import <Appkit/Appkit.h>
#import "view_controller.h"

void* C_ViewController_Alloc() {
    return [NSViewController alloc];
}

void* C_NSViewController_InitWithNibName_Bundle(void* ptr, void* nibNameOrNil, void* nibBundleOrNil) {
    NSViewController* nSViewController = (NSViewController*)ptr;
    NSViewController* result_ = [nSViewController initWithNibName:(NSString*)nibNameOrNil bundle:(NSBundle*)nibBundleOrNil];
    return result_;
}

void* C_NSViewController_InitWithCoder(void* ptr, void* coder) {
    NSViewController* nSViewController = (NSViewController*)ptr;
    NSViewController* result_ = [nSViewController initWithCoder:(NSCoder*)coder];
    return result_;
}

void* C_NSViewController_Init(void* ptr) {
    NSViewController* nSViewController = (NSViewController*)ptr;
    NSViewController* result_ = [nSViewController init];
    return result_;
}

void C_NSViewController_LoadView(void* ptr) {
    NSViewController* nSViewController = (NSViewController*)ptr;
    [nSViewController loadView];
}

bool C_NSViewController_CommitEditing(void* ptr) {
    NSViewController* nSViewController = (NSViewController*)ptr;
    BOOL result_ = [nSViewController commitEditing];
    return result_;
}

void C_NSViewController_DiscardEditing(void* ptr) {
    NSViewController* nSViewController = (NSViewController*)ptr;
    [nSViewController discardEditing];
}

void C_NSViewController_DismissController(void* ptr, void* sender) {
    NSViewController* nSViewController = (NSViewController*)ptr;
    [nSViewController dismissController:(id)sender];
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

void C_NSViewController_PresentViewController_Animator(void* ptr, void* viewController, void* animator) {
    NSViewController* nSViewController = (NSViewController*)ptr;
    [nSViewController presentViewController:(NSViewController*)viewController animator:(id)animator];
}

void C_NSViewController_DismissViewController(void* ptr, void* viewController) {
    NSViewController* nSViewController = (NSViewController*)ptr;
    [nSViewController dismissViewController:(NSViewController*)viewController];
}

void C_NSViewController_PresentViewController_AsPopoverRelativeToRect_OfView_PreferredEdge_Behavior(void* ptr, void* viewController, CGRect positioningRect, void* positioningView, unsigned int preferredEdge, int behavior) {
    NSViewController* nSViewController = (NSViewController*)ptr;
    [nSViewController presentViewController:(NSViewController*)viewController asPopoverRelativeToRect:positioningRect ofView:(NSView*)positioningView preferredEdge:preferredEdge behavior:behavior];
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
    id result_ = [nSViewController representedObject];
    return result_;
}

void C_NSViewController_SetRepresentedObject(void* ptr, void* value) {
    NSViewController* nSViewController = (NSViewController*)ptr;
    [nSViewController setRepresentedObject:(id)value];
}

void* C_NSViewController_NibBundle(void* ptr) {
    NSViewController* nSViewController = (NSViewController*)ptr;
    NSBundle* result_ = [nSViewController nibBundle];
    return result_;
}

void* C_NSViewController_NibName(void* ptr) {
    NSViewController* nSViewController = (NSViewController*)ptr;
    NSNibName result_ = [nSViewController nibName];
    return result_;
}

void* C_NSViewController_Title(void* ptr) {
    NSViewController* nSViewController = (NSViewController*)ptr;
    NSString* result_ = [nSViewController title];
    return result_;
}

void C_NSViewController_SetTitle(void* ptr, void* value) {
    NSViewController* nSViewController = (NSViewController*)ptr;
    [nSViewController setTitle:(NSString*)value];
}

void* C_NSViewController_Storyboard(void* ptr) {
    NSViewController* nSViewController = (NSViewController*)ptr;
    NSStoryboard* result_ = [nSViewController storyboard];
    return result_;
}

bool C_NSViewController_IsViewLoaded(void* ptr) {
    NSViewController* nSViewController = (NSViewController*)ptr;
    BOOL result_ = [nSViewController isViewLoaded];
    return result_;
}

CGSize C_NSViewController_PreferredContentSize(void* ptr) {
    NSViewController* nSViewController = (NSViewController*)ptr;
    NSSize result_ = [nSViewController preferredContentSize];
    return result_;
}

void C_NSViewController_SetPreferredContentSize(void* ptr, CGSize value) {
    NSViewController* nSViewController = (NSViewController*)ptr;
    [nSViewController setPreferredContentSize:value];
}

Array C_NSViewController_ChildViewControllers(void* ptr) {
    NSViewController* nSViewController = (NSViewController*)ptr;
    NSArray* result_ = [nSViewController childViewControllers];
    Array result_Array;
    int result_count = [result_ count];
    if (result_count > 0) {
    	void** result_Data = malloc(result_count * sizeof(void*));
    	for (int i = 0; i < result_count; i++) {
    		 void* p = [result_ objectAtIndex:i];
    		 result_Data[i] = p;
    	}
    	result_Array.data = result_Data;
    	result_Array.len = result_count;
    }
    return result_Array;
}

void C_NSViewController_SetChildViewControllers(void* ptr, Array value) {
    NSViewController* nSViewController = (NSViewController*)ptr;
    NSMutableArray* objcValue = [[NSMutableArray alloc] init];
    if (value.len > 0) {
    	void** valueData = (void**)value.data;
    	for (int i = 0; i < value.len; i++) {
    		void* p = valueData[i];
    		[objcValue addObject:(NSViewController*)(NSViewController*)p];
    	}
    }
    [nSViewController setChildViewControllers:objcValue];
}

void* C_NSViewController_ParentViewController(void* ptr) {
    NSViewController* nSViewController = (NSViewController*)ptr;
    NSViewController* result_ = [nSViewController parentViewController];
    return result_;
}

Array C_NSViewController_PresentedViewControllers(void* ptr) {
    NSViewController* nSViewController = (NSViewController*)ptr;
    NSArray* result_ = [nSViewController presentedViewControllers];
    Array result_Array;
    int result_count = [result_ count];
    if (result_count > 0) {
    	void** result_Data = malloc(result_count * sizeof(void*));
    	for (int i = 0; i < result_count; i++) {
    		 void* p = [result_ objectAtIndex:i];
    		 result_Data[i] = p;
    	}
    	result_Array.data = result_Data;
    	result_Array.len = result_count;
    }
    return result_Array;
}

void* C_NSViewController_PresentingViewController(void* ptr) {
    NSViewController* nSViewController = (NSViewController*)ptr;
    NSViewController* result_ = [nSViewController presentingViewController];
    return result_;
}

void* C_NSViewController_ExtensionContext(void* ptr) {
    NSViewController* nSViewController = (NSViewController*)ptr;
    NSExtensionContext* result_ = [nSViewController extensionContext];
    return result_;
}

CGPoint C_NSViewController_PreferredScreenOrigin(void* ptr) {
    NSViewController* nSViewController = (NSViewController*)ptr;
    NSPoint result_ = [nSViewController preferredScreenOrigin];
    return result_;
}

void C_NSViewController_SetPreferredScreenOrigin(void* ptr, CGPoint value) {
    NSViewController* nSViewController = (NSViewController*)ptr;
    [nSViewController setPreferredScreenOrigin:value];
}

CGSize C_NSViewController_PreferredMaximumSize(void* ptr) {
    NSViewController* nSViewController = (NSViewController*)ptr;
    NSSize result_ = [nSViewController preferredMaximumSize];
    return result_;
}

CGSize C_NSViewController_PreferredMinimumSize(void* ptr) {
    NSViewController* nSViewController = (NSViewController*)ptr;
    NSSize result_ = [nSViewController preferredMinimumSize];
    return result_;
}
