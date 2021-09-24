#import <Foundation/NSGeometry.h>
#import <Foundation/NSRange.h>
#import <stdbool.h>
#import <stdint.h>
#import <stdlib.h>
#import <utils.h>

void* C_SharingServicePicker_Alloc();

void* C_NSSharingServicePicker_InitWithItems(void* ptr, Array items);
void* C_NSSharingServicePicker_AllocSharingServicePicker();
void* C_NSSharingServicePicker_Autorelease(void* ptr);
void* C_NSSharingServicePicker_Retain(void* ptr);
void C_NSSharingServicePicker_ShowRelativeToRect_OfView_PreferredEdge(void* ptr, CGRect rect, void* view, unsigned int preferredEdge);
void* C_NSSharingServicePicker_Delegate(void* ptr);
void C_NSSharingServicePicker_SetDelegate(void* ptr, void* value);
