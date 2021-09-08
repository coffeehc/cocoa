#import <Appkit/Appkit.h>
#import "font_manager.h"

void* C_FontManager_Alloc() {
    return [NSFontManager alloc];
}

void* C_NSFontManager_AllocFontManager() {
    NSFontManager* result_ = [NSFontManager alloc];
    return result_;
}

void* C_NSFontManager_Init(void* ptr) {
    NSFontManager* nSFontManager = (NSFontManager*)ptr;
    NSFontManager* result_ = [nSFontManager init];
    return result_;
}

void* C_NSFontManager_NewFontManager() {
    NSFontManager* result_ = [NSFontManager new];
    return result_;
}

void* C_NSFontManager_Autorelease(void* ptr) {
    NSFontManager* nSFontManager = (NSFontManager*)ptr;
    NSFontManager* result_ = [nSFontManager autorelease];
    return result_;
}

void* C_NSFontManager_Retain(void* ptr) {
    NSFontManager* nSFontManager = (NSFontManager*)ptr;
    NSFontManager* result_ = [nSFontManager retain];
    return result_;
}

Array C_NSFontManager_AvailableFontNamesWithTraits(void* ptr, unsigned int someTraits) {
    NSFontManager* nSFontManager = (NSFontManager*)ptr;
    NSArray* result_ = [nSFontManager availableFontNamesWithTraits:someTraits];
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

void C_NSFontManager_SetSelectedFont_IsMultiple(void* ptr, void* fontObj, bool flag) {
    NSFontManager* nSFontManager = (NSFontManager*)ptr;
    [nSFontManager setSelectedFont:(NSFont*)fontObj isMultiple:flag];
}

bool C_NSFontManager_SendAction(void* ptr) {
    NSFontManager* nSFontManager = (NSFontManager*)ptr;
    BOOL result_ = [nSFontManager sendAction];
    return result_;
}

void* C_NSFontManager_LocalizedNameForFamily_Face(void* ptr, void* family, void* faceKey) {
    NSFontManager* nSFontManager = (NSFontManager*)ptr;
    NSString* result_ = [nSFontManager localizedNameForFamily:(NSString*)family face:(NSString*)faceKey];
    return result_;
}

void C_NSFontManager_AddFontTrait(void* ptr, void* sender) {
    NSFontManager* nSFontManager = (NSFontManager*)ptr;
    [nSFontManager addFontTrait:(id)sender];
}

void C_NSFontManager_RemoveFontTrait(void* ptr, void* sender) {
    NSFontManager* nSFontManager = (NSFontManager*)ptr;
    [nSFontManager removeFontTrait:(id)sender];
}

void C_NSFontManager_ModifyFont(void* ptr, void* sender) {
    NSFontManager* nSFontManager = (NSFontManager*)ptr;
    [nSFontManager modifyFont:(id)sender];
}

void C_NSFontManager_ModifyFontViaPanel(void* ptr, void* sender) {
    NSFontManager* nSFontManager = (NSFontManager*)ptr;
    [nSFontManager modifyFontViaPanel:(id)sender];
}

void C_NSFontManager_OrderFrontStylesPanel(void* ptr, void* sender) {
    NSFontManager* nSFontManager = (NSFontManager*)ptr;
    [nSFontManager orderFrontStylesPanel:(id)sender];
}

void C_NSFontManager_OrderFrontFontPanel(void* ptr, void* sender) {
    NSFontManager* nSFontManager = (NSFontManager*)ptr;
    [nSFontManager orderFrontFontPanel:(id)sender];
}

void* C_NSFontManager_ConvertFont(void* ptr, void* fontObj) {
    NSFontManager* nSFontManager = (NSFontManager*)ptr;
    NSFont* result_ = [nSFontManager convertFont:(NSFont*)fontObj];
    return result_;
}

void* C_NSFontManager_ConvertFont_ToFace(void* ptr, void* fontObj, void* typeface) {
    NSFontManager* nSFontManager = (NSFontManager*)ptr;
    NSFont* result_ = [nSFontManager convertFont:(NSFont*)fontObj toFace:(NSString*)typeface];
    return result_;
}

void* C_NSFontManager_ConvertFont_ToFamily(void* ptr, void* fontObj, void* family) {
    NSFontManager* nSFontManager = (NSFontManager*)ptr;
    NSFont* result_ = [nSFontManager convertFont:(NSFont*)fontObj toFamily:(NSString*)family];
    return result_;
}

void* C_NSFontManager_ConvertFont_ToHaveTrait(void* ptr, void* fontObj, unsigned int trait) {
    NSFontManager* nSFontManager = (NSFontManager*)ptr;
    NSFont* result_ = [nSFontManager convertFont:(NSFont*)fontObj toHaveTrait:trait];
    return result_;
}

void* C_NSFontManager_ConvertFont_ToNotHaveTrait(void* ptr, void* fontObj, unsigned int trait) {
    NSFontManager* nSFontManager = (NSFontManager*)ptr;
    NSFont* result_ = [nSFontManager convertFont:(NSFont*)fontObj toNotHaveTrait:trait];
    return result_;
}

void* C_NSFontManager_ConvertFont_ToSize(void* ptr, void* fontObj, double size) {
    NSFontManager* nSFontManager = (NSFontManager*)ptr;
    NSFont* result_ = [nSFontManager convertFont:(NSFont*)fontObj toSize:size];
    return result_;
}

void* C_NSFontManager_ConvertWeight_OfFont(void* ptr, bool upFlag, void* fontObj) {
    NSFontManager* nSFontManager = (NSFontManager*)ptr;
    NSFont* result_ = [nSFontManager convertWeight:upFlag ofFont:(NSFont*)fontObj];
    return result_;
}

unsigned int C_NSFontManager_ConvertFontTraits(void* ptr, unsigned int traits) {
    NSFontManager* nSFontManager = (NSFontManager*)ptr;
    NSFontTraitMask result_ = [nSFontManager convertFontTraits:traits];
    return result_;
}

void* C_NSFontManager_FontWithFamily_Traits_Weight_Size(void* ptr, void* family, unsigned int traits, int weight, double size) {
    NSFontManager* nSFontManager = (NSFontManager*)ptr;
    NSFont* result_ = [nSFontManager fontWithFamily:(NSString*)family traits:traits weight:weight size:size];
    return result_;
}

unsigned int C_NSFontManager_TraitsOfFont(void* ptr, void* fontObj) {
    NSFontManager* nSFontManager = (NSFontManager*)ptr;
    NSFontTraitMask result_ = [nSFontManager traitsOfFont:(NSFont*)fontObj];
    return result_;
}

bool C_NSFontManager_FontNamed_HasTraits(void* ptr, void* fName, unsigned int someTraits) {
    NSFontManager* nSFontManager = (NSFontManager*)ptr;
    BOOL result_ = [nSFontManager fontNamed:(NSString*)fName hasTraits:someTraits];
    return result_;
}

int C_NSFontManager_WeightOfFont(void* ptr, void* fontObj) {
    NSFontManager* nSFontManager = (NSFontManager*)ptr;
    NSInteger result_ = [nSFontManager weightOfFont:(NSFont*)fontObj];
    return result_;
}

void* C_NSFontManager_FontPanel(void* ptr, bool create) {
    NSFontManager* nSFontManager = (NSFontManager*)ptr;
    NSFontPanel* result_ = [nSFontManager fontPanel:create];
    return result_;
}

void C_NSFontManager_SetFontMenu(void* ptr, void* newMenu) {
    NSFontManager* nSFontManager = (NSFontManager*)ptr;
    [nSFontManager setFontMenu:(NSMenu*)newMenu];
}

void* C_NSFontManager_FontMenu(void* ptr, bool create) {
    NSFontManager* nSFontManager = (NSFontManager*)ptr;
    NSMenu* result_ = [nSFontManager fontMenu:create];
    return result_;
}

void C_NSFontManager_SetSelectedAttributes_IsMultiple(void* ptr, Dictionary attributes, bool flag) {
    NSFontManager* nSFontManager = (NSFontManager*)ptr;
    NSMutableDictionary* objcAttributes = [NSMutableDictionary dictionaryWithCapacity:attributes.len];
    if (attributes.len > 0) {
    	void** attributesKeyData = (void**)attributes.key_data;
    	void** attributesValueData = (void**)attributes.value_data;
    	for (int i = 0; i < attributes.len; i++) {
    		void* kp = attributesKeyData[i];
    		void* vp = attributesValueData[i];
    		[objcAttributes setObject:(NSString*)(NSString*)kp forKey:(id)(NSString*)vp];
    	}
    }
    [nSFontManager setSelectedAttributes:objcAttributes isMultiple:flag];
}

Dictionary C_NSFontManager_ConvertAttributes(void* ptr, Dictionary attributes) {
    NSFontManager* nSFontManager = (NSFontManager*)ptr;
    NSMutableDictionary* objcAttributes = [NSMutableDictionary dictionaryWithCapacity:attributes.len];
    if (attributes.len > 0) {
    	void** attributesKeyData = (void**)attributes.key_data;
    	void** attributesValueData = (void**)attributes.value_data;
    	for (int i = 0; i < attributes.len; i++) {
    		void* kp = attributesKeyData[i];
    		void* vp = attributesValueData[i];
    		[objcAttributes setObject:(NSString*)(NSString*)kp forKey:(id)(NSString*)vp];
    	}
    }
    NSDictionary* result_ = [nSFontManager convertAttributes:objcAttributes];
    Dictionary result_Array;
    NSArray * result_Keys = [result_ allKeys];
    int result_Count = [result_Keys count];
    if (result_Count > 0) {
    	void** result_KeyData = malloc(result_Count * sizeof(void*));
    	void** result_ValueData = malloc(result_Count * sizeof(void*));
    	for (int i = 0; i < result_Count; i++) {
    		NSString* kp = [result_Keys objectAtIndex:i];
    		id vp = result_[kp];
    		 result_KeyData[i] = kp;
    		 result_ValueData[i] = vp;
    	}
    	result_Array.key_data = result_KeyData;
    	result_Array.value_data = result_ValueData;
    	result_Array.len = result_Count;
    }
    return result_Array;
}

void* C_NSFontManager_SharedFontManager() {
    NSFontManager* result_ = [NSFontManager sharedFontManager];
    return result_;
}

Array C_NSFontManager_AvailableFonts(void* ptr) {
    NSFontManager* nSFontManager = (NSFontManager*)ptr;
    NSArray* result_ = [nSFontManager availableFonts];
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

Array C_NSFontManager_AvailableFontFamilies(void* ptr) {
    NSFontManager* nSFontManager = (NSFontManager*)ptr;
    NSArray* result_ = [nSFontManager availableFontFamilies];
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

void* C_NSFontManager_SelectedFont(void* ptr) {
    NSFontManager* nSFontManager = (NSFontManager*)ptr;
    NSFont* result_ = [nSFontManager selectedFont];
    return result_;
}

bool C_NSFontManager_IsMultiple(void* ptr) {
    NSFontManager* nSFontManager = (NSFontManager*)ptr;
    BOOL result_ = [nSFontManager isMultiple];
    return result_;
}

unsigned int C_NSFontManager_CurrentFontAction(void* ptr) {
    NSFontManager* nSFontManager = (NSFontManager*)ptr;
    NSFontAction result_ = [nSFontManager currentFontAction];
    return result_;
}

bool C_NSFontManager_IsEnabled(void* ptr) {
    NSFontManager* nSFontManager = (NSFontManager*)ptr;
    BOOL result_ = [nSFontManager isEnabled];
    return result_;
}

void C_NSFontManager_SetEnabled(void* ptr, bool value) {
    NSFontManager* nSFontManager = (NSFontManager*)ptr;
    [nSFontManager setEnabled:value];
}

void* C_NSFontManager_Action(void* ptr) {
    NSFontManager* nSFontManager = (NSFontManager*)ptr;
    SEL result_ = [nSFontManager action];
    return result_;
}

void C_NSFontManager_SetAction(void* ptr, void* value) {
    NSFontManager* nSFontManager = (NSFontManager*)ptr;
    [nSFontManager setAction:(SEL)value];
}

void* C_NSFontManager_Target(void* ptr) {
    NSFontManager* nSFontManager = (NSFontManager*)ptr;
    id result_ = [nSFontManager target];
    return result_;
}

void C_NSFontManager_SetTarget(void* ptr, void* value) {
    NSFontManager* nSFontManager = (NSFontManager*)ptr;
    [nSFontManager setTarget:(id)value];
}
