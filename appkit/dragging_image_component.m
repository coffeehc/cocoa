#import <Appkit/Appkit.h>
#import "dragging_image_component.h"

void* C_DraggingImageComponent_Alloc() {
    return [NSDraggingImageComponent alloc];
}

void* C_NSDraggingImageComponent_InitWithKey(void* ptr, void* key) {
    NSDraggingImageComponent* nSDraggingImageComponent = (NSDraggingImageComponent*)ptr;
    NSDraggingImageComponent* result_ = [nSDraggingImageComponent initWithKey:(NSString*)key];
    return result_;
}

void* C_NSDraggingImageComponent_DraggingImageComponentWithKey(void* key) {
    NSDraggingImageComponent* result_ = [NSDraggingImageComponent draggingImageComponentWithKey:(NSString*)key];
    return result_;
}

void* C_NSDraggingImageComponent_Key(void* ptr) {
    NSDraggingImageComponent* nSDraggingImageComponent = (NSDraggingImageComponent*)ptr;
    NSDraggingImageComponentKey result_ = [nSDraggingImageComponent key];
    return result_;
}

void C_NSDraggingImageComponent_SetKey(void* ptr, void* value) {
    NSDraggingImageComponent* nSDraggingImageComponent = (NSDraggingImageComponent*)ptr;
    [nSDraggingImageComponent setKey:(NSString*)value];
}

void* C_NSDraggingImageComponent_Contents(void* ptr) {
    NSDraggingImageComponent* nSDraggingImageComponent = (NSDraggingImageComponent*)ptr;
    id result_ = [nSDraggingImageComponent contents];
    return result_;
}

void C_NSDraggingImageComponent_SetContents(void* ptr, void* value) {
    NSDraggingImageComponent* nSDraggingImageComponent = (NSDraggingImageComponent*)ptr;
    [nSDraggingImageComponent setContents:(id)value];
}

CGRect C_NSDraggingImageComponent_Frame(void* ptr) {
    NSDraggingImageComponent* nSDraggingImageComponent = (NSDraggingImageComponent*)ptr;
    NSRect result_ = [nSDraggingImageComponent frame];
    return result_;
}

void C_NSDraggingImageComponent_SetFrame(void* ptr, CGRect value) {
    NSDraggingImageComponent* nSDraggingImageComponent = (NSDraggingImageComponent*)ptr;
    [nSDraggingImageComponent setFrame:value];
}
