#import <stdbool.h>
#import <stdlib.h>
#import <utils.h>
#import <Appkit/Appkit.h>

void* C_ViewController_Alloc();

void* C_NSViewController_InitWithNibName_Bundle(void* ptr, void* nibNameOrNil, void* nibBundleOrNil);
void* C_NSViewController_InitWithCoder(void* ptr, void* coder);
void* C_NSViewController_Init(void* ptr);
void C_NSViewController_LoadView(void* ptr);
bool C_NSViewController_CommitEditing(void* ptr);
void C_NSViewController_DiscardEditing(void* ptr);
void C_NSViewController_ViewDidLoad(void* ptr);
void C_NSViewController_ViewWillAppear(void* ptr);
void C_NSViewController_ViewDidAppear(void* ptr);
void C_NSViewController_ViewWillDisappear(void* ptr);
void C_NSViewController_ViewDidDisappear(void* ptr);
void C_NSViewController_UpdateViewConstraints(void* ptr);
void C_NSViewController_ViewWillLayout(void* ptr);
void C_NSViewController_ViewDidLayout(void* ptr);
void C_NSViewController_AddChildViewController(void* ptr, void* childViewController);
void C_NSViewController_InsertChildViewController_AtIndex(void* ptr, void* childViewController, int index);
void C_NSViewController_RemoveChildViewControllerAtIndex(void* ptr, int index);
void C_NSViewController_RemoveFromParentViewController(void* ptr);
void C_NSViewController_PreferredContentSizeDidChangeForViewController(void* ptr, void* viewController);
void C_NSViewController_DismissViewController(void* ptr, void* viewController);
void C_NSViewController_PresentViewControllerAsModalWindow(void* ptr, void* viewController);
void C_NSViewController_PresentViewControllerAsSheet(void* ptr, void* viewController);
void C_NSViewController_ViewWillTransitionToSize(void* ptr, CGSize newSize);
void* C_NSViewController_RepresentedObject(void* ptr);
void C_NSViewController_SetRepresentedObject(void* ptr, void* value);
void* C_NSViewController_NibBundle(void* ptr);
void* C_NSViewController_NibName(void* ptr);
void* C_NSViewController_Title(void* ptr);
void C_NSViewController_SetTitle(void* ptr, void* value);
void* C_NSViewController_Storyboard(void* ptr);
bool C_NSViewController_IsViewLoaded(void* ptr);
CGSize C_NSViewController_PreferredContentSize(void* ptr);
void C_NSViewController_SetPreferredContentSize(void* ptr, CGSize value);
void* C_NSViewController_ParentViewController(void* ptr);
void* C_NSViewController_PresentingViewController(void* ptr);
void* C_NSViewController_ExtensionContext(void* ptr);
CGPoint C_NSViewController_PreferredScreenOrigin(void* ptr);
void C_NSViewController_SetPreferredScreenOrigin(void* ptr, CGPoint value);
CGSize C_NSViewController_PreferredMaximumSize(void* ptr);
CGSize C_NSViewController_PreferredMinimumSize(void* ptr);
