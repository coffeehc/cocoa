#import <Appkit/Appkit.h>
#import "sharing_service_picker.h"

void* C_SharingServicePicker_Alloc() {
    return [NSSharingServicePicker alloc];
}

void* C_NSSharingServicePicker_InitWithItems(void* ptr, Array items) {
    NSSharingServicePicker* nSSharingServicePicker = (NSSharingServicePicker*)ptr;
    NSMutableArray* objcItems = [[NSMutableArray alloc] init];
    void** itemsData = (void**)items.data;
    for (int i = 0; i < items.len; i++) {
    	void* p = itemsData[i];
    	[objcItems addObject:(NSObject*)(NSObject*)p];
    }
    NSSharingServicePicker* result_ = [nSSharingServicePicker initWithItems:objcItems];
    return result_;
}

void C_NSSharingServicePicker_ShowRelativeToRect_OfView_PreferredEdge(void* ptr, CGRect rect, void* view, unsigned int preferredEdge) {
    NSSharingServicePicker* nSSharingServicePicker = (NSSharingServicePicker*)ptr;
    [nSSharingServicePicker showRelativeToRect:rect ofView:(NSView*)view preferredEdge:preferredEdge];
}

void* C_NSSharingServicePicker_Delegate(void* ptr) {
    NSSharingServicePicker* nSSharingServicePicker = (NSSharingServicePicker*)ptr;
    id result_ = [nSSharingServicePicker delegate];
    return result_;
}

void C_NSSharingServicePicker_SetDelegate(void* ptr, void* value) {
    NSSharingServicePicker* nSSharingServicePicker = (NSSharingServicePicker*)ptr;
    [nSSharingServicePicker setDelegate:(id)value];
}
