#import <Appkit/Appkit.h>
#import "text_input_context.h"

void* C_TextInputContext_Alloc() {
    return [NSTextInputContext alloc];
}

void* C_NSTextInputContext_InitWithClient(void* ptr, void* client) {
    NSTextInputContext* nSTextInputContext = (NSTextInputContext*)ptr;
    NSTextInputContext* result_ = [nSTextInputContext initWithClient:(id)client];
    return result_;
}

void* C_NSTextInputContext_AllocTextInputContext() {
    NSTextInputContext* result_ = [NSTextInputContext alloc];
    return result_;
}

void* C_NSTextInputContext_Autorelease(void* ptr) {
    NSTextInputContext* nSTextInputContext = (NSTextInputContext*)ptr;
    NSTextInputContext* result_ = [nSTextInputContext autorelease];
    return result_;
}

void* C_NSTextInputContext_Retain(void* ptr) {
    NSTextInputContext* nSTextInputContext = (NSTextInputContext*)ptr;
    NSTextInputContext* result_ = [nSTextInputContext retain];
    return result_;
}

void C_NSTextInputContext_Activate(void* ptr) {
    NSTextInputContext* nSTextInputContext = (NSTextInputContext*)ptr;
    [nSTextInputContext activate];
}

void C_NSTextInputContext_Deactivate(void* ptr) {
    NSTextInputContext* nSTextInputContext = (NSTextInputContext*)ptr;
    [nSTextInputContext deactivate];
}

bool C_NSTextInputContext_HandleEvent(void* ptr, void* event) {
    NSTextInputContext* nSTextInputContext = (NSTextInputContext*)ptr;
    BOOL result_ = [nSTextInputContext handleEvent:(NSEvent*)event];
    return result_;
}

void C_NSTextInputContext_DiscardMarkedText(void* ptr) {
    NSTextInputContext* nSTextInputContext = (NSTextInputContext*)ptr;
    [nSTextInputContext discardMarkedText];
}

void C_NSTextInputContext_InvalidateCharacterCoordinates(void* ptr) {
    NSTextInputContext* nSTextInputContext = (NSTextInputContext*)ptr;
    [nSTextInputContext invalidateCharacterCoordinates];
}

void* C_NSTextInputContext_TextInputContext_LocalizedNameForInputSource(void* inputSourceIdentifier) {
    NSString* result_ = [NSTextInputContext localizedNameForInputSource:(NSString*)inputSourceIdentifier];
    return result_;
}

void* C_NSTextInputContext_TextInputContext_CurrentInputContext() {
    NSTextInputContext* result_ = [NSTextInputContext currentInputContext];
    return result_;
}

void* C_NSTextInputContext_Client(void* ptr) {
    NSTextInputContext* nSTextInputContext = (NSTextInputContext*)ptr;
    id result_ = [nSTextInputContext client];
    return result_;
}

bool C_NSTextInputContext_AcceptsGlyphInfo(void* ptr) {
    NSTextInputContext* nSTextInputContext = (NSTextInputContext*)ptr;
    BOOL result_ = [nSTextInputContext acceptsGlyphInfo];
    return result_;
}

void C_NSTextInputContext_SetAcceptsGlyphInfo(void* ptr, bool value) {
    NSTextInputContext* nSTextInputContext = (NSTextInputContext*)ptr;
    [nSTextInputContext setAcceptsGlyphInfo:value];
}

Array C_NSTextInputContext_AllowedInputSourceLocales(void* ptr) {
    NSTextInputContext* nSTextInputContext = (NSTextInputContext*)ptr;
    NSArray* result_ = [nSTextInputContext allowedInputSourceLocales];
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

void C_NSTextInputContext_SetAllowedInputSourceLocales(void* ptr, Array value) {
    NSTextInputContext* nSTextInputContext = (NSTextInputContext*)ptr;
    NSMutableArray* objcValue = [NSMutableArray arrayWithCapacity:value.len];
    if (value.len > 0) {
    	void** valueData = (void**)value.data;
    	for (int i = 0; i < value.len; i++) {
    		void* p = valueData[i];
    		[objcValue addObject:(NSString*)p];
    	}
    }
    [nSTextInputContext setAllowedInputSourceLocales:objcValue];
}

Array C_NSTextInputContext_KeyboardInputSources(void* ptr) {
    NSTextInputContext* nSTextInputContext = (NSTextInputContext*)ptr;
    NSArray* result_ = [nSTextInputContext keyboardInputSources];
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

void* C_NSTextInputContext_SelectedKeyboardInputSource(void* ptr) {
    NSTextInputContext* nSTextInputContext = (NSTextInputContext*)ptr;
    NSTextInputSourceIdentifier result_ = [nSTextInputContext selectedKeyboardInputSource];
    return result_;
}

void C_NSTextInputContext_SetSelectedKeyboardInputSource(void* ptr, void* value) {
    NSTextInputContext* nSTextInputContext = (NSTextInputContext*)ptr;
    [nSTextInputContext setSelectedKeyboardInputSource:(NSString*)value];
}
