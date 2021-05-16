#import <stdbool.h>
#import <stdlib.h>
#import <utils.h>
#import <Appkit/Appkit.h>

void* C_SharingServicePicker_Alloc();

void* C_NSSharingServicePicker_InitWithItems(void* ptr, Array items);
void C_NSSharingServicePicker_ShowRelativeToRect_OfView_PreferredEdge(void* ptr, CGRect rect, void* view, unsigned int preferredEdge);
void* C_NSSharingServicePicker_Delegate(void* ptr);
void C_NSSharingServicePicker_SetDelegate(void* ptr, void* value);
